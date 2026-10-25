package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 int, p1 uint16) string {
    v4 := GenS1{F0: 0}
    var v5 float32 = ((0.0 - 0.0) / (1e10 * ((1e10 / 1.0) * 1.5)))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ""
}
func genFn1(p0 uint8, p1 int16, p2 int32) uint8 {
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            if false {
                p1 += (p1 / ((p1) | 1))
                p1 -= p1
                v8 := GenS1{F0: 80}
                _ = v8
            }
            _ = fi7
        }
        _ = fi6
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 4294967295
    var v3 uint = 18446744073709551615
    v2 -= (^(v3))
    var v9 string = ("hello" + "\n\t")
    v2 += (v3 - (^(v3)))
    if false {
        var v10 int32 = 2147483647
        fmt.Println(v9, v2)
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v2 += (v2 - 1)
        var v12 bool = true
        v13 := GenS2{F0: 101, F1: ((v9 + "a") + "hello")}
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v9, v3)
    v14 := GenS1{F0: 2}
    v15 := GenS2{F0: 42, F1: v9}
    if false {
        var v16 int = 1
        _ = v16
    } else {
        v2 -= (v3 & (v2 / (((v3 | v3)) | 1)))
        var v17 int32 = 21
        _ = v17
    }
    v9 = v9
    var v18 uint64 = 2
    v18 -= v18
    if true {
        switch (v9 + "") {
        case "hello":
            if true {
                v2 += v3
                v2 += v3
                var v19 bool = (!false)
                _ = v19
            } else {
                for fi20 := 0; fi20 < 13; fi20++ {
                    switch "foo bar" {
                    case "a":
                        v21 := []int{339, fi20}
                        _ = v21
                    default:
                        v2 = 576
                    }
                    v18 = v18
                    _ = fi20
                }
                var v22 uint16 = 907
                _ = v22
            }
        case "foo bar":
            switch 2147483647 {
            case 128:
                if ((v9 + "foo bar") <= v9) {
                    v23 := GenS2{F0: 2, F1: (("hello" + "foo bar") + (("\n\t" + (v9 + "\n\t")) + "foo bar"))}
                    v9 = v9
                    _ = v23
                }
            default:
                var v24 string = v9
                _ = v24
            }
        }
        v2 = (v3 ^ v3)
    } else {
        if false {
            if false {
                var v25 uint64 = v18
                _ = v25
            }
            var v26 string = v9
            v27 := []int{}
            _ = v26
            _ = v27
        } else {
            fmt.Println(v9)
        }
        var v28 int8 = 1
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
