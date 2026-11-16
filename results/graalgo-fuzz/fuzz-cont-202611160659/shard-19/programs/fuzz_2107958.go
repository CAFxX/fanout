package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }
type GenS3 struct { F0 uint64; F1 bool; F2 uint16 }
func genFn0() string {
    var v4 uint32 = 0
    var v5 float32 = 1e10
    var v6 float64 = 1.5
    fmt.Println(v6, v6)
    var v7 int8 = 101
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return ("\n\t" + "foo bar")
}
func genFn1() bool {
    if (2 != 65535) {
        var v8 int = 127
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := []int{(v8 << ((714) & 63)), -1, fi9}
            _ = fi9
            _ = v10
        }
        var v11 float64 = 3.14159
        _ = v11
        _ = v8
    } else {
        v12 := GenS1{F0: 4294967295}
        var v13 uint = 46
        var v14 uint = (^(v13))
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        switch fi15 {
        case 42:
            if true {
                var v16 string = ("foo bar" + "\n\t")
                v17 := []int{-2147483648, (fi15 - 32768), (fi15 | 615)}
                for fi18 := 0; fi18 < 17; fi18++ {
                    if false {
                        v16 = v16
                        if true {
                            if (!true) {
                                if ((!(true && (true || true))) || (true || true)) {
                                    _ = len(v17)
                                } else {
                                    switch ("\n\t" + (("foo bar" + (v16 + "")) + v16)) {
                                    case "foo bar":
                                        v16 = ((("a" + v16) + v16) + v16)
                                    default:
                                        var v19 float32 = float32(3.14159)
                                        _ = v19
                                    }
                                    switch v16 {
                                    case "hello":
                                        v20 := GenS1{F0: 4294967295}
                                        _ = v20
                                    case "\n\t":
                                        if true {
                                            var v21 bool = false
                                            v17 = append(v17, fi18)
                                            _ = len(v17)
                                            if false {
                                                var v22 float64 = float64(float32(1e10))
                                                v16 = ""
                                                _ = v22
                                            }
                                            v21 = v21
                                            _ = v21
                                        } else {
                                            v16 = (v16 + "a")
                                            if len(v17) > 0 {
                                                v17[0] = fi18
                                            }
                                        }
                                    }
                                }
                                var v23 float32 = 0.1
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                _ = v23
                            }
                            switch v16 {
                            case "\n\t":
                                var v24 uint64 = 4294967295
                                _ = v24
                            default:
                                var v25 float64 = (0.1 / ((1e10 + 1.5) - 0.0))
                                _ = v25
                            }
                        } else {
                            if (-45 >= -81) {
                                if false {
                                    var v26 float32 = -2.25
                                    _ = v26
                                } else {
                                    switch (fi18 / (((fi18 << ((270) & 63))) | 1)) {
                                    case -242:
                                        if true {
                                            v27 := GenS2{F0: 710, F1: 223, F2: (2 < 4294967295)}
                                            _ = len(v17)
                                            v17 = append(v17, (fi15 | fi18))
                                            _ = v27
                                        } else {
                                            var v28 uint32 = 65535
                                            var v29 int = 140
                                            fmt.Println(fi18, fi18)
                                            _ = v28
                                            _ = v29
                                        }
                                    case 626:
                                        if true {
                                            v16 = "foo bar"
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                        }
                                    default:
                                        var v30 bool = false
                                        _ = v30
                                    }
                                    var v31 float64 = (((1e10 * 3.14159) * 1e10) + 1.5)
                                    _ = len(v17)
                                    var v32 int64 = -1
                                    v17 = append(v17, fi15)
                                    _ = v31
                                    _ = v32
                                }
                                v16 = ""
                                v16 = "foo bar"
                                var v33 string = (v16 + v16)
                                var v34 int16 = 32767
                                _ = v33
                                _ = v34
                            } else {
                                var v35 uint32 = 2147483648
                                _ = v35
                            }
                            if false {
                                v16 = (v16 + v16)
                                var v36 uint64 = 42
                                _ = v36
                            } else {
                                if false {
                                    if (4294967295 == 1) {
                                        var v37 uint32 = 2147483648
                                        if false {
                                            switch "hello" {
                                            case "foo bar":
                                                switch 128 {
                                                case -2147483648:
                                                    v37 -= 2
                                                case -956:
                                                    v37 -= (v37 & (v37 * v37))
                                                case 4294967295:
                                                    if true {
                                                        var v38 int = fi15
                                                        fmt.Println(v38)
                                                        if len(v17) >= 2 {
                                                            v17 = v17[1:len(v17)]
                                                        }
                                                        _ = v38
                                                    }
                                                }
                                            case "a":
                                                switch fi18 {
                                                case 9223372036854775807:
                                                    var v39 uint16 = 193
                                                    _ = v39
                                                default:
                                                    v40 := GenS1{F0: 42}
                                                    _ = v40
                                                }
                                                v17 = append(v17, int(v37))
                                            default:
                                                var v41 bool = true
                                                _ = v41
                                            }
                                            var v42 int32 = 193
                                            v42 -= 2147483647
                                            _ = v42
                                        } else {
                                            v17 = append(v17, fi15)
                                            fmt.Println(v16, v37)
                                            if true {
                                                v37 += v37
                                                var v43 bool = true
                                                v16 = "hello"
                                                _ = v43
                                            }
                                            if true {
                                                v37 += v37
                                            }
                                            v17 = append(v17, fi15)
                                        }
                                        v16 = v16
                                        _ = v37
                                    } else {
                                        v17 = append(v17, -1)
                                        fmt.Println(v16, v16)
                                    }
                                }
                                var v44 int16 = 42
                                var v45 uint32 = 2
                                v46 := GenS3{F0: 0, F1: (!true), F2: 2}
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                        }
                    } else {
                        v47 := GenS1{F0: 1}
                        v17 = append(v17, (^(fi15)))
                        _ = v47
                    }
                    if false {
                        switch (fi15 | fi15) {
                        case 472:
                            if false {
                                fmt.Println(fi18)
                                var v48 int = fi18
                                fmt.Println(fi15, fi18, fi15)
                                _ = v48
                            }
                        case 660:
                            fmt.Println(fi15)
                        case 316:
                            var v49 bool = false
                            _ = v49
                        default:
                            var v50 uint16 = 1
                            _ = v50
                        }
                        fmt.Println(fi18)
                        fmt.Println(v16, fi18, fi18)
                    }
                    v16 = "a"
                    _ = fi18
                }
                _ = v16
                _ = v17
            }
        case 965:
            for fi51 := 0; fi51 < 8; fi51++ {
                if (!false) {
                    v52 := []int{(fi15 + fi15)}
                    fmt.Println(fi15)
                    _ = v52
                }
                var v53 float64 = 1e10
                fmt.Println(v53, fi15, v53)
                fmt.Println(fi51, fi15)
                _ = fi51
                _ = v53
            }
        case 128:
            var v54 float32 = 1.5
            _ = v54
        }
        for fi55 := 0; fi55 < 6; fi55++ {
            v56 := []int{(fi15 + (fi55 | fi15)), fi15, -134}
            _ = fi55
            _ = v56
        }
        _ = fi15
    }
    for fi57 := 0; fi57 < 10; fi57++ {
        var v58 uint8 = 1
        if true {
            v58 += 99
            switch ((("\n\t" + "a") + "foo bar") + "a") {
            case "hello":
                var v59 uint64 = 4294967295
                _ = v59
            case "foo bar":
                if false {
                    v58 -= 127
                    for fi60 := 0; fi60 < 20; fi60++ {
                        v61 := []int{(^(fi60)), 128, fi57}
                        v58 += 127
                        var v62 uint32 = 376
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                    var v63 int = fi57
                    _ = v63
                } else {
                    if (true && (false && ((!true) || true))) {
                        v58 = 2
                    }
                    fmt.Println(v58, v58)
                }
            }
        } else {
            if true {
                if true {
                    v58 += v58
                }
                v64 := []int{-2147483648, fi57, fi57}
                _ = v64
            }
            switch (fi57 % (((fi57 << ((65535) & 63))) | 1)) {
            case 615:
                for fi65 := 0; fi65 < 16; fi65++ {
                    v66 := []int{fi65, (fi65 << ((4294967295) & 63)), 267, 32767}
                    _ = len(v66)
                    switch int(v58) {
                    case -463:
                        var v67 uint8 = v58
                        _ = v67
                    case -579:
                        var v68 uint64 = 2
                        _ = v68
                    case 0:
                        fmt.Println(fi65, fi65)
                    default:
                        v69 := GenS3{F0: 18446744073709551615, F1: false, F2: 255}
                        if len(v66) >= 2 {
                            v66 = v66[1:len(v66)]
                        }
                        var v70 int8 = 1
                        v70 = 42
                        _ = v69
                        _ = v70
                    }
                    _ = fi65
                    _ = v66
                }
            }
            var v71 uint32 = 535
            _ = v71
        }
        v72 := GenS3{F0: 18446744073709551615, F1: (!true), F2: 885}
        _ = fi57
        _ = v58
        _ = v72
    }
    var v73 int = -392
    _ = v73
    return (true || true)
}

func main() {
    var v1 string = "a"
    v2 := []string{(("" + ((v1 + v1) + ("hello" + v1))) + v1), ("hello" + (v1 + v1)), (((("foo bar" + v1) + ((v1 + v1) + v1)) + (("" + v1) + "a")) + (("a" + "foo bar") + v1))}
    v3 := make(map[string]int)
    var v74 int32 = 706
    var v75 uint8 = 127
    v1 = v1
    v76 := GenS2{F0: 255, F1: v75, F2: true}
    var v77 int32 = (v74 ^ -1)
    var v78 uint32 = 151
    var v79 int16 = 32767
    var v80 uint8 = v75
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
