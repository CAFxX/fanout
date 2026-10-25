package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint16) uint8 {
    v5 := []int{32767, 9223372036854775807, 9223372036854775807, -2147483648}
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 = p0
        if true {
            switch fi6 {
            case -2147483648:
                fmt.Println(fi6)
            }
            p0 -= p0
        }
        p0 += 1
        _ = fi6
    }
    _ = p0
    _ = v5
    return 128
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 int16 = 32767
    if true {
        switch ((("foo bar" + ("" + (("" + ("\n\t" + ((((("a" + "foo bar") + "foo bar") + ("foo bar" + "a")) + ("\n\t" + "hello")) + ""))) + (("a" + (("hello" + ("" + ("hello" + ""))) + "foo bar")) + "a")))) + ((("foo bar" + ("\n\t" + "")) + ("" + "hello")) + ("" + "\n\t"))) + "foo bar") {
        case "a":
            fmt.Println(v1, v4)
        case "hello":
            if (0 >= 0) {
                v4 = v4
                fmt.Println(v4)
                fmt.Println(v4, v1)
            }
        case "foo bar":
            var v7 int16 = (^(v4))
            _ = v7
        }
        v2["hello"] = "a"
    }
    v8 := GenS1{F0: -127, F1: 18446744073709551615, F2: (v1 + (v1 - 2147483648))}
    for fi9 := 0; fi9 < 4; fi9++ {
        switch fi9 {
        case 2147483647:
            var v10 uint = 42
            _ = v10
        case 32767:
            v11 := GenS1{F0: -1, F1: 4294967295, F2: v1}
            _ = v11
        default:
            if false {
                var v12 int32 = -322
                _ = v12
            }
        }
        if true {
            var v13 int8 = -24
            if ("" < "hello") {
                _ = len(v3)
                v14 := GenS1{F0: -128, F1: 18446744073709551615, F2: (^(v1))}
                _ = v14
            }
            if false {
                var v15 uint = 2
                fmt.Println(v15)
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS1{F0: v13, F1: (v15 & 1), F2: (^(v1))}
                    var v18 int64 = 42
                    var v19 bool = (true && false)
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = v15
            } else {
                var v20 int8 = (v13 / ((-106) | 1))
                var v21 uint = 810
                _ = v20
                _ = v21
            }
            _ = v13
        } else {
            switch (fi9 ^ (fi9 - (fi9 * fi9))) {
            case 9223372036854775807:
                v22 := []int{(fi9 + -998), 32767, 32767, 32767}
                _ = v22
            case -1:
                v1 = 4294967295
            case 1:
                v23 := GenS1{F0: 1, F1: 18446744073709551615, F2: 446}
                _ = v23
            default:
                if true {
                    var v24 int8 = 28
                    _ = v24
                } else {
                    v4 = (-(v4))
                }
            }
            v4 -= (v4 + v4)
        }
        _ = fi9
    }
    v1 += v1
    fmt.Println(v4, v4)
    switch -1 {
    case 42:
        switch -9223372036854775808 {
        case 0:
            v25 := GenS1{F0: 42, F1: 0, F2: v1}
            _ = v25
        case 128:
            v26 := []int{32767, 1, 725}
            _ = v26
        }
    case -426:
        v4 += v4
    }
    if (!(65535 == 192)) {
        var v27 int16 = (v4 ^ (v4 - (v4 * v4)))
        v28 := GenS2{F0: 42, F1: -2147483648}
        v1 += (v1 ^ v1)
        _ = v27
        _ = v28
    } else {
        v29 := []int{}
        var v30 uint8 = 0
        _ = v29
        _ = v30
    }
    v1 -= (v1 >> ((828) & 63))
    v4 += v4
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS2{F0: 1, F1: int32(v4)}
        _ = fi31
        _ = v32
    }
    switch 0 {
    case -950:
        var v33 float64 = (1.5 / (-2.25 / (1.5 / 1.0)))
        _ = v33
    case 32767:
        v34 := GenS3{F0: 367, F1: -128, F2: 4294967295}
        _ = v34
    case -90:
        if (true || (42 <= 2)) {
            _ = len(v3)
            v1 -= (v1 ^ v1)
            var v35 string = ""
            _ = v35
        }
    default:
        if (false && (false && (v1 <= (v1 + 313)))) {
            if true {
                if true {
                    var v36 int32 = -1
                    _ = v36
                } else {
                    fmt.Println(v1, v4)
                    if (true && (32767 != 0)) {
                        for fi37 := 0; fi37 < 18; fi37++ {
                            v2["\n\t"] = ((((("foo bar" + "a") + ((("foo bar" + "a") + "foo bar") + ("\n\t" + "a"))) + "\n\t") + "a") + "\n\t")
                            v38 := GenS1{F0: 0, F1: 482, F2: (v1 >> ((42) & 63))}
                            _ = fi37
                            _ = v38
                        }
                    } else {
                        v2[""] = ("\n\t" + "foo bar")
                    }
                    var v39 int64 = 9223372036854775807
                    _ = v39
                }
            }
        }
    }
    fmt.Println(v4, v4)
    var v40 float64 = (1.5 / (1e10 / -2.25))
    v41 := GenS1{F0: -45, F1: 65535, F2: (v1 << ((65535) & 63))}
    if (65535 > 32767) {
        delete(v2, "a")
    } else {
        v1 = (v1 >> ((825) & 63))
        for fi42 := 0; fi42 < 4; fi42++ {
            v43 := GenS3{F0: 1, F1: 42, F2: 4294967295}
            var v44 int16 = v4
            v40 += v40
            _ = fi42
            _ = v43
            _ = v44
        }
        fmt.Println(v4)
    }
    if (true || (42 < 2)) {
        var v45 uint16 = 42
        _ = v45
    } else {
        v1 = v1
    }
    switch "foo bar" {
    case "hello":
        v46 := GenS3{F0: 32767, F1: int8(v4), F2: 42}
        _ = v46
    case "foo bar":
        switch 32768 {
        case 128:
            var v47 uint8 = 2
            _ = v47
        default:
            if ((0.0 / -2.25) >= 0.0) {
                for fi48 := 0; fi48 < 5; fi48++ {
                    v49 := GenS4{F0: 74, F1: 675, F2: 236}
                    _ = fi48
                    _ = v49
                }
                v40 = (0.1 / float64(((1e10 / -2.25) * 0.1)))
            } else {
                v50 := GenS5{F0: 28, F1: 127}
                _ = v50
            }
        }
    case "":
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v8
}
