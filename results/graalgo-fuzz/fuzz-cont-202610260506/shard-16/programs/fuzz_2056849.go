package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float32 = -2.25
    var v2 int = 704
    var v3 int = v2
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    var v6 uint = 42
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: uint16(v6), F1: 0}
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int8 = -1
    for fi11 := 0; fi11 < 4; fi11++ {
        v3 = v3
        var v12 uint8 = 127
        _ = fi11
        _ = v12
    }
    var v13 bool = false
    switch v3 {
    case -9223372036854775808:
        for fi14 := 0; fi14 < 15; fi14++ {
            fmt.Println(fi14, v1, v6)
            switch fi14 {
            case 1:
                v1 += v1
            case 32767:
                switch -381 {
                case 32768:
                    v1 += 0.1
                }
            }
            delete(v4, ("\n\t" + ("" + "")))
            _ = fi14
        }
    case 18:
        v15 := GenS1{F0: 1, F1: 32767}
        _ = v15
    default:
        v3 = 0
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v13 = (!false)
        _ = fi16
    }
    var v17 int8 = v10
    v3 += 886
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS2{F0: -2147483648}
            _ = fi19
            _ = v20
        }
        var v21 bool = (-1 < 999)
        fmt.Println(v1, v17, v10)
        _ = fi18
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
