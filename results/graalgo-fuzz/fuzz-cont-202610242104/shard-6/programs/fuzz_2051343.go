package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16) bool {
    p0 = p0
    if ((86 == 190) && (true || false)) {
        v7 := GenS1{F0: 2147483648, F1: -651}
        p0 += p0
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 65535, F1: 0}
            var v10 int32 = int32(p0)
            for fi11 := 0; fi11 < 8; fi11++ {
                if (32767 != -393) {
                    v10 = v10
                    p0 = 581
                }
                _ = fi11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        if false {
            var v12 int8 = 1
            _ = v12
        } else {
            var v13 uint8 = 1
            p0 -= (p0 | p0)
            _ = v13
        }
        p0 -= (^(p0))
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: 2147483647, F1: 42}
        _ = fi14
        _ = v15
    }
    _ = p0
    return true
}

func main() {
    var v1 bool = false
    var v2 uint32 = 766
    var v3 float32 = 0.0
    var v4 int64 = -1
    var v5 int = 9223372036854775807
    v6 := make(map[string]string)
    var v16 int32 = 1
    var v17 int16 = 620
    v6[""] = "foo bar"
    switch (v5 & (v5 + v5)) {
    case -603:
        fmt.Println(v2, v3)
    default:
        v18 := GenS2{F0: -32768, F1: 255}
        _ = v18
    }
    v17 += -785
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := []int{}
        if (!true) {
            v1 = (!(!false))
        } else {
            for fi21 := 0; fi21 < 15; fi21++ {
                v4 = v4
                if v1 {
                    v22 := GenS3{F0: v16}
                    _ = v22
                } else {
                    v23 := GenS1{F0: v2, F1: (v4 >> ((1) & 63))}
                    v24 := GenS3{F0: 1}
                    fmt.Println(v4, v16)
                    _ = v23
                    _ = v24
                }
                var v25 int8 = -128
                v20 = append(v20, (fi19 | v5))
                _ = fi21
                _ = v25
            }
            v26 := v6[""]
            _ = v26
        }
        v1 = v1
        _ = fi19
        _ = v20
    }
    if (!v1) {
        switch -198 {
        case 2147483647:
            fmt.Println(v17)
        case 944:
            v3 += (1.5 + 0.1)
        default:
            var v27 float64 = float64(v3)
            _ = v27
        }
    }
    if v1 {
        if v1 {
            delete(v6, "")
        }
    }
    v1 = v1
    var v28 float32 = v3
    var v29 bool = (v1 || (false || false))
    delete(v6, "hello")
    if v29 {
        v5 = (v5 * v5)
        if v1 {
            v30 := []int{v5, (v5 + (v5 >> ((65535) & 63)))}
            _ = v30
        }
        var v31 int32 = -966
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
