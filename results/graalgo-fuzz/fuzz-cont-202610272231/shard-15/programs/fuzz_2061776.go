package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint8 = 2
    var v2 uint16 = 853
    var v3 float32 = ((float32((0.0 / (float64(1.5) + 1.5))) + -2.25) + (1e10 * 0.0))
    var v4 float64 = -2.25
    v5 := []int{}
    v3 -= 0.0
    _ = len(v5)
    v6 := GenS1{F0: 0}
    v4 -= (-2.25 - v4)
    var v7 string = "\n\t"
    v1 += (v1 >> ((18446744073709551615) & 63))
    v4 += float64(3.14159)
    _ = len(v5)
    if true {
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: 18446744073709551615}
                _ = fi8
                _ = v9
            }
            v7 = ((v7 + (v7 + ("\n\t" + v7))) + "a")
        } else {
            var v10 uint8 = (v1 / ((v1) | 1))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v11 := GenS1{F0: 1}
            _ = v10
            _ = v11
        }
    } else {
        var v12 int64 = 0
        var v13 int16 = 1
        for fi14 := 0; fi14 < 10; fi14++ {
            if true {
                var v15 uint8 = (v1 + 1)
                v16 := GenS1{F0: 65535}
                var v17 string = v7
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    v2 = v2
    v5 = append(v5, 127)
    for fi18 := 0; fi18 < 9; fi18++ {
        if len(v5) > 0 {
            v5[0] = fi18
        }
        v2 = (v2 - v2)
        _ = fi18
    }
    var v19 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
