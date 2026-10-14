package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int16 }
func genFn0(p0 int, p1 uint32, p2 int8) uint16 {
    switch p0 {
    case 32767:
        for fi4 := 0; fi4 < 4; fi4++ {
            for fi5 := 0; fi5 < 15; fi5++ {
                p2 -= (p2 >> ((4294967295) & 63))
                var v6 uint16 = 497
                fmt.Println(v6, fi5, fi5)
                _ = fi5
                _ = v6
            }
            if true {
                v7 := GenS1{F0: -592, F1: 0}
                _ = v7
            }
            if (-278 <= 1) {
                if true {
                    for fi8 := 0; fi8 < 15; fi8++ {
                        var v9 int = 57
                        var v10 float32 = (1.5 * ((0.1 - 1e10) - -2.25))
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    v11 := GenS1{F0: 32767, F1: 0}
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 uint16 = 1
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
                v14 := []int{(p0 - (fi4 - 128)), fi4, (fi4 + fi4), (fi4 << ((861) & 63))}
                v15 := GenS2{F0: 9223372036854775807, F1: -140}
                _ = v14
                _ = v15
            }
            _ = fi4
        }
    case -1:
        p0 = (p0 - p0)
    }
    fmt.Println(p1, p0, p1)
    v16 := GenS3{F0: 1, F1: 18446744073709551615, F2: 255}
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := []int{fi18, (p0 * p0), fi17}
            p1 -= 4294967295
            if (-1 <= 1) {
                if false {
                    var v20 uint32 = p1
                    switch "foo bar" {
                    case "hello":
                        _ = len(v19)
                    case "a":
                        _ = len(v19)
                    case "":
                        v20 = 4294967295
                    default:
                        fmt.Println(fi18)
                    }
                    var v21 int = int(p2)
                    if (!(0.0 == (((1.5 - 3.14159) * ((((3.14159 + -2.25) + (0.0 - 1.5)) + 1.5) * 1e10)) * (3.14159 - (1.5 + 0.1))))) {
                        if len(v19) > 0 {
                            v19[0] = 0
                        }
                        if true {
                            fmt.Println(fi18)
                        }
                        v21 -= fi17
                    }
                    _ = v20
                    _ = v21
                }
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 32767
}
func genFn1() int {
    if (65535 == 708) {
        v22 := []int{839, 32767}
        if false {
            var v23 int64 = 658
            _ = v23
        }
        _ = v22
    }
    var v24 uint16 = 834
    v25 := GenS1{F0: -32768, F1: -1}
    _ = v24
    _ = v25
    return -148
}

func main() {
    var v1 uint = 4294967295
    var v2 uint = v1
    v3 := make(map[uint32]int)
    for fi26 := 0; fi26 < 13; fi26++ {
        v1 -= (^(v2))
        v1 += (v1 % (((v1 >> ((v1) & 63))) | 1))
        _ = fi26
    }
    v1 -= v1
    v2 = 1
    if true {
        v2 = (v1 * v1)
    } else {
        v27 := []int{}
        _ = v27
    }
    fmt.Println(v2, v2)
    var v28 int = 1
    if true {
        v28 -= v28
        var v29 int16 = -277
        v30 := GenS4{F0: 978, F1: (v29 & v29)}
        _ = v29
        _ = v30
    }
    v1 += (v2 ^ v1)
    v31 := []int{(-(v28)), -1}
    if (false && true) {
        var v32 bool = (900 == 32767)
        var v33 uint = v2
        v33 = (v1 / ((2) | 1))
        _ = v32
        _ = v33
    }
    v34 := GenS1{F0: 0, F1: int64(v2)}
    v2 = 1
    if (!(9223372036854775807 != 9223372036854775807)) {
        v2 = v2
        var v35 int8 = 100
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        v2 += (^(v1))
        if (false && (((453 <= 895) && true) && (42 > -1))) {
            v28 -= (fi36 >> ((v1) & 63))
            var v37 uint64 = 1
            v38 := GenS4{F0: 4294967295, F1: 42}
            _ = v37
            _ = v38
        }
        _ = fi36
    }
    switch (-(v28)) {
    case 907:
        var v39 int64 = -9223372036854775808
        _ = v39
    default:
        delete(v3, 970)
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 int = (^(v28))
        _ = fi40
        _ = v41
    }
    v28 -= (v28 - 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v34
}
