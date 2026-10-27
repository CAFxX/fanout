package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 string; F1 int64; F2 uint16 }
func genFn0() int16 {
    var v6 uint32 = 4294967295
    var v7 int = 32767
    _ = v6
    _ = v7
    return -32768
}
func genFn1(p0 uint16) uint32 {
    if false {
        v8 := []int{}
        _ = v8
    }
    p0 = 1
    p0 = p0
    var v9 int8 = 42
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := []int{fi10, 894, (fi10 ^ (^(fi10)))}
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int64 = -202
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v9
    return uint32(p0)
}

func main() {
    var v1 bool = (true && ((42 < -1) || (1 < 2147483647)))
    var v2 int16 = 32767
    v3 := []string{(((("foo bar" + "a") + "hello") + "hello") + "hello"), "\n\t", ("" + ("" + "")), "a"}
    var v4 uint64 = 2
    var v5 int32 = -1
    fmt.Println(v1, v5)
    v14 := GenS1{F0: 4294967295, F1: (v5 / ((-846) | 1)), F2: (v5 & (v5 * (v5 - v5)))}
    if ((!(42 <= uint16(v4))) || (2147483647 > (v5 ^ 1))) {
        var v15 uint32 = genFn1(1)
        v5 = v5
        _ = v15
    }
    v16 := GenS1{F0: 2, F1: (v5 % ((2147483647) | 1)), F2: (v5 - v5)}
    for fi17 := 0; fi17 < 20; fi17++ {
        v1 = ((("\n\t" != ("" + "a")) && (1 > 150)) || (true && (((true || v1) || (false && false)) || false)))
        _ = fi17
    }
    if v1 {
        var v18 int64 = -533
        for fi19 := 0; fi19 < 8; fi19++ {
            v4 = (v4 % (((v4 << ((872) & 63))) | 1))
            for fi20 := 0; fi20 < 4; fi20++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v21 := GenS1{F0: 0, F1: 0, F2: v5}
                if (v4 < v4) {
                    fmt.Println(fi19, v2)
                    if ((!true) || true) {
                        v3 = append(v3, "")
                        v2 = -676
                    } else {
                        fmt.Println(fi20, v2)
                    }
                    var v22 int16 = (v2 - v2)
                    _ = v22
                } else {
                    fmt.Println(fi19, v1)
                    if v1 {
                        if (!v1) {
                            if (!v1) {
                                if ((-2.25 / (-2.25 / 0.1)) != 0.0) {
                                    v3 = append(v3, "hello")
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v23 := GenS1{F0: 831, F1: v5, F2: (v5 | v5)}
                                    _ = v23
                                }
                                v4 += 42
                                var v24 int = 42
                                _ = v24
                            }
                            v18 = (v18 / ((922) | 1))
                            var v25 float64 = 0.0
                            _ = v25
                        } else {
                            switch -1 {
                            case 1:
                                if (((0.1 + -2.25) >= (0.1 + (((1e10 / (-2.25 + 1e10)) / (-2.25 / 1.5)) / ((3.14159 * 1e10) - (1.5 * (0.0 * float64(-2.25))))))) || ((v4 * 31) < 2)) {
                                    if v1 {
                                        var v26 uint16 = 42
                                        _ = v26
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    if ((!v1) || ((v18 == v18) && true)) {
                                        v27 := GenS1{F0: 2, F1: (v5 - (^(v5))), F2: 1}
                                        _ = v27
                                    }
                                    v28 := GenS1{F0: 18446744073709551615, F1: v5, F2: (v5 % ((v5) | 1))}
                                    _ = v28
                                }
                            default:
                                v4 += (v4 << ((4294967295) & 63))
                            }
                            fmt.Println(fi19, v18)
                        }
                    } else {
                        v2 = v2
                    }
                }
                if v1 {
                    var v29 uint16 = 255
                    if v1 {
                        v30 := GenS1{F0: 947, F1: v5, F2: v5}
                        _ = v30
                    }
                    var v31 uint = 42
                    _ = v29
                    _ = v31
                }
                var v32 int32 = v5
                var v33 uint64 = v4
                _ = fi20
                _ = v21
                _ = v32
                _ = v33
            }
            _ = fi19
        }
        _ = v18
    }
    if (0 < 0) {
        var v34 uint16 = 795
        if v1 {
            switch (("foo bar" + "hello") + "a") {
            case "a":
                var v35 int = -1
                _ = v35
            case "foo bar":
                v36 := GenS1{F0: 18446744073709551615, F1: v5, F2: -1}
                _ = v36
            case "":
                if true {
                    if v1 {
                        var v37 uint16 = v34
                        v38 := GenS2{F0: -9223372036854775808, F1: v4, F2: v5}
                        _ = v37
                        _ = v38
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ((("" + "\n\t") + "foo bar") + "a")
                        }
                        v39 := GenS3{F0: ("\n\t" + "hello"), F1: 662, F2: 255}
                        v40 := GenS3{F0: "", F1: 42, F2: v34}
                        _ = v39
                        _ = v40
                    }
                    v4 = (^(v4))
                    v41 := GenS1{F0: 255, F1: -2147483648, F2: (v5 / ((v5) | 1))}
                    _ = v41
                }
            }
        } else {
            for fi42 := 0; fi42 < 17; fi42++ {
                switch ("" + (("" + "hello") + "a")) {
                case "\n\t":
                    for fi43 := 0; fi43 < 16; fi43++ {
                        switch "\n\t" {
                        case "":
                            var v44 uint = 0
                            _ = v44
                        case "\n\t":
                            v34 = 32767
                        default:
                            var v45 uint16 = v34
                            _ = v45
                        }
                        _ = fi43
                    }
                case "a":
                    v4 -= 2
                default:
                    for fi46 := 0; fi46 < 20; fi46++ {
                        v3 = append(v3, "foo bar")
                        _ = fi46
                    }
                }
                _ = fi42
            }
        }
        if len(v3) > 0 {
            v3[0] = ""
        }
        _ = v34
    }
    if v1 {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        v47 := GenS1{F0: 1, F1: v5, F2: (v5 >> ((42) & 63))}
        _ = v47
    }
    var v48 string = ""
    fmt.Println(v48, v48, v4)
    var v49 int64 = 9223372036854775807
    var v50 uint32 = 626
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
