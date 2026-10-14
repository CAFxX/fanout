package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 int64, p2 bool) int {
    fmt.Println(p1, p1)
    v4 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -660
}
func genFn1() string {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 float32 = (float32(-2.25) / -2.25)
        var v7 bool = (979 != 0)
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS2{F0: true, F1: -128, F2: -9223372036854775808}
        var v10 int16 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{}
    var v12 int = 4294967295
    _ = v11
    _ = v12
    return "foo bar"
}

func main() {
    var v1 string = ("" + "\n\t")
    var v2 int64 = -9223372036854775808
    var v3 bool = ((52 >= 128) || true)
    switch 4294967295 {
    case -2147483648:
        v13 := GenS3{F0: 0}
        _ = v13
    case -1:
        v3 = (false || (!((1e10 + float32(0.1)) <= (float32((3.14159 - 0.1)) + (1e10 + -2.25)))))
    case 127:
        v14 := GenS3{F0: 171}
        _ = v14
    }
    fmt.Println(v2, v1, v2)
    v15 := GenS2{F0: (1 < 742), F1: -1, F2: 9223372036854775807}
    var v16 int8 = -9
    for fi17 := 0; fi17 < 11; fi17++ {
        v2 = (-(v2))
        switch "\n\t" {
        case "hello":
            v18 := []int{(fi17 ^ 202), (fi17 << ((294) & 63)), (fi17 % (((^(fi17))) | 1)), 127}
            _ = v18
        case "foo bar":
            fmt.Println(fi17, v2, v3)
        }
        _ = fi17
    }
    v19 := []int{genFn0(v2, (-(v2)), false), 42, 0}
    var v20 bool = (!false)
    for fi21 := 0; fi21 < 5; fi21++ {
        v20 = (false && (2 == 0))
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
