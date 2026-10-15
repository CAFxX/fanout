package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 string }

func main() {
    var v1 int16 = -373
    v2 := []int{128, -1}
    var v3 uint = 0
    v4 := make(map[string]bool)
    var v5 string = "foo bar"
    var v6 float64 = 1e10
    if true {
        fmt.Println(v5, v5)
        _ = len(v4)
    } else {
        if false {
            delete(v4, v5)
            v3 += 18446744073709551615
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 bool = false
                if false {
                    v6 -= (((float64(3.14159) * (0.1 / 1e10)) / v6) + (-2.25 + 1.5))
                    switch fi7 {
                    case -363:
                        v9 := v4["a"]
                        _ = v9
                    case 32768:
                        switch v5 {
                        case "a":
                            v3 += (^(v3))
                        case "hello":
                            v10 := GenS1{F0: 0, F1: "hello", F2: 42}
                            _ = v10
                        default:
                            var v11 uint64 = 4294967295
                            _ = v11
                        }
                    case 453:
                        if false {
                            switch fi7 {
                            case -11:
                                if (v3 < (v3 + 117)) {
                                    var v12 uint8 = 246
                                    v13 := GenS1{F0: 258, F1: v5, F2: 1}
                                    _ = v12
                                    _ = v13
                                } else {
                                    var v14 int = 0
                                    v1 += -32768
                                    _ = v14
                                }
                            case 42:
                                if v8 {
                                    fmt.Println(v8)
                                    if v8 {
                                        var v15 uint16 = 0
                                        var v16 uint32 = 2147483648
                                        _ = v15
                                        _ = v16
                                    }
                                    v8 = (1e10 > (float64(0.0) + ((1e10 * v6) - (0.1 / v6))))
                                }
                            case 112:
                                for fi17 := 0; fi17 < 15; fi17++ {
                                    v3 = (v3 - (v3 - v3))
                                    var v18 int16 = 42
                                    var v19 bool = false
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            }
                        } else {
                            var v20 float32 = float32(3.14159)
                            _ = v20
                        }
                    default:
                        v8 = ((!(v6 < v6)) || false)
                    }
                    if true {
                        v5 = ""
                    }
                } else {
                    v6 += v6
                    v21 := GenS1{F0: 65535, F1: ((((v5 + ("a" + "a")) + v5) + v5) + ((v5 + v5) + "a")), F2: 42}
                    _ = v21
                }
                _ = fi7
                _ = v8
            }
        } else {
            v5 = "foo bar"
            var v22 uint32 = 625
            _ = v22
        }
    }
    v6 += v6
    v23 := GenS2{F0: -532, F1: 9223372036854775807, F2: (("" + "\n\t") + "hello")}
    var v24 uint8 = 255
    v25 := GenS3{F0: 9223372036854775807}
    var v26 int16 = v1
    if true {
        for fi27 := 0; fi27 < 17; fi27++ {
            if true {
                fmt.Println(v24)
            } else {
                v5 = ((v5 + v5) + "\n\t")
                v1 -= (v1 * v26)
            }
            v6 = -2.25
            v28 := GenS3{F0: fi27}
            _ = fi27
            _ = v28
        }
        var v29 float32 = (((0.0 / 1.0) + 3.14159) + 0.0)
        _ = v29
    } else {
        for fi30 := 0; fi30 < 8; fi30++ {
            switch fi30 {
            case -2147483648:
                var v31 int16 = 1000
                _ = v31
            case 1:
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 string = "a"
                    v24 += 42
                    _ = fi32
                    _ = v33
                }
            case 127:
                for fi34 := 0; fi34 < 8; fi34++ {
                    var v35 uint = v3
                    v3 += (v3 & 4294967295)
                    if false {
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if (530 < 1) {
                                    fmt.Println(v3, v24)
                                    delete(v4, ((("a" + v5) + v5) + v5))
                                    var v36 int32 = -1
                                    _ = v36
                                }
                            } else {
                                v24 -= (v24 / ((v24) | 1))
                                var v37 uint8 = 255
                                v38 := GenS4{F0: 698, F1: v5}
                                _ = v37
                                _ = v38
                            }
                        case "a":
                            switch (-(fi34)) {
                            case -2147483648:
                                _ = len(v2)
                            }
                        }
                        fmt.Println(v35, v6)
                    }
                    fmt.Println(fi30, v1)
                    v1 -= (v26 * v1)
                    _ = len(v4)
                    var v39 int8 = 100
                    delete(v4, v5)
                    v3 -= (v3 / (((v3 & 18446744073709551615)) | 1))
                    _ = fi34
                    _ = v35
                    _ = v39
                }
            }
            _ = fi30
        }
        var v40 float64 = v6
        _ = v40
    }
    v41 := GenS3{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
