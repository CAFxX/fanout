package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0() int32 {
    if true {
        var v7 float32 = (((0.0 + float32(1.5)) - 3.14159) - 1.5)
        _ = v7
    } else {
        v8 := []int{1, 32768}
        var v9 int8 = 42
        var v10 bool = true
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint8 = 255
    _ = v11
    return 42
}
func genFn1() uint16 {
    if (!false) {
        var v12 bool = (!(1 < 84))
        _ = v12
    }
    v13 := []int{127, 32768, 127}
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint16 = 65535
        v16 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int64 = -1
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            v17 -= 720
            _ = fi19
        }
        _ = fi18
    }
    _ = v13
    _ = v17
    return 832
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 int8 = -1
    var v4 uint16 = 181
    var v5 uint32 = 2147483647
    v6 := make(map[string]bool)
    v20 := GenS1{F0: 9223372036854775807, F1: 0}
    fmt.Println(v3)
    fmt.Println(v3)
    for fi21 := 0; fi21 < 14; fi21++ {
        delete(v6, "\n\t")
        v22 := []int{fi21}
        _ = fi21
        _ = v22
    }
    switch 127 {
    case 32767:
        v23 := []int{128, 32768}
        _ = v23
    default:
        v24 := []int{}
        _ = v24
    }
    switch ("a" + "foo bar") {
    case "a":
        fmt.Println(v3)
    case "foo bar":
        fmt.Println(v1)
    default:
        var v25 float32 = 0.0
        _ = v25
    }
    v26 := []int{int(v3), 128, -1}
    var v27 float32 = float32((0.1 * 0.1))
    if len(v26) > 0 {
        v26[0] = 42
    }
    var v28 float64 = 1e10
    v6[(((("\n\t" + ("" + "\n\t")) + "a") + "hello") + ("hello" + ((("" + ("" + "hello")) + "\n\t") + "a")))] = true
    var v29 uint64 = 1
    v30 := GenS2{F0: true}
    var v31 int16 = 773
    v6["a"] = true
    v28 -= 1.5
    if (!(!(-649 > 0))) {
        var v32 string = "a"
        v1 = (127 < 2147483647)
        if ((v29 >> ((275) & 63)) != v29) {
            for fi33 := 0; fi33 < 12; fi33++ {
                v4 = v4
                switch (^(fi33)) {
                case -9223372036854775808:
                    if (210 <= 42) {
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        v1 = (!(genFn0() != 0))
                        for fi34 := 0; fi34 < 10; fi34++ {
                            if (((-1 >= 0) && v1) || (v27 <= (v27 / (v27 / v27)))) {
                                var v35 float64 = v28
                                _ = v35
                            }
                            if (int16(v5) == (^(v31))) {
                                fmt.Println(v5)
                                v1 = false
                            }
                            var v36 int64 = -794
                            _ = fi34
                            _ = v36
                        }
                    }
                case 32767:
                    for fi37 := 0; fi37 < 5; fi37++ {
                        fmt.Println(v3, v29)
                        var v38 float64 = (v28 + (0.1 + (-2.25 + (v28 / ((float64(-2.25) / v28) + (v28 / (-2.25 / v28)))))))
                        var v39 uint64 = (v29 + 679)
                        v38 = v38
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                case 127:
                    for fi40 := 0; fi40 < 17; fi40++ {
                        if ((2 != 183) && v1) {
                            v41 := GenS3{F0: (v3 % ((100) | 1)), F1: v4, F2: -145}
                            _ = v41
                        }
                        v42 := GenS3{F0: v3, F1: v4, F2: v31}
                        _ = fi40
                        _ = v42
                    }
                default:
                    v31 -= 180
                }
                for fi43 := 0; fi43 < 16; fi43++ {
                    v44 := GenS4{F0: 255, F1: (v31 % (((v31 | v31)) | 1))}
                    v45 := GenS1{F0: -1, F1: -146}
                    v46 := GenS3{F0: (v3 ^ 0), F1: 515, F2: (v31 - v31)}
                    if (!(v1 && v1)) {
                        v4 -= v4
                        fmt.Println(fi33, v1)
                        var v47 uint64 = 266
                        _ = v47
                    }
                    _ = len(v2)
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
                _ = fi33
            }
        } else {
            v1 = ("a" <= ("\n\t" + v32))
            _ = len(v26)
        }
        _ = v32
    } else {
        if (false && (!v1)) {
            for fi48 := 0; fi48 < 15; fi48++ {
                v4 = genFn1()
                fmt.Println(v3)
                _ = fi48
            }
            if len(v26) > 0 {
                v26[0] = 128
            }
            fmt.Println(v1, v4, v29)
        } else {
            _ = len(v2)
            var v49 uint16 = 770
            _ = v49
        }
        switch (("foo bar" + ("\n\t" + "hello")) + ("" + "")) {
        case "hello":
            var v50 int16 = v31
            _ = v50
        case "a":
            v51 := GenS1{F0: 9223372036854775807, F1: -656}
            _ = v51
        case "":
            var v52 int8 = v3
            _ = v52
        default:
            fmt.Println(v31)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
