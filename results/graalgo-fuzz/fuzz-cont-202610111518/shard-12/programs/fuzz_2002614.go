package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint; F2 int8 }

func main() {
    var v1 int = 4294967295
    v2 := []uint32{}
    v3 := make(map[string]bool)
    var v4 float64 = float64(1.5)
    v5 := make(map[string]uint32)
    v4 = v4
    v4 -= 0.1
    for fi6 := 0; fi6 < 4; fi6++ {
        delete(v3, "\n\t")
        _ = fi6
    }
    delete(v5, "a")
    switch "\n\t" {
    case "hello":
        if (!false) {
            v7 := GenS1{F0: 18446744073709551615, F1: 647}
            v8 := GenS1{F0: 42, F1: 365}
            _ = v7
            _ = v8
        } else {
            v4 -= v4
        }
    }
    switch v1 {
    case 42:
        if (!false) {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    v4 = float64(3.14159)
                }
                var v9 float64 = (v4 * v4)
                _ = v9
            }
            switch "" {
            case "hello":
                var v10 int32 = -2147483648
                _ = v10
            }
        } else {
            v11 := GenS2{F0: 0}
            fmt.Println(v4, v4)
            _ = v11
        }
    case -1:
        if ((1 <= 42) || false) {
            delete(v3, "\n\t")
            var v12 int8 = 100
            _ = v12
        } else {
            fmt.Println(v1, v4)
            v4 = v4
            fmt.Println(v1, v1)
        }
    }
    switch "a" {
    case "\n\t":
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int64 = 1
            _ = fi13
            _ = v14
        }
    default:
        if (!(1 < 623)) {
            if false {
                var v15 int64 = -76
                _ = v15
            }
        } else {
            switch ("\n\t" + "hello") {
            case "":
                if true {
                    v1 = v1
                }
            case "a":
                v4 -= v4
            case "hello":
                var v16 float64 = ((v4 - v4) - 3.14159)
                _ = v16
            }
            v1 = (v1 / ((-326) | 1))
        }
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        if (!(0.1 > v4)) {
            v4 = v4
            if true {
                var v18 int = -572
                _ = v18
            }
        }
        _ = fi17
    }
    v19 := GenS3{F0: 4294967295, F1: 58}
    var v20 uint32 = 612
    for fi21 := 0; fi21 < 2; fi21++ {
        switch (fi21 % (((v1 >> ((uint(v20)) & 63))) | 1)) {
        case 273:
            for fi22 := 0; fi22 < 19; fi22++ {
                v4 -= (v4 - v4)
                switch ("" + "") {
                case "foo bar":
                    _ = len(v2)
                case "":
                    var v23 bool = true
                    _ = v23
                default:
                    if (!false) {
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    } else {
                        v20 = v20
                    }
                }
                _ = fi22
            }
        default:
            _ = len(v5)
        }
        var v24 int8 = -1
        _ = fi21
        _ = v24
    }
    var v25 int = v1
    switch "foo bar" {
    case "a":
        v2 = append(v2, (v20 << ((683) & 63)))
    }
    if (1.5 >= (0.1 / ((-2.25 / (-2.25 - 3.14159)) * 1e10))) {
        v26 := v5["\n\t"]
        switch "hello" {
        case "":
            v27 := GenS4{F0: 1, F1: 165}
            _ = v27
        case "hello":
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 string = "a"
                _ = fi28
                _ = v29
            }
        default:
            var v30 bool = (!false)
            _ = v30
        }
        _ = v26
    } else {
        v31 := GenS5{F0: -624}
        _ = v31
    }
    _ = len(v3)
    switch "" {
    case "\n\t":
        v4 = -2.25
    case "":
        fmt.Println(v4, v20)
    default:
        v20 -= v20
    }
    v32 := GenS6{F0: "a", F1: 347, F2: -128}
    if true {
        v4 += (0.1 + ((v4 * v4) / v4))
    } else {
        v33 := GenS5{F0: -9223372036854775808}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
