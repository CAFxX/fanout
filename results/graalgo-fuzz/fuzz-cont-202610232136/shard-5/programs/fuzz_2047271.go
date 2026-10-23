package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int }
func genFn0() bool {
    var v5 uint = 948
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(fi6, fi6)
        if (!(false || false)) {
            var v7 uint = (v5 & (^(v5)))
            if true {
                switch (fi6 + fi6) {
                case 512:
                    if (32767 == -32768) {
                        v8 := []int{fi6}
                        var v9 uint8 = 255
                        _ = v8
                        _ = v9
                    } else {
                        v5 -= (^(v7))
                        var v10 int16 = -840
                        _ = v10
                    }
                case 114:
                    var v11 int8 = 1
                    _ = v11
                default:
                    var v12 int = 162
                    _ = v12
                }
            } else {
                v13 := []int{9223372036854775807, fi6}
                _ = v13
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                fmt.Println(fi6, v7)
                v7 = 42
                _ = fi14
            }
            _ = v7
        } else {
            v15 := GenS1{F0: v5}
            v5 = (v5 - (v5 & (v5 << ((88) & 63))))
            _ = v15
        }
        _ = fi6
    }
    switch "foo bar" {
    case "a":
        var v16 bool = true
        _ = v16
    case "\n\t":
        var v17 int64 = 504
        _ = v17
    case "hello":
        v18 := GenS2{F0: uint8(v5)}
        _ = v18
    default:
        v5 = v5
    }
    v19 := []int{118, -764, -835, 32768}
    _ = v19
    _ = v5
    return ((131 >= 1) || true)
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -1
    v3 := []int{754, 32768, 820, 4294967295}
    var v4 uint8 = 42
    var v20 int32 = -1
    v21 := GenS1{F0: 1}
    if false {
        v22 := GenS1{F0: 42}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 19; fi23++ {
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 int64 = 0
                v26 := GenS2{F0: 42}
                v27 := v1[fi24]
                v2 = v2
                v4 += v4
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            switch ("hello" + "") {
            case "hello":
                fmt.Println(v4)
            case "a":
                if len(v3) > 0 {
                    v3[0] = (^(fi23))
                }
            case "\n\t":
                _ = len(v1)
            default:
                v28 := GenS2{F0: (^(v4))}
                _ = v28
            }
            delete(v1, fi23)
            _ = fi23
        }
        fmt.Println(v20, v4, v2)
        v20 += (^(v20))
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 uint32 = 2147483648
        var v31 int8 = 19
        _ = fi29
        _ = v30
        _ = v31
    }
    v4 = (v4 / ((v4) | 1))
    for fi32 := 0; fi32 < 15; fi32++ {
        for fi33 := 0; fi33 < 19; fi33++ {
            v20 += -130
            _ = fi33
        }
        v2 = v2
        v2 = int16(v4)
        _ = fi32
    }
    v34 := v1[848]
    v35 := GenS2{F0: 77}
    if false {
        if true {
            if false {
                v4 += v4
            } else {
                for fi36 := 0; fi36 < 14; fi36++ {
                    if ((-(v20)) >= (v20 - (-(v20)))) {
                        v4 = 242
                        fmt.Println(v20)
                        var v37 int64 = -9223372036854775808
                        _ = v37
                    }
                    switch v34 {
                    case "\n\t":
                        var v38 float32 = 1e10
                        _ = v38
                    case "":
                        v39 := GenS3{F0: v20, F1: 42}
                        _ = v39
                    }
                    _ = fi36
                }
            }
        }
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 uint64 = 512
            _ = fi40
            _ = v41
        }
        v34 = (v34 + v34)
    }
    var v42 int32 = v20
    for fi43 := 0; fi43 < 1; fi43++ {
        v2 -= int16(v4)
        v4 += (v4 * v4)
        v44 := GenS2{F0: (^(v4))}
        _ = fi43
        _ = v44
    }
    if ((v2 << ((2) & 63)) != v2) {
        for fi45 := 0; fi45 < 12; fi45++ {
            if false {
                for fi46 := 0; fi46 < 6; fi46++ {
                    var v47 string = "\n\t"
                    _ = fi46
                    _ = v47
                }
                if true {
                    var v48 int16 = (v2 - 1000)
                    _ = v48
                }
            } else {
                fmt.Println(v2)
                v2 = 1000
                delete(v1, fi45)
            }
            var v49 bool = (901 == 65535)
            _ = fi45
            _ = v49
        }
        fmt.Println(v42)
        v50 := GenS4{F0: 4294967295, F1: -6}
        _ = v50
    }
    var v51 int64 = 140
    var v52 uint32 = 42
    v53 := GenS5{F0: 127}
    switch 32767 {
    case 191:
        v34 = "hello"
    case -2147483648:
        fmt.Println(v2, v20, v2)
    }
    var v54 bool = (42 >= (v42 + (v42 / (((v42 >> ((255) & 63))) | 1))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
