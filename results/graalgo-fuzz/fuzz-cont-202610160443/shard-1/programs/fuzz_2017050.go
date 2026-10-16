package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 bool }
type GenS3 struct { F0 string }
func genFn0() uint32 {
    v7 := GenS1{F0: 2147483647, F1: 610, F2: 254}
    var v8 int32 = 0
    switch -664 {
    case 32768:
        var v9 int8 = -1
        _ = v9
    case -204:
        fmt.Println(v8)
    }
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int16) int32 {
    v10 := []int{992, -1, -50, -9223372036854775808}
    switch "foo bar" {
    case "":
        p0 -= (p0 * 0)
    case "foo bar":
        _ = len(v10)
    case "a":
        if (!false) {
            p0 = p0
        } else {
            var v11 uint64 = 0
            v11 -= (v11 ^ (v11 ^ v11))
            _ = v11
        }
    }
    v12 := GenS2{F0: 552, F1: 1, F2: ((65535 == 255) && true)}
    var v13 float32 = 1e10
    _ = p0
    _ = v10
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 uint64 = 725
    var v2 int64 = 9223372036854775807
    var v3 int32 = -2147483648
    var v4 bool = false
    v5 := []string{"\n\t", ((("hello" + ("a" + "foo bar")) + "") + "hello"), "\n\t", "hello"}
    v6 := []int{2147483647, -760}
    v4 = v4
    if ((2 != 1) && ("\n\t" != "\n\t")) {
        v5 = append(v5, "foo bar")
        if false {
            v14 := GenS1{F0: (v3 | (v3 % ((915) | 1))), F1: 821, F2: (v1 / ((0) | 1))}
            if len(v6) > 0 {
                v6[0] = 32768
            }
            _ = v14
        }
        v1 = 4294967295
    } else {
        var v15 int16 = 42
        _ = v15
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (!false) {
        v4 = v4
        v6 = append(v6, 32768)
        for fi16 := 0; fi16 < 6; fi16++ {
            v5 = append(v5, (("\n\t" + "foo bar") + "foo bar"))
            _ = fi16
        }
    }
    v3 -= (v3 ^ genFn1(1000))
    var v17 float64 = (float64(0.1) / (0.1 - 0.0))
    switch 127 {
    case 837:
        for fi18 := 0; fi18 < 2; fi18++ {
            switch 2147483647 {
            case -1:
                switch 32768 {
                case -9223372036854775808:
                    var v19 uint = 18446744073709551615
                    _ = v19
                }
            case -662:
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 uint64 = v1
                    _ = fi20
                    _ = v21
                }
            case 1:
                var v22 int32 = (v3 ^ (v3 | genFn1(-1)))
                _ = v22
            }
            v23 := GenS1{F0: genFn1(0), F1: 65535, F2: v1}
            var v24 int16 = 32767
            _ = fi18
            _ = v23
            _ = v24
        }
    case 2147483647:
        var v25 float32 = (0.0 + 0.0)
        _ = v25
    default:
        v6 = append(v6, -1)
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint32 = genFn0()
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        if v4 {
            switch fi26 {
            case 921:
                v2 -= -1
            case -971:
                v28 := GenS3{F0: ("hello" + "hello")}
                _ = v28
            default:
                v29 := GenS1{F0: v3, F1: 42, F2: 42}
                _ = v29
            }
        }
        _ = fi26
        _ = v27
    }
    var v30 bool = (65535 < 665)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
