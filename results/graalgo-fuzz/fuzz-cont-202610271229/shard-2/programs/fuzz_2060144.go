package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0() uint16 {
    if false {
        v5 := []int{128, 0, -693}
        var v6 uint8 = 0
        v6 = (v6 + v6)
        _ = v5
        _ = v6
    } else {
        if true {
            var v7 uint64 = 1
            v7 += 1
            _ = v7
        } else {
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 int16 = 32767
                v10 := GenS1{F0: 42, F1: 42, F2: 100}
                fmt.Println(v9, v9)
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 uint64 = 42
            var v12 float32 = (float32((1.5 * (1e10 - 1e10))) * float32(-2.25))
            _ = v11
            _ = v12
        }
        v13 := []int{-424, 0, 32768}
        _ = v13
    }
    switch "foo bar" {
    case "":
        switch "hello" {
        case "foo bar":
            for fi14 := 0; fi14 < 14; fi14++ {
                switch "foo bar" {
                case "foo bar":
                    var v15 int16 = -1
                    _ = v15
                case "hello":
                    var v16 uint32 = 385
                    _ = v16
                case "\n\t":
                    if false {
                        var v17 int16 = 1000
                        v17 = v17
                        _ = v17
                    } else {
                        v18 := GenS1{F0: 2, F1: 0, F2: -27}
                        var v19 string = "hello"
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi14
            }
        case "a":
            v20 := GenS2{F0: 1, F1: ("\n\t" + "\n\t")}
            _ = v20
        default:
            if (65535 == 2147483648) {
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    v21 := []int{-776}
                    _ = v21
                case "hello":
                    var v22 float32 = ((3.14159 + (0.0 / 0.1)) - 3.14159)
                    _ = v22
                case "":
                    var v23 int16 = 42
                    _ = v23
                default:
                    v24 := []int{128}
                    _ = v24
                }
                v25 := GenS2{F0: 723, F1: ("a" + "\n\t")}
                var v26 string = "foo bar"
                _ = v25
                _ = v26
            }
        }
    }
    v27 := []int{9223372036854775807}
    var v28 uint16 = 65535
    _ = v27
    _ = v28
    return (v28 - (v28 - (v28 & (^(v28)))))
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 bool = ("a" == "foo bar")
    v4 := []int{}
    v29 := GenS3{F0: 0}
    switch "a" {
    case "\n\t":
        fmt.Println(v3)
    }
    var v30 uint16 = 65535
    if v3 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v30 -= 2
        v30 = 1
    }
    var v31 uint16 = (v30 * genFn0())
    for fi32 := 0; fi32 < 20; fi32++ {
        if (1 == v30) {
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 string = "a"
                _ = fi33
                _ = v34
            }
        }
        delete(v2, fi32)
        fmt.Println(v3, v30, v31)
        _ = fi32
    }
    var v35 string = ("" + "a")
    v35 = ("\n\t" + v35)
    var v36 uint64 = 18446744073709551615
    v30 += 32767
    fmt.Println(v3, v3, v36)
    var v37 int = -1
    for fi38 := 0; fi38 < 5; fi38++ {
        for fi39 := 0; fi39 < 10; fi39++ {
            if (-884 == -2147483648) {
                fmt.Println(v35)
            } else {
                var v40 int64 = -1
                v1[v35] = true
                var v41 int32 = 42
                _ = v40
                _ = v41
            }
            v1[(("\n\t" + v35) + (v35 + v35))] = (42 < -1)
            v42 := GenS4{F0: 1, F1: 42}
            _ = fi39
            _ = v42
        }
        var v43 uint8 = 1
        _ = fi38
        _ = v43
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        _ = len(v1)
        v35 = v35
        _ = fi44
    }
    if v3 {
        fmt.Println(v37, v31, v31)
    } else {
        v35 = (("\n\t" + v35) + "a")
    }
    var v45 bool = (v3 || (127 != 31))
    if (1000 <= -32768) {
        for fi46 := 0; fi46 < 11; fi46++ {
            if v45 {
                v31 += v31
                for fi47 := 0; fi47 < 10; fi47++ {
                    _ = len(v4)
                    fmt.Println(v35)
                    _ = fi47
                }
            } else {
                switch v37 {
                case 32767:
                    switch ((v35 + v35) + "\n\t") {
                    case "\n\t":
                        switch "" {
                        case "a":
                            v48 := GenS1{F0: v31, F1: 202, F2: 0}
                            _ = v48
                        }
                    default:
                        if false {
                            switch ((v35 + "hello") + ((v35 + v35) + v35)) {
                            case "\n\t":
                                v31 -= (v31 | (v30 ^ v30))
                            case "a":
                                v45 = (0 > 0)
                            default:
                                v49 := GenS4{F0: -1, F1: -128}
                                _ = v49
                            }
                        } else {
                            switch v37 {
                            case 0:
                                if (0 <= 634) {
                                    fmt.Println(v36, v35, fi46)
                                    fmt.Println(v3, v36)
                                    var v50 int64 = -9223372036854775808
                                    _ = v50
                                }
                            case 127:
                                var v51 uint8 = 1
                                _ = v51
                            case 32768:
                                for fi52 := 0; fi52 < 8; fi52++ {
                                    v35 = ((((v35 + v35) + "a") + "a") + (("\n\t" + v35) + (v35 + (v35 + "foo bar"))))
                                    v35 = v35
                                    _ = fi52
                                }
                            default:
                                v53 := GenS2{F0: (^(v36)), F1: v35}
                                _ = v53
                            }
                        }
                    }
                case -548:
                    var v54 float32 = -2.25
                    _ = v54
                case -336:
                    v30 = v30
                default:
                    fmt.Println(v35, fi46)
                }
                switch v35 {
                case "a":
                    var v55 int16 = 1000
                    _ = v55
                default:
                    v1[v35] = v3
                }
                for fi56 := 0; fi56 < 15; fi56++ {
                    if v3 {
                        fmt.Println(v31, fi46)
                        v37 += fi46
                    }
                    _ = fi56
                }
            }
            _ = fi46
        }
        var v57 uint32 = 1
        v31 -= v30
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v45
}
