package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 string }
type GenS2 struct { F0 string; F1 int32; F2 bool }
func genFn0(p0 int64, p1 int64, p2 string) int {
    for fi4 := 0; fi4 < 20; fi4++ {
        if (false || true) {
            var v5 uint64 = 2
            _ = v5
        }
        if true {
            var v6 uint8 = 50
            v7 := GenS1{F0: 4294967295, F1: fi4, F2: "foo bar"}
            _ = v6
            _ = v7
        } else {
            var v8 uint64 = 1
            switch (fi4 / ((fi4) | 1)) {
            case -2147483648:
                p1 += p1
            default:
                for fi9 := 0; fi9 < 8; fi9++ {
                    v8 -= v8
                    _ = fi9
                }
            }
            p1 += (p0 << ((255) & 63))
            _ = v8
        }
        v10 := GenS1{F0: 0, F1: (fi4 ^ (fi4 - fi4)), F2: p2}
        _ = fi4
        _ = v10
    }
    var v11 int16 = 1000
    var v12 int8 = -128
    p0 += 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 127
}
func genFn1(p0 int32, p1 uint, p2 int16) uint16 {
    var v13 uint8 = 2
    fmt.Println(v13, p2)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 397
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = 0.1
    var v3 uint32 = 34
    v14 := []int{127, 2147483647, -825, 889}
    if (0.1 <= (1e10 * v2)) {
        v14 = append(v14, -2147483648)
        v2 += v2
        switch "foo bar" {
        case "":
            switch -837 {
            case 9223372036854775807:
                v15 := GenS2{F0: ((("\n\t" + ("\n\t" + "")) + "a") + (v1 + ((("hello" + v1) + v1) + ((("" + (("" + v1) + (v1 + ("hello" + "")))) + "") + v1)))), F1: -299, F2: true}
                _ = v15
            case -2147483648:
                v2 += v2
            }
        case "a":
            v3 += (v3 ^ 1)
        case "\n\t":
            v16 := GenS1{F0: 18446744073709551615, F1: 32768, F2: v1}
            _ = v16
        }
    } else {
        var v17 float64 = 0.0
        for fi18 := 0; fi18 < 15; fi18++ {
            fmt.Println(v1)
            var v19 int16 = 0
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v1 = ""
            _ = fi21
        }
        if ((v1 + "a") != "\n\t") {
            fmt.Println(v2, fi20)
        } else {
            if false {
                v1 = "hello"
                switch 843 {
                case 56:
                    if (42 >= 100) {
                        if (42 < 2147483647) {
                            if (!((true || true) && false)) {
                                var v22 string = v1
                                _ = v22
                            } else {
                                var v23 uint = 1
                                _ = v23
                            }
                            v2 += (v2 * v2)
                            var v24 bool = false
                            _ = v24
                        } else {
                            fmt.Println(v1, v1, v2)
                            v25 := GenS1{F0: 0, F1: (^(fi20)), F2: v1}
                            _ = v25
                        }
                        var v26 float64 = float64((v2 + float32(3.14159)))
                        v1 = v1
                        _ = v26
                    }
                default:
                    for fi27 := 0; fi27 < 10; fi27++ {
                        var v28 uint64 = 2
                        var v29 uint16 = 2
                        v30 := GenS1{F0: v28, F1: (fi20 * fi20), F2: "a"}
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
                var v31 uint32 = (^(v3))
                _ = v31
            }
            var v32 int32 = -2147483648
            _ = v32
        }
        _ = fi20
    }
    var v33 bool = false
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v34 float64 = 3.14159
    var v35 float64 = float64(v2)
    v14 = append(v14, 128)
    if ((v3 >> ((639) & 63)) <= (v3 * v3)) {
        var v36 int = 2147483647
        var v37 int32 = -1
        v38 := GenS2{F0: "\n\t", F1: 1, F2: v33}
        _ = v36
        _ = v37
        _ = v38
    } else {
        var v39 int = 32768
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
