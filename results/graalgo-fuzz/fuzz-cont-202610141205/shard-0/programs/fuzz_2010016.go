package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int32; F2 int8 }
type GenS5 struct { F0 int; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 string = "hello"
    v2 := make(map[string]int)
    var v3 string = v1
    var v4 uint = 356
    v5 := make(map[string]bool)
    v6 := make(map[uint32]string)
    var v7 uint8 = 2
    var v8 bool = (42 > 42)
    var v9 int16 = -260
    v10 := GenS1{F0: 0, F1: 0, F2: 0}
    v11 := []int{}
    if ((0 == 42) || (2 >= 517)) {
        v3 = ("hello" + (v1 + (v3 + v1)))
        fmt.Println(v7, v9)
        v8 = (-32768 > (v9 ^ int16(v4)))
    } else {
        v9 = v9
        switch 133 {
        case 0:
            for fi12 := 0; fi12 < 9; fi12++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                var v13 int64 = 42
                v14 := GenS2{F0: -9223372036854775808, F1: v9}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    }
    if true {
        switch -2147483648 {
        case 995:
            var v15 uint8 = 2
            _ = v15
        case -1:
            fmt.Println(v4, v4, v4)
        case 32767:
            v4 = v4
        default:
            fmt.Println(v3)
        }
        v16 := GenS3{F0: -9223372036854775808}
        _ = v16
    } else {
        _ = len(v11)
        _ = len(v11)
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            switch fi18 {
            case -1:
                v7 = v7
            case 145:
                if v8 {
                    var v19 uint16 = 2
                    fmt.Println(fi17, v7, fi18)
                    var v20 string = v3
                    _ = v19
                    _ = v20
                } else {
                    if (!(0 != 0)) {
                        var v21 float32 = 3.14159
                        v9 -= v9
                        v22 := GenS3{F0: -516}
                        _ = v21
                        _ = v22
                    } else {
                        if (1 == -697) {
                            v8 = (1e10 != 1e10)
                            v7 += 127
                            var v23 float64 = 0.0
                            _ = v23
                        }
                    }
                }
            case 32768:
                if (42 < 18446744073709551615) {
                    fmt.Println(v1)
                    v8 = v8
                    v24 := GenS4{F0: 4294967295, F1: 0, F2: 1}
                    _ = v24
                }
            }
            if v8 {
                if (!(1 >= 42)) {
                    switch "a" {
                    case "hello":
                        switch 0 {
                        case 32768:
                            v4 = (^(v4))
                            switch "a" {
                            case "foo bar":
                                if (!(((0.0 * -2.25) / ((float64(1e10) - 0.0) / 1.5)) < 0.1)) {
                                    v7 -= 127
                                    v4 += (v4 % (((^(v4))) | 1))
                                }
                            case "hello":
                                v25 := GenS4{F0: 2, F1: -8, F2: 127}
                                _ = v25
                            default:
                                var v26 int32 = -2147483648
                                _ = v26
                            }
                        case 2147483647:
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        default:
                            v27 := GenS2{F0: fi17, F1: (v9 / ((914) | 1))}
                            _ = v27
                        }
                    case "":
                        fmt.Println(v4)
                    default:
                        var v28 uint32 = 865
                        _ = v28
                    }
                    switch -2147483648 {
                    case 466:
                        if len(v11) > 0 {
                            v11[0] = fi18
                        }
                    }
                    fmt.Println(fi18)
                } else {
                    var v29 int64 = -1
                    v4 = (v4 << ((v4) & 63))
                    v4 = (v4 / ((v4) | 1))
                    v29 -= (v29 << ((v4) & 63))
                    _ = v29
                }
                v7 = (v7 * v7)
            }
            _ = fi18
        }
        var v30 string = ((((v1 + ((v3 + (v1 + v3)) + "foo bar")) + v3) + (("a" + "\n\t") + ("a" + v1))) + (v3 + ("a" + v3)))
        v4 = (v4 ^ 811)
        _ = fi17
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 bool = (42 < 42)
        _ = fi31
        _ = v32
    }
    if (!v8) {
        if (v8 || true) {
            v33 := GenS5{F0: int(v4), F1: 0, F2: 65535}
            switch (v3 + v1) {
            case "a":
                v34 := GenS5{F0: 32767, F1: 371, F2: 32767}
                _ = v34
            case "":
                v9 += 0
            default:
                if v8 {
                    var v35 int32 = -1
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 9; fi36++ {
                        switch 0 {
                        case 1:
                            for fi37 := 0; fi37 < 6; fi37++ {
                                fmt.Println(v3, v7)
                                if (127 >= -123) {
                                    var v38 float64 = (float64(1e10) * float64(1e10))
                                    var v39 string = v3
                                    if (!(-2147483648 <= 2147483647)) {
                                        var v40 uint64 = 42
                                        switch ("\n\t" + "hello") {
                                        case "a":
                                            _ = len(v11)
                                        case "":
                                            var v41 int64 = -181
                                            _ = v41
                                        default:
                                            switch v3 {
                                            case "":
                                                fmt.Println(v8, v39, v1)
                                            default:
                                                v39 = v3
                                            }
                                        }
                                        _ = v40
                                    } else {
                                        v42 := GenS6{F0: 1, F1: 990}
                                        _ = v42
                                    }
                                    var v43 uint64 = 1
                                    _ = v38
                                    _ = v39
                                    _ = v43
                                } else {
                                    v44 := GenS7{F0: (v7 << ((65535) & 63))}
                                    var v45 int8 = 127
                                    _ = v44
                                    _ = v45
                                }
                                v46 := GenS8{F0: 0, F1: 65535}
                                _ = fi37
                                _ = v46
                            }
                        case 240:
                            for fi47 := 0; fi47 < 20; fi47++ {
                                v48 := v6[4294967295]
                                var v49 uint8 = (^(v7))
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                        default:
                            v6[42] = (v1 + (v3 + v3))
                        }
                        v11 = append(v11, (fi36 - -758))
                        v9 = (v9 % (((v9 * (^(v9)))) | 1))
                        _ = fi36
                    }
                    v3 = "hello"
                }
            }
            _ = v33
        } else {
            var v50 int16 = (v9 << (((v4 - (v4 ^ 0))) & 63))
            var v51 int64 = 0
            _ = len(v5)
            _ = v50
            _ = v51
        }
        fmt.Println(v4, v3)
        v3 = v1
    }
    v4 = v4
    switch 977 {
    case 42:
        var v52 int64 = 9223372036854775807
        _ = v52
    }
    for fi53 := 0; fi53 < 11; fi53++ {
        if true {
            _ = len(v11)
            fmt.Println(v7)
        }
        var v54 int32 = -2147483648
        v55 := GenS1{F0: (v54 * (^(v54))), F1: 55, F2: -128}
        _ = fi53
        _ = v54
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
