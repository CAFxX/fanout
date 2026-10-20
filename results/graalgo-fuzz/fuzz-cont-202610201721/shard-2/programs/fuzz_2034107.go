package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 int32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int }
func genFn0(p0 int) bool {
    v4 := GenS1{F0: 118, F1: 1}
    var v5 uint32 = 1
    if true {
        var v6 int8 = 0
        _ = v6
    } else {
        var v7 float32 = -2.25
        var v8 uint64 = 1
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int32 = 42
            v11 := []int{}
            switch (fi9 | 0) {
            case 0:
                switch ("\n\t" + "hello") {
                case "a":
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 float32 = v7
                        v13 += ((((v13 - ((-2.25 * v7) / 0.1)) - (((v13 + -2.25) - (v7 * (0.0 / v13))) / float32((0.1 * (1e10 / float64(float32(3.14159))))))) / v7) * 1.5)
                        switch "a" {
                        case "\n\t":
                            v5 = v5
                        default:
                            if ((("" + "a") + "foo bar") > "") {
                                fmt.Println(v7, v8, v10)
                                v5 -= (v5 | (v5 + 0))
                            }
                        }
                        _ = fi12
                        _ = v13
                    }
                case "":
                    v5 = 0
                case "\n\t":
                    v10 = v10
                }
            case -1:
                v14 := GenS2{F0: 65535, F1: 0, F2: v10}
                _ = v14
            case 682:
                fmt.Println(v5, fi9, fi9)
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 int32, p1 int) int8 {
    switch -628 {
    case -127:
        fmt.Println(p1, p0)
    case 0:
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS1{F0: 194, F1: 2}
            p1 += fi15
            var v17 float64 = (0.1 * 3.14159)
            _ = fi15
            _ = v16
            _ = v17
        }
    case 32768:
        v18 := GenS2{F0: 65535, F1: 65535, F2: 42}
        _ = v18
    default:
        for fi19 := 0; fi19 < 17; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                p0 = -2147483648
                p0 = -783
                v21 := GenS2{F0: 1, F1: 2, F2: p0}
                p1 += (p1 << ((626) & 63))
                _ = fi20
                _ = v21
            }
            var v22 int64 = -591
            p1 = (fi19 & (fi19 - fi19))
            _ = fi19
            _ = v22
        }
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint16 = 255
    var v2 uint16 = 1
    v3 := make(map[uint32]bool)
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS1{F0: 255, F1: 42}
        fmt.Println(v2, fi23, v1)
        v3[42] = (0 != -1)
        _ = fi23
        _ = v24
    }
    v2 -= v1
    v25 := v3[0]
    var v26 int32 = -2147483648
    v1 = (v1 >> ((0) & 63))
    fmt.Println(v25, v26, v25)
    var v27 uint16 = (v1 % ((v2) | 1))
    v28 := []int{875}
    var v29 uint64 = 4294967295
    delete(v3, 960)
    v29 -= 42
    for fi30 := 0; fi30 < 16; fi30++ {
        switch fi30 {
        case -9223372036854775808:
            if (v26 == v26) {
                fmt.Println(v25, v2, v29)
                var v31 string = "a"
                _ = v31
            } else {
                delete(v3, 718)
                var v32 int = fi30
                switch int(v2) {
                case 128:
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                case -422:
                    for fi33 := 0; fi33 < 11; fi33++ {
                        switch 316 {
                        case 9223372036854775807:
                            fmt.Println(v32, fi33, v26)
                        case 252:
                            v34 := GenS3{F0: 1000, F1: 18446744073709551615}
                            _ = v34
                        }
                        switch "hello" {
                        case "a":
                            v35 := GenS4{F0: (fi33 ^ v32)}
                            _ = v35
                        case "hello":
                            var v36 uint32 = 399
                            _ = v36
                        case "":
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                        default:
                            switch (((("\n\t" + "a") + "foo bar") + "") + ("a" + "")) {
                            case "a":
                                v37 := GenS2{F0: uint(v2), F1: v1, F2: 0}
                                _ = v37
                            }
                        }
                        _ = fi33
                    }
                }
                _ = v32
            }
        default:
            _ = len(v28)
        }
        switch (-(fi30)) {
        case 2147483647:
            if (true && true) {
                var v38 string = "a"
                _ = v38
            } else {
                var v39 uint = 79
                _ = v39
            }
        case 42:
            if (42 != genFn1(42, fi30)) {
                v29 = v29
                var v40 string = ("" + ("" + "a"))
                v25 = (42 < 0)
                _ = v40
            } else {
                v41 := GenS1{F0: 229, F1: 248}
                _ = v41
            }
        default:
            switch "\n\t" {
            case "":
                if (!(v25 || (false && ("a" != "\n\t")))) {
                    fmt.Println(v1)
                } else {
                    v29 = 624
                }
            default:
                v2 += 255
            }
        }
        var v42 uint16 = v1
        _ = fi30
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
