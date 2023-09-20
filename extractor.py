from itertools import zip_longest
import re
import json

# Read HTML content from 'companies-table.txt'
with open('companies-table.txt', 'r', encoding='utf-8') as f:
    html_string = f.read()


# Regular expression patterns
company_url_pattern = r'<a role="link" class="component--field-formatter accent ng-star-inserted"[^>]*href="(/organization/[^"]+)"'
company_name_pattern = r'<div hoveroverelement="" class="identifier-label">\s*([^<]+)\s*</div>'
total_funding_amount_pattern = r'<a onboardinghighlight="" class="component--field-formatter field-type-money[^>]*>([^<]+)</a>'
cb_rank_organization_pattern = r'<a onboardinghighlight="" class="component--field-formatter field-type-integer[^>]*>([^<]+)</a>'
trend_score_30_days_pattern = r'<span class="component--field-formatter field-type-decimal[^>]*title="([^"]+)"'

# Find all instances
company_urls = re.findall(company_url_pattern, html_string)
company_names = re.findall(company_name_pattern, html_string)
total_funding_amounts = re.findall(total_funding_amount_pattern, html_string)
cb_rank_organizations = re.findall(cb_rank_organization_pattern, html_string)
trend_score_30_days = re.findall(trend_score_30_days_pattern, html_string)

# Create JSON object
companies = []

# Loop through all the lists, filling in None for missing values
for url, name, funding, rank, trend in zip_longest(
    company_urls, company_names, total_funding_amounts, 
    cb_rank_organizations, trend_score_30_days):

    # Initialize missing values to '0' or some default
    funding = 'N/A' if funding is None else funding

    company = {
        "company_name": name,
        "cb_rank_organization": rank,
        "company_url": "https://www.crunchbase.com" + url,
        "total_funding_amount": funding,
        "trend_score_30_days": trend
    }

    companies.append(company)

# Save JSON object to a file in human-readable format
with open("data/companies.json", "w") as f:
    json.dump(companies, f, indent=2)