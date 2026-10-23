package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 int }
type GenS11 struct { F0 uint }

func main() {
    var v1 uint64 = 42
    var v2 float64 = 1e10
    var v3 uint16 = 32767
    var v4 int = 127
    var v5 uint64 = v1
    var v6 bool = ("hello" <= "a")
    v5 = (v1 ^ v5)
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: v5, F1: 2147483647}
        for fi9 := 0; fi9 < 1; fi9++ {
            if false {
                var v10 uint16 = v3
                v11 := GenS2{F0: "foo bar", F1: 42}
                _ = v10
                _ = v11
            } else {
                var v12 float64 = v2
                _ = v12
            }
            fmt.Println(v4, v2, v5)
            fmt.Println(v3, fi7, v1)
            switch ("" + "hello") {
            case "hello":
                var v13 int8 = 48
                _ = v13
            case "":
                if v6 {
                    v4 = (v4 / (((fi7 * v4)) | 1))
                    var v14 int8 = -68
                    var v15 uint = 65535
                    v16 := []int{}
                    v6 = (true && v6)
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    v17 := GenS2{F0: "", F1: 128}
                    var v18 int16 = 1000
                    fmt.Println(v5)
                    _ = v17
                    _ = v18
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v5 = (v5 - v5)
    v19 := GenS2{F0: ((((("hello" + ("" + (("hello" + ("hello" + "a")) + "a"))) + "foo bar") + "hello") + "a") + "hello"), F1: 2}
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS3{F0: 100, F1: (^(v5)), F2: 2}
        for fi22 := 0; fi22 < 7; fi22++ {
            switch "foo bar" {
            case "hello":
                v23 := GenS3{F0: 30, F1: v5, F2: 0}
                _ = v23
            case "":
                if (!(0 < 255)) {
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        var v24 uint32 = 2147483647
                        _ = v24
                    case "a":
                        if true {
                            v25 := []int{v4, (v4 ^ fi22), fi22}
                            _ = v25
                        }
                    case "":
                        var v26 int16 = 0
                        switch 1 {
                        case -1:
                            fmt.Println(v3, v5)
                        case 9223372036854775807:
                            v6 = (1.5 <= ((1e10 * float32(3.14159)) * -2.25))
                        case 32767:
                            v27 := GenS3{F0: -108, F1: (v1 >> ((429) & 63)), F2: 2}
                            _ = v27
                        }
                        v28 := GenS4{F0: (^(v5)), F1: 32767}
                        if ((!v6) && (-208 != -1)) {
                            var v29 string = "hello"
                            _ = v29
                        }
                        _ = v26
                        _ = v28
                    }
                } else {
                    switch fi20 {
                    case 310:
                        v6 = v6
                        v30 := []int{(fi20 - (fi22 - v4))}
                        v31 := GenS5{F0: (v1 << ((4294967295) & 63)), F1: 1}
                        var v32 uint16 = 2
                        _ = v30
                        _ = v31
                        _ = v32
                    case 669:
                        v6 = v6
                    case 506:
                        fmt.Println(v5, v4)
                    default:
                        var v33 int = 9223372036854775807
                        _ = v33
                    }
                }
            default:
                v34 := GenS6{F0: 0, F1: (^(v1))}
                _ = v34
            }
            v35 := []int{4294967295}
            _ = fi22
            _ = v35
        }
        fmt.Println(fi20, v4, v2)
        _ = fi20
        _ = v21
    }
    if (v6 && false) {
        var v36 int8 = 127
        for fi37 := 0; fi37 < 19; fi37++ {
            if ((0 < -9223372036854775808) || true) {
                v38 := GenS6{F0: uint32(v5), F1: v5}
                _ = v38
            } else {
                var v39 uint8 = 29
                for fi40 := 0; fi40 < 9; fi40++ {
                    v41 := GenS7{F0: 663, F1: 983}
                    _ = fi40
                    _ = v41
                }
                _ = v39
            }
            var v42 int32 = -318
            _ = fi37
            _ = v42
        }
        _ = v36
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        var v44 uint8 = 1
        for fi45 := 0; fi45 < 9; fi45++ {
            var v46 int32 = 1
            v47 := []int{v4, v4}
            if v6 {
                v48 := GenS8{F0: 32767}
                v3 = (v3 | (^(v3)))
                _ = v48
            }
            _ = len(v47)
            switch "a" {
            case "hello":
                v47 = append(v47, (fi43 % (((v4 | 32768)) | 1)))
                var v49 uint = 255
                _ = v49
            }
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi43
        _ = v44
    }
    v3 += (v3 & v3)
    v2 += v2
    if true {
        var v50 uint32 = 2147483647
        var v51 uint8 = 128
        _ = v50
        _ = v51
    }
    v52 := []int{v4, v4, (v4 - v4), (v4 / ((v4) | 1))}
    v4 = v4
    fmt.Println(v4)
    v1 -= v1
    fmt.Println(v2, v3)
    if true {
        fmt.Println(v1, v4, v2)
        var v53 int16 = 1000
        _ = v53
    } else {
        for fi54 := 0; fi54 < 20; fi54++ {
            if (32767 >= 541) {
                v1 = 4294967295
                for fi55 := 0; fi55 < 10; fi55++ {
                    if v6 {
                        switch "\n\t" {
                        case "a":
                            var v56 uint32 = 0
                            _ = v56
                        case "hello":
                            v1 -= v5
                        case "foo bar":
                            v3 = v3
                        }
                    }
                    v57 := GenS6{F0: 1, F1: 18446744073709551615}
                    switch (fi55 + 60) {
                    case -831:
                        var v58 uint8 = 1
                        _ = v58
                    case 1:
                        v59 := GenS9{F0: 128}
                        _ = v59
                    case 128:
                        v3 = (v3 | v3)
                    }
                    switch 32767 {
                    case 9223372036854775807:
                        v60 := GenS10{F0: (v4 * v4), F1: (v4 | 9223372036854775807)}
                        v61 := GenS11{F0: 4294967295}
                        _ = v60
                        _ = v61
                    case 127:
                        v4 += 1
                    case 238:
                        _ = len(v52)
                    }
                    switch fi55 {
                    case 127:
                        var v62 bool = (!v6)
                        _ = v62
                    }
                    _ = fi55
                    _ = v57
                }
                var v63 uint = uint(v5)
                _ = v63
            }
            var v64 int32 = -326
            var v65 uint8 = 181
            _ = fi54
            _ = v64
            _ = v65
        }
    }
    var v66 int32 = 0
    v5 += (v1 * 185)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v66
}
