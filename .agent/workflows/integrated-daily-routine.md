# integrated-daily-routine (Core 8-Step v1.5)

## 목표
TTO(시간대별 운영), 정밀 포렌식, 차트 분석을 결합하여 개별 종목 매매 전략 혹은 **시장 주도 10개 종목 리스트**를 수립함.

## Pipeline (Strict 8-SEQ)

### [Phase 1: Intelligence & Discovery]
- **01. 매크로**: 지수 이격도 및 유동성 에너지 실측.
- **02. 이벤트**: 현재 TTO 시점의 뉴스 영향력 분석.
- **03. 시장/기술**: 섹터 주도권 판별 및 **(종목 미지정 시) 유망 10개 종목 발굴**.

### [Phase 2: Deep Dive]
- **04. 가치**: [GP-ERN] 실적 기반 밸류에이션 리레이팅 여부 판별.
- **05. 수급**: [GP-FORENSIC] 창구별 손바뀜 및 검은머리 외인 추적.
- **06. 리스크**: 기술적 과열 및 규제 공시 리스크 필터링.

### [Phase 3: Action]
- **07. 합성**: 근거 집대성 및 **트레이딩 관점(TP)** 확정.
- **08. 결정**: 최종 MAD 의사결정 및 매매 지침 발행.

### [Phase 4: Finalize]
- **Git Persistence**: [GP-GIT]에 따라 로컬 및 배포 저장소 동기화 후 일괄 커밋.

---
## 아카이빙 및 동기화
- **저장 경로 1**: `reports/YYYY-MM-DD/` (로컬 보관)
- **저장 경로 2**: `../trading-pulse/_posts/YYYY-MM-DD/` (배포 보관)
- **파일명 규격**: `YYYY-MM-DD-HHmm-SEQ-한글_제목.md` 엄수.
- **검증**: 배포 전 순번(SEQ) 및 KST 시각 기록 여부 확인.
