package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    var v1 int32 = 42
    var v2 string = ("a" + "\n\t")
    v3 := []uint32{}
    v4 := GenS1{F0: (-387 < 1000), F1: v2, F2: v2}
    v2 = (v2 + v2)
    if (!false) {
        fmt.Println(v1)
        v5 := GenS1{F0: (false && (!((!true) || true))), F1: "a", F2: ""}
        if (!(0 == 42)) {
            if (!(true || (!(false || false)))) {
                v1 = 2147483647
                switch 83 {
                case 131:
                    var v6 float64 = (float64((0.0 / 1e10)) * 3.14159)
                    _ = v6
                }
            }
            fmt.Println(v2)
        }
        _ = v5
    } else {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    }
    var v7 int8 = 127
    var v8 uint64 = 31
    switch v2 {
    case "":
        v3 = append(v3, 0)
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: false, F1: "foo bar", F2: "hello"}
            v2 = v2
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            v1 += (v1 - v1)
            v8 += 0
            fmt.Println(fi9, v2)
            _ = fi12
        }
        _ = fi9
    }
    v13 := GenS2{F0: -1, F1: (v1 & (v1 ^ -2147483648))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
