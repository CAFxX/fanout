package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 string }
func genFn0(p0 int64, p1 int16, p2 int) int8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: p1, F1: (!(-128 > -19)), F2: 478}
        v6 := []int{-951, 4294967295, (fi4 >> ((1) & 63))}
        var v7 int = (-(fi4))
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        switch "\n\t" {
        case "hello":
            var v9 int64 = (p0 & 9223372036854775807)
            _ = v9
        default:
            v10 := []int{(^(fi8)), p2, (p2 >> ((0) & 63))}
            _ = v10
        }
        v11 := []int{}
        fmt.Println(fi8, fi8, fi8)
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]string)
    v3 := []string{(("a" + (("foo bar" + "hello") + ("" + "a"))) + ("hello" + "a")), "\n\t", "", "\n\t"}
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 string = "foo bar"
        fmt.Println(fi12)
        _ = fi12
        _ = v13
    }
    v1 = (v1 ^ genFn0(408, 0, 127))
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 float32 = 3.14159
        _ = fi14
        _ = v15
    }
    fmt.Println(v1)
    v1 = 100
    if true {
        v1 -= (^(v1))
        if (!true) {
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 uint = 4294967295
                _ = len(v3)
                if (!false) {
                    var v18 uint = (v17 + 2)
                    switch "foo bar" {
                    case "":
                        var v19 uint = 18446744073709551615
                        _ = v19
                    }
                    fmt.Println(fi16, fi16)
                    _ = v18
                } else {
                    if (false || (!false)) {
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            var v20 uint64 = 4294967295
                            _ = v20
                        case "hello":
                            if (true && (0 <= 0)) {
                                if (!true) {
                                    for fi21 := 0; fi21 < 16; fi21++ {
                                        var v22 float32 = float32(((-2.25 + 1.5) / 0.1))
                                        _ = fi21
                                        _ = v22
                                    }
                                    for fi23 := 0; fi23 < 9; fi23++ {
                                        v17 += v17
                                        if (127 != v1) {
                                            v1 = genFn0(0, -1, fi23)
                                        } else {
                                            fmt.Println(v17, v1)
                                            var v24 uint32 = 2147483647
                                            var v25 uint32 = v24
                                            _ = v24
                                            _ = v25
                                        }
                                        fmt.Println(fi23, fi16, v17)
                                        _ = fi23
                                    }
                                } else {
                                    for fi26 := 0; fi26 < 2; fi26++ {
                                        v27 := GenS1{F0: -1, F1: true, F2: 0}
                                        v17 -= 42
                                        if true {
                                            var v28 int32 = 1
                                            _ = v28
                                        }
                                        var v29 float64 = 1e10
                                        if false {
                                            v3 = append(v3, "\n\t")
                                            v1 = v1
                                        }
                                        _ = fi26
                                        _ = v27
                                        _ = v29
                                    }
                                    _ = len(v2)
                                }
                                var v30 bool = false
                                _ = v30
                            }
                        }
                        fmt.Println(v1, v1, v1)
                        var v31 uint = (v17 & 734)
                        _ = v31
                    }
                    fmt.Println(v17, fi16)
                }
                _ = fi16
                _ = v17
            }
        }
        for fi32 := 0; fi32 < 19; fi32++ {
            for fi33 := 0; fi33 < 16; fi33++ {
                v1 = (v1 << ((941) & 63))
                fmt.Println(fi32, fi32, fi33)
                switch (fi32 * 2147483647) {
                case 9223372036854775807:
                    v1 = v1
                }
                _ = fi33
            }
            switch "a" {
            case "":
                v34 := GenS2{F0: (!false), F1: ("hello" + "hello")}
                _ = v34
            }
            _ = fi32
        }
    }
    if false {
        fmt.Println(v1)
        fmt.Println(v1)
    }
    v1 -= (v1 & v1)
    v1 = 42
    v1 = genFn0(1, 1, -569)
    fmt.Println(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
}
