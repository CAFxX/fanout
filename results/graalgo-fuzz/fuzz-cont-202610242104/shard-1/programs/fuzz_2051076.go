package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint32; F2 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 uint64; F2 string }
type GenS10 struct { F0 uint32; F1 int }
type GenS11 struct { F0 uint16; F1 uint64 }
func genFn0() uint32 {
    var v4 int64 = 140
    fmt.Println(v4)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int16 = 42
        if true {
            switch (fi5 & fi5) {
            case 32767:
                var v7 float32 = 0.1
                _ = v7
            case 125:
                switch fi5 {
                case 9223372036854775807:
                    fmt.Println(fi5, v4)
                case 42:
                    var v8 float32 = (3.14159 * float32((float64(1.5) - 0.1)))
                    _ = v8
                default:
                    v4 += v4
                }
            case 9223372036854775807:
                var v9 int32 = 42
                _ = v9
            default:
                fmt.Println(v6)
            }
        } else {
            fmt.Println(v6, fi5)
        }
        _ = fi5
        _ = v6
    }
    v10 := GenS1{F0: ""}
    _ = v10
    _ = v4
    return 4294967295
}
func genFn1() string {
    switch 4294967295 {
    case 32768:
        if false {
            var v11 int32 = -1
            if (!false) {
                var v12 int64 = 0
                switch "a" {
                case "":
                    v11 += (v11 << ((18446744073709551615) & 63))
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        v13 := []int{-2147483648, 42}
                        _ = v13
                    default:
                        switch "hello" {
                        case "hello":
                            var v14 int32 = v11
                            _ = v14
                        case "foo bar":
                            var v15 uint64 = 1
                            _ = v15
                        default:
                            var v16 int16 = -32768
                            _ = v16
                        }
                    }
                case "\n\t":
                    fmt.Println(v12)
                default:
                    if true {
                        v17 := GenS2{F0: 128}
                        v18 := []int{}
                        _ = v17
                        _ = v18
                    } else {
                        var v19 bool = false
                        var v20 float32 = ((0.1 + 0.0) - (float32((0.1 - 1e10)) * 1.5))
                        var v21 uint16 = 78
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
                var v22 bool = false
                _ = v12
                _ = v22
            }
            _ = v11
        }
    case 480:
        v23 := GenS3{F0: 866, F1: 0, F2: 32767}
        _ = v23
    default:
        v24 := []int{}
        _ = v24
    }
    v25 := []int{2147483647, 42, 2147483647}
    var v26 int64 = 922
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS4{F0: 1}
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    _ = v25
    _ = v26
    return "hello"
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 3.14159
    var v3 string = "a"
    v30 := GenS5{F0: 1}
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 int64 = -1
        var v33 float64 = 3.14159
        delete(v1, v3)
        _ = fi31
        _ = v32
        _ = v33
    }
    if true {
        var v34 uint32 = 0
        fmt.Println(v3, v3)
        _ = v34
    } else {
        fmt.Println(v3)
    }
    fmt.Println(v3, v3)
    if (-787 > -23) {
        if true {
            for fi35 := 0; fi35 < 17; fi35++ {
                v3 = "hello"
                _ = fi35
            }
            for fi36 := 0; fi36 < 13; fi36++ {
                fmt.Println(v3)
                if true {
                    for fi37 := 0; fi37 < 9; fi37++ {
                        var v38 uint64 = 2
                        var v39 float64 = 1.5
                        switch (v3 + v3) {
                        case "":
                            v40 := GenS5{F0: 810}
                            _ = v40
                        case "\n\t":
                            _ = len(v1)
                        }
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                } else {
                    var v41 uint = 282
                    switch ("\n\t" + (((v3 + "a") + v3) + v3)) {
                    case "":
                        if false {
                            fmt.Println(v3, v2, v41)
                            v41 = 0
                            if (115 > 753) {
                                switch (-(fi36)) {
                                case -9223372036854775808:
                                    v3 = v3
                                case 128:
                                    var v42 int32 = -2147483648
                                    _ = v42
                                }
                                v41 -= (v41 % (((v41 >> (((^(v41))) & 63))) | 1))
                                fmt.Println(fi36, v41, fi36)
                            }
                        }
                    case "\n\t":
                        v43 := []int{(fi36 - (fi36 % ((fi36) | 1))), -369, fi36}
                        _ = v43
                    default:
                        _ = len(v1)
                    }
                    _ = v41
                }
                _ = fi36
            }
            v2 = (float64(float32(3.14159)) / v2)
        }
        fmt.Println(v3)
        var v44 int = 4294967295
        _ = v44
    }
    switch 127 {
    case -2147483648:
        var v45 bool = true
        _ = v45
    default:
        for fi46 := 0; fi46 < 10; fi46++ {
            var v47 float32 = (((0.0 / 1e10) - -2.25) * -2.25)
            v48 := GenS6{F0: true, F1: 2, F2: v3}
            for fi49 := 0; fi49 < 4; fi49++ {
                switch v3 {
                case "\n\t":
                    v50 := GenS4{F0: 42}
                    _ = v50
                case "foo bar":
                    fmt.Println(v47, v47, fi46)
                default:
                    v3 = v3
                }
                switch 9223372036854775807 {
                case 4294967295:
                    var v51 uint16 = 42
                    _ = v51
                case -643:
                    var v52 int64 = 0
                    var v53 int8 = 83
                    _ = v52
                    _ = v53
                default:
                    if true {
                        v3 = "a"
                        v3 = v3
                        if true {
                            v47 = (v47 + v47)
                        } else {
                            v54 := []int{-569}
                            v2 = v2
                            fmt.Println(v2)
                            v2 = float64(1e10)
                            v55 := GenS7{F0: "\n\t"}
                            if (0 > 990) {
                                var v56 int64 = -623
                                _ = v56
                            }
                            _ = v54
                            _ = v55
                        }
                        if true {
                            if false {
                                v57 := GenS1{F0: "a"}
                                _ = v57
                            }
                        }
                        if true {
                            if false {
                                var v58 int = (fi46 ^ (fi46 << ((65535) & 63)))
                                v3 = (("foo bar" + v3) + v3)
                                v59 := GenS6{F0: false, F1: 4294967295, F2: (v3 + "\n\t")}
                                v47 = ((v47 + ((v47 + v47) + 0.0)) - float32(v2))
                                if (false || false) {
                                    fmt.Println(v58, v58)
                                    if (v3 >= (v3 + v3)) {
                                        v60 := GenS8{F0: 1000}
                                        if true {
                                            var v61 uint64 = 4294967295
                                            fmt.Println(fi49)
                                            _ = v61
                                        } else {
                                            var v62 int64 = -259
                                            if false {
                                                v58 = (fi49 << ((291) & 63))
                                            }
                                            v63 := GenS6{F0: (v2 <= (float64(v47) * 3.14159)), F1: 42, F2: v3}
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = v60
                                    }
                                    switch (v3 + v3) {
                                    case "a":
                                        switch fi49 {
                                        case 42:
                                            if ((true && false) && true) {
                                                v3 = "\n\t"
                                                v64 := GenS9{F0: 42, F1: 737, F2: v3}
                                                v65 := []int{}
                                                _ = v64
                                                _ = v65
                                            } else {
                                                v66 := GenS10{F0: 42, F1: (fi46 + (v58 - (v58 * v58)))}
                                                _ = v66
                                            }
                                        case -9223372036854775808:
                                            var v67 int = 1
                                            _ = v67
                                        case 1:
                                            switch (fi46 >> ((1) & 63)) {
                                            case 42:
                                                var v68 uint8 = 0
                                                _ = v68
                                            case 37:
                                                v58 -= (v58 - fi46)
                                            }
                                        }
                                    case "foo bar":
                                        var v69 int = (fi49 ^ (v58 >> ((65535) & 63)))
                                        _ = v69
                                    }
                                    var v70 int8 = 60
                                    _ = v70
                                }
                                if true {
                                    switch (fi46 / (((v58 >> ((2) & 63))) | 1)) {
                                    case -9223372036854775808:
                                        v71 := GenS3{F0: fi46, F1: 1, F2: 32767}
                                        _ = v71
                                    case 32768:
                                        var v72 string = v3
                                        _ = v72
                                    }
                                } else {
                                    v73 := GenS11{F0: 42, F1: 0}
                                    var v74 uint = 445
                                    if (!true) {
                                        var v75 int = fi49
                                        if (v2 <= v2) {
                                            var v76 uint8 = 128
                                            v77 := []int{}
                                            var v78 bool = ((((580 > (fi49 << (((v74 + v74)) & 63))) && (false || (true || true))) && (171 < 611)) || (v3 > v3))
                                            _ = v76
                                            _ = v77
                                            _ = v78
                                        } else {
                                            var v79 float64 = (-2.25 * float64(v47))
                                            _ = v79
                                        }
                                        _ = v75
                                    }
                                    _ = v73
                                    _ = v74
                                }
                                var v80 uint64 = 2
                                _ = v58
                                _ = v59
                                _ = v80
                            }
                        }
                    } else {
                        v81 := GenS10{F0: 780, F1: 1}
                        _ = v81
                    }
                }
                _ = fi49
            }
            _ = fi46
            _ = v47
            _ = v48
        }
    }
    fmt.Println(v3)
    var v82 string = v3
    v83 := GenS1{F0: v3}
    v82 = v3
    v82 = ""
    var v84 float64 = v2
    v3 = "foo bar"
    fmt.Println(v82)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v82:", v82)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v82
    _ = v83
    _ = v84
}
