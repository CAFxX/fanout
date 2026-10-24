package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
func genFn0(p0 uint64) uint16 {
    p0 += p0
    if true {
        fmt.Println(p0)
    }
    _ = p0
    return 255
}

func main() {
    v1 := []uint32{42}
    v2 := make(map[int]string)
    var v3 uint32 = 2147483648
    v4 := make(map[string]string)
    var v5 int64 = 9223372036854775807
    var v6 int = 32768
    if (-1 < 127) {
        v7 := GenS1{F0: 1, F1: v3, F2: ("\n\t" + ("a" + "foo bar"))}
        _ = v7
    } else {
        if (209 > 1) {
            fmt.Println(v6, v5)
            fmt.Println(v5)
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint8 = 42
            _ = len(v1)
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := v2[-2147483648]
        v12 := GenS2{F0: true, F1: 255, F2: 1}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch v6 {
    case 1:
        var v13 float64 = ((0.0 * 0.0) * -2.25)
        _ = v13
    case 127:
        fmt.Println(v5)
    case 42:
        v5 -= 1
    default:
        v3 += 931
    }
    v3 -= 65535
    v3 = v3
    fmt.Println(v5, v5, v3)
    switch "" {
    case "a":
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                if (0 >= 65535) {
                    v6 -= v6
                } else {
                    switch "hello" {
                    case "\n\t":
                        v3 -= v3
                    case "hello":
                        var v16 uint8 = 180
                        _ = v16
                    case "":
                        v6 = fi14
                    }
                    v6 += fi14
                    v4["hello"] = ("\n\t" + ("a" + "foo bar"))
                }
                _ = fi15
            }
            _ = fi14
        }
    case "hello":
        v5 = -9223372036854775808
    case "":
        var v17 uint16 = 95
        _ = v17
    }
    switch ("hello" + "hello") {
    case "":
        v18 := v2[343]
        _ = v18
    case "hello":
        _ = len(v4)
    case "a":
        v5 = -9223372036854775808
    }
    v6 += (v6 | (v6 << ((2) & 63)))
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 int8 = -1
        _ = len(v2)
        fmt.Println(v3, v6, v5)
        _ = fi19
        _ = v20
    }
    var v21 bool = (!(!(true || (true || ((42 <= 249) || false)))))
    var v22 uint32 = 0
    var v23 int64 = v5
    fmt.Println(v23, v5)
    v6 = (v6 << ((1) & 63))
    var v24 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
