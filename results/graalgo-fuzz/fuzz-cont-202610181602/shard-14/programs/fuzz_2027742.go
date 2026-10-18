package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := []int{-183}
    var v2 int32 = -2147483648
    var v3 uint16 = 32767
    if true {
        var v4 uint32 = 65535
        var v5 int64 = 0
        if (2 <= 177) {
            if len(v1) > 0 {
                v1[0] = 128
            }
            var v6 uint = 890
            v7 := GenS1{F0: -92, F1: 0}
            _ = v6
            _ = v7
        }
        _ = v4
        _ = v5
    }
    v2 -= (^(v2))
    var v8 int64 = -1
    switch 42 {
    case 435:
        fmt.Println(v2, v2)
    case -277:
        var v9 int32 = -522
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v3, v3)
        fmt.Println(v3, v2)
        _ = fi10
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 uint = 442
    if true {
        if false {
            if len(v1) > 0 {
                v1[0] = 128
            }
            v1 = append(v1, 127)
        } else {
            if true {
                for fi12 := 0; fi12 < 20; fi12++ {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 uint32 = 65535
                        switch "\n\t" {
                        case "a":
                            v3 = uint16(v11)
                        default:
                            v11 = v11
                        }
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
            } else {
                switch -1 {
                case -539:
                    var v15 uint16 = v3
                    _ = v15
                case 0:
                    v11 -= v11
                }
            }
        }
        v16 := GenS2{F0: (-(v2)), F1: 0, F2: false}
        _ = v16
    } else {
        v17 := GenS2{F0: int32(v11), F1: 57, F2: ((!false) || false)}
        v18 := GenS3{F0: v2, F1: 721}
        _ = v17
        _ = v18
    }
    v19 := GenS4{F0: 318}
    v2 += -2147483648
    for fi20 := 0; fi20 < 16; fi20++ {
        v1 = append(v1, 128)
        fmt.Println(fi20, v11, v8)
        v3 = 657
        _ = fi20
    }
    v8 -= v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
