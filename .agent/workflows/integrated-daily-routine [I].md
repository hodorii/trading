---
description: 매크로, 이벤트, 섹터 분석을 통해 유망 종목을 추출하고, 개별 종목의 전략적 근거를 집대성하는 팩트 중심 마스터 워크플로우
---

# 🔄 integrated-daily-routine [I] (MoE 전문가 믹스 루틴)

## 🎯 목적
매크로(거시)부터 리스크(위험)까지 8개의 특화된 **전문가(Expert)** 모델을 통해 데이터를 정밀 실측하고, 각 단계의 **수치적 근거와 추론 과정(Chain of Thought)**을 누락 없이 기록하여 전략 수립을 위한 최종 변론서를 완성합니다.

## 🚨 전문가 공통 준수 사항 (Global Protocol)
- **근거의 충실함**: "좋다/나쁘다"의 결론보다 **"왜(Why)"**에 집중하며, 수치(Price, Vol, Ratio)와 시각(Intraday Time)을 반드시 포함합니다.
- **파일 네이밍**: **`GEMINI.md`**의 컨벤션을 준수합니다.

---

## 🛠️ 전문가 믹스 파이프라인 (8-Step Expert Pipeline)

### [Phase 1: Market Intelligence - 시장 정보 분석]

### 1단계: 매크로 전문가 (Macro Expert)
-   **Expert Skill**: `macro-event-scanner [M]`
-   **Output Tag**: `[매크로]`
-   **Task**: 글로벌 지수 및 **미국 마켓 임팩트 팩터(금리, DXY, VIX) 민감도 분석**. 
    > *단순 지수 나열이 아닌, 전일 대비 변동폭과 그에 따른 한국 시장의 시초가 영향력을 수치로 제시.*

### 2단계: 이벤트 전문가 (Event Expert)
-   **Expert Skill**: `event-driven [M]`
-   **Output Tag**: `[이벤트]`
-   **Task**: 뉴스/공시 팩트 체크 및 **국내 섹터 전이 경로(Spillover)** 심층 분석. 
    > *이벤트의 '시의성'과 '파급력'을 상/중/하로 구분하고 근거 뉴스 링크 또는 공시 번호 병기.*

### 3단계: 섹터 전문가 (Sector Expert)
-   **Expert Skill**: `market-sector-scanner [M]`
-   **Output Tag**: `[시장]`
-   **Task**: **미국 감응 섹터(US-Linked)** 및 시장 주도 테마 발굴.
    > *섹터 내 대장주와 부대장주의 거래대금 비중 및 상승 강도를 비교 분석하여 기록.*

### [Phase 2: Deep Dive - 종목 심층 분석]

### 4단계: 가치 전문가 (Value Expert)
-   **Expert Skill**: `fundamental-analysis [I]`
-   **Output Tag**: `[가치]`
-   **Task**: 재무 적정성 및 **밸류에이션(PER/PBR) 추이 분석**.
    > *최근 3개년 평균 대비 현재 위치를 수치화하고, 적정 주가 산출의 산식(Formula)이나 논리를 상세히 기록.*

### 5단계: 수급 전문가 (Flow Expert)
-   **Expert Skill**: `investor-matrix [I]`
-   **Output Tag**: `[수급]`
-   **Task**: 메이저 주체 매매 패턴 및 **'검은머리 외국인'** 판별.
    > *특정 창구의 집중 매수/매도 시각(Time-stamp)과 프로그램 매매와의 동조율을 실측 데이터로 기록.*

### 6단계: 리스크 전문가 (Risk Expert)
-   **Expert Skill**: `risk-strategist [I]`
-   **Output Tag**: `[리스크]`
-   **Task**: **De-coupling Risk(개별 악재 vs 매크로 호재)** 진단.
    > *신용 잔고, 공매도 대차잔고의 실시간 변동 수치를 근거로 하방 압력 강도를 상세히 서술.*

### [Phase 3: Tactical Synth - 전략적 합성]

### 7단계: 전략 합성 전문가 (Tactical Synthesis)
-   **Expert Task**: 1~6단계의 모든 핵심 근거(Key Evidences)를 집대성한 **'최종 변론서'** 작성.
-   **Output Tag**: `[합성]`
-   **Task**:
    -   **근거 함축 금지**: 각 단계의 핵심 수치와 논리를 MAD가 판단할 수 있도록 상세히 나열.
    -   **모순점 도출**: 가치와 수급이 충돌하거나, 이벤트와 리스크가 상충하는 지점을 명확히 짚어냄.

### 8단계: 세션 요약 리포터 (Session Summary Reporter)
-   **Output Tag**: `[요약]`
-   **Task**: 전체 세션을 한눈에 파악할 수 있는 인덱스 역할.

---
**Workflow ID**: `IDR`
