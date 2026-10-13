package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
func genFn0(p0 int) uint64 {
    v7 := []int{}
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(p0)
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 2
    var v2 uint16 = 102
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    var v5 bool = (!(false && (65535 >= 65535)))
    v6 := []uint32{574, 561, 2, 306}
    v1 -= v1
    _ = len(v4)
    fmt.Println(v5, v5, v1)
    var v10 uint64 = 302
    if v5 {
        if (-2147483648 <= 42) {
            if (0 != 1) {
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := GenS1{F0: v10, F1: (!v5)}
                    var v13 uint8 = v1
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v6 = append(v6, 65535)
            } else {
                for fi14 := 0; fi14 < 14; fi14++ {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v1 = v1
                        var v16 float32 = 0.1
                        var v17 uint64 = (v10 & (v10 + genFn0((fi14 + fi15))))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                v18 := GenS1{F0: v10, F1: v5}
                switch 291 {
                case 1:
                    var v19 int64 = -562
                    _ = v19
                case 42:
                    if (278 <= 32768) {
                        v20 := v3[-9223372036854775808]
                        _ = v20
                    } else {
                        fmt.Println(v1)
                    }
                case -1:
                    if len(v6) > 0 {
                        v6[0] = 226
                    }
                default:
                    if true {
                        v21 := GenS1{F0: (^(v10)), F1: v5}
                        v2 -= v2
                        _ = v21
                    }
                }
                _ = v18
            }
            var v22 float64 = ((1.5 * (1.5 + (1e10 / 1e10))) + -2.25)
            var v23 uint8 = (v1 / ((v1) | 1))
            _ = v22
            _ = v23
        }
        v1 -= 238
        fmt.Println(v5, v5)
    }
    v5 = (v1 == 255)
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint32 = 618
        _ = fi24
        _ = v25
    }
    v2 = v2
    v5 = v5
    for fi26 := 0; fi26 < 5; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            switch (fi26 + 2147483647) {
            case 32767:
                var v28 uint = 4294967295
                _ = v28
            case 32768:
                fmt.Println(v2)
            case 42:
                v29 := v3[(^(fi26))]
                _ = v29
            }
            switch ("foo bar" + ("\n\t" + "\n\t")) {
            case "":
                fmt.Println(v1)
            case "foo bar":
                v10 = (v10 & v10)
            }
            v10 = genFn0(-1)
            _ = fi27
        }
        var v30 uint = 109
        _ = fi26
        _ = v30
    }
    v5 = v5
    var v31 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
