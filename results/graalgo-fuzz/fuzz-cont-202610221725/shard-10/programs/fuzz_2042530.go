package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS4 struct { F0 int; F1 uint8 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v4 := []int{32767, -1, -9223372036854775808, 388}
    v5 := make(map[uint32]bool)
    var v6 uint8 = 235
    var v7 uint16 = 424
    switch 985 {
    case 1:
        v1 = -1
    case 32768:
        var v8 uint16 = 42
        _ = v8
    case -275:
        v9 := v3[110]
        _ = v9
    }
    var v10 uint = 4294967295
    v7 += v7
    delete(v2, ((("\n\t" + (("hello" + ("a" + "foo bar")) + "hello")) + ("a" + ("foo bar" + "hello"))) + "hello"))
    switch -188 {
    case 128:
        var v11 float32 = 3.14159
        _ = v11
    default:
        fmt.Println(v1)
    }
    v7 += 65535
    var v12 int8 = 71
    switch "a" {
    case "":
        fmt.Println(v6)
    case "foo bar":
        if true {
            for fi13 := 0; fi13 < 6; fi13++ {
                for fi14 := 0; fi14 < 20; fi14++ {
                    v15 := GenS1{F0: 4294967295}
                    fmt.Println(v1, v1)
                    if (!((v7 / ((v7) | 1)) < v7)) {
                        v16 := GenS2{F0: (v6 << (((v10 & 1)) & 63)), F1: 243}
                        _ = v16
                    } else {
                        if (((!((((((3.14159 / 1.5) * 1e10) / 1.0) * 3.14159) - (-2.25 + ((3.14159 - (0.0 * 1.5)) / 1.5))) == 0.1)) && (!(1.5 != (3.14159 - float64(1.5))))) || true) {
                            delete(v5, 42)
                            v17 := GenS1{F0: (v10 >> (((v10 - (v10 + v10))) & 63))}
                            switch "\n\t" {
                            case "foo bar":
                                if false {
                                    v12 = 0
                                }
                            }
                            v1 = 42
                            _ = v17
                        }
                        v18 := GenS3{F0: 556, F1: 18446744073709551615, F2: v12}
                        switch "\n\t" {
                        case "\n\t":
                            var v19 uint64 = 763
                            _ = v19
                        }
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
        } else {
            var v20 int8 = (v12 / ((v12) | 1))
            fmt.Println(v12, v10)
            _ = v20
        }
    case "a":
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int32 = -145
            _ = fi21
            _ = v22
        }
    default:
        v12 = v12
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v23 uint16 = v7
    if (v6 > v6) {
        for fi24 := 0; fi24 < 16; fi24++ {
            delete(v5, 515)
            v25 := GenS3{F0: 554, F1: uint64(v23), F2: v12}
            _ = fi24
            _ = v25
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 float64 = (1.5 / (0.1 * 1e10))
            v28 := GenS4{F0: int(v6), F1: v6}
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
