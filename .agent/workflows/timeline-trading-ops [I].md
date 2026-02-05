---
description: 타임라인에 따른 최적화된 워크플로우 실행 및 데일리 오퍼레이션 통합 가이드
---

# 🕒 timeline-trading-ops [I] (시점별 트레이딩 오퍼레이션)

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Time-Sensitive**: 시장의 각 시점(프리/정규/애프터)에 최적화된 워크플로우를 가동합니다.
- **Context-First**: 현재 시각의 시장 상태를 먼저 확인하고 분석 모드를 결정합니다.
- **Code-Accuracy**: 모든 분석 시작 전 `KIS find_stock_code`와 `Naver search_stock`을 병행하여 정확한 6자리 종목 코드를 확정합니다. KIS에서 검색되지 않을 경우 네이버 금융 검색 결과를 우선합니다.

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 워크플로우는 **`GEMINI.md`**에 정의된 **'공통 보고서 프로토콜'** 및 **'핵심 원칙'**을 엄격히 준수합니다.
- **WorkflowID**: `TTO` 식별자를 사용합니다.

## 🎯 워크플로우 특화 규칙 (Workflow Rules)
1. **시간 엄수**: 현재 마켓 타임 테이블에 따른 전용 분석 도구를 우선 호출합니다.
2. **데이터 무결성**: 분석 시작 전 `NQ=F`, `ES=F` 실시간 등락률을 최우선으로 확보합니다.
3. **종목 코드 검증**: `KIS`와 `Naver`를 병행하여 6자리 코드를 확정합니다.
## 🛠️ 시점별 분석 단계 (Time-based Process)

### 1단계: 장전 브리핑 (08:00 ~ 09:00) - [모드: M]
-   **핵심 스킬**: `macro-event-scanner [M]` → `market-sector-scanner [M]`
-   **Action**: 
  - 밤사이 미 증시 결과 및 나스닥/S&P 선물 등락률 실측.
  - **글로벌 뉴스 체인 점검**: Yahoo Finance 및 주요 외신 헤드라인을 통해 글로벌 센티먼트 파악.
  - 전일 시간외 수급 및 당일 테마 형성 여부 확인.
  - 프리마켓(NXT) 체결가를 통한 시초가 강도 예측.

### 2단계: 장중 실전 매매 (09:00 ~ 15:30) - [모드: T]
-   **핵심 스킬**: `day-trading [T]` or `swing-trading [T]`
-   **Action**:
  - 분봉 지지/저항선 돌파 여부 실시간 모니터링.
  - `investor-matrix [I]`(Expert Skill)를 통한 메이저 창구(기관/외인) 진입 가격 추적.
  - VI 발생 시 비정상적 과열/투매 리스크 확인.

### 3단계: 장후 수사 및 야간 대응 (15:30 ~ 20:00) - [모드: T/I]
-   **핵심 스킬**: `after-market-trading [T]` → `risk-strategist [I]`
-   **Action**:
  - 장 마감 후 발생한 DART 공시 및 뉴스 시의성 검증.
  - NXT 애프터마켓 수급 강도 실측 및 오버나이트 비중 조절.
  - `risk-strategist [I]`(Expert Skill)를 통한 익일 리스크 사전 예고.

### 4단계: 데일리 통합 전략 수립 (20:00 ~ 심야) - [모드: I]
-   **Sequence**: `integrated-daily-routine [I]` (팩트 생산) ➡️ `multi-agent-decision [I]` (전략 결정)
-   **Action**:
  - `IDR [I]`을 실행하여 6대 전문가의 팩트 리포트 세트 완성.
  - 생성된 팩트 리포트를 기반으로 `MAD [I]`을 호출하여 에이전트 간 토론 및 최종 전략 수립.
  - 마침내 전문가용 상세 전략과 투자자용 **'실전 쉬운 가이드'** 패키징 완성.

---

## 저장 및 네이밍
`GEMINI.md` 공통 프로토콜 준수 (ID: `TTO`)

---
**Standard Operating Procedure**: `Antigravity Timeline Ops`
