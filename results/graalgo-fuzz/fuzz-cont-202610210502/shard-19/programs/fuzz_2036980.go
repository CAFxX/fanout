package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint64) uint32 {
    p1 -= (p1 | p1)
    fmt.Println(p0, p0)
    var v6 int32 = 706
    _ = p0
    _ = p1
    _ = v6
    return 2147483647
}
func genFn1() uint64 {
    var v7 uint = 4294967295
    if (!(!(false || false))) {
        v8 := GenS1{F0: 0}
        if (false && false) {
            if true {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 string = "a"
                    _ = fi9
                    _ = v10
                }
            }
        } else {
            v7 = (v7 ^ 0)
        }
        v7 = (^(v7))
        _ = v8
    }
    fmt.Println(v7)
    _ = v7
    return 4294967295
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]bool)
    var v3 bool = false
    var v4 int8 = -128
    var v5 uint16 = 65535
    v11 := GenS2{F0: 2}
    var v12 int32 = -690
    if ((^(v5)) > (v5 >> ((255) & 63))) {
        v13 := []int{9223372036854775807}
        if len(v13) > 0 {
            v13[0] = -2147483648
        }
        var v14 uint16 = v5
        _ = v13
        _ = v14
    } else {
        delete(v2, "hello")
        v15 := []int{32768, 128, -9223372036854775808}
        _ = v15
    }
    var v16 uint = 65535
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float32 = v1
            _ = fi18
            _ = v19
        }
        var v20 int64 = -9223372036854775808
        var v21 uint64 = 935
        _ = fi17
        _ = v20
        _ = v21
    }
    var v22 int8 = 124
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 int64 = 42
        if (!((v5 * (v5 / ((65535) | 1))) <= (v5 << (((v16 << ((102) & 63))) & 63)))) {
            var v25 float64 = 3.14159
            _ = v25
        } else {
            switch "hello" {
            case "\n\t":
                var v26 string = ("foo bar" + (("\n\t" + "\n\t") + ("\n\t" + "foo bar")))
                _ = v26
            case "":
                var v27 int64 = v24
                _ = v27
            default:
                var v28 float32 = float32(((0.1 * (1e10 - 0.1)) - ((1.5 + (float64(1e10) + 3.14159)) / float64(v1))))
                _ = v28
            }
            v5 -= (v5 - v5)
        }
        v4 = (v4 >> ((42) & 63))
        _ = fi23
        _ = v24
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        v1 = float32(-2.25)
        _ = fi29
    }
    v30 := GenS1{F0: -1}
    var v31 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
