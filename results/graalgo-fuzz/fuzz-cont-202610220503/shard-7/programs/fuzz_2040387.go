package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
func genFn0() uint {
    var v6 bool = (!((("foo bar" + ("foo bar" + "foo bar")) + (("hello" + "a") + "\n\t")) < "\n\t"))
    var v7 uint32 = 666
    if (((-2.25 * (0.0 * 1.5)) == (1.5 * 1e10)) && v6) {
        v6 = (!(-9223372036854775808 == 0))
        v7 += (^(v7))
    }
    v6 = (1e10 < 0.0)
    fmt.Println(v6, v7)
    _ = v6
    _ = v7
    return 976
}

func main() {
    v1 := []int{513, 32767, 1}
    var v2 string = "hello"
    var v3 int32 = 2147483647
    var v4 uint16 = 0
    v5 := []int{42}
    var v8 uint32 = 2147483648
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    fmt.Println(v4, v2)
    v4 = 984
    v9 := GenS1{F0: 174}
    if (18446744073709551615 != 42) {
        var v10 uint64 = 42
        var v11 uint16 = 410
        _ = v10
        _ = v11
    }
    var v12 bool = true
    if false {
        switch 2147483647 {
        case -517:
            var v13 uint64 = 0
            _ = v13
        }
    }
    switch ((v2 + v2) + v2) {
    case "hello":
        v2 = v2
    default:
        v2 = "a"
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        if (v12 && ((3.14159 * (1.5 * float64(float32(0.0)))) <= 0.0)) {
            _ = len(v5)
            var v15 bool = v12
            v16 := GenS2{F0: 42, F1: 0, F2: -85}
            _ = v15
            _ = v16
        } else {
            _ = len(v5)
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := GenS2{F0: -32768, F1: -9223372036854775808, F2: 42}
                _ = fi17
                _ = v18
            }
        }
        v19 := GenS2{F0: 32767, F1: 1, F2: 42}
        if len(v1) > 0 {
            v1[0] = int(v3)
        }
        _ = fi14
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
