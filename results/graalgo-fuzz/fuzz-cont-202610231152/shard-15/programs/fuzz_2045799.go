package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 int32 }
func genFn0() int32 {
    if (("a" + ("foo bar" + "foo bar")) < (((("" + "foo bar") + "\n\t") + "") + ("" + "\n\t"))) {
        var v6 int32 = 2147483647
        v6 -= (v6 << ((216) & 63))
        _ = v6
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            switch ("hello" + "foo bar") {
            case "a":
                if false {
                    v8 := GenS1{F0: (fi7 >> ((0) & 63)), F1: 1}
                    var v9 uint8 = 0
                    fmt.Println(fi7, v9)
                    _ = v8
                    _ = v9
                }
            case "hello":
                if ((0.0 + -2.25) >= (0.0 - ((float32(-2.25) - 1.5) + ((0.0 / 3.14159) + float32(1e10))))) {
                    var v10 float64 = 1.5
                    v11 := GenS1{F0: 2147483647, F1: 18446744073709551615}
                    _ = v10
                    _ = v11
                }
            default:
                var v12 uint = 2
                _ = v12
            }
            _ = fi7
        }
        v13 := []int{}
        var v14 int32 = 0
        _ = v13
        _ = v14
    }
    var v15 uint16 = 65535
    var v16 uint32 = 2147483648
    if false {
        v17 := GenS2{F0: 6, F1: 127, F2: 42}
        var v18 string = "foo bar"
        switch v18 {
        case "hello":
            v19 := []int{-362, 32768}
            _ = v19
        default:
            v20 := []int{0}
            _ = v20
        }
        _ = v17
        _ = v18
    }
    _ = v15
    _ = v16
    return 209
}
func genFn1() uint {
    if (!true) {
        var v21 int8 = 0
        _ = v21
    } else {
        switch 42 {
        case 127:
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := []int{(-(fi22)), fi22, (fi22 / ((0) | 1))}
                _ = fi22
                _ = v23
            }
        case 2147483647:
            var v24 string = ""
            _ = v24
        }
    }
    v25 := GenS3{F0: 255, F1: 42, F2: 42}
    var v26 uint8 = 2
    _ = v25
    _ = v26
    return 401
}

func main() {
    var v1 int8 = -81
    var v2 bool = true
    v3 := []uint32{794, 4294967295, 4294967295, 0}
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    var v27 string = ("hello" + "a")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = fi28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        v27 = "hello"
        v30 := GenS1{F0: 667, F1: 2}
        var v31 bool = false
        _ = fi29
        _ = v30
        _ = v31
    }
    if false {
        var v32 uint32 = 764
        _ = v32
    } else {
        switch -2147483648 {
        case -1:
            for fi33 := 0; fi33 < 18; fi33++ {
                var v34 int = (fi33 + fi33)
                v2 = (42 <= 446)
                v35 := GenS3{F0: 0, F1: genFn0(), F2: v1}
                _ = fi33
                _ = v34
                _ = v35
            }
        case 42:
            v36 := GenS4{F0: 653, F1: 461, F2: 2}
            _ = v36
        default:
            for fi37 := 0; fi37 < 6; fi37++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi37
            }
        }
    }
    if (!(415 < -273)) {
        if v2 {
            var v38 bool = (995 != 53)
            v39 := GenS3{F0: 42, F1: 42, F2: (v1 / ((v1) | 1))}
            v40 := GenS2{F0: 42, F1: (^(v1)), F2: 42}
            _ = v38
            _ = v39
            _ = v40
        } else {
            var v41 int16 = 1000
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v41
        }
        for fi42 := 0; fi42 < 15; fi42++ {
            v43 := GenS1{F0: (fi42 << ((140) & 63)), F1: 446}
            fmt.Println(v27)
            _ = fi42
            _ = v43
        }
        fmt.Println(v27)
    } else {
        v44 := GenS5{F0: genFn1()}
        v1 = 0
        v2 = true
        _ = v44
    }
    switch ("" + "") {
    case "\n\t":
        v27 = (((v27 + v27) + v27) + "\n\t")
    case "":
        var v45 string = ("\n\t" + v27)
        _ = v45
    default:
        switch v27 {
        case "\n\t":
            for fi46 := 0; fi46 < 17; fi46++ {
                v1 += (v1 * v1)
                v1 = (v1 | (v1 + -1))
                for fi47 := 0; fi47 < 4; fi47++ {
                    v2 = true
                    _ = fi47
                }
                _ = fi46
            }
        case "foo bar":
            v2 = (469 < -154)
        default:
            v5[v27] = (("" + ("" + (v27 + (v27 + v27)))) + (((((v27 + "hello") + v27) + (v27 + ("hello" + v27))) + v27) + v27))
        }
    }
    var v48 uint32 = 0
    if (v48 >= 2147483647) {
        for fi49 := 0; fi49 < 20; fi49++ {
            for fi50 := 0; fi50 < 3; fi50++ {
                var v51 int32 = 377
                _ = fi50
                _ = v51
            }
            for fi52 := 0; fi52 < 10; fi52++ {
                v53 := GenS6{F0: int32(fi49)}
                if (!v2) {
                    v27 = v27
                } else {
                    if (false || (3.14159 <= -2.25)) {
                        var v54 int32 = 519
                        _ = v54
                    }
                    v1 = (v1 * 100)
                }
                fmt.Println(fi49)
                _ = fi52
                _ = v53
            }
            switch 128 {
            case -706:
                _ = len(v3)
            case 42:
                var v55 uint = 255
                _ = v55
            case 0:
                var v56 int16 = -698
                _ = v56
            }
            _ = fi49
        }
        if v2 {
            var v57 bool = v2
            _ = v57
        }
    } else {
        fmt.Println(v27, v48)
        v58 := GenS7{F0: v48, F1: 0}
        if v2 {
            if (!(v2 || (genFn0() <= -1))) {
                v48 -= (v48 ^ v48)
            }
            if ((!v2) && (1.5 >= -2.25)) {
                fmt.Println(v1)
                fmt.Println(v2, v2, v2)
                v27 = v27
            }
        }
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
