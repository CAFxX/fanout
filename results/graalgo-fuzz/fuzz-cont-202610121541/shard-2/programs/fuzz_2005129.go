package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 uint16; F2 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 int16 = -1
    var v4 uint8 = 42
    var v5 int16 = (v3 & (v3 + 1))
    v6 := []int{427, -950}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 2}
        var v9 float64 = (0.0 - 3.14159)
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: 411}
            v5 = v5
            switch (fi11 / (((fi11 << ((935) & 63))) | 1)) {
            case 888:
                if true {
                    v4 = (v4 - (v4 % ((110) | 1)))
                    v3 = (v5 ^ 1000)
                    var v13 uint32 = 65535
                    _ = v13
                } else {
                    var v14 int8 = -100
                    switch (fi11 >> ((439) & 63)) {
                    case 32767:
                        v4 -= (^(v4))
                    case 127:
                        switch fi11 {
                        case 42:
                            var v15 int16 = 42
                            _ = v15
                        case -639:
                            var v16 int64 = 1
                            _ = v16
                        case 32767:
                            if (!(42 < 4294967295)) {
                                v14 += (v14 * v14)
                                if true {
                                    if false {
                                        var v17 bool = (2147483648 > 4294967295)
                                        if v17 {
                                            var v18 uint8 = (v4 % ((0) | 1))
                                            if v17 {
                                                v4 = (^(v4))
                                                v19 := GenS2{F0: (-(v14))}
                                                v20 := v1[(fi10 | fi10)]
                                                v17 = v17
                                                _ = v19
                                                _ = v20
                                            } else {
                                                v4 = v4
                                                var v21 string = (("foo bar" + "\n\t") + "foo bar")
                                                _ = v21
                                            }
                                            _ = v18
                                        } else {
                                            switch fi10 {
                                            case 9223372036854775807:
                                                if v17 {
                                                    v3 -= (-(v5))
                                                    v22 := GenS3{F0: 255, F1: -245, F2: (v4 - v4)}
                                                    _ = v22
                                                }
                                            case 614:
                                                switch "a" {
                                                case "a":
                                                    if v17 {
                                                        fmt.Println(fi11, v3, v14)
                                                        if (924 > 1) {
                                                            v3 = (v5 | v3)
                                                            fmt.Println(fi10, v3)
                                                        }
                                                        var v23 int32 = 2147483647
                                                        _ = v23
                                                    }
                                                case "foo bar":
                                                    v3 = v3
                                                }
                                            case 32768:
                                                var v24 uint32 = 65535
                                                _ = v24
                                            }
                                            if ((42 >= 1) && v17) {
                                                switch ((((("\n\t" + "a") + ("a" + ("a" + ""))) + (("hello" + "\n\t") + "\n\t")) + "a") + ("\n\t" + ("hello" + "foo bar"))) {
                                                case "":
                                                    var v25 float64 = float64(1.5)
                                                    _ = v25
                                                case "\n\t":
                                                    v26 := GenS3{F0: (^(v4)), F1: -2147483648, F2: v4}
                                                    _ = v26
                                                default:
                                                    _ = len(v6)
                                                }
                                                v17 = v17
                                                switch "a" {
                                                case "foo bar":
                                                    v5 = -32768
                                                case "\n\t":
                                                    if false {
                                                        fmt.Println(v4, v14)
                                                        if true {
                                                            switch fi10 {
                                                            case -2147483648:
                                                                var v27 int64 = 9223372036854775807
                                                                _ = v27
                                                            case 9223372036854775807:
                                                                v28 := GenS3{F0: (v4 | v4), F1: 781, F2: v4}
                                                                _ = v28
                                                            default:
                                                                v17 = (v17 && false)
                                                                var v29 int32 = 2147483647
                                                                _ = v29
                                                            }
                                                        } else {
                                                            if v17 {
                                                                fmt.Println(v3, v3)
                                                                switch ((("foo bar" + "foo bar") + ("\n\t" + ("foo bar" + ("foo bar" + ("\n\t" + "foo bar"))))) + "") {
                                                                case "hello":
                                                                    var v30 uint32 = 2
                                                                    _ = v30
                                                                default:
                                                                    v31 := v2["hello"]
                                                                    _ = v31
                                                                }
                                                            } else {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                v17 = v17
                                                            }
                                                            v5 += v5
                                                        }
                                                    }
                                                case "hello":
                                                    var v32 int64 = -1
                                                    _ = v32
                                                }
                                            } else {
                                                if false {
                                                    v3 = (v3 & v3)
                                                }
                                                v33 := GenS2{F0: v14}
                                                delete(v2, ((((("foo bar" + "") + "foo bar") + "foo bar") + ("a" + ("a" + "hello"))) + ""))
                                                _ = v33
                                            }
                                            v14 -= v14
                                            v4 = v4
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v34 int16 = 878
                                                _ = v34
                                            }
                                        }
                                        v17 = (4294967295 == 4294967295)
                                        _ = v17
                                    } else {
                                        v4 = (v4 % (((v4 | v4)) | 1))
                                        v14 += v14
                                        if true {
                                            if false {
                                                v2["hello"] = "hello"
                                            }
                                            var v35 int64 = 42
                                            v14 = v14
                                            _ = v35
                                        }
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        v6 = append(v6, fi10)
                                    case "a":
                                        v5 += (v3 ^ v3)
                                    case "hello":
                                        switch fi11 {
                                        case 949:
                                            var v36 uint = 646
                                            _ = v36
                                        }
                                    }
                                }
                                fmt.Println(v3)
                            } else {
                                var v37 int = fi10
                                _ = v37
                            }
                        }
                    case -2147483648:
                        v4 -= (v4 ^ (v4 + 36))
                    }
                    _ = v14
                }
            case 4294967295:
                v4 = (^(v4))
            default:
                v4 += 1
            }
            _ = fi11
            _ = v12
        }
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 float64 = 3.14159
            var v40 uint = 42
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi10
    }
    v41 := GenS4{F0: 477}
    v42 := GenS1{F0: 0}
    v3 = v5
    var v43 uint8 = v4
    fmt.Println(v5)
    _ = len(v2)
    switch 42 {
    case -445:
        fmt.Println(v43, v43)
    case 9223372036854775807:
        var v44 uint32 = 2147483648
        _ = v44
    default:
        if (true || (0 < 0)) {
            delete(v1, 126)
            v43 = (v43 - (v4 ^ v4))
            v4 -= (v4 * 141)
        } else {
            v45 := GenS5{F0: (v5 * (v3 - v5)), F1: 2, F2: ("hello" + (("" + "hello") + "a"))}
            var v46 int = 32768
            _ = v45
            _ = v46
        }
    }
    v2[""] = "foo bar"
    _ = len(v2)
    for fi47 := 0; fi47 < 13; fi47++ {
        if len(v6) > 0 {
            v6[0] = fi47
        }
        _ = fi47
    }
    var v48 uint8 = (^(v43))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v48
    _ = v5
    _ = v6
}
