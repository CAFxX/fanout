package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint16; F2 bool }

func main() {
    var v1 float64 = 1.5
    var v2 uint32 = 65535
    v3 := make(map[string]int)
    var v4 uint = 42
    var v5 int16 = 0
    var v6 bool = true
    switch "hello" {
    case "a":
        var v7 int16 = -198
        _ = v7
    }
    if v6 {
        if (0.0 == v1) {
            v8 := GenS1{F0: 1, F1: -40}
            switch 128 {
            case 1:
                if v6 {
                    var v9 int8 = -1
                    switch -997 {
                    case 319:
                        v10 := GenS2{F0: (v2 << ((v4) & 63))}
                        _ = v10
                    default:
                        var v11 string = ("foo bar" + "foo bar")
                        _ = v11
                    }
                    switch 4294967295 {
                    case 42:
                        fmt.Println(v6, v1, v4)
                    case 0:
                        v12 := GenS1{F0: -196, F1: v9}
                        _ = v12
                    case 32767:
                        var v13 float32 = 0.1
                        _ = v13
                    default:
                        if v6 {
                            v14 := GenS3{F0: -9223372036854775808}
                            var v15 float32 = -2.25
                            var v16 int32 = -887
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            v17 := v3[("\n\t" + "hello")]
                            var v18 int64 = -1
                            var v19 string = "a"
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    }
                    _ = v9
                }
            case 127:
                fmt.Println(v6, v5, v1)
            }
            _ = v8
        } else {
            if (true && v6) {
                fmt.Println(v4)
                switch "hello" {
                case "foo bar":
                    delete(v3, "a")
                case "\n\t":
                    fmt.Println(v6, v2, v4)
                case "a":
                    _ = len(v3)
                default:
                    if false {
                        v20 := GenS4{F0: 2147483647, F1: v2, F2: v2}
                        v5 -= v5
                        _ = v20
                    }
                }
                for fi21 := 0; fi21 < 15; fi21++ {
                    v22 := GenS3{F0: -329}
                    v23 := GenS4{F0: fi21, F1: (v2 & 557), F2: (v2 & v2)}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            } else {
                if ((((v5 * 1000) != -645) || (!((!(!(false && (-128 > 42)))) && (v6 && (100 == -100))))) && v6) {
                    var v24 uint64 = 2
                    v2 += 4294967295
                    var v25 string = ("\n\t" + ("" + (("\n\t" + ((("" + "\n\t") + ("\n\t" + "\n\t")) + "\n\t")) + "hello")))
                    _ = v24
                    _ = v25
                } else {
                    var v26 string = ""
                    var v27 float64 = ((1e10 / v1) / v1)
                    var v28 int8 = -128
                    _ = v26
                    _ = v27
                    _ = v28
                }
                if true {
                    v5 = 1
                    switch ("hello" + "a") {
                    case "foo bar":
                        for fi29 := 0; fi29 < 4; fi29++ {
                            switch fi29 {
                            case -196:
                                var v30 int32 = 0
                                _ = v30
                            case 128:
                                var v31 int = (^(fi29))
                                _ = v31
                            }
                            v32 := GenS2{F0: (v2 % ((134) | 1))}
                            _ = fi29
                            _ = v32
                        }
                    case "\n\t":
                        var v33 int = 127
                        _ = v33
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            switch "foo bar" {
                            case "hello":
                                v6 = v6
                            default:
                                if (-128 > 42) {
                                    fmt.Println(v6, v6)
                                    v1 = v1
                                    for fi34 := 0; fi34 < 7; fi34++ {
                                        var v35 float64 = 0.1
                                        _ = fi34
                                        _ = v35
                                    }
                                } else {
                                    var v36 uint64 = 4294967295
                                    for fi37 := 0; fi37 < 14; fi37++ {
                                        v1 += v1
                                        v4 += 502
                                        _ = fi37
                                    }
                                    _ = v36
                                }
                            }
                        default:
                            var v38 uint8 = 0
                            _ = v38
                        }
                    }
                } else {
                    v39 := GenS2{F0: (v2 * 4294967295)}
                    _ = v39
                }
            }
            var v40 int32 = 696
            if v6 {
                v41 := v3["\n\t"]
                if ((!(32767 == v5)) && v6) {
                    var v42 uint = 0
                    v43 := GenS3{F0: 669}
                    v4 += (v42 / (((v42 * v42)) | 1))
                    _ = v42
                    _ = v43
                } else {
                    var v44 int8 = 42
                    _ = v44
                }
                _ = v41
            }
            _ = v40
        }
        v45 := v3["hello"]
        v46 := GenS5{F0: (v5 - int16(v2))}
        _ = v45
        _ = v46
    }
    fmt.Println(v1)
    v47 := []int{4294967295}
    _ = len(v47)
    var v48 uint8 = 255
    var v49 int32 = 0
    if v6 {
        fmt.Println(v2)
        if ((false && (!false)) || true) {
            if v6 {
                fmt.Println(v1)
                v1 -= 0.0
            } else {
                v49 -= 693
                var v50 uint64 = 736
                var v51 int32 = 0
                _ = v50
                _ = v51
            }
            for fi52 := 0; fi52 < 6; fi52++ {
                switch "a" {
                case "foo bar":
                    if (!(42 > -128)) {
                        switch (fi52 & fi52) {
                        case 383:
                            switch "\n\t" {
                            case "hello":
                                switch 0 {
                                case 961:
                                    if (-2.25 <= (((3.14159 / (1.5 / 1.0)) * 0.0) + -2.25)) {
                                        var v53 int32 = (v49 % ((0) | 1))
                                        _ = v53
                                    }
                                }
                            }
                        }
                    } else {
                        v54 := v3["a"]
                        _ = v54
                    }
                case "":
                    var v55 uint32 = (v2 % ((v2) | 1))
                    _ = v55
                case "a":
                    v1 -= v1
                }
                switch fi52 {
                case 32767:
                    delete(v3, "\n\t")
                default:
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        for fi56 := 0; fi56 < 6; fi56++ {
                            fmt.Println(v48)
                            if true {
                                switch fi56 {
                                case 1:
                                    if v6 {
                                        if v6 {
                                            var v57 uint8 = v48
                                            _ = v57
                                        }
                                    }
                                case 32767:
                                    v47 = append(v47, (fi52 >> ((42) & 63)))
                                case -470:
                                    var v58 int64 = 9223372036854775807
                                    _ = v58
                                default:
                                    if false {
                                        var v59 int32 = -1
                                        _ = v59
                                    } else {
                                        if (!v6) {
                                            switch (("foo bar" + ((("\n\t" + ("" + "\n\t")) + "\n\t") + "a")) + "a") {
                                            case "foo bar":
                                                var v60 string = "foo bar"
                                                _ = v60
                                            case "\n\t":
                                                fmt.Println(v6, v48)
                                            case "a":
                                                var v61 int64 = 42
                                                _ = v61
                                            default:
                                                if (!((^(v48)) < v48)) {
                                                    var v62 uint16 = 326
                                                    switch "a" {
                                                    case "a":
                                                        fmt.Println(fi56)
                                                    }
                                                    _ = v62
                                                }
                                            }
                                            var v63 uint8 = (v48 << (((v4 & v4)) & 63))
                                            if true {
                                                v64 := GenS6{F0: (fi52 - (fi56 >> ((v4) & 63))), F1: 255, F2: (true && v6)}
                                                var v65 bool = ((!(0 != 42)) && (!(1 < 255)))
                                                var v66 int = fi56
                                                v6 = v65
                                                v3[("foo bar" + "hello")] = (^(fi52))
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            } else {
                                                var v67 bool = v6
                                                _ = v67
                                            }
                                            _ = v63
                                        }
                                        v49 += v49
                                        var v68 float64 = 1.5
                                        _ = v68
                                    }
                                }
                                v49 = 42
                                v69 := GenS2{F0: (v2 & v2)}
                                _ = v69
                            }
                            _ = fi56
                        }
                    default:
                        if v6 {
                            v49 = -2147483648
                        } else {
                            fmt.Println(v1)
                            for fi70 := 0; fi70 < 16; fi70++ {
                                v2 = (v2 + 0)
                                _ = fi70
                            }
                        }
                    }
                }
                if v6 {
                    var v71 bool = (false && true)
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v48, v1)
                    case "\n\t":
                        switch fi52 {
                        case -818:
                            v72 := GenS2{F0: 65535}
                            _ = v72
                        case 2147483647:
                            var v73 int64 = -882
                            _ = v73
                        default:
                            var v74 int32 = v49
                            _ = v74
                        }
                    default:
                        _ = len(v3)
                    }
                    var v75 uint8 = 2
                    _ = v71
                    _ = v75
                } else {
                    v5 += v5
                }
                _ = fi52
            }
        } else {
            fmt.Println(v2)
        }
    }
    var v76 uint = (v4 << ((v4) & 63))
    var v77 int8 = 42
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v78 string = ("\n\t" + "a")
        _ = v78
    case -639:
        v76 = (v76 >> ((974) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v76
    _ = v77
}
