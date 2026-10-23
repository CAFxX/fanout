package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64 }

func main() {
    v1 := []string{(("a" + ("" + "hello")) + "a"), "a", ("foo bar" + ((((("a" + (("\n\t" + "\n\t") + "")) + "hello") + "foo bar") + "a") + "foo bar"))}
    var v2 uint64 = 4294967295
    var v3 float32 = (-2.25 * ((0.1 - ((-2.25 / 1e10) - (0.0 * 0.1))) - 1.5))
    var v4 int64 = int64(v2)
    v2 -= v2
    var v5 bool = (-909 != -703)
    for fi6 := 0; fi6 < 13; fi6++ {
        v4 = int64(v2)
        var v7 uint16 = 32767
        _ = fi6
        _ = v7
    }
    var v8 string = (("" + "\n\t") + "a")
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint64 = (^(v2))
        v11 := GenS1{F0: (125 != (^(v10)))}
        if v5 {
            var v12 uint64 = 2
            fmt.Println(v10)
            v1 = append(v1, v8)
            _ = v12
        } else {
            fmt.Println(fi9)
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v13 string = (v8 + (v8 + "a"))
    switch v8 {
    case "":
        var v14 uint64 = 2
        _ = v14
    }
    v4 = -297
    for fi15 := 0; fi15 < 15; fi15++ {
        if len(v1) > 0 {
            v1[0] = (v8 + v13)
        }
        _ = fi15
    }
    if v5 {
        if v5 {
            var v16 int = 354
            _ = len(v1)
            _ = v16
        } else {
            fmt.Println(v4, v8, v5)
        }
        v8 = ("foo bar" + v13)
    }
    if (0 < 2) {
        for fi17 := 0; fi17 < 2; fi17++ {
            fmt.Println(v5)
            var v18 bool = v5
            v5 = v5
            _ = fi17
            _ = v18
        }
        switch v13 {
        case "foo bar":
            var v19 int64 = (v4 - (-(v4)))
            _ = v19
        case "":
            switch v8 {
            case "a":
                v20 := GenS2{F0: 308, F1: 151}
                _ = v20
            case "foo bar":
                if (0 != v4) {
                    fmt.Println(v4)
                    if v5 {
                        fmt.Println(v13, v2)
                        v21 := GenS3{F0: 2147483647, F1: v4, F2: -9223372036854775808}
                        v22 := GenS3{F0: 0, F1: v4, F2: 128}
                        _ = v21
                        _ = v22
                    } else {
                        var v23 int32 = 42
                        for fi24 := 0; fi24 < 17; fi24++ {
                            var v25 int32 = v23
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    }
                }
            case "\n\t":
                if ("foo bar" == v8) {
                    v26 := GenS4{F0: 786, F1: 42}
                    var v27 float64 = 1e10
                    _ = v26
                    _ = v27
                } else {
                    var v28 float64 = -2.25
                    _ = v28
                }
            }
        }
        if (!v5) {
            var v29 int32 = 42
            v3 = (3.14159 - v3)
            _ = v29
        }
    } else {
        if (!false) {
            for fi30 := 0; fi30 < 5; fi30++ {
                if (fi30 > (^(fi30))) {
                    var v31 int64 = v4
                    _ = len(v1)
                    v32 := GenS4{F0: 32767, F1: 0}
                    _ = v31
                    _ = v32
                } else {
                    v33 := GenS5{F0: 1}
                    v1 = append(v1, ("a" + (("hello" + v13) + ("" + ""))))
                    if (!(v5 && (600 > 32767))) {
                        v34 := GenS6{F0: 519, F1: 42}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v2)
                        _ = v34
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    }
                    _ = v33
                }
                switch -120 {
                case 2147483647:
                    v13 = v13
                }
                switch ("" + (v8 + (((((v13 + "") + v8) + v13) + v13) + v13))) {
                case "":
                    if (!(0 == 0)) {
                        var v35 uint16 = 878
                        fmt.Println(fi30, v35, v3)
                        v36 := GenS7{F0: 907}
                        _ = v35
                        _ = v36
                    }
                default:
                    var v37 int64 = (v4 & (v4 >> ((65535) & 63)))
                    _ = v37
                }
                _ = fi30
            }
        }
        for fi38 := 0; fi38 < 8; fi38++ {
            if (v5 && v5) {
                var v39 int8 = 42
                v5 = v5
                _ = v39
            } else {
                v4 += -391
            }
            _ = fi38
        }
    }
    fmt.Println(v8)
    var v40 int16 = -1
    for fi41 := 0; fi41 < 15; fi41++ {
        v8 = (((((v13 + v13) + v8) + ((v13 + "hello") + v13)) + ("foo bar" + (v8 + v13))) + v13)
        switch v13 {
        case "a":
            v4 = 9223372036854775807
        case "":
            var v42 uint16 = 42
            _ = v42
        case "\n\t":
            switch ("foo bar" + v8) {
            case "\n\t":
                for fi43 := 0; fi43 < 14; fi43++ {
                    fmt.Println(v40)
                    _ = fi43
                }
            default:
                if v5 {
                    var v44 float32 = (v3 * (v3 / v3))
                    var v45 uint16 = 401
                    if len(v1) > 0 {
                        v1[0] = v8
                    }
                    _ = v44
                    _ = v45
                } else {
                    v46 := GenS3{F0: fi41, F1: 42, F2: 0}
                    _ = v46
                }
            }
        }
        _ = fi41
    }
    var v47 int32 = -1
    v48 := GenS1{F0: true}
    if (v2 == v2) {
        var v49 float64 = 1.5
        v50 := GenS2{F0: 4294967295, F1: 0}
        if false {
            for fi51 := 0; fi51 < 11; fi51++ {
                if (v2 <= (v2 + v2)) {
                    v5 = (65535 <= 1)
                } else {
                    v40 = v40
                    v8 = ((((v13 + (((("" + ((v8 + (v13 + v8)) + "\n\t")) + ("foo bar" + v13)) + "foo bar") + v8)) + v13) + "") + v13)
                }
                v52 := GenS8{F0: 987}
                _ = fi51
                _ = v52
            }
            v8 = v8
            var v53 bool = v5
            _ = v53
        } else {
            v13 = "\n\t"
            if v5 {
                fmt.Println(v3)
            }
        }
        _ = v49
        _ = v50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v48
    _ = v5
    _ = v8
}
