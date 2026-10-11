package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int64 }
type GenS2 struct { F0 bool }
func genFn0() string {
    v4 := GenS1{F0: ("\n\t" + ""), F1: -692, F2: -519}
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: "hello", F1: 1, F2: -9223372036854775808}
        var v7 uint16 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return "\n\t"
}
func genFn1(p0 int32) uint8 {
    p0 += 2147483647
    for fi8 := 0; fi8 < 20; fi8++ {
        p0 = (p0 & p0)
        _ = fi8
    }
    if false {
        p0 -= -2147483648
    }
    var v9 uint16 = 710
    _ = p0
    _ = v9
    return 222
}

func main() {
    v1 := []int{32767, 128, 1, -222}
    var v2 uint64 = 449
    var v3 uint8 = 42
    if (!(2147483648 >= 42)) {
        if (65535 < 65535) {
            if (527 <= 18446744073709551615) {
                v3 = genFn1(-1)
                v3 += 2
                fmt.Println(v3)
            } else {
                v10 := GenS2{F0: true}
                fmt.Println(v3)
                switch "\n\t" {
                case "hello":
                    v2 = 18446744073709551615
                }
                _ = v10
            }
            v3 -= v3
        } else {
            var v11 uint32 = 2
            fmt.Println(v11)
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS2{F0: true}
            switch "a" {
            case "a":
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(fi12, v3, fi14)
                    _ = fi14
                }
            }
            var v15 uint32 = 2147483648
            _ = fi12
            _ = v13
            _ = v15
        }
    }
    v2 = (v2 << ((1) & 63))
    v16 := GenS2{F0: (62 <= 65535)}
    var v17 string = "hello"
    var v18 uint8 = (^(v3))
    v3 += (^(v3))
    if (287 >= v2) {
        if (-1 <= -16) {
            var v19 int8 = -128
            _ = v19
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v20 uint8 = genFn1(0)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
