---
name: day-trading [T]
description: 장중 분봉 및 실시간 수급 실측을 통한 당일 매매 전술 워크플로우
---

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 스킬은 **`GEMINI.md`**에 정의된 **'공통 보고서 프로토콜'** 및 **'핵심 원칙'**을 엄격히 준수합니다.
- **실시간성**: `inquire_price`, `inquire_asking_price_exp_ccn` 등을 사용하여 최신 데이터 유지.

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Trend is Friend**: 당일 주도 섹터와 거래대금 상위 종목에만 집중합니다.
- **Risk Managed**: 손절가(Stop-loss) 미준수 시 어떠한 매매도 정당화될 수 없습니다.

# ⚔️ day-trading [T] (데이 트레이딩)

## 🎯 분석 목적
- 장중 발생하는 변동성을 수익으로 연결.
- 주요 지지/저항선 돌파 및 이탈 여부 실시간 감시.

---

## 🛠️ 활용 도구 및 API
- **KIS API**: `inquire_time_itemchartprice` (분봉), `inquire_asking_price_exp_ccn` (호가)
- **Naver**: `get_brokerage_info` (실시간 창구)

---

## 📈 분석 프로세스 (Process)

### 1단계: 시초가 및 초기 흐름 파악 (09:00~09:30)
- **Action**: 예상 체결가 대비 시초가 갭 강도 확인.
- **Check**: 전일 고가/저가 돌파 여부 및 거래량 동반 유무.

### 2단계: 거래대금 및 수급 실측
- **Action**: 특정 창구(키움, 미래 등)의 대량 주문 유입 확인.
- **Logic**: 대형 창구의 순매수 전환 시 추세 강화 가능성 높음.

### 3단계: 기술적 타점 선정
- **Action**: 3분봉/15분봉 기준 눌림목 및 돌파 타점 계산.
- **Decision**: 매수(Buy) / 관망(Wait) / 손절(Cut) 여부 결정.

---

## 📄 리포트 템플릿 (Report Template)
1. **실시간 현황**: 현재가, 등락률, 체결강도
2. **차트 분석**: 주요 분봉 이평선 및 지지/저항선
3. **창구 분석**: 주도 매수 창구 및 외국인 가집계
4. **매매 계획**: 진입가, 목표가, 손절가
5. **결론**: 실행(Execution) / 대기(Standby)
6. **참고 자료 (Reference)**: [필수] 실시간 호가 및 분봉 데이터 요약

---

**Next Action**: 매매 종료 후 `integrated-daily-routine`의 복기 단계로 이동.
