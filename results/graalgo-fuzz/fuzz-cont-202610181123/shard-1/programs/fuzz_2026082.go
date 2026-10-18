package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int8; F1 int8; F2 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 uint32; F2 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 int64; F1 uint32 }

func main() {
    v1 := []int{128, 42, 2147483647, -2147483648}
    v2 := make(map[string]string)
    var v3 float32 = ((3.14159 - 1e10) - (1.5 + (((((-2.25 + 1e10) + 1.5) + 0.0) / 1e10) * 3.14159)))
    var v4 string = ""
    var v5 bool = true
    var v6 uint = 42
    v7 := GenS1{F0: 0}
    fmt.Println(v3)
    if false {
        v4 = v4
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint64 = 2
            v6 = 2
            _ = fi8
            _ = v9
        }
        switch ("\n\t" + v4) {
        case "\n\t":
            for fi10 := 0; fi10 < 15; fi10++ {
                switch fi10 {
                case -332:
                    v11 := GenS2{F0: -609, F1: 153}
                    _ = v11
                case 1:
                    var v12 bool = false
                    _ = v12
                case -2147483648:
                    var v13 uint16 = 32767
                    _ = v13
                }
                _ = fi10
            }
        case "a":
            v6 -= 767
        case "hello":
            v14 := GenS3{F0: 254, F1: (v6 >> (((v6 >> ((v6) & 63))) & 63))}
            _ = v14
        }
    }
    v15 := GenS4{F0: 127, F1: -118, F2: 42}
    switch 0 {
    case 9223372036854775807:
        v16 := GenS5{F0: 1}
        _ = v16
    }
    var v17 uint8 = 255
    v1 = append(v1, 2147483647)
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint64 = 0
        v17 = (v17 >> (((v6 * (v6 | (v6 >> ((v6) & 63))))) & 63))
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v5, v5)
            v3 -= float32(0.0)
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    v4 = (v4 + ("foo bar" + v4))
    var v21 bool = (-1 > 604)
    v5 = (1 != 2)
    fmt.Println(v17, v17)
    v22 := v2["a"]
    v23 := GenS6{F0: 100, F1: 4294967295, F2: v6}
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := GenS7{F0: (fi24 * fi24)}
        _ = fi24
        _ = v25
    }
    if (-2147483648 != 32768) {
        v5 = (32767 <= 65535)
        var v26 uint = v6
        fmt.Println(v6)
        _ = v26
    } else {
        fmt.Println(v17)
        switch "a" {
        case "foo bar":
            fmt.Println(v21, v17)
        case "\n\t":
            v4 = (((v4 + "foo bar") + "") + v22)
        case "a":
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 bool = (65535 <= 585)
                _ = fi27
                _ = v28
            }
        default:
            var v29 string = "a"
            _ = v29
        }
        v5 = v5
    }
    v21 = (v5 && true)
    if (v4 >= "\n\t") {
        for fi30 := 0; fi30 < 18; fi30++ {
            switch (fi30 % (((^(fi30))) | 1)) {
            case 698:
                _ = len(v2)
            case 9223372036854775807:
                v31 := v2[("hello" + (v22 + v22))]
                _ = v31
            case 127:
                for fi32 := 0; fi32 < 5; fi32++ {
                    v3 = float32((1.5 * 0.0))
                    switch (fi30 ^ (fi30 << ((v6) & 63))) {
                    case 9223372036854775807:
                        if (v21 || v5) {
                            v33 := GenS1{F0: -403}
                            if ((1 == 42) && (v5 && v5)) {
                                _ = len(v1)
                                if (v3 == v3) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v6 -= v6
                                    if true {
                                        v34 := GenS5{F0: (v6 - v6)}
                                        if (((fi32 % (((fi32 ^ (fi30 & (fi30 * fi32)))) | 1)) < (fi32 << ((v6) & 63))) && (!v5)) {
                                            v21 = (fi32 >= fi30)
                                            var v35 uint8 = v17
                                            if (513 != 773) {
                                                v36 := GenS8{F0: 42, F1: 0}
                                                v37 := GenS8{F0: 42, F1: 0}
                                                v6 = (v6 * v6)
                                                v38 := GenS9{F0: 32768}
                                                switch (fi32 % (((fi30 | (-(fi30)))) | 1)) {
                                                case 0:
                                                    v6 = v6
                                                case 128:
                                                    fmt.Println(fi30, fi32, v3)
                                                case -264:
                                                    v6 = (v6 >> (((^(v6))) & 63))
                                                }
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            }
                                            _ = v35
                                        } else {
                                            var v39 uint8 = v17
                                            v40 := GenS10{F0: 42}
                                            if (false || (!false)) {
                                                fmt.Println(v22, v4, v17)
                                            } else {
                                                v41 := GenS11{F0: v22, F1: (((v4 + (("a" + (v22 + "foo bar")) + v4)) + v4) + "\n\t")}
                                                _ = v41
                                            }
                                            _ = v39
                                            _ = v40
                                        }
                                        _ = v34
                                    }
                                } else {
                                    v1 = append(v1, fi30)
                                    if (2147483647 != fi30) {
                                        var v42 uint = (v6 << ((v6) & 63))
                                        var v43 uint16 = 255
                                        _ = v42
                                        _ = v43
                                    } else {
                                        switch (fi32 * fi30) {
                                        case -437:
                                            fmt.Println(v3, v3, v3)
                                        }
                                        switch fi30 {
                                        case 4294967295:
                                            v3 = 1.5
                                        case 1:
                                            var v44 uint8 = 128
                                            _ = v44
                                        }
                                    }
                                }
                            }
                            _ = v33
                        }
                    }
                    if (807 <= 2147483648) {
                        var v45 string = v22
                        _ = v45
                    } else {
                        v22 = ((v4 + v4) + ("" + v22))
                        if true {
                            var v46 float64 = 1.5
                            _ = v46
                        } else {
                            switch fi30 {
                            case 42:
                                v3 = v3
                            case 32767:
                                v5 = v21
                            default:
                                switch fi30 {
                                case 0:
                                    v17 = (^(v17))
                                case 32768:
                                    var v47 uint32 = 367
                                    _ = v47
                                case 670:
                                    var v48 int32 = -836
                                    if v21 {
                                        switch fi32 {
                                        case 32768:
                                            v49 := GenS12{F0: 42, F1: 2147483648}
                                            _ = v49
                                        case 20:
                                            var v50 uint8 = v17
                                            var v51 int32 = (-(v48))
                                            _ = v50
                                            _ = v51
                                        }
                                        var v52 uint8 = (v17 | (v17 ^ (v17 & v17)))
                                        var v53 uint = (v6 % ((18446744073709551615) | 1))
                                        _ = v52
                                        _ = v53
                                    } else {
                                        var v54 float32 = 0.1
                                        _ = v54
                                    }
                                    var v55 uint8 = 1
                                    _ = v48
                                    _ = v55
                                }
                            }
                        }
                    }
                    var v56 float32 = 1e10
                    v6 -= 42
                    _ = fi32
                    _ = v56
                }
            default:
                if v5 {
                    fmt.Println(v3, v21, v6)
                    switch 32768 {
                    case -2147483648:
                        switch (v22 + v4) {
                        case "":
                            v57 := GenS9{F0: 32767}
                            _ = v57
                        case "\n\t":
                            v5 = (2147483647 > -50)
                        }
                    case -9223372036854775808:
                        v1 = append(v1, fi30)
                    case 32767:
                        var v58 uint16 = 799
                        _ = v58
                    default:
                        for fi59 := 0; fi59 < 8; fi59++ {
                            _ = len(v1)
                            fmt.Println(v17, v3)
                            _ = fi59
                        }
                    }
                }
            }
            v17 = v17
            var v60 float64 = 0.0
            _ = fi30
            _ = v60
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
