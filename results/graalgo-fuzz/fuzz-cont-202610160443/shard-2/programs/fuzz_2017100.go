package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []uint32{42, 721, 65535, 2}
    var v2 uint16 = 104
    v3 := []uint32{42, 65535, 4294967295}
    var v4 float64 = -2.25
    var v5 int64 = 1
    v4 = 1e10
    v2 = v2
    var v6 uint32 = 795
    _ = len(v3)
    v5 = v5
    var v7 int64 = 877
    if (32767 > 359) {
        if ((^(v5)) == -11) {
            if false {
                v3 = append(v3, (v6 | v6))
                v8 := GenS1{F0: -32768, F1: 575, F2: -128}
                v5 -= v7
                _ = v8
            }
            if len(v1) > 0 {
                v1[0] = v6
            }
        } else {
            if true {
                v5 = 0
            } else {
                v9 := GenS1{F0: -32768, F1: 1, F2: -128}
                _ = v9
            }
            v7 = v5
            if true {
                fmt.Println(v6, v6, v4)
            } else {
                var v10 uint = 1
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            v6 -= v6
            v12 := GenS1{F0: -32768, F1: 1, F2: -128}
            if ((0 >= 1000) && false) {
                var v13 uint32 = (v6 << ((10) & 63))
                if false {
                    var v14 int8 = -91
                    _ = v14
                }
                v4 -= (3.14159 * (v4 - v4))
                _ = v13
            } else {
                if len(v3) > 0 {
                    v3[0] = v6
                }
            }
            _ = fi11
            _ = v12
        }
        if len(v1) > 0 {
            v1[0] = v6
        }
    } else {
        if false {
            var v15 float32 = 1.5
            var v16 int32 = 2147483647
            _ = v15
            _ = v16
        } else {
            switch 745 {
            case -1:
                for fi17 := 0; fi17 < 16; fi17++ {
                    switch "hello" {
                    case "foo bar":
                        if ((true && true) && (42 > 0)) {
                            v5 -= v7
                            v2 -= 42
                        }
                    case "\n\t":
                        fmt.Println(v7, v5, v7)
                    case "":
                        if len(v3) > 0 {
                            v3[0] = uint32(v2)
                        }
                    default:
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                    }
                    _ = fi17
                }
            case 9223372036854775807:
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v4 = v4
                        var v20 int64 = v7
                        fmt.Println(v2, v20, v2)
                        _ = fi19
                        _ = v20
                    }
                    v21 := GenS1{F0: -719, F1: 255, F2: 42}
                    _ = fi18
                    _ = v21
                }
            case 4294967295:
                v5 = (v7 | (v7 / ((v7) | 1)))
            }
            v6 = v6
        }
    }
    var v22 uint8 = 128
    for fi23 := 0; fi23 < 15; fi23++ {
        v2 -= 255
        _ = fi23
    }
    if (((-412 < 0) || true) || (v6 == (v6 << ((42) & 63)))) {
        v1 = append(v1, v6)
        if false {
            v7 += v5
        }
        v22 -= (v22 | (^(v22)))
    }
    v7 += v5
    for fi24 := 0; fi24 < 11; fi24++ {
        switch "hello" {
        case "a":
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v22, v6, v2)
            }
        }
        _ = fi24
    }
    if ((v2 / ((1) | 1)) >= v2) {
        if (v22 > v22) {
            v25 := GenS1{F0: 534, F1: 226, F2: -58}
            v5 = 1
            v5 = v5
            _ = v25
        } else {
            for fi26 := 0; fi26 < 3; fi26++ {
                v27 := GenS1{F0: 32767, F1: uint(v6), F2: -128}
                _ = fi26
                _ = v27
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v22)
        }
        var v28 int64 = (v5 ^ v5)
        _ = len(v3)
        _ = v28
    }
    var v29 uint = 1
    if ((true && true) && false) {
        v1 = append(v1, (v6 * (v6 - (^(v6)))))
    } else {
        for fi30 := 0; fi30 < 2; fi30++ {
            switch 35 {
            case 128:
                v3 = append(v3, (v6 & uint32(v29)))
            case 32767:
                for fi31 := 0; fi31 < 6; fi31++ {
                    if true {
                        var v32 float64 = 0.1
                        _ = v32
                    } else {
                        if true {
                            v5 -= v7
                        }
                    }
                    v33 := GenS2{F0: 42}
                    _ = fi31
                    _ = v33
                }
            }
            v22 = (v22 + v22)
            _ = fi30
        }
        v4 = v4
    }
    fmt.Println(v29, v7)
    v4 -= 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
