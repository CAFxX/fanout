package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint, p1 int, p2 int) int {
    p2 += p1
    if false {
        v4 := GenS1{F0: -9223372036854775808}
        if false {
            p1 -= 1
            v5 := GenS1{F0: -1}
            switch 42 {
            case 2147483647:
                v6 := GenS2{F0: 1, F1: 127}
                _ = v6
            }
            _ = v5
        }
        _ = v4
    }
    p2 = (p2 / ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    return (p1 / ((p1) | 1))
}

func main() {
    var v1 int64 = -761
    v2 := []uint32{65535, 0}
    v3 := []int{-1}
    if len(v3) > 0 {
        v3[0] = 32768
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS2{F0: 2, F1: 1}
        if (!true) {
            fmt.Println(v1)
        } else {
            v3 = append(v3, (fi7 * genFn0(4294967295, (fi7 * genFn0(42, genFn0(598, -551, (-(fi7))), (^(fi7)))), genFn0(255, 2147483647, fi7))))
            fmt.Println(fi7, fi7)
            v9 := GenS1{F0: (v1 - 9223372036854775807)}
            _ = v9
        }
        var v10 int32 = 452
        _ = fi7
        _ = v10
        _ = v8
    }
    var v11 int = -9223372036854775808
    var v12 uint16 = 190
    if true {
        switch "\n\t" {
        case "a":
            _ = len(v2)
        }
    } else {
        v11 += -2147483648
        var v13 int64 = (-(v1))
        if false {
            v14 := GenS1{F0: (v13 * (v13 * v13))}
            for fi15 := 0; fi15 < 20; fi15++ {
                if len(v3) > 0 {
                    v3[0] = v11
                }
                v12 += (v12 & (v12 << ((42) & 63)))
                _ = fi15
            }
            v2 = append(v2, 251)
            _ = v14
        }
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = 443
    }
    fmt.Println(v11, v11)
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v11 = (^(fi16))
            if false {
                _ = len(v2)
            } else {
                var v18 uint = 0
                var v19 int16 = -32768
                var v20 string = "\n\t"
                _ = v18
                _ = v19
                _ = v20
            }
            var v21 int = (v11 / ((980) | 1))
            v12 = v12
            var v22 int16 = -1
            _ = fi17
            _ = v21
            _ = v22
        }
        _ = fi16
    }
    v11 = -120
    v12 = (v12 << ((729) & 63))
    var v23 int64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
}
