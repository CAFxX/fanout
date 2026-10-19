package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 int; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 18446744073709551615
    var v3 uint32 = 42
    if (204 <= 0) {
        if (-2.25 == float64((float32(1e10) * (((1e10 / 0.1) * 3.14159) / float32(1e10))))) {
            v4 := []int{1, 127}
            _ = v4
        } else {
            v2 = (v2 % ((255) | 1))
            if false {
                if (0.1 <= 0.0) {
                    v2 = v2
                }
            }
            v5 := GenS1{F0: (1 <= 866)}
            _ = v5
        }
        fmt.Println(v2)
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        v2 += v2
        var v7 int32 = -2147483648
        switch (fi6 / ((107) | 1)) {
        case -2147483648:
            switch (-(fi6)) {
            case 42:
                var v8 float64 = (1e10 / (-2.25 - 0.0))
                _ = v8
            default:
                if ("" < "hello") {
                    v9 := GenS2{F0: 956, F1: 9223372036854775807, F2: (false || true)}
                    var v10 int = 9223372036854775807
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v7 = 42
                        v3 -= (^(v3))
                        _ = fi11
                    }
                    _ = v10
                    _ = v9
                }
            }
        case 2147483647:
            _ = len(v1)
        default:
            var v12 uint16 = 1
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    v13 := GenS3{F0: -544, F1: true, F2: 0}
    if (true || (false && false)) {
        v3 -= (^(v3))
        switch "foo bar" {
        case "foo bar":
            switch "" {
            case "":
                for fi14 := 0; fi14 < 7; fi14++ {
                    v2 += v2
                    v15 := []int{}
                    _ = fi14
                    _ = v15
                }
            case "hello":
                fmt.Println(v3)
            case "foo bar":
                switch "a" {
                case "foo bar":
                    v3 = v3
                case "a":
                    v16 := GenS3{F0: 9223372036854775807, F1: (!false), F2: 79}
                    _ = v16
                case "":
                    var v17 bool = true
                    _ = v17
                default:
                    v1[("a" + "foo bar")] = true
                }
            }
        case "hello":
            fmt.Println(v2, v2)
        case "\n\t":
            if (42 <= (v3 >> (((v2 / ((v2) | 1))) & 63))) {
                v3 -= 65535
                v18 := GenS2{F0: (v2 << ((4294967295) & 63)), F1: -569, F2: false}
                delete(v1, "a")
                _ = v18
            }
        default:
            var v19 uint16 = 42
            _ = v19
        }
        v3 = uint32(v2)
    }
    v2 -= (v2 + 18446744073709551615)
    v2 = 97
    switch "\n\t" {
    case "hello":
        v3 = v3
    default:
        if (!true) {
            fmt.Println(v3)
        }
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        switch fi20 {
        case -9223372036854775808:
            for fi21 := 0; fi21 < 17; fi21++ {
                var v22 string = ("foo bar" + "\n\t")
                var v23 bool = (-2147483648 < -1)
                _ = fi21
                _ = v22
                _ = v23
            }
        case 1:
            switch "hello" {
            case "hello":
                v3 = 0
            case "":
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS2{F0: (v2 & (^(v2))), F1: 42, F2: false}
                    var v26 int16 = 42
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            default:
                if (100 >= -1) {
                    v27 := []int{(fi20 / ((1) | 1)), (-(fi20))}
                    var v28 int16 = 0
                    _ = v27
                    _ = v28
                }
            }
        case 93:
            fmt.Println(v2, fi20)
        default:
            for fi29 := 0; fi29 < 11; fi29++ {
                fmt.Println(v2)
                _ = fi29
            }
        }
        if false {
            var v30 float64 = 3.14159
            v30 -= 0.1
            _ = v30
        } else {
            for fi31 := 0; fi31 < 6; fi31++ {
                v32 := []int{}
                fmt.Println(fi31, v3, fi20)
                _ = fi31
                _ = v32
            }
            if true {
                var v33 uint64 = 2
                _ = v33
            }
            v3 = (v3 ^ (v3 * v3))
        }
        for fi34 := 0; fi34 < 11; fi34++ {
            if false {
                v1["hello"] = true
                v2 = v2
                var v35 int64 = 577
                _ = v35
            }
            if true {
                v36 := GenS4{F0: 32767}
                var v37 uint64 = 233
                _ = v36
                _ = v37
            } else {
                v38 := GenS5{F0: 1, F1: 0, F2: 0}
                var v39 uint32 = (v3 + 1)
                switch "a" {
                case "":
                    var v40 uint8 = 15
                    _ = v40
                case "a":
                    delete(v1, "foo bar")
                    if true {
                        delete(v1, ("hello" + "\n\t"))
                        v3 = v39
                    }
                    var v41 int16 = 1000
                    _ = v41
                case "foo bar":
                    v42 := []int{}
                    _ = v42
                }
                fmt.Println(v39, v39)
                _ = v38
                _ = v39
            }
            _ = fi34
        }
        _ = fi20
    }
    var v43 float32 = -2.25
    v3 = 2
    if false {
        v3 += v3
        v43 -= float32(0.0)
        for fi44 := 0; fi44 < 4; fi44++ {
            fmt.Println(fi44, v2)
            _ = fi44
        }
    }
    var v45 bool = false
    var v46 uint8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v43
    _ = v45
    _ = v46
}
