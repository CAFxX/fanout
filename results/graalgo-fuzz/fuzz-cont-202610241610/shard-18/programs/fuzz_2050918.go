package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int8; F2 uint }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]string)
    var v3 int32 = 0
    for fi4 := 0; fi4 < 8; fi4++ {
        v1 += v1
        _ = fi4
    }
    if true {
        v1 -= (^(v1))
        var v5 string = "hello"
        _ = v5
    }
    fmt.Println(v3, v3)
    fmt.Println(v3, v3)
    v6 := GenS1{F0: -1, F1: -1, F2: "a"}
    if true {
        v1 = v1
        v7 := []int{v1, -372, (v1 >> ((4294967295) & 63))}
        if false {
            switch (("hello" + ((("\n\t" + ("foo bar" + ("a" + "a"))) + "hello") + "\n\t")) + "\n\t") {
            case "foo bar":
                fmt.Println(v1, v1)
            case "a":
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 int16 = -32768
                    _ = fi8
                    _ = v9
                }
            default:
                v3 = (v3 + 886)
            }
            v10 := GenS1{F0: 0, F1: -107, F2: "a"}
            _ = v10
        } else {
            var v11 int = (-(v1))
            _ = v11
        }
        _ = v7
    }
    v12 := GenS2{F0: 575}
    switch 601 {
    case 128:
        v13 := GenS3{F0: 0, F1: -1, F2: 2}
        _ = v13
    case 464:
        v14 := GenS3{F0: 90, F1: 68, F2: 2}
        _ = v14
    default:
        var v15 uint = 255
        _ = v15
    }
    if (!(-1 > 0)) {
        if true {
            var v16 uint8 = 127
            v16 = (^(v16))
            fmt.Println(v1)
            _ = v16
        }
        v17 := GenS2{F0: 32767}
        _ = v17
    }
    var v18 float64 = ((0.1 + 0.0) / 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v6
}
