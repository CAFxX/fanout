package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }

func main() {
    var v1 string = (("" + ("foo bar" + (("hello" + "foo bar") + "a"))) + "")
    var v2 int64 = -9223372036854775808
    var v3 int64 = (v2 + (v2 / (((v2 << ((374) & 63))) | 1)))
    var v4 uint64 = 2
    var v5 uint64 = (v4 / ((v4) | 1))
    var v6 string = v1
    fmt.Println(v6, v2, v5)
    if true {
        var v7 float64 = 1e10
        v3 -= 42
        _ = v7
    } else {
        v8 := GenS1{F0: 2147483647, F1: 1, F2: 602}
        var v9 int32 = -1
        _ = v8
        _ = v9
    }
    v10 := []int{int(v5), 4294967295}
    v2 = -4
    if len(v10) > 0 {
        v10[0] = 32768
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: 2, F1: -111, F2: 65535}
        v6 = v1
        if false {
            v10 = append(v10, 32768)
        }
        _ = fi11
        _ = v12
    }
    v3 -= v3
    var v13 int64 = 9223372036854775807
    switch 0 {
    case -63:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case -1:
        var v14 int = -1
        _ = v14
    case 93:
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                v2 = 0
                var v17 int8 = 0
                v13 = (v3 * (v2 / (((v13 - -9223372036854775808)) | 1)))
                if true {
                    v17 = (-(v17))
                    fmt.Println(v6, v1, v5)
                    v2 = (v3 << ((2) & 63))
                }
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS1{F0: 937, F1: -128, F2: 0}
                var v20 uint = 7
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi15
        }
    default:
        switch 42 {
        case 128:
            if true {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
        case 2147483647:
            if true {
                for fi21 := 0; fi21 < 18; fi21++ {
                    fmt.Println(v1, v5)
                    fmt.Println(v5)
                    var v22 int64 = v2
                    _ = fi21
                    _ = v22
                }
            } else {
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int8 = 41
                    var v25 uint16 = 0
                    for fi26 := 0; fi26 < 6; fi26++ {
                        switch -9223372036854775808 {
                        case 1:
                            if (4294967295 <= 1) {
                                var v27 bool = true
                                _ = v27
                            }
                        default:
                            var v28 uint8 = 127
                            _ = v28
                        }
                        _ = fi26
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                var v29 uint32 = 955
                _ = v29
            }
        default:
            v13 -= -296
        }
    }
    v30 := GenS1{F0: 4294967295, F1: 42, F2: 862}
    for fi31 := 0; fi31 < 2; fi31++ {
        if false {
            fmt.Println(v6, v13, v4)
            var v32 uint64 = (v5 / ((v5) | 1))
            for fi33 := 0; fi33 < 16; fi33++ {
                v1 = v6
                _ = fi33
            }
            _ = v32
        } else {
            var v34 uint32 = 2147483648
            if true {
                var v35 float32 = 0.1
                fmt.Println(v4)
                switch 9223372036854775807 {
                case 579:
                    v36 := GenS1{F0: uint32(v4), F1: -3, F2: 482}
                    _ = v36
                default:
                    if len(v10) > 0 {
                        v10[0] = (^(fi31))
                    }
                }
                _ = v35
            }
            switch (fi31 & 127) {
            case -247:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case -8:
                var v37 uint16 = 42
                _ = v37
            default:
                _ = len(v10)
            }
            _ = v34
        }
        var v38 uint64 = v5
        _ = fi31
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
