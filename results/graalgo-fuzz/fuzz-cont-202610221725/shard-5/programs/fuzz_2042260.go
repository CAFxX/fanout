package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 float64 = 1e10
    var v4 uint = 2
    v3 += 0.1
    var v5 int16 = -402
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 32767}
        v8 := GenS2{F0: ("\n\t" + "a")}
        v9 := []int{fi6, fi6, fi6, (fi6 + fi6)}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 bool = false
    var v11 bool = (2 == 255)
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int = (fi12 - (fi12 >> (((v4 * 4294967295)) & 63)))
        switch fi12 {
        case -545:
            var v14 int8 = -128
            _ = v14
        default:
            if (int(v5) == 9223372036854775807) {
                v2["foo bar"] = "\n\t"
            }
        }
        v2[("a" + "foo bar")] = "a"
        _ = fi12
        _ = v13
    }
    v1[1] = 2147483647
    var v15 int8 = 0
    v16 := GenS2{F0: "a"}
    v17 := []int{128}
    v18 := GenS2{F0: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
