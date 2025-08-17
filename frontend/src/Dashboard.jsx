import React, { useState, useEffect } from "react";

// URLCard Component
const URLCard = ({ url, onViewChanges, onDelete }) => {
  const [isHovered, setIsHovered] = useState(false);

  const formatDate = (dateString) => {
    if (!dateString) return "Never";
    return new Date(dateString).toLocaleString();
  };

  const getStatusColor = (lastChecked) => {
    if (!lastChecked) return "#ffc107"; // Yellow for never checked
    const now = new Date();
    const checked = new Date(lastChecked);
    const diffMinutes = (now - checked) / (1000 * 60);

    if (diffMinutes < 30) return "#28a745"; // Green for recent
    if (diffMinutes < 120) return "#ffc107"; // Yellow for moderate
    return "#dc3545"; // Red for old
  };

  const cardStyle = {
    background: "linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%)",
    border: "1px solid #e9ecef",
    borderRadius: "12px",
    padding: "20px",
    marginBottom: "16px",
    boxShadow: isHovered
      ? "0 8px 25px rgba(0,0,0,0.1)"
      : "0 2px 10px rgba(0,0,0,0.05)",
    transition: "all 0.3s ease",
    transform: isHovered ? "translateY(-2px)" : "translateY(0)",
    position: "relative",
    overflow: "hidden",
  };

  const statusIndicatorStyle = {
    position: "absolute",
    top: "16px",
    right: "16px",
    width: "10px",
    height: "10px",
    borderRadius: "50%",
    backgroundColor: getStatusColor(url.lastCheckAt),
    boxShadow: `0 0 0 3px ${getStatusColor(url.lastCheckAt)}20`,
  };

  const titleStyle = {
    margin: "0 0 12px 0",
    color: "#2c3e50",
    fontSize: "18px",
    fontWeight: "600",
    paddingRight: "30px",
  };

  const infoStyle = {
    margin: "6px 0",
    color: "#6c757d",
    fontSize: "14px",
    display: "flex",
    alignItems: "center",
  };

  const buttonContainerStyle = {
    marginTop: "16px",
    display: "flex",
    gap: "10px",
  };

  const baseButtonStyle = {
    padding: "8px 16px",
    border: "none",
    borderRadius: "6px",
    cursor: "pointer",
    fontSize: "14px",
    fontWeight: "500",
    transition: "all 0.2s ease",
    boxShadow: "0 2px 4px rgba(0,0,0,0.1)",
  };

  const viewButtonStyle = {
    ...baseButtonStyle,
    background: "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
    color: "white",
  };

  const deleteButtonStyle = {
    ...baseButtonStyle,
    background: "linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%)",
    color: "white",
  };

  return (
    <div
      style={cardStyle}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      <div style={statusIndicatorStyle}></div>

      <h3 style={titleStyle}>{url.description}</h3>

      <div style={infoStyle}>
        <span style={{ marginRight: "8px" }}>🔗</span>
        <span
          style={{
            fontFamily: "monospace",
            backgroundColor: "#f8f9fa",
            padding: "2px 6px",
            borderRadius: "4px",
          }}
        >
          {url.url}
        </span>
      </div>

      <div style={infoStyle}>
        <span style={{ marginRight: "8px" }}>⏰</span>
        <span>Last checked: {formatDate(url.lastCheckAt)}</span>
      </div>

      <div style={infoStyle}>
        <span style={{ marginRight: "8px" }}>🔄</span>
        <span>Every {url.checkInterval} seconds</span>
      </div>

      <div style={buttonContainerStyle}>
        <button
          style={viewButtonStyle}
          onClick={() => onViewChanges(url)}
          onMouseOver={(e) => (e.target.style.transform = "scale(1.05)")}
          onMouseOut={(e) => (e.target.style.transform = "scale(1)")}
        >
          📊 View Changes
        </button>
        <button
          style={deleteButtonStyle}
          onClick={() => onDelete(url.id)}
          onMouseOver={(e) => (e.target.style.transform = "scale(1.05)")}
          onMouseOut={(e) => (e.target.style.transform = "scale(1)")}
        >
          🗑️ Delete
        </button>
      </div>
    </div>
  );
};

// AddURLForm Component
const AddURLForm = ({ onAddURL, isLoading }) => {
  const [formData, setFormData] = useState({
    url: "",
    description: "",
    checkInterval: 60,
  });
  const [isFocused, setIsFocused] = useState(false);

  const handleSubmit = () => {
    if (!formData.url || !formData.description) {
      alert("Please fill in all required fields");
      return;
    }
    onAddURL(formData);
    setFormData({ url: "", description: "", checkInterval: 60 });
  };

  const formStyle = {
    background: "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
    borderRadius: "16px",
    padding: "24px",
    marginBottom: "32px",
    boxShadow: "0 10px 30px rgba(102, 126, 234, 0.3)",
    color: "white",
  };

  const inputStyle = {
    width: "100%",
    padding: "12px 16px",
    marginBottom: "16px",
    border: "none",
    borderRadius: "8px",
    fontSize: "14px",
    backgroundColor: "rgba(255, 255, 255, 0.95)",
    backdropFilter: "blur(10px)",
    boxShadow: "0 2px 10px rgba(0,0,0,0.1)",
    transition: "all 0.3s ease",
    boxSizing: "border-box",
  };

  const buttonStyle = {
    padding: "12px 24px",
    background: isLoading
      ? "#6c757d"
      : "linear-gradient(135deg, #28a745 0%, #20c997 100%)",
    color: "white",
    border: "none",
    borderRadius: "8px",
    cursor: isLoading ? "not-allowed" : "pointer",
    fontSize: "16px",
    fontWeight: "600",
    transition: "all 0.3s ease",
    boxShadow: "0 4px 15px rgba(40, 167, 69, 0.4)",
    transform: isLoading ? "scale(1)" : "scale(1.02)",
  };

  return (
    <div style={formStyle}>
      <h2
        style={{
          margin: "0 0 20px 0",
          fontSize: "24px",
          fontWeight: "700",
          textAlign: "center",
        }}
      >
        ✨ Add New URL to Track
      </h2>

      <input
        type="url"
        placeholder="🔗 URL (e.g., https://github.com/user/repo)"
        value={formData.url}
        onChange={(e) => setFormData({ ...formData, url: e.target.value })}
        style={inputStyle}
        onFocus={() => setIsFocused(true)}
        onBlur={() => setIsFocused(false)}
      />

      <input
        type="text"
        placeholder="📝 Description (e.g., 'Summer 2025 Internships')"
        value={formData.description}
        onChange={(e) =>
          setFormData({ ...formData, description: e.target.value })
        }
        style={inputStyle}
      />

      <input
        type="number"
        placeholder="⏱️ Check interval in seconds"
        value={formData.checkInterval}
        onChange={(e) =>
          setFormData({ ...formData, checkInterval: parseInt(e.target.value) })
        }
        style={inputStyle}
        min="1"
      />

      <div style={{ textAlign: "center" }}>
        <button
          onClick={handleSubmit}
          style={buttonStyle}
          disabled={isLoading}
          onMouseOver={(e) =>
            !isLoading && (e.target.style.transform = "scale(1.05)")
          }
          onMouseOut={(e) =>
            !isLoading && (e.target.style.transform = "scale(1.02)")
          }
        >
          {isLoading ? "🔄 Adding..." : "🚀 Add URL"}
        </button>
      </div>
    </div>
  );
};

// ChangeLogModal Component
const ChangeLogModal = ({ url, changes, onClose }) => {
  if (!url) return null;

  const modalOverlayStyle = {
    position: "fixed",
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: "rgba(0,0,0,0.6)",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    zIndex: 1000,
    backdropFilter: "blur(5px)",
  };

  const modalStyle = {
    background: "linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%)",
    padding: "32px",
    borderRadius: "20px",
    maxWidth: "85vw",
    maxHeight: "85vh",
    overflow: "auto",
    position: "relative",
    boxShadow: "0 20px 60px rgba(0,0,0,0.3)",
    animation: "modalSlideIn 0.3s ease",
  };

  const closeButtonStyle = {
    position: "absolute",
    top: "20px",
    right: "25px",
    background: "linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%)",
    border: "none",
    fontSize: "18px",
    cursor: "pointer",
    color: "white",
    width: "35px",
    height: "35px",
    borderRadius: "50%",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    transition: "all 0.2s ease",
    boxShadow: "0 2px 10px rgba(255, 107, 107, 0.3)",
  };

  const changeStyle = {
    border: "1px solid #e9ecef",
    borderRadius: "10px",
    padding: "16px",
    marginBottom: "16px",
    background: "linear-gradient(135deg, #f8f9fa 0%, #ffffff 100%)",
    boxShadow: "0 2px 8px rgba(0,0,0,0.05)",
    position: "relative",
    borderLeft: "4px solid #667eea",
  };

  const headerStyle = {
    margin: "0 0 24px 0",
    color: "#2c3e50",
    fontSize: "28px",
    fontWeight: "700",
    paddingRight: "50px",
  };

  if (url.lastKnownHash === "INVALID_COULD_NOT_EXTRACT_WEB_CONTENT") {
    return (
      <div style={modalOverlayStyle} onClick={onClose}>
        <div style={modalStyle} onClick={(e) => e.stopPropagation()}>
          <button style={closeButtonStyle} onClick={onClose}>
            ×
          </button>
          <h2 style={headerStyle}>📈 Change Log: {url.description}</h2>

          <div
            style={{
              textAlign: "center",
              padding: "40px 20px",
              color: "#6c757d",
            }}
          >
            <div style={{ fontSize: "48px", marginBottom: "16px" }}>😔</div>
            <p style={{ fontSize: "18px", margin: "0" }}>
              Something went wrong with the checking of this url.
            </p>
            <p style={{ fontSize: "14px", margin: "8px 0 0 0" }}>
              I would love to hear your feedback, feel free to email me
              @tyarciniaga@gmail.com with a screenshot and url for me to fix
              this bug ASAP :)
            </p>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div style={modalOverlayStyle} onClick={onClose}>
      <div style={modalStyle} onClick={(e) => e.stopPropagation()}>
        <style>{`
          @keyframes modalSlideIn {
            from { transform: translateY(-20px); opacity: 0; }
            to { transform: translateY(0); opacity: 1; }
          }
        `}</style>

        <button
          style={closeButtonStyle}
          onClick={onClose}
          onMouseOver={(e) => (e.target.style.transform = "scale(1.1)")}
          onMouseOut={(e) => (e.target.style.transform = "scale(1)")}
        >
          ×
        </button>

        <h2 style={headerStyle}>📈 Change Log: {url.description}</h2>

        <div
          style={{
            backgroundColor: "#f8f9fa",
            padding: "12px 16px",
            borderRadius: "8px",
            marginBottom: "24px",
            border: "1px solid #e9ecef",
          }}
        >
          <p style={{ margin: "0", color: "#6c757d", fontSize: "14px" }}>
            <strong>🔗 URL:</strong>{" "}
            <span style={{ fontFamily: "monospace" }}>{url.url}</span>
          </p>
        </div>

        {!changes || changes.length === 0 ? (
          <div
            style={{
              textAlign: "center",
              padding: "40px 20px",
              color: "#6c757d",
            }}
          >
            <div style={{ fontSize: "48px", marginBottom: "16px" }}>🔍</div>
            <p style={{ fontSize: "18px", margin: "0" }}>
              No changes detected yet.
            </p>
            <p style={{ fontSize: "14px", margin: "8px 0 0 0" }}>
              We'll notify you when something changes!
            </p>
          </div>
        ) : (
          <div>
            <h3
              style={{
                margin: "0 0 20px 0",
                color: "#2c3e50",
                fontSize: "20px",
                fontWeight: "600",
              }}
            >
              📊 Recent Changes ({changes.length}):
            </h3>
            {changes
              .slice()
              .reverse()
              .map((change, index) => (
                <div key={index} style={changeStyle}>
                  <div
                    style={{
                      position: "absolute",
                      top: "16px",
                      right: "16px",
                      backgroundColor: "#667eea",
                      color: "white",
                      padding: "4px 8px",
                      borderRadius: "12px",
                      fontSize: "12px",
                      fontWeight: "500",
                    }}
                  >
                    #{changes.length - index}
                  </div>

                  <p
                    style={{
                      margin: "0 0 12px 0",
                      fontSize: "14px",
                      color: "#6c757d",
                      display: "flex",
                      alignItems: "center",
                    }}
                  >
                    <span style={{ marginRight: "8px" }}>🕐</span>
                    <strong>Detected:</strong>
                    <span style={{ marginLeft: "8px" }}>
                      {new Date(change.timestamp).toLocaleString()}
                    </span>
                  </p>

                  <p
                    style={{
                      margin: "0 0 12px 0",
                      fontSize: "15px",
                      color: "#2c3e50",
                      display: "flex",
                      alignItems: "center",
                    }}
                  >
                    <span style={{ marginRight: "8px" }}>⚡</span>
                    <strong>Changes:</strong>
                    <span style={{ marginLeft: "8px" }}>
                      {change.diffsummary || "Content modified"}
                    </span>
                  </p>

                  {change.added && (
                    <div style={{ marginTop: "12px" }}>
                      <p
                        style={{
                          margin: "0 0 8px 0",
                          fontSize: "13px",
                          color: "#6c757d",
                          fontWeight: "500",
                        }}
                      >
                        📝 Details:
                      </p>
                      <div
                        style={{
                          backgroundColor: "#2c3e50",
                          color: "#ffffff",
                          padding: "12px",
                          borderRadius: "6px",
                          fontSize: "12px",
                          overflow: "auto",
                          fontFamily: "Consolas, Monaco, monospace",
                          lineHeight: "1.4",
                        }}
                      >
                        {change.added.map((job, i) => (
                          <div key={i} style={{ marginBottom: "12px" }}>
                            {job.fields?.map((field, j) => (
                              <div key={j}>{field}</div>
                            ))}
                          </div>
                        ))}
                      </div>
                    </div>
                  )}
                </div>
              ))}
          </div>
        )}

        {url.lastKnownContent && (
          <div
            style={{
              marginTop: "32px",
              borderTop: "2px solid #e9ecef",
              paddingTop: "24px",
            }}
          >
            <h3
              style={{
                margin: "0 0 16px 0",
                color: "#2c3e50",
                fontSize: "20px",
                fontWeight: "600",
              }}
            >
              📄 Last Known Content:
            </h3>
            <div
              style={{
                backgroundColor: "#2c3e50",
                color: "#ffffff",
                padding: "16px",
                borderRadius: "8px",
                fontSize: "12px",
                overflow: "auto",
                maxHeight: "250px",
                fontFamily: "Consolas, Monaco, monospace",
                lineHeight: "1.4",
              }}
            >
              {url.lastKnownContent.map((job, index) => (
                <div key={index} style={{ marginBottom: "12px" }}>
                  {job.fields?.map((field, i) => (
                    <div key={i}>{field}</div>
                  ))}
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

// Main App Component
const Dashboard = ({ onLogout }) => {
  const [urls, setUrls] = useState([]);
  const [loading, setLoading] = useState(true);
  const [addingURL, setAddingURL] = useState(false);
  const [selectedURL, setSelectedURL] = useState(null);
  const [changes, setChanges] = useState([]);
  const [error, setError] = useState("");

  // TODO: Replace with actual backend URL in PROD
  const API_BASE_URL = "http://localhost:8080";
  const accessToken = localStorage.getItem("authToken");
  if (accessToken == null) {
    onLogout();
  }
  const user = localStorage.getItem("userData");

  // Fetch all tracked URLs
  const fetchURLs = async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/urls`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      });
      if (!response.ok) throw new Error("Failed to fetch URLs");
      const data = await response.json();
      setUrls(data || []);
    } catch (err) {
      setError("Failed to load URLs: " + err.message);
    } finally {
      setLoading(false);
    }
  };

  // Add a new URL
  const addURL = async (urlData) => {
    setAddingURL(true);
    try {
      const response = await fetch(`${API_BASE_URL}/urls`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
        body: JSON.stringify(urlData),
      });

      if (!response.ok) throw new Error("Failed to add URL");

      await fetchURLs(); // Refresh the list
      setError("");
    } catch (err) {
      setError("Failed to add URL: " + err.message);
    } finally {
      setAddingURL(false);
    }
  };

  // Delete a URL
  const deleteURL = async (id) => {
    if (!window.confirm("Are you sure you want to delete this URL?")) return;

    try {
      const response = await fetch(`${API_BASE_URL}/urls/${id}`, {
        method: "DELETE",
      });

      if (!response.ok) throw new Error("Failed to delete URL");

      await fetchURLs(); // Refresh the list
      setError("");
    } catch (err) {
      setError("Failed to delete URL: " + err.message);
    }
  };

  // View changes for a URL
  const viewChanges = async (url) => {
    try {
      const response = await fetch(`${API_BASE_URL}/changelog/${url.id}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      });
      if (!response.ok) throw new Error("Failed to fetch changes");
      const data = await response.json();
      setChanges(data);
      setSelectedURL(url);
    } catch (err) {
      setError("Failed to load changes: " + err.message);
    }
  };

  // Handle logout
  const handleLogout = () => {
    if (window.confirm("Are you sure you want to logout?")) {
      onLogout();
    }
  };

  // Load URLs on component mount
  useEffect(() => {
    fetchURLs();
  }, []);

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

  const headerStyle = {
    textAlign: "center",
    color: "white",
    marginBottom: "40px",
    background: "rgba(255, 255, 255, 0.1)",
    backdropFilter: "blur(10px)",
    borderRadius: "20px",
    padding: "32px",
    boxShadow: "0 8px 32px rgba(0,0,0,0.1)",
    position: "relative",
  };

  const logoutButtonStyle = {
    position: "absolute",
    top: "20px",
    right: "20px",
    background: "rgba(255, 255, 255, 0.2)",
    border: "2px solid rgba(255, 255, 255, 0.3)",
    borderRadius: "12px",
    color: "white",
    padding: "10px 20px",
    fontSize: "14px",
    fontWeight: "600",
    cursor: "pointer",
    backdropFilter: "blur(10px)",
    transition: "all 0.3s ease",
    display: "flex",
    alignItems: "center",
    gap: "8px",
  };

  const mainTitleStyle = {
    fontSize: "48px",
    fontWeight: "800",
    margin: "0 0 16px 0",
    background: "linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%)",
    WebkitBackgroundClip: "text",
    textShadow: "0 2px 4px rgba(0,0,0,0.1)",
  };

  const subtitleStyle = {
    fontSize: "18px",
    margin: "0",
    opacity: "0.9",
    fontWeight: "400",
  };

  const dashboardStyle = {
    background: "rgba(255, 255, 255, 0.95)",
    backdropFilter: "blur(10px)",
    borderRadius: "20px",
    padding: "32px",
    boxShadow: "0 10px 40px rgba(0,0,0,0.1)",
  };

  const errorStyle = {
    background: "linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%)",
    color: "white",
    padding: "16px 20px",
    borderRadius: "10px",
    marginBottom: "20px",
    boxShadow: "0 4px 20px rgba(255, 107, 107, 0.3)",
    display: "flex",
    alignItems: "center",
  };

  if (loading) {
    return (
      <div style={containerStyle}>
        <div style={contentStyle}>
          <div
            style={{
              textAlign: "center",
              padding: "100px 20px",
              color: "white",
            }}
          >
            <div style={{ fontSize: "48px", marginBottom: "20px" }}>🔄</div>
            <p style={{ fontSize: "24px", margin: "0", fontWeight: "600" }}>
              Loading...
            </p>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div style={containerStyle}>
      <div style={contentStyle}>
        <div style={headerStyle}>
          {/* Logout Button */}
          <button
            onClick={handleLogout}
            style={logoutButtonStyle}
            onMouseEnter={(e) => {
              e.target.style.background = "rgba(255, 255, 255, 0.3)";
              e.target.style.transform = "translateY(-2px)";
            }}
            onMouseLeave={(e) => {
              e.target.style.background = "rgba(255, 255, 255, 0.2)";
              e.target.style.transform = "translateY(0)";
            }}
          >
            <span>🚪</span>
            Logout
          </button>

          <h1 style={mainTitleStyle}>🎯 SWE Sniper</h1>
          <p style={subtitleStyle}>
            Precision job tracking and change detection for early-bird SWE
            internship hunters.
          </p>
        </div>

        {error && (
          <div style={errorStyle}>
            <span style={{ marginRight: "12px", fontSize: "20px" }}>⚠️</span>
            {error}
          </div>
        )}

        <AddURLForm onAddURL={addURL} isLoading={addingURL} />

        <div style={dashboardStyle}>
          <h2
            style={{
              color: "#2c3e50",
              marginBottom: "24px",
              fontSize: "24px",
              fontWeight: "700",
              display: "flex",
              alignItems: "center",
            }}
          >
            <span style={{ marginRight: "12px" }}>📊</span>
            Tracked URLs ({urls.length})
          </h2>

          {urls.length === 0 ? (
            <div
              style={{
                textAlign: "center",
                padding: "60px 20px",
                color: "#6c757d",
              }}
            >
              <div style={{ fontSize: "64px", marginBottom: "20px" }}>🔍</div>
              <p
                style={{
                  fontSize: "20px",
                  margin: "0 0 8px 0",
                  fontWeight: "600",
                }}
              >
                No URLs being tracked yet
              </p>
              <p style={{ fontSize: "16px", margin: "0", opacity: "0.8" }}>
                Add one above to get started!
              </p>
            </div>
          ) : (
            urls.map((url) => (
              <URLCard
                key={url.id}
                url={url}
                onViewChanges={viewChanges}
                onDelete={deleteURL}
              />
            ))
          )}
        </div>

        <ChangeLogModal
          url={selectedURL}
          changes={changes}
          onClose={() => setSelectedURL(null)}
        />
      </div>
    </div>
  );
};

export default Dashboard;
