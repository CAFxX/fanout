package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    var v1 int16 = 32767
    var v2 int8 = -1
    var v3 float64 = (3.14159 * float64(((1e10 - 1.5) * 3.14159)))
    v4 := make(map[uint32]string)
    var v5 uint64 = 544
    v6 := make(map[string]string)
    v7 := []int{}
    var v8 bool = false
    v9 := v4[972]
    var v10 int = 4294967295
    var v11 int16 = (v1 / ((v1) | 1))
    fmt.Println(v5, v11)
    v4[623] = ("foo bar" + (((v9 + v9) + "a") + (v9 + ((("" + (v9 + v9)) + "hello") + "\n\t"))))
    for fi12 := 0; fi12 < 13; fi12++ {
        if v8 {
            v4[0] = v9
        } else {
            fmt.Println(fi12, v8)
        }
        v13 := GenS1{F0: 42, F1: int16(v5), F2: 339}
        switch 0 {
        case 32767:
            _ = len(v6)
        case 311:
            v5 = v5
        case 32768:
            if true {
                for fi14 := 0; fi14 < 16; fi14++ {
                    fmt.Println(v8, fi14)
                    delete(v6, (v9 + v9))
                    _ = fi14
                }
                if ((v5 + v5) <= (v5 + (v5 + (v5 >> ((415) & 63))))) {
                    v15 := GenS2{F0: 975, F1: 168}
                    v7 = append(v7, v10)
                    _ = v15
                } else {
                    v10 = fi12
                    var v16 int16 = v11
                    _ = v16
                }
                for fi17 := 0; fi17 < 13; fi17++ {
                    if v8 {
                        v3 = -2.25
                        if (-9223372036854775808 == -1) {
                            var v18 uint8 = 0
                            _ = v18
                        } else {
                            fmt.Println(v2, fi17)
                        }
                    }
                    _ = fi17
                }
            } else {
                v7 = append(v7, (v10 / ((v10) | 1)))
            }
        default:
            var v19 uint32 = 2147483647
            _ = v19
        }
        _ = fi12
        _ = v13
    }
    var v20 uint16 = 916
    _ = len(v7)
    v11 += (v1 >> ((18446744073709551615) & 63))
    v7 = append(v7, v10)
    var v21 string = "\n\t"
    v22 := GenS2{F0: 14, F1: 615}
    if v8 {
        if v8 {
            v11 += -966
            var v23 uint16 = v20
            if (!v8) {
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 int8 = v2
                    _ = fi24
                    _ = v25
                }
            }
            _ = v23
        }
        var v26 float32 = ((1e10 + float32(float64(((1e10 / float32(float64(1.5))) / (3.14159 - (0.0 + (float32((v3 - 1.5)) + 3.14159))))))) / float32((v3 / -2.25)))
        v20 = v20
        _ = v26
    }
    if ((226 != 178) && v8) {
        v27 := GenS2{F0: 2147483647, F1: 255}
        _ = v27
    }
    v10 += (v10 | (-(v10)))
    v20 += (v20 * v20)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
