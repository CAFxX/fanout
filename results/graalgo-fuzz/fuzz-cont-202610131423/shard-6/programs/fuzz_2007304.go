package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint64 }
func genFn0() uint32 {
    v5 := []int{322, 9223372036854775807, -2147483648, 9223372036854775807}
    if true {
        switch ("" + ("hello" + ("\n\t" + ("hello" + ("\n\t" + ("a" + "\n\t")))))) {
        case "hello":
            var v6 bool = (42 >= 2147483648)
            _ = v6
        case "":
            for fi7 := 0; fi7 < 11; fi7++ {
                v5 = append(v5, fi7)
                for fi8 := 0; fi8 < 6; fi8++ {
                    switch fi7 {
                    case 42:
                        switch "foo bar" {
                        case "a":
                            var v9 uint32 = 56
                            _ = v9
                        case "":
                            var v10 int16 = 765
                            _ = v10
                        }
                    default:
                        v11 := GenS1{F0: "hello", F1: 59, F2: 1000}
                        var v12 int16 = 624
                        _ = v11
                        _ = v12
                    }
                    var v13 float32 = 0.0
                    v14 := GenS1{F0: "a", F1: 0, F2: -32768}
                    _ = fi8
                    _ = v13
                    _ = v14
                }
                _ = fi7
            }
        case "a":
            v15 := GenS1{F0: ("foo bar" + ""), F1: 60, F2: 32767}
            _ = v15
        }
        var v16 uint = 287
        _ = v16
    }
    v17 := GenS1{F0: "hello", F1: 0, F2: 1000}
    var v18 uint16 = 255
    _ = v17
    _ = v18
    _ = v5
    return 0
}
func genFn1() uint {
    switch 128 {
    case 962:
        if ((genFn0() == 4294967295) || false) {
            v19 := GenS1{F0: "", F1: 1, F2: -707}
            var v20 string = "hello"
            v21 := GenS2{F0: 927, F1: "hello", F2: 888}
            _ = v19
            _ = v20
            _ = v21
        } else {
            if (42 > 2) {
                var v22 uint8 = 2
                v22 = (v22 % ((v22) | 1))
                v22 = 128
                _ = v22
            } else {
                if (-946 == 1000) {
                    var v23 int64 = -9223372036854775808
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v25 := GenS3{F0: 1, F1: 0}
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                } else {
                    var v26 bool = (627 > 2)
                    _ = v26
                }
                var v27 uint16 = 940
                v28 := GenS4{F0: true, F1: 1}
                _ = v27
                _ = v28
            }
        }
    }
    if false {
        switch "a" {
        case "hello":
            var v29 int = 127
            _ = v29
        case "a":
            v30 := []int{0, 0, 32767, 1}
            _ = v30
        default:
            var v31 int32 = 0
            _ = v31
        }
        var v32 float64 = -2.25
        v33 := []int{0}
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 19; fi34++ {
            if true {
                for fi35 := 0; fi35 < 5; fi35++ {
                    var v36 float64 = 0.1
                    if (false && false) {
                        v36 = v36
                        var v37 int32 = 783
                        _ = v37
                    }
                    v38 := []int{(fi34 & fi35)}
                    _ = fi35
                    _ = v36
                    _ = v38
                }
                v39 := GenS4{F0: (("foo bar" == "a") && false), F1: 0}
                _ = v39
            }
            var v40 string = (("\n\t" + "") + "a")
            _ = fi34
            _ = v40
        }
        var v41 int16 = 1000
        if (false || false) {
            for fi42 := 0; fi42 < 5; fi42++ {
                v41 += (-(v41))
                var v43 float64 = float64(0.0)
                var v44 float64 = float64(-2.25)
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        _ = v41
    }
    var v45 uint = 18446744073709551615
    _ = v45
    return v45
}

func main() {
    var v1 bool = (-428 <= 1)
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 uint64 = 645
    fmt.Println(v1)
    v2[-101] = 221
    for fi46 := 0; fi46 < 10; fi46++ {
        if v1 {
            v47 := []int{62}
            _ = v47
        }
        if v1 {
            switch (fi46 >> ((genFn1()) & 63)) {
            case 127:
                v48 := []int{42}
                _ = v48
            case 335:
                if v1 {
                    for fi49 := 0; fi49 < 7; fi49++ {
                        var v50 int = fi49
                        v51 := GenS3{F0: (v4 >> ((0) & 63)), F1: -569}
                        v1 = (65535 <= 1)
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    if v1 {
                        fmt.Println(v1)
                    }
                    v4 = 2
                } else {
                    var v52 int64 = -815
                    v53 := []int{(fi46 & 32767)}
                    var v54 uint16 = 1
                    _ = v52
                    _ = v53
                    _ = v54
                }
            case 42:
                var v55 uint32 = 4294967295
                _ = v55
            default:
                v56 := []int{(fi46 % ((660) | 1)), fi46, -224, fi46}
                _ = v56
            }
            v1 = v1
            v57 := GenS3{F0: v4, F1: 542}
            _ = v57
        } else {
            switch fi46 {
            case 169:
                v58 := GenS1{F0: "\n\t", F1: 2, F2: 1}
                _ = v58
            }
        }
        switch (fi46 + (fi46 & (-(fi46)))) {
        case 0:
            v4 += (v4 + 2)
        case -2147483648:
            fmt.Println(v4, v4)
        }
        _ = fi46
    }
    v59 := GenS2{F0: v4, F1: ("hello" + ("\n\t" + "hello")), F2: 614}
    switch ("\n\t" + "foo bar") {
    case "":
        v3[2147483647] = "a"
    case "\n\t":
        v60 := []int{-9223372036854775808}
        _ = v60
    case "a":
        delete(v3, 501)
    }
    for fi61 := 0; fi61 < 17; fi61++ {
        switch -175 {
        case 9223372036854775807:
            v1 = (!("hello" >= ("" + "\n\t")))
        case -729:
            for fi62 := 0; fi62 < 5; fi62++ {
                v4 += v4
                _ = fi62
            }
        case 0:
            var v63 uint32 = 42
            _ = v63
        }
        _ = fi61
    }
    v64 := GenS3{F0: 1, F1: 440}
    v4 = 1
    fmt.Println(v1, v1)
    v65 := GenS1{F0: "", F1: 138, F2: 42}
    v66 := GenS5{F0: 566, F1: 42}
    if ((231 < 42) || v1) {
        if v1 {
            var v67 int8 = 42
            var v68 int8 = (v67 >> ((42) & 63))
            for fi69 := 0; fi69 < 3; fi69++ {
                v68 = (v67 % (((v68 >> ((0) & 63))) | 1))
                _ = fi69
            }
            _ = v67
            _ = v68
        } else {
            v1 = (701 >= 42)
            var v70 int8 = 120
            _ = v70
        }
        v1 = v1
    } else {
        if (1 <= 891) {
            var v71 int64 = -1
            _ = v71
        } else {
            v3[2147483648] = ("\n\t" + "foo bar")
        }
    }
    switch "foo bar" {
    case "a":
        v72 := GenS5{F0: 2, F1: 32767}
        _ = v72
    case "\n\t":
        _ = len(v3)
    case "foo bar":
        fmt.Println(v4)
    }
    v73 := GenS6{F0: v4}
    v74 := GenS5{F0: 498, F1: -653}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v64
    _ = v65
    _ = v66
    _ = v73
    _ = v74
}
