package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint, p2 string) uint64 {
    p2 = (((p2 + ((p2 + "\n\t") + p2)) + (p2 + p2)) + p2)
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 int, p1 int64, p2 bool) uint8 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{}
        switch "" {
        case "a":
            v9 := GenS1{F0: 100, F1: 1, F2: 0}
            _ = v9
        case "":
            if p2 {
                v10 := GenS2{F0: 4294967295}
                _ = v10
            }
        case "hello":
            var v11 int8 = 127
            _ = v11
        default:
            for fi12 := 0; fi12 < 13; fi12++ {
                if (!p2) {
                    p1 -= (-(p1))
                    p1 = (^(p1))
                }
                _ = fi12
            }
        }
        _ = len(v8)
        _ = fi7
        _ = v8
    }
    var v13 uint32 = 2147483648
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 0
}

func main() {
    var v1 bool = true
    var v2 int16 = 1
    var v3 int16 = v2
    var v4 int8 = 100
    var v5 uint8 = 127
    var v6 int8 = -128
    fmt.Println(v2)
    var v14 float32 = -2.25
    var v15 uint8 = (v5 | (v5 | 255))
    var v16 int = -201
    v1 = ("a" > ("a" + ("" + "\n\t")))
    for fi17 := 0; fi17 < 6; fi17++ {
        v3 += 1
        _ = fi17
    }
    var v18 uint32 = 465
    if v1 {
        if v1 {
            for fi19 := 0; fi19 < 19; fi19++ {
                fmt.Println(v6)
                switch ("\n\t" + "") {
                case "":
                    v20 := []int{}
                    _ = v20
                }
                var v21 uint32 = v18
                _ = fi19
                _ = v21
            }
        }
        v16 -= (^(v16))
    } else {
        var v22 uint32 = 1
        for fi23 := 0; fi23 < 7; fi23++ {
            v18 -= v18
            v24 := GenS1{F0: int8(v15), F1: uint64(v15), F2: 42}
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v6 = v4
    v25 := []int{0}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
