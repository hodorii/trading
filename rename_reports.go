package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 네이밍 룰: YYYYMMDD_HHMM_[분류]_대상_보고서명.md
// 예: 20260127_1330_[이벤트]_주요이벤트_분석보고서.md

func main() {
	reportsDir := "reports"

	// 네이밍 룰 정규식: YYYYMMDD_HHMM_[분류]_대상_보고서명
	validPattern := regexp.MustCompile(`^\d{8}_\d{4}_\[.+?\]_.+\.md$`)

	// 잘못된 형식 패턴 (YYMMDD 형식)
	wrongPattern := regexp.MustCompile(`^(\d{6})_(\d{4})_(.+)\.md$`)

	var renames []struct {
		old string
		new string
	}

	err := filepath.WalkDir(reportsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}

		filename := d.Name()

		// 이미 올바른 형식이면 스킵
		if validPattern.MatchString(filename) {
			return nil
		}

		// YYMMDD 형식을 YYYYMMDD로 변환
		if matches := wrongPattern.FindStringSubmatch(filename); matches != nil {
			yymmdd := matches[1]
			hhmm := matches[2]
			rest := matches[3]

			// YY를 20YY로 변환
			yyyymmdd := "20" + yymmdd
			newFilename := fmt.Sprintf("%s_%s_%s.md", yyyymmdd, hhmm, rest)

			oldPath := path
			newPath := filepath.Join(filepath.Dir(path), newFilename)

			renames = append(renames, struct {
				old string
				new string
			}{oldPath, newPath})

			fmt.Printf("변경 예정: %s -> %s\n", filename, newFilename)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("오류: %v\n", err)
		return
	}

	if len(renames) == 0 {
		fmt.Println("변경할 파일이 없습니다. 모든 파일이 네이밍 룰을 준수합니다.")
		return
	}

	fmt.Printf("\n총 %d개 파일을 변경합니다.\n", len(renames))
	fmt.Print("계속하시겠습니까? (y/n): ")

	var response string
	fmt.Scanln(&response)

	if strings.ToLower(response) != "y" {
		fmt.Println("취소되었습니다.")
		return
	}

	// 파일명 변경 실행
	for _, r := range renames {
		if err := os.Rename(r.old, r.new); err != nil {
			fmt.Printf("실패: %s -> %s (오류: %v)\n", r.old, r.new, err)
		} else {
			fmt.Printf("완료: %s -> %s\n", filepath.Base(r.old), filepath.Base(r.new))
		}
	}

	fmt.Println("\n파일명 변경이 완료되었습니다!")
}
