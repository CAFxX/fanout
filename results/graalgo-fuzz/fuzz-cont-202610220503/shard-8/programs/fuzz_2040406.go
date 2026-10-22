package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint }

func main() {
    v1 := []string{("" + ("" + "a")), ""}
    var v2 uint32 = 716
    v3 := make(map[uint32]string)
    var v4 string = "\n\t"
    var v5 int64 = -1
    v6 := make(map[string]string)
    switch v4 {
    case "":
        var v7 bool = (2 <= 42)
        _ = v7
    case "\n\t":
        v8 := GenS1{F0: -876, F1: 18446744073709551615, F2: 1}
        _ = v8
    }
    switch v4 {
    case "hello":
        fmt.Println(v5)
    }
    v5 = 761
    if true {
        _ = len(v3)
        v4 = (v4 + "hello")
    }
    var v9 int8 = -44
    v2 += 0
    v10 := GenS1{F0: 128, F1: 884, F2: 2}
    if false {
        var v11 string = v4
        var v12 bool = false
        var v13 float32 = 0.0
        _ = v11
        _ = v12
        _ = v13
    } else {
        var v14 uint64 = 4294967295
        v9 -= (v9 % ((v9) | 1))
        _ = v14
    }
    var v15 uint16 = 32767
    for fi16 := 0; fi16 < 12; fi16++ {
        switch v4 {
        case "\n\t":
            _ = len(v3)
        case "hello":
            for fi17 := 0; fi17 < 12; fi17++ {
                fmt.Println(fi16, v9)
                switch 162 {
                case -2147483648:
                    switch fi17 {
                    case 590:
                        fmt.Println(fi16)
                    default:
                        var v18 int64 = (v5 ^ (v5 >> ((42) & 63)))
                        _ = v18
                    }
                }
                var v19 string = (v4 + (("hello" + v4) + "\n\t"))
                if false {
                    if (0.0 != float64(-2.25)) {
                        v19 = v4
                        v20 := GenS1{F0: (fi16 + fi16), F1: 0, F2: 89}
                        if (!(!false)) {
                            v21 := GenS2{F0: (v9 * (v9 % ((107) | 1))), F1: 42}
                            var v22 uint32 = (v2 | v2)
                            var v23 uint32 = (v22 >> ((1) & 63))
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v24 := GenS3{F0: 693}
                                _ = v24
                            }
                            var v25 int = fi17
                            _ = v25
                        }
                        v26 := GenS1{F0: -823, F1: 980, F2: 0}
                        _ = v20
                        _ = v26
                    } else {
                        fmt.Println(v4, v2)
                        v9 -= v9
                    }
                    v4 = "hello"
                    v9 -= v9
                } else {
                    fmt.Println(fi16, fi16, v5)
                    v1 = append(v1, "")
                }
                v19 = v19
                var v27 uint16 = 0
                v4 = (((v4 + (v19 + ("hello" + ""))) + (v4 + (v4 + (v4 + v19)))) + ("\n\t" + "foo bar"))
                if len(v1) > 0 {
                    v1[0] = v19
                }
                _ = fi17
                _ = v19
                _ = v27
            }
        }
        _ = fi16
    }
    var v28 int64 = v5
    if ((1.5 < 1e10) && true) {
        var v29 uint16 = 633
        switch 28 {
        case 9223372036854775807:
            v30 := GenS2{F0: v9, F1: 2147483647}
            _ = v30
        case 42:
            v31 := GenS2{F0: (v9 ^ 1), F1: 42}
            _ = v31
        }
        _ = v29
    }
    v32 := GenS2{F0: v9, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
