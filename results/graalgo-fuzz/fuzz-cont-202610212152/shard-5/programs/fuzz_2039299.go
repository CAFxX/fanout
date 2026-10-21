package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 bool, p1 int, p2 string) uint64 {
    v6 := []int{p1, p1, (-(p1)), p1}
    var v7 uint = 940
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 741
}

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 2
    var v3 uint16 = v2
    var v4 uint8 = 1
    v5 := make(map[uint32]string)
    var v8 uint8 = (^(v4))
    v9 := GenS1{F0: -152, F1: v4}
    v10 := GenS1{F0: 42, F1: 235}
    fmt.Println(v1)
    v11 := []int{32768, 0, 6, -1}
    if (v3 > 32767) {
        v12 := GenS1{F0: 1, F1: (v8 ^ v4)}
        v1 -= 3.14159
        _ = v12
    }
    if (!(9223372036854775807 > -970)) {
        v2 = v2
    }
    switch 437 {
    case -491:
        v8 = 255
    case -9223372036854775808:
        fmt.Println(v8, v2, v8)
    }
    switch ("\n\t" + "\n\t") {
    case "":
        var v13 uint64 = 0
        _ = v13
    case "\n\t":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    v1 += v1
    for fi14 := 0; fi14 < 19; fi14++ {
        if false {
            if (42 <= 42) {
                for fi15 := 0; fi15 < 7; fi15++ {
                    v2 = 0
                    v16 := GenS1{F0: -583, F1: uint8(v3)}
                    _ = fi15
                    _ = v16
                }
                v11 = append(v11, (^(fi14)))
            }
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 int32 = 0
                _ = fi17
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 string = "a"
                var v21 bool = false
                v4 += (v8 >> ((255) & 63))
                _ = fi19
                _ = v20
                _ = v21
            }
            switch fi14 {
            case 0:
                fmt.Println(v1, fi14, v1)
            default:
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 int8 = 1
                    if false {
                        v8 -= (v8 ^ v8)
                    }
                    _ = fi22
                    _ = v23
                }
            }
        }
        var v24 uint = 255
        v25 := GenS1{F0: -2147483648, F1: (^(v4))}
        _ = fi14
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
