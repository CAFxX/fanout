package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint64 = 2
    v3 := []string{((("a" + ("" + "")) + "hello") + "\n\t"), ("a" + "foo bar"), "foo bar"}
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 65
    var v6 int = 2147483647
    delete(v4, 98)
    var v7 int = (v6 + (v6 * 32767))
    var v8 string = ("" + "foo bar")
    switch (-(v6)) {
    case 2147483647:
        fmt.Println(v8)
    case -9223372036854775808:
        v9 := GenS1{F0: (v8 + (((("hello" + ((v8 + v8) + v8)) + v8) + v8) + "foo bar")), F1: 32767, F2: false}
        _ = v9
    default:
        var v10 uint64 = 0
        _ = v10
    }
    v5 -= 255
    fmt.Println(v8, v7)
    v8 = ""
    v11 := GenS2{F0: 127, F1: v8}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 float64 = 3.14159
    if (!false) {
        var v13 uint = 279
        if true {
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS3{F0: -128, F1: false}
                if false {
                    if (42 > 0) {
                        var v16 string = "a"
                        _ = v16
                    } else {
                        var v17 int32 = -1
                        var v18 uint16 = v5
                        _ = v17
                        _ = v18
                    }
                    var v19 int32 = 1
                    _ = v19
                }
                _ = fi14
                _ = v15
            }
            v13 = v13
        }
        switch v8 {
        case "hello":
            switch v8 {
            case "foo bar":
                v8 = (((("\n\t" + (v8 + "foo bar")) + v8) + "foo bar") + "\n\t")
            case "a":
                var v20 float32 = float32((0.1 - float64(1e10)))
                _ = v20
            case "hello":
                switch v8 {
                case "foo bar":
                    if true {
                        for fi21 := 0; fi21 < 3; fi21++ {
                            for fi22 := 0; fi22 < 5; fi22++ {
                                fmt.Println(fi21, fi22, v12)
                                if len(v3) > 0 {
                                    v3[0] = ((v8 + (v8 + (v8 + v8))) + v8)
                                }
                                _ = fi22
                            }
                            _ = fi21
                        }
                    } else {
                        for fi23 := 0; fi23 < 14; fi23++ {
                            switch fi23 {
                            case 0:
                                switch v8 {
                                case "a":
                                    v5 -= (^(v5))
                                case "\n\t":
                                    delete(v4, 2147483648)
                                }
                            case 9223372036854775807:
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    v8 = "foo bar"
                                    _ = fi24
                                }
                            case 4294967295:
                                if (!(0 == -101)) {
                                    switch fi23 {
                                    case -1:
                                        v6 -= (v6 / ((v6) | 1))
                                    case 32768:
                                        v25 := GenS2{F0: -128, F1: "a"}
                                        _ = v25
                                    default:
                                        fmt.Println(v7)
                                    }
                                }
                            }
                            var v26 uint16 = v5
                            var v27 int16 = 32767
                            _ = fi23
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS4{F0: 1}
                        _ = v28
                    }
                default:
                    fmt.Println(v12)
                }
            }
        }
        _ = v13
    }
    var v29 float64 = ((v12 / v12) - v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
