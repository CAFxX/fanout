package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint64; F2 string }
func genFn0(p0 uint64, p1 int32) int {
    switch -2147483648 {
    case 32767:
        fmt.Println(p0, p1)
    case 9223372036854775807:
        v5 := []int{-302, 128, 4294967295}
        _ = v5
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint64 = 704
    v2 := []uint32{42, 2147483648, 42, 793}
    v3 := make(map[uint32]int)
    var v4 int16 = 42
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 float32 = 1.5
        v1 = 2
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v8 := GenS1{F0: 65535, F1: 84, F2: -9223372036854775808}
            switch 116 {
            case 9223372036854775807:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 127:
                fmt.Println(v4, v7)
            }
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    var v9 bool = true
    var v10 float64 = (1.5 / (1.5 / -2.25))
    v11 := v3[42]
    v12 := v3[42]
    for fi13 := 0; fi13 < 7; fi13++ {
        if v9 {
            if (100 == 117) {
                switch ("\n\t" + "") {
                case "hello":
                    v3[65535] = int(v4)
                default:
                    var v14 uint8 = 252
                    _ = v14
                }
                switch (("foo bar" + "") + "foo bar") {
                case "hello":
                    if v9 {
                        v10 -= 3.14159
                        if v9 {
                            var v15 uint = 1
                            _ = v15
                        } else {
                            var v16 uint = 65535
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v16
                        }
                    }
                case "a":
                    v4 += v4
                case "\n\t":
                    fmt.Println(v9, v9, v12)
                default:
                    switch (v11 >> ((4294967295) & 63)) {
                    case 32767:
                        if (65535 != 802) {
                            var v17 float64 = ((v10 - v10) - ((1.5 / 1.5) * v10))
                            for fi18 := 0; fi18 < 14; fi18++ {
                                v12 -= genFn0(v1, -981)
                                v19 := GenS1{F0: 113, F1: 1, F2: 0}
                                v20 := GenS2{F0: 128}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            switch "" {
                            case "\n\t":
                                var v21 int8 = 42
                                _ = v21
                            case "a":
                                var v22 int8 = 91
                                _ = v22
                            case "foo bar":
                                v23 := GenS3{F0: ("foo bar" + ("\n\t" + "a")), F1: 714, F2: ("hello" + "hello")}
                                _ = v23
                            }
                            _ = v17
                        }
                    case 4294967295:
                        v24 := GenS2{F0: 2}
                        _ = v24
                    }
                }
            } else {
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 int8 = -1
                    _ = fi25
                    _ = v26
                }
                var v27 uint = 961
                _ = v27
            }
        } else {
            fmt.Println(v1, v10)
            v28 := GenS1{F0: 2147483647, F1: 0, F2: -9223372036854775808}
            _ = v28
        }
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            for fi29 := 0; fi29 < 18; fi29++ {
                v10 = v10
                v4 -= 86
                _ = fi29
            }
        }
        _ = fi13
    }
    var v30 uint64 = 4294967295
    switch (((("a" + ("a" + ("foo bar" + (("foo bar" + (("\n\t" + ("\n\t" + ("\n\t" + "a"))) + "hello")) + "foo bar")))) + "hello") + "hello") + "a") {
    case "hello":
        v12 = 365
    case "a":
        v30 -= v1
    case "\n\t":
        if ((!(!(553 >= 32767))) && (1 != 42)) {
            var v31 uint32 = 536
            fmt.Println(v30, v9, v12)
            _ = len(v2)
            _ = v31
        }
    default:
        if (v1 < (v1 * v30)) {
            _ = len(v3)
            v12 -= v12
            for fi32 := 0; fi32 < 1; fi32++ {
                switch "foo bar" {
                case "\n\t":
                    var v33 uint8 = 128
                    _ = v33
                default:
                    fmt.Println(v10, v9, v11)
                }
                _ = fi32
            }
        } else {
            for fi34 := 0; fi34 < 13; fi34++ {
                v10 = (v10 / v10)
                _ = fi34
            }
            switch "\n\t" {
            case "foo bar":
                var v35 uint16 = 0
                _ = v35
            case "a":
                var v36 int = (v11 ^ v12)
                _ = v36
            case "\n\t":
                v30 = 337
            }
            var v37 uint = 18446744073709551615
            _ = v37
        }
    }
    fmt.Println(v12, v4, v12)
    v4 -= -240
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
