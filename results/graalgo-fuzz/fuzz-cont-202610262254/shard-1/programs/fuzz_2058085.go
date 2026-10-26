package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
type GenS2 struct { F0 int32; F1 int32; F2 int8 }
type GenS3 struct { F0 uint32; F1 string; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 uint16; F2 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 int8; F2 bool }
type GenS8 struct { F0 int16; F1 int64; F2 bool }
type GenS9 struct { F0 int32 }
func genFn0(p0 bool, p1 uint64) int8 {
    p0 = false
    var v5 int32 = -2147483648
    switch 42 {
    case 42:
        v6 := GenS1{F0: -9223372036854775808, F1: 42, F2: 0}
        _ = v6
    case 128:
        var v7 float64 = -2.25
        _ = v7
    case 127:
        for fi8 := 0; fi8 < 19; fi8++ {
            v5 = 857
            fmt.Println(p0, fi8)
            v5 -= 0
            _ = fi8
        }
    default:
        v9 := GenS2{F0: (^(v5)), F1: (v5 - (v5 + (v5 >> ((255) & 63)))), F2: 127}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1(p0 uint32, p1 bool) string {
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            if p1 {
                var v12 int = 1
                _ = v12
            } else {
                if (421 != 1) {
                    var v13 float64 = 1e10
                    v13 -= float64(-2.25)
                    p0 = p0
                    _ = v13
                } else {
                    var v14 float32 = float32(3.14159)
                    _ = v14
                }
                v15 := []int{-2147483648}
                var v16 int8 = -108
                if p1 {
                    v17 := GenS3{F0: (p0 << ((169) & 63)), F1: "foo bar", F2: "a"}
                    var v18 uint8 = 255
                    p1 = (p1 || (!(308 >= (p0 >> ((42) & 63)))))
                    _ = v17
                    _ = v18
                }
                var v19 uint32 = (p0 >> ((255) & 63))
                _ = v15
                _ = v16
                _ = v19
            }
            fmt.Println(p0, p0, fi10)
            if ((p1 || (true && (42 < 622))) && true) {
                var v20 string = "foo bar"
                if true {
                    v21 := GenS1{F0: -1, F1: 65535, F2: 1}
                    switch v20 {
                    case "a":
                        if p1 {
                            v22 := []int{(fi10 << ((727) & 63)), 32767, fi10}
                            _ = v22
                        } else {
                            if p1 {
                                var v23 uint = 626
                                var v24 string = v20
                                if ((0.0 + 0.1) == ((float64((((3.14159 - 1e10) / 0.1) + 0.1)) - (3.14159 + 1.5)) * 3.14159)) {
                                    if false {
                                        var v25 int16 = -245
                                        fmt.Println(v25, p0)
                                        var v26 uint8 = uint8(p0)
                                        _ = v25
                                        _ = v26
                                    }
                                } else {
                                    var v27 int16 = 620
                                    if ((^(v23)) == v23) {
                                        p1 = ((-1 >= 42) || true)
                                        p0 = p0
                                        v28 := GenS3{F0: p0, F1: v20, F2: v20}
                                        _ = v28
                                    }
                                    _ = v27
                                }
                                _ = v23
                                _ = v24
                            }
                            if p1 {
                                fmt.Println(p0, fi11, p0)
                            } else {
                                v29 := []int{(fi10 % (((fi11 - (fi11 ^ fi11))) | 1)), -2147483648, fi10, 127}
                                p1 = p1
                                _ = v29
                            }
                        }
                    case "foo bar":
                        v20 = v20
                    case "":
                        var v30 float64 = 1e10
                        _ = v30
                    }
                    v31 := GenS2{F0: 0, F1: -393, F2: genFn0(p1, 1)}
                    _ = v21
                    _ = v31
                } else {
                    if ((p1 && (35 >= (^(p0)))) && p1) {
                        v32 := GenS1{F0: 1, F1: 42, F2: 598}
                        _ = v32
                    } else {
                        if true {
                            fmt.Println(p0)
                            fmt.Println(fi10, fi11, fi10)
                        }
                    }
                }
                _ = v20
            } else {
                if ((0.1 < 0.1) && (9223372036854775807 < -253)) {
                    fmt.Println(p1, p0, fi10)
                    v33 := GenS4{F0: "hello"}
                    v34 := GenS2{F0: 42, F1: 1, F2: genFn0(false, 1)}
                    _ = v33
                    _ = v34
                }
                switch (((("" + ("hello" + "hello")) + "foo bar") + "") + "\n\t") {
                case "":
                    var v35 int16 = 32767
                    _ = v35
                case "\n\t":
                    var v36 int32 = 0
                    _ = v36
                case "foo bar":
                    switch fi11 {
                    case -469:
                        v37 := []int{}
                        _ = v37
                    }
                default:
                    p1 = (!p1)
                }
                v38 := []int{42}
                _ = v38
            }
            _ = fi11
        }
        _ = fi10
    }
    switch "foo bar" {
    case "\n\t":
        for fi39 := 0; fi39 < 20; fi39++ {
            v40 := []int{0, -1, 32767}
            for fi41 := 0; fi41 < 14; fi41++ {
                if (fi39 > fi39) {
                    fmt.Println(fi41, fi41)
                    var v42 int16 = 32767
                    _ = v42
                } else {
                    var v43 float32 = 1.5
                    var v44 uint16 = 0
                    _ = len(v40)
                    _ = v43
                    _ = v44
                }
                _ = fi41
            }
            v45 := GenS5{F0: ("a" + "a"), F1: 32767, F2: 4294967295}
            _ = fi39
            _ = v40
            _ = v45
        }
    case "foo bar":
        v46 := GenS6{F0: -593}
        _ = v46
    case "":
        var v47 uint = 1
        _ = v47
    default:
        v48 := GenS7{F0: 310, F1: 127, F2: p1}
        _ = v48
    }
    p1 = true
    v49 := []int{42, 1, -2147483648}
    v50 := GenS8{F0: -32768, F1: -304, F2: p1}
    _ = p0
    _ = p1
    _ = v49
    _ = v50
    return ("foo bar" + "a")
}

func main() {
    var v1 uint16 = 696
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 int16 = 169
    var v51 int64 = 42
    if false {
        if false {
            fmt.Println(v4)
            for fi52 := 0; fi52 < 2; fi52++ {
                switch (((("\n\t" + "hello") + "") + "\n\t") + ("a" + (("" + "hello") + "\n\t"))) {
                case "a":
                    var v53 int16 = (^(v4))
                    _ = v53
                case "":
                    switch -2147483648 {
                    case -40:
                        var v54 uint64 = 18446744073709551615
                        _ = v54
                    case 9223372036854775807:
                        v55 := []int{530}
                        _ = v55
                    }
                case "foo bar":
                    var v56 int64 = 370
                    _ = v56
                }
                v4 = (v4 % ((v4) | 1))
                _ = fi52
            }
            v1 -= 624
        }
        v57 := GenS4{F0: ("a" + ("hello" + "a"))}
        _ = v57
    } else {
        for fi58 := 0; fi58 < 10; fi58++ {
            fmt.Println(v51, v1)
            if false {
                switch "" {
                case "\n\t":
                    var v59 float32 = -2.25
                    _ = v59
                default:
                    v1 = v1
                }
                v60 := []int{fi58, (fi58 ^ fi58)}
                _ = v60
            } else {
                var v61 uint64 = 672
                _ = v61
            }
            _ = len(v2)
            _ = fi58
        }
        v62 := v3[("foo bar" + "hello")]
        _ = v62
    }
    switch "hello" {
    case "":
        _ = len(v3)
    default:
        if false {
            var v63 float32 = 0.1
            v64 := []int{}
            v65 := GenS1{F0: (v51 - v51), F1: (v1 ^ v1), F2: int64(v1)}
            _ = v63
            _ = v64
            _ = v65
        } else {
            v66 := GenS9{F0: -2147483648}
            _ = v66
        }
    }
    if false {
        var v67 uint32 = 4294967295
        v67 -= v67
        _ = v67
    } else {
        var v68 uint32 = 937
        v1 = v1
        _ = v68
    }
    fmt.Println(v1)
    var v69 uint8 = 0
    fmt.Println(v4)
    v70 := GenS7{F0: v1, F1: -15, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v69
    _ = v70
}
