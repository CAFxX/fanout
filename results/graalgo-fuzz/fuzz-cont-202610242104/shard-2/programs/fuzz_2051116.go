package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint }
type GenS8 struct { F0 uint; F1 bool }
func genFn0(p0 string, p1 int32) int16 {
    v4 := []int{42, 2147483647}
    v5 := GenS1{F0: 127, F1: 175, F2: 0}
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: (fi7 - fi6), F1: -9223372036854775808, F2: 18446744073709551615}
            var v9 uint8 = 24
            p1 -= p1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    switch 2147483647 {
    case 42:
        var v10 uint8 = 151
        _ = v10
    default:
        fmt.Println(p0)
    }
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -32768
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 42
    var v3 uint8 = 42
    fmt.Println(v2)
    v2 += (v2 % ((34) | 1))
    fmt.Println(v2)
    var v11 float32 = -2.25
    v12 := GenS2{F0: 826}
    v13 := GenS1{F0: 127, F1: 32767, F2: 42}
    if (2 >= 385) {
        v14 := []int{32768}
        var v15 float32 = v11
        var v16 int64 = 1
        _ = v14
        _ = v15
        _ = v16
    } else {
        var v17 uint32 = 2147483648
        _ = v17
    }
    fmt.Println(v11, v11, v3)
    v3 += 2
    var v18 uint = 18446744073709551615
    v3 = (v3 - 128)
    if (888 == 0) {
        if (-1 == -9223372036854775808) {
            v19 := GenS3{F0: genFn0("\n\t", 1)}
            _ = v19
        }
        var v20 uint64 = 18446744073709551615
        for fi21 := 0; fi21 < 12; fi21++ {
            switch int(v20) {
            case 32767:
                var v22 uint64 = (^(v20))
                _ = v22
            case 240:
                delete(v1, "a")
            case 4294967295:
                v18 = (v18 ^ v18)
            default:
                v18 += v18
            }
            if true {
                fmt.Println(v18, v2)
            } else {
                fmt.Println(v2, v3, v3)
            }
            delete(v1, ("hello" + "hello"))
            _ = fi21
        }
        _ = v20
    } else {
        switch ("" + ("" + (("" + "foo bar") + "\n\t"))) {
        case "a":
            var v23 float64 = 0.0
            _ = v23
        }
        if true {
            v24 := GenS4{F0: 1000, F1: 0}
            if (18446744073709551615 >= (v18 | (v18 ^ 42))) {
                v3 = (v3 % ((131) | 1))
                if (false || (v11 == (v11 * v11))) {
                    fmt.Println(v11, v18, v3)
                    var v25 int16 = genFn0("", 1)
                    switch "foo bar" {
                    case "a":
                        if ((v25 ^ (v25 % ((genFn0("", -240)) | 1))) <= -582) {
                            var v26 int16 = (v25 / ((genFn0("hello", 0)) | 1))
                            _ = v26
                        } else {
                            fmt.Println(v3)
                            var v27 uint = 74
                            _ = v27
                        }
                    case "\n\t":
                        delete(v1, "")
                    case "hello":
                        v28 := GenS4{F0: (^(v25)), F1: 9223372036854775807}
                        _ = v28
                    }
                    _ = v25
                } else {
                    var v29 int32 = 1
                    v30 := []int{-9223372036854775808, 32767, -9223372036854775808}
                    if (!(1000 < 32767)) {
                        v3 = (v3 - v3)
                    } else {
                        var v31 int32 = v29
                        _ = v31
                    }
                    _ = v29
                    _ = v30
                }
            } else {
                if ((!(false || (false && (true || true)))) || (true && true)) {
                    var v32 float64 = (1.5 * (3.14159 - (((((0.1 - 0.1) * -2.25) / 1.0) - 1e10) - (((1.5 * 1e10) + (((1e10 - 0.0) * 0.1) - (((1e10 * 3.14159) + 1.5) * 3.14159))) - 1e10))))
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v34 := []int{(fi33 << ((4294967295) & 63)), fi33}
                        if false {
                            var v35 int16 = -1
                            if (!true) {
                                v36 := GenS1{F0: 2147483647, F1: int(v18), F2: 162}
                                var v37 int16 = v35
                                for fi38 := 0; fi38 < 7; fi38++ {
                                    v39 := GenS5{F0: 2147483647, F1: -2147483648}
                                    if false {
                                        var v40 uint16 = 32767
                                        var v41 uint = v18
                                        if true {
                                            v35 = (v35 * int16(v3))
                                            var v42 uint16 = v40
                                            delete(v1, "\n\t")
                                            _ = v42
                                        }
                                        _ = v40
                                        _ = v41
                                    }
                                    v35 = 1000
                                    _ = fi38
                                    _ = v39
                                }
                                _ = v36
                                _ = v37
                            } else {
                                switch ("" + "hello") {
                                case "a":
                                    fmt.Println(v35)
                                }
                                var v43 int32 = 36
                                v11 = v11
                                _ = v43
                            }
                            v18 -= v18
                            _ = v35
                        } else {
                            switch (fi33 | (-(fi33))) {
                            case 4294967295:
                                var v44 float32 = 3.14159
                                _ = v44
                            case 42:
                                if false {
                                    v18 = 1
                                    var v45 int32 = -2147483648
                                    _ = v45
                                }
                            case -2147483648:
                                if len(v34) > 0 {
                                    v34[0] = fi33
                                }
                            }
                        }
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                }
            }
            for fi46 := 0; fi46 < 8; fi46++ {
                if (1000 < genFn0(("" + ("\n\t" + "a")), -1)) {
                    if (2 == 2) {
                        for fi47 := 0; fi47 < 13; fi47++ {
                            if (!true) {
                                var v48 int8 = (v2 >> ((2) & 63))
                                var v49 int16 = -1
                                if true {
                                    v50 := GenS6{F0: 2147483648, F1: (^(v18))}
                                    _ = v50
                                }
                                if ((v49 + genFn0("", 0)) < (v49 - (^(v49)))) {
                                    v51 := v1[(("\n\t" + "a") + "")]
                                    switch "a" {
                                    case "hello":
                                        var v52 uint = 548
                                        _ = v52
                                    }
                                    _ = v51
                                }
                                _ = v48
                                _ = v49
                            }
                            v18 = v18
                            delete(v1, ("a" + ""))
                            _ = fi47
                        }
                        delete(v1, "a")
                        var v53 uint8 = 2
                        _ = v53
                    }
                    v11 = (float32(1e10) - v11)
                    var v54 int8 = (v2 | (v2 << ((0) & 63)))
                    _ = v54
                } else {
                    switch 337 {
                    case -524:
                        v18 = v18
                    case -1:
                        if ("\n\t" > "foo bar") {
                            v11 -= ((-2.25 / v11) + v11)
                            var v55 uint64 = 694
                            if true {
                                v56 := GenS1{F0: fi46, F1: -1, F2: 2}
                                _ = v56
                            }
                            _ = v55
                        } else {
                            v2 = v2
                            fmt.Println(v3, v11)
                        }
                    default:
                        v57 := v1["foo bar"]
                        _ = v57
                    }
                    v58 := v1[""]
                    v59 := []int{fi46}
                    _ = v58
                    _ = v59
                }
                v60 := GenS7{F0: 378, F1: (^(v3)), F2: (v18 % ((v18) | 1))}
                _ = fi46
                _ = v60
            }
            _ = v24
        } else {
            var v61 int32 = 42
            if (("hello" + "a") > (("a" + ("hello" + "\n\t")) + "\n\t")) {
                v61 -= (v61 | int32(v3))
                var v62 int8 = 42
                if ((v62 ^ (v2 + (v2 * v62))) > -106) {
                    v63 := GenS8{F0: (v18 + (v18 & 18446744073709551615)), F1: (true || false)}
                    if true {
                        v62 = v62
                        v2 += v62
                    }
                    v64 := []int{32767}
                    _ = v63
                    _ = v64
                }
                _ = v62
            } else {
                fmt.Println(v3)
                var v65 bool = true
                _ = v65
            }
            _ = v61
        }
    }
    v18 += v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
}
