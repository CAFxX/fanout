package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 bool = false
    v2 := []string{("a" + ""), "hello", "\n\t"}
    v3 := make(map[uint32]bool)
    if v1 {
        if ((v1 && ((2 == 1) && (0 != 1))) && (!(false && true))) {
            if len(v2) > 0 {
                v2[0] = ("" + "foo bar")
            }
        } else {
            switch 127 {
            case -759:
                for fi4 := 0; fi4 < 7; fi4++ {
                    var v5 int16 = 1000
                    _ = fi4
                    _ = v5
                }
            default:
                for fi6 := 0; fi6 < 8; fi6++ {
                    var v7 uint16 = 1
                    switch "hello" {
                    case "\n\t":
                        switch ("hello" + ("a" + ("" + ("hello" + ("a" + "\n\t"))))) {
                        case "hello":
                            var v8 uint8 = 1
                            _ = v8
                        }
                    case "foo bar":
                        for fi9 := 0; fi9 < 12; fi9++ {
                            if (!true) {
                                var v10 string = (("\n\t" + ("hello" + ("" + "hello"))) + "\n\t")
                                var v11 string = v10
                                var v12 int32 = 42
                                delete(v3, 42)
                                v13 := v3[284]
                                switch v10 {
                                case "hello":
                                    var v14 int8 = 1
                                    _ = v14
                                default:
                                    var v15 uint = 37
                                    _ = v15
                                }
                                _ = v10
                                _ = v11
                                _ = v12
                                _ = v13
                            }
                            _ = fi9
                        }
                    case "a":
                        var v16 float32 = (float32(1.5) * 0.0)
                        _ = v16
                    }
                    _ = fi6
                    _ = v7
                }
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                switch "hello" {
                case "":
                    var v18 uint = 1
                    _ = v18
                case "\n\t":
                    v19 := GenS1{F0: 565}
                    _ = v19
                }
                fmt.Println(fi17, fi17)
                var v20 float32 = 3.14159
                _ = fi17
                _ = v20
            }
            var v21 int64 = 375
            _ = v21
        }
    } else {
        var v22 string = (("hello" + "foo bar") + "\n\t")
        if true {
            switch "" {
            case "a":
                fmt.Println(v22, v22)
            case "foo bar":
                fmt.Println(v1, v22)
            case "":
                delete(v3, 467)
            }
            var v23 uint64 = 18446744073709551615
            _ = v23
        }
        _ = len(v2)
        _ = v22
    }
    v24 := GenS1{F0: 2}
    for fi25 := 0; fi25 < 3; fi25++ {
        fmt.Println(fi25, v1)
        _ = fi25
    }
    fmt.Println(v1)
    v26 := GenS1{F0: 509}
    var v27 uint16 = 42
    var v28 uint8 = 81
    for fi29 := 0; fi29 < 3; fi29++ {
        fmt.Println(v27, fi29, fi29)
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
