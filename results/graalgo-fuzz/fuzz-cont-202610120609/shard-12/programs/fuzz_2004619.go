package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    v1 := []string{"", "foo bar"}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    v4 := GenS1{F0: "", F1: "hello"}
    var v5 uint64 = 1
    for fi6 := 0; fi6 < 17; fi6++ {
        switch "a" {
        case "":
            _ = len(v1)
        case "foo bar":
            switch 32767 {
            case 532:
                _ = len(v1)
            case -836:
                v7 := GenS1{F0: "\n\t", F1: ""}
                _ = v7
            case 482:
                if (255 >= 248) {
                    v8 := GenS1{F0: ((("hello" + ("foo bar" + "")) + "\n\t") + (("a" + (((((("foo bar" + ("hello" + ("\n\t" + "\n\t"))) + "") + "\n\t") + "a") + "\n\t") + "hello")) + ("foo bar" + "hello"))), F1: "\n\t"}
                    v2[(fi6 + fi6)] = (128 < 2)
                    if false {
                        for fi9 := 0; fi9 < 20; fi9++ {
                            var v10 int32 = -2147483648
                            _ = fi9
                            _ = v10
                        }
                    }
                    _ = v8
                }
            }
        case "\n\t":
            v1 = append(v1, "hello")
        default:
            _ = len(v2)
        }
        _ = fi6
    }
    v11 := GenS1{F0: "\n\t", F1: "foo bar"}
    switch 128 {
    case 9223372036854775807:
        v12 := GenS1{F0: "hello", F1: ("" + "foo bar")}
        _ = v12
    default:
        var v13 float64 = 1e10
        _ = v13
    }
    v5 = 1
    v14 := GenS1{F0: "\n\t", F1: "a"}
    v5 += 42
    v15 := GenS1{F0: "foo bar", F1: "\n\t"}
    var v16 int8 = 0
    if false {
        v17 := GenS1{F0: "foo bar", F1: (((("hello" + "a") + "hello") + "hello") + ("\n\t" + ""))}
        for fi18 := 0; fi18 < 13; fi18++ {
            if false {
                if (int64(v16) <= 1) {
                    fmt.Println(v5, fi18, v5)
                    if false {
                        fmt.Println(v16, v16, v16)
                    } else {
                        if (2 != 795) {
                            fmt.Println(v5, v16, v16)
                        } else {
                            for fi19 := 0; fi19 < 13; fi19++ {
                                v5 = v5
                                v20 := v2[1]
                                if false {
                                    v20 = (384 == 242)
                                }
                                fmt.Println(fi18)
                                _ = fi19
                                _ = v20
                            }
                        }
                    }
                    fmt.Println(v16, fi18)
                } else {
                    _ = len(v3)
                    fmt.Println(v16, fi18)
                    var v21 bool = (!("foo bar" > "\n\t"))
                    _ = v21
                }
                v22 := GenS1{F0: ("\n\t" + "a"), F1: "a"}
                v23 := GenS1{F0: "\n\t", F1: "foo bar"}
                _ = v22
                _ = v23
            }
            v24 := GenS1{F0: "", F1: "hello"}
            if false {
                fmt.Println(fi18, v16, v5)
            } else {
                var v25 uint = 2
                _ = v25
            }
            _ = fi18
            _ = v24
        }
        _ = v17
    }
    if (v5 <= v5) {
        v26 := GenS1{F0: "\n\t", F1: "a"}
        _ = v26
    }
    if true {
        v5 -= (v5 & (v5 & v5))
    }
    if false {
        for fi27 := 0; fi27 < 10; fi27++ {
            v16 -= v16
            v28 := GenS1{F0: "a", F1: ("a" + "hello")}
            _ = fi27
            _ = v28
        }
        v29 := GenS1{F0: "hello", F1: (("\n\t" + "") + "\n\t")}
        v5 = 42
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
