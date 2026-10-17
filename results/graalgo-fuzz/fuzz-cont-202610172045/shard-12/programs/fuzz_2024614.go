package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int; F1 int32 }
func genFn0(p0 bool, p1 uint16) string {
    fmt.Println(p1, p0)
    if true {
        v4 := GenS1{F0: 254, F1: 42}
        for fi5 := 0; fi5 < 7; fi5++ {
            if false {
                var v6 uint8 = uint8(p1)
                _ = v6
            }
            p0 = (!(((1 >= 227) && (65535 < 2147483648)) && ((^(p1)) > p1)))
            _ = fi5
        }
        _ = v4
    } else {
        v7 := GenS1{F0: 1, F1: 127}
        _ = v7
    }
    fmt.Println(p0, p1)
    if p0 {
        p1 -= 42
        for fi8 := 0; fi8 < 15; fi8++ {
            if p0 {
                if (-204 >= 869) {
                    var v9 int16 = -755
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        if (4294967295 != 2147483648) {
                            var v11 uint8 = 159
                            fmt.Println(v11, p0)
                            p1 = (p1 >> ((121) & 63))
                            _ = v11
                        }
                        v12 := []int{fi8, 0}
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        v12 = append(v12, (fi8 * -9223372036854775808))
                        var v13 uint = 2
                        _ = fi10
                        _ = v12
                        _ = v13
                    }
                }
                p0 = false
            } else {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v15 := []int{(fi8 & fi14), -2147483648, (fi8 * 4294967295)}
                    if (!((124 >= (fi8 >> ((18446744073709551615) & 63))) && false)) {
                        if (1 < -1) {
                            var v16 uint8 = 127
                            v15 = append(v15, fi14)
                            var v17 uint8 = v16
                            _ = v16
                            _ = v17
                        } else {
                            v15 = append(v15, (fi14 ^ (-(fi14))))
                            fmt.Println(fi14, fi14, fi14)
                            switch (fi14 & fi14) {
                            case 128:
                                p1 = p1
                            case 32767:
                                p1 = p1
                            case 9223372036854775807:
                                p0 = p0
                            }
                        }
                    }
                    fmt.Println(p0)
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = fi14
                    _ = v15
                }
                if p0 {
                    if ((((1e10 - (-2.25 - 3.14159)) / 1.0) + -2.25) < 0.1) {
                        v18 := []int{(-(fi8))}
                        _ = v18
                    }
                } else {
                    v19 := []int{}
                    if false {
                        for fi20 := 0; fi20 < 14; fi20++ {
                            fmt.Println(fi20, fi20)
                            _ = fi20
                        }
                        p0 = p0
                        v21 := GenS1{F0: 1, F1: 247}
                        _ = v21
                    }
                    var v22 uint8 = 255
                    _ = v19
                    _ = v22
                }
                v23 := GenS2{F0: (fi8 % ((fi8) | 1)), F1: -1}
                _ = v23
            }
            v24 := []int{fi8, fi8, -640, (fi8 / (((fi8 | (-(fi8)))) | 1))}
            var v25 float64 = (0.0 - 0.1)
            _ = fi8
            _ = v24
            _ = v25
        }
    } else {
        var v26 int8 = 127
        _ = v26
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    var v1 int32 = 42
    v2 := []string{("a" + "\n\t")}
    var v3 string = (("a" + "foo bar") + ("\n\t" + ((("hello" + "") + (("a" + "foo bar") + ("a" + "\n\t"))) + "hello")))
    var v27 uint16 = 1
    v3 = v3
    if true {
        v1 -= int32(v27)
        v27 += (v27 - 0)
    } else {
        var v28 int64 = 9223372036854775807
        v1 += v1
        _ = v28
    }
    v29 := GenS2{F0: 1, F1: 439}
    if (891 > 833) {
        var v30 bool = (!true)
        _ = v30
    } else {
        var v31 uint32 = 2
        _ = v31
    }
    fmt.Println(v3, v1)
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 int32 = 1
        _ = fi32
        _ = v33
    }
    v34 := GenS2{F0: -740, F1: v1}
    var v35 string = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v34
    _ = v35
}
