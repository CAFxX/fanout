package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int16; F2 uint }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 int16 }
func genFn0() int {
    var v4 int32 = -376
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            fmt.Println(v4)
            _ = fi6
        }
        var v7 string = "a"
        _ = fi5
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 32767}
    if (1.5 < (float32((0.0 - (((float64(1e10) + 3.14159) - 1.5) * ((1e10 - (0.0 - 1.5)) - ((1e10 + (3.14159 * (0.1 + float64(float32((((3.14159 - 0.0) - 1e10) + 1e10)))))) - 1.5))))) * 3.14159)) {
        if (0.0 == (-2.25 * (float32((-2.25 + -2.25)) * ((-2.25 - (-2.25 * (0.0 / 1e10))) - 0.0)))) {
            v9 := GenS2{F0: 255}
            _ = v9
        } else {
            v10 := []int{}
            v11 := GenS3{F0: 935}
            v12 := GenS4{F0: 32767, F1: 24, F2: 973}
            _ = v10
            _ = v11
            _ = v12
        }
    }
    if (2147483647 >= -9223372036854775808) {
        v4 = v4
    } else {
        v4 += v4
        if (false && false) {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v4)
            case "a":
                fmt.Println(v4)
            }
            v13 := []int{2147483647, 1, 128, -2147483648}
            var v14 uint = 2
            _ = v13
            _ = v14
        } else {
            v4 = v4
        }
    }
    _ = v4
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 0
    var v3 int64 = 9223372036854775807
    var v15 uint16 = 52
    v16 := GenS4{F0: 434, F1: -32768, F2: v1}
    v17 := GenS5{F0: 0, F1: (3.14159 == 1.5)}
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 uint = uint(v15)
        v2 -= 0
        _ = fi18
        _ = v19
    }
    v15 += v15
    if (100 <= 24) {
        v3 += 42
        for fi20 := 0; fi20 < 5; fi20++ {
            v2 = v2
            v21 := GenS1{F0: 661, F1: (fi20 * (fi20 * (-(fi20))))}
            var v22 uint8 = v2
            _ = fi20
            _ = v21
            _ = v22
        }
        if true {
            v3 -= (^(v3))
            if true {
                if (true && false) {
                    var v23 int8 = -128
                    _ = v23
                }
            }
        } else {
            var v24 uint32 = 42
            _ = v24
        }
    } else {
        if (32768 != 28) {
            var v25 int64 = (v3 >> ((v1) & 63))
            v26 := GenS6{F0: 32767}
            v25 = (^(v25))
            _ = v25
            _ = v26
        } else {
            if false {
                v3 = (^(v3))
                v27 := []int{}
                for fi28 := 0; fi28 < 2; fi28++ {
                    fmt.Println(fi28, v1, v2)
                    var v29 int32 = 2147483647
                    var v30 uint = v1
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v27
            } else {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v3 = 42
                    if (4294967295 < 255) {
                        var v32 uint16 = 2
                        v2 -= v2
                        _ = v32
                    } else {
                        var v33 uint32 = uint32(v2)
                        for fi34 := 0; fi34 < 12; fi34++ {
                            fmt.Println(fi34)
                            _ = fi34
                        }
                        switch genFn0() {
                        case 9223372036854775807:
                            for fi35 := 0; fi35 < 14; fi35++ {
                                v36 := []int{fi35, fi35, (fi31 << (((v1 + (v1 | uint(v33)))) & 63))}
                                if (true || false) {
                                    v1 -= 957
                                    var v37 int = (fi31 * (fi31 << ((0) & 63)))
                                    var v38 float32 = -2.25
                                    _ = v37
                                    _ = v38
                                }
                                _ = fi35
                                _ = v36
                            }
                        }
                        _ = v33
                    }
                    _ = fi31
                }
                switch (("a" + "foo bar") + "a") {
                case "hello":
                    var v39 float32 = 0.0
                    _ = v39
                case "\n\t":
                    v15 -= (^(v15))
                }
            }
            v2 = (v2 ^ v2)
        }
        v15 = 1
    }
    v1 += (v1 | uint(v2))
    var v40 int64 = -1
    v41 := GenS7{F0: false, F1: -218}
    v40 = (v3 & v3)
    v1 = 1
    var v42 float64 = (1.5 * 1e10)
    var v43 uint64 = 855
    fmt.Println(v3, v3)
    if (-1 < -1) {
        v2 += (^(v2))
        switch 32767 {
        case 32768:
            v44 := []int{}
            _ = v44
        }
        switch "\n\t" {
        case "hello":
            fmt.Println(v43, v2)
        }
    } else {
        v40 += v40
        var v45 string = "foo bar"
        v46 := []int{32767, 32768, genFn0()}
        _ = v45
        _ = v46
    }
    var v47 uint8 = v2
    v48 := []int{4294967295, -51, 128, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v47
    _ = v48
}
