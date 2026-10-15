package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 int32; F2 uint16 }

func main() {
    v1 := []uint32{531, 2, 911}
    var v2 uint8 = 127
    v3 := []int{9223372036854775807, 2147483647, 32767}
    var v4 uint64 = 979
    v5 := []string{("hello" + "\n\t"), (((("\n\t" + ("hello" + "")) + "hello") + "hello") + ""), "\n\t"}
    var v6 uint = 65535
    var v7 int = 408
    v8 := GenS1{F0: v4, F1: 535, F2: v2}
    if (!true) {
        var v9 int32 = 755
        v4 += (v4 / (((v4 & v4)) | 1))
        _ = v9
    } else {
        v2 = v2
    }
    switch ("\n\t" + "\n\t") {
    case "":
        var v10 uint64 = v4
        _ = v10
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch ((("\n\t" + (("\n\t" + (("hello" + "foo bar") + (("a" + ("a" + (("hello" + "a") + "\n\t"))) + ("a" + "\n\t")))) + "")) + "a") + ("a" + "foo bar")) {
    case "hello":
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 6; fi12++ {
                v2 -= uint8(v4)
                var v13 uint32 = 45
                _ = fi12
                _ = v13
            }
            if (!(1 < 646)) {
                switch -932 {
                case -1:
                    var v14 bool = (-2147483648 < -2147483648)
                    _ = v14
                case 376:
                    fmt.Println(fi11, v7, v4)
                case 127:
                    for fi15 := 0; fi15 < 3; fi15++ {
                        if len(v3) > 0 {
                            v3[0] = (v7 % ((0) | 1))
                        }
                        _ = fi15
                    }
                default:
                    for fi16 := 0; fi16 < 3; fi16++ {
                        fmt.Println(fi11, fi16)
                        _ = fi16
                    }
                }
                for fi17 := 0; fi17 < 18; fi17++ {
                    if ((v4 << (((^(v6))) & 63)) >= (v4 & v4)) {
                        var v18 float32 = 1e10
                        v4 += v4
                        switch "hello" {
                        case "":
                            v19 := GenS1{F0: v4, F1: 368, F2: uint8(v4)}
                            _ = v19
                        default:
                            if (v4 <= 42) {
                                if len(v3) > 0 {
                                    v3[0] = 4294967295
                                }
                                v6 = 65535
                                v20 := GenS1{F0: v4, F1: 749, F2: v2}
                                if (false && false) {
                                    var v21 uint = (^(v6))
                                    _ = v21
                                } else {
                                    _ = len(v1)
                                    if (0 != 1) {
                                        v22 := GenS1{F0: v4, F1: 0, F2: (v2 | (v2 | 128))}
                                        v2 = 185
                                        var v23 int = fi17
                                        v3 = append(v3, (-(fi11)))
                                        var v24 int = (fi11 * 42)
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    } else {
                                        v2 = 178
                                        v1 = append(v1, 42)
                                    }
                                }
                                v25 := GenS2{F0: false, F1: fi17}
                                _ = v20
                                _ = v25
                            } else {
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch "" {
                                    case "a":
                                        var v26 float32 = v18
                                        _ = v26
                                    case "hello":
                                        _ = len(v5)
                                    case "foo bar":
                                        v27 := GenS3{F0: (v4 | v4), F1: 42, F2: 881}
                                        _ = v27
                                    }
                                    _ = len(v1)
                                } else {
                                    switch 9223372036854775807 {
                                    case 4294967295:
                                        var v28 float32 = 3.14159
                                        _ = v28
                                    case 9223372036854775807:
                                        var v29 float32 = v18
                                        _ = v29
                                    case -9223372036854775808:
                                        if false {
                                            switch (("a" + "hello") + "a") {
                                            case "":
                                                fmt.Println(v2, fi17, v2)
                                            case "hello":
                                                if len(v3) > 0 {
                                                    v3[0] = -297
                                                }
                                            case "a":
                                                fmt.Println(v7, v18, fi11)
                                                var v30 int64 = 0
                                                _ = v30
                                            default:
                                                if (false || true) {
                                                    var v31 bool = (-1 != int64(v6))
                                                    _ = v31
                                                } else {
                                                    v18 += (0.1 - v18)
                                                    switch fi11 {
                                                    case -9223372036854775808:
                                                        if ((-979 != 0) || ((!false) && false)) {
                                                            var v32 uint = v6
                                                            v6 = (^(v6))
                                                            v33 := GenS2{F0: ("" > "foo bar"), F1: (-(v7))}
                                                            _ = v32
                                                            _ = v33
                                                        } else {
                                                            var v34 bool = ((v2 % ((v2) | 1)) != (v2 + 242))
                                                            switch (^(v7)) {
                                                            case 836:
                                                                var v35 float64 = 0.1
                                                                _ = v35
                                                            case -1:
                                                                v36 := GenS2{F0: false, F1: (fi17 / ((-104) | 1))}
                                                                _ = v36
                                                            case 4294967295:
                                                                var v37 int64 = 42
                                                                _ = v37
                                                            }
                                                            var v38 string = ("\n\t" + ("hello" + ("\n\t" + "")))
                                                            _ = v34
                                                            _ = v38
                                                        }
                                                    case 0:
                                                        switch 127 {
                                                        case 2147483647:
                                                            switch ("" + "") {
                                                            case "a":
                                                                fmt.Println(fi11, v2, v6)
                                                            }
                                                        case -1:
                                                            var v39 int64 = -9223372036854775808
                                                            _ = v39
                                                        default:
                                                            fmt.Println(fi17, v4, fi11)
                                                            v7 = (-(fi17))
                                                        }
                                                    case -1:
                                                        var v40 bool = (!true)
                                                        _ = v40
                                                    }
                                                    v41 := GenS4{F0: 65535, F1: "a"}
                                                    _ = v41
                                                }
                                                v42 := GenS5{F0: 1, F1: 712}
                                                if true {
                                                    if (true || false) {
                                                        v43 := GenS3{F0: v4, F1: (v4 & 0), F2: int64(v2)}
                                                        _ = v43
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = (^(fi17))
                                                        }
                                                        v2 = 42
                                                        var v44 uint64 = v4
                                                        _ = v44
                                                    }
                                                }
                                                fmt.Println(v18, v4, v6)
                                                _ = v42
                                            }
                                            v7 += (fi11 - (^(fi11)))
                                            if true {
                                                v2 -= 42
                                            } else {
                                                v7 = v7
                                            }
                                        } else {
                                            v45 := GenS4{F0: 1, F1: ("\n\t" + ("foo bar" + ("hello" + "")))}
                                            _ = v45
                                        }
                                    }
                                    var v46 int64 = 9223372036854775807
                                    _ = len(v1)
                                    _ = v46
                                }
                                var v47 float64 = -2.25
                                if (-821 >= -32768) {
                                    fmt.Println(fi11)
                                    switch 127 {
                                    case -265:
                                        var v48 float32 = 0.1
                                        _ = v48
                                    case 414:
                                        switch 128 {
                                        case 127:
                                            switch "a" {
                                            case "":
                                                var v49 int = fi17
                                                _ = v49
                                            }
                                        }
                                    default:
                                        var v50 int64 = 0
                                        v51 := GenS1{F0: v4, F1: 1, F2: v2}
                                        _ = v50
                                        _ = v51
                                    }
                                }
                                _ = v47
                            }
                        }
                        _ = v18
                    }
                    v4 -= (v4 - 986)
                    _ = fi17
                }
                var v52 int = (v7 >> ((255) & 63))
                _ = v52
            }
            var v53 int = 0
            _ = fi11
            _ = v53
        }
    }
    for fi54 := 0; fi54 < 5; fi54++ {
        if len(v3) > 0 {
            v3[0] = v7
        }
        _ = fi54
    }
    if true {
        fmt.Println(v4, v4)
        v55 := GenS6{F0: 774}
        v2 -= 86
        _ = v55
    }
    var v56 uint = v6
    if len(v1) > 0 {
        v1[0] = 1
    }
    v4 = v4
    switch ("\n\t" + "") {
    case "":
        if false {
            switch (-(v7)) {
            case 185:
                v4 -= (v4 | (v4 / (((v4 + (v4 << ((v56) & 63)))) | 1)))
            }
        }
    case "a":
        fmt.Println(v7)
    }
    fmt.Println(v6)
    v3 = append(v3, (^(v7)))
    v7 = (v7 + (v7 >> ((573) & 63)))
    if true {
        v5 = append(v5, "a")
        v57 := GenS7{F0: -15, F1: 1, F2: 32767}
        for fi58 := 0; fi58 < 11; fi58++ {
            if false {
                switch (^(fi58)) {
                case 775:
                    v56 = 4294967295
                case -9223372036854775808:
                    v59 := GenS5{F0: 42, F1: 1}
                    _ = v59
                default:
                    switch ((("foo bar" + "foo bar") + "hello") + "\n\t") {
                    case "hello":
                        _ = len(v3)
                    case "\n\t":
                        v6 -= 688
                    }
                }
                var v60 int64 = 9223372036854775807
                _ = v60
            } else {
                for fi61 := 0; fi61 < 20; fi61++ {
                    var v62 int32 = -799
                    _ = fi61
                    _ = v62
                }
                fmt.Println(v7)
            }
            fmt.Println(fi58)
            v4 += 4294967295
            _ = fi58
        }
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v7
    _ = v8
}
