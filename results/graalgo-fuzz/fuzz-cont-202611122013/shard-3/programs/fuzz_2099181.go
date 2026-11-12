package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 int32 }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0(p0 bool, p1 int32, p2 int16) uint {
    if (42 != -1) {
        for fi4 := 0; fi4 < 6; fi4++ {
            v5 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807, F2: 2147483648}
            v6 := GenS1{F0: 1, F1: fi4, F2: 484}
            _ = fi4
            _ = v5
            _ = v6
        }
        if (p0 && p0) {
            p1 -= -1
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := []int{fi7, 42, (fi7 >> ((193) & 63)), fi7}
                v8 = append(v8, (fi7 | 32768))
                _ = fi7
                _ = v8
            }
        } else {
            v9 := GenS1{F0: 4294967295, F1: 0, F2: 0}
            var v10 int = -369
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            p2 = (p2 * (p2 >> ((433) & 63)))
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int16 = p2
                if p0 {
                    var v14 bool = (float32(((1e10 - 3.14159) / (-2.25 + 1.5))) >= 3.14159)
                    var v15 uint16 = 65535
                    _ = v14
                    _ = v15
                }
                if (1 <= 189) {
                    fmt.Println(p2, p2)
                }
                _ = fi12
                _ = v13
            }
            var v16 float32 = 0.0
            _ = fi11
            _ = v16
        }
    }
    if ((-(p2)) < p2) {
        fmt.Println(p2, p1, p1)
        var v17 int16 = p2
        _ = v17
    }
    var v18 uint64 = 976
    var v19 int = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 556
}
func genFn1(p0 uint16) int8 {
    if true {
        var v20 uint8 = 2
        _ = v20
    }
    switch -438 {
    case 127:
        fmt.Println(p0)
    case 9223372036854775807:
        p0 -= (p0 & p0)
    }
    var v21 uint64 = 2
    _ = p0
    _ = v21
    return -128
}

func main() {
    var v1 uint8 = 255
    var v2 int64 = 0
    var v3 uint32 = 0
    switch "hello" {
    case "a":
        if (1.5 > 3.14159) {
            var v22 float64 = (3.14159 - 0.0)
            if false {
                v3 += v3
            } else {
                var v23 int32 = 1
                _ = v23
            }
            _ = v22
        } else {
            var v24 int16 = -593
            v25 := []int{128, 32768, 1}
            _ = v24
            _ = v25
        }
    case "hello":
        v26 := GenS2{F0: 255, F1: 186, F2: 666}
        _ = v26
    case "foo bar":
        v3 = v3
    }
    fmt.Println(v1, v3)
    var v27 int64 = (v2 - 0)
    v28 := []int{9223372036854775807}
    fmt.Println(v1, v27, v3)
    if (3.14159 > 1.5) {
        fmt.Println(v2, v27)
        v3 -= 65535
        if ("a" > ((("" + ("hello" + ("\n\t" + "a"))) + "a") + "hello")) {
            v29 := GenS3{F0: v1, F1: 431, F2: 42}
            fmt.Println(v27, v27, v3)
            _ = v29
        } else {
            var v30 uint16 = 349
            v31 := GenS3{F0: (v1 & (v1 >> ((65535) & 63))), F1: (v2 << ((1) & 63)), F2: -2147483648}
            v27 = (-(v27))
            _ = v30
            _ = v31
        }
    } else {
        v3 = v3
        v32 := GenS3{F0: v1, F1: (v2 + 950), F2: 2147483647}
        _ = v32
    }
    if false {
        if len(v28) > 0 {
            v28[0] = 32767
        }
        v1 = v1
    } else {
        for fi33 := 0; fi33 < 17; fi33++ {
            var v34 uint16 = 0
            _ = fi33
            _ = v34
        }
        v35 := GenS3{F0: v1, F1: 9223372036854775807, F2: 0}
        switch "a" {
        case "\n\t":
            var v36 string = ("hello" + "a")
            _ = v36
        case "":
            v37 := GenS4{F0: 2, F1: 9223372036854775807}
            _ = v37
        case "a":
            if len(v28) > 0 {
                v28[0] = -78
            }
        default:
            v38 := GenS2{F0: 0, F1: 804, F2: 0}
            _ = v38
        }
        _ = v35
    }
    v27 -= (v27 & v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
