package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int64; F1 int16; F2 uint8 }

func main() {
    var v1 uint16 = 65535
    var v2 float32 = -2.25
    var v3 float32 = ((float32((-2.25 / (float64((3.14159 + 0.1)) / 3.14159))) * v2) * 3.14159)
    var v4 int64 = -357
    v5 := []uint32{529}
    var v6 int64 = v4
    v7 := GenS1{F0: 638, F1: -682, F2: 127}
    var v8 int = 128
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int32 = -847
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int64 = (v4 ^ 0)
            switch fi11 {
            case 1:
                var v13 uint = 4294967295
                _ = v13
            case 4294967295:
                switch (fi9 + (fi11 | (^(fi11)))) {
                case 0:
                    v10 -= 0
                case -797:
                    v1 -= v1
                }
                if true {
                    if (!(18446744073709551615 == 42)) {
                        v14 := GenS2{F0: (("foo bar" + ("\n\t" + "")) + "hello"), F1: 1}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v14
                    } else {
                        fmt.Println(v1, fi11)
                        v10 += 233
                    }
                    fmt.Println(v4, v12)
                    if (!(v1 > v1)) {
                        _ = len(v5)
                    }
                }
                if true {
                    var v15 float32 = v3
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                    if ((((("foo bar" + "") + "hello") + "foo bar") + "hello") == ("hello" + "")) {
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                        v3 += ((float32(((1e10 * 0.0) * (((-2.25 / 1.0) + ((3.14159 - 0.1) * 0.0)) - -2.25))) / float32(1.5)) - (v3 * v3))
                        if (v6 >= 42) {
                            var v16 uint32 = 4294967295
                            _ = v16
                        } else {
                            var v17 int8 = 0
                            if (true || (((4294967295 > 700) && ((int(v1) == (-(fi9))) || (true || false))) && true)) {
                                if false {
                                    fmt.Println(v15)
                                } else {
                                    v12 = v6
                                    switch (("" + ("\n\t" + "foo bar")) + "foo bar") {
                                    case "foo bar":
                                        switch fi9 {
                                        case 246:
                                            var v18 uint8 = 42
                                            _ = v18
                                        case 982:
                                            var v19 uint32 = 2
                                            _ = v19
                                        }
                                    case "\n\t":
                                        var v20 float64 = 0.1
                                        _ = v20
                                    }
                                }
                            } else {
                                var v21 float64 = 3.14159
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v22 := GenS2{F0: ("hello" + "foo bar"), F1: -1}
                                v17 += (v17 + v17)
                                _ = v21
                                _ = v22
                            }
                            _ = v17
                        }
                    }
                    v6 = v12
                    _ = v15
                }
            }
            _ = fi11
            _ = v12
        }
        v23 := GenS3{F0: false, F1: v10}
        _ = fi9
        _ = v10
        _ = v23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        switch (-(fi24)) {
        case -9223372036854775808:
            v1 += (v1 / (((v1 - 0)) | 1))
        case -1:
            v1 = 1
        }
        v2 -= v3
        var v25 int32 = 42
        _ = fi24
        _ = v25
    }
    fmt.Println(v8)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v26 int64 = v4
    if len(v5) > 0 {
        v5[0] = 0
    }
    var v27 float32 = 1.5
    var v28 uint32 = 398
    var v29 float32 = (float32(-2.25) * float32((1.5 + 3.14159)))
    var v30 uint16 = (v1 + 1)
    var v31 bool = true
    switch "\n\t" {
    case "hello":
        fmt.Println(v29, v1)
    case "a":
        var v32 int64 = 346
        _ = v32
    default:
        v33 := GenS1{F0: (v8 | (v8 << ((255) & 63))), F1: 0, F2: 4294967295}
        _ = v33
    }
    fmt.Println(v26)
    if (14 >= 127) {
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 string = ("a" + ("" + ("\n\t" + "")))
            _ = fi34
            _ = v35
        }
    } else {
        v36 := GenS4{F0: (-(v6)), F1: int16(v30), F2: 128}
        _ = v36
    }
    _ = len(v5)
    if v31 {
        var v37 int32 = 1
        var v38 bool = v31
        _ = v37
        _ = v38
    } else {
        if (!true) {
            if len(v5) > 0 {
                v5[0] = uint32(v30)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
