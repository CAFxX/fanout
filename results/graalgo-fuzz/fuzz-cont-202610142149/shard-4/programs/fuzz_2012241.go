package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint64) uint32 {
    v5 := []int{128, -1}
    v5 = append(v5, -481)
    _ = p0
    _ = v5
    return 2147483647
}
func genFn1(p0 int32, p1 uint) uint32 {
    v6 := []int{1, 104, -958, -9223372036854775808}
    switch "hello" {
    case "\n\t":
        v6 = append(v6, 724)
    }
    var v7 int64 = -9223372036854775808
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v8 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]bool)
    var v3 float32 = (0.0 * 1e10)
    v4 := make(map[string]bool)
    var v9 uint32 = 42
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 float64 = float64(3.14159)
        delete(v2, 65535)
        _ = fi10
        _ = v11
    }
    var v12 int64 = 9223372036854775807
    var v13 uint = 255
    var v14 float32 = ((v3 + (0.1 / v3)) * (((float32(v1) * float32(v1)) * ((float32(float64(0.0)) / v3) * 0.1)) / v3))
    if ((((((v3 / v3) / -2.25) / v14) / 3.14159) <= ((v14 / ((float32(3.14159) + 3.14159) / ((v14 + -2.25) + 0.0))) / 0.1)) || false) {
        var v15 uint64 = 42
        _ = v15
    }
    switch 32767 {
    case -927:
        var v16 uint16 = 662
        _ = v16
    case 0:
        fmt.Println(v1, v13)
    case 4294967295:
        v9 -= (v9 & v9)
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        switch "" {
        case "":
            for fi18 := 0; fi18 < 18; fi18++ {
                fmt.Println(v9, v14, v9)
                switch (fi17 % ((128) | 1)) {
                case -1:
                    if (true && (false || true)) {
                        if false {
                            var v19 float32 = ((0.1 / v14) / float32(v1))
                            switch 9223372036854775807 {
                            case 32768:
                                var v20 uint16 = 339
                                _ = v20
                            case 9223372036854775807:
                                if false {
                                    v4["\n\t"] = false
                                } else {
                                    v3 = (v3 + float32(-2.25))
                                    v19 = v14
                                }
                            case -9223372036854775808:
                                v12 = (v12 * v12)
                            }
                            v12 -= v12
                            _ = len(v2)
                            _ = v19
                        } else {
                            delete(v2, (v9 ^ genFn0(2)))
                            v21 := []int{fi18, 127, (fi17 >> ((4294967295) & 63)), -8}
                            _ = v21
                        }
                    }
                case -2147483648:
                    var v22 int32 = -39
                    _ = v22
                }
                var v23 int64 = int64(fi18)
                _ = fi18
                _ = v23
            }
        }
        _ = fi17
    }
    v1 = (v1 - (((v1 * ((v1 / 1e10) + (v1 - v1))) / v1) - v1))
    v3 = 3.14159
    for fi24 := 0; fi24 < 2; fi24++ {
        v14 += ((float32((v1 - v1)) - 1e10) + v3)
        switch 42 {
        case -2147483648:
            v13 -= 0
        case 9223372036854775807:
            var v25 uint16 = 65535
            _ = v25
        }
        _ = fi24
    }
    var v26 uint = v13
    v27 := GenS1{F0: 147}
    var v28 uint32 = v9
    var v29 int = 454
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
