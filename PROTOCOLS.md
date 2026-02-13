# 📜 Trading Agent Global Protocols (v1.1)

본 저장소의 모든 에이전트는 아래 규정을 강제 준수한다.

### 1. [GP-ETO] 시각 관리 (Execution-Time vs Session-Time)
- **파일명**: `YYYYMMDD_HHmm_[TAG]_NAME.md` 형식을 사용하며, `HHmm`은 해당 분석 **세션의 시작 시각**으로 통일하여 리포팅 서버에서 그룹화되도록 한다.
- **본문 시각**: 리포트 내부의 '발행 시각'은 `date` 명령어의 **실제 호출 시각(실측치)**을 정직하게 기록한다.
- 미래 시점 기입 금지. 데이터 정직성 최우선.

### 2. [GP-SPEC] KIS API 명세 우선 (Spec First)
- 호출 전 `find_api_detail` 필수 수행.
- **필수 인자 준수**: `env_dv`('real'/'demo') 등 API별 필수 파라미터 누락 방지를 위해 반환된 명세를 엄격히 준수한다.
- 추측 금지. 사용 API명 및 필수 `fid_` 인자 리포트 하단 적시.

### 3. [GP-EBR] 근거 중심 리포팅 (Evidence-Based)
- "좋음/나쁨" 금지. 수치(Price/Vol/Ratio) + 시각(Time) + 추론(CoT) 필수.
- 모든 보고서 상단에 **Quick Scan Matrix (🔴/🟡/🟢)** 포함.

### 4. [GP-VFY] 데이터 교차 검증 (Cross-Check)
- KIS 데이터 기본 + **Yahoo Finance/Naver Stock MCP** 대조 필수.
- **Yahoo (get_current_stock_price)**: 나스닥 선물(`NQ=F`), S&P500 선물(`ES=F`), 환율(`USDKRW=X`) 실시간 확인.
- **Naver (search_stock)**: 국내 종목 테마 및 실시간 수급 보완.
- 괴리 발생 시 `[리스크]` 보고서에 수치 기입 및 경고.

### 5. [GP-REV] 복기 우선 (Review First)
- 분석 시작 전 "어제 가설 vs 오늘 결과" 대조 및 가중치 보정.

### 6. [GP-KIS-EX] KIS API 호출 예시 (Golden Examples)
- **잔고 조회**: `domestic_stock({ "api_type": "inquire_balance", "params": { "afhr_flpr_yn": "N", "env_dv": "real", "fncg_amt_auto_rdpt_yn": "N", "fund_sttl_icld_yn": "N", "inqr_dvsn": "02", "prcs_dvsn": "01", "tr_cont": "", "unpr_dvsn": "01" } })`
- **투자자별 매매동향**: `domestic_stock({ "api_type": "inquire_investor", "params": { "env_dv": "real", "fid_cond_mrkt_div_code": "J", "fid_input_iscd": "005930" } })`
- **시간외/NX 시세**: `domestic_stock({ "api_type": "inquire_overtime_price", "params": { "fid_cond_mrkt_div_code": "J", "fid_input_iscd": "005930" } })` (NX 시장은 `fid_cond_mrkt_div_code: "NX"`)
- **일자별 가격**: `domestic_stock({ "api_type": "inquire_daily_price", "params": { "env_dv": "real", "fid_cond_mrkt_div_code": "J", "fid_input_iscd": "005930", "fid_org_adj_prc": "1", "fid_period_div_code": "D" } })`
