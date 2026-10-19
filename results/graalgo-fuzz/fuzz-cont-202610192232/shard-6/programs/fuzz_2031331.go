package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 uint; F1 int64; F2 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    v4 := []int{-1}
    v5 := []int{}
    v6 := make(map[int]int)
    var v7 float32 = 1.5
    if false {
        var v8 int16 = 32767
        _ = len(v5)
        _ = v8
    } else {
        switch 1 {
        case 670:
            v9 := v6[-261]
            _ = v9
        case -200:
            switch "hello" {
            case "\n\t":
                v10 := GenS1{F0: 1, F1: "a", F2: ("foo bar" + ((("" + "foo bar") + ("hello" + "foo bar")) + "hello"))}
                _ = v10
            case "hello":
                var v11 uint8 = 144
                _ = v11
            case "foo bar":
                var v12 string = "hello"
                _ = v12
            default:
                switch 993 {
                case 128:
                    var v13 int = 4294967295
                    _ = v13
                case 32768:
                    v7 -= float32(0.1)
                case -709:
                    var v14 uint32 = 342
                    _ = v14
                default:
                    if (false && ((4294967295 >= 32768) || (-1 > -762))) {
                        v7 += 0.1
                        switch ("" + "foo bar") {
                        case "":
                            for fi15 := 0; fi15 < 8; fi15++ {
                                if (!true) {
                                    switch 42 {
                                    case 340:
                                        _ = len(v6)
                                    case 4294967295:
                                        if true {
                                            if ((114 > 1000) && ("a" >= "\n\t")) {
                                                v16 := GenS1{F0: 1, F1: "", F2: "hello"}
                                                v7 = v7
                                                _ = v16
                                            } else {
                                                var v17 uint8 = 71
                                                _ = v17
                                            }
                                            for fi18 := 0; fi18 < 2; fi18++ {
                                                v19 := GenS1{F0: 9223372036854775807, F1: ("a" + "a"), F2: "hello"}
                                                if false {
                                                    fmt.Println(fi15)
                                                    v7 += v7
                                                }
                                                var v20 int = fi18
                                                _ = len(v1)
                                                v20 = (fi18 + (fi18 + fi15))
                                                _ = fi18
                                                _ = v19
                                                _ = v20
                                            }
                                            fmt.Println(fi15, v7)
                                        } else {
                                            var v21 int = fi15
                                            var v22 float64 = (0.1 / (((0.1 - 0.0) / 0.1) / 1.5))
                                            switch fi15 {
                                            case 32767:
                                                fmt.Println(v21, v21)
                                            }
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                }
                                fmt.Println(v7)
                                _ = fi15
                            }
                        case "a":
                            var v23 string = "a"
                            _ = v23
                        }
                    }
                }
            }
        }
    }
    switch 128 {
    case 127:
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 14; fi25++ {
                v26 := GenS1{F0: 217, F1: ("hello" + ((("a" + "") + ("foo bar" + "")) + "\n\t")), F2: "a"}
                _ = len(v4)
                _ = fi25
                _ = v26
            }
            if ((18446744073709551615 >= 255) && true) {
                if true {
                    var v27 int32 = 0
                    _ = v27
                }
                v7 = v7
                for fi28 := 0; fi28 < 18; fi28++ {
                    if (false && false) {
                        if false {
                            v7 += v7
                        } else {
                            var v29 string = "hello"
                            fmt.Println(fi24, fi28)
                            if true {
                                fmt.Println(v7, fi28, v7)
                                if false {
                                    v30 := GenS2{F0: 255, F1: -32, F2: 4294967295}
                                    var v31 int16 = 692
                                    _ = v30
                                    _ = v31
                                }
                                switch v29 {
                                case "\n\t":
                                    v7 = 0.1
                                }
                            }
                            _ = v29
                        }
                    } else {
                        v32 := GenS1{F0: 345, F1: ("" + "hello"), F2: "\n\t"}
                        _ = v32
                    }
                    v7 = v7
                    fmt.Println(fi28)
                    _ = fi28
                }
            }
            _ = fi24
        }
    case 128:
        fmt.Println(v7)
    case 658:
        v6[-9223372036854775808] = 127
    }
    switch "\n\t" {
    case "hello":
        if false {
            v7 = v7
        }
    }
    v33 := GenS1{F0: -9223372036854775808, F1: "a", F2: "hello"}
    switch 1 {
    case 148:
        var v34 float64 = 1e10
        _ = v34
    case -446:
        v35 := GenS1{F0: 696, F1: "hello", F2: ""}
        _ = v35
    case 32768:
        var v36 int8 = 42
        _ = v36
    }
    var v37 uint8 = 35
    delete(v3, ("a" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
