package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0(p0 int32, p1 uint) bool {
    fmt.Println(p0, p1)
    p0 = (p0 >> ((417) & 63))
    var v6 string = "hello"
    var v7 uint64 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}
func genFn1() bool {
    v8 := GenS1{F0: 255}
    if false {
        if (18446744073709551615 != 255) {
            v9 := []int{42}
            _ = v9
        } else {
            var v10 float64 = (1.5 + 3.14159)
            v11 := GenS2{F0: 100, F1: 2147483647}
            _ = v10
            _ = v11
        }
    }
    _ = v8
    return true
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    var v4 bool = true
    var v5 float64 = (-2.25 / (1e10 - -2.25))
    if (false || v4) {
        v12 := []int{626}
        _ = v12
    }
    v13 := GenS1{F0: 255}
    switch "hello" {
    case "hello":
        v14 := v2[1]
        _ = v14
    }
    switch -1 {
    case 2147483647:
        if (1 > 32767) {
            v15 := v2[uint32(v1)]
            _ = v15
        } else {
            fmt.Println(v3, v3)
        }
    case -929:
        v16 := v2[682]
        _ = v16
    case 9223372036854775807:
        if true {
            switch "foo bar" {
            case "a":
                if v4 {
                    var v17 int16 = -833
                    v18 := GenS1{F0: 433}
                    _ = v17
                    _ = v18
                } else {
                    var v19 int = 4294967295
                    v20 := []int{610}
                    _ = v19
                    _ = v20
                }
            case "foo bar":
                switch "\n\t" {
                case "":
                    v4 = (v4 && true)
                case "a":
                    v2[818] = (("\n\t" + "\n\t") + ("\n\t" + (("\n\t" + ((("foo bar" + "hello") + "a") + "foo bar")) + "foo bar")))
                case "\n\t":
                    var v21 int32 = 1
                    _ = v21
                }
            case "\n\t":
                switch "\n\t" {
                case "":
                    var v22 float64 = v5
                    _ = v22
                default:
                    v1 += 579
                }
            }
            var v23 int64 = -1
            v24 := GenS3{F0: 4294967295, F1: 65535}
            _ = v23
            _ = v24
        } else {
            var v25 uint8 = 128
            var v26 uint = v1
            _ = v25
            _ = v26
        }
    default:
        v4 = false
    }
    var v27 float32 = 0.0
    v28 := GenS2{F0: 34, F1: -33}
    var v29 string = ("a" + "a")
    var v30 uint32 = 4294967295
    v29 = ((v29 + v29) + (v29 + v29))
    v5 -= (v5 * 1.5)
    v27 -= (v27 * v27)
    switch 9223372036854775807 {
    case 32768:
        _ = len(v2)
    case 1:
        var v31 int16 = 42
        _ = v31
    }
    var v32 int16 = 46
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
}
