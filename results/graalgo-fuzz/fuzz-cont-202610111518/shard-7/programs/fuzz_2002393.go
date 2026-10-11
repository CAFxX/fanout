package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint8, p1 uint) int64 {
    switch 757 {
    case 2147483647:
        v4 := GenS1{F0: 0, F1: 0}
        _ = v4
    case 9223372036854775807:
        v5 := GenS1{F0: 4294967295, F1: 42}
        _ = v5
    default:
        p0 -= p0
    }
    v6 := []int{128, 2147483647, 2147483647}
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v7 uint8 = 255
    var v8 uint = 174
    v8 = v8
    if false {
        var v9 bool = true
        if (862 >= 4294967295) {
            var v10 int32 = -201
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint = 2
                var v13 float32 = -2.25
                if v9 {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        switch 4294967295 {
                        case 1:
                            fmt.Println(fi11, fi11, v9)
                        case 42:
                            var v15 int16 = 32767
                            _ = v15
                        case 9223372036854775807:
                            var v16 int64 = -9223372036854775808
                            _ = v16
                        default:
                            var v17 int16 = -32768
                            v17 = v17
                            _ = v17
                        }
                        v7 += (v7 & (v7 & v7))
                        delete(v2, (fi14 | (^(fi11))))
                        _ = fi14
                    }
                    delete(v2, (fi11 + fi11))
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        } else {
            v8 = (v8 << ((uint(v7)) & 63))
            v8 = (v8 & (v8 << ((v8) & 63)))
        }
        _ = v9
    }
    v18 := GenS1{F0: 1, F1: 1}
    v19 := v1[0]
    switch 613 {
    case -135:
        if (-1 != -796) {
            var v20 bool = (!(-2147483648 >= 2147483647))
            v21 := GenS1{F0: 18446744073709551615, F1: -1}
            _ = v20
            _ = v21
        } else {
            v8 = (v8 & v8)
            var v22 string = "hello"
            for fi23 := 0; fi23 < 11; fi23++ {
                if true {
                    var v24 int16 = 32767
                    v8 -= (v8 >> ((v8) & 63))
                    _ = v24
                }
                var v25 int32 = 304
                _ = fi23
                _ = v25
            }
            _ = v22
        }
    case 4294967295:
        var v26 bool = ((42 > 18446744073709551615) && false)
        _ = v26
    case -9223372036854775808:
        fmt.Println(v7, v19)
    default:
        v19 = v19
    }
    fmt.Println(v8, v8, v7)
    v27 := []int{32767}
    _ = len(v1)
    if (((((v19 + (v19 + ("" + (v19 + "a")))) + v19) + v19) >= ("" + "\n\t")) && (v19 != v19)) {
        v19 = v19
        var v28 uint64 = 42
        _ = v28
    } else {
        v19 = v19
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        if (float32(0.0) > 0.1) {
            var v30 uint64 = 18446744073709551615
            v30 += v30
            if false {
                fmt.Println(v30)
            } else {
                var v31 int8 = -128
                for fi32 := 0; fi32 < 16; fi32++ {
                    v30 += (v30 * v30)
                    var v33 uint32 = 348
                    v7 = 2
                    _ = fi32
                    _ = v33
                }
                fmt.Println(v7, fi29, v30)
                _ = v31
            }
            _ = v30
        }
        v34 := GenS1{F0: 317, F1: 42}
        var v35 string = ("foo bar" + (v19 + ((v19 + v19) + v19)))
        _ = fi29
        _ = v34
        _ = v35
    }
    v8 += v8
    for fi36 := 0; fi36 < 9; fi36++ {
        _ = len(v27)
        var v37 int = 2147483647
        _ = fi36
        _ = v37
    }
    v7 = v7
    for fi38 := 0; fi38 < 3; fi38++ {
        v39 := GenS1{F0: 42, F1: 0}
        _ = fi38
        _ = v39
    }
    switch 0 {
    case 0:
        var v40 float32 = (0.1 - 1.5)
        _ = v40
    default:
        v19 = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v7
    _ = v8
}
