package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint16 = 294
    var v2 float64 = 0.1
    v3 := make(map[uint32]string)
    var v4 float64 = (0.1 * (1e10 / 1.0))
    if true {
        var v5 int32 = -1
        v2 = (float64(1.5) + v2)
        _ = v5
    } else {
        for fi6 := 0; fi6 < 14; fi6++ {
            switch (fi6 * fi6) {
            case -466:
                var v7 float64 = v2
                _ = v7
            case -109:
                v8 := GenS1{F0: 54}
                _ = v8
            case 128:
                v1 += (v1 & v1)
            }
            if false {
                var v9 int = (fi6 & fi6)
                v10 := []int{(fi6 << ((699) & 63)), (^(v9))}
                _ = v10
                _ = v9
            }
            v1 = 1
            _ = fi6
        }
    }
    switch 4294967295 {
    case 32767:
        v11 := []int{-224}
        _ = v11
    default:
        var v12 int8 = -1
        _ = v12
    }
    v3[2147483647] = "a"
    var v13 int = -563
    v14 := GenS1{F0: 100}
    var v15 int64 = 9223372036854775807
    switch "foo bar" {
    case "foo bar":
        var v16 float64 = 0.1
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 9; fi17++ {
            v15 -= int64(v13)
            _ = fi17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
