package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int; F1 int64; F2 int16 }
type GenS5 struct { F0 int8; F1 uint32; F2 string }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 bool; F1 uint64; F2 int64 }
type GenS8 struct { F0 int8; F1 int32 }
func genFn0(p0 int32, p1 int32) int {
    var v6 string = "\n\t"
    if (!true) {
        fmt.Println(p0, p1)
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{(fi7 + (fi7 % (((^(fi7))) | 1))), (fi7 << ((4294967295) & 63))}
            var v9 int32 = (p1 - p0)
            if true {
                switch "\n\t" {
                case "hello":
                    var v10 int16 = -528
                    _ = v10
                case "":
                    var v11 int16 = 32767
                    _ = v11
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        v12 := GenS1{F0: 0, F1: 0, F2: 548}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{65535, 2, 280}
    var v3 int64 = 42
    var v4 int8 = 1
    v5 := make(map[int]string)
    fmt.Println(v4, v3)
    v13 := GenS2{F0: ("hello" + "\n\t"), F1: 1}
    if (0 >= 128) {
        v14 := v1[65535]
        _ = v14
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint32 = 65535
            _ = fi15
            _ = v16
        }
        if false {
            var v17 string = (("a" + ("hello" + ("a" + "a"))) + "")
            if true {
                if false {
                    var v18 int8 = v4
                    var v19 bool = false
                    if (!false) {
                        var v20 int = 32768
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            var v22 int8 = 100
                            _ = fi21
                            _ = v22
                        }
                        var v23 int16 = -715
                        v24 := GenS3{F0: v18, F1: "\n\t"}
                        _ = v23
                        _ = v24
                    }
                    _ = v18
                    _ = v19
                } else {
                    var v25 bool = (42 < 65535)
                    if (1e10 >= 0.1) {
                        var v26 uint = 4294967295
                        _ = v26
                    }
                    if false {
                        if (0.0 == (0.1 - 1.5)) {
                            v27 := GenS1{F0: 42, F1: v4, F2: 65535}
                            var v28 int32 = -279
                            var v29 int = -402
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        v17 = v17
                    }
                    _ = v25
                }
                switch ("\n\t" + (v17 + ((v17 + "") + ("foo bar" + "foo bar")))) {
                case "":
                    var v30 string = (v17 + (v17 + v17))
                    _ = v30
                default:
                    if true {
                        v17 = (v17 + v17)
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        v4 = (v4 << ((2) & 63))
                    } else {
                        if (2147483647 <= 1) {
                            fmt.Println(v17)
                            var v31 int64 = v3
                            switch 42 {
                            case 475:
                                if true {
                                    for fi32 := 0; fi32 < 5; fi32++ {
                                        var v33 string = (v17 + v17)
                                        var v34 uint16 = 2
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            default:
                                switch -2147483648 {
                                case -947:
                                    v35 := GenS4{F0: 2147483647, F1: 736, F2: 1000}
                                    _ = v35
                                case 578:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    v31 += (v31 ^ (^(v3)))
                                }
                            }
                            _ = v31
                        }
                        for fi36 := 0; fi36 < 17; fi36++ {
                            for fi37 := 0; fi37 < 10; fi37++ {
                                switch v17 {
                                case "foo bar":
                                    v3 -= (v3 | (v3 % ((v3) | 1)))
                                default:
                                    if (false || (!true)) {
                                        if true {
                                            v38 := GenS5{F0: -128, F1: 2147483647, F2: v17}
                                            fmt.Println(fi36)
                                            delete(v5, fi37)
                                            _ = v38
                                        } else {
                                            var v39 int64 = 0
                                            v39 = v3
                                            if (2 != 112) {
                                                if (0 != 127) {
                                                    var v40 int64 = v3
                                                    _ = v40
                                                }
                                                var v41 int8 = (v4 ^ 127)
                                                _ = v41
                                            } else {
                                                v2 = append(v2, 993)
                                            }
                                            _ = v39
                                        }
                                    } else {
                                        v4 += -94
                                        switch fi36 {
                                        case 128:
                                            _ = len(v5)
                                        case 4294967295:
                                            switch ((v17 + v17) + v17) {
                                            case "\n\t":
                                                fmt.Println(fi36)
                                            case "foo bar":
                                                if false {
                                                    fmt.Println(v3, fi36)
                                                } else {
                                                    v17 = "hello"
                                                }
                                            }
                                        }
                                        v2 = append(v2, 2147483647)
                                    }
                                }
                                switch "a" {
                                case "a":
                                    var v42 float64 = 1.5
                                    _ = v42
                                case "":
                                    var v43 uint = 65535
                                    _ = v43
                                }
                                _ = fi37
                            }
                            _ = fi36
                        }
                        fmt.Println(v3, v3)
                    }
                }
            }
            _ = v17
        }
        v5[42] = ("" + "\n\t")
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        var v45 int32 = 1
        if true {
            var v46 float32 = 3.14159
            _ = v46
        }
        v45 += v45
        _ = fi44
        _ = v45
    }
    if false {
        switch 32767 {
        case 32767:
            for fi47 := 0; fi47 < 10; fi47++ {
                for fi48 := 0; fi48 < 10; fi48++ {
                    switch (fi48 - genFn0(394, 2147483647)) {
                    case 0:
                        v49 := GenS6{F0: fi48, F1: 2147483647}
                        _ = v49
                    case 885:
                        v2 = append(v2, 65535)
                        v50 := GenS3{F0: (^(v4)), F1: "foo bar"}
                        _ = v50
                    default:
                        switch "a" {
                        case "hello":
                            delete(v5, (fi48 % ((856) | 1)))
                        case "foo bar":
                            v51 := GenS3{F0: (^(v4)), F1: ("a" + "a")}
                            _ = v51
                        default:
                            fmt.Println(v3, fi47)
                        }
                    }
                    if (441 < 995) {
                        if true {
                            v52 := GenS7{F0: (0 != (v3 - -105)), F1: 1, F2: v3}
                            if false {
                                var v53 uint8 = 128
                                if false {
                                    fmt.Println(v3, fi48)
                                    v53 -= (v53 + (v53 & 19))
                                    var v54 int32 = 1
                                    _ = v54
                                } else {
                                    v4 = (v4 + v4)
                                    var v55 string = "a"
                                    v55 = v55
                                    _ = v55
                                }
                                _ = v53
                            }
                            _ = v52
                        }
                        var v56 string = ""
                        v56 = v56
                        _ = v56
                    }
                    if (!false) {
                        var v57 int = fi48
                        _ = v57
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v4 = 42
                        v3 = (v3 * v3)
                    }
                    _ = fi48
                }
                for fi58 := 0; fi58 < 17; fi58++ {
                    v59 := v5[genFn0(-630, -368)]
                    _ = fi58
                    _ = v59
                }
                _ = fi47
            }
        }
        v4 -= v4
        for fi60 := 0; fi60 < 7; fi60++ {
            if true {
                fmt.Println(v4, v4)
            }
            var v61 int = -1
            _ = fi60
            _ = v61
        }
    }
    switch "hello" {
    case "a":
        var v62 uint64 = 6
        _ = v62
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            v63 := GenS7{F0: (false || (v4 != -128)), F1: 18446744073709551615, F2: (v3 >> ((2) & 63))}
            _ = v63
        }
        var v64 uint16 = 0
        _ = v64
    } else {
        fmt.Println(v4)
    }
    v3 = v3
    var v65 string = "hello"
    var v66 float64 = float64(0.0)
    if false {
        v67 := GenS7{F0: false, F1: 18446744073709551615, F2: v3}
        v68 := GenS8{F0: 100, F1: 1}
        _ = v67
        _ = v68
    } else {
        _ = len(v2)
        v4 = (v4 / ((29) | 1))
        v66 = (v66 / float64(3.14159))
    }
    switch v65 {
    case "\n\t":
        v69 := v5[genFn0(1, 42)]
        _ = v69
    }
    for fi70 := 0; fi70 < 10; fi70++ {
        v3 = (^(v3))
        v66 += v66
        _ = fi70
    }
    var v71 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v71
}
