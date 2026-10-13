package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint; F2 string }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint16; F1 bool; F2 string }
func genFn0(p0 string, p1 uint16, p2 uint8) string {
    fmt.Println(p0, p2, p0)
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return p0
}
func genFn1() int16 {
    if false {
        var v5 int64 = 1
        fmt.Println(v5)
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        switch (("foo bar" + "") + (("" + "foo bar") + "foo bar")) {
        case "a":
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 uint = 626
                var v9 int16 = 1000
                _ = fi7
                _ = v8
                _ = v9
            }
        case "foo bar":
            v10 := GenS1{F0: -279, F1: 1, F2: 255}
            _ = v10
        }
        _ = fi6
    }
    return 42
}

func main() {
    var v1 int = 1
    v2 := make(map[uint32]bool)
    var v3 uint16 = 742
    var v4 uint16 = (^(v3))
    if (!(false || false)) {
        var v11 int32 = -1
        _ = v11
    } else {
        switch (v1 - (v1 & 32767)) {
        case 127:
            v3 += 55
        case 4294967295:
            for fi12 := 0; fi12 < 2; fi12++ {
                v2[2] = (-9223372036854775808 >= -220)
                fmt.Println(v4, v3, v3)
                _ = fi12
            }
        default:
            var v13 int32 = -2147483648
            _ = v13
        }
        v1 -= -1
        _ = len(v2)
    }
    switch -756 {
    case 1:
        v14 := []int{(v1 ^ 9223372036854775807), v1, v1}
        _ = v14
    case 603:
        v3 = (v3 / ((v3) | 1))
    }
    fmt.Println(v1, v4, v3)
    switch ((("a" + "\n\t") + "hello") + "foo bar") {
    case "hello":
        v1 = (v1 << ((18446744073709551615) & 63))
    case "":
        v1 += v1
    case "a":
        var v15 int = (v1 << ((4294967295) & 63))
        _ = v15
    default:
        if true {
            v2[2147483648] = true
            var v16 float64 = (((((float64(1.5) * 3.14159) - -2.25) + float64((1e10 + 3.14159))) + (float64(float32((((0.0 / 1.0) + -2.25) - 0.0))) / (0.0 + float64((1.5 / -2.25))))) + 1.5)
            v4 += (v4 * 409)
            _ = v16
        } else {
            var v17 bool = false
            fmt.Println(v1, v4)
            _ = v17
        }
    }
    switch -972 {
    case -654:
        v1 = v1
    case 127:
        var v18 uint32 = 42
        _ = v18
    case -574:
        v1 -= 32768
    default:
        v19 := GenS2{F0: ""}
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        switch -231 {
        case -57:
            var v21 float32 = 3.14159
            _ = v21
        case 32767:
            switch -2147483648 {
            case 32767:
                var v22 bool = true
                _ = v22
            case 42:
                if (942 < 32767) {
                    for fi23 := 0; fi23 < 19; fi23++ {
                        var v24 int = -57
                        var v25 int = (fi23 ^ (fi23 >> ((464) & 63)))
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                } else {
                    if true {
                        fmt.Println(fi20)
                        v26 := []int{fi20, -550, (^(v1))}
                        _ = v26
                    } else {
                        var v27 float64 = 3.14159
                        if false {
                            v28 := []int{fi20, fi20}
                            v29 := GenS1{F0: 701, F1: 42, F2: 18446744073709551615}
                            _ = v28
                            _ = v29
                        } else {
                            for fi30 := 0; fi30 < 14; fi30++ {
                                v4 -= v3
                                var v31 string = (("hello" + "foo bar") + "a")
                                var v32 int64 = 0
                                switch v31 {
                                case "foo bar":
                                    if (((v32 < 42) || (!(false || (!false)))) || true) {
                                        v31 = ""
                                        if (979 > (fi20 | -622)) {
                                            var v33 uint = 255
                                            _ = v33
                                        }
                                    } else {
                                        if true {
                                            v34 := []int{177, fi20}
                                            _ = v34
                                        }
                                        var v35 uint64 = 1
                                        var v36 uint64 = 4294967295
                                        _ = v35
                                        _ = v36
                                    }
                                case "\n\t":
                                    if ("hello" != ("" + v31)) {
                                        v1 -= -9223372036854775808
                                        var v37 uint16 = (v3 % ((v4) | 1))
                                        fmt.Println(v3, v37)
                                        _ = v37
                                    } else {
                                        v31 = v31
                                        v31 = v31
                                    }
                                default:
                                    var v38 int8 = 100
                                    _ = v38
                                }
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            fmt.Println(v3, v1)
                            v27 += v27
                        }
                        _ = v27
                    }
                }
            case 127:
                v39 := GenS3{F0: 2147483647}
                _ = v39
            }
        case 421:
            v4 -= 0
        }
        var v40 int32 = -248
        _ = fi20
        _ = v40
    }
    if false {
        if (!(true && false)) {
            v41 := GenS4{F0: -877, F1: 0, F2: "foo bar"}
            _ = v41
        }
        fmt.Println(v1, v1)
        delete(v2, 4294967295)
    }
    var v42 float64 = ((-2.25 - 1.5) * (-2.25 - float64(0.1)))
    v43 := GenS5{F0: 9223372036854775807, F1: 0}
    v1 = 1
    if true {
        switch ((("" + "foo bar") + "hello") + ((("hello" + "hello") + "\n\t") + "\n\t")) {
        case "":
            for fi44 := 0; fi44 < 16; fi44++ {
                if false {
                    v45 := GenS3{F0: -1}
                    _ = v45
                }
                if true {
                    v3 = v4
                    v1 -= 128
                    v4 -= v4
                }
                var v46 string = "a"
                _ = fi44
                _ = v46
            }
        default:
            fmt.Println(v42)
        }
    }
    for fi47 := 0; fi47 < 6; fi47++ {
        switch v1 {
        case 0:
            if false {
                fmt.Println(v1, v4, v3)
                var v48 uint32 = 4294967295
                switch "foo bar" {
                case "a":
                    v49 := GenS6{F0: (v3 << ((1) & 63)), F1: true, F2: "\n\t"}
                    _ = v49
                case "foo bar":
                    for fi50 := 0; fi50 < 13; fi50++ {
                        if ((((989 >= 65535) && false) && false) && false) {
                            fmt.Println(fi50, v1)
                            if false {
                                v51 := GenS6{F0: (v3 % ((33) | 1)), F1: true, F2: "a"}
                                _ = v51
                            } else {
                                var v52 string = "hello"
                                v53 := GenS1{F0: 32767, F1: 2, F2: 888}
                                _ = v52
                                _ = v53
                            }
                            v54 := GenS2{F0: "\n\t"}
                            _ = v54
                        }
                        v42 -= 1.5
                        _ = fi50
                    }
                }
                _ = v48
            }
        case -2147483648:
            fmt.Println(fi47)
        case 128:
            v42 = float64(1e10)
        }
        var v55 int32 = 42
        v56 := GenS2{F0: "foo bar"}
        _ = fi47
        _ = v55
        _ = v56
    }
    v57 := []int{}
    var v58 uint32 = 0
    fmt.Println(v3, v58)
    v59 := GenS5{F0: 9223372036854775807, F1: 42}
    var v60 uint16 = v3
    switch (v1 ^ 42) {
    case -1:
        v60 += (v60 | uint16(v58))
    case 32768:
        fmt.Println(v58)
    default:
        v60 += v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v57
    _ = v58
    _ = v59
    _ = v60
}
