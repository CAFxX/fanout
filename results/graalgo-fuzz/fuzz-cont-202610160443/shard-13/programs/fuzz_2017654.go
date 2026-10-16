package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    var v1 string = (("\n\t" + "foo bar") + ("a" + ("\n\t" + ("foo bar" + "foo bar"))))
    v2 := []string{"\n\t", ((v1 + (v1 + "hello")) + v1), "\n\t"}
    var v3 string = "a"
    v4 := make(map[uint32]string)
    var v5 int16 = 0
    v6 := []string{"hello", ""}
    v7 := GenS1{F0: v1, F1: (("" + ((v3 + v1) + (v3 + ((v3 + ("" + "\n\t")) + v3)))) + "")}
    for fi8 := 0; fi8 < 20; fi8++ {
        v5 = v5
        _ = fi8
    }
    switch v3 {
    case "\n\t":
        v1 = v1
    }
    v5 = v5
    v9 := v4[110]
    switch 32767 {
    case -555:
        var v10 uint32 = 65535
        _ = v10
    }
    v1 = ("" + "a")
    if (!(!(683 >= -846))) {
        var v11 uint8 = 255
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int16 = 32767
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        fmt.Println(v1, v1)
        switch 341 {
        case 128:
            var v14 int = 9223372036854775807
            _ = v14
        default:
            switch v3 {
            case "\n\t":
                for fi15 := 0; fi15 < 5; fi15++ {
                    v6 = append(v6, "a")
                    v16 := v4[1]
                    _ = fi15
                    _ = v16
                }
            case "":
                switch v3 {
                case "a":
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v1 = v1
                        _ = fi17
                    }
                case "":
                    v18 := GenS1{F0: (("hello" + (v1 + "")) + ("a" + v1)), F1: v1}
                    _ = v18
                default:
                    fmt.Println(v9, v9, v1)
                }
            default:
                switch -581 {
                case 127:
                    fmt.Println(v3)
                case -1:
                    v19 := GenS1{F0: "", F1: v3}
                    _ = v19
                case -9223372036854775808:
                    v9 = ""
                default:
                    if false {
                        v3 = v9
                        v20 := GenS1{F0: v3, F1: "foo bar"}
                        if (42 != 0) {
                            for fi21 := 0; fi21 < 8; fi21++ {
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    var v23 int64 = int64(fi21)
                                    var v24 bool = true
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                _ = fi21
                            }
                        }
                        _ = v20
                    }
                }
            }
        }
        var v25 uint64 = 42
        _ = v25
    }
    v9 = "a"
    var v26 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
