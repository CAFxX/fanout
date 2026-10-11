package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int8; F1 int16 }
func genFn0() bool {
    var v4 int64 = 1
    switch 127 {
    case 4294967295:
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := GenS1{F0: (fi5 << ((0) & 63))}
            for fi7 := 0; fi7 < 1; fi7++ {
                if false {
                    v8 := GenS1{F0: int(v4)}
                    if false {
                        v9 := []int{fi7, (fi5 % (((fi5 - fi7)) | 1))}
                        if (false && true) {
                            fmt.Println(fi7)
                            var v10 int16 = 42
                            _ = v10
                        }
                        _ = v9
                    }
                    _ = v8
                } else {
                    v11 := GenS2{F0: fi7}
                    v4 = (v4 & 42)
                    var v12 float64 = (((1.5 + (0.0 / ((float64(1e10) - ((((-2.25 / -2.25) * (0.1 - 0.0)) * 3.14159) + (-2.25 * (1.5 - 0.0)))) * 0.1))) + ((0.1 / (float64(1.5) - -2.25)) + 1.5)) * ((((1.5 + (3.14159 - 0.0)) + 1.5) - ((1e10 * 0.0) * 1e10)) - -2.25))
                    _ = v11
                    _ = v12
                }
                if false {
                    if (((false && (1e10 < -2.25)) || (!(!(42 == -172)))) && (255 == 42)) {
                        v13 := GenS1{F0: -9223372036854775808}
                        fmt.Println(v4)
                        var v14 int64 = 1
                        _ = v13
                        _ = v14
                    } else {
                        if (true || (true && (!true))) {
                            if (42 >= 42) {
                                v4 -= v4
                                var v15 string = ("\n\t" + (("foo bar" + ("hello" + "\n\t")) + (("foo bar" + ("" + "hello")) + ("a" + ("\n\t" + "a")))))
                                v16 := []int{fi7, -265, fi7}
                                _ = v15
                                _ = v16
                            }
                            if (v4 == (^(v4))) {
                                v17 := []int{(fi7 | fi7), (fi7 >> ((65535) & 63)), 32768}
                                var v18 uint32 = 340
                                _ = len(v17)
                                v19 := GenS2{F0: (fi7 << ((42) & 63))}
                                v20 := GenS3{F0: 127}
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        } else {
                            fmt.Println(v4)
                        }
                        fmt.Println(v4)
                        v21 := GenS4{F0: true, F1: 255}
                        _ = v21
                    }
                    var v22 bool = true
                    _ = v22
                } else {
                    v23 := []int{(fi7 + (fi5 / ((fi5) | 1))), 369, fi5, fi7}
                    fmt.Println(fi5, v4)
                    var v24 int = fi5
                    _ = v23
                    _ = v24
                }
                v25 := []int{(fi5 ^ (fi7 ^ (fi7 | 127))), (fi7 ^ fi7)}
                _ = fi7
                _ = v25
            }
            v26 := GenS5{F0: 127, F1: ("" + "a")}
            _ = fi5
            _ = v26
            _ = v6
        }
    case 1:
        var v27 float64 = 1e10
        _ = v27
    default:
        v28 := []int{-326}
        _ = v28
    }
    fmt.Println(v4)
    _ = v4
    return false
}
func genFn1() int8 {
    switch "a" {
    case "\n\t":
        var v29 int64 = -1
        _ = v29
    default:
        var v30 int64 = 42
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        if false {
            if (true && (2147483647 > 0)) {
                v32 := GenS1{F0: (fi31 * -2147483648)}
                v33 := GenS5{F0: fi31, F1: ""}
                _ = v32
                _ = v33
            }
        }
        _ = fi31
    }
    return -128
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 42
    var v3 int8 = 49
    var v34 int64 = 0
    var v35 int8 = v3
    v3 = 127
    var v36 int = 9223372036854775807
    delete(v1, v36)
    v37 := []int{(-(v36)), v36, 200, 449}
    if false {
        var v38 int8 = v3
        if ((v2 < (^(v2))) || (!false)) {
            v2 -= 1
            var v39 int16 = -947
            for fi40 := 0; fi40 < 12; fi40++ {
                for fi41 := 0; fi41 < 19; fi41++ {
                    v42 := GenS6{F0: v35, F1: (v39 | -323)}
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
            _ = v39
        }
        _ = v38
    }
    fmt.Println(v2, v36, v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
