package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint; F1 int64; F2 int }
type GenS3 struct { F0 int32 }
func genFn0() int64 {
    var v7 int32 = 467
    v8 := GenS1{F0: 2147483647, F1: 0}
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 4294967295
    var v5 float32 = -2.25
    v6 := []string{"foo bar"}
    for fi9 := 0; fi9 < 13; fi9++ {
        v5 = v5
        _ = fi9
    }
    v5 += v5
    v10 := GenS1{F0: -1, F1: -710}
    var v11 float32 = v5
    if true {
        v12 := GenS1{F0: 927, F1: -1}
        _ = v12
    } else {
        if false {
            if ((!("" > "a")) || (!((v4 <= v4) || (true || true)))) {
                var v13 float64 = (0.0 * 3.14159)
                v14 := v2[""]
                _ = v13
                _ = v14
            }
        } else {
            var v15 uint = 105
            var v16 int32 = 592
            var v17 int16 = 537
            _ = v15
            _ = v16
            _ = v17
        }
        v11 += v5
    }
    var v18 string = "a"
    var v19 bool = true
    v19 = ((v4 * v4) != v4)
    var v20 uint16 = 65535
    v4 = (v4 << ((42) & 63))
    if ((true && v19) && (true && (v19 || v19))) {
        var v21 bool = false
        var v22 string = "a"
        _ = v21
        _ = v22
    }
    fmt.Println(v18, v19, v11)
    if (-460 < -912) {
        if (!v19) {
            if len(v6) > 0 {
                v6[0] = v18
            }
            v18 = v18
            var v23 float64 = 0.1
            _ = v23
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        fmt.Println(v5, v11)
    }
    v4 = v4
    if v19 {
        v5 = 1.5
        v24 := GenS2{F0: 383, F1: -1, F2: 0}
        if (v18 <= v18) {
            v25 := GenS3{F0: -323}
            _ = v25
        }
        _ = v24
    }
    var v26 uint16 = (v20 << ((501) & 63))
    v5 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
