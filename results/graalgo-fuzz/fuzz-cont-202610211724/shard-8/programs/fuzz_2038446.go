package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint; F1 uint16; F2 int }
type GenS3 struct { F0 bool; F1 uint32; F2 uint32 }
func genFn0(p0 int16, p1 uint16, p2 uint32) int {
    switch -1 {
    case 128:
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{(fi7 >> ((2) & 63)), (fi7 >> ((42) & 63)), 959}
            if false {
                p0 -= -288
            } else {
                p2 += p2
                var v9 uint16 = (p1 + 65535)
                _ = v9
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(p0, p1, p0)
    var v10 int = int(p0)
    switch (^(v10)) {
    case 127:
        v11 := GenS1{F0: 355, F1: uint(p1)}
        _ = v11
    case 128:
        v12 := GenS1{F0: 887, F1: 42}
        _ = v12
    case 834:
        fmt.Println(p1)
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 int16 = p0
            if false {
                if (((0.1 + 0.0) <= 0.0) || (4294967295 != 0)) {
                    p2 -= 2147483648
                }
                v15 := GenS1{F0: 1, F1: 65535}
                switch ("hello" + "hello") {
                case "a":
                    var v16 int16 = v14
                    _ = v16
                case "\n\t":
                    var v17 bool = false
                    _ = v17
                }
                _ = v15
            }
            p0 = v14
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return v10
}

func main() {
    var v1 int32 = 2147483647
    var v2 int = 0
    v3 := make(map[string]int)
    var v4 uint32 = 1
    var v5 uint8 = 1
    v6 := []uint32{(v4 % ((v4) | 1)), 42, v4, (v4 << ((1) & 63))}
    v2 -= v2
    fmt.Println(v2)
    if (!false) {
        v1 += (v1 | v1)
        v6 = append(v6, v4)
    }
    var v18 uint16 = 0
    switch (v2 | v2) {
    case -100:
        v19 := GenS2{F0: 4294967295, F1: 255, F2: -9223372036854775808}
        _ = v19
    case 2147483647:
        switch "foo bar" {
        case "\n\t":
            var v20 float32 = (1e10 / 3.14159)
            _ = v20
        default:
            v6 = append(v6, v4)
        }
    }
    v2 -= (-(v2))
    if ((!(!((500 < -544) && true))) || (-80 > -117)) {
        v4 = v4
    }
    var v21 uint32 = 2
    v1 -= v1
    switch -612 {
    case -1:
        for fi22 := 0; fi22 < 19; fi22++ {
            v18 += 1
            var v23 int = 407
            _ = fi22
            _ = v23
        }
    case 42:
        for fi24 := 0; fi24 < 7; fi24++ {
            if (!false) {
                for fi25 := 0; fi25 < 9; fi25++ {
                    _ = len(v6)
                    v3[""] = fi24
                    _ = fi25
                }
                v1 = v1
            } else {
                fmt.Println(v1)
                if true {
                    if (127 >= (v2 + genFn0(98, (v18 / ((32767) | 1)), 2147483647))) {
                        fmt.Println(fi24, v18, fi24)
                        for fi26 := 0; fi26 < 1; fi26++ {
                            fmt.Println(fi26, v4)
                            _ = fi26
                        }
                        v18 += (v18 + v18)
                    }
                }
            }
            _ = fi24
        }
    default:
        fmt.Println(v4, v2, v5)
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v1 = v1
        _ = fi27
    }
    v2 = (v2 | genFn0(-32768, v18, uint32(v18)))
    delete(v3, "")
    switch (v2 & (v2 % ((9223372036854775807) | 1))) {
    case -9223372036854775808:
        switch "" {
        case "a":
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS2{F0: 235, F1: 42, F2: 2147483647}
                var v30 int16 = 1
                for fi31 := 0; fi31 < 16; fi31++ {
                    v32 := GenS3{F0: true, F1: (v4 - 2147483648), F2: (v21 * v21)}
                    _ = len(v6)
                    _ = fi31
                    _ = v32
                }
                _ = fi28
                _ = v29
                _ = v30
            }
        case "foo bar":
            v33 := GenS1{F0: 467, F1: 4294967295}
            _ = v33
        case "":
            _ = len(v6)
        default:
            fmt.Println(v5, v1, v5)
        }
    case -609:
        var v34 int32 = (v1 << ((42) & 63))
        _ = v34
    case 9223372036854775807:
        v5 = (v5 + (^(v5)))
    default:
        v18 += v18
    }
    switch (v2 << ((65535) & 63)) {
    case -1:
        v35 := GenS1{F0: 549, F1: 0}
        _ = v35
    }
    fmt.Println(v18)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
