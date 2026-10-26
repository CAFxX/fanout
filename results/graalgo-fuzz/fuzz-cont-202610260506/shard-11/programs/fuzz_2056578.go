package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int8; F2 int8 }

func main() {
    v1 := []int{964, 32768}
    v2 := []string{"hello"}
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 1
    fmt.Println(v5)
    if ((v5 >> ((2) & 63)) == (v5 / (((^(v5))) | 1))) {
        if (0.0 == ((3.14159 * ((1.5 / ((3.14159 + 3.14159) + 1.5)) + float64((0.0 * 0.1)))) - float64(1.5))) {
            fmt.Println(v5)
            var v6 uint64 = 1
            _ = v6
        }
    }
    fmt.Println(v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 int = 128
    v8 := GenS1{F0: 133, F1: (false || ((689 != 42) || true)), F2: 807}
    fmt.Println(v7, v7)
    var v9 int32 = -30
    for fi10 := 0; fi10 < 10; fi10++ {
        if len(v1) > 0 {
            v1[0] = int(v9)
        }
        v9 += -222
        _ = fi10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            if false {
                var v13 int = v7
                _ = v13
            } else {
                v14 := GenS1{F0: 0, F1: (0 >= 42), F2: 4294967295}
                _ = v14
            }
            if false {
                v5 = (v5 & (v5 % ((897) | 1)))
                delete(v4, 0)
            } else {
                v15 := GenS2{F0: (^(v5)), F1: (v5 >> ((1) & 63))}
                fmt.Println(fi11, v7)
                _ = v15
            }
            fmt.Println(v7)
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println(v7)
    var v16 uint8 = 1
    switch v7 {
    case -2147483648:
        v16 = (v16 | 127)
    case 127:
        if (1e10 <= (1.5 * (float32(3.14159) - 1.5))) {
            for fi17 := 0; fi17 < 18; fi17++ {
                if true {
                    var v18 uint32 = 703
                    _ = v18
                } else {
                    if (100 >= 38) {
                        v19 := v4[930]
                        _ = v19
                    }
                }
                v20 := GenS1{F0: v16, F1: false, F2: 2}
                _ = fi17
                _ = v20
            }
        }
    case 2147483647:
        var v21 string = "hello"
        _ = v21
    default:
        switch v7 {
        case 572:
            v3[1] = 276
        case -1:
            v22 := GenS2{F0: 42, F1: (v5 / ((401) | 1))}
            _ = v22
        case 127:
            v16 = (^(v16))
        default:
            if (965 == 9223372036854775807) {
                v23 := GenS3{F0: 2, F1: 0, F2: -1}
                _ = v23
            }
        }
    }
    var v24 string = ""
    v2 = append(v2, v24)
    if (!(100 >= 127)) {
        fmt.Println(v9, v5)
    }
    delete(v3, 4294967295)
    v5 += (^(v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
