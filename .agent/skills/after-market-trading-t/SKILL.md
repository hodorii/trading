---
name: after-market-trading [T]
description: 장 마감 후 시간외 거래(NXTrade) 및 공시 분석을 통한 익일 대응 전략 수립
---

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 스킬은 **`GEMINI.md`**에 정의된 **'공통 보고서 프로토콜'** 및 **'핵심 원칙'**을 엄격히 준수합니다.

## 🛠️ 도구 사용 지침
- **데이터 대조 필수 (Cross-Check)**: `inquire_overtime_price` 호출 시 `fid_cond_mrkt_div_code='NX'`(단일가/야간)와 `'J'`(일반 시간외) 데이터를 반드시 동시에 조회하여 대조합니다.
- **체결량 확인**: `ovtm_untp_vol`이 0인 경우 해당 시간대의 데이터가 아직 생성되지 않았거나 거래가 없는 것이므로, 이를 장중 종가와 혼동하여 보고하지 않습니다.

## 🎯 분석 목적
- 장 마감 후 발생한 공시/뉴스의 실제 수급 반영 여부 확인.
- 시간외(NXTrade) 가격 변동과 호가 잔량을 통한 익일 시초가 신뢰도 측정.

---

## 🛠️ 활용 도구 및 API
- **KIS API**: `inquire_overtime_price` (시간외/NXTrade 가격), `news_title` (최신 공시)
- **Naver**: `get_disclosure_list`, `get_stock_news`

---

## 📈 분석 프로세스 (Process)

### 0단계: 데이터 무결성 검증 (Data Validation) - [신설]
- **Action**: `fid_cond_mrkt_div_code`를 `J`와 `NX`로 각각 호출.
- **Check**: `NX` 데이터의 `ovtm_untp_prpr`(체결가)가 장중 종가와 동일하고 `vol`이 0이라면, 아직 시간외 거래 데이터가 유효하지 않은 상태임을 명시.

### 1단계: 뉴스 및 공시 스크리닝
- **Action**: 장 마감(15:30) 이후 발생한 DART 공시 및 핵심 뉴스 필터링.
- **Extract Dates**: 공시 내용 중 주주명부폐쇄일, 배당기준일, 변경상장일, 합병기일 등 **핵심 일정(Key Dates)**을 반드시 추출하여 '예정일'로 기록.
- **Check**: 재료의 시의성(Timeliness)과 시간외 수급과의 상관관계 분석.

### 2단계: 시간외(NXTrade) 실측 수급 분석
- **Action**: 시간외 단일가 및 야간 거래의 체결 강도와 **매수/매도 호가 잔량(Bid/Ask Spread)** 확인.
- **Check**: 허수 주문 여부 판별 및 실제 돈의 흐름(Smart Money) 추적.

### 3단계: 익일 시나리오 및 전술 수립
- **Action**: 실제 체결된 시간외 평균 단가를 기준으로 익일 시초가 예상 범위(Gap Range) 설정.
- **Strategy**: 갭 상승/하락 폭에 따른 비중 조절 및 손절선 재설정.

---

## 📄 리포트 템플릿 (Report Template)
1. **데이터 검증**: 현재 참조 중인 시간외 데이터의 유효성 (NX vs J 대조 결과)
2. **시간외 실측**: NX 체결가, 등락률, 호가 잔량 상태
3. **공시/재료 분석**: 
    - 장후 발표된 재료의 호재/악재 강도
    - **📅 주요 예정일 (Schedule)**: 배당/분할/합병 등 캘린더 이벤트 명시
4. **익일 시나리오**: 시초가 예측 및 장전 대응 구간
5. **결론**: 오버나이트 전략 및 비중 관리 제언

---

**Next Action**: 분석 완료 후 `integrated-daily-routine` 및 `multi-agent-decision`의 기초 팩트로 활용.
