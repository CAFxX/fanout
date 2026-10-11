package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 int64 }
func genFn0() int {
    if (false && (999 < 1)) {
        var v4 int16 = 32
        if false {
            fmt.Println(v4)
        }
        _ = v4
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int16 = -32768
        v7 := GenS1{F0: (fi5 + fi5)}
        if true {
            if (1 > 4294967295) {
                for fi8 := 0; fi8 < 12; fi8++ {
                    v6 -= v6
                    _ = fi8
                }
                v6 = (-(v6))
                fmt.Println(fi5, fi5)
            }
            fmt.Println(v6, fi5)
            fmt.Println(v6, v6)
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v9 int32 = -1
    switch "foo bar" {
    case "":
        if true {
            switch "foo bar" {
            case "":
                var v10 int = 188
                _ = v10
            }
        } else {
            var v11 int = -1
            if (18446744073709551615 <= 2) {
                v12 := GenS2{F0: 2, F1: 726, F2: -812}
                var v13 uint16 = 42
                _ = v12
                _ = v13
            }
            _ = v11
        }
    case "a":
        switch 32767 {
        case 462:
            fmt.Println(v9)
        case 408:
            v14 := GenS1{F0: 828}
            _ = v14
        }
    case "\n\t":
        fmt.Println(v9)
    }
    _ = v9
    return 0
}
func genFn1() int {
    switch ("a" + ("foo bar" + "foo bar")) {
    case "foo bar":
        for fi15 := 0; fi15 < 3; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 float32 = 0.0
                var v18 float32 = (v17 * v17)
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
    case "\n\t":
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 uint32 = 2
            v20 = v20
            _ = fi19
            _ = v20
        }
    case "":
        var v21 string = ("\n\t" + "hello")
        _ = v21
    }
    v22 := []int{9223372036854775807, -1, 32768}
    var v23 uint32 = 1
    fmt.Println(v23)
    _ = v22
    _ = v23
    return 128
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 int16 = -360
    v3 := make(map[int]bool)
    delete(v3, -9223372036854775808)
    var v24 int32 = 42
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 uint8 = 42
        var v27 uint32 = 2147483648
        _ = fi25
        _ = v26
        _ = v27
    }
    if (963 <= 2147483648) {
        _ = len(v1)
        v24 -= v24
        v3[644] = true
    } else {
        if (129 == 262) {
            switch 127 {
            case 32768:
                switch "hello" {
                case "\n\t":
                    v2 = (v2 / (((^(v2))) | 1))
                case "a":
                    var v28 int16 = 408
                    _ = v28
                case "":
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 uint16 = 0
                        delete(v3, fi29)
                        v2 = (^(v2))
                        _ = fi29
                        _ = v30
                    }
                }
            case 32767:
                v31 := v3[int(v24)]
                _ = v31
            default:
                v2 -= v2
            }
            switch 674 {
            case -633:
                v2 = (^(v2))
            default:
                v2 = -1
            }
        } else {
            v24 = -191
        }
    }
    v32 := GenS3{F0: true}
    var v33 uint16 = 2
    v34 := GenS4{F0: 738, F1: 666}
    v35 := GenS1{F0: 127}
    _ = len(v3)
    var v36 uint32 = 310
    switch 395 {
    case -763:
        delete(v3, 0)
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        for fi38 := 0; fi38 < 6; fi38++ {
            var v39 int32 = (v24 | v24)
            v24 = -2147483648
            if false {
                var v40 uint = 42
                _ = v40
            }
            _ = fi38
            _ = v39
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            v42 := GenS5{F0: 1, F1: 62}
            _ = fi41
            _ = v42
        }
        for fi43 := 0; fi43 < 12; fi43++ {
            switch "a" {
            case "a":
                v24 += v24
                var v44 int8 = 42
                _ = v44
            case "":
                fmt.Println(v24)
            case "foo bar":
                v1 = append(v1, 0)
            default:
                if ((-(v24)) == (v24 + v24)) {
                    v45 := GenS1{F0: fi43}
                    var v46 int8 = 0
                    var v47 int32 = v24
                    _ = v45
                    _ = v46
                    _ = v47
                }
            }
            var v48 string = ""
            v2 = v2
            _ = fi43
            _ = v48
        }
        _ = fi37
    }
    for fi49 := 0; fi49 < 3; fi49++ {
        for fi50 := 0; fi50 < 7; fi50++ {
            var v51 int64 = 137
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    if (((!false) || false) || (v24 == v24)) {
        v33 = uint16(v36)
        v52 := GenS5{F0: -143, F1: 127}
        _ = v52
    }
    for fi53 := 0; fi53 < 3; fi53++ {
        var v54 float64 = 0.0
        _ = fi53
        _ = v54
    }
    v55 := GenS6{F0: -9223372036854775808}
    var v56 float64 = (1.5 - 0.1)
    v2 -= -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v55
    _ = v56
}
