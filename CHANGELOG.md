# CHANGELOG.md (antigravity 작업 기록)

본 파일은 `antigravity` 에이전트가 `d:\dev\trading` 워크스페이스 내 `.agent` 및 `.gemini` 디렉터리에 가한 주요 변경 사항과 전체 작업 내역을 기록합니다.

## [2026-02-05]
### 추가
- **통합 데일리 루틴 분석 수행**: `integrated-daily-routine [I]` 워크플로우를 활용한 종목 분석 수행.
- **Git 저장소 관리**: 워크스페이스 Git 초기화 및 `.gitignore`, `CHANGELOG.md` 생성. (.agent, .gemini 포함)

## [이전 주요 변경 사항] - Antigravity 내역
### `.agent/` (지능형 워크플로우 및 전문가 스킬)
- **MoE(Mixture of Experts) 아키텍처 구축**:
  - `integrated-daily-routine [I]` 워크플로우 설계: 매크로, 이벤트, 섹터, 가치, 수급, 리스크를 순차적으로 분석하는 8단계 파이프라인 구축.
  - 전문가 모델(Skills) SRP 적용: 각 전문가(Macro, Fundamental, Investor, Risk 등)가 단일 책임 원칙을 따르도록 스킬 파일 고도화.
- **네이밍 컨벤션 도입**: `YYYYMMDD_HHmm_WorkflowID_[태그]_종목명` 형식을 도입하여 분석 자산의 파편화 방지 및 검색성 향상.
- **데이터 체인 로딩**: [국내 주가 <-> 미 지수 선물 <-> 환율/금리]를 동시 분석하는 입체적 데이터 로직 반영.

### `.gemini/` (에이전트 페르소나 및 원칙)
- **최상위 절대 원칙 수립 (`GEMINI.md`)**:
  - **한국어 우선(Korean-First)**: 모든 사고와 아웃풋을 한국어로 수행하는 강력한 커뮤니케이션 규약 적용.
  - **사실 기반(Fact-Based)**: 주관적 판단 배제 및 로우 데이터(DART, 야후파이낸스 등) 중심의 분석 원칙 확립.
- **연결 최적화 (`mcp_config.json`)**: KIS Trading, Naver Finance 등 외부 MCP 도구와의 원활한 연결 설정.

---
*기록자: antigravity (Agentic AI)*
