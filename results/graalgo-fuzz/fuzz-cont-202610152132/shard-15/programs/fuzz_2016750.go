package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string; F1 int16; F2 uint }
func genFn0(p0 uint, p1 uint8, p2 bool) int64 {
    var v7 float32 = 0.0
    switch 32768 {
    case 127:
        v8 := []int{32767}
        _ = v8
    case -2147483648:
        p1 = 128
    }
    switch "hello" {
    case "hello":
        v9 := []int{452, 32768}
        _ = v9
    default:
        if p2 {
            v10 := GenS1{F0: 4294967295, F1: -155}
            _ = v10
        } else {
            var v11 float64 = -2.25
            var v12 bool = (!p2)
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int32 = -645
    var v2 int64 = -284
    v3 := make(map[int]int)
    var v4 int8 = 0
    var v5 float32 = -2.25
    v6 := make(map[uint32]uint32)
    switch ("hello" + "a") {
    case "hello":
        v2 += 461
    default:
        var v13 uint64 = 630
        _ = v13
    }
    switch ("" + (("\n\t" + "a") + "a")) {
    case "foo bar":
        var v14 uint = 65535
        _ = v14
    }
    if false {
        if (!true) {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint64 = 110
                _ = fi15
                _ = v16
            }
        }
        v4 += v4
    } else {
        var v17 int8 = v4
        var v18 int64 = (^(v2))
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: -1, F1: (v1 ^ -1)}
    v4 += v4
    var v20 uint8 = 2
    var v21 uint32 = 2
    v22 := GenS1{F0: 32767, F1: v1}
    if false {
        if true {
            v23 := GenS1{F0: 0, F1: v1}
            var v24 uint32 = v21
            _ = v23
            _ = v24
        } else {
            var v25 int = 127
            fmt.Println(v25, v21)
            _ = v25
        }
        switch -115 {
        case -630:
            var v26 uint = 4294967295
            _ = v26
        case 1:
            v27 := []int{128, 127}
            _ = v27
        default:
            v28 := GenS1{F0: 32768, F1: -2147483648}
            _ = v28
        }
    } else {
        var v29 uint64 = 4294967295
        _ = v29
    }
    v30 := []int{}
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS2{F0: "a", F1: 34, F2: 392}
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        if (v1 >= 60) {
            var v33 int32 = (v1 - (v1 % (((v1 ^ -2147483648)) | 1)))
            var v34 float32 = v5
            _ = v33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    var v35 float64 = -2.25
    for fi36 := 0; fi36 < 18; fi36++ {
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        _ = fi36
    }
    v2 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
