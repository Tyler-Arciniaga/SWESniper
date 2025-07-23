import React from "react";
import { Link } from "react-router-dom";

/**
 * Landing Page Component
 * - First page users see when visiting the app
 * - Provides navigation to login/signup pages
 * - Showcases app features and value proposition
 * - Matches the style of your existing dashboard
 */
const LandingPage = () => {
  const containerStyle = {
    minHeight: "100vh",
    background: "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
    padding: "0",
    margin: "0",
    fontFamily:
      '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif',
  };

  const contentStyle = {
    maxWidth: "900px",
    margin: "0 auto",
    padding: "40px 20px",
    minHeight: "100vh",
    position: "relative",
  };

  const navStyle = {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
    marginBottom: "60px",
    background: "rgba(255, 255, 255, 0.1)",
    backdropFilter: "blur(10px)",
    borderRadius: "16px",
    padding: "20px 32px",
    boxShadow: "0 8px 32px rgba(0,0,0,0.1)",
  };

  const logoStyle = {
    color: "white",
    fontSize: "24px",
    fontWeight: "800",
    margin: "0",
  };

  const navButtonsStyle = {
    display: "flex",
    gap: "16px",
  };

  const baseButtonStyle = {
    padding: "12px 24px",
    borderRadius: "8px",
    textDecoration: "none",
    fontSize: "16px",
    fontWeight: "600",
    transition: "all 0.3s ease",
    border: "none",
    cursor: "pointer",
    boxShadow: "0 4px 15px rgba(0,0,0,0.2)",
  };

  const loginButtonStyle = {
    ...baseButtonStyle,
    background: "rgba(255, 255, 255, 0.2)",
    color: "white",
    border: "2px solid rgba(255, 255, 255, 0.3)",
  };

  const signupButtonStyle = {
    ...baseButtonStyle,
    background: "linear-gradient(135deg, #28a745 0%, #20c997 100%)",
    color: "white",
  };

  const heroStyle = {
    textAlign: "center",
    color: "white",
    marginBottom: "60px",
    background: "rgba(255, 255, 255, 0.1)",
    backdropFilter: "blur(10px)",
    borderRadius: "20px",
    padding: "60px 32px",
    boxShadow: "0 8px 32px rgba(0,0,0,0.1)",
  };

  const mainTitleStyle = {
    fontSize: "64px",
    fontWeight: "800",
    margin: "0 0 24px 0",
    background: "linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%)",
    WebkitBackgroundClip: "text",
    textShadow: "0 2px 4px rgba(0,0,0,0.1)",
  };

  const subtitleStyle = {
    fontSize: "24px",
    margin: "0 0 32px 0",
    opacity: "0.9",
    fontWeight: "400",
    lineHeight: "1.4",
  };

  const ctaButtonStyle = {
    ...baseButtonStyle,
    background: "linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%)",
    color: "white",
    fontSize: "20px",
    padding: "16px 32px",
    textDecoration: "none",
    display: "inline-block",
    marginTop: "16px",
  };

  const featuresStyle = {
    background: "rgba(255, 255, 255, 0.95)",
    backdropFilter: "blur(10px)",
    borderRadius: "20px",
    padding: "40px",
    boxShadow: "0 10px 40px rgba(0,0,0,0.1)",
  };

  const featuresGridStyle = {
    display: "grid",
    gridTemplateColumns: "repeat(auto-fit, minmax(250px, 1fr))",
    gap: "24px",
    marginTop: "32px",
  };

  const featureCardStyle = {
    background: "linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%)",
    border: "1px solid #e9ecef",
    borderRadius: "12px",
    padding: "24px",
    textAlign: "center",
    boxShadow: "0 2px 10px rgba(0,0,0,0.05)",
    transition: "all 0.3s ease",
  };

  const featureIconStyle = {
    fontSize: "48px",
    marginBottom: "16px",
    display: "block",
  };

  const featureTitleStyle = {
    color: "#2c3e50",
    fontSize: "20px",
    fontWeight: "700",
    marginBottom: "12px",
    margin: "0 0 12px 0",
  };

  const featureDescStyle = {
    color: "#6c757d",
    fontSize: "14px",
    margin: "0",
    lineHeight: "1.5",
  };

  return (
    <div style={containerStyle}>
      <div style={contentStyle}>
        {/* Navigation */}
        <nav style={navStyle}>
          <h1 style={logoStyle}>🎯 SWE Sniper</h1>
          <div style={navButtonsStyle}>
            <Link
              to="/login"
              style={loginButtonStyle}
              onMouseOver={(e) => {
                e.target.style.background = "rgba(255, 255, 255, 0.3)";
                e.target.style.transform = "translateY(-2px)";
              }}
              onMouseOut={(e) => {
                e.target.style.background = "rgba(255, 255, 255, 0.2)";
                e.target.style.transform = "translateY(0)";
              }}
            >
              Login
            </Link>
            <Link
              to="/signup"
              style={signupButtonStyle}
              onMouseOver={(e) => {
                e.target.style.transform = "translateY(-2px)";
                e.target.style.boxShadow = "0 6px 20px rgba(40, 167, 69, 0.4)";
              }}
              onMouseOut={(e) => {
                e.target.style.transform = "translateY(0)";
                e.target.style.boxShadow = "0 4px 15px rgba(0,0,0,0.2)";
              }}
            >
              Sign Up
            </Link>
          </div>
        </nav>

        {/* Hero Section */}
        <div style={heroStyle}>
          <h1 style={mainTitleStyle}>🎯 SWE Sniper</h1>
          <p style={subtitleStyle}>
            Precision job tracking and change detection for early-bird SWE
            internship hunters.
          </p>
          <p style={{ fontSize: "18px", margin: "0 0 24px 0", opacity: "0.8" }}>
            Never miss an internship opportunity again. Get instant
            notifications when new positions open up.
          </p>
          <Link
            to="/signup"
            style={ctaButtonStyle}
            onMouseOver={(e) => {
              e.target.style.transform = "translateY(-2px)";
              e.target.style.boxShadow = "0 8px 25px rgba(255, 107, 107, 0.4)";
            }}
            onMouseOut={(e) => {
              e.target.style.transform = "translateY(0)";
              e.target.style.boxShadow = "0 4px 15px rgba(0,0,0,0.2)";
            }}
          >
            🚀 Start Tracking Jobs
          </Link>
        </div>

        {/* Features Section */}
        <div style={featuresStyle}>
          <h2
            style={{
              color: "#2c3e50",
              fontSize: "32px",
              fontWeight: "700",
              textAlign: "center",
              margin: "0 0 16px 0",
            }}
          >
            ✨ Why Choose SWE Sniper? ✨
          </h2>
          <p
            style={{
              color: "#6c757d",
              fontSize: "18px",
              textAlign: "center",
              margin: "0",
            }}
          >
            Built by a CS student for CS student. Trust me I know the struggle
            ...
          </p>

          <div style={featuresGridStyle}>
            <div
              style={featureCardStyle}
              onMouseOver={(e) => {
                e.currentTarget.style.transform = "translateY(-4px)";
                e.currentTarget.style.boxShadow = "0 8px 25px rgba(0,0,0,0.1)";
              }}
              onMouseOut={(e) => {
                e.currentTarget.style.transform = "translateY(0)";
                e.currentTarget.style.boxShadow = "0 2px 10px rgba(0,0,0,0.05)";
              }}
            >
              <span style={featureIconStyle}>⚡</span>
              <h3 style={featureTitleStyle}>Real-time Monitoring</h3>
              <p style={featureDescStyle}>
                Track job boards 24/7 and get alerts instantly or on your
                schedule.
              </p>
            </div>

            <div
              style={featureCardStyle}
              onMouseOver={(e) => {
                e.currentTarget.style.transform = "translateY(-4px)";
                e.currentTarget.style.boxShadow = "0 8px 25px rgba(0,0,0,0.1)";
              }}
              onMouseOut={(e) => {
                e.currentTarget.style.transform = "translateY(0)";
                e.currentTarget.style.boxShadow = "0 2px 10px rgba(0,0,0,0.05)";
              }}
            >
              <span style={featureIconStyle}>🎯</span>
              <h3 style={featureTitleStyle}>Precision Targeting</h3>
              <p style={featureDescStyle}>
                Focus on exactly what you want, from specific companies to role
                types.
              </p>
            </div>

            <div
              style={featureCardStyle}
              onMouseOver={(e) => {
                e.currentTarget.style.transform = "translateY(-4px)";
                e.currentTarget.style.boxShadow = "0 8px 25px rgba(0,0,0,0.1)";
              }}
              onMouseOut={(e) => {
                e.currentTarget.style.transform = "translateY(0)";
                e.currentTarget.style.boxShadow = "0 2px 10px rgba(0,0,0,0.05)";
              }}
            >
              <span style={featureIconStyle}>📊</span>
              <h3 style={featureTitleStyle}>Smart Analytics</h3>
              <p style={featureDescStyle}>
                See exactly what changed, when it changed, and never miss a
                detail.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LandingPage;
