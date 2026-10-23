package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 uint16; F1 string; F2 bool }
func genFn0(p0 uint16) uint32 {
    switch 4294967295 {
    case -14:
        var v4 int64 = -1
        _ = v4
    case -2147483648:
        var v5 int8 = -81
        _ = v5
    case 2147483647:
        p0 += p0
    }
    v6 := GenS1{F0: 331, F1: p0}
    var v7 int32 = int32(p0)
    v8 := []int{325, -335}
    p0 = (p0 - 116)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 798
}
func genFn1() bool {
    var v9 int64 = 9223372036854775807
    var v10 int = -1
    switch (v10 * (v10 << ((448) & 63))) {
    case 42:
        if true {
            v9 = v9
        } else {
            v11 := []int{}
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := GenS1{F0: 42, F1: 32767}
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    v10 -= (v10 / ((v10) | 1))
    v9 = v9
    _ = v10
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", "\n\t", (("hello" + "hello") + "hello")}
    if false {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint64 = 18446744073709551615
            _ = fi15
            _ = v16
        }
        switch "" {
        case "a":
            _ = len(v3)
        case "hello":
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS2{F0: (fi17 >> ((1) & 63))}
                _ = fi17
                _ = v18
            }
        default:
            var v19 uint16 = 2
            _ = v19
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 string = (("hello" + "foo bar") + (("hello" + "foo bar") + ("\n\t" + ("hello" + "foo bar"))))
            v22 := GenS1{F0: 18446744073709551615, F1: 65535}
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 int64 = 9223372036854775807
        _ = v23
    }
    var v24 string = ("foo bar" + "a")
    v25 := GenS3{F0: 4294967295}
    v24 = v24
    v26 := GenS4{F0: 563, F1: 42, F2: 0}
    var v27 float32 = 1.5
    switch (("" + v24) + "\n\t") {
    case "foo bar":
        switch 1 {
        case 991:
            switch v24 {
            case "a":
                v27 += (v27 - v27)
            case "hello":
                fmt.Println(v24, v24)
            case "":
                var v28 uint64 = 42
                _ = v28
            }
        case 9223372036854775807:
            var v29 uint16 = 32767
            _ = v29
        default:
            switch 127 {
            case -1:
                switch v24 {
                case "":
                    if (false || ((v24 + "hello") == (v24 + "\n\t"))) {
                        var v30 uint64 = 0
                        if false {
                            fmt.Println(v24)
                        } else {
                            v31 := GenS1{F0: (^(v30)), F1: 1}
                            v24 = ("\n\t" + v24)
                            v24 = "hello"
                            _ = v31
                        }
                        if (!((1e10 - (0.1 + 1.5)) != 1e10)) {
                            var v32 int16 = -32768
                            for fi33 := 0; fi33 < 20; fi33++ {
                                var v34 float32 = v27
                                delete(v2, genFn0(828))
                                _ = fi33
                                _ = v34
                            }
                            _ = v32
                        }
                        _ = v30
                    }
                }
            case 0:
                fmt.Println(v27, v24)
            case 4294967295:
                for fi35 := 0; fi35 < 2; fi35++ {
                    v24 = v24
                    fmt.Println(v27, v24, fi35)
                    _ = fi35
                }
            }
        }
    case "hello":
        var v36 uint64 = 18446744073709551615
        _ = v36
    }
    fmt.Println(v24)
    v37 := GenS5{F0: 953}
    for fi38 := 0; fi38 < 7; fi38++ {
        v39 := GenS6{F0: 196, F1: -2147483648}
        var v40 int8 = -128
        v41 := GenS7{F0: 2, F1: (v24 + ((v24 + (((("hello" + v24) + v24) + v24) + "\n\t")) + ("" + v24))), F2: true}
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
}
