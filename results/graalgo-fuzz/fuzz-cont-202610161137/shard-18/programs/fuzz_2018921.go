package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 int; F1 bool; F2 uint }
func genFn0(p0 int32, p1 int8) uint {
    var v7 int64 = 9223372036854775807
    var v8 uint8 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 343
}
func genFn1(p0 uint64) int64 {
    v9 := GenS1{F0: 1, F1: uint(p0), F2: 127}
    var v10 int32 = 1
    for fi11 := 0; fi11 < 5; fi11++ {
        if (true || (((("\n\t" + "hello") + "") + ("a" + "hello")) != "hello")) {
            var v12 int8 = 0
            _ = v12
        }
        fmt.Println(v10)
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := []int{-9223372036854775808, -390, 128, -583}
    var v2 int64 = 1
    var v3 int16 = 1
    v4 := []string{"hello", ("a" + "hello"), (((("a" + "hello") + (("hello" + (("\n\t" + "") + ("foo bar" + ("hello" + ("foo bar" + ""))))) + "\n\t")) + "") + ""), ((("a" + (("foo bar" + "hello") + "")) + "") + (((("foo bar" + ("" + "\n\t")) + "a") + "hello") + ("hello" + "")))}
    var v5 uint16 = 75
    var v6 bool = true
    var v13 float64 = 3.14159
    v6 = ((959 != -2147483648) && (v5 != v5))
    v6 = (v6 && false)
    if (!v6) {
        var v14 uint16 = (v5 / ((32767) | 1))
        v6 = v6
        v15 := GenS1{F0: 723, F1: genFn0(-1, -128), F2: -128}
        _ = v14
        _ = v15
    }
    var v16 int64 = v2
    for fi17 := 0; fi17 < 13; fi17++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi17
    }
    v6 = (42 == 236)
    v18 := GenS1{F0: 42, F1: uint(v5), F2: 1}
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            if (1000 <= (v3 | (-(v3)))) {
                v16 += 1
                if (!(!v6)) {
                    v21 := GenS2{F0: -1, F1: 2147483648, F2: -128}
                    _ = v21
                }
                _ = len(v4)
                var v22 uint32 = 656
                _ = v22
            }
            _ = fi20
        }
        v16 += (v16 & v2)
        _ = fi19
    }
    fmt.Println(v13, v13, v2)
    for fi23 := 0; fi23 < 15; fi23++ {
        v13 -= -2.25
        _ = fi23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        fmt.Println(v16)
        var v25 int16 = v3
        fmt.Println(v3, v6)
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 16; fi27++ {
            fmt.Println(v3)
            var v28 int64 = genFn1(0)
            var v29 uint16 = v5
            _ = fi27
            _ = v28
            _ = v29
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi26
    }
    v30 := GenS3{F0: 2147483647, F1: true, F2: uint(v5)}
    var v31 int = 743
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
