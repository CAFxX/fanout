package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 int64 = 162
    var v2 int = -2147483648
    v3 := []string{"a"}
    v4 := make(map[int]int)
    var v5 uint16 = 1
    fmt.Println(v2, v1, v5)
    fmt.Println(v5, v2)
    delete(v4, v2)
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            v2 -= (fi7 >> ((638) & 63))
            var v8 uint8 = 92
            v2 += fi6
            _ = fi7
            _ = v8
        }
        if (!(!(!((fi6 - fi6) > 32767)))) {
            var v9 string = "foo bar"
            var v10 uint16 = 2
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    if false {
        v11 := GenS1{F0: -152, F1: (v1 & (-(v1))), F2: (v5 + 2)}
        v12 := GenS1{F0: 1, F1: v1, F2: (^(v5))}
        v13 := GenS2{F0: 42}
        _ = v11
        _ = v12
        _ = v13
    } else {
        v5 = 2
    }
    if true {
        var v14 float64 = (0.0 / -2.25)
        _ = v14
    } else {
        if false {
            v2 += 32768
            var v15 float32 = (float32(0.1) / (((0.0 + 3.14159) + (0.1 + float32(1e10))) + 0.0))
            if true {
                v16 := GenS3{F0: -682}
                v2 += (v2 / ((-1) | 1))
                fmt.Println(v2, v15)
                _ = v16
            } else {
                fmt.Println(v15)
                v17 := GenS2{F0: 42}
                v18 := GenS3{F0: -816}
                _ = v17
                _ = v18
            }
            _ = v15
        } else {
            v3 = append(v3, "foo bar")
        }
    }
    delete(v4, 42)
    for fi19 := 0; fi19 < 1; fi19++ {
        v2 += (v2 | v2)
        v20 := GenS3{F0: 1}
        fmt.Println(v1)
        _ = fi19
        _ = v20
    }
    if (!((!true) && (!true))) {
        v21 := GenS4{F0: 2}
        fmt.Println(v1, v2)
        v2 = 128
        _ = v21
    }
    v2 = v2
    v22 := GenS5{F0: (v1 & 1)}
    var v23 float64 = 3.14159
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := v4[-528]
        switch fi24 {
        case 32768:
            switch "" {
            case "hello":
                for fi26 := 0; fi26 < 15; fi26++ {
                    v27 := GenS5{F0: (v1 | (v1 & v1))}
                    _ = fi26
                    _ = v27
                }
            case "\n\t":
                v28 := GenS6{F0: 2147483647}
                _ = v28
            default:
                v29 := GenS4{F0: 1}
                _ = v29
            }
        case 2147483647:
            v25 += 32767
        default:
            fmt.Println(v5, v25)
        }
        var v30 int32 = 42
        _ = fi24
        _ = v25
        _ = v30
    }
    v5 += (v5 + 0)
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 uint32 = 65535
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 uint32 = 2147483647
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
