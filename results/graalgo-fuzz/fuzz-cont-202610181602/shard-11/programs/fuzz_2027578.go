package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int32; F1 string; F2 string }
type GenS5 struct { F0 uint8; F1 uint64 }
func genFn0() uint32 {
    if false {
        var v7 int8 = 91
        fmt.Println(v7)
        v8 := GenS1{F0: -719, F1: 32767}
        _ = v7
        _ = v8
    } else {
        var v9 string = (("hello" + "foo bar") + "a")
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS2{F0: 523}
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    switch 4294967295 {
    case 2147483647:
        var v12 bool = (1e10 != 1.5)
        _ = v12
    case 32768:
        if false {
            if false {
                var v13 uint = 0
                var v14 int = -346
                _ = v13
                _ = v14
            } else {
                var v15 string = ""
                for fi16 := 0; fi16 < 19; fi16++ {
                    fmt.Println(v15, v15)
                    v15 = (v15 + (v15 + v15))
                    _ = fi16
                }
                v15 = "a"
                _ = v15
            }
        } else {
            var v17 int32 = 1
            _ = v17
        }
    case -2147483648:
        var v18 int8 = -46
        _ = v18
    }
    switch -969 {
    case 2147483647:
        var v19 uint32 = 1
        _ = v19
    case 0:
        var v20 int64 = 9223372036854775807
        _ = v20
    case 335:
        var v21 uint16 = 42
        _ = v21
    default:
        var v22 int64 = 807
        _ = v22
    }
    var v23 uint16 = 32767
    _ = v23
    return 42
}

func main() {
    var v1 uint32 = 0
    v2 := []string{}
    var v3 int = 42
    v4 := make(map[int]string)
    var v5 float32 = ((3.14159 + 3.14159) * 1.5)
    var v6 uint16 = 65535
    var v24 int32 = 42
    _ = len(v2)
    var v25 float64 = 3.14159
    v5 = (float32((v25 + v25)) + ((3.14159 * v5) / v5))
    fmt.Println(v3)
    v3 -= v3
    for fi26 := 0; fi26 < 13; fi26++ {
        v3 = fi26
        switch "" {
        case "":
            switch v3 {
            case -425:
                v27 := GenS2{F0: v6}
                _ = v27
            }
        case "a":
            var v28 int16 = 1000
            _ = v28
        case "\n\t":
            fmt.Println(fi26, v6, v1)
        default:
            v29 := GenS3{F0: 1, F1: 1}
            _ = v29
        }
        _ = fi26
    }
    var v30 string = "foo bar"
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS1{F0: 1, F1: 28}
        switch v3 {
        case 128:
            fmt.Println(v25, v3)
        case 1:
            if (1 <= -1) {
                fmt.Println(v5, v30)
            }
        default:
            switch v30 {
            case "a":
                _ = len(v2)
            case "hello":
                var v33 uint = 4294967295
                _ = v33
            case "\n\t":
                if false {
                    if true {
                        var v34 int16 = 1
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 1; fi35++ {
                            var v36 float64 = ((((v25 * -2.25) + v25) * ((0.0 - 0.0) + 0.0)) - 3.14159)
                            _ = fi35
                            _ = v36
                        }
                        var v37 int16 = -225
                        _ = v37
                    }
                } else {
                    v6 -= 1
                }
            default:
                switch (("foo bar" + "foo bar") + ("hello" + v30)) {
                case "\n\t":
                    v25 = v25
                default:
                    v30 = ((("foo bar" + "") + v30) + v30)
                }
            }
        }
        _ = fi31
        _ = v32
    }
    v38 := v4[(-(v3))]
    v24 += (v24 % ((v24) | 1))
    if false {
        var v39 int64 = -9223372036854775808
        v40 := GenS4{F0: v24, F1: v38, F2: v38}
        _ = v39
        _ = v40
    } else {
        switch (v3 << ((439) & 63)) {
        case 128:
            v30 = v38
        case 0:
            for fi41 := 0; fi41 < 18; fi41++ {
                var v42 int16 = 1
                v43 := GenS5{F0: 138, F1: 42}
                switch ("" + "foo bar") {
                case "":
                    if len(v2) > 0 {
                        v2[0] = v38
                    }
                default:
                    switch v38 {
                    case "":
                        var v44 float64 = 0.1
                        _ = v44
                    }
                }
                _ = fi41
                _ = v42
                _ = v43
            }
        case -845:
            var v45 uint64 = 42
            _ = v45
        default:
            if len(v2) > 0 {
                v2[0] = v30
            }
        }
    }
    var v46 int8 = int8(v24)
    v47 := GenS3{F0: 2, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
