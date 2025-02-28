# MarketBill Messaging Service

MarketBill Messaging Service는 비즈니스 크리티컬한 알림 메시지를 안정적으로 전송하기 위한 서버리스 메시징 시스템입니다.

## 기술 스택

- Go 1.19+
- AWS Lambda
- PostgreSQL
- Naver Cloud SENS
- GORM
- AWS API Gateway

## 아키텍처 설계 원칙

### 1. 서버리스 아키텍처 채택 이유

- **비용 효율성**: 메시지 발송은 간헐적으로 발생하는 이벤트이므로, 상시 운영되는 서버 대신 서버리스 아키텍처 채택
- **자동 스케일링**: 대량 메시지 발송 시에도 AWS Lambda의 자동 스케일링으로 안정적인 처리 보장
- **운영 부담 감소**: 인프라 관리 복잡성 최소화

### 2. 계층화된 아키텍처

```
Controller -> Service -> Repository
```

- **관심사의 분리**: 각 계층이 독립적인 책임을 가져 유지보수성 향상
- **테스트 용이성**: 계층별 독립적인 테스트 가능
- **코드 재사용**: 서비스 로직의 모듈화로 재사용성 확보

### 3. 템플릿 기반 메시지 관리

- **일관성 보장**: 정해진 템플릿으로 메시지 포맷 통일
- **유지보수성**: 메시지 내용 변경 시 템플릿만 수정하면 되어 유지보수 용이
- **검증 로직**: 템플릿별 필수 파라미터 검증으로 오류 방지

### 4. 보안 설계

- **환경별 제어**: 개발 환경에서는 특정 전화번호로만 발송 가능하도록 제한
- **시그니처 인증**: HMAC-SHA256 기반의 API 요청 인증
- **민감정보 보호**: 환경변수를 통한 보안 정보 관리

### 5. 안정성 확보

- **로깅 시스템**: 모든 메시지 발송 기록을 DB에 저장하여 추적성 확보
- **예외 처리**: Panic Recovery를 통한 안정적인 에러 처리
- **메시지 타입 자동 전환**: 길이에 따라 SMS/LMS 자동 전환

## 주요 기능

- 기본 SMS 발송
- 템플릿 기반 메시지 발송
- SMS/LMS 자동 전환
- 발송 이력 관리
- 개발 환경 제한적 발송

## 프로젝트 특징

1. **확장성**
   - 새로운 메시지 템플릿 추가가 용이한 구조
   - 다양한 메시징 채널(카카오톡 등) 추가 가능한 인터페이스 설계

2. **유지보수성**
   - 명확한 계층 구조로 코드 이해도 향상
   - 독립적인 모듈로 분리되어 수정 영향도 최소화

3. **안정성**
   - 모든 발송 기록 DB 저장
   - 장애 상황에 대한 예외 처리
   - 개발 환경 보호 장치

4. **비용 효율성**
   - 서버리스 아키텍처로 유휴 리소스 제거
   - 사용량 기반 과금으로 비용 최적화
