package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    var v3 uint64 = 2
    v4 := make(map[uint32]int)
    var v5 string = "\n\t"
    var v6 uint8 = uint8(v3)
    fmt.Println(v3)
    for fi7 := 0; fi7 < 19; fi7++ {
        if (3.14159 != ((0.0 - 0.1) + 0.1)) {
            var v8 uint = 18446744073709551615
            _ = v8
        } else {
            v6 -= (v6 / (((v6 ^ 7)) | 1))
            _ = len(v4)
        }
        v3 = uint64(v6)
        _ = len(v2)
        _ = fi7
    }
    fmt.Println(v6, v3, v6)
    fmt.Println(v5, v5, v3)
    v4[0] = 0
    v6 += 1
    var v9 string = ((((("foo bar" + v5) + v5) + v5) + "foo bar") + ("hello" + v5))
    v5 = "\n\t"
    v6 -= (v6 >> ((607) & 63))
    var v10 float64 = 0.0
    switch 9223372036854775807 {
    case 694:
        v3 = v3
    case 54:
        fmt.Println(v9)
    case -38:
        for fi11 := 0; fi11 < 8; fi11++ {
            if (((float64((-2.25 / 1.0)) / v10) != 1.5) && (0 == 0)) {
                for fi12 := 0; fi12 < 6; fi12++ {
                    switch fi11 {
                    case 32768:
                        var v13 int32 = -1
                        _ = v13
                    }
                    v9 = (((v5 + v9) + v9) + v9)
                    v14 := GenS1{F0: 65535, F1: 42}
                    v9 = ("hello" + (v5 + v5))
                    var v15 int64 = -440
                    _ = fi12
                    _ = v14
                    _ = v15
                }
            }
            switch (v5 + v9) {
            case "foo bar":
                _ = len(v2)
            case "a":
                v16 := GenS2{F0: v9}
                _ = v16
            case "":
                for fi17 := 0; fi17 < 10; fi17++ {
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    if false {
                        if (2147483648 > 782) {
                            v10 -= v10
                            switch (v9 + (v5 + "a")) {
                            case "":
                                v6 = (v6 * (v6 & (^(v6))))
                            }
                            var v18 int32 = 2147483647
                            _ = v18
                        } else {
                            switch 1 {
                            case 32768:
                                var v19 string = "\n\t"
                                _ = v19
                            case 2147483647:
                                if false {
                                    v9 = v5
                                }
                            default:
                                v3 += 915
                            }
                        }
                        v20 := GenS1{F0: 2147483647, F1: 42}
                        _ = v20
                    } else {
                        v6 += (^(v6))
                        v3 = (v3 * v3)
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                    }
                    _ = fi17
                }
            default:
                v6 += 2
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                v5 = "\n\t"
                fmt.Println(v6, v3, v3)
                _ = fi21
            }
            _ = fi11
        }
    }
    v3 -= 0
    switch 4294967295 {
    case -2147483648:
        v6 -= (v6 % ((v6) | 1))
    case 32768:
        if false {
            v6 = (v6 | v6)
            if true {
                var v22 float64 = 3.14159
                _ = v22
            }
        }
    }
    v10 += float64(3.14159)
    v6 = 255
    var v23 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
