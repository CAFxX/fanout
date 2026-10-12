package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0(p0 int64, p1 int16) uint8 {
    var v6 uint64 = 0
    var v7 uint32 = 0
    v6 -= v6
    var v8 int16 = (^(p1))
    v9 := GenS1{F0: -2147483648, F1: 255}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 2
    var v2 uint8 = v1
    var v3 uint16 = 980
    var v4 int = -24
    var v5 int64 = -9223372036854775808
    fmt.Println(v5, v4)
    var v10 int64 = v5
    for fi11 := 0; fi11 < 8; fi11++ {
        v1 = uint8(v3)
        if false {
            v1 = v1
            var v12 int64 = v5
            _ = v12
        }
        v3 -= v3
        _ = fi11
    }
    v13 := []int{v4, 4294967295, 32768}
    if (!true) {
        var v14 int8 = 42
        switch ("\n\t" + "") {
        case "hello":
            v1 = (v2 << ((877) & 63))
        case "a":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        case "foo bar":
            var v15 int32 = -2147483648
            _ = v15
        }
        switch v4 {
        case -80:
            if false {
                v14 = (v14 << ((1) & 63))
            }
        case -442:
            v2 += v1
        case 791:
            v16 := GenS1{F0: int32(v14), F1: 65535}
            _ = v16
        }
        _ = v14
    }
    switch "foo bar" {
    case "":
        v17 := GenS1{F0: -914, F1: 18446744073709551615}
        _ = v17
    case "a":
        switch (-(v4)) {
        case 127:
            var v18 uint32 = 42
            _ = v18
        case 128:
            var v19 uint8 = (v2 + (v2 * genFn0(9223372036854775807, 1)))
            _ = v19
        }
    }
    v3 = 0
    _ = len(v13)
    var v20 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
