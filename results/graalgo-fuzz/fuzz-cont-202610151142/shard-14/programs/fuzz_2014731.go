package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 uint32, p1 uint8) uint32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        switch ("" + ("" + ("\n\t" + ""))) {
        case "foo bar":
            switch (fi7 >> ((267) & 63)) {
            case 717:
                p0 -= p0
            }
        default:
            v8 := []int{(fi7 - (fi7 ^ (^(fi7)))), (fi7 + fi7), fi7, (fi7 & 9223372036854775807)}
            _ = v8
        }
        var v9 uint32 = 2
        v9 += 968
        _ = fi7
        _ = v9
    }
    var v10 uint = 4294967295
    v10 = (v10 & (v10 - v10))
    _ = p0
    _ = p1
    _ = v10
    return (p0 << ((4294967295) & 63))
}

func main() {
    var v1 float32 = (0.1 * 0.0)
    var v2 float64 = (3.14159 / (0.0 + 1e10))
    var v3 uint8 = 1
    v4 := make(map[int]bool)
    var v5 string = "hello"
    var v6 int32 = -2147483648
    v11 := []int{42, 358, int(v6)}
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            delete(v4, fi12)
            if (-420 > 42) {
                v3 += v3
                var v14 uint16 = 316
                switch v5 {
                case "foo bar":
                    var v15 int = 4294967295
                    _ = v15
                }
                _ = v14
            } else {
                var v16 bool = false
                _ = v16
            }
            _ = fi13
        }
        if (42 == 255) {
            v1 = 1e10
            switch v5 {
            case "foo bar":
                if false {
                    v17 := v4[fi12]
                    _ = v17
                }
            case "":
                v18 := GenS1{F0: (v5 + "a"), F1: 562}
                _ = v18
            default:
                fmt.Println(v3, v6, fi12)
            }
        } else {
            var v19 float64 = (((float64((0.1 + 0.0)) - (0.1 - v2)) + (v2 * v2)) + (v2 * (v2 - v2)))
            _ = v19
        }
        fmt.Println(v6)
        _ = fi12
    }
    var v20 float32 = v1
    v3 = 0
    _ = len(v11)
    v11 = append(v11, 154)
    for fi21 := 0; fi21 < 7; fi21++ {
        fmt.Println(v3)
        _ = fi21
    }
    var v22 uint64 = 0
    for fi23 := 0; fi23 < 1; fi23++ {
        v6 -= v6
        v5 = (v5 + "\n\t")
        _ = fi23
    }
    switch 32768 {
    case -226:
        v11 = append(v11, -14)
    case 0:
        var v24 string = (((v5 + "a") + (("\n\t" + (v5 + (v5 + v5))) + v5)) + (((v5 + v5) + "") + v5))
        _ = v24
    case 9223372036854775807:
        if true {
            var v25 uint32 = 2147483648
            _ = v25
        } else {
            fmt.Println(v2, v3, v20)
        }
    default:
        fmt.Println(v3, v3)
    }
    v26 := GenS1{F0: ("a" + (v5 + v5)), F1: v22}
    v27 := GenS2{F0: true, F1: 42}
    v28 := GenS1{F0: v5, F1: v22}
    var v29 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
