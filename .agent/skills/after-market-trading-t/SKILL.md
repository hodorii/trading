---
name: after-market-trading [T]
description: 장 마감 후 시간외 거래(NXTrade) 및 공시 분석을 통한 익일 대응 전략 수립
---

## 🚨 글로벌 규약 (Global Rules)
1. **한국어 우선 (Korean-First)**: 모든 분석 결과는 한국어로 작성합니다.
2. **참고 자료 의무화 (Mandatory Reference)**: DART 공시 링크 및 NXTrade 체결 화면 요약을 포함합니다.
3. **폴더 규칙 준수**: `d:\dev\trading\reports\YYYY-MM-DD\` 경로에 저장합니다.

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Sustainability**: 시간외 상승이 단발성 뉴스인지, 지속 가능한 모멘텀인지 판별합니다.
- **NX vs Regular**: 정규장 종가와 NXTrade 체결가 자금 흐름의 괴리를 분석합니다.

## 📋 파일명 네이밍 컨벤션
### 형식
`YYYYMMDD_HHmm_[태그]_종목명_설명.md`
(예: `20260127_1630_[애프터]_우리기술_시간외분석보고서.md`)

# 🌙 after-market-trading [T] (시간외 분석)

## 🛠️ 도구 사용 지침
- **NXTrade**: `inquire_overtime_price` 호출 시 `fid_cond_mrkt_div_code='NX'` 필수.

## 🎯 분석 목적
- 장 마감 후 발생한 공시/뉴스의 임팩트 측정.
- 시간외 수급 주체 및 가격 변동을 통한 내일 시초가 예측.

---

## 🛠️ 활용 도구 및 API
- **KIS API**: `inquire_overtime_price` (시간외 가격), `news_title` (최신 공시)
- **Naver**: `get_disclosure_list`, `get_stock_news`

---

## 📈 분석 프로세스 (Process)

### 1단계: 뉴스 및 공시 스크리닝
- **Action**: 장 마감 후 15:30 이후 발생한 모든 공시 및 중요 뉴스 취합.
- **Check**: 유상증자, 단일판매공급계약, 최대주주변경 등 주가 영향력 판단.

### 2단계: 시간외(NXTrade) 수급 추적
- **Action**: 시간외 단일가(16:00~18:00) 체결량 및 가격 추이 모니터링.
- **Check**: 거래량을 동반한 상승인지, 소량으로 띄운 허수인지 판별.

### 3단계: 익일 시나리오 수립
- **Action**: 시간외 종가를 기준으로 내일 예상 시초가 범위 설정.
- **Strategy**: 갭 상승 시 추격 매수 금지 / 갭 하락 시 지지선 매수 등 전술 수립.

---

## 📄 리포트 템플릿 (Report Template)
1. **시간외 요약**: NX 종가, 등락률, 거래량
2. **뉴스/공시 분석**: 주요 재료 요약 및 호재/악재 판별
3. **익일 시뮬레이션**: 시초가 예측 및 대응 구간 설정
4. **결론**: 오버나이트(보유) / 비중 축소 / 시초가 매도
5. **참고 자료 (Reference)**: [필수] DART 공시 링크, NXTrade 체결 데이터

---

**Next Action**: 분석 완료 후 `integrated-daily-routine`의 최종 단계로 전달.
