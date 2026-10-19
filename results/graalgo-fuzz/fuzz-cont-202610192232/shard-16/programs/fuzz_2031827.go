package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    var v1 int64 = 939
    v2 := []int{}
    v3 := make(map[string]int)
    var v4 int16 = 32767
    v5 := []uint32{83, 42, 0}
    var v6 int64 = 1
    v7 := GenS1{F0: ("foo bar" + "\n\t"), F1: v6, F2: 655}
    v2 = append(v2, 1)
    var v8 bool = true
    fmt.Println(v4, v6, v8)
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            if ((-2.25 == 1e10) && false) {
                v11 := GenS1{F0: "hello", F1: 42, F2: 255}
                _ = v11
            } else {
                _ = len(v2)
                fmt.Println(v1)
            }
            var v12 float32 = 1.5
            v13 := GenS2{F0: "\n\t", F1: 42}
            v6 -= 0
            _ = fi10
            _ = v12
            _ = v13
        }
        var v14 int8 = 0
        v1 = v1
        _ = fi9
        _ = v14
    }
    var v15 int = 9223372036854775807
    v1 = (v6 / ((v1) | 1))
    v16 := GenS1{F0: (("hello" + ("\n\t" + ("\n\t" + "a"))) + ("foo bar" + "a")), F1: (^(v1)), F2: 2147483648}
    if ((3.14159 - (3.14159 - (0.1 - 1e10))) < -2.25) {
        _ = len(v5)
        if v8 {
            v3["foo bar"] = 0
        }
        var v17 float32 = 3.14159
        _ = v17
    } else {
        var v18 int = (-(v15))
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 int64 = v6
            _ = fi19
            _ = v20
        }
        var v21 uint64 = 4294967295
        _ = v18
        _ = v21
    }
    v1 = v6
    switch 42 {
    case 128:
        v8 = (!v8)
    case -9223372036854775808:
        v6 = v6
    default:
        if ((2147483648 >= 736) || v8) {
            var v22 uint8 = 127
            if v8 {
                v23 := GenS2{F0: "\n\t", F1: 113}
                for fi24 := 0; fi24 < 3; fi24++ {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v6)
                    case "\n\t":
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            fmt.Println(v22, v15, v1)
                        default:
                            v25 := GenS1{F0: "hello", F1: v1, F2: uint32(v22)}
                            _ = v25
                        }
                    default:
                        v15 += (fi24 % (((fi24 + 128)) | 1))
                    }
                    _ = fi24
                }
                _ = v23
            } else {
                var v26 uint32 = 256
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 float32 = 0.0
                    if (0 == -641) {
                        fmt.Println(v26)
                        if true {
                            for fi29 := 0; fi29 < 12; fi29++ {
                                fmt.Println(v8, v22, v15)
                                _ = fi29
                            }
                        } else {
                            for fi30 := 0; fi30 < 5; fi30++ {
                                var v31 int32 = -1
                                var v32 uint8 = (v22 * v22)
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            var v33 uint64 = 2
                            fmt.Println(v28)
                            _ = v33
                        }
                        fmt.Println(fi27, v8, v6)
                    }
                    switch fi27 {
                    case 128:
                        v15 += 127
                    case 439:
                        var v34 float64 = -2.25
                        _ = v34
                    case 32767:
                        fmt.Println(v26)
                    }
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
            switch (((("foo bar" + "hello") + "hello") + "a") + "hello") {
            case "\n\t":
                var v35 uint16 = 32767
                _ = v35
            case "foo bar":
                fmt.Println(v1, v15, v8)
            }
            _ = v22
        }
    }
    var v36 float64 = 0.0
    v4 = (v4 + (v4 % ((v4) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
