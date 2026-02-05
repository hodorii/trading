---
name: event-driven [M]
description: 주요 경제 지표 발표 기업 공시 뉴스 등 이벤트 발생 시 시장의 즉각적 반응과 섹터 수급 이동을 분석하는 워크플로우
---

## 🚨 글로벌 규약 (Global Rules)
1. **한국어 우선 (Korean-First)**: 모든 리포트의 헤더와 핵심 용어는 '한국어'로 작성합니다.
2. **참고 자료 의무화 (Mandatory Reference)**: 모든 리포트 하단에 `## 참고 자료 (Reference)` 섹션을 만들고, 링크를 명시합니다.
3. **폴더 규칙 준수**: `d:\dev\trading\reports\YYYY-MM-DD\` 경로에 저장합니다.

## 📊 핵심 매매 원칙 (Core Trading Principles)
- **Global-First**: 모든 이벤트는 글로벌 지수(나스닥 선물 등)와의 연동성을 먼저 살핍니다.
- **Recency-First**: 뉴스의 '최초 보도 시각'을 확인하여 이미 주가에 반영된 재료인지(Priced-in), 새로운 충격인지 구분합니다.

## 📋 파일명 네이밍 컨벤션
### 형식
`YYYYMMDD_HHmm_WorkflowID_[태그]_종목명_설명.md`
(예: `20260127_1230_EVT_[이벤트]_우리기술_주요이벤트_분석보고서.md`)

## 🔗 상호 참조 (Cross-Reference)
- 동일 세션(`YYYYMMDD_HHmm`)에서 생성되는 다른 전문가들의 보고서 파일명을 예상하여 `## 참고 자료` 섹션에 수동으로 링크를 포함합니다.

# ⚡ event-driven [M] (이벤트 드리븐 스캐너)

## 🛠️ 도구 사용 지침
- **명세 우선 확인**: KIS API 호출 전 `find_api_detail` 필수.
- **필수값 누락주의**: `env_dv` 파라미터 필수.

## 🎯 분석 목적
- 뉴스/공시 발생 시 시장의 '진짜 반응'과 '가짜 반응'을 구분.
- 섹터 간 자금 이동(Rotation) 포착.

---

## 🛠️ 활용 도구 및 API
- **Macro**: `search_web` (뉴스 검색), `get_current_stock_price` (NQ=F, ES=F, USDKRW=X)
- **Domestic**: `news_title`, `fluctuation`, `get_stock_news`
- **NXT**: `inquire_overtime_price` (야간 시장 반응)

---

## 📈 분석 프로세스 (Process)

### 1단계: 글로벌 지수 연동 및 뉴스/공시 오염도 체크
- **Action**: 미 선물 지수 등락률 확인 및 뉴스/공시 신선도(Recency) 검증.
- **Check**: "이미 다 아는 뉴스인가?" 혹은 "이미 발표된 공시인가?" (기반영 여부 판단).
- **Tool**: `get_disclosure_list` (공시 목록 확인), `get_stock_news` (뉴스 목록 확인)
- **Table**: 뉴스·공시 토픽 / 발표 시각 / 현재 상태 / 오염도 판별(Low/Mid/High) 표 작성.

### 2단계: 공시 내용 심층 분석 (Disclosure Deep Dive)
- **Action**: 주요 공시(잠정실적, 무상증자, 인허가 획득 등)의 세부 내용 확인.
- **Check**: 공시 내용이 기업 가치(EPS)에 직접적인 영향을 주는지, 혹은 단순 경영상 절차인지 구분.
- **Tool**: `get_disclosure_content` (공시 세부 텍스트 분석)

### 2단계: 섹터별 파급력 분석
- **Action**: 해당 이벤트가 [수혜] 섹터와 [피해] 섹터에 미치는 영향 분석.
- **Check**: 단순 테마성인지, 펀더멘털(실적) 변화인지 구분.

### 3단계: NXT(야간) 및 수급 가상 시뮬레이션
- **Action**: 정규장 마감 후 야간 시장(NXT) 반응 확인.
- **Logic**: 
  - 호재 뉴스인데 야간 시장 하락? → 재료 소멸(Sell the news).
  - 악재 뉴스인데 야간 시장 상승? → 저가 매수 기회(Buy the dip).

### 4단계: 대응 전략 수립
- **Action**: 내일 장초반 시나리오(갭 상승/하락)별 대응 전술 작성.

---

## 📄 리포트 템플릿 (Report Template)
1. **이벤트 개요**: 핵심 키워드, 관련 섹터, D-Day
2. **뉴스 및 공시 타임라인 (Timeline)**: 주요 뉴스/공시 발생 시각 및 제목 정리
3. **신선도 및 오염도 체크**: 오염도 판별 테이블 (공시 포함)
4. **섹터별 파급력**: 수혜(Positive) vs 피해(Negative)
5. **NXT(야간) 및 수급 시뮬레이션**: 상황/예측/전술
6. **대응 전략**: 시나리오별 구체적 행동 강령
7. **결론**
8. **참고 자료 (Reference)**: [필수] DART 공시 링크, 관련 뉴스 링크 및 세션 내 다른 보고서 연동

---

**Next Action**: 분석된 주도 섹터를 대상으로 `day-trading [T]` 또는 `total-trading-strategy [I]` 실행.
