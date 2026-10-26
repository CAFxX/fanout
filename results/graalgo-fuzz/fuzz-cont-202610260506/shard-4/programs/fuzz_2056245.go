package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint = 42
    var v2 int16 = -32768
    var v3 uint = 567
    v4 := make(map[string]bool)
    if (((!false) && true) || false) {
        var v5 uint16 = 401
        v6 := GenS1{F0: true}
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            if (float64(3.14159) < (((float64(-2.25) * (1e10 + float64((3.14159 * ((1e10 / 0.1) - 1e10))))) / 0.1) - 1e10)) {
                if ((0.0 / ((3.14159 - float64(-2.25)) / 1.5)) != ((0.1 + (((3.14159 / 1.0) / 1e10) + 1e10)) - -2.25)) {
                    v2 += (v2 >> ((61) & 63))
                    if true {
                        v1 = v1
                    }
                    v8 := GenS1{F0: (0 > -204)}
                    _ = v8
                } else {
                    if (!(128 > 42)) {
                        v1 -= v1
                        v9 := []int{}
                        _ = v9
                    }
                }
                var v10 bool = (!true)
                var v11 uint64 = 476
                _ = v10
                _ = v11
            }
            switch (fi7 % ((int(v2)) | 1)) {
            case 128:
                v12 := GenS2{F0: 42, F1: -937}
                _ = v12
            case 633:
                if true {
                    var v13 int64 = 19
                    _ = v13
                }
            }
            _ = fi7
        }
    }
    v14 := GenS2{F0: 128, F1: 739}
    v2 = v2
    v2 += (v2 ^ v2)
    v1 += v1
    fmt.Println(v3, v2)
    fmt.Println(v2, v1)
    var v15 uint8 = 127
    var v16 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
