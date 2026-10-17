package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint64 }
func genFn0(p0 int8, p1 int16) uint {
    fmt.Println(p1, p0)
    if (-1 < (p0 >> ((4294967295) & 63))) {
        v7 := []int{-1, 9223372036854775807}
        switch 32767 {
        case 42:
            var v8 float32 = 1e10
            _ = v8
        }
        v9 := GenS1{F0: 32767, F1: 4294967295, F2: 1}
        _ = v7
        _ = v9
    }
    v10 := []int{71, 0, 932, 9223372036854775807}
    v11 := GenS1{F0: 2, F1: 65535, F2: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 uint8 = 2
    var v2 bool = (!true)
    var v3 string = "foo bar"
    var v4 bool = (((^(v1)) > (^(v1))) || (-648 < -1))
    v5 := []string{}
    var v6 uint8 = (v1 + (v1 << ((232) & 63)))
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 float64 = -2.25
        _ = fi12
        _ = v13
    }
    _ = len(v5)
    fmt.Println(v1, v1, v1)
    if true {
        fmt.Println(v3)
        fmt.Println(v4, v1)
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 bool = true
            switch (fi14 % ((fi14) | 1)) {
            case 1:
                v2 = v4
            case 127:
                v15 = v15
            default:
                switch 32768 {
                case 626:
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 string = v3
                        if v4 {
                            v3 = v3
                            switch "" {
                            case "hello":
                                v2 = v2
                            case "a":
                                fmt.Println(v1)
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                case -2147483648:
                    var v18 uint32 = 0
                    _ = v18
                }
            }
            v1 += v1
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v2, v2)
    v6 += v6
    var v19 uint16 = 351
    v20 := GenS1{F0: (v19 & (v19 << ((genFn0(-128, int16(v19))) & 63))), F1: 1, F2: 2}
    v5 = append(v5, v3)
    if v2 {
        v3 = v3
    } else {
        var v21 int16 = 1
        var v22 float32 = 0.0
        v21 += v21
        _ = v21
        _ = v22
    }
    v19 = 32767
    _ = len(v5)
    v4 = (255 <= (v19 & v19))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
