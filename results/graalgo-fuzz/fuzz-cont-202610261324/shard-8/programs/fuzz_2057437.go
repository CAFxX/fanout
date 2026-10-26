package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int }
type GenS2 struct { F0 int64; F1 int32; F2 uint32 }
type GenS3 struct { F0 bool; F1 string }

func main() {
    v1 := []uint32{}
    var v2 int64 = -9223372036854775808
    v3 := []uint32{65535, 646, 143}
    if true {
        for fi4 := 0; fi4 < 13; fi4++ {
            if true {
                if (127 > 10) {
                    v2 -= (-(v2))
                } else {
                    v5 := GenS1{F0: int8(v2), F1: 32767, F2: fi4}
                    if true {
                        fmt.Println(v2, fi4)
                    }
                    _ = v5
                }
                v2 -= (^(v2))
            }
            _ = len(v1)
            _ = fi4
        }
    }
    v2 -= v2
    v6 := GenS1{F0: 42, F1: 1, F2: 32768}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 0, F1: 2, F2: fi7}
        switch (fi7 / (((fi7 + (fi7 << ((0) & 63)))) | 1)) {
        case 127:
            for fi9 := 0; fi9 < 18; fi9++ {
                fmt.Println(fi7, fi7, fi9)
                var v10 uint = 883
                if (!(32767 > (fi9 | (fi9 << ((v10) & 63))))) {
                    v11 := GenS2{F0: (v2 & 0), F1: 2147483647, F2: 995}
                    _ = v11
                } else {
                    fmt.Println(v10, fi7, fi7)
                    v10 += v10
                    v2 -= (v2 / ((v2) | 1))
                    if true {
                        v2 = v2
                        fmt.Println(v2, v10)
                        v10 = 996
                        fmt.Println(v10, v2)
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 65535
                        }
                        var v12 int = (fi7 * (-(fi9)))
                        _ = v12
                    }
                }
                _ = fi9
                _ = v10
            }
        }
        _ = fi7
        _ = v8
    }
    v2 += (v2 / (((-(v2))) | 1))
    v13 := GenS3{F0: (2147483647 > 619), F1: (("" + (("foo bar" + "hello") + "hello")) + "a")}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint64 = 2
        _ = fi14
        _ = v15
    }
    if ((((2 < 255) && (v2 < (v2 ^ v2))) || true) || false) {
        var v16 uint8 = 43
        fmt.Println(v2, v2)
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 uint64 = 986
            _ = fi17
            _ = v18
        }
    } else {
        v2 = (v2 & -9223372036854775808)
    }
    v2 -= (^(v2))
    var v19 int = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
