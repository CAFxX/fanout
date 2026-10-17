package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint64) uint8 {
    if (!false) {
        var v6 int64 = -843
        _ = v6
    } else {
        p0 -= (p0 << ((18446744073709551615) & 63))
        fmt.Println(p0)
    }
    if false {
        p0 += p0
    }
    var v7 bool = (1 < 407)
    _ = p0
    _ = v7
    return uint8(p0)
}
func genFn1(p0 bool, p1 uint64) uint8 {
    v8 := GenS1{F0: 32767, F1: "\n\t"}
    fmt.Println(p0, p0)
    p0 = true
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{("a" + (((("foo bar" + (("" + "") + "\n\t")) + "hello") + ("\n\t" + ("a" + ("hello" + ("a" + "\n\t"))))) + "foo bar")), ""}
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    var v5 uint16 = 32767
    var v9 int8 = 38
    v5 = (v5 & (v5 >> ((65535) & 63)))
    fmt.Println(v5, v5)
    v5 = v5
    v9 -= (v9 << ((65535) & 63))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    delete(v4, "a")
    if (181 >= 127) {
        v10 := v1["hello"]
        v9 = v9
        v11 := GenS1{F0: 0, F1: "\n\t"}
        _ = v10
        _ = v11
    } else {
        var v12 uint16 = 32767
        for fi13 := 0; fi13 < 5; fi13++ {
            fmt.Println(v5, v9)
            fmt.Println(v9, v9)
            fmt.Println(v12, v9, v12)
            _ = fi13
        }
        switch 9223372036854775807 {
        case 634:
            fmt.Println(v9)
        }
        _ = v12
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if (!false) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, ("foo bar" + "foo bar"))
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 float64 = -2.25
                _ = fi15
                _ = v16
            }
        } else {
            switch "hello" {
            case "\n\t":
                var v17 int = fi14
                _ = v17
            case "":
                var v18 int8 = -36
                _ = v18
            }
            v9 = v9
        }
        _ = fi14
    }
    v19 := v1[(("" + "\n\t") + "\n\t")]
    v5 += (v5 ^ v5)
    v5 = (^(v5))
    v5 -= 318
    v20 := GenS2{F0: (4294967295 == 2147483647), F1: 255}
    if false {
        v21 := GenS3{F0: ((v19 + "hello") + ((v19 + v19) + "foo bar")), F1: (!((-1 != -1) && (false || ((!false) || false))))}
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 string = (v19 + "hello")
            v24 := GenS4{F0: -674, F1: (^(v9))}
            for fi25 := 0; fi25 < 6; fi25++ {
                if (((127 < 116) || true) || (!(!(false && true)))) {
                    var v26 float32 = 0.1
                    v27 := GenS2{F0: true, F1: 12}
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS5{F0: v5}
                    var v29 int32 = 693
                    switch fi25 {
                    case -1:
                        if (!(-2.25 != 0.0)) {
                            delete(v1, v23)
                        }
                    case -578:
                        v19 = "hello"
                    case 42:
                        var v30 int16 = -32768
                        _ = v30
                    default:
                        fmt.Println(v23, v5, v19)
                    }
                    v2 = append(v2, v19)
                    _ = v28
                    _ = v29
                }
                _ = fi25
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        v3[((v19 + "\n\t") + (("a" + v19) + "hello"))] = -9223372036854775808
        _ = v21
    } else {
        var v31 uint64 = 931
        _ = v31
    }
    v32 := GenS3{F0: v19, F1: (!false)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
