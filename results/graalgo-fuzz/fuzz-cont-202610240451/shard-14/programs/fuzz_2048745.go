package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    var v1 float32 = (0.1 - 1e10)
    var v2 uint32 = 2147483648
    var v3 uint = 42
    var v4 int = 9223372036854775807
    v5 := []int{v4, 1, 1, 50}
    var v6 float64 = (0.0 + 1e10)
    v4 -= 733
    v2 -= (^(v2))
    v3 = (v3 & 949)
    v7 := GenS1{F0: 1, F1: (v4 / (((v4 - int(v2))) | 1)), F2: 893}
    if false {
        v1 += (v1 * v1)
        v8 := GenS1{F0: 937, F1: v4, F2: -2147483648}
        if (1 >= 87) {
            v9 := GenS2{F0: -1, F1: (^(v2))}
            _ = v9
        }
        _ = v8
    }
    if ((^(v2)) == (v2 ^ v2)) {
        v10 := GenS1{F0: 0, F1: 145, F2: 42}
        if (1e10 <= ((0.1 - v1) - ((v1 - float32(3.14159)) - 1e10))) {
            v11 := GenS1{F0: -358, F1: (v4 / (((v4 & v4)) | 1)), F2: -1}
            _ = v11
        } else {
            var v12 uint8 = 128
            if (!((!true) && false)) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v13 int32 = 1
                _ = v13
            } else {
                v6 += (float64(0.1) / (0.1 + v6))
                var v14 int32 = -1
                v1 -= (v1 * 3.14159)
                _ = v14
            }
            _ = v12
        }
        v15 := GenS2{F0: -113, F1: v2}
        _ = v10
        _ = v15
    }
    fmt.Println(v1, v3)
    v6 -= (v6 + v6)
    v6 = (v6 + v6)
    fmt.Println(v3)
    fmt.Println(v6, v3, v3)
    v16 := GenS1{F0: -520, F1: -148, F2: -2147483648}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        _ = len(v5)
        v3 = v3
        _ = fi17
    }
    fmt.Println(v2, v6)
    fmt.Println(v6, v2, v2)
    switch v4 {
    case 4294967295:
        switch (v4 / ((32768) | 1)) {
        case -2147483648:
            switch 4294967295 {
            case 0:
                if len(v5) > 0 {
                    v5[0] = int(v2)
                }
            }
        default:
            var v18 int8 = 100
            _ = v18
        }
    case 32767:
        var v19 int16 = -26
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
