package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint; F1 bool }
func genFn0(p0 int) uint64 {
    v4 := GenS1{F0: 42, F1: "a"}
    p0 = (p0 - (^(p0)))
    fmt.Println(p0)
    if true {
        var v5 uint = 512
        var v6 int16 = 0
        _ = v5
        _ = v6
    } else {
        var v7 int16 = 32767
        var v8 int = -247
        v7 = (v7 - v7)
        _ = v7
        _ = v8
    }
    if false {
        fmt.Println(p0)
        var v9 int32 = -2147483648
        var v10 uint16 = 255
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 uint64, p1 int16) uint64 {
    fmt.Println(p0, p1)
    v11 := GenS1{F0: 42, F1: "hello"}
    p0 -= p0
    v12 := []int{310}
    var v13 int16 = (-(p1))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return (p0 + 42)
}

func main() {
    var v1 bool = true
    v2 := []string{"foo bar", (("" + "hello") + "hello")}
    v3 := []int{351}
    if v1 {
        v1 = (!v1)
    } else {
        _ = len(v2)
    }
    var v14 int32 = 0
    if v1 {
        v14 = (v14 | v14)
        switch (("a" + (((("hello" + "") + "") + "hello") + "")) + ("foo bar" + (("foo bar" + "a") + "hello"))) {
        case "hello":
            v15 := GenS2{F0: 128, F1: ("\n\t" + ("" + ""))}
            _ = v15
        case "foo bar":
            var v16 float64 = ((float64(0.0) / (1e10 + 0.0)) - float64(0.1))
            _ = v16
        default:
            switch ("foo bar" + "a") {
            case "foo bar":
                var v17 uint32 = 65535
                _ = v17
            case "hello":
                if v1 {
                    switch "" {
                    case "a":
                        v18 := GenS3{F0: 0, F1: v1}
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 1; fi19++ {
                            switch ("" + "foo bar") {
                            case "foo bar":
                                var v20 int = 128
                                _ = v20
                            case "a":
                                var v21 uint16 = 2
                                _ = v21
                            }
                            _ = fi19
                        }
                    }
                }
            case "a":
                v14 += 465
            }
        }
        var v22 uint32 = 2147483648
        _ = v22
    }
    v14 = (v14 * (v14 - (v14 & (v14 >> ((4294967295) & 63)))))
    fmt.Println(v1)
    if len(v2) > 0 {
        v2[0] = ("hello" + ("a" + (("\n\t" + "foo bar") + ("\n\t" + "a"))))
    }
    _ = len(v3)
    v3 = append(v3, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
