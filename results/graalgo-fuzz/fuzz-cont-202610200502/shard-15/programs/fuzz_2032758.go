package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []string{}
    var v2 int8 = -128
    var v3 float32 = 1.5
    v4 := []string{"", "hello", ""}
    var v5 int64 = -1
    var v6 int = -9223372036854775808
    v2 = v2
    if len(v1) > 0 {
        v1[0] = ""
    }
    if (!(42 > 32767)) {
        fmt.Println(v3)
        switch 128 {
        case -2147483648:
            if (true || true) {
                _ = len(v4)
                v7 := GenS1{F0: 198}
                _ = v7
            }
        case 531:
            if true {
                var v8 int32 = 643
                for fi9 := 0; fi9 < 10; fi9++ {
                    if len(v4) > 0 {
                        v4[0] = ("" + ("hello" + (("foo bar" + ("hello" + "")) + "\n\t")))
                    }
                    switch (-(v6)) {
                    case -218:
                        v6 = v6
                    }
                    switch "hello" {
                    case "foo bar":
                        switch (v6 - 127) {
                        case -423:
                            v10 := GenS1{F0: int16(v2)}
                            _ = v10
                        case 2147483647:
                            v11 := GenS1{F0: 3}
                            _ = v11
                        case 32768:
                            var v12 string = "\n\t"
                            _ = v12
                        default:
                            v8 = (v8 * 2147483647)
                        }
                    }
                    _ = fi9
                }
                _ = v8
            }
        default:
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 uint = 712
                _ = fi13
                _ = v14
            }
        }
        v3 += v3
    }
    switch v6 {
    case 634:
        var v15 int = v6
        _ = v15
    case -9223372036854775808:
        if (42 <= -1) {
            fmt.Println(v6)
            if true {
                _ = len(v1)
                if (!true) {
                    fmt.Println(v5, v2)
                } else {
                    if false {
                        var v16 int32 = 42
                        _ = v16
                    }
                    v6 = 999
                }
            } else {
                v6 -= 32767
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                v4 = append(v4, "foo bar")
            }
        } else {
            v6 = v6
        }
    case 32768:
        var v17 int16 = -1
        _ = v17
    }
    var v18 uint32 = 2
    if (4294967295 > 42) {
        v6 += v6
        if (2 > 2) {
            v19 := GenS1{F0: 32767}
            _ = v19
        } else {
            v6 -= v6
            v2 = 123
            if len(v4) > 0 {
                v4[0] = ("hello" + "")
            }
        }
    }
    var v20 int32 = 2147483647
    v21 := GenS2{F0: 4294967295, F1: v18}
    for fi22 := 0; fi22 < 9; fi22++ {
        if true {
            for fi23 := 0; fi23 < 2; fi23++ {
                if false {
                    v4 = append(v4, "a")
                } else {
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "foo bar")
                    }
                }
                _ = fi23
            }
            v20 -= v20
            var v24 int64 = v5
            _ = v24
        } else {
            if true {
                var v25 uint = 1
                _ = v25
            }
            var v26 uint32 = 508
            _ = v26
        }
        if false {
            fmt.Println(v20)
        }
        _ = fi22
    }
    switch v6 {
    case -426:
        for fi27 := 0; fi27 < 14; fi27++ {
            var v28 uint8 = 127
            v20 -= 1
            v3 = float32(0.1)
            _ = fi27
            _ = v28
        }
    default:
        var v29 float64 = 3.14159
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 int8 = (-(v2))
        var v32 string = (("a" + (((("" + "\n\t") + "hello") + (("\n\t" + "hello") + ("" + ("\n\t" + "hello")))) + "\n\t")) + ("\n\t" + ("" + ("\n\t" + "hello"))))
        v5 -= v5
        _ = fi30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        v6 -= (v6 << ((18446744073709551615) & 63))
        var v34 int64 = (-(v5))
        switch ("foo bar" + ((((("hello" + "foo bar") + (("a" + ("hello" + "\n\t")) + "\n\t")) + ("a" + "foo bar")) + "a") + "hello")) {
        case "a":
            v35 := GenS1{F0: 42}
            _ = v35
        default:
            v36 := GenS1{F0: 1000}
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    var v37 uint16 = 2
    if ((-1 < int16(v37)) || ((true || true) && true)) {
        for fi38 := 0; fi38 < 6; fi38++ {
            if (1e10 == ((0.0 - (3.14159 * (float64((v3 / (float32(((-2.25 * float64(v3)) - ((-2.25 + (0.1 + (3.14159 / ((0.1 * 0.1) / 1e10)))) / -2.25))) + v3))) * (-2.25 + 1.5)))) - (0.0 - ((-2.25 - 0.0) - -2.25)))) {
                var v39 int8 = v2
                for fi40 := 0; fi40 < 7; fi40++ {
                    v4 = append(v4, ((("foo bar" + "\n\t") + "hello") + "\n\t"))
                    v6 -= fi38
                    fmt.Println(fi38, v5)
                    var v41 uint64 = 1
                    _ = fi40
                    _ = v41
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v39
            }
            v42 := GenS3{F0: v5}
            for fi43 := 0; fi43 < 2; fi43++ {
                var v44 int8 = v2
                var v45 string = "foo bar"
                _ = fi43
                _ = v44
                _ = v45
            }
            _ = fi38
            _ = v42
        }
    } else {
        for fi46 := 0; fi46 < 8; fi46++ {
            for fi47 := 0; fi47 < 10; fi47++ {
                fmt.Println(v3)
                _ = fi47
            }
            _ = fi46
        }
        v2 -= (-(v2))
        for fi48 := 0; fi48 < 3; fi48++ {
            v4 = append(v4, "\n\t")
            v20 -= v20
            _ = fi48
        }
    }
    switch (^(v6)) {
    case 667:
        if true {
            v1 = append(v1, "hello")
            if (false || (false && false)) {
                v2 = v2
            }
        }
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        fmt.Println(v5, v37, v3)
        v1 = append(v1, ("" + "\n\t"))
        var v50 uint8 = 204
        _ = fi49
        _ = v50
    }
    var v51 uint16 = (v37 + (v37 >> ((2) & 63)))
    if (false && ((18446744073709551615 <= 18446744073709551615) && (((1.5 - 3.14159) / 0.1) >= (1e10 * 3.14159)))) {
        for fi52 := 0; fi52 < 1; fi52++ {
            for fi53 := 0; fi53 < 17; fi53++ {
                v54 := GenS1{F0: -462}
                var v55 uint32 = v18
                v56 := GenS1{F0: 0}
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
            }
            _ = len(v1)
            _ = fi52
        }
        if (994 > 4294967295) {
            v3 = ((v3 - -2.25) * v3)
            for fi57 := 0; fi57 < 10; fi57++ {
                v20 = -529
                _ = fi57
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
