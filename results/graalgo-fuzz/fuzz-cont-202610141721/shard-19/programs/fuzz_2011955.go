package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint16 = 65535
    var v2 int64 = 42
    v3 := []string{(("\n\t" + ("foo bar" + ("hello" + ""))) + "\n\t"), "\n\t"}
    var v4 int = 128
    var v5 uint8 = 128
    v6 := make(map[int]uint32)
    v1 -= v1
    var v7 float32 = float32(((-2.25 + ((((-2.25 * (0.0 / 1.5)) * 0.0) * (float64((-2.25 - (1e10 / 1e10))) * -2.25)) - float64((1e10 - float32(float64(((1.5 * 3.14159) - float32((-2.25 - float64(0.1)))))))))) / -2.25))
    v8 := GenS1{F0: 0}
    v5 = v5
    if (!(!(1e10 <= 0.1))) {
        var v9 string = ((("a" + "\n\t") + "a") + "hello")
        _ = v9
    }
    fmt.Println(v1, v4, v7)
    v5 += (v5 & v5)
    var v10 int = (v4 - v4)
    if ((v4 & (v10 | v10)) <= (v4 ^ 4294967295)) {
        var v11 int64 = 1
        _ = v11
    }
    var v12 int32 = 1
    switch "a" {
    case "\n\t":
        if false {
            var v13 uint = 312
            _ = v13
        } else {
            var v14 float32 = v7
            switch v4 {
            case 4294967295:
                for fi15 := 0; fi15 < 15; fi15++ {
                    v4 -= 32767
                    _ = fi15
                }
            case -9223372036854775808:
                v5 = v5
            default:
                fmt.Println(v7)
            }
            _ = v14
        }
    default:
        switch "foo bar" {
        case "":
            var v16 int32 = v12
            _ = v16
        case "\n\t":
            switch (-(v10)) {
            case 4294967295:
                for fi17 := 0; fi17 < 14; fi17++ {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = fi17
                }
            case 1:
                v18 := GenS1{F0: v12}
                _ = v18
            case 171:
                fmt.Println(v10)
            default:
                var v19 int16 = 42
                _ = v19
            }
        case "a":
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
    }
    fmt.Println(v1, v2)
    v6[(v4 | -1)] = 4294967295
    var v20 uint32 = 0
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := v6[int(v12)]
        _ = fi21
        _ = v22
    }
    fmt.Println(v10, v12, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
