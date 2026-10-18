package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool, p1 int16) string {
    if p0 {
        if (!(0 != -2147483648)) {
            v4 := []int{-682, -1, int(p1), 32768}
            _ = v4
        }
        switch "a" {
        case "\n\t":
            if false {
                var v5 uint64 = 0
                p0 = ("" > ("a" + "a"))
                p1 = (p1 ^ (p1 & (p1 | p1)))
                _ = v5
            } else {
                var v6 float32 = (float32((((0.0 / 1e10) + (-2.25 - 1.5)) - 1e10)) / 1.0)
                v6 = v6
                if p0 {
                    if p0 {
                        p1 -= (p1 + -32768)
                        v7 := []int{9223372036854775807, -734, -9223372036854775808, -1}
                        v6 += float32(3.14159)
                        _ = v7
                    } else {
                        if true {
                            v6 += 1e10
                            if true {
                                switch -1 {
                                case 127:
                                    v6 += 1e10
                                case 1:
                                    v6 = (1.5 - ((v6 * 3.14159) - 1.5))
                                case 32768:
                                    v6 = (v6 / ((v6 + v6) - (float32(1.5) * ((float32(1e10) + (((v6 * 1.5) - ((v6 + v6) - float32(0.0))) * (3.14159 * float32(1.5)))) * (((1.5 - (0.0 + ((v6 * v6) - v6))) * (v6 + float32((0.0 - 0.1)))) - (((1e10 / v6) - 3.14159) / (v6 / (0.0 - (float32((1.5 + 0.0)) / 1.0)))))))))
                                }
                                if p0 {
                                    v8 := GenS1{F0: 0}
                                    var v9 string = "\n\t"
                                    _ = v8
                                    _ = v9
                                }
                            }
                            var v10 int32 = -2147483648
                            _ = v10
                        } else {
                            var v11 uint64 = 2
                            if p0 {
                                v12 := GenS1{F0: 549}
                                v13 := []int{9223372036854775807}
                                _ = v12
                                _ = v13
                            }
                            fmt.Println(v6, v11, v11)
                            _ = v11
                        }
                        if p0 {
                            p0 = (0 >= 545)
                        }
                    }
                    v14 := []int{-533, 127, -9223372036854775808, 42}
                    _ = v14
                }
                _ = v6
            }
        default:
            fmt.Println(p1)
        }
    } else {
        v15 := GenS1{F0: 120}
        if false {
            p0 = p0
            p1 -= p1
        }
        _ = v15
    }
    if ((0 >= 1) && (1 > 272)) {
        fmt.Println(p1, p1)
    } else {
        p1 += -398
    }
    var v16 int8 = 127
    _ = p0
    _ = p1
    _ = v16
    return ("foo bar" + "a")
}
func genFn1(p0 int8) uint32 {
    fmt.Println(p0)
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := GenS1{F0: 18446744073709551615}
            var v20 bool = false
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: 0}
        _ = fi17
        _ = v21
    }
    _ = p0
    return 1
}

func main() {
    v1 := []uint32{42}
    var v2 int32 = -792
    v3 := []uint32{841, 310, 2147483648, 65535}
    v22 := GenS2{F0: 0}
    v23 := GenS2{F0: 1}
    v2 = (v2 >> ((1) & 63))
    fmt.Println(v2)
    var v24 int32 = v2
    v25 := GenS1{F0: 65535}
    v24 -= v24
    switch ((("\n\t" + "foo bar") + ("hello" + "")) + ("a" + "a")) {
    case "hello":
        v3 = append(v3, 0)
    case "foo bar":
        var v26 int16 = 0
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
