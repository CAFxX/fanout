package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int64; F2 int }
func genFn0(p0 int16, p1 string) int32 {
    switch 0 {
    case 9223372036854775807:
        fmt.Println(p1, p1)
    }
    p1 = (p1 + "")
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 12; fi5++ {
            p1 = (p1 + "foo bar")
            p1 = "foo bar"
            _ = fi5
        }
        switch "foo bar" {
        case "\n\t":
            v6 := GenS1{F0: 42}
            _ = v6
        case "foo bar":
            var v7 int16 = 42
            _ = v7
        }
        _ = fi4
    }
    switch p1 {
    case "foo bar":
        var v8 float32 = -2.25
        _ = v8
    case "a":
        var v9 uint16 = 969
        _ = v9
    case "":
        var v10 uint32 = 2147483648
        _ = v10
    default:
        if (false || false) {
            v11 := GenS1{F0: 65535}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 15; fi12++ {
                switch 343 {
                case -234:
                    var v13 uint32 = 2147483647
                    _ = v13
                case -1:
                    if false {
                        p1 = ""
                    }
                case 4294967295:
                    if false {
                        if true {
                            if (!false) {
                                p1 = "foo bar"
                                switch (fi12 ^ (fi12 | fi12)) {
                                case 862:
                                    fmt.Println(p0, fi12)
                                case 976:
                                    p1 = "hello"
                                }
                            }
                        }
                    } else {
                        if ((!((false && (0 < 847)) || true)) && (0 > 0)) {
                            v14 := []int{128, fi12, (fi12 << ((1) & 63))}
                            _ = len(v14)
                            _ = v14
                        }
                        fmt.Println(fi12, p1)
                    }
                }
                _ = fi12
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 uint16 = 559
                    fmt.Println(p0, fi16, v17)
                    _ = fi16
                    _ = v17
                }
                p0 += (p0 - 42)
                if true {
                    switch fi15 {
                    case 2147483647:
                        if (-2147483648 >= -2147483648) {
                            var v18 int64 = 42
                            v19 := GenS2{F0: (939 < 2), F1: 42, F2: 128}
                            _ = v18
                            _ = v19
                        }
                    case 716:
                        if false {
                            if (!true) {
                                if false {
                                    p0 = 388
                                    switch (p1 + p1) {
                                    case "a":
                                        v20 := GenS1{F0: 245}
                                        _ = v20
                                    default:
                                        var v21 uint32 = 4294967295
                                        _ = v21
                                    }
                                    var v22 int16 = (p0 >> ((4294967295) & 63))
                                    _ = v22
                                }
                                var v23 float64 = 3.14159
                                if true {
                                    for fi24 := 0; fi24 < 2; fi24++ {
                                        var v25 uint = 4294967295
                                        if (false && false) {
                                            v23 += (v23 * (v23 * 3.14159))
                                            fmt.Println(p0)
                                        }
                                        _ = fi24
                                        _ = v25
                                    }
                                } else {
                                    var v26 uint16 = 255
                                    if false {
                                        var v27 int64 = 9223372036854775807
                                        _ = v27
                                    }
                                    _ = v26
                                }
                                _ = v23
                            }
                            v28 := []int{(fi15 - fi15), 0}
                            v29 := GenS2{F0: ((1e10 * -2.25) == (-2.25 / (((0.1 - 1.5) / 1e10) / float32(-2.25)))), F1: 255, F2: 2}
                            _ = v28
                            _ = v29
                        } else {
                            var v30 uint = 503
                            var v31 uint64 = 674
                            if (245 > 74) {
                                p0 = (p0 - (p0 << ((v30) & 63)))
                            }
                            _ = v30
                            _ = v31
                        }
                    case 127:
                        v32 := []int{fi15}
                        _ = v32
                    }
                } else {
                    v33 := []int{(fi15 << ((2) & 63)), (fi15 + (fi15 - (fi15 & fi15)))}
                    _ = v33
                }
                _ = fi15
            }
        }
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int8) uint8 {
    var v34 string = "hello"
    if true {
        switch v34 {
        case "hello":
            p0 += -1
        case "":
            var v35 uint16 = 750
            _ = v35
        case "a":
            for fi36 := 0; fi36 < 13; fi36++ {
                var v37 bool = true
                _ = fi36
                _ = v37
            }
        default:
            if false {
                p0 = (p0 - (^(p0)))
            } else {
                switch 704 {
                case -2147483648:
                    v34 = ""
                }
            }
        }
    }
    if false {
        p0 += (p0 >> ((255) & 63))
        var v38 bool = true
        _ = v38
    }
    fmt.Println(v34, p0)
    _ = p0
    _ = v34
    return 0
}

func main() {
    var v1 int = 0
    v2 := []string{"", "foo bar", "a"}
    var v3 uint16 = 1
    v39 := GenS1{F0: 165}
    if false {
        v1 = (v1 & (v1 << ((4294967295) & 63)))
    } else {
        var v40 int8 = 100
        if (-234 < -1) {
            v40 = v40
            var v41 uint = 65535
            _ = v41
        }
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        _ = v40
    }
    v42 := GenS3{F0: (v3 + (v3 + (^(v3))))}
    v43 := GenS2{F0: true, F1: 255, F2: 190}
    fmt.Println(v1)
    if true {
        v1 += (v1 + 32768)
    }
    fmt.Println(v3, v3)
    v44 := GenS2{F0: (!false), F1: 96, F2: 0}
    fmt.Println(v1)
    for fi45 := 0; fi45 < 11; fi45++ {
        var v46 int8 = 42
        if false {
            switch v1 {
            case -9223372036854775808:
                v47 := GenS4{F0: v46}
                _ = v47
            case 703:
                v48 := GenS2{F0: (3.14159 >= -2.25), F1: 42, F2: 95}
                _ = v48
            default:
                v49 := GenS4{F0: v46}
                _ = v49
            }
        }
        _ = fi45
        _ = v46
    }
    var v50 int = v1
    var v51 uint32 = 65535
    v1 = v50
    switch "foo bar" {
    case "a":
        switch (v50 & (v50 & v50)) {
        case -2147483648:
            v52 := GenS1{F0: 1}
            _ = v52
        }
    default:
        switch (("foo bar" + "hello") + "a") {
        case "\n\t":
            v53 := GenS5{F0: -1}
            _ = v53
        case "hello":
            v1 = v1
        }
    }
    v1 = v1
    switch (v1 & -537) {
    case 127:
        v3 = 65535
    }
    var v54 int = -589
    for fi55 := 0; fi55 < 18; fi55++ {
        v50 = v1
        var v56 int16 = 0
        if false {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            if (1e10 == 0.1) {
                v57 := GenS6{F0: true, F1: 42, F2: fi55}
                _ = v57
            } else {
                v2 = append(v2, "a")
            }
        }
        _ = fi55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v42
    _ = v43
    _ = v44
    _ = v50
    _ = v51
    _ = v54
}
