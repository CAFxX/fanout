package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 uint }
func genFn0(p0 string) int {
    fmt.Println(p0)
    p0 = "a"
    v4 := GenS1{F0: 127, F1: p0, F2: ((("foo bar" + (p0 + p0)) + ((("\n\t" + "") + ("\n\t" + ((p0 + ((p0 + "hello") + p0)) + p0))) + "hello")) + (p0 + p0))}
    v5 := GenS2{F0: 245, F1: p0}
    _ = p0
    _ = v4
    _ = v5
    return 32768
}
func genFn1() uint {
    if (1 > 2147483647) {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 uint16 = 1
            _ = fi6
            _ = v7
        }
    } else {
        var v8 bool = (587 == -1)
        v9 := []int{127, -221, 0, genFn0("hello")}
        _ = v8
        _ = v9
    }
    v10 := []int{genFn0("a")}
    switch -219 {
    case -1:
        for fi11 := 0; fi11 < 15; fi11++ {
            if (!true) {
                if false {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    switch ("" + "foo bar") {
                    case "foo bar":
                        fmt.Println(fi11)
                    case "a":
                        var v12 bool = (((!(!(1e10 <= 3.14159))) && true) || true)
                        _ = v12
                    case "":
                        for fi13 := 0; fi13 < 4; fi13++ {
                            var v14 int64 = -1
                            v15 := GenS3{F0: "a"}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                } else {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        var v17 uint16 = 378
                        fmt.Println(fi11)
                        v17 -= (^(v17))
                        _ = fi16
                        _ = v17
                    }
                    var v18 uint16 = 0
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 int32 = 828
                        _ = len(v10)
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                }
                var v21 int64 = -232
                v21 -= 9223372036854775807
                _ = v21
            } else {
                var v22 float64 = 0.1
                if false {
                    fmt.Println(fi11, fi11)
                } else {
                    v22 = (v22 * 0.0)
                    if false {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        var v23 int64 = -9223372036854775808
                        v24 := GenS4{F0: 46}
                        _ = v23
                        _ = v24
                    } else {
                        var v25 bool = (137 < 42)
                        v10 = append(v10, -416)
                        _ = v25
                    }
                    var v26 string = ("hello" + "hello")
                    _ = v26
                }
                var v27 uint8 = 2
                _ = v22
                _ = v27
            }
            fmt.Println(fi11)
            if true {
                v28 := GenS5{F0: false}
                for fi29 := 0; fi29 < 5; fi29++ {
                    fmt.Println(fi11, fi29)
                    var v30 uint8 = 0
                    _ = fi29
                    _ = v30
                }
                var v31 uint = 255
                _ = v28
                _ = v31
            } else {
                for fi32 := 0; fi32 < 8; fi32++ {
                    var v33 int64 = 42
                    var v34 uint8 = 192
                    v33 = (v33 - (v33 / ((v33) | 1)))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            _ = fi11
        }
    default:
        var v35 int8 = 0
        _ = v35
    }
    _ = v10
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 476
    v3 := make(map[uint32]string)
    var v36 int32 = 660
    switch ("" + "\n\t") {
    case "\n\t":
        v37 := GenS6{F0: false}
        _ = v37
    }
    var v38 bool = (!(!((1 <= 218) && (((-2.25 / 0.1) / -2.25) >= (0.0 / 1e10)))))
    for fi39 := 0; fi39 < 16; fi39++ {
        v36 = v36
        _ = fi39
    }
    if (!false) {
        if (68 >= 127) {
            var v40 uint8 = 0
            v41 := []int{32767, 368, 4294967295, int(v40)}
            _ = v40
            _ = v41
        }
    } else {
        switch 157 {
        case -267:
            v2 += (^(v2))
        }
        v38 = v38
        for fi42 := 0; fi42 < 4; fi42++ {
            fmt.Println(v36)
            fmt.Println(v38, v38)
            _ = fi42
        }
    }
    var v43 int = 127
    v38 = (447 < 1)
    fmt.Println(v43, v38, v38)
    var v44 int16 = 1000
    if v38 {
        v45 := []int{v43, genFn0("hello"), -9223372036854775808, (v43 << ((42) & 63))}
        if false {
            v36 = v36
            var v46 uint64 = 42
            _ = v46
        }
        _ = v45
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        var v48 bool = (0.0 == (0.1 + (float64(0.0) * 1e10)))
        v49 := GenS7{F0: "foo bar", F1: 0}
        _ = fi47
        _ = v48
        _ = v49
    }
    var v50 uint = genFn1()
    var v51 int8 = 17
    if (true || ((!(!v38)) || v38)) {
        var v52 int8 = (v51 - (v51 ^ 81))
        switch "\n\t" {
        case "hello":
            v51 = v51
        case "":
            for fi53 := 0; fi53 < 19; fi53++ {
                var v54 uint32 = uint32(v2)
                _ = fi53
                _ = v54
            }
        case "\n\t":
            fmt.Println(v52, v52, v52)
        default:
            for fi55 := 0; fi55 < 4; fi55++ {
                fmt.Println(v36, v2, v2)
                switch ("foo bar" + "a") {
                case "foo bar":
                    var v56 bool = v38
                    _ = v56
                case "":
                    fmt.Println(v52)
                default:
                    v57 := GenS4{F0: 0}
                    _ = v57
                }
                _ = fi55
            }
        }
        v58 := v1["a"]
        _ = v52
        _ = v58
    } else {
        v43 -= (-(v43))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v43
    _ = v44
    _ = v50
    _ = v51
}
