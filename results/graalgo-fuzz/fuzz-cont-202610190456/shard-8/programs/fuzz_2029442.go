package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 string }

func main() {
    v1 := []string{}
    var v2 string = ""
    var v3 uint32 = 37
    var v4 int64 = 9223372036854775807
    var v5 bool = false
    if true {
        if true {
            v6 := GenS1{F0: 0, F1: -882}
            if (!v5) {
                _ = len(v1)
                for fi7 := 0; fi7 < 18; fi7++ {
                    var v8 uint = 775
                    _ = fi7
                    _ = v8
                }
                v5 = v5
            } else {
                fmt.Println(v3, v4, v5)
            }
            _ = v6
        } else {
            var v9 float64 = ((1.5 * (0.0 / 1.0)) - ((3.14159 + 0.1) + 0.0))
            _ = v9
        }
    } else {
        switch v2 {
        case "hello":
            for fi10 := 0; fi10 < 2; fi10++ {
                v3 -= v3
                v3 = 2147483648
                _ = fi10
            }
        default:
            switch 127 {
            case -255:
                if (-1 < 42) {
                    v11 := GenS1{F0: -594, F1: 32767}
                    fmt.Println(v3)
                    _ = v11
                }
            case -1:
                var v12 float32 = 0.0
                _ = v12
            case 4294967295:
                v13 := GenS1{F0: 42, F1: -1}
                _ = v13
            default:
                if (1 >= 151) {
                    if true {
                        var v14 float32 = -2.25
                        for fi15 := 0; fi15 < 6; fi15++ {
                            var v16 int = fi15
                            v3 -= 42
                            _ = fi15
                            _ = v16
                        }
                        var v17 int32 = 0
                        _ = v14
                        _ = v17
                    }
                }
            }
        }
        var v18 int = 128
        _ = v18
    }
    v5 = (356 == uint16(v3))
    v19 := GenS2{F0: (v3 >> ((121) & 63))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v20 := GenS2{F0: 4294967295}
    var v21 string = ((("\n\t" + v2) + "\n\t") + (v2 + ""))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        switch v21 {
        case "":
            v3 = v3
        }
        v21 = v21
        var v23 string = ""
        _ = fi22
        _ = v23
    }
    switch -295 {
    case -818:
        var v24 int = int(v3)
        _ = v24
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -1:
        var v25 int16 = 990
        _ = v25
    default:
        v21 = v2
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        switch (-(fi26)) {
        case 9223372036854775807:
            var v27 int = fi26
            _ = v27
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v28 int64 = v4
            _ = v28
        }
        var v29 float64 = (float64(0.1) * -2.25)
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 float32 = float32(v29)
            v32 := GenS3{F0: (v5 && (1 != 128))}
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi26
        _ = v29
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        v34 := GenS1{F0: -1, F1: 32767}
        if v5 {
            switch v21 {
            case "":
                v2 = ""
            case "foo bar":
                v35 := GenS4{F0: 1, F1: 255}
                _ = v35
            case "hello":
                for fi36 := 0; fi36 < 11; fi36++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if v5 {
                        if v5 {
                            var v37 uint = 990
                            var v38 int8 = 0
                            _ = v37
                            _ = v38
                        }
                        v5 = (uint8(v3) <= 255)
                        switch "\n\t" {
                        case "foo bar":
                            var v39 bool = ("\n\t" >= (("hello" + "a") + ("" + "")))
                            v5 = (0.1 >= 1e10)
                            v5 = false
                            _ = v39
                        case "hello":
                            v40 := GenS5{F0: v21}
                            _ = v40
                        case "\n\t":
                            var v41 uint = uint(v3)
                            _ = v41
                        default:
                            var v42 float64 = float64(0.0)
                            _ = v42
                        }
                    } else {
                        if ((^(v4)) != v4) {
                            v43 := GenS4{F0: 18446744073709551615, F1: 65535}
                            v44 := GenS5{F0: v21}
                            if (v3 < 2147483648) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = v43
                            _ = v44
                        } else {
                            fmt.Println(v21, fi33)
                        }
                        fmt.Println(v2)
                    }
                    _ = fi36
                }
            }
        }
        _ = fi33
        _ = v34
    }
    if len(v1) > 0 {
        v1[0] = v2
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        v3 -= v3
        _ = fi45
    }
    var v46 uint8 = 255
    v5 = false
    var v47 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
