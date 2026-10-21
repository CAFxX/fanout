package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 int; F1 int64; F2 uint32 }
type GenS8 struct { F0 int16; F1 bool }
func genFn0(p0 uint16, p1 uint32, p2 uint32) int32 {
    v4 := GenS1{F0: 95, F1: (p2 << ((18446744073709551615) & 63)), F2: 128}
    v5 := GenS2{F0: 42}
    var v6 bool = false
    v7 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8, p1 int8, p2 int8) int64 {
    p0 += 128
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            var v9 uint16 = 0
            v10 := []int{}
            if false {
                v9 = 65535
                p1 += (p1 / (((p1 + p1)) | 1))
                if true {
                    v11 := GenS1{F0: 9223372036854775807, F1: 2147483648, F2: fi8}
                    _ = v11
                }
            }
            _ = v10
            _ = v9
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            if false {
                var v13 int = (^(fi12))
                if (!(!true)) {
                    p1 = (p1 / (((p1 << ((18446744073709551615) & 63))) | 1))
                } else {
                    var v14 int = (v13 * v13)
                    if true {
                        if true {
                            if false {
                                switch "a" {
                                case "foo bar":
                                    var v15 int64 = 0
                                    _ = v15
                                }
                            } else {
                                v16 := []int{}
                                v13 -= (fi8 / (((v13 ^ (v14 | 42))) | 1))
                                _ = v16
                            }
                        }
                    } else {
                        var v17 uint = 18446744073709551615
                        var v18 string = "a"
                        _ = v17
                        _ = v18
                    }
                    v19 := []int{(v13 % ((-2147483648) | 1)), fi8}
                    _ = v14
                    _ = v19
                }
                _ = v13
            } else {
                var v20 string = "a"
                p2 = (p2 ^ p1)
                var v21 float32 = (0.0 + (-2.25 - (((((float32(0.0) + -2.25) + 0.0) + 0.0) * float32(-2.25)) - 0.1)))
                _ = v20
                _ = v21
            }
            v22 := []int{(fi8 % ((4294967295) | 1)), (fi12 << ((18446744073709551615) & 63)), (fi8 | fi12)}
            fmt.Println(p2)
            _ = fi12
            _ = v22
        }
        p2 = p2
        _ = fi8
    }
    v23 := []int{-9223372036854775808, -940, 1, 32767}
    p1 += (p2 << ((0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 0
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    var v3 int8 = 111
    var v24 uint16 = 255
    if (!(65535 < (v24 << ((0) & 63)))) {
        v24 += v24
    }
    v25 := GenS1{F0: -9223372036854775808, F1: 515, F2: int(v3)}
    v26 := GenS3{F0: 2, F1: 282}
    v27 := GenS4{F0: 375}
    var v28 uint8 = 1
    if (0.1 != 1e10) {
        v29 := GenS3{F0: 162, F1: 255}
        for fi30 := 0; fi30 < 15; fi30++ {
            v31 := GenS3{F0: v28, F1: 2}
            var v32 string = "foo bar"
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    } else {
        v2[1] = 1
        for fi33 := 0; fi33 < 6; fi33++ {
            v28 = v28
            v3 -= v3
            _ = fi33
        }
        v3 = v3
    }
    v34 := GenS5{F0: ("hello" + "foo bar")}
    switch ("hello" + "hello") {
    case "":
        var v35 bool = (true && (1 < 847))
        _ = v35
    }
    var v36 uint = 2
    switch (("foo bar" + ("foo bar" + "")) + "\n\t") {
    case "hello":
        if true {
            var v37 uint16 = 729
            switch int(v24) {
            case 2147483647:
                v38 := GenS1{F0: -236, F1: 0, F2: 2147483647}
                _ = v38
            default:
                if (!(0.1 != 0.1)) {
                    var v39 int8 = v3
                    _ = v39
                }
            }
            var v40 int32 = -2147483648
            _ = v37
            _ = v40
        } else {
            v28 = v28
            var v41 int = 127
            for fi42 := 0; fi42 < 10; fi42++ {
                var v43 uint32 = 986
                var v44 float64 = 0.0
                if ((255 > (v28 ^ v28)) && (false && (v24 == (v24 >> ((v36) & 63))))) {
                    v43 += (v43 >> ((201) & 63))
                    if false {
                        fmt.Println(fi42)
                        v45 := v2[(v43 + v43)]
                        var v46 uint16 = (v24 - (v24 % (((v24 ^ (v24 & (^(v24))))) | 1)))
                        _ = v45
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 4; fi47++ {
                        v48 := GenS6{F0: (^(v43)), F1: 127}
                        var v49 int = fi47
                        v41 = v41
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                }
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = v41
        }
    case "foo bar":
        if (((float32(1.5) * (3.14159 - (float32((float64(0.0) / 1.0)) / 0.1))) * (1e10 * (3.14159 / -2.25))) >= 1e10) {
            switch "hello" {
            case "hello":
                v24 = (v24 | (^(v24)))
            case "\n\t":
                if false {
                    v1 = append(v1, 128)
                    fmt.Println(v3)
                }
            case "foo bar":
                switch ("\n\t" + "") {
                case "a":
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        var v50 int64 = genFn1((v28 & v28), (v3 & (v3 + v3)), 42)
                        _ = v50
                    }
                case "foo bar":
                    if false {
                        v36 = v36
                        v28 = 62
                    }
                case "":
                    v51 := GenS7{F0: 2147483647, F1: 0, F2: 49}
                    _ = v51
                }
            default:
                fmt.Println(v3, v3, v36)
            }
            _ = len(v2)
            v52 := GenS8{F0: 0, F1: true}
            _ = v52
        }
    case "":
        var v53 uint32 = 407
        _ = v53
    default:
        v54 := GenS7{F0: -804, F1: 733, F2: 267}
        _ = v54
    }
    switch 0 {
    case 32767:
        var v55 uint8 = v28
        _ = v55
    case 42:
        v36 -= v36
    default:
        v56 := GenS6{F0: 0, F1: v3}
        _ = v56
    }
    v36 = 18446744073709551615
    fmt.Println(v24)
    if false {
        v28 -= (v28 | (v28 + 255))
        var v57 int32 = -2147483648
        _ = v57
    } else {
        v28 += (v28 << ((v36) & 63))
        switch "a" {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "a":
            v36 = 604
        }
        fmt.Println(v24, v28, v24)
    }
    fmt.Println(v28, v36, v28)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v36
}
