package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    var v1 int16 = -632
    var v2 int64 = 771
    var v3 int8 = -86
    var v4 int16 = (v1 - v1)
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: 127, F1: 591}
    v4 = 1000
    if ((("\n\t" + ((("hello" + "hello") + ("" + "\n\t")) + "hello")) + ("\n\t" + "foo bar")) < "") {
        v7 := []int{1, 1}
        v3 -= (v3 + v3)
        if (1 != 949) {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int16 = -32768
                switch fi8 {
                case -678:
                    v4 = v1
                }
                _ = fi8
                _ = v9
            }
        } else {
            v7 = append(v7, 254)
        }
        _ = v7
    } else {
        v10 := GenS2{F0: 2, F1: (v4 + v1)}
        _ = v10
    }
    fmt.Println(v4)
    v11 := []int{0, 32767, 404}
    if (-32768 >= (^(v4))) {
        _ = len(v5)
    }
    var v12 int32 = int32(v4)
    v2 += -9223372036854775808
    v3 = (v3 * 0)
    if ((3.14159 == 0.0) && ("hello" >= "hello")) {
        _ = len(v5)
        switch "a" {
        case "":
            var v13 bool = false
            _ = v13
        }
    }
    var v14 uint8 = 255
    switch 32767 {
    case 592:
        for fi15 := 0; fi15 < 10; fi15++ {
            switch "" {
            case "a":
                v16 := GenS3{F0: 0, F1: 1, F2: 820}
                _ = v16
            default:
                v12 += v12
            }
            _ = fi15
        }
    case 9223372036854775807:
        v11 = append(v11, 127)
    default:
        v4 = (v1 & v1)
    }
    var v17 int8 = -6
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
