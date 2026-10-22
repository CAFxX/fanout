package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 string; F1 bool; F2 uint }
type GenS5 struct { F0 int16; F1 uint; F2 int64 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8 }

func main() {
    var v1 bool = true
    var v2 bool = ((128 >= 127) || (1 <= 21))
    v3 := []string{""}
    var v4 int64 = 1
    v5 := make(map[string]int)
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 bool = false
        if v2 {
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(v1)
                _ = fi9
            }
            switch fi7 {
            case 32767:
                v2 = v2
            case 1:
                v10 := GenS1{F0: "a", F1: 229, F2: 255}
                _ = v10
            case 128:
                fmt.Println(v1, fi7, v1)
            }
            if v8 {
                v8 = (v4 > (v4 ^ 42))
                for fi11 := 0; fi11 < 6; fi11++ {
                    switch -9223372036854775808 {
                    case 2147483647:
                        switch fi7 {
                        case 42:
                            if (0 > 244) {
                                v12 := GenS1{F0: "hello", F1: 18446744073709551615, F2: 65535}
                                switch fi7 {
                                case 32767:
                                    v4 += v4
                                case -796:
                                    v5["hello"] = -403
                                default:
                                    var v13 uint32 = 821
                                    _ = v13
                                }
                                _ = v12
                            } else {
                                if v2 {
                                    var v14 float32 = 1.5
                                    _ = v14
                                }
                                var v15 float32 = 0.1
                                delete(v5, "\n\t")
                                _ = v15
                            }
                            if v8 {
                                fmt.Println(v1, fi7)
                            }
                        }
                    }
                    _ = fi11
                }
            }
        }
        _ = fi7
        _ = v8
    }
    var v16 int32 = 0
    v17 := GenS1{F0: "foo bar", F1: 1, F2: 65535}
    var v18 int8 = 42
    if (42 > 113) {
        if (((1.5 * 0.1) >= (3.14159 - float32(0.0))) && v1) {
            for fi19 := 0; fi19 < 10; fi19++ {
                v3 = append(v3, "a")
                v16 = -604
                _ = fi19
            }
            fmt.Println(v16, v18, v16)
        }
        fmt.Println(v18)
        var v20 int16 = 317
        _ = v20
    } else {
        var v21 uint64 = 639
        v16 = (v16 | v16)
        _ = v21
    }
    v2 = true
    v22 := GenS2{F0: 764}
    switch ("a" + ("\n\t" + (("a" + "foo bar") + ""))) {
    case "\n\t":
        if (1 != 921) {
            fmt.Println(v16, v2, v18)
            v4 += 1
            switch 127 {
            case -2147483648:
                switch ("" + ("\n\t" + "foo bar")) {
                case "\n\t":
                    if (v2 || (!(127 == 9223372036854775807))) {
                        v23 := GenS3{F0: v4, F1: -206}
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v18 = -121
                            v18 = v18
                            v25 := GenS4{F0: "a", F1: (v4 == (v4 << ((4294967295) & 63))), F2: 1}
                            _ = fi24
                            _ = v25
                        }
                        v26 := GenS5{F0: -32768, F1: 740, F2: (^(v4))}
                        _ = v23
                        _ = v26
                    }
                case "a":
                    var v27 string = ""
                    _ = v27
                default:
                    v2 = v1
                }
            case 32768:
                fmt.Println(v18)
            case -9223372036854775808:
                switch 649 {
                case 878:
                    switch "hello" {
                    case "\n\t":
                        var v28 uint = 105
                        _ = v28
                    case "foo bar":
                        switch 2147483647 {
                        case 4294967295:
                            if v1 {
                                fmt.Println(v4, v1)
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    switch ((("a" + "hello") + (("a" + "foo bar") + ("a" + ("" + ("\n\t" + (("foo bar" + "hello") + ("" + "foo bar"))))))) + "") {
                                    case "hello":
                                        for fi30 := 0; fi30 < 15; fi30++ {
                                            if false {
                                                v31 := v5["a"]
                                                v32 := GenS6{F0: 369, F1: v16}
                                                if len(v3) > 0 {
                                                    v3[0] = "\n\t"
                                                }
                                                fmt.Println(fi30, fi30, v1)
                                                v31 -= fi30
                                                v1 = true
                                                switch "hello" {
                                                case "\n\t":
                                                    var v33 string = ""
                                                    _ = v33
                                                default:
                                                    v2 = ((-1 > (fi30 | fi29)) && ((-14 >= 1000) && false))
                                                }
                                                var v34 float32 = (((3.14159 / (1.5 + ((1.5 / (0.1 + ((3.14159 - 1e10) - 0.0))) - 3.14159))) + ((1e10 * 1e10) * float32(3.14159))) + 0.1)
                                                switch (^(v31)) {
                                                case 127:
                                                    var v35 int8 = 1
                                                    _ = v35
                                                case -9223372036854775808:
                                                    v3 = append(v3, "\n\t")
                                                }
                                                v18 += v18
                                                _ = v31
                                                _ = v32
                                                _ = v34
                                            } else {
                                                switch (fi30 << ((220) & 63)) {
                                                case 1:
                                                    v36 := GenS1{F0: ("hello" + "\n\t"), F1: 358, F2: 32767}
                                                    _ = v36
                                                case -9223372036854775808:
                                                    v16 = v16
                                                default:
                                                    fmt.Println(fi30)
                                                }
                                                if (255 == 255) {
                                                    v18 = 42
                                                    if (!v2) {
                                                        v2 = (260 != 414)
                                                    } else {
                                                        fmt.Println(v18, v1)
                                                        if ((float32((3.14159 * -2.25)) * (float32(float64(0.0)) * (0.1 - 0.0))) == (0.1 / 0.1)) {
                                                            v37 := GenS7{F0: 0, F1: (v4 + 42), F2: 1}
                                                            var v38 float32 = 0.1
                                                            _ = v37
                                                            _ = v38
                                                        } else {
                                                            var v39 uint = 4294967295
                                                            v40 := GenS6{F0: -1, F1: (v16 & (^(v16)))}
                                                            v41 := GenS8{F0: 59}
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        var v42 int = (fi30 * fi30)
                                                        var v43 int16 = int16(v18)
                                                        if (!false) {
                                                            var v44 float64 = 1.5
                                                            _ = v44
                                                        }
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    var v45 uint8 = 105
                                                    _ = v45
                                                } else {
                                                    var v46 int32 = 1
                                                    _ = v46
                                                }
                                                fmt.Println(v16)
                                                var v47 int32 = (v16 << ((1) & 63))
                                                var v48 float64 = 0.0
                                                if (v1 && v1) {
                                                    var v49 float64 = v48
                                                    _ = v49
                                                } else {
                                                    v50 := GenS3{F0: v4, F1: int(v47)}
                                                    _ = v50
                                                }
                                                v47 += v16
                                                _ = v47
                                                _ = v48
                                            }
                                            if (((-1 != (v18 + v18)) && true) || (true || (403 <= fi29))) {
                                                v3 = append(v3, "a")
                                                var v51 uint32 = 0
                                                _ = v51
                                            } else {
                                                v2 = (601 == 255)
                                            }
                                            var v52 int64 = v4
                                            _ = fi30
                                            _ = v52
                                        }
                                    case "a":
                                        _ = len(v5)
                                    }
                                    _ = fi29
                                }
                            }
                        case 627:
                            for fi53 := 0; fi53 < 20; fi53++ {
                                if v1 {
                                    var v54 int8 = 100
                                    _ = v54
                                } else {
                                    _ = len(v3)
                                    for fi55 := 0; fi55 < 1; fi55++ {
                                        v3 = append(v3, "\n\t")
                                        fmt.Println(v2, v18, v18)
                                        _ = fi55
                                    }
                                }
                                if (!(true && true)) {
                                    var v56 uint32 = 65535
                                    for fi57 := 0; fi57 < 8; fi57++ {
                                        fmt.Println(v56, fi57)
                                        var v58 uint8 = 127
                                        if true {
                                            fmt.Println(v18, v56)
                                            v59 := GenS8{F0: 1}
                                            v60 := GenS8{F0: -288}
                                            _ = v59
                                            _ = v60
                                        }
                                        _ = fi57
                                        _ = v58
                                    }
                                    v61 := v6[4294967295]
                                    _ = v56
                                    _ = v61
                                } else {
                                    v16 = (^(v16))
                                    fmt.Println(v16, v16, v4)
                                }
                                for fi62 := 0; fi62 < 15; fi62++ {
                                    switch ("hello" + ("a" + "\n\t")) {
                                    case "foo bar":
                                        v16 = v16
                                    case "":
                                        v63 := GenS1{F0: "a", F1: 148, F2: 42}
                                        _ = v63
                                    case "hello":
                                        var v64 float64 = float64(((0.0 + 1.5) + 0.0))
                                        v3 = append(v3, "\n\t")
                                        v16 -= -218
                                        _ = v64
                                    }
                                    v16 -= -550
                                    _ = fi62
                                }
                                _ = fi53
                            }
                        case -1:
                            var v65 uint16 = 255
                            _ = v65
                        }
                    case "a":
                        v5["hello"] = 127
                    default:
                        var v66 uint = 221
                        _ = v66
                    }
                case 0:
                    var v67 int8 = -1
                    _ = v67
                case -9223372036854775808:
                    for fi68 := 0; fi68 < 20; fi68++ {
                        v16 += v16
                        _ = len(v3)
                        v4 += v4
                        _ = fi68
                    }
                }
            default:
                for fi69 := 0; fi69 < 9; fi69++ {
                    _ = len(v3)
                    _ = fi69
                }
            }
        }
    case "a":
        if true {
            var v70 uint = 2
            for fi71 := 0; fi71 < 9; fi71++ {
                v18 += (v18 | v18)
                _ = fi71
            }
            v72 := GenS9{F0: 606}
            _ = v70
            _ = v72
        }
    default:
        switch -9223372036854775808 {
        case -1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 128:
            var v73 uint64 = 18446744073709551615
            _ = v73
        }
    }
    var v74 uint64 = 4294967295
    _ = len(v5)
    if (1 < 1000) {
        if v2 {
            var v75 uint64 = (v74 & (v74 - (^(v74))))
            _ = v75
        } else {
            var v76 uint32 = 469
            v1 = v1
            var v77 uint = 2
            _ = v76
            _ = v77
        }
        v2 = false
        var v78 float32 = 1e10
        _ = v78
    }
    for fi79 := 0; fi79 < 15; fi79++ {
        fmt.Println(v16, v4)
        v80 := GenS10{F0: 0}
        _ = fi79
        _ = v80
    }
    v3 = append(v3, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
