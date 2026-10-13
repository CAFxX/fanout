package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }

func main() {
    var v1 uint8 = 2
    var v2 int = 0
    var v3 float32 = 3.14159
    var v4 int32 = -514
    var v5 int16 = -404
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := []int{(v2 + (fi6 / ((32767) | 1)))}
            if true {
                v2 = (v2 * (-(fi6)))
                var v9 string = "hello"
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if false {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := []int{-1}
            var v12 float64 = 1.5
            for fi13 := 0; fi13 < 17; fi13++ {
                v3 += (1.5 / v3)
                v14 := GenS1{F0: "a", F1: (v4 - (v4 | (v4 / ((v4) | 1)))), F2: 65535}
                _ = fi13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        var v15 uint = 589
        _ = v15
    }
    var v16 float32 = ((v3 - ((((v3 - (float32((1.5 - 0.1)) / (v3 * ((float32(1.5) - v3) * (((v3 + v3) + 0.1) * (3.14159 * 0.1)))))) - ((float32((((-2.25 * 1e10) / 3.14159) + 0.0)) * (1.5 + v3)) + (float32(3.14159) * (float32(1.5) - -2.25)))) / (v3 + v3)) + float32(0.1))) + v3)
    fmt.Println(v3, v4)
    v17 := []int{(v2 ^ (v2 ^ v2))}
    var v18 uint16 = 173
    v19 := GenS1{F0: ("hello" + "a"), F1: -317, F2: 428}
    if (!true) {
        fmt.Println(v4, v2, v1)
    } else {
        v17 = append(v17, 128)
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        v4 += 0
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = fi20
    }
    if (255 < (^(v1))) {
        var v21 int = v2
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int64 = 1
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    var v24 int8 = -128
    v16 -= (1.5 * (v16 * -2.25))
    v1 = (v1 / ((v1) | 1))
    v3 += v3
    v24 = (^(v24))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
