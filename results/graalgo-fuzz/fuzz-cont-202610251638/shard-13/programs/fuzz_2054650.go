package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 bool; F1 uint; F2 int }
type GenS5 struct { F0 int64; F1 uint; F2 uint16 }
func genFn0(p0 uint32, p1 int32) int {
    if (true && (!(!true))) {
        p0 = p0
        fmt.Println(p0)
    }
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 255, F1: 635}
        _ = fi4
        _ = v5
    }
    if false {
        p0 = 0
        switch 0 {
        case 0:
            p0 += p0
        case -646:
            fmt.Println(p1, p1)
        case 42:
            fmt.Println(p0)
        }
    } else {
        if false {
            var v6 float64 = 0.1
            var v7 uint8 = 42
            v6 = 1e10
            _ = v6
            _ = v7
        }
        switch -885 {
        case 42:
            var v8 int32 = 1
            _ = v8
        case 263:
            fmt.Println(p1, p0)
        }
        var v9 uint = 0
        _ = v9
    }
    v10 := GenS2{F0: (p1 & (p1 | p1))}
    if false {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint32 = (p0 % (((p0 | p0)) | 1))
            v13 := []int{}
            if true {
                v12 = (p0 - v12)
            } else {
                if false {
                    if ("foo bar" < "") {
                        for fi14 := 0; fi14 < 1; fi14++ {
                            switch 127 {
                            case 127:
                                var v15 int64 = -784
                                _ = v15
                            default:
                                switch ((("foo bar" + "\n\t") + "") + "a") {
                                case "":
                                    switch (fi11 | fi14) {
                                    case 42:
                                        fmt.Println(p0, fi11, fi11)
                                    case 464:
                                        if (4 > 127) {
                                            var v16 uint = 241
                                            _ = v16
                                        } else {
                                            v12 = 104
                                            var v17 bool = true
                                            _ = v17
                                        }
                                    case 128:
                                        p1 += 493
                                    }
                                default:
                                    fmt.Println(p0)
                                }
                            }
                            _ = fi14
                        }
                    }
                    v13 = append(v13, (^(fi11)))
                } else {
                    v13 = append(v13, (fi11 * fi11))
                    switch int(p1) {
                    case 127:
                        switch -9223372036854775808 {
                        case 127:
                            var v18 int = (fi11 / ((32768) | 1))
                            _ = v18
                        case 42:
                            var v19 int16 = -1
                            _ = v19
                        case 0:
                            fmt.Println(p0, fi11)
                        default:
                            v13 = append(v13, 1)
                        }
                    }
                }
                fmt.Println(p0, fi11)
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        var v20 int8 = 1
        switch ("foo bar" + (("hello" + (("\n\t" + "\n\t") + ("\n\t" + "hello"))) + "\n\t")) {
        case "":
            switch "" {
            case "foo bar":
                if false {
                    if (-9223372036854775808 == -1) {
                        v21 := GenS1{F0: 4294967295, F1: 448}
                        _ = v21
                    }
                } else {
                    var v22 string = ("a" + "foo bar")
                    _ = v22
                }
            case "a":
                if true {
                    var v23 uint8 = 127
                    var v24 int = 42
                    _ = v23
                    _ = v24
                }
            case "\n\t":
                if false {
                    var v25 int32 = (p1 | p1)
                    v26 := []int{32767, 32767, 4294967295, 9223372036854775807}
                    _ = v25
                    _ = v26
                }
            }
        }
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v10
    return -382
}
func genFn1() uint8 {
    var v27 int64 = 555
    v27 -= (v27 / (((v27 + (v27 >> ((255) & 63)))) | 1))
    v27 -= (v27 | (v27 | (v27 % (((v27 | v27)) | 1))))
    _ = v27
    return 0
}

func main() {
    var v1 int8 = 1
    var v2 uint16 = 887
    var v3 int8 = (v1 - v1)
    if (true && false) {
        v1 -= (v1 | v1)
        switch genFn0(992, 1) {
        case 2147483647:
            fmt.Println(v2, v2, v3)
        case -28:
            for fi28 := 0; fi28 < 7; fi28++ {
                fmt.Println(v3)
                _ = fi28
            }
        case 160:
            for fi29 := 0; fi29 < 19; fi29++ {
                v1 -= 0
                var v30 int64 = -9223372036854775808
                _ = fi29
                _ = v30
            }
        default:
            var v31 uint64 = 1
            _ = v31
        }
    }
    v1 = v1
    for fi32 := 0; fi32 < 6; fi32++ {
        if true {
            v2 -= v2
            v33 := GenS3{F0: 42, F1: 1}
            _ = v33
        }
        v2 += 65535
        _ = fi32
    }
    fmt.Println(v1, v2, v1)
    v34 := GenS4{F0: (2147483647 > 2), F1: 716, F2: -2147483648}
    v35 := GenS5{F0: 1, F1: 65535, F2: v2}
    fmt.Println(v3)
    if false {
        var v36 float32 = -2.25
        for fi37 := 0; fi37 < 1; fi37++ {
            v3 -= (^(v3))
            _ = fi37
        }
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
