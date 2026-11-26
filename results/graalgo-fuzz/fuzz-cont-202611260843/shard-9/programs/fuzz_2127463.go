package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint = 921
    v2 := make(map[string]string)
    var v3 uint = (v1 << (((v1 >> ((v1) & 63))) & 63))
    var v4 float32 = 1e10
    if (487 < (^(v3))) {
        v5 := GenS1{F0: 9223372036854775807, F1: (false || true)}
        v4 -= v4
        delete(v2, "foo bar")
        _ = v5
    } else {
        if true {
            for fi6 := 0; fi6 < 15; fi6++ {
                v4 -= v4
                v4 = (1e10 + v4)
                _ = fi6
            }
        }
        var v7 int32 = 2147483647
        v4 = 1e10
        _ = v7
    }
    var v8 uint16 = 945
    v9 := GenS2{F0: 830, F1: 1}
    switch 128 {
    case 2147483647:
        switch 1 {
        case -739:
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int = fi10
                for fi12 := 0; fi12 < 6; fi12++ {
                    v2["foo bar"] = "foo bar"
                    _ = fi12
                }
                _ = fi10
                _ = v11
            }
        case 440:
            v13 := v2["foo bar"]
            _ = v13
        }
    case 0:
        v8 += 2
    case 32768:
        if false {
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(v4)
                _ = fi14
            }
        } else {
            var v15 uint = v1
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 8; fi16++ {
            v3 += v1
            v17 := GenS1{F0: 9223372036854775807, F1: (false || (337 >= -32768))}
            v18 := GenS1{F0: -1, F1: ((1.5 / (0.1 / 3.14159)) == float64(v4))}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    fmt.Println(v4, v3)
    fmt.Println(v1, v1, v4)
    for fi19 := 0; fi19 < 11; fi19++ {
        switch ("" + "a") {
        case "hello":
            var v20 uint32 = 348
            _ = v20
        case "":
            var v21 float32 = float32(0.1)
            _ = v21
        default:
            var v22 int16 = 42
            _ = v22
        }
        _ = fi19
    }
    var v23 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
