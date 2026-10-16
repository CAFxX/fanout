package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint64) uint32 {
    p1 = p1
    var v7 uint = 522
    _ = p0
    _ = p1
    _ = v7
    return 2147483648
}

func main() {
    var v1 int = 0
    v2 := []string{(("\n\t" + "") + "foo bar"), "a", "foo bar", "a"}
    v3 := make(map[uint32]int)
    var v4 uint16 = 32767
    var v5 uint8 = 2
    v6 := make(map[uint32]string)
    switch "foo bar" {
    case "foo bar":
        v8 := v3[203]
        _ = v8
    case "hello":
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int8 = 26
            _ = fi9
            _ = v10
        }
    case "\n\t":
        if false {
            if false {
                var v11 uint16 = (v4 - (v4 ^ v4))
                v12 := GenS1{F0: true, F1: 100, F2: 4294967295}
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                if (true || (false || false)) {
                    v5 -= (v5 & 210)
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        _ = fi14
                    }
                } else {
                    v15 := GenS2{F0: 2, F1: (^(v1))}
                    if (0 == 1) {
                        v1 = (v1 & 42)
                        v16 := GenS3{F0: (fi13 >> ((42) & 63)), F1: 1, F2: 4294967295}
                        _ = v16
                    }
                    _ = v15
                }
                v4 += v4
                v5 -= 2
                _ = fi13
            }
            switch (v1 * (^(v1))) {
            case 4294967295:
                var v17 uint16 = 42
                _ = v17
            case -1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        } else {
            var v18 int = int(v5)
            var v19 uint64 = 1
            _ = v18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 9; fi20++ {
            switch fi20 {
            case -2147483648:
                fmt.Println(fi20, v1)
            case -9223372036854775808:
                fmt.Println(v4)
            case 9223372036854775807:
                v4 -= 0
            default:
                var v21 uint = 0
                _ = v21
            }
            if false {
                v3[2147483648] = v1
                v22 := GenS4{F0: 472}
                _ = v22
            } else {
                switch ((("foo bar" + "\n\t") + ("a" + ("\n\t" + ("" + "a")))) + ("hello" + "foo bar")) {
                case "\n\t":
                    v4 = (^(v4))
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            _ = fi20
        }
    }
    v5 = v5
    var v23 int16 = -1
    v1 = v1
    var v24 uint16 = (v4 * v4)
    v1 -= (v1 & (-(v1)))
    var v25 int16 = (v23 - -32768)
    delete(v3, 65535)
    switch "hello" {
    case "hello":
        var v26 bool = (((v1 ^ int(v25)) <= 4294967295) && true)
        _ = v26
    case "":
        fmt.Println(v24, v25)
    case "foo bar":
        var v27 int16 = v25
        _ = v27
    default:
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 uint8 = uint8(v4)
            _ = fi28
            _ = v29
        }
    }
    v6[4294967295] = (("a" + "foo bar") + "foo bar")
    fmt.Println(v1, v25, v23)
    if (!(true || (18446744073709551615 == 489))) {
        v25 = v25
    } else {
        fmt.Println(v5, v23)
        v4 -= v4
    }
    var v30 uint = 1
    if (v5 < v5) {
        v24 += v4
    } else {
        switch v1 {
        case 32767:
            if (-236 <= -1) {
                v5 -= v5
            }
        case 0:
            if false {
                var v31 float64 = float64((0.0 + (float32(0.0) + 0.1)))
                _ = v31
            }
        case -1:
            delete(v6, genFn0(-344, 543))
        }
        switch (v1 - 1) {
        case 4294967295:
            var v32 int16 = v25
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
