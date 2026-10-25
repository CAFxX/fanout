package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int16 }
func genFn0(p0 uint, p1 int64, p2 int) string {
    v7 := GenS1{F0: ((0 < 2) && (42 == 127)), F1: (!(341 < 0)), F2: 42}
    var v8 int32 = 42
    p2 = p2
    v9 := GenS1{F0: ((4294967295 == 4294967295) && (-2.25 > float32((float64(0.1) * (3.14159 * (0.0 + 1.5)))))), F1: false, F2: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 uint32 = 4294967295
    v5 := make(map[int]bool)
    v6 := []uint32{(^(v4)), v4}
    v4 += v4
    v10 := GenS1{F0: ((!(false && true)) && false), F1: false, F2: 233}
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v1 = v1
            if true {
                switch 128 {
                case 0:
                    switch ("" + ("a" + ((("" + "hello") + "foo bar") + "a"))) {
                    case "hello":
                        fmt.Println(fi11, v1, fi11)
                    case "foo bar":
                        v1 -= 42
                    }
                case -980:
                    var v13 int8 = -1
                    _ = v13
                case 9223372036854775807:
                    var v14 uint16 = 1
                    _ = v14
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                fmt.Println(fi12)
            } else {
                if (true && false) {
                    v4 -= v4
                }
                fmt.Println(v4)
            }
            var v15 uint64 = 2
            _ = fi12
            _ = v15
        }
        var v16 float64 = 1.5
        for fi17 := 0; fi17 < 18; fi17++ {
            v1 = v1
            if (!true) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            v5[fi11] = (!false)
            _ = fi17
        }
        _ = fi11
        _ = v16
    }
    var v18 float32 = 0.1
    v19 := GenS1{F0: (0.1 != ((1.5 - float64(1e10)) - -2.25)), F1: true, F2: -1}
    v18 += 1e10
    v1 -= v1
    var v20 uint64 = 317
    if len(v6) > 0 {
        v6[0] = 65535
    }
    fmt.Println(v4, v4)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
