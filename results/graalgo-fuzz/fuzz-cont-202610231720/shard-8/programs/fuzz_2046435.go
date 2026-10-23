package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
type GenS5 struct { F0 int8 }
func genFn0(p0 int, p1 bool, p2 int64) int {
    switch (p0 << ((0) & 63)) {
    case -9223372036854775808:
        if (0 > 65535) {
            var v6 int8 = -72
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 uint8 = 22
                _ = fi7
                _ = v8
            }
            var v9 int8 = (v6 - -128)
            _ = v6
            _ = v9
        } else {
            var v10 uint8 = 0
            _ = v10
        }
    case 1:
        p0 += (^(p0))
    case 0:
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: ("a" + "hello")}
            v13 := GenS2{F0: 781, F1: (p2 << ((65535) & 63)), F2: 2}
            var v14 int16 = 1
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return (p0 ^ p0)
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 42
    var v3 string = (("" + (("" + "") + "hello")) + (("" + ("foo bar" + "a")) + "hello"))
    var v4 string = (((v3 + v3) + v3) + v3)
    var v5 uint32 = 2
    v15 := []int{-9223372036854775808}
    var v16 uint64 = 0
    v17 := GenS1{F0: v3}
    fmt.Println(v5, v5, v4)
    v2 -= v2
    if len(v15) > 0 {
        v15[0] = int(v5)
    }
    v18 := GenS2{F0: v5, F1: (v1 * int64(v2)), F2: v16}
    switch 766 {
    case -188:
        v19 := GenS2{F0: v5, F1: (v1 ^ v1), F2: v16}
        _ = v19
    default:
        switch "\n\t" {
        case "a":
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS3{F0: 107, F1: v4}
                var v22 string = "hello"
                if false {
                    _ = len(v15)
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        case "hello":
            if (887 < -32768) {
                for fi23 := 0; fi23 < 19; fi23++ {
                    switch (fi23 | fi23) {
                    case -512:
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    default:
                        v2 -= v2
                    }
                    v24 := GenS4{F0: v2, F1: 0, F2: (42 != 128)}
                    if false {
                        var v25 bool = (v16 <= (^(v16)))
                        var v26 string = (((v4 + "\n\t") + v3) + ("\n\t" + ("foo bar" + (v4 + v3))))
                        _ = v25
                        _ = v26
                    }
                    _ = fi23
                    _ = v24
                }
                fmt.Println(v1, v2, v2)
            } else {
                var v27 float64 = -2.25
                for fi28 := 0; fi28 < 14; fi28++ {
                    v29 := GenS5{F0: 20}
                    var v30 int16 = 42
                    var v31 int = (fi28 % (((fi28 + fi28)) | 1))
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = v27
            }
        }
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v16 += (v16 | (v16 ^ 2))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
