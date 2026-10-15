package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int32; F2 bool }
func genFn0(p0 uint16, p1 uint64, p2 uint8) string {
    switch -2147483648 {
    case 4294967295:
        v7 := GenS1{F0: 2, F1: ("foo bar" + "hello")}
        _ = v7
    case 32768:
        v8 := GenS1{F0: 1, F1: "hello"}
        _ = v8
    case -2147483648:
        if true {
            var v9 float32 = 0.1
            _ = v9
        } else {
            var v10 string = "foo bar"
            v11 := GenS2{F0: 1}
            _ = v10
            _ = v11
        }
    default:
        if ((1e10 < -2.25) && false) {
            v12 := []int{-1, 128, 9223372036854775807}
            _ = v12
        } else {
            v13 := []int{9223372036854775807, -627, 902, 0}
            _ = v13
        }
    }
    var v14 uint64 = (p1 / ((4294967295) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return ("foo bar" + "hello")
}

func main() {
    var v1 uint16 = 571
    v2 := make(map[int]int)
    var v3 float64 = 0.1
    var v4 uint16 = v1
    var v5 uint16 = v1
    v6 := make(map[uint32]bool)
    if (2 >= 761) {
        v15 := []int{4294967295, 9223372036854775807, 9223372036854775807, -2147483648}
        var v16 uint32 = 1
        v17 := GenS3{F0: -110, F1: -9223372036854775808}
        _ = v15
        _ = v16
        _ = v17
    }
    v1 = (v4 % (((v4 * 514)) | 1))
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 bool = (((239 >= 2) && ((true || true) && (!("hello" < "")))) && (fi18 == 4294967295))
        switch (fi18 % (((fi18 / (((fi18 - (fi18 ^ 32768))) | 1))) | 1)) {
        case 954:
            var v20 uint32 = 2
            _ = v20
        case -2147483648:
            fmt.Println(fi18, v3, v5)
        }
        v21 := v2[fi18]
        _ = fi18
        _ = v19
        _ = v21
    }
    v22 := v6[535]
    if v22 {
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint16 = v5
            _ = fi23
            _ = v24
        }
        v1 = 62
        v4 = v5
    }
    if v22 {
        v25 := GenS1{F0: 42, F1: "a"}
        switch ("hello" + ("\n\t" + ("hello" + "foo bar"))) {
        case "\n\t":
            v22 = true
        default:
            switch "a" {
            case "foo bar":
                var v26 int = 749
                _ = v26
            }
        }
        v27 := []int{4294967295, 1, -320, int(v5)}
        _ = v25
        _ = v27
    } else {
        fmt.Println(v3, v4)
        for fi28 := 0; fi28 < 7; fi28++ {
            for fi29 := 0; fi29 < 3; fi29++ {
                v30 := GenS4{F0: 0}
                fmt.Println(v5, v3)
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
        switch -660 {
        case -2147483648:
            v3 -= -2.25
        case -9223372036854775808:
            v4 = v4
        case 32767:
            var v31 int = 9223372036854775807
            _ = v31
        }
    }
    var v32 float32 = 0.0
    var v33 int32 = -721
    switch 42 {
    case -9223372036854775808:
        var v34 int = 4294967295
        _ = v34
    }
    if ((v32 + (0.0 / v32)) < (v32 / v32)) {
        v35 := GenS1{F0: 255, F1: "hello"}
        if ((((v32 == v32) && v22) || false) && v22) {
            fmt.Println(v22, v4)
        }
        v36 := v6[42]
        _ = v35
        _ = v36
    }
    v4 += (v4 * v4)
    switch -9223372036854775808 {
    case 9223372036854775807:
        for fi37 := 0; fi37 < 12; fi37++ {
            v1 += (v1 * (v4 * 215))
            _ = fi37
        }
    default:
        v38 := GenS5{F0: ((^(v33)) >= 42), F1: v33, F2: ("foo bar" != "\n\t")}
        _ = v38
    }
    delete(v2, 128)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
