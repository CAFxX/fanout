package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
func genFn0(p0 uint, p1 int8) int16 {
    for fi4 := 0; fi4 < 13; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            p0 = 255
            if (2147483647 != 65535) {
                var v6 uint64 = 0
                v6 = 1
                v7 := GenS1{F0: 1000, F1: ((v6 / ((v6) | 1)) <= v6), F2: 42}
                _ = v6
                _ = v7
            }
            _ = fi5
        }
        _ = fi4
    }
    if (65535 > 2) {
        v8 := GenS1{F0: 912, F1: (0 >= 2147483647), F2: 2147483648}
        v9 := []int{127, -658, 32768, 1}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return -442
}
func genFn1(p0 int32, p1 string, p2 uint16) uint64 {
    switch -1 {
    case 32767:
        p2 -= 32767
    case 42:
        v10 := GenS1{F0: 32767, F1: true, F2: 65535}
        _ = v10
    }
    p0 = (p0 ^ (-(p0)))
    v11 := GenS1{F0: -717, F1: (!false), F2: 2147483647}
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(fi12)
        var v13 uint32 = 929
        _ = fi12
        _ = v13
    }
    if false {
        p1 = (p1 + (("hello" + "foo bar") + p1))
        var v14 int32 = p0
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return uint64(p2)
}

func main() {
    var v1 int64 = 850
    var v2 int16 = 42
    v3 := make(map[string]int)
    v1 += 9223372036854775807
    fmt.Println(v2)
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(v1, fi15)
        _ = fi15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 bool = true
        v1 = (v1 << ((2) & 63))
        v2 -= (v2 / ((42) | 1))
        _ = fi16
        _ = v17
    }
    var v18 uint = 255
    if true {
        var v19 uint16 = 827
        fmt.Println(v18)
        _ = v19
    }
    v20 := GenS1{F0: (v2 - (v2 + 118)), F1: true, F2: 1}
    if (!(!(42 == 707))) {
        v21 := GenS2{F0: true, F1: 242, F2: 958}
        fmt.Println(v18, v18, v18)
        var v22 float32 = 0.1
        _ = v21
        _ = v22
    }
    var v23 uint = v18
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 3; fi25++ {
            v18 = (^(v18))
            switch (("foo bar" + "\n\t") + "foo bar") {
            case "a":
                if (false || (0 >= 42)) {
                    if (!true) {
                        v3["a"] = 128
                    } else {
                        v23 = v23
                        v26 := []int{-709}
                        v1 = v1
                        _ = v26
                    }
                } else {
                    v23 = v18
                    switch fi25 {
                    case 2147483647:
                        if (-1 == 55) {
                            if false {
                                if true {
                                    fmt.Println(fi25, fi25)
                                    var v27 int8 = -1
                                    var v28 uint32 = 0
                                    v29 := GenS1{F0: int16(v27), F1: (v2 != v2), F2: (^(v28))}
                                    v23 = 0
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                } else {
                                    v30 := []int{(fi24 - fi25), (fi25 + (fi24 | (fi24 + 2147483647)))}
                                    _ = v30
                                }
                                var v31 int32 = 42
                                _ = v31
                            }
                            if ((v2 == genFn0(2, -4)) && (2147483647 >= 2147483647)) {
                                var v32 int64 = (v1 / ((v1) | 1))
                                if (0 < 124) {
                                    v33 := GenS2{F0: (true || false), F1: 2, F2: 2}
                                    switch fi24 {
                                    case 128:
                                        v34 := GenS2{F0: true, F1: 255, F2: 182}
                                        _ = v34
                                    case 32768:
                                        v3["foo bar"] = -2147483648
                                    case -2147483648:
                                        v3[("foo bar" + ("a" + "\n\t"))] = (fi24 / ((fi24) | 1))
                                    default:
                                        v1 = (v32 % ((v1) | 1))
                                    }
                                    _ = v33
                                }
                                _ = v32
                            } else {
                                v35 := []int{-2147483648, (fi25 % ((fi25) | 1)), (fi24 + -9223372036854775808)}
                                v36 := GenS2{F0: ((!true) && (!false)), F1: 995, F2: 2}
                                v3["hello"] = fi24
                                _ = v35
                                _ = v36
                            }
                            v1 += 42
                            if (!((fi24 * -2147483648) != (^(fi24)))) {
                                v37 := []int{fi25, fi24}
                                _ = v37
                            }
                        }
                    case 230:
                        v23 += 255
                    case -170:
                        v1 = (v1 & 1)
                        var v38 float64 = 1.5
                        _ = v38
                    default:
                        v2 = v2
                    }
                    var v39 string = ("hello" + "foo bar")
                    _ = v39
                }
            case "":
                v3["hello"] = (fi24 % (((fi24 + fi24)) | 1))
            case "foo bar":
                v18 = (v18 >> ((v23) & 63))
            default:
                fmt.Println(fi24)
            }
            v40 := GenS2{F0: false, F1: 42, F2: 0}
            v41 := []int{}
            _ = fi25
            _ = v40
            _ = v41
        }
        v42 := []int{42}
        switch (fi24 + 9223372036854775807) {
        case 0:
            v2 = v2
        }
        _ = fi24
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
