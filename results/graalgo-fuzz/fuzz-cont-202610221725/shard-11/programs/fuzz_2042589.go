package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int16; F2 uint64 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 uint = 322
    if true {
        v5 := []int{32767, 9223372036854775807, 9223372036854775807, 293}
        if (42 <= v4) {
            v6 := GenS1{F0: 2, F1: -2147483648, F2: 32767}
            v4 = 18446744073709551615
            if true {
                if len(v5) > 0 {
                    v5[0] = 32768
                }
                v4 -= v4
                switch "a" {
                case "a":
                    for fi7 := 0; fi7 < 16; fi7++ {
                        for fi8 := 0; fi8 < 5; fi8++ {
                            fmt.Println(fi7, fi7, v4)
                            if false {
                                var v9 bool = (!true)
                                v4 += v4
                                _ = v9
                            } else {
                                var v10 string = ""
                                _ = v10
                            }
                            var v11 float64 = (0.1 - 1e10)
                            _ = fi8
                            _ = v11
                        }
                        v5 = append(v5, fi7)
                        _ = fi7
                    }
                case "":
                    if true {
                        var v12 uint = (^(v4))
                        _ = v12
                    } else {
                        var v13 int32 = 42
                        if (false && (4294967295 >= 1)) {
                            v14 := GenS1{F0: 255, F1: 4294967295, F2: 1}
                            v13 += v13
                            _ = v14
                        } else {
                            if false {
                                v15 := GenS2{F0: 18446744073709551615}
                                _ = v15
                            } else {
                                switch -594 {
                                case -187:
                                    var v16 uint16 = 32767
                                    _ = v16
                                case 9223372036854775807:
                                    for fi17 := 0; fi17 < 10; fi17++ {
                                        v4 += v4
                                        _ = fi17
                                    }
                                }
                                _ = len(v1)
                            }
                            v13 += -35
                        }
                        _ = v13
                    }
                }
            }
            _ = v6
        }
        _ = v5
    } else {
        var v18 uint16 = 1
        v19 := GenS1{F0: 255, F1: 4294967295, F2: int16(v4)}
        _ = v18
        _ = v19
    }
    fmt.Println(v4)
    v20 := []int{1, 499, -1}
    v2[65535] = "a"
    switch "a" {
    case "foo bar":
        if len(v20) > 0 {
            v20[0] = 0
        }
    case "hello":
        switch ((((("\n\t" + "foo bar") + (((("hello" + ("hello" + (("foo bar" + "a") + "foo bar"))) + "foo bar") + "foo bar") + ("foo bar" + "foo bar"))) + "\n\t") + ((("\n\t" + "hello") + ("hello" + "foo bar")) + "foo bar")) + "foo bar") {
        case "foo bar":
            switch -658 {
            case 185:
                fmt.Println(v4)
            default:
                var v21 int32 = -2147483648
                _ = v21
            }
        case "\n\t":
            for fi22 := 0; fi22 < 2; fi22++ {
                fmt.Println(fi22, v4)
                v4 -= 207
                switch (^(fi22)) {
                case 534:
                    var v23 uint32 = 2147483648
                    _ = v23
                case 9223372036854775807:
                    if (0 == 54) {
                        v24 := GenS3{F0: "\n\t", F1: -576, F2: 876}
                        for fi25 := 0; fi25 < 7; fi25++ {
                            if false {
                                v4 += v4
                                if ("a" == "hello") {
                                    var v26 float64 = 1.5
                                    var v27 uint64 = 2
                                    _ = v26
                                    _ = v27
                                }
                                v4 += v4
                            }
                            var v28 int8 = 127
                            _ = fi25
                            _ = v28
                        }
                        var v29 int = 0
                        _ = v24
                        _ = v29
                    } else {
                        v4 -= 221
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v4 = (^(v4))
                            _ = len(v2)
                            _ = fi30
                        }
                        _ = len(v20)
                    }
                default:
                    _ = len(v20)
                }
                _ = fi22
            }
        case "":
            if true {
                fmt.Println(v4)
            }
        default:
            var v31 uint8 = 0
            _ = v31
        }
    default:
        var v32 uint8 = 0
        _ = v32
    }
    v33 := GenS4{F0: -2147483648}
    switch "a" {
    case "\n\t":
        v4 -= v4
    case "":
        var v34 uint64 = 42
        _ = v34
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        for fi36 := 0; fi36 < 8; fi36++ {
            v4 += (^(v4))
            _ = fi36
        }
        _ = fi35
    }
    var v37 uint32 = 0
    var v38 uint64 = 4294967295
    switch "hello" {
    case "hello":
        v39 := GenS2{F0: v38}
        _ = v39
    }
    if len(v20) > 0 {
        v20[0] = 1
    }
    v40 := GenS3{F0: "a", F1: -436, F2: v38}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v40
}
