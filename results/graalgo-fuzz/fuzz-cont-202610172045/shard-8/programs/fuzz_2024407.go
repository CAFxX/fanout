package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0(p0 uint32, p1 uint8) int32 {
    p1 -= p1
    p1 += 255
    _ = p0
    _ = p1
    return -2147483648
}
func genFn1(p0 uint) string {
    p0 = p0
    if (4294967295 >= 32768) {
        p0 = 164
        v5 := GenS1{F0: ("" + ((("\n\t" + "") + ("" + "\n\t")) + "foo bar")), F1: 4294967295}
        _ = v5
    }
    switch "foo bar" {
    case "hello":
        p0 += 438
    case "a":
        if true {
            v6 := []int{}
            for fi7 := 0; fi7 < 19; fi7++ {
                switch ((("hello" + "hello") + "") + "a") {
                case "hello":
                    for fi8 := 0; fi8 < 20; fi8++ {
                        var v9 int = fi7
                        _ = fi8
                        _ = v9
                    }
                default:
                    v10 := GenS2{F0: -638, F1: 177}
                    _ = v10
                }
                switch "foo bar" {
                case "":
                    var v11 int64 = -193
                    _ = v11
                case "hello":
                    var v12 bool = false
                    _ = v12
                case "a":
                    fmt.Println(fi7, p0)
                }
                _ = fi7
            }
            fmt.Println(p0)
            _ = v6
        } else {
            v13 := []int{-711}
            _ = v13
        }
    }
    var v14 int = 0
    _ = p0
    _ = v14
    return (((("hello" + "a") + "") + "\n\t") + "")
}

func main() {
    var v1 uint = 42
    v2 := []string{"\n\t", "a"}
    var v3 int64 = 1
    var v4 float64 = 1.5
    v15 := GenS3{F0: (float64(0.1) != v4), F1: -754}
    var v16 int64 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 uint16 = 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        if false {
            var v19 int64 = -953
            _ = v19
        }
        v17 += v17
        var v20 int8 = 100
        _ = fi18
        _ = v20
    }
    var v21 int16 = 0
    if false {
        fmt.Println(v16, v21, v3)
    } else {
        v17 = v17
        fmt.Println(v4, v1, v21)
        v17 = v17
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v4)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi22
    }
    if (!(0 < 1)) {
        fmt.Println(v1, v17)
    }
    switch 2147483647 {
    case -84:
        fmt.Println(v16, v4, v1)
    case -823:
        fmt.Println(v21)
    default:
        v3 += (v3 * v3)
    }
    v16 = v3
    var v23 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
