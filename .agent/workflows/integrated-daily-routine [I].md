---
description: 매크로, 이벤트, 섹터, 가치, 수급, 리스크, 전략을 총망라하여 글로벌 규칙 준수 여부를 검증하는 마스터 루틴
---

# 🔄 integrated-daily-routine [I] (MoE 전문가 믹스 루틴)

## 🎯 목적
매크로(거시)부터 리스크(위험)까지 7개의 특화된 **전문가(Expert)** 모델을 순차적으로 로딩하여, 데이터 오염 없이 최적의 의사결정을 도출하는 MoE 기반 마스터 워크플로우입니다.

## 🚨 SRP 전문가 및 네이밍 준수 사항 (Strict Rules)
1.  **시작 시각 고정**: 워크플로우를 시작하는 시점의 시각(HHmm)을 모든 산출물의 파일명에 동일하게 적용합니다. (예: 12시 30분 시작 시 모든 파일은 `20260127_1230_...` 으로 시작)
2.  **파일명 네이밍 컨벤션**: 반드시 **`YYYYMMDD_HHmm_IDR_[태그]_종목명_설명.md`** 형식을 준수합니다. (`IDR`은 본 워크플로우 고유 식별자입니다)
3.  **단일 책임(SRP)**: 각 전문가는 고유 영역의 데이터만 분석하며 타 영역의 판단에 간섭하지 않습니다.
4.  **지수 입체 분석**: 모든 전문가는 실행 전 [미 현물 종가 vs 실시간 선물] 데이터 체인을 로딩합니다.
5.  **폴더 규칙**: 모든 산출물은 `d:\dev\trading\reports\YYYY-MM-DD\` 에 저장합니다.
6.  **상호 참조 (Cross-reference)**: 모든 보고서 하단 `## 참고 자료` 섹션에 본 세션에서 생성될(혹은 생성된) 다른 보고서들의 링크를 예상 파일명 규칙에 따라 미리 포함하여 세트임을 명시합니다.
7.  **종목 코드 검증 (Code-Accuracy)**: 분석 시작 전 반드시 `KIS find_stock_code`와 `Naver search_stock`을 교차 확인하여 정확한 6자리 코드를 확정합니다. KIS 검색 실패 시 네이버 결과를 우선합니다.

---

## 🛠️ 전문가 믹스 파이프라인 (MoE Pipeline)

### 1단계: 매크로 전문가 (Macro Expert Load)
-   **Expert Skill**: `macro-event-scanner [M]`
-   **Output**: `YYYYMMDD_HHmm_IDR_[매크로]_전체_매크로보고서.md`
-   **Task**: 오늘의 시장 기상도 확정 (KRX-NXT 가격 괴리 및 08:30 거래량 급증 테마 선제 포착).

### 2단계: 이벤트 전문가 (Event Expert Load)
-   **Expert Skill**: `event-driven [M]`
-   **Output**: `YYYYMMDD_HHmm_IDR_[이벤트]_주요이벤트_분석보고서.md`
-   **Task**: 뉴스 신선도 및 정치적 오염도 판별. **(필수: Yahoo News를 통한 글로벌 시각 교차 검증)**

### 3단계: 섹터 전문가 (Sector Expert Load)
-   **Expert Skill**: `market-sector-scanner [M]`
-   **Output**: `YYYYMMDD_HHmm_IDR_[시장]_전체_섹터스캐너.md`
-   **Task**: 주도 섹터 발굴 및 글로벌 피어(Peer) 동조화 확인.

### 4단계: 가치 전문가 (Value Expert Load)
-   **Expert Skill**: `fundamental-analysis [I]`
-   **Output**: `YYYYMMDD_HHmm_IDR_[가치]_종목명_기업분석보고서.md`
-   **Task**: 철저히 재무제표 및 글로벌 밸류에이션 기반 적정가 산출. **(필수: 국내외 애널리스트 의견 괴리 분석 포함)**

### 5단계: 수급 전문가 (Flow Expert Load)
-   **Expert Skill**: `investor-matrix [I]`
-   **Output**: `YYYYMMDD_HHmm_IDR_[수급]_종목명_수급분석보고서.md`
-   **Task**: 실시간 외국인/기관 매매 대조 및 세력 평단가 추적.

### 6단계: 리스크 전문가 (Risk Expert Load)
-   **Expert Skill**: `risk-strategist [I]`
-   **Output**: `YYYYMMDD_HHmm_IDR_[리스크]_종목명_위험분석보고서.md`
-   **Task**: 공매도, 대차, 오버행 및 하락 시나리오 수립.

### 7단계: 전략 통합가 (Final Strategist Load)
-   **Expert Skill**: `total-trading-strategy [I]`
-   **Output**: 
    1. `YYYYMMDD_HHmm_IDR_[결정]_종목명_최종통합전략.md`
    2. `YYYYMMDD_HHmm_IDR_[가이드]_종목명_실전쉬운가이드.md`
-   **Task**: 위 6개 전문가 보고서를 통합하여 최종 투자 비중 확정.

### 8단계: 통합 리포터 (Total Reporter Load)
-   **Expert Skill**: `total-trading-strategy [I]` (Extended Reporting)
-   **Output**: `YYYYMMDD_HHmm_IDR_[통합]_전체_종합분석보고서.md`
-   **Task**: 모든 전문가의 분석 결과(이벤트 영향, 섹터 스캔, 자금 이동, 10선 추천 등)를 하나로 집대성한 최종 요약본 생성.

---


## 📝 최종 점검 (Final Check)
- [ ] 모든 파일이 동일한 `YYYYMMDD_HHmm` 접두사로 통일되었는가?
- [ ] 7단계 전문가의 SRP가 정확히 이행되었는가?
- [ ] Bull/Bear 관점이 균형 있게 포함되었는가?
