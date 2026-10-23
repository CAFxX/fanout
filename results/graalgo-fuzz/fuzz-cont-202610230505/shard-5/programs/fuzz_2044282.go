package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 uint }

func main() {
    var v1 uint8 = 41
    var v2 int = 1
    var v3 uint64 = 0
    var v4 float64 = 1.5
    var v5 uint16 = 255
    var v6 string = ("a" + (("\n\t" + ("\n\t" + "foo bar")) + "\n\t"))
    if true {
        v7 := GenS1{F0: (v3 - (v3 + v3))}
        v4 = 3.14159
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{127}
        v2 = v2
        v10 := GenS2{F0: v5}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS1{F0: (v3 << ((18446744073709551615) & 63))}
        fmt.Println(v5)
        v13 := []int{}
        _ = fi11
        _ = v12
        _ = v13
    }
    v2 = (v2 << ((267) & 63))
    var v14 uint8 = v1
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := []int{v2, (^(v2)), (-(v2))}
        switch v6 {
        case "":
            v5 -= v5
        default:
            if len(v16) > 0 {
                v16[0] = fi15
            }
        }
        _ = fi15
        _ = v16
    }
    var v17 uint8 = 38
    v18 := []int{32767, v2, v2, v2}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v19 uint = 255
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int = v2
        if (true && ((fi20 * -1) <= fi20)) {
            if len(v18) > 0 {
                v18[0] = (v2 >> ((v19) & 63))
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                v23 := GenS3{F0: (v2 / (((-(fi20))) | 1)), F1: 1}
                var v24 int = v2
                _ = fi22
                _ = v23
                _ = v24
            }
            v14 = (v1 | v17)
        } else {
            v17 = v14
            if true {
                v25 := GenS4{F0: v19}
                _ = v25
            } else {
                v3 += (v3 * v3)
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                if true {
                    fmt.Println(fi26, v14)
                }
                _ = fi26
            }
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v3, v17, v17)
    fmt.Println(v6, v6)
    v6 = "hello"
    if (false && false) {
        v4 = float64(0.0)
        v18 = append(v18, v2)
        if (255 >= 255) {
            v27 := GenS5{F0: (float64(1e10) <= v4), F1: uint(v17)}
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
