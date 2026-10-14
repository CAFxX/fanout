package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 bool }
func genFn0(p0 bool) uint64 {
    if p0 {
        switch 9223372036854775807 {
        case -699:
            p0 = (255 == 255)
        }
        for fi4 := 0; fi4 < 1; fi4++ {
            for fi5 := 0; fi5 < 12; fi5++ {
                fmt.Println(fi5)
                _ = fi5
            }
            v6 := GenS1{F0: -9223372036854775808}
            switch (fi4 & (-(fi4))) {
            case -259:
                var v7 float32 = 0.0
                _ = v7
            case -2147483648:
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := GenS1{F0: -9223372036854775808}
                    v10 := GenS2{F0: -1, F1: 1, F2: (fi4 - (fi4 ^ (fi4 % (((fi8 ^ (fi8 * (fi4 << ((697) & 63))))) | 1))))}
                    switch (fi8 - (fi8 ^ 32768)) {
                    case 128:
                        v11 := GenS3{F0: 277}
                        _ = v11
                    case 32767:
                        var v12 uint64 = 42
                        _ = v12
                    default:
                        switch "hello" {
                        case "a":
                            var v13 uint64 = 18446744073709551615
                            _ = v13
                        case "foo bar":
                            p0 = false
                        default:
                            p0 = (42 > 255)
                        }
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
            _ = fi4
            _ = v6
        }
    }
    var v14 uint = 255
    var v15 int32 = -2147483648
    _ = p0
    _ = v14
    _ = v15
    return 1
}
func genFn1(p0 uint, p1 int16) int16 {
    switch -554 {
    case 0:
        v16 := []int{-603}
        _ = v16
    case -1:
        if (true && (1.5 != ((0.1 + (1.5 * (1.5 - float64(1.5)))) - -2.25))) {
            v17 := []int{32767, 127, 4294967295, 4294967295}
            fmt.Println(p1)
            _ = v17
        }
    case -2147483648:
        v18 := []int{9223372036854775807, 636}
        _ = v18
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    var v3 int8 = -12
    v3 -= 42
    var v19 float32 = 1e10
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 uint32 = 0
            switch "" {
            case "a":
                fmt.Println(fi21, v3)
            case "foo bar":
                switch ("a" + ("a" + "\n\t")) {
                case "":
                    var v23 int16 = -862
                    _ = v23
                }
                var v24 uint64 = 4294967295
                switch "" {
                case "":
                    var v25 string = "a"
                    _ = v25
                default:
                    if (!(27 == 250)) {
                        v19 -= float32(-2.25)
                        v26 := []int{(fi21 | (fi20 * fi21)), (fi21 + (fi21 + fi21))}
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        _ = v26
                    } else {
                        var v27 uint64 = v24
                        fmt.Println(v3, v24)
                        var v28 float64 = 3.14159
                        _ = v27
                        _ = v28
                    }
                }
                fmt.Println(fi21, fi20, v24)
                _ = v24
            default:
                switch fi20 {
                case 42:
                    if ((42 > v3) || (!true)) {
                        if true {
                            var v29 int32 = -973
                            fmt.Println(v29, v1)
                            v30 := v2[v22]
                            _ = v29
                            _ = v30
                        }
                        switch "a" {
                        case "a":
                            if (v1 < (v1 - genFn1(18446744073709551615, v1))) {
                                var v31 int16 = genFn1(301, (v1 % (((v1 << ((35) & 63))) | 1)))
                                v32 := []int{(fi20 << ((2) & 63)), (fi21 >> ((42) & 63))}
                                fmt.Println(v31)
                                _ = v31
                                _ = v32
                            } else {
                                if ((!((false || ((false && false) || ((^(v3)) > (v3 * v3)))) || false)) && (!(!true))) {
                                    v19 = v19
                                    switch "a" {
                                    case "\n\t":
                                        var v33 uint = 18446744073709551615
                                        _ = v33
                                    }
                                    v34 := []int{(fi20 ^ 195), fi20}
                                    _ = v34
                                }
                                if (int32(v1) < 42) {
                                    v35 := []int{(fi21 ^ (fi21 << ((65535) & 63))), fi21, fi21, 726}
                                    v19 += (0.0 * (v19 * (1.5 * ((v19 * (v19 * v19)) + v19))))
                                    v1 -= (v1 << ((65535) & 63))
                                    v22 = v22
                                    _ = v35
                                }
                                v36 := []int{42}
                                _ = v36
                            }
                        default:
                            var v37 int = fi21
                            _ = v37
                        }
                        fmt.Println(v19, v22)
                        v3 = 0
                    }
                case -9223372036854775808:
                    v22 -= v22
                }
            }
            _ = fi21
            _ = v22
        }
        fmt.Println(v1, fi20, fi20)
        for fi38 := 0; fi38 < 19; fi38++ {
            switch 127 {
            case -9223372036854775808:
                var v39 uint64 = genFn0(false)
                _ = v39
            case 42:
                v1 += v1
                fmt.Println(v1)
            default:
                if false {
                    if true {
                        v19 -= v19
                        if false {
                            v2[2147483647] = 2147483647
                            fmt.Println(v1)
                        }
                        v19 -= 1e10
                    } else {
                        v19 += (1.5 / 1.5)
                    }
                    var v40 int32 = 1
                    v41 := GenS4{F0: 0, F1: 42, F2: (!true)}
                    _ = v40
                    _ = v41
                }
            }
            fmt.Println(v19, fi20, v1)
            _ = fi38
        }
        _ = fi20
    }
    var v42 int64 = 1
    v43 := []int{-919}
    if len(v43) >= 2 {
        v43 = v43[1:len(v43)]
    }
    v42 += v42
    delete(v2, 1)
    delete(v2, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
