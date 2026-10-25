package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 string }
func genFn0(p0 uint64) uint8 {
    p0 = p0
    switch "" {
    case "a":
        p0 += p0
    }
    _ = p0
    return 127
}
func genFn1(p0 uint, p1 uint) int8 {
    var v6 int32 = -2147483648
    var v7 float64 = 1e10
    p1 = (p1 - p1)
    if (!(!((-664 < 1) || true))) {
        v6 = (v6 ^ v6)
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 float32 = 0.0
                _ = fi9
                _ = v10
            }
            var v11 int64 = 1
            _ = fi8
            _ = v11
        }
    }
    var v12 int16 = -339
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    _ = v7
    return -70
}

func main() {
    v1 := []int{358, -1, -1}
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    v4 := []int{}
    v5 := make(map[int]bool)
    var v13 int32 = 0
    fmt.Println(v13)
    var v14 uint = 317
    switch ((("hello" + ("\n\t" + ((("\n\t" + ("hello" + ("foo bar" + "foo bar"))) + "a") + "\n\t"))) + ("" + "foo bar")) + "foo bar") {
    case "a":
        fmt.Println(v13)
    case "":
        var v15 string = ("" + ("\n\t" + ("hello" + "")))
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int8 = 100
            v17 = v17
            _ = fi16
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 int8 = -128
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(fi18, fi18, v13)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    var v21 int32 = (v13 + -2147483648)
    v21 -= v13
    var v22 bool = true
    var v23 int64 = 900
    var v24 uint8 = 128
    v25 := GenS1{F0: v13, F1: 87}
    if ((3.14159 * 0.1) == 0.1) {
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS2{F0: 255, F1: 42}
            v22 = v22
            for fi28 := 0; fi28 < 6; fi28++ {
                switch "a" {
                case "\n\t":
                    v23 -= (v23 / ((v23) | 1))
                case "foo bar":
                    delete(v2, 65535)
                case "a":
                    var v29 uint8 = 127
                    _ = v29
                default:
                    fmt.Println(v13, fi26)
                }
                fmt.Println(v22)
                fmt.Println(v21)
                _ = fi28
            }
            _ = fi26
            _ = v27
        }
        v30 := GenS3{F0: 1, F1: v23, F2: ""}
        _ = v30
    } else {
        _ = len(v2)
        v31 := GenS3{F0: (v23 + v23), F1: v23, F2: ""}
        _ = v31
    }
    v32 := GenS2{F0: 4294967295, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
