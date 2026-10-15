package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 int32 }
func genFn0() int32 {
    var v4 uint = 65535
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 65535, F1: 1}
        v7 := GenS1{F0: 386, F1: 42}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: 65535, F1: 88}
            if (1 != 42) {
                if (true && (false && true)) {
                    var v11 uint64 = 4294967295
                    if false {
                        v11 += v11
                        v12 := GenS1{F0: 42, F1: 878}
                        _ = v12
                    }
                    v13 := []int{fi8}
                    _ = v11
                    _ = v13
                } else {
                    if ((false && (2 <= 128)) && ((128 >= 127) && (!(("hello" + "foo bar") <= "\n\t")))) {
                        v4 += v4
                    }
                    v4 += v4
                    var v14 string = ("hello" + "\n\t")
                    _ = v14
                }
                v15 := GenS1{F0: 2, F1: 265}
                fmt.Println(fi8, fi9)
                _ = v15
            } else {
                v4 -= 42
            }
            var v16 string = "a"
            _ = fi9
            _ = v10
            _ = v16
        }
        v4 -= 4
        _ = fi8
    }
    v17 := GenS1{F0: 42, F1: 409}
    if ((!(!true)) && (!(270 > 42))) {
        fmt.Println(v4)
    } else {
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 int = (fi18 / ((fi18) | 1))
            var v20 int8 = 100
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: 832, F1: 0}
        var v22 int16 = 0
        _ = v21
        _ = v22
    }
    _ = v17
    _ = v4
    return -184
}
func genFn1(p0 uint32, p1 uint8) uint {
    switch ("\n\t" + "\n\t") {
    case "a":
        fmt.Println(p1)
    case "hello":
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS1{F0: 255, F1: 638}
            switch "" {
            case "foo bar":
                var v25 string = (("foo bar" + "a") + "foo bar")
                _ = v25
            default:
                var v26 int = fi23
                _ = v26
            }
            _ = fi23
            _ = v24
        }
    }
    var v27 int32 = -522
    v28 := GenS1{F0: 1, F1: 359}
    var v29 uint = 42
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    _ = v29
    return (v29 * (v29 & v29))
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v1 = (-34 != -9223372036854775808)
    for fi30 := 0; fi30 < 9; fi30++ {
        v1 = true
        _ = fi30
    }
    v31 := GenS2{F0: 2147483648, F1: 0}
    switch "a" {
    case "a":
        var v32 uint32 = 0
        _ = v32
    default:
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 float64 = -2.25
            switch ("hello" + "foo bar") {
            case "a":
                var v35 int16 = 1000
                _ = v35
            case "foo bar":
                _ = len(v2)
            default:
                var v36 float64 = v34
                _ = v36
            }
            v37 := []int{(fi33 - fi33), 32767, 703, fi33}
            _ = fi33
            _ = v34
            _ = v37
        }
    }
    switch "hello" {
    case "":
        if (3.14159 > (3.14159 * (3.14159 * 1e10))) {
            var v38 uint32 = 897
            var v39 string = ""
            for fi40 := 0; fi40 < 12; fi40++ {
                for fi41 := 0; fi41 < 11; fi41++ {
                    switch 128 {
                    case -9223372036854775808:
                        switch 4294967295 {
                        case 32767:
                            var v42 uint = 860
                            _ = v42
                        case 0:
                            v1 = v1
                        case 4294967295:
                            v1 = v1
                        default:
                            fmt.Println(v39)
                        }
                    }
                    var v43 string = (v39 + v39)
                    if (-2147483648 < -1) {
                        fmt.Println(v43, v1, fi41)
                    }
                    v44 := GenS3{F0: 4294967295, F1: v39}
                    v45 := GenS4{F0: 42}
                    _ = fi41
                    _ = v43
                    _ = v44
                    _ = v45
                }
                var v46 bool = (42 < 2)
                _ = fi40
                _ = v46
            }
            _ = v38
            _ = v39
        } else {
            v1 = false
        }
    default:
        v47 := GenS5{F0: 0, F1: 0}
        _ = v47
    }
    if true {
        delete(v3, -1)
        v1 = true
        switch -695 {
        case 1:
            var v48 int16 = 32767
            _ = v48
        case -1:
            fmt.Println(v1)
        }
    } else {
        var v49 float32 = (((0.0 * ((((1.5 - 0.0) / 1e10) / 1.5) * 1.5)) / (0.1 / -2.25)) * ((0.1 + (((1.5 / -2.25) + (1e10 + 3.14159)) - float32(0.1))) - (1.5 + 1e10)))
        if true {
            v2[0] = 128
            var v50 string = ""
            _ = v50
        }
        _ = v49
    }
    var v51 int = -788
    v52 := GenS6{F0: 262}
    v51 = v51
    v53 := GenS7{F0: -1, F1: -394}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v51
    _ = v52
    _ = v53
}
