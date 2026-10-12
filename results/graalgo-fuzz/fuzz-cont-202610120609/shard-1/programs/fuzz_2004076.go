package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 int32; F1 int8; F2 string }
type GenS4 struct { F0 int8; F1 bool; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 int64; F2 int }
type GenS11 struct { F0 int32; F1 uint; F2 uint8 }
func genFn0(p0 int8) bool {
    if (1 >= 671) {
        for fi4 := 0; fi4 < 5; fi4++ {
            switch (fi4 | fi4) {
            case 9223372036854775807:
                var v5 bool = true
                _ = v5
            case -438:
                for fi6 := 0; fi6 < 14; fi6++ {
                    v7 := []int{(fi6 ^ fi6), 4294967295, fi4}
                    fmt.Println(fi6, p0)
                    _ = len(v7)
                    if false {
                        if (true && true) {
                            _ = len(v7)
                            if true {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                if len(v7) > 0 {
                                    v7[0] = fi6
                                }
                                p0 = p0
                            }
                            if (true || (!false)) {
                                if len(v7) > 0 {
                                    v7[0] = (fi6 | 9223372036854775807)
                                }
                                p0 += (p0 << ((0) & 63))
                                fmt.Println(fi6)
                            }
                            fmt.Println(fi4, fi4, p0)
                            if len(v7) > 0 {
                                v7[0] = (fi4 >> ((18446744073709551615) & 63))
                            }
                        }
                    }
                    p0 = (-(p0))
                    _ = fi6
                    _ = v7
                }
            case -960:
                p0 = p0
            }
            _ = fi4
        }
        var v8 float32 = float32(float64(-2.25))
        p0 -= 42
        _ = v8
    }
    fmt.Println(p0)
    v9 := GenS1{F0: 42, F1: 32767}
    _ = p0
    _ = v9
    return (32767 > -12)
}
func genFn1(p0 uint, p1 uint32) uint32 {
    var v10 float32 = 1.5
    if false {
        if (0 >= 18446744073709551615) {
            v11 := GenS2{F0: 4294967295, F1: 18446744073709551615, F2: ("" + "hello")}
            _ = v11
        }
        var v12 int16 = -1
        v13 := []int{2147483647, 127}
        _ = v12
        _ = v13
    } else {
        v14 := GenS3{F0: 42, F1: 0, F2: "hello"}
        v15 := GenS4{F0: 1, F1: (42 == -128), F2: 1}
        fmt.Println(p1)
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    return 116
}

func main() {
    v1 := []string{("a" + "hello")}
    var v2 bool = (0 == 350)
    var v3 uint8 = 128
    v16 := GenS2{F0: 136, F1: 4294967295, F2: "a"}
    var v17 uint = 255
    for fi18 := 0; fi18 < 17; fi18++ {
        if v2 {
            if (32767 <= 42) {
                var v19 uint64 = 364
                v20 := GenS5{F0: 0}
                _ = v19
                _ = v20
            }
        }
        var v21 int32 = 442
        switch "hello" {
        case "\n\t":
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 uint16 = 65535
                switch ("" + ("hello" + "a")) {
                case "\n\t":
                    var v24 string = "foo bar"
                    _ = v24
                }
                v23 = 255
                _ = fi22
                _ = v23
            }
        case "foo bar":
            var v25 int8 = 1
            _ = v25
        case "a":
            v17 += v17
        }
        _ = fi18
        _ = v21
    }
    switch "" {
    case "hello":
        var v26 int8 = 1
        _ = v26
    case "a":
        for fi27 := 0; fi27 < 11; fi27++ {
            for fi28 := 0; fi28 < 19; fi28++ {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                switch "foo bar" {
                case "foo bar":
                    switch ((("a" + (("hello" + "") + "\n\t")) + ("\n\t" + ("foo bar" + (("a" + ("hello" + "\n\t")) + "")))) + "a") {
                    case "":
                        var v29 int = fi28
                        _ = v29
                    case "a":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if v2 {
                            var v30 int64 = -64
                            v31 := GenS1{F0: v30, F1: 393}
                            v32 := GenS6{F0: fi27, F1: (v30 + (^(v30)))}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                default:
                    v17 += v17
                }
                _ = fi28
            }
            v33 := GenS7{F0: (fi27 % ((fi27) | 1))}
            _ = fi27
            _ = v33
        }
    case "foo bar":
        var v34 uint = (v17 * 0)
        _ = v34
    default:
        fmt.Println(v2)
    }
    switch -251 {
    case 127:
        var v35 string = ((("hello" + "\n\t") + "foo bar") + ("\n\t" + "foo bar"))
        _ = v35
    case 9223372036854775807:
        var v36 int = -1
        _ = v36
    case -140:
        v37 := GenS6{F0: -2147483648, F1: 1}
        _ = v37
    }
    _ = len(v1)
    for fi38 := 0; fi38 < 4; fi38++ {
        v17 = 2
        if v2 {
            var v39 float32 = float32((((1e10 + 0.1) / (float64(3.14159) - (0.1 / ((0.1 - -2.25) / -2.25)))) / 0.1))
            _ = v39
        } else {
            if (!(-32768 < -1)) {
                var v40 uint64 = 1
                _ = v40
            }
            for fi41 := 0; fi41 < 13; fi41++ {
                fmt.Println(v3)
                var v42 int32 = -2147483648
                v2 = (895 >= 1)
                _ = fi41
                _ = v42
            }
            v1 = append(v1, ("" + ""))
        }
        _ = fi38
    }
    v17 -= 0
    for fi43 := 0; fi43 < 18; fi43++ {
        if false {
            switch fi43 {
            case 42:
                _ = len(v1)
            case 32768:
                var v44 bool = (!(v2 && (!(v17 == v17))))
                _ = v44
            }
        }
        _ = fi43
    }
    fmt.Println(v17, v2, v3)
    v2 = (!v2)
    var v45 int64 = -571
    var v46 int32 = 42
    fmt.Println(v3, v2)
    var v47 uint64 = 711
    for fi48 := 0; fi48 < 3; fi48++ {
        if (!((2 != 36) && ("foo bar" >= (((("" + "") + "") + "hello") + ("" + ("" + "a")))))) {
            _ = len(v1)
            fmt.Println(v17, v2, v3)
        } else {
            if (false || (!((v17 + (v17 << (((v17 >> (((v17 ^ v17)) & 63))) & 63))) != (^(v17))))) {
                if false {
                    v17 = 1
                } else {
                    for fi49 := 0; fi49 < 17; fi49++ {
                        switch (-(fi48)) {
                        case 127:
                            var v50 uint32 = uint32(v3)
                            _ = v50
                        case 32767:
                            fmt.Println(v2, v17, fi49)
                        default:
                            if v2 {
                                switch (fi48 % ((32768) | 1)) {
                                case -581:
                                    if (3.14159 != 1.5) {
                                        var v51 uint32 = genFn1(uint(v47), 4294967295)
                                        _ = v51
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = ("\n\t" + ("a" + ("foo bar" + "a")))
                                        }
                                        if v2 {
                                            fmt.Println(v2, v3, v47)
                                            var v52 float32 = float32(1e10)
                                            var v53 float32 = v52
                                            _ = v52
                                            _ = v53
                                        }
                                    }
                                case 32767:
                                    if (v46 != 42) {
                                        if (v2 && (26 == 65535)) {
                                            var v54 uint = v17
                                            if false {
                                                v45 = (v45 << ((42) & 63))
                                                _ = len(v1)
                                            }
                                            _ = v54
                                        } else {
                                            if v2 {
                                                fmt.Println(v3, v47, v45)
                                            }
                                            fmt.Println(v3, v45, v2)
                                        }
                                        fmt.Println(v45)
                                        v46 -= (v46 ^ v46)
                                    }
                                case -2147483648:
                                    switch "hello" {
                                    case "foo bar":
                                        v3 = v3
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                        v55 := GenS8{F0: 42}
                                        _ = v55
                                    case "a":
                                        switch -986 {
                                        case 4294967295:
                                            v3 -= (v3 + v3)
                                        case 1:
                                            v56 := GenS9{F0: (v3 / ((127) | 1))}
                                            _ = v56
                                        case -1:
                                            v57 := GenS2{F0: 2, F1: (^(v47)), F2: "\n\t"}
                                            _ = v57
                                        default:
                                            var v58 int16 = 42
                                            _ = v58
                                        }
                                    default:
                                        v59 := GenS10{F0: fi48, F1: v45, F2: (-(fi49))}
                                        _ = v59
                                    }
                                    v45 = v45
                                    v45 = -410
                                    fmt.Println(v17, v3)
                                default:
                                    var v60 int32 = v46
                                    _ = v60
                                }
                                v2 = v2
                                v61 := GenS11{F0: (v46 >> (((v17 ^ (v17 ^ 2))) & 63)), F1: v17, F2: 1}
                                _ = v61
                            }
                        }
                        _ = fi49
                    }
                    v3 = 255
                    switch 42 {
                    case -9223372036854775808:
                        if v2 {
                            var v62 int16 = 1
                            _ = v62
                        }
                    }
                }
            }
        }
        _ = fi48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
}
