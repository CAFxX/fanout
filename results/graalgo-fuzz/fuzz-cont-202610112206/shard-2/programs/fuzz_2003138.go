package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
func genFn0() uint64 {
    var v5 uint8 = 18
    v6 := GenS1{F0: 315, F1: -128}
    _ = v5
    _ = v6
    return 496
}

func main() {
    var v1 uint = 0
    v2 := []int{-9223372036854775808, 42}
    var v3 uint = 1
    var v4 uint32 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 := GenS1{F0: 42, F1: 42}
    v4 -= v4
    v1 = 4294967295
    v8 := GenS2{F0: 1, F1: (true && (2 != 780)), F2: "\n\t"}
    fmt.Println(v1)
    fmt.Println(v3, v4)
    fmt.Println(v3, v3, v4)
    if len(v2) > 0 {
        v2[0] = 32768
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v1 -= (v1 >> (((v3 - 887)) & 63))
        _ = fi9
    }
    v10 := GenS2{F0: 103, F1: (!false), F2: ("\n\t" + (("a" + "hello") + ("\n\t" + "a")))}
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    var v11 int16 = -32768
    for fi12 := 0; fi12 < 14; fi12++ {
        switch fi12 {
        case -2147483648:
            v1 = v1
        }
        v11 += (v11 * (v11 >> ((18446744073709551615) & 63)))
        _ = fi12
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
