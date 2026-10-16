package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int; F2 string }
func genFn0() int64 {
    v6 := GenS1{F0: 0, F1: (("foo bar" + "hello") + ("foo bar" + "\n\t"))}
    switch (("\n\t" + "hello") + "hello") {
    case "foo bar":
        var v7 uint16 = 2
        _ = v7
    default:
        var v8 float32 = (3.14159 + -2.25)
        _ = v8
    }
    var v9 int8 = -10
    var v10 uint16 = 42
    var v11 uint = 18446744073709551615
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 int16, p1 int64, p2 int32) int32 {
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS2{F0: 18446744073709551615}
        _ = fi12
        _ = v13
    }
    v14 := GenS3{F0: (-(p0)), F1: 127, F2: "a"}
    fmt.Println(p2, p1, p2)
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS2{F0: 1}
        var v17 float32 = (3.14159 - 3.14159)
        switch 128 {
        case 250:
            v17 = 1.5
        case 42:
            v18 := []int{fi15, (fi15 % ((150) | 1))}
            _ = v18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return int32(p0)
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -52
    var v3 uint32 = 4294967295
    var v4 float64 = (1.5 + (-2.25 / 1.5))
    v5 := make(map[uint32]uint32)
    if true {
        v2 = (v2 ^ 4294967295)
        if (845 == 2) {
            v2 += (v2 * v2)
        } else {
            v3 -= v3
            switch ("a" + ("\n\t" + "")) {
            case "a":
                fmt.Println(v3, v4, v4)
            case "foo bar":
                fmt.Println(v3)
            }
        }
    } else {
        v3 += (^(v3))
        var v19 string = ((("hello" + ("hello" + "foo bar")) + "") + (("\n\t" + "") + "hello"))
        var v20 int16 = 98
        _ = v19
        _ = v20
    }
    v3 = (v3 ^ 722)
    fmt.Println(v3, v3, v2)
    if (v3 <= (^(v3))) {
        v2 = v2
    }
    v4 += v4
    v1[669] = "foo bar"
    fmt.Println(v4, v4, v4)
    v3 = v3
    v3 -= v3
    var v21 float32 = -2.25
    v22 := []int{v2, v2, (v2 ^ (v2 & (v2 ^ (v2 >> ((0) & 63))))), v2}
    var v23 int = v2
    var v24 uint16 = 65535
    v23 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
