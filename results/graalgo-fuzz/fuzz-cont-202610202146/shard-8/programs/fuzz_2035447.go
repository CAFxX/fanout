package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    var v1 int16 = 665
    var v2 uint16 = 2
    var v3 int8 = 81
    v4 := []uint32{894, 42}
    v1 -= int16(v2)
    v2 = v2
    fmt.Println(v3)
    v3 -= -1
    v5 := GenS1{F0: ("\n\t" + "hello"), F1: "a"}
    switch 127 {
    case -497:
        var v6 bool = (-658 >= -2147483648)
        _ = v6
    case 0:
        v2 += (v2 % ((v2) | 1))
    case 128:
        v1 -= -1
    }
    var v7 uint = 632
    var v8 int64 = -630
    var v9 float32 = 0.0
    var v10 float32 = (3.14159 * (((1.5 - (v9 * 0.0)) + (float32((3.14159 / 0.1)) * 1e10)) - v9))
    if len(v4) > 0 {
        v4[0] = 1
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v9 = ((v10 - v10) - ((v10 + (float32(1.5) + ((v9 / (((v9 / v9) * (-2.25 - (((v10 - v9) + (v9 - v10)) / v10))) + (v10 / v10))) / v10))) - v9))
        _ = fi11
    }
    if (4294967295 >= 4294967295) {
        v8 += 805
        var v12 uint8 = 2
        _ = v12
    } else {
        if true {
            var v13 int64 = 42
            if false {
                if (v9 <= (v10 + float32(0.1))) {
                    v13 += v8
                    var v14 float64 = 3.14159
                    v3 = v3
                    _ = v14
                }
            }
            _ = v13
        }
        v4 = append(v4, 441)
    }
    if true {
        if ((2 == 219) || (!true)) {
            v8 = v8
            switch -1 {
            case 1:
                for fi15 := 0; fi15 < 8; fi15++ {
                    fmt.Println(v2, v3, v2)
                    _ = fi15
                }
            }
        }
        v8 = -1
    }
    var v16 uint = v7
    if false {
        for fi17 := 0; fi17 < 16; fi17++ {
            v9 += ((((1e10 - v10) - v10) * (((float32(0.1) / 0.1) / (float32(((3.14159 - 3.14159) + (1.5 + (-2.25 + 0.0)))) / v9)) * v9)) * (3.14159 * 1e10))
            v16 = v16
            _ = fi17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
