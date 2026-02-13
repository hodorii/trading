---
description: 장전-장중-장후-심야로 이어지는 데이터 내러티브를 관리하고 단계별 최적 워크플로우를 연결하는 마스터 오퍼레이션
---

# 🕒 timeline-trading-ops [I] (시점별 트레이딩 오퍼레이션)

## 🎯 목적
현재 시각을 기준으로 데이터의 연속성(Context Hand-over)을 유지하며 최적의 워크플로우를 실행합니다. **"오늘의 가설(장전) ➡️ 실측(장중) ➡️ 검증(장후) ➡️ 결론(IDR/MAD)"**의 완결된 내러티브를 구축합니다.

---

## 🛠️ 시점별 오퍼레이션 시퀀스 (Operation Sequence)

### 1단계: 장전 브리핑 (08:00 ~ 09:00) - [가설 수립]
- **Action**: `macro-event-scanner [M]` ➡️ `market-sector-scanner [M]`.
- **Context**: 전일 미 증시 기반 금일 시초가 주도 테마 및 등락 범위 가설 수립.

### 2단계: 장중 실전 대응 (09:00 ~ 15:30) - [데이터 실측]
- **Action**: `day-trading [T]` or `swing-trading [T]`.
- **Context**: 장전 가설과 실제 수급 대조. 주도 창구의 실시간 평단가(VWAP) 추적.

### 3단계: 장후 수사 (15:30 ~ 20:00) - [현상 검증]
- **Action**: `after-market-trading [T]` ➡️ `risk-strategist [I]`.
- **Context**: 
    - 당일 공시의 시의성 및 시간외 단일가 허수 여부 판별.
    - 장중 발생한 변동성의 '진위'를 수급 데이터로 최종 검증하여 IDR로 전달.

### 4단계: 데일리 마스터 루틴 (20:00 ~ 심야) - [최종 결론]
- **Action**: `IDR [I]` (근거 집대성) ➡️ `MAD [I]` (판단 및 확약).
- **Context**: 
    - **Context Hand-over**: 1~3단계에서 실측된 모든 팩트를 IDR 전문가들에게 주입.
    - **Final Commitment**: 수집된 근거 사이의 모순을 해결하고 익일 실행 전략 확정.

---
**Standard Operating Procedure**: `Antigravity Timeline Ops`
**Workflow ID**: `TTO`
