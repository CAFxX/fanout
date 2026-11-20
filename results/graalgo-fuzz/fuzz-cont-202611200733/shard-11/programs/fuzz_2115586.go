package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint16; F2 int32 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 string, p1 int16) bool {
    v7 := GenS1{F0: -9223372036854775808}
    var v8 uint16 = 0
    switch "a" {
    case "":
        v9 := GenS1{F0: 4294967295}
        _ = v9
    default:
        switch p0 {
        case "foo bar":
            v10 := []int{457, -9223372036854775808, 18, 1}
            _ = v10
        }
    }
    if (!false) {
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS2{F0: (false || (18446744073709551615 == 65535)), F1: 781, F2: 0}
                fmt.Println(p0, v8, fi12)
                v14 := []int{fi11, 9223372036854775807, (fi12 % ((fi12) | 1)), (fi12 >> ((65535) & 63))}
                p1 = (-(p1))
                fmt.Println(fi11, v8, fi11)
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS1{F0: fi11}
                p0 = (p0 + ("" + ""))
                switch (fi15 % (((fi15 * fi15)) | 1)) {
                case -28:
                    if (247 <= 255) {
                        v8 += v8
                        fmt.Println(fi11)
                        p0 = p0
                    }
                case 32768:
                    p1 += (p1 & p1)
                case -1:
                    if (2 >= 1) {
                        var v17 uint32 = 2147483648
                        v18 := []int{(-(fi11)), (fi11 | 1), fi15, int(v17)}
                        var v19 string = p0
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                default:
                    p1 = (p1 * (p1 / ((p1) | 1)))
                }
                _ = fi15
                _ = v16
            }
            var v20 string = "hello"
            _ = fi11
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ((1e10 / 0.1) > 0.0)
}
func genFn1(p0 string, p1 uint, p2 uint) int16 {
    if true {
        for fi21 := 0; fi21 < 8; fi21++ {
            fmt.Println(p2, p1)
            _ = fi21
        }
    }
    var v22 uint16 = 0
    p2 -= (p1 >> ((p2) & 63))
    var v23 uint = (p1 ^ (p1 & (p1 - p2)))
    if (false && true) {
        var v24 int32 = 1
        v25 := []int{-661, 32768}
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 1; fi26++ {
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 string = ("a" + "foo bar")
                _ = fi27
                _ = v28
            }
            var v29 int64 = 9223372036854775807
            v30 := GenS1{F0: fi26}
            _ = fi26
            _ = v29
            _ = v30
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 32767
}

func main() {
    v1 := []uint32{}
    var v2 int16 = -737
    v3 := []int{506}
    var v4 uint32 = 32
    var v5 uint8 = 127
    v6 := []uint32{(v4 % (((v4 + v4)) | 1))}
    for fi31 := 0; fi31 < 5; fi31++ {
        fmt.Println(v2, v4)
        v4 = (v4 >> ((2) & 63))
        _ = fi31
    }
    var v32 int32 = 42
    if len(v1) > 0 {
        v1[0] = v4
    }
    v4 -= 4294967295
    var v33 int8 = 0
    if (!((!(65535 <= 257)) && false)) {
        if ((v33 * (v33 | (v33 + v33))) == -80) {
            for fi34 := 0; fi34 < 3; fi34++ {
                v33 = (^(v33))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v3) > 0 {
                    v3[0] = fi34
                }
                _ = fi34
            }
            v35 := GenS2{F0: true, F1: 60, F2: -2147483648}
            v36 := GenS1{F0: 42}
            _ = v35
            _ = v36
        } else {
            fmt.Println(v5, v5, v32)
            var v37 string = "a"
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        v39 := GenS1{F0: 4294967295}
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 uint64 = 62
            _ = fi40
            _ = v41
        }
        var v42 uint = 260
        _ = fi38
        _ = v39
        _ = v42
    }
    var v43 int16 = genFn1("\n\t", 18446744073709551615, 2)
    var v44 bool = false
    if v44 {
        fmt.Println(v44, v5)
        var v45 bool = (false && v44)
        fmt.Println(v33, v32)
        _ = v45
    }
    if (v44 && false) {
        v46 := GenS1{F0: 1}
        fmt.Println(v2, v32, v5)
        for fi47 := 0; fi47 < 11; fi47++ {
            switch ("a" + "foo bar") {
            case "hello":
                fmt.Println(v33)
            default:
                var v48 float64 = -2.25
                _ = v48
            }
            fmt.Println(v44, fi47)
            if v44 {
                _ = len(v6)
                switch "foo bar" {
                case "foo bar":
                    var v49 uint = 18446744073709551615
                    _ = v49
                case "\n\t":
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        var v50 float64 = (1.5 - -2.25)
                        _ = v50
                    default:
                        var v51 float32 = 1e10
                        _ = v51
                    }
                case "a":
                    v33 = (v33 + v33)
                default:
                    v52 := GenS3{F0: "", F1: -31}
                    _ = v52
                }
                switch "a" {
                case "":
                    switch "foo bar" {
                    case "a":
                        var v53 uint32 = (v4 & (^(v4)))
                        _ = v53
                    default:
                        var v54 int32 = v32
                        _ = v54
                    }
                case "\n\t":
                    if len(v6) > 0 {
                        v6[0] = v4
                    }
                case "foo bar":
                    var v55 uint32 = 1
                    _ = v55
                }
            } else {
                var v56 uint8 = (v5 - 127)
                v56 = (v5 * v56)
                var v57 uint8 = v56
                _ = v56
                _ = v57
            }
            _ = fi47
        }
        _ = v46
    } else {
        var v58 bool = v44
        _ = v58
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
