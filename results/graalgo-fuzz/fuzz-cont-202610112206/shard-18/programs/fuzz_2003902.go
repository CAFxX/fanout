package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }

func main() {
    v1 := []string{"hello", "", "foo bar", ("\n\t" + ("a" + "a"))}
    var v2 uint = 42
    var v3 int64 = -1
    v4 := make(map[uint32]string)
    switch "\n\t" {
    case "\n\t":
        if true {
            v1 = append(v1, ("\n\t" + ("foo bar" + "hello")))
            if (!true) {
                for fi5 := 0; fi5 < 14; fi5++ {
                    v3 -= (^(v3))
                    var v6 bool = true
                    var v7 int16 = 0
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            }
        } else {
            _ = len(v1)
            v3 = v3
            switch 4294967295 {
            case -826:
                var v8 int32 = -1
                _ = v8
            case 4294967295:
                v3 = (v3 << ((255) & 63))
            case 128:
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := GenS1{F0: "", F1: 251}
                    fmt.Println(fi9)
                    var v11 int32 = 1
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                v12 := GenS1{F0: "foo bar", F1: -9223372036854775808}
                _ = v12
            }
        }
    default:
        var v13 uint8 = 42
        _ = v13
    }
    var v14 uint64 = 42
    var v15 int16 = 707
    var v16 uint = (v2 + (^(v2)))
    switch ("\n\t" + ("" + ("\n\t" + ("a" + "\n\t")))) {
    case "hello":
        v16 = v16
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            if len(v1) > 0 {
                v1[0] = ("a" + ((("a" + "") + (("hello" + ("" + "hello")) + ("\n\t" + "\n\t"))) + "hello"))
            }
            for fi18 := 0; fi18 < 5; fi18++ {
                if false {
                    var v19 int = fi17
                    var v20 int64 = 1
                    _ = v19
                    _ = v20
                } else {
                    v2 += 506
                    fmt.Println(v16)
                    v21 := GenS1{F0: "hello", F1: v3}
                    _ = v21
                }
                if (v15 < int16(v3)) {
                    if false {
                        fmt.Println(v15)
                        var v22 float32 = (1e10 / 3.14159)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v22
                    }
                    var v23 string = "hello"
                    _ = v23
                }
                var v24 uint32 = 0
                _ = fi18
                _ = v24
            }
            v3 = 0
            _ = fi17
        }
    }
    v14 -= (v14 | (v14 & (v14 * v14)))
    var v25 bool = (("foo bar" + "foo bar") == ("a" + "foo bar"))
    v25 = (v25 || (false && (false && (127 <= 255))))
    for fi26 := 0; fi26 < 11; fi26++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if v25 {
            if (936 == 477) {
                var v27 int32 = -2147483648
                _ = v27
            }
            v28 := GenS1{F0: "a", F1: v3}
            var v29 uint64 = 1
            _ = v28
            _ = v29
        } else {
            var v30 int16 = v15
            _ = v30
        }
        var v31 bool = ((!v25) && ((v16 >> ((v16) & 63)) > 4294967295))
        _ = fi26
        _ = v31
    }
    v3 -= -1
    v2 = 0
    var v32 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v4
}
