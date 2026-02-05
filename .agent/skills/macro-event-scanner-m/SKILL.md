---
name: macro-event-scanner [M]
description: 거시적 이벤트(뉴스/공시) 발생 시 섹터 간 수급 이동 및 NXT 실시간 반응을 분석하는 워크플로우
---

## 🚨 글로벌 규약 (Global Rules)
1. **한국어 우선 (Korean-First)**: 모든 리포트의 헤더와 핵심 용어는 '한국어'로 작성합니다. (예: 'Overview' -> '개요')
2. **참고 자료 의무화 (Mandatory Reference)**: 모든 리포트 하단에 `## 참고 자료 (Reference)` 섹션을 만들고, 분석에 사용된 링크(Yahoo, DART 등)를 명시합니다.
3. **폴더 규칙 준수**: `d:\dev\trading\reports\YYYY-MM-DD\` 경로에 저장합니다.

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Fact-Based**: 모든 전망은 수치(환율, 금리, 선물 지수)를 근거로 해야 합니다.
- **Global Correlation**: 한국 시장은 글로벌 변수(미국장, 환율)에 종속적이므로, 이를 최우선으로 분석합니다.

## 📋 파일명 네이밍 컨벤션
### 형식
`YYYYMMDD_HHmm_WorkflowID_[태그]_종목명_분석명.md`
(예: `20260125_1930_IDR_[매크로]_전체_매크로보고서.md`)

## 🔗 상호 참조 (Cross-Reference)
- 동일 세션(`YYYYMMDD_HHmm`)에서 생성되는 다른 전문가들의 보고서 파일명을 예상하여 `## 참고 자료` 섹션에 수동으로 링크를 포함합니다.

## 🛠️ 도구 사용 지침 (Tool Discipline)
- **명세 우선 확인**: KIS API 호출 전 `find_api_detail` 조회가 필수입니다.
- **필수값 누락주의**: `env_dv` (실전/모의 구분) 파라미터를 절대 누락하지 마십시오.

## 🎯 분석 목적
- 뉴스, 공시, 매크로 지표 변화 등 **'포괄적 이벤트'** 발생 시 시장의 질적 변화 감지.
- 정규장과 NXT 애프터마켓의 괴리를 통해 내일의 주도 섹터를 선점.

---

## 🛠️ 활용 도구 및 API
- **KIS API**: `news_title` (시장 전체 뉴스), `fluctuation` (섹터별 등락), `inquire_overtime_price` (NXT 가격)
- **Naver Finance**: `get_stock_news` (종목별 심리), `get_trading_trends` (수급 주체 확인)
- **Macro**: `search_web` (환율, 금리, 미 반도체 관세 및 지수 선물 등 외부 변수)
- **Yahoo Finance**: 
  - **현물(Spot)**: 나스닥 종합(`^IXIC`), S&P 500(`^GSPC`), 다우 존스(`^DJI`), 필라델피아 반도체(`^SOX`) 실시간 및 전일 종가 확인.
  - **선물(Future)**: 나스닥 100 선물(`NQ=F`), S&P 500 선물(`ES=F`), USD/KRW(`USDKRW=X`), 비트코인(`BTC-USD`) 실시간 확인.

---

## 📈 분석 프로세스 (Process)

### 1단계: 글로벌 상황 및 임팩트 이벤트 식별
- **Action**: 전일 미 증시 현물(Spot) 종가와 현재 선물(Future) 등락률을 비교 분석.
- **Index Check**: 필라델피아 반도체(`^SOX`) 지수와 나스닥 종합(`^IXIC`) 지수의 전일 종가를 기준으로 국내 반도체/테크 섹터 영향력 예측.
- **Evidence**: 리포트 상단에 [현물 종가 vs 선물 실시간] 표를 구성하여 시계열적 흐름 제시.

### 2단계: 섹터 및 수급 분석 (Neutral Perspective)
- **Action**: 주도 섹터의 상승 동력 검증 시, 반드시 **'하락 가능성(Bear Case)'** 혹은 **'재료 오염도'**를 함께 체크.
- **Check**: 외국인/기관의 매도세가 특정 섹터에 집중되는지, 혹은 다른 섹터로의 '회피 매수'가 발생하는지 확인.

### 3단계: NXT 및 선물 시장 실측
- **Action**: 정규장 마감 후 NXT 애프터마켓 및 글로벌 지수 선물의 움직임을 통한 실거래 심리 확인.
- **Logic**: 
  - 현물(Spot) 마감가 대비 선물(Future)이 낮은 경우 → 시초가 하락 압력(보수적).
  - 환율(`USDKRW=X`) 변동폭이 10원 이상인 경우 → 수급 변동성 주의.

### 4단계: 대응 시나리오 수립
- **Action**: [강세 시나리오]와 [약세 시나리오]를 균형 있게 작성. 특정 섹터에 편중된 긍정론 지양.

---

## 📄 리포트 템플릿 (Report Template)
1. **이벤트 요약**: 날짜, 시간, 핵심 키워드
2. **글로벌 지수 현황 (Market Status)**: 현물(^SOX, ^IXIC 등 전일 종가) 및 실시간 선물 지표 표 정리
3. **주요 뉴스 및 이벤트 심층 분석**: 팩트체크 및 양면적 시사점(Bull/Bear Implication)
4. **대응 시나리오**: 공격적(추격 매수) vs 보수적(낙폭 과대 대기) vs 대안(인버스/금)
5. **요약 및 결론**
6. **참고 자료 (Reference)**: [필수] Yahoo Finance 지수 링크 및 뉴스 탭

---

**Next Action**: 분석된 결과를 바탕으로 `market-sector-scanner [M]` 또는 `total-trading-strategy [I]`로 연결.
