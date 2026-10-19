package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }

func main() {
    var v1 uint16 = 1
    var v2 int16 = 42
    var v3 uint16 = (v1 << ((386) & 63))
    v1 = 1
    var v4 uint16 = v1
    var v5 uint32 = 1
    if true {
        switch int(v5) {
        case 127:
            for fi6 := 0; fi6 < 5; fi6++ {
                v1 += (v4 ^ (v4 << ((1) & 63)))
                _ = fi6
            }
        }
    } else {
        v3 += (v3 << ((4294967295) & 63))
    }
    switch "a" {
    case "a":
        if (1 >= 42) {
            var v7 float32 = 3.14159
            switch 357 {
            case 32768:
                var v8 float64 = 0.1
                _ = v8
            }
            _ = v7
        } else {
            v4 = 631
        }
    case "\n\t":
        v9 := GenS1{F0: -1, F1: -1}
        _ = v9
    }
    var v10 uint = 4294967295
    switch -988 {
    case 32767:
        switch -1 {
        case -965:
            var v11 uint32 = (v5 * (v5 % (((v5 | (v5 + uint32(v1)))) | 1)))
            _ = v11
        }
    }
    v10 += v10
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v2, v10, v1)
        v13 := []int{(fi12 / (((fi12 - fi12)) | 1)), fi12, (fi12 - (fi12 & (fi12 & (fi12 % (((fi12 >> ((v10) & 63))) | 1)))))}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint64 = 289
        var v16 uint16 = 255
        v17 := []int{fi14, int(v2)}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    switch "foo bar" {
    case "a":
        if true {
            var v18 uint8 = 117
            if false {
                var v19 uint = 2
                _ = v19
            }
            if true {
                for fi20 := 0; fi20 < 19; fi20++ {
                    if true {
                        fmt.Println(v10)
                        if (0 >= -1) {
                            for fi21 := 0; fi21 < 2; fi21++ {
                                v22 := []int{int(v2), -286}
                                var v23 uint64 = 1
                                switch fi21 {
                                case 32768:
                                    var v24 int32 = 0
                                    _ = v24
                                default:
                                    fmt.Println(v5, v10)
                                }
                                if true {
                                    v25 := GenS1{F0: 927, F1: 98}
                                    _ = v25
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    v4 = v1
                                case "":
                                    fmt.Println(fi20)
                                }
                                fmt.Println(v5, v18, v4)
                                v23 = 4294967295
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            v26 := GenS2{F0: 637}
                            _ = v26
                        }
                    }
                    _ = fi20
                }
                v27 := GenS3{F0: true, F1: (v10 == 18446744073709551615), F2: (^(v18))}
                var v28 int8 = 65
                _ = v27
                _ = v28
            } else {
                v4 -= v4
                v4 = (v1 + v4)
            }
            _ = v18
        }
    case "":
        switch -9223372036854775808 {
        case 127:
            var v29 int64 = -794
            _ = v29
        case 32767:
            fmt.Println(v3, v4, v5)
        case 1:
            if false {
                v30 := []int{-720, -481, 42}
                switch ("foo bar" + "") {
                case "hello":
                    _ = len(v30)
                default:
                    var v31 uint = v10
                    _ = v31
                }
                _ = v30
            } else {
                fmt.Println(v2, v1, v1)
                var v32 string = ("foo bar" + "\n\t")
                _ = v32
            }
        default:
            if false {
                switch "\n\t" {
                case "a":
                    v10 -= v10
                }
            } else {
                switch 32767 {
                case 2147483647:
                    var v33 int = 4294967295
                    _ = v33
                case -2147483648:
                    if false {
                        var v34 uint64 = 590
                        v35 := []int{}
                        _ = v34
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 2; fi36++ {
                            v37 := GenS2{F0: -1}
                            _ = fi36
                            _ = v37
                        }
                        var v38 string = "hello"
                        v38 = (v38 + "hello")
                        _ = v38
                    }
                default:
                    var v39 uint = (v10 | 926)
                    _ = v39
                }
                fmt.Println(v2, v10)
            }
        }
    case "hello":
        v40 := GenS2{F0: -9223372036854775808}
        _ = v40
    default:
        v2 += (v2 & -226)
    }
    v41 := []int{0, 32767, 32767, -776}
    v42 := GenS3{F0: (-9223372036854775808 < 762), F1: true, F2: 204}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
