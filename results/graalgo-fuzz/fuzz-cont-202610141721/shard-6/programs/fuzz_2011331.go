package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int64; F1 int }
func genFn0(p0 uint, p1 bool, p2 uint8) int16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        p0 += 42
        var v8 string = ""
        if ((p0 - p0) <= 234) {
            v9 := []int{(fi7 << ((p0) & 63)), (fi7 & 9223372036854775807), 2147483647, 458}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v10 uint32 = 2147483647
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            switch -9223372036854775808 {
            case 127:
                var v13 int8 = 72
                _ = v13
            case 747:
                v10 += uint32(p2)
                fmt.Println(p1, p2)
            }
            _ = fi12
        }
        var v14 uint16 = 42
        _ = fi11
        _ = v14
    }
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return int16(v10)
}

func main() {
    var v1 float32 = 1e10
    v2 := []int{-25}
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]uint32)
    var v6 uint16 = 1
    if ((249 == 457) && false) {
        var v15 int8 = 1
        for fi16 := 0; fi16 < 16; fi16++ {
            v15 = (v15 ^ (v15 % (((v15 % ((v15) | 1))) | 1)))
            _ = fi16
        }
        if (795 >= 652) {
            if (true || true) {
                var v17 int32 = 2147483647
                _ = v17
            } else {
                v18 := GenS1{F0: (v6 / (((^(v6))) | 1))}
                if true {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if false {
                            fmt.Println(v6)
                        } else {
                            switch "a" {
                            case "foo bar":
                                v15 -= v15
                            default:
                                v20 := GenS1{F0: (v6 ^ 0)}
                                _ = v20
                            }
                        }
                        for fi21 := 0; fi21 < 17; fi21++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi21
                        }
                        _ = fi19
                    }
                    v22 := GenS2{F0: 255, F1: 0, F2: true}
                    _ = v22
                }
                _ = v18
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                if true {
                    v6 += 659
                    if (!true) {
                        v24 := GenS3{F0: 127}
                        fmt.Println(v1, fi23)
                        _ = v24
                    }
                }
                if true {
                    if len(v2) > 0 {
                        v2[0] = (-(fi23))
                    }
                } else {
                    switch 32767 {
                    case 42:
                        var v25 float32 = ((v1 / v1) - float32(0.1))
                        _ = v25
                    }
                }
                var v26 uint64 = 18446744073709551615
                _ = fi23
                _ = v26
            }
        } else {
            delete(v4, 65535)
        }
        _ = v15
    }
    v1 -= (1e10 * v1)
    _ = len(v5)
    fmt.Println(v6, v6)
    switch 128 {
    case -496:
        var v27 uint = 255
        _ = v27
    case 799:
        var v28 float64 = 0.1
        _ = v28
    case 959:
        var v29 int = 9223372036854775807
        _ = v29
    default:
        switch (((("" + "hello") + "foo bar") + ((("foo bar" + (("foo bar" + ("hello" + "a")) + "foo bar")) + "\n\t") + ("\n\t" + ((("foo bar" + (("\n\t" + ("" + ("hello" + ""))) + "hello")) + ("\n\t" + "foo bar")) + ("a" + ""))))) + "") {
        case "hello":
            v30 := GenS4{F0: 0, F1: 32767, F2: 42}
            _ = v30
        default:
            v31 := GenS5{F0: ((("\n\t" + "\n\t") + "hello") + "hello")}
            _ = v31
        }
    }
    _ = len(v4)
    fmt.Println(v6, v1)
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 bool = (1 < 950)
        for fi34 := 0; fi34 < 2; fi34++ {
            fmt.Println(v1)
            v33 = v33
            _ = fi34
        }
        switch "\n\t" {
        case "":
            fmt.Println(v1)
        case "foo bar":
            var v35 int16 = genFn0(2, (32767 > -210), 2)
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    v2 = append(v2, -69)
    if true {
        var v36 uint64 = 4294967295
        v36 -= v36
        _ = v36
    }
    if false {
        if false {
            _ = len(v2)
            if (551 == 693) {
                var v37 string = ""
                var v38 uint8 = 86
                var v39 int16 = 329
                _ = v37
                _ = v38
                _ = v39
            } else {
                switch "" {
                case "foo bar":
                    var v40 float32 = v1
                    _ = v40
                }
                v6 = (v6 % (((v6 & v6)) | 1))
            }
        }
        var v41 int8 = 0
        var v42 bool = ((v41 >> ((0) & 63)) <= v41)
        _ = v41
        _ = v42
    }
    v6 += (v6 * (v6 << ((157) & 63)))
    v6 = (v6 << ((4294967295) & 63))
    if true {
        v5["foo bar"] = 797
    } else {
        var v43 bool = (-1 == 42)
        _ = v43
    }
    v1 = v1
    for fi44 := 0; fi44 < 15; fi44++ {
        v6 -= (v6 + 255)
        fmt.Println(v6, v6, fi44)
        if (genFn0(2, (false || true), 1) <= 430) {
            if true {
                var v45 uint = 1
                fmt.Println(v1, v45)
                var v46 int64 = 9223372036854775807
                _ = v45
                _ = v46
            } else {
                v47 := GenS6{F0: 121, F1: v6}
                fmt.Println(v6, v6)
                var v48 float32 = -2.25
                _ = v47
                _ = v48
            }
            v49 := GenS7{F0: -42, F1: 953}
            _ = v49
        }
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
