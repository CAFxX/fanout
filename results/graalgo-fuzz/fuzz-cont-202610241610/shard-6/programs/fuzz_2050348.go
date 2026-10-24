package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
func genFn0(p0 int32) int64 {
    if true {
        if true {
            v6 := []int{127}
            var v7 bool = false
            _ = v6
            _ = v7
        } else {
            p0 += p0
            p0 += (p0 | (p0 ^ p0))
        }
        v8 := []int{-9223372036854775808}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                v10 := []int{585, fi9, fi9, fi9}
                v11 := GenS1{F0: 2, F1: 32767, F2: (-(p0))}
                v12 := GenS2{F0: 42}
                _ = v10
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: 18446744073709551615, F1: 790, F2: (p0 & (p0 | (p0 << ((733) & 63))))}
            var v14 float64 = ((0.1 - ((3.14159 * ((-2.25 * 0.1) * 3.14159)) * -2.25)) * ((3.14159 / 1e10) * 0.0))
            _ = fi9
            _ = v13
            _ = v14
        }
        switch -378 {
        case 1:
            p0 = (p0 << ((255) & 63))
        case -1:
            var v15 int = 9223372036854775807
            _ = v15
        case 128:
            v16 := GenS2{F0: 626}
            _ = v16
        default:
            switch "a" {
            case "hello":
                switch 298 {
                case 2147483647:
                    var v17 uint = 623
                    _ = v17
                default:
                    v18 := []int{-2147483648}
                    _ = v18
                }
            }
        }
        var v19 uint8 = 42
        _ = v19
    }
    v20 := []int{-273, 9223372036854775807}
    _ = len(v20)
    _ = p0
    _ = v20
    return 782
}
func genFn1(p0 int16) uint64 {
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS3{F0: 2, F1: (p0 % ((p0) | 1))}
        p0 += (p0 - -747)
        p0 = (-(p0))
        _ = fi21
        _ = v22
    }
    switch 32768 {
    case 42:
        p0 -= (p0 - (p0 >> ((2) & 63)))
    }
    _ = p0
    return 42
}

func main() {
    var v1 uint16 = 255
    var v2 uint64 = 875
    var v3 bool = (-2147483648 == 32767)
    var v4 int32 = 244
    var v5 bool = ("" != "a")
    if (127 != uint8(v1)) {
        switch -107 {
        case 128:
            switch "a" {
            case "":
                v23 := []int{1, 1, 32768}
                _ = v23
            case "a":
                v24 := GenS1{F0: 0, F1: (^(v1)), F2: (v4 + (^(v4)))}
                _ = v24
            case "\n\t":
                v3 = (!(-1 != 9223372036854775807))
            }
        case 354:
            fmt.Println(v3, v2, v5)
        case 9223372036854775807:
            fmt.Println(v3, v1, v5)
        default:
            var v25 uint8 = 127
            _ = v25
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            v2 = v2
        case "a":
            v26 := []int{}
            _ = v26
        case "hello":
            v27 := []int{1}
            _ = v27
        }
    }
    v28 := []int{32767}
    if (("foo bar" + "") == "hello") {
        fmt.Println(v5, v4)
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 int16 = -32768
            v3 = v3
            _ = fi29
            _ = v30
        }
        v4 = v4
    }
    v1 = v1
    switch 128 {
    case -9223372036854775808:
        v2 = (^(v2))
    case 42:
        v1 += uint16(v2)
    }
    v2 = (v2 << ((65535) & 63))
    var v31 int8 = -1
    var v32 uint32 = 1
    switch (("a" + ("hello" + "a")) + ("" + "hello")) {
    case "hello":
        for fi33 := 0; fi33 < 13; fi33++ {
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            if (!true) {
                if len(v28) > 0 {
                    v28[0] = fi33
                }
            } else {
                var v34 int64 = 270
                _ = v34
            }
            for fi35 := 0; fi35 < 10; fi35++ {
                v4 = v4
                _ = fi35
            }
            _ = fi33
        }
    case "foo bar":
        v28 = append(v28, 128)
    default:
        for fi36 := 0; fi36 < 17; fi36++ {
            fmt.Println(v4, v2, v32)
            if (761 == 0) {
                if v5 {
                    _ = len(v28)
                } else {
                    var v37 bool = (true || (!(v31 > (^(v31)))))
                    fmt.Println(fi36, v37)
                    v38 := GenS4{F0: 42}
                    _ = v37
                    _ = v38
                }
                if (v31 == (v31 + (v31 ^ v31))) {
                    var v39 int64 = int64(v32)
                    var v40 bool = false
                    _ = v39
                    _ = v40
                }
                var v41 bool = (!false)
                _ = v41
            }
            v28 = append(v28, fi36)
            _ = fi36
        }
    }
    if len(v28) > 0 {
        v28[0] = 9223372036854775807
    }
    if (v3 && (v3 || (128 <= 2))) {
        if v3 {
            v28 = append(v28, 9223372036854775807)
        }
        v42 := GenS1{F0: v2, F1: v1, F2: v4}
        _ = v42
    }
    v43 := GenS3{F0: genFn1(32767), F1: 1}
    v44 := GenS5{F0: (v4 > -267)}
    switch 127 {
    case -152:
        if false {
            var v45 uint = 822
            switch -2147483648 {
            case -9223372036854775808:
                fmt.Println(v4, v32)
            case 0:
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
            default:
                v28 = append(v28, -9223372036854775808)
            }
            _ = v45
        }
    case -9223372036854775808:
        fmt.Println(v1, v32)
    case 128:
        var v46 uint32 = 285
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
