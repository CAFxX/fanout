package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint }
type GenS2 struct { F0 int64 }
func genFn0() bool {
    var v4 uint32 = 82
    v4 = v4
    var v5 string = "a"
    var v6 int = -118
    _ = v4
    _ = v5
    _ = v6
    return (!(42 < 1))
}

func main() {
    var v1 uint16 = 2
    v2 := []int{807, 42, 920, 42}
    v3 := []int{}
    v1 = v1
    v7 := GenS1{F0: 127, F1: 374, F2: 898}
    v8 := GenS1{F0: 128, F1: 0, F2: 2}
    if (uint64(v1) < 2) {
        if false {
            var v9 float64 = 0.0
            if (1 <= 625) {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                var v10 int8 = 127
                var v11 float32 = 1.5
                _ = v10
                _ = v11
            } else {
                switch ("a" + "foo bar") {
                case "foo bar":
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v1 -= v1
                        _ = fi12
                    }
                case "\n\t":
                    fmt.Println(v1, v1)
                case "":
                    var v13 bool = (2 != 811)
                    _ = v13
                }
            }
            switch 42 {
            case 734:
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 uint16 = 255
                    _ = fi14
                    _ = v15
                }
            case -777:
                v16 := GenS1{F0: 2, F1: 4294967295, F2: 18446744073709551615}
                _ = v16
            }
            _ = v9
        }
        var v17 int64 = -211
        for fi18 := 0; fi18 < 4; fi18++ {
            fmt.Println(v1)
            if true {
                v19 := GenS2{F0: (v17 % (((v17 << ((42) & 63))) | 1))}
                for fi20 := 0; fi20 < 14; fi20++ {
                    v1 = 65535
                    var v21 float64 = 0.0
                    v17 -= v17
                    _ = fi20
                    _ = v21
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v19
            } else {
                if true {
                    var v22 bool = (4294967295 <= 42)
                    if (!v22) {
                        if (1 < 255) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v2 = append(v2, fi18)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        fmt.Println(v1, fi18)
                        v2 = append(v2, fi18)
                    }
                    for fi23 := 0; fi23 < 5; fi23++ {
                        switch (^(fi18)) {
                        case 9223372036854775807:
                            v17 += (v17 % ((599) | 1))
                        case 0:
                            if len(v2) > 0 {
                                v2[0] = fi23
                            }
                        case 127:
                            fmt.Println(fi23)
                        }
                        var v24 string = "\n\t"
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
            }
            _ = fi18
        }
        _ = v17
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 float64 = 0.0
        for fi27 := 0; fi27 < 1; fi27++ {
            if ("a" > "a") {
                switch -1 {
                case 42:
                    var v28 float64 = (v26 + (0.1 / (((((0.1 + 3.14159) / 0.1) + 0.0) - v26) / (v26 * float64(1.5)))))
                    _ = v28
                }
                v1 -= 185
                v3 = append(v3, fi25)
            }
            _ = fi27
        }
        v1 -= 1
        _ = fi25
        _ = v26
    }
    v1 += (v1 - v1)
    v29 := GenS1{F0: 202, F1: 1, F2: 825}
    if (84 < 164) {
        v1 += (v1 - 32767)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v1 = 546
        _ = fi30
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        for fi32 := 0; fi32 < 6; fi32++ {
            if true {
                var v33 uint64 = 1
                v1 = 2
                var v34 int8 = 0
                _ = v33
                _ = v34
            }
            _ = fi32
        }
        v1 -= v1
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v7
    _ = v8
}
