package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint; F2 int64 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int32; F1 int16; F2 string }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 int) int8 {
    p0 -= p0
    v5 := GenS1{F0: 727, F1: 100}
    _ = p0
    _ = v5
    return -1
}
func genFn1(p0 int8) uint32 {
    p0 = (p0 + p0)
    fmt.Println(p0)
    fmt.Println(p0)
    if (4294967295 < 1) {
        if false {
            var v6 string = "hello"
            if false {
                v7 := []int{-616}
                if (float64((float32(-2.25) + ((0.0 + 3.14159) - 0.0))) < 3.14159) {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v9 := GenS2{F0: p0}
                        p0 = p0
                        _ = fi8
                        _ = v9
                    }
                    _ = len(v7)
                    if len(v7) > 0 {
                        v7[0] = 9223372036854775807
                    }
                }
                if false {
                    var v10 string = v6
                    var v11 int16 = -87
                    v12 := GenS1{F0: 589, F1: p0}
                    _ = v10
                    _ = v11
                    _ = v12
                }
                _ = v7
            }
            if false {
                v13 := GenS2{F0: (p0 % (((p0 - (-(p0)))) | 1))}
                v14 := []int{32768, 2147483647, -189, 347}
                _ = v13
                _ = v14
            }
            _ = v6
        } else {
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 uint8 = 255
                switch 2147483647 {
                case -2147483648:
                    fmt.Println(v16, p0, fi15)
                case 2147483647:
                    var v17 string = ""
                    _ = v17
                case 127:
                    if true {
                        if (!false) {
                            v18 := GenS3{F0: -1, F1: 1, F2: -883}
                            var v19 float32 = 3.14159
                            var v20 float64 = -2.25
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            if (!(v16 == (v16 + v16))) {
                                fmt.Println(p0)
                                p0 = genFn0(fi15)
                            } else {
                                var v21 int32 = -1
                                _ = v21
                            }
                            p0 = p0
                            var v22 float32 = 1e10
                            _ = v22
                        }
                        v23 := GenS3{F0: 1000, F1: 65535, F2: -151}
                        p0 = 89
                        _ = v23
                    }
                default:
                    var v24 string = ("foo bar" + "")
                    _ = v24
                }
                _ = fi15
                _ = v16
            }
            v25 := GenS4{F0: 1, F1: "\n\t"}
            for fi26 := 0; fi26 < 16; fi26++ {
                if (-1 > 0) {
                    v27 := GenS5{F0: 2147483647, F1: 1000, F2: ""}
                    _ = v27
                }
                _ = fi26
            }
            _ = v25
        }
    } else {
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := GenS6{F0: 42, F1: (fi28 % (((fi29 | 1)) | 1))}
                v31 := []int{}
                _ = fi29
                _ = v30
                _ = v31
            }
            if false {
                v32 := GenS7{F0: 42, F1: 0}
                _ = v32
            }
            _ = fi28
        }
    }
    _ = p0
    return 542
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    var v3 string = "a"
    var v4 int32 = 1
    var v33 float32 = 0.0
    v1 = ((-1 < genFn0(1)) || v1)
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 uint = 2
        var v36 uint32 = genFn1(127)
        switch fi34 {
        case 42:
            if ((-(v4)) >= v4) {
                v37 := GenS8{F0: (("hello" + "") + (v3 + v3)), F1: (v4 + (v4 / ((v4) | 1)))}
                _ = v37
            } else {
                v35 = 1
                for fi38 := 0; fi38 < 9; fi38++ {
                    var v39 float64 = (1.5 / (1.5 / (((1.5 - 0.1) - 0.0) / 1.5)))
                    switch v3 {
                    case "hello":
                        v1 = v1
                    case "\n\t":
                        if (1 < 0) {
                            v39 += v39
                            v40 := GenS5{F0: v4, F1: -32768, F2: "a"}
                            _ = v40
                        }
                    }
                    v4 = (-(v4))
                    _ = fi38
                    _ = v39
                }
                var v41 uint16 = 32767
                _ = v41
            }
        case -1:
            v42 := GenS6{F0: 2, F1: fi34}
            _ = v42
        case 370:
            v33 -= 1e10
        default:
            v43 := GenS3{F0: 42, F1: (v35 << (((v35 & v35)) & 63)), F2: -9223372036854775808}
            _ = v43
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    switch v3 {
    case "foo bar":
        if ((282 < 425) && (1 == 255)) {
            fmt.Println(v3, v4)
        }
    }
    v44 := GenS7{F0: 2147483647, F1: 32767}
    for fi45 := 0; fi45 < 17; fi45++ {
        for fi46 := 0; fi46 < 3; fi46++ {
            v1 = (!v1)
            v47 := GenS9{F0: 255}
            if (v1 && v1) {
                v3 = v3
            } else {
                v3 = v3
                v1 = true
            }
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    v48 := []int{-857, 32768}
    fmt.Println(v3)
    var v49 uint = 550
    v49 -= v49
    switch -9223372036854775808 {
    case -465:
        for fi50 := 0; fi50 < 6; fi50++ {
            var v51 uint16 = 884
            v51 += v51
            v48 = append(v48, 4294967295)
            _ = fi50
            _ = v51
        }
    default:
        v1 = (!(847 == 169))
    }
    var v52 int64 = -1
    v49 = 255
    fmt.Println(v52, v52)
    v48 = append(v48, 755)
    if ((!((v33 * v33) > v33)) && (!((v1 || (v4 < v4)) || (v52 == v52)))) {
        if false {
            fmt.Println(v49, v3, v49)
        }
    } else {
        _ = len(v48)
        fmt.Println(v3)
        switch v3 {
        case "hello":
            v53 := GenS2{F0: 0}
            _ = v53
        case "\n\t":
            var v54 bool = (true || (130 < 127))
            _ = v54
        case "a":
            var v55 uint = (v49 >> ((65535) & 63))
            _ = v55
        default:
            for fi56 := 0; fi56 < 11; fi56++ {
                if (!(!((-(v52)) == (-(v52))))) {
                    _ = len(v2)
                }
                v3 = ("\n\t" + v3)
                switch (fi56 << ((v49) & 63)) {
                case 32767:
                    v57 := GenS5{F0: (-(v4)), F1: -1, F2: ""}
                    _ = v57
                case -2147483648:
                    v4 = v4
                default:
                    var v58 uint64 = 42
                    _ = v58
                }
                _ = fi56
            }
        }
    }
    v59 := GenS10{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v44
    _ = v48
    _ = v49
    _ = v52
    _ = v59
}
