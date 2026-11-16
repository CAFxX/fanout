package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 int64 }

func main() {
    var v1 uint16 = 562
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 int8 = -32
    if (1 != 42) {
        fmt.Println(v4, v1)
        var v5 uint = 18446744073709551615
        var v6 uint32 = 1
        _ = v5
        _ = v6
    } else {
        _ = len(v2)
        if ((true && false) || (!(-641 <= 686))) {
            var v7 string = (("\n\t" + "a") + ("hello" + "a"))
            if ((-2.25 / float32(1.5)) < 1.5) {
                var v8 int16 = 1
                _ = v8
            }
            _ = v7
        } else {
            var v9 int16 = 42
            _ = v9
        }
        fmt.Println(v1, v1)
    }
    v4 = v4
    v10 := GenS1{F0: v1, F1: 32, F2: 1}
    v11 := GenS1{F0: 65535, F1: 128, F2: uint(v1)}
    var v12 bool = (true || (4294967295 > 1))
    if true {
        fmt.Println(v1)
    } else {
        if v12 {
            var v13 float32 = 3.14159
            v13 = ((v13 * 0.0) / (v13 * v13))
            _ = v13
        } else {
            var v14 uint64 = 42
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 float32 = (0.1 + 0.1)
                var v17 uint32 = 2
                v14 -= v14
                _ = fi15
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                switch 32768 {
                case 4294967295:
                    v19 := GenS1{F0: v1, F1: (fi18 % (((fi18 | (fi18 % (((fi18 ^ fi18)) | 1)))) | 1)), F2: uint(v14)}
                    _ = v19
                case 2147483647:
                    switch (fi18 + (^(fi18))) {
                    case 127:
                        switch (fi18 % ((fi18) | 1)) {
                        case -57:
                            fmt.Println(v14, v14, v12)
                        case -1:
                            fmt.Println(v1, v14)
                        case 9223372036854775807:
                            var v20 uint8 = 128
                            _ = v20
                        default:
                            if v12 {
                                var v21 uint64 = (v14 & v14)
                                var v22 string = "foo bar"
                                _ = v21
                                _ = v22
                            } else {
                                var v23 uint64 = 590
                                _ = v23
                            }
                        }
                    case -1:
                        var v24 uint64 = (v14 ^ v14)
                        _ = v24
                    case 4294967295:
                        v1 = v1
                    default:
                        switch ((("a" + "\n\t") + "hello") + "") {
                        case "\n\t":
                            v25 := v2[fi18]
                            _ = v25
                        case "hello":
                            for fi26 := 0; fi26 < 12; fi26++ {
                                var v27 int64 = 42
                                _ = fi26
                                _ = v27
                            }
                        default:
                            if (42 >= (v4 ^ -1)) {
                                var v28 int8 = (v4 * 1)
                                var v29 uint16 = v1
                                _ = v28
                                _ = v29
                            } else {
                                if (!(3.14159 != (-2.25 - (float32(-2.25) / 1.5)))) {
                                    v30 := GenS2{F0: 2}
                                    v31 := GenS3{F0: int(v14)}
                                    for fi32 := 0; fi32 < 17; fi32++ {
                                        var v33 uint32 = 2147483648
                                        var v34 int8 = v4
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = v30
                                    _ = v31
                                } else {
                                    var v35 int8 = (v4 & v4)
                                    delete(v2, 127)
                                    switch (fi18 & fi18) {
                                    case -438:
                                        var v36 float64 = 1.5
                                        _ = v36
                                    default:
                                        v37 := []int{fi18, fi18, (^(fi18))}
                                        _ = v37
                                    }
                                    _ = v35
                                }
                            }
                        }
                    }
                default:
                    v38 := GenS4{F0: true, F1: 1}
                    _ = v38
                }
                for fi39 := 0; fi39 < 19; fi39++ {
                    switch 127 {
                    case -209:
                        var v40 int8 = 22
                        _ = v40
                    case 728:
                        var v41 int32 = -1
                        _ = v41
                    }
                    _ = fi39
                }
                v42 := []int{(fi18 >> ((18446744073709551615) & 63)), fi18, fi18}
                _ = fi18
                _ = v42
            }
            _ = v14
        }
        switch "a" {
        case "":
            switch -451 {
            case 9223372036854775807:
                fmt.Println(v1, v4, v4)
            case 0:
                v4 = (v4 - (-(v4)))
            default:
                switch 234 {
                case 331:
                    switch -454 {
                    case 4294967295:
                        var v43 float64 = 3.14159
                        _ = v43
                    case 1:
                        switch -887 {
                        case 42:
                            for fi44 := 0; fi44 < 9; fi44++ {
                                fmt.Println(v4, v1)
                                _ = fi44
                            }
                        case 128:
                            v45 := GenS5{F0: 0}
                            _ = v45
                        case -2147483648:
                            v12 = v12
                        }
                    }
                case 374:
                    v46 := GenS6{F0: v1}
                    _ = v46
                case 42:
                    fmt.Println(v4, v4)
                default:
                    var v47 int64 = -1
                    _ = v47
                }
            }
        case "\n\t":
            for fi48 := 0; fi48 < 10; fi48++ {
                v49 := GenS7{F0: 19, F1: 127}
                fmt.Println(v12, fi48)
                _ = fi48
                _ = v49
            }
        case "a":
            v1 = (v1 ^ 1)
        }
        v12 = (!true)
    }
    v12 = (!(v1 > v1))
    for fi50 := 0; fi50 < 20; fi50++ {
        for fi51 := 0; fi51 < 20; fi51++ {
            if (("foo bar" + ("\n\t" + "a")) != "foo bar") {
                var v52 string = ("\n\t" + "foo bar")
                v53 := GenS2{F0: 2}
                _ = v52
                _ = v53
            } else {
                var v54 int16 = 28
                _ = v54
            }
            v12 = (("hello" > "\n\t") && v12)
            v1 += v1
            _ = fi51
        }
        _ = fi50
    }
    for fi55 := 0; fi55 < 3; fi55++ {
        fmt.Println(fi55)
        _ = fi55
    }
    v56 := GenS8{F0: 2, F1: 127}
    if (uint32(v1) != 55) {
        var v57 uint8 = 255
        var v58 uint16 = v1
        _ = v57
        _ = v58
    } else {
        v1 = (v1 * 471)
    }
    if v12 {
        var v59 uint = 193
        var v60 int64 = 1
        var v61 float32 = (((3.14159 + (3.14159 - (3.14159 * float32(float64(float32(((3.14159 / 0.1) - 1e10))))))) - (1e10 - (0.0 * 3.14159))) + 3.14159)
        _ = v59
        _ = v60
        _ = v61
    }
    if v12 {
        v62 := GenS9{F0: -9223372036854775808}
        var v63 uint64 = 4294967295
        _ = v62
        _ = v63
    }
    v4 = (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v56
}
