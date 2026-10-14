package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 int32 }
func genFn0(p0 uint32, p1 bool) bool {
    switch 0 {
    case 555:
        var v6 uint16 = 65535
        _ = v6
    }
    var v7 int8 = 0
    fmt.Println(v7, p1, p0)
    _ = p0
    _ = p1
    _ = v7
    return (!p1)
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 206
    v3 := []int{1, 2147483647, 0}
    var v4 float32 = 0.1
    var v5 bool = false
    switch "a" {
    case "":
        var v8 float32 = 0.1
        _ = v8
    }
    v1 = v1
    v9 := GenS1{F0: 0, F1: 168, F2: (!false)}
    v4 += 1e10
    var v10 uint64 = 1
    if (("\n\t" != "foo bar") && v5) {
        if v5 {
            var v11 float64 = (1e10 + 3.14159)
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int64 = 136
                _ = fi12
                _ = v13
            }
            fmt.Println(v10, v10)
            _ = v11
        }
        v10 += (v10 & 1)
    } else {
        v2 = 86
        switch -26 {
        case 0:
            switch ("foo bar" + "") {
            case "":
                var v14 int64 = 9223372036854775807
                _ = v14
            default:
                switch 1 {
                case 905:
                    if (-376 > 0) {
                        fmt.Println(v2)
                        var v15 uint32 = 65535
                        _ = v15
                    } else {
                        v3 = append(v3, 9223372036854775807)
                        if v5 {
                            v1 = v1
                        } else {
                            for fi16 := 0; fi16 < 10; fi16++ {
                                if (!(v5 && false)) {
                                    fmt.Println(v10, v10, v5)
                                    for fi17 := 0; fi17 < 20; fi17++ {
                                        var v18 uint16 = 778
                                        var v19 int8 = 1
                                        if len(v3) > 0 {
                                            v3[0] = fi16
                                        }
                                        v18 = 42
                                        _ = fi17
                                        _ = v18
                                        _ = v19
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    v4 = -2.25
                                }
                                _ = fi16
                            }
                            var v20 int64 = -9223372036854775808
                            v21 := GenS1{F0: 1, F1: 0, F2: false}
                            _ = v20
                            _ = v21
                        }
                        if v5 {
                            for fi22 := 0; fi22 < 3; fi22++ {
                                if (v5 || (42 != -82)) {
                                    if ((v1 / (((v1 ^ v1)) | 1)) < (v1 / ((285) | 1))) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v1 += v1
                                        switch "hello" {
                                        case "a":
                                            for fi23 := 0; fi23 < 11; fi23++ {
                                                fmt.Println(v4, v10)
                                                var v24 int64 = -1
                                                var v25 string = "hello"
                                                _ = fi23
                                                _ = v24
                                                _ = v25
                                            }
                                        case "\n\t":
                                            var v26 int = (fi22 << ((65535) & 63))
                                            _ = v26
                                        default:
                                            switch "\n\t" {
                                            case "a":
                                                v1 -= v1
                                            case "":
                                                v1 += v1
                                            default:
                                                v3 = append(v3, -2147483648)
                                            }
                                        }
                                    }
                                    v27 := GenS1{F0: 42, F1: 42, F2: (65535 >= 4294967295)}
                                    _ = v27
                                }
                                v28 := GenS2{F0: v2, F1: false, F2: 205}
                                var v29 uint = 42
                                _ = fi22
                                _ = v28
                                _ = v29
                            }
                        }
                    }
                }
            }
        }
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v5 = ((0 == 42) || (!(true && (!v5))))
        _ = fi30
    }
    var v31 float32 = v4
    if true {
        for fi32 := 0; fi32 < 1; fi32++ {
            for fi33 := 0; fi33 < 6; fi33++ {
                if true {
                    var v34 uint16 = (v2 << ((1) & 63))
                    if len(v3) > 0 {
                        v3[0] = (fi32 >> ((4294967295) & 63))
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v35 uint64 = (v10 + (v10 - (v10 << ((255) & 63))))
                        _ = v35
                    default:
                        v1 = (-(v1))
                    }
                    _ = v34
                }
                var v36 bool = (v5 || true)
                var v37 uint64 = (v10 << ((4294967295) & 63))
                _ = fi33
                _ = v36
                _ = v37
            }
            if true {
                v1 = 32767
                v2 -= (v2 << ((18446744073709551615) & 63))
            }
            _ = fi32
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 float32 = -2.25
            v40 := GenS3{F0: "\n\t", F1: (-(v1)), F2: 0}
            _ = fi38
            _ = v39
            _ = v40
        }
    } else {
        var v41 int32 = 392
        v2 = uint16(v10)
        v1 += (v1 >> ((967) & 63))
        _ = v41
    }
    fmt.Println(v31, v4, v5)
    for fi42 := 0; fi42 < 9; fi42++ {
        var v43 int32 = 1
        v44 := GenS3{F0: (("\n\t" + "\n\t") + ("\n\t" + "foo bar")), F1: 1, F2: v43}
        _ = fi42
        _ = v43
        _ = v44
    }
    var v45 uint = 150
    fmt.Println(v31)
    var v46 uint64 = (v10 | uint64(v2))
    var v47 uint = 65535
    var v48 int64 = 0
    for fi49 := 0; fi49 < 3; fi49++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
        var v50 uint16 = v2
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v9
}
