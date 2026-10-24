package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint64; F1 bool; F2 uint64 }
func genFn0() uint8 {
    var v6 uint16 = 2
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 14; fi7++ {
            if (1 < 1) {
                var v8 bool = (false || (116 == 255))
                _ = v8
            } else {
                v6 = v6
            }
            fmt.Println(v6, v6)
            var v9 bool = false
            _ = fi7
            _ = v9
        }
    }
    _ = v6
    return 255
}

func main() {
    var v1 int = 76
    var v2 string = (("foo bar" + (("" + (("hello" + "hello") + "\n\t")) + "\n\t")) + ("\n\t" + ("a" + "\n\t")))
    v3 := make(map[uint32]int)
    var v4 int = (^(v1))
    v5 := make(map[int]string)
    v10 := GenS1{F0: 0, F1: v1}
    switch -9223372036854775808 {
    case 4294967295:
        var v11 float32 = (1.5 * (((float32(3.14159) * ((((0.1 * 1.5) / 1.0) - (((1.5 + 0.1) - float32(-2.25)) + ((0.0 / 1.0) * 0.1))) * float32(1.5))) * 1.5) - (3.14159 / (0.1 + -2.25))))
        _ = v11
    }
    v12 := GenS2{F0: 248, F1: true, F2: 263}
    _ = len(v3)
    if true {
        var v13 float64 = float64(3.14159)
        _ = v13
    }
    v4 += -9223372036854775808
    var v14 bool = (-32768 <= 564)
    var v15 bool = v14
    v4 -= v1
    v14 = (!false)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
