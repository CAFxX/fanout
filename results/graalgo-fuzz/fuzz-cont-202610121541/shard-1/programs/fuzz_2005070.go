package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int32; F2 int8 }
func genFn0(p0 int32) int16 {
    switch 573 {
    case -2147483648:
        p0 -= p0
    default:
        v7 := GenS1{F0: (p0 << ((4294967295) & 63)), F1: 60}
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return 6
}
func genFn1(p0 int32) int8 {
    if false {
        v8 := []int{1, 40, int(p0), -9223372036854775808}
        _ = v8
    }
    p0 -= (p0 % ((-1) | 1))
    _ = p0
    return -1
}

func main() {
    var v1 int64 = 1
    var v2 uint = 846
    var v3 int32 = 2147483647
    var v4 bool = (4294967295 == 2147483647)
    v5 := make(map[string]int)
    v6 := make(map[int]string)
    v9 := []int{}
    if (v4 || (!v4)) {
        v2 = (v2 << (((v2 / ((v2) | 1))) & 63))
        v10 := GenS1{F0: (v3 & v3), F1: 67}
        _ = v10
    } else {
        _ = len(v9)
        v11 := GenS1{F0: (v3 ^ 0), F1: 42}
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: -2147483648, F1: 1}
            var v14 int64 = v1
            v1 += (v14 + (-(v1)))
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    var v15 int16 = -137
    if false {
        fmt.Println(v3, v3, v15)
        switch -1 {
        case 1:
            switch ("" + "foo bar") {
            case "foo bar":
                v3 -= (v3 + (v3 ^ (v3 & (v3 % ((v3) | 1)))))
            case "hello":
                fmt.Println(v1)
            }
        case -1:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        case 478:
            for fi16 := 0; fi16 < 16; fi16++ {
                v17 := GenS1{F0: 42, F1: 127}
                v2 += v2
                _ = fi16
                _ = v17
            }
        }
    } else {
        var v18 int = -1
        if false {
            var v19 int64 = (v1 ^ (v1 & 9223372036854775807))
            _ = v19
        }
        var v20 int16 = (^(v15))
        _ = v18
        _ = v20
    }
    v21 := GenS2{F0: "\n\t"}
    v15 -= -9
    v3 -= 1
    fmt.Println(v15)
    var v22 uint32 = 0
    if ((-690 == 4294967295) && v4) {
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 int = 32768
            v25 := GenS2{F0: "foo bar"}
            v26 := GenS2{F0: ("foo bar" + "a")}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    } else {
        var v27 bool = ((genFn1(42) > 100) || true)
        v15 -= -565
        _ = v27
    }
    var v28 int16 = (v15 + v15)
    fmt.Println(v15, v22, v15)
    switch "foo bar" {
    case "hello":
        var v29 float32 = 1.5
        _ = v29
    case "":
        v3 -= -2147483648
    }
    v30 := GenS3{F0: (v2 / (((^(v2))) | 1)), F1: (v3 >> ((v2) & 63)), F2: 114}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
