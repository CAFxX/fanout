package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint16; F1 uint8; F2 string }
type GenS5 struct { F0 int; F1 int8; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 uint8 }
func genFn0() int16 {
    switch -1 {
    case -2147483648:
        var v6 int8 = 0
        _ = v6
    case 678:
        switch -991 {
        case -771:
            var v7 int64 = 817
            _ = v7
        case 32768:
            switch 128 {
            case 32767:
                var v8 uint16 = 304
                _ = v8
            default:
                switch 127 {
                case 32767:
                    var v9 int32 = 1
                    _ = v9
                }
            }
        case 42:
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 int64 = -9223372036854775808
                fmt.Println(v11, v11)
                v12 := GenS1{F0: 1, F1: 1, F2: ((!(v11 != -1)) && true)}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: 159, F1: 32767, F2: true}
        _ = fi13
        _ = v14
    }
    switch 127 {
    case 32767:
        var v15 float32 = (0.1 + float32(0.1))
        _ = v15
    case 1:
        v16 := GenS2{F0: -1, F1: 255}
        _ = v16
    default:
        switch 32768 {
        case -1:
            if true {
                switch 32768 {
                case 32768:
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v18 := GenS1{F0: 0, F1: -32768, F2: ((127 >= 42) && (((1 <= 1) && (32767 < -903)) && false))}
                        for fi19 := 0; fi19 < 7; fi19++ {
                            var v20 int16 = -1
                            v21 := GenS2{F0: (^(fi17)), F1: 1}
                            v20 = v20
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 20; fi22++ {
                            fmt.Println(fi17, fi17)
                            var v23 bool = (4294967295 >= 42)
                            switch ("foo bar" + "") {
                            case "foo bar":
                                v23 = (v23 && false)
                            default:
                                v24 := GenS3{F0: 2, F1: 65535}
                                _ = v24
                            }
                            _ = fi22
                            _ = v23
                        }
                        _ = fi17
                        _ = v18
                    }
                case 32767:
                    switch 128 {
                    case 127:
                        for fi25 := 0; fi25 < 5; fi25++ {
                            var v26 float32 = -2.25
                            var v27 string = ""
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                case 9223372036854775807:
                    var v28 int8 = 114
                    _ = v28
                }
            }
        case 9223372036854775807:
            var v29 uint = 4294967295
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 uint32 = 993
        v32 := []int{(fi30 ^ fi30), (fi30 % ((-1) | 1)), fi30, (fi30 | 0)}
        if true {
            v33 := GenS2{F0: -963, F1: 255}
            if false {
                if (!false) {
                    var v34 float64 = 0.0
                    v35 := GenS4{F0: 929, F1: 255, F2: ("hello" + ("hello" + ("a" + "hello")))}
                    _ = v34
                    _ = v35
                } else {
                    var v36 uint64 = 1
                    _ = v36
                }
                v37 := GenS5{F0: (fi30 | (^(fi30))), F1: 100, F2: 255}
                _ = v37
            } else {
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
            }
            var v38 int = (fi30 * (fi30 >> ((42) & 63)))
            _ = v33
            _ = v38
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    return 32767
}
func genFn1(p0 uint32, p1 int64) bool {
    var v39 int = -238
    switch ("a" + (((("" + "foo bar") + "\n\t") + "\n\t") + ("a" + ""))) {
    case "foo bar":
        switch v39 {
        case 32768:
            switch 42 {
            case 9223372036854775807:
                v39 += v39
            case 2147483647:
                p0 = (^(p0))
            case 148:
                for fi40 := 0; fi40 < 3; fi40++ {
                    v41 := GenS6{F0: -364}
                    if (!(false && true)) {
                        for fi42 := 0; fi42 < 19; fi42++ {
                            v39 = (v39 ^ (fi40 ^ int(p0)))
                            _ = fi42
                        }
                        var v43 uint = 255
                        p1 = (p1 * p1)
                        _ = v43
                    }
                    _ = fi40
                    _ = v41
                }
            default:
                v39 = (v39 % ((-9223372036854775808) | 1))
            }
        case 42:
            v44 := []int{}
            _ = v44
        case -913:
            v45 := []int{127, (v39 + v39)}
            _ = v45
        }
    }
    _ = p0
    _ = p1
    _ = v39
    return (true && (0.1 < 3.14159))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    v3 := []int{127}
    v4 := []uint32{185, 237, 4294967295, 540}
    v5 := make(map[string]string)
    var v46 uint8 = 2
    switch "\n\t" {
    case "hello":
        v46 = 255
    case "foo bar":
        v46 = (v46 - v46)
    }
    v46 -= 0
    v47 := GenS1{F0: 0, F1: 42, F2: (true || true)}
    v46 = (v46 - (^(v46)))
    for fi48 := 0; fi48 < 5; fi48++ {
        v49 := GenS2{F0: (fi48 + (fi48 ^ (^(fi48)))), F1: 2}
        _ = fi48
        _ = v49
    }
    v46 = (v46 ^ v46)
    if ((3.14159 + ((0.0 / -2.25) + 3.14159)) >= 3.14159) {
        v1["hello"] = 41
        var v50 uint16 = 79
        _ = v50
    } else {
        v46 = (^(v46))
    }
    var v51 bool = false
    var v52 float32 = -2.25
    if (!v51) {
        for fi53 := 0; fi53 < 2; fi53++ {
            var v54 string = "hello"
            v51 = v51
            _ = fi53
            _ = v54
        }
        for fi55 := 0; fi55 < 18; fi55++ {
            v56 := GenS7{F0: -551, F1: -1}
            for fi57 := 0; fi57 < 18; fi57++ {
                v58 := GenS6{F0: -2147483648}
                var v59 int8 = -79
                if (!true) {
                    var v60 int16 = genFn0()
                    v52 = v52
                    _ = v60
                } else {
                    v61 := GenS4{F0: 533, F1: 42, F2: "hello"}
                    v59 += 121
                    v62 := GenS8{F0: v46}
                    _ = v61
                    _ = v62
                }
                _ = fi57
                _ = v58
                _ = v59
            }
            v46 += (v46 + v46)
            _ = fi55
            _ = v56
        }
        for fi63 := 0; fi63 < 18; fi63++ {
            v64 := GenS2{F0: 42, F1: 2}
            v52 = 3.14159
            v46 = 164
            _ = fi63
            _ = v64
        }
    } else {
        v52 -= v52
        fmt.Println(v52)
        var v65 int8 = 37
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v51
    _ = v52
}
