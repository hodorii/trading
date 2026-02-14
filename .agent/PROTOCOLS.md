# 📜 Trading Agent Global Protocols (v1.2)

본 저장소의 모든 에이전트는 아래 규정을 강제 준수한다.

### 1. [GP-ETO] 시각 관리 및 KST 준수 (Time Management & KST)
- **기준 시각**: 모든 시각은 **한국 표준시(KST, UTC+9)**를 기준으로 하며, 파일명과 리포트 내부에 이를 반영한다.
- **파일명 (GP-FNAME)**: `YYYY-MM-DD-HHmm-SEQ-Title.md` 형식을 엄수한다.
- **세션 ID (GP-SESS)**: 동일 주제나 15분 이내의 연속된 분석은 동일한 `HHmm`을 공유하여 그룹화한다.
- **본문 시각**: 리포트 내부의 시각은 실제 실측 시각을 KST로 기록한다.

### 2. [GP-SEQ] 워크플로우 순번 규정 (Workflow Sequencing)
보고서 생성 시 아래 순번에 따라 `SEQ`를 부여하여 아카이브 정렬 순서를 보장한다.
- **01 (Review)**: 전일 분석 복기 및 성과 실측.
- **02 (Macro)**: 매크로 지수 및 환율 분석.
- **03 (Event)**: 주요 뉴스 및 공시 팩트체크.
- **04 (Market)**: 섹터 주도권 및 ETF 수급 동향.
- **05 (Value)**: 밸류에이션 및 실적 분석.
- **06 (Supply)**: 외인/기관 수급 포렌식.
- **07 (Risk)**: 단기 과열 및 데이터 괴리 진단.
- **08 (Synthesis)**: 전략 집대성 및 추천 종목 확정.
- **09 (Summary)**: 3줄 핵심 요약 가이드.
- **10 (TSM)**: Trading Score Matrix 수치화.
- **11-13 (MAD/Guide/Unified)**: 의사결정 및 실전 가이드.

### 3. [GP-TINT] 제목 무결성 및 명명 규칙 (Title Integrity)
- **파일명/제목**: 대괄호 `[` `]`를 제거하고, 공백 대신 언더바 `_`를 구분자로 사용한다.
- **특수문자**: 파일 시스템 호환성을 위해 최소한의 특수문자만 사용한다.

### 4. [GP-META] 메타데이터 및 구조 유지 (Metadata & Structure)
- **YAML Front Matter**: 모든 리포트 상단에 `session_id`, `session_order`, `title`, `layout: post`를 포함한다.
- **로컬 아카이브**: `reports/YYYY-MM-DD/` 폴더 구조를 유지하며 원본 파일을 보관한다.
- **배포 아카이브**: 외부 저장소(`trading-pulse`)의 `_posts/` 폴더에 배포용 파일을 동기화한다.

### 5. [GP-SPEC] KIS API 명세 우선 (Spec First)
- 호출 전 `find_api_detail` 필수 수행 및 `env_dv`('real'/'demo') 확인.
- 추측 금지. 사용 API명 및 필수 `fid_` 인자 리포트 하단 적시.

### 6. [GP-EBR] 근거 중심 리포팅 (Evidence-Based)
- 수치(Price/Vol/Ratio) + 시각(Time) + 추론(CoT) 필수.
- 모든 보고서 상단에 **Quick Scan Matrix (🔴/🟡/🟢)** 포함.

### 7. [GP-REV] 복기 우선 (Review First)
- 분석 시작 전 "어제 가설 vs 오늘 결과" 대조 및 가중치 보정.

### 8. [GP-TSM] 수치화 평가 규정 (Trading Score Matrix)
1. **[EQS] 실적 품질 점수**: YoY 영업익 변화율 기준.
2. **[SMR] 수급 모멘텀 비율**: 외인/기관 수급 흡수율 기준.
3. **[VPR] 가치 프리미엄 비율**: 3년 평균 PBR 대비 위치 기준.
- **종합 점수(Total Score)**: 합산 점수에 따른 비중 조절 가이드 강제.

### 9. [GP-GIT] Git 반영 정책 (Git Persistence)
- **리포트 발행**: 로컬 `reports/`와 외부 `trading-pulse`를 동시에 업데이트하며, 세션 종료 시 일괄 커밋 및 푸시한다.
- **KST 타임스탬프**: 파일 수정 시간(`mtime`)은 가급적 생성 시각과 일치시킨다.
