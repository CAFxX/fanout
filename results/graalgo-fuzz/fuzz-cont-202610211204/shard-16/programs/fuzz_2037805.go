package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 255
    var v2 int16 = 1000
    var v3 string = ""
    v4 := make(map[uint32]uint32)
    v5 := v4[51]
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(v1)
        switch v3 {
        case "foo bar":
            var v7 uint8 = 121
            _ = v7
        default:
            switch fi6 {
            case 32768:
                v4[v5] = 65535
            case 536:
                switch 952 {
                case 32767:
                    fmt.Println(v1)
                }
            }
        }
        if false {
            v8 := []int{fi6}
            var v9 int8 = 118
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        if (((2 > 65535) || (2 == v1)) && false) {
            v3 = ((v3 + "a") + "foo bar")
        }
        var v11 float64 = 0.1
        _ = fi10
        _ = v11
    }
    if (false || false) {
        switch "\n\t" {
        case "":
            fmt.Println(v2, v5)
        case "\n\t":
            delete(v4, 0)
        }
    }
    var v12 uint32 = (^(v5))
    for fi13 := 0; fi13 < 12; fi13++ {
        v5 -= uint32(v1)
        _ = fi13
    }
    fmt.Println(v5, v1)
    if false {
        var v14 float64 = (1.5 * (1.5 + 0.0))
        v15 := []int{2147483647, 0, 42}
        var v16 float32 = ((1.5 * (0.1 + 3.14159)) * -2.25)
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint8 = 2
    v18 := v4[(v12 - (v12 | 4294967295))]
    var v19 uint16 = (v1 >> ((1) & 63))
    if (100 <= 0) {
        var v20 string = ("foo bar" + "foo bar")
        _ = v20
    } else {
        if ((475 == -2147483648) && true) {
            if ((v3 + v3) > v3) {
                v21 := []int{}
                _ = v21
            }
            v2 -= (^(v2))
        }
        v22 := v4[(v5 % (((v18 | (v5 / ((2147483648) | 1)))) | 1))]
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
