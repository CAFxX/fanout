package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint16 }
func genFn0(p0 int8) int32 {
    fmt.Println(p0)
    v6 := []int{32768, 127, 741, 42}
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 int8, p1 int64, p2 int16) uint32 {
    var v7 uint32 = 704
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}

func main() {
    var v1 uint64 = 757
    var v2 uint32 = 0
    var v3 uint64 = v1
    v4 := make(map[uint32]bool)
    var v5 string = "\n\t"
    v8 := GenS1{F0: "foo bar", F1: v1, F2: 42}
    fmt.Println(v2, v2)
    fmt.Println(v2, v3)
    fmt.Println(v1, v5)
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 int8 = 1
            var v12 uint32 = v2
            v12 += (v2 << ((227) & 63))
            _ = fi10
            _ = v11
            _ = v12
        }
        if true {
            v13 := GenS1{F0: "foo bar", F1: v1, F2: 0}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                fmt.Println(v1)
                if false {
                    v15 := GenS2{F0: 255, F1: 0}
                    _ = v15
                }
                if false {
                    v3 -= (v3 % ((1) | 1))
                    var v16 uint8 = 1
                    _ = v16
                }
                _ = fi14
            }
            var v17 int = fi9
            var v18 int16 = -463
            _ = v17
            _ = v18
        }
        if (-9223372036854775808 >= -9223372036854775808) {
            if true {
                for fi19 := 0; fi19 < 3; fi19++ {
                    v20 := GenS3{F0: 0}
                    _ = fi19
                    _ = v20
                }
            } else {
                v2 -= v2
                var v21 int8 = -11
                var v22 int16 = -342
                _ = v21
                _ = v22
            }
            v5 = ((v5 + v5) + (("hello" + v5) + v5))
            v23 := []int{int(v1), int(v2)}
            _ = v23
        }
        _ = fi9
    }
    v5 = v5
    fmt.Println(v1)
    var v24 uint = 2
    var v25 int16 = 0
    switch 1 {
    case -177:
        if ((false || (float32(0.0) == 0.0)) && (128 != 242)) {
            v2 = (v2 ^ genFn1(1, 73, v25))
            var v26 float64 = -2.25
            var v27 bool = false
            _ = v26
            _ = v27
        } else {
            fmt.Println(v5, v24, v3)
            v28 := []int{32768, 128, 127}
            for fi29 := 0; fi29 < 7; fi29++ {
                _ = len(v28)
                _ = fi29
            }
            _ = v28
        }
    default:
        if (259 == uint32(v1)) {
            for fi30 := 0; fi30 < 16; fi30++ {
                if false {
                    var v31 int8 = -128
                    _ = v31
                }
                v1 = (^(v1))
                switch ("foo bar" + (v5 + "\n\t")) {
                case "":
                    for fi32 := 0; fi32 < 14; fi32++ {
                        var v33 int32 = -1
                        v34 := []int{(fi30 >> ((65535) & 63)), 0}
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                default:
                    for fi35 := 0; fi35 < 2; fi35++ {
                        if true {
                            v36 := GenS2{F0: (v24 % ((4294967295) | 1)), F1: 252}
                            v25 -= 1
                            _ = v36
                        }
                        v3 += v3
                        _ = fi35
                    }
                }
                _ = fi30
            }
        } else {
            _ = len(v4)
            var v37 string = ((v5 + v5) + "a")
            _ = v37
        }
    }
    var v38 int16 = (v25 << (((^(v24))) & 63))
    v38 = v38
    switch "hello" {
    case "\n\t":
        var v39 int = -9223372036854775808
        _ = v39
    case "foo bar":
        v40 := []int{0}
        _ = v40
    case "":
        if false {
            var v41 int64 = 0
            fmt.Println(v5, v25, v3)
            var v42 float64 = 1e10
            _ = v41
            _ = v42
        }
    default:
        v2 = v2
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        for fi44 := 0; fi44 < 12; fi44++ {
            if true {
                var v45 uint16 = 2
                if (!("foo bar" > (v5 + (("" + "a") + (((v5 + v5) + "foo bar") + v5))))) {
                    v46 := GenS2{F0: uint(v2), F1: 1}
                    var v47 uint64 = v3
                    _ = v46
                    _ = v47
                }
                if ((^(v2)) > v2) {
                    v5 = "foo bar"
                    v48 := []int{fi44, (fi44 & fi44), fi43}
                    fmt.Println(v5, v1)
                    if (125 != 1) {
                        if true {
                            _ = len(v4)
                            if (((^(v24)) <= (v24 << (((^(v24))) & 63))) || true) {
                                fmt.Println(v25, fi44)
                                _ = len(v4)
                            } else {
                                v5 = v5
                                v3 += 2
                                v49 := GenS1{F0: "", F1: v1, F2: -1}
                                _ = v49
                            }
                            v3 -= v3
                        } else {
                            v24 -= (^(v24))
                            v38 -= v25
                        }
                        v2 -= 259
                        fmt.Println(fi43, v45, v2)
                    } else {
                        if true {
                            var v50 bool = false
                            _ = v50
                        }
                        switch (((v5 + v5) + (v5 + ("a" + (v5 + v5)))) + (v5 + v5)) {
                        case "foo bar":
                            var v51 float32 = (-2.25 / 1e10)
                            _ = v51
                        }
                        v52 := v4[42]
                        _ = v52
                    }
                    var v53 uint16 = (v45 | v45)
                    _ = v48
                    _ = v53
                }
                _ = v45
            }
            var v54 float64 = ((-2.25 * 0.1) - ((1.5 - 0.1) + 0.0))
            _ = fi44
            _ = v54
        }
        if true {
            v55 := GenS3{F0: v2}
            _ = v55
        }
        if ((42 != 1) && false) {
            if true {
                v3 = 1
                for fi56 := 0; fi56 < 16; fi56++ {
                    v25 = (-(v38))
                    if true {
                        v5 = ((("a" + v5) + v5) + "")
                        v25 = v38
                        v24 = (v24 % ((4294967295) | 1))
                        v38 = v38
                    } else {
                        v57 := []int{fi56, (fi43 & fi43), (fi56 | fi43), 59}
                        v58 := GenS3{F0: (v2 ^ (v2 % ((42) | 1)))}
                        _ = v57
                        _ = v58
                    }
                    v4[v2] = false
                    _ = fi56
                }
            }
        }
        _ = fi43
    }
    v25 -= -1
    if false {
        v59 := GenS4{F0: 593, F1: uint(v2), F2: genFn1(14, -465, 1000)}
        v38 -= v25
        v60 := GenS5{F0: int8(v25), F1: (v2 / (((^(v2))) | 1)), F2: 2}
        _ = v59
        _ = v60
    } else {
        fmt.Println(v5, v24, v5)
    }
    v2 = (v2 & v2)
    var v61 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v61
    _ = v8
}
