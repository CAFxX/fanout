package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int8 }
func genFn0(p0 bool) uint {
    var v6 uint8 = 127
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint8, p1 int8) string {
    if false {
        fmt.Println(p0)
        var v7 string = ("a" + "\n\t")
        _ = v7
    }
    v8 := GenS1{F0: int32(p1), F1: 2147483648, F2: 683}
    fmt.Println(p1, p0)
    var v9 uint16 = 478
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    v4 := []int{128, 1}
    var v5 bool = true
    if true {
        fmt.Println(v5, v1)
    } else {
        fmt.Println(v5)
        for fi10 := 0; fi10 < 17; fi10++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v5 = ((!(2147483647 < 65535)) && (-9223372036854775808 != -1))
            for fi11 := 0; fi11 < 3; fi11++ {
                v1 = v1
                _ = fi11
            }
            _ = fi10
        }
    }
    v5 = (true && false)
    if v5 {
        switch 32768 {
        case 32767:
            for fi12 := 0; fi12 < 20; fi12++ {
                v1 = (^(v1))
                if (1e10 > -2.25) {
                    v5 = v5
                }
                v5 = (v5 || true)
                _ = fi12
            }
        default:
            for fi13 := 0; fi13 < 5; fi13++ {
                v5 = (-9223372036854775808 > -9223372036854775808)
                v5 = (v5 && (true && ((1 >= -58) || false)))
                _ = fi13
            }
        }
        if ((18446744073709551615 != 18446744073709551615) && (!false)) {
            var v14 bool = v5
            if v5 {
                for fi15 := 0; fi15 < 17; fi15++ {
                    fmt.Println(fi15, v14)
                    switch fi15 {
                    case 937:
                        switch ("" + ("foo bar" + "")) {
                        case "hello":
                            v16 := GenS1{F0: -2147483648, F1: 0, F2: 115}
                            _ = v16
                        case "\n\t":
                            delete(v3, 2147483647)
                        case "":
                            var v17 string = "hello"
                            _ = v17
                        default:
                            fmt.Println(v1)
                        }
                    case -9223372036854775808:
                        v3[2147483647] = fi15
                    case 9223372036854775807:
                        v1 -= v1
                    }
                    delete(v2, fi15)
                    _ = fi15
                }
                if (957 < 1) {
                    var v18 uint64 = 454
                    _ = v18
                } else {
                    v19 := GenS2{F0: -1}
                    var v20 int8 = 42
                    _ = v19
                    _ = v20
                }
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v21 int64 = 1
                _ = v21
            }
            switch "a" {
            case "foo bar":
                fmt.Println(v5)
            case "a":
                if false {
                    var v22 int8 = 28
                    _ = v22
                }
            case "hello":
                var v23 int = 128
                _ = v23
            default:
                v24 := GenS3{F0: 0}
                _ = v24
            }
            _ = v14
        }
    } else {
        for fi25 := 0; fi25 < 7; fi25++ {
            v26 := GenS1{F0: -2147483648, F1: 2147483647, F2: -1}
            for fi27 := 0; fi27 < 7; fi27++ {
                if (171 == 0) {
                    v28 := GenS2{F0: 1}
                    var v29 uint = v1
                    _ = v28
                    _ = v29
                }
                var v30 int64 = 42
                _ = fi27
                _ = v30
            }
            _ = fi25
            _ = v26
        }
        var v31 int8 = -1
        _ = v31
    }
    var v32 int16 = 42
    if v5 {
        switch "foo bar" {
        case "foo bar":
            v2[32768] = 1
        case "a":
            for fi33 := 0; fi33 < 18; fi33++ {
                v32 = v32
                _ = fi33
            }
        }
    }
    switch ("foo bar" + "a") {
    case "hello":
        if ((0.1 + 1e10) != 1.5) {
            fmt.Println(v32, v5)
            for fi34 := 0; fi34 < 4; fi34++ {
                for fi35 := 0; fi35 < 18; fi35++ {
                    var v36 uint32 = 2147483647
                    _ = fi35
                    _ = v36
                }
                fmt.Println(fi34)
                _ = fi34
            }
        }
    case "\n\t":
        switch 997 {
        case -218:
            var v37 int = int(v1)
            _ = v37
        case -773:
            for fi38 := 0; fi38 < 2; fi38++ {
                v1 = (^(v1))
                for fi39 := 0; fi39 < 7; fi39++ {
                    if v5 {
                        var v40 uint = (v1 | (v1 * 255))
                        v5 = (1 != 24)
                        var v41 int16 = (v32 + v32)
                        _ = v40
                        _ = v41
                    }
                    var v42 int32 = 909
                    _ = fi39
                    _ = v42
                }
                _ = fi38
            }
        default:
            if v5 {
                for fi43 := 0; fi43 < 11; fi43++ {
                    for fi44 := 0; fi44 < 20; fi44++ {
                        var v45 uint32 = 2
                        _ = fi44
                        _ = v45
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi43
                }
                v46 := GenS3{F0: -35}
                v5 = (v5 && true)
                _ = v46
            } else {
                var v47 int = 577
                fmt.Println(v1, v47)
                v48 := GenS4{F0: v1, F1: 0, F2: 157}
                _ = v47
                _ = v48
            }
        }
    default:
        for fi49 := 0; fi49 < 9; fi49++ {
            v50 := GenS1{F0: 2147483647, F1: 758, F2: -1}
            v32 -= (^(v32))
            _ = fi49
            _ = v50
        }
    }
    v51 := GenS5{F0: 2147483648, F1: 1, F2: 127}
    for fi52 := 0; fi52 < 15; fi52++ {
        if (!(4294967295 == 2)) {
            v5 = v5
        } else {
            var v53 int8 = -43
            v1 = 1
            v2[fi52] = 1
            _ = v53
        }
        for fi54 := 0; fi54 < 6; fi54++ {
            v5 = (!v5)
            switch fi52 {
            case -966:
                v2[(fi52 << ((42) & 63))] = 149
            default:
                fmt.Println(fi54, fi52)
            }
            v1 -= (v1 ^ (v1 & v1))
            var v55 float64 = (((1.5 * 0.0) / 1.0) + ((3.14159 + -2.25) - float64(-2.25)))
            _ = fi54
            _ = v55
        }
        _ = fi52
    }
    fmt.Println(v5)
    fmt.Println(v32, v32, v1)
    fmt.Println(v1, v32)
    var v56 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v51
    _ = v56
}
