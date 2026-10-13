package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int8; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint; F1 int16; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int16; F1 uint; F2 uint8 }
type GenS7 struct { F0 string; F1 int64; F2 int64 }
type GenS8 struct { F0 int64; F1 int16; F2 uint8 }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 int32 }
type GenS12 struct { F0 uint32; F1 uint32 }
func genFn0() string {
    var v6 uint = 4294967295
    if (876 <= 313) {
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS1{F0: 42, F1: true}
                _ = fi8
                _ = v9
            }
            v10 := GenS2{F0: -128, F1: 816, F2: 4294967295}
            var v11 int32 = -1
            _ = fi7
            _ = v10
            _ = v11
        }
    } else {
        v6 = 255
        switch -9223372036854775808 {
        case -848:
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 string = (("foo bar" + ((("\n\t" + "a") + ("hello" + "hello")) + "")) + "hello")
                v6 = (v6 << ((42) & 63))
                v13 = v13
                _ = fi12
                _ = v13
            }
        default:
            v6 = v6
        }
    }
    if false {
        var v14 int8 = -1
        if false {
            var v15 uint = 163
            var v16 int8 = 1
            _ = v15
            _ = v16
        } else {
            if false {
                v17 := []int{0, 32767, 2147483647, 32768}
                _ = v17
            } else {
                if (255 > 255) {
                    var v18 int16 = -158
                    var v19 float32 = 3.14159
                    for fi20 := 0; fi20 < 7; fi20++ {
                        for fi21 := 0; fi21 < 19; fi21++ {
                            fmt.Println(v18)
                            v22 := GenS2{F0: int8(v6), F1: v6, F2: 4294967295}
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                    _ = v18
                    _ = v19
                } else {
                    v23 := GenS1{F0: 4294967295, F1: false}
                    if true {
                        var v24 uint16 = 255
                        for fi25 := 0; fi25 < 2; fi25++ {
                            v6 = 1
                            var v26 uint32 = 65535
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    } else {
                        v27 := []int{-1, 9223372036854775807}
                        v6 += v6
                        if false {
                            v14 += v14
                            v28 := GenS3{F0: (v6 + v6), F1: 0, F2: -987}
                            fmt.Println(v14, v6)
                            _ = v28
                        }
                        _ = v27
                    }
                    switch "a" {
                    case "":
                        var v29 uint = (v6 * v6)
                        _ = v29
                    default:
                        switch 0 {
                        case 32767:
                            v30 := GenS3{F0: v6, F1: 32767, F2: -320}
                            _ = v30
                        default:
                            v6 += (^(v6))
                        }
                    }
                    _ = v23
                }
            }
            var v31 int64 = -673
            var v32 uint64 = 42
            _ = v31
            _ = v32
        }
        v33 := GenS4{F0: 922}
        _ = v14
        _ = v33
    } else {
        var v34 int64 = 676
        switch "\n\t" {
        case "foo bar":
            switch -271 {
            case 42:
                fmt.Println(v6, v6)
            }
        }
        _ = v34
    }
    _ = v6
    return ((("" + "") + "a") + "")
}
func genFn1(p0 uint8, p1 int, p2 uint64) uint16 {
    p2 = 0
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := []int{4294967295, 128, 32768, 1}
    v2 := []int{-1}
    v3 := make(map[uint32]uint32)
    v4 := []string{"foo bar", "hello", "foo bar", "a"}
    v5 := make(map[uint32]bool)
    switch "foo bar" {
    case "foo bar":
        var v35 uint8 = 128
        _ = v35
    case "":
        var v36 string = ("a" + "foo bar")
        _ = v36
    case "hello":
        var v37 string = "hello"
        _ = v37
    }
    var v38 int8 = -128
    delete(v5, 2147483647)
    switch "\n\t" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v40 int32 = 42
            _ = v40
        } else {
            v41 := GenS5{F0: (0.1 <= 1e10), F1: 18446744073709551615}
            _ = v41
        }
        v42 := GenS6{F0: 574, F1: 42, F2: 42}
        switch fi39 {
        case 0:
            switch "" {
            case "\n\t":
                v38 += 0
            case "a":
                var v43 int = fi39
                _ = v43
            case "hello":
                v44 := GenS7{F0: "hello", F1: 0, F2: 0}
                _ = v44
            default:
                var v45 string = ""
                _ = v45
            }
        case -678:
            v46 := GenS5{F0: (false && false), F1: 215}
            _ = v46
        case 4294967295:
            for fi47 := 0; fi47 < 10; fi47++ {
                v48 := GenS7{F0: "", F1: 785, F2: 1}
                _ = fi47
                _ = v48
            }
        default:
            v38 -= -1
        }
        _ = fi39
        _ = v42
    }
    var v49 int8 = (v38 << ((65535) & 63))
    fmt.Println(v38, v49)
    v50 := GenS8{F0: 42, F1: 1, F2: 128}
    fmt.Println(v38, v49)
    v51 := GenS6{F0: -32768, F1: 820, F2: 128}
    switch 0 {
    case 1:
        var v52 uint64 = 42
        _ = v52
    default:
        v53 := GenS9{F0: 1, F1: v38}
        _ = v53
    }
    if true {
        var v54 uint64 = 891
        for fi55 := 0; fi55 < 17; fi55++ {
            for fi56 := 0; fi56 < 4; fi56++ {
                if (!false) {
                    v54 = 1
                    fmt.Println(fi56, fi56)
                    v57 := GenS10{F0: 1000}
                    _ = v57
                }
                v54 -= v54
                v49 += 42
                _ = fi56
            }
            if (v38 < v38) {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                fmt.Println(fi55, fi55, v38)
                delete(v5, 760)
            }
            _ = fi55
        }
        v58 := GenS11{F0: -101, F1: 42}
        _ = v54
        _ = v58
    }
    if false {
        v59 := GenS12{F0: 2, F1: 0}
        if (true || false) {
            var v60 uint32 = 42
            if true {
                fmt.Println(v49)
            }
            _ = v60
        } else {
            var v61 uint = 843
            var v62 float32 = 0.1
            _ = v61
            _ = v62
        }
        v49 = (^(v38))
        _ = v59
    }
    for fi63 := 0; fi63 < 20; fi63++ {
        v64 := GenS3{F0: 0, F1: 42, F2: 42}
        _ = fi63
        _ = v64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
