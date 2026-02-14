#!/bin/bash
# Trading Pulse Archive Validator

POSTS_DIR="../trading-pulse/_posts"
REPORTS_DIR="reports"
ERROR_COUNT=0

echo "🔍 [Archive Validation] 분석 시작..."

# 1. 파일명 규칙 검사 (GP-FNAME)
for file in $(find $POSTS_DIR -name "*.md"); do
    filename=$(basename "$file")
    if [[ ! $filename =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{4}-[0-9]{2}-.*\.md$ ]]; then
        echo "❌ [GP-FNAME 위반] 파일명 형식이 잘못됨: $filename"
        ((ERROR_COUNT++))
    fi
done

# 2. 메타데이터(Front Matter) 및 필수 내용 검사
for file in $(find $POSTS_DIR -name "*.md"); do
    if ! grep -q "session_id:" "$file"; then echo "❌ [GP-META 위반] session_id 누락: $(basename "$file")"; ((ERROR_COUNT++)); fi
    if ! grep -q "session_order:" "$file"; then echo "❌ [GP-META 위반] session_order 누락: $(basename "$file")"; ((ERROR_COUNT++)); fi
    
    # 합성 보고서(SEQ 08) 내 TSM 표 존재 여부 검사
    if [[ $(basename "$file") =~ -08- ]]; then
        if ! grep -q "TSM" "$file"; then
            echo "⚠️ [Content 위반] 합성 보고서에 TSM 매트릭스 누락: $(basename "$file")"
            ((ERROR_COUNT++))
        fi
    fi
done

# 3. 로컬 동기화 검사
for file in $(find $POSTS_DIR -name "*.md"); do
    rel_path=$(echo "$file" | sed 's|.*/_posts/||')
    date_dir=$(echo "$rel_path" | cut -d'-' -f1-3)
    if [ ! -f "$REPORTS_DIR/$date_dir/$rel_path" ]; then
        echo "❌ [Sync 위반] 로컬 reports 폴더에 파일이 없음: $rel_path"
        ((ERROR_COUNT++))
    fi
done

if [ $ERROR_COUNT -eq 0 ]; then
    echo "✅ [SUCCESS] 모든 아카이빙 원칙이 준수되었습니다."
    exit 0
else
    echo "🚨 [FAILURE] 총 $ERROR_COUNT건의 원칙 위반이 발견되었습니다."
    exit 1
fi
