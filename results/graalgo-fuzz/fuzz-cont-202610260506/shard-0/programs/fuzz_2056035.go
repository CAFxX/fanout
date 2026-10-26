package main

import (
    "fmt"
)


func main() {
    var v1 float64 = (0.0 + 0.1)
    var v2 uint64 = 1
    var v3 int16 = 32767
    var v4 float32 = ((1e10 * 0.1) * float32((v1 - float64(0.0))))
    v5 := []int{257}
    _ = len(v5)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            if false {
                v3 = (v3 * (v3 - v3))
                switch (fi6 + 9223372036854775807) {
                case 32768:
                    v3 = -800
                case -2147483648:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case -1:
                    var v8 uint = 42
                    _ = v8
                }
                var v9 uint64 = (v2 << ((65535) & 63))
                _ = v9
            }
            _ = fi7
        }
        if true {
            var v10 int16 = (-(v3))
            _ = v10
        } else {
            var v11 uint = 2
            var v12 int16 = v3
            if (true || true) {
                v12 += (v3 - v12)
                fmt.Println(v3, v12)
                if false {
                    if false {
                        v11 -= (v11 + 203)
                        v2 += (v2 ^ (v2 / (((v2 << (((v11 + (v11 * (v11 - (v11 ^ v11))))) & 63))) | 1)))
                    }
                    var v13 string = "\n\t"
                    if true {
                        if (false && true) {
                            fmt.Println(v1)
                        } else {
                            var v14 int16 = v3
                            var v15 uint16 = 731
                            _ = v14
                            _ = v15
                        }
                    }
                    _ = v13
                }
            } else {
                v11 -= 65535
                v4 += ((0.1 + (float32(v1) - -2.25)) / v4)
                v5 = append(v5, (fi6 * fi6))
            }
            _ = v11
            _ = v12
        }
        if len(v5) > 0 {
            v5[0] = fi6
        }
        _ = fi6
    }
    var v16 uint32 = 1
    var v17 uint8 = 255
    var v18 uint32 = (^(v16))
    fmt.Println(v4)
    for fi19 := 0; fi19 < 8; fi19++ {
        _ = len(v5)
        fmt.Println(v1, v3)
        if len(v5) > 0 {
            v5[0] = (fi19 << ((19) & 63))
        }
        _ = fi19
    }
    if true {
        var v20 int16 = -760
        for fi21 := 0; fi21 < 18; fi21++ {
            v18 += (^(v18))
            _ = fi21
        }
        if len(v5) > 0 {
            v5[0] = 42
        }
        _ = v20
    } else {
        var v22 float32 = (v4 * ((0.1 + v4) - 1e10))
        var v23 uint64 = (v2 - (v2 | 2))
        var v24 uint64 = v23
        _ = v22
        _ = v23
        _ = v24
    }
    if (true && false) {
        fmt.Println(v1, v2, v2)
        var v25 float32 = 3.14159
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
