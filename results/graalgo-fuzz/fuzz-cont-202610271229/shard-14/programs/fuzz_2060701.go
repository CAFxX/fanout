package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0(p0 int, p1 string) bool {
    if (!(false || (!(2 >= 557)))) {
        v7 := GenS1{F0: -607, F1: 2147483648, F2: true}
        _ = v7
    }
    p1 = "foo bar"
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(p0, fi8, p1)
        _ = fi8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        p0 -= fi9
        _ = fi9
    }
    _ = p0
    _ = p1
    return (0.1 == 0.0)
}
func genFn1(p0 int8) int32 {
    var v10 int16 = 1
    v10 = v10
    _ = p0
    _ = v10
    return 369
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 uint64 = 189
    var v6 uint = 349
    fmt.Println(v2, v2)
    v11 := []int{2147483647, 2147483647, 32768}
    v1 = (v1 | 2147483647)
    switch "\n\t" {
    case "":
        fmt.Println(v5)
    case "\n\t":
        var v12 int = 32767
        _ = v12
    case "a":
        v13 := GenS1{F0: 292, F1: v1, F2: true}
        _ = v13
    default:
        if ("" > "hello") {
            v11 = append(v11, 1)
        }
    }
    switch 4294967295 {
    case -9223372036854775808:
        v1 -= (v1 << ((v6) & 63))
    case 0:
        v3[-1] = v1
    default:
        if (3.14159 >= 0.0) {
            for fi14 := 0; fi14 < 6; fi14++ {
                v2 -= v2
                if (!(-1 >= (v2 & v2))) {
                    v15 := GenS1{F0: 1, F1: v1, F2: false}
                    v16 := GenS1{F0: -2147483648, F1: v1, F2: (v6 >= (^(v6)))}
                    _ = v15
                    _ = v16
                } else {
                    if (0 > 32767) {
                        v17 := GenS1{F0: 42, F1: (v1 & (v1 << (((v6 ^ v6)) & 63))), F2: true}
                        switch (fi14 % ((fi14) | 1)) {
                        case 9223372036854775807:
                            v18 := GenS1{F0: genFn1(127), F1: 0, F2: ((2 == 17) && false)}
                            _ = v18
                        default:
                            for fi19 := 0; fi19 < 6; fi19++ {
                                v1 = (v1 ^ (v1 + v1))
                                _ = fi19
                            }
                        }
                        _ = v17
                    }
                    if true {
                        v1 += v1
                    }
                    v11 = append(v11, (fi14 ^ 128))
                }
                var v20 uint32 = (v1 ^ 2)
                _ = fi14
                _ = v20
            }
        }
    }
    v21 := GenS1{F0: 0, F1: v1, F2: (!((v1 == (v1 + uint32(v6))) && (true || true)))}
    v2 = v2
    v1 = uint32(v5)
    if false {
        fmt.Println(v1)
    } else {
        var v22 uint32 = (^(v1))
        switch ("hello" + "") {
        case "hello":
            switch (("\n\t" + (("" + "a") + "\n\t")) + "a") {
            case "\n\t":
                v2 = v2
            default:
                v23 := GenS2{F0: 0, F1: (v22 + v22), F2: (-(v2))}
                _ = v23
            }
        }
        fmt.Println(v1, v6)
        _ = v22
    }
    switch 4294967295 {
    case 396:
        var v24 uint64 = 106
        _ = v24
    }
    if false {
        switch -1 {
        case 127:
            v6 = v6
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            fmt.Println(fi25)
            _ = fi25
        }
        v2 += (v2 & (v2 | v2))
    }
    if true {
        var v26 float32 = (1e10 + 1e10)
        if false {
            fmt.Println(v2, v2, v2)
            v11 = append(v11, 9223372036854775807)
            var v27 int64 = v2
            _ = v27
        }
        _ = v26
    } else {
        var v28 int8 = -81
        v5 += (v5 / (((v5 >> ((42) & 63))) | 1))
        fmt.Println(v6, v28, v2)
        _ = v28
    }
    var v29 uint = (v6 & 2)
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 uint8 = 234
        switch 128 {
        case 686:
            _ = len(v4)
        case 467:
            for fi32 := 0; fi32 < 14; fi32++ {
                v1 -= v1
                _ = fi32
            }
        case 32768:
            v33 := GenS2{F0: 728, F1: v1, F2: (v2 / ((v2) | 1))}
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    v34 := GenS3{F0: v1, F1: uint16(v1), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
