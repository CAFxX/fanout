package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int64) uint32 {
    v6 := GenS1{F0: (-(p0)), F1: "foo bar"}
    var v7 uint32 = 42
    _ = p0
    _ = v6
    _ = v7
    return (v7 & (v7 | (^(v7))))
}
func genFn1(p0 uint16, p1 bool, p2 int32) uint32 {
    switch ("hello" + ("hello" + (("" + (("" + "a") + "foo bar")) + ("hello" + "hello")))) {
    case "hello":
        v8 := GenS2{F0: 127, F1: genFn0(-590)}
        _ = v8
    default:
        var v9 uint16 = (^(p0))
        _ = v9
    }
    var v10 uint64 = 11
    var v11 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 577
}

func main() {
    var v1 uint8 = 127
    var v2 float64 = 1.5
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 int = 127
    var v12 float32 = float32(v2)
    var v13 uint64 = 4294967295
    switch (v5 + -1) {
    case 32768:
        v12 = float32((v2 + (v2 * (float64((3.14159 - v12)) * v2))))
    default:
        var v14 uint = 65535
        _ = v14
    }
    if ((v2 * float64(((v12 * 0.0) + 0.1))) == 3.14159) {
        switch "hello" {
        case "foo bar":
            var v15 uint = 18446744073709551615
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        if (((v5 & -103) == fi16) || ((!false) || false)) {
            fmt.Println(v13, v2, v5)
            v17 := GenS1{F0: 1, F1: ("hello" + "foo bar")}
            v18 := GenS2{F0: 185, F1: 1}
            _ = v17
            _ = v18
        }
        var v19 uint8 = v1
        _ = fi16
        _ = v19
    }
    v1 -= v1
    for fi20 := 0; fi20 < 4; fi20++ {
        v3[""] = (("a" + "a") + "")
        _ = fi20
    }
    v21 := []int{(v5 - (v5 >> ((0) & 63))), v5, v5, (v5 & v5)}
    var v22 uint64 = (v13 / ((v13) | 1))
    if true {
        v23 := GenS3{F0: v22, F1: 2}
        if (v12 >= v12) {
            var v24 uint = 0
            for fi25 := 0; fi25 < 6; fi25++ {
                switch ("a" + "a") {
                case "a":
                    var v26 int32 = -2147483648
                    _ = v26
                case "hello":
                    var v27 string = ""
                    _ = v27
                }
                _ = fi25
            }
            _ = v24
        }
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 float32 = (float32(3.14159) / 1.0)
            v29 = float32(v2)
            _ = fi28
            _ = v29
        }
        _ = v23
    } else {
        fmt.Println(v2, v13, v13)
        v5 += (v5 | v5)
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    var v30 bool = (32767 <= 42)
    v31 := GenS1{F0: -9223372036854775808, F1: ""}
    v1 = v1
    if (false || true) {
        delete(v3, ("hello" + "\n\t"))
    }
    switch "a" {
    case "hello":
        var v32 float64 = float64(v12)
        _ = v32
    case "":
        _ = len(v3)
    case "foo bar":
        v33 := GenS1{F0: 9223372036854775807, F1: ""}
        _ = v33
    }
    var v34 uint32 = genFn1(32767, (1 >= -1), -1)
    var v35 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
