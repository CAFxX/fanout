package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }

func main() {
    var v1 float64 = (0.0 - 3.14159)
    var v2 float32 = -2.25
    var v3 bool = false
    var v4 uint32 = 65535
    var v5 uint16 = 842
    var v6 bool = ((0 >= 494) && (2 >= 4294967295))
    v1 -= v1
    fmt.Println(v4)
    v7 := GenS1{F0: -1}
    var v8 int = -2147483648
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := []int{(v8 / ((fi9) | 1)), 0}
        v11 := GenS1{F0: -1}
        _ = fi9
        _ = v10
        _ = v11
    }
    switch 0 {
    case -678:
        v2 = 1.5
    default:
        switch v8 {
        case 32767:
            var v12 uint16 = (v5 << ((18446744073709551615) & 63))
            _ = v12
        case 42:
            v13 := GenS1{F0: -2147483648}
            _ = v13
        }
    }
    if v6 {
        v3 = true
        v14 := []int{}
        if ((float64((((3.14159 + 0.1) * 1.5) - 1e10)) <= 0.0) && false) {
            if false {
                v6 = true
                var v15 uint8 = 0
                switch 42 {
                case -1:
                    var v16 bool = v3
                    _ = v16
                case 32768:
                    v14 = append(v14, 814)
                default:
                    fmt.Println(v15)
                }
                _ = v15
            }
        } else {
            v17 := GenS1{F0: 2147483647}
            _ = v17
        }
        _ = v14
    } else {
        v3 = false
    }
    v18 := []int{42, (v8 >> ((18446744073709551615) & 63))}
    v2 -= 1.5
    v19 := GenS2{F0: (v8 + v8)}
    var v20 int8 = 1
    v21 := GenS1{F0: -806}
    switch "foo bar" {
    case "":
        _ = len(v18)
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v2 = v2
        v23 := GenS2{F0: (v8 + (fi22 % (((^(fi22))) | 1)))}
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
