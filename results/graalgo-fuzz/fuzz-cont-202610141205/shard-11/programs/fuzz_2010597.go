package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32, p1 uint8, p2 string) uint32 {
    var v7 int8 = 127
    v7 = (v7 ^ v7)
    if ((!(!(1 != 1000))) || ((p0 * p0) <= p0)) {
        var v8 string = ((p2 + p2) + ((p2 + p2) + (p2 + p2)))
        _ = v8
    } else {
        v9 := []int{}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}
func genFn1() int64 {
    var v10 float32 = (3.14159 * 1e10)
    v11 := []int{}
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float32 = (1e10 + 1.5)
        _ = fi12
        _ = v13
    }
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 int16 = -429
    v2 := make(map[string]int)
    var v3 uint16 = 32767
    var v4 bool = true
    v5 := []uint32{0, 42, 4294967295}
    var v6 float32 = 0.0
    if v4 {
        for fi14 := 0; fi14 < 12; fi14++ {
            if v4 {
                for fi15 := 0; fi15 < 18; fi15++ {
                    v16 := v2[("\n\t" + ("foo bar" + "a"))]
                    _ = fi15
                    _ = v16
                }
                v1 -= v1
            }
            _ = fi14
        }
        v17 := GenS1{F0: 1, F1: 1}
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v1 = (v1 ^ v1)
        v19 := GenS2{F0: 0, F1: 127}
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 18446744073709551615
    var v21 bool = true
    v4 = (!true)
    switch "" {
    case "\n\t":
        if ((v6 + v6) >= v6) {
            v3 += (^(v3))
        } else {
            var v22 bool = v21
            _ = v22
        }
    case "hello":
        _ = len(v5)
    default:
        var v23 uint32 = uint32(v3)
        _ = v23
    }
    v4 = true
    v6 += v6
    fmt.Println(v3, v6, v3)
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 uint64 = v20
        v26 := GenS3{F0: 525}
        var v27 int32 = 0
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v4 = false
    v28 := GenS1{F0: 456, F1: 71}
    v29 := GenS3{F0: 687}
    v30 := v2[((("a" + (("\n\t" + "foo bar") + "foo bar")) + "a") + (("\n\t" + "foo bar") + ("hello" + "")))]
    v6 -= -2.25
    v31 := GenS3{F0: v30}
    v3 -= 619
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 19; fi33++ {
            v34 := GenS2{F0: -1, F1: -128}
            v35 := GenS2{F0: 2147483647, F1: 100}
            switch 42 {
            case 128:
                v4 = (genFn0(2147483647, 0, "hello") != 901)
            case 0:
                fmt.Println(v1, v21)
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        if v4 {
            if (v3 != 687) {
                v36 := v2[((("hello" + "foo bar") + "a") + ("hello" + "hello"))]
                v21 = (v36 != (fi32 & (v30 % (((v36 + (v30 / ((-2147483648) | 1)))) | 1))))
                _ = v36
            }
        } else {
            var v37 uint8 = 0
            _ = v37
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
