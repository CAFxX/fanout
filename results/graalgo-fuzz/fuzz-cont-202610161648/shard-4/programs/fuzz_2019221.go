package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 uint }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint32, p1 uint64, p2 int) int8 {
    fmt.Println(p0)
    var v7 int16 = 32767
    var v8 uint8 = 151
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint64) uint8 {
    p0 += p0
    p0 = 2
    switch "hello" {
    case "":
        var v9 int = -2147483648
        _ = v9
    case "\n\t":
        var v10 int8 = genFn0(65535, 42, 42)
        _ = v10
    default:
        for fi11 := 0; fi11 < 13; fi11++ {
            if false {
                var v12 string = "a"
                fmt.Println(v12, p0)
                _ = v12
            } else {
                var v13 int64 = -1
                var v14 int64 = v13
                switch (("\n\t" + "hello") + "hello") {
                case "\n\t":
                    if false {
                        var v15 int8 = -128
                        _ = v15
                    } else {
                        var v16 uint32 = 0
                        var v17 uint64 = 18446744073709551615
                        v18 := GenS1{F0: 146, F1: false}
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                case "foo bar":
                    v13 -= v13
                }
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    }
    _ = p0
    return 0
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]bool)
    var v3 float64 = -2.25
    var v4 uint = 255
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]string)
    if (698 < 2147483648) {
        if (true || true) {
            switch -1 {
            case 127:
                switch v1 {
                case "foo bar":
                    if true {
                        v1 = v1
                    }
                case "\n\t":
                    v19 := GenS2{F0: 127, F1: 903}
                    _ = v19
                }
            default:
                v20 := GenS2{F0: -853, F1: 2147483648}
                _ = v20
            }
            v21 := GenS1{F0: -32768, F1: (0.1 > (float32(0.0) / float32(v3)))}
            _ = v21
        } else {
            var v22 bool = true
            _ = v22
        }
    }
    var v23 float64 = 0.0
    v23 = v3
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 float64 = ((0.1 / (0.0 - (0.1 * 3.14159))) + v3)
        _ = fi24
        _ = v25
    }
    switch v1 {
    case "hello":
        fmt.Println(v4, v4, v23)
    case "\n\t":
        if true {
            if true {
                var v26 uint64 = 2
                fmt.Println(v3)
                v4 = v4
                _ = v26
            } else {
                v27 := GenS2{F0: 1, F1: 1}
                var v28 bool = true
                if false {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if v28 {
                            fmt.Println(v4)
                        } else {
                            var v30 bool = true
                            _ = v30
                        }
                        _ = fi29
                    }
                } else {
                    var v31 bool = v28
                    if (false && true) {
                        fmt.Println(v31)
                        if false {
                            delete(v2, 712)
                            for fi32 := 0; fi32 < 19; fi32++ {
                                v33 := GenS1{F0: 32767, F1: true}
                                v23 = ((v23 / 1.0) - v3)
                                _ = fi32
                                _ = v33
                            }
                        }
                    } else {
                        v34 := []int{128}
                        _ = v34
                    }
                    switch 105 {
                    case 737:
                        var v35 uint16 = 255
                        _ = v35
                    default:
                        for fi36 := 0; fi36 < 13; fi36++ {
                            if (!v31) {
                                v37 := []int{fi36, (fi36 & fi36), (fi36 | 305), fi36}
                                _ = v37
                            } else {
                                v31 = (1e10 >= -2.25)
                                for fi38 := 0; fi38 < 4; fi38++ {
                                    v39 := GenS3{F0: "", F1: false, F2: (v4 ^ (v4 / (((v4 * v4)) | 1)))}
                                    fmt.Println(v23)
                                    _ = fi38
                                    _ = v39
                                }
                                for fi40 := 0; fi40 < 16; fi40++ {
                                    v41 := GenS1{F0: -32768, F1: (((1 >= 1) || false) && (0 <= 1))}
                                    _ = fi40
                                    _ = v41
                                }
                            }
                            fmt.Println(v31, fi36)
                            if true {
                                v3 -= float64((3.14159 * -2.25))
                            } else {
                                var v42 float32 = 0.0
                                _ = v42
                            }
                            _ = fi36
                        }
                    }
                    _ = v31
                }
                _ = v27
                _ = v28
            }
        }
    default:
        var v43 uint = (^(v4))
        _ = v43
    }
    fmt.Println(v3, v4, v4)
    v1 = ((v1 + ((v1 + "hello") + "")) + v1)
    if (v1 >= ("\n\t" + (v1 + (v1 + v1)))) {
        v6[4294967295] = v1
        v1 = v1
    } else {
        if (!(0.1 <= -2.25)) {
            var v44 bool = (32767 <= 32767)
            var v45 int = -2147483648
            for fi46 := 0; fi46 < 18; fi46++ {
                switch (v45 % (((fi46 / ((v45) | 1))) | 1)) {
                case 32768:
                    var v47 uint = (v4 >> ((255) & 63))
                    _ = v47
                case -250:
                    if (-1 > -2147483648) {
                        switch v1 {
                        case "\n\t":
                            var v48 string = ((v1 + "hello") + "hello")
                            _ = v48
                        default:
                            fmt.Println(fi46, v1)
                        }
                    } else {
                        var v49 float32 = 1.5
                        _ = v49
                    }
                default:
                    var v50 bool = false
                    _ = v50
                }
                v51 := []int{v45, fi46, (fi46 | (fi46 - fi46))}
                for fi52 := 0; fi52 < 7; fi52++ {
                    var v53 bool = v44
                    if v53 {
                        switch v1 {
                        case "foo bar":
                            if (1 != -1) {
                                v54 := GenS4{F0: 1, F1: 2147483647}
                                switch (fi52 - 32768) {
                                case 109:
                                    v4 -= (v4 ^ v4)
                                default:
                                    v55 := GenS3{F0: (v1 + "a"), F1: v44, F2: 172}
                                    _ = v55
                                }
                                _ = v54
                            }
                        default:
                            v4 -= v4
                        }
                        v44 = true
                    } else {
                        v3 += v3
                        if true {
                            v51 = append(v51, 32767)
                        }
                        switch v1 {
                        case "foo bar":
                            fmt.Println(v3, fi52)
                        case "a":
                            if v53 {
                                v56 := GenS3{F0: ((((v1 + "a") + v1) + ("hello" + v1)) + ((v1 + v1) + "hello")), F1: true, F2: (v4 & v4)}
                                v1 = ("foo bar" + v1)
                                _ = v56
                            }
                        default:
                            if (((float64(((float32((v3 - (float64((1.5 * 3.14159)) - ((((3.14159 / -2.25) + 1.5) + ((0.1 / float64(0.1)) * v23)) * (v3 / (v3 * float64(float32(float64((0.1 + (3.14159 * 1.5))))))))))) - 1e10) / (0.0 - 0.1))) - (v23 / v3)) - (v3 / v3)) > ((v3 + (v3 / v3)) - v23)) {
                                fmt.Println(v4, fi52)
                                if (3.14159 < 0.1) {
                                    v57 := GenS5{F0: 0, F1: 127}
                                    v45 = (fi52 + fi52)
                                    _ = v57
                                }
                                fmt.Println(v4)
                            }
                        }
                    }
                    _ = fi52
                    _ = v53
                }
                _ = fi46
                _ = v51
            }
            _ = v44
            _ = v45
        }
        var v58 uint = 18446744073709551615
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
