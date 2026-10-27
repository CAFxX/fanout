package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int64 }

func main() {
    var v1 int = 127
    var v2 float64 = float64((-2.25 + 3.14159))
    var v3 int16 = 1000
    var v4 bool = (!(!(!(!false))))
    v5 := GenS1{F0: 0, F1: 32768, F2: -120}
    v6 := GenS1{F0: -1, F1: (-(v1)), F2: 42}
    var v7 int16 = -32768
    if (!(100 == 0)) {
        v3 -= (v7 % ((1000) | 1))
    } else {
        v8 := []int{(v1 - v1)}
        switch "\n\t" {
        case "":
            var v9 int32 = -378
            _ = v9
        default:
            var v10 int8 = 42
            _ = v10
        }
        if (v1 <= (v1 % ((-9223372036854775808) | 1))) {
            if true {
                v11 := GenS1{F0: 9223372036854775807, F1: v1, F2: 329}
                _ = v11
            }
            switch (v1 * 4294967295) {
            case 127:
                v1 += v1
            case 0:
                var v12 int = v1
                _ = v12
            }
            v2 = 3.14159
        } else {
            var v13 uint8 = 0
            var v14 bool = (127 != -38)
            _ = v13
            _ = v14
        }
        _ = v8
    }
    v2 += ((0.1 * v2) * (v2 - v2))
    var v15 bool = (("hello" == ("hello" + ("a" + (("" + "hello") + "a")))) && (!true))
    if v15 {
        if false {
            var v16 int64 = -283
            if v15 {
                if v4 {
                    if v15 {
                        v17 := GenS1{F0: (v16 + 0), F1: v1, F2: -90}
                        for fi18 := 0; fi18 < 11; fi18++ {
                            fmt.Println(v1, v3, v7)
                            switch 9223372036854775807 {
                            case 532:
                                v7 = v7
                            }
                            _ = fi18
                        }
                        _ = v17
                    } else {
                        var v19 uint32 = 81
                        _ = v19
                    }
                }
            } else {
                v20 := []int{v1, v1, 128}
                _ = v20
            }
            if (true && (((float64((float32(v2) - float32(v2))) - 0.0) + v2) >= (v2 + (0.1 - v2)))) {
                v3 -= v3
            } else {
                for fi21 := 0; fi21 < 3; fi21++ {
                    v22 := GenS1{F0: -9223372036854775808, F1: fi21, F2: v16}
                    _ = fi21
                    _ = v22
                }
                v23 := GenS1{F0: v16, F1: (v1 & (v1 / (((v1 << ((255) & 63))) | 1))), F2: (v16 >> ((409) & 63))}
                v1 = (v1 % ((v1) | 1))
                _ = v23
            }
            _ = v16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
