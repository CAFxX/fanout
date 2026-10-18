package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string; F1 int }

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := []string{("hello" + (((("" + "foo bar") + "foo bar") + ("" + ("" + (("hello" + ("a" + ("\n\t" + "a"))) + "a")))) + "a")), "foo bar", "\n\t"}
    var v4 string = "foo bar"
    var v5 uint = 4294967295
    var v6 int = 4294967295
    v3 = append(v3, v4)
    var v7 float64 = 1e10
    var v8 float64 = ((0.1 * 0.0) * v7)
    v9 := GenS1{F0: 1}
    if (0 <= 42) {
        var v10 string = ""
        fmt.Println(v7, v10)
        _ = v10
    }
    switch -1 {
    case 127:
        for fi11 := 0; fi11 < 1; fi11++ {
            switch fi11 {
            case -402:
                switch "" {
                case "\n\t":
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 int = 496
                        _ = fi12
                        _ = v13
                    }
                }
            default:
                switch v4 {
                case "":
                    for fi14 := 0; fi14 < 10; fi14++ {
                        var v15 uint64 = 42
                        var v16 string = (v4 + (v4 + "a"))
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
            }
            v2[fi11] = 42
            _ = fi11
        }
    case 128:
        var v17 int64 = 303
        _ = v17
    case -2147483648:
        for fi18 := 0; fi18 < 16; fi18++ {
            switch "hello" {
            case "a":
                var v19 uint64 = 1
                _ = v19
            case "":
                var v20 int16 = 42
                _ = v20
            }
            for fi21 := 0; fi21 < 8; fi21++ {
                v8 = v7
                v4 = "hello"
                var v22 int8 = -128
                _ = fi21
                _ = v22
            }
            _ = fi18
        }
    }
    v23 := GenS2{F0: -956, F1: 1}
    var v24 uint8 = 181
    v8 = v8
    fmt.Println(v7)
    if (127 >= -89) {
        delete(v2, v6)
    }
    delete(v2, (-(v6)))
    if (true && v1) {
        v25 := v2[v6]
        var v26 int = 1
        if (v5 == v5) {
            if ((-2.25 * 1e10) > -2.25) {
                var v27 uint16 = 0
                _ = v27
            } else {
                var v28 uint64 = 1
                _ = v28
            }
            switch v6 {
            case 686:
                if false {
                    switch ("\n\t" + ("hello" + v4)) {
                    case "foo bar":
                        if v1 {
                            v29 := v2[v6]
                            v2[(v6 - (-(v26)))] = 1
                            _ = v29
                        }
                    }
                    v30 := GenS2{F0: 42, F1: v25}
                    fmt.Println(v6)
                    _ = v30
                }
            case -1:
                v31 := v2[(-(v6))]
                _ = v31
            }
            for fi32 := 0; fi32 < 5; fi32++ {
                switch (((((v4 + "\n\t") + v4) + v4) + (v4 + v4)) + "foo bar") {
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        if true {
                            if (v1 && v1) {
                                for fi33 := 0; fi33 < 4; fi33++ {
                                    fmt.Println(v7, v7)
                                    _ = fi33
                                }
                                var v34 uint32 = 1
                                if ((v24 >= v24) || true) {
                                    delete(v2, (v6 ^ 0))
                                } else {
                                    var v35 float64 = v8
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v36 := GenS3{F0: (("" + "") + (v4 + v4)), F1: (v26 + -662)}
                                    _ = v35
                                    _ = v36
                                }
                                _ = v34
                            } else {
                                v26 += v6
                            }
                        }
                    case "":
                        if (false || false) {
                            var v37 uint8 = (^(v24))
                            _ = len(v2)
                            _ = v37
                        }
                    }
                case "foo bar":
                    for fi38 := 0; fi38 < 20; fi38++ {
                        var v39 int8 = 125
                        var v40 string = (("\n\t" + "") + v4)
                        v41 := GenS1{F0: 42}
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                }
                switch v4 {
                case "a":
                    v7 += 0.1
                case "\n\t":
                    if ((481 != -525) && true) {
                        v42 := GenS3{F0: ("foo bar" + (("\n\t" + (v4 + (v4 + "hello"))) + (v4 + "hello"))), F1: v26}
                        _ = v42
                    }
                }
                _ = fi32
            }
        }
        _ = v25
        _ = v26
    }
    if len(v3) > 0 {
        v3[0] = v4
    }
    v4 = v4
    var v43 bool = (1 <= 0)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
