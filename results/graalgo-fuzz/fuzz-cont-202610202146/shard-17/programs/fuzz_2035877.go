package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int32; F2 string }
func genFn0(p0 bool, p1 int16) string {
    var v6 int16 = p1
    v7 := []int{0}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 32767
        _ = fi8
        _ = v9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ("\n\t" + ("" + "\n\t"))
}
func genFn1() uint8 {
    v10 := GenS1{F0: 0}
    for fi11 := 0; fi11 < 19; fi11++ {
        if true {
            v12 := GenS1{F0: 0}
            _ = v12
        } else {
            switch 127 {
            case 128:
                var v13 int = (fi11 * (fi11 + (fi11 % (((fi11 >> ((448) & 63))) | 1))))
                _ = v13
            case -2147483648:
                fmt.Println(fi11)
            default:
                var v14 float64 = -2.25
                _ = v14
            }
            for fi15 := 0; fi15 < 2; fi15++ {
                fmt.Println(fi15)
                var v16 float64 = 3.14159
                v16 = v16
                _ = fi15
                _ = v16
            }
        }
        var v17 uint8 = 2
        _ = fi11
        _ = v17
    }
    var v18 bool = (true && false)
    _ = v10
    _ = v18
    return 0
}

func main() {
    var v1 int64 = 0
    v2 := []string{"hello", "\n\t", "foo bar"}
    v3 := make(map[uint32]string)
    var v4 bool = (4294967295 == 456)
    var v5 uint32 = 1
    if v4 {
        v5 = v5
        for fi19 := 0; fi19 < 3; fi19++ {
            fmt.Println(fi19, v5)
            if (-2.25 >= (3.14159 + 3.14159)) {
                if (!v4) {
                    v4 = v4
                    var v20 int64 = v1
                    _ = v20
                } else {
                    var v21 uint = 2
                    v22 := v3[v5]
                    _ = v21
                    _ = v22
                }
                if (!(v4 || ((!v4) || (508 >= uint16(v5))))) {
                    if ((255 != 0) || (true && v4)) {
                        var v23 uint32 = (v5 | v5)
                        _ = len(v3)
                        _ = v23
                    } else {
                        var v24 string = "foo bar"
                        v5 = v5
                        _ = v24
                    }
                    v4 = true
                    v1 = (v1 + 1)
                }
                for fi25 := 0; fi25 < 8; fi25++ {
                    var v26 int8 = -99
                    switch "\n\t" {
                    case "\n\t":
                        delete(v3, 4294967295)
                    default:
                        v26 -= 1
                    }
                    v4 = (!(0.0 == (1e10 / (0.0 - -2.25))))
                    _ = fi25
                    _ = v26
                }
            }
            _ = fi19
        }
    }
    v5 -= (v5 & (v5 % ((2) | 1)))
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 bool = (((1e10 - 1e10) * -2.25) <= -2.25)
        _ = fi27
        _ = v28
    }
    v3[(v5 - v5)] = "a"
    v29 := GenS1{F0: -108}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v30 := GenS2{F0: 0}
    var v31 float32 = (-2.25 - (1.5 / float32((float64(-2.25) - 0.1))))
    switch "hello" {
    case "a":
        if v4 {
            if v4 {
                var v32 float64 = 3.14159
                _ = v32
            } else {
                v2 = append(v2, "\n\t")
                v1 += (v1 ^ (^(v1)))
                var v33 string = (("hello" + "foo bar") + ("a" + "hello"))
                _ = v33
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v34 uint = 4294967295
            _ = v34
        }
    case "\n\t":
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 float32 = 0.1
            fmt.Println(v5, v1, v36)
            v5 -= (^(v5))
            _ = fi35
            _ = v36
        }
    default:
        if ("\n\t" == ("" + "a")) {
            switch -706 {
            case 208:
                v37 := GenS3{F0: 0, F1: 387, F2: ""}
                _ = v37
            case -174:
                v4 = (18446744073709551615 >= 2)
            case 0:
                fmt.Println(v5, v31, v1)
            }
            fmt.Println(v5, v1)
            _ = len(v2)
        } else {
            fmt.Println(v31)
            v38 := GenS3{F0: 127, F1: 35, F2: "foo bar"}
            fmt.Println(v4, v4)
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
