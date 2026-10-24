package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int16; F2 string }
func genFn0(p0 int16) uint64 {
    switch -344 {
    case -741:
        p0 -= 0
    case -87:
        switch -638 {
        case 128:
            if (2 != 9) {
                if ((!true) && (true && true)) {
                    switch -1 {
                    case 127:
                        v7 := GenS1{F0: 42, F1: 55}
                        _ = v7
                    }
                    p0 += p0
                } else {
                    v8 := []int{-9223372036854775808, -9223372036854775808, 127, -137}
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    for fi9 := 0; fi9 < 5; fi9++ {
                        switch "a" {
                        case "hello":
                            p0 -= (p0 - (p0 / (((p0 >> ((722) & 63))) | 1)))
                        case "foo bar":
                            var v10 float32 = (-2.25 - 1e10)
                            _ = v10
                        case "\n\t":
                            fmt.Println(p0, p0)
                        default:
                            for fi11 := 0; fi11 < 15; fi11++ {
                                p0 -= -32768
                                switch 2147483647 {
                                case 9223372036854775807:
                                    var v12 uint = 42
                                    v12 = v12
                                    _ = v12
                                case 103:
                                    var v13 string = "foo bar"
                                    _ = v13
                                }
                                var v14 uint = 413
                                _ = fi11
                                _ = v14
                            }
                        }
                        v15 := GenS1{F0: 2147483647, F1: 1}
                        p0 = (p0 % ((610) | 1))
                        _ = fi9
                        _ = v15
                    }
                    _ = v8
                }
                var v16 bool = ((279 <= 634) || (-77 != -123))
                _ = v16
            }
        case -1:
            var v17 uint32 = 862
            _ = v17
        }
    case 690:
        fmt.Println(p0)
    }
    fmt.Println(p0)
    var v18 uint64 = 818
    v19 := GenS1{F0: 631, F1: 42}
    v20 := []int{1, 127, 4294967295, -2147483648}
    _ = p0
    _ = v18
    _ = v19
    _ = v20
    return (v18 - v18)
}

func main() {
    var v1 bool = false
    v2 := []uint32{379, 615, 2147483648, 2}
    v3 := make(map[string]uint32)
    v4 := []int{2147483647, 4294967295, 761}
    v5 := []string{("foo bar" + "foo bar")}
    var v6 int = 1
    if len(v2) > 0 {
        v2[0] = 857
    }
    v21 := GenS1{F0: 0, F1: 0}
    v1 = (142 == 42)
    v6 = (v6 & v6)
    v1 = (v1 || (!(71 == -1)))
    _ = len(v5)
    if v1 {
        fmt.Println(v1, v1)
    } else {
        if v1 {
            v22 := GenS1{F0: 2147483648, F1: -8}
            if v1 {
                var v23 int32 = 2147483647
                for fi24 := 0; fi24 < 18; fi24++ {
                    v23 = 42
                    _ = fi24
                }
                v23 = v23
                _ = v23
            }
            _ = v22
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            for fi26 := 0; fi26 < 7; fi26++ {
                switch (("" + ("hello" + "foo bar")) + "foo bar") {
                case "\n\t":
                    v6 -= fi26
                case "hello":
                    var v27 float64 = (3.14159 / -2.25)
                    _ = v27
                }
                _ = fi26
            }
            _ = fi25
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            v3["hello"] = 0
            switch fi28 {
            case 795:
                for fi29 := 0; fi29 < 3; fi29++ {
                    fmt.Println(v1)
                    var v30 uint32 = 2
                    _ = fi29
                    _ = v30
                }
            case 399:
                switch "hello" {
                case "\n\t":
                    var v31 float64 = ((0.0 / -2.25) - (0.0 + 1.5))
                    _ = v31
                case "":
                    var v32 int64 = 42
                    _ = v32
                case "a":
                    var v33 uint = 42
                    _ = v33
                default:
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v6 += fi28
                        _ = fi34
                    }
                }
            default:
                fmt.Println(v6, v1, v1)
            }
            v35 := GenS2{F0: 9223372036854775807, F1: v1}
            _ = fi28
            _ = v35
        }
    }
    if v1 {
        switch v6 {
        case 4294967295:
            v1 = v1
        case 227:
            v3["\n\t"] = 2
        default:
            if true {
                v36 := v3["hello"]
                _ = v36
            } else {
                v37 := GenS3{F0: 310}
                _ = v37
            }
        }
    } else {
        switch v6 {
        case 389:
            for fi38 := 0; fi38 < 13; fi38++ {
                var v39 int32 = 0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi38
                _ = v39
            }
        }
    }
    var v40 bool = true
    v41 := GenS4{F0: -27, F1: 0, F2: "\n\t"}
    var v42 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
