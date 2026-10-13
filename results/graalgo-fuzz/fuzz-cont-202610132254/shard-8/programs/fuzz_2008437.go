package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint; F1 string }
func genFn0() uint16 {
    switch ("" + "\n\t") {
    case "\n\t":
        var v5 int32 = 1
        _ = v5
    case "foo bar":
        switch 4294967295 {
        case 2147483647:
            for fi6 := 0; fi6 < 16; fi6++ {
                v7 := GenS1{F0: false, F1: 42}
                var v8 bool = true
                _ = fi6
                _ = v7
                _ = v8
            }
        case 697:
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 uint = 65535
                switch "foo bar" {
                case "foo bar":
                    v10 -= (v10 + (v10 - (v10 | 65535)))
                case "a":
                    v10 += v10
                case "":
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v12 := []int{(fi11 % ((fi11) | 1))}
                        _ = fi11
                        _ = v12
                    }
                }
                _ = fi9
                _ = v10
            }
        default:
            if false {
                var v13 uint = 0
                for fi14 := 0; fi14 < 15; fi14++ {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v16 := GenS1{F0: (0 != 2), F1: 0}
                        var v17 bool = (!true)
                        if true {
                            v18 := GenS1{F0: v17, F1: 1}
                            v17 = v17
                            v13 += v13
                            _ = v18
                        } else {
                            v13 -= v13
                        }
                        fmt.Println(fi14)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                fmt.Println(v13)
                _ = v13
            }
        }
    }
    v19 := GenS2{F0: -9223372036854775808, F1: 2147483647, F2: 127}
    var v20 uint8 = 1
    v20 = 2
    _ = v19
    _ = v20
    return 1
}
func genFn1(p0 int32, p1 int64) uint16 {
    p0 += p0
    p0 += p0
    var v21 string = (("\n\t" + "") + ("\n\t" + ("" + ("hello" + (("a" + "") + "hello")))))
    v21 = (v21 + v21)
    _ = p0
    _ = p1
    _ = v21
    return genFn0()
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"hello", ("hello" + ("\n\t" + "\n\t"))}
    var v3 uint32 = 65535
    v4 := []uint32{(v3 + (v3 | v3)), v3, (v3 + v3), (v3 >> ((18446744073709551615) & 63))}
    v3 += 2147483647
    fmt.Println(v3)
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if false {
        v22 := GenS3{F0: 1, F1: ("" + "hello")}
        v1[v3] = ""
        _ = v22
    }
    var v23 int8 = -1
    if (2 > v3) {
        v24 := GenS3{F0: 918, F1: ""}
        v4 = append(v4, (v3 * (v3 * 551)))
        switch 1 {
        case -2147483648:
            fmt.Println(v3)
        case 42:
            v25 := GenS2{F0: -9223372036854775808, F1: 36, F2: 115}
            _ = v25
        case 127:
            fmt.Println(v23, v3)
        }
        _ = v24
    }
    v23 -= v23
    switch "foo bar" {
    case "\n\t":
        delete(v1, 2)
    case "foo bar":
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := GenS2{F0: -1, F1: 768, F2: 255}
            fmt.Println(v23)
            v3 += 2147483648
            _ = fi26
            _ = v27
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
