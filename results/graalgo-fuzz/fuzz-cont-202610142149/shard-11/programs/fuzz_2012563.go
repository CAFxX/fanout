package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int8; F1 int; F2 uint }
func genFn0(p0 int) int32 {
    v7 := []int{-1}
    var v8 int = (-(p0))
    v8 += (v8 / ((p0) | 1))
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            v8 -= (fi9 - 127)
            _ = fi10
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 string = ""
    v2 := make(map[int]string)
    var v3 int16 = 42
    var v4 float64 = (1.5 + 3.14159)
    v5 := []string{}
    var v6 uint64 = 4294967295
    if (255 != 128) {
        var v11 uint16 = 0
        switch 32767 {
        case -112:
            v12 := GenS1{F0: ""}
            _ = v12
        case -1:
            var v13 int8 = 100
            _ = v13
        default:
            var v14 uint64 = v6
            _ = v14
        }
        v6 = (^(v6))
        _ = v11
    } else {
        switch 32768 {
        case -2147483648:
            _ = len(v5)
        case 225:
            delete(v2, 37)
        case 127:
            if true {
                v15 := GenS1{F0: ("hello" + "a")}
                _ = v15
            } else {
                fmt.Println(v1, v6)
                for fi16 := 0; fi16 < 14; fi16++ {
                    v2[int(v3)] = v1
                    _ = fi16
                }
                v17 := GenS2{F0: 255}
                _ = v17
            }
        }
        fmt.Println(v1, v4, v4)
        v18 := GenS1{F0: "a"}
        _ = v18
    }
    var v19 uint32 = 65535
    v19 = v19
    v19 -= (v19 & v19)
    v5 = append(v5, "hello")
    v3 = -32768
    switch 0 {
    case -2147483648:
        var v20 float32 = float32((1e10 / 3.14159))
        _ = v20
    default:
        v21 := GenS1{F0: v1}
        _ = v21
    }
    if (!(false && (-2147483648 <= genFn0(128)))) {
        v6 = v6
        v3 += v3
    }
    switch v1 {
    case "a":
        fmt.Println(v3)
    default:
        if len(v5) > 0 {
            v5[0] = v1
        }
    }
    v6 = 0
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            if true {
                fmt.Println(v6, v1, v19)
            }
            if (false || ((v19 & v19) <= v19)) {
                v24 := GenS2{F0: 42}
                _ = len(v5)
                v25 := GenS2{F0: 88}
                _ = v24
                _ = v25
            } else {
                var v26 uint = 336
                _ = v26
            }
            fmt.Println(fi22, v6)
            _ = fi23
        }
        _ = fi22
    }
    if true {
        v27 := GenS1{F0: ""}
        fmt.Println(v1, v19)
        _ = v27
    } else {
        var v28 int = 4294967295
        if true {
            var v29 uint64 = (v6 % ((1) | 1))
            v30 := GenS3{F0: int64(v19)}
            var v31 float64 = 0.1
            _ = v29
            _ = v30
            _ = v31
        }
        v4 += 0.0
        _ = v28
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        v6 = v6
        v6 = 4294967295
        for fi33 := 0; fi33 < 5; fi33++ {
            if (1e10 < 1e10) {
                fmt.Println(v19, v1)
            }
            v34 := GenS4{F0: int64(v3), F1: 192, F2: (v19 + uint32(v6))}
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    v35 := GenS4{F0: -368, F1: 127, F2: v19}
    var v36 int32 = -2147483648
    switch 32768 {
    case 4294967295:
        switch v1 {
        case "\n\t":
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 int16 = v3
                fmt.Println(v38, v4, v36)
                switch 752 {
                case 2147483647:
                    v4 -= (float64((((0.1 / float32(((v4 + (v4 + v4)) + v4))) + 0.0) - (1.5 + 1.5))) * v4)
                case 335:
                    var v39 bool = true
                    _ = v39
                case 9223372036854775807:
                    var v40 int8 = -114
                    _ = v40
                default:
                    v41 := GenS5{F0: 100, F1: -957, F2: 18446744073709551615}
                    _ = v41
                }
                _ = fi37
                _ = v38
            }
        default:
            if (1e10 == ((0.0 + 0.1) - 1e10)) {
                v36 = -832
            } else {
                var v42 uint16 = 65535
                v1 = v1
                var v43 uint16 = v42
                _ = v42
                _ = v43
            }
        }
    }
    var v44 float32 = (1e10 + float32(v4))
    v36 += v36
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
