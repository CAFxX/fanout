package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int64; F2 int8 }
type GenS5 struct { F0 int64; F1 uint8; F2 int }
type GenS6 struct { F0 uint; F1 int; F2 string }
func genFn0(p0 string, p1 int32, p2 uint) uint {
    if (uint32(p2) == 2) {
        switch 126 {
        case 24:
            p1 += p1
        case 0:
            v4 := GenS1{F0: 716, F1: ((p0 + "a") + (p0 + ((p0 + "a") + ("foo bar" + p0))))}
            _ = v4
        }
        v5 := []int{436}
        _ = v5
    } else {
        if true {
            var v6 uint32 = 0
            v7 := GenS1{F0: (p2 * (p2 + p2)), F1: (p0 + p0)}
            if false {
                var v8 string = "foo bar"
                _ = v8
            } else {
                v9 := GenS1{F0: 65535, F1: p0}
                v10 := GenS1{F0: (p2 << (((p2 | uint(v6))) & 63)), F1: "a"}
                _ = v10
                _ = v9
            }
            _ = v6
            _ = v7
        }
        if (((p2 >> (((p2 ^ p2)) & 63)) >= p2) && ((false && false) && false)) {
            var v11 int = 9223372036854775807
            var v12 bool = true
            fmt.Println(v11)
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: 1, F1: 42}
    fmt.Println(p0, p2)
    switch 127 {
    case 42:
        if true {
            var v14 uint64 = 4294967295
            _ = v14
        } else {
            if false {
                if false {
                    var v15 float64 = 0.0
                    v16 := GenS2{F0: 941, F1: 255}
                    _ = v15
                    _ = v16
                }
                var v17 bool = (0.0 >= ((-2.25 * ((-2.25 - ((3.14159 * 1.5) - 0.1)) * ((float32(1e10) - 0.1) - 1e10))) / 1.0))
                p1 -= 309
                _ = v17
            } else {
                var v18 bool = (882 <= 32767)
                if false {
                    var v19 int = -1
                    v20 := GenS3{F0: p1, F1: 2147483647}
                    v21 := GenS2{F0: 1000, F1: 65535}
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    v22 := []int{}
                    fmt.Println(p1, p1, p2)
                    _ = v22
                }
                fmt.Println(p0)
                _ = v18
            }
            if (true && (32767 > 2)) {
                if (!(42 < 32767)) {
                    v23 := GenS3{F0: 0, F1: 268}
                    v24 := GenS3{F0: p1, F1: 42}
                    _ = v23
                    _ = v24
                } else {
                    p1 = (p1 - -2147483648)
                    var v25 uint = p2
                    _ = v25
                }
                p2 += p2
            } else {
                p0 = "a"
            }
        }
    case 32768:
        switch 0 {
        case 128:
            p0 = (p0 + p0)
        case 29:
            if (204 < 83) {
                p2 = (p2 ^ 255)
                v26 := GenS4{F0: 0, F1: 0, F2: -128}
                _ = v26
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (^(p2))
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = 38
    var v3 bool = false
    var v27 uint16 = 1
    v1 -= 0
    v28 := []int{128, -1, 42, 1}
    for fi29 := 0; fi29 < 5; fi29++ {
        switch "foo bar" {
        case "":
            fmt.Println(v2, v1, fi29)
        case "hello":
            v3 = (-878 >= 42)
        }
        var v30 uint64 = v1
        _ = fi29
        _ = v30
    }
    var v31 int32 = 887
    v3 = v3
    v32 := GenS5{F0: v2, F1: 255, F2: -193}
    v31 = 1
    if true {
        v31 = 0
        var v33 int64 = int64(v1)
        v28 = append(v28, 4294967295)
        _ = v33
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 uint32 = 2147483647
        var v36 uint32 = v35
        var v37 uint16 = (^(v27))
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    v28 = append(v28, 1)
    switch 1 {
    case 4294967295:
        v31 += v31
    }
    var v38 uint64 = v1
    switch ("a" + "foo bar") {
    case "hello":
        v38 = (v38 ^ v38)
    case "foo bar":
        v39 := GenS6{F0: 439, F1: 127, F2: (("\n\t" + ("hello" + "hello")) + "\n\t")}
        _ = v39
    case "":
        v3 = ((("a" + "foo bar") + "foo bar") == "hello")
    default:
        v3 = (!v3)
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v31, v38, v3)
    default:
        v3 = ((((-2.25 / 1.5) < 1.5) && v3) || (v27 >= v27))
    }
    v38 += v38
    var v40 int8 = -81
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v40
}
