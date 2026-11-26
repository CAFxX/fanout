package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v4 int16 = 1
    if true {
        switch -483 {
        case 4294967295:
            v4 -= 1
        }
        for fi5 := 0; fi5 < 11; fi5++ {
            switch ("\n\t" + "foo bar") {
            case "a":
                v4 = 65
            case "":
                var v6 int8 = -98
                _ = v6
            case "\n\t":
                if (!true) {
                    if (9223372036854775807 != 42) {
                        v7 := GenS1{F0: "foo bar", F1: 255, F2: 1}
                        var v8 float32 = (float32(float64(0.1)) / ((3.14159 / 3.14159) * 3.14159))
                        _ = v7
                        _ = v8
                    } else {
                        v4 = (v4 * v4)
                    }
                }
            }
            _ = fi5
        }
        v9 := v2[("\n\t" + "a")]
        _ = v9
    }
    v4 = (-(v4))
    var v10 int16 = (^(v4))
    v10 -= -653
    fmt.Println(v10)
    v4 -= v10
    v11 := []int{}
    var v12 uint16 = 813
    var v13 string = "a"
    fmt.Println(v12, v10)
    v13 = v13
    switch 4294967295 {
    case -2147483648:
        if (1 != -9223372036854775808) {
            delete(v2, "hello")
            var v14 uint16 = (v12 | v12)
            fmt.Println(v14, v12, v10)
            _ = v14
        } else {
            v4 = v10
            _ = len(v11)
        }
    case 127:
        v15 := v3[2]
        _ = v15
    }
    var v16 uint8 = 232
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
