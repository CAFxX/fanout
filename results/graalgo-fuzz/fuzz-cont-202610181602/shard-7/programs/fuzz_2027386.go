package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 0
    var v3 int64 = 937
    var v4 int = 4294967295
    var v5 int = v4
    var v6 int64 = (v3 | 0)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int16 = 42
        v9 := GenS1{F0: 0}
        v10 := []int{fi7, 9223372036854775807, (v5 | v5)}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 bool = (3.14159 != -2.25)
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v6)
        fmt.Println(v4, fi12, v2)
        _ = fi12
    }
    var v13 int = -1
    var v14 int = (-(v4))
    var v15 uint8 = 1
    v11 = ((1.5 > 0.1) || (v11 && v11))
    if ((!(324 != v2)) && (2 < 95)) {
        v4 += (v4 << ((42) & 63))
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 float64 = (1e10 * -2.25)
            if (false && v11) {
                var v18 uint32 = 975
                _ = v18
            } else {
                if v11 {
                    delete(v1, 0)
                    if (!v11) {
                        v14 -= (-(v13))
                        v19 := GenS1{F0: (v2 / ((v2) | 1))}
                        if (2 != 42) {
                            switch ("a" + "hello") {
                            case "\n\t":
                                for fi20 := 0; fi20 < 15; fi20++ {
                                    var v21 float32 = 1e10
                                    _ = fi20
                                    _ = v21
                                }
                            case "a":
                                v15 = v15
                            }
                            for fi22 := 0; fi22 < 16; fi22++ {
                                v6 -= (v3 << (((v2 % (((v2 << ((v2) & 63))) | 1))) & 63))
                                v23 := GenS2{F0: (((v2 % ((v2) | 1)) >= (^(v2))) || false), F1: "foo bar"}
                                _ = fi22
                                _ = v23
                            }
                            v5 = (v5 << (((^(v2))) & 63))
                        }
                        _ = v19
                    }
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v25 := []int{1, (-(v4)), (fi24 & fi16)}
                        v6 = (v3 >> ((v2) & 63))
                        var v26 int8 = 1
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
                for fi27 := 0; fi27 < 13; fi27++ {
                    v3 = (v6 | v3)
                    v28 := GenS2{F0: (!v11), F1: "foo bar"}
                    _ = fi27
                    _ = v28
                }
                v15 -= (v15 - 2)
            }
            _ = fi16
            _ = v17
        }
        v13 = (v5 + v5)
    }
    var v29 uint64 = uint64(v2)
    if true {
        v30 := GenS3{F0: -1, F1: (float32(-2.25) <= float32(1e10))}
        _ = v30
    }
    v29 -= 2
    for fi31 := 0; fi31 < 3; fi31++ {
        v32 := GenS3{F0: 100, F1: v11}
        var v33 uint32 = uint32(v29)
        v6 += (v6 / ((v6) | 1))
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(v11, v3)
        _ = fi34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        if (v2 == v2) {
            v36 := v1[4294967295]
            switch "a" {
            case "hello":
                v5 -= 413
            case "":
                if v11 {
                    delete(v1, 501)
                    switch -9223372036854775808 {
                    case 32767:
                        for fi37 := 0; fi37 < 4; fi37++ {
                            v13 -= (fi35 + -837)
                            v1[2147483647] = true
                            _ = fi37
                        }
                    default:
                        for fi38 := 0; fi38 < 15; fi38++ {
                            if v36 {
                                v39 := []int{(v5 >> ((v2) & 63)), fi38, fi38, v4}
                                v36 = (uint16(v15) != 95)
                                v14 += (fi35 + (v4 / ((v14) | 1)))
                                _ = v39
                            } else {
                                if (42 <= 1) {
                                    var v40 int64 = 42
                                    var v41 float64 = ((0.1 / -2.25) / 0.1)
                                    _ = v40
                                    _ = v41
                                }
                                v42 := GenS4{F0: 217}
                                fmt.Println(v13, v11)
                                _ = v42
                            }
                            v1[0] = false
                            _ = fi38
                        }
                    }
                    v2 += v2
                }
            }
            _ = v36
        } else {
            var v43 int8 = 42
            _ = v43
        }
        v44 := GenS3{F0: 100, F1: ((("a" + "hello") + "") == ("foo bar" + "foo bar"))}
        v2 = (^(v2))
        _ = fi35
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
