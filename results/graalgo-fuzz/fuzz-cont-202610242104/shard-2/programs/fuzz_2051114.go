package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint8 = 67
    var v2 float32 = 0.1
    var v3 int8 = 100
    var v4 int16 = -32768
    var v5 bool = (v1 >= 255)
    fmt.Println(v2, v4)
    v6 := GenS1{F0: 0, F1: (v1 - (v1 ^ v1)), F2: v1}
    v2 -= v2
    v1 += 0
    var v7 bool = (!v5)
    if true {
        v8 := GenS2{F0: ((!(!(359 != 18446744073709551615))) || false), F1: 4294967295}
        var v9 int8 = v3
        if true {
            var v10 uint32 = 4294967295
            if (true && v5) {
                v11 := GenS2{F0: v5, F1: (v10 << ((42) & 63))}
                fmt.Println(v3)
                switch ("foo bar" + "foo bar") {
                case "a":
                    fmt.Println(v3, v9)
                case "":
                    for fi12 := 0; fi12 < 18; fi12++ {
                        if v7 {
                            v1 = 2
                        }
                        v1 -= 0
                        v1 = v1
                        _ = fi12
                    }
                }
                _ = v11
            } else {
                v13 := GenS2{F0: (2 == 0), F1: v10}
                v2 = float32(1.5)
                _ = v13
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                if v7 {
                    if false {
                        var v15 uint32 = 4294967295
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v5 = v5
                        _ = fi16
                    }
                }
                _ = fi14
            }
            _ = v10
        }
        _ = v8
        _ = v9
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint8 = v1
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int8 = v3
            var v21 uint = 331
            var v22 float64 = ((float64(v2) / 1.0) - ((1.5 + (float64((v2 / v2)) + (-2.25 / float64((v2 + v2))))) / 3.14159))
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        v7 = v5
        _ = fi17
        _ = v18
    }
    if (((v7 && (!((2 <= v1) || (false || (367 < 1))))) && ((-271 > 128) && false)) || true) {
        fmt.Println(v5, v7)
    }
    var v23 int8 = (v3 / (((v3 % (((v3 >> ((18446744073709551615) & 63))) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
