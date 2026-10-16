package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 bool }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint8 = 1
    var v2 string = "a"
    var v3 uint64 = 522
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 uint32 = 2
            v7 := []int{fi5, -147, -2147483648}
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 int32 = -138
        _ = fi4
        _ = v8
    }
    fmt.Println(v3, v1, v3)
    if false {
        v3 += 42
    } else {
        var v9 int = 447
        var v10 bool = false
        _ = v10
        _ = v9
    }
    if false {
        v11 := GenS1{F0: 257}
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if (true || (18446744073709551615 >= 276)) {
            fmt.Println(v2)
        } else {
            v13 := []int{42, fi12, 32768, fi12}
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(v1)
                switch (fi12 * 4294967295) {
                case 32768:
                    v15 := GenS2{F0: v3, F1: v2, F2: (v1 <= (v1 - v1))}
                    v1 -= v1
                    if (-2.25 == 3.14159) {
                        if ((0 <= 623) || (!false)) {
                            var v16 int16 = 728
                            _ = v16
                        }
                        if false {
                            v3 = v3
                            var v17 float64 = 3.14159
                            var v18 int32 = -827
                            _ = v17
                            _ = v18
                        }
                        switch "" {
                        case "foo bar":
                            v19 := GenS1{F0: 42}
                            var v20 bool = false
                            _ = v19
                            _ = v20
                        }
                        fmt.Println(v3, v3)
                    } else {
                        v21 := GenS2{F0: (v3 * (^(v3))), F1: ((v2 + ("\n\t" + v2)) + v2), F2: true}
                        _ = v21
                    }
                    _ = v15
                case -1:
                    var v22 float64 = (3.14159 - (1.5 * (1.5 / float64(((0.0 / 1.0) + 1.5)))))
                    _ = v22
                }
                _ = fi14
            }
            var v23 uint32 = 2147483648
            _ = v13
            _ = v23
        }
        fmt.Println(v1, v1, v3)
        _ = fi12
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 float64 = 1.5
        switch (("" + (v2 + v2)) + (((v2 + "") + "hello") + ((("a" + v2) + (v2 + ((v2 + "foo bar") + v2))) + v2))) {
        case "":
            var v26 uint32 = 314
            _ = v26
        case "\n\t":
            if false {
                switch -9223372036854775808 {
                case 32768:
                    var v27 int64 = 217
                    _ = v27
                case -2147483648:
                    v28 := GenS1{F0: 2}
                    _ = v28
                default:
                    v29 := GenS2{F0: (v3 << ((1) & 63)), F1: v2, F2: true}
                    _ = v29
                }
                v25 = float64((((0.0 / (3.14159 + ((-2.25 * (1e10 * 1e10)) / 1e10))) / -2.25) - 1.5))
                if (-459 >= -1) {
                    v30 := []int{fi24, fi24}
                    _ = v30
                }
            } else {
                var v31 int = (fi24 - fi24)
                v32 := []int{fi24}
                _ = v31
                _ = v32
            }
        }
        v33 := GenS3{F0: true}
        _ = fi24
        _ = v25
        _ = v33
    }
    if (1 == 4294967295) {
        var v34 float32 = 3.14159
        _ = v34
    }
    var v35 uint64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}
