package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = -2
    v3 := make(map[uint32]bool)
    var v4 uint16 = 255
    var v5 int32 = 2147483647
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{fi7, fi7, 9223372036854775807}
        _ = fi7
        _ = v8
    }
    v5 -= (v5 / ((v5) | 1))
    v9 := []int{9223372036854775807, 2147483647, 4294967295, -647}
    if false {
        v4 = (v4 & (v4 % (((v4 ^ (^(v4)))) | 1)))
        v2 = 420
        switch -137 {
        case 509:
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := GenS1{F0: 829}
                var v12 int8 = -88
                _ = fi10
                _ = v11
                _ = v12
            }
        case 2147483647:
            v2 = v2
        case -469:
            switch "hello" {
            case "":
                v5 = (v5 ^ v5)
            case "\n\t":
                var v13 string = "hello"
                _ = v13
            }
        default:
            if true {
                if (-1 <= int8(v5)) {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 string = "hello"
                        _ = fi14
                        _ = v15
                    }
                    v4 -= v4
                }
            }
        }
    } else {
        var v16 uint8 = 1
        v9 = append(v9, -1)
        switch 646 {
        case 430:
            var v17 float64 = 1.5
            _ = v17
        }
        _ = v16
    }
    _ = len(v6)
    v18 := GenS2{F0: 127}
    var v19 string = "a"
    var v20 uint = 482
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
