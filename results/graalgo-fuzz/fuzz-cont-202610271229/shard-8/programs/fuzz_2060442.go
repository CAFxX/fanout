package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }
type GenS3 struct { F0 string }
func genFn0(p0 uint16) uint16 {
    var v7 uint64 = 0
    var v8 float32 = ((-2.25 * 1.5) + (1.5 + -2.25))
    for fi9 := 0; fi9 < 13; fi9++ {
        if (244 != 0) {
            for fi10 := 0; fi10 < 9; fi10++ {
                v7 = (v7 & (v7 | 0))
                _ = fi10
            }
            switch "foo bar" {
            case "":
                fmt.Println(v7, v7)
            case "\n\t":
                v11 := []int{(fi9 >> ((839) & 63)), fi9, int(v7), 127}
                _ = v11
            }
            var v12 float32 = v8
            _ = v12
        } else {
            v13 := GenS1{F0: 0, F1: fi9, F2: 813}
            _ = v13
        }
        if false {
            switch "hello" {
            case "":
                v8 = v8
            case "foo bar":
                if true {
                    p0 += (p0 * (p0 >> ((10) & 63)))
                    p0 = (p0 ^ (p0 ^ (p0 % (((p0 - p0)) | 1))))
                    p0 += (p0 / ((p0) | 1))
                }
            default:
                switch "foo bar" {
                case "\n\t":
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v15 := GenS1{F0: 255, F1: (fi9 & (fi9 | (fi14 >> ((262) & 63)))), F2: 205}
                        v16 := GenS1{F0: 127, F1: fi14, F2: 32767}
                        p0 -= 687
                        v8 = (1e10 * 1e10)
                        var v17 int8 = 100
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
            }
        } else {
            var v18 uint = 2
            v18 = v18
            switch ("foo bar" + (((("hello" + "hello") + "\n\t") + "\n\t") + "\n\t")) {
            case "":
                var v19 uint = v18
                _ = v19
            case "hello":
                v20 := GenS2{F0: 255, F1: -128, F2: 859}
                _ = v20
            case "\n\t":
                p0 -= p0
            }
            _ = v18
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return (^(p0))
}
func genFn1(p0 int64) uint8 {
    p0 += (p0 << ((2) & 63))
    v21 := GenS3{F0: (("a" + ("hello" + ("" + (((("" + "foo bar") + "foo bar") + "a") + "a")))) + "hello")}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v21
    return 183
}

func main() {
    var v1 float32 = 1.5
    v2 := []string{""}
    v3 := []string{"a", (("" + "hello") + ("a" + "a")), "foo bar"}
    var v4 float64 = 3.14159
    var v5 string = ""
    var v6 float32 = 0.1
    v22 := GenS2{F0: 137, F1: -73, F2: -9223372036854775808}
    if (275 < 0) {
        v4 += (v4 / -2.25)
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 float32 = v6
            v1 += v24
            _ = fi23
            _ = v24
        }
        v25 := GenS1{F0: 255, F1: -465, F2: 32767}
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 uint32 = 65535
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi26
        _ = v27
    }
    v4 -= float64(0.1)
    v4 = v4
    for fi28 := 0; fi28 < 14; fi28++ {
        fmt.Println(fi28, fi28)
        _ = fi28
    }
    var v29 int8 = 42
    if ((0 < 1000) && false) {
        v29 -= (v29 + 42)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
