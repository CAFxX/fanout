package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0() int64 {
    var v6 int32 = 105
    var v7 float64 = (0.1 / 0.1)
    if (true || (690 == 201)) {
        if (!false) {
            if false {
                fmt.Println(v7, v7)
                var v8 int16 = 1000
                var v9 int64 = 678
                _ = v8
                _ = v9
            } else {
                if true {
                    var v10 int16 = 32767
                    _ = v10
                }
                switch -2147483648 {
                case 140:
                    var v11 int64 = -1
                    _ = v11
                default:
                    v6 -= (v6 - v6)
                }
            }
        }
    }
    _ = v6
    _ = v7
    return -278
}
func genFn1() int32 {
    v12 := GenS1{F0: 1, F1: 0, F2: 1}
    if (0.0 == -2.25) {
        var v13 uint = 255
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: 452, F1: -215, F2: -32768}
        var v16 float64 = ((1.5 - -2.25) * 0.1)
        v17 := []int{(-(fi14)), (fi14 & fi14), fi14}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 2147483647, F1: (42 <= 100), F2: 2}
    switch ("hello" + ("hello" + "")) {
    case "hello":
        var v19 float64 = (1.5 * -2.25)
        _ = v19
    default:
        if true {
            v20 := []int{-2147483648}
            _ = v20
        }
    }
    _ = v12
    _ = v18
    return 42
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = 3.14159
    var v3 int8 = -1
    v4 := make(map[string]int)
    v5 := []string{"", "foo bar", "hello", ""}
    delete(v4, ("\n\t" + ("\n\t" + "\n\t")))
    v21 := GenS3{F0: 1}
    v2 -= (0.1 - (v2 - (v2 - v2)))
    v22 := GenS3{F0: 32767}
    fmt.Println(v1, v1)
    if (((120 >= genFn0()) || (v3 < (v3 - -128))) && (127 != 1)) {
        v2 += (v2 - v2)
        for fi23 := 0; fi23 < 3; fi23++ {
            for fi24 := 0; fi24 < 14; fi24++ {
                if false {
                    v25 := GenS2{F0: int32(fi24), F1: false, F2: 0}
                    v2 = v2
                    _ = v25
                }
                var v26 int8 = (v3 ^ -117)
                v4[("a" + "")] = fi24
                _ = fi24
                _ = v26
            }
            _ = fi23
        }
    } else {
        fmt.Println(v2, v3, v3)
        v2 = 1e10
        var v27 uint64 = 18446744073709551615
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        v3 = (v3 | 42)
        _ = fi28
    }
    var v29 int32 = v1
    fmt.Println(v3, v3)
    v4["a"] = 127
    switch "" {
    case "hello":
        var v30 bool = false
        _ = v30
    default:
        var v31 int64 = 0
        _ = v31
    }
    var v32 int = 382
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
