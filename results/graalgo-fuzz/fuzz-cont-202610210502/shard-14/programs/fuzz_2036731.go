package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0() uint64 {
    var v6 uint64 = 0
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(fi7, fi7)
        var v8 int8 = -115
        v6 = 740
        _ = fi7
        _ = v8
    }
    _ = v6
    return v6
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -9223372036854775808
    v3 := []string{}
    v4 := []int{178, -922, 298}
    v5 := []int{4294967295, 379}
    v2 = (v2 ^ 1)
    if (!true) {
        var v9 uint8 = 131
        var v10 int64 = (-(v2))
        _ = v10
        _ = v9
    }
    v2 = v2
    if (42 == 978) {
        var v11 int = 210
        var v12 float32 = 0.0
        _ = v11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch -9223372036854775808 {
            case 0:
                delete(v1, "hello")
            case -505:
                v2 = v2
            }
            var v14 float64 = 0.1
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 149, F1: ("a" + "hello"), F2: 42}
        v1["foo bar"] = 2147483648
        _ = v15
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 := GenS1{F0: 127, F1: ("\n\t" + ""), F2: 1}
    v2 += (v2 & v2)
    var v17 float32 = 1.5
    v18 := GenS1{F0: 2, F1: "hello", F2: 0}
    switch 1 {
    case -1:
        var v19 int16 = -32768
        _ = v19
    }
    switch "hello" {
    case "\n\t":
        var v20 int64 = 9223372036854775807
        _ = v20
    case "a":
        switch ("" + "a") {
        case "hello":
            switch 541 {
            case 0:
                switch "foo bar" {
                case "foo bar":
                    v21 := GenS1{F0: 1, F1: ("hello" + (("a" + "foo bar") + "")), F2: -1}
                    _ = v21
                case "hello":
                    var v22 uint8 = 2
                    _ = v22
                case "a":
                    v17 = (v17 / v17)
                }
            default:
                var v23 uint = 18446744073709551615
                _ = v23
            }
        case "":
            v24 := GenS2{F0: 82, F1: 42}
            _ = v24
        case "foo bar":
            var v25 int = 450
            _ = v25
        default:
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v17, v2)
            case "hello":
                v17 = (float32(float64(v17)) / 1.0)
            }
        }
    }
    delete(v1, "")
    var v26 int = 4294967295
    if true {
        switch "foo bar" {
        case "hello":
            if false {
                switch "foo bar" {
                case "a":
                    if (216 >= -821) {
                        var v27 float64 = -2.25
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 14; fi28++ {
                            switch (fi28 - 1) {
                            case 545:
                                v2 += (v2 - int64(v26))
                            case 32768:
                                var v29 uint32 = 1
                                _ = v29
                            case 1:
                                if true {
                                    v1[""] = 639
                                    var v30 uint16 = 510
                                    _ = v30
                                } else {
                                    for fi31 := 0; fi31 < 4; fi31++ {
                                        switch "\n\t" {
                                        case "foo bar":
                                            delete(v1, "hello")
                                        case "hello":
                                            if (!(2147483648 <= 42)) {
                                                v17 = v17
                                                var v32 int32 = 2147483647
                                                v26 = (v26 ^ fi31)
                                                v26 += -502
                                                _ = v32
                                            }
                                        }
                                        if (false || (65535 != 65535)) {
                                            var v33 uint = 42
                                            _ = v33
                                        }
                                        _ = fi31
                                    }
                                }
                            default:
                                var v34 uint16 = 255
                                _ = v34
                            }
                            v35 := GenS3{F0: 113}
                            _ = fi28
                            _ = v35
                        }
                        v26 -= 0
                    }
                case "hello":
                    v36 := GenS2{F0: 100, F1: 170}
                    _ = v36
                default:
                    v17 += v17
                }
                var v37 int8 = 1
                _ = v37
            }
        default:
            for fi38 := 0; fi38 < 10; fi38++ {
                for fi39 := 0; fi39 < 8; fi39++ {
                    switch "a" {
                    case "hello":
                        var v40 uint8 = 0
                        _ = v40
                    default:
                        v17 = v17
                    }
                    _ = fi39
                }
                _ = fi38
            }
        }
        switch v26 {
        case 32767:
            var v41 int32 = 42
            _ = v41
        case 9223372036854775807:
            var v42 int16 = 0
            _ = v42
        default:
            var v43 uint64 = 1
            _ = v43
        }
        var v44 uint16 = 2
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
