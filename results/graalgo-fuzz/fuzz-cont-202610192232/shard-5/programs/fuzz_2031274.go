package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (0.1 - float64(0.1))
    var v3 uint16 = 165
    switch 127 {
    case -1:
        var v4 int8 = 42
        _ = v4
    }
    v5 := GenS1{F0: 174}
    if (0 < -1) {
        var v6 int8 = 77
        var v7 uint64 = 1
        _ = v6
        _ = v7
    } else {
        if (2147483648 != 762) {
            fmt.Println(v3, v2)
            var v8 int8 = 100
            _ = v8
        } else {
            var v9 uint8 = 0
            _ = v9
        }
        switch "\n\t" {
        case "foo bar":
            for fi10 := 0; fi10 < 3; fi10++ {
                for fi11 := 0; fi11 < 1; fi11++ {
                    if ("foo bar" < "foo bar") {
                        var v12 int = (fi10 & (fi11 ^ (-(fi11))))
                        var v13 string = "hello"
                        _ = v12
                        _ = v13
                    } else {
                        var v14 uint16 = (v3 - v3)
                        v15 := GenS2{F0: (42 != 0), F1: 42, F2: "a"}
                        _ = v14
                        _ = v15
                    }
                    var v16 bool = false
                    switch (fi11 + 128) {
                    case -1:
                        fmt.Println(fi11)
                    case 42:
                        if ((!(("\n\t" + ("" + (("hello" + ("hello" + ("\n\t" + "foo bar"))) + ""))) >= (((("foo bar" + "") + ("" + "foo bar")) + "hello") + "hello"))) && (!v16)) {
                            v1[0] = "a"
                            if v16 {
                                switch "" {
                                case "hello":
                                    var v17 float32 = float32((0.0 + (v2 - (v2 / float64((((-2.25 - 3.14159) + 0.1) + 3.14159))))))
                                    _ = v17
                                case "foo bar":
                                    v3 = (^(v3))
                                    v18 := GenS1{F0: 42}
                                    v2 += v2
                                    _ = v18
                                case "\n\t":
                                    var v19 uint32 = 2147483647
                                    _ = v19
                                default:
                                    var v20 int = fi10
                                    _ = v20
                                }
                                var v21 int64 = -9223372036854775808
                                if (1 >= 0) {
                                    if (!(42 >= 42)) {
                                        if v16 {
                                            var v22 uint8 = 128
                                            v21 -= v21
                                            fmt.Println(v16)
                                            _ = v22
                                        } else {
                                            v23 := []int{fi11, (fi11 - 862), fi11}
                                            _ = v23
                                        }
                                        v2 += (v2 * v2)
                                    } else {
                                        var v24 bool = (v16 || (false && true))
                                        var v25 uint = 42
                                        var v26 uint16 = (v3 % ((v3) | 1))
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    switch ("\n\t" + ("a" + "a")) {
                                    case "":
                                        fmt.Println(v16)
                                    case "foo bar":
                                        v1[fi10] = "\n\t"
                                    default:
                                        if false {
                                            if (v16 && (-128 != 127)) {
                                                if v16 {
                                                    v27 := GenS3{F0: 255}
                                                    _ = v27
                                                }
                                            }
                                            fmt.Println(v2, fi10, v2)
                                            if (164 < 1) {
                                                fmt.Println(v21)
                                                var v28 float32 = (0.1 / 1e10)
                                                _ = v28
                                            }
                                            v29 := []int{(-(fi11)), (fi11 << ((255) & 63))}
                                            _ = v29
                                        }
                                    }
                                    v16 = v16
                                } else {
                                    v16 = v16
                                    v30 := GenS4{F0: v21}
                                    var v31 int8 = 42
                                    _ = v30
                                    _ = v31
                                }
                                _ = v21
                            }
                            var v32 uint8 = 176
                            _ = v32
                        }
                    default:
                        v33 := []int{}
                        _ = v33
                    }
                    _ = fi11
                    _ = v16
                }
                _ = fi10
            }
        case "\n\t":
            var v34 uint16 = v3
            _ = v34
        }
    }
    v35 := []int{-2147483648, 0}
    v2 += (-2.25 - ((v2 - v2) - (((1.5 - 0.1) + -2.25) - 0.1)))
    fmt.Println(v3, v2)
    var v36 string = "hello"
    v37 := GenS4{F0: 524}
    v38 := GenS4{F0: -975}
    var v39 uint = 42
    fmt.Println(v2, v39, v36)
    v3 = v3
    var v40 int8 = 1
    if false {
        fmt.Println(v39, v39)
        for fi41 := 0; fi41 < 11; fi41++ {
            fmt.Println(fi41)
            v3 += (v3 | 814)
            fmt.Println(fi41, v40, v2)
            _ = fi41
        }
    } else {
        v39 = v39
        var v42 uint32 = 830
        _ = v42
    }
    var v43 uint = v39
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v43
    _ = v5
}
