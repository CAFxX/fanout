package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint64) uint16 {
    v5 := []int{-2147483648, -475, -9223372036854775808}
    if (1 == 65535) {
        v5 = append(v5, 42)
        switch 9223372036854775807 {
        case -2147483648:
            switch "hello" {
            case "\n\t":
                p0 = 1
            case "hello":
                switch 42 {
                case 32768:
                    switch ("a" + "") {
                    case "a":
                        var v6 float64 = -2.25
                        _ = v6
                    default:
                        p0 -= p0
                    }
                }
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = 42
                }
            default:
                if len(v5) > 0 {
                    v5[0] = 876
                }
            }
        }
        var v7 int = 792
        _ = v7
    } else {
        fmt.Println(p0)
    }
    _ = p0
    _ = v5
    return 32767
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"foo bar", (("a" + "foo bar") + ""), "\n\t", "a"}
    v3 := make(map[uint32]string)
    var v4 int64 = 9223372036854775807
    switch "\n\t" {
    case "a":
        v4 += (v4 / ((v4) | 1))
    case "hello":
        v4 -= 1
    case "":
        fmt.Println(v4)
    default:
        var v8 uint32 = 4294967295
        _ = v8
    }
    v9 := GenS1{F0: (("hello" + "foo bar") + "a")}
    var v10 uint8 = 2
    v11 := GenS1{F0: "a"}
    var v12 uint8 = v10
    v3[420] = (("hello" + "foo bar") + "foo bar")
    v10 = 128
    _ = len(v2)
    v12 = v10
    var v13 int32 = 42
    v14 := GenS1{F0: ""}
    var v15 uint8 = (v12 / (((^(v10))) | 1))
    v12 = (v15 | (v15 * v10))
    switch ("a" + "") {
    case "\n\t":
        for fi16 := 0; fi16 < 1; fi16++ {
            v2 = append(v2, "a")
            v13 += v13
            if true {
                fmt.Println(v4)
                switch "a" {
                case "hello":
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v12 -= v15
                        _ = fi17
                    }
                case "foo bar":
                    v10 -= 42
                case "a":
                    v12 += v12
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 int16 = -32768
                    if false {
                        var v20 int64 = (v4 >> ((0) & 63))
                        _ = v20
                    } else {
                        if false {
                            fmt.Println(v19, v10, fi16)
                            switch "a" {
                            case "hello":
                                v13 = 54
                            case "foo bar":
                                if (491 < 18446744073709551615) {
                                    fmt.Println(v13, v4)
                                    v19 = -88
                                } else {
                                    switch fi16 {
                                    case 9223372036854775807:
                                        if ((v13 * (v13 % ((v13) | 1))) < (v13 / ((433) | 1))) {
                                            v19 = (^(v19))
                                            var v21 string = "a"
                                            var v22 float64 = (0.0 * 0.0)
                                            var v23 int64 = v4
                                            v22 -= v22
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        }
                                    default:
                                        if false {
                                            v24 := GenS1{F0: "a"}
                                            _ = v24
                                        }
                                    }
                                    _ = len(v2)
                                    var v25 uint8 = v15
                                    _ = v25
                                }
                            case "":
                                v15 -= v10
                            default:
                                delete(v1, (fi18 - -45))
                            }
                            var v26 int8 = 127
                            _ = v26
                        }
                        v19 = -32768
                        var v27 uint = 4294967295
                        _ = v27
                    }
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi16
        }
    default:
        for fi28 := 0; fi28 < 4; fi28++ {
            switch "a" {
            case "hello":
                var v29 int32 = v13
                _ = v29
            default:
                v15 += 1
            }
            _ = fi28
        }
    }
    var v30 float64 = (1e10 + 0.0)
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
