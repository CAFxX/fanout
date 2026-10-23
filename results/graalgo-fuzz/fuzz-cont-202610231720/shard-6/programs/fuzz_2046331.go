package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 bool; F2 int32 }
type GenS3 struct { F0 uint16; F1 int16; F2 int32 }

func main() {
    var v1 float32 = 0.1
    var v2 bool = (!((v1 * ((float32(0.0) + v1) - float32((1e10 + float64(1e10))))) >= v1))
    var v3 int64 = -9223372036854775808
    fmt.Println(v1, v3)
    var v4 int16 = 844
    var v5 float32 = ((float32((1.5 / (0.0 - float64(3.14159)))) + 0.0) * v1)
    if (("hello" + "\n\t") > "hello") {
        v6 := []int{}
        v3 -= (^(v3))
        switch 1 {
        case 9223372036854775807:
            var v7 uint16 = 1
            _ = v7
        case -1:
            var v8 string = ""
            _ = v8
        case 128:
            for fi9 := 0; fi9 < 1; fi9++ {
                if v2 {
                    fmt.Println(fi9)
                    v10 := GenS1{F0: 18446744073709551615}
                    _ = v10
                } else {
                    var v11 uint = 42
                    var v12 int32 = -189
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
        }
        _ = v6
    } else {
        var v13 uint32 = 1
        if false {
            if v2 {
                if (!v2) {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        if (18446744073709551615 >= 46) {
                            fmt.Println(v4, v1)
                            fmt.Println(v2, v2)
                        }
                        v15 := []int{(-(fi14)), (-(fi14)), (fi14 - fi14)}
                        var v16 int16 = -497
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                } else {
                    v17 := []int{9223372036854775807, 552}
                    v18 := GenS2{F0: -305, F1: (0 > 2), F2: 42}
                    _ = v17
                    _ = v18
                }
                var v19 int32 = -1
                fmt.Println(v13, v2)
                _ = v19
            }
        }
        _ = v13
    }
    var v20 uint8 = 159
    var v21 uint8 = (v20 | (v20 / ((207) | 1)))
    v22 := GenS3{F0: 0, F1: v4, F2: -2147483648}
    fmt.Println(v21)
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS2{F0: (-(fi23)), F1: (!false), F2: -891}
        v5 = (v1 * (3.14159 * ((v1 - v1) * (3.14159 - v5))))
        _ = fi23
        _ = v24
    }
    var v25 float64 = -2.25
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 uint8 = 127
            v29 := GenS1{F0: 658}
            var v30 string = "hello"
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        v2 = v2
        _ = fi26
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        fmt.Println(v5, v4)
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
