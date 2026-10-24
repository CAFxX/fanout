package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint32 }
func genFn0(p0 uint64, p1 bool) uint8 {
    v6 := GenS1{F0: 750, F1: 32767}
    p0 -= 349
    for fi7 := 0; fi7 < 4; fi7++ {
        p1 = p1
        if p1 {
            var v8 uint = 705
            _ = v8
        }
        _ = fi7
    }
    var v9 int32 = 668
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return 255
}

func main() {
    var v1 int8 = -49
    v2 := []uint32{725, 2147483647, 2147483647, 4294967295}
    var v3 bool = (true && false)
    var v4 uint8 = 0
    v5 := make(map[uint32]bool)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "" {
    case "foo bar":
        v3 = (!(42 < 42))
    case "\n\t":
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v3)
            v4 += (^(v4))
            _ = fi10
        }
    case "a":
        for fi11 := 0; fi11 < 14; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                v5[457] = (0.1 == (3.14159 / 1.5))
                _ = fi12
            }
            _ = fi11
        }
    }
    var v13 int32 = 1
    if (!(!(uint64(v4) != 42))) {
        v14 := v5[8]
        switch -2147483648 {
        case 445:
            if v14 {
                v4 += v4
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 float32 = 0.0
                    if (!(true && (!(673 != 828)))) {
                        for fi17 := 0; fi17 < 17; fi17++ {
                            fmt.Println(v13, v4)
                            v18 := GenS1{F0: 0, F1: 1}
                            _ = fi17
                            _ = v18
                        }
                        var v19 uint16 = 668
                        var v20 uint32 = 65535
                        _ = v19
                        _ = v20
                    }
                    _ = fi15
                    _ = v16
                }
                switch "hello" {
                case "":
                    if v14 {
                        v1 -= 0
                        fmt.Println(v4)
                    } else {
                        if ((18446744073709551615 == 788) || v3) {
                            fmt.Println(v3, v4, v4)
                        } else {
                            v5[4294967295] = false
                            v3 = v3
                        }
                    }
                }
            } else {
                v3 = ((v13 + (v13 | v13)) < v13)
            }
        case -9223372036854775808:
            switch ("a" + "foo bar") {
            case "foo bar":
                v3 = false
            default:
                v13 -= v13
            }
        case -1:
            var v21 uint64 = 0
            _ = v21
        default:
            _ = len(v2)
        }
        if true {
            v3 = ("" <= ("" + ("\n\t" + "a")))
            var v22 uint8 = (v4 + v4)
            _ = v22
        } else {
            if v14 {
                var v23 int = 9223372036854775807
                _ = v23
            } else {
                v24 := GenS2{F0: 32767, F1: -937}
                for fi25 := 0; fi25 < 12; fi25++ {
                    if (-1 >= (v13 / (((v13 ^ (v13 << ((2) & 63)))) | 1))) {
                        var v26 uint64 = 642
                        v14 = true
                        v27 := GenS3{F0: 18446744073709551615}
                        _ = v26
                        _ = v27
                    } else {
                        var v28 uint = 72
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v13 -= 42
                        _ = fi29
                    }
                    fmt.Println(v4, fi25, v1)
                    _ = fi25
                }
                if v3 {
                    var v30 float32 = 0.0
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 1; fi31++ {
                        var v32 int8 = -32
                        _ = fi31
                        _ = v32
                    }
                    var v33 uint64 = 760
                    var v34 int32 = -878
                    _ = v33
                    _ = v34
                }
                _ = v24
            }
            v3 = v14
            for fi35 := 0; fi35 < 7; fi35++ {
                fmt.Println(v3, v4)
                _ = fi35
            }
        }
        _ = v14
    } else {
        v36 := GenS1{F0: v13, F1: 159}
        _ = v36
    }
    _ = len(v2)
    var v37 uint8 = 255
    if (1 >= 255) {
        var v38 int16 = 1000
        v13 += (-(v13))
        var v39 uint64 = 18446744073709551615
        _ = v38
        _ = v39
    } else {
        if (v4 >= (^(v4))) {
            switch -297 {
            case 127:
                for fi40 := 0; fi40 < 19; fi40++ {
                    v41 := GenS1{F0: (v13 * (^(v13))), F1: 32767}
                    var v42 int = (fi40 - (fi40 >> ((42) & 63)))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            case 9223372036854775807:
                fmt.Println(v13, v37)
            case -359:
                _ = len(v5)
            default:
                var v43 uint = 255
                _ = v43
            }
        } else {
            v1 = (v1 ^ v1)
        }
        v44 := GenS3{F0: 4294967295}
        _ = len(v2)
        _ = v44
    }
    if false {
        v37 -= v37
        if (v37 >= (^(v37))) {
            v3 = ((128 <= 563) || v3)
            switch "hello" {
            case "foo bar":
                v1 -= -1
            case "hello":
                if v3 {
                    switch "\n\t" {
                    case "":
                        v2 = append(v2, 4294967295)
                    default:
                        var v45 string = "a"
                        _ = v45
                    }
                    v3 = v3
                }
            }
        } else {
            v46 := GenS3{F0: 0}
            if (!v3) {
                var v47 string = (("" + "foo bar") + "\n\t")
                _ = v47
            }
            var v48 bool = (1.5 >= -2.25)
            _ = v46
            _ = v48
        }
        delete(v5, 733)
    } else {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        var v49 bool = v3
        v50 := GenS4{F0: 1000, F1: 16, F2: 1}
        _ = v49
        _ = v50
    }
    fmt.Println(v37, v1, v37)
    switch ("\n\t" + "") {
    case "":
        v13 = v13
    default:
        fmt.Println(v3, v3, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
