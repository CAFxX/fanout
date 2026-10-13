package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 uint; F2 string }
func genFn0(p0 bool, p1 uint16, p2 bool) uint32 {
    v5 := GenS1{F0: 0}
    switch -2147483648 {
    case -1:
        v6 := GenS1{F0: 42}
        _ = v6
    }
    var v7 uint = 699
    switch "hello" {
    case "\n\t":
        var v8 float32 = ((float32(float64(1.5)) - float32((0.0 * ((0.1 + 0.0) / 0.1)))) + 0.0)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 int8) int {
    p1 += p1
    v9 := GenS1{F0: -433}
    v10 := GenS2{F0: 128, F1: 102}
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 453
    v2 := []uint32{65535, 42, 2147483648, 2147483648}
    var v3 uint8 = 128
    var v4 uint8 = 128
    for fi12 := 0; fi12 < 12; fi12++ {
        if false {
            v4 = (v3 ^ 175)
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                v1 -= 42
                var v14 string = ("a" + "\n\t")
                _ = fi13
                _ = v14
            }
        }
        v1 = (v1 ^ (v1 ^ v1))
        v15 := GenS3{F0: "\n\t", F1: 731, F2: "hello"}
        _ = fi12
        _ = v15
    }
    var v16 int32 = 1
    if (143 != 625) {
        switch -9223372036854775808 {
        case 4294967295:
            _ = len(v2)
        default:
            if true {
                var v17 int16 = 856
                v1 = v1
                var v18 int = 230
                _ = v17
                _ = v18
            }
        }
    } else {
        v19 := GenS2{F0: 740, F1: (v3 + v4)}
        _ = v19
    }
    fmt.Println(v4, v3, v3)
    var v20 string = "\n\t"
    fmt.Println(v1)
    for fi21 := 0; fi21 < 14; fi21++ {
        v22 := GenS1{F0: (-(v16))}
        var v23 int8 = -14
        _ = fi21
        _ = v22
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = uint32(v1)
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            fmt.Println(fi25, v1)
            var v26 uint = 0
            v16 += v16
            v4 -= v4
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1)
            fmt.Println(fi24, fi24)
            var v27 int8 = -92
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := GenS3{F0: ("foo bar" + v20), F1: 42, F2: v20}
            var v30 int8 = 95
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi24
    }
    fmt.Println(v3, v3, v4)
    fmt.Println(v1)
    var v31 string = "\n\t"
    for fi32 := 0; fi32 < 20; fi32++ {
        v3 -= 218
        fmt.Println(v4, v3)
        _ = fi32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        if (((true || ((v4 | (v4 << ((255) & 63))) != (^(v4)))) || true) && false) {
            v34 := GenS2{F0: (fi33 * (fi33 ^ 572)), F1: v4}
            if (false || (!(false || true))) {
                v2 = append(v2, 1)
                fmt.Println(v3, fi33)
            }
            var v35 uint32 = 2147483647
            _ = v34
            _ = v35
        } else {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            fmt.Println(v4, v16)
        }
        _ = fi33
    }
    switch 32767 {
    case -2147483648:
        var v36 uint8 = 1
        _ = v36
    case 506:
        if (1e10 == 0.1) {
            if (-9223372036854775808 > 0) {
                v16 -= -417
                v16 = (v16 ^ (v16 & (-(v16))))
            }
            if (0 != -32768) {
                fmt.Println(v20)
                var v37 string = (v20 + "hello")
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 float64 = (0.1 / 0.1)
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
            fmt.Println(v4, v3)
        }
    default:
        fmt.Println(v4)
    }
    v3 -= 2
    if ("a" >= v20) {
        switch v20 {
        case "foo bar":
            v3 = 13
        default:
            v3 = (v4 - (v4 << ((65535) & 63)))
        }
    } else {
        v40 := GenS2{F0: 127, F1: 42}
        var v41 uint16 = 678
        v31 = ""
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
}
