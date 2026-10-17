package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 bool }
func genFn0() int8 {
    var v4 uint32 = 743
    v4 -= 65535
    _ = v4
    return 42
}

func main() {
    v1 := []uint32{4294967295, 174, 495, 1}
    var v2 bool = false
    var v3 uint32 = 990
    if v2 {
        v2 = (3.14159 >= ((0.0 - 3.14159) - ((3.14159 / -2.25) + (-2.25 / 3.14159))))
    }
    var v5 int8 = 0
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint16 = 43
        switch (fi6 - 127) {
        case -682:
            fmt.Println(v5, v2)
        case -527:
            for fi8 := 0; fi8 < 13; fi8++ {
                _ = len(v1)
                v3 += v3
                v5 -= genFn0()
                _ = fi8
            }
        case -736:
            v7 -= (v7 & v7)
        default:
            var v9 int16 = 32767
            _ = v9
        }
        v5 = (v5 << ((891) & 63))
        _ = fi6
        _ = v7
    }
    var v10 uint = 124
    fmt.Println(v10, v3, v10)
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: 247, F1: 32767, F2: -862}
        for fi13 := 0; fi13 < 4; fi13++ {
            fmt.Println(v3, v3)
            fmt.Println(v2, v5, fi11)
            v14 := GenS1{F0: uint8(v3), F1: 237, F2: -646}
            fmt.Println(fi11, v5, v2)
            var v15 string = "foo bar"
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    v16 := GenS1{F0: 5, F1: 65535, F2: 1}
    if v2 {
        v17 := GenS1{F0: 2, F1: 32767, F2: 1}
        _ = v17
    } else {
        v18 := GenS1{F0: 2, F1: 255, F2: 2147483647}
        if len(v1) > 0 {
            v1[0] = v3
        }
        _ = v18
    }
    v19 := GenS1{F0: 233, F1: 32767, F2: int32(v10)}
    fmt.Println(v5, v5)
    var v20 uint8 = 164
    v2 = true
    for fi21 := 0; fi21 < 6; fi21++ {
        switch fi21 {
        case 2147483647:
            v3 = 4294967295
        default:
            for fi22 := 0; fi22 < 12; fi22++ {
                if (0 > (^(v10))) {
                    switch (fi21 + fi22) {
                    case 987:
                        v20 = (v20 << (((v10 << ((2) & 63))) & 63))
                    }
                }
                _ = fi22
            }
        }
        if (0.0 != 1e10) {
            if (!((1 >= 1000) || (1 < 1))) {
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 float32 = (0.1 * 0.1)
                    var v25 int16 = -402
                    var v26 uint32 = (v3 ^ v3)
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                var v27 uint16 = 0
                _ = v27
            } else {
                var v28 uint64 = 1
                _ = v28
            }
        } else {
            v29 := GenS1{F0: v20, F1: uint16(v3), F2: 1}
            v3 = (v3 >> ((42) & 63))
            var v30 uint16 = 255
            _ = v29
            _ = v30
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v31 int64 = -9223372036854775808
    _ = len(v1)
    if v2 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        fmt.Println(v20, v2, v3)
        switch ("foo bar" + ("foo bar" + "hello")) {
        case "a":
            switch -980 {
            case 4294967295:
                switch 2147483647 {
                case 32767:
                    v31 = 1
                }
            case 42:
                v32 := GenS1{F0: v20, F1: 0, F2: 1}
                _ = v32
            case 127:
                if v2 {
                    var v33 uint16 = 255
                    switch int(v20) {
                    case 32768:
                        v3 = v3
                    case 4294967295:
                        for fi34 := 0; fi34 < 1; fi34++ {
                            v31 += 42
                            _ = fi34
                        }
                    case 128:
                        _ = len(v1)
                    }
                    _ = v33
                }
            default:
                var v35 bool = (-128 < (v5 * 100))
                _ = v35
            }
        case "":
            v36 := GenS2{F0: v3, F1: v5}
            _ = v36
        }
        switch -279 {
        case -1:
            if v2 {
                if (0.1 < (0.1 - (0.0 * -2.25))) {
                    if ((-696 < v31) || true) {
                        var v37 int = 42
                        v38 := GenS2{F0: v3, F1: (v5 - -1)}
                        fmt.Println(v2, v31, v31)
                        _ = v37
                        _ = v38
                    }
                    var v39 string = "hello"
                    _ = len(v1)
                    _ = v39
                }
            } else {
                var v40 int16 = -177
                var v41 float64 = 0.1
                _ = v40
                _ = v41
            }
        case 32768:
            switch "foo bar" {
            case "":
                fmt.Println(v3)
            case "\n\t":
                switch "\n\t" {
                case "a":
                    v42 := GenS2{F0: v3, F1: v5}
                    _ = v42
                default:
                    fmt.Println(v3, v20, v31)
                }
            case "a":
                if len(v1) > 0 {
                    v1[0] = (v3 % (((v3 >> ((uint(v20)) & 63))) | 1))
                }
            }
        case 2147483647:
            fmt.Println(v31, v3, v5)
        default:
            if len(v1) > 0 {
                v1[0] = v3
            }
        }
    }
    v43 := GenS3{F0: (1e10 == 0.0)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v43
    _ = v5
}
