package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0(p0 uint32, p1 int8) uint8 {
    fmt.Println(p0)
    var v7 uint = 2
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]string)
    v3 := []int{-260}
    var v4 uint16 = 42
    v5 := make(map[uint32]int)
    var v6 uint32 = 2147483648
    var v8 string = "hello"
    v4 = v4
    v5[v6] = 485
    v9 := GenS1{F0: v4}
    for fi10 := 0; fi10 < 16; fi10++ {
        if (!(((float64((1.5 / 3.14159)) - 3.14159) < 0.1) && ((v1 % ((v1) | 1)) >= (v1 % (((v1 & (v1 + 65535))) | 1))))) {
            if true {
                for fi11 := 0; fi11 < 16; fi11++ {
                    v4 = (^(v4))
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    v4 += v4
                    _ = fi11
                }
                v12 := GenS1{F0: (v4 - v4)}
                v6 -= v6
                _ = v12
            }
            switch ("" + v8) {
            case "a":
                v13 := v5[(v6 / ((0) | 1))]
                _ = v13
            case "hello":
                switch "a" {
                case "":
                    v3 = append(v3, fi10)
                }
            default:
                v3 = append(v3, fi10)
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v14 uint64 = 0
        _ = fi10
        _ = v14
    }
    delete(v2, -2147483648)
    fmt.Println(v4)
    v15 := GenS1{F0: (v4 + 311)}
    v16 := GenS2{F0: (!false), F1: -1}
    v3 = append(v3, 4294967295)
    v17 := GenS2{F0: true, F1: -9223372036854775808}
    v4 += v4
    if false {
        for fi18 := 0; fi18 < 3; fi18++ {
            v1 = v1
            _ = fi18
        }
        v8 = v8
    } else {
        v1 -= v1
        var v19 uint16 = 0
        _ = v19
    }
    _ = len(v2)
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 14; fi21++ {
            if (!(fi21 >= (fi21 + (fi21 / ((1) | 1))))) {
                _ = len(v3)
            }
            v6 += v6
            _ = fi21
        }
        v1 = v1
        _ = fi20
    }
    v2[-9223372036854775808] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
