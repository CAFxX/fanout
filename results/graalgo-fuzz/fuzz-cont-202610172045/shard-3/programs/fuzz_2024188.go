package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int8, p1 int16) uint64 {
    var v5 string = "\n\t"
    fmt.Println(v5)
    p1 -= p1
    v6 := []int{-196, 128, 4294967295}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 bool) int64 {
    switch -2147483648 {
    case 32768:
        if true {
            for fi7 := 0; fi7 < 13; fi7++ {
                p0 = (255 != 32767)
                _ = fi7
            }
            var v8 int = 42
            _ = v8
        } else {
            for fi9 := 0; fi9 < 14; fi9++ {
                switch "hello" {
                case "a":
                    v10 := GenS1{F0: 0, F1: 255}
                    _ = v10
                case "\n\t":
                    v11 := GenS2{F0: 128, F1: 104}
                    _ = v11
                }
                fmt.Println(p0, fi9)
                _ = fi9
            }
            p0 = ((!(628 > 2)) || (1 >= -3))
            var v12 int = 0
            _ = v12
        }
    case 796:
        for fi13 := 0; fi13 < 14; fi13++ {
            if ((("hello" + ("hello" + ("foo bar" + "a"))) + "hello") != "foo bar") {
                p0 = p0
            }
            v14 := []int{fi13}
            var v15 float64 = 0.0
            _ = fi13
            _ = v14
            _ = v15
        }
    case 127:
        if false {
            fmt.Println(p0)
        } else {
            var v16 bool = (!((-9223372036854775808 >= 42) || (!false)))
            fmt.Println(p0, p0)
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(fi17)
        var v18 uint = 255
        _ = fi17
        _ = v18
    }
    p0 = (127 < 1)
    _ = p0
    return 223
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{0, 32768}
    v3 := make(map[int]string)
    var v4 string = ("hello" + "a")
    switch -650 {
    case -551:
        var v19 uint32 = 455
        _ = v19
    case 128:
        var v20 float64 = 3.14159
        _ = v20
    case 4294967295:
        var v21 uint8 = 1
        _ = v21
    }
    switch 128 {
    case 32768:
        var v22 uint = 65535
        _ = v22
    case 421:
        var v23 bool = (32768 <= 32768)
        _ = v23
    }
    v4 = v4
    if (9223372036854775807 <= 764) {
        v24 := GenS2{F0: 255, F1: 1}
        _ = v24
    }
    if (-1 == -656) {
        for fi25 := 0; fi25 < 10; fi25++ {
            v3[fi25] = "\n\t"
            var v26 bool = false
            var v27 int16 = 32767
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    if (false || false) {
        v4 = v4
    }
    v28 := GenS3{F0: 127, F1: 1}
    var v29 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
