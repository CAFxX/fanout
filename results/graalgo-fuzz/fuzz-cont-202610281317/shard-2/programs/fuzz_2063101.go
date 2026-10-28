package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint32 = 519
    v2 := make(map[uint32]string)
    var v3 uint = 0
    v4 := make(map[string]string)
    var v5 int = 420
    if true {
        var v6 uint64 = 18446744073709551615
        if false {
            switch (("hello" + (((("a" + "\n\t") + ("a" + "\n\t")) + ("hello" + "a")) + ("a" + "\n\t"))) + "foo bar") {
            case "hello":
                switch (("\n\t" + ("hello" + ("" + (("foo bar" + "") + ("a" + "foo bar"))))) + ("\n\t" + "hello")) {
                case "":
                    fmt.Println(v6, v5, v6)
                case "foo bar":
                    v6 = (v6 & uint64(v1))
                }
            case "a":
                var v7 int16 = -245
                _ = v7
            default:
                switch 127 {
                case 0:
                    v8 := GenS1{F0: 32767, F1: v5, F2: (v5 + (v5 >> ((18446744073709551615) & 63)))}
                    _ = v8
                case -794:
                    var v9 uint8 = 128
                    _ = v9
                case 32767:
                    v3 = v3
                default:
                    if false {
                        var v10 int32 = 42
                        if true {
                            v11 := GenS2{F0: v3}
                            _ = v11
                        } else {
                            if (128 > uint8(v1)) {
                                v12 := GenS2{F0: v3}
                                _ = v12
                            } else {
                                v1 += (^(v1))
                                v3 -= v3
                                v5 -= v5
                            }
                            v3 = (v3 * (^(v3)))
                            for fi13 := 0; fi13 < 1; fi13++ {
                                fmt.Println(v10, v1, v1)
                                var v14 bool = true
                                _ = fi13
                                _ = v14
                            }
                        }
                        var v15 int8 = -128
                        _ = v10
                        _ = v15
                    }
                }
            }
            var v16 int16 = 1000
            _ = v16
        }
        v17 := GenS1{F0: 32767, F1: 0, F2: (v5 % ((v5) | 1))}
        _ = v17
        _ = v6
    } else {
        if false {
            if (1 <= 1000) {
                v18 := GenS2{F0: v3}
                _ = v18
            } else {
                if (false || false) {
                    var v19 uint = (v3 * v3)
                    var v20 float64 = 1.5
                    v20 -= (v20 / 1.0)
                    _ = v19
                    _ = v20
                }
            }
        }
        v21 := []int{(v5 % ((-648) | 1))}
        _ = v21
    }
    var v22 float64 = 1.5
    v5 -= (v5 + (v5 >> ((v3) & 63)))
    if true {
        if true {
            v23 := GenS1{F0: 569, F1: (v5 / (((-(v5))) | 1)), F2: v5}
            var v24 float32 = ((1.5 * 0.0) - 1.5)
            _ = v23
            _ = v24
        } else {
            v5 = v5
        }
        var v25 int64 = int64(v3)
        if false {
            v22 = (v22 + 3.14159)
        } else {
            v25 += (v25 | (v25 % ((9223372036854775807) | 1)))
            if ((^(v1)) > (v1 << ((42) & 63))) {
                v26 := []int{-1}
                if false {
                    var v27 int8 = -128
                    _ = v27
                } else {
                    fmt.Println(v25)
                }
                _ = v26
            } else {
                var v28 uint64 = 1
                v4[""] = "a"
                _ = v28
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                for fi30 := 0; fi30 < 6; fi30++ {
                    delete(v4, "a")
                    if (!true) {
                        switch "hello" {
                        case "\n\t":
                            switch ("foo bar" + ("" + "")) {
                            case "hello":
                                switch fi30 {
                                case 4294967295:
                                    if ((2147483647 == -2147483648) || false) {
                                        if false {
                                            if (-70 >= -20) {
                                                v31 := GenS2{F0: 2}
                                                v32 := []int{-677, (^(v5)), fi29}
                                                var v33 uint16 = 65535
                                                var v34 int64 = -763
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            }
                                            var v35 uint64 = 156
                                            v4["hello"] = ("a" + "\n\t")
                                            _ = v35
                                        }
                                        fmt.Println(fi29, v1)
                                        v25 -= 0
                                    }
                                default:
                                    v36 := []int{(v5 / ((v5) | 1)), (fi30 | 42), (fi30 << (((v3 + (v3 % ((1) | 1)))) & 63)), (fi30 + (v5 / ((1) | 1)))}
                                    _ = v36
                                }
                            case "a":
                                v5 -= fi29
                            case "":
                                v37 := GenS3{F0: 558}
                                _ = v37
                            }
                        case "a":
                            fmt.Println(fi30)
                        }
                    } else {
                        v38 := []int{}
                        fmt.Println(v3, fi29)
                        _ = v38
                    }
                    _ = fi30
                }
                _ = fi29
            }
        }
        _ = v25
    } else {
        v39 := GenS4{F0: 2}
        _ = v39
    }
    var v40 string = "foo bar"
    switch (v40 + "") {
    case "a":
        switch ("hello" + "\n\t") {
        case "foo bar":
            fmt.Println(v3, v22, v5)
        default:
            v3 = 289
        }
    case "\n\t":
        fmt.Println(v3, v3, v40)
    default:
        if (uint8(v3) <= uint8(v1)) {
            var v41 uint = v3
            _ = v41
        }
    }
    var v42 uint = v3
    v43 := []int{v5, v5, v5}
    var v44 uint8 = 231
    if (-1 <= int8(v5)) {
        var v45 uint64 = 489
        _ = v45
    } else {
        v3 -= v42
        for fi46 := 0; fi46 < 6; fi46++ {
            var v47 bool = ((v40 + "hello") >= v40)
            var v48 uint64 = 18446744073709551615
            _ = fi46
            _ = v47
            _ = v48
        }
    }
    var v49 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v49
    _ = v5
}
