package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16, p1 int) string {
    var v5 bool = (!(0 <= -128))
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int16 = -32768
        p0 = p0
        _ = fi6
        _ = v7
    }
    v5 = true
    _ = p0
    _ = p1
    _ = v5
    return ("foo bar" + "\n\t")
}
func genFn1(p0 uint16) uint {
    p0 = p0
    p0 -= 147
    p0 = p0
    _ = p0
    return 2
}

func main() {
    var v1 uint16 = 255
    var v2 uint64 = 2
    v3 := make(map[string]uint32)
    var v4 uint = 0
    if true {
        v4 = uint(v1)
        if (!("foo bar" >= (("hello" + "\n\t") + ("foo bar" + "")))) {
            if true {
                fmt.Println(v1, v4, v1)
                fmt.Println(v1, v2, v4)
                if false {
                    v4 = genFn1(v1)
                    v4 = genFn1((^(v1)))
                    if ("a" > ("a" + "")) {
                        v1 = (^(v1))
                        v2 = 2
                        if true {
                            switch "foo bar" {
                            case "hello":
                                delete(v3, ("hello" + ""))
                            case "a":
                                if true {
                                    var v8 int16 = 1000
                                    _ = v8
                                } else {
                                    v9 := []int{}
                                    _ = v9
                                }
                            case "foo bar":
                                var v10 bool = (127 < -2147483648)
                                _ = v10
                            }
                        }
                    } else {
                        for fi11 := 0; fi11 < 5; fi11++ {
                            v12 := GenS1{F0: 42}
                            if false {
                                v2 = (v2 & (v2 * (^(v2))))
                                if (fi11 >= (fi11 << ((v4) & 63))) {
                                    for fi13 := 0; fi13 < 1; fi13++ {
                                        v2 = v2
                                        _ = fi13
                                    }
                                    for fi14 := 0; fi14 < 3; fi14++ {
                                        var v15 float32 = (0.0 - (1.5 + 0.1))
                                        switch "a" {
                                        case "":
                                            var v16 int64 = 90
                                            _ = v16
                                        }
                                        var v17 int = (^(fi11))
                                        _ = fi14
                                        _ = v15
                                        _ = v17
                                    }
                                    var v18 uint32 = uint32(v1)
                                    _ = v18
                                }
                                fmt.Println(v4)
                            }
                            _ = fi11
                            _ = v12
                        }
                        v19 := []int{}
                        _ = v19
                    }
                } else {
                    v20 := GenS1{F0: -1}
                    var v21 int32 = 1
                    _ = v20
                    _ = v21
                }
            } else {
                switch -96 {
                case 0:
                    v22 := []int{-1, -9223372036854775808}
                    _ = v22
                }
            }
            v23 := GenS1{F0: -26}
            var v24 int8 = -106
            _ = v23
            _ = v24
        } else {
            if (true && true) {
                fmt.Println(v4, v2, v4)
                v1 -= (^(v1))
                switch -102 {
                case 0:
                    fmt.Println(v1, v1)
                case 482:
                    var v25 uint = (v4 - (^(v4)))
                    _ = v25
                }
            }
        }
    }
    v4 = (v4 & (v4 ^ (v4 << ((397) & 63))))
    v26 := GenS2{F0: ((2147483648 >= 268) || true), F1: -1}
    var v27 uint8 = 128
    v28 := GenS1{F0: int8(v1)}
    for fi29 := 0; fi29 < 10; fi29++ {
        if false {
            if (uint64(v1) == v2) {
                if true {
                    fmt.Println(v4, v4, fi29)
                }
                delete(v3, "\n\t")
            } else {
                v30 := []int{(fi29 * 1)}
                if false {
                    switch (fi29 * fi29) {
                    case -512:
                        v2 = (^(v2))
                    case 780:
                        for fi31 := 0; fi31 < 16; fi31++ {
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                            if (!((((1.5 * (1.5 * 0.0)) * 1e10) - float64(1e10)) >= ((3.14159 / 1.0) + 1.5))) {
                                var v32 uint64 = v2
                                var v33 float32 = 0.0
                                var v34 uint64 = v2
                                if true {
                                    var v35 uint8 = 119
                                    _ = v35
                                } else {
                                    if false {
                                        v36 := GenS2{F0: ((v33 / (((v33 / -2.25) * 0.0) + (((0.1 / v33) / v33) - ((v33 * -2.25) + 0.1)))) > v33), F1: -1}
                                        switch -2147483648 {
                                        case 1:
                                            fmt.Println(v4)
                                        }
                                        v37 := GenS1{F0: 100}
                                        v38 := GenS3{F0: -849}
                                        var v39 string = "hello"
                                        var v40 float32 = (0.1 * (((1.5 * v33) * v33) * v33))
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                    if true {
                                        v41 := GenS2{F0: ("foo bar" < ("foo bar" + "hello")), F1: 843}
                                        _ = v41
                                    }
                                }
                                v42 := GenS3{F0: (fi29 >> ((genFn1((v1 & (v1 * v1)))) & 63))}
                                _ = v32
                                _ = v33
                                _ = v34
                                _ = v42
                            }
                            var v43 bool = false
                            v2 = (v2 % (((v2 % ((v2) | 1))) | 1))
                            _ = fi31
                            _ = v43
                        }
                    case -2147483648:
                        var v44 uint16 = v1
                        _ = v44
                    default:
                        var v45 float32 = 0.1
                        _ = v45
                    }
                    fmt.Println(v2)
                    fmt.Println(v2, fi29)
                } else {
                    if len(v30) > 0 {
                        v30[0] = (fi29 & fi29)
                    }
                    v4 = v4
                }
                fmt.Println(v4, v27, fi29)
                _ = v30
            }
            v2 = (v2 * (v2 - (^(v2))))
            var v46 bool = (!((false || (!(-115 > -562))) || (0 != -883)))
            _ = v46
        } else {
            if ((296 <= v1) || (127 == -128)) {
                v47 := v3[("a" + "")]
                v48 := GenS3{F0: (fi29 * int(v1))}
                _ = v47
                _ = v48
            } else {
                var v49 uint64 = v2
                _ = v49
            }
        }
        if (!false) {
            if true {
                fmt.Println(fi29)
                var v50 float64 = (3.14159 / 1.5)
                _ = v50
            }
            v27 = v27
            v2 = (v2 % ((v2) | 1))
        } else {
            if (42 == 338) {
                if ("\n\t" <= "\n\t") {
                    v27 += (^(v27))
                    v51 := GenS2{F0: false, F1: 157}
                    _ = v51
                }
                v4 += v4
                if (2147483648 != 2) {
                    fmt.Println(v1, v2)
                    _ = len(v3)
                    var v52 string = "hello"
                    _ = v52
                }
            } else {
                v27 = 213
                var v53 int8 = 1
                _ = v53
            }
            v27 += v27
        }
        var v54 float64 = 1.5
        _ = fi29
        _ = v54
    }
    var v55 int = 2147483647
    v55 += (v55 ^ v55)
    fmt.Println(v2, v27, v27)
    var v56 float32 = ((1.5 - (1.5 * 1e10)) / 3.14159)
    v1 = 255
    var v57 int = v55
    delete(v3, "a")
    v58 := GenS4{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
