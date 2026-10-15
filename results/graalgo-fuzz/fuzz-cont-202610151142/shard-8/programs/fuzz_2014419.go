package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 int32 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []int{4294967295, 704}
    var v2 int32 = 0
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 3; fi4++ {
        fmt.Println(v2)
        _ = fi4
    }
    var v5 uint = 591
    fmt.Println(v5, v5)
    v3[(("a" + "") + "")] = 42
    var v6 uint = 818
    v7 := GenS1{F0: 42, F1: 4294967295, F2: 65535}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint64 = 696
        if (false || (false && (-128 == -128))) {
            v10 := GenS2{F0: 128, F1: 459, F2: int32(v5)}
            var v11 int32 = (v2 & v2)
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: 137, F1: 9223372036854775807, F2: (-(v2))}
            v13 := GenS1{F0: (v9 << ((65535) & 63)), F1: 0, F2: 2}
            var v14 uint = (v5 - (v6 << ((v6) & 63)))
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v6, v2, v5)
    v2 = -444
    fmt.Println(v5)
    if false {
        v15 := GenS2{F0: 128, F1: 300, F2: v2}
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 13; fi17++ {
                v2 = (v2 ^ v2)
                v3["\n\t"] = 2147483647
                _ = fi17
            }
            var v18 uint32 = 312
            _ = fi16
            _ = v18
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                if (646 < 42) {
                    v2 += (v2 - (^(v2)))
                    if true {
                        v21 := v3[("foo bar" + "hello")]
                        var v22 int16 = -1
                        _ = v21
                        _ = v22
                    }
                } else {
                    _ = len(v1)
                    v23 := GenS1{F0: 42, F1: 65535, F2: 2}
                    _ = v23
                }
                _ = fi20
            }
            if true {
                switch (fi19 % ((963) | 1)) {
                case 9223372036854775807:
                    fmt.Println(v5)
                }
            }
            _ = fi19
        }
        _ = v15
    }
    v24 := GenS3{F0: 255}
    fmt.Println(v2, v2, v6)
    if (0 != 0) {
        v2 = int32(v5)
    } else {
        fmt.Println(v6, v2)
    }
    v2 = v2
    var v25 int64 = -1
    fmt.Println(v5, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
