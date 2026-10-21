package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0(p0 string, p1 int32, p2 int8) uint16 {
    fmt.Println(p1)
    p2 += (p2 - (p2 << ((1) & 63)))
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int32) uint8 {
    v4 := []int{0, 9223372036854775807}
    var v5 int8 = 0
    switch int(p0) {
    case 42:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 2147483647:
        var v6 int16 = 457
        _ = v6
    }
    var v7 int8 = v5
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return 36
}

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 352
    var v3 uint32 = 509
    if true {
        v9 := GenS1{F0: 0, F1: "foo bar", F2: false}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            if true {
                var v11 uint32 = (v3 * v3)
                v1 = v1
                _ = v11
            }
            v1 -= -2.25
            switch "foo bar" {
            case "a":
                if (1.5 >= 0.1) {
                    var v12 uint = 828
                    v13 := GenS2{F0: v12, F1: ((((("" + (("foo bar" + "foo bar") + ("" + ("" + ("foo bar" + "a"))))) + ("hello" + "a")) + (("" + "hello") + "")) + "hello") + ("a" + "foo bar"))}
                    _ = v12
                    _ = v13
                }
            default:
                var v14 int16 = 294
                _ = v14
            }
            _ = fi10
        }
        v15 := []int{}
        if (250 <= 42) {
            var v16 uint16 = 32767
            v1 -= (3.14159 - 3.14159)
            _ = v16
        } else {
            v15 = append(v15, 1)
            v1 += 0.1
            var v17 uint64 = v2
            _ = v17
        }
        _ = v15
    }
    v18 := GenS2{F0: uint(v3), F1: ""}
    v19 := GenS3{F0: 128}
    v20 := GenS1{F0: 32768, F1: "foo bar", F2: true}
    var v21 int32 = 0
    v22 := []int{}
    v21 += (v21 + v21)
    v22 = append(v22, int(v21))
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            if (!((true && true) || false)) {
                v21 = int32(fi23)
                if false {
                    if (((42 >= 127) && false) || false) {
                        v25 := GenS3{F0: 255}
                        v22 = append(v22, (fi23 / (((-(fi24))) | 1)))
                        v1 = v1
                        if true {
                            v26 := GenS1{F0: 1, F1: "\n\t", F2: false}
                            switch "\n\t" {
                            case "hello":
                                var v27 bool = (v21 == (v21 & 42))
                                _ = v27
                            case "foo bar":
                                v22 = append(v22, 127)
                            case "\n\t":
                                fmt.Println(v1, v3)
                            default:
                                _ = len(v22)
                            }
                            _ = v26
                        }
                        _ = v25
                    }
                    if len(v22) > 0 {
                        v22[0] = (fi24 * -524)
                    }
                    var v28 int16 = -363
                    _ = v28
                }
            }
            if true {
                if (!(false || true)) {
                    fmt.Println(v3)
                }
                var v29 int64 = -1
                _ = v29
            } else {
                var v30 string = "\n\t"
                v21 = (-(v21))
                switch int(v3) {
                case 1:
                    fmt.Println(v1)
                }
                _ = v30
            }
            if true {
                v31 := GenS4{F0: 255}
                var v32 float64 = v1
                _ = v31
                _ = v32
            }
            _ = fi24
        }
        switch ("\n\t" + "hello") {
        case "\n\t":
            v33 := GenS5{F0: v3, F1: (v3 | (v3 - 4294967295)), F2: 100}
            _ = v33
        case "foo bar":
            v21 = (v21 >> ((873) & 63))
        default:
            var v34 float32 = 1.5
            _ = v34
        }
        var v35 float64 = v1
        _ = fi23
        _ = v35
    }
    var v36 uint16 = 403
    switch 127 {
    case 0:
        var v37 string = "\n\t"
        _ = v37
    case -2147483648:
        var v38 float32 = float32(v1)
        _ = v38
    case -215:
        switch ("hello" + "foo bar") {
        case "\n\t":
            _ = len(v22)
        }
    default:
        if true {
            var v39 int8 = int8(v2)
            _ = v39
        } else {
            switch "" {
            case "a":
                fmt.Println(v2, v1, v36)
            case "":
                v3 -= (v3 + v3)
            case "foo bar":
                if (false || ("" <= "")) {
                    if false {
                        var v40 int8 = 127
                        var v41 uint = 760
                        _ = v40
                        _ = v41
                    }
                    fmt.Println(v1)
                    var v42 uint16 = (v36 >> ((uint(v36)) & 63))
                    _ = v42
                }
            }
            v2 -= (v2 | v2)
            if len(v22) > 0 {
                v22[0] = -9223372036854775808
            }
        }
    }
    var v43 int32 = 2147483647
    fmt.Println(v1, v43, v21)
    if ((genFn1(v43) == 42) && false) {
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        if ((255 >= 4294967295) && (true && (0 > -2147483648))) {
            for fi44 := 0; fi44 < 16; fi44++ {
                var v45 uint8 = 255
                v3 = 357
                v46 := GenS3{F0: 127}
                _ = fi44
                _ = v45
                _ = v46
            }
            if len(v22) > 0 {
                v22[0] = 128
            }
        } else {
            if false {
                var v47 float64 = v1
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = v47
            } else {
                var v48 bool = false
                v48 = v48
                _ = v48
            }
            v49 := GenS4{F0: 0}
            if true {
                switch "hello" {
                case "a":
                    var v50 float64 = v1
                    _ = v50
                default:
                    fmt.Println(v3, v43, v2)
                }
            }
            _ = v49
        }
    }
    if true {
        v51 := GenS6{F0: v36, F1: -1, F2: -86}
        fmt.Println(v1, v36)
        _ = v51
    } else {
        fmt.Println(v3, v21, v43)
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        if len(v22) > 0 {
            v22[0] = -468
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v43
}
