package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int64; F2 int32 }
type GenS5 struct { F0 string }
func genFn0() uint32 {
    var v7 int64 = 1
    if false {
        switch (("foo bar" + "a") + "\n\t") {
        case "":
            v7 -= 42
        case "a":
            var v8 string = "\n\t"
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 11; fi9++ {
                switch "" {
                case "\n\t":
                    var v10 uint32 = 1
                    _ = v10
                }
                _ = fi9
            }
        default:
            if true {
                v11 := GenS1{F0: 541, F1: 0, F2: 659}
                if false {
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 uint64 = 437
                        _ = fi12
                        _ = v13
                    }
                    if (false || (!(!true))) {
                        v14 := GenS1{F0: 409, F1: 42, F2: -2147483648}
                        switch 408 {
                        case -600:
                            switch "a" {
                            case "foo bar":
                                switch 1 {
                                case 1:
                                    for fi15 := 0; fi15 < 20; fi15++ {
                                        v7 += v7
                                        _ = fi15
                                    }
                                }
                            case "":
                                var v16 int32 = -2147483648
                                _ = v16
                            case "a":
                                if ((true || true) || true) {
                                    for fi17 := 0; fi17 < 14; fi17++ {
                                        var v18 uint32 = 42
                                        switch fi17 {
                                        case 819:
                                            fmt.Println(fi17)
                                        case 2147483647:
                                            var v19 bool = (true || true)
                                            _ = v19
                                        case -2147483648:
                                            switch fi17 {
                                            case 127:
                                                for fi20 := 0; fi20 < 4; fi20++ {
                                                    v21 := GenS2{F0: "a"}
                                                    _ = fi20
                                                    _ = v21
                                                }
                                            default:
                                                v7 -= -9223372036854775808
                                            }
                                        }
                                        _ = fi17
                                        _ = v18
                                    }
                                    if true {
                                        v22 := GenS3{F0: 1}
                                        if false {
                                            var v23 uint = 65535
                                            v7 = (^(v7))
                                            if (((true || true) || false) || false) {
                                                v24 := GenS3{F0: 403}
                                                v23 += 4294967295
                                                _ = v24
                                            }
                                            _ = v23
                                        }
                                        v7 += v7
                                        _ = v22
                                    } else {
                                        var v25 uint32 = 65535
                                        var v26 int64 = v7
                                        _ = v25
                                        _ = v26
                                    }
                                    fmt.Println(v7)
                                }
                            }
                        case 2147483647:
                            var v27 float32 = -2.25
                            _ = v27
                        }
                        _ = v14
                    } else {
                        fmt.Println(v7)
                        v28 := []int{-692, -1, 9223372036854775807}
                        _ = v28
                    }
                }
                v29 := GenS4{F0: 701, F1: -9223372036854775808, F2: -1}
                _ = v11
                _ = v29
            } else {
                fmt.Println(v7)
                var v30 int = 32768
                switch ("a" + ("a" + "hello")) {
                case "\n\t":
                    v30 += -2147483648
                case "hello":
                    var v31 int = v30
                    _ = v31
                case "a":
                    fmt.Println(v30, v7)
                }
                _ = v30
            }
        }
        v32 := []int{4294967295}
        v33 := GenS5{F0: ""}
        _ = v32
        _ = v33
    } else {
        var v34 int8 = -128
        _ = v34
    }
    v7 -= v7
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -2147483648
    var v3 int8 = 112
    var v4 uint64 = 2
    var v5 uint64 = 1
    v6 := []int{}
    v2 -= (-(v2))
    for fi35 := 0; fi35 < 3; fi35++ {
        v5 = (v4 | 42)
        _ = fi35
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v36 string = "hello"
    var v37 float64 = 1e10
    v38 := GenS5{F0: (v36 + v36)}
    if false {
        var v39 uint = 42
        _ = v39
    } else {
        v40 := GenS4{F0: 255, F1: 9223372036854775807, F2: -746}
        _ = v40
    }
    _ = len(v1)
    if true {
        v37 += v37
        var v41 int = (v2 << ((331) & 63))
        _ = v41
    }
    var v42 uint32 = 4294967295
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
