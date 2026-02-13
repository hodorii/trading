name: kis-api-expert [S]
description: 한국투자증권(KIS) API 호출의 정확성을 보장하기 위한 파라미터 맵핑 및 가이드 스킬

# 🚨 KIS API 호출의 절대 원칙
1. **명세 우선 확인**: 새로운 API를 사용하기 전 무조건 `find_api_detail`을 실행한다.
2. **파라미터 엄격 준수**: `find_api_detail`의 `params` 목록에 없는 인자(특히 `env_dv`)를 임의로 추가하여 호출하지 않는다.
3. **필수 인자 누락 금지**: 상세 명세에 `required: true`로 표시된 모든 `fid_...` 인자는 반드시 포함한다.
4. **데이터 타입 준수**: 문자열인지 숫자인지 확인하여 정확한 형식으로 전달한다.

# 🛠️ 주요 API별 완벽 호출 가이드

## 1. 시간외/NX 현재가 (`inquire_overtime_price`)
- **목적**: 정규장 이후 시간외 단일가 또는 넥스트레이드(NX) 시세 확인.
- **필수 파라미터**:
    - `fid_cond_mrkt_div_code`: 시장 구분 (J: KRX, NX: NXT)
    - `fid_input_iscd`: 종목코드 (6자리)
- **주의**: **`env_dv` 인자를 지원하지 않음**. 포함 시 `TypeError` 발생.
- **예시**:
  ```python
  domestic_stock(api_type="inquire_overtime_price", params={
      "fid_cond_mrkt_div_code": "J",
      "fid_input_iscd": "005930"
  })
  ```

## 2. 시가총액 상위 (`market_cap`)
- **목적**: 시장별 시가총액 순위 및 규모 확인.
- **필수 파라미터**:
    - `fid_cond_mrkt_div_code`: 'J' (KRX)
    - `fid_cond_scr_div_code`: '20174' (고정값)
    - `fid_div_cls_code`: '0' (전체)
    - `fid_input_iscd`: '0000' (전체)
    - `fid_trgt_cls_code`: '0'
    - `fid_trgt_exls_cls_code`: '0'
    - `fid_input_price_1`: '0'
    - `fid_input_price_2`: '0'
    - `fid_vol_cnt`: '0'
- **주의**: `env_dv`는 포함하지 않음(오류 발생함).

## 3. 등락률 순위 (`fluctuation`)
- **목적**: 당일 상승/하락 상위 종목 포착.
- **필수 파라미터**:
    - `fid_cond_mrkt_div_code`: 'J' (KRX)
    - `fid_cond_scr_div_code`: '20170' (고정값)
    - `fid_input_iscd`: '0000' (전체)
    - `fid_rank_sort_cls_code`: '0000' (등락률순)
    - `fid_input_cnt_1`: '0' (또는 조회 건수)
    - `fid_prc_cls_code`: '0'
    - `fid_input_price_1`: '0'
    - `fid_input_price_2`: '0'
    - `fid_vol_cnt`: '0'
    - `fid_trgt_cls_code`: '0'
    - `fid_trgt_exls_cls_code`: '0'
    - `fid_div_cls_code`: '0'
    - `fid_rsfl_rate1`: '0.0'
    - `fid_rsfl_rate2`: '0.0'

## 4. 종합 시황/공시 (`news_title`)
- **목적**: 실시간 뉴스 및 기업 공시 제목 조회.
- **필수 파라미터**:
    - `fid_news_ofer_entp_code`: 뉴스사 코드 (제공업체 확인 필요)
    - `fid_cond_mrkt_cls_code`: 시장 구분 코드
    - `fid_input_iscd`: 종목코드 (또는 전체)
    - `fid_titl_cntt`: 제목 검색어 (없으면 공란)
    - `fid_input_date_1`: 날짜
    - `fid_input_hour_1`: 시간
    - `fid_rank_sort_cls_code`: 정렬 구분
    - `fid_input_srno`: 일련번호 (처음 조회 시 '0' 또는 공란)

# 🚦 오류 발생 시 자가 진단 (Self-Check)
1. `TypeError: ... got an unexpected keyword argument 'env_dv'`: 해당 API는 `env_dv`를 지원하지 않음. 제거 후 재호출.
2. `missing X required positional arguments`: `find_api_detail`에서 `required: true`인 모든 인자를 채웠는지 확인.
3. `ERROR INVALID FID_COND_MRKT_DIV_CODE`: 코스피는 'J'가 아닌 다른 값이 필요한지 명세 확인(보통 KRX는 'J').

# 🔄 업데이트 가이드
신규 API 추가 시 반드시 `find_api_detail` 결과를 이 문서에 업데이트하여 지식을 자산화한다.
