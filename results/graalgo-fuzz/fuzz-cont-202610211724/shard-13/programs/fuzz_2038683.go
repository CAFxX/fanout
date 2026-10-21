package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int; F2 int }
func genFn0(p0 uint32) int16 {
    v7 := []int{1, 0, 127}
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(fi8, p0)
        _ = fi8
    }
    v7 = append(v7, 293)
    _ = p0
    _ = v7
    return 89
}
func genFn1(p0 bool) uint8 {
    p0 = true
    for fi9 := 0; fi9 < 11; fi9++ {
        if (p0 || ((fi9 + fi9) >= (fi9 / (((fi9 | (^(fi9)))) | 1)))) {
            p0 = (0.0 != 1e10)
        }
        fmt.Println(fi9, fi9)
        _ = fi9
    }
    var v10 uint8 = 2
    p0 = ((!(("foo bar" == "\n\t") && (!(((false || true) && p0) || p0)))) || (((-138 != 0) && (-969 >= -1)) || false))
    _ = p0
    _ = v10
    return (v10 * v10)
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = false
    v3 := []string{(("\n\t" + "hello") + "\n\t"), "", (("foo bar" + "\n\t") + "foo bar")}
    var v4 string = "foo bar"
    var v5 string = "foo bar"
    v6 := []uint32{65535}
    v5 = v5
    v2 = (true || v2)
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            v4 = (v4 + v5)
            _ = fi12
        }
        v13 := GenS1{F0: 1}
        _ = fi11
        _ = v13
    }
    switch 547 {
    case 32767:
        v2 = (42 == 2)
    default:
        v4 = v5
    }
    if (42 <= 65535) {
        var v14 float32 = 0.1
        var v15 uint16 = 891
        v16 := GenS1{F0: 9223372036854775807}
        _ = v14
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := v1[(^(fi17))]
            if v2 {
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 uint8 = genFn1(((fi19 - fi17) < (-(fi17))))
                    v3 = append(v3, v4)
                    delete(v1, (fi19 & 128))
                    _ = fi19
                    _ = v20
                }
            } else {
                var v21 uint16 = 143
                var v22 uint64 = 1
                _ = v21
                _ = v22
            }
            _ = fi17
            _ = v18
        }
    }
    v2 = v2
    if true {
        fmt.Println(v4, v2)
        var v23 uint32 = 0
        v24 := GenS2{F0: -9223372036854775808, F1: 708, F2: -145}
        _ = v23
        _ = v24
    } else {
        v5 = ""
        var v25 int32 = 1
        v5 = (v5 + (("hello" + ((v5 + (v4 + v4)) + "a")) + (v5 + (("hello" + v4) + v5))))
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            _ = len(v1)
            _ = fi27
        }
        v4 = v4
        var v28 float32 = ((0.1 * 0.1) / 1.0)
        _ = fi26
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
