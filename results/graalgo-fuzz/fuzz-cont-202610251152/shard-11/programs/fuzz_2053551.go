package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []uint32{}
    v2 := []string{"foo bar", "", "foo bar"}
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    var v5 int8 = -103
    v2 = append(v2, "")
    fmt.Println(v5)
    switch "a" {
    case "hello":
        v6 := GenS1{F0: 1000}
        _ = v6
    case "":
        for fi7 := 0; fi7 < 2; fi7++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 string = (("" + "foo bar") + "a")
                _ = fi8
                _ = v9
            }
            var v10 int8 = (v5 | (v5 >> ((2) & 63)))
            _ = fi7
            _ = v10
        }
    }
    v5 += v5
    var v11 int64 = 42
    if false {
        if false {
            if true {
                fmt.Println(v5, v11)
                if ("foo bar" > ("\n\t" + "foo bar")) {
                    v12 := GenS1{F0: 0}
                    if (3.14159 < 3.14159) {
                        v5 = (-(v5))
                    }
                    v11 -= (^(v11))
                    _ = v12
                }
                var v13 uint32 = 4294967295
                _ = v13
            }
            switch -9223372036854775808 {
            case 9223372036854775807:
                switch -322 {
                case 83:
                    var v14 uint16 = 65535
                    _ = v14
                case -9223372036854775808:
                    fmt.Println(v11, v11)
                default:
                    v15 := v3[-9223372036854775808]
                    _ = v15
                }
            case 32767:
                for fi16 := 0; fi16 < 6; fi16++ {
                    switch fi16 {
                    case -1:
                        fmt.Println(fi16, fi16)
                    case 210:
                        if false {
                            for fi17 := 0; fi17 < 12; fi17++ {
                                switch fi16 {
                                case 128:
                                    _ = len(v4)
                                default:
                                    var v18 float32 = 1.5
                                    _ = v18
                                }
                                if (!true) {
                                    var v19 bool = ((-279 > 0) || (!(!(((^(v11)) == -1) && (78 <= (v5 + 100))))))
                                    _ = v19
                                } else {
                                    v5 = v5
                                    var v20 int16 = 625
                                    var v21 uint16 = 42
                                    fmt.Println(fi16, v5, v20)
                                    _ = v20
                                    _ = v21
                                }
                                v2 = append(v2, (("\n\t" + "\n\t") + "\n\t"))
                                _ = fi17
                            }
                        }
                    default:
                        var v22 int16 = 0
                        _ = v22
                    }
                    if true {
                        v23 := GenS1{F0: 32767}
                        if false {
                            v5 += v5
                            v11 = v11
                        }
                        v5 = v5
                        _ = v23
                    } else {
                        if ((!false) && (false && false)) {
                            var v24 uint32 = 1
                            var v25 bool = true
                            switch (("a" + ("foo bar" + ("\n\t" + ""))) + "hello") {
                            case "hello":
                                v26 := GenS1{F0: -375}
                                _ = v26
                            case "foo bar":
                                for fi27 := 0; fi27 < 18; fi27++ {
                                    var v28 bool = (v25 || (("a" != "hello") || v25))
                                    _ = fi27
                                    _ = v28
                                }
                            case "":
                                for fi29 := 0; fi29 < 8; fi29++ {
                                    if v25 {
                                        v24 = v24
                                        if ((v5 / (((v5 + v5)) | 1)) < v5) {
                                            switch 344 {
                                            case -2147483648:
                                                fmt.Println(fi29)
                                            case 826:
                                                var v30 uint = 65535
                                                _ = v30
                                            }
                                        }
                                    } else {
                                        _ = len(v2)
                                        var v31 int32 = 42
                                        if v25 {
                                            if v25 {
                                                if v25 {
                                                    var v32 float64 = 3.14159
                                                    var v33 int = (fi16 * -2147483648)
                                                    switch "" {
                                                    case "a":
                                                        if (((v25 || false) && true) && (int32(v5) <= v31)) {
                                                            fmt.Println(v32)
                                                            v33 += 1
                                                            fmt.Println(v31)
                                                        } else {
                                                            if true {
                                                                v2 = append(v2, "foo bar")
                                                                switch (^(v33)) {
                                                                case 32768:
                                                                    var v34 int32 = int32(v5)
                                                                    _ = v34
                                                                case 2147483647:
                                                                    var v35 int16 = 0
                                                                    _ = len(v2)
                                                                    v11 -= (v11 >> ((1) & 63))
                                                                    _ = v35
                                                                default:
                                                                    var v36 float64 = (float64(1.5) + (1.5 + v32))
                                                                    _ = v36
                                                                }
                                                                var v37 int16 = -32768
                                                                v25 = (!(((1.5 * v32) + 1.5) != 0.1))
                                                                _ = v37
                                                            }
                                                            var v38 string = ""
                                                            _ = v38
                                                        }
                                                    case "":
                                                        fmt.Println(v32, fi16)
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    v39 := GenS1{F0: -32768}
                                                    if ((("" + (("foo bar" + "") + (("" + "\n\t") + "hello"))) + ("a" + "hello")) > "hello") {
                                                        var v40 uint64 = 1
                                                        v25 = v25
                                                        _ = v40
                                                    }
                                                    v41 := GenS1{F0: 999}
                                                    _ = v39
                                                    _ = v41
                                                }
                                            }
                                            var v42 int = 884
                                            _ = v42
                                        } else {
                                            v5 -= (-(v5))
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = (v24 | v24)
                                        }
                                        switch fi16 {
                                        case 1:
                                            switch ((("\n\t" + ("hello" + "a")) + ("foo bar" + "a")) + "a") {
                                            case "foo bar":
                                                var v43 uint16 = 497
                                                _ = v43
                                            }
                                            v5 = 0
                                            v24 = 2147483648
                                        case 9223372036854775807:
                                            v31 = (v31 | -2147483648)
                                        default:
                                            var v44 float32 = 0.0
                                            _ = v44
                                        }
                                        var v45 float32 = (0.0 - (float32(3.14159) * 3.14159))
                                        _ = v31
                                        _ = v45
                                    }
                                    _ = fi29
                                }
                            }
                            _ = v24
                            _ = v25
                        }
                        fmt.Println(fi16)
                    }
                    switch fi16 {
                    case -2147483648:
                        v46 := GenS2{F0: 285}
                        _ = v46
                    case 128:
                        switch (fi16 * (^(fi16))) {
                        case 42:
                            v5 = v5
                        case -1:
                            for fi47 := 0; fi47 < 16; fi47++ {
                                v1 = append(v1, 894)
                                var v48 string = ("a" + "hello")
                                v11 = (v11 * 1)
                                _ = fi47
                                _ = v48
                            }
                        case 2147483647:
                            _ = len(v2)
                        default:
                            if true {
                                switch ("a" + "foo bar") {
                                case "a":
                                    for fi49 := 0; fi49 < 5; fi49++ {
                                        fmt.Println(v5, fi16)
                                        _ = fi49
                                    }
                                case "hello":
                                    delete(v4, fi16)
                                default:
                                    for fi50 := 0; fi50 < 1; fi50++ {
                                        var v51 string = ("a" + "")
                                        var v52 uint32 = 1
                                        v11 = (v11 | (v11 * (v11 * 0)))
                                        _ = fi50
                                        _ = v51
                                        _ = v52
                                    }
                                }
                            }
                        }
                    }
                    _ = fi16
                }
            case 0:
                v5 += -126
            }
            var v53 int16 = 1
            _ = v53
        }
        var v54 string = ""
        fmt.Println(v11, v5)
        _ = v54
    }
    v3[-1] = ("a" + "")
    for fi55 := 0; fi55 < 5; fi55++ {
        var v56 int32 = 636
        _ = fi55
        _ = v56
    }
    v2 = append(v2, "\n\t")
    for fi57 := 0; fi57 < 16; fi57++ {
        fmt.Println(fi57, v5)
        fmt.Println(fi57, v5)
        _ = fi57
    }
    v58 := v4[9223372036854775807]
    v5 = (v5 ^ (v5 >> ((2) & 63)))
    v59 := GenS2{F0: 0}
    v5 -= int8(v11)
    v11 += 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
}
