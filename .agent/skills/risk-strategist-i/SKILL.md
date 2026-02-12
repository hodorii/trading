---
name: risk-strategist [I]
description: 공매도 대차거래 및 매크로 위험 요인을 판별하는 전략 리서치 워크플로우
---

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 스킬은 **`GEMINI.md`**에 정의된 **'공통 보고서 프로토콜'** 및 **'핵심 원칙'**을 엄격히 준수합니다.

# 🛡️ risk-strategist [I] (리스크 전략가)

## 🛠️ 도구 사용 지침
- **필수 조회**: `daily_short_sale`(공매도), `daily_loan_trans`(대차) 전 `find_api_detail` 확인.

## 🎯 분석 목적
- 공매도, 대차거래, 신용잔고 등 수급상의 잠재적 매도 압력 분석.
- 전환사채(CB), 유상증자 등 오버행(대량 대기 매물) 리스크 점검.
- 글로벌/매크로 변수에 따른 하락 시나리오(Bear Case) 수립.

---

## 🛠️ 활용 도구 및 API
- **Macro**: `mcp_yahoo-finance_get_current_stock_price` (`^VIX`, `NQ=F`)
- **Short/Loan**: `daily_short_sale`, `daily_loan_trans`
- **Disclosure**: `search_stock_info`, `mcp_naver-finance_get_disclosure_list`

---

## 📈 분석 프로세스 (Process)

### 1단계: 시장 위험도 측정 (Systemic Risk)
- **Action**: VIX 지수 및 코스피 변동성 지수 확인.
- **Check**: 시장 전체가 공포 구간인지 탐욕 구간인지 판단.

### 2단계: 수급 파생 리스크 (Short & Loan)
- **Action**: 대차잔고 증감 추이 및 공매도 비중(일별) 확인.
- **Logic**: 대차잔고 급증 = 공매도 대기 물량 증가 (하락 압력).

### 3단계: 오버행 및 재무 리스크 (Overhang)
- **Action**: 미상환 전환사채(CB), 신주인수권부사채(BW) 잔액 및 전환가액 확인.
- **Calc**: 현재 주가가 전환가액보다 높다면 물량 출회 가능성 높음.

### 4단계: 매크로 민감도 테스트
- **Action**: 환율/금리 급등 시 기업 이익 훼손 가능성 시뮬레이션.

### 5단계: De-coupling Check (개별 vs 매크로 괴리)
- **Logic**: 매크로(미국장/섹터) 호재에도 불구하고 개별 기업 악재(소송, 지배구조, 기술 결함)가 발생한 경우, **개별 악재를 최우선(Override)**으로 반영.
- **Action**: 글로벌 동종 업체 주가와 국내 기업 주가 간 괴리율(Divergence)이 확대되면 'Structural Risk'로 간주.

---

## 📄 리포트 템플릿 (Report Template)
1. **위험 요약 (Executive Lowlights)**: 리스크 점수 (1~10점) 및 핵심 경고
2. **De-coupling 진단**: 매크로 호재와 개별 악재 간 충돌 여부 (괴리 심화 시 **'Strong Sell'**)
3. **수급 리스크**: 공매도/대차/신용잔고 현황 (검은머리 외국인 포함)
4. **오버행 리스크**: CB/BW 잔액 및 희석 가능성
5. **매크로 민감도**: 환율/유가/금리 등 외부 충격 시나리오
6. **방어 전략**: 손절가 및 헤지(Hedge) 방안
7. **참고 자료 (Reference)**: [필수] KRX 공매도 통계, DART 오버행 공시 링크

---

**Next Action**: 리스크 분석 완료 후 `total-trading-strategy`로 이동하여 최종 의사결정.
