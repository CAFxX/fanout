package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[uint32]string)
    v3 := []string{"a", "foo bar"}
    v4 := []int{96, 128, 32768}
    v5 := GenS1{F0: -889, F1: ""}
    v1 = append(v1, "foo bar")
    v6 := GenS1{F0: 82, F1: "a"}
    var v7 int16 = 0
    for fi8 := 0; fi8 < 13; fi8++ {
        if (false || (!false)) {
            var v9 bool = true
            v9 = v9
            _ = v9
        } else {
            var v10 uint8 = 7
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS1{F0: -1, F1: (("hello" + (("hello" + "a") + "a")) + ("hello" + "foo bar"))}
                if (1e10 < 1e10) {
                    switch "" {
                    case "\n\t":
                        if (!(!false)) {
                            var v13 float32 = 1.5
                            if (367 <= 543) {
                                _ = len(v3)
                                switch "" {
                                case "\n\t":
                                    if (true && false) {
                                        var v14 int = fi11
                                        if (((false || false) || (112 < 42)) || (true || true)) {
                                            var v15 int16 = v7
                                            var v16 float64 = 1e10
                                            var v17 bool = false
                                            if true {
                                                switch v14 {
                                                case -283:
                                                    if v17 {
                                                        fmt.Println(fi8)
                                                    } else {
                                                        fmt.Println(v16)
                                                        var v18 int8 = 127
                                                        v2[uint32(v10)] = ("foo bar" + ("foo bar" + "a"))
                                                        _ = v18
                                                    }
                                                case 1:
                                                    fmt.Println(v15, fi11)
                                                    _ = len(v2)
                                                    v16 = v16
                                                case 32767:
                                                    fmt.Println(v14, fi8, v15)
                                                }
                                                v19 := GenS1{F0: 1, F1: "foo bar"}
                                                fmt.Println(v10, v14)
                                                _ = v19
                                            }
                                            var v20 uint64 = 18446744073709551615
                                            v21 := GenS1{F0: -958, F1: "foo bar"}
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                            _ = v20
                                            _ = v21
                                        } else {
                                            var v22 uint8 = 127
                                            v13 -= v13
                                            _ = v22
                                        }
                                        _ = v14
                                    } else {
                                        v2[65535] = "hello"
                                        var v23 int32 = -2147483648
                                        var v24 uint64 = 4294967295
                                        _ = v23
                                        _ = v24
                                    }
                                case "a":
                                    switch 32768 {
                                    case 9223372036854775807:
                                        fmt.Println(fi11)
                                        v7 = (v7 + v7)
                                    case -517:
                                        var v25 uint8 = v10
                                        _ = v25
                                    case 1:
                                        v26 := GenS1{F0: 404, F1: ""}
                                        _ = v26
                                    }
                                case "hello":
                                    v27 := GenS1{F0: 753, F1: "\n\t"}
                                    _ = v27
                                default:
                                    v28 := GenS1{F0: 0, F1: ("" + "foo bar")}
                                    _ = v28
                                }
                            }
                            v10 = v10
                            _ = v13
                        }
                    default:
                        v10 += (v10 * (v10 >> ((42) & 63)))
                    }
                }
                v29 := GenS2{F0: 727, F1: 129}
                var v30 uint16 = 274
                _ = fi11
                _ = v12
                _ = v29
                _ = v30
            }
            _ = len(v2)
            _ = v10
        }
        _ = fi8
    }
    fmt.Println(v7)
    switch 4294967295 {
    case -624:
        v31 := GenS1{F0: 0, F1: ("" + "\n\t")}
        _ = v31
    case 4294967295:
        var v32 uint16 = 42
        _ = v32
    case 32767:
        var v33 string = "foo bar"
        _ = v33
    default:
        switch ("foo bar" + ("" + "a")) {
        case "\n\t":
            var v34 int16 = v7
            _ = v34
        case "foo bar":
            var v35 uint = 527
            _ = v35
        default:
            for fi36 := 0; fi36 < 12; fi36++ {
                switch (fi36 | 4294967295) {
                case -2147483648:
                    if (-631 > -1) {
                        var v37 uint32 = 2
                        v38 := GenS1{F0: 0, F1: "foo bar"}
                        _ = v37
                        _ = v38
                    }
                case -1:
                    v7 += v7
                case 9223372036854775807:
                    v7 += 32767
                default:
                    fmt.Println(fi36, v7)
                }
                _ = fi36
            }
        }
    }
    var v39 bool = (false || (393 != 2))
    fmt.Println(v39, v39)
    var v40 int = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
