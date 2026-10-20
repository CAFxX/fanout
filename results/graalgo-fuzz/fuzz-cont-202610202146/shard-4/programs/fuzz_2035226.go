package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 6; fi6++ {
        switch "hello" {
        case "a":
            v7 := []int{9223372036854775807}
            _ = v7
        }
        v8 := GenS1{F0: 2, F1: 4294967295}
        _ = fi6
        _ = v8
    }
    var v9 int32 = 2147483647
    _ = v9
    return 636
}
func genFn1() uint64 {
    var v10 int = 75
    v11 := GenS2{F0: 2, F1: -31}
    var v12 float32 = float32((float64(0.1) * 1e10))
    switch "a" {
    case "\n\t":
        v12 -= v12
    }
    v12 += -2.25
    _ = v10
    _ = v11
    _ = v12
    return 407
}

func main() {
    v1 := []string{}
    v2 := make(map[string]string)
    var v3 float64 = -2.25
    v4 := make(map[uint32]string)
    var v5 string = "a"
    v5 = v5
    if (!(65535 != 1)) {
        fmt.Println(v3, v3)
    }
    v3 -= v3
    v13 := v4[2147483647]
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            switch v5 {
            case "foo bar":
                fmt.Println(v3, v5, v5)
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = (v5 + v5)
                }
            default:
                var v16 string = v5
                _ = v16
            }
            v13 = (v5 + "foo bar")
            _ = fi15
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint8 = 1
            fmt.Println(v18, fi17)
            _ = fi17
            _ = v18
        }
        switch "\n\t" {
        case "a":
            if (!(!(42 >= -128))) {
                if (true || (1e10 >= 1.5)) {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 string = "a"
                        switch fi19 {
                        case 2147483647:
                            var v21 float64 = v3
                            _ = v21
                        case 147:
                            fmt.Println(fi19)
                        case 9223372036854775807:
                            if false {
                                switch "hello" {
                                case "foo bar":
                                    v22 := GenS2{F0: 42, F1: 42}
                                    _ = v22
                                case "a":
                                    var v23 uint = 0
                                    _ = v23
                                case "hello":
                                    switch 770 {
                                    case 1:
                                        v24 := GenS3{F0: 602}
                                        _ = v24
                                    case -79:
                                        var v25 uint8 = 128
                                        _ = v25
                                    }
                                }
                                v26 := GenS2{F0: 1, F1: 100}
                                if ((255 < 42) || true) {
                                    v27 := GenS4{F0: 2, F1: -1}
                                    fmt.Println(fi14, fi19)
                                    v28 := v4[2147483647]
                                    _ = v27
                                    _ = v28
                                }
                                _ = v26
                            } else {
                                v29 := GenS1{F0: 255, F1: 1}
                                _ = v29
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                    var v30 bool = (((173 >= 42) && true) || true)
                    var v31 int64 = -9223372036854775808
                    _ = v30
                    _ = v31
                } else {
                    if false {
                        var v32 uint64 = 18446744073709551615
                        if len(v1) > 0 {
                            v1[0] = (("hello" + v5) + "hello")
                        }
                        _ = v32
                    }
                    delete(v4, 693)
                }
                v5 = "a"
            }
        case "":
            delete(v2, "hello")
        }
        _ = fi14
    }
    v33 := GenS2{F0: 18446744073709551615, F1: -128}
    v34 := GenS4{F0: 101, F1: 42}
    v35 := GenS4{F0: 65535, F1: 157}
    var v36 uint64 = 1
    var v37 int32 = 42
    fmt.Println(v3, v5, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
