package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string; F1 uint64; F2 int64 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 int32; F2 uint }
type GenS7 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 uint64) bool {
    var v4 int16 = -576
    p0 -= p0
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(fi5, v4)
        v6 := GenS1{F0: 127, F1: (81 != 55)}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 142, F1: true}
    v8 := GenS1{F0: 108, F1: true}
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 float32 = 0.0
    v2 := []string{("hello" + "a"), ("a" + "\n\t")}
    var v3 uint32 = 2147483647
    v9 := GenS2{F0: "a", F1: 108, F2: -1}
    for fi10 := 0; fi10 < 2; fi10++ {
        if true {
            if false {
                v3 = 76
            } else {
                v1 = 0.1
                v3 = v3
                var v11 bool = false
                _ = v11
            }
            fmt.Println(fi10)
            if true {
                if len(v2) > 0 {
                    v2[0] = ("a" + "")
                }
                var v12 float32 = 0.0
                v12 -= 1.5
                _ = v12
            } else {
                v13 := GenS3{F0: 98, F1: 421}
                for fi14 := 0; fi14 < 5; fi14++ {
                    v3 -= 895
                    _ = fi14
                }
                _ = v13
            }
        }
        v15 := GenS1{F0: 128, F1: true}
        _ = fi10
        _ = v15
    }
    var v16 int8 = -23
    for fi17 := 0; fi17 < 19; fi17++ {
        if (-2.25 < ((3.14159 + -2.25) / float64(v1))) {
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS2{F0: "hello", F1: 42, F2: 0}
                v20 := GenS4{F0: 32767, F1: 0, F2: 1}
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            v21 := GenS3{F0: (v16 ^ (v16 & (v16 << ((872) & 63)))), F1: 4294967295}
            v22 := GenS2{F0: "hello", F1: 529, F2: -691}
            var v23 int8 = -128
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi17
    }
    if false {
        v24 := GenS4{F0: 42, F1: 0, F2: 65535}
        fmt.Println(v16, v1)
        if false {
            for fi25 := 0; fi25 < 6; fi25++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
                if true {
                    var v26 uint16 = 635
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v27 := GenS5{F0: (v3 % ((v3) | 1))}
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS4{F0: 255, F1: -9223372036854775808, F2: 42}
                    _ = v28
                }
                _ = fi25
            }
            v3 += 0
            var v29 float64 = ((3.14159 - (-2.25 + 1e10)) + 1e10)
            _ = v29
        }
        _ = v24
    } else {
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 int8 = int8(v3)
            if (127 < 127) {
                var v32 int8 = 0
                var v33 string = "hello"
                var v34 int8 = -45
                _ = v32
                _ = v33
                _ = v34
            } else {
                var v35 uint8 = 2
                _ = v35
            }
            var v36 uint8 = 2
            _ = fi30
            _ = v31
            _ = v36
        }
        v37 := GenS6{F0: ("\n\t" + "foo bar"), F1: -1, F2: 435}
        _ = v37
    }
    if (!(false && false)) {
        if (!true) {
            for fi38 := 0; fi38 < 11; fi38++ {
                v3 = v3
                if (42 < 54) {
                    v39 := GenS1{F0: 59, F1: (3.14159 == float64((v1 * (0.1 / (0.1 * 1.5)))))}
                    for fi40 := 0; fi40 < 8; fi40++ {
                        v1 += -2.25
                        v2 = append(v2, "a")
                        _ = fi40
                    }
                    _ = v39
                }
                v2 = append(v2, "\n\t")
                _ = fi38
            }
            if true {
                if ((v3 <= (v3 % ((2147483647) | 1))) || false) {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                }
                switch 32767 {
                case 128:
                    var v41 uint8 = 2
                    _ = v41
                case 32767:
                    for fi42 := 0; fi42 < 16; fi42++ {
                        v43 := GenS7{F0: -243, F1: 255, F2: 79}
                        v3 = v3
                        fmt.Println(v1, fi42)
                        _ = fi42
                        _ = v43
                    }
                }
                switch -547 {
                case 127:
                    switch (("foo bar" + "a") + "foo bar") {
                    case "\n\t":
                        var v44 int32 = 0
                        _ = v44
                    default:
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    }
                default:
                    v1 += v1
                }
            } else {
                v45 := GenS4{F0: 255, F1: -9223372036854775808, F2: 2}
                _ = v45
            }
        }
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        var v47 bool = (v3 < v3)
        var v48 float64 = 1e10
        _ = fi46
        _ = v47
        _ = v48
    }
    if false {
        v16 = (-(v16))
        v16 = 1
        for fi49 := 0; fi49 < 4; fi49++ {
            if true {
                _ = len(v2)
            } else {
                v50 := GenS8{F0: 1}
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = v50
            }
            fmt.Println(v16)
            _ = fi49
        }
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    if (true && ((v3 ^ 10) > 2147483647)) {
        v51 := GenS4{F0: 65535, F1: 172, F2: 0}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v51
    } else {
        var v52 int8 = v16
        var v53 uint32 = v3
        _ = v52
        _ = v53
    }
    v16 = v16
    var v54 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v54
    _ = v9
}
