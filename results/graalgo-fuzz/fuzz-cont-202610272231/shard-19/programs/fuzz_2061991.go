package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int = 42
    var v2 uint8 = 255
    var v3 uint64 = 42
    var v4 int16 = 1
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int64 = 9223372036854775807
        v1 -= v1
        if (true || (!true)) {
            if false {
                v7 := []int{}
                v8 := GenS1{F0: 4294967295}
                _ = v7
                _ = v8
            }
            v1 += (fi5 / ((-481) | 1))
            v1 += (v1 % ((v1) | 1))
        }
        _ = fi5
        _ = v6
    }
    if (-93 <= 127) {
        var v9 int64 = -950
        v4 += v4
        var v10 bool = true
        _ = v10
        _ = v9
    }
    var v11 uint16 = 65535
    if (255 > v11) {
        var v12 int64 = 9223372036854775807
        _ = v12
    }
    v4 -= 145
    if (!true) {
        v11 -= v11
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := []int{}
            switch fi13 {
            case 0:
                if (0.1 <= (0.1 + (1.5 - 0.1))) {
                    var v15 uint8 = v2
                    _ = v15
                }
            case 128:
                v14 = append(v14, 32767)
            default:
                _ = len(v14)
            }
            _ = fi13
            _ = v14
        }
        v16 := GenS2{F0: -104}
        _ = v16
    }
    var v17 int16 = (v4 & v4)
    var v18 int = v1
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 int8 = 1
        fmt.Println(v1, v20)
        _ = fi19
        _ = v20
    }
    var v21 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
