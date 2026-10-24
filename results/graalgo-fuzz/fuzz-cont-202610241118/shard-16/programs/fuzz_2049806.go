package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 string; F1 bool; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int8 }
func genFn0() int8 {
    var v7 bool = (((false && (!false)) && true) && (507 != 2))
    switch "a" {
    case "hello":
        for fi8 := 0; fi8 < 16; fi8++ {
            v7 = (1 < -47)
            _ = fi8
        }
    case "":
        v9 := []int{1, 32767}
        _ = v9
    case "foo bar":
        if v7 {
            var v10 int32 = -467
            _ = v10
        }
    default:
        if true {
            var v11 int16 = -1
            for fi12 := 0; fi12 < 14; fi12++ {
                switch (fi12 >> ((18446744073709551615) & 63)) {
                case -159:
                    v13 := []int{(fi12 / ((fi12) | 1)), 127}
                    _ = v13
                case -9223372036854775808:
                    var v14 float64 = 1e10
                    _ = v14
                default:
                    v11 = v11
                }
                v15 := []int{234}
                _ = fi12
                _ = v15
            }
            _ = v11
        }
    }
    _ = v7
    return 100
}
func genFn1() uint32 {
    switch 1 {
    case -442:
        v16 := GenS1{F0: 42, F1: 0, F2: -1}
        _ = v16
    case 98:
        v17 := GenS2{F0: ("a" + "hello"), F1: (!true), F2: false}
        _ = v17
    default:
        var v18 string = ("a" + ("\n\t" + ""))
        _ = v18
    }
    var v19 uint16 = 798
    v19 += v19
    fmt.Println(v19)
    fmt.Println(v19)
    _ = v19
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 1
    var v3 int64 = 1
    var v4 uint = v2
    var v5 uint = (v4 / (((^(v2))) | 1))
    v6 := []string{"hello", "\n\t", "foo bar", ("" + (("hello" + "hello") + "\n\t"))}
    var v20 bool = true
    v5 = v4
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            v23 := v1[(fi22 | 0)]
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v20 = (false && (570 != genFn1()))
    v2 -= (v2 - 4294967295)
    if v20 {
        v4 -= (v2 >> ((1) & 63))
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 4; fi25++ {
                v2 = (^(v4))
                if len(v6) > 0 {
                    v6[0] = "foo bar"
                }
                if (genFn0() >= 42) {
                    v4 = (v2 % ((v4) | 1))
                    var v26 int32 = -2147483648
                    var v27 uint = (v4 ^ 0)
                    v28 := GenS1{F0: (v2 % ((v27) | 1)), F1: v26, F2: v3}
                    v26 = (v26 - v26)
                    fmt.Println(v27, fi24)
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    var v29 uint16 = 65535
                    switch 0 {
                    case 128:
                        var v30 uint8 = 153
                        _ = v30
                    case -9223372036854775808:
                        var v31 int64 = (^(v3))
                        _ = v31
                    }
                    fmt.Println(v20, fi24, v5)
                    _ = v29
                }
                var v32 int64 = 42
                var v33 uint32 = 4294967295
                if v20 {
                    switch fi24 {
                    case 547:
                        var v34 uint = v5
                        _ = v34
                    case 4294967295:
                        if v20 {
                            v3 += v3
                            v35 := v1[-1]
                            _ = v35
                        }
                    }
                } else {
                    if len(v6) > 0 {
                        v6[0] = (("a" + "a") + ("" + ("a" + "\n\t")))
                    }
                    if ((false || true) && v20) {
                        v32 = 42
                        switch (("hello" + ("hello" + "foo bar")) + "hello") {
                        case "foo bar":
                            var v36 uint8 = 146
                            _ = v36
                        }
                    } else {
                        if v20 {
                            var v37 int32 = 0
                            v38 := GenS2{F0: (("hello" + "") + "foo bar"), F1: false, F2: (v20 || (-925 >= (fi24 | fi25)))}
                            _ = v37
                            _ = v38
                        } else {
                            v4 -= (v4 * (v2 ^ v4))
                        }
                        var v39 int16 = 0
                        _ = v39
                    }
                }
                _ = fi25
                _ = v32
                _ = v33
            }
            _ = fi24
        }
    }
    var v40 float32 = 3.14159
    v5 = 0
    if (v20 || ((!(31 < 43)) && false)) {
        var v41 uint = v5
        v42 := GenS3{F0: -633}
        _ = v41
        _ = v42
    } else {
        for fi43 := 0; fi43 < 15; fi43++ {
            _ = len(v6)
            v44 := GenS4{F0: (false && ((1.5 - float64((v40 / v40))) != 3.14159)), F1: (!((-(fi43)) == (-(fi43))))}
            _ = fi43
            _ = v44
        }
        v45 := GenS3{F0: 7}
        fmt.Println(v20)
        _ = v45
    }
    v6 = append(v6, "\n\t")
    v3 += int64(v4)
    fmt.Println(v2, v2, v40)
    v20 = v20
    for fi46 := 0; fi46 < 4; fi46++ {
        for fi47 := 0; fi47 < 15; fi47++ {
            var v48 bool = true
            if v48 {
                var v49 uint32 = 4294967295
                var v50 uint16 = 1
                v51 := GenS5{F0: genFn0()}
                _ = v49
                _ = v50
                _ = v51
            }
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
