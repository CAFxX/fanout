package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
func genFn0() uint64 {
    switch "hello" {
    case "foo bar":
        v6 := GenS1{F0: 512}
        _ = v6
    default:
        var v7 uint8 = 127
        _ = v7
    }
    v8 := []int{1}
    var v9 bool = (!false)
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []string{"\n\t", "\n\t", "\n\t"}
    v2 := make(map[string]string)
    v3 := []string{}
    var v4 float64 = 0.1
    var v5 int8 = 74
    v5 -= v5
    fmt.Println(v5, v5)
    for fi10 := 0; fi10 < 14; fi10++ {
        switch fi10 {
        case -2147483648:
            v2["foo bar"] = ""
        }
        var v11 bool = true
        _ = fi10
        _ = v11
    }
    var v12 int16 = -14
    var v13 int16 = -32768
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int64 = int64(v13)
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(fi14)
            switch fi14 {
            case 42:
                v13 += 42
            }
            if (true || false) {
                var v17 bool = ((v15 / (((v15 ^ (v15 >> ((255) & 63)))) | 1)) != 9223372036854775807)
                var v18 float32 = 1e10
                _ = v17
                _ = v18
            } else {
                v13 = 32767
                v19 := GenS1{F0: 864}
                _ = v19
            }
            _ = fi16
        }
        if true {
            fmt.Println(v15, v4, v15)
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "")
            }
            v20 := GenS2{F0: v15, F1: -76, F2: 2}
            _ = v20
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v4)
    var v21 string = ("\n\t" + "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
