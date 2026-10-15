package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint32 }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int32 }

func main() {
    var v1 uint32 = 4294967295
    var v2 int = 521
    var v3 int = v2
    v4 := make(map[int]uint32)
    var v5 uint32 = 2147483648
    v3 -= 827
    v6 := GenS1{F0: false, F1: (255 != 2), F2: (v1 * v5)}
    var v7 uint8 = 1
    switch (v3 % ((-9223372036854775808) | 1)) {
    case -562:
        v8 := GenS1{F0: (!(-2.25 <= 3.14159)), F1: (255 != 42), F2: (v1 % ((v1) | 1))}
        _ = v8
    case -9223372036854775808:
        for fi9 := 0; fi9 < 11; fi9++ {
            if (!(!(42 > 65535))) {
                v10 := GenS2{F0: 42, F1: 775, F2: 133}
                var v11 float32 = (1.5 + (((1e10 + 1e10) / (3.14159 * 1.5)) * ((-2.25 + (float32((1e10 * 0.1)) + 0.0)) + ((3.14159 * 0.0) / (3.14159 * (((float32((float64((0.1 * 0.0)) + 0.1)) / (0.1 / 1.0)) / float32(((1e10 / float64(1.5)) / 1.0))) / -2.25))))))
                for fi12 := 0; fi12 < 19; fi12++ {
                    v5 = (v5 >> ((1) & 63))
                    _ = fi12
                }
                _ = v10
                _ = v11
            } else {
                fmt.Println(v1)
            }
            var v13 int64 = 9223372036854775807
            _ = fi9
            _ = v13
        }
    default:
        v3 = v3
    }
    fmt.Println(v3, v3, v5)
    var v14 uint32 = v1
    fmt.Println(v7, v1, v1)
    fmt.Println(v5)
    v15 := GenS3{F0: 989, F1: v7, F2: 0}
    fmt.Println(v2, v3)
    if (1 == 993) {
        var v16 float32 = (0.0 * 0.1)
        v17 := GenS3{F0: 782, F1: (v7 - (v7 << ((1) & 63))), F2: -1}
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
