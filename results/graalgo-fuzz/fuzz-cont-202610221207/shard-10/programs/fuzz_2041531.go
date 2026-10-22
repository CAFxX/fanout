package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint }

func main() {
    var v1 int = 127
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    for fi4 := 0; fi4 < 18; fi4++ {
        switch 640 {
        case 316:
            switch v2 {
            case "\n\t":
                delete(v3, 43)
            case "foo bar":
                if (886 == 18446744073709551615) {
                    for fi5 := 0; fi5 < 20; fi5++ {
                        _ = len(v3)
                        v6 := []int{}
                        v1 = fi4
                        _ = fi5
                        _ = v6
                    }
                    var v7 uint16 = 255
                    _ = v7
                }
            default:
                v8 := GenS1{F0: "a", F1: 854, F2: 311}
                _ = v8
            }
        default:
            var v9 bool = false
            _ = v9
        }
        if (true && ((120 < 123) || (127 != 2))) {
            var v10 int8 = 42
            _ = v10
        }
        _ = fi4
    }
    v2 = (v2 + (((v2 + "foo bar") + (v2 + v2)) + "hello"))
    v1 += v1
    v11 := []int{v1}
    if false {
        if (false && true) {
            var v12 uint32 = 2
            v13 := GenS2{F0: 2}
            var v14 uint16 = 91
            _ = v12
            _ = v13
            _ = v14
        }
        var v15 bool = false
        _ = v15
    } else {
        switch (v1 | v1) {
        case 531:
            switch v2 {
            case "foo bar":
                var v16 bool = (false || (-573 > -32768))
                _ = v16
            case "":
                v2 = (((((v2 + (("" + "a") + (v2 + "\n\t"))) + (v2 + v2)) + ((v2 + (((("a" + (((v2 + (v2 + v2)) + v2) + (((((v2 + ("" + "a")) + (v2 + "")) + v2) + v2) + ((v2 + v2) + (("" + "a") + v2))))) + (("a" + v2) + (v2 + (v2 + v2)))) + (v2 + v2)) + "foo bar")) + v2)) + v2) + "\n\t")
            }
        case 9223372036854775807:
            switch (v1 % (((v1 >> ((42) & 63))) | 1)) {
            case 0:
                for fi17 := 0; fi17 < 12; fi17++ {
                    var v18 float32 = (float32((0.0 * -2.25)) * ((1e10 - 1.5) * 0.1))
                    switch "hello" {
                    case "foo bar":
                        for fi19 := 0; fi19 < 3; fi19++ {
                            v20 := GenS1{F0: v2, F1: 844, F2: 545}
                            _ = fi19
                            _ = v20
                        }
                    case "hello":
                        v18 = v18
                    case "a":
                        var v21 int = fi17
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 16; fi22++ {
                            v1 = (v1 | -530)
                            var v23 uint16 = 32767
                            _ = fi22
                            _ = v23
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            case -393:
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 uint = 4294967295
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v2 = ((((v2 + (v2 + v2)) + (v2 + "")) + v2) + v2)
                        if true {
                            fmt.Println(fi24)
                        }
                        _ = fi26
                    }
                    _ = fi24
                    _ = v25
                }
            default:
                v2 = v2
            }
        default:
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 float64 = 0.0
                _ = fi27
                _ = v28
            }
        }
        v29 := GenS2{F0: 1}
        _ = v29
    }
    switch v2 {
    case "a":
        var v30 string = v2
        _ = v30
    case "hello":
        fmt.Println(v1)
    }
    if len(v11) > 0 {
        v11[0] = v1
    }
    fmt.Println(v2)
    if false {
        v1 = -844
        _ = len(v11)
    } else {
        var v31 int32 = -1
        _ = v31
    }
    var v32 uint = 65535
    var v33 int64 = 368
    v34 := GenS3{F0: false, F1: 2, F2: (v32 * 4294967295)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
