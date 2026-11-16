package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 bool; F2 int16 }
func genFn0(p0 uint16, p1 int64) uint32 {
    v4 := []int{-164}
    var v5 int8 = -1
    var v6 int64 = p1
    if ((1e10 / (((1.5 - float64(-2.25)) * ((3.14159 * 1e10) - 1e10)) - 1e10)) < 0.1) {
        var v7 int = 4294967295
        _ = v7
    }
    var v8 uint16 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]int)
    v3 := []uint32{338, 909, v1}
    v9 := GenS1{F0: 65535, F1: -9223372036854775808}
    v1 += 376
    var v10 string = ("" + "a")
    v2[-749] = 42
    if (184 < 1) {
        var v11 uint32 = (^(v1))
        for fi12 := 0; fi12 < 12; fi12++ {
            if false {
                switch ("\n\t" + ((v10 + v10) + "a")) {
                case "a":
                    v1 = v11
                case "hello":
                    v13 := GenS1{F0: 1, F1: 9223372036854775807}
                    _ = v13
                case "foo bar":
                    if true {
                        var v14 string = ""
                        for fi15 := 0; fi15 < 9; fi15++ {
                            switch -9223372036854775808 {
                            case 818:
                                v1 -= 2147483648
                            case -9223372036854775808:
                                var v16 uint16 = 2
                                _ = v16
                            }
                            fmt.Println(v11)
                            _ = fi15
                        }
                        _ = v14
                    }
                }
                _ = len(v3)
                if true {
                    v10 = ""
                    v17 := GenS1{F0: 911, F1: 234}
                    _ = v17
                } else {
                    v18 := GenS2{F0: -32768, F1: 1, F2: 99}
                    var v19 uint8 = 0
                    v1 += 2147483647
                    _ = v18
                    _ = v19
                }
            } else {
                v10 = v10
                v2[fi12] = (-(fi12))
            }
            var v20 float64 = (0.1 * -2.25)
            _ = fi12
            _ = v20
        }
        v1 = (v1 << ((0) & 63))
        _ = v11
    } else {
        var v21 bool = (255 != 42)
        delete(v2, 42)
        _ = v21
    }
    var v22 uint8 = 0
    fmt.Println(v10, v10, v22)
    fmt.Println(v22, v10)
    switch "\n\t" {
    case "a":
        v22 = (v22 >> ((42) & 63))
    case "\n\t":
        v10 = v10
    }
    var v23 float64 = 0.0
    switch "foo bar" {
    case "\n\t":
        v24 := GenS2{F0: 1, F1: 127, F2: v22}
        _ = v24
    case "hello":
        var v25 uint = 613
        _ = v25
    case "":
        v26 := GenS3{F0: v1, F1: false, F2: -608}
        _ = v26
    }
    var v27 float64 = v23
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v28 int64 = 9223372036854775807
    var v29 string = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v9
}
