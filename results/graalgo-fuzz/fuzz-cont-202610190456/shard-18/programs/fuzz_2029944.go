package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint16 = 65535
    var v2 int32 = 0
    var v3 float64 = (0.0 - (1.5 * (1e10 + 1.5)))
    var v4 uint32 = 65535
    v3 += v3
    if true {
        for fi5 := 0; fi5 < 16; fi5++ {
            fmt.Println(fi5, v3, v3)
            if true {
                fmt.Println(v4, v4)
            } else {
                for fi6 := 0; fi6 < 12; fi6++ {
                    var v7 int16 = 224
                    if (("" + "foo bar") <= "\n\t") {
                        if (!true) {
                            v2 = v2
                            if (false || (true || (!false))) {
                                v8 := GenS1{F0: 608}
                                _ = v8
                            } else {
                                v9 := GenS1{F0: (v4 * v4)}
                                _ = v9
                            }
                        } else {
                            v10 := []int{fi6, (fi6 | fi5), (fi5 / (((fi6 << ((42) & 63))) | 1))}
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v10
                        }
                        if true {
                            var v11 uint64 = 1
                            v3 = 1e10
                            var v12 int32 = 2147483647
                            _ = v11
                            _ = v12
                        } else {
                            v13 := GenS1{F0: (v4 ^ v4)}
                            _ = v13
                        }
                    } else {
                        v3 -= 0.1
                    }
                    v14 := []int{fi6, -43}
                    _ = fi6
                    _ = v14
                    _ = v7
                }
                v15 := GenS1{F0: 1}
                switch fi5 {
                case -1:
                    v16 := GenS1{F0: v4}
                    _ = v16
                }
                _ = v15
            }
            _ = fi5
        }
    }
    v17 := []int{636, -9223372036854775808}
    for fi18 := 0; fi18 < 3; fi18++ {
        v3 = (float64(((0.0 + (1e10 / 3.14159)) * 0.1)) + float64((1e10 * (3.14159 * ((3.14159 - 3.14159) + 1.5)))))
        _ = fi18
    }
    if ((!false) || false) {
        v3 -= (float64(1e10) + float64((1e10 + 0.0)))
    }
    v19 := GenS1{F0: 2147483647}
    if true {
        fmt.Println(v4, v1)
        fmt.Println(v1, v2, v3)
        v4 -= v4
    } else {
        if (false && true) {
            switch 1 {
            case 127:
                if len(v17) > 0 {
                    v17[0] = 32768
                }
            case -866:
                if true {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v3 -= 1.5
                        fmt.Println(v2, v2)
                        v2 = (-(v2))
                        _ = fi20
                    }
                    fmt.Println(v4)
                }
            }
            fmt.Println(v2)
            fmt.Println(v2)
        }
        v2 = v2
        for fi21 := 0; fi21 < 16; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                _ = len(v17)
                _ = fi22
            }
            switch fi21 {
            case -1:
                if len(v17) > 0 {
                    v17[0] = 9223372036854775807
                }
            }
            _ = fi21
        }
    }
    var v23 int16 = 32767
    v24 := GenS2{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
