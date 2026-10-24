package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int; F2 uint }
type GenS4 struct { F0 int32; F1 int8; F2 uint32 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint32 = 593
    var v2 uint16 = 960
    var v3 float64 = 3.14159
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint32 = 4294967295
        var v6 bool = (1.5 >= -2.25)
        v7 := GenS1{F0: (fi4 * fi4), F1: v2, F2: (v2 & 42)}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    if false {
        var v8 string = "hello"
        var v9 float64 = v3
        _ = v8
        _ = v9
    } else {
        var v10 uint8 = uint8(v1)
        var v11 string = "hello"
        _ = v10
        _ = v11
    }
    v2 -= (v2 - v2)
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int64 = -255
        v2 -= (v2 % (((v2 % ((255) | 1))) | 1))
        _ = fi12
        _ = v13
    }
    v1 = v1
    v14 := GenS2{F0: true}
    for fi15 := 0; fi15 < 9; fi15++ {
        v1 += v1
        _ = fi15
    }
    var v16 bool = true
    if v16 {
        switch 4294967295 {
        case 128:
            for fi17 := 0; fi17 < 2; fi17++ {
                v3 = v3
                v18 := []int{-2147483648, int(v2)}
                v19 := GenS2{F0: v16}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        var v20 uint16 = uint16(v1)
        _ = v20
    } else {
        fmt.Println(v1)
        for fi21 := 0; fi21 < 12; fi21++ {
            if (!(-1 > 644)) {
                v1 = 205
                v22 := []int{}
                _ = v22
            } else {
                var v23 uint16 = (v2 / (((v2 << ((uint(v2)) & 63))) | 1))
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 float32 = 0.0
                    var v26 bool = v16
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = v23
            }
            _ = fi21
        }
        if v16 {
            v2 -= (^(v2))
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := GenS1{F0: fi27, F1: 736, F2: v2}
                var v29 uint16 = v2
                v16 = v16
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 uint32 = (v1 << ((65535) & 63))
            _ = v30
        } else {
            v2 = v2
            if (((4294967295 < 1) && v16) && (744 <= 0)) {
                for fi31 := 0; fi31 < 15; fi31++ {
                    if true {
                        var v32 float32 = ((((3.14159 + (-2.25 * (1e10 - (0.1 + (1.5 - float32((1.5 - (v3 + 0.1)))))))) + (0.1 / 3.14159)) + 1.5) / (1.5 * 3.14159))
                        v33 := GenS1{F0: (fi31 << ((702) & 63)), F1: (v2 & 2), F2: v2}
                        v32 = v32
                        _ = v32
                        _ = v33
                    } else {
                        fmt.Println(v3)
                        for fi34 := 0; fi34 < 16; fi34++ {
                            v2 = 42
                            _ = fi34
                        }
                        var v35 uint64 = 549
                        _ = v35
                    }
                    var v36 uint16 = v2
                    fmt.Println(v36, v16, v2)
                    _ = fi31
                    _ = v36
                }
                if (int8(v2) != 100) {
                    for fi37 := 0; fi37 < 2; fi37++ {
                        if v16 {
                            v38 := []int{fi37, (fi37 + 128)}
                            for fi39 := 0; fi39 < 7; fi39++ {
                                v1 += 983
                                v16 = ((-507 <= -9223372036854775808) && v16)
                                _ = fi39
                            }
                            _ = v38
                        } else {
                            v40 := GenS3{F0: -1, F1: 2147483647, F2: uint(v1)}
                            v1 = (^(v1))
                            _ = v40
                        }
                        v1 = v1
                        fmt.Println(v16, v16, v2)
                        _ = fi37
                    }
                    fmt.Println(v16, v3, v1)
                } else {
                    if (127 == 2) {
                        var v41 uint32 = 4294967295
                        _ = v41
                    }
                    switch "a" {
                    case "\n\t":
                        for fi42 := 0; fi42 < 7; fi42++ {
                            if (-2.25 >= ((0.1 - (3.14159 + ((float32(0.0) + float32(1e10)) * -2.25))) / 0.1)) {
                                switch ("hello" + "") {
                                case "hello":
                                    fmt.Println(v16, v3, v3)
                                case "a":
                                    if v16 {
                                        v2 += v2
                                        v3 += -2.25
                                    } else {
                                        v16 = (1000 >= -367)
                                        v43 := GenS2{F0: (!(!(false || v16)))}
                                        for fi44 := 0; fi44 < 1; fi44++ {
                                            var v45 string = "foo bar"
                                            v45 = v45
                                            _ = fi44
                                            _ = v45
                                        }
                                        _ = v43
                                    }
                                }
                                for fi46 := 0; fi46 < 6; fi46++ {
                                    v2 -= v2
                                    _ = fi46
                                }
                            }
                            var v47 float64 = v3
                            _ = fi42
                            _ = v47
                        }
                    case "foo bar":
                        v48 := GenS2{F0: (!(true || false))}
                        _ = v48
                    }
                }
                for fi49 := 0; fi49 < 9; fi49++ {
                    for fi50 := 0; fi50 < 10; fi50++ {
                        v3 = v3
                        var v51 int32 = -825
                        v51 = 190
                        _ = fi50
                        _ = v51
                    }
                    var v52 string = "a"
                    _ = fi49
                    _ = v52
                }
            } else {
                for fi53 := 0; fi53 < 9; fi53++ {
                    if (float32(v3) >= 3.14159) {
                        fmt.Println(v1, fi53, v16)
                    }
                    v54 := []int{fi53, (fi53 / (((fi53 >> ((4294967295) & 63))) | 1))}
                    var v55 uint = 727
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                for fi56 := 0; fi56 < 6; fi56++ {
                    v57 := []int{(fi56 % ((fi56) | 1)), fi56, fi56}
                    v3 = float64(3.14159)
                    _ = len(v57)
                    _ = fi56
                    _ = v57
                }
                switch int(v1) {
                case -1:
                    v58 := GenS1{F0: 2147483647, F1: v2, F2: v2}
                    _ = v58
                case 4294967295:
                    v59 := GenS4{F0: 2147483647, F1: 127, F2: (v1 - v1)}
                    _ = v59
                }
            }
        }
    }
    fmt.Println(v1, v2, v1)
    v16 = (42 <= 757)
    v60 := GenS5{F0: -1}
    v61 := GenS5{F0: -9223372036854775808}
    v3 = (0.1 * -2.25)
    v2 = v2
    v62 := []int{42}
    switch 179 {
    case 127:
        var v63 float32 = -2.25
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v62
}
