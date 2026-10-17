package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    var v1 bool = ((469 < 81) && true)
    v2 := []string{"hello"}
    v3 := make(map[int]bool)
    v1 = false
    v1 = (621 <= 0)
    v1 = v1
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 bool = v1
        v6 := GenS1{F0: 42, F1: 42, F2: (true || (32767 <= 510))}
        if v1 {
            v7 := GenS1{F0: 18446744073709551615, F1: 0, F2: true}
            _ = v7
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    v1 = (0 < 32767)
    v8 := GenS1{F0: 42, F1: -22, F2: (42 != 2147483648)}
    var v9 uint = 0
    v10 := v3[32768]
    var v11 int64 = 42
    if true {
        v12 := GenS2{F0: true, F1: v9, F2: -128}
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        if true {
            delete(v3, 42)
            switch 139 {
            case -488:
                v11 = (v11 & (v11 << ((v9) & 63)))
            case 2147483647:
                v10 = v1
            case -2147483648:
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 string = ("a" + "foo bar")
                    _ = fi14
                    _ = v15
                }
            default:
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := GenS3{F0: 22, F1: ((1.5 / 1.5) == (-2.25 - float32((-2.25 * ((float64(0.1) / -2.25) - 1e10)))))}
                    var v18 int64 = v11
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        } else {
            if len(v2) > 0 {
                v2[0] = ""
            }
        }
        var v19 int8 = -128
        _ = fi13
        _ = v19
    }
    _ = len(v2)
    fmt.Println(v1, v1, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
