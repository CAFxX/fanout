package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int64; F1 uint64; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int16, p1 int64) uint64 {
    v5 := GenS1{F0: 309, F1: (p0 + -681)}
    v6 := GenS2{F0: 408}
    if (-16 >= 1) {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{fi7, -71, -9223372036854775808}
            _ = fi7
            _ = v8
        }
        var v9 int32 = -1
        fmt.Println(p1)
        _ = v9
    }
    var v10 uint8 = 100
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 bool = (!((true && (0.0 == float32(3.14159))) && true))
    var v2 int = 4294967295
    var v3 uint32 = 65535
    v4 := make(map[int]string)
    v11 := []int{v2, v2, 9223372036854775807}
    for fi12 := 0; fi12 < 19; fi12++ {
        _ = len(v4)
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int64 = 905
            v15 := GenS3{F0: 127, F1: v1}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    if v1 {
        for fi16 := 0; fi16 < 6; fi16++ {
            if (!(0.1 > 1e10)) {
                var v17 float64 = 0.0
                _ = v17
            } else {
                fmt.Println(v1, v3)
            }
            _ = fi16
        }
        v3 = v3
        var v18 string = "a"
        _ = v18
    }
    v3 += v3
    v19 := GenS4{F0: 9223372036854775807, F1: 4294967295, F2: ("\n\t" + (("a" + (("hello" + ("foo bar" + ((("a" + ("hello" + "hello")) + "\n\t") + "hello"))) + ("foo bar" + (("" + "foo bar") + "\n\t")))) + ""))}
    if v1 {
        v20 := GenS5{F0: 2147483648}
        if (1 > v3) {
            var v21 uint32 = v3
            _ = len(v4)
            _ = v21
        }
        _ = v20
    }
    v22 := GenS6{F0: -1, F1: 42}
    v23 := v4[-534]
    switch -1 {
    case 2147483647:
        v24 := GenS5{F0: v3}
        _ = v24
    }
    v23 = (v23 + "\n\t")
    var v25 int32 = 2147483647
    fmt.Println(v25, v3, v23)
    if v1 {
        v2 = (v2 % (((^(v2))) | 1))
    }
    v2 -= 812
    if (false || false) {
        fmt.Println(v25, v25)
        v4[(v2 * v2)] = (v23 + "")
    } else {
        v26 := GenS5{F0: (v3 | 65535)}
        _ = v26
    }
    v27 := GenS7{F0: genFn0(1000, -1), F1: 807}
    fmt.Println(v2, v25, v1)
    var v28 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
