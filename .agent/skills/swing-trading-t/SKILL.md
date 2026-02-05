---
name: swing-trading [T]
description: 1~2주 보유하며 추세의 무릎에서 어깨까지 수익을 추구하는 스윙 매매 워크플로우
---

## 🚨 글로벌 규약 (Global Rules)
1. **한국어 우선 (Korean-First)**: 분석 내용은 한국어로 기술합니다.
2. **참고 자료 의무화 (Mandatory Reference)**: 주봉/일봉 차트 및 기관/외인 누적 수급표를 첨부합니다.
3. **폴더 규칙 준수**: `d:\dev\trading\reports\YYYY-MM-DD\` 경로에 저장합니다.

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Trend Following**: 이미 형성된 우상향 추세에 올라타는 것을 기본으로 합니다.
- **Cycle Analysis**: 업황이나 섹터의 순환 주기를 고려합니다.

## 📋 파일명 네이밍 컨벤션
### 형식
`YYYYMMDD_HHmm_[태그]_종목명_설명.md`
(예: `20260127_1100_[스윙]_우리기술_스윙운용전략.md`)

# 🌊 swing-trading [T] (스윙 트레이딩)

## 🎯 분석 목적
- 안정적인 우상향 추세 종목의 수익 극대화.
- 시장 노이즈를 배제한 본질적 추세 추종.

---

## 🛠️ 활용 도구 및 API
- **KIS API**: `inquire_daily_itemchartprice` (일봉/주봉), `inquire_investor`
- **Naver**: `get_trading_trends`, `get_stock_news`

---

## 📈 분석 프로세스 (Process)

### 1단계: 주봉 기반 대추세 확인
- **Action**: 역배열에서 정배열로의 전환 여부 확인.
- **Check**: 중장기 이평선(60일, 120일) 지지 유무.

### 2단계: 기관/외인 누적 매집량 분석
- **Action**: 최근 1개월~3개월간의 메이저 주체 누적 순매수 확인.
- **Logic**: 박스권 내에서의 누적 매집은 돌파의 전조임.

### 3단계: 피드백 및 목표가 보정
- **Action**: 시장 상황(지수)에 따른 목표가 상향 또는 하향 조정.

---

## 📄 리포트 템플릿 (Report Template)
1. **추세 분석**: 주봉/일봉 이평선 배열, 추세선 작도
2. **누적 수급**: 기관/외인 기간별 매집 현황
3. **운영 계획**: 진입 타이밍, 비중 조절 계획
4. **목표/손절**: 중기 목표가 및 추세 붕괴 손절가
5. **결론**: 비중 확대 / 보유 / 매도
6. **참고 자료 (Reference)**: [필수] 차트 스크린샷 및 수급 링크

---

**Next Action**: 분석 결과가 `fundamental-analysis`와 일치하는지 최종 대조.
