package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int8; F2 int16 }
func genFn0(p0 uint, p1 int32, p2 int16) int {
    switch ("" + ("a" + "hello")) {
    case "a":
        var v5 uint8 = 2
        _ = v5
    case "\n\t":
        var v6 int32 = 42
        _ = v6
    }
    v7 := []int{4294967295, 333}
    p1 = (p1 << ((p0) & 63))
    v8 := GenS1{F0: false, F1: 4294967295, F2: -1}
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 bool = true
    v2 := []uint32{2, 288}
    var v3 float32 = 1e10
    v4 := []string{"", "\n\t", ("" + "\n\t"), "hello"}
    var v9 uint64 = 18446744073709551615
    v3 = 1e10
    if (false && (true || true)) {
        if (255 <= 237) {
            fmt.Println(v1, v1)
        } else {
            v4 = append(v4, ("hello" + ("" + "hello")))
            for fi10 := 0; fi10 < 13; fi10++ {
                v1 = v1
                fmt.Println(v3, fi10)
                _ = fi10
            }
        }
        var v11 int16 = 1000
        _ = v11
    } else {
        if (v1 && v1) {
            var v12 string = "a"
            _ = len(v2)
            _ = v12
        } else {
            v3 = (v3 - v3)
        }
        var v13 int = 0
        fmt.Println(v13, v1)
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v9 = 1
        if true {
            var v15 int = fi14
            v16 := GenS2{F0: 0, F1: 534, F2: 127}
            _ = v15
            _ = v16
        } else {
            switch "hello" {
            case "a":
                switch (fi14 | (fi14 << ((2) & 63))) {
                case -9223372036854775808:
                    if v1 {
                        v17 := GenS3{F0: v9, F1: 255}
                        var v18 uint = 978
                        var v19 string = ("\n\t" + "a")
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        var v20 string = (("" + ((("hello" + ("" + (("" + "") + "\n\t"))) + "foo bar") + "hello")) + "")
                        for fi21 := 0; fi21 < 14; fi21++ {
                            v22 := GenS1{F0: (v1 || false), F1: 2147483648, F2: -128}
                            fmt.Println(v3, v3, v1)
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                default:
                    var v23 uint64 = (v9 / ((18446744073709551615) | 1))
                    _ = v23
                }
            }
        }
        v1 = false
        _ = fi14
    }
    var v24 int = 32767
    v25 := GenS3{F0: v9, F1: 2}
    fmt.Println(v3, v1, v24)
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 int = (fi26 * 0)
        switch ("" + ("hello" + "a")) {
        case "":
            if true {
                v28 := GenS4{F0: -2147483648, F1: v9}
                for fi29 := 0; fi29 < 4; fi29++ {
                    v30 := GenS3{F0: v9, F1: 42}
                    v31 := GenS5{F0: genFn0(232, 1, 1)}
                    if (v1 && ((v9 ^ 301) == 42)) {
                        if (v3 == ((1.5 / 1e10) - v3)) {
                            v3 = 0.1
                            if (genFn0(453, -942, -832) >= genFn0(619, -2147483648, -1)) {
                                v32 := GenS1{F0: v1, F1: 65535, F2: 0}
                                if (v1 || v1) {
                                    var v33 uint = 338
                                    v27 = -94
                                    _ = len(v2)
                                    _ = v33
                                } else {
                                    _ = len(v4)
                                }
                                _ = v32
                            }
                        }
                        v1 = (-526 > 1)
                        v34 := GenS3{F0: (v9 / (((v9 << ((18446744073709551615) & 63))) | 1)), F1: 42}
                        _ = v34
                    } else {
                        if v1 {
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                            } else {
                                v3 += v3
                            }
                            if (!v1) {
                                var v35 float32 = v3
                                _ = len(v4)
                                switch -2147483648 {
                                case 217:
                                    v2 = append(v2, 0)
                                case -1:
                                    var v36 float64 = 1.5
                                    _ = v36
                                default:
                                    v37 := GenS3{F0: (v9 / ((42) | 1)), F1: 0}
                                    _ = v37
                                }
                                _ = v35
                            } else {
                                fmt.Println(v1)
                            }
                        }
                        v27 = 32767
                        if (!(v1 && v1)) {
                            var v38 int16 = 0
                            _ = len(v4)
                            v39 := GenS6{F0: 127, F1: 42, F2: v38}
                            _ = v38
                            _ = v39
                        } else {
                            var v40 bool = v1
                            _ = v40
                        }
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = v28
            } else {
                var v41 uint16 = 42
                v9 -= v9
                _ = v41
            }
        case "\n\t":
            v24 = v27
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
