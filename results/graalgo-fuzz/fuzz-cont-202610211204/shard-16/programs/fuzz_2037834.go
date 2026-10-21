package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8; F1 int32; F2 int }
type GenS3 struct { F0 string; F1 int32; F2 int64 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int; F1 uint64; F2 int64 }
type GenS6 struct { F0 int64; F1 uint64; F2 int32 }
type GenS7 struct { F0 int8; F1 bool; F2 uint64 }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := []string{"a", "\n\t", ((("foo bar" + "foo bar") + "a") + ("hello" + (("foo bar" + "a") + ("" + "hello")))), "foo bar"}
    var v4 int8 = 0
    if false {
        for fi5 := 0; fi5 < 18; fi5++ {
            if ((((3.14159 * 0.1) + 1.5) * ((3.14159 + float32(-2.25)) - ((1e10 * 0.0) - 0.0))) > float32(0.1)) {
                v6 := GenS1{F0: -1, F1: 42}
                v4 -= v4
                _ = v6
            } else {
                if true {
                    v4 = v4
                    v4 += (v4 << ((65535) & 63))
                    for fi7 := 0; fi7 < 14; fi7++ {
                        v8 := GenS1{F0: 1, F1: 42}
                        _ = fi7
                        _ = v8
                    }
                } else {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        v4 -= (-(v4))
                        v4 = (^(v4))
                        _ = fi9
                    }
                    v4 -= 1
                }
                var v10 uint8 = 1
                if false {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        var v12 uint8 = v10
                        _ = fi11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 12; fi13++ {
                        delete(v1, (^(fi13)))
                        _ = fi13
                    }
                    v14 := GenS2{F0: (v10 | (v10 ^ v10)), F1: 1, F2: fi5}
                    _ = v14
                } else {
                    v1[fi5] = ("a" + "hello")
                }
                _ = v10
            }
            _ = fi5
        }
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        switch "a" {
        case "\n\t":
            for fi16 := 0; fi16 < 19; fi16++ {
                if (!(v4 == 42)) {
                    var v17 uint64 = 4294967295
                    _ = v17
                } else {
                    var v18 float32 = float32((1.5 + (((1e10 / -2.25) / float64((((float32(0.1) + (0.1 - 1.5)) * (0.1 * (3.14159 - -2.25))) * 1e10))) + float64(-2.25))))
                    _ = v18
                }
                _ = fi16
            }
        }
        v19 := GenS2{F0: 152, F1: 42, F2: (fi15 * (fi15 & fi15))}
        fmt.Println(v4, v4)
        _ = fi15
        _ = v19
    }
    if (42 < 0) {
        fmt.Println(v4)
        v20 := GenS3{F0: "foo bar", F1: -276, F2: -9223372036854775808}
        _ = v20
    } else {
        var v21 uint32 = 360
        _ = v21
    }
    switch 0 {
    case 9223372036854775807:
        v22 := GenS3{F0: "", F1: 746, F2: 808}
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(fi23)
        v24 := GenS3{F0: "hello", F1: -2147483648, F2: 9223372036854775807}
        _ = fi23
        _ = v24
    }
    if (true && false) {
        var v25 int16 = -1
        var v26 int = 732
        v26 = v26
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 9; fi27++ {
            switch "\n\t" {
            case "\n\t":
                v4 = (v4 / ((v4) | 1))
            case "hello":
                v28 := GenS1{F0: int16(fi27), F1: 352}
                _ = v28
            }
            _ = fi27
        }
        fmt.Println(v4)
        var v29 uint8 = 1
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        if false {
            for fi31 := 0; fi31 < 19; fi31++ {
                v4 = (^(v4))
                v4 -= -128
                _ = fi31
            }
            v32 := GenS3{F0: "hello", F1: -1, F2: -9223372036854775808}
            _ = v32
        } else {
            if ((!true) || (fi30 >= fi30)) {
                switch "hello" {
                case "":
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v34 := GenS4{F0: fi33, F1: -292}
                        fmt.Println(fi30, fi30)
                        var v35 int8 = v4
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                default:
                    v36 := v1[(-(fi30))]
                    _ = v36
                }
            } else {
                var v37 int8 = v4
                fmt.Println(v37)
                v38 := GenS5{F0: -654, F1: 0, F2: -906}
                _ = v37
                _ = v38
            }
        }
        fmt.Println(v4, v4)
        _ = fi30
    }
    var v39 uint32 = 1
    var v40 float32 = 0.0
    for fi41 := 0; fi41 < 6; fi41++ {
        if len(v3) > 0 {
            v3[0] = ("\n\t" + (("hello" + "") + "foo bar"))
        }
        v42 := GenS6{F0: 9223372036854775807, F1: 0, F2: -1}
        _ = fi41
        _ = v42
    }
    v43 := GenS7{F0: v4, F1: (true && true), F2: 42}
    for fi44 := 0; fi44 < 18; fi44++ {
        if true {
            if (true || ((v40 + v40) != 1.5)) {
                var v45 uint8 = 42
                if true {
                    v3 = append(v3, ((("hello" + (("a" + "a") + ("" + "foo bar"))) + "\n\t") + "hello"))
                    v4 += int8(fi44)
                    switch "a" {
                    case "":
                        fmt.Println(v4, fi44, v40)
                    }
                } else {
                    fmt.Println(v40)
                }
                _ = v45
            }
            fmt.Println(v4, v40, v4)
            for fi46 := 0; fi46 < 1; fi46++ {
                switch "\n\t" {
                case "":
                    var v47 string = (("a" + "") + ("a" + ""))
                    _ = v47
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = ("" + "a")
                    }
                    v48 := GenS2{F0: 0, F1: 0, F2: -9223372036854775808}
                    v49 := v1[(fi44 - (-(fi46)))]
                    _ = v48
                    _ = v49
                case "\n\t":
                    v50 := GenS8{F0: 19, F1: v39}
                    _ = v50
                default:
                    var v51 int8 = 43
                    _ = v51
                }
                var v52 int = fi46
                if (!(42 <= 42)) {
                    switch (fi44 | fi44) {
                    case 32767:
                        switch (fi46 + fi44) {
                        case 128:
                            v4 = v4
                        default:
                            switch -2147483648 {
                            case 42:
                                v53 := GenS9{F0: 42, F1: 0}
                                _ = v53
                            default:
                                v40 -= v40
                            }
                        }
                    }
                    _ = len(v2)
                }
                _ = fi46
                _ = v52
            }
        }
        for fi54 := 0; fi54 < 15; fi54++ {
            if true {
                if false {
                    v4 += v4
                } else {
                    v55 := GenS5{F0: (-(fi54)), F1: 42, F2: -591}
                    _ = v55
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch (fi54 + (fi54 * 1)) {
                case 128:
                    if len(v3) > 0 {
                        v3[0] = ("" + ("a" + ("hello" + "")))
                    }
                case -988:
                    fmt.Println(v39, v39)
                }
                var v56 int64 = -9223372036854775808
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v57 := GenS6{F0: v56, F1: 2, F2: 42}
                _ = v56
                _ = v57
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi54
        }
        v4 -= v4
        _ = fi44
    }
    v3 = append(v3, "a")
    v4 = v4
    if (-1 < -572) {
        switch 32768 {
        case 9223372036854775807:
            switch 2147483647 {
            case -2147483648:
                if (9223372036854775807 == -1) {
                    for fi58 := 0; fi58 < 14; fi58++ {
                        v1[-333] = "a"
                        _ = fi58
                    }
                } else {
                    delete(v2, v39)
                }
            case 32767:
                switch -1 {
                case 1:
                    v3 = append(v3, "a")
                case 32768:
                    v59 := GenS10{F0: "a", F1: v39}
                    _ = v59
                case 42:
                    var v60 uint = 42
                    _ = v60
                }
            case 4294967295:
                var v61 int32 = 42
                _ = v61
            }
        }
        v62 := GenS2{F0: 242, F1: -2147483648, F2: -846}
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
}
