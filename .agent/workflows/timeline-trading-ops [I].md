---
description: 현재 시각을 기준으로 장전 장중 장후 심야 단계에 최적화된 워크플로우를 호출하여 공백 없는 데일리 오퍼레이션을 수행하는 마스터 워크플로우
---

# 🕒 timeline-trading-ops [I] (시점별 트레이딩 오퍼레이션)

## 🎯 목적
현재 시각을 기준으로 트레이딩의 각 단계(장전/장중/장후/심야)에 최적화된 마스터 워크플로우를 호출하여 공백 없는 데일리 오퍼레이션을 수행합니다.

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 워크플로우는 **`GEMINI.md`** 및 **`integrated-daily-routine [I]`**의 세션 관리 규칙을 엄격히 준수합니다.

---

## 🛠️ 시점별 오퍼레이션 시퀀스 (Operation Sequence)

### 1단계: 장전 브리핑 (08:00 ~ 09:00)
- **Mode**: 시장 탐색 및 등락 예측.
- **Action Chain**: `macro-event-scanner [M]` ➡️ `market-sector-scanner [M]`.
- **Task**: 미 증시 종가 확인, 현물/선물 괴리 분석, NXT 프리마켓 체결가 기반 시초가 테마 포착.

### 2단계: 장중 실전 대응 (09:00 ~ 15:30)
- **Mode**: 실시간 추세 추종 및 수급 실측.
- **Action Chain**: `day-trading [T]` or `swing-trading [T]`.
- **Task**: 분봉 지지/저항 실시간 체크, `investor-matrix [I]`를 통한 메이저 창구 평단가 추적.

### 3단계: 장후 수사 및 야간 브리핑 (15:30 ~ 20:00)
- **Mode**: 공시 분석 및 익일 리스크 사전 진단.
- **Action Chain**: `after-market-trading [T]` ➡️ `risk-strategist [I]`.
- **Task**: 
    1. 당일 공시/뉴스 시의성 검증 및 **주요 일정(배당/분할 예정일 등) 추출**.
    2. **데이터 소스 검증**: `inquire_overtime_price` 결과값이 장중 종가와 단순히 동일한지 확인하고, 실제 체결량(`vol`)이 발생한 데이터인지 반드시 실측.
    3. NXT 애프터마켓 수급 확인 및 익일 예상 리스크 리포팅.

### 4단계: 데일리 마스터 루틴 (20:00 ~ 심야)
- **Mode**: 데이터 집대성 및 최종 투자의견 결정.
- **Action Chain**: `integrated-daily-routine [I]` (팩트 생산) ➡️ `multi-agent-decision [I]` (전략 결정).
- **Task**:
    1. **Fact Check (필수)**: `mcp_kis-trading-remote_domestic_stock` API를 사용하여 정규장 종가(`inquire_price`)와 시간외 단일가(`inquire_overtime_price`)를 직접 실측(Direct Verification)합니다.
        > *단순 종가 확인이 아니라, 거래량(`vol`)과 매수/매도 잔량을 확인하여 허수 여부를 판별해야 합니다.*
    2. **Cross-Check**: `mcp_naver-finance` 도구를 활용하여 기관/외국인 수급 추이(`get_trading_trends`)와 컨센서스 괴리율을 교차 검증합니다.
    3. **De-coupling Check**: 매크로(미국장) 호재와 개별 종목 악재(소송, 상장 철회 등) 충돌 시, **개별 악재 리스크(Idiosyncratic Risk)**를 최우선으로 반영합니다.
    4. `IDR [I]` 실행: 위 검증된 팩트를 기반으로 전문가 리포트 8종 세트를 완성합니다.
    5. `MAD [I]` 토론: 검증된 팩트를 토대로 최종 [통합전략] 및 [실전가이드] 도출.

---

## 📝 최종 점검 (Final Check)
- [ ] 현재 시각에 맞는 워크플로우 시퀀스가 선택되었는가?
- [ ] 모든 분석 전 KIS API 실측 데이터(정규장/시간외)가 확보되었는가? (심야 필수)
- [ ] **[심야 모드]** 실행 시 IDR과 MAD가 하나의 세션 ID를 공유하며 연속 실행되었는가?

---
**Standard Operating Procedure**: `Antigravity Timeline Ops`
**Workflow ID**: `TTO`
