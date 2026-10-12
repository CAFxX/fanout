package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int64 }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }
func genFn0(p0 int8, p1 uint32) uint16 {
    if false {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{fi7}
            _ = len(v8)
            _ = fi7
            _ = v8
        }
        v9 := []int{-188}
        _ = v9
    }
    var v10 float64 = 1.5
    switch ("\n\t" + ("\n\t" + "a")) {
    case "":
        p0 += (-(p0))
    case "hello":
        v10 = v10
    case "a":
        if (!true) {
            switch -326 {
            case 128:
                v11 := GenS1{F0: 881, F1: ((("foo bar" + "hello") + "foo bar") + (("hello" + "hello") + "foo bar")), F2: 42}
                _ = v11
            case 32768:
                for fi12 := 0; fi12 < 12; fi12++ {
                    p1 = (p1 & 1)
                    v13 := []int{fi12, fi12}
                    _ = fi12
                    _ = v13
                }
            case 858:
                p0 += 127
            }
            fmt.Println(v10, v10)
        }
    }
    var v14 int16 = -32768
    var v15 uint16 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    _ = v15
    return (v15 ^ v15)
}
func genFn1(p0 string, p1 bool, p2 uint64) uint64 {
    fmt.Println(p2, p2)
    p0 = p0
    if p1 {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 uint16 = genFn0(-128, 0)
            p2 = 1
            _ = fi16
            _ = v17
        }
        p0 = "a"
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        p2 -= 42
        p2 = p2
        _ = fi18
    }
    v19 := GenS2{F0: 85, F1: 65535, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return p2
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{-1, 32767, 403, -1}
    var v3 int32 = 2147483647
    var v4 bool = (-506 > 42)
    var v5 string = ((("a" + "") + (("a" + "") + "hello")) + ("foo bar" + "\n\t"))
    var v6 bool = v4
    _ = len(v2)
    for fi20 := 0; fi20 < 11; fi20++ {
        if (1 <= 2) {
            v1 = 699
        }
        v4 = (-971 >= 32768)
        v21 := GenS2{F0: -123, F1: 18446744073709551615, F2: 925}
        _ = fi20
        _ = v21
    }
    fmt.Println(v5)
    var v22 float32 = 0.0
    v4 = false
    v4 = true
    v23 := GenS2{F0: -60, F1: v1, F2: genFn0(0, 2)}
    var v24 float32 = (3.14159 * v22)
    v3 -= (v3 / (((v3 * (v3 / ((int32(v1)) | 1)))) | 1))
    v1 += 790
    v24 += float32(float64(v22))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
