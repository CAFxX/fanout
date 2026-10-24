package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 int32; F2 uint16 }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 1.5
    var v3 uint32 = 0
    v4 := []int{675, 4294967295, -250, 128}
    var v5 uint = 65535
    var v6 int16 = 42
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v2, v6, v2)
        var v8 uint8 = 128
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 42
    if (0 < -1) {
        var v10 int32 = 711
        var v11 uint8 = 119
        _ = len(v1)
        _ = v10
        _ = v11
    }
    v5 += (^(v5))
    fmt.Println(v2)
    if (v9 > uint16(v5)) {
        if false {
            fmt.Println(v6, v2)
            v12 := GenS1{F0: -177, F1: uint32(v9)}
            _ = v12
        } else {
            v9 = v9
            if false {
                var v13 bool = false
                _ = v13
            }
        }
        switch 2147483647 {
        case 143:
            v2 += v2
        case -1:
            if (("hello" + ("hello" + "foo bar")) <= "") {
                _ = len(v4)
                v14 := GenS2{F0: (((1 != 2) && (0 == (^(v6)))) && (!((!(false && (!true))) || (v9 <= (v9 % (((v9 / ((180) | 1))) | 1)))))), F1: int32(v6), F2: uint16(v3)}
                var v15 float64 = (((v2 / 0.1) - 3.14159) + ((v2 * ((0.1 + (float64(float32(float64(1.5))) - float64((0.1 + 1e10)))) / float64(1.5))) / (v2 - (((((0.1 - -2.25) - (v2 * 1.5)) - (((3.14159 - (1e10 + (v2 * v2))) + v2) * v2)) * (v2 - v2)) / -2.25))))
                _ = v14
                _ = v15
            }
        default:
            _ = len(v4)
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            switch fi16 {
            case 32768:
                v17 := v1[42]
                _ = v17
            }
            _ = fi16
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 9223372036854775807
        }
        v5 = 65535
        v18 := GenS2{F0: (false || (true || false)), F1: 2147483647, F2: (v9 & v9)}
        _ = v18
    }
    var v19 int16 = v6
    v20 := GenS3{F0: 0, F1: 0}
    v21 := GenS3{F0: 750, F1: 19}
    v22 := GenS3{F0: -2147483648, F1: int8(v6)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
