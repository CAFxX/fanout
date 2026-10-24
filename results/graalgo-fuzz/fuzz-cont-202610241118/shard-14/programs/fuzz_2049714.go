package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 string; F1 int; F2 int8 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint, p1 int64) uint64 {
    if true {
        v6 := []int{132, 9223372036854775807, -2147483648}
        if len(v6) > 0 {
            v6[0] = 99
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 776}
            var v9 uint64 = 4294967295
            var v10 int64 = p1
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = v6
    }
    var v11 uint16 = 1
    switch "\n\t" {
    case "hello":
        for fi12 := 0; fi12 < 11; fi12++ {
            p1 = (p1 & -9223372036854775808)
            _ = fi12
        }
    default:
        if false {
            fmt.Println(v11, v11, p0)
            var v13 uint32 = 42
            p1 += 0
            _ = v13
        }
    }
    v14 := GenS2{F0: 923}
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    return 2
}
func genFn1() uint {
    v15 := GenS3{F0: 0}
    switch 34 {
    case 128:
        switch "foo bar" {
        case "hello":
            var v16 int16 = -32768
            _ = v16
        case "foo bar":
            var v17 uint8 = 84
            _ = v17
        case "\n\t":
            var v18 uint64 = 1
            _ = v18
        }
    default:
        switch -9223372036854775808 {
        case 42:
            if (1e10 == (3.14159 - 0.1)) {
                switch -2147483648 {
                case 32767:
                    switch ("" + "") {
                    case "hello":
                        var v19 uint16 = 377
                        _ = v19
                    case "\n\t":
                        v20 := []int{-36, 1, 32768}
                        _ = v20
                    case "foo bar":
                        switch ("" + "") {
                        case "":
                            v21 := GenS3{F0: 42}
                            _ = v21
                        case "a":
                            var v22 float32 = float32(float64((3.14159 * 1e10)))
                            _ = v22
                        case "foo bar":
                            if true {
                                v23 := GenS4{F0: 270, F1: 0}
                                var v24 uint32 = 65535
                                if true {
                                    v24 -= v24
                                    if (((127 < 90) && (((("foo bar" + ("a" + ("foo bar" + "hello"))) + "hello") + "\n\t") <= ((("" + "a") + "") + "\n\t"))) && false) {
                                        v25 := []int{}
                                        v26 := GenS1{F0: 772}
                                        v24 += v24
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v27 := GenS2{F0: 2147483647}
                                        _ = v27
                                    }
                                    v28 := GenS5{F0: ("" + (("" + "\n\t") + (("a" + "hello") + (("" + "") + ("\n\t" + "a"))))), F1: -2147483648, F2: 0}
                                    _ = v28
                                }
                                _ = v23
                                _ = v24
                            }
                        }
                    default:
                        v29 := []int{32767, 128, -817}
                        _ = v29
                    }
                case 4294967295:
                    v30 := GenS1{F0: 2}
                    _ = v30
                default:
                    var v31 int16 = 0
                    _ = v31
                }
                v32 := []int{597, 9223372036854775807}
                _ = v32
            }
        case 859:
            var v33 uint64 = 18446744073709551615
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        if true {
            var v35 bool = (1 >= 219)
            _ = v35
        }
        switch "hello" {
        case "hello":
            var v36 uint8 = 0
            _ = v36
        case "a":
            var v37 uint64 = 42
            _ = v37
        default:
            var v38 int = (fi34 / ((fi34) | 1))
            _ = v38
        }
        switch "foo bar" {
        case "foo bar":
            var v39 int = 127
            _ = v39
        case "a":
            var v40 int16 = 540
            _ = v40
        }
        _ = fi34
    }
    _ = v15
    return 1
}

func main() {
    var v1 uint = 255
    var v2 string = ((("foo bar" + "") + "") + ("hello" + "hello"))
    var v3 float64 = (float64(1e10) + (1e10 - (((float64(1e10) - ((3.14159 / (-2.25 / 1.0)) * 1.5)) - 0.1) - 1e10)))
    var v4 uint = v1
    v5 := make(map[string]bool)
    switch 42 {
    case 11:
        var v41 bool = (!true)
        _ = v41
    case 9223372036854775807:
        fmt.Println(v1)
    case 2147483647:
        v2 = "foo bar"
    }
    v1 = 837
    v42 := GenS4{F0: v1, F1: -630}
    for fi43 := 0; fi43 < 6; fi43++ {
        var v44 int32 = 572
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 20; fi45++ {
        var v46 uint8 = 255
        _ = fi45
        _ = v46
    }
    v3 = ((((3.14159 + v3) / (v3 - ((v3 / 0.1) - (0.1 / v3)))) * float64(0.0)) - (v3 + (((1e10 + (float64(((-2.25 / float32((v3 - 0.1))) * (1.5 - 0.0))) / v3)) + v3) - ((v3 / (v3 / v3)) + v3))))
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 int64 = -1
        var v49 int32 = 2147483647
        _ = fi47
        _ = v48
        _ = v49
    }
    var v50 bool = (v2 != v2)
    v51 := []int{1, 2147483647, -9223372036854775808, -723}
    if ((!((float32(v3) / float32(v3)) <= 1.5)) || (-1 < -9223372036854775808)) {
        v52 := GenS6{F0: -9223372036854775808}
        _ = v52
    }
    switch (v2 + "") {
    case "hello":
        delete(v5, v2)
    case "":
        for fi53 := 0; fi53 < 9; fi53++ {
            for fi54 := 0; fi54 < 7; fi54++ {
                v55 := GenS5{F0: ("foo bar" + v2), F1: (fi53 & fi54), F2: 127}
                _ = fi54
                _ = v55
            }
            v3 = float64(1.5)
            for fi56 := 0; fi56 < 3; fi56++ {
                v57 := GenS1{F0: 65535}
                _ = fi56
                _ = v57
            }
            _ = fi53
        }
    default:
        fmt.Println(v2)
    }
    for fi58 := 0; fi58 < 3; fi58++ {
        for fi59 := 0; fi59 < 12; fi59++ {
            v5["hello"] = (127 != 0)
            v60 := GenS2{F0: 2}
            _ = fi59
            _ = v60
        }
        _ = fi58
    }
    var v61 uint = (v1 * (v1 >> (((^(v1))) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v50
    _ = v51
    _ = v61
}
