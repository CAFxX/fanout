package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 bool = false
    var v2 int8 = 100
    var v3 int32 = 2147483647
    v4 := make(map[int]uint32)
    var v5 bool = v1
    v2 -= v2
    v6 := GenS1{F0: "hello", F1: 32767, F2: 916}
    v7 := GenS2{F0: 42}
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := []int{-9223372036854775808}
        _ = fi8
        _ = v9
    }
    v10 := v4[-362]
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            v2 += v2
            _ = fi12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{(fi13 >> ((42) & 63)), (fi13 % (((fi13 - 42)) | 1)), (fi13 - fi13)}
        _ = fi13
        _ = v14
    }
    v1 = (v5 || (!true))
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := []int{32768, (fi15 + fi15)}
        for fi17 := 0; fi17 < 18; fi17++ {
            if v1 {
                v3 = int32(fi15)
                var v18 int8 = (v2 % ((v2) | 1))
                if ((^(v10)) == v10) {
                    v19 := GenS2{F0: (v10 >> ((42) & 63))}
                    _ = len(v16)
                    v18 -= 100
                    var v20 int64 = 0
                    _ = v19
                    _ = v20
                }
                _ = v18
            } else {
                fmt.Println(v2)
                switch fi15 {
                case -718:
                    if len(v16) > 0 {
                        v16[0] = (fi15 + -254)
                    }
                case 9223372036854775807:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                case 251:
                    var v21 uint64 = 42
                    _ = v21
                }
                var v22 int64 = 0
                v3 += (v3 + 0)
                var v23 float64 = -2.25
                _ = v22
                _ = v23
            }
            switch fi15 {
            case -9223372036854775808:
                v1 = v1
            case 1:
                var v24 int32 = 2147483647
                _ = v24
            case 0:
                v10 = (v10 & (^(v10)))
            default:
                var v25 int16 = 32767
                _ = v25
            }
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
