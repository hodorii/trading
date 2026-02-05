---
name: after-market-trading [T]
description: 장 마감 후 시간외 거래(NXTrade) 및 공시 분석을 통한 익일 대응 전략 수립
---

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 스킬은 **`GEMINI.md`**에 정의된 **'공통 보고서 프로토콜'** 및 **'핵심 원칙'**을 엄격히 준수합니다.

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
