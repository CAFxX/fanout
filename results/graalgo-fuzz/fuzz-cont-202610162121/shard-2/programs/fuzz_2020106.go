package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int16; F2 int16 }
func genFn0(p0 uint16) int32 {
    v5 := GenS1{F0: (("\n\t" + "") + "foo bar"), F1: -86, F2: 2}
    v6 := []int{}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int16 = int16(p0)
        for fi9 := 0; fi9 < 3; fi9++ {
            switch ("" + "hello") {
            case "":
                v10 := GenS1{F0: "", F1: -1, F2: 2}
                _ = v10
            case "hello":
                if false {
                    var v11 float64 = 1e10
                    var v12 bool = false
                    _ = v11
                    _ = v12
                } else {
                    if false {
                        if (-128 == 0) {
                            fmt.Println(fi9, fi9, p0)
                        }
                        p0 = p0
                        if (0 >= 0) {
                            v13 := GenS1{F0: "a", F1: 127, F2: 1}
                            p0 = (p0 % ((p0) | 1))
                            var v14 int = fi7
                            _ = v13
                            _ = v14
                        }
                    }
                    if true {
                        v15 := GenS1{F0: "foo bar", F1: 42, F2: 196}
                        switch fi7 {
                        case -9223372036854775808:
                            v8 = (v8 - v8)
                        default:
                            if len(v6) > 0 {
                                v6[0] = (fi7 ^ (fi9 << ((18446744073709551615) & 63)))
                            }
                        }
                        var v16 bool = (!((fi9 >> ((1) & 63)) > (fi7 + (fi7 % ((439) | 1)))))
                        _ = v15
                        _ = v16
                    } else {
                        if false {
                            var v17 int = (^(fi7))
                            _ = v17
                        } else {
                            fmt.Println(v8)
                        }
                        v8 = v8
                        p0 = (p0 * p0)
                        if len(v6) > 0 {
                            v6[0] = (^(fi9))
                        }
                        p0 -= 0
                        var v18 uint16 = (^(p0))
                        var v19 float32 = -2.25
                        _ = v18
                        _ = v19
                    }
                    p0 = (p0 & p0)
                    if false {
                        if len(v6) > 0 {
                            v6[0] = fi9
                        }
                        v8 = -369
                    }
                }
            }
            switch ("hello" + "hello") {
            case "hello":
                v20 := GenS2{F0: 65535}
                _ = v20
            case "":
                var v21 int8 = -20
                _ = v21
            case "foo bar":
                _ = len(v6)
            }
            p0 = (^(p0))
            _ = fi9
        }
        var v22 uint = 237
        _ = fi7
        _ = v22
        _ = v8
    }
    switch "foo bar" {
    case "hello":
        p0 = (^(p0))
    case "\n\t":
        if true {
            fmt.Println(p0)
            v23 := GenS3{F0: -9223372036854775808}
            _ = v23
        } else {
            fmt.Println(p0)
            var v24 uint16 = p0
            _ = v24
        }
    default:
        var v25 float64 = (((3.14159 / 1.5) + ((0.1 * 3.14159) * -2.25)) + 1e10)
        _ = v25
    }
    _ = p0
    _ = v5
    _ = v6
    return 839
}
func genFn1() int {
    var v26 uint = 4294967295
    v27 := []int{9223372036854775807, -1}
    _ = v26
    _ = v27
    return 4294967295
}

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]string)
    var v3 float64 = 3.14159
    var v4 uint64 = 42
    for fi28 := 0; fi28 < 12; fi28++ {
        switch ("" + "a") {
        case "hello":
            switch ("foo bar" + ("a" + "\n\t")) {
            case "hello":
                switch "\n\t" {
                case "hello":
                    switch ("hello" + ((("a" + "a") + "foo bar") + "")) {
                    case "foo bar":
                        var v29 uint = 311
                        _ = v29
                    case "\n\t":
                        fmt.Println(v4)
                    }
                default:
                    v30 := GenS3{F0: (v1 / ((v1) | 1))}
                    _ = v30
                }
            case "foo bar":
                switch (("\n\t" + ("a" + "a")) + "foo bar") {
                case "a":
                    var v31 int8 = -1
                    _ = v31
                case "hello":
                    var v32 uint32 = 2147483647
                    _ = v32
                case "foo bar":
                    v33 := GenS3{F0: v1}
                    _ = v33
                default:
                    var v34 bool = false
                    _ = v34
                }
            case "\n\t":
                switch fi28 {
                case 458:
                    v35 := []int{fi28, (fi28 + (fi28 - -2147483648)), (-(fi28))}
                    _ = v35
                case -9223372036854775808:
                    v36 := []int{(fi28 ^ fi28), fi28}
                    _ = v36
                default:
                    for fi37 := 0; fi37 < 11; fi37++ {
                        var v38 int8 = 127
                        v38 = (^(v38))
                        v1 -= int64(v4)
                        _ = len(v2)
                        _ = fi37
                        _ = v38
                    }
                }
            default:
                switch genFn1() {
                case 32767:
                    v39 := GenS4{F0: -2147483648}
                    _ = v39
                case 9223372036854775807:
                    var v40 uint64 = v4
                    _ = v40
                case 42:
                    v1 += -1
                }
            }
        case "foo bar":
            switch (fi28 + genFn1()) {
            case 4294967295:
                var v41 float32 = 1.5
                _ = v41
            }
        }
        switch "foo bar" {
        case "\n\t":
            var v42 float32 = 0.0
            _ = v42
        case "hello":
            if true {
                fmt.Println(v4, v3, v1)
                v43 := GenS5{F0: -54, F1: 1000, F2: -151}
                fmt.Println(v1, v3, v4)
                _ = v43
            }
        }
        _ = fi28
    }
    var v44 int16 = 558
    v45 := GenS3{F0: -1}
    var v46 bool = true
    if v46 {
        var v47 uint8 = 13
        switch genFn1() {
        case -2147483648:
            fmt.Println(v4, v47)
        case -1:
            var v48 uint64 = (v4 | v4)
            _ = v48
        case 4294967295:
            v4 += v4
        }
        _ = v47
    } else {
        for fi49 := 0; fi49 < 19; fi49++ {
            var v50 uint = 304
            var v51 int = fi49
            _ = fi49
            _ = v50
            _ = v51
        }
        for fi52 := 0; fi52 < 3; fi52++ {
            var v53 float64 = float64((float32(float64(0.1)) * (1e10 - 3.14159)))
            switch ("a" + ("a" + ("a" + "foo bar"))) {
            case "foo bar":
                v54 := v2[2147483647]
                _ = v54
            case "\n\t":
                v44 -= v44
            case "":
                v55 := GenS4{F0: -897}
                _ = v55
            }
            _ = fi52
            _ = v53
        }
    }
    var v56 uint = 261
    for fi57 := 0; fi57 < 9; fi57++ {
        v44 = v44
        _ = fi57
    }
    var v58 int16 = v44
    fmt.Println(v46)
    var v59 uint64 = v4
    for fi60 := 0; fi60 < 4; fi60++ {
        v56 = (v56 / ((v56) | 1))
        _ = fi60
    }
    var v61 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v56
    _ = v58
    _ = v59
    _ = v61
}
