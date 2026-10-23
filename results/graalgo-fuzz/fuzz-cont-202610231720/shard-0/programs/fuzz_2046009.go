package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint8) int8 {
    var v7 uint = 628
    var v8 uint8 = (p0 ^ 239)
    if (4294967295 == 2) {
        v8 = (v8 | 23)
        v9 := GenS1{F0: 18446744073709551615, F1: 42, F2: 65535}
        p0 -= (v8 & 0)
        _ = v9
    } else {
        fmt.Println(v7, p0)
    }
    _ = p0
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int16 = 32767
    v2 := []string{"foo bar", (("" + "a") + "hello")}
    var v3 uint32 = 2
    var v4 uint8 = 127
    var v5 uint16 = 2
    var v6 float32 = (0.1 - (1.5 + 1.5))
    v10 := GenS1{F0: 18446744073709551615, F1: 0, F2: 4294967295}
    for fi11 := 0; fi11 < 10; fi11++ {
        _ = len(v2)
        fmt.Println(v6)
        switch (fi11 ^ fi11) {
        case -495:
            v4 = (^(v4))
        }
        _ = fi11
    }
    _ = len(v2)
    v1 += v1
    var v12 uint64 = uint64(v3)
    if true {
        v1 = v1
        var v13 uint32 = 0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v13
    }
    switch -965 {
    case 127:
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: 2, F1: int8(v1)}
            _ = fi14
            _ = v15
        }
    case -825:
        v3 += 556
    }
    v12 -= (v12 ^ (v12 >> ((2) & 63)))
    v12 = uint64(v5)
    var v16 uint16 = (v5 >> ((1) & 63))
    v12 -= (v12 / ((2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
