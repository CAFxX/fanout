package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int; F2 int16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint8, p2 uint8) uint32 {
    p0 = p0
    v4 := GenS1{F0: 42}
    v5 := GenS2{F0: (p2 ^ p2), F1: 4294967295, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 int = 2147483647
    var v2 uint16 = 65535
    v3 := []string{("" + "foo bar"), "\n\t", "\n\t"}
    var v6 int32 = -710
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v6 = 1000
        _ = fi7
    }
    if false {
        v3 = append(v3, (("" + ("\n\t" + ("\n\t" + "foo bar"))) + ("hello" + "\n\t")))
        var v8 int = v1
        _ = v8
    }
    v9 := GenS2{F0: 127, F1: 18446744073709551615, F2: 603}
    v10 := GenS3{F0: 21, F1: 1, F2: 1}
    v2 = (v2 ^ 42)
    var v11 float64 = 0.0
    if (v11 != (v11 * (v11 - (v11 + v11)))) {
        var v12 int64 = -308
        var v13 int32 = v6
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(v11, v6)
            for fi15 := 0; fi15 < 13; fi15++ {
                v1 -= (fi14 << ((255) & 63))
                fmt.Println(fi15, fi15, fi15)
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                _ = fi15
            }
            if (0 > 364) {
                fmt.Println(fi14, fi14)
            } else {
                if true {
                    if len(v3) > 0 {
                        v3[0] = ("hello" + "foo bar")
                    }
                }
                _ = len(v3)
            }
            _ = fi14
        }
        var v16 uint16 = v2
        var v17 uint32 = 889
        _ = v16
        _ = v17
    }
    fmt.Println(v11, v6)
    v11 = v11
    switch (v1 % ((32767) | 1)) {
    case -1:
        var v18 int16 = 32767
        _ = v18
    default:
        v1 -= -495
    }
    var v19 uint = 834
    v20 := GenS4{F0: (v2 << (((v19 % (((v19 * (v19 + 18446744073709551615))) | 1))) & 63)), F1: (v1 << ((873) & 63)), F2: 1000}
    var v21 uint = v19
    v2 = v2
    for fi22 := 0; fi22 < 16; fi22++ {
        v23 := GenS5{F0: 135}
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v6
    _ = v9
}
