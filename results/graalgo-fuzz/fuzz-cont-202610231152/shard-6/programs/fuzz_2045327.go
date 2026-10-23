package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }

func main() {
    var v1 int8 = 104
    var v2 int = 0
    v3 := make(map[string]int)
    var v4 float64 = 1e10
    v5 := make(map[string]bool)
    v2 = 127
    var v6 uint8 = 175
    v1 += v1
    v1 += (v1 << ((65535) & 63))
    v7 := GenS1{F0: v2}
    if false {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS2{F0: "a"}
            _ = fi8
            _ = v9
        }
        v1 = v1
        for fi10 := 0; fi10 < 16; fi10++ {
            _ = len(v5)
            var v11 float32 = 3.14159
            for fi12 := 0; fi12 < 12; fi12++ {
                v3[("" + "a")] = fi10
                switch "foo bar" {
                case "":
                    v13 := GenS2{F0: ("hello" + "\n\t")}
                    _ = v13
                case "\n\t":
                    v2 += (fi10 >> ((831) & 63))
                default:
                    var v14 int64 = 0
                    _ = v14
                }
                v15 := []int{}
                _ = fi12
                _ = v15
            }
            _ = fi10
            _ = v11
        }
    } else {
        if (true || false) {
            _ = len(v5)
        } else {
            switch 32768 {
            case -943:
                for fi16 := 0; fi16 < 6; fi16++ {
                    v2 -= v2
                    _ = fi16
                }
            case 32768:
                v4 = v4
            }
            v6 -= (v6 & 42)
        }
        var v17 uint8 = (v6 * v6)
        _ = v17
    }
    switch (^(v2)) {
    case -48:
        fmt.Println(v2, v2)
    case 32767:
        v18 := []int{v2, (v2 | -2147483648)}
        _ = v18
    }
    v4 -= v4
    v19 := []int{-724, (v2 & v2), v2}
    v4 = v4
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
