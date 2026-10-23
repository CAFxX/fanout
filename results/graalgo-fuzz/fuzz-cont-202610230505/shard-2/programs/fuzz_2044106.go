package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint32; F2 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint8; F2 uint32 }
func genFn0(p0 int32, p1 uint16) uint32 {
    p1 = p1
    p0 -= (p0 + 0)
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := []uint32{563, 4294967295, 42, 1}
    var v2 uint16 = 1
    var v3 int64 = 0
    v4 := make(map[int]int)
    if false {
        if len(v1) > 0 {
            v1[0] = 548
        }
    } else {
        fmt.Println(v3, v2)
    }
    v5 := GenS1{F0: ((((("foo bar" + ("\n\t" + "")) + "") + "\n\t") + "\n\t") + "a"), F1: -585, F2: -1}
    var v6 int64 = v3
    var v7 int64 = v3
    if (91 < 0) {
        v8 := GenS2{F0: 0}
        v7 -= (v3 + v3)
        _ = v8
    } else {
        v6 = (v6 % ((42) | 1))
        v9 := GenS3{F0: (v2 % ((v2) | 1)), F1: "a", F2: 65535}
        _ = v9
    }
    if (!(v2 != v2)) {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        v10 := GenS2{F0: 2147483647}
        v1 = append(v1, 486)
        _ = v10
    }
    var v11 int = -2147483648
    v3 = v6
    v12 := GenS4{F0: 1, F1: 291, F2: false}
    v6 -= (v6 + v7)
    _ = len(v1)
    fmt.Println(v11, v7, v11)
    fmt.Println(v11, v6, v11)
    fmt.Println(v7, v3)
    if true {
        switch "" {
        case "foo bar":
            var v13 int32 = -2147483648
            _ = v13
        case "a":
            for fi14 := 0; fi14 < 1; fi14++ {
                for fi15 := 0; fi15 < 18; fi15++ {
                    if false {
                        if (false && true) {
                            v3 += v3
                        }
                    }
                    if (v3 != 9223372036854775807) {
                        v16 := GenS1{F0: "a", F1: fi14, F2: 42}
                        v17 := GenS5{F0: 32767}
                        _ = v16
                        _ = v17
                    }
                    _ = fi15
                }
                _ = fi14
            }
        default:
            v2 = v2
        }
        if true {
            if true {
                switch v11 {
                case -100:
                    var v18 uint8 = 128
                    _ = v18
                default:
                    v19 := GenS1{F0: "", F1: (v11 + -768), F2: 0}
                    _ = v19
                }
                var v20 float64 = -2.25
                var v21 bool = ((("" >= "a") && (32767 > 111)) && (2147483648 >= 2))
                _ = v20
                _ = v21
            }
            v6 = v3
        } else {
            var v22 int8 = 1
            switch (v11 & v11) {
            case 42:
                v2 = v2
            case 2147483647:
                v23 := GenS3{F0: v2, F1: "a", F2: 65535}
                _ = v23
            default:
                v24 := GenS6{F0: 128, F1: 42, F2: 1}
                _ = v24
            }
            if (true || false) {
                if len(v1) > 0 {
                    v1[0] = 1
                }
            }
            _ = v22
        }
        var v25 uint32 = genFn0(254, (v2 | v2))
        _ = v25
    } else {
        v6 += v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
