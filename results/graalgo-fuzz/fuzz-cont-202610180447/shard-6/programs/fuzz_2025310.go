package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 int32) uint8 {
    p0 += (^(p1))
    var v7 int = -2147483648
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int32 = (p0 * p1)
        v10 := []int{(v7 - (v7 % (((v7 >> ((42) & 63))) | 1)))}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    v1 := []string{("" + (("\n\t" + ("foo bar" + "\n\t")) + ("hello" + ("" + "a"))))}
    var v2 int8 = -81
    var v3 int8 = (v2 ^ (v2 + (v2 & v2)))
    var v4 uint32 = 4294967295
    var v5 uint8 = 128
    var v6 uint64 = 1000
    switch 32767 {
    case -2147483648:
        v6 -= (^(v6))
    case 724:
        _ = len(v1)
    case 0:
        for fi11 := 0; fi11 < 2; fi11++ {
            switch (-(fi11)) {
            case 32767:
                var v12 uint32 = 0
                _ = v12
            case 128:
                fmt.Println(v2)
            default:
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            }
            _ = fi11
        }
    default:
        var v13 float64 = float64(1.5)
        _ = v13
    }
    switch int(v3) {
    case 9223372036854775807:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -950:
        var v14 int32 = 2147483647
        _ = v14
    case -2147483648:
        var v15 int16 = 0
        _ = v15
    }
    v4 += 2
    var v16 float64 = 0.0
    fmt.Println(v6, v6, v5)
    var v17 int16 = -370
    fmt.Println(v5, v3)
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint16 = 65535
        v2 -= (v2 % (((-(v3))) | 1))
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 65535
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    var v21 string = ("foo bar" + "a")
    if len(v1) > 0 {
        v1[0] = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
