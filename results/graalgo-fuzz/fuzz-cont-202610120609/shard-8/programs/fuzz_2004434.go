package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint = 255
    var v2 float32 = float32((((-2.25 * -2.25) * (0.0 / 1.5)) - -2.25))
    var v3 uint16 = 1
    v4 := make(map[uint32]bool)
    var v5 int32 = -1
    var v6 uint32 = 0
    fmt.Println(v1, v2)
    switch -1 {
    case 42:
        v6 -= (v6 & (v6 * (v6 | (v6 / (((v6 / ((2147483647) | 1))) | 1)))))
    case 32767:
        if true {
            var v7 float64 = 1.5
            v2 = v2
            var v8 uint64 = 518
            _ = v7
            _ = v8
        }
    default:
        v3 = 0
    }
    v9 := GenS1{F0: (v3 >> (((v1 - v1)) & 63))}
    var v10 uint8 = uint8(v1)
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := []int{}
        var v13 int8 = 0
        if (v5 >= v5) {
            _ = len(v4)
            var v14 float64 = -2.25
            var v15 uint = v1
            _ = v14
            _ = v15
        } else {
            v16 := GenS1{F0: (v3 ^ 65535)}
            if (uint32(v1) > 279) {
                var v17 int64 = 9223372036854775807
                v6 -= (^(v6))
                _ = v17
            } else {
                var v18 bool = true
                var v19 int64 = 896
                switch fi11 {
                case 847:
                    var v20 string = "a"
                    _ = v20
                case -949:
                    v19 -= (-(v19))
                }
                _ = v18
                _ = v19
            }
            _ = len(v12)
            _ = v16
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v21 uint16 = 0
    v1 -= 0
    delete(v4, (v6 * v6))
    v2 = (v2 + -2.25)
    switch ("a" + ((("\n\t" + "foo bar") + "\n\t") + "\n\t")) {
    case "hello":
        var v22 int32 = (-(v5))
        _ = v22
    case "foo bar":
        v21 += 42
    }
    v2 -= 1.5
    var v23 bool = false
    if v23 {
        v4[v6] = (("hello" + "foo bar") != "foo bar")
    } else {
        v24 := GenS1{F0: v21}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
