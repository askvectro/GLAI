# 🚀 GLAI: AI-Powered Go Assistant (Web + iOS)  
**Build, debug, and share Go code from anywhere** – on your phone or browser.  
*"Write once, run anywhere" for the Go ecosystem*  

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)  
[![Discord](https://img.shields.io/badge/Join-Discord-blue)](https://discord.gg/your-invite-link)  

<div align="center">  
  <img src="https://example.com/gla-screenshot.png" width="800" alt="Web/iOS Screenshot">  
</div>  

---

## ✨ Features  
| **Web App**                      | **iOS App**                     | **Shared Core**                |  
|----------------------------------|---------------------------------|--------------------------------|  
| Real-time collaborative editing  | Voice-to-code with Siri Shortcuts| GPT-4 powered code generation  |  
| Browser-based Go IDE             | AR code visualization (ARKit)   | Cross-platform snippet sync    |  
| Chrome DevTools integration      | Offline mode support            | Community voting/forks         |  
| Export to Git repos              | Camera-driven UML diagramming   | AI-driven debugging assistant  |  

---

## 🛠️ Tech Stack  
**Shared Core**  
- Go 1.21+ • Yaegi (embedded runtime) • Redis • Firebase Realtime DB  
- AI/ML: GPT-4o API + fine-tuned CodeLlama-34b (Go specialization)  

**Web Frontend**  
- Next.js 14 (TypeScript) • Tailwind CSS • WebAssembly (TinyGo)  
- Monaco Editor • WebSocket (real-time sync)  

**iOS Frontend**  
- SwiftUI • CoreML (on-device models) • ARKit • CoreData  
- CGo bindings for shared Go modules  

**DevOps**  
- AWS ECS/Fargate • Terraform • GitHub Actions  
- Prometheus/Grafana monitoring • Fastlane (iOS CI/CD)  

---

## 🚀 Getting Started  

### Prerequisites  
- Go 1.21+ • Node 18+ • Xcode 15+ (for iOS)  

### Installation  
```bash  
# Web App  
git clone https://github.com/vectro-ai/gla-web  
cd gla-web && npm install  

# iOS App  
git clone https://github.com/vectro-ai/gla-ios  
open gla-ios/GLAI.xcodeproj  

# Shared Core  
git clone https://github.com/vectro-ai/gla-core  
cd gla-core && go mod download
