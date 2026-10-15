package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }

func main() {
    var v1 float32 = (3.14159 / ((0.1 - float32(3.14159)) / 0.1))
    var v2 int8 = -1
    var v3 string = "hello"
    var v4 int = -9223372036854775808
    v1 = v1
    var v5 int16 = 42
    v2 -= v2
    for fi6 := 0; fi6 < 1; fi6++ {
        v5 = int16(fi6)
        _ = fi6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint64 = 66
        if (91 >= 193) {
            v9 := []int{}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    v5 = -642
    var v10 uint64 = 24
    v5 += v5
    if false {
        v11 := GenS1{F0: v5, F1: 0, F2: 42}
        _ = v11
    }
    if false {
        v5 -= v5
        v12 := []int{int(v2), (v4 % ((v4) | 1)), -2147483648}
        if (((v10 << ((uint(v10)) & 63)) > v10) && ((("" + "a") + v3) != v3)) {
            if true {
                fmt.Println(v1, v1)
            }
            if (true && true) {
                _ = len(v12)
            } else {
                switch v3 {
                case "":
                    v13 := GenS1{F0: (^(v5)), F1: 153, F2: 255}
                    _ = v13
                }
                if ((!(true || false)) || (1 > 2)) {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v4 -= fi14
                        v10 -= v10
                        _ = fi14
                    }
                    var v15 int = v4
                    v10 = v10
                    _ = v15
                } else {
                    fmt.Println(v3)
                }
            }
            fmt.Println(v2)
        } else {
            fmt.Println(v5, v3)
        }
        _ = v12
    }
    v3 = v3
    var v16 uint64 = (v10 % (((v10 + (v10 * 4294967295))) | 1))
    v17 := GenS1{F0: (^(v5)), F1: 42, F2: 145}
    var v18 uint8 = 1
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int16 = (^(v5))
        v21 := []int{(^(fi19)), (fi19 / ((int(v10)) | 1)), (fi19 - (fi19 << ((868) & 63)))}
        v2 = 9
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
