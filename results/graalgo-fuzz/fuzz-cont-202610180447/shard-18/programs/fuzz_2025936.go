package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int16 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 bool, p1 string, p2 bool) uint64 {
    switch 2147483647 {
    case 314:
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := []int{(fi7 >> ((1) & 63)), (fi7 >> ((255) & 63)), fi7, 2147483647}
            switch fi7 {
            case 9223372036854775807:
                fmt.Println(p0, fi7, p2)
            default:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = fi7
            _ = v8
        }
    case 2147483647:
        p2 = p2
    case 32767:
        p0 = p2
    default:
        var v9 float32 = 1.5
        _ = v9
    }
    switch "\n\t" {
    case "":
        p1 = "\n\t"
    case "a":
        p0 = p2
    default:
        p1 = ((((p1 + (p1 + "\n\t")) + (((p1 + p1) + p1) + "")) + ((p1 + "foo bar") + (((((((p1 + p1) + p1) + p1) + p1) + ("" + "hello")) + (p1 + p1)) + ((p1 + (p1 + p1)) + p1)))) + ("\n\t" + ((p1 + "foo bar") + p1)))
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 bool = false
    var v2 int32 = 0
    var v3 float32 = float32(3.14159)
    v4 := []uint32{631, 0, 2147483648}
    var v5 int32 = v2
    var v6 uint16 = 639
    v4 = append(v4, 229)
    switch (((("" + ((("\n\t" + "a") + "\n\t") + "foo bar")) + ("a" + "")) + "") + ("hello" + "foo bar")) {
    case "\n\t":
        v3 -= v3
    case "":
        if v1 {
            if v1 {
                for fi10 := 0; fi10 < 17; fi10++ {
                    v1 = v1
                    _ = fi10
                }
            } else {
                if ((("\n\t" + (("" + "foo bar") + "foo bar")) + ("foo bar" + ("\n\t" + "\n\t"))) != "\n\t") {
                    var v11 int64 = 9223372036854775807
                    _ = v11
                }
                switch "\n\t" {
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                case "hello":
                    v12 := GenS1{F0: true, F1: "\n\t", F2: 875}
                    _ = v12
                }
                v3 = -2.25
            }
            v4 = append(v4, 134)
        }
    case "hello":
        if v1 {
            fmt.Println(v2, v1)
        } else {
            v3 = 0.1
            v1 = false
            var v13 int8 = -128
            _ = v13
        }
    }
    switch -9223372036854775808 {
    case 128:
        v5 += (v5 + 1)
    case -2147483648:
        v2 = v2
    case 409:
        if v1 {
            v14 := GenS2{F0: -667, F1: "\n\t"}
            fmt.Println(v3, v3)
            _ = v14
        }
    }
    var v15 int = -9223372036854775808
    fmt.Println(v1, v5, v1)
    v15 = v15
    _ = len(v4)
    v5 = (v2 / ((int32(v15)) | 1))
    v16 := GenS1{F0: v1, F1: "hello", F2: 560}
    _ = len(v4)
    if false {
        var v17 int = 32768
        v3 = -2.25
        v18 := GenS2{F0: v15, F1: "\n\t"}
        _ = v17
        _ = v18
    }
    var v19 uint = 255
    var v20 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
