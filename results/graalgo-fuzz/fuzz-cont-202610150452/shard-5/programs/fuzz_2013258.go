package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int8; F2 string }
type GenS5 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := []string{"a", "\n\t", (("foo bar" + "") + (("" + "a") + "")), "hello"}
    var v2 uint64 = 565
    v3 := make(map[string]uint32)
    var v4 int = 32768
    v5 := GenS1{F0: 2147483647, F1: 0}
    if true {
        if false {
            fmt.Println(v4, v2)
            v4 = v4
            switch "\n\t" {
            case "hello":
                v4 = (v4 & 2147483647)
            }
        }
        var v6 int = 0
        _ = v6
    } else {
        fmt.Println(v2, v4)
        v2 -= v2
        var v7 int8 = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        if false {
            var v9 uint8 = 2
            var v10 int8 = -21
            if true {
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 string = "hello"
                    v4 = fi8
                    _ = fi11
                    _ = v12
                }
            }
            _ = v10
            _ = v9
        } else {
            for fi13 := 0; fi13 < 12; fi13++ {
                v3[""] = 4294967295
                v14 := GenS1{F0: 1, F1: 1000}
                _ = fi13
                _ = v14
            }
            if (91 > 0) {
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 int16 = 0
                    _ = fi15
                    _ = v16
                }
            } else {
                _ = len(v3)
                switch "foo bar" {
                case "hello":
                    v17 := GenS2{F0: 1, F1: 672}
                    _ = v17
                case "a":
                    var v18 float32 = 3.14159
                    _ = v18
                default:
                    fmt.Println(v4)
                }
            }
            if len(v1) > 0 {
                v1[0] = ("\n\t" + "foo bar")
            }
        }
        v2 = 0
        switch 0 {
        case -9223372036854775808:
            v19 := GenS3{F0: 42}
            _ = v19
        case 4294967295:
            _ = len(v3)
        default:
            switch "hello" {
            case "hello":
                v20 := GenS3{F0: -1}
                _ = v20
            case "foo bar":
                switch (fi8 % (((v4 & fi8)) | 1)) {
                case -258:
                    v2 = 1
                }
            case "":
                v21 := GenS3{F0: 921}
                _ = v21
            }
        }
        _ = fi8
    }
    v4 = (v4 >> ((248) & 63))
    var v22 uint = 98
    if true {
        v1 = append(v1, "foo bar")
        switch (^(v4)) {
        case 772:
            switch (v4 >> ((v22) & 63)) {
            case 32767:
                if true {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        fmt.Println(v22)
                        var v24 string = "foo bar"
                        v25 := GenS4{F0: -1, F1: -77, F2: ""}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    var v26 int8 = 100
                    v26 = (v26 & (v26 * v26))
                    _ = v26
                }
            case -2147483648:
                v22 -= (^(v22))
            default:
                v2 = (v2 / (((v2 | v2)) | 1))
            }
        case 32768:
            v4 += v4
        default:
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 float32 = 1.5
                v4 -= fi27
                _ = fi27
                _ = v28
            }
        }
        v4 = v4
    }
    fmt.Println(v2, v2)
    for fi29 := 0; fi29 < 3; fi29++ {
        if false {
            var v30 uint = (v22 << ((v22) & 63))
            _ = v30
        } else {
            if ((v22 ^ 65535) <= 18446744073709551615) {
                switch v4 {
                case 32767:
                    var v31 uint64 = v2
                    _ = v31
                default:
                    v32 := GenS4{F0: -743, F1: -128, F2: (("a" + "") + (("hello" + ("hello" + "")) + "foo bar"))}
                    _ = v32
                }
                var v33 int16 = -32768
                v34 := v3[""]
                _ = v33
                _ = v34
            } else {
                v22 = (v22 >> ((v22) & 63))
                var v35 uint64 = 0
                v2 = 42
                _ = v35
            }
        }
        if (!(0 < 255)) {
            for fi36 := 0; fi36 < 11; fi36++ {
                v37 := GenS5{F0: 253, F1: 42, F2: 4294967295}
                _ = fi36
                _ = v37
            }
            if (int64(v22) >= -9223372036854775808) {
                v4 -= v4
                v38 := v3[""]
                _ = v38
            } else {
                for fi39 := 0; fi39 < 9; fi39++ {
                    switch v4 {
                    case -771:
                        v3["hello"] = 2147483648
                    case 2147483647:
                        delete(v3, "a")
                    default:
                        v2 -= (v2 % ((4294967295) | 1))
                    }
                    v40 := v3["foo bar"]
                    switch ("\n\t" + "\n\t") {
                    case "":
                        fmt.Println(v40, v4, v4)
                    case "hello":
                        fmt.Println(v22, v4, v40)
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if (42 < 92) {
                        v40 -= v40
                    } else {
                        fmt.Println(v22, fi29, v4)
                        if false {
                            var v41 float64 = float64(1.5)
                            v42 := GenS6{F0: 42}
                            _ = v41
                            _ = v42
                        }
                        v2 = v2
                        _ = len(v1)
                    }
                    _ = fi39
                    _ = v40
                }
                if (!false) {
                    if true {
                        var v43 float64 = float64(0.0)
                        _ = v43
                    }
                }
            }
        }
        _ = fi29
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        switch (fi44 ^ (v4 & v4)) {
        case -549:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -327:
            if len(v1) > 0 {
                v1[0] = (("" + ("foo bar" + "hello")) + "\n\t")
            }
        case 4294967295:
            var v45 int = v4
            _ = v45
        }
        var v46 uint16 = 42
        _ = fi44
        _ = v46
    }
    var v47 uint32 = 2
    var v48 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
