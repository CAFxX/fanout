package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint; F1 uint64; F2 string }
type GenS3 struct { F0 int8; F1 uint; F2 string }
type GenS4 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := []string{"", "hello"}
    var v2 uint8 = 127
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    v2 = 128
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: -32768, F1: 354}
        switch fi6 {
        case -258:
            v1 = append(v1, "hello")
        case -1:
            v8 := v3[(fi6 ^ fi6)]
            _ = v8
        case 9223372036854775807:
            var v9 bool = (!((687 <= 2) || (0 > -50)))
            _ = v9
        default:
            var v10 uint64 = 730
            _ = v10
        }
        v11 := GenS1{F0: -1, F1: -1}
        _ = fi6
        _ = v11
        _ = v7
    }
    v12 := GenS1{F0: -32768, F1: 0}
    var v13 int = 0
    var v14 uint = 147
    if len(v1) > 0 {
        v1[0] = (("a" + ("a" + ("hello" + ("foo bar" + "hello")))) + "foo bar")
    }
    v15 := GenS2{F0: (^(v14)), F1: 1, F2: "hello"}
    fmt.Println(v14, v13)
    v2 += 1
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS2{F0: (v14 + 255), F1: 654, F2: "hello"}
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 uint = (v14 - (v14 << (((^(v14))) & 63)))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 = uint8(v19)
            v14 += uint(v2)
            var v20 int = fi16
            _ = fi18
            _ = v19
            _ = v20
        }
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            var v21 uint64 = 1
            _ = v21
        case "\n\t":
            for fi22 := 0; fi22 < 4; fi22++ {
                switch ("hello" + "a") {
                case "foo bar":
                    _ = len(v5)
                case "\n\t":
                    if false {
                        var v23 int8 = -1
                        v23 += 0
                        _ = v23
                    } else {
                        var v24 int = 128
                        v14 = (^(v14))
                        var v25 float64 = (3.14159 - (1e10 * 0.1))
                        _ = v24
                        _ = v25
                    }
                }
                var v26 int32 = 0
                _ = fi22
                _ = v26
            }
        }
        _ = fi16
        _ = v17
    }
    v4[608] = (v13 - 42)
    if (!((121 >= -128) && (2 != 42))) {
        if true {
            v2 -= 255
            var v27 uint32 = 2147483648
            switch ((("foo bar" + "foo bar") + "") + "\n\t") {
            case "":
                for fi28 := 0; fi28 < 16; fi28++ {
                    switch "hello" {
                    case "foo bar":
                        if true {
                            v2 = v2
                        }
                    case "hello":
                        for fi29 := 0; fi29 < 5; fi29++ {
                            var v30 uint64 = 2
                            _ = fi29
                            _ = v30
                        }
                    default:
                        if len(v1) > 0 {
                            v1[0] = (("hello" + "\n\t") + "a")
                        }
                    }
                    v13 += 921
                    fmt.Println(v2, v27, v13)
                    _ = fi28
                }
            }
            _ = v27
        } else {
            v13 = -294
            if true {
                var v31 int32 = 1
                fmt.Println(v13)
                fmt.Println(v31, v2, v31)
                _ = v31
            }
        }
        fmt.Println(v13, v2)
        v14 = v14
    }
    var v32 int16 = -519
    for fi33 := 0; fi33 < 5; fi33++ {
        switch 527 {
        case -9223372036854775808:
            v34 := GenS3{F0: -10, F1: (v14 | 375), F2: ("\n\t" + "foo bar")}
            _ = v34
        case 128:
            switch (v13 + int(v2)) {
            case -478:
                _ = len(v1)
            case 436:
                switch "" {
                case "hello":
                    var v35 uint8 = v2
                    _ = v35
                case "foo bar":
                    v4[2147483647] = v13
                }
            case -9223372036854775808:
                v2 += v2
            }
        case 0:
            fmt.Println(fi33, v2, fi33)
        default:
            fmt.Println(fi33, v32, v14)
        }
        var v36 uint16 = 0
        var v37 float32 = ((((1.5 - 3.14159) * 0.1) + 3.14159) - 1e10)
        _ = fi33
        _ = v36
        _ = v37
    }
    v2 -= v2
    v38 := GenS1{F0: -765, F1: -2147483648}
    for fi39 := 0; fi39 < 7; fi39++ {
        if false {
            if false {
                if false {
                    for fi40 := 0; fi40 < 6; fi40++ {
                        v13 += v13
                        _ = fi40
                    }
                    var v41 uint8 = uint8(v14)
                    fmt.Println(v32)
                    _ = v41
                } else {
                    var v42 float32 = ((3.14159 + 1.5) - -2.25)
                    v13 += v13
                    _ = v42
                }
                var v43 bool = false
                _ = v43
            }
        } else {
            for fi44 := 0; fi44 < 17; fi44++ {
                if (189 < 1) {
                    v5[816] = "foo bar"
                    switch "" {
                    case "a":
                        v2 = (v2 * v2)
                    case "foo bar":
                        var v45 int16 = (v32 + 1)
                        _ = v45
                    case "hello":
                        var v46 float64 = 0.0
                        _ = v46
                    }
                    switch fi39 {
                    case 639:
                        fmt.Println(v2, v2, v14)
                    default:
                        if (!true) {
                            var v47 int64 = -503
                            fmt.Println(fi39, v47, v14)
                            _ = v47
                        }
                    }
                    _ = len(v1)
                }
                var v48 string = ("" + ("hello" + ("a" + "\n\t")))
                _ = fi44
                _ = v48
            }
        }
        fmt.Println(v2)
        if false {
            for fi49 := 0; fi49 < 13; fi49++ {
                v50 := v3[(fi39 ^ (fi49 + (-(fi39))))]
                fmt.Println(v32, fi39)
                _ = fi49
                _ = v50
            }
            v51 := GenS4{F0: 245, F1: 215}
            for fi52 := 0; fi52 < 4; fi52++ {
                v13 += (fi52 | fi52)
                switch ("\n\t" + "") {
                case "":
                    v32 -= v32
                case "\n\t":
                    if (v2 < 127) {
                        _ = len(v4)
                    }
                }
                fmt.Println(v2)
                _ = fi52
            }
            _ = v51
        }
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
}
