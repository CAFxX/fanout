package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
type GenS2 struct { F0 bool; F1 bool; F2 uint16 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0() int32 {
    v5 := GenS1{F0: 127, F1: 4294967295, F2: -9223372036854775808}
    var v6 int8 = 42
    var v7 float64 = float64(0.0)
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 int32, p2 string) int16 {
    v8 := []int{-2147483648}
    switch 0 {
    case -2147483648:
        p1 = (p1 * p1)
    case -414:
        var v9 int32 = (p0 * (p0 << ((4294967295) & 63)))
        _ = v9
    case 9223372036854775807:
        v10 := GenS1{F0: 254, F1: 492, F2: 9223372036854775807}
        _ = v10
    }
    switch 0 {
    case 4294967295:
        if false {
            p2 = "hello"
            switch p2 {
            case "a":
                p1 = p0
            default:
                var v11 int64 = 1
                _ = v11
            }
        } else {
            v12 := GenS1{F0: 128, F1: 957, F2: -721}
            p2 = p2
            switch "foo bar" {
            case "a":
                fmt.Println(p1)
            case "\n\t":
                var v13 float32 = 0.1
                _ = v13
            case "":
                p1 += (p0 / ((genFn0()) | 1))
            }
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = 0
    var v3 bool = true
    var v4 string = ((("" + "") + ("hello" + ("foo bar" + "foo bar"))) + ("" + ("hello" + "a")))
    for fi14 := 0; fi14 < 9; fi14++ {
        v4 = (v4 + v4)
        switch fi14 {
        case -763:
            var v15 uint8 = 69
            _ = v15
        default:
            v2 += int8(fi14)
        }
        switch v4 {
        case "a":
            var v16 int16 = -32768
            _ = v16
        case "hello":
            v2 = 0
        }
        _ = fi14
    }
    if (-59 == genFn1(-199, v1, ((("hello" + "\n\t") + v4) + v4))) {
        v17 := []int{}
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(v1, v2)
            _ = fi19
        }
        _ = fi18
    }
    if (genFn0() == -487) {
        for fi20 := 0; fi20 < 18; fi20++ {
            v1 = -74
            _ = fi20
        }
        if (0 >= 1000) {
            switch ((v4 + "") + v4) {
            case "":
                if v3 {
                    fmt.Println(v1, v2, v2)
                    v21 := GenS1{F0: 9223372036854775807, F1: 1, F2: 244}
                    var v22 uint64 = 0
                    _ = v21
                    _ = v22
                }
            case "foo bar":
                v4 = (v4 + ("" + (v4 + "")))
            case "\n\t":
                v23 := GenS2{F0: (v3 || v3), F1: false, F2: 42}
                _ = v23
            default:
                v4 = "foo bar"
            }
            v2 = v2
        }
    } else {
        if (-562 <= -1) {
            switch -91 {
            case 0:
                v24 := []int{-2147483648, 32767, 127}
                _ = v24
            case 127:
                v25 := []int{-561, 1, 833}
                _ = v25
            case 398:
                v26 := GenS1{F0: 489, F1: 1, F2: -258}
                _ = v26
            default:
                var v27 int16 = 0
                _ = v27
            }
            v4 = v4
            v1 = 744
        } else {
            v28 := GenS2{F0: v3, F1: false, F2: 255}
            _ = v28
        }
    }
    switch ("hello" + "") {
    case "":
        if (v3 || v3) {
            v2 -= v2
            v4 = "\n\t"
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                v2 = -1
                _ = fi29
            }
            fmt.Println(v3, v4, v3)
            v30 := GenS3{F0: 2147483647, F1: false}
            _ = v30
        }
    case "a":
        if (!(1 > 42)) {
            var v31 float32 = 1.5
            _ = v31
        }
    case "foo bar":
        var v32 bool = ((42 != 32767) || ((!(!(!false))) || (v3 || true)))
        _ = v32
    default:
        if (!(v4 >= v4)) {
            switch 127 {
            case -682:
                v3 = true
            }
        } else {
            v4 = v4
            if (0 < -1) {
                switch 127 {
                case 954:
                    v33 := GenS3{F0: 3, F1: false}
                    _ = v33
                default:
                    v2 = int8(v1)
                }
            } else {
                v4 = (v4 + "\n\t")
                switch 9223372036854775807 {
                case -9223372036854775808:
                    fmt.Println(v1, v4)
                }
                switch 128 {
                case 128:
                    for fi34 := 0; fi34 < 13; fi34++ {
                        var v35 uint16 = 65535
                        var v36 string = "hello"
                        var v37 uint16 = v35
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
            }
        }
    }
    var v38 uint32 = 419
    if v3 {
        var v39 bool = true
        if (v3 || (false || true)) {
            v4 = ("foo bar" + "")
            switch "hello" {
            case "foo bar":
                if (203 != 128) {
                    switch 32767 {
                    case 32767:
                        var v40 string = "\n\t"
                        _ = v40
                    default:
                        v41 := GenS3{F0: v1, F1: (!false)}
                        _ = v41
                    }
                    fmt.Println(v39, v1)
                } else {
                    v42 := []int{-1, 42}
                    _ = v42
                }
            default:
                fmt.Println(v1, v38)
            }
        }
        var v43 int16 = genFn1(v1, 2147483647, ((v4 + v4) + ("foo bar" + "foo bar")))
        _ = v39
        _ = v43
    }
    v44 := GenS1{F0: 4294967295, F1: 4294967295, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
}
