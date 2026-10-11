package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 761
    var v3 int32 = -2147483648
    v4 := make(map[string]string)
    v5 := []uint32{v2, v2, (v2 - v2)}
    fmt.Println(v1, v1, v2)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int32 = (v3 + (^(v3)))
        if false {
            v1 -= v1
            v1 = (v1 >> ((2) & 63))
            v4[("\n\t" + "hello")] = "a"
        }
        if false {
            fmt.Println(v1, v2, v3)
            var v8 uint8 = 255
            var v9 int32 = (v3 / ((v3) | 1))
            _ = v8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v2 -= (v2 - v2)
    for fi10 := 0; fi10 < 9; fi10++ {
        if false {
            v11 := GenS1{F0: 738, F1: 4294967295, F2: ((("a" + "a") + "\n\t") + "foo bar")}
            v1 -= (v1 / ((v1) | 1))
            _ = v11
        } else {
            var v12 int = fi10
            _ = len(v4)
            if false {
                v5 = append(v5, v2)
            }
            _ = v12
        }
        _ = len(v5)
        v13 := GenS1{F0: 42, F1: 18446744073709551615, F2: "hello"}
        _ = fi10
        _ = v13
    }
    v2 += (^(v2))
    var v14 string = "foo bar"
    switch ("\n\t" + "") {
    case "\n\t":
        var v15 int16 = 32767
        _ = v15
    }
    v16 := GenS1{F0: 9223372036854775807, F1: 18, F2: (v14 + "\n\t")}
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(v1, v14, v2)
        if false {
            var v18 int16 = -11
            if false {
                v2 = (v2 << ((18446744073709551615) & 63))
                v19 := GenS1{F0: -1, F1: 42, F2: v14}
                if (65535 != 1) {
                    v14 = (((("\n\t" + "\n\t") + ("foo bar" + ("a" + v14))) + "\n\t") + (("foo bar" + (v14 + v14)) + v14))
                    v2 = (v2 + uint32(v1))
                }
                _ = v19
            }
            _ = v18
        }
        _ = fi17
    }
    v2 = v2
    v20 := GenS2{F0: 638}
    fmt.Println(v3, v3)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
