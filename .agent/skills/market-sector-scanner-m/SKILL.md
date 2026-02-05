---
name: market-sector-scanner [M]
description: 금일 주식 시황 전체 분석 및 주도 섹터와 종목 추출을 위한 마켓 스캐너
---

## 🚨 글로벌 규약 (Global Rules)
1. **한국어 우선 (Korean-First)**: 리포트 내 영어 사용을 지양하고 한국어 용어를 우선 사용합니다.
2. **참고 자료 의무화 (Mandatory Reference)**: 리포트 하단에 데이터 출처(Reference)를 반드시 명시합니다.
3. **폴더 규칙 준수**: `d:\dev\trading\reports\YYYY-MM-DD\` 경로에 저장합니다.

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Top-Down**: 지수 -> 섹터 -> 종목 순으로 분석합니다.
- **Global Linkage**: 국내 주도주는 반드시 글로벌 동종 기업(Peer)과 비교 분석해야 합니다. (커플링/디커플링 확인)

## 📋 파일명 네이밍 컨벤션
### 형식
`YYYYMMDD_HHmm_WorkflowID_[태그]_종목명_설명.md`
(예: `20260127_1230_EVT_[시장]_전체_섹터스캐너.md`)

## 🔗 상호 참조 (Cross-Reference)
- 동일 세션(`YYYYMMDD_HHmm`)에서 생성되는 다른 전문가들의 보고서 파일명을 예상하여 `## 참고 자료` 섹션에 수동으로 링크를 포함합니다.

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

### 1단계: 시장 요약 및 글로벌 연동성 체크
- **Action**: 코스피/코스닥 지수 및 수급(외국인/기관) 현황 파악.
- **Check**: 환율(USDKRW) 변동에 따른 외국인 수급 영향 분석.

### 2단계: 주도 섹터 및 테마 발굴
- **Action**: 등락률 상위 섹터 도출.
- **Ranking**: 1위~3위 주도 섹터 선정 및 상승 동력(뉴스) 매칭.

### 3단계: 주도 종목 정밀 필터링 (Top Pick)
- **Action**: 거래대금 터진 대장주 위주 선정.
- **Global Peer Check (중요)**: 선정된 종목의 글로벌 경쟁사(예: SMR - Oklo, NuScale) 주가 흐름 확인.
    - 같이 오르면? → 강력한 테마/업황 호조.
    - 미 주식 빠졌는데 국장만 오르면? → **독자적 호재(Decoupling)** 인지 아니면 **후행적 하락 리스크**인지 정밀 분석 필요.

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
