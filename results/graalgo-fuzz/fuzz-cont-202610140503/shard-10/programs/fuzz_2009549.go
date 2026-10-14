package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int8) int32 {
    if ((true && (!(!(787 >= 65535)))) || (p0 < (p0 / (((p0 << ((271) & 63))) | 1)))) {
        if (-2147483648 > -1) {
            var v4 uint32 = 65535
            _ = v4
        }
        p0 = (p0 >> ((4294967295) & 63))
    } else {
        p0 -= p0
        switch "\n\t" {
        case "a":
            var v5 uint8 = 16
            _ = v5
        }
        p0 += 1
    }
    v6 := []int{128}
    var v7 float32 = 1e10
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int32) uint {
    p0 = p0
    var v8 bool = (p0 == (-(p0)))
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            if (v8 || v8) {
                p0 = p0
                p0 = p0
            } else {
                if (v8 || false) {
                    v11 := GenS1{F0: 18446744073709551615, F1: 255}
                    _ = v11
                }
                v8 = (1000 == 901)
            }
            var v12 string = ""
            v8 = ("a" != v12)
            _ = fi10
            _ = v12
        }
        _ = fi9
    }
    v13 := []int{-2147483648, 1, -901}
    v14 := GenS1{F0: 18446744073709551615, F1: 42}
    _ = p0
    _ = v13
    _ = v14
    _ = v8
    return 346
}

func main() {
    var v1 string = ""
    var v2 int16 = -201
    v3 := []uint32{0, 1}
    fmt.Println(v1, v2)
    v1 = v1
    if (42 != 2) {
        fmt.Println(v2, v2)
    } else {
        v2 = v2
        var v15 uint16 = 953
        if false {
            for fi16 := 0; fi16 < 16; fi16++ {
                v1 = (v1 + ((((v1 + v1) + ("a" + v1)) + ((v1 + ("hello" + v1)) + v1)) + "hello"))
                var v17 uint64 = 301
                if (true || false) {
                    var v18 bool = true
                    var v19 float64 = float64(-2.25)
                    _ = v18
                    _ = v19
                } else {
                    if (!true) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        for fi20 := 0; fi20 < 3; fi20++ {
                            switch (v1 + v1) {
                            case "\n\t":
                                v2 = (v2 & (v2 & v2))
                            default:
                                v21 := GenS1{F0: v17, F1: v15}
                                _ = v21
                            }
                            _ = fi20
                        }
                    }
                }
                _ = fi16
                _ = v17
            }
        }
        _ = v15
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v1 = v1
        v2 = v2
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 float64 = 1e10
            if true {
                var v25 int16 = (v2 * v2)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v26 int16 = int16(fi22)
                _ = v25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    var v27 string = (((v1 + "foo bar") + "hello") + (v1 + (((v1 + "a") + v1) + "hello")))
    v2 -= v2
    var v28 int64 = 1
    if false {
        v29 := GenS1{F0: 18446744073709551615, F1: 42}
        _ = v29
    } else {
        if true {
            _ = len(v3)
            v1 = v27
        }
        v30 := GenS1{F0: 18446744073709551615, F1: 1}
        for fi31 := 0; fi31 < 20; fi31++ {
            fmt.Println(v1)
            v32 := GenS1{F0: 2, F1: 927}
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        switch "\n\t" {
        case "a":
            _ = len(v3)
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 65535
            }
        }
        v1 = v1
        v1 = "hello"
        _ = fi33
    }
    fmt.Println(v28, v1, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
