package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
func genFn0(p0 string) bool {
    if (!true) {
        v5 := GenS1{F0: 1, F1: -1}
        _ = v5
    } else {
        v6 := GenS1{F0: 1000, F1: -128}
        p0 = p0
        if false {
            switch 64 {
            case 42:
                if false {
                    p0 = (("\n\t" + (p0 + "a")) + p0)
                    p0 = p0
                    var v7 int8 = 0
                    _ = v7
                } else {
                    var v8 float64 = 0.1
                    v8 += 0.1
                    fmt.Println(p0)
                    _ = v8
                }
            case 128:
                v9 := GenS1{F0: 175, F1: -128}
                _ = v9
            default:
                if (404 <= 554) {
                    var v10 uint8 = 1
                    v10 -= 128
                    v11 := []int{-221, -2147483648}
                    _ = v10
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 17; fi12++ {
                        v13 := GenS2{F0: 137}
                        _ = fi12
                        _ = v13
                    }
                    p0 = ("hello" + p0)
                    v14 := []int{32768, -1, 544}
                    _ = v14
                }
            }
            var v15 int16 = 1000
            _ = v15
        }
        _ = v6
    }
    v16 := GenS2{F0: 127}
    v17 := GenS3{F0: ((("\n\t" + p0) + ("foo bar" + p0)) + p0)}
    v18 := GenS3{F0: ((("\n\t" + p0) + (((p0 + "foo bar") + (p0 + "hello")) + "foo bar")) + (p0 + (p0 + p0)))}
    for fi19 := 0; fi19 < 13; fi19++ {
        switch p0 {
        case "\n\t":
            v20 := []int{fi19, -9223372036854775808}
            _ = v20
        case "hello":
            p0 = p0
        case "":
            v21 := []int{-581}
            _ = v21
        default:
            var v22 uint = 18446744073709551615
            _ = v22
        }
        p0 = "\n\t"
        _ = fi19
    }
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return true
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 32767
    v3 := make(map[string]string)
    v4 := []uint32{4294967295, 80, 677, 2}
    _ = len(v3)
    fmt.Println(v2)
    v2 += -113
    var v23 uint8 = 127
    v2 -= v2
    v24 := v1[("\n\t" + ("hello" + "foo bar"))]
    var v25 int64 = 287
    delete(v1, ("" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
