package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{296}
    var v3 int32 = -1
    var v4 uint8 = 2
    if (!((true || true) && ("\n\t" <= (("hello" + "a") + "foo bar")))) {
        v5 := GenS1{F0: -36, F1: 924, F2: 2}
        switch 1 {
        case 127:
            var v6 int16 = -600
            _ = v6
        case 268:
            var v7 bool = (-980 == 467)
            _ = v7
        }
        _ = v5
    } else {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            for fi8 := 0; fi8 < 9; fi8++ {
                switch (^(fi8)) {
                case -181:
                    v9 := GenS2{F0: 9223372036854775807, F1: 42}
                    _ = v9
                case -1:
                    if (false || (!true)) {
                        v3 += v3
                    }
                default:
                    if (-2.25 >= 1.5) {
                        if false {
                            var v10 uint8 = v4
                            _ = v10
                        }
                    }
                }
                var v11 int8 = -1
                if len(v2) > 0 {
                    v2[0] = fi8
                }
                _ = fi8
                _ = v11
            }
        case "hello":
            switch (("a" + (("hello" + "foo bar") + ("hello" + "hello"))) + "foo bar") {
            case "":
                switch ("a" + "a") {
                case "\n\t":
                    if true {
                        var v12 uint32 = 974
                        var v13 uint64 = 4294967295
                        _ = v12
                        _ = v13
                    } else {
                        v4 = (v4 >> ((65535) & 63))
                    }
                case "a":
                    v3 -= -260
                }
            }
        }
    }
    switch "foo bar" {
    case "":
        if false {
            for fi14 := 0; fi14 < 6; fi14++ {
                if (9223372036854775807 == 0) {
                    v15 := GenS3{F0: (!((fi14 << ((385) & 63)) <= fi14)), F1: -128}
                    _ = v15
                } else {
                    v4 = v4
                }
                _ = fi14
            }
            var v16 float64 = (1e10 + 0.0)
            _ = v16
        } else {
            v3 = v3
        }
    case "a":
        for fi17 := 0; fi17 < 14; fi17++ {
            if len(v2) > 0 {
                v2[0] = (fi17 >> ((18446744073709551615) & 63))
            }
            _ = fi17
        }
    default:
        v4 = (v4 / ((v4) | 1))
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = (v3 << ((4294967295) & 63))
    if (18446744073709551615 == 4294967295) {
        v18 := GenS2{F0: 579, F1: 42}
        var v19 bool = (-524 != 128)
        switch 122 {
        case -105:
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 float32 = (0.0 / (1.5 / -2.25))
                if (255 >= 2) {
                    var v22 int32 = (v3 + (v3 << ((42) & 63)))
                    _ = v22
                }
                if (3.14159 != (float64(float32(0.0)) * (0.0 * (-2.25 * float64(v21))))) {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        var v24 string = ("a" + "foo bar")
                        v25 := v1["foo bar"]
                        v4 += 128
                        switch (("\n\t" + v24) + v24) {
                        case "hello":
                            fmt.Println(v4)
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    var v26 uint = uint(v4)
                    _ = v26
                } else {
                    var v27 uint64 = 2
                    switch "hello" {
                    case "a":
                        fmt.Println(v4)
                    case "":
                        fmt.Println(fi20)
                    case "\n\t":
                        if (-1 > -620) {
                            var v28 int64 = -1
                            var v29 float32 = v21
                            if (v28 == (v28 % (((^(v28))) | 1))) {
                                if v19 {
                                    v1[(("\n\t" + ("" + "hello")) + "foo bar")] = 1
                                }
                                fmt.Println(v27, fi20)
                            }
                            _ = v28
                            _ = v29
                        }
                    default:
                        var v30 int16 = 1
                        _ = v30
                    }
                    _ = v27
                }
                _ = fi20
                _ = v21
            }
        case 306:
            _ = len(v1)
        case 2147483647:
            var v31 uint32 = 650
            _ = v31
        }
        _ = v18
        _ = v19
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v32 string = "a"
    fmt.Println(v3, v4, v3)
    v3 = (^(v3))
    for fi33 := 0; fi33 < 13; fi33++ {
        v4 += (v4 + v4)
        for fi34 := 0; fi34 < 6; fi34++ {
            delete(v1, ((("hello" + ("foo bar" + v32)) + "") + (v32 + v32)))
            var v35 int = -2147483648
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    fmt.Println(v3, v3)
    var v36 uint = 42
    if (!true) {
        switch ("a" + ((v32 + v32) + "\n\t")) {
        case "a":
            v2 = append(v2, 164)
        default:
            for fi37 := 0; fi37 < 18; fi37++ {
                v38 := GenS1{F0: 100, F1: v36, F2: 4294967295}
                v1["\n\t"] = 4294967295
                var v39 uint = v36
                _ = fi37
                _ = v38
                _ = v39
            }
        }
        v3 -= v3
        if false {
            switch ((v32 + "foo bar") + "\n\t") {
            case "foo bar":
                var v40 uint64 = 42
                _ = v40
            case "hello":
                var v41 string = (v32 + v32)
                _ = v41
            case "\n\t":
                _ = len(v1)
            default:
                var v42 int = 66
                _ = v42
            }
        }
    } else {
        _ = len(v2)
        var v43 int32 = v3
        _ = v43
    }
    fmt.Println(v3, v36, v32)
    var v44 int64 = -1
    if (2 != 42) {
        var v45 int16 = -1
        _ = v45
    } else {
        var v46 int = 128
        _ = v46
    }
    switch 128 {
    case -2147483648:
        for fi47 := 0; fi47 < 16; fi47++ {
            v48 := GenS3{F0: false, F1: -128}
            fmt.Println(v44, fi47)
            fmt.Println(v32)
            _ = fi47
            _ = v48
        }
    case -283:
        v1[v32] = 1
    default:
        switch (((v32 + "a") + "foo bar") + "hello") {
        case "\n\t":
            v49 := GenS4{F0: -32768, F1: 18446744073709551615}
            _ = v49
        case "hello":
            v50 := v1[(v32 + ((v32 + v32) + ""))]
            _ = v50
        case "":
            v32 = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v44
}
