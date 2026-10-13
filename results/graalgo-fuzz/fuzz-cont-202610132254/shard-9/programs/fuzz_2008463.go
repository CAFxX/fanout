package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    var v1 int32 = 1
    var v2 uint8 = 1
    v3 := []string{"foo bar", ("\n\t" + "foo bar"), "foo bar", "\n\t"}
    var v4 int64 = -1
    var v5 int8 = -10
    v6 := GenS1{F0: 2, F1: ((-386 != 32767) || (1 > 4294967295)), F2: 1}
    v5 -= -128
    for fi7 := 0; fi7 < 1; fi7++ {
        if (int64(v5) == (^(v4))) {
            var v8 int32 = v1
            _ = v8
        } else {
            for fi9 := 0; fi9 < 13; fi9++ {
                v1 = v1
                v5 = int8(fi9)
                v3 = append(v3, (("hello" + ("" + "a")) + ""))
                var v10 uint32 = 606
                _ = fi9
                _ = v10
            }
            if (255 >= 2) {
                v2 -= v2
            } else {
                if (v5 == v5) {
                    v1 = int32(v4)
                    _ = len(v3)
                    if false {
                        var v11 int64 = -715
                        v12 := GenS1{F0: 0, F1: (v11 < (^(v4))), F2: 603}
                        _ = v11
                        _ = v12
                    }
                }
                fmt.Println(fi7, v4, v2)
                for fi13 := 0; fi13 < 15; fi13++ {
                    _ = len(v3)
                    var v14 int64 = v4
                    v2 += 0
                    _ = fi13
                    _ = v14
                }
            }
        }
        v3 = append(v3, "a")
        v5 = (v5 << ((0) & 63))
        _ = fi7
    }
    v15 := GenS1{F0: 42, F1: (v1 < (v1 & (-(v1)))), F2: 917}
    switch "a" {
    case "":
        v3 = append(v3, "foo bar")
    case "a":
        v16 := GenS2{F0: true, F1: (v4 | v4)}
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 3; fi17++ {
            v4 -= v4
            for fi18 := 0; fi18 < 9; fi18++ {
                v5 = v5
                if (!(true && false)) {
                    v1 -= (v1 / ((-2147483648) | 1))
                    v2 = 128
                    v2 = 2
                }
                _ = fi18
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                if false {
                    fmt.Println(v2)
                } else {
                    switch 9223372036854775807 {
                    case 42:
                        if false {
                            v20 := GenS1{F0: 1, F1: (1 <= 1), F2: 4294967295}
                            _ = v20
                        } else {
                            v4 += v4
                            v3 = append(v3, "foo bar")
                        }
                    case 9223372036854775807:
                        if (1.5 < 1e10) {
                            var v21 int = fi17
                            _ = v21
                        }
                    case 128:
                        fmt.Println(v4, v5)
                    default:
                        var v22 int64 = v4
                        _ = v22
                    }
                    if false {
                        fmt.Println(v4, v1)
                        fmt.Println(v2, v4, fi19)
                    } else {
                        fmt.Println(v5)
                        var v23 string = ("foo bar" + "foo bar")
                        _ = v23
                    }
                    v5 = 52
                    v1 -= 1
                }
                _ = fi19
            }
            _ = fi17
        }
    }
    _ = len(v3)
    var v24 uint32 = 716
    if false {
        v25 := GenS1{F0: 1, F1: (false && false), F2: 0}
        var v26 float64 = 0.0
        _ = v25
        _ = v26
    } else {
        var v27 bool = false
        _ = v27
    }
    v5 += (v5 >> ((50) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
