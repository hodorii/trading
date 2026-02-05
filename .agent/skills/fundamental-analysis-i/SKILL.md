---
name: fundamental-analysis [I]
description: 기업 본질 가치 및 글로벌 밸류에이션 딥다이브를 위한 펀더멘털 분석 워크플로우
---

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 스킬은 **`GEMINI.md`**에 정의된 **'공통 보고서 프로토콜'** 및 **'핵심 원칙'**을 엄격히 준수합니다.
- **원천 데이터 링크**: 분석에 사용된 네이버/야후/DART의 개별 페이지 URL을 각 섹션 하단에 배치.

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Value Investing**: 가격(Price)과 가치(Value)의 괴리를 찾습니다.
- **Global Peer Check**: 국내 기업의 밸류에이션 적정성은 글로벌 1등 기업과의 비교(Peer Valuation)를 통해 검증합니다.
- **Why-First (핵심)**: 변동 원인을 요약하되 본문 헤더(#)에는 세션 정보를 포함하지 않습니다. 결론뿐 아니라 근거 수치와 데이터를 상세 표 형태로 반드시 제시해야 합니다.

## 📄 데이터 증거 요건 (Data Evidence Requirement)
- **재무 테이블**: 최근 3개년 및 4분기 주요 재무지표(매출, 영업이익, 당기순이익, OPM)를 마크다운 표로 작성.
- **밸류에이션 지표**: PER, PBR, ROE 등을 피어 그룹과 일대일 대조한 수치 데이터 포함.

# 💎 fundamental-analysis [I] (가치 분석)

## 🛠️ 도구 사용 지침
- **교차 검증**: API 데이터와 실제 공시(DART) 데이터의 정합성을 확인합니다.

## 🎯 분석 목적
- 기업의 재무 건전성(안전마진) 및 성장성(Upside) 확인.
- 저평가 여부 판단 및 목표 주가(Fair Value) 산출.

---

## 🛠️ 활용 도구 및 API
- **Global**: `mcp_yahoo-finance_get_income_statement`, `get_current_stock_price` (해외 경쟁사)
- **Domestic**: `mcp_naver-finance_get_stock_price` (PER/PBR), `get_disclosure_list` (최신 공시), `get_stock_news` (최신 뉴스)
- **Consensus**: `mcp_yahoo-finance_get_recommendations`

---

## 📈 분석 프로세스 (Process)

### 1단계: 글로벌 피어 밸류에이션 (Peer Valuation)
- **Action**: 분석 대상 기업의 글로벌 경쟁사(예: 삼성전자 vs 마이크론) PER/PBR 비교.
- **Logic**: 글로벌 피어 대비 과도한 할인(Discount)인지 정당한 할인인지 판별.

### 2단계: 실적 성장성 및 건전성 (Growth & Health)
- **Action**: 최근 4분기 매출/영업이익 추이(YoY, QoQ) 확인.
- **Check**: 부채비율, 유보율, 영업이익률(OPM) 추세 분석.

### 3단계: 뉴스 및 공시 기반 가치 임팩트 분석 (Catalyst Check)
- **Action**: 최근 3개월 내 핵심 뉴스/공시(M&A, 수주, 특허, 실적 가이던스 등)가 펀더멘털에 미치는 영향 수치화.
- **Check**: 일시적 호재인가, 영구적인 이익 구조 개선인가 판별.

### 4단계: 적정 주가 및 글로벌 컨센서스 (Target Price & Global Consensus)
- **Action**: PER 멀티플, PBR 밴드, 혹은 국내외 애널리스트 컨센서스 기반 적정 주가 추정.
- **IB Report Check**: (필수) **골드만삭스, JP모간, 모건스탠리** 등 글로벌 IB의 해당 종목 최신 리포트 요약 및 목표주가 변동 추이 확인.
- **Gap**: 현재 주가와 적정 주가의 괴리율(%) 계산 (상승 여력).

---

## 📄 리포트 템플릿 (Report Template)
1. **가치 요약**: 적정 주가, 상승 여력, 투자 매력도
2. **펀더멘털 체크**: 성장성(매출/이익), 수익성(OPM), 안정성(부채)
3. **글로벌 피어 비교**: 경쟁사 대비 밸류에이션 매력도
4. **최신 뉴스/공시 임팩트**: 최근 주요 재료가 가치에 미치는 영향 (Bull/Bear 임팩트)
5. **컨센서스**: 증권사 투자의견 및 목표가 추이
6. **결론**: 저평가(매수) / 적정가(보유) / 고평가(매도)
7. **참고 자료 (Reference)**: [필수] 네이버 금융 종목분석, Yahoo 재무 탭, 주요 공시 링크

---

**Next Action**: 펀더멘털 분석 후 `investor-matrix`로 이동하여 해당 가치를 수급이 알아봐주고 있는지 확인.
