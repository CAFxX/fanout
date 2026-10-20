package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int; F1 uint32; F2 int }

func main() {
    var v1 float64 = (float64((1e10 + (0.1 - ((float32(1e10) + -2.25) * (0.0 - 1.5))))) * -2.25)
    v2 := make(map[int]string)
    var v3 int64 = 42
    var v4 uint16 = 2
    fmt.Println(v3, v4)
    var v5 int64 = (v3 >> ((186) & 63))
    if false {
        v6 := GenS1{F0: 84}
        var v7 uint16 = 97
        var v8 float32 = float32(3.14159)
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 int64 = (v5 << ((18446744073709551615) & 63))
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int32 = -1
            v3 = -1
            fmt.Println(v4)
            _ = fi10
            _ = v11
        }
        if (42 >= 0) {
            v12 := v2[128]
            v9 = -9223372036854775808
            var v13 int32 = -1
            _ = v12
            _ = v13
        }
        _ = v9
    }
    v3 = (^(v3))
    if ((true && false) && (!true)) {
        v4 = (v4 << ((780) & 63))
        v14 := v2[0]
        _ = v14
    } else {
        v2[42] = "foo bar"
        v15 := []int{-177, -9223372036854775808, 1, 4294967295}
        switch "" {
        case "hello":
            var v16 int8 = -13
            _ = v16
        case "a":
            var v17 float64 = float64(1e10)
            _ = v17
        case "":
            var v18 float32 = 0.1
            _ = v18
        default:
            var v19 int64 = -570
            _ = v19
        }
        _ = v15
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v5 -= (v5 | int64(v4))
        var v21 uint32 = 65535
        for fi22 := 0; fi22 < 17; fi22++ {
            v3 = (-(v3))
            if (878 < v4) {
                v3 = v3
                v23 := GenS1{F0: -81}
                _ = v23
            } else {
                v1 = (1.5 / v1)
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    if (0 <= 2) {
        if true {
            for fi24 := 0; fi24 < 9; fi24++ {
                v1 = (1e10 - 1.5)
                _ = fi24
            }
            fmt.Println(v1, v4, v5)
        } else {
            for fi25 := 0; fi25 < 5; fi25++ {
                v4 -= v4
                v26 := GenS2{F0: 1000}
                v5 = -9223372036854775808
                _ = fi25
                _ = v26
            }
        }
        v2[128] = (("foo bar" + "") + "a")
    }
    if true {
        if true {
            for fi27 := 0; fi27 < 11; fi27++ {
                v4 = 84
                _ = fi27
            }
            v28 := []int{-1, 32767, 9223372036854775807}
            _ = v28
        } else {
            if true {
                v29 := GenS3{F0: (v3 | -1), F1: "a"}
                var v30 int = 128
                for fi31 := 0; fi31 < 14; fi31++ {
                    v32 := GenS4{F0: (v30 + (fi31 / ((4294967295) | 1))), F1: 0, F2: fi31}
                    var v33 float64 = v1
                    fmt.Println(fi31, fi31, v33)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = v29
                _ = v30
            }
            var v34 int64 = v3
            v35 := []int{}
            _ = v34
            _ = v35
        }
        if (true && true) {
            var v36 uint8 = 254
            if ("foo bar" > "hello") {
                var v37 float32 = float32(v1)
                _ = v37
            }
            fmt.Println(v4, v1)
            _ = v36
        }
    }
    _ = len(v2)
    v1 = (v1 - 0.0)
    switch 0 {
    case 780:
        var v38 int16 = -32768
        _ = v38
    case 127:
        v3 -= (v3 << ((289) & 63))
    case -215:
        v3 -= 116
    }
    var v39 bool = (!false)
    if (true && v39) {
        v40 := []int{}
        fmt.Println(v5, v1, v5)
        for fi41 := 0; fi41 < 12; fi41++ {
            var v42 int16 = 399
            fmt.Println(v42, v3)
            if (163 < 59) {
                v3 = v3
                fmt.Println(v1, fi41)
                for fi43 := 0; fi43 < 14; fi43++ {
                    var v44 uint32 = 65535
                    v45 := GenS1{F0: 73}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            }
            _ = fi41
            _ = v42
        }
        _ = v40
    } else {
        for fi46 := 0; fi46 < 10; fi46++ {
            v39 = (42 <= 2147483647)
            if (0 < 1000) {
                v47 := GenS4{F0: 32767, F1: 1, F2: 1}
                _ = v47
            } else {
                for fi48 := 0; fi48 < 2; fi48++ {
                    var v49 uint8 = 0
                    if true {
                        var v50 string = "a"
                        v51 := GenS3{F0: v3, F1: "foo bar"}
                        v52 := GenS2{F0: -1}
                        _ = v50
                        _ = v51
                        _ = v52
                    } else {
                        var v53 int32 = 0
                        _ = v53
                    }
                    v4 += v4
                    _ = fi48
                    _ = v49
                }
            }
            _ = fi46
        }
        v1 = (float64(float32(v1)) * v1)
        var v54 int16 = -1
        _ = v54
    }
    v55 := GenS4{F0: 763, F1: 2147483648, F2: 2147483647}
    var v56 bool = (v39 && (849 == 2))
    var v57 uint = 1
    v4 = (v4 + v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
}
