package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
type GenS3 struct { F0 int; F1 string; F2 string }
type GenS4 struct { F0 uint16; F1 bool; F2 int32 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint }
type GenS6 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("foo bar" + "\n\t")
    var v3 float64 = 3.14159
    var v4 string = ("hello" + v2)
    var v5 int32 = 2147483647
    if false {
        delete(v1, 1)
        if ((1 >= 510) || true) {
            v1[65535] = 2
            v3 = (float64(3.14159) - ((float64(0.1) / v3) / v3))
        } else {
            if (!((false || (32767 >= 42)) || false)) {
                if (!(0 < 185)) {
                    v3 = 0.0
                    v6 := v1[2]
                    _ = v6
                } else {
                    var v7 int8 = 1
                    if ((true || false) && (!true)) {
                        var v8 uint = 116
                        var v9 int16 = 962
                        _ = v8
                        _ = v9
                    }
                    _ = v7
                }
            }
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(fi10, fi10)
                _ = fi10
            }
        }
        if (65535 >= 255) {
            v3 += v3
        } else {
            v11 := GenS1{F0: 487}
            v12 := GenS1{F0: v5}
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v2 = v4
        v14 := []int{(fi13 % (((fi13 << ((65535) & 63))) | 1))}
        v2 = v2
        _ = fi13
        _ = v14
    }
    if false {
        fmt.Println(v2)
        if false {
            if false {
                if (!false) {
                    if (-9223372036854775808 != -908) {
                        v15 := []int{541}
                        _ = v15
                    }
                }
                if true {
                    fmt.Println(v3, v2)
                    fmt.Println(v4, v4, v3)
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v2 = (v4 + "")
                        _ = fi16
                    }
                } else {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 float64 = -2.25
                        _ = fi17
                        _ = v18
                    }
                }
            } else {
                fmt.Println(v3)
                v3 = 1e10
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 bool = (true || ((false && (0.1 != 0.0)) && ((true || true) && (!(fi19 == fi19)))))
                    for fi21 := 0; fi21 < 10; fi21++ {
                        switch fi19 {
                        case -801:
                            var v22 int32 = v5
                            _ = v22
                        case 918:
                            if v20 {
                                v23 := GenS2{F0: 998, F1: 32767, F2: (v5 >> ((937) & 63))}
                                fmt.Println(v2)
                                _ = v23
                            }
                        case 2147483647:
                            if v20 {
                                var v24 string = v4
                                v25 := v1[2147483648]
                                var v26 string = v4
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        }
                        var v27 int64 = -1
                        v20 = (false || v20)
                        _ = fi21
                        _ = v27
                    }
                    _ = fi19
                    _ = v20
                }
            }
            v28 := GenS3{F0: -1, F1: "\n\t", F2: ((v4 + v4) + "a")}
            _ = v28
        }
    } else {
        switch v4 {
        case "":
            v2 = ((v2 + "foo bar") + "a")
        default:
            for fi29 := 0; fi29 < 12; fi29++ {
                v3 = ((float64((1.5 - -2.25)) + (v3 * v3)) - (v3 * v3))
                var v30 float64 = (1e10 - v3)
                if true {
                    var v31 int16 = 42
                    _ = v31
                } else {
                    v32 := GenS4{F0: 1, F1: false, F2: (-(v5))}
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
        }
        var v33 int8 = 127
        fmt.Println(v5, v33, v4)
        _ = v33
    }
    v5 = (v5 / (((v5 + (v5 - (v5 << ((672) & 63))))) | 1))
    v4 = v4
    if (0 != 2) {
        var v34 uint = 948
        _ = v34
    } else {
        v1[273] = 2
        switch 269 {
        case -2147483648:
            v5 += v5
        case 32768:
            var v35 int16 = -276
            _ = v35
        case -9223372036854775808:
            for fi36 := 0; fi36 < 14; fi36++ {
                var v37 float64 = -2.25
                _ = fi36
                _ = v37
            }
        default:
            switch (v4 + "a") {
            case "\n\t":
                v1[42] = 2147483647
            case "hello":
                v38 := GenS5{F0: 17, F1: 42, F2: 65535}
                _ = v38
            case "a":
                v39 := []int{0}
                _ = v39
            }
        }
    }
    switch 420 {
    case -988:
        var v40 int16 = -232
        _ = v40
    case 42:
        v41 := []int{}
        _ = v41
    }
    switch -929 {
    case 127:
        _ = len(v1)
    }
    v42 := GenS6{F0: -1, F1: (!false)}
    switch (("a" + v4) + v4) {
    case "foo bar":
        fmt.Println(v5, v3)
    }
    v3 -= 1e10
    v43 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
