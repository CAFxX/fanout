package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint; F2 bool }
type GenS5 struct { F0 int; F1 uint8 }
func genFn0(p0 int) int32 {
    v6 := GenS1{F0: 2}
    p0 -= (p0 * (p0 - 32767))
    _ = p0
    _ = v6
    return 425
}
func genFn1() uint64 {
    var v7 bool = (!false)
    if (false && ((true && v7) || (144 >= 91))) {
        var v8 uint8 = 42
        _ = v8
    }
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 32
    var v2 int8 = 100
    var v3 uint = 0
    var v4 string = "a"
    v5 := []int{-397}
    var v9 float64 = -2.25
    _ = len(v5)
    v10 := GenS2{F0: v3, F1: true}
    if false {
        switch 9223372036854775807 {
        case 42:
            v2 = (^(v2))
        case -77:
            fmt.Println(v2, v1, v1)
        case 1:
            for fi11 := 0; fi11 < 10; fi11++ {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v5 = append(v5, fi11)
                    if len(v5) > 0 {
                        v5[0] = fi11
                    }
                    _ = fi12
                }
                v13 := GenS1{F0: 121}
                _ = len(v5)
                _ = fi11
                _ = v13
            }
        }
        v4 = "hello"
    }
    v4 = "\n\t"
    var v14 int32 = 549
    if (true && true) {
        v15 := GenS3{F0: 0}
        v9 = 0.1
        _ = v15
    }
    var v16 int = 127
    if true {
        switch (v16 / ((-680) | 1)) {
        case -9223372036854775808:
            for fi17 := 0; fi17 < 10; fi17++ {
                v2 = (^(v2))
                fmt.Println(v4, v4)
                _ = fi17
            }
        case 2147483647:
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS4{F0: v1, F1: 4294967295, F2: ((!(v14 >= (v14 | v14))) && ((v9 - ((v9 + ((((((float64((1e10 - 0.0)) - v9) + (float64(-2.25) * (v9 / v9))) - 3.14159) + v9) + float64(1e10)) * 1e10)) + (v9 * -2.25))) >= ((0.0 + float64(1e10)) / 0.1)))}
                var v20 int16 = int16(v14)
                v14 += (v14 * (^(v14)))
                _ = fi18
                _ = v19
                _ = v20
            }
        case 631:
            v16 = (^(v16))
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            if false {
                v5 = append(v5, (fi21 | 32768))
            }
            _ = fi21
        }
    } else {
        v22 := GenS1{F0: 998}
        switch -476 {
        case -2147483648:
            switch (v16 * v16) {
            case 623:
                v5 = append(v5, (v16 + (v16 - (v16 ^ v16))))
            case -9223372036854775808:
                switch (v4 + v4) {
                case "\n\t":
                    if ((v16 | (v16 & int(v1))) == v16) {
                        v3 += (v3 >> (((v3 << (((v3 | v3)) & 63))) & 63))
                    }
                }
            default:
                v23 := GenS5{F0: (v16 - 4294967295), F1: 1}
                _ = v23
            }
        case -1:
            if (9223372036854775807 == 1) {
                fmt.Println(v14)
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 int8 = (v2 % (((v2 >> ((v3) & 63))) | 1))
                    var v26 uint = (v3 | (^(v3)))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                v4 = ((v4 + v4) + v4)
            }
        }
        _ = v22
    }
    if true {
        var v27 float32 = float32(0.1)
        switch 127 {
        case -2147483648:
            fmt.Println(v4, v27, v2)
        default:
            if false {
                _ = len(v5)
                fmt.Println(v16, v2)
                v28 := GenS1{F0: 333}
                _ = v28
            } else {
                var v29 string = ((v4 + v4) + ("hello" + v4))
                var v30 int = v16
                _ = v29
                _ = v30
            }
        }
        _ = v27
    } else {
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
