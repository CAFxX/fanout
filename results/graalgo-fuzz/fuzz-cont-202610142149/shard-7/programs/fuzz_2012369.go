package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 string; F1 string }

func main() {
    var v1 uint = 255
    v2 := make(map[uint32]bool)
    var v3 bool = true
    v4 := []string{"\n\t"}
    var v5 int64 = 9223372036854775807
    var v6 float32 = 3.14159
    if (-2147483648 >= -2147483648) {
        var v7 int = -1
        _ = v7
    }
    v8 := GenS1{F0: 2147483647, F1: ((("a" + "a") + "\n\t") + "foo bar"), F2: -9223372036854775808}
    fmt.Println(v1, v1, v1)
    v6 += float32(-2.25)
    v6 = (v6 * float32(0.1))
    var v9 int64 = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v10 int64 = (v5 | (v9 >> (((^(v1))) & 63)))
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := v2[2147483648]
        if v3 {
            v1 -= v1
            v13 := v2[4294967295]
            var v14 uint16 = 355
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: 1, F1: "\n\t", F2: (fi11 * (fi11 - fi11))}
        _ = fi11
        _ = v12
        _ = v15
    }
    switch -604 {
    case 145:
        var v16 float32 = float32((float64(v6) / (1e10 / 0.1)))
        _ = v16
    case -1:
        var v17 bool = v3
        _ = v17
    case -408:
        switch 1 {
        case 0:
            var v18 bool = v3
            _ = v18
        case 127:
            v19 := GenS1{F0: 2147483647, F1: ("foo bar" + ("" + "a")), F2: -78}
            _ = v19
        }
    default:
        v10 -= 9223372036854775807
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v20 int32 = 42
    v3 = v3
    var v21 uint16 = 42
    if (v3 || false) {
        _ = len(v4)
        v22 := GenS1{F0: (v20 >> (((v1 + v1)) & 63)), F1: "a", F2: 830}
        _ = v22
    }
    v23 := GenS2{F0: ("hello" + "hello"), F1: ("a" + "")}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
