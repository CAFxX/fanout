package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 bool }
func genFn0(p0 int64, p1 int16) int8 {
    v4 := []int{0, 2147483647}
    var v5 int64 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 string, p1 bool) int32 {
    switch 9223372036854775807 {
    case -9223372036854775808:
        if (false && true) {
            if (2 >= 1) {
                fmt.Println(p1, p1)
                v6 := []int{-2147483648, 0, 127}
                p1 = true
                _ = v6
            } else {
                p1 = (p1 || (1000 >= -32768))
                var v7 int64 = 527
                v8 := []int{-9223372036854775808, 128, 42, 0}
                _ = v7
                _ = v8
            }
            switch p0 {
            case "hello":
                v9 := GenS1{F0: p1, F1: (1 <= 809), F2: 429}
                _ = v9
            }
            var v10 int64 = -9223372036854775808
            _ = v10
        }
    case 32768:
        v11 := []int{584}
        _ = v11
    case -1:
        if p1 {
            v12 := []int{2147483647, 229, 32767}
            var v13 uint = 263
            _ = v12
            _ = v13
        } else {
            fmt.Println(p0)
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 bool = (false || false)
                _ = fi14
                _ = v15
            }
            if false {
                v16 := []int{}
                var v17 int = 576
                _ = v16
                _ = v17
            } else {
                var v18 uint = 390
                _ = v18
            }
        }
    }
    p1 = (!p1)
    var v19 uint = 2
    _ = p0
    _ = p1
    _ = v19
    return -1
}

func main() {
    var v1 float32 = (1e10 * -2.25)
    var v2 bool = (!(-1 < 42))
    var v3 int = -106
    fmt.Println(v3, v2, v3)
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            v22 := GenS1{F0: false, F1: (-70 != 1), F2: -1}
            _ = fi21
            _ = v22
        }
        v2 = (!(9223372036854775807 == 954))
        _ = fi20
    }
    v23 := []int{-9223372036854775808}
    switch (-(v3)) {
    case 549:
        v3 = (v3 ^ v3)
    case 592:
        var v24 int16 = 237
        _ = v24
    case 32767:
        for fi25 := 0; fi25 < 5; fi25++ {
            v23 = append(v23, 32768)
            _ = fi25
        }
    default:
        for fi26 := 0; fi26 < 1; fi26++ {
            v27 := GenS1{F0: (0.1 == v1), F1: v2, F2: -925}
            fmt.Println(v3)
            _ = fi26
            _ = v27
        }
    }
    switch (v3 / ((v3) | 1)) {
    case 42:
        _ = len(v23)
    case -482:
        if v2 {
            for fi28 := 0; fi28 < 2; fi28++ {
                v1 = v1
                _ = fi28
            }
            if true {
                v3 = v3
            }
        } else {
            var v29 uint16 = 255
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        v31 := GenS2{F0: 116, F1: ""}
        _ = fi30
        _ = v31
    }
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    var v32 int = (^(v3))
    fmt.Println(v2, v3, v3)
    v2 = true
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 uint = 1
        v35 := GenS3{F0: (((v34 | v34) < (v34 - (v34 % ((v34) | 1)))) || false)}
        if len(v23) > 0 {
            v23[0] = (^(fi33))
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: (65535 >= 2)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v36
}
