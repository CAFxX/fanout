package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int32; F1 int8; F2 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 uint16, p2 int) uint8 {
    var v4 uint16 = (p1 | 1)
    p2 += 4294967295
    switch (p2 | p0) {
    case 32767:
        if (-128 >= -36) {
            if true {
                switch -1 {
                case -906:
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        if false {
                            v5 := []int{p0, (-(p0)), (^(p2))}
                            if (true && ((^(v4)) < 255)) {
                                var v6 int8 = 127
                                v7 := GenS1{F0: true, F1: 65535}
                                var v8 int16 = 1000
                                _ = v6
                                _ = v7
                                _ = v8
                            }
                            v9 := GenS1{F0: true, F1: 4294967295}
                            _ = v5
                            _ = v9
                        } else {
                            if true {
                                p2 -= p0
                            } else {
                                var v10 float64 = 1e10
                                _ = v10
                            }
                            p0 -= (p2 - -2147483648)
                        }
                    default:
                        v11 := []int{p0}
                        _ = v11
                    }
                default:
                    v12 := GenS1{F0: true, F1: 2147483648}
                    _ = v12
                }
                var v13 bool = false
                for fi14 := 0; fi14 < 5; fi14++ {
                    p0 += (^(p2))
                    v15 := []int{0, p0}
                    switch (p0 / ((p0) | 1)) {
                    case 9223372036854775807:
                        v13 = ((!v13) || (48 < 1))
                    case 0:
                        p0 = 9223372036854775807
                    case -9223372036854775808:
                        fmt.Println(v4, v4, p1)
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v13
            } else {
                if false {
                    var v16 uint64 = 2
                    var v17 uint8 = uint8(v16)
                    v18 := GenS2{F0: 2147483647, F1: 1, F2: 0}
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    if false {
                        if (!false) {
                            fmt.Println(v4, p2)
                        }
                    } else {
                        v19 := []int{p0, (p2 ^ 128), p0, p0}
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint32 = 2147483648
                    _ = fi20
                    _ = v21
                }
            }
            var v22 bool = true
            _ = v22
        }
    case 286:
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS1{F0: false, F1: 2}
            switch "foo bar" {
            case "foo bar":
                var v25 int = (^(fi23))
                _ = v25
            case "\n\t":
                fmt.Println(p1, p2, p0)
            case "hello":
                v26 := []int{p0, (p2 ^ 2147483647), fi23, -327}
                _ = v26
            }
            var v27 int32 = -1
            _ = fi23
            _ = v24
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 11; fi28++ {
            if false {
                v29 := []int{}
                var v30 string = ((("a" + "a") + "foo bar") + ("" + "foo bar"))
                _ = v29
                _ = v30
            }
            v31 := []int{(p2 >> ((65535) & 63)), 4294967295, (p2 - (^(fi28))), 964}
            _ = fi28
            _ = v31
        }
    }
    var v32 uint32 = 2147483647
    switch p0 {
    case -872:
        if true {
            fmt.Println(p1, p1, p2)
            v33 := []int{(p0 ^ 413), p2, (^(p0))}
            var v34 uint = 82
            _ = v33
            _ = v34
        } else {
            var v35 float32 = -2.25
            v36 := GenS3{F0: int32(p1)}
            v37 := []int{}
            _ = v35
            _ = v36
            _ = v37
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    _ = v4
    return 42
}

func main() {
    var v1 int = 9223372036854775807
    var v2 bool = (0 <= 127)
    v3 := []int{(v1 * -640)}
    if len(v3) > 0 {
        v3[0] = 127
    }
    fmt.Println(v2, v2)
    fmt.Println(v2, v2)
    var v38 int64 = 1
    fmt.Println(v2, v38, v2)
    v39 := GenS2{F0: 250, F1: -128, F2: int32(v1)}
    var v40 int16 = 566
    v41 := GenS2{F0: -2147483648, F1: 42, F2: -573}
    var v42 int16 = (v40 / (((v40 + (v40 >> ((2) & 63)))) | 1))
    var v43 float32 = 1.5
    var v44 int64 = 42
    for fi45 := 0; fi45 < 18; fi45++ {
        if v2 {
            var v46 string = "foo bar"
            _ = v46
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
