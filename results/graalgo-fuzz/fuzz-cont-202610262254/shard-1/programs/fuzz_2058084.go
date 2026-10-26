package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
type GenS5 struct { F0 int16 }
func genFn0() bool {
    switch "foo bar" {
    case "\n\t":
        for fi5 := 0; fi5 < 15; fi5++ {
            fmt.Println(fi5)
            if (42 >= 85) {
                var v6 float64 = -2.25
                _ = v6
            } else {
                fmt.Println(fi5)
            }
            v7 := []int{9223372036854775807}
            _ = fi5
            _ = v7
        }
    case "foo bar":
        v8 := []int{-2147483648, -58}
        _ = v8
    case "a":
        var v9 int64 = 0
        _ = v9
    }
    var v10 int32 = 169
    v11 := GenS1{F0: (v10 | (^(v10))), F1: 1}
    v12 := []int{-9223372036854775808, -9223372036854775808, 1}
    _ = v10
    _ = v11
    _ = v12
    return true
}
func genFn1() uint8 {
    var v13 uint64 = 1
    var v14 uint8 = 175
    var v15 uint64 = 0
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    v1 := []int{32768}
    var v2 uint8 = 42
    v3 := []uint32{366, 926, 930, 0}
    v4 := make(map[string]string)
    var v16 uint32 = 227
    v2 = (^(v2))
    if true {
        if len(v3) > 0 {
            v3[0] = 65535
        }
        if false {
            v2 = v2
        }
    } else {
        v2 = (v2 - genFn1())
        v17 := GenS1{F0: 1, F1: 13}
        _ = v17
    }
    v2 -= v2
    if (-32768 > 334) {
        var v18 uint32 = v16
        if ((false || (((23 >= 1) || true) && true)) || false) {
            var v19 int8 = 100
            var v20 uint64 = 0
            switch 1 {
            case -2147483648:
                var v21 uint8 = (v2 >> ((4294967295) & 63))
                _ = v21
            case 29:
                if (false || (236 >= 907)) {
                    v22 := GenS1{F0: -75, F1: 409}
                    v18 = 42
                    fmt.Println(v2)
                    _ = v22
                }
            case 127:
                switch "hello" {
                case "\n\t":
                    delete(v4, ("" + ""))
                }
            default:
                for fi23 := 0; fi23 < 6; fi23++ {
                    v19 = v19
                    for fi24 := 0; fi24 < 3; fi24++ {
                        var v25 int16 = -1
                        _ = len(v4)
                        var v26 float64 = (3.14159 + ((0.0 * ((3.14159 / 1e10) + (-2.25 * 0.1))) - 0.1))
                        v27 := v4[(("foo bar" + (("foo bar" + ("hello" + "a")) + ("hello" + ("" + "hello")))) + "a")]
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    if true {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            fmt.Println(v16, fi23, v19)
                            _ = fi28
                        }
                    }
                    _ = fi23
                }
            }
            _ = v19
            _ = v20
        }
        _ = v18
    }
    v29 := GenS1{F0: 42, F1: 502}
    var v30 uint8 = 42
    if (!((1 > 65535) && ((1e10 < 1.5) && ((-2147483648 <= 2147483647) || (float64(((0.1 + 3.14159) + 0.0)) <= 0.0))))) {
        var v31 uint = 1
        _ = v31
    } else {
        var v32 float64 = 3.14159
        var v33 int8 = 98
        _ = v32
        _ = v33
    }
    v2 = v30
    v16 -= v16
    if false {
        if (0 == 0) {
            switch ("foo bar" + ("a" + "a")) {
            case "\n\t":
                v2 -= v30
            case "a":
                var v34 uint32 = 1
                _ = v34
            case "":
                v1 = append(v1, 32767)
            default:
                for fi35 := 0; fi35 < 12; fi35++ {
                    v36 := GenS2{F0: "\n\t", F1: genFn1(), F2: 1}
                    _ = fi35
                    _ = v36
                }
            }
            v37 := GenS2{F0: "a", F1: (v2 + (^(v2))), F2: 42}
            if false {
                var v38 int32 = -276
                v39 := GenS3{F0: v2, F1: 42}
                v16 = (^(v16))
                _ = v38
                _ = v39
            } else {
                _ = len(v4)
                _ = len(v3)
                var v40 int32 = -791
                _ = v40
            }
            _ = v37
        } else {
            var v41 int64 = 653
            if (true && false) {
                v42 := GenS3{F0: 163, F1: -32768}
                var v43 int64 = -129
                fmt.Println(v41, v30)
                _ = v42
                _ = v43
            }
            _ = v41
        }
        v30 -= v2
        var v44 bool = true
        _ = v44
    } else {
        var v45 int = 127
        delete(v4, ("foo bar" + ""))
        _ = v45
    }
    v46 := GenS4{F0: 2, F1: -1, F2: false}
    for fi47 := 0; fi47 < 1; fi47++ {
        if (v30 == v30) {
            _ = len(v4)
            v48 := GenS3{F0: 2, F1: 1000}
            _ = v48
        }
        _ = fi47
    }
    var v49 float32 = 0.1
    switch ("a" + "") {
    case "\n\t":
        switch "foo bar" {
        case "":
            _ = len(v4)
        case "a":
            v16 -= v16
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case "a":
        v50 := GenS5{F0: -538}
        _ = v50
    case "":
        if true {
            var v51 int16 = 42
            v49 -= (float32((1.5 + 3.14159)) / v49)
            switch "hello" {
            case "a":
                for fi52 := 0; fi52 < 9; fi52++ {
                    v2 = 127
                    _ = fi52
                }
            case "foo bar":
                var v53 string = "foo bar"
                _ = v53
            default:
                fmt.Println(v49, v16, v49)
            }
            _ = v51
        } else {
            var v54 string = "\n\t"
            if ((-72 <= 42) && true) {
                var v55 uint8 = (v30 + genFn1())
                _ = v55
            }
            _ = v54
        }
    }
    v3 = append(v3, (^(v16)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v46
    _ = v49
}
