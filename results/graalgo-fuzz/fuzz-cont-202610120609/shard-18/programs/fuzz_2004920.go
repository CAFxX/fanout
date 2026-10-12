package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint32 = 188
    v2 := []uint32{v1, (v1 % ((v1) | 1)), 46, v1}
    v3 := make(map[string]uint32)
    var v4 uint8 = 255
    var v5 uint16 = 255
    var v6 uint32 = v1
    if true {
        var v7 uint8 = 1
        _ = v7
    }
    v5 += 426
    v8 := GenS1{F0: "\n\t"}
    fmt.Println(v5, v4, v6)
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint16 = (v5 >> ((1) & 63))
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 string = "\n\t"
            var v13 bool = (!((v6 >> ((390) & 63)) > 2147483648))
            if v13 {
                v13 = v13
                v4 -= (v4 - 255)
                v6 -= (v1 * 42)
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: "a"}
        _ = fi9
        _ = v10
        _ = v14
    }
    var v15 uint = 1
    v4 -= 255
    for fi16 := 0; fi16 < 15; fi16++ {
        if false {
            var v17 int64 = 42
            switch (fi16 - -9223372036854775808) {
            case 0:
                var v18 uint8 = 111
                _ = v18
            case 32767:
                if false {
                    fmt.Println(v15, v17, v1)
                }
            default:
                fmt.Println(v6, v4)
            }
            switch (fi16 - fi16) {
            case 128:
                v2 = append(v2, uint32(v4))
            case 102:
                v19 := GenS2{F0: int8(v5)}
                _ = v19
            default:
                var v20 int = (fi16 - fi16)
                _ = v20
            }
            _ = v17
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 uint = 255
            v23 := GenS2{F0: 56}
            var v24 float64 = 0.1
            if true {
                if len(v2) > 0 {
                    v2[0] = (v6 % (((v1 * (v6 % ((2147483648) | 1)))) | 1))
                }
                var v25 uint = uint(v4)
                if ((uint8(v5) == (^(v4))) || false) {
                    var v26 int32 = 2147483647
                    _ = v26
                }
                fmt.Println(v22, v1)
                _ = v25
            } else {
                var v27 bool = true
                if true {
                    _ = len(v2)
                    delete(v3, "foo bar")
                } else {
                    v6 += 2147483648
                    fmt.Println(v1, v4, v22)
                    v27 = (false && false)
                    v28 := GenS1{F0: "hello"}
                    _ = v28
                }
                var v29 int32 = 2147483647
                var v30 uint32 = 0
                v5 = (^(v5))
                _ = v27
                _ = v29
                _ = v30
            }
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi16
    }
    var v31 uint64 = 336
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 uint16 = v5
        _ = fi32
        _ = v33
    }
    v34 := v3[("hello" + "\n\t")]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
