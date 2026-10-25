package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
func genFn0(p0 uint16) uint8 {
    p0 += (p0 % (((p0 & 1)) | 1))
    switch 42 {
    case -2147483648:
        if false {
            var v6 uint8 = 156
            v7 := []int{-225, 127, 32767, -9223372036854775808}
            if false {
                v6 = (v6 + (^(v6)))
            } else {
                var v8 int = 4294967295
                _ = v8
            }
            _ = v6
            _ = v7
        } else {
            p0 = 255
            p0 = (p0 * p0)
        }
    case 32767:
        fmt.Println(p0)
    case -937:
        p0 = (p0 + 255)
    default:
        var v9 int32 = -1
        _ = v9
    }
    _ = p0
    return 173
}
func genFn1() bool {
    var v10 string = "hello"
    v10 = v10
    _ = v10
    return (0 <= 32767)
}

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 2
    var v3 uint16 = 554
    var v4 uint = 69
    var v5 uint8 = (v2 | 42)
    fmt.Println(v4, v5, v5)
    v4 = v4
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := []int{-530, fi11, fi11}
        v5 -= (v5 & (v5 << ((572) & 63)))
        _ = fi11
        _ = v12
    }
    fmt.Println(v2)
    var v13 uint = v4
    if false {
        v14 := []int{-516, 4294967295, 755, 0}
        v13 -= 836
        if (0 == 344) {
            if true {
                switch "" {
                case "hello":
                    v13 = v4
                }
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                var v15 int64 = -9223372036854775808
                _ = v15
            }
        }
        _ = v14
    } else {
        switch (("" + "\n\t") + "") {
        case "a":
            v5 = (v2 & (^(v2)))
        case "foo bar":
            var v16 float64 = (v1 * float64(1e10))
            _ = v16
        default:
            v17 := []int{0, 0, 127}
            _ = v17
        }
        var v18 uint16 = (v3 + (v3 & v3))
        switch 2147483647 {
        case 128:
            var v19 uint8 = 0
            _ = v19
        default:
            v20 := GenS1{F0: v3, F1: 0, F2: v5}
            _ = v20
        }
        _ = v18
    }
    v21 := []int{1}
    fmt.Println(v4, v2)
    if len(v21) > 0 {
        v21[0] = -1
    }
    v21 = append(v21, 390)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
