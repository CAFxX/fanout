package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 string, p1 bool) uint16 {
    var v7 float32 = 0.1
    var v8 uint64 = 927
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32767
}
func genFn1() uint64 {
    var v9 uint64 = 942
    var v10 int64 = int64(v9)
    _ = v10
    _ = v9
    return (^(v9))
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{0, 128, 952}
    var v3 float32 = (1.5 / 1.0)
    var v4 int32 = -2147483648
    v5 := []int{32767, 2147483647}
    v6 := make(map[uint32]uint32)
    v11 := GenS1{F0: 1, F1: 0}
    if (0 < 255) {
        v12 := GenS1{F0: 1, F1: 4294967295}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 6; fi14++ {
                v5 = append(v5, fi13)
                _ = fi14
            }
            v3 += v3
            var v15 bool = (false && (!true))
            _ = fi13
            _ = v15
        }
    }
    switch int(v4) {
    case 202:
        if true {
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 int32 = (v4 ^ (v4 << ((0) & 63)))
                for fi18 := 0; fi18 < 1; fi18++ {
                    fmt.Println(v4)
                    var v19 uint16 = 255
                    _ = fi18
                    _ = v19
                }
                switch fi16 {
                case 549:
                    for fi20 := 0; fi20 < 14; fi20++ {
                        fmt.Println(v4, v4)
                        _ = fi20
                    }
                default:
                    delete(v1, (fi16 << ((255) & 63)))
                }
                _ = fi16
                _ = v17
            }
            _ = len(v2)
            switch 128 {
            case -518:
                var v21 bool = ((0 <= 350) && (!(!((-1 > -250) || (128 >= 42)))))
                _ = v21
            case 822:
                var v22 int = -11
                _ = v22
            }
        } else {
            fmt.Println(v3, v4)
            v4 = (v4 << ((255) & 63))
        }
    default:
        v3 = v3
    }
    var v23 int8 = 0
    var v24 int64 = 9223372036854775807
    var v25 int8 = (v23 >> ((255) & 63))
    switch -530 {
    case 32767:
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := GenS1{F0: 26, F1: genFn1()}
            if false {
                var v28 uint8 = 1
                _ = v28
            } else {
                switch (fi26 - 4294967295) {
                case -654:
                    _ = len(v1)
                default:
                    v29 := GenS1{F0: 2147483647, F1: 81}
                    _ = v29
                }
                v30 := v1[fi26]
                _ = v30
            }
            _ = fi26
            _ = v27
        }
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if true {
            _ = len(v2)
            for fi32 := 0; fi32 < 5; fi32++ {
                var v33 uint = 305
                var v34 int16 = -1
                _ = fi32
                _ = v33
                _ = v34
            }
        } else {
            v6[4294967295] = 42
        }
        v25 = 36
        _ = fi31
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        for fi36 := 0; fi36 < 1; fi36++ {
            v4 = -1
            _ = fi36
        }
        if true {
            v37 := GenS1{F0: 1, F1: 18446744073709551615}
            if (!((fi35 ^ (fi35 & (fi35 & 286))) < (fi35 >> ((42) & 63)))) {
                for fi38 := 0; fi38 < 9; fi38++ {
                    v39 := v1[314]
                    v3 -= (float32(-2.25) - v3)
                    _ = fi38
                    _ = v39
                }
            }
            if len(v5) > 0 {
                v5[0] = fi35
            }
            _ = v37
        } else {
            for fi40 := 0; fi40 < 7; fi40++ {
                v41 := GenS2{F0: ""}
                switch "\n\t" {
                case "a":
                    switch fi35 {
                    case 127:
                        v4 = -1
                    case 1:
                        v4 = v4
                    default:
                        if ((("foo bar" + ("a" + ((("hello" + "hello") + "foo bar") + "a"))) + "a") == "a") {
                            var v42 float64 = -2.25
                            _ = v42
                        }
                    }
                case "foo bar":
                    if (!false) {
                        var v43 uint32 = 276
                        if (false && false) {
                            var v44 string = ((("" + "\n\t") + ("a" + "a")) + "hello")
                            _ = v44
                        } else {
                            v45 := GenS3{F0: 127, F1: 1}
                            if (true && true) {
                                v3 = v3
                                v46 := GenS3{F0: 1, F1: v4}
                                _ = v46
                            }
                            _ = v45
                        }
                        _ = v43
                    }
                case "\n\t":
                    fmt.Println(v3, v3, v3)
                }
                _ = fi40
                _ = v41
            }
            if len(v5) > 0 {
                v5[0] = fi35
            }
            if false {
                fmt.Println(v25, v23)
                var v47 int8 = (v25 << ((18446744073709551615) & 63))
                _ = v47
            }
        }
        for fi48 := 0; fi48 < 10; fi48++ {
            if false {
                var v49 string = "hello"
                _ = v49
            }
            v50 := v6[617]
            fmt.Println(v25, fi35)
            _ = fi48
            _ = v50
        }
        _ = fi35
    }
    v25 += 41
    var v51 uint8 = 255
    if false {
        var v52 bool = ((v4 ^ (v4 - v4)) <= v4)
        _ = v52
    }
    for fi53 := 0; fi53 < 6; fi53++ {
        _ = len(v2)
        var v54 bool = (v51 == v51)
        _ = fi53
        _ = v54
    }
    delete(v6, 2147483648)
    fmt.Println(v23, v23, v24)
    v55 := GenS3{F0: v51, F1: -2147483648}
    var v56 float64 = (0.0 + 3.14159)
    v3 = (float32(0.0) - float32((float64(v3) * (v56 - -2.25))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v55
    _ = v56
    _ = v6
}
