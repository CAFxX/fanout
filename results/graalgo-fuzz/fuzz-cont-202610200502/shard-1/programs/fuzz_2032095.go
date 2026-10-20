package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int16; F2 string }
func genFn0(p0 string, p1 int) bool {
    v6 := []int{p1, p1, (p1 + (^(p1))), (p1 | 2147483647)}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint16 = 65535
        var v9 int64 = 952
        v10 := GenS1{F0: true, F1: 457}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        p0 = p0
        _ = fi11
    }
    var v12 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return false
}
func genFn1(p0 uint16) int32 {
    if false {
        v13 := GenS2{F0: "\n\t"}
        p0 += 0
        v14 := GenS1{F0: true, F1: 65535}
        _ = v13
        _ = v14
    }
    p0 += (p0 - p0)
    var v15 uint8 = 87
    var v16 int16 = -437
    _ = p0
    _ = v15
    _ = v16
    return -830
}

func main() {
    var v1 uint32 = 2147483647
    var v2 bool = true
    v3 := make(map[uint32]int)
    v4 := []int{32768, -1}
    v5 := []int{0, -62, 825}
    var v17 bool = v2
    v2 = v2
    v2 = true
    fmt.Println(v2, v1)
    switch "" {
    case "a":
        if true {
            v18 := GenS1{F0: (0 != -512), F1: 0}
            v2 = ("\n\t" == "")
            _ = v18
        } else {
            var v19 uint8 = 1
            v2 = v2
            _ = v19
        }
    case "foo bar":
        var v20 int = 42
        _ = v20
    }
    v2 = ((((0.0 / 0.1) * 1.5) / float64(0.1)) == ((0.1 - (0.1 - 0.1)) - 3.14159))
    switch 32767 {
    case 32768:
        for fi21 := 0; fi21 < 20; fi21++ {
            delete(v3, (v1 - (v1 % (((v1 + (v1 * v1))) | 1))))
            _ = fi21
        }
    }
    v2 = true
    if ((v2 && ((v2 && true) || (false && true))) || v17) {
        v1 = (v1 + (v1 | (v1 | (v1 << ((1) & 63)))))
        var v22 string = "foo bar"
        var v23 int64 = -985
        _ = v22
        _ = v23
    } else {
        if false {
            var v24 float64 = -2.25
            _ = v24
        }
        if v17 {
            v1 -= 2147483648
        }
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v3[0] = (-(fi25))
        v1 -= 1
        _ = fi25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        fmt.Println(fi26)
        _ = fi26
    }
    if (("\n\t" + "hello") != "hello") {
        if (v17 || (!v2)) {
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
            v2 = ((42 > 42) || (v2 || v17))
            v1 = 2147483648
        } else {
            v17 = (32767 != 255)
            var v27 int8 = 122
            v2 = (1 <= -32768)
            _ = v27
        }
    }
    if v17 {
        if false {
            v28 := GenS1{F0: ((3.14159 / (float32((1e10 / -2.25)) + float32(1.5))) >= (1e10 / float32(float64(float32((1e10 / ((((((0.1 + 0.1) * 0.1) / 1.0) * 0.1) * (float64(1e10) * 0.0)) - -2.25))))))), F1: 42}
            var v29 int = 200
            if ((v17 || (!v2)) && v17) {
                switch ("a" + "") {
                case "foo bar":
                    for fi30 := 0; fi30 < 3; fi30++ {
                        var v31 uint = 65535
                        fmt.Println(fi30, v31)
                        _ = fi30
                        _ = v31
                    }
                case "a":
                    v17 = v2
                default:
                    var v32 float32 = ((1e10 + ((-2.25 + 0.0) * (((0.0 - ((3.14159 / 0.1) - (3.14159 * 1e10))) / float32((0.1 - 3.14159))) / (-2.25 + (0.1 / 1.5))))) / 3.14159)
                    _ = v32
                }
            }
            _ = v28
            _ = v29
        } else {
            for fi33 := 0; fi33 < 7; fi33++ {
                switch -673 {
                case 610:
                    for fi34 := 0; fi34 < 20; fi34++ {
                        v35 := GenS3{F0: 84, F1: 523, F2: 1}
                        fmt.Println(v1)
                        v5 = append(v5, (fi34 >> ((65535) & 63)))
                        if (-2147483648 != fi34) {
                            v1 -= v1
                            fmt.Println(fi33, fi34, v17)
                            v17 = v2
                        }
                        v36 := v3[0]
                        var v37 string = ("\n\t" + "a")
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                case 4294967295:
                    v1 = 65535
                case -423:
                    for fi38 := 0; fi38 < 12; fi38++ {
                        fmt.Println(v2, v2)
                        v5 = append(v5, fi33)
                        v2 = ((2 <= 1) || v17)
                        _ = fi38
                    }
                default:
                    v39 := GenS4{F0: 895, F1: 1000, F2: "foo bar"}
                    _ = v39
                }
                _ = fi33
            }
            for fi40 := 0; fi40 < 15; fi40++ {
                _ = len(v4)
                _ = fi40
            }
        }
        v17 = v2
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 float64 = -2.25
        v1 -= v1
        if v17 {
            var v43 int16 = 0
            _ = v43
        }
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
