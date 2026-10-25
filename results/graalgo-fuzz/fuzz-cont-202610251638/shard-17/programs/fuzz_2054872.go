package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint16 }
type GenS2 struct { F0 string }
func genFn0() int32 {
    var v5 bool = (128 >= 63)
    var v6 uint32 = 46
    v7 := GenS1{F0: 42, F1: (!false), F2: 649}
    _ = v5
    _ = v6
    _ = v7
    return -994
}
func genFn1(p0 string, p1 bool, p2 uint) uint32 {
    for fi8 := 0; fi8 < 7; fi8++ {
        p1 = (!false)
        v9 := GenS1{F0: 2147483647, F1: ((!false) || p1), F2: 250}
        switch 32767 {
        case -9223372036854775808:
            if (4294967295 <= (p2 + p2)) {
                var v10 int8 = 84
                for fi11 := 0; fi11 < 6; fi11++ {
                    v12 := []int{fi8, 463, 826, (fi11 % ((fi11) | 1))}
                    v12 = append(v12, fi11)
                    var v13 int8 = (v10 ^ (v10 & 0))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                var v14 int64 = -1
                _ = v10
                _ = v14
            }
        case 619:
            var v15 uint32 = 2
            _ = v15
        case 822:
            p0 = p0
        default:
            fmt.Println(p0, p2, p1)
        }
        _ = fi8
        _ = v9
    }
    if false {
        var v16 int = 9223372036854775807
        if p1 {
            p2 = p2
        } else {
            if p1 {
                v16 += v16
                for fi17 := 0; fi17 < 13; fi17++ {
                    v16 -= (^(v16))
                    p1 = p1
                    _ = fi17
                }
            } else {
                for fi18 := 0; fi18 < 13; fi18++ {
                    if (223 > (^(p2))) {
                        var v19 uint16 = 2
                        v20 := GenS2{F0: "hello"}
                        _ = v19
                        _ = v20
                    }
                    p2 -= p2
                    _ = fi18
                }
            }
            var v21 int16 = 350
            for fi22 := 0; fi22 < 12; fi22++ {
                fmt.Println(p1, p0, p1)
                for fi23 := 0; fi23 < 6; fi23++ {
                    fmt.Println(v21, fi22)
                    fmt.Println(v16, p2, p1)
                    _ = fi23
                }
                v21 += v21
                _ = fi22
            }
            _ = v21
        }
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := []int{564, 4294967295}
    var v2 uint8 = 138
    var v3 int = 1
    var v4 int64 = 983
    var v24 uint32 = genFn1("\n\t", (("a" != "a") || false), 65535)
    v24 = 2147483647
    _ = len(v1)
    v24 -= v24
    var v25 uint32 = (v24 << ((266) & 63))
    var v26 string = "a"
    switch (v3 << ((18446744073709551615) & 63)) {
    case 330:
        v26 = v26
    case 0:
        _ = len(v1)
    case -168:
        switch v26 {
        case "\n\t":
            v24 = v24
        case "foo bar":
            var v27 uint16 = 42
            _ = v27
        }
    }
    fmt.Println(v3)
    var v28 uint8 = v2
    var v29 float64 = 0.1
    var v30 int = v3
    v25 += 1
    switch ((v26 + ("a" + v26)) + v26) {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = v3
        }
    case "foo bar":
        for fi31 := 0; fi31 < 19; fi31++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi31
        }
    }
    v29 -= v29
    v1 = append(v1, v30)
    var v32 float32 = (3.14159 / ((((0.1 + 3.14159) - 3.14159) - 0.1) / 1.5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
