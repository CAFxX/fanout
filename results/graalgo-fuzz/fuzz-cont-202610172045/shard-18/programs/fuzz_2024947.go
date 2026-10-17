package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16, p1 bool) int64 {
    if false {
        var v7 bool = true
        var v8 int = 4294967295
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 1, F1: 0}
        v10 := []int{42, -489, -1}
        fmt.Println(p0, p1)
        _ = v10
        _ = v9
    }
    v11 := []int{724, -2147483648, 42}
    _ = p0
    _ = p1
    _ = v11
    return 1
}
func genFn1(p0 uint64, p1 uint8) uint16 {
    v12 := GenS1{F0: 42, F1: 0}
    fmt.Println(p1, p0)
    v13 := GenS2{F0: p0, F1: p0}
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(fi14, p1, fi15)
            p1 = 1
            _ = fi15
        }
        var v16 int = 42
        _ = fi14
        _ = v16
    }
    var v17 uint16 = 32767
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v17
    return v17
}

func main() {
    v1 := []string{("" + ""), "a", "\n\t", ("hello" + "hello")}
    v2 := make(map[int]int)
    v3 := []string{("a" + ("foo bar" + "foo bar")), "", "a", "\n\t"}
    v4 := []string{"a", "", "a"}
    v5 := []int{163}
    var v6 int32 = 680
    var v18 int64 = -1
    v6 = 38
    v19 := GenS1{F0: 65535, F1: 1}
    switch -2147483648 {
    case 32768:
        fmt.Println(v18, v18)
    case 9223372036854775807:
        fmt.Println(v18)
    case 32767:
        v20 := GenS1{F0: 255, F1: int8(v18)}
        _ = v20
    }
    fmt.Println(v6)
    v21 := GenS1{F0: 554, F1: 1}
    v18 = 942
    var v22 uint16 = 32767
    v6 += -1
    v22 = v22
    var v23 int = 802
    if (!(v23 >= 32768)) {
        var v24 int64 = 752
        _ = v24
    } else {
        switch "foo bar" {
        case "hello":
            delete(v2, 32767)
        case "\n\t":
            switch "a" {
            case "hello":
                v23 -= -2147483648
            }
        }
    }
    if ((!(!true)) || (v6 == v6)) {
        v23 = (v23 & 499)
        for fi25 := 0; fi25 < 5; fi25++ {
            v26 := GenS3{F0: 1}
            var v27 bool = true
            if ((v18 | -332) < 1) {
                v28 := v2[432]
                _ = v28
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            v6 += v6
            v18 = v18
            var v31 uint16 = (^(v22))
            if ("" == "foo bar") {
                var v32 int = (fi30 << ((4294967295) & 63))
                _ = v32
            } else {
                fmt.Println(v6, v31)
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                var v33 int16 = -1
                _ = v33
            }
            var v34 float32 = float32(1.5)
            _ = fi30
            _ = v31
            _ = v34
        }
        v6 += (v6 - (v6 - v6))
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
