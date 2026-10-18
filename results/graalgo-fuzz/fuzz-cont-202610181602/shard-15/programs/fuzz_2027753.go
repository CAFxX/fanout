package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
type GenS3 struct { F0 int32; F1 uint }
func genFn0() int8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 string = ("hello" + "")
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        var v9 uint16 = 258
        _ = v9
    case "foo bar":
        if (4294967295 >= 2147483648) {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint32 = 1
                _ = fi10
                _ = v11
            }
            var v12 int = -9223372036854775808
            _ = v12
        } else {
            v13 := []int{42, -1}
            _ = v13
        }
    }
    var v14 uint64 = 18446744073709551615
    v15 := GenS1{F0: "\n\t", F1: 0}
    var v16 int16 = -555
    _ = v14
    _ = v15
    _ = v16
    return 127
}
func genFn1(p0 uint) int8 {
    var v17 uint = 1
    switch 32767 {
    case 0:
        if false {
            v17 -= p0
        } else {
            var v18 uint = (^(v17))
            var v19 int64 = 631
            _ = v18
            _ = v19
        }
    case 4294967295:
        if false {
            p0 = (^(p0))
            p0 += (p0 << ((p0) & 63))
        }
    default:
        if true {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 int16 = -247
                _ = fi20
                _ = v21
            }
            if true {
                v17 += (v17 | (p0 << ((p0) & 63)))
            }
        } else {
            v22 := GenS1{F0: ("a" + ""), F1: 127}
            _ = v22
        }
    }
    _ = p0
    _ = v17
    return 0
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    var v5 float64 = 1e10
    v6 := []string{v1}
    v5 = v5
    if true {
        v23 := GenS1{F0: v1, F1: -1}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 12; fi24++ {
            v5 -= (v5 + v5)
            v3 += v3
            _ = fi24
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 float64 = ((-2.25 / v5) / (v5 / v5))
            v27 := GenS2{F0: -85, F1: 65535, F2: "a"}
            var v28 int32 = 2147483647
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
    }
    var v29 string = v1
    var v30 uint64 = 2
    v31 := GenS1{F0: ("hello" + v1), F1: 124}
    v5 = v5
    _ = len(v6)
    v3 = ((v3 + 3.14159) * 1.5)
    if true {
        fmt.Println(v5, v29)
    } else {
        if (100 > 43) {
            var v32 float64 = ((v5 / -2.25) / (v5 * 3.14159))
            _ = v32
        } else {
            var v33 uint = 2
            var v34 uint16 = 2
            _ = v33
            _ = v34
        }
    }
    var v35 bool = (v3 >= (v3 + (1e10 / 0.1)))
    fmt.Println(v1, v29)
    _ = len(v6)
    fmt.Println(v1)
    v5 += 1e10
    fmt.Println(v1, v5, v5)
    v36 := GenS3{F0: 247, F1: 255}
    if false {
        for fi37 := 0; fi37 < 16; fi37++ {
            var v38 uint32 = 384
            v35 = v35
            var v39 int64 = 0
            _ = fi37
            _ = v38
            _ = v39
        }
        switch v1 {
        case "":
            fmt.Println(v5, v35)
        case "hello":
            for fi40 := 0; fi40 < 15; fi40++ {
                var v41 uint64 = v30
                _ = fi40
                _ = v41
            }
        case "\n\t":
            v30 = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
