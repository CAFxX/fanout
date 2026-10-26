package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int16; F1 uint16 }
func genFn0() bool {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 int16 = -728
        var v6 int32 = -2147483648
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{-849, 523}
    if false {
        v8 := GenS1{F0: 4294967295}
        v9 := GenS1{F0: 99}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(fi10)
        if false {
            v11 := GenS1{F0: 65535}
            _ = v11
        } else {
            if true {
                v12 := GenS1{F0: 255}
                _ = v12
            }
        }
        var v13 float32 = float32(-2.25)
        _ = fi10
        _ = v13
    }
    switch "a" {
    case "a":
        var v14 int = 9223372036854775807
        _ = v14
    case "hello":
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int8 = 1
            v7 = append(v7, fi15)
            _ = fi15
            _ = v16
        }
    default:
        if (1 != -32768) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        } else {
            _ = len(v7)
        }
    }
    _ = v7
    return (2 >= 46)
}
func genFn1() uint16 {
    var v17 uint16 = 110
    v18 := []int{1, 149, -136}
    v19 := GenS2{F0: 42}
    v18 = append(v18, 758)
    _ = v17
    _ = v18
    _ = v19
    return v17
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    var v3 int64 = -1
    if false {
        switch "foo bar" {
        case "\n\t":
            if (1 != -1) {
                fmt.Println(v3)
                if true {
                    if (559 > v3) {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 int8 = 100
                            v22 := []int{fi20, -1, fi20, fi20}
                            var v23 float64 = -2.25
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    }
                    for fi24 := 0; fi24 < 8; fi24++ {
                        var v25 float32 = 0.1
                        _ = fi24
                        _ = v25
                    }
                }
            } else {
                var v26 int8 = -128
                fmt.Println(v2, v26, v26)
                switch 2147483647 {
                case 0:
                    fmt.Println(v2, v2)
                case 32767:
                    switch 127 {
                    case -827:
                        for fi27 := 0; fi27 < 3; fi27++ {
                            if true {
                                v3 = 955
                                var v28 int64 = (v3 | (v3 << ((313) & 63)))
                                for fi29 := 0; fi29 < 20; fi29++ {
                                    v30 := GenS3{F0: fi27}
                                    v26 = v26
                                    var v31 int = (fi27 + 0)
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                _ = v28
                            }
                            fmt.Println(v26, v3)
                            _ = fi27
                        }
                    default:
                        var v32 uint16 = 994
                        _ = v32
                    }
                case -846:
                    v33 := GenS3{F0: 127}
                    _ = v33
                default:
                    if (((false && (!false)) && (!true)) || false) {
                        v3 += v3
                        v3 += (v3 ^ -1)
                        var v34 uint8 = 1
                        _ = v34
                    }
                }
                _ = v26
            }
        case "hello":
            var v35 float64 = (3.14159 * (-2.25 - 1e10))
            _ = v35
        }
    }
    var v36 int16 = 1
    v2 -= (^(v2))
    v37 := GenS4{F0: -2147483648, F1: 215, F2: (v36 << ((65535) & 63))}
    for fi38 := 0; fi38 < 5; fi38++ {
        var v39 uint16 = 0
        v40 := []int{fi38, fi38}
        if (true || (1000 < v36)) {
            v41 := v1[fi38]
            for fi42 := 0; fi42 < 5; fi42++ {
                fmt.Println(fi42)
                _ = fi42
            }
            _ = v41
        } else {
            v36 -= (v36 ^ v36)
            v39 -= v39
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    if (42 < 95) {
        v43 := GenS5{F0: 103, F1: (v3 | v3)}
        delete(v1, 32768)
        if false {
            v36 += v36
            v44 := GenS4{F0: 42, F1: 2, F2: v36}
            v36 += (v36 << ((995) & 63))
            _ = v44
        } else {
            if true {
                v45 := GenS1{F0: 65535}
                v46 := GenS6{F0: v36, F1: 1}
                _ = v45
                _ = v46
            } else {
                for fi47 := 0; fi47 < 8; fi47++ {
                    for fi48 := 0; fi48 < 9; fi48++ {
                        v3 -= (-(v3))
                        fmt.Println(v3, v36, v3)
                        var v49 int16 = (v36 | 516)
                        _ = fi48
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 20; fi50++ {
                        v2 -= (v2 & v2)
                        var v51 uint8 = (v2 + (v2 ^ (v2 >> ((2) & 63))))
                        _ = len(v1)
                        _ = fi50
                        _ = v51
                    }
                    _ = fi47
                }
                if (2 != 42) {
                    v2 = (v2 >> ((65535) & 63))
                    for fi52 := 0; fi52 < 1; fi52++ {
                        v53 := []int{(-(fi52))}
                        _ = fi52
                        _ = v53
                    }
                    v1[-538] = ((("" + (("\n\t" + ("hello" + "foo bar")) + "hello")) + ("\n\t" + "")) + "hello")
                }
                if (552 >= 269) {
                    var v54 uint64 = 1
                    if true {
                        fmt.Println(v3, v54)
                    } else {
                        for fi55 := 0; fi55 < 13; fi55++ {
                            _ = len(v1)
                            _ = fi55
                        }
                        var v56 uint = 42
                        _ = v56
                    }
                    fmt.Println(v3)
                    _ = v54
                }
            }
            v2 -= v2
            _ = len(v1)
        }
        _ = v43
    } else {
        var v57 string = ("a" + "foo bar")
        _ = v57
    }
    var v58 int32 = 522
    for fi59 := 0; fi59 < 1; fi59++ {
        var v60 int16 = v36
        v61 := GenS1{F0: 451}
        switch 0 {
        case 128:
            if (!false) {
                v58 -= (v58 << ((42) & 63))
            }
        case 0:
            fmt.Println(v58, v60, v60)
        case 4294967295:
            var v62 bool = ((v36 * (v60 / ((1000) | 1))) > (^(v36)))
            _ = v62
        }
        _ = fi59
        _ = v60
        _ = v61
    }
    v58 = -437
    var v63 float32 = -2.25
    v64 := GenS2{F0: 0}
    if ((((!true) && false) || false) && (1.5 != ((-2.25 * -2.25) + 0.0))) {
        var v65 uint8 = (v2 / ((1) | 1))
        v66 := GenS1{F0: 1}
        for fi67 := 0; fi67 < 15; fi67++ {
            v3 += (v3 << ((1) & 63))
            fmt.Println(v65)
            _ = fi67
        }
        _ = v65
        _ = v66
    }
    for fi68 := 0; fi68 < 2; fi68++ {
        var v69 string = ""
        var v70 float32 = v63
        _ = fi68
        _ = v69
        _ = v70
    }
    v71 := []int{-362, -479, -189, 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v58:", v58)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v58
    _ = v63
    _ = v64
    _ = v71
}
