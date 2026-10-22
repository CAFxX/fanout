package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint; F1 bool; F2 int8 }
type GenS5 struct { F0 int; F1 uint32; F2 uint }
type GenS6 struct { F0 uint64; F1 int16 }

func main() {
    var v1 bool = (!(!false))
    v2 := make(map[uint32]uint32)
    v3 := []string{"hello", "hello"}
    v4 := make(map[int]string)
    v5 := make(map[uint32]int)
    if (false || (1 <= 32768)) {
        if (v1 && (!false)) {
            _ = len(v3)
            var v6 float32 = (((-2.25 + (0.0 / -2.25)) * 0.0) - (-2.25 + 0.1))
            switch "a" {
            case "\n\t":
                if v1 {
                    for fi7 := 0; fi7 < 9; fi7++ {
                        v8 := GenS1{F0: "hello", F1: 258}
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v6)
                        case "":
                            var v9 uint16 = 2
                            _ = v9
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                        }
                        v10 := GenS2{F0: 0, F1: -128}
                        _ = fi7
                        _ = v10
                        _ = v8
                    }
                } else {
                    v11 := GenS1{F0: "foo bar", F1: 0}
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v13 := GenS2{F0: -1, F1: 25}
                        var v14 string = ("\n\t" + "hello")
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    if true {
                        v15 := GenS2{F0: -1, F1: 127}
                        v1 = v1
                        for fi16 := 0; fi16 < 17; fi16++ {
                            v17 := GenS3{F0: true, F1: 945, F2: 42}
                            v1 = true
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    } else {
                        v1 = v1
                        var v18 uint = 1
                        _ = v18
                    }
                    _ = v11
                }
            case "a":
                v19 := GenS4{F0: 1, F1: true, F2: 100}
                _ = v19
            default:
                for fi20 := 0; fi20 < 18; fi20++ {
                    fmt.Println(v6, fi20, v1)
                    var v21 uint8 = 128
                    _ = fi20
                    _ = v21
                }
            }
            _ = v6
        }
    } else {
        v1 = (v1 && false)
    }
    v1 = v1
    var v22 uint64 = 831
    if (32767 < -32768) {
        v1 = v1
        v22 += v22
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := v4[-2147483648]
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 int64 = 9223372036854775807
                v26 = 1
                _ = fi25
                _ = v26
            }
            fmt.Println(v1, v24)
            _ = fi23
            _ = v24
        }
    } else {
        v1 = true
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 5; fi28++ {
            fmt.Println(fi28, v22)
            _ = fi28
        }
        fmt.Println(v22, v22)
        for fi29 := 0; fi29 < 20; fi29++ {
            _ = len(v3)
            if (v1 || v1) {
                _ = len(v3)
                delete(v2, 2147483648)
                if (-859 == -2147483648) {
                    var v30 int64 = -850
                    if v1 {
                        v31 := GenS5{F0: fi27, F1: 2147483647, F2: 65535}
                        if (true && (false || false)) {
                            if v1 {
                                v32 := GenS2{F0: -1, F1: 83}
                                fmt.Println(fi27, fi29, v22)
                                _ = v32
                            }
                        }
                        switch ("a" + ("foo bar" + "hello")) {
                        case "hello":
                            var v33 int8 = -128
                            delete(v2, 717)
                            _ = v33
                        case "":
                            _ = len(v3)
                        case "foo bar":
                            fmt.Println(v1)
                        default:
                            v34 := GenS3{F0: (v1 || (!(!(15 != 26)))), F1: 1, F2: 2}
                            _ = v34
                        }
                        _ = v31
                    } else {
                        v35 := GenS6{F0: 362, F1: -1}
                        var v36 int8 = 127
                        fmt.Println(fi29, fi29)
                        _ = v35
                        _ = v36
                    }
                    _ = v30
                } else {
                    switch fi29 {
                    case -2147483648:
                        v22 -= v22
                    default:
                        v1 = v1
                    }
                }
            } else {
                var v37 float32 = 0.1
                v22 -= v22
                _ = v37
            }
            _ = fi29
        }
        _ = fi27
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        _ = len(v2)
        _ = fi38
    }
    v39 := GenS6{F0: v22, F1: 42}
    if v1 {
        _ = len(v3)
        switch ("a" + (((("\n\t" + "") + ("" + ("\n\t" + "\n\t"))) + ((("foo bar" + ((("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) + "") + ("hello" + ""))) + "hello") + "hello")) + "\n\t")) {
        case "hello":
            var v40 uint = 4294967295
            _ = v40
        }
    }
    v41 := GenS5{F0: 4294967295, F1: 28, F2: 129}
    var v42 uint32 = 4294967295
    v1 = v1
    fmt.Println(v22)
    v22 = 714
    v1 = (255 <= 2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
