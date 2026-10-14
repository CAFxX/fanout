package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
func genFn0(p0 int32, p1 uint8) int16 {
    p1 += (p1 - (p1 << ((58) & 63)))
    switch 9223372036854775807 {
    case 943:
        v6 := GenS1{F0: 36}
        _ = v6
    }
    _ = p0
    _ = p1
    return 1000
}
func genFn1(p0 int64, p1 string, p2 int16) uint64 {
    if (!(true && (false && false))) {
        var v7 int32 = -2147483648
        v8 := GenS2{F0: 0, F1: 301, F2: 100}
        _ = v7
        _ = v8
    }
    p2 = p2
    fmt.Println(p1, p2)
    var v9 uint = 1
    var v10 int16 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 189
}

func main() {
    var v1 int8 = 13
    var v2 uint16 = 73
    v3 := []string{(("foo bar" + "a") + ("\n\t" + ((("\n\t" + "foo bar") + "a") + ((("foo bar" + "hello") + (("hello" + "hello") + ("hello" + ("" + ("foo bar" + "foo bar"))))) + "hello"))))}
    var v4 int8 = (-(v1))
    var v5 uint16 = 1
    if false {
        fmt.Println(v5, v1)
        v1 -= v4
    }
    var v11 float64 = -2.25
    var v12 int8 = v1
    var v13 float64 = 1e10
    var v14 uint32 = 0
    v15 := GenS2{F0: (v4 - 0), F1: int64(v4), F2: (v4 % (((v1 >> ((866) & 63))) | 1))}
    var v16 uint = 18446744073709551615
    for fi17 := 0; fi17 < 2; fi17++ {
        switch (fi17 ^ (^(fi17))) {
        case -1:
            var v18 float32 = float32((v11 + v13))
            _ = v18
        case 42:
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 int8 = int8(v14)
                _ = fi19
                _ = v20
            }
        }
        var v21 int64 = int64(v16)
        var v22 uint64 = genFn1((-(v21)), ("foo bar" + ""), 32767)
        _ = fi17
        _ = v21
        _ = v22
    }
    fmt.Println(v13, v14, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
