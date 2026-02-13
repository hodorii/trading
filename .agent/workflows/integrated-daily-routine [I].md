---
description: 매크로, 이벤트, 섹터 분석을 통해 유망 종목을 추출하고, 개별 종목의 전략적 근거를 집대성하는 팩트 중심 마스터 워크플로우
---

# 🔄 integrated-daily-routine [I] (MoE 전문가 믹스 루틴)

## 🎯 목적
과거 결정의 복기(Review)부터 현재 데이터의 정밀 실측, 그리고 전략적 합성(Synthesis)까지 데이터의 유실 없는 분석을 통해 최종 의사결정을 지원합니다.

## 🚨 전문가 공통 준수 사항 (Global Protocol)
- **실시간 시각 동기화 (ETO Rule)**: 모든 보고서의 발행 시각은 반드시 `date` 명령어의 실행 결과 또는 시스템 메타데이터(`ADDITIONAL_METADATA`)의 시각을 그대로 사용합니다. **예상 완료 시각이나 워크플로우 권장 시각을 임의로 기입하는 행위를 엄격히 금지**하며, 데이터의 정직성을 최우선으로 합니다.
- **복기 우선 (Review First)**: 모든 분석은 어제의 예측과 오늘의 실측 데이터 대조에서 시작합니다.
- **교차 검증 (Cross-Check)**: KIS API 데이터를 기본으로 하되, Naver/Yahoo 데이터를 통해 수급 및 가격의 괴리율을 실측하고 이상 발생 시 `[리스크]` 보고서에 즉시 반영합니다.
- **시각화 의무 (Visual Matrix)**: 모든 분석 보고서 상단에는 현 상태를 한눈에 파악할 수 있는 **'Quick Scan Matrix (🔴/🟡/🟢)'** 테이블을 반드시 포함합니다.
- **근거의 충실함**: 수치(Price, Vol, Ratio)와 시각(Intraday Time), 그리고 추론 과정(CoT)을 상세히 기록합니다.

---

## 🛠️ 전문가 믹스 파이프라인 (9-Step Expert Pipeline)

### [Phase 0: Accountability - 책임 및 복기]

### 0단계: 복기 전문가 (Review Expert)
-   **Task**: "어제의 [결정] vs 오늘의 [결과]" 대조 분석.
-   **Output**: 어제 전략의 유효성 점검 및 오늘 분석의 보정 가중치 산출.

### [Phase 1: Market Intelligence - 시장 정보 분석]

### 1단계: 매크로 전문가 (Macro Expert)
-   **Output Tag**: `[매크로]`
-   **Task**: 글로벌 지수 및 미국 마켓 임팩트 팩터 민감도 분석.

### 2단계: 이벤트 전문가 (Event Expert)
-   **Output Tag**: `[이벤트]`
-   **Task**: 뉴스/공시 팩트 체크 및 국내 섹터 전이 경로(Spillover) 분석.

### 3단계: 섹터 전문가 (Sector Expert)
-   **Output Tag**: `[시장]`
-   **Task**: 미국 감응 섹터(US-Linked) 및 시장 주도 테마 발굴.

### [Phase 2: Deep Dive - 종목 심층 분석]

### 4단계: 가치 전문가 (Value Expert)
-   **Output Tag**: `[가치]`
-   **Task**: 재무 적정성 및 밸류에이션 추이 분석 (최근 3개년 대조).

### 5단계: 수급 전문가 (Flow Expert)
-   **Output Tag**: `[수급]`
-   **Task**: 메이저 주체 매매 패턴 및 **'검은머리 외국인'** 판별 (창구 유입 시각 실측).

### 6단계: 리스크 전문가 (Risk Expert)
-   **Output Tag**: `[리스크]`
-   **Task**: **De-coupling Risk** 및 데이터 괴리율(KIS vs Naver) 진단.

### [Phase 3: Tactical Synth - 전략적 합성]

### 7단계: 전략 합성 전문가 (Tactical Synthesis)
-   **Output Tag**: **`[합성]`**
-   **Task**: 0~6단계의 모든 핵심 근거(Key Evidences)를 집대성한 **'최종 변론서'** 작성. (함축 금지)

### 8단계: 세션 요약 리포터 (Session Summary Reporter)
-   **Output Tag**: `[요약]`
-   **Task**: 전체 세션 인덱스 및 핵심 Why-First 요약.

---
**Workflow ID**: `IDR`
