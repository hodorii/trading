---
name: market-sector-scanner [M]
description: 금일 주식 시황 전체 분석 및 주도 섹터와 종목 추출을 위한 마켓 스캐너
---

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 스킬은 **`GEMINI.md`**에 정의된 **'공통 보고서 프로토콜'** 및 **'핵심 원칙'**을 엄격히 준수합니다.

# 🔍 market-sector-scanner [M]

## 🛠️ 도구 사용 지침
- API 호출 전 `find_api_detail` 확인 필수.

## 🎯 분석 목적
- 금일 시장을 주도한 '돈의 흐름' 파악.
- 상승한 섹터가 '진짜'인지 확인하기 위한 글로벌 피어 체크.
- 내일 공략할 후보 종목군 추출.

---

## 🛠️ 활용 도구 및 API
- **KIS API**: `fluctuation`, `volume_rank`, `market_cap`
- **Naver**: `get_trading_trends`, `get_stock_news`
- **Global**: `search_web` (미국/글로벌 경쟁사 주가 확인)

---

## 📈 분석 프로세스 (Process)

### 1단계: 시장 요약 및 매크로 연동성 체크 (Market & Macro Link)
- **Input**: `macro-event-scanner`에서 도출된 **'미국 감응 유망 섹터'**를 최우선 분석 대상으로 설정.
- **Action**: 코스피/코스닥 지수 및 외국인 수급 현황 파악.
- **Check**: 매크로 변수(환율, 금리, 미 선물)가 특정 섹터(반도체, 자동차 등)에 미치는 영향 분석.

### 2단계: 주도 섹터 및 테마 발굴 (Dual Track Screening)
- **Track A (Macro-Driven)**: 미국 빅테크/정책 수혜 섹터 (예: AI 반도체, 전력, 방산)
    - **Peer Check**: 관련 미국 종목(Nvidia, Tesla, GE Vernova 등)의 주가 흐름과 동조화(Coupling) 여부 확인.
- **Track B (Market-Driven)**: 당일 국내 시장 거래대금/등락률 상위 주도 테마 (예: 바이오, 로봇, 정치인 테마 등)
    - **Ranking**: 상승률 상위 섹터의 지속성 및 재료의 크기(Size) 평가.

### 3단계: 주도 종목 정밀 필터링 (Top Pick Selection)
- **Filtering**:
    - **대장주(Leader)**: 거래대금 상위, 신고가 경신, 외국인 수급 집중 종목.
    - **낙수효과(Laggard)**: 대장주 상승 시 후발 주자로 부각될 수 있는 2등주 식별.
- **US-Linked Impact**: 'SaaSpocalypse'나 '관세' 등 글로벌 이슈에 노출된 종목은 **리스크 요인**으로 별도 분류.

### 4단계: 다음 관심 섹터 및 전략
- **Action**: 순환매가 예상되는 차기 섹터(낙폭 과대 등) 선정.

---

## 📄 리포트 템플릿 (Report Template)
1. **시장 요약 (Market Summary)**: 지수, 거래대금, 수급, 특이사항
2. **주도 섹터 및 테마**: 섹터명, 등락률, 핵심 종목, 상승 동력
3. **주도 종목 정밀 필터링**: 대형주/중소형주 최선호주
4. **다음 관심 섹터**: 순환매 길목 지키기
5. **최종 추천 리스트**: 통합 전략 수립을 위한 후보군
6. **참고 자료 (Reference)**: [필수] 네이버 금융 및 글로벌 증시 링크

---

**Next Action**: 발굴된 종목에 대해 `total-trading-strategy [I]`를 통해 세부 전략 수립.
