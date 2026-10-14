package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint8 }
func genFn0(p0 uint32, p1 uint8, p2 uint32) uint16 {
    p0 = (p0 + p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint32 = (^(p0))
        v7 := []int{}
        var v8 int8 = 100
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    switch "" {
    case "hello":
        fmt.Println(p2, p0, p2)
    case "foo bar":
        v9 := GenS1{F0: -9223372036854775808}
        _ = v9
    default:
        fmt.Println(p0, p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []string{(("foo bar" + "hello") + "foo bar"), "hello", ("" + "foo bar"), "\n\t"}
    var v3 float32 = 0.0
    var v4 uint64 = v1
    v1 += v1
    if true {
        for fi10 := 0; fi10 < 18; fi10++ {
            switch 2147483647 {
            case 9223372036854775807:
                switch 803 {
                case 4294967295:
                    for fi11 := 0; fi11 < 10; fi11++ {
                        if true {
                            v1 = v4
                            fmt.Println(v4, v3, v3)
                        }
                        _ = fi11
                    }
                case 128:
                    fmt.Println(v3)
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            default:
                var v12 int16 = 42
                _ = v12
            }
            _ = fi10
        }
        v13 := GenS2{F0: true, F1: 42, F2: -2147483648}
        _ = v13
    }
    if ((772 == 0) || (!(((!(-401 != 42)) || (1 >= -1)) || false))) {
        var v14 int16 = -32768
        var v15 uint = 2
        _ = v14
        _ = v15
    } else {
        v16 := GenS3{F0: 2, F1: 63}
        var v17 int32 = 861
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: true, F1: 42, F2: 536}
    fmt.Println(v1)
    fmt.Println(v1, v4)
    v3 -= v3
    if true {
        v3 += (((3.14159 + 0.1) + v3) * ((0.0 / (v3 / v3)) / (v3 / ((v3 + 1.5) / float32(1e10)))))
    }
    if (0 == 0) {
        v19 := GenS4{F0: 1, F1: (v1 ^ 0), F2: uint8(v4)}
        _ = v19
    } else {
        switch "hello" {
        case "":
            _ = len(v2)
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v3 += v3
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS2{F0: false, F1: 9223372036854775807, F2: fi20}
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v3 = 1.5
    switch 878 {
    case 207:
        v1 -= v4
    case 358:
        switch "a" {
        case "hello":
            if (true && true) {
                var v23 uint8 = 1
                _ = v23
            }
        default:
            var v24 uint32 = 65535
            _ = v24
        }
    default:
        for fi25 := 0; fi25 < 20; fi25++ {
            if true {
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 int = (fi25 * (fi25 + -1))
                    if true {
                        v27 = (fi25 - 127)
                    } else {
                        var v28 int8 = -128
                        v1 -= v4
                        v3 += (v3 * v3)
                        switch "\n\t" {
                        case "\n\t":
                            var v29 uint8 = 1
                            _ = v29
                        case "a":
                            fmt.Println(fi26)
                        case "hello":
                            var v30 uint16 = 0
                            _ = v30
                        default:
                            v28 = 100
                        }
                        _ = v28
                    }
                    _ = len(v2)
                    _ = fi26
                    _ = v27
                }
                var v31 int = fi25
                _ = v31
            } else {
                v4 += (v4 << ((775) & 63))
                var v32 bool = ((true && (!true)) && (!(65535 != 42)))
                if true {
                    if true {
                        var v33 uint8 = 2
                        v34 := GenS4{F0: 457, F1: (v1 >> ((586) & 63)), F2: v33}
                        _ = v33
                        _ = v34
                    }
                    if (v1 <= v4) {
                        v32 = (584 > 1)
                        for fi35 := 0; fi35 < 6; fi35++ {
                            var v36 string = ("" + "")
                            _ = fi35
                            _ = v36
                        }
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                    }
                    for fi37 := 0; fi37 < 13; fi37++ {
                        var v38 int64 = 897
                        v39 := GenS2{F0: v32, F1: (v38 + (v38 % ((9223372036854775807) | 1))), F2: 128}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
                _ = v32
            }
            _ = fi25
        }
    }
    if true {
        v1 = (v1 + v4)
        if len(v2) > 0 {
            v2[0] = ""
        }
    }
    var v40 bool = false
    var v41 uint64 = v4
    var v42 float32 = float32((-2.25 / (0.1 - -2.25)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
