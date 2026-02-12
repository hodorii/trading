# 📋 Spec-Driven Development (SDD) 프롬프트

## 🎯 목적
- 기능 아이디어 → 체계적 설계 문서 + 구현 계획 변환
- Property-Based Testing(PBT) 기반 정확성 검증
- 고품질 소프트웨어 개발 방법론

## 🚨 핵심 원칙
- **사용자 진실 확립**: 각 단계별 사용자 승인 필수
- **반복적 개선**: 요구사항 ↔ 연구 자유 이동
- **정형 명세**: 실행 가능한 정확성 속성 검증
- **단계별 실행**: 1작업 집중 → 완료 → 사용자 검토

---

## 📁 폴더 구조 (Folder Structure)

### 기본 디렉토리 구조
```
.kiro/
└── specs/
    └── {feature-name}/          # kebab-case 형식 (예: user-authentication)
        ├── requirements.md      # 요구사항 문서
        ├── process.md          # 업무 프로세스 설계 (BPMN 2.0 기반)
        ├── design.md           # 기술 설계 문서
        ├── ui-design.md        # UI/UX 설계 문서
        └── tasks.md            # 구현 작업 목록
```

### 기능명 명명 규칙
- **형식**: kebab-case (소문자, 하이픈으로 구분)
- **예시**: `user-authentication`, `payment-processing`, `data-analytics`
- **원칙**: 간결하고 기능을 명확히 표현

---

## 🛠️ 워크플로우 단계 (Workflow Steps)

### Phase 1: 요구사항 수집 (Requirements Gathering)

#### 1단계: 기능명 결정
```
기능 아이디어 → kebab-case 기능명 생성
예: "사용자 로그인 시스템" → "user-authentication"
```

#### 2단계: requirements.md 생성
```markdown
# {기능명} 요구사항

## 개요
- 목적: [핵심 목적 1줄]
- 범위: [포함/제외 범위]
- 대상 사용자: [주요 사용자 그룹]

## 사용자 스토리
| ID | 사용자 | 기능 | 목적 | 우선순위 |
|----|--------|------|------|----------|
| US1 | [사용자 유형] | [원하는 기능] | [달성 목적] | High/Medium/Low |
| US2 | [사용자 유형] | [원하는 기능] | [달성 목적] | High/Medium/Low |

## 수용 기준
| 스토리 ID | 기준 | 검증 방법 |
|-----------|------|-----------|
| US1 | [구체적 조건] | [테스트 방법] |
| US2 | [구체적 조건] | [테스트 방법] |

## 제약사항
| 유형 | 제약사항 | 영향도 |
|------|----------|--------|
| 기술적 | [기술 제약] | [영향 범위] |
| 비즈니스 | [비즈니스 제약] | [영향 범위] |
| 법적/규제 | [규제 제약] | [영향 범위] |

## 비기능 요구사항
| 항목 | 요구사항 | 측정 기준 |
|------|----------|-----------|
| 성능 | [성능 요구사항] | [측정 방법] |
| 보안 | [보안 요구사항] | [검증 방법] |
| 확장성 | [확장성 요구사항] | [평가 기준] |
| 사용성 | [사용성 요구사항] | [측정 지표] |
```

### Phase 2: 업무 프로세스 설계 (Business Process Design)

#### 3단계: process.md 생성 (BPMN 2.0 기반)
```markdown
# {기능명} 업무 프로세스 설계

## 1. 프로세스 개요 (Process Overview)
- 비즈니스 프로세스의 목적과 범위
- 주요 이해관계자 (Stakeholders)
- 프로세스 성과 지표 (KPIs)

## 2. 프로세스 맵 (Process Map)
### 2.1 메인 프로세스 (Main Process)
```
[시작] → [활동1] → [게이트웨이] → [활동2] → [종료]
```

### 2.2 서브 프로세스 (Sub Processes)
- 각 주요 활동별 세부 프로세스

## 3. 역할 정의 (Role Definitions)
### 3.1 역할별 책임 (Role Responsibilities)
- **역할명**: 책임 범위, 권한, 의사결정 권한

### 3.2 RACI 매트릭스
| 작업 | 역할1 | 역할2 | 역할3 |
|------|-------|-------|-------|
| 작업1 | R | A | C |
| 작업2 | A | R | I |

## 4. 계층화된 작업 분해 (Hierarchical Task Breakdown)
### 4.1 구조 표기법 (Structure Notation)
```
P1: 프로세스명 (Process)
  R1: 역할명 (Role)
    T1: 작업명 (Task)
      FG1: 기능그룹명 (Function Group/UI)
        S1: 단계명 (Step)
          DS1: 세부단계명 (Detail Step)
            L1: 로직명 (Logic/AST)
          DS2: 세부단계명
            L2: 로직명
        S2: 단계명
          DS3: 세부단계명
            L3: 로직명
      FG2: 기능그룹명
        S3: 단계명
          DS4: 세부단계명
            L4: 로직명
    T2: 작업명
      FG3: 기능그룹명
        S4: 단계명
          DS5: 세부단계명
            L5: 로직명
  R2: 역할명
    T3: 작업명
      FG4: 기능그룹명
        S5: 단계명
          DS6: 세부단계명
            L6: 로직명
```

### 4.2 실제 작업 분해 예시 (Example Breakdown)
```
P1: 사용자 인증 프로세스 (User Authentication Process)
  목적: 안전한 사용자 로그인 및 세션 관리
  입력: 사용자 자격증명 (이메일, 비밀번호)
  출력: 인증 토큰 및 사용자 세션
  소유자: 보안팀

  R1: 프론트엔드 개발자 (Frontend Developer)
    책임: 사용자 인터페이스 구현 및 클라이언트 측 검증
    권한: UI/UX 설계 결정, 클라이언트 로직 구현
    스킬: React/Vue.js, JavaScript, CSS, UX 디자인

    T1: 로그인 폼 구현 (Login Form Implementation)
      설명: 사용자 입력을 받는 로그인 인터페이스 구현
      선행조건: UI 디자인 완료, API 스펙 확정
      완료조건: 폼 검증 및 제출 기능 완료
      예상시간: 2일

      FG1: 입력 폼 컴포넌트 (Input Form Component)
        UI 컴포넌트: EmailInput, PasswordInput, SubmitButton
        기능 목록: 입력 검증, 에러 표시, 자동완성
        사용자 인터랙션: 키보드 입력, 마우스 클릭, 터치

        S1: 이메일 입력 검증 (Email Input Validation)
          액션: 사용자 이메일 입력 시 실시간 검증
          입력: 이메일 문자열
          검증: 이메일 형식, 길이 제한
          예외처리: 잘못된 형식 시 에러 메시지 표시

          DS1: 이메일 형식 체크 (Email Format Check)
            세부 액션: 정규식을 이용한 이메일 패턴 매칭
            데이터 변환: 입력 문자열 → 불린 값
            비즈니스 규칙: RFC 5322 표준 준수

            L1: 이메일 정규식 검증 로직 (Email Regex Logic)
              알고리즘: 정규식 패턴 매칭
              의사결정 트리: 
                IF 빈 문자열 → 필수 입력 에러
                IF 형식 불일치 → 형식 에러
                ELSE → 검증 통과
              데이터 구조: String → Boolean
              복잡도: O(n) 시간, O(1) 공간

          DS2: 중복 이메일 체크 (Duplicate Email Check)
            세부 액션: 서버 API 호출하여 이메일 중복 확인
            데이터 변환: 이메일 → API 요청 → 불린 응답
            비즈니스 규칙: 기존 사용자와 중복 불가

            L2: 비동기 중복 체크 로직 (Async Duplicate Check)
              알고리즘: Debounced API 호출
              의사결정 트리:
                IF API 에러 → 네트워크 에러 표시
                IF 중복 존재 → 중복 에러 표시
                ELSE → 사용 가능 표시
              데이터 구조: Promise<Boolean>
              복잡도: O(1) 시간, O(1) 공간

        S2: 비밀번호 입력 검증 (Password Input Validation)
          액션: 비밀번호 강도 검증 및 보안 표시
          입력: 비밀번호 문자열
          검증: 길이, 복잡도, 특수문자 포함
          예외처리: 약한 비밀번호 시 경고 표시

          DS3: 비밀번호 강도 계산 (Password Strength Calculation)
            세부 액션: 비밀번호 복잡도 점수 계산
            데이터 변환: 문자열 → 강도 점수 (0-100)
            비즈니스 규칙: 최소 8자, 대소문자+숫자+특수문자

            L3: 비밀번호 점수 알고리즘 (Password Score Algorithm)
              알고리즘: 가중치 기반 점수 계산
              의사결정 트리:
                길이 점수 (length * 4)
                대문자 점수 (uppercase * 2)
                소문자 점수 (lowercase * 2)
                숫자 점수 (numbers * 4)
                특수문자 점수 (symbols * 6)
              데이터 구조: String → Number
              복잡도: O(n) 시간, O(1) 공간

      FG2: 제출 버튼 컴포넌트 (Submit Button Component)
        UI 컴포넌트: LoadingButton, ErrorMessage
        기능 목록: 폼 제출, 로딩 상태 표시, 에러 처리
        사용자 인터랙션: 클릭, 키보드 엔터

        S3: 폼 제출 처리 (Form Submit Handling)
          액션: 검증된 데이터를 서버로 전송
          입력: 검증된 폼 데이터
          검증: 모든 필드 유효성 확인
          예외처리: 네트워크 에러, 서버 에러 처리

          DS4: API 호출 및 응답 처리 (API Call & Response Handling)
            세부 액션: HTTP POST 요청 및 응답 파싱
            데이터 변환: FormData → JSON → AuthToken
            비즈니스 규칙: 토큰 저장, 리다이렉트 처리

            L4: 인증 API 통신 로직 (Auth API Communication Logic)
              알고리즘: HTTP 클라이언트 요청/응답 처리
              의사결정 트리:
                IF 200 OK → 토큰 저장 후 대시보드 이동
                IF 401 Unauthorized → 인증 실패 메시지
                IF 422 Validation Error → 필드별 에러 표시
                ELSE → 일반 에러 메시지
              데이터 구조: Promise<AuthResponse>
              복잡도: O(1) 시간, O(1) 공간

    T2: 세션 관리 구현 (Session Management Implementation)
      설명: 사용자 세션 상태 관리 및 자동 로그아웃
      선행조건: 인증 토큰 발급 완료
      완료조건: 세션 만료 처리 및 자동 갱신 완료
      예상시간: 1일

      FG3: 세션 스토어 컴포넌트 (Session Store Component)
        UI 컴포넌트: SessionProvider, AuthGuard
        기능 목록: 토큰 저장, 자동 갱신, 만료 처리
        사용자 인터랙션: 자동 처리 (사용자 개입 없음)

        S4: 토큰 자동 갱신 (Token Auto Refresh)
          액션: 토큰 만료 전 자동 갱신 요청
          입력: 현재 토큰, 만료 시간
          검증: 토큰 유효성, 갱신 가능 여부
          예외처리: 갱신 실패 시 로그아웃 처리

          DS5: 토큰 만료 감지 (Token Expiry Detection)
            세부 액션: JWT 토큰 만료 시간 파싱 및 타이머 설정
            데이터 변환: JWT → 만료시간 → 타이머 설정
            비즈니스 규칙: 만료 5분 전 갱신 시도

            L5: JWT 만료 타이머 로직 (JWT Expiry Timer Logic)
              알고리즘: JWT 디코딩 및 setTimeout 설정
              의사결정 트리:
                IF 토큰 없음 → 로그인 페이지 이동
                IF 만료 임박 → 갱신 API 호출
                IF 갱신 성공 → 새 토큰으로 타이머 재설정
                IF 갱신 실패 → 강제 로그아웃
              데이터 구조: JWT → Date → Timer
              복잡도: O(1) 시간, O(1) 공간

  R2: 백엔드 개발자 (Backend Developer)
    책임: 서버 측 인증 로직 및 API 구현
    권한: 데이터베이스 스키마 설계, 보안 정책 구현
    스킬: Node.js/Python, JWT, 암호화, 데이터베이스

    T3: 인증 API 구현 (Authentication API Implementation)
      설명: 사용자 인증 및 토큰 발급 API 개발
      선행조건: 데이터베이스 스키마 완료, 보안 정책 수립
      완료조건: 로그인/로그아웃 API 완료, 보안 테스트 통과
      예상시간: 3일

      FG4: 인증 엔드포인트 (Authentication Endpoints)
        UI 컴포넌트: REST API (POST /auth/login, POST /auth/logout)
        기능 목록: 로그인, 로그아웃, 토큰 갱신, 비밀번호 재설정
        사용자 인터랙션: HTTP 요청/응답

        S5: 로그인 요청 처리 (Login Request Processing)
          액션: 사용자 자격증명 검증 및 토큰 발급
          입력: 이메일, 비밀번호
          검증: 사용자 존재 여부, 비밀번호 일치
          예외처리: 잘못된 자격증명, 계정 잠금 처리

          DS6: 비밀번호 해시 검증 (Password Hash Verification)
            세부 액션: 입력 비밀번호와 저장된 해시 비교
            데이터 변환: 평문 비밀번호 → 해시 → 비교 결과
            비즈니스 규칙: bcrypt 사용, 솔트 적용

            L6: bcrypt 해시 비교 로직 (bcrypt Hash Comparison Logic)
              알고리즘: bcrypt.compare() 함수 사용
              의사결정 트리:
                IF 사용자 없음 → 404 에러
                IF 비밀번호 불일치 → 401 에러
                IF 계정 잠금 → 423 에러
                ELSE → JWT 토큰 발급
              데이터 구조: String → Promise<Boolean>
              복잡도: O(1) 시간, O(1) 공간 (해시 함수 특성)
```

## 5. 프로세스 플로우 (Process Flow)
### 5.1 해피 패스 (Happy Path)
```
시작 → 입력검증 → 비즈니스로직 → 결과저장 → 종료
```

### 5.2 예외 시나리오 (Exception Scenarios)
- **시나리오 1**: [예외 상황]
  - **트리거**: [발생 조건]
  - **처리**: [대응 방법]
  - **복구**: [복구 절차]

## 6. 비즈니스 규칙 (Business Rules)
### 6.1 검증 규칙 (Validation Rules)
- **VR1**: [규칙명] - [규칙 설명]

### 6.2 계산 규칙 (Calculation Rules)
- **CR1**: [규칙명] - [계산 공식]

### 6.3 의사결정 규칙 (Decision Rules)
- **DR1**: [규칙명] - [의사결정 기준]

## 7. 데이터 플로우 (Data Flow)
### 7.1 데이터 입력 (Data Input)
- **소스**: [데이터 출처]
- **형식**: [데이터 형식]
- **검증**: [검증 기준]

### 7.2 데이터 변환 (Data Transformation)
- **변환 규칙**: [데이터 변환 방식]
- **매핑**: [필드 매핑]

### 7.3 데이터 출력 (Data Output)
- **대상**: [데이터 목적지]
- **형식**: [출력 형식]

## 8. 성능 요구사항 (Performance Requirements)
- **처리량**: [초당 처리 건수]
- **응답시간**: [최대 응답 시간]
- **동시사용자**: [최대 동시 사용자 수]

## 9. 모니터링 및 측정 (Monitoring & Metrics)
### 9.1 프로세스 KPI
- **효율성**: [처리 시간, 처리량]
- **품질**: [오류율, 정확도]
- **만족도**: [사용자 만족도]

### 9.2 알림 및 에스컬레이션
- **임계값**: [알림 기준]
- **에스컬레이션**: [단계별 대응]

## 10. 리스크 및 완화방안 (Risks & Mitigation)
| 리스크 | 영향도 | 확률 | 완화방안 |
|--------|--------|------|----------|
| 리스크1 | 높음 | 중간 | 대응방안1 |
| 리스크2 | 중간 | 낮음 | 대응방안2 |
```

### Phase 3: 기술 설계 (Technical Design)

#### 4단계: design.md 생성
```markdown
# {기능명} 기술 설계

## 아키텍처 개요
- 시스템 구조: [전체 아키텍처]
- 핵심 컴포넌트: [주요 구성요소]
- 기술 제약사항: [UI/UX 설계 시 고려사항]

## 데이터 모델
| 엔티티 | 속성 | 관계 | 제약조건 |
|--------|------|------|----------|
| [엔티티명] | [주요 속성] | [관계 정보] | [제약사항] |

## API 설계
| 엔드포인트 | 메서드 | 요청 | 응답 | 목적 |
|------------|--------|------|------|------|
| [URL] | [HTTP 메서드] | [요청 형식] | [응답 형식] | [기능 설명] |

## 정확성 속성
| 속성 ID | 속성 설명 | 검증 방법 | 연결 요구사항 |
|---------|----------|-----------|---------------|
| P1 | [속성 내용] | [PBT 방법] | [Requirements ID] |

## 기술 스택
| 계층 | 기술 | 버전 | 선택 이유 |
|------|------|------|----------|
| Frontend | [기술명] | [버전] | [선택 근거] |
| Backend | [기술명] | [버전] | [선택 근거] |
| Database | [기술명] | [버전] | [선택 근거] |

## 보안 고려사항
| 영역 | 위험 | 대응방안 | 구현 방법 |
|------|------|----------|-----------|
| 인증 | [보안 위험] | [보안 대책] | [구현 상세] |
| 권한 | [보안 위험] | [보안 대책] | [구현 상세] |
| 데이터 | [보안 위험] | [보안 대책] | [구현 상세] |
```

### Phase 4: UI/UX 설계 (UI/UX Design)

#### 5단계: ui-design.md 생성
```markdown
# {기능명} UI/UX 설계

## 1. 설계 개요 (Design Overview)
- UI/UX 설계 목표 및 원칙
- 타겟 사용자 및 사용 환경
- 디자인 시스템 및 브랜드 가이드라인

## 2. 사용자 여정 맵 (User Journey Map)
### 2.1 사용자 페르소나 (User Personas)
- **페르소나 1**: [이름] - [특성 및 니즈]
- **페르소나 2**: [이름] - [특성 및 니즈]

### 2.2 사용자 플로우 (User Flow)
```
시작점 → 액션1 → 결정점 → 액션2 → 목표 달성
```

### 2.3 터치포인트 분석 (Touchpoint Analysis)
- **진입점**: 사용자가 기능에 접근하는 방법
- **상호작용점**: 주요 사용자 액션 지점
- **이탈점**: 사용자가 이탈할 수 있는 지점

## 3. 정보 아키텍처 (Information Architecture)
### 3.1 사이트맵 (Sitemap)
```
홈
├── 기능 A
│   ├── 서브기능 A1
│   └── 서브기능 A2
├── 기능 B
└── 설정
```

### 3.2 네비게이션 구조 (Navigation Structure)
- **주 네비게이션**: 핵심 기능 접근
- **보조 네비게이션**: 부가 기능 및 설정
- **브레드크럼**: 현재 위치 표시

## 4. 와이어프레임 (Wireframes)
### 4.1 로우파이 와이어프레임 (Low-fidelity Wireframes)
```
[페이지명 1]
┌─────────────────────────────────┐
│ Header                          │
├─────────────────────────────────┤
│ Navigation                      │
├─────────────────────────────────┤
│ Main Content Area               │
│ ┌─────────────┐ ┌─────────────┐ │
│ │   Section   │ │   Sidebar   │ │
│ │             │ │             │ │
│ └─────────────┘ └─────────────┘ │
├─────────────────────────────────┤
│ Footer                          │
└─────────────────────────────────┘
```

### 4.2 하이파이 와이어프레임 (High-fidelity Wireframes)
- 상세한 레이아웃 및 컴포넌트 배치
- 실제 콘텐츠 및 이미지 영역 정의
- 반응형 디자인 고려사항

## 5. UI 컴포넌트 설계 (UI Component Design)
### 5.1 컴포넌트 라이브러리 (Component Library)
```
기본 컴포넌트 (Atoms)
├── Button
│   ├── Primary Button
│   ├── Secondary Button
│   └── Icon Button
├── Input
│   ├── Text Input
│   ├── Password Input
│   └── Search Input
├── Typography
│   ├── Heading (H1-H6)
│   ├── Body Text
│   └── Caption
└── Icon
    ├── Action Icons
    └── Status Icons

복합 컴포넌트 (Molecules)
├── Form Field
│   ├── Label + Input + Error Message
│   └── Label + Input + Help Text
├── Card
│   ├── Basic Card
│   └── Action Card
└── Navigation Item
    ├── Menu Item
    └── Tab Item

페이지 컴포넌트 (Organisms)
├── Header
│   ├── Logo + Navigation + User Menu
│   └── Search Bar + Notifications
├── Form
│   ├── Login Form
│   └── Registration Form
└── Content Section
    ├── Article List
    └── Detail View
```

### 5.2 컴포넌트 명세 (Component Specifications)
```
Button 컴포넌트
├── Props
│   ├── variant: 'primary' | 'secondary' | 'outline'
│   ├── size: 'small' | 'medium' | 'large'
│   ├── disabled: boolean
│   └── onClick: () => void
├── States
│   ├── Default
│   ├── Hover
│   ├── Active
│   ├── Focus
│   └── Disabled
└── Accessibility
    ├── ARIA labels
    ├── Keyboard navigation
    └── Screen reader support
```

## 6. 인터랙션 설계 (Interaction Design)
### 6.1 마이크로 인터랙션 (Micro-interactions)
- **버튼 클릭**: 시각적 피드백 및 상태 변화
- **폼 입력**: 실시간 검증 및 에러 표시
- **로딩 상태**: 프로그레스 인디케이터 및 스켈레톤 UI

### 6.2 애니메이션 및 전환 (Animations & Transitions)
```
페이지 전환
├── Fade In/Out (300ms ease-in-out)
├── Slide Left/Right (250ms ease-in-out)
└── Scale Up/Down (200ms ease-in-out)

컴포넌트 애니메이션
├── Button Hover (150ms ease-in-out)
├── Modal Open/Close (300ms ease-in-out)
└── Dropdown Expand (200ms ease-in-out)
```

### 6.3 제스처 및 입력 (Gestures & Input)
- **터치 제스처**: 스와이프, 핀치, 탭
- **키보드 단축키**: 주요 액션에 대한 단축키
- **마우스 인터랙션**: 호버, 클릭, 드래그

## 7. 반응형 디자인 (Responsive Design)
### 7.1 브레이크포인트 (Breakpoints)
```
Mobile: 320px - 767px
├── Single column layout
├── Touch-optimized controls
└── Simplified navigation

Tablet: 768px - 1023px
├── Two column layout
├── Hybrid touch/mouse controls
└── Collapsible sidebar

Desktop: 1024px+
├── Multi-column layout
├── Mouse-optimized controls
└── Full navigation menu
```

### 7.2 적응형 컴포넌트 (Adaptive Components)
- **네비게이션**: 햄버거 메뉴 ↔ 풀 메뉴
- **데이터 테이블**: 스크롤 ↔ 카드 뷰
- **폼 레이아웃**: 세로 배치 ↔ 가로 배치

## 8. 접근성 설계 (Accessibility Design)
### 8.1 WCAG 2.1 준수사항
- **인식 가능성**: 색상 대비, 텍스트 크기, 이미지 대체 텍스트
- **운용 가능성**: 키보드 네비게이션, 포커스 관리
- **이해 가능성**: 명확한 라벨, 일관된 네비게이션
- **견고성**: 스크린 리더 호환성, 시맨틱 HTML

### 8.2 접근성 체크리스트
- [ ] 색상만으로 정보 전달하지 않음
- [ ] 충분한 색상 대비 (4.5:1 이상)
- [ ] 키보드로 모든 기능 접근 가능
- [ ] 포커스 인디케이터 명확히 표시
- [ ] 스크린 리더용 ARIA 라벨 제공

## 9. 디자인 시스템 (Design System)
### 9.1 컬러 팔레트 (Color Palette)
```
Primary Colors
├── Primary-50: #f0f9ff
├── Primary-500: #3b82f6 (Main)
└── Primary-900: #1e3a8a

Secondary Colors
├── Secondary-50: #f8fafc
├── Secondary-500: #64748b (Main)
└── Secondary-900: #0f172a

Semantic Colors
├── Success: #10b981
├── Warning: #f59e0b
├── Error: #ef4444
└── Info: #3b82f6
```

### 9.2 타이포그래피 (Typography)
```
Font Family
├── Primary: 'Inter', sans-serif
└── Monospace: 'JetBrains Mono', monospace

Font Scale
├── H1: 2.25rem (36px) / Bold
├── H2: 1.875rem (30px) / Bold
├── H3: 1.5rem (24px) / Semibold
├── Body: 1rem (16px) / Regular
└── Caption: 0.875rem (14px) / Regular
```

### 9.3 간격 시스템 (Spacing System)
```
Spacing Scale (8px base)
├── xs: 4px (0.25rem)
├── sm: 8px (0.5rem)
├── md: 16px (1rem)
├── lg: 24px (1.5rem)
├── xl: 32px (2rem)
└── 2xl: 48px (3rem)
```

## 10. 프로토타입 (Prototype)
### 10.1 인터랙티브 프로토타입
- **도구**: Figma, Adobe XD, Sketch
- **범위**: 핵심 사용자 플로우
- **피드백**: 사용자 테스트 결과 반영

### 10.2 프로토타입 시나리오
```
시나리오 1: 신규 사용자 온보딩
├── 랜딩 페이지 진입
├── 회원가입 플로우
├── 이메일 인증
└── 초기 설정 완료

시나리오 2: 기존 사용자 로그인
├── 로그인 페이지 진입
├── 자격증명 입력
├── 2FA 인증 (선택)
└── 대시보드 이동
```

## 11. 사용성 테스트 계획 (Usability Testing Plan)
### 11.1 테스트 목표
- 사용자 플로우의 직관성 검증
- 인터페이스 요소의 명확성 확인
- 접근성 및 사용성 문제 발견

### 11.2 테스트 방법론
- **A/B 테스트**: 대안 디자인 비교
- **사용자 인터뷰**: 정성적 피드백 수집
- **태스크 분석**: 완료율 및 소요시간 측정

## 12. 구현 가이드라인 (Implementation Guidelines)
### 12.1 개발자 핸드오프
- **디자인 토큰**: CSS 변수 및 상수 정의
- **컴포넌트 명세**: Props 및 상태 정의
- **애니메이션 스펙**: 지속시간 및 이징 함수

### 12.2 품질 보증
- **디자인 QA**: 구현된 UI와 디자인 일치성 검증
- **크로스 브라우저 테스트**: 다양한 환경에서 동작 확인
- **성능 최적화**: 이미지 압축, 코드 분할
```

### Phase 5: 작업 계획 (Task Planning)

#### 6단계: tasks.md 생성
```markdown
# {기능명} 구현 작업

## 작업 상태 표기법
- `- [ ]` = 미시작 (필수 작업)
- `- [ ]*` = 미시작 (선택 작업)
- `- [~]` = 대기중
- `- [-]` = 진행중
- `- [x]` = 완료

## 구현 작업 목록

### 1. 환경 설정
- [ ] 1.1 프로젝트 초기 설정
- [ ] 1.2 의존성 설치
- [ ] 1.3 개발 환경 구성

### 2. 핵심 기능 구현
- [ ] 2.1 데이터 모델 구현
- [ ] 2.2 비즈니스 로직 구현
- [ ] 2.3 API 엔드포인트 구현

### 3. 테스트 구현
- [ ] 3.1 단위 테스트 작성
- [ ] 3.2 Property-Based 테스트 작성
  - **검증 대상**: Requirements 1.1, 1.2
- [ ] 3.3 통합 테스트 작성

### 4. 사용자 인터페이스
- [ ] 4.1 UI 컴포넌트 구현
- [ ] 4.2 사용자 플로우 구현
- [ ]* 4.3 접근성 개선 (선택사항)

### 5. 배포 및 운영
- [ ] 5.1 배포 스크립트 작성
- [ ]* 5.2 모니터링 설정 (선택사항)
```

---

## 🧪 Property-Based Testing 통합

### PBT 작업 형식
```markdown
- [ ] X.Y Property-Based 테스트: [속성명]
  - **검증 대상**: Requirements A.B, A.C
  - **속성**: [구체적인 속성 설명]
  - **생성기**: [테스트 데이터 생성 전략]
```

### PBT 상태 관리
- `passed`: 테스트 성공
- `failed`: 테스트 실패 (반례 포함)
- `unexpected_pass`: 버그 탐지 테스트가 예상과 달리 통과
- `not_run`: 미실행

---

## 📋 실행 지침 (Execution Guidelines)

### 단일 작업 실행 모드
1. **사전 준비**: requirements.md, design.md, tasks.md 읽기
2. **작업 선택**: 하나의 작업에만 집중
3. **구현**: 코드 작성 → 테스트 → 검증
4. **완료 보고**: 사용자 검토 대기

### 전체 작업 실행 모드
1. **작업 파싱**: 미완료 필수 작업 식별
2. **순차 실행**: 하위 작업부터 상위 작업 순으로
3. **상태 업데이트**: 각 단계별 진행 상황 업데이트
4. **오류 처리**: 실패 시 중단 및 사용자 보고

### 테스트 가이드라인
- **단위 테스트**: 구체적 예시와 경계값 검증
- **Property-Based 테스트**: 범용 속성의 모든 입력에 대한 검증
- **최대 시도**: 테스트 실패 시 최대 2회 수정 시도
- **실제 기능**: 모킹 없이 실제 기능 검증

---

## 🔄 버그 수정 워크플로우

### 버그 탐지 테스트 (Task 1)
```markdown
- [ ] 1.1 버그 조건 탐지 Property 테스트 작성
  - **목적**: 버그 존재 확인
  - **예상 결과**: 테스트 실패 (버그 발견)
  - **예상외 통과**: 버그 재조사 필요
```

### 상태별 처리
- **실패 (예상)**: `passed` 상태로 기록, 반례 문서화
- **통과 (예상외)**: `unexpected_pass` 상태, 사용자 선택 요청
  - "계속 진행" vs "재조사"

---

## 📝 파일 참조 시스템

### 외부 파일 참조
```markdown
#[[file:openapi.yaml]]
#[[file:schema.graphql]]
#[[file:config.json]]
```

### 사용 사례
- OpenAPI 명세서 참조
- GraphQL 스키마 참조
- 설정 파일 참조

---

## ✅ 품질 체크리스트

### 요구사항 문서
- [ ] 사용자 스토리가 명확한가?
- [ ] 수용 기준이 테스트 가능한가?
- [ ] 제약사항이 구체적인가?

### 프로세스 문서
- [ ] BPMN 2.0 표준을 준수하는가?
- [ ] 7단계 드릴다운이 완성되었는가?
- [ ] 비즈니스 규칙이 명확히 정의되었는가?
- [ ] 예외 시나리오가 포함되었는가?
- [ ] 성능 요구사항이 구체적인가?

### 기술 설계 문서
- [ ] 아키텍처가 요구사항을 만족하는가?
- [ ] 기술 제약사항이 UI 설계에 반영되었는가?
- [ ] 정확성 속성이 실행 가능한가?
- [ ] 보안 고려사항이 포함되었는가?

### UI 설계 문서
- [ ] 사용자 여정 맵이 명확한가?
- [ ] 와이어프레임이 완성되었는가?
- [ ] UI 컴포넌트가 체계적으로 설계되었는가?
- [ ] 반응형 디자인이 고려되었는가?
- [ ] 접근성 가이드라인을 준수하는가?
- [ ] 기술 제약사항 내에서 설계되었는가?

### 기술 설계 문서
- [ ] 아키텍처가 요구사항을 만족하는가?
- [ ] 정확성 속성이 실행 가능한가?
- [ ] 보안 고려사항이 포함되었는가?

### 작업 목록
- [ ] 작업이 구체적이고 실행 가능한가?
- [ ] 의존성이 올바르게 정의되었는가?
- [ ] 테스트 작업이 포함되었는가?

---

## 🎯 성공 기준

### 완성된 스펙의 조건
1. **포괄적 요구사항**: 모든 기능과 제약사항 문서화
2. **체계적 프로세스 설계**: BPMN 2.0 기반 7단계 드릴다운 완료
3. **사용자 중심 UI 설계**: 와이어프레임, 컴포넌트, 인터랙션 설계 완료
4. **실행 가능한 기술 설계**: 구현 가능한 구체적 기술 설계
5. **검증 가능한 속성**: PBT로 검증할 수 있는 정확성 속성
6. **실행 가능한 계획**: 단계별 구현 작업 목록
7. **사용자 승인**: 각 단계별 사용자 검토 완료

### 구현 완료의 조건
1. **모든 필수 작업 완료**: tasks.md의 필수 작업 100% 완료
2. **테스트 통과**: 단위 테스트 및 PBT 모두 통과
3. **요구사항 만족**: 모든 수용 기준 충족
4. **코드 품질**: 보안, 성능, 유지보수성 확보

---

## 🚀 시작하기

### 새 스펙 생성
```
"[기능 아이디어]에 대한 스펙을 만들어주세요"
```

### 기존 스펙 업데이트
```
"[기능명] 스펙의 프로세스 문서를 업데이트해주세요"
"[기능명] 스펙의 기술 설계 문서를 업데이트해주세요"
"[기능명] 스펙의 UI 설계 문서를 업데이트해주세요"
"[기능명] 스펙의 작업 목록을 새로고침해주세요"
```

### 작업 실행
```
"[기능명] 스펙의 [작업번호] 작업을 실행해주세요"
"[기능명] 스펙의 모든 작업을 실행해주세요"
```

---

**Methodology**: Spec-Driven Development with Property-Based Testing
**Version**: 2.0
**Last Updated**: 2026-02-06