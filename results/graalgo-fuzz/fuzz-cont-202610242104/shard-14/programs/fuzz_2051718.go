package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := []string{}
    var v4 uint8 = 111
    v5 := GenS1{F0: (("" + ("hello" + "\n\t")) + ("hello" + "hello"))}
    switch -2147483648 {
    case 32768:
        v4 -= (v4 ^ v4)
    case 42:
        if ((!(65535 < 42)) && ((float64((1e10 + 3.14159)) == ((-2.25 - (0.0 * -2.25)) * float64(3.14159))) && (true && true))) {
            v6 := v2["\n\t"]
            _ = v6
        }
    default:
        if false {
            var v7 int64 = -292
            switch 285 {
            case -895:
                var v8 int8 = 1
                _ = v8
            }
            v9 := GenS2{F0: (v4 + (v4 << ((478) & 63))), F1: 0}
            _ = v7
            _ = v9
        } else {
            v4 = 255
        }
    }
    v1[0] = ("a" + "\n\t")
    var v10 float64 = (0.0 + -2.25)
    if (-2.25 != 1e10) {
        v10 += (v10 + v10)
    }
    var v11 uint = 669
    var v12 uint64 = 1
    var v13 int16 = 42
    var v14 uint8 = (v4 & 244)
    v15 := GenS2{F0: 128, F1: v12}
    v16 := v2[""]
    var v17 uint32 = 0
    v3 = append(v3, "")
    fmt.Println(v13, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
