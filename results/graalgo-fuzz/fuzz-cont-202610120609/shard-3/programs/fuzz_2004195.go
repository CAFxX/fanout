package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 bool; F2 string }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 int8, p1 uint16, p2 int64) int {
    var v4 uint64 = 2
    var v5 int64 = 0
    v6 := []int{815, 127}
    var v7 int32 = 2147483647
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint64 = 42
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    var v3 int64 = -82
    v2 += (v2 * 2)
    v10 := []int{128, genFn0(-37, 625, v3), 1}
    v3 = (v3 >> ((700) & 63))
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int = fi11
            switch v13 {
            case 1:
                fmt.Println(v13)
            default:
                v13 += (v13 >> ((953) & 63))
                if (32767 != 42) {
                    switch (fi11 >> ((2) & 63)) {
                    case -1:
                        if (((0.0 * ((3.14159 / float32(3.14159)) - ((0.1 / 1.0) * ((float32(float64(3.14159)) * ((float32(0.1) + (float32(float64((0.0 / ((3.14159 + (0.0 + (3.14159 / 1e10))) / 1e10)))) / ((0.1 * -2.25) / (float32(-2.25) / (float32(1e10) / 1e10))))) * 0.0)) * ((3.14159 - (0.1 / (((0.0 + (0.1 / (float32((0.1 + 1.5)) / ((1e10 * 0.1) + -2.25)))) + 1.5) - (1.5 / 1.0)))) * 0.0))))) < 1.5) && (!((-2.25 - 0.0) != ((0.1 * (3.14159 * 0.0)) / (3.14159 + 1.5))))) {
                            var v14 bool = false
                            _ = v14
                        }
                        v15 := GenS1{F0: 42}
                        v16 := GenS1{F0: 1}
                        if (!true) {
                            v3 = -44
                            fmt.Println(v3, v13, v3)
                            fmt.Println(fi11, fi11)
                        }
                        v13 = genFn0(127, 1, 42)
                        v17 := GenS1{F0: 1}
                        v18 := GenS1{F0: 18446744073709551615}
                        v13 += genFn0(-104, 2, -1)
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    case 128:
                        fmt.Println(fi12, v3, v13)
                    case 42:
                        fmt.Println(fi12)
                    default:
                        var v19 int8 = 100
                        _ = v19
                    }
                    fmt.Println(v3, fi11)
                }
            }
            fmt.Println(fi11, fi12, v3)
            _ = fi12
            _ = v13
        }
        var v20 int32 = 2147483647
        switch "hello" {
        case "\n\t":
            v10 = append(v10, fi11)
        }
        _ = fi11
        _ = v20
    }
    v21 := GenS2{F0: 2, F1: (42 == 2147483648), F2: "\n\t"}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v22 := GenS2{F0: 2, F1: (2 > 0), F2: ""}
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            if true {
                v3 -= (v3 / ((9223372036854775807) | 1))
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            } else {
                fmt.Println(fi24)
            }
            v2 = v2
            var v25 int16 = 1
            _ = fi24
            _ = v25
        }
        var v26 uint64 = 678
        _ = fi23
        _ = v26
    }
    if true {
        var v27 string = ("foo bar" + "\n\t")
        if true {
            if false {
                v28 := GenS2{F0: 0, F1: true, F2: ((v27 + "\n\t") + "foo bar")}
                if ((-2147483648 == 42) || true) {
                    fmt.Println(v3, v27)
                    if true {
                        v29 := GenS2{F0: 18446744073709551615, F1: (32767 >= -32768), F2: (v27 + v27)}
                        var v30 float64 = 1e10
                        v10 = append(v10, 32767)
                        _ = v29
                        _ = v30
                    }
                }
                _ = v28
            } else {
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 int64 = v3
                    switch (v27 + v27) {
                    case "hello":
                        fmt.Println(v2, v27)
                    }
                    _ = fi31
                    _ = v32
                }
            }
            var v33 bool = false
            _ = v33
        } else {
            v34 := GenS1{F0: 18446744073709551615}
            _ = v34
        }
        _ = v27
    }
    if false {
        for fi35 := 0; fi35 < 9; fi35++ {
            fmt.Println(fi35, fi35, v2)
            _ = fi35
        }
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        var v37 float32 = float32(1.5)
        var v38 uint16 = 65535
        v39 := GenS3{F0: 255, F1: 0}
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    v40 := GenS3{F0: (v2 | v2), F1: 65535}
    v2 += v2
    var v41 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v40
    _ = v41
}
