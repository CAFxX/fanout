package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 int64; F2 uint }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint32; F2 int8 }
func genFn0(p0 int16) bool {
    switch ("a" + ("foo bar" + ("hello" + "a"))) {
    case "":
        fmt.Println(p0)
    }
    v6 := []int{-963, 128}
    if (!false) {
        var v7 bool = false
        v6 = append(v6, 32768)
        fmt.Println(p0, p0)
        _ = v7
    } else {
        var v8 int32 = 1
        if (-2147483648 > (v8 - 0)) {
            if false {
                fmt.Println(p0)
                v9 := GenS1{F0: 18446744073709551615}
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v9
            } else {
                fmt.Println(v8, p0)
            }
        } else {
            var v10 bool = (true || (65535 != 2))
            var v11 uint = 89
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 4294967295}
        _ = v12
        _ = v8
    }
    p0 += (p0 >> ((255) & 63))
    fmt.Println(p0)
    _ = p0
    _ = v6
    return true
}
func genFn1() string {
    if ("" >= "foo bar") {
        var v13 int64 = 0
        v14 := GenS1{F0: 288}
        var v15 int16 = 1000
        _ = v13
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 20; fi16++ {
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 int8 = -1
                v19 := []int{fi17}
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 int8 = 0
            var v21 int32 = 2147483647
            _ = fi16
            _ = v20
            _ = v21
        }
    }
    var v22 float64 = 3.14159
    v23 := GenS2{F0: -503}
    if false {
        var v24 int = -838
        _ = v24
    }
    _ = v22
    _ = v23
    return ("\n\t" + "a")
}

func main() {
    var v1 int = 128
    var v2 int16 = 514
    var v3 int16 = v2
    var v4 float32 = 3.14159
    var v5 int32 = 1
    if true {
        v1 += v1
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 string = "foo bar"
                var v28 uint32 = 758
                _ = fi26
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 15; fi29++ {
                switch 0 {
                case 9223372036854775807:
                    v30 := GenS3{F0: 127}
                    _ = v30
                case -1:
                    fmt.Println(fi29, fi29)
                case 32767:
                    fmt.Println(fi25)
                }
                v31 := GenS3{F0: 251}
                _ = fi29
                _ = v31
            }
            var v32 uint8 = 57
            _ = fi25
            _ = v32
        }
    }
    var v33 int32 = v5
    var v34 int16 = v2
    v5 = v33
    if true {
        v35 := GenS4{F0: (-(v34)), F1: 829, F2: 264}
        var v36 uint16 = 32767
        if true {
            fmt.Println(v36)
            v37 := []int{}
            var v38 int64 = 267
            _ = v37
            _ = v38
        } else {
            if true {
                for fi39 := 0; fi39 < 14; fi39++ {
                    switch ("a" + "hello") {
                    case "foo bar":
                        var v40 uint16 = v36
                        _ = v40
                    case "hello":
                        for fi41 := 0; fi41 < 16; fi41++ {
                            v4 = (float32(0.1) - 0.1)
                            v3 = (v3 ^ (v2 >> ((174) & 63)))
                            v42 := GenS2{F0: v3}
                            v43 := GenS5{F0: 1, F1: 18446744073709551615}
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    case "":
                        for fi44 := 0; fi44 < 19; fi44++ {
                            v34 += v34
                            if false {
                                switch "\n\t" {
                                case "hello":
                                    var v45 int64 = 1
                                    _ = v45
                                case "\n\t":
                                    var v46 float64 = (-2.25 * (0.1 + (0.0 / 1.0)))
                                    _ = v46
                                case "":
                                    switch (^(fi39)) {
                                    case 42:
                                        v4 = v4
                                    default:
                                        v47 := GenS5{F0: 399, F1: 167}
                                        _ = v47
                                    }
                                default:
                                    var v48 uint = 0
                                    _ = v48
                                }
                                v5 = v5
                            } else {
                                if (false || true) {
                                    switch ("" + "hello") {
                                    case "hello":
                                        v2 += 32767
                                    case "a":
                                        v49 := GenS6{F0: 261, F1: 1, F2: 42}
                                        _ = v49
                                    case "":
                                        v50 := []int{42}
                                        _ = v50
                                    }
                                    var v51 int16 = v3
                                    _ = v51
                                } else {
                                    fmt.Println(v34)
                                }
                                v52 := []int{}
                                fmt.Println(v4, v3)
                                v33 += (-(v33))
                                v52 = append(v52, fi44)
                                _ = v52
                            }
                            v53 := GenS1{F0: 472}
                            _ = fi44
                            _ = v53
                        }
                    default:
                        switch v1 {
                        case -588:
                            fmt.Println(v1, v33)
                        case 32768:
                            v54 := []int{fi39}
                            _ = v54
                        }
                    }
                    _ = fi39
                }
                var v55 uint = 255
                _ = v55
            }
        }
        _ = v35
        _ = v36
    }
    fmt.Println(v4, v2, v3)
    v4 = 0.1
    var v56 int64 = 653
    fmt.Println(v34)
    v33 = 143
    v57 := GenS6{F0: 65535, F1: 0, F2: 0}
    var v58 int16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
}
