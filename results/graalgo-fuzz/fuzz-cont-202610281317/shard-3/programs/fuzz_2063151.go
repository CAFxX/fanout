package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint {
    if (false || (-389 >= -444)) {
        if (128 < 127) {
            var v4 float32 = 1e10
            _ = v4
        } else {
            if true {
                var v5 float32 = (float32(3.14159) - float32(0.0))
                var v6 uint64 = 0
                _ = v5
                _ = v6
            }
            v7 := []int{127, 660, 127}
            for fi8 := 0; fi8 < 13; fi8++ {
                fmt.Println(fi8)
                if false {
                    var v9 uint64 = 2
                    v9 -= (v9 | v9)
                    _ = v9
                } else {
                    if (-2147483648 <= 137) {
                        var v10 uint64 = 618
                        var v11 uint64 = v10
                        if ((3.14159 <= 1e10) && false) {
                            v11 = v10
                        } else {
                            var v12 string = (("\n\t" + "a") + "foo bar")
                            _ = v12
                        }
                        _ = v10
                        _ = v11
                    }
                    for fi13 := 0; fi13 < 13; fi13++ {
                        var v14 float32 = float32(0.0)
                        if len(v7) > 0 {
                            v7[0] = fi13
                        }
                        _ = fi13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v16 := GenS1{F0: 287}
                        var v17 int8 = 127
                        if ((v17 | 26) != (v17 * -1)) {
                            v17 = v17
                            var v18 uint16 = 32767
                            v17 = v17
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                _ = fi8
            }
            _ = v7
        }
        var v19 float32 = (-2.25 + ((0.0 / float32((1e10 - -2.25))) - ((1.5 * ((float32(0.0) + 0.0) - -2.25)) + 1e10)))
        _ = v19
    } else {
        var v20 int16 = 564
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 bool = (!(-913 == (fi21 - (fi21 | (fi21 << ((0) & 63))))))
        if v22 {
            fmt.Println(fi21, v22)
            switch fi21 {
            case 2147483647:
                v22 = ((!false) && true)
            case -976:
                var v23 uint = 263
                _ = v23
            case 0:
                fmt.Println(v22)
            }
        }
        var v24 uint16 = 397
        _ = fi21
        _ = v22
        _ = v24
    }
    return 261
}
func genFn1(p0 string) string {
    switch 127 {
    case 724:
        p0 = "hello"
    default:
        if true {
            v25 := []int{9223372036854775807, 286}
            p0 = (p0 + p0)
            v26 := GenS1{F0: 2}
            _ = v25
            _ = v26
        }
    }
    v27 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = v27
    return (p0 + p0)
}

func main() {
    v1 := []int{32767, -1, -9223372036854775808, 127}
    var v2 int16 = 99
    v3 := make(map[int]string)
    switch "a" {
    case "hello":
        if false {
            v2 += (v2 | v2)
        } else {
            fmt.Println(v2)
        }
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 uint16 = 0
        _ = fi28
        _ = v29
    }
    switch "foo bar" {
    case "hello":
        var v30 int32 = 369
        _ = v30
    case "":
        v2 = v2
    }
    var v31 int64 = -1
    switch "" {
    case "\n\t":
        v32 := GenS1{F0: 0}
        _ = v32
    case "foo bar":
        fmt.Println(v2)
    case "":
        var v33 float32 = ((1.5 + -2.25) / 0.1)
        _ = v33
    default:
        var v34 float64 = (0.1 / 3.14159)
        _ = v34
    }
    var v35 bool = (((!true) && (false && true)) || true)
    v2 = -32768
    var v36 uint16 = 32767
    if (1 <= 100) {
        var v37 int32 = 459
        _ = v37
    } else {
        v31 = (^(v31))
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        var v39 float64 = -2.25
        _ = fi38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
}
