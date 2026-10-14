package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 string }
func genFn0(p0 int32) int {
    v4 := []int{0}
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: "a"}
        for fi7 := 0; fi7 < 7; fi7++ {
            p0 = (p0 & (p0 % ((p0) | 1)))
            var v8 uint = 768
            _ = fi7
            _ = v8
        }
        var v9 int16 = -1
        _ = fi5
        _ = v6
        _ = v9
    }
    _ = p0
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 int16 = 1
    var v2 int8 = 127
    var v3 uint32 = 696
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(fi10, v2)
        switch ((((("a" + (("\n\t" + ("a" + "")) + "foo bar")) + (((("foo bar" + "hello") + "") + ("hello" + "hello")) + (("a" + "hello") + (((("" + "\n\t") + "hello") + "hello") + ((((("foo bar" + "foo bar") + (("hello" + "\n\t") + "")) + "\n\t") + ("a" + "a")) + "foo bar"))))) + "a") + ("\n\t" + "foo bar")) + ("" + "foo bar")) {
        case "a":
            v11 := []int{(fi10 & 42), (fi10 + -496)}
            _ = v11
        case "hello":
            v3 += v3
        default:
            fmt.Println(v1, fi10)
        }
        var v12 uint32 = v3
        _ = fi10
        _ = v12
    }
    var v13 uint = 175
    v14 := []int{127, 255, 4294967295, 42}
    v15 := GenS2{F0: 1, F1: true}
    if false {
        v16 := GenS3{F0: (^(v3)), F1: 274}
        if (v13 == (v13 << (((v13 >> ((4294967295) & 63))) & 63))) {
            fmt.Println(v2, v1)
        } else {
            v1 = (-(v1))
        }
        for fi17 := 0; fi17 < 8; fi17++ {
            v3 += v3
            v1 = int16(fi17)
            var v18 uint32 = (v3 >> (((v13 & (^(v13)))) & 63))
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    var v19 uint16 = 167
    v1 += v1
    var v20 uint = 18446744073709551615
    if (0.1 < 1e10) {
        v21 := GenS4{F0: "hello"}
        v1 -= v1
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
