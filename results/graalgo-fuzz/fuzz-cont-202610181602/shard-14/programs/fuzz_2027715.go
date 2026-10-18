package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 bool }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    var v3 uint = 1
    var v4 int8 = 1
    var v5 float32 = -2.25
    switch 2147483647 {
    case 336:
        if ((770 <= 2147483648) && (1 != 641)) {
            v3 = (v3 << (((v3 % (((v3 + v3)) | 1))) & 63))
            var v6 float64 = (3.14159 * 1e10)
            _ = v6
        }
    case 9223372036854775807:
        for fi7 := 0; fi7 < 18; fi7++ {
            v5 += float32(3.14159)
            v1 = "a"
            _ = fi7
        }
    }
    var v8 uint = v3
    if (611 != 42) {
        var v9 int16 = int16(v4)
        var v10 uint64 = 18446744073709551615
        _ = v10
        _ = v9
    }
    var v11 float64 = (1.5 / 3.14159)
    for fi12 := 0; fi12 < 1; fi12++ {
        if (!false) {
            var v13 uint16 = 32767
            _ = v13
        } else {
            switch v1 {
            case "":
                switch fi12 {
                case -7:
                    for fi14 := 0; fi14 < 12; fi14++ {
                        delete(v2, 698)
                        var v15 int16 = -169
                        v5 = 1e10
                        _ = fi14
                        _ = v15
                    }
                case 0:
                    var v16 uint = 630
                    _ = v16
                }
            }
        }
        v1 = ""
        var v17 float32 = 0.1
        _ = fi12
        _ = v17
    }
    v3 = (v3 << ((v8) & 63))
    v18 := GenS1{F0: 42}
    if (2147483648 > 2147483647) {
        var v19 float64 = v11
        v20 := []int{0}
        if false {
            for fi21 := 0; fi21 < 11; fi21++ {
                fmt.Println(v5, v3, v4)
                var v22 bool = true
                var v23 uint32 = 2147483647
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        _ = v19
        _ = v20
    }
    v11 += v11
    fmt.Println(v4, v5)
    delete(v2, 834)
    v8 = (^(v8))
    for fi24 := 0; fi24 < 1; fi24++ {
        v11 += -2.25
        v25 := GenS2{F0: int8(v3), F1: 83, F2: (((((^(v3)) == 278) || ((-2.25 >= -2.25) && false)) || (!false)) && (!true))}
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
