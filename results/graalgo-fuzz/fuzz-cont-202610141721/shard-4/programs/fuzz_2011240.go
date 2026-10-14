package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 65535
    var v3 string = "a"
    var v4 int16 = 1000
    v2 = (v2 % (((^(v2))) | 1))
    switch v3 {
    case "a":
        if ((1 == 1) && false) {
            v5 := GenS1{F0: 263, F1: -300, F2: v2}
            if false {
                v6 := []int{128, 563, -342}
                if (1e10 <= (((((-2.25 * (-2.25 * ((1.5 - 1.5) - (1e10 + 1.5)))) / 1e10) + (0.0 - -2.25)) * 0.0) + -2.25)) {
                    var v7 int8 = -26
                    delete(v1, 65535)
                    _ = v7
                } else {
                    if (!(1 == -978)) {
                        for fi8 := 0; fi8 < 13; fi8++ {
                            var v9 int64 = -542
                            switch (("\n\t" + (("" + v3) + "hello")) + "") {
                            case "a":
                                v9 -= v9
                            case "\n\t":
                                var v10 uint32 = 0
                                _ = v10
                            default:
                                v11 := GenS1{F0: v4, F1: 2147483647, F2: (v2 * (v2 % (((^(v2))) | 1)))}
                                _ = v11
                            }
                            _ = fi8
                            _ = v9
                        }
                        v12 := v1[4294967295]
                        _ = v12
                    }
                    if true {
                        if (!(1 < -9223372036854775808)) {
                            v4 -= -32768
                        }
                        for fi13 := 0; fi13 < 1; fi13++ {
                            var v14 int = -1
                            _ = fi13
                            _ = v14
                        }
                        var v15 uint8 = 1
                        _ = v15
                    }
                }
                _ = v6
            }
            v16 := GenS1{F0: (^(v4)), F1: -2147483648, F2: (v2 >> ((427) & 63))}
            _ = v16
            _ = v5
        }
    default:
        v4 = 32767
    }
    var v17 int = -1
    v18 := GenS1{F0: v4, F1: 0, F2: v2}
    v4 = 1000
    fmt.Println(v4)
    if (42 == v2) {
        v19 := []int{int(v2), (v17 / (((v17 >> ((105) & 63))) | 1)), (v17 << ((255) & 63)), int(v4)}
        var v20 int32 = 99
        _ = v19
        _ = v20
    }
    v4 = v4
    fmt.Println(v3, v4)
    v21 := []int{v17, v17, (v17 / ((v17) | 1))}
    var v22 uint = 1
    v23 := GenS1{F0: v4, F1: 1, F2: v2}
    v3 = v3
    var v24 int64 = 1
    var v25 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
