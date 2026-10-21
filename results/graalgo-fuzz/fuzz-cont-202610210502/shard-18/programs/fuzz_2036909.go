package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 int16, p1 string, p2 int64) uint16 {
    if false {
        p2 += p2
        for fi7 := 0; fi7 < 16; fi7++ {
            p1 = "\n\t"
            if (64 != fi7) {
                v8 := GenS1{F0: 42}
                _ = v8
            } else {
                var v9 uint32 = 2147483647
                for fi10 := 0; fi10 < 19; fi10++ {
                    if true {
                        v9 -= (^(v9))
                        var v11 float64 = (((0.1 * (float64(float32((float64(3.14159) * (0.0 * 1.5)))) - (1.5 / 1.0))) / (((0.1 + 3.14159) / (1.5 - ((3.14159 * float64(-2.25)) * 3.14159))) / 3.14159)) / (3.14159 - float64(float32(((1e10 - -2.25) / -2.25)))))
                        v9 -= v9
                        _ = v11
                    }
                    _ = fi10
                }
                _ = v9
            }
            fmt.Println(p1)
            _ = fi7
        }
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        p0 += (p0 * (p0 + (p0 << ((25) & 63))))
        for fi13 := 0; fi13 < 2; fi13++ {
            p2 = (p2 - (p2 >> ((65535) & 63)))
            v14 := []int{(fi13 | fi13), (^(fi12)), (-(fi13)), fi12}
            v15 := GenS1{F0: 43}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() uint {
    var v16 int8 = -128
    var v17 int32 = -1
    v18 := GenS2{F0: (^(v17)), F1: true}
    _ = v16
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 128
    v3 := make(map[int]uint32)
    var v4 float64 = -2.25
    var v5 string = "foo bar"
    v6 := make(map[string]uint32)
    v4 = v4
    for fi19 := 0; fi19 < 7; fi19++ {
        switch "hello" {
        case "hello":
            v20 := GenS1{F0: 42}
            _ = v20
        case "":
            v3[v2] = 519
        default:
            v21 := GenS2{F0: 42, F1: true}
            _ = v21
        }
        _ = fi19
    }
    switch "foo bar" {
    case "":
        v22 := GenS2{F0: -19, F1: false}
        _ = v22
    case "\n\t":
        v4 = 3.14159
    default:
        delete(v6, (("" + (v5 + (v5 + (("a" + (((v5 + v5) + v5) + (("a" + (((v5 + "foo bar") + (v5 + v5)) + (v5 + v5))) + ((v5 + ("hello" + v5)) + v5)))) + "foo bar")))) + v5))
    }
    var v23 int64 = 42
    v5 = v5
    v23 -= v23
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint64 = v1
        _ = fi24
        _ = v25
    }
    var v26 int = v2
    fmt.Println(v2, v4)
    fmt.Println(v4)
    v5 = v5
    for fi27 := 0; fi27 < 17; fi27++ {
        switch v5 {
        case "foo bar":
            var v28 int = -874
            _ = v28
        }
        _ = fi27
    }
    switch (v2 + (v2 + 4294967295)) {
    case 4294967295:
        var v29 uint16 = 2
        _ = v29
    default:
        if true {
            v23 = v23
            v30 := GenS2{F0: -331, F1: (((182 >= 2) && false) || false)}
            _ = v30
        } else {
            v31 := GenS2{F0: 2147483647, F1: (!(true && (!true)))}
            v6[v5] = uint32(v1)
            _ = v31
        }
    }
    v32 := GenS1{F0: 18446744073709551615}
    v1 -= (v1 >> ((0) & 63))
    v33 := []int{192, -280, v2, v26}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
