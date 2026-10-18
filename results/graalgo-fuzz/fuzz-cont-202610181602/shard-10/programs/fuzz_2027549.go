package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 string; F2 string }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint64; F1 int; F2 uint8 }
func genFn0(p0 uint64, p1 uint8, p2 int32) int8 {
    v4 := []int{-9223372036854775808, 1}
    switch ("hello" + "hello") {
    case "":
        v5 := GenS1{F0: "hello"}
        _ = v5
    case "foo bar":
        var v6 float32 = 1.5
        _ = v6
    default:
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: "hello"}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi7
            _ = v8
        }
    }
    if (3.14159 == 1e10) {
        switch "\n\t" {
        case "hello":
            for fi9 := 0; fi9 < 3; fi9++ {
                p2 = 0
                if (2147483647 == 782) {
                    fmt.Println(p0, fi9, p1)
                } else {
                    switch fi9 {
                    case 0:
                        if len(v4) > 0 {
                            v4[0] = fi9
                        }
                    default:
                        for fi10 := 0; fi10 < 8; fi10++ {
                            p1 = p1
                            _ = fi10
                        }
                    }
                }
                p2 -= p2
                _ = fi9
            }
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        var v11 uint16 = 42
        if (!((true && (float64(-2.25) < 0.0)) || true)) {
            v12 := GenS2{F0: -9223372036854775808, F1: (v11 < (^(v11)))}
            v13 := GenS3{F0: 0, F1: 255}
            var v14 float64 = 0.0
            _ = v12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    var v15 float32 = (float32(-2.25) + 1.5)
    switch (("foo bar" + "foo bar") + ("hello" + "a")) {
    case "\n\t":
        if (42 > -273) {
            for fi16 := 0; fi16 < 7; fi16++ {
                fmt.Println(v15, fi16, p2)
                for fi17 := 0; fi17 < 4; fi17++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v18 bool = false
                    _ = fi17
                    _ = v18
                }
                fmt.Println(p1, fi16)
                _ = fi16
            }
        }
    case "hello":
        if len(v4) > 0 {
            v4[0] = -39
        }
    case "foo bar":
        for fi19 := 0; fi19 < 16; fi19++ {
            p2 -= (p2 * 2147483647)
            _ = fi19
        }
    default:
        var v20 float64 = 0.1
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v4
    return -105
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    v3 := make(map[string]bool)
    for fi21 := 0; fi21 < 11; fi21++ {
        if false {
            fmt.Println(v2, fi21)
            for fi22 := 0; fi22 < 14; fi22++ {
                v2 = ((("hello" + "hello") + ("hello" + "a")) != "")
                v23 := []int{fi21, fi22}
                v24 := GenS1{F0: "a"}
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        if (!v2) {
            delete(v3, "foo bar")
        }
        _ = fi21
    }
    v2 = ((1 == 495) && v2)
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 uint32 = 65535
            _ = fi26
            _ = v27
        }
        v2 = (0 > 0)
        v28 := GenS4{F0: 2, F1: (454 == 1)}
        _ = fi25
        _ = v28
    }
    v2 = v2
    _ = len(v1)
    if (-249 != 240) {
        v29 := GenS5{F0: (0.1 >= (0.1 * 1e10))}
        _ = v29
    }
    v30 := []int{}
    for fi31 := 0; fi31 < 10; fi31++ {
        if (!v2) {
            if true {
                if false {
                    var v32 uint16 = 938
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 18; fi33++ {
                        var v34 float32 = 1.5
                        var v35 int8 = 1
                        if false {
                            v35 -= (v35 << ((18446744073709551615) & 63))
                            v35 += v35
                            v36 := GenS6{F0: ((!v2) && v2), F1: "hello", F2: ("foo bar" + (("hello" + "hello") + "a"))}
                            _ = v36
                        }
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
                v37 := GenS7{F0: 2, F1: 1000}
                _ = v37
            } else {
                v38 := GenS6{F0: v2, F1: ("a" + "hello"), F2: ("" + "a")}
                for fi39 := 0; fi39 < 9; fi39++ {
                    v40 := GenS8{F0: 9223372036854775807, F1: 128}
                    _ = fi39
                    _ = v40
                }
                var v41 int32 = -248
                _ = v38
                _ = v41
            }
            fmt.Println(v2, fi31)
        } else {
            v2 = (42 != 255)
        }
        v3["foo bar"] = (float64(((1e10 / (1e10 / 0.1)) * 1.5)) >= -2.25)
        v2 = v2
        _ = fi31
    }
    switch 42 {
    case 42:
        if v2 {
            v2 = (v2 || (!(v2 || (false || v2))))
        } else {
            fmt.Println(v2)
        }
    case 32768:
        switch 9223372036854775807 {
        case -563:
            fmt.Println(v2)
        }
    case 0:
        fmt.Println(v2)
    default:
        if (2 < 255) {
            for fi42 := 0; fi42 < 10; fi42++ {
                var v43 bool = (2 > 503)
                v2 = (v2 && ((0.1 >= -2.25) && (!v43)))
                _ = len(v3)
                _ = fi42
                _ = v43
            }
            var v44 uint8 = 69
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            _ = v44
        } else {
            if v2 {
                var v45 uint64 = 0
                _ = v45
            }
        }
    }
    for fi46 := 0; fi46 < 3; fi46++ {
        var v47 int32 = 1
        if v2 {
            if false {
                if v2 {
                    v2 = (891 > 65535)
                    switch (fi46 & (fi46 ^ (fi46 ^ 872))) {
                    case 32768:
                        fmt.Println(fi46, fi46, v47)
                    default:
                        v2 = ((float64(0.1) / 1.0) < 1e10)
                    }
                }
                fmt.Println(v2, v2)
            }
        } else {
            v48 := GenS4{F0: 674, F1: false}
            _ = v48
        }
        for fi49 := 0; fi49 < 10; fi49++ {
            fmt.Println(v47)
            if (596 != 1) {
                v50 := v3["hello"]
                _ = v50
            } else {
                v51 := GenS9{F0: 42, F1: fi49, F2: 255}
                if (-1 != 42) {
                    v52 := v1[(-(fi49))]
                    var v53 uint32 = 2147483647
                    _ = v52
                    _ = v53
                }
                switch (fi49 - fi46) {
                case 271:
                    if v2 {
                        if false {
                            var v54 uint8 = 0
                            _ = v54
                        }
                        v47 = int32(fi49)
                        var v55 float32 = (0.0 * 1e10)
                        _ = v55
                    } else {
                        v47 = -2147483648
                    }
                case -2147483648:
                    switch (fi46 ^ -623) {
                    case 127:
                        var v56 bool = v2
                        _ = v56
                    case 645:
                        v2 = v2
                    default:
                        v47 = 42
                    }
                case 32767:
                    v2 = v2
                }
                v47 = v47
                if v2 {
                    fmt.Println(fi46, v2)
                } else {
                    fmt.Println(fi46, v47)
                }
                _ = v51
            }
            if v2 {
                if v2 {
                    v2 = (4294967295 == 723)
                    if (false && v2) {
                        var v57 uint = 65535
                        _ = v57
                    } else {
                        if v2 {
                            var v58 int = fi49
                            _ = v58
                        } else {
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                        }
                        var v59 uint32 = 303
                        v59 += (^(v59))
                        v59 = v59
                        _ = v59
                    }
                    v2 = (true && (!((v47 | (v47 - v47)) <= -413)))
                }
                var v60 string = ""
                _ = v60
            }
            _ = fi49
        }
        _ = fi46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
