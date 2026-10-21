package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
func genFn0(p0 int16, p1 string) string {
    var v4 bool = true
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 = (p0 >> ((1) & 63))
        var v6 int8 = 0
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{(fi7 >> ((856) & 63)), fi7, (fi7 & fi7), fi7}
        _ = fi7
        _ = v8
    }
    switch (p1 + ("" + ((p1 + "a") + "hello"))) {
    case "hello":
        var v9 float32 = float32(1e10)
        _ = v9
    case "":
        var v10 uint64 = 2
        _ = v10
    case "a":
        if (-1 < -128) {
            if (2147483647 < 173) {
                for fi11 := 0; fi11 < 4; fi11++ {
                    v4 = v4
                    if v4 {
                        var v12 string = p1
                        _ = v12
                    }
                    _ = fi11
                }
            } else {
                if v4 {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v14 := []int{}
                        var v15 uint8 = 173
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 6; fi16++ {
                        p1 = (p1 + "hello")
                        var v17 int64 = 0
                        switch p1 {
                        case "":
                            for fi18 := 0; fi18 < 4; fi18++ {
                                var v19 uint = 255
                                v20 := GenS1{F0: 139}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                        case "a":
                            switch p1 {
                            case "hello":
                                v21 := []int{fi16, (^(fi16)), fi16}
                                _ = v21
                            case "\n\t":
                                var v22 float32 = (float32((-2.25 / 1.0)) + (3.14159 + 0.0))
                                _ = v22
                            case "":
                                v23 := GenS2{F0: -59}
                                _ = v23
                            default:
                                v24 := []int{fi16, 738}
                                _ = v24
                            }
                        default:
                            for fi25 := 0; fi25 < 13; fi25++ {
                                v26 := []int{fi16, (fi16 << ((65535) & 63)), fi16}
                                v27 := GenS3{F0: -2147483648, F1: 1, F2: 417}
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                    if true {
                        if v4 {
                            fmt.Println(p0, p1, v4)
                            v4 = (0 != 634)
                            v28 := []int{4294967295, 393, 32767}
                            _ = v28
                        }
                        var v29 int64 = -1
                        switch 42 {
                        case 0:
                            v30 := GenS3{F0: 0, F1: 42, F2: 18446744073709551615}
                            _ = v30
                        case -925:
                            v31 := []int{128}
                            _ = v31
                        case -816:
                            if (!(2 != 2147483648)) {
                                if (2147483648 == 4294967295) {
                                    fmt.Println(p0, p1)
                                } else {
                                    fmt.Println(v29)
                                    fmt.Println(v29, v4)
                                    switch 0 {
                                    case 32768:
                                        var v32 string = p1
                                        _ = v32
                                    default:
                                        for fi33 := 0; fi33 < 1; fi33++ {
                                            v29 = -890
                                            v29 = v29
                                            fmt.Println(p0, fi33)
                                            _ = fi33
                                        }
                                    }
                                }
                            }
                        }
                        _ = v29
                    } else {
                        for fi34 := 0; fi34 < 2; fi34++ {
                            v35 := []int{(fi34 - fi34), fi34, fi34, fi34}
                            if len(v35) > 0 {
                                v35[0] = fi34
                            }
                            _ = fi34
                            _ = v35
                        }
                    }
                } else {
                    var v36 uint8 = 85
                    _ = v36
                }
            }
            v37 := []int{}
            fmt.Println(p0, v4, p0)
            _ = v37
        } else {
            var v38 uint = 65535
            _ = v38
        }
    default:
        if v4 {
            for fi39 := 0; fi39 < 16; fi39++ {
                var v40 int = (^(fi39))
                var v41 uint64 = 1
                _ = fi39
                _ = v40
                _ = v41
            }
            var v42 uint64 = 634
            fmt.Println(p0)
            _ = v42
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return "\n\t"
}

func main() {
    var v1 uint = 1
    var v2 float64 = 0.1
    var v3 uint8 = 190
    var v43 float64 = ((v2 - v2) - 3.14159)
    for fi44 := 0; fi44 < 16; fi44++ {
        var v45 uint8 = v3
        _ = fi44
        _ = v45
    }
    v46 := GenS2{F0: 0}
    if true {
        if ((false && false) && (v1 < (v1 ^ (v1 + (v1 | v1))))) {
            v47 := []int{}
            for fi48 := 0; fi48 < 13; fi48++ {
                switch "foo bar" {
                case "":
                    fmt.Println(v2)
                case "\n\t":
                    v47 = append(v47, (-(fi48)))
                default:
                    var v49 string = "hello"
                    _ = v49
                }
                _ = fi48
            }
            _ = v47
        } else {
            v50 := []int{36, 32767, -1}
            fmt.Println(v43)
            var v51 float32 = (1.5 - (((float32(v43) / ((1e10 * ((-2.25 * float32(v43)) + 0.0)) / -2.25)) - 3.14159) / 1.0))
            _ = v50
            _ = v51
        }
        v52 := []int{}
        if (100 != -27) {
            v43 -= v2
            var v53 int = -1
            var v54 uint16 = 267
            _ = v53
            _ = v54
        } else {
            for fi55 := 0; fi55 < 4; fi55++ {
                var v56 string = "a"
                _ = fi55
                _ = v56
            }
            var v57 uint8 = (v3 << (((v1 ^ v1)) & 63))
            _ = v57
        }
        _ = v52
    } else {
        var v58 float64 = (1e10 * v43)
        fmt.Println(v3)
        _ = v58
    }
    v59 := []int{859, 1, 128, 4294967295}
    var v60 float64 = v43
    v61 := GenS2{F0: 1}
    v62 := GenS4{F0: 0, F1: (v3 * 42)}
    for fi63 := 0; fi63 < 13; fi63++ {
        v43 = (v43 + 1.5)
        if true {
            v64 := GenS5{F0: -2147483648}
            _ = v64
        } else {
            switch (fi63 % (((fi63 + (fi63 & (-(fi63))))) | 1)) {
            case 32768:
                v3 = 127
            default:
                switch ("foo bar" + "foo bar") {
                case "hello":
                    for fi65 := 0; fi65 < 19; fi65++ {
                        var v66 float32 = float32(0.1)
                        _ = len(v59)
                        _ = fi65
                        _ = v66
                    }
                case "a":
                    v1 -= 42
                case "foo bar":
                    v67 := GenS4{F0: 32767, F1: (v3 % ((v3) | 1))}
                    _ = v67
                default:
                    if (false && (!true)) {
                        fmt.Println(v60, v2, v1)
                        if true {
                            v60 -= v60
                            if ((fi63 % (((fi63 * 32768)) | 1)) != (-(fi63))) {
                                var v68 int16 = -32768
                                _ = v68
                            }
                        }
                    }
                }
            }
            fmt.Println(v2)
            v69 := GenS6{F0: false}
            _ = v69
        }
        _ = fi63
    }
    v70 := GenS7{F0: 100}
    v3 = (v3 + 87)
    var v71 uint8 = (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v60:", v60)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v46
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v70
    _ = v71
}
