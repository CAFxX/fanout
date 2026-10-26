package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
func genFn0(p0 int) uint8 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS1{F0: -240, F1: -1}
            _ = fi6
            _ = v7
        }
        fmt.Println(p0, p0)
        _ = fi5
    }
    switch (p0 + -739) {
    case 2147483647:
        if (!(p0 < 82)) {
            p0 += p0
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS1{F0: 163, F1: -128}
                var v10 int64 = -1
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            v11 := []int{32767}
            for fi12 := 0; fi12 < 7; fi12++ {
                _ = len(v11)
                switch p0 {
                case 2147483647:
                    if len(v11) > 0 {
                        v11[0] = (^(fi12))
                    }
                case -356:
                    v13 := GenS2{F0: 2, F1: 2147483648}
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 17; fi14++ {
                        fmt.Println(fi14, fi14, fi14)
                        _ = fi14
                    }
                }
                var v15 bool = (!true)
                _ = fi12
                _ = v15
            }
            _ = v11
        }
    case -9223372036854775808:
        var v16 string = "a"
        _ = v16
    default:
        if true {
            v17 := GenS2{F0: 2147483648, F1: 0}
            v18 := GenS1{F0: 1000, F1: -10}
            var v19 string = ""
            _ = v17
            _ = v18
            _ = v19
        }
    }
    p0 -= p0
    if true {
        v20 := GenS3{F0: 1}
        _ = v20
    }
    _ = p0
    return 195
}
func genFn1(p0 bool, p1 int8) uint64 {
    var v21 uint64 = 206
    var v22 uint16 = 701
    switch (("foo bar" + ("\n\t" + ("hello" + "\n\t"))) + "foo bar") {
    case "foo bar":
        for fi23 := 0; fi23 < 12; fi23++ {
            switch (-(fi23)) {
            case -794:
                v24 := []int{fi23}
                _ = v24
            case 127:
                var v25 int = (fi23 + fi23)
                _ = v25
            case 911:
                for fi26 := 0; fi26 < 15; fi26++ {
                    if false {
                        v27 := []int{-9223372036854775808, (fi23 ^ (fi23 % ((fi23) | 1)))}
                        if len(v27) >= 2 {
                            v27 = v27[1:len(v27)]
                        }
                        var v28 uint16 = (v22 ^ v22)
                        _ = v27
                        _ = v28
                    }
                    _ = fi26
                }
            }
            v29 := GenS3{F0: (v21 ^ (v21 / (((v21 | (^(v21)))) | 1)))}
            _ = fi23
            _ = v29
        }
    case "":
        if (!p0) {
            v21 -= (v21 >> ((4294967295) & 63))
            v30 := GenS3{F0: 18446744073709551615}
            v31 := GenS3{F0: (v21 ^ (v21 ^ (^(v21))))}
            _ = v30
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 uint8 = 1
            v21 -= v21
            _ = fi32
            _ = v33
        }
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return v21
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]bool)
    var v3 uint8 = 128
    v4 := make(map[int]string)
    v34 := GenS1{F0: 1, F1: 42}
    v1 = (v1 * (v1 - v1))
    _ = len(v2)
    v3 -= (v3 / ((v3) | 1))
    switch "" {
    case "foo bar":
        fmt.Println(v1, v1)
    default:
        v35 := GenS4{F0: -9223372036854775808}
        _ = v35
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        var v37 float32 = 1e10
        _ = fi36
        _ = v37
    }
    v2[("foo bar" + "hello")] = ((0.1 == 1e10) || (!false))
    var v38 uint16 = (v1 ^ (v1 | v1))
    var v39 int64 = 0
    v38 = v38
    v40 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
