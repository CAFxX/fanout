package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int64) uint16 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 += p0
        if false {
            var v8 int8 = -106
            switch "" {
            case "":
                switch fi7 {
                case 127:
                    p0 = p0
                case 0:
                    var v9 uint = 65535
                    _ = v9
                }
            }
            if true {
                var v10 uint8 = 127
                fmt.Println(fi7)
                var v11 float64 = ((((((1.5 + 0.1) * -2.25) * (1.5 * -2.25)) + 0.1) / float64((((1.5 - 0.1) / -2.25) / -2.25))) + ((float64(0.0) + 1e10) / 1e10))
                _ = v10
                _ = v11
            } else {
                if false {
                    if (32767 <= 1) {
                        v8 += v8
                        var v12 int8 = v8
                        _ = v12
                    }
                    if false {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            v8 = v8
                            var v14 float64 = 3.14159
                            var v15 int = (fi13 / ((-496) | 1))
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        fmt.Println(p0, v8, v8)
                    }
                    fmt.Println(fi7)
                }
                var v16 bool = (374 >= 420)
                _ = v16
            }
            _ = v8
        } else {
            v17 := []int{42, 32768}
            if (!true) {
                p0 -= p0
                p0 -= p0
                var v18 uint32 = 2147483647
                _ = v18
            }
            switch (fi7 << ((1) & 63)) {
            case 0:
                fmt.Println(p0, p0)
            case 475:
                for fi19 := 0; fi19 < 2; fi19++ {
                    p0 = 9223372036854775807
                    _ = fi19
                }
            }
            _ = v17
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            p0 -= (p0 % ((-112) | 1))
            _ = fi20
        }
        _ = fi7
    }
    var v21 uint8 = 0
    var v22 int16 = -132
    _ = p0
    _ = v21
    _ = v22
    return 0
}

func main() {
    var v1 uint8 = 42
    var v2 int16 = 0
    var v3 float32 = (1e10 + 0.0)
    v4 := []int{-539, 191, -2147483648}
    var v5 bool = (4294967295 <= 0)
    var v6 int32 = 2147483647
    if (3.14159 >= v3) {
        v23 := GenS1{F0: -40}
        v1 = 2
        _ = v23
    }
    var v24 int64 = -9223372036854775808
    var v25 uint32 = 42
    for fi26 := 0; fi26 < 5; fi26++ {
        for fi27 := 0; fi27 < 3; fi27++ {
            v25 -= 1
            _ = fi27
        }
        switch (fi26 << ((uint(v25)) & 63)) {
        case 1:
            var v28 int16 = 1000
            _ = v28
        default:
            v1 = 128
        }
        for fi29 := 0; fi29 < 14; fi29++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v24 += (v24 - 1)
            _ = fi29
        }
        _ = fi26
    }
    v30 := GenS1{F0: 124}
    v1 += 255
    v25 -= (^(v25))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
