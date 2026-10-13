package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint8) int32 {
    p0 -= (^(p0))
    var v7 uint16 = 255
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1() uint8 {
    if true {
        for fi8 := 0; fi8 < 10; fi8++ {
            switch (fi8 ^ fi8) {
            case 42:
                var v9 uint = 4294967295
                _ = v9
            case -1:
                fmt.Println(fi8)
            default:
                var v10 uint32 = 2
                _ = v10
            }
            _ = fi8
        }
        var v11 int8 = 44
        switch "hello" {
        case "foo bar":
            if true {
                v12 := []int{32768, 9223372036854775807, 1}
                _ = v12
            }
        }
        _ = v11
    }
    v13 := GenS1{F0: 65535}
    var v14 uint8 = 128
    _ = v13
    _ = v14
    return v14
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]string)
    var v3 int64 = -1
    var v4 float64 = (0.1 - 1e10)
    var v5 int32 = -1
    var v6 uint = 18446744073709551615
    v15 := GenS1{F0: 42}
    var v16 int8 = 1
    delete(v2, -9223372036854775808)
    if (-300 > 2147483647) {
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := []int{127}
            switch "foo bar" {
            case "hello":
                var v19 uint8 = 127
                _ = v19
            case "foo bar":
                v20 := GenS1{F0: 954}
                _ = v20
            case "":
                var v21 int16 = v1
                _ = v21
            default:
                v3 += (v3 % ((v3) | 1))
            }
            switch "hello" {
            case "":
                fmt.Println(v1)
            case "a":
                delete(v2, fi17)
            case "foo bar":
                if false {
                    if len(v18) > 0 {
                        v18[0] = (fi17 >> (((v6 | 42)) & 63))
                    }
                    for fi22 := 0; fi22 < 8; fi22++ {
                        v16 = v16
                        if (!false) {
                            v6 = v6
                            var v23 uint16 = 42
                            _ = v23
                        } else {
                            if false {
                                if (1 == 32767) {
                                    switch "hello" {
                                    case "\n\t":
                                        var v24 int16 = -910
                                        _ = v24
                                    case "":
                                        if ((fi17 << ((1) & 63)) == -1) {
                                            var v25 string = (("foo bar" + ("\n\t" + ("hello" + "\n\t"))) + "hello")
                                            _ = v25
                                        } else {
                                            switch fi17 {
                                            case -480:
                                                v1 -= (v1 << ((0) & 63))
                                            default:
                                                v4 = float64(0.0)
                                            }
                                            var v26 uint = (v6 + (v6 % ((65535) | 1)))
                                            _ = v26
                                        }
                                    }
                                    var v27 string = "hello"
                                    v4 -= (((((v4 + v4) / v4) + (3.14159 + v4)) - v4) - (float64(3.14159) * (v4 - 1.5)))
                                    _ = v27
                                } else {
                                    delete(v2, (^(fi17)))
                                }
                                v3 = (v3 << (((v6 % ((4294967295) | 1))) & 63))
                            } else {
                                fmt.Println(v1)
                                var v28 int = fi17
                                _ = v28
                            }
                        }
                        if true {
                            if (!(!(0.1 >= ((float32((((v4 / v4) - v4) * v4)) / 3.14159) / -2.25)))) {
                                var v29 uint = (v6 ^ (^(v6)))
                                v4 += 0.1
                                _ = v29
                            } else {
                                var v30 uint64 = 171
                                var v31 int64 = -576
                                _ = v30
                                _ = v31
                            }
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v5, v4)
                            case "foo bar":
                                _ = len(v18)
                            default:
                                v32 := v2[fi22]
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                if false {
                                    var v33 int64 = 0
                                    if true {
                                        var v34 uint64 = 42
                                        _ = v34
                                    }
                                    _ = v33
                                } else {
                                    v2[fi17] = v32
                                    v35 := GenS2{F0: v32}
                                    var v36 uint8 = genFn1()
                                    if (!true) {
                                        var v37 int = (fi17 / ((fi22) | 1))
                                        v1 = v1
                                        _ = v37
                                    } else {
                                        fmt.Println(v32, v5)
                                        v38 := GenS3{F0: ""}
                                        v18 = append(v18, (fi22 - (fi22 / (((fi22 ^ 42)) | 1))))
                                        _ = len(v18)
                                        _ = v38
                                    }
                                    _ = v35
                                    _ = v36
                                }
                                _ = v32
                            }
                        }
                        _ = fi22
                    }
                    for fi39 := 0; fi39 < 15; fi39++ {
                        var v40 uint8 = genFn1()
                        v1 = -32768
                        if true {
                            var v41 int64 = v3
                            v41 = 9223372036854775807
                            var v42 float64 = (((v4 - ((1e10 - float64(((3.14159 * (0.0 * 3.14159)) - 0.0))) / ((float64(((1e10 + 0.1) / ((1.5 / 1.0) + 1e10))) * -2.25) * (0.1 - (v4 / v4))))) / v4) * 0.1)
                            _ = v41
                            _ = v42
                        } else {
                            if (v6 != 4294967295) {
                                var v43 uint16 = 190
                                switch fi17 {
                                case 2147483647:
                                    v44 := GenS4{F0: (v43 << (((v6 * (v6 << ((v6) & 63)))) & 63))}
                                    _ = v44
                                case 32768:
                                    _ = len(v2)
                                    v45 := GenS1{F0: 407}
                                    _ = v45
                                }
                                _ = v43
                            } else {
                                var v46 int16 = v1
                                if (1 == 41) {
                                    v6 -= 34
                                    if len(v18) > 0 {
                                        v18[0] = (fi39 >> ((570) & 63))
                                    }
                                    fmt.Println(v40, fi39, fi39)
                                } else {
                                    v1 += v1
                                }
                                _ = v46
                            }
                            v6 -= 4294967295
                        }
                        if len(v18) > 0 {
                            v18[0] = 32767
                        }
                        v40 = 2
                        switch "a" {
                        case "a":
                            v47 := GenS1{F0: 77}
                            _ = v47
                        case "hello":
                            delete(v2, fi39)
                        case "":
                            var v48 uint = 696
                            _ = v48
                        default:
                            switch "a" {
                            case "foo bar":
                                v49 := GenS3{F0: ((("\n\t" + ("a" + (("" + "foo bar") + ("\n\t" + "foo bar")))) + "foo bar") + "")}
                                _ = v49
                            default:
                                var v50 bool = false
                                _ = v50
                            }
                        }
                        _ = fi39
                        _ = v40
                    }
                }
            default:
                fmt.Println(v5)
            }
            _ = fi17
            _ = v18
        }
        v5 += (v5 | (v5 << ((42) & 63)))
        if (127 <= -417) {
            var v51 uint32 = 1
            for fi52 := 0; fi52 < 19; fi52++ {
                v16 = v16
                v53 := v2[int(v5)]
                _ = fi52
                _ = v53
            }
            var v54 string = "hello"
            _ = v51
            _ = v54
        }
    } else {
        fmt.Println(v6)
    }
    var v55 uint16 = 1
    v16 += (v16 << ((v6) & 63))
    if (!true) {
        v2[-123] = "foo bar"
    }
    fmt.Println(v55, v5, v55)
    if false {
        if ((("a" <= "foo bar") && false) && true) {
            v56 := GenS4{F0: v55}
            var v57 float64 = 1e10
            _ = v56
            _ = v57
        } else {
            v1 = (-(v1))
            v58 := v2[1]
            v55 = (v55 - v55)
            _ = v58
        }
    } else {
        var v59 int8 = (-(v16))
        fmt.Println(v1)
        v55 = v55
        _ = v59
    }
    var v60 bool = true
    v61 := []int{-408, 0}
    v1 -= v1
    var v62 int64 = v3
    if v60 {
        switch -280 {
        case 128:
            if (1e10 < -2.25) {
                fmt.Println(v5, v55, v6)
                fmt.Println(v55, v60, v1)
            }
        case 783:
            v63 := GenS2{F0: "hello"}
            _ = v63
        }
        v60 = ((v1 & v1) >= v1)
        v1 -= v1
    } else {
        fmt.Println(v1, v16, v60)
    }
    switch -696 {
    case 750:
        var v64 uint8 = 128
        _ = v64
    default:
        if (0.0 >= ((float32(0.1) + 3.14159) * 0.0)) {
            fmt.Println(v60, v16)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
