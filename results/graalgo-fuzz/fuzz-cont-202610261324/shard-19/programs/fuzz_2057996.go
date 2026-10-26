package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int32 = 0
    var v2 int8 = 100
    v3 := []int{32767, 9223372036854775807, -717}
    var v4 uint16 = 2
    v5 := GenS1{F0: 128, F1: v2}
    v3 = append(v3, 4294967295)
    var v6 uint64 = 566
    var v7 bool = ((v4 >> ((42) & 63)) > (v4 >> ((802) & 63)))
    if (((v7 || (!v7)) || v7) || (!((v7 || ((255 == 730) || v7)) || (uint8(v4) <= 128)))) {
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS2{F0: 42}
                v4 = v4
                _ = fi9
                _ = v10
            }
            if (!v7) {
                v4 = v4
                var v11 int64 = 1
                _ = v11
            } else {
                if (!v7) {
                    fmt.Println(v1, v4, v6)
                    v12 := GenS1{F0: 1, F1: v2}
                    var v13 uint = 442
                    _ = v12
                    _ = v13
                } else {
                    switch fi8 {
                    case 4294967295:
                        switch ("hello" + ("a" + "")) {
                        case "hello":
                            v3 = append(v3, (fi8 >> ((uint(v6)) & 63)))
                        }
                    case 337:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 0:
                        fmt.Println(v1)
                    default:
                        v14 := GenS3{F0: v6}
                        _ = v14
                    }
                }
            }
            _ = fi8
        }
    }
    v15 := GenS1{F0: 2, F1: v2}
    v2 = (v2 / ((v2) | 1))
    fmt.Println(v2)
    var v16 int16 = 1000
    if v7 {
        v1 = v1
    }
    var v17 uint = 217
    var v18 uint16 = 2
    var v19 int8 = (-(v2))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
