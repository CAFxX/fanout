package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    v3 := []int{}
    v2 -= v2
    var v4 bool = true
    switch -234 {
    case 2147483647:
        v4 = true
    case -674:
        switch -626 {
        case 4294967295:
            fmt.Println(v2, v2)
        case 2147483647:
            if len(v3) > 0 {
                v3[0] = 42
            }
        default:
            var v5 uint16 = 1
            _ = v5
        }
    }
    switch 42 {
    case -300:
        v6 := GenS1{F0: "a"}
        _ = v6
    }
    v4 = (true || (!(32767 <= -9223372036854775808)))
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch fi7 {
            case -165:
                var v9 uint32 = 77
                var v10 bool = (!false)
                _ = v10
                _ = v9
            case 32767:
                v11 := v1[1]
                v12 := GenS2{F0: 448, F1: 383, F2: 542}
                v13 := GenS1{F0: (v11 + v11)}
                _ = v11
                _ = v12
                _ = v13
            }
            switch fi8 {
            case 2147483647:
                var v14 uint64 = 18446744073709551615
                _ = v14
            case 4294967295:
                if ((0.0 + ((float64(((3.14159 * 1.5) + 1.5)) - 3.14159) + 0.0)) > 3.14159) {
                    v4 = v4
                    v2 = (v2 ^ 191)
                }
            case 0:
                v4 = (!v4)
            default:
                var v15 int64 = 0
                _ = v15
            }
            switch fi7 {
            case -2147483648:
                fmt.Println(v4, fi8)
            default:
                v3 = append(v3, (fi7 + fi8))
            }
            fmt.Println(v2, v4, fi7)
            if true {
                v16 := GenS1{F0: "foo bar"}
                _ = v16
            } else {
                switch "\n\t" {
                case "\n\t":
                    v4 = (-1 > -85)
                case "a":
                    v17 := GenS1{F0: "\n\t"}
                    _ = v17
                }
                if v4 {
                    v4 = true
                }
                var v18 uint8 = (v2 + v2)
                _ = v18
            }
            var v19 bool = v4
            _ = fi8
            _ = v19
        }
        _ = fi7
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := v1[2147483648]
            v4 = false
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    switch 4294967295 {
    case -9223372036854775808:
        for fi23 := 0; fi23 < 18; fi23++ {
            if true {
                var v24 int32 = 672
                v3 = append(v3, fi23)
                v24 -= 0
                _ = v24
            } else {
                v4 = (4294967295 >= fi23)
            }
            if false {
                v2 = v2
            } else {
                if v4 {
                    fmt.Println(v4, v4, fi23)
                    var v25 string = ((("foo bar" + "a") + "hello") + "foo bar")
                    _ = v25
                }
                fmt.Println(v4, fi23, v2)
                v2 = (v2 << ((uint(v2)) & 63))
            }
            var v26 string = "foo bar"
            _ = fi23
            _ = v26
        }
    case 127:
        v4 = (float32((0.0 / 3.14159)) < 3.14159)
    }
    var v27 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
