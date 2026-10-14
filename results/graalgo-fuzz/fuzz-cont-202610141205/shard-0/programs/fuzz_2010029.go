package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int32, p1 uint32) string {
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        if true {
            if ((!(!false)) || ("\n\t" > "\n\t")) {
                var v9 uint64 = 1
                v7 = v7
                _ = v9
            }
        } else {
            p0 = p0
            var v10 string = "foo bar"
            _ = v10
        }
        _ = fi8
    }
    var v11 float32 = 0.1
    var v12 int8 = 0
    v11 -= v11
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v7
    return "\n\t"
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = 0.1
    var v3 uint16 = 0
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    var v6 uint = 0
    v13 := []int{4294967295, 128, 659, 32767}
    if (1 > -9223372036854775808) {
        for fi14 := 0; fi14 < 9; fi14++ {
            if len(v13) > 0 {
                v13[0] = (fi14 + (fi14 & fi14))
            }
            var v15 int = (fi14 * -2147483648)
            switch int(v6) {
            case 32768:
                if false {
                    var v16 uint16 = v3
                    for fi17 := 0; fi17 < 15; fi17++ {
                        v3 -= 1
                        fmt.Println(fi14)
                        _ = fi17
                    }
                    switch int(v1) {
                    case 127:
                        for fi18 := 0; fi18 < 18; fi18++ {
                            fmt.Println(v2, v16, v3)
                            v15 += v15
                            v19 := GenS1{F0: -128}
                            _ = fi18
                            _ = v19
                        }
                    case 2147483647:
                        _ = len(v4)
                    case 786:
                        if len(v13) > 0 {
                            v13[0] = v15
                        }
                    default:
                        var v20 int16 = 1000
                        _ = v20
                    }
                    _ = v16
                }
            }
            _ = fi14
            _ = v15
        }
        fmt.Println(v1, v2, v3)
        v21 := GenS2{F0: v1, F1: 1}
        _ = v21
    }
    v6 += v6
    fmt.Println(v2, v2, v1)
    if false {
        var v22 uint8 = uint8(v3)
        _ = v22
    }
    v23 := GenS3{F0: 1}
    var v24 int32 = (v1 << (((v6 - 255)) & 63))
    var v25 int64 = 762
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 string = ("" + "hello")
        fmt.Println(v27)
        fmt.Println(v24, v1, v1)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
