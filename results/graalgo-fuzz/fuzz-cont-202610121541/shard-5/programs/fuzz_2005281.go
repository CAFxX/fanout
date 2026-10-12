package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 int8 = -105
    var v4 int64 = -831
    var v5 float64 = (0.0 / 1.5)
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            if (false && ((!true) && false)) {
                var v8 string = ("a" + "")
                _ = v8
            } else {
                v5 = v5
            }
            v3 = v3
            if (65535 > 471) {
                fmt.Println(fi6)
                fmt.Println(fi7)
            } else {
                if false {
                    var v9 uint8 = 127
                    fmt.Println(v4)
                    v3 = (v3 + (v3 << ((18446744073709551615) & 63)))
                    _ = v9
                }
            }
            _ = fi7
        }
        v10 := GenS1{F0: 170}
        v11 := GenS1{F0: 2}
        _ = fi6
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 439, F1: 0}
    v3 = (v3 & (v3 >> ((65535) & 63)))
    var v13 float64 = (v5 + v5)
    if (!true) {
        switch ((("foo bar" + "\n\t") + "a") + "") {
        case "":
            if (28 < 69) {
                fmt.Println(v3)
            }
        case "a":
            fmt.Println(v3, v5)
        default:
            var v14 uint64 = 18446744073709551615
            _ = v14
        }
        switch 0 {
        case 2147483647:
            var v15 uint32 = 42
            _ = v15
        case 42:
            var v16 float64 = v13
            _ = v16
        default:
            v17 := []int{0, 127, 32768}
            _ = v17
        }
    }
    v18 := []int{128}
    _ = len(v1)
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS1{F0: 127}
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
