package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }
type GenS4 struct { F0 int32; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int16; F1 int; F2 int16 }
type GenS8 struct { F0 int32; F1 int32 }
func genFn0() int8 {
    var v7 int = -189
    v7 = 32768
    v8 := []int{v7, 1, v7}
    _ = v7
    _ = v8
    return -128
}
func genFn1() string {
    var v9 string = (("a" + "hello") + ("hello" + ("foo bar" + "\n\t")))
    v9 = ("hello" + v9)
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint64 = 4294967295
        _ = fi10
        _ = v11
    }
    if true {
        var v12 uint8 = 128
        var v13 int8 = genFn0()
        var v14 int8 = (v13 / ((v13) | 1))
        _ = v12
        _ = v13
        _ = v14
    } else {
        switch 127 {
        case -690:
            var v15 float32 = (3.14159 * float32(3.14159))
            _ = v15
        }
    }
    if (!true) {
        var v16 int64 = 9223372036854775807
        switch v9 {
        case "a":
            for fi17 := 0; fi17 < 18; fi17++ {
                if true {
                    fmt.Println(fi17)
                    v18 := GenS1{F0: v9}
                    v19 := GenS2{F0: 641}
                    _ = v18
                    _ = v19
                }
                switch (fi17 ^ fi17) {
                case 32767:
                    var v20 string = v9
                    _ = v20
                case -1:
                    if (!((fi17 % ((fi17) | 1)) >= 215)) {
                        var v21 int8 = 42
                        var v22 uint16 = 0
                        var v23 string = "hello"
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        var v24 uint64 = 0
                        v16 = (v16 >> ((864) & 63))
                        v16 = (v16 >> ((18446744073709551615) & 63))
                        _ = v24
                    }
                case 42:
                    var v25 uint = 2
                    _ = v25
                }
                _ = fi17
            }
        case "":
            v16 -= -1
        }
        _ = v16
    } else {
        fmt.Println(v9)
    }
    _ = v9
    return "hello"
}

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 uint16 = 1
    v5 := []uint32{4294967295, 2147483647, 2147483648}
    v6 := make(map[string]bool)
    fmt.Println(v1, v4)
    v4 = (v4 - 923)
    var v26 int16 = -458
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 float64 = 0.1
        v29 := GenS3{F0: (-(v26)), F1: 208, F2: int32(v26)}
        var v30 int64 = 1
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    if ((v1 && (!v1)) && (v26 < (v26 % (((v26 | v26)) | 1)))) {
        for fi31 := 0; fi31 < 9; fi31++ {
            switch "" {
            case "\n\t":
                var v32 int16 = (v26 >> ((18446744073709551615) & 63))
                _ = v32
            default:
                var v33 int32 = -2147483648
                _ = v33
            }
            _ = fi31
        }
        v1 = (172 <= 128)
    }
    fmt.Println(v4, v26, v1)
    v34 := v6[("foo bar" + "hello")]
    if true {
        if false {
            v35 := GenS4{F0: 517, F1: -699, F2: 2147483648}
            _ = len(v3)
            _ = v35
        }
        var v36 int16 = v26
        fmt.Println(v34, v4)
        _ = v36
    } else {
        if v1 {
            v37 := GenS5{F0: 313, F1: uint8(v4)}
            _ = v37
        }
        _ = len(v6)
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        for fi39 := 0; fi39 < 7; fi39++ {
            v40 := GenS6{F0: v26, F1: (((((-2.25 / 0.1) == (((1.5 - -2.25) / 0.1) * 1e10)) || v34) || v1) && true)}
            _ = len(v5)
            if (!v1) {
                fmt.Println(v26, v4)
                var v41 float64 = float64(0.1)
                if v1 {
                    var v42 int8 = -4
                    _ = v42
                } else {
                    var v43 uint32 = 4294967295
                    _ = v43
                }
                _ = v41
            } else {
                fmt.Println(fi39, v34, fi38)
                v4 = v4
                v44 := GenS7{F0: v26, F1: fi39, F2: v26}
                var v45 uint64 = 583
                _ = v44
                _ = v45
            }
            if len(v5) > 0 {
                v5[0] = 2
            }
            v46 := GenS3{F0: v26, F1: 42, F2: -188}
            v34 = v1
            v4 -= v4
            if (v34 || (451 <= -613)) {
                switch -2147483648 {
                case -2147483648:
                    _ = len(v6)
                case 127:
                    var v47 bool = false
                    _ = v47
                case 9223372036854775807:
                    if v34 {
                        v26 += (v26 - 0)
                    }
                }
            } else {
                fmt.Println(v1, v26, v26)
                var v48 int32 = 2147483647
                v49 := GenS1{F0: "a"}
                _ = v48
                _ = v49
            }
            _ = fi39
            _ = v40
            _ = v46
        }
        if v1 {
            var v50 uint32 = 361
            var v51 int32 = -12
            _ = v50
            _ = v51
        }
        _ = fi38
    }
    v26 -= -261
    var v52 int8 = -26
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1 = ((!(37 == 42)) && true)
    for fi53 := 0; fi53 < 19; fi53++ {
        v54 := GenS8{F0: -221, F1: -866}
        fmt.Println(v52)
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
