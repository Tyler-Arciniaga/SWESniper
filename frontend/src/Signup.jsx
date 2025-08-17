import React, { useState } from "react";
import { Link } from "react-router-dom";

const Signup = ({ onSignup }) => {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    password: "",
    confirmPassword: "",
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");
  const [isHovered, setIsHovered] = useState(false);

  //TODO: change in final PROD!!!
  const API_BASE_URL = "http://localhost:8080";

  // Handle form submission
  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError("");

    // Basic validation
    if (
      !formData.name ||
      !formData.email ||
      !formData.password ||
      !formData.confirmPassword
    ) {
      setError("Please fill in all fields");
      setIsLoading(false);
      return;
    }

    if (formData.password !== formData.confirmPassword) {
      setError("Passwords do not match");
      setIsLoading(false);
      return;
    }

    if (formData.password.length < 6) {
      setError("Password must be at least 6 characters long");
      setIsLoading(false);
      return;
    }

    // Email validation
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) {
      setError("Please enter a valid email address");
      setIsLoading(false);
      return;
    }

    try {
      // Call backend signup API
      const response = await fetch(`${API_BASE_URL}/signup`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: formData.email,
          password: formData.password,
          email_confirm: true,
        }),
      });

      if (!response.ok) {
        const errData = await response.json();
        console.log(errData);
        throw new Error(errData?.message || "Failed to sign up user");
      }

      const userData = await response.json();
      onSignup(userData); // Only call after successful signup
    } catch (err) {
      setError(
        "Failed to sign up user: " +
          "perhaps account already exists with given email"
      );
    } finally {
      setIsLoading(false);
    }
  };

  // Handle input changes
  const handleInputChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
    // Clear error when user starts typing
    if (error) setError("");
  };

  // Styles matching your dashboard theme
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
    maxWidth: "450px",
    width: "100%",
    margin: "0 auto",
    padding: "20px",
  };

  const formContainerStyle = {
    background: "rgba(255, 255, 255, 0.95)",
    backdropFilter: "blur(10px)",
    borderRadius: "20px",
    padding: "40px",
    boxShadow: "0 10px 40px rgba(0,0,0,0.1)",
  };

  const headerStyle = {
    textAlign: "center",
    marginBottom: "32px",
  };

  const titleStyle = {
    fontSize: "32px",
    fontWeight: "700",
    color: "#2c3e50",
    margin: "0 0 8px 0",
  };

  const subtitleStyle = {
    fontSize: "16px",
    color: "#6c757d",
    margin: "0",
  };

  const formStyle = {
    width: "100%",
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
    padding: "12px 16px",
    border: "1px solid #e9ecef",
    borderRadius: "8px",
    fontSize: "16px",
    backgroundColor: "#ffffff",
    transition: "all 0.3s ease",
    boxSizing: "border-box",
  };

  const inputFocusStyle = {
    ...inputStyle,
    borderColor: "#667eea",
    boxShadow: "0 0 0 3px rgba(102, 126, 234, 0.1)",
    outline: "none",
  };

  const errorStyle = {
    background: "linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%)",
    color: "white",
    padding: "12px 16px",
    borderRadius: "8px",
    marginBottom: "20px",
    fontSize: "14px",
    display: "flex",
    alignItems: "center",
  };

  const buttonStyle = {
    width: "100%",
    padding: "14px",
    background: isLoading
      ? "#6c757d"
      : isHovered
      ? "linear-gradient(135deg, #5a6fd8 0%, #6b5b95 100%)"
      : "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
    color: "white",
    border: "none",
    borderRadius: "8px",
    fontSize: "16px",
    fontWeight: "600",
    cursor: isLoading ? "not-allowed" : "pointer",
    transition: "all 0.3s ease",
    transform: isHovered && !isLoading ? "translateY(-1px)" : "translateY(0)",
    boxShadow:
      isHovered && !isLoading
        ? "0 6px 20px rgba(102, 126, 234, 0.4)"
        : "0 2px 10px rgba(102, 126, 234, 0.2)",
  };

  const linkContainerStyle = {
    textAlign: "center",
    marginTop: "24px",
    padding: "16px",
    background: "rgba(102, 126, 234, 0.1)",
    borderRadius: "8px",
  };

  const linkTextStyle = {
    fontSize: "14px",
    color: "#6c757d",
    margin: "0",
  };

  const linkStyle = {
    color: "#667eea",
    textDecoration: "none",
    fontWeight: "600",
    transition: "color 0.2s ease",
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
        <div style={formContainerStyle}>
          {/* Header */}
          <div style={headerStyle}>
            <h1 style={titleStyle}>🚀 Join SWE Sniper</h1>
            <p style={subtitleStyle}>
              Create your account and start tracking opportunities
            </p>
          </div>

          {/* Error message */}
          {error && (
            <div style={errorStyle}>
              <span style={{ marginRight: "8px" }}>⚠️</span>
              {error}
            </div>
          )}

          {/* Signup Form */}
          <form onSubmit={handleSubmit} style={formStyle}>
            {/* Name Input */}
            <div style={inputGroupStyle}>
              <label htmlFor="name" style={labelStyle}>
                Full Name
              </label>
              <input
                type="text"
                id="name"
                name="name"
                value={formData.name}
                onChange={handleInputChange}
                style={inputStyle}
                onFocus={(e) => Object.assign(e.target.style, inputFocusStyle)}
                onBlur={(e) => Object.assign(e.target.style, inputStyle)}
                placeholder="Enter your full name"
                disabled={isLoading}
              />
            </div>

            {/* Email Input */}
            <div style={inputGroupStyle}>
              <label htmlFor="email" style={labelStyle}>
                Email Address
              </label>
              <input
                type="email"
                id="email"
                name="email"
                value={formData.email}
                onChange={handleInputChange}
                style={inputStyle}
                onFocus={(e) => Object.assign(e.target.style, inputFocusStyle)}
                onBlur={(e) => Object.assign(e.target.style, inputStyle)}
                placeholder="Enter your email address"
                disabled={isLoading}
              />
            </div>

            {/* Password Input */}
            <div style={inputGroupStyle}>
              <label htmlFor="password" style={labelStyle}>
                Password
              </label>
              <input
                type="password"
                id="password"
                name="password"
                value={formData.password}
                onChange={handleInputChange}
                style={inputStyle}
                onFocus={(e) => Object.assign(e.target.style, inputFocusStyle)}
                onBlur={(e) => Object.assign(e.target.style, inputStyle)}
                placeholder="Create a password (min. 6 characters)"
                disabled={isLoading}
              />
            </div>

            {/* Confirm Password Input */}
            <div style={inputGroupStyle}>
              <label htmlFor="confirmPassword" style={labelStyle}>
                Confirm Password
              </label>
              <input
                type="password"
                id="confirmPassword"
                name="confirmPassword"
                value={formData.confirmPassword}
                onChange={handleInputChange}
                style={inputStyle}
                onFocus={(e) => Object.assign(e.target.style, inputFocusStyle)}
                onBlur={(e) => Object.assign(e.target.style, inputStyle)}
                placeholder="Confirm your password"
                disabled={isLoading}
              />
            </div>

            {/* Submit Button */}
            <button
              type="submit"
              style={buttonStyle}
              disabled={isLoading}
              onMouseEnter={() => setIsHovered(true)}
              onMouseLeave={() => setIsHovered(false)}
            >
              {isLoading ? "🔄 Creating Account..." : "🎯 Create Account"}
            </button>
          </form>

          {/* Login Link */}
          <div style={linkContainerStyle}>
            <p style={linkTextStyle}>
              Already have an account?{" "}
              <Link
                to="/login"
                style={linkStyle}
                onMouseOver={(e) => (e.target.style.color = "#5a6fd8")}
                onMouseOut={(e) => (e.target.style.color = "#667eea")}
              >
                Sign in here
              </Link>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Signup;
