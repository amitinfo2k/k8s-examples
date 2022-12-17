package main

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/onosproject/onos-lib-go/pkg/logging"
    "gorm.io/gorm"
)

var log = logging.GetLogger("model_0_0_0")

type AetherServer struct {
   
   DB *gorm.DB

}
func (w *AetherServer) DeleteConnectivityServices(ctx echo.Context, target Target) error {
    fmt.Println("Target ", target)
    panic("implement me")
}

func (w *AetherServer) GetConnectivityServices(ctx echo.Context, target Target) error {
    fmt.Println("Target ", target)
    log.Infof("GetConnectivityServices")
    panic("implement me")
}

func (w *AetherServer) GetConnectivityServicesConnectivityServiceList(ctx echo.Context, target Target) error {
    panic("implement me")
}

func (w *AetherServer) DeleteConnectivityServicesConnectivityService(ctx echo.Context, target Target, connectivityServiceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetConnectivityServicesConnectivityService(ctx echo.Context, target Target, connectivityServiceId string) error {
    panic("implement me")
}

func (w *AetherServer) PostConnectivityServicesConnectivityService(ctx echo.Context, target Target, connectivityServiceId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprises(ctx echo.Context, target Target) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprises(ctx echo.Context, target Target) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprises(ctx echo.Context, target Target) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseList(ctx echo.Context, target Target) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterprise(ctx echo.Context, target Target, enterpriseId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterprise(ctx echo.Context, target Target, enterpriseId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterprise(ctx echo.Context, target Target, enterpriseId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseApplicationList(ctx echo.Context, target Target, enterpriseId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseApplication(ctx echo.Context, target Target, enterpriseId string, applicationId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseApplication(ctx echo.Context, target Target, enterpriseId string, applicationId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseApplication(ctx echo.Context, target Target, enterpriseId string, applicationId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseApplicationEndpointList(ctx echo.Context, target Target, enterpriseId string, applicationId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseApplicationEndpoint(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseApplicationEndpoint(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseApplicationEndpoint(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseApplicationEndpointMbr(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseApplicationEndpointMbr(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseApplicationEndpointMbr(ctx echo.Context, target Target, enterpriseId string, applicationId string, endpointId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseConnectivityServiceList(ctx echo.Context, target Target, enterpriseId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseConnectivityService(ctx echo.Context, target Target, enterpriseId string, connectivityService string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseConnectivityService(ctx echo.Context, target Target, enterpriseId string, connectivityService string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseConnectivityService(ctx echo.Context, target Target, enterpriseId string, connectivityService string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteList(ctx echo.Context, target Target, enterpriseId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSite(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSite(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSite(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteDeviceList(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteDeviceGroupList(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteDeviceGroupDeviceList(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteDeviceGroupDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string, deviceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteDeviceGroupDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string, deviceId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteDeviceGroupDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string, deviceId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteDeviceGroupMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteDeviceGroupMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteDeviceGroupMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceGroupId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, deviceId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteImsiDefinition(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteImsiDefinition(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteImsiDefinition(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteIpDomainList(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteIpDomain(ctx echo.Context, target Target, enterpriseId string, siteId string, ipDomainId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteIpDomain(ctx echo.Context, target Target, enterpriseId string, siteId string, ipDomainId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteIpDomain(ctx echo.Context, target Target, enterpriseId string, siteId string, ipDomainId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteMonitoring(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteMonitoring(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteMonitoring(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteMonitoringEdgeDeviceList(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteMonitoringEdgeDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, edgeDeviceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteMonitoringEdgeDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, edgeDeviceId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteMonitoringEdgeDevice(ctx echo.Context, target Target, enterpriseId string, siteId string, edgeDeviceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSimCardList(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteSimCard(ctx echo.Context, target Target, enterpriseId string, siteId string, simId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSimCard(ctx echo.Context, target Target, enterpriseId string, siteId string, simId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteSimCard(ctx echo.Context, target Target, enterpriseId string, siteId string, simId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSliceList(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteSlice(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSlice(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteSlice(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSliceDeviceGroupList(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteSliceDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, deviceGroup string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSliceDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, deviceGroup string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteSliceDeviceGroup(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, deviceGroup string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSliceFilterList(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteSliceFilter(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, application string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSliceFilter(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, application string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteSliceFilter(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, application string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteSliceMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSliceMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteSliceMbr(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSlicePriorityTrafficRuleList(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteSlicePriorityTrafficRule(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, priorityTrafficRuleId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSlicePriorityTrafficRule(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, priorityTrafficRuleId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteSlicePriorityTrafficRule(ctx echo.Context, target Target, enterpriseId string, siteId string, sliceId string, priorityTrafficRuleId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSmallCellList(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteSmallCell(ctx echo.Context, target Target, enterpriseId string, siteId string, smallCellId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteSmallCell(ctx echo.Context, target Target, enterpriseId string, siteId string, smallCellId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteSmallCell(ctx echo.Context, target Target, enterpriseId string, siteId string, smallCellId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteUpfList(ctx echo.Context, target Target, enterpriseId string, siteId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseSiteUpf(ctx echo.Context, target Target, enterpriseId string, siteId string, upfId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseSiteUpf(ctx echo.Context, target Target, enterpriseId string, siteId string, upfId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseSiteUpf(ctx echo.Context, target Target, enterpriseId string, siteId string, upfId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseTemplateList(ctx echo.Context, target Target, enterpriseId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseTemplate(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseTemplate(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseTemplate(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseTemplateMbr(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseTemplateMbr(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseTemplateMbr(ctx echo.Context, target Target, enterpriseId string, templateId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseTrafficClassList(ctx echo.Context, target Target, enterpriseId string) error {
    panic("implement me")
}

func (w *AetherServer) DeleteEnterprisesEnterpriseTrafficClass(ctx echo.Context, target Target, enterpriseId string, trafficClassId string) error {
    panic("implement me")
}

func (w *AetherServer) GetEnterprisesEnterpriseTrafficClass(ctx echo.Context, target Target, enterpriseId string, trafficClassId string) error {
    panic("implement me")
}

func (w *AetherServer) PostEnterprisesEnterpriseTrafficClass(ctx echo.Context, target Target, enterpriseId string, trafficClassId string) error {
    panic("implement me")
}

/*func (w *AetherServer) PostConnectivityServices(ctx echo.Context, target Target, enterpriseId string, trafficClassId string) error {
    var err error
    // ------------- Path parameter "target" -------------
    fmt.Println("enterpriseId ", enterpriseId)
    fmt.Println("Target ", target)

    return err
}*/

func (w *AetherServer) PostConnectivityServices(ctx echo.Context, target Target) error {
    fmt.Println("Target ", target)
    panic("implement me")
}

func main() {
    var s = AetherServer{}
    // Echo instance
    echoRouter := echo.New()

    // Middleware
    echoRouter.Use(middleware.Logger())
    echoRouter.Use(middleware.Recover())

    RegisterHandlers(echoRouter, &s)

    // Start server
    echoRouter.Logger.Fatal(echoRouter.Start(":8080"))
}




