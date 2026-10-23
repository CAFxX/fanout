package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint16; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int; F2 uint64 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int32; F1 string; F2 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16 }
func genFn0(p0 uint32) uint8 {
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := []int{fi4, 42, (fi4 - (fi4 % (((fi4 & 127)) | 1)))}
        var v6 uint16 = 756
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: (("foo bar" + "a") + "hello")}
    v8 := []int{}
    switch "a" {
    case "hello":
        var v9 uint8 = 128
        _ = v9
    default:
        switch "hello" {
        case "hello":
            if (!(!false)) {
                p0 += p0
            } else {
                v10 := GenS2{F0: 65535}
                var v11 bool = (-1 != -1)
                _ = v10
                _ = v11
            }
        case "\n\t":
            if false {
                switch ("hello" + "foo bar") {
                case "foo bar":
                    switch "a" {
                    case "":
                        p0 += p0
                    case "foo bar":
                        fmt.Println(p0)
                    case "a":
                        switch "foo bar" {
                        case "":
                            var v12 int8 = -27
                            _ = v12
                        case "a":
                            v13 := GenS1{F0: ""}
                            _ = v13
                        }
                    default:
                        p0 = 42
                    }
                case "hello":
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 uint32 = p0
                        if false {
                            v16 := GenS2{F0: 42}
                            fmt.Println(v15)
                            _ = v16
                        }
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = fi14
                        _ = v15
                    }
                default:
                    p0 = p0
                }
                var v17 string = ("a" + "a")
                var v18 int32 = -940
                _ = v17
                _ = v18
            } else {
                switch "a" {
                case "\n\t":
                    for fi19 := 0; fi19 < 14; fi19++ {
                        fmt.Println(fi19, fi19)
                        _ = fi19
                    }
                case "a":
                    if false {
                        v20 := GenS3{F0: ("" + "foo bar"), F1: 798, F2: 42}
                        _ = v20
                    } else {
                        var v21 bool = (!(("a" < "hello") || false))
                        _ = v21
                    }
                case "foo bar":
                    p0 = (p0 & (p0 | 1))
                default:
                    _ = len(v8)
                }
                if false {
                    if len(v8) > 0 {
                        v8[0] = -828
                    }
                }
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint16) int64 {
    v22 := GenS4{F0: 1}
    switch 500 {
    case 449:
        v23 := []int{}
        _ = v23
    case 4294967295:
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS5{F0: 42, F1: (fi24 / (((fi24 | (fi24 * fi24))) | 1)), F2: 528}
            p0 += p0
            _ = fi24
            _ = v25
        }
    case 2147483647:
        p0 -= p0
    }
    _ = p0
    _ = v22
    return 0
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{2147483647, 0}
    var v3 uint32 = 4294967295
    v26 := GenS6{F0: -1, F1: ""}
    v27 := GenS7{F0: 42, F1: "a", F2: false}
    switch 128 {
    case 170:
        switch 2147483647 {
        case 32767:
            if true {
                v3 = (v3 / ((1) | 1))
                var v28 bool = true
                fmt.Println(v1, v3, v28)
                _ = v28
            }
        case 128:
            var v29 int16 = 452
            _ = v29
        case -646:
            if (1e10 >= float64(-2.25)) {
                var v30 uint32 = (^(v3))
                var v31 float64 = 3.14159
                fmt.Println(v1, v30)
                _ = v30
                _ = v31
            }
        default:
            v1 -= (v1 | (v1 ^ (^(v1))))
        }
    case 32767:
        v32 := GenS3{F0: "a", F1: 1, F2: 2147483647}
        _ = v32
    }
    var v33 float32 = float32(-2.25)
    var v34 uint = 2
    switch 42 {
    case 2147483647:
        if (true || (2 < (v1 | 709))) {
            var v35 bool = (false && (-128 <= 66))
            _ = v35
        }
    default:
        v34 += (^(v34))
    }
    v36 := GenS5{F0: 2, F1: 32768, F2: 0}
    if ((0 > -9223372036854775808) || (false || false)) {
        for fi37 := 0; fi37 < 8; fi37++ {
            if (v3 > (v3 + v3)) {
                switch "" {
                case "foo bar":
                    var v38 int8 = 1
                    _ = v38
                case "\n\t":
                    v39 := GenS8{F0: 127}
                    _ = v39
                case "hello":
                    v3 = (v3 % ((v3) | 1))
                }
                switch 32768 {
                case 128:
                    fmt.Println(fi37, v33, v1)
                case 127:
                    for fi40 := 0; fi40 < 17; fi40++ {
                        if len(v2) > 0 {
                            v2[0] = v3
                        }
                        _ = fi40
                    }
                case 132:
                    var v41 float64 = ((float64(v33) * float64((float32((-2.25 + (1e10 * -2.25))) - (v33 - ((v33 * float32(1.5)) - 1.5))))) / 1.5)
                    _ = v41
                default:
                    var v42 int8 = 1
                    _ = v42
                }
            }
            for fi43 := 0; fi43 < 17; fi43++ {
                if false {
                    v44 := GenS5{F0: (^(v1)), F1: fi37, F2: v1}
                    if true {
                        if (false || (3.14159 < -2.25)) {
                            var v45 int64 = 1
                            fmt.Println(v45, fi37)
                            _ = v45
                        } else {
                            var v46 int = fi37
                            var v47 string = ((("hello" + ("foo bar" + ("hello" + "foo bar"))) + ("foo bar" + "")) + "hello")
                            fmt.Println(v33)
                            _ = len(v2)
                            _ = v46
                            _ = v47
                        }
                    } else {
                        switch ("foo bar" + "") {
                        case "foo bar":
                            v33 += v33
                        default:
                            var v48 uint8 = 128
                            _ = v48
                        }
                        var v49 int32 = 1
                        _ = v49
                    }
                    if len(v2) > 0 {
                        v2[0] = (v3 ^ 955)
                    }
                    _ = v44
                }
                _ = fi43
            }
            _ = fi37
        }
    }
    if (32767 == 0) {
        var v50 uint16 = 255
        _ = v50
    } else {
        v1 += v1
    }
    var v51 string = "foo bar"
    fmt.Println(v1, v3)
    switch -2147483648 {
    case 32767:
        v52 := GenS9{F0: 0}
        _ = v52
    default:
        if (!((v1 ^ (v1 / ((v1) | 1))) == (v1 >> ((2) & 63)))) {
            var v53 int8 = 42
            if true {
                var v54 bool = (330 < 255)
                _ = len(v2)
                var v55 string = v51
                _ = v54
                _ = v55
            } else {
                if ((v51 + "hello") > "foo bar") {
                    var v56 float32 = v33
                    for fi57 := 0; fi57 < 16; fi57++ {
                        v1 = 634
                        _ = fi57
                    }
                    _ = v56
                }
            }
            _ = v53
        }
    }
    v58 := GenS3{F0: "hello", F1: 99, F2: 1}
    var v59 int16 = 1
    if false {
        for fi60 := 0; fi60 < 7; fi60++ {
            var v61 string = ((("\n\t" + "hello") + (v51 + ("a" + ((("a" + ("a" + v51)) + "") + v51)))) + v51)
            _ = fi60
            _ = v61
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v51:", v51)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v51
    _ = v58
    _ = v59
}
