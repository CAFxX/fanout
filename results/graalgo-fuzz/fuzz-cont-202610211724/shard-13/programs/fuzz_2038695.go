package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0(p0 uint32, p1 uint64) bool {
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{fi8}
        p0 = 65535
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return false
}
func genFn1(p0 int16) uint {
    var v11 int16 = p0
    var v12 float64 = 1e10
    v13 := []int{}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 1
    v3 := make(map[int]string)
    var v4 uint8 = 0
    v5 := []int{9223372036854775807, 617}
    var v6 uint8 = (v4 * (v4 % ((v4) | 1)))
    fmt.Println(v4)
    if (1 == -1) {
        _ = len(v5)
        fmt.Println(v6)
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        if (1000 == -1) {
            v15 := GenS1{F0: (fi14 >> ((18446744073709551615) & 63)), F1: 1}
            var v16 uint = 1
            switch (-(fi14)) {
            case 132:
                switch "a" {
                case "foo bar":
                    for fi17 := 0; fi17 < 14; fi17++ {
                        delete(v3, -731)
                        v18 := GenS1{F0: (fi17 / (((fi14 / ((fi14) | 1))) | 1)), F1: -9223372036854775808}
                        _ = fi17
                        _ = v18
                    }
                case "":
                    fmt.Println(v16)
                default:
                    v4 -= (v6 / ((205) | 1))
                }
            }
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    var v19 bool = (-11 == 0)
    if v19 {
        fmt.Println(v19)
    } else {
        v19 = (-1 < 2147483647)
        v4 = 127
        if v19 {
            var v20 uint16 = 73
            v21 := GenS2{F0: 0, F1: v20}
            v6 = v4
            _ = v20
            _ = v21
        } else {
            var v22 uint32 = 4294967295
            _ = v22
        }
    }
    v19 = ((v2 != uint64(v6)) && (v19 || (!(!(0.1 == -2.25)))))
    if (0.1 > -2.25) {
        for fi23 := 0; fi23 < 1; fi23++ {
            if true {
                fmt.Println(v19, v4, v4)
            }
            switch (("a" + "") + "") {
            case "a":
                switch (^(fi23)) {
                case 217:
                    for fi24 := 0; fi24 < 7; fi24++ {
                        if false {
                            var v25 int64 = 42
                            if (fi23 != (-(fi24))) {
                                v2 += (^(v2))
                            }
                            _ = v25
                        }
                        v26 := GenS2{F0: (fi23 * fi23), F1: 65535}
                        v6 = v4
                        _ = fi24
                        _ = v26
                    }
                case 128:
                    var v27 int64 = 9223372036854775807
                    _ = v27
                case -269:
                    if false {
                        delete(v1, 2147483648)
                        v28 := GenS2{F0: (fi23 ^ fi23), F1: 1}
                        v6 += (v4 / ((v4) | 1))
                        _ = v28
                    }
                }
            case "\n\t":
                v6 += (^(v6))
            case "hello":
                if (2147483647 < 0) {
                    var v29 uint64 = (^(v2))
                    var v30 float64 = 1e10
                    fmt.Println(fi23, v29, v2)
                    _ = v29
                    _ = v30
                } else {
                    v19 = ((0 < 4294967295) && v19)
                }
            default:
                if v19 {
                    if (32767 == 532) {
                        for fi31 := 0; fi31 < 11; fi31++ {
                            fmt.Println(v4, fi23)
                            v19 = true
                            var v32 string = ""
                            _ = fi31
                            _ = v32
                        }
                        v6 = v6
                    } else {
                        v33 := v1[405]
                        _ = v33
                    }
                    var v34 int64 = -94
                    _ = v34
                }
            }
            v19 = true
            _ = fi23
        }
        v4 = v6
    } else {
        for fi35 := 0; fi35 < 10; fi35++ {
            v2 = (v2 >> ((uint(v6)) & 63))
            if v19 {
                for fi36 := 0; fi36 < 3; fi36++ {
                    fmt.Println(v4, v2, fi35)
                    _ = fi36
                }
                var v37 uint = uint(v6)
                v19 = true
                _ = v37
            }
            var v38 uint32 = 0
            _ = fi35
            _ = v38
        }
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
