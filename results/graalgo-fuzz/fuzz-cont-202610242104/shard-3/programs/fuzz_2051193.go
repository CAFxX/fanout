package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint32; F2 int }
func genFn0(p0 int8, p1 string, p2 int8) string {
    if (!(-859 < 42)) {
        if true {
            p2 -= (p2 >> ((1) & 63))
        } else {
            for fi7 := 0; fi7 < 3; fi7++ {
                for fi8 := 0; fi8 < 10; fi8++ {
                    p2 -= p0
                    _ = fi8
                }
                _ = fi7
            }
            p2 = p2
            v9 := GenS1{F0: 65535}
            _ = v9
        }
        switch -1 {
        case -258:
            for fi10 := 0; fi10 < 12; fi10++ {
                for fi11 := 0; fi11 < 2; fi11++ {
                    if false {
                        var v12 int = (fi10 / ((128) | 1))
                        p0 = p0
                        v12 = (fi10 >> ((18446744073709551615) & 63))
                        _ = v12
                    }
                    _ = fi11
                }
                _ = fi10
            }
        case -537:
            v13 := GenS2{F0: 127}
            _ = v13
        case 1:
            switch (((("foo bar" + "") + (p1 + (p1 + p1))) + p1) + p1) {
            case "":
                if false {
                    p0 += p0
                    for fi14 := 0; fi14 < 5; fi14++ {
                        v15 := GenS2{F0: 0}
                        for fi16 := 0; fi16 < 19; fi16++ {
                            var v17 uint16 = 1
                            switch p1 {
                            case "a":
                                switch p1 {
                                case "\n\t":
                                    v18 := GenS1{F0: 1}
                                    _ = v18
                                case "hello":
                                    if (p1 > (("\n\t" + "") + p1)) {
                                        var v19 int64 = -1
                                        var v20 uint32 = 957
                                        v19 = (v19 ^ v19)
                                        _ = v19
                                        _ = v20
                                    }
                                }
                            default:
                                v21 := []int{fi14, 32768}
                                _ = v21
                            }
                            _ = fi16
                            _ = v17
                        }
                        _ = fi14
                        _ = v15
                    }
                }
            case "a":
                v22 := GenS1{F0: 683}
                _ = v22
            default:
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 uint8 = 42
                    var v25 uint16 = 32767
                    v25 -= (v25 - (^(v25)))
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        default:
            p0 = p2
        }
        v26 := GenS3{F0: false, F1: 42}
        _ = v26
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 int64 = -138
            _ = fi27
            _ = v28
        }
        var v29 uint64 = 889
        _ = v29
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 uint8 = 2
        v32 := []int{fi30}
        _ = fi30
        _ = v31
        _ = v32
    }
    if (p1 <= p1) {
        if (-9223372036854775808 > -1) {
            if false {
                var v33 bool = (1e10 < (3.14159 / -2.25))
                v33 = true
                _ = v33
            }
            v34 := []int{32768, -2147483648}
            var v35 uint = 517
            _ = v34
            _ = v35
        }
        p1 = ("foo bar" + p1)
        var v36 uint8 = 2
        _ = v36
    }
    if (false && false) {
        v37 := []int{127}
        if (!false) {
            p0 = (p2 - -108)
        }
        _ = v37
    } else {
        v38 := GenS2{F0: 2}
        v39 := []int{896, 42, -2147483648, 2147483647}
        _ = v38
        _ = v39
    }
    v40 := GenS4{F0: 127, F1: 1000}
    _ = p0
    _ = p1
    _ = p2
    _ = v40
    return p1
}
func genFn1(p0 uint64) uint8 {
    for fi41 := 0; fi41 < 5; fi41++ {
        var v42 uint = 42
        var v43 float64 = float64(3.14159)
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 float64 = 3.14159
    v45 := GenS5{F0: 136}
    v46 := GenS2{F0: 42}
    _ = p0
    _ = v44
    _ = v45
    _ = v46
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 0
    var v3 uint64 = 2
    var v4 uint64 = (v3 ^ (^(v3)))
    v5 := []string{"foo bar", "", (("\n\t" + ("a" + ((("\n\t" + ("a" + "foo bar")) + "") + "a"))) + "\n\t"), ""}
    var v6 int8 = -116
    if false {
        v3 = (v4 * 4294967295)
        v3 += (^(v4))
    } else {
        v3 -= (^(v4))
    }
    v3 -= (v4 >> ((1) & 63))
    v47 := GenS4{F0: 4294967295, F1: v2}
    switch "hello" {
    case "a":
        v48 := GenS6{F0: ("\n\t" + "\n\t"), F1: 65535}
        _ = v48
    case "foo bar":
        for fi49 := 0; fi49 < 17; fi49++ {
            fmt.Println(v3, v6, v3)
            for fi50 := 0; fi50 < 1; fi50++ {
                fmt.Println(fi49, fi50)
                v51 := GenS3{F0: (true || (2147483647 > 752)), F1: 1}
                v2 = v2
                if (42 < 595) {
                    v6 = (v6 ^ v6)
                    var v52 uint64 = (^(v4))
                    if (uint32(v52) < 916) {
                        var v53 int32 = 1
                        _ = len(v1)
                        var v54 int64 = 42
                        _ = v53
                        _ = v54
                    }
                    _ = v52
                } else {
                    if false {
                        fmt.Println(fi50, v3)
                    }
                }
                _ = fi50
                _ = v51
            }
            _ = fi49
        }
    }
    var v55 uint64 = (v3 ^ v3)
    var v56 uint = 1
    v4 = v55
    switch "hello" {
    case "a":
        v57 := GenS7{F0: 0, F1: 65535, F2: 32768}
        _ = v57
    default:
        if false {
            _ = len(v5)
            var v58 string = "hello"
            v3 -= 123
            _ = v58
        } else {
            for fi59 := 0; fi59 < 12; fi59++ {
                var v60 string = ""
                _ = fi59
                _ = v60
            }
            fmt.Println(v6, v55)
            var v61 uint16 = 65535
            _ = v61
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v55
    _ = v56
    _ = v6
}
