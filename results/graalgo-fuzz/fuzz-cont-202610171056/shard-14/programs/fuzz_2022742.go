package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }

func main() {
    var v1 int = 112
    var v2 int8 = 1
    v3 := []int{(-(v1))}
    v4 := []string{}
    v5 := make(map[int]int)
    fmt.Println(v1, v2)
    v6 := v5[(-(v1))]
    fmt.Println(v2, v1)
    fmt.Println(v1)
    switch (v1 * v6) {
    case -1:
        var v7 float32 = 0.0
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        if false {
            fmt.Println(v6, v1, fi8)
            v5[v1] = v6
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v6 = v1
            switch (fi8 << ((18446744073709551615) & 63)) {
            case 42:
                if false {
                    v9 := GenS1{F0: 31}
                    v10 := v5[(^(v6))]
                    switch v10 {
                    case 34:
                        v11 := GenS2{F0: v2, F1: 65535, F2: v2}
                        _ = v11
                    case 951:
                        var v12 float64 = float64(float32((0.1 / 3.14159)))
                        _ = v12
                    }
                    _ = v10
                    _ = v9
                } else {
                    if (true || true) {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            v2 = 0
                            switch "" {
                            case "foo bar":
                                v1 -= (v6 ^ fi8)
                            case "hello":
                                var v14 int = (v1 >> ((0) & 63))
                                _ = v14
                            case "a":
                                v15 := GenS1{F0: (v2 | 127)}
                                _ = v15
                            }
                            var v16 int32 = 42
                            _ = fi13
                            _ = v16
                        }
                    }
                    var v17 float32 = (0.0 * 3.14159)
                    v6 = (v1 / ((v1) | 1))
                    _ = v17
                }
            default:
                _ = len(v3)
            }
        }
        v4 = append(v4, "hello")
        if ((0.1 - float32(float64(0.0))) > (((1e10 / 1.5) - ((-2.25 / 1.5) / 1e10)) - -2.25)) {
            var v18 bool = (42 < 32767)
            if len(v3) > 0 {
                v3[0] = (v1 | (v1 | fi8))
            }
            var v19 int16 = -1
            _ = v18
            _ = v19
        } else {
            v4 = append(v4, (("\n\t" + "") + "hello"))
            v2 -= v2
            v2 -= v2
        }
        _ = fi8
    }
    var v20 int8 = v2
    switch ("hello" + "\n\t") {
    case "foo bar":
        v4 = append(v4, ((("hello" + "foo bar") + ("\n\t" + "")) + "\n\t"))
    case "\n\t":
        var v21 uint64 = 42
        _ = v21
    case "a":
        v22 := GenS3{F0: 1, F1: (-(v2)), F2: 715}
        _ = v22
    }
    v2 -= (v2 % ((v2) | 1))
    v6 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
