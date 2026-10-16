package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint; F2 string }
func genFn0() uint64 {
    switch "\n\t" {
    case "foo bar":
        var v6 int32 = 693
        _ = v6
    case "\n\t":
        var v7 float64 = 0.1
        _ = v7
    default:
        v8 := GenS1{F0: 0}
        _ = v8
    }
    v9 := GenS2{F0: -1, F1: 255, F2: "\n\t"}
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{377, 42, 826, 0}
    var v3 float64 = (1e10 - -2.25)
    v4 := []int{-15, 32767}
    var v5 uint8 = 2
    v10 := GenS2{F0: -599, F1: 833, F2: "a"}
    var v11 string = "foo bar"
    v5 += v5
    if (v3 >= 1e10) {
        v3 += 0.1
        for fi12 := 0; fi12 < 17; fi12++ {
            v11 = ("hello" + v11)
            v11 = ("\n\t" + ("a" + ""))
            var v13 float32 = 1.5
            _ = fi12
            _ = v13
        }
    } else {
        var v14 float64 = 0.0
        _ = v14
    }
    v5 = (v5 * 255)
    var v15 int8 = 4
    v3 = (v3 - v3)
    fmt.Println(v5, v3, v11)
    if (false || false) {
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 uint = 1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi16
            _ = v17
        }
        var v18 float32 = -2.25
        _ = v18
    } else {
        v15 = v15
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := v1[v11]
            fmt.Println(v5, v20, v11)
            var v21 string = (v20 + v11)
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
