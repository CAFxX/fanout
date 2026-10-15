package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    v1 := []uint32{1, 0, 42}
    var v2 string = "\n\t"
    v3 := []int{1, -934, 127}
    v2 = (("" + v2) + ((v2 + v2) + ((v2 + ("" + "hello")) + v2)))
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int32 = -1
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            v2 = v2
            _ = fi7
        }
        v2 = v2
        for fi8 := 0; fi8 < 8; fi8++ {
            switch fi8 {
            case 768:
                if (!("" <= (v2 + v2))) {
                    v2 = ("" + ((v2 + "foo bar") + (v2 + v2)))
                    _ = len(v3)
                    v2 = (v2 + v2)
                } else {
                    var v9 uint = 65535
                    var v10 int16 = 32767
                    _ = v10
                    _ = v9
                }
            default:
                if (2147483648 >= 596) {
                    v11 := GenS1{F0: -1, F1: 0}
                    if true {
                        switch (fi8 + -2147483648) {
                        case -460:
                            fmt.Println(fi6, fi6)
                        }
                        v1 = append(v1, 65535)
                        v2 = "a"
                        var v12 uint8 = 42
                        switch (v2 + "a") {
                        case "a":
                            var v13 int64 = 0
                            _ = v13
                        case "hello":
                            fmt.Println(v12, v2)
                        case "foo bar":
                            switch (v2 + ("" + v2)) {
                            case "a":
                                var v14 float32 = 0.0
                                _ = v14
                            case "\n\t":
                                switch fi8 {
                                case -8:
                                    fmt.Println(v2)
                                case 42:
                                    var v15 uint32 = 0
                                    _ = v15
                                }
                            case "":
                                switch "a" {
                                case "a":
                                    var v16 uint32 = 663
                                    _ = v16
                                }
                                switch 32768 {
                                case -281:
                                    _ = len(v1)
                                case 1:
                                    var v17 uint16 = 2
                                    _ = v17
                                default:
                                    if ((-(fi8)) != -9223372036854775808) {
                                        fmt.Println(fi8, v12)
                                        v12 = v12
                                        v2 = v2
                                    }
                                }
                            default:
                                var v18 int16 = -933
                                _ = v18
                            }
                        }
                        v19 := GenS1{F0: 1000, F1: 1}
                        v2 = "hello"
                        if true {
                            v20 := GenS2{F0: 0, F1: -766}
                            v2 = v2
                            _ = v20
                        }
                        _ = v12
                        _ = v19
                    } else {
                        var v21 string = ("hello" + "")
                        v3 = append(v3, 32767)
                        if true {
                            var v22 uint64 = 18446744073709551615
                            _ = len(v1)
                            var v23 float32 = -2.25
                            _ = v22
                            _ = v23
                        } else {
                            var v24 uint32 = 2147483648
                            _ = v24
                        }
                        _ = v21
                    }
                    _ = v11
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 42
                }
                fmt.Println(fi6, fi8)
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi8
        }
        _ = fi6
    }
    if (!false) {
        for fi25 := 0; fi25 < 1; fi25++ {
            fmt.Println(v2, v2)
            _ = fi25
        }
        if ((true && (42 <= 43)) || true) {
            v2 = "hello"
        }
        v2 = (((("a" + "foo bar") + v2) + v2) + ("\n\t" + "foo bar"))
    }
    fmt.Println(v2)
    v2 = v2
    v2 = "\n\t"
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
}
