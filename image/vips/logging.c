#include "logging.h"

static void fyntrix_logging_handler(const gchar *log_domain,
                                   GLogLevelFlags log_level,
                                   const gchar *message, gpointer user_data) {
 fyntrixLoggingHandler((char *)log_domain, (int)log_level, (char *)message);
}

static void null_logging_handler(const gchar *log_domain,
                                 GLogLevelFlags log_level, const gchar *message,
                                 gpointer user_data) {}

void vips_set_logging_handler(void) {
  g_log_set_default_handler(fyntrix_logging_handler, NULL);
}

void vips_default_logging_handler(void) {
  g_log_set_default_handler(g_log_default_handler, NULL);
}
