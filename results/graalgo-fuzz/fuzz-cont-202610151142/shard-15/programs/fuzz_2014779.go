package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 string; F1 int64; F2 string }
type GenS3 struct { F0 int16; F1 int32; F2 string }
type GenS4 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS5 struct { F0 int; F1 int; F2 uint }
type GenS6 struct { F0 uint16; F1 bool }
func genFn0(p0 uint16) uint64 {
    fmt.Println(p0)
    var v5 uint64 = 18446744073709551615
    _ = p0
    _ = v5
    return 18446744073709551615
}
func genFn1() uint32 {
    if false {
        v6 := GenS1{F0: "a", F1: 9223372036854775807}
        _ = v6
    }
    v7 := GenS1{F0: ("hello" + ("\n\t" + "")), F1: 0}
    var v8 int8 = 127
    v8 += (v8 * (-(v8)))
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []string{"a", (("foo bar" + ("" + "")) + "\n\t")}
    var v2 uint32 = 560
    v3 := []string{(("" + ("foo bar" + "hello")) + "foo bar"), "", ("\n\t" + "")}
    var v4 int8 = -1
    v2 = (v2 << ((940) & 63))
    var v9 string = (("" + "\n\t") + "hello")
    v1 = append(v1, "\n\t")
    if len(v1) > 0 {
        v1[0] = v9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v10 int32 = -2147483648
    var v11 bool = true
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v12 uint16 = 539
    fmt.Println(v12, v2, v12)
    if (-1 != 0) {
        var v13 uint16 = v12
        var v14 uint64 = 0
        fmt.Println(v12, v14)
        _ = v13
        _ = v14
    } else {
        switch 850 {
        case 128:
            _ = len(v3)
        case 127:
            v2 = 2147483647
        case 32767:
            fmt.Println(v2, v10, v9)
        default:
            if true {
                var v15 int16 = 835
                var v16 uint = 2
                _ = v15
                _ = v16
            }
        }
        v17 := GenS2{F0: (v9 + ((v9 + "\n\t") + (v9 + ""))), F1: 0, F2: "a"}
        _ = v17
    }
    switch 240 {
    case 32768:
        switch 42 {
        case 9223372036854775807:
            fmt.Println(v9, v9, v12)
        case -1:
            _ = len(v3)
        case 2147483647:
            v9 = ((v9 + v9) + v9)
        default:
            for fi18 := 0; fi18 < 11; fi18++ {
                for fi19 := 0; fi19 < 4; fi19++ {
                    v4 = (v4 % ((int8(v12)) | 1))
                    var v20 int32 = 0
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        }
    default:
        if v11 {
            v21 := GenS1{F0: "", F1: 2147483647}
            if len(v1) > 0 {
                v1[0] = ((v9 + v9) + "")
            }
            switch -531 {
            case 42:
                v22 := GenS1{F0: "\n\t", F1: -2147483648}
                _ = v22
            case 32767:
                var v23 int16 = 219
                _ = v23
            case 9223372036854775807:
                v11 = true
            default:
                var v24 uint8 = 59
                _ = v24
            }
            _ = v21
        } else {
            if true {
                if v11 {
                    if (!(!true)) {
                        v2 = v2
                    }
                    switch "a" {
                    case "hello":
                        var v25 uint = 2
                        _ = v25
                    default:
                        if (-2.25 < 3.14159) {
                            var v26 uint32 = (v2 - 2)
                            _ = v26
                        }
                    }
                    v11 = v11
                } else {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v28 := GenS3{F0: -1, F1: (-(v10)), F2: (v9 + "a")}
                        var v29 uint8 = 101
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v4 = 0
                        v10 = v10
                        if (!false) {
                            v31 := GenS4{F0: (v2 % (((v2 - 4294967295)) | 1)), F1: 0, F2: 0}
                            _ = v31
                        } else {
                            v32 := GenS5{F0: 0, F1: fi30, F2: 611}
                            _ = v32
                        }
                        _ = fi30
                    }
                }
                for fi33 := 0; fi33 < 14; fi33++ {
                    v2 -= v2
                    _ = fi33
                }
                switch ("" + "foo bar") {
                case "a":
                    for fi34 := 0; fi34 < 12; fi34++ {
                        fmt.Println(v9, v2)
                        for fi35 := 0; fi35 < 19; fi35++ {
                            v10 = int32(v4)
                            _ = fi35
                        }
                        _ = fi34
                    }
                }
            } else {
                for fi36 := 0; fi36 < 14; fi36++ {
                    _ = len(v3)
                    v37 := GenS6{F0: (v12 * v12), F1: (42 <= 24)}
                    _ = fi36
                    _ = v37
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
