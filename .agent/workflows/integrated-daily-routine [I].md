---
description: 매크로, 이벤트, 섹터 분석을 통해 유망 종목(3-10선)을 추출하고, 개별 종목의 전략적 트레이딩(데이/스윙/단기)을 설계하는 마스터 워크플로우
---

# 🔄 integrated-daily-routine [I] (MoE 전문가 믹스 루틴)

## 🎯 목적
매크로(거시)부터 리스크(위험)까지 8개의 특화된 **전문가(Expert)** 모델을 순차적으로 로딩하여, 데이터 오염 없이 최적의 의사결정을 도출하고 개별 보고서를 생성하는 MoE 기반 마스터 워크플로우입니다.

## 🚨 전문가 공통 준수 사항 (Global Protocol)
본 워크플로우는 **`GEMINI.md`**에 정의된 **'공통 보고서 프로토콜'** 및 **'파일 네이밍 컨벤션'**을 엄격히 준수합니다.

## 🎯 전문가별 SRP 및 리포팅 원칙 (Strict Rules)
1. **개별 보고서 생성**: 각 전문가(Step)는 반드시 자신의 태그가 붙은 **독립적인 리서치 보고서**를 생성해야 합니다.
2. **데이터 체인**: 모든 분석 시작 전 [미 현물 종가 vs 실시간 선물] 데이터를 대조 로딩하여 전역 컨텍스트를 동기화하며, ADDITIONAL_METADATA의 현재 시각을 기준으로 본문 내 발행 시각을 정합성 있게 기록합니다. (미래 시점 기입 금지)

---

## 🛠️ 전문가 믹스 파이프라인 (8-Step Expert Pipeline)

### [Phase 1: Market Intelligence - 시장 정보 분석]

### 1단계: 매크로 전문가 (Macro Expert)
-   **Expert Skill**: `macro-event-scanner [M]`
-   **Output Tag**: `[매크로]`
-   **Task**: 글로벌 지수 및 **미국 마켓 임팩트 팩터(금리, DXY, VIX) 민감도 분석**을 통해 대내외 리스크와 기회 요인 도출.

### 2단계: 이벤트 전문가 (Event Expert)
-   **Expert Skill**: `event-driven [M]`
-   **Output Tag**: `[이벤트]`
-   **Task**: 뉴스/공시 팩트 체크(D-Day 및 예정일 명시) 및 **미국발 이벤트(SaaS 위기/AI 인프라)의 국내 섹터 전이 경로(Spillover)** 심층 분석.

### 3단계: 섹터 전문가 (Sector Expert)
-   **Expert Skill**: `market-sector-scanner [M]`
-   **Output Tag**: `[시장]`
-   **Task**: 매크로 분석 결과를 수용하여 **미국 감응 섹터(US-Linked)** 및 시장 주도 테마를 발굴하고 **Top 3~5 추천 종목** 선정.

### [Phase 2: Deep Dive - 종목 심층 분석]
*(사용자가 지정한 종목이 없을 경우, **Phase 1에서 선정된 Top Pick 종목**을 대상으로 **'Auto-Discovery Mode'**로 자동 전환하여 수행합니다)*

### 4단계: 가치 전문가 (Value Expert)
-   **Expert Skill**: `fundamental-analysis [I]`
-   **Output Tag**: `[가치]`
-   **Task**: 재무 적정성, 밸류에이션(PER/PBR) 상단/하단 도출. 
    > *다수 종목(3개 이상) 분석 시, 개별 리포트 대신 **'가치 비교 매트릭스(Value Matrix)' 통합 리포트** 생성을 허용합니다.*

### 5단계: 수급 전문가 (Flow Expert)
-   **Expert Skill**: `investor-matrix [I]`
-   **Output Tag**: `[수급]`
-   **Task**: 메이저 수급 주체 매매 패턴 및 **'검은머리 외국인'** 판별.
    > *다수 종목 분석 시, **'수급 주체별 매트리스(Flow Matrix)' 통합 리포트** 생성을 허용합니다.*

### 6단계: 리스크 전문가 (Risk Expert)
-   **Expert Skill**: `risk-strategist [I]`
-   **Output Tag**: `[리스크]`
-   **Task**: 공매도/대차거래 추이 및 **De-coupling Risk(개별 악재 vs 매크로 호재)** 진단.
    > *다수 종목 분석 시, **'위험 요인 비교(Risk Comparison)' 통합 리포트** 생성을 허용합니다.*

### [Phase 3: Tactical Synth - 전략적 합성]

### 7단계: 통합 전략가 (Final Strategist)
-   **Expert Skill**: `total-trading-strategy [I]`
-   **Output Tag**: `[결정]`
-   **Task**: 앞선 팩트 보고서를 통합하여 **[당일/단기/중기]** 매매 성격 판별 및 타점(진입/목표/손절) 설계.

### 8단계: 세션 요약 리포터 (Session Summary Reporter)
-   **Expert Task**: 전체 세션의 8단계 분석을 한눈에 파악할 수 있도록 핵심만 추출한 요약 보고서 생성.
-   **Output Tag**: `[요약]`
-   **Task**: 
    -   금일 변동 핵심 사유를 3가지로 압축.
    -   생성된 7개 전문가 보고서를 Phase별로 그룹화하여 괄호 안에 한 줄 요약 포함.
    -   최종 전략 제언을 진입가/목표가/손절가와 함께 명확히 제시.

---

## 📝 최종 요약 양식 (Summary Template)
워크플로우 완료 후 제공되는 요약입니다.

### [요약]
#### 📌 금일 변동 핵심 사유 (Why-First)
- 요약 내용 1
- 요약 내용 2

#### 📂 생성된 분석 보고서 (reports/YYYY-MM-DD/)
- **[매크로/이벤트/시장/가치/수급/리스크/결정]** 리스트 나열 (각 보고서마다 괄호 안에 핵심 내용 1줄 요약)

#### ⚖️ 최종 전략 제언
- 핵심 결론 및 대응 방향.

---

## 📝 최종 점검 (Final Check)
- [ ] 8단계 전문가의 SRP가 정확히 이행되어 각 파일이 생성되었는가?
- [ ] 모든 보고서 최상단에 'Why-First'가 명시되었는가?
- [ ] 외국인 창구 분석 시 '단타성' 여부 판별이 포함되었는가?
