package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int8; F2 int8 }
func genFn0() uint32 {
    var v7 int = 0
    v8 := []int{v7, v7}
    var v9 int = 9223372036854775807
    switch -92 {
    case 0:
        v10 := GenS1{F0: v9, F1: true, F2: (^(v9))}
        _ = v10
    default:
        v9 -= (v7 * (v7 & 1))
    }
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]string)
    v3 := []int{-2147483648, 32767, 40, 2147483647}
    var v4 float64 = (0.0 * 0.0)
    var v5 int64 = 42
    v6 := []string{"\n\t", "\n\t", (("" + "\n\t") + ("foo bar" + ("hello" + ""))), "hello"}
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: 1, F1: false, F2: (^(fi11))}
        v1 -= genFn0()
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 128, F1: false, F2: 1}
    for fi14 := 0; fi14 < 6; fi14++ {
        if (!true) {
            v15 := GenS2{F0: true}
            _ = v15
        } else {
            if true {
                v6 = append(v6, ("\n\t" + ("foo bar" + "\n\t")))
            } else {
                v1 += (v1 % (((v1 >> ((121) & 63))) | 1))
            }
            v6 = append(v6, "foo bar")
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint64 = 42
            if len(v6) > 0 {
                v6[0] = ""
            }
            _ = fi16
            _ = v17
        }
        if true {
            fmt.Println(v1)
            if (740 >= 18446744073709551615) {
                var v18 uint16 = 1
                _ = v18
            }
        } else {
            v4 = (0.0 - (1.5 * float64(1e10)))
        }
        _ = fi14
    }
    v5 = (v5 << ((2) & 63))
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 string = ("a" + "")
        _ = fi19
        _ = v20
    }
    if false {
        v4 = 3.14159
        fmt.Println(v4, v5)
        if (255 <= 255) {
            if len(v6) > 0 {
                v6[0] = "\n\t"
            }
            if (0 != 159) {
                v21 := v2["a"]
                for fi22 := 0; fi22 < 4; fi22++ {
                    delete(v2, v21)
                    if true {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            var v24 int32 = 649
                            v4 -= v4
                            _ = fi23
                            _ = v24
                        }
                        v25 := GenS1{F0: fi22, F1: true, F2: fi22}
                        fmt.Println(v5, v21)
                        _ = v25
                    }
                    var v26 float64 = v4
                    _ = fi22
                    _ = v26
                }
                _ = v21
            }
        }
    }
    v2["hello"] = "\n\t"
    fmt.Println(v1, v4, v1)
    v27 := v2["foo bar"]
    if false {
        v28 := GenS3{F0: v27, F1: 86, F2: 127}
        if (-398 >= -9223372036854775808) {
            v27 = (v27 + "")
            v29 := v2[(("a" + "a") + "foo bar")]
            v29 = v29
            _ = v29
        }
        _ = v28
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v30 := GenS2{F0: (((0.1 - -2.25) / (-2.25 / float32(v4))) == (3.14159 + 0.0))}
    switch (v27 + "") {
    case "\n\t":
        for fi31 := 0; fi31 < 19; fi31++ {
            v32 := GenS2{F0: false}
            v5 = (v5 % ((0) | 1))
            _ = fi31
            _ = v32
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        v4 = v4
        v4 = (v4 * 3.14159)
        if (1.5 >= (1e10 / float64(1.5))) {
            delete(v2, v27)
            var v34 bool = (false && true)
            _ = v34
        }
        _ = fi33
    }
    var v35 int32 = -28
    if (!true) {
        for fi36 := 0; fi36 < 7; fi36++ {
            if true {
                var v37 int64 = -994
                _ = v37
            }
            var v38 uint64 = 0
            var v39 int = 696
            _ = fi36
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            var v41 uint8 = 128
            _ = fi40
            _ = v41
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            for fi42 := 0; fi42 < 3; fi42++ {
                v1 = (v1 + v1)
                v1 = v1
                _ = fi42
            }
        case "hello":
            _ = len(v3)
        }
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        v1 = (v1 - v1)
        var v44 bool = true
        if len(v6) > 0 {
            v6[0] = (v27 + "")
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
