package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int64; F1 int; F2 string }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 int8; F1 int32 }
func genFn0() uint {
    v6 := []int{-2147483648, 127, 32768, 128}
    var v7 uint8 = 255
    for fi8 := 0; fi8 < 18; fi8++ {
        if (42 != -32768) {
            v9 := GenS1{F0: (("foo bar" + "hello") + "hello")}
            _ = v9
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            switch fi8 {
            case 454:
                if true {
                    fmt.Println(v7)
                    v11 := GenS2{F0: 1, F1: 0}
                    _ = v11
                } else {
                    v7 = (v7 >> ((18446744073709551615) & 63))
                    v7 -= (v7 - (v7 + (v7 + v7)))
                    switch (-(fi8)) {
                    case 1:
                        if (0 < 746) {
                            var v12 uint16 = 395
                            _ = v12
                        } else {
                            v13 := GenS1{F0: "\n\t"}
                            var v14 uint = 0
                            v14 = v14
                            _ = v13
                            _ = v14
                        }
                    case -2147483648:
                        var v15 uint16 = 109
                        _ = v15
                    default:
                        v7 = (v7 << ((1) & 63))
                        var v16 bool = true
                        v7 = v7
                        v16 = (false && v16)
                        _ = len(v6)
                        _ = v16
                    }
                    _ = len(v6)
                    var v17 uint = 400
                    v17 += 255
                    _ = v17
                }
                _ = len(v6)
            case 1:
                v7 -= v7
            case 683:
                if true {
                    fmt.Println(fi8, v7)
                    v18 := GenS2{F0: 100, F1: 152}
                    _ = v18
                } else {
                    v6 = append(v6, fi10)
                }
            }
            v19 := GenS3{F0: 42, F1: true}
            _ = fi10
            _ = v19
        }
        _ = fi8
    }
    v7 = v7
    _ = v6
    _ = v7
    return 42
}
func genFn1() int32 {
    var v20 uint32 = 2147483648
    v20 += v20
    fmt.Println(v20)
    switch 42 {
    case -316:
        switch "foo bar" {
        case "foo bar":
            switch "hello" {
            case "foo bar":
                if false {
                    if (!((float64(3.14159) + -2.25) <= float64((0.0 / 1.5)))) {
                        v20 = 2147483648
                    } else {
                        v21 := GenS1{F0: ("hello" + "a")}
                        _ = v21
                    }
                } else {
                    if (535 != 2) {
                        switch 4294967295 {
                        case -1:
                            var v22 bool = false
                            _ = v22
                        default:
                            v23 := GenS4{F0: -266, F1: 9223372036854775807, F2: ("foo bar" + ("\n\t" + "\n\t"))}
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 18; fi24++ {
                            var v25 int32 = 42
                            _ = fi24
                            _ = v25
                        }
                        v20 -= 42
                    }
                    v26 := GenS5{F0: 1, F1: 2}
                    if (42 > -9223372036854775808) {
                        var v27 float64 = float64(0.1)
                        _ = v27
                    }
                    _ = v26
                }
            case "a":
                if (!(65535 != (^(v20)))) {
                    v20 = (v20 - v20)
                }
            }
        case "":
            if false {
                for fi28 := 0; fi28 < 16; fi28++ {
                    if false {
                        v29 := []int{(fi28 << ((1) & 63)), (fi28 / ((fi28) | 1)), fi28, int(v20)}
                        v20 = (v20 | (v20 - 4294967295))
                        for fi30 := 0; fi30 < 14; fi30++ {
                            v31 := GenS6{F0: 127}
                            v20 += v20
                            var v32 uint32 = (v20 / ((841) | 1))
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        _ = v29
                    } else {
                        v33 := GenS7{F0: "foo bar"}
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := GenS1{F0: "\n\t"}
                        var v36 uint64 = 462
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    _ = fi28
                }
                for fi37 := 0; fi37 < 8; fi37++ {
                    fmt.Println(fi37)
                    fmt.Println(fi37, v20)
                    _ = fi37
                }
            } else {
                v20 = 2
            }
        case "a":
            v38 := []int{4294967295, -516}
            _ = v38
        }
    case -161:
        v39 := []int{}
        _ = v39
    }
    if ((!false) && (((!((!false) && (!true))) || false) || (true || (-32768 >= 42)))) {
        v40 := []int{32767}
        switch 32767 {
        case 0:
            if (("foo bar" + ("" + "a")) < "hello") {
                fmt.Println(v20)
            }
        default:
            v20 = (v20 << ((4294967295) & 63))
        }
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
        _ = v40
    }
    _ = v20
    return 739
}

func main() {
    var v1 float64 = 3.14159
    var v2 int8 = 46
    v3 := make(map[string]bool)
    var v4 bool = (42 == 2147483647)
    var v5 uint16 = 690
    v41 := GenS8{F0: v5, F1: 336, F2: v2}
    var v42 uint8 = 42
    v2 -= v2
    v4 = v4
    v43 := GenS1{F0: "a"}
    v3[""] = (v4 && (0 < -1))
    v2 = (v2 * v2)
    if (v5 != (v5 & v5)) {
        var v44 int16 = 32767
        _ = v44
    }
    var v45 float64 = 1e10
    if ((!v4) || (727 >= 656)) {
        var v46 uint64 = 143
        v47 := GenS7{F0: "\n\t"}
        _ = v46
        _ = v47
    }
    if false {
        v48 := GenS9{F0: 42, F1: 169}
        _ = v48
    } else {
        _ = len(v3)
        for fi49 := 0; fi49 < 3; fi49++ {
            var v50 int32 = -833
            _ = fi49
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 20; fi51++ {
        var v52 uint16 = (v5 % (((v5 ^ 212)) | 1))
        var v53 int8 = v2
        v54 := GenS10{F0: v2, F1: 1}
        _ = fi51
        _ = v52
        _ = v53
        _ = v54
    }
    for fi55 := 0; fi55 < 17; fi55++ {
        v56 := []int{(fi55 & -561)}
        _ = fi55
        _ = v56
    }
    switch -809 {
    case 42:
        v57 := []int{}
        _ = v57
    case 32768:
        v58 := GenS6{F0: v42}
        _ = v58
    case -519:
        var v59 int64 = int64(v5)
        _ = v59
    }
    var v60 uint64 = 831
    var v61 uint64 = 2
    var v62 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v5
    _ = v60
    _ = v61
    _ = v62
}
