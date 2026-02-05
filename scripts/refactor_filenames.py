import os
import re

def rename_reports(base_dir):
    pattern = re.compile(r'^(\d{8}_\d{4})_\[')
    count = 0
    for root, dirs, files in os.walk(base_dir):
        # 2026-02-04, 2026-02-05 폴더만 대상
        if '2026-02-04' not in root and '2026-02-05' not in root:
            continue
            
        for filename in files:
            if not filename.endswith('.md'):
                continue
            
            # WorkflowID가 누락된 경우 (YYYYMMDD_HHmm_ 뒤에 바로 [ 가 오는 경우)
            if pattern.match(filename):
                new_filename = pattern.sub(r'\1_IDR_[', filename)
                old_path = os.path.join(root, filename)
                new_path = os.path.join(root, new_filename)
                
                if not os.path.exists(new_path):
                    try:
                        os.rename(old_path, new_path)
                        print(f"Renamed: {filename} -> {new_filename}")
                        count += 1
                    except Exception as e:
                        print(f"Error renaming {filename}: {e}")
                else:
                    print(f"Skip (Already exists): {filename}")
            
            # EVT 식별자를 IDR로 통합 (사용자 요청 기반)
            if '_EVT_[' in filename:
                new_filename = filename.replace('_EVT_[', '_IDR_[')
                old_path = os.path.join(root, filename)
                new_path = os.path.join(root, new_filename)
                if not os.path.exists(new_path):
                    os.rename(old_path, new_path)
                    print(f"Unified: {filename} -> {new_filename}")
                    count += 1
                    
    print(f"Total {count} files refactored.")

if __name__ == "__main__":
    rename_reports(r"d:\dev\trading\reports")
