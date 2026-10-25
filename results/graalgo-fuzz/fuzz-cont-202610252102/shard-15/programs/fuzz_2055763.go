package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint16 }
func genFn0() string {
    var v5 uint64 = 1
    v5 += 268
    v5 -= v5
    _ = v5
    return ""
}

func main() {
    var v1 uint16 = 65535
    var v2 float32 = 0.1
    var v3 int = -167
    var v4 int64 = -1
    fmt.Println(v3, v1)
    switch (v3 << ((uint(v1)) & 63)) {
    case -283:
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint64 = 0
            _ = fi6
            _ = v7
        }
    }
    v4 -= v4
    v1 += (v1 | (^(v1)))
    for fi8 := 0; fi8 < 5; fi8++ {
        switch ("" + "foo bar") {
        case "hello":
            if true {
                for fi9 := 0; fi9 < 1; fi9++ {
                    v2 = 1.5
                    if false {
                        if (814 == 0) {
                            v4 -= 211
                            v10 := []int{fi9, fi8, 2147483647, v3}
                            _ = v10
                        } else {
                            var v11 int32 = -524
                            _ = v11
                        }
                        v12 := GenS1{F0: -18, F1: 0}
                        fmt.Println(v1, v1, v4)
                        _ = v12
                    }
                    var v13 int32 = 2147483647
                    _ = fi9
                    _ = v13
                }
                var v14 uint = 4294967295
                v2 -= -2.25
                _ = v14
            }
        case "foo bar":
            switch "a" {
            case "a":
                v15 := []int{(fi8 - 1), fi8, (v3 & v3)}
                _ = v15
            }
        case "":
            v1 -= 30
        default:
            v16 := GenS1{F0: int8(fi8), F1: int8(v4)}
            _ = v16
        }
        _ = fi8
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v3 = v3
        var v18 int8 = 0
        v3 = (v3 + int(v4))
        _ = fi17
        _ = v18
    }
    var v19 float32 = v2
    v20 := []int{(v3 * (^(v3))), 32768}
    if len(v20) > 0 {
        v20[0] = v3
    }
    var v21 uint8 = 85
    if (v4 >= (v4 << ((4294967295) & 63))) {
        v2 += ((v2 / v19) * 1.5)
        v22 := GenS1{F0: -128, F1: -13}
        v4 -= (^(v4))
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        if (true || ((0.1 + ((3.14159 - v19) * (v2 - -2.25))) >= 0.0)) {
            var v24 int8 = 108
            _ = v24
        } else {
            fmt.Println(v21)
            for fi25 := 0; fi25 < 15; fi25++ {
                v1 -= v1
                v26 := GenS1{F0: 1, F1: 1}
                v19 += ((v19 * 1e10) - v19)
                _ = fi25
                _ = v26
            }
        }
        v27 := GenS1{F0: 100, F1: -2}
        if (-1 > 127) {
            v28 := GenS1{F0: 0, F1: -53}
            v29 := GenS1{F0: 100, F1: 0}
            _ = v28
            _ = v29
        } else {
            if (!(v2 < v19)) {
                v1 = v1
            } else {
                v20 = append(v20, int(v1))
                if (!true) {
                    if ((-2.25 - (float32(0.1) - v2)) <= (3.14159 / ((0.1 / 1.0) / (v2 - v2)))) {
                        v30 := GenS2{F0: (v1 - (v1 - v1)), F1: (false && true), F2: 2}
                        var v31 uint = 0
                        _ = v30
                        _ = v31
                    } else {
                        v32 := GenS1{F0: -30, F1: 127}
                        fmt.Println(v1, v4)
                        for fi33 := 0; fi33 < 15; fi33++ {
                            var v34 uint32 = 0
                            v19 = (v19 + (v19 - ((v19 * v19) / 1.0)))
                            _ = len(v20)
                            v21 = 0
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    }
                    v1 = (v1 * (v1 & (v1 + (v1 + v1))))
                }
            }
        }
        _ = fi23
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
