package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint }
type GenS4 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint; F1 int16 }
func genFn0() uint16 {
    var v7 bool = false
    v7 = false
    if v7 {
        if v7 {
            fmt.Println(v7)
            v7 = v7
            for fi8 := 0; fi8 < 18; fi8++ {
                switch (fi8 | -471) {
                case 42:
                    v7 = ((!v7) && (152 <= 255))
                case 18:
                    var v9 float32 = 0.1
                    _ = v9
                case 9223372036854775807:
                    fmt.Println(fi8, v7)
                }
                _ = fi8
            }
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(fi10)
            _ = fi10
        }
    }
    var v11 int16 = 32767
    _ = v11
    _ = v7
    return 42
}
func genFn1(p0 uint16) int64 {
    if ((42 >= 1) && false) {
        v12 := []int{}
        var v13 uint32 = uint32(p0)
        if true {
            p0 = (p0 & (p0 & (p0 * (p0 >> ((uint(v13)) & 63)))))
            var v14 uint = 255
            _ = v14
        }
        _ = v12
        _ = v13
    } else {
        var v15 float64 = 1e10
        _ = v15
    }
    fmt.Println(p0)
    var v16 uint32 = 42
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := []int{127, (fi17 << ((289) & 63)), 224, (fi17 >> ((4294967295) & 63))}
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]int)
    var v3 uint8 = 0
    v4 := []string{"hello", "", ""}
    v5 := make(map[uint32]string)
    v6 := []string{}
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS1{F0: 1}
        if (22 < 0) {
            v21 := GenS2{F0: (false || (!false))}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 10; fi22++ {
                if false {
                    v3 -= (v3 >> ((42) & 63))
                    v23 := GenS1{F0: 1}
                    var v24 int16 = 1000
                    _ = v23
                    _ = v24
                }
                var v25 uint16 = 255
                if (true && false) {
                    v26 := GenS1{F0: 0}
                    _ = v26
                } else {
                    var v27 int64 = -101
                    v28 := GenS1{F0: 77}
                    fmt.Println(v1, v25)
                    _ = v27
                    _ = v28
                }
                if (42 <= 0) {
                    switch "foo bar" {
                    case "":
                        v25 -= v25
                    case "hello":
                        v4 = append(v4, "hello")
                    }
                    switch (fi19 / (((fi22 & fi19)) | 1)) {
                    case -1:
                        var v29 uint16 = genFn0()
                        _ = v29
                    default:
                        fmt.Println(fi19, fi19, fi22)
                        if true {
                            var v30 uint = uint(v25)
                            fmt.Println(v25)
                            var v31 float32 = 3.14159
                            v31 += ((1e10 + (v31 - (float32(-2.25) - (v31 / (float32(1e10) + v31))))) - float32(0.1))
                            if ((fi22 == -9223372036854775808) && true) {
                                var v32 float64 = -2.25
                                var v33 int32 = 1
                                _ = v32
                                _ = v33
                            } else {
                                v30 = (^(v30))
                                if false {
                                    v34 := GenS3{F0: 42, F1: uint32(v30), F2: v30}
                                    _ = v34
                                }
                            }
                            _ = v30
                            _ = v31
                        }
                    }
                } else {
                    var v35 bool = ("" < "\n\t")
                    _ = v35
                }
                v25 -= 629
                var v36 int64 = genFn1(782)
                _ = fi22
                _ = v25
                _ = v36
            }
            var v37 int8 = 0
            delete(v5, 963)
            _ = v37
        }
        _ = fi19
        _ = v20
    }
    if ((false || (true || true)) || true) {
        var v38 int8 = -1
        _ = v38
    }
    v1 -= v1
    if (false || false) {
        v1 += (v1 + (^(v1)))
        v3 += (v3 - v3)
    }
    v3 -= v3
    for fi39 := 0; fi39 < 4; fi39++ {
        if len(v4) > 0 {
            v4[0] = "a"
        }
        _ = fi39
    }
    v40 := GenS1{F0: 42}
    var v41 float64 = 1e10
    if ("hello" > "") {
        var v42 string = "foo bar"
        _ = v42
    } else {
        v43 := GenS4{F0: v1, F1: (v3 != v3), F2: v3}
        _ = v43
    }
    if false {
        var v44 int = 1
        switch v44 {
        case 128:
            var v45 float32 = 0.1
            _ = v45
        }
        var v46 int64 = 0
        _ = v44
        _ = v46
    } else {
        var v47 uint32 = 2147483647
        if true {
            for fi48 := 0; fi48 < 10; fi48++ {
                fmt.Println(v1)
                _ = fi48
            }
            v3 += v3
            if true {
                fmt.Println(v3, v3)
                for fi49 := 0; fi49 < 19; fi49++ {
                    if (genFn1(32767) < 19) {
                        var v50 uint64 = v1
                        v47 = v47
                        for fi51 := 0; fi51 < 19; fi51++ {
                            v52 := GenS2{F0: true}
                            v2[((("foo bar" + "a") + ("hello" + ("" + ("hello" + "")))) + ("\n\t" + ("hello" + "hello")))] = (fi51 * fi51)
                            _ = fi51
                            _ = v52
                        }
                        _ = v50
                    }
                    v53 := GenS5{F0: 255, F1: 1000}
                    _ = fi49
                    _ = v53
                }
                v3 += v3
            }
        }
        v41 = v41
        _ = v47
    }
    switch "foo bar" {
    case "foo bar":
        _ = len(v5)
    case "hello":
        fmt.Println(v41, v1)
    case "a":
        if (2147483648 <= 925) {
            v2["foo bar"] = -2147483648
            switch (((("a" + ("a" + "")) + "a") + "foo bar") + "") {
            case "\n\t":
                v54 := GenS4{F0: (v1 + uint64(v3)), F1: (!true), F2: (v3 & (^(v3)))}
                _ = v54
            case "foo bar":
                switch "a" {
                case "a":
                    fmt.Println(v41)
                default:
                    v41 = v41
                }
            case "hello":
                v1 = 4294967295
            default:
                delete(v5, 775)
            }
            switch -229 {
            case 1:
                for fi55 := 0; fi55 < 13; fi55++ {
                    var v56 int32 = 1
                    delete(v5, 4294967295)
                    _ = fi55
                    _ = v56
                }
            case 9223372036854775807:
                for fi57 := 0; fi57 < 7; fi57++ {
                    _ = len(v6)
                    var v58 bool = (0 != 0)
                    _ = fi57
                    _ = v58
                }
            }
        } else {
            for fi59 := 0; fi59 < 13; fi59++ {
                v1 = (v1 * (v1 / (((v1 + v1)) | 1)))
                if (((((3.14159 + 0.0) * (1e10 - 0.0)) - (float32(v41) + (float32(3.14159) - 1.5))) - 0.1) != 1.5) {
                    switch 4294967295 {
                    case -381:
                        var v60 int8 = 100
                        _ = v60
                    default:
                        if true {
                            var v61 int8 = -128
                            for fi62 := 0; fi62 < 14; fi62++ {
                                fmt.Println(v61, v61)
                                _ = fi62
                            }
                            var v63 uint8 = v3
                            _ = v61
                            _ = v63
                        }
                    }
                } else {
                    if true {
                        var v64 uint = 42
                        if (!(1.5 <= float32(v41))) {
                            for fi65 := 0; fi65 < 11; fi65++ {
                                v3 -= (v3 * (v3 + uint8(v64)))
                                switch fi65 {
                                case 127:
                                    if false {
                                        v3 += 2
                                    }
                                case 2147483647:
                                    if true {
                                        v66 := GenS1{F0: 127}
                                        v67 := GenS1{F0: 125}
                                        v1 += (v1 << ((v64) & 63))
                                        _ = v66
                                        _ = v67
                                    }
                                case 32767:
                                    if false {
                                        fmt.Println(v41, v3)
                                        v5[42] = "foo bar"
                                    }
                                }
                                _ = fi65
                            }
                            v4 = append(v4, "")
                            var v68 int64 = 9223372036854775807
                            _ = v68
                        }
                        _ = v64
                    } else {
                        fmt.Println(fi59)
                    }
                    switch (fi59 | fi59) {
                    case 4294967295:
                        for fi69 := 0; fi69 < 9; fi69++ {
                            v70 := GenS1{F0: int8(fi69)}
                            v1 += (v1 * (^(v1)))
                            v1 = (^(v1))
                            _ = fi69
                            _ = v70
                        }
                    case 195:
                        if (!false) {
                            v1 = (^(v1))
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                            fmt.Println(v3)
                        }
                    }
                }
                _ = fi59
            }
            v1 -= v1
        }
    }
    v71 := GenS2{F0: false}
    v72 := v2["foo bar"]
    fmt.Println(v3, v72, v3)
    delete(v2, "")
    if true {
        var v73 uint16 = 1
        v1 += 1
        var v74 float32 = 1e10
        _ = v73
        _ = v74
    }
    v72 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v71
    _ = v72
}
