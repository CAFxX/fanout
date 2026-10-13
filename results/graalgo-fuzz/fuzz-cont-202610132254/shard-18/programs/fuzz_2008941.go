package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
func genFn0(p0 uint64, p1 uint, p2 uint) uint {
    var v4 bool = false
    fmt.Println(p2)
    var v5 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]int)
    var v3 uint32 = v1
    var v6 int8 = 0
    fmt.Println(v1)
    if (0.0 > 0.0) {
        if false {
            var v7 int32 = -1
            _ = v7
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 float32 = (0.0 - 3.14159)
                var v10 uint = 2
                var v11 uint16 = 42
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v2[""] = 42
            var v12 uint8 = 213
            _ = v12
        }
    } else {
        var v13 bool = false
        var v14 uint64 = 18446744073709551615
        _ = v13
        _ = v14
    }
    if (!true) {
        v15 := GenS1{F0: 32767, F1: 42, F2: 32768}
        _ = v15
    } else {
        v1 = (v3 % ((v1) | 1))
        var v16 string = "hello"
        _ = v16
    }
    var v17 int = 9223372036854775807
    if ((0 > 33) && ("" != "hello")) {
        v1 += 230
    } else {
        v18 := []int{v17, v17, v17, (^(v17))}
        _ = v18
    }
    v19 := []int{v17, (v17 * (v17 + v17)), 0, v17}
    switch v17 {
    case -439:
        var v20 uint16 = 1
        _ = v20
    case -2147483648:
        fmt.Println(v6, v3, v3)
    case 4294967295:
        v3 += 2
    default:
        for fi21 := 0; fi21 < 12; fi21++ {
            if false {
                if (true || (fi21 > -501)) {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        switch (("a" + "hello") + ("a" + "foo bar")) {
                        case "\n\t":
                            fmt.Println(v6)
                        }
                        v1 -= 2147483648
                        v23 := GenS1{F0: fi22, F1: 1, F2: fi22}
                        v3 -= v3
                        _ = fi22
                        _ = v23
                    }
                }
                if len(v19) > 0 {
                    v19[0] = -1
                }
                fmt.Println(v1, v6)
            } else {
                v6 = (v6 + v6)
                var v24 uint32 = 265
                _ = v24
            }
            _ = fi21
        }
    }
    if false {
        delete(v2, ("\n\t" + "foo bar"))
        for fi25 := 0; fi25 < 17; fi25++ {
            fmt.Println(v3, v6, v1)
            var v26 int32 = 42
            var v27 uint32 = v1
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v28 := v2["a"]
    var v29 int8 = v6
    v28 -= (v28 + v28)
    v2[("foo bar" + "hello")] = v17
    v30 := GenS1{F0: (v28 ^ (v28 / ((v28) | 1))), F1: 1, F2: (v17 | (^(v17)))}
    var v31 int8 = (v29 * (v6 - v6))
    v32 := GenS1{F0: (v28 ^ (v17 >> ((42) & 63))), F1: 9223372036854775807, F2: (v28 / (((v28 / ((v28) | 1))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v6
}
