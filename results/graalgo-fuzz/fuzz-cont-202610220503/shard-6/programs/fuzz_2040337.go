package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 100
    var v3 int32 = -2147483648
    var v4 bool = true
    v3 = (v3 ^ 0)
    fmt.Println(v4, v3)
    v4 = v4
    v4 = v4
    switch -1 {
    case 0:
        v2 += (v2 >> ((4294967295) & 63))
    case 128:
        var v5 uint8 = 1
        _ = v5
    default:
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(v3, fi6)
            v1 = (v1 + (v1 << ((65535) & 63)))
            if true {
                v7 := []int{fi6, fi6, fi6, fi6}
                for fi8 := 0; fi8 < 11; fi8++ {
                    if len(v7) > 0 {
                        v7[0] = (fi8 >> ((uint(v2)) & 63))
                    }
                    var v9 int = (fi8 / ((fi8) | 1))
                    _ = fi8
                    _ = v9
                }
                _ = v7
            }
            _ = fi6
        }
    }
    var v10 int16 = 638
    for fi11 := 0; fi11 < 8; fi11++ {
        switch fi11 {
        case 819:
            var v12 float32 = (1e10 * 0.0)
            _ = v12
        case -1:
            switch (fi11 % (((fi11 * int(v3))) | 1)) {
            case 32768:
                v13 := []int{(fi11 - fi11), (fi11 >> ((1) & 63)), fi11, fi11}
                _ = v13
            case -508:
                switch 2147483647 {
                case 1:
                    if (0 >= 0) {
                        for fi14 := 0; fi14 < 18; fi14++ {
                            if (0 != 57) {
                                if (false && (!v4)) {
                                    var v15 int8 = 127
                                    var v16 uint16 = 1
                                    _ = v15
                                    _ = v16
                                }
                                var v17 string = ("" + "")
                                if ((3.14159 - 0.1) <= 3.14159) {
                                    var v18 string = "hello"
                                    _ = v18
                                }
                                _ = v17
                            } else {
                                v19 := GenS1{F0: (fi11 * fi11), F1: 0, F2: 190}
                                _ = v19
                            }
                            v20 := []int{127, (fi14 & (fi14 & fi11))}
                            switch ("" + ("" + "hello")) {
                            case "hello":
                                switch "\n\t" {
                                case "hello":
                                    v4 = (v4 || ((v3 + 1) > v3))
                                default:
                                    var v21 string = "a"
                                    _ = v21
                                }
                            case "\n\t":
                                v22 := GenS2{F0: 127}
                                _ = v22
                            default:
                                v23 := GenS1{F0: (^(fi11)), F1: 0, F2: 0}
                                _ = v23
                            }
                            var v24 uint = 234
                            _ = fi14
                            _ = v20
                            _ = v24
                        }
                        var v25 bool = (42 < 42)
                        v26 := GenS3{F0: (v1 ^ (v1 & v1)), F1: int64(v3)}
                        _ = v25
                        _ = v26
                    }
                case 492:
                    for fi27 := 0; fi27 < 2; fi27++ {
                        switch int(v3) {
                        case 42:
                            fmt.Println(fi11)
                        case -449:
                            v10 += v10
                        case 127:
                            switch "foo bar" {
                            case "a":
                                if (true || (0 < -308)) {
                                    v28 := GenS4{F0: (v2 - v2)}
                                    var v29 int = -294
                                    var v30 float64 = 1e10
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                            case "":
                                fmt.Println(fi27, v3)
                            case "foo bar":
                                v31 := []int{(fi27 << ((299) & 63)), (^(fi11)), (fi27 ^ fi27), (fi27 ^ (fi11 % ((fi27) | 1)))}
                                _ = v31
                            }
                        default:
                            v32 := []int{646, fi11, (fi27 & (fi27 & 9223372036854775807))}
                            _ = v32
                        }
                        var v33 int = fi27
                        if v4 {
                            v34 := GenS5{F0: (v2 >> ((9) & 63))}
                            _ = v34
                        }
                        v35 := GenS3{F0: (^(v1)), F1: 0}
                        _ = fi27
                        _ = v33
                        _ = v35
                    }
                case -1:
                    v36 := GenS6{F0: (^(v1)), F1: v4}
                    _ = v36
                }
            case 4294967295:
                var v37 uint8 = 94
                _ = v37
            }
        case 4294967295:
            v4 = (false || (859 == 0))
        default:
            if v4 {
                v38 := []int{0}
                _ = v38
            }
        }
        for fi39 := 0; fi39 < 8; fi39++ {
            fmt.Println(v10)
            _ = fi39
        }
        _ = fi11
    }
    v10 -= v10
    var v40 int8 = -122
    var v41 bool = v4
    if (!(0.1 >= 1.5)) {
        v4 = ((v2 <= 508) && ((128 > 255) && (29 < 42)))
        v3 -= (v3 * 693)
        v42 := GenS5{F0: (v2 | (v2 % (((v2 % (((v2 - v2)) | 1))) | 1)))}
        _ = v42
    } else {
        if (8 <= 1) {
            v1 += (v1 & (v1 - v1))
            var v43 uint64 = 18446744073709551615
            _ = v43
        }
        v44 := []int{42, 127, 279}
        if false {
            v3 = (^(v3))
            v44 = append(v44, 2147483647)
            if len(v44) >= 2 {
                v44 = v44[1:len(v44)]
            }
        }
        _ = v44
    }
    if v41 {
        switch "hello" {
        case "a":
            v45 := []int{}
            _ = v45
        }
    }
    var v46 uint = 65535
    for fi47 := 0; fi47 < 20; fi47++ {
        fmt.Println(v3, v40, v10)
        var v48 int8 = int8(v2)
        _ = fi47
        _ = v48
    }
    v49 := GenS7{F0: (v2 ^ 2), F1: uint8(v46)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v49
}
