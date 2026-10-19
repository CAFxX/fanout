package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1.5
    var v3 uint32 = 2
    v4 := make(map[uint32]bool)
    var v5 int16 = -1
    if true {
        v6 := []int{1, 127}
        switch -145 {
        case 1:
            fmt.Println(v5, v3)
        case 9223372036854775807:
            v3 += (v3 * 2147483647)
        case 128:
            switch (("a" + "\n\t") + "hello") {
            case "\n\t":
                for fi7 := 0; fi7 < 12; fi7++ {
                    v2 = (v2 - 3.14159)
                    _ = fi7
                }
            }
        default:
            if false {
                if true {
                    v5 = -32768
                    switch "" {
                    case "foo bar":
                        v3 = (v3 >> ((608) & 63))
                    case "hello":
                        for fi8 := 0; fi8 < 2; fi8++ {
                            switch fi8 {
                            case 2147483647:
                                if (255 == 383) {
                                    v6 = append(v6, fi8)
                                } else {
                                    v9 := GenS1{F0: "\n\t"}
                                    v10 := GenS2{F0: 9223372036854775807}
                                    _ = v10
                                    _ = v9
                                }
                            }
                            fmt.Println(v3, v5)
                            _ = fi8
                        }
                    default:
                        if len(v6) > 0 {
                            v6[0] = 42
                        }
                    }
                }
                v2 -= 0.0
            } else {
                if ((v3 - (v3 | (v3 ^ v3))) == (v3 | (^(v3)))) {
                    v4[83] = (18446744073709551615 != 0)
                    v2 -= (-2.25 + 0.0)
                    if len(v6) > 0 {
                        v6[0] = -9223372036854775808
                    }
                }
            }
        }
        _ = v6
    } else {
        var v11 uint8 = 128
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int = 9223372036854775807
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    v14 := []int{9223372036854775807, 9223372036854775807, 32767, int(v5)}
    fmt.Println(v5, v3, v3)
    delete(v1, 437)
    v2 = v2
    fmt.Println(v5, v2)
    if (18446744073709551615 != 2) {
        var v15 uint64 = 18446744073709551615
        _ = v15
    }
    v16 := GenS3{F0: (v5 - v5), F1: 42}
    if (2 <= 42) {
        var v17 uint64 = 1
        var v18 bool = false
        _ = len(v1)
        _ = v17
        _ = v18
    }
    if (((!true) && false) && true) {
        for fi19 := 0; fi19 < 20; fi19++ {
            v2 += float64(0.1)
            _ = fi19
        }
        v5 = v5
    } else {
        var v20 string = "a"
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v3 -= (v3 * 264)
        v3 = (v3 ^ v3)
        _ = fi21
    }
    v2 -= (float64(3.14159) * v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
