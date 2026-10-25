package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }
func genFn0(p0 int32, p1 uint16) int64 {
    var v7 uint = 398
    v8 := []int{32767, 42, -836}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1() uint {
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 0, F1: 42, F2: 130}
        fmt.Println(fi9)
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 255
    if true {
        var v12 int8 = 0
        v13 := []int{1, 4294967295, 1, 42}
        _ = v12
        _ = v13
    } else {
        v14 := GenS2{F0: 42, F1: 1}
        var v15 float64 = (0.1 - 1.5)
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := []int{fi16}
        _ = fi16
        _ = v17
    }
    v11 = 11
    _ = v11
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 4294967295
    var v3 float32 = 1e10
    v4 := []uint32{}
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]string)
    var v18 uint8 = 172
    if (-943 > 1) {
        v1[("hello" + ("\n\t" + "foo bar"))] = (v2 ^ v2)
    }
    var v19 uint64 = 1
    v19 -= 42
    for fi20 := 0; fi20 < 9; fi20++ {
        delete(v1, "")
        if (-1 != 1) {
            if (true && true) {
                v19 = (^(v19))
                var v21 uint32 = (v2 << ((2) & 63))
                var v22 int16 = 1
                _ = v21
                _ = v22
            }
            if (int(v19) >= 4294967295) {
                v2 += (v2 | v2)
            } else {
                if len(v4) > 0 {
                    v4[0] = (v2 / ((439) | 1))
                }
                for fi23 := 0; fi23 < 1; fi23++ {
                    if (!(v18 == (v18 % ((v18) | 1)))) {
                        v24 := GenS2{F0: -1, F1: (v19 / ((18446744073709551615) | 1))}
                        switch "a" {
                        case "a":
                            v18 += 2
                        }
                        _ = v24
                    }
                    _ = fi23
                }
            }
        }
        _ = fi20
    }
    if (-95 == -128) {
        fmt.Println(v18)
        v25 := GenS1{F0: 1000, F1: 1, F2: 32767}
        switch 128 {
        case -198:
            fmt.Println(v19, v19)
        default:
            var v26 float64 = 0.0
            _ = v26
        }
        _ = v25
    } else {
        for fi27 := 0; fi27 < 14; fi27++ {
            var v28 uint = 0
            v3 = (3.14159 - (v3 * 0.1))
            _ = fi27
            _ = v28
        }
        v29 := GenS1{F0: 0, F1: -1, F2: 32767}
        _ = v29
    }
    var v30 int = int(v19)
    fmt.Println(v30, v19)
    for fi31 := 0; fi31 < 13; fi31++ {
        v4 = append(v4, (v2 - (v2 ^ (^(v2)))))
        var v32 float64 = 0.1
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 uint = 0
        v3 = v3
        _ = fi33
        _ = v34
    }
    var v35 uint64 = (v19 % ((v19) | 1))
    fmt.Println(v30, v35, v18)
    if (42 <= 0) {
        delete(v5, "foo bar")
    } else {
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 float64 = 0.1
            _ = len(v5)
            if false {
                v38 := GenS2{F0: -1, F1: v35}
                switch (fi36 ^ 4294967295) {
                case 1:
                    for fi39 := 0; fi39 < 2; fi39++ {
                        fmt.Println(v37, v3)
                        v40 := GenS3{F0: 365, F1: 0, F2: 42}
                        _ = fi39
                        _ = v40
                    }
                case 127:
                    v18 -= v18
                }
                _ = v38
            }
            _ = fi36
            _ = v37
        }
        var v41 float32 = (-2.25 - 3.14159)
        fmt.Println(v35, v19)
        _ = v41
    }
    v2 = (v2 & v2)
    if true {
        fmt.Println(v35, v19)
    } else {
        var v42 int16 = -1
        var v43 float32 = 3.14159
        delete(v5, "")
        _ = v42
        _ = v43
    }
    var v44 int = 32768
    v45 := GenS2{F0: 1, F1: 2}
    v46 := v5[("foo bar" + "")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
