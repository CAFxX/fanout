package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint }
func genFn0() string {
    v7 := GenS1{F0: -9223372036854775808}
    var v8 uint = 431
    v8 -= v8
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = true
    var v3 uint32 = 773
    var v4 int32 = 475
    v5 := make(map[string]int)
    v6 := make(map[string]uint32)
    v3 = 2
    v9 := GenS1{F0: -1}
    if v2 {
        var v10 string = (("\n\t" + "hello") + ("" + "a"))
        v10 = v10
        _ = v10
    } else {
        v2 = v2
    }
    if v2 {
        v11 := GenS2{F0: 0, F1: v3}
        var v12 int64 = 9223372036854775807
        for fi13 := 0; fi13 < 5; fi13++ {
            if v2 {
                v4 -= v4
                if v2 {
                    v2 = false
                }
            }
            for fi14 := 0; fi14 < 8; fi14++ {
                v12 -= (-(v12))
                _ = fi14
            }
            fmt.Println(v4, v2)
            _ = fi13
        }
        _ = v11
        _ = v12
    } else {
        if (!(44 == 45)) {
            _ = len(v6)
        }
        var v15 int = -922
        _ = v15
    }
    var v16 float64 = 0.0
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := v1[(v3 * v3)]
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int8 = 1
        for fi22 := 0; fi22 < 12; fi22++ {
            switch fi20 {
            case -1:
                fmt.Println(v2)
            }
            var v23 float32 = (((1.5 + (-2.25 / 3.14159)) - (float32(1.5) - 0.1)) / (((1e10 / -2.25) + (0.0 + (3.14159 / 1.0))) * 1e10))
            _ = fi22
            _ = v23
        }
        switch (fi20 << ((4294967295) & 63)) {
        case 4294967295:
            if (775 >= 0) {
                if v2 {
                    v24 := GenS1{F0: 9223372036854775807}
                    v3 = (v3 << ((0) & 63))
                    _ = v24
                }
            } else {
                if v2 {
                    var v25 int64 = 9223372036854775807
                    v26 := []int{(fi20 - (fi20 << ((2) & 63))), (fi20 % (((fi20 % ((2147483647) | 1))) | 1)), fi20}
                    var v27 float32 = 1e10
                    _ = v25
                    _ = v26
                    _ = v27
                }
                if (false || v2) {
                    for fi28 := 0; fi28 < 9; fi28++ {
                        var v29 uint32 = (v3 >> ((228) & 63))
                        _ = fi28
                        _ = v29
                    }
                    var v30 int16 = -32768
                    v1[(v3 & (v3 - (v3 << ((493) & 63))))] = "a"
                    _ = v30
                }
            }
        case -1:
            v31 := []int{(fi20 ^ 127), (fi20 + fi20)}
            _ = v31
        }
        _ = fi20
        _ = v21
    }
    v32 := v1[v3]
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int64 = -248
        _ = fi33
        _ = v34
    }
    v35 := GenS2{F0: 0, F1: v3}
    v16 = float64((0.0 / 3.14159))
    switch 32767 {
    case 42:
        switch 941 {
        case 978:
            var v36 int64 = -1
            _ = v36
        }
    case -10:
        _ = len(v6)
    default:
        switch -2147483648 {
        case -65:
            v37 := GenS3{F0: 29, F1: 42, F2: 0}
            _ = v37
        case 128:
            if v2 {
                var v38 string = ("foo bar" + "a")
                _ = len(v5)
                if (255 >= 65535) {
                    switch 9223372036854775807 {
                    case 42:
                        if v2 {
                            if ((v16 < v16) || (697 != 67)) {
                                v39 := GenS1{F0: 0}
                                _ = v39
                            }
                            if ("hello" != "foo bar") {
                                v2 = true
                                fmt.Println(v16, v32, v38)
                                v4 -= (^(v4))
                            }
                            for fi40 := 0; fi40 < 18; fi40++ {
                                var v41 float32 = 1.5
                                v42 := []int{fi40, fi40}
                                for fi43 := 0; fi43 < 15; fi43++ {
                                    v16 = ((v16 / v16) / ((0.0 - float64((0.1 + ((-2.25 / 1.5) - v41)))) + v16))
                                    _ = fi43
                                }
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                        } else {
                            v44 := []int{32768, 1}
                            switch "a" {
                            case "":
                                if (false || false) {
                                    var v45 int = 0
                                    var v46 uint16 = 2
                                    _ = v45
                                    _ = v46
                                }
                            }
                            _ = v44
                        }
                    default:
                        for fi47 := 0; fi47 < 10; fi47++ {
                            var v48 uint8 = 2
                            _ = fi47
                            _ = v48
                        }
                    }
                    fmt.Println(v16, v2, v16)
                }
                _ = v38
            } else {
                for fi49 := 0; fi49 < 15; fi49++ {
                    if false {
                        var v50 bool = v2
                        var v51 int64 = -9223372036854775808
                        v1[v3] = (((v32 + "\n\t") + v32) + v32)
                        _ = v50
                        _ = v51
                    } else {
                        v32 = ("hello" + "foo bar")
                        switch "foo bar" {
                        case "hello":
                            v32 = v32
                        case "\n\t":
                            v2 = (-723 > -32768)
                        case "a":
                            v52 := GenS4{F0: -595, F1: 42}
                            _ = v52
                        default:
                            var v53 string = v32
                            _ = v53
                        }
                    }
                    v54 := []int{(fi49 % ((fi49) | 1))}
                    v55 := GenS5{F0: 9}
                    _ = fi49
                    _ = v54
                    _ = v55
                }
                switch 32768 {
                case 9223372036854775807:
                    v3 += (v3 * (v3 + v3))
                case -2147483648:
                    switch 433 {
                    case -2147483648:
                        v16 -= ((-2.25 - (((v16 + float64(1e10)) - (1e10 / 1e10)) - (float64((((-2.25 + 1e10) + (((((1e10 / ((((1.5 * (float32(1e10) + -2.25)) * 3.14159) - 1e10) / 0.1)) - -2.25) / 1.0) + 1e10) * ((0.1 + -2.25) + 0.1))) * (((0.0 + 1e10) - ((float32(v16) / (0.0 - -2.25)) * 1e10)) + 0.1))) / 1.0))) * (-2.25 / v16))
                    default:
                        var v56 float32 = -2.25
                        _ = v56
                    }
                case 1:
                    v57 := GenS1{F0: -9223372036854775808}
                    _ = v57
                }
                switch ((v32 + v32) + "\n\t") {
                case "hello":
                    var v58 int8 = 58
                    _ = v58
                case "":
                    for fi59 := 0; fi59 < 2; fi59++ {
                        v60 := []int{}
                        _ = fi59
                        _ = v60
                    }
                }
            }
        case -2147483648:
            var v61 uint = 4294967295
            _ = v61
        }
    }
    delete(v1, 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
