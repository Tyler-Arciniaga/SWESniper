import React, { useState } from "react";
import { Link } from "react-router-dom";

/**
 * Login Page Component
 * - Handles user login
 * - Matches the visual style of your dashboard
 * - Includes form validation and error handling
 */
const LoginPage = ({ onLogin }) => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState({});
  const [error, setError] = useState("");

  // Handle form input changes
  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
    // Clear error when user starts typing
    if (error) setError("");
  };

  // Handle form submission
  const handleSubmit = async (e) => {
    e.preventDefault();

    // Basic validation
    if (!formData.email || !formData.password) {
      setError("Please fill in all fields");
      return;
    }

    setLoading(true);
    setError("");

    try {
      // TODO: Replace with actual API call in production
      // For now, we'll simulate a login with demo credentials
      //make API call to Supabase auth service
      //NOTE: had to make a seperate env file within root of frontend directory, need to make sure this does not interfere with
      //the extract of env variables in backend go logic
      fetch(
        `https://${process.env.REACT_APP_SUPABASE_PROJECT_REF}.supabase.co/auth/v1/token?grant_type=password`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Connection: "keep-alive",
            apikey: process.env.REACT_APP_SUPABASE_ANON_KEY,
          },
          body: JSON.stringify({
            email: formData.email,
            password: formData.password,
          }),
        }
      )
        .then((response) => {
          if (!response.ok) {
            throw new Error(
              `HTTP error trying to retrieve access token from supabase: ${response.status}`
            );
          }
          console.log(response.body);
          return response.json();
        })
        .then((data) => {
          const accessToken = data.access_token;
          const user = {
            email: formData.email,
          };

          onLogin(user, accessToken);
        });

      /*
        if (
        formData.email === "demo@example.com" &&
        formData.password === "password"
      ) {
        // Simulate successful login
        const userData = {
          id: 1,
          email: formData.email,
          name: "Demo User",
        };
        const token = "demo-jwt-token-" + Date.now();

        onLogin(userData, token);
      } else {
        // For demo purposes, any other credentials will "work"
        const userData = {
          id: Math.floor(Math.random() * 1000),
          email: formData.email,
          name: formData.email.split("@")[0],
        };
        const token = "demo-jwt-token-" + Date.now();

        onLogin(userData, token);
      }
        */
    } catch (err) {
      setError("Login failed. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  const containerStyle = {
    minHeight: "100vh",
    background: "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
    padding: "0",
    margin: "0",
    fontFamily:
      '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif',
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  };

  const contentStyle = {
    width: "100%",
    maxWidth: "450px",
    padding: "20px",
  };

  const cardStyle = {
    background: "rgba(255, 255, 255, 0.95)",
    backdropFilter: "blur(10px)",
    borderRadius: "20px",
    padding: "40px",
    boxShadow: "0 10px 40px rgba(0,0,0,0.2)",
    textAlign: "center",
  };

  const headerStyle = {
    marginBottom: "32px",
  };

  const titleStyle = {
    fontSize: "32px",
    fontWeight: "800",
    margin: "0 0 8px 0",
    color: "#2c3e50",
  };

  const subtitleStyle = {
    fontSize: "16px",
    color: "#6c757d",
    margin: "0",
  };

  const formStyle = {
    textAlign: "left",
  };

  const inputGroupStyle = {
    marginBottom: "20px",
  };

  const labelStyle = {
    display: "block",
    fontSize: "14px",
    fontWeight: "600",
    color: "#2c3e50",
    marginBottom: "8px",
  };

  const inputStyle = {
    width: "100%",
    padding: "14px 16px",
    border: "2px solid #e9ecef",
    borderRadius: "8px",
    fontSize: "16px",
    backgroundColor: "#ffffff",
    transition: "all 0.3s ease",
    boxSizing: "border-box",
  };

  const buttonStyle = {
    width: "100%",
    padding: "14px",
    background: loading
      ? "#6c757d"
      : "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
    color: "white",
    border: "none",
    borderRadius: "8px",
    fontSize: "16px",
    fontWeight: "600",
    cursor: loading ? "not-allowed" : "pointer",
    transition: "all 0.3s ease",
    boxShadow: "0 4px 15px rgba(102, 126, 234, 0.4)",
    marginBottom: "24px",
  };

  const errorStyle = {
    background: "linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%)",
    color: "white",
    padding: "12px 16px",
    borderRadius: "8px",
    marginBottom: "20px",
    fontSize: "14px",
    textAlign: "center",
    boxShadow: "0 4px 15px rgba(255, 107, 107, 0.3)",
  };

  const linkStyle = {
    color: "#667eea",
    textDecoration: "none",
    fontWeight: "600",
  };

  const backLinkStyle = {
    position: "absolute",
    top: "20px",
    left: "20px",
    color: "white",
    textDecoration: "none",
    fontSize: "14px",
    fontWeight: "500",
    display: "flex",
    alignItems: "center",
    padding: "8px 16px",
    background: "rgba(255, 255, 255, 0.1)",
    borderRadius: "20px",
    transition: "all 0.2s ease",
  };

  return (
    <div style={containerStyle}>
      {/* Back to home link */}
      <Link
        to="/"
        style={backLinkStyle}
        onMouseOver={(e) => {
          e.target.style.background = "rgba(255, 255, 255, 0.2)";
          e.target.style.transform = "translateX(-2px)";
        }}
        onMouseOut={(e) => {
          e.target.style.background = "rgba(255, 255, 255, 0.1)";
          e.target.style.transform = "translateX(0)";
        }}
      >
        ← Back to Home
      </Link>

      <div style={contentStyle}>
        <div style={cardStyle}>
          <div style={headerStyle}>
            <h1 style={titleStyle}>🎯 Welcome Back</h1>
            <p style={subtitleStyle}>Sign in to access your dashboard</p>
          </div>

          {error && <div style={errorStyle}>⚠️ {error}</div>}

          <form style={formStyle} onSubmit={handleSubmit}>
            <div style={inputGroupStyle}>
              <label style={labelStyle} htmlFor="email">
                📧 Email Address
              </label>
              <input
                type="email"
                id="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                style={inputStyle}
                placeholder="Enter your email"
                onFocus={(e) => {
                  e.target.style.borderColor = "#667eea";
                  e.target.style.boxShadow =
                    "0 0 0 3px rgba(102, 126, 234, 0.1)";
                }}
                onBlur={(e) => {
                  e.target.style.borderColor = "#e9ecef";
                  e.target.style.boxShadow = "none";
                }}
              />
            </div>

            <div style={inputGroupStyle}>
              <label style={labelStyle} htmlFor="password">
                🔒 Password
              </label>
              <input
                type="password"
                id="password"
                name="password"
                value={formData.password}
                onChange={handleChange}
                style={inputStyle}
                placeholder="Enter your password"
                onFocus={(e) => {
                  e.target.style.borderColor = "#667eea";
                  e.target.style.boxShadow =
                    "0 0 0 3px rgba(102, 126, 234, 0.1)";
                }}
                onBlur={(e) => {
                  e.target.style.borderColor = "#e9ecef";
                  e.target.style.boxShadow = "none";
                }}
              />
            </div>

            <button
              type="submit"
              style={buttonStyle}
              disabled={loading}
              onMouseOver={(e) => {
                if (!loading) {
                  e.target.style.transform = "translateY(-2px)";
                  e.target.style.boxShadow =
                    "0 6px 20px rgba(102, 126, 234, 0.5)";
                }
              }}
              onMouseOut={(e) => {
                if (!loading) {
                  e.target.style.transform = "translateY(0)";
                  e.target.style.boxShadow =
                    "0 4px 15px rgba(102, 126, 234, 0.4)";
                }
              }}
            >
              {loading ? "🔄 Signing In..." : "🚀 Sign In"}
            </button>

            <p style={{ textAlign: "center", color: "#6c757d", margin: "0" }}>
              Don't have an account?{" "}
              <Link
                to="/signup"
                style={linkStyle}
                onMouseOver={(e) =>
                  (e.target.style.textDecoration = "underline")
                }
                onMouseOut={(e) => (e.target.style.textDecoration = "none")}
              >
                Sign up here
              </Link>
            </p>
          </form>

          {/* Demo credentials info */}
          <div
            style={{
              marginTop: "24px",
              padding: "16px",
              backgroundColor: "#f8f9fa",
              borderRadius: "8px",
              fontSize: "12px",
              color: "#6c757d",
              textAlign: "left",
            }}
          >
            <strong>Demo Mode:</strong>
            <br />
            You can use any email/password to login.
            <br />
            Try: demo@example.com / password
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
