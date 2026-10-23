package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64 }
func genFn0() uint8 {
    var v4 float32 = 1e10
    v5 := []int{127, 1, 32768, 0}
    switch ((("foo bar" + "") + ("" + "\n\t")) + "foo bar") {
    case "foo bar":
        v6 := GenS1{F0: 282, F1: 299, F2: 932}
        _ = v6
    }
    if (-9223372036854775808 >= 0) {
        if (127 < 1) {
            v4 = (float32(0.1) - 1e10)
        }
    } else {
        v4 = v4
        var v7 uint = 65535
        v7 = 2
        _ = v7
    }
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 string = "\n\t"
    var v2 uint8 = 2
    v3 := make(map[string]int)
    var v8 float32 = float32(0.1)
    v2 = v2
    _ = len(v3)
    var v9 uint = 0
    v10 := []int{-2147483648}
    v11 := GenS2{F0: -1}
    var v12 bool = (0 <= -604)
    if (v12 || ((!(("a" + v1) < v1)) && (42 <= 0))) {
        v8 = -2.25
        fmt.Println(v2, v1, v12)
        if false {
            v13 := GenS2{F0: 9223372036854775807}
            for fi14 := 0; fi14 < 10; fi14++ {
                fmt.Println(v9, v8, v12)
                v12 = (v9 == 0)
                v8 -= -2.25
                _ = fi14
            }
            _ = v13
        } else {
            for fi15 := 0; fi15 < 1; fi15++ {
                for fi16 := 0; fi16 < 9; fi16++ {
                    v8 += v8
                    _ = fi16
                }
                var v17 float32 = ((((((v8 * 1.5) - 1e10) - -2.25) - (float32(0.1) * (1.5 - float32(0.1)))) + v8) * 3.14159)
                _ = fi15
                _ = v17
            }
            switch 4294967295 {
            case 769:
                if v12 {
                    var v18 int32 = 1
                    v2 = 151
                    v19 := GenS2{F0: 9223372036854775807}
                    _ = v18
                    _ = v19
                } else {
                    var v20 float64 = -2.25
                    _ = v20
                }
            case 128:
                var v21 int8 = 10
                _ = v21
            default:
                switch "" {
                case "":
                    var v22 int32 = -1
                    _ = v22
                default:
                    v23 := GenS2{F0: -9223372036854775808}
                    _ = v23
                }
            }
            var v24 int8 = 42
            _ = v24
        }
    }
    var v25 uint32 = 65535
    switch 127 {
    case 127:
        v8 += v8
    default:
        v1 = v1
    }
    v1 = (v1 + v1)
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 uint8 = v2
            v2 = (v2 >> (((^(v9))) & 63))
            if v12 {
                switch v1 {
                case "foo bar":
                    v28 -= 8
                default:
                    if (!v12) {
                        if (false && ((!v12) || v12)) {
                            v8 += v8
                            v10 = append(v10, (fi26 ^ 9223372036854775807))
                        } else {
                            fmt.Println(fi27)
                            var v29 int32 = int32(fi27)
                            _ = len(v3)
                            _ = v29
                        }
                        v25 -= uint32(v28)
                    } else {
                        v30 := v3[v1]
                        var v31 int8 = -85
                        _ = v30
                        _ = v31
                    }
                }
                v32 := GenS1{F0: 0, F1: 565, F2: 42}
                v12 = v12
                _ = v32
            }
            var v33 uint16 = 799
            fmt.Println(v8, v8, v1)
            _ = fi27
            _ = v28
            _ = v33
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v8
    _ = v9
}
