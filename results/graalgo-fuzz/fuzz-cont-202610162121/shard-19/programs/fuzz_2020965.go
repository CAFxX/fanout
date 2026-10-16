package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint16; F1 int }

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 int = 32768
    v1 = v1
    if false {
        v4 += v4
        for fi5 := 0; fi5 < 3; fi5++ {
            delete(v3, "foo bar")
            _ = fi5
        }
        v1 = (!false)
    }
    switch "a" {
    case "foo bar":
        switch "" {
        case "":
            switch "foo bar" {
            case "a":
                var v6 uint = 0
                _ = v6
            default:
                fmt.Println(v1, v4)
            }
        case "hello":
            var v7 float32 = (3.14159 / 1e10)
            _ = v7
        default:
            if v1 {
                if v1 {
                    v4 = v4
                } else {
                    fmt.Println(v4, v4)
                    v8 := GenS1{F0: 4294967295, F1: 4294967295, F2: 0}
                    v9 := []int{32768}
                    _ = v8
                    _ = v9
                }
                v10 := GenS2{F0: 0, F1: false, F2: 65535}
                _ = v10
            } else {
                v1 = (!((((0.1 + 0.1) * (-2.25 - ((-2.25 - -2.25) + 0.1))) * -2.25) > (float32(1e10) - -2.25)))
                for fi11 := 0; fi11 < 8; fi11++ {
                    v4 = (v4 + fi11)
                    _ = fi11
                }
            }
        }
    case "":
        if true {
            var v12 bool = true
            v1 = v1
            switch (v4 << ((0) & 63)) {
            case -9223372036854775808:
                if v12 {
                    var v13 float32 = 1.5
                    var v14 int16 = 0
                    var v15 uint = 254
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case 440:
                if v1 {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        if (v1 || (!(!true))) {
                            fmt.Println(v4)
                            var v17 uint8 = 42
                            if (!v1) {
                                v18 := GenS3{F0: fi16, F1: "foo bar"}
                                fmt.Println(v12, v17, v1)
                                _ = v18
                            }
                            _ = v17
                        } else {
                            var v19 uint = 405
                            var v20 uint = v19
                            var v21 int16 = 1000
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        switch fi16 {
                        case -244:
                            var v22 int8 = 100
                            _ = v22
                        case -215:
                            v23 := []int{(fi16 * 882)}
                            _ = v23
                        default:
                            var v24 int16 = 336
                            _ = v24
                        }
                        _ = fi16
                    }
                    var v25 uint = 0
                    for fi26 := 0; fi26 < 14; fi26++ {
                        v1 = v1
                        _ = fi26
                    }
                    _ = v25
                } else {
                    v27 := GenS4{F0: 2}
                    _ = v27
                }
            default:
                if v1 {
                    v28 := []int{(v4 ^ (v4 ^ (v4 ^ v4)))}
                    _ = v28
                } else {
                    fmt.Println(v1)
                }
            }
            _ = v12
        } else {
            switch 984 {
            case 32767:
                var v29 uint32 = 2147483647
                _ = v29
            case 502:
                for fi30 := 0; fi30 < 4; fi30++ {
                    fmt.Println(fi30, fi30)
                    _ = fi30
                }
            case 72:
                v4 = 602
            }
        }
    case "\n\t":
        if (!true) {
            switch v4 {
            case -2147483648:
                var v31 uint64 = 210
                _ = v31
            case -383:
                fmt.Println(v4, v1)
            case 32767:
                v1 = (0.1 <= float64((((1.5 / (0.0 - 1.5)) * 0.1) * (0.1 + -2.25))))
            default:
                switch "foo bar" {
                case "hello":
                    if v1 {
                        if true {
                            v2["a"] = v4
                        } else {
                            v32 := v3["a"]
                            v33 := GenS2{F0: 32767, F1: (float32(0.0) == 0.1), F2: 1}
                            v34 := GenS5{F0: 2, F1: 32767}
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        var v35 int8 = 42
                        _ = v35
                    } else {
                        switch 128 {
                        case -2147483648:
                            v36 := GenS6{F0: 242, F1: 18446744073709551615}
                            _ = v36
                        }
                    }
                case "foo bar":
                    v4 -= 32768
                }
            }
            v4 = 0
        } else {
            if (v1 || ("a" != ("a" + "a"))) {
                var v37 int32 = 871
                var v38 uint = 42
                _ = v37
                _ = v38
            }
        }
    }
    v1 = (-1 >= 0)
    v1 = (((65535 != 42) || (3.14159 > 0.1)) && v1)
    v1 = ((1.5 / ((0.0 - 1e10) / 3.14159)) >= 0.1)
    if v1 {
        if (false && true) {
            v39 := GenS4{F0: 42}
            if false {
                v40 := v2[""]
                v1 = (v1 || (277 != 65535))
                var v41 uint32 = 2147483647
                _ = v40
                _ = v41
            } else {
                delete(v2, ("" + ("\n\t" + ("hello" + "foo bar"))))
                var v42 float64 = (3.14159 + 0.0)
                _ = v42
            }
            if v1 {
                if v1 {
                    v43 := []int{(v4 | (v4 - v4)), 4294967295, v4}
                    v4 = (-(v4))
                    _ = v43
                }
                fmt.Println(v1, v1)
            } else {
                if v1 {
                    v1 = v1
                    switch -999 {
                    case 42:
                        v44 := GenS7{F0: 733, F1: v4}
                        _ = v44
                    case 32767:
                        v2["\n\t"] = (v4 % ((v4) | 1))
                    }
                    var v45 float32 = 0.1
                    _ = v45
                } else {
                    var v46 uint8 = 128
                    var v47 uint32 = 2147483648
                    _ = v46
                    _ = v47
                }
            }
            _ = v39
        }
        v48 := GenS2{F0: -666, F1: (30 == 208), F2: 2147483647}
        _ = v48
    }
    if false {
        var v49 uint = 1
        v4 = 2147483647
        _ = v49
    } else {
        var v50 int16 = -32768
        _ = v50
    }
    var v51 float64 = 3.14159
    switch ("hello" + "\n\t") {
    case "":
        v52 := GenS3{F0: (-(v4)), F1: ("" + "")}
        _ = v52
    case "\n\t":
        var v53 int16 = -764
        _ = v53
    case "foo bar":
        switch v4 {
        case 32768:
            var v54 int = -1
            _ = v54
        case 870:
            v55 := GenS7{F0: 416, F1: v4}
            _ = v55
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
}
