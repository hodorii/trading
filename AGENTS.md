# AGENTS.md - Antigravity Trading Intelligence System

## Build/Lint/Test Commands

This is an **Agentic AI System** with no traditional build/lint/test commands. The system operates through:

### Core Execution Commands
- **Main Workflows**: Execute via `.agent/workflows/*.md` files
  - `integrated-daily-routine.md` - IDR (Integrated Daily Routine) 
  - `multi-agent-decision.md` - MAD (Multi-Agent Decision)
  - `timeline-trading-ops.md` - TTO (Timeline Trading Ops)
- **Skills**: Execute via `.agent/skills/*/SKILL.md` files
- **MCP Servers**: Configure via `.gemini/mcp_config.json` and `.agent/mcp_config.json`

### MCP Server Commands
```bash
# Naver Finance API
d:/dev/naver-trading/naver-finance-mcp.exe

# Yahoo Finance API  
uvx mcp-yahoo-finance

# KIS Trading Remote
npx -y mcp-remote http://localhost:3000/sse

# Sequential Thinking (disabled by default)
npx -y @modelcontextprotocol/server-sequential-thinking
```

### Python Scripts
```bash
# Refactor report filenames
python scripts/refactor_filenames.py
```

## Code Style Guidelines

### Language Policy
- **Korean-First**: ALL content, analysis, reports, and communication MUST be in Korean
- **No English**: English expressions like "Event Analysis" are forbidden - use "이벤트 분석"
- **Pure Korean**: Use pure Korean expressions in YAML descriptions and all text content

### File Naming Convention
```
YYYYMMDD_HHmm_WorkflowID_[태그]_종목명_설명.md
```

#### Components
- **YYYYMMDD**: Execution date (e.g., 20260205)
- **HHmm**: Session ID using execution time (e.g., 1956)
- **WorkflowID**: Workflow identifier (IDR, MAD, TTO)
- **[태그]**: Report type identifier
- **종목명**: Analysis target or "전체" for market analysis
- **설명**: Report content summary

#### Tag System by Workflow
| Workflow | WorkflowID | 태그 | Description | Example |
|----------|-----------|------|-------------|---------|
| IDR | IDR | [매크로] | Macro index analysis | `20260205_1956_IDR_매크로_전체_매크로지수분석.md` |
| IDR | IDR | [이벤트] | Major event analysis | `20260205_1956_IDR_이벤트_전체_주요이벤트분석.md` |
| IDR | IDR | [시장] | Sector scanner | `20260205_1956_IDR_시장_전체_섹터스캐너.md` |
| IDR | IDR | [가치] | Stock value analysis | `20260205_1956_IDR_가치_삼성전자_가치분석.md` |
| IDR | IDR | [수급] | Stock supply-demand analysis | `20260205_1956_IDR_수급_삼성전자_수급분석.md` |
| IDR | IDR | [리스크] | Stock risk analysis | `20260205_1956_IDR_리스크_삼성전자_위험분석.md` |
| IDR | IDR | [결정] | Integrated stock strategy | `20260205_1956_IDR_결정_삼성전자_최종통합전략.md` |
| MAD | MAD | [결정] | Final integrated strategy | `20260205_1956_MAD_결정_삼성전자_최종통합전략.md` |
| MAD | MAD | [가이드] | Practical easy guide | `20260205_1956_MAD_가이드_삼성전자_실전가이드.md` |
| TTO | TTO | [장전] | Pre-market briefing | `20260205_0830_TTO_장전_전체_시장브리핑.md` |

### Directory Structure & File Storage
- **Reports**: Stored in `reports/YYYY-MM-DD/` (from `REPORTS_DIR` env var)
- **Workflows**: `.agent/workflows/*.md`
- **Skills**: `.agent/skills/*/SKILL.md`
- **Config**: `.gemini/` and `.agent/` for MCP settings
- **Scripts**: `scripts/` for utility scripts

### YAML Discipline
- **Pure Korean**: Use only Korean text in YAML description fields
- **No Special Characters**: Avoid colons, brackets, quotes in YAML text fields
- **Minimal Quotes**: Use double quotes only when absolutely necessary (except for tickers, variable names)
- **Clean Parsing**: Prevent YAML parsing errors with careful character usage

### Data & Content Requirements
- **Data Richness**: Include raw data, markdown tables, and absolute numerical values
- **Evidence-Based**: All analysis must include source data URLs and cross-references
- **Global-Domestic Sync**: Correlate Korean data (Naver/DART) with global data (Yahoo/IB reports)
- **Why-First**: Start reports with direct market drivers in 3 sentences or less
- **No Hardcoding**: Use environment variables or relative paths for all file paths

### Error Handling & Validation
- **Session Integrity**: File timestamps must match `ADDITIONAL_METADATA` current local time
- **Future Timestamps Forbidden**: Never use future timestamps in report content
- **Data Cross-Check**: Validate KIS API data against Naver/DART sources
- **Auto-Discovery**: Automatically discover top 3-5 promising stocks when no specific target provided

### Import & Code Conventions
- **Python Scripts**: Use standard Python conventions (PEP 8)
- **MCP Configuration**: JSON format with server command definitions
- **Markdown Files**: Korean content with structured headers and data tables

### Testing & Quality Assurance
- **No Traditional Tests**: Quality is ensured through data validation and report generation
- **Report Validation**: Check for Korean content, proper naming, data completeness
- **Cross-Reference Validation**: Ensure all referenced reports and data sources exist
- **Session Consistency**: Verify all reports in a session use consistent timestamps

## Environment Variables
- `REPORTS_DIR`: Base directory for storing generated reports

## MCP Servers Configuration
Primary data sources configured in `.gemini/mcp_config.json`:
- `naver-finance`: Korean financial data
- `yahoo-finance`: Global market data  
- `kis-trading-remote`: Real-time trading API
- `sequential-thinking`: Advanced reasoning (disabled by default)

## Steering Instructions
1. **Workflows First**: Always reference `.agent/workflows/*.md` files for workflow execution
2. **Korean Language Policy**: Default language is Korean - all outputs must follow this
3. **SRP (Single Responsibility)**: Each expert/skill focuses on its specific domain only
4. **Data Integrity**: Never rely on single data source - always cross-validate
5. **Fact-Based Analysis**: All judgments must be based on verified facts and data