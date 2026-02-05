---
name: investor-matrix [I]
description: 돈의 경로를 추적하는 수급 전문 수사 보고서 워크플로우
---

## 🚨 글로벌 규약 (Global Rules)
1. **한국어 우선 (Korean-First)**: 모든 리포트 헤더와 용어는 한국어로 작성합니다.
2. **참고 자료 의무화 (Mandatory Reference)**: 리포트 하단에 분석에 활용된 '네이버 금융 수급' 및 'Yahoo Finance' 링크를 반드시 포함합니다.
3. **폴더 규칙 준수**: `d:\dev\trading\reports\YYYY-MM-DD\` 경로에 저장합니다.

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Flow Follows Value**: 돈은 결국 가치가 있는 곳, 혹은 이슈가 있는 곳으로 흐릅니다.
- **Trace the Smart Money**: 주가를 움직이는 '진짜 주체(Main Actor)'를 찾아냅니다.

## 📋 파일명 네이밍 컨벤션
### 형식
`YYYYMMDD_HHmm_WorkflowID_[태그]_종목명_분석명.md`
(예: `20260127_1230_EVT_[수급]_우리기술_수급분석보고서.md`)

## 🔗 상호 참조 (Cross-Reference)
- 동일 세션(`YYYYMMDD_HHmm`)에서 생성되는 다른 전문가들의 보고서 파일명을 예상하여 `## 참고 자료` 섹션에 수동으로 링크를 포함합니다.

# 🕵️ investor-matrix [I] (수급 포렌식)

## 🛠️ 도구 사용 지침
- **명세 우선 확인**: `find_api_detail` 조회 필수.
- **필수값 누락주의**: `env_dv` 파라미터 필수.

## 🎯 분석 목적
- 현재 주가 상승/하락을 주도하는 메이저 주체(외국인/기관) 특정.
- 주포의 **평균 단가(VWAP)** 및 **손익분기점(BEP)** 추정.
- 수급의 질(Quality) 분석 (단타성 vs 매집성).

---

## 🛠️ 활용 도구 및 API
- **Global**: `mcp_yahoo-finance_get_current_stock_price` (`NQ=F`, `ES=F`)
- **수급 추적**: `mcp_naver-finance_get_trading_trends`, `mcp_naver-finance_get_brokerage_info`
- **KIS**: `inquire_investor_trend` (투자자별 매매동향)

---

## 📈 분석 프로세스 (Process)

### 1단계: 글로벌 수급 환경 (Global Context)
- **Action**: 나스닥 선물(`NQ=F`) 및 S&P 500 선물(`ES=F`) 흐름 확인.
- **Logic**: 글로벌 Risk-On 환경이어야 외국인 수급 유입이 원활함.

### 2단계: 메이저 주체 식별 및 당일 흐름 (Identify Main Actor)
- **Action**: 최근 5일/20일 누적 순매수 주체와 함께 **'금일 실시간(Intraday)'** 순매수/순매도 추이 확인.
- **Check**: 외국인과 기관의 매매 방향이 일치하는지(쌍끌이), 혹은 대립하는지(손바뀜) 확인.
- **Sector Mapping**: 외국인이 매도하는 섹터를 기관이 방어하는지(방어력 테스트), 혹은 동반 매도하는지(급락 리스크) 대조 분석.

### 3단계: 주포 평단가 및 수익 현황 (Cost Basis)
- **Action**: 메이저 주체의 추정 평단가와 현재가의 괴리율을 산출하여 **'차익 실현 압력'** 혹은 **'물타기/추가매수 가능성'** 분석.
- **Inference**: 현재 주가가 세력 평단가 대비 과도하게 올랐다면(Bullish) 단기 조정 리스크(Bearish)를 함께 경고.

### 4단계: 수급 이탈 및 심리 지표 확인
- **Action**: 거래대금 상위 증권사(창구) 분석을 통해 단타성 외인(검은머리 외국인 등)의 이탈 징후 포착.
- **Risk Check**: 정치적/사회적 이슈 발생 시 외국인의 '패닉 셀' 발생 여부 실시간 모니터링.

---

## 📄 리포트 템플릿 (Report Template)
1. **수급 요약**: 주도 주체, 외국인 보유율 변화
2. **주포 분석**: 메이저 주체 특정 및 성격 분석
3. **평단가 추정**: 세력 추정 매집 단가 vs 현재가 괴리율
4. **결론**: 수급 우수(매수) / 수급 꼬임(관망) / 수급 이탈(매도)
5. **참고 자료 (Reference)**: [필수] 네이버 금융 투자자별 매매동향 링크

---

**Next Action**: 수급 분석 결과를 바탕으로 `risk-strategist`로 넘어가 리스크 점검 수행.
