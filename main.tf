provider "assembla"{
}


resource "assembla_webhookbasic" "webhook45"{
    space_name="terraformprovider"
    title="webhook45"
    http_method=0
    content_type="application/json"
    post_about_scrum_reports="1"
    external_url="https://krunalshimpi.com/configure?param1=value2"
}
