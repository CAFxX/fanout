package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 string }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]int)
    var v3 bool = (!true)
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    switch "\n\t" {
    case "":
        v6 := GenS1{F0: 16}
        _ = v6
    }
    v3 = v3
    v7 := v2[(("\n\t" + "\n\t") + "hello")]
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 bool = (!(!v3))
        v9 = v3
        var v10 int16 = -1
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 -= (-(v1))
    fmt.Println(v7, v1)
    if v3 {
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(v1)
            switch (("a" + ((("hello" + "a") + "a") + ("a" + "foo bar"))) + (("hello" + "foo bar") + "a")) {
            case "foo bar":
                delete(v5, v7)
            case "hello":
                if (v3 && true) {
                    if true {
                        v1 = (v1 + (v1 - (v1 & v1)))
                    } else {
                        if (true && v3) {
                            switch 128 {
                            case 32767:
                                _ = len(v5)
                            case 32768:
                                switch "" {
                                case "":
                                    switch v7 {
                                    case -9223372036854775808:
                                        _ = len(v4)
                                    case 128:
                                        for fi12 := 0; fi12 < 14; fi12++ {
                                            v13 := GenS1{F0: 241}
                                            _ = fi12
                                            _ = v13
                                        }
                                    default:
                                        for fi14 := 0; fi14 < 14; fi14++ {
                                            v3 = (!false)
                                            _ = fi14
                                        }
                                    }
                                }
                            case 1:
                                delete(v2, ("" + "foo bar"))
                            }
                            var v15 int32 = 1
                            if (0.1 <= float64(1e10)) {
                                for fi16 := 0; fi16 < 12; fi16++ {
                                    var v17 int32 = 462
                                    _ = fi16
                                    _ = v17
                                }
                                fmt.Println(v3, v3, fi11)
                            } else {
                                v1 -= v1
                            }
                            _ = v15
                        } else {
                            var v18 uint8 = 115
                            v19 := GenS1{F0: (^(v18))}
                            v20 := []int{390, 2147483647, (fi11 % ((fi11) | 1)), (fi11 + (v7 / ((fi11) | 1)))}
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        fmt.Println(v1)
                        switch fi11 {
                        case 128:
                            var v21 int32 = -124
                            _ = v21
                        case 32768:
                            if v3 {
                                for fi22 := 0; fi22 < 9; fi22++ {
                                    fmt.Println(v7, v1)
                                    v23 := GenS2{F0: 547, F1: "foo bar"}
                                    _ = fi22
                                    _ = v23
                                }
                            } else {
                                v1 += (v1 & v1)
                                var v24 int32 = 1
                                _ = v24
                            }
                        default:
                            v1 = v1
                        }
                    }
                    if (v3 && ((9223372036854775807 == 0) || ((v3 || v3) || v3))) {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            v26 := GenS2{F0: 2147483647, F1: "hello"}
                            v1 = v1
                            var v27 int16 = 1
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        v7 = v7
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v29 := v4[(-(v7))]
                            _ = fi28
                            _ = v29
                        }
                    }
                    fmt.Println(v3, v7)
                } else {
                    v30 := v2["hello"]
                    if (float32(1.5) > 0.1) {
                        var v31 uint32 = 42
                        _ = v31
                    } else {
                        v3 = (v3 || (2 >= 18446744073709551615))
                        v30 = v30
                        fmt.Println(v1, v3)
                    }
                    v7 = (v7 * (-(v7)))
                    _ = v30
                }
            default:
                var v32 string = ""
                _ = v32
            }
            _ = fi11
        }
        if v3 {
            if false {
                v3 = false
            } else {
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 int64 = 1
                    var v35 uint64 = 18446744073709551615
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
            if false {
                v36 := []int{(v7 >> ((4294967295) & 63))}
                _ = v36
            }
            switch "" {
            case "\n\t":
                switch (("\n\t" + "hello") + ("hello" + "")) {
                case "hello":
                    switch 164 {
                    case 128:
                        var v37 int = 736
                        _ = v37
                    default:
                        var v38 uint64 = 18446744073709551615
                        _ = v38
                    }
                case "foo bar":
                    v39 := GenS1{F0: 42}
                    _ = v39
                }
            }
        }
    } else {
        v7 = -231
        v40 := GenS2{F0: -785, F1: "hello"}
        fmt.Println(v7, v7)
        _ = v40
    }
    var v41 float64 = 3.14159
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v7
}
