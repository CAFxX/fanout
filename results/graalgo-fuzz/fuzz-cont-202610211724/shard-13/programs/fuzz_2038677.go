package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 bool = true
    var v2 int16 = -314
    var v3 int16 = (v2 % ((-1) | 1))
    v4 := []uint32{2, 2, 65535, 1}
    v5 := make(map[string]bool)
    switch "hello" {
    case "":
        for fi6 := 0; fi6 < 13; fi6++ {
            if len(v4) > 0 {
                v4[0] = 42
            }
            _ = fi6
        }
    default:
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    var v8 float64 = 0.0
    v3 = (v2 << ((255) & 63))
    fmt.Println(v8, v2)
    v9 := GenS1{F0: 18446744073709551615, F1: -28, F2: int(v2)}
    var v10 int8 = 100
    v11 := GenS1{F0: 1, F1: (v10 % ((v10) | 1)), F2: 128}
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS2{F0: 126, F1: 42}
        v14 := GenS3{F0: 1}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float64 = (1e10 + (v8 + 0.0))
    for fi16 := 0; fi16 < 19; fi16++ {
        if (43 == 990) {
            for fi17 := 0; fi17 < 4; fi17++ {
                fmt.Println(v1, fi16)
                if v1 {
                    fmt.Println(v1, v15)
                }
                switch fi16 {
                case -1:
                    if v1 {
                        v18 := GenS2{F0: 255, F1: 42}
                        _ = v18
                    } else {
                        v5["a"] = (76 >= 128)
                        v19 := GenS2{F0: 1, F1: -9223372036854775808}
                        _ = v19
                    }
                default:
                    if ("a" != "hello") {
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                    } else {
                        var v20 uint = 42
                        _ = v20
                    }
                }
                _ = fi17
            }
        } else {
            fmt.Println(fi16, fi16)
            switch (fi16 / ((fi16) | 1)) {
            case -501:
                switch int(v10) {
                case 4294967295:
                    v21 := GenS4{F0: (fi16 - -916), F1: -907}
                    _ = v21
                case 2147483647:
                    v22 := GenS3{F0: 2}
                    _ = v22
                }
            case 32768:
                fmt.Println(v8, v8, v1)
            }
            if (65535 <= 205) {
                switch fi16 {
                case 628:
                    v2 -= 416
                }
            } else {
                var v23 string = ""
                var v24 float32 = (1e10 * 1e10)
                var v25 int = (-(fi16))
                _ = v23
                _ = v24
                _ = v25
            }
        }
        v26 := GenS5{F0: 2}
        fmt.Println(v1)
        _ = fi16
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
