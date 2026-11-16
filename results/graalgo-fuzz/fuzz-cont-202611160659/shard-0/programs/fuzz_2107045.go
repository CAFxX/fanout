package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]bool)
    v3 := []string{}
    var v4 int8 = -89
    v4 += (^(v4))
    v5 := GenS1{F0: 0, F1: 0, F2: 964}
    v6 := GenS1{F0: 18446744073709551615, F1: 150, F2: 735}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 += (v4 / (((v4 - (v4 << ((2) & 63)))) | 1))
    v2[("hello" + ("foo bar" + ""))] = ((v1 == (v1 % ((v1) | 1))) && true)
    v1 -= (^(v1))
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int8 = v4
            _ = fi8
            _ = v9
        }
        v1 = v1
        _ = fi7
    }
    var v10 uint32 = 65535
    v11 := GenS2{F0: -764}
    v12 := GenS3{F0: 65535}
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 int = 32768
            var v16 uint8 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int32 = -888
        if len(v3) > 0 {
            v3[0] = ("hello" + (("a" + "a") + ("" + "a")))
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v10, v1)
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
