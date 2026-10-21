package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
func genFn0() uint32 {
    var v7 bool = true
    var v8 uint16 = 965
    v7 = (2 > 42)
    if true {
        fmt.Println(v8, v8)
        v9 := GenS1{F0: 1, F1: 2}
        v10 := []int{277, -284, 4294967295, -22}
        _ = v10
        _ = v9
    } else {
        var v11 uint16 = 584
        _ = v11
    }
    v7 = (-756 >= -1)
    _ = v7
    _ = v8
    return 276
}
func genFn1(p0 int16) uint64 {
    if true {
        v12 := GenS1{F0: 1, F1: 42}
        var v13 int8 = 78
        var v14 int = 32767
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        p0 = (p0 | p0)
        if false {
            v16 := []int{fi15, (fi15 << ((601) & 63))}
            p0 += -32768
            _ = v16
        } else {
            var v17 uint = 2
            switch "hello" {
            case "a":
                v18 := []int{1, fi15, 275, (-(fi15))}
                _ = v18
            case "foo bar":
                v19 := GenS1{F0: 4294967295, F1: v17}
                _ = v19
            default:
                if (false && true) {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        var v21 string = ""
                        switch (-(fi20)) {
                        case -973:
                            switch v21 {
                            case "":
                                v21 = v21
                                var v22 float64 = 1.5
                                fmt.Println(v17)
                                _ = v22
                            default:
                                fmt.Println(p0, p0, fi15)
                            }
                        case 503:
                            var v23 int = fi20
                            _ = v23
                        }
                        _ = fi20
                        _ = v21
                    }
                    switch 9223372036854775807 {
                    case -2147483648:
                        for fi24 := 0; fi24 < 7; fi24++ {
                            v25 := GenS1{F0: 2, F1: (v17 / ((4294967295) | 1))}
                            v26 := GenS1{F0: 4294967295, F1: 0}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    case -1:
                        v27 := GenS2{F0: 903, F1: 4294967295}
                        _ = v27
                    case 2147483647:
                        v28 := GenS3{F0: 65535}
                        _ = v28
                    }
                    if true {
                        var v29 uint = (^(v17))
                        fmt.Println(fi15, p0)
                        _ = v29
                    } else {
                        v30 := []int{(fi15 % ((fi15) | 1)), -2147483648, (fi15 - 981), 955}
                        if false {
                            for fi31 := 0; fi31 < 19; fi31++ {
                                switch (("a" + ("a" + ("foo bar" + ("\n\t" + "a")))) + "foo bar") {
                                case "foo bar":
                                    var v32 uint64 = 42
                                    _ = v32
                                case "\n\t":
                                    fmt.Println(fi31, v17)
                                }
                                v17 = v17
                                _ = fi31
                            }
                            v33 := GenS4{F0: 18446744073709551615, F1: ("" + "a"), F2: -329}
                            v34 := GenS2{F0: -188, F1: 42}
                            _ = v33
                            _ = v34
                        }
                        _ = v30
                    }
                }
            }
            _ = v17
        }
        _ = fi15
    }
    var v35 bool = true
    _ = p0
    _ = v35
    return 2
}

func main() {
    var v1 int32 = -676
    var v2 int64 = 42
    var v3 int16 = 775
    var v4 uint8 = 127
    var v5 string = "a"
    var v6 int = 265
    var v36 uint64 = 4294967295
    v36 += v36
    var v37 int64 = (v2 / ((-1) | 1))
    v38 := []int{42, (v6 | 461), (v6 + v6), int(v4)}
    _ = len(v38)
    if true {
        switch v6 {
        case 127:
            v37 -= -1
        case -9223372036854775808:
            var v39 int = int(v3)
            _ = v39
        default:
            var v40 uint16 = 2
            _ = v40
        }
        v1 = 2147483647
    }
    var v41 uint64 = (v36 << ((65535) & 63))
    if true {
        _ = len(v38)
        v41 += 2
        var v42 uint64 = 0
        _ = v42
    } else {
        if (!((v3 * v3) >= v3)) {
            var v43 uint64 = 732
            _ = v43
        }
        switch 32768 {
        case 42:
            if len(v38) >= 2 {
                v38 = v38[1:len(v38)]
            }
        default:
            v41 += 18446744073709551615
        }
        fmt.Println(v6, v4, v2)
    }
    fmt.Println(v1, v36)
    v36 += (^(v36))
    for fi44 := 0; fi44 < 20; fi44++ {
        var v45 int64 = (v37 | v37)
        var v46 int8 = 127
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
