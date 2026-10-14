package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int64; F1 int16; F2 int8 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint }
func genFn0(p0 uint) string {
    p0 += 65535
    v4 := GenS1{F0: -1, F1: 0}
    _ = p0
    _ = v4
    return ("foo bar" + ("\n\t" + "hello"))
}
func genFn1(p0 uint, p1 int16) uint8 {
    fmt.Println(p1)
    v5 := GenS1{F0: p1, F1: 66}
    var v6 int16 = 32767
    for fi7 := 0; fi7 < 13; fi7++ {
        switch fi7 {
        case 32768:
            p0 = (p0 << ((673) & 63))
        }
        if ((-(p1)) == v6) {
            var v8 uint64 = 1
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint32 = 446
                _ = fi9
                _ = v10
            }
            var v11 uint = p0
            _ = v11
            _ = v8
        }
        v12 := []int{-684}
        _ = fi7
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS2{F0: 9223372036854775807, F1: p1, F2: -20}
        if false {
            var v15 int32 = -633
            _ = v15
        }
        v16 := GenS1{F0: (v6 << ((18446744073709551615) & 63)), F1: 42}
        _ = fi13
        _ = v14
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int64 = 1
    v2 := []uint32{42, 215}
    v3 := []string{"", ("" + "a"), "a", ("a" + "")}
    fmt.Println(v1)
    if true {
        if (!(42 > 4294967295)) {
            var v17 int8 = -22
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v17
        } else {
            if true {
                v1 += (v1 >> ((2) & 63))
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            }
            v1 = v1
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            _ = len(v3)
            for fi19 := 0; fi19 < 2; fi19++ {
                _ = len(v3)
                _ = fi19
            }
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = fi18
        }
        if false {
            if (!((!(false || (!(!false)))) && false)) {
                v2 = append(v2, 4294967295)
                v20 := GenS2{F0: (v1 * 388), F1: 1, F2: 16}
                v21 := GenS3{F0: v1, F1: 32767, F2: 2}
                _ = v20
                _ = v21
            }
            _ = len(v3)
        } else {
            v22 := GenS3{F0: v1, F1: 2, F2: 283}
            _ = len(v3)
            _ = v22
        }
    } else {
        if false {
            var v23 uint32 = 2147483648
            v1 = v1
            v24 := GenS3{F0: -1, F1: 555, F2: v23}
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            if false {
                var v26 string = ("" + "\n\t")
                _ = v26
            } else {
                if false {
                    _ = len(v3)
                    v1 += (v1 - v1)
                    var v27 float32 = 1.5
                    _ = v27
                } else {
                    var v28 uint = 4294967295
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v28
                }
                switch fi25 {
                case 32768:
                    switch "hello" {
                    case "":
                        v29 := GenS2{F0: v1, F1: 32767, F2: 0}
                        _ = v29
                    default:
                        for fi30 := 0; fi30 < 7; fi30++ {
                            var v31 int64 = v1
                            v32 := GenS3{F0: 616, F1: 0, F2: 2}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    }
                case 2147483647:
                    fmt.Println(fi25)
                }
                v1 += (v1 | (v1 & (v1 * 1)))
            }
            _ = fi25
        }
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        v1 += (v1 >> ((224) & 63))
        switch "a" {
        case "hello":
            fmt.Println(v1)
        default:
            if len(v3) > 0 {
                v3[0] = "a"
            }
        }
        _ = fi33
    }
    var v34 bool = true
    if true {
        v3 = append(v3, "hello")
        fmt.Println(v1, v1)
    } else {
        v34 = (!(32767 >= -32768))
        var v35 float64 = 0.0
        v3 = append(v3, "a")
        _ = v35
    }
    fmt.Println(v1)
    if (4294967295 < -9223372036854775808) {
        v2 = append(v2, 2147483648)
        for fi36 := 0; fi36 < 2; fi36++ {
            _ = len(v2)
            var v37 uint64 = 1
            for fi38 := 0; fi38 < 15; fi38++ {
                fmt.Println(v1)
                v34 = (42 == 0)
                v1 -= (v1 & 42)
                v39 := GenS1{F0: 1, F1: 128}
                _ = fi38
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        v40 := GenS1{F0: 0, F1: 0}
        _ = v40
    } else {
        switch "hello" {
        case "":
            var v41 int32 = 0
            _ = v41
        case "a":
            var v42 float64 = 3.14159
            _ = v42
        }
        if (-128 < 85) {
            v1 += v1
            if (v34 && (!(1 < 1))) {
                if v34 {
                    var v43 float32 = 1e10
                    _ = v43
                } else {
                    switch "foo bar" {
                    case "":
                        var v44 bool = true
                        _ = v44
                    }
                }
                v34 = (!(true && v34))
                var v45 int32 = 1
                _ = v45
            }
        } else {
            fmt.Println(v34, v34)
        }
        v34 = true
    }
    if v34 {
        var v46 int32 = 0
        _ = v46
    } else {
        if (!(false || (4294967295 < 1))) {
            v34 = (((0.1 - 1.5) + 0.1) > (1e10 * 0.1))
            switch "" {
            case "a":
                v47 := GenS4{F0: 18}
                _ = v47
            default:
                for fi48 := 0; fi48 < 7; fi48++ {
                    var v49 float64 = (float64(0.1) / float64(float32(1.5)))
                    _ = fi48
                    _ = v49
                }
            }
        } else {
            var v50 int8 = -86
            _ = v50
        }
    }
    v34 = v34
    v1 += (v1 * 0)
    fmt.Println(v1, v1)
    v51 := GenS5{F0: 2, F1: -683, F2: 765}
    v34 = (!(953 < 18446744073709551615))
    switch -2147483648 {
    case 882:
        var v52 int64 = 0
        _ = v52
    default:
        fmt.Println(v34)
    }
    var v53 bool = (true && true)
    switch "hello" {
    case "\n\t":
        fmt.Println(v1)
    case "":
        var v54 uint = 18446744073709551615
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v51
    _ = v53
}
