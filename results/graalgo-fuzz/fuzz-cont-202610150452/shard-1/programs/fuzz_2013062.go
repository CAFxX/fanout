package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 uint8 = 2
    v2 := []string{"", ("\n\t" + "foo bar"), "\n\t"}
    v3 := []string{"a", "a", ("a" + "")}
    var v4 uint8 = v1
    var v5 int8 = 100
    for fi6 := 0; fi6 < 14; fi6++ {
        switch (fi6 >> ((698) & 63)) {
        case 32768:
            switch "\n\t" {
            case "":
                v7 := GenS1{F0: "", F1: true, F2: 0}
                _ = v7
            case "foo bar":
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    var v8 int16 = 655
                    _ = v8
                }
            case "a":
                for fi9 := 0; fi9 < 8; fi9++ {
                    v4 = v4
                    switch ("" + "a") {
                    case "foo bar":
                        v10 := GenS2{F0: -2147483648}
                        _ = v10
                    case "hello":
                        switch (fi9 << ((42) & 63)) {
                        case 128:
                            var v11 int8 = 95
                            _ = v11
                        case 32768:
                            var v12 uint32 = 787
                            _ = v12
                        default:
                            if true {
                                if (!false) {
                                    var v13 int8 = (v5 * v5)
                                    switch -2147483648 {
                                    case 457:
                                        if (v1 > 0) {
                                            var v14 string = "foo bar"
                                            _ = v14
                                        } else {
                                            fmt.Println(fi6)
                                            fmt.Println(v13, v4, v13)
                                            v5 = v13
                                        }
                                    case -9223372036854775808:
                                        fmt.Println(v1, fi9)
                                    default:
                                        var v15 int = 0
                                        _ = v15
                                    }
                                    switch fi6 {
                                    case 0:
                                        v16 := GenS3{F0: 284, F1: 2}
                                        _ = v16
                                    case 9223372036854775807:
                                        if true {
                                            v13 -= v5
                                        } else {
                                            v5 = v13
                                        }
                                    default:
                                        switch ((("a" + ((((("hello" + ("a" + "a")) + "hello") + "hello") + "foo bar") + "")) + "\n\t") + "hello") {
                                        case "":
                                            var v17 int16 = 32767
                                            _ = v17
                                        }
                                    }
                                    _ = v13
                                } else {
                                    v1 = v1
                                }
                                if len(v3) > 0 {
                                    v3[0] = ("foo bar" + ("hello" + "\n\t"))
                                }
                            } else {
                                v18 := GenS4{F0: 418, F1: 309, F2: v4}
                                if len(v3) > 0 {
                                    v3[0] = ("hello" + ("\n\t" + ((("a" + ("\n\t" + ("foo bar" + "\n\t"))) + "") + "")))
                                }
                                _ = v18
                            }
                        }
                    }
                    _ = fi9
                }
            default:
                v5 = (v5 << ((429) & 63))
            }
        }
        v19 := GenS4{F0: 42, F1: 65535, F2: 128}
        v20 := GenS5{F0: 2}
        _ = fi6
        _ = v19
        _ = v20
    }
    fmt.Println(v1)
    v1 = (v4 - (v4 % (((v1 & (^(v1)))) | 1)))
    v21 := GenS3{F0: 4294967295, F1: v1}
    var v22 int8 = (-(v5))
    v4 = (v1 << ((375) & 63))
    for fi23 := 0; fi23 < 18; fi23++ {
        v3 = append(v3, "foo bar")
        _ = fi23
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
