package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int8, p1 string) uint32 {
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 2, F1: 1, F2: "\n\t"}
        v10 := GenS1{F0: 391, F1: 2, F2: ("" + "\n\t")}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 0 {
    case 300:
        v11 := []int{-583, 42, 0}
        _ = v11
    case 2147483647:
        fmt.Println(p1)
    case -869:
        v12 := GenS2{F0: 127}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return 327
}

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 uint = 2
    var v3 uint8 = 0
    var v4 int16 = 315
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    var v13 uint16 = 32767
    v14 := []int{-973, 32768, 127, -490}
    if (!true) {
        v3 -= v3
        fmt.Println(v4, v1)
    }
    fmt.Println(v1)
    v2 -= v2
    if false {
        var v15 uint16 = (v13 / ((v13) | 1))
        _ = v15
    } else {
        if ((false || true) || (!(!true))) {
            for fi16 := 0; fi16 < 8; fi16++ {
                v2 = 18446744073709551615
                for fi17 := 0; fi17 < 12; fi17++ {
                    switch (fi17 + 115) {
                    case -9223372036854775808:
                        var v18 int8 = 55
                        _ = v18
                    case 491:
                        v13 = (v13 >> ((v2) & 63))
                    default:
                        v13 += 481
                    }
                    _ = fi17
                }
                _ = fi16
            }
        }
    }
    fmt.Println(v3, v2, v2)
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS2{F0: fi19}
        v21 := GenS3{F0: 42}
        _ = fi19
        _ = v20
        _ = v21
    }
    if false {
        v1 = v1
        if false {
            v4 += (v4 >> ((1) & 63))
            v22 := GenS3{F0: (v3 ^ 42)}
            var v23 bool = true
            _ = v22
            _ = v23
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    } else {
        v4 -= 42
        switch 4294967295 {
        case -1:
            v13 = 2
        case 97:
            for fi24 := 0; fi24 < 2; fi24++ {
                if false {
                    v25 := GenS1{F0: 0, F1: 4294967295, F2: v1}
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    var v26 int32 = 0
                    _ = v25
                    _ = v26
                }
                if (v1 >= (("a" + "foo bar") + "")) {
                    var v27 int64 = 9223372036854775807
                    v3 = v3
                    _ = v27
                } else {
                    _ = len(v6)
                }
                v28 := GenS2{F0: (-(fi24))}
                _ = fi24
                _ = v28
            }
        case 0:
            var v29 int = 42
            _ = v29
        default:
            v30 := GenS3{F0: 255}
            _ = v30
        }
        v1 = ""
    }
    var v31 int = -2147483648
    var v32 string = ((v1 + v1) + "hello")
    if false {
        v5[v31] = (v31 << ((uint(v13)) & 63))
        if len(v14) > 0 {
            v14[0] = 2147483647
        }
    } else {
        var v33 uint8 = 1
        _ = v33
    }
    v32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
