package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int32; F1 uint32; F2 string }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 int; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool }
func genFn0(p0 int32, p1 uint16, p2 bool) string {
    p0 = (-(p0))
    p2 = (!(!false))
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 1, F1: 4294967295}
        if p2 {
            var v6 int = (fi4 & fi4)
            p2 = p2
            _ = v6
        }
        if p2 {
            fmt.Println(fi4, p2)
            var v7 int = (^(fi4))
            fmt.Println(fi4, v7, p1)
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}
func genFn1(p0 uint8, p1 int32, p2 uint) bool {
    p2 = (^(p2))
    var v8 int32 = (p1 | p1)
    var v9 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]int)
    var v3 uint = 0
    var v10 int8 = 84
    switch 358 {
    case 4294967295:
        var v11 int = 9223372036854775807
        _ = v11
    default:
        for fi12 := 0; fi12 < 4; fi12++ {
            v10 += v10
            if (("foo bar" + "a") > "foo bar") {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v3 = (v3 - v3)
                    _ = fi13
                }
            } else {
                switch "foo bar" {
                case "":
                    var v14 uint64 = 0
                    _ = v14
                }
                v15 := GenS2{F0: -1, F1: true}
                _ = v15
            }
            _ = fi12
        }
    }
    v10 = (v10 | v10)
    switch -137 {
    case 322:
        v1 = (v1 / ((0) | 1))
    case 1:
        for fi16 := 0; fi16 < 9; fi16++ {
            switch fi16 {
            case -9223372036854775808:
                v17 := GenS2{F0: 1, F1: false}
                _ = v17
            case 127:
                v10 += v10
            }
            fmt.Println(v10, fi16)
            _ = fi16
        }
    default:
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS3{F0: 42, F1: (fi18 + (fi18 / ((127) | 1)))}
            for fi20 := 0; fi20 < 1; fi20++ {
                switch "hello" {
                case "foo bar":
                    switch (("\n\t" + "\n\t") + "\n\t") {
                    case "hello":
                        v1 = v1
                    }
                case "a":
                    v10 = v10
                default:
                    v3 = 18446744073709551615
                    var v21 uint64 = 1
                    _ = v21
                }
                v3 = (v3 + (v3 << ((v3) & 63)))
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    var v22 bool = (false || false)
    if ((-2.25 + 0.1) > 0.1) {
        for fi23 := 0; fi23 < 10; fi23++ {
            if v22 {
                v24 := GenS4{F0: 0, F1: 2147483648, F2: ("\n\t" + "hello")}
                if (v22 || v22) {
                    if (473 < 65535) {
                        var v25 uint16 = 660
                        _ = v25
                    }
                } else {
                    var v26 int = fi23
                    _ = v26
                }
                _ = v24
            }
            _ = fi23
        }
        v27 := GenS1{F0: -715, F1: 0}
        _ = v27
    } else {
        switch "hello" {
        case "foo bar":
            v28 := GenS5{F0: v3, F1: ""}
            _ = v28
        default:
            var v29 int8 = (v10 & (-(v10)))
            _ = v29
        }
        v30 := []int{808, int(v3), -429, 127}
        switch (("" + (("" + "") + "\n\t")) + "foo bar") {
        case "":
            if (!v22) {
                var v31 int = 4294967295
                _ = v31
            } else {
                var v32 float64 = float64(-2.25)
                _ = v32
            }
        case "\n\t":
            if (3.14159 > 0.0) {
                var v33 string = "foo bar"
                var v34 int = -1
                _ = v33
                _ = v34
            } else {
                var v35 float32 = ((1.5 - 0.1) / (0.1 - -2.25))
                v3 += (v3 / (((^(v3))) | 1))
                v35 = ((((v35 - v35) * (v35 / ((v35 * (((v35 + v35) + (float32((-2.25 - 3.14159)) * float32(1.5))) - v35)) * 0.1))) - v35) - 1e10)
                _ = v35
            }
        case "hello":
            if true {
                switch "\n\t" {
                case "hello":
                    for fi36 := 0; fi36 < 7; fi36++ {
                        fmt.Println(v10)
                        var v37 uint32 = 2
                        _ = fi36
                        _ = v37
                    }
                case "foo bar":
                    v1 = (-(v1))
                case "a":
                    var v38 uint32 = 2147483648
                    _ = v38
                default:
                    switch -678 {
                    case 127:
                        if v22 {
                            v39 := GenS6{F0: 547, F1: 1}
                            _ = v39
                        } else {
                            var v40 int64 = v1
                            for fi41 := 0; fi41 < 6; fi41++ {
                                if true {
                                    for fi42 := 0; fi42 < 6; fi42++ {
                                        v43 := GenS5{F0: v3, F1: (("hello" + "a") + "a")}
                                        _ = fi42
                                        _ = v43
                                    }
                                } else {
                                    fmt.Println(v1)
                                }
                                switch fi41 {
                                case -2147483648:
                                    v44 := GenS7{F0: 735, F1: 255, F2: 128}
                                    _ = v44
                                case 2147483647:
                                    var v45 int64 = (v1 | v40)
                                    _ = v45
                                default:
                                    var v46 uint32 = 1
                                    _ = v46
                                }
                                var v47 int = (fi41 - (fi41 ^ (fi41 | fi41)))
                                _ = fi41
                                _ = v47
                            }
                            _ = v40
                        }
                    case -9223372036854775808:
                        var v48 int8 = 1
                        _ = v48
                    }
                }
                v49 := GenS3{F0: 635, F1: 32768}
                _ = v49
            } else {
                var v50 int64 = (-(v1))
                fmt.Println(v3, v50)
                v51 := GenS7{F0: 127, F1: uint8(v3), F2: 255}
                _ = v50
                _ = v51
            }
        default:
            v52 := GenS8{F0: 1000}
            _ = v52
        }
        _ = v30
    }
    switch 978 {
    case -380:
        fmt.Println(v22)
    default:
        for fi53 := 0; fi53 < 1; fi53++ {
            _ = len(v2)
            v10 -= (v10 | v10)
            v54 := GenS9{F0: (0 == 2147483647)}
            _ = fi53
            _ = v54
        }
    }
    delete(v2, 42)
    var v55 uint64 = 630
    v55 = (v55 * 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v55
}
