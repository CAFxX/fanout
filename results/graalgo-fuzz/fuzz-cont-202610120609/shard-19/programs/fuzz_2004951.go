package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }
type GenS3 struct { F0 uint; F1 uint; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint32; F2 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 1
    v3 := []string{("\n\t" + (("hello" + (("hello" + ("a" + "")) + "a")) + ("a" + ((("hello" + ("a" + "hello")) + (("a" + "hello") + "")) + "")))), ("\n\t" + ("foo bar" + ((((("" + "") + ("a" + (("hello" + "hello") + "a"))) + "") + "hello") + ("" + "hello")))), ("a" + "\n\t"), "a"}
    v4 := []int{9223372036854775807}
    v5 := make(map[uint32]uint32)
    v2 -= v2
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(v2, fi6)
            _ = fi7
        }
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi6
    }
    fmt.Println(v2)
    if (1 == 2) {
        v8 := GenS1{F0: (v2 - 0)}
        if (28 >= 42) {
            fmt.Println(v2)
            if ("foo bar" > "\n\t") {
                v2 += (v2 / ((v2) | 1))
                v2 -= 2
            } else {
                switch "\n\t" {
                case "hello":
                    switch "hello" {
                    case "":
                        fmt.Println(v2)
                    case "foo bar":
                        v9 := GenS1{F0: v2}
                        _ = v9
                    }
                case "a":
                    for fi10 := 0; fi10 < 18; fi10++ {
                        if (("\n\t" <= "hello") || true) {
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                if (-403 > -914) {
                                    v2 -= v2
                                    fmt.Println(v2, fi10)
                                    fmt.Println(fi10)
                                }
                            case "a":
                                for fi11 := 0; fi11 < 16; fi11++ {
                                    var v12 uint8 = uint8(v2)
                                    _ = fi11
                                    _ = v12
                                }
                            case "":
                                v2 = (v2 / ((v2) | 1))
                            default:
                                switch "a" {
                                case "":
                                    v13 := GenS2{F0: 1, F1: 42, F2: (fi10 - (fi10 / ((int(v2)) | 1)))}
                                    _ = v13
                                case "hello":
                                    fmt.Println(fi10)
                                case "a":
                                    v14 := GenS3{F0: 255, F1: 2, F2: 42}
                                    _ = v14
                                default:
                                    v2 = v2
                                }
                            }
                            if false {
                                v15 := GenS3{F0: 42, F1: 871, F2: 311}
                                var v16 int = (^(fi10))
                                var v17 uint32 = 4294967295
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        } else {
                            switch fi10 {
                            case 4294967295:
                                if false {
                                    var v18 float32 = 0.0
                                    _ = v18
                                }
                            case 9223372036854775807:
                                v19 := GenS1{F0: (^(v2))}
                                _ = v19
                            case 591:
                                if true {
                                    switch (fi10 * (fi10 / ((fi10) | 1))) {
                                    case 32767:
                                        fmt.Println(v2, fi10)
                                    case 1:
                                        v3 = append(v3, "foo bar")
                                    case 42:
                                        for fi20 := 0; fi20 < 15; fi20++ {
                                            var v21 int16 = -844
                                            fmt.Println(v2, v21, v2)
                                            if (-1 > -516) {
                                                var v22 int8 = 42
                                                var v23 int8 = (-(v22))
                                                switch (^(fi20)) {
                                                case 2147483647:
                                                    v22 = v22
                                                case -495:
                                                    v2 -= v2
                                                case 128:
                                                    var v24 uint16 = (v2 * 653)
                                                    _ = v24
                                                default:
                                                    if (true || true) {
                                                        v25 := GenS1{F0: (v2 + v2)}
                                                        var v26 uint32 = 0
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        var v27 uint = 2
                                                        if len(v4) > 0 {
                                                            v4[0] = 0
                                                        }
                                                        _ = v27
                                                    }
                                                }
                                                _ = v22
                                                _ = v23
                                            }
                                            _ = fi20
                                            _ = v21
                                        }
                                    }
                                } else {
                                    for fi28 := 0; fi28 < 14; fi28++ {
                                        switch fi28 {
                                        case 0:
                                            fmt.Println(fi28)
                                        case 127:
                                            switch ("\n\t" + "foo bar") {
                                            case "a":
                                                var v29 uint64 = 4294967295
                                                _ = v29
                                            case "hello":
                                                var v30 int16 = -1
                                                _ = v30
                                            case "":
                                                v2 += v2
                                            }
                                        }
                                        _ = fi28
                                    }
                                    fmt.Println(fi10, v2)
                                    var v31 uint32 = 148
                                    _ = v31
                                }
                            }
                            var v32 bool = (1 > 288)
                            _ = v32
                        }
                        _ = fi10
                    }
                default:
                    _ = len(v4)
                }
            }
        }
        var v33 uint32 = 4294967295
        _ = v33
        _ = v8
    }
    if false {
        var v34 bool = false
        _ = v34
    } else {
        fmt.Println(v2)
        v35 := GenS3{F0: 829, F1: 666, F2: 208}
        _ = v35
    }
    v36 := GenS4{F0: -1, F1: 2, F2: "a"}
    switch 0 {
    case 878:
        v1[687] = "a"
    }
    fmt.Println(v2)
    var v37 int8 = -128
    switch ((("hello" + "") + ("foo bar" + (("foo bar" + ("a" + ("hello" + ("a" + "")))) + "\n\t"))) + "") {
    case "hello":
        v2 -= (^(v2))
    }
    switch "foo bar" {
    case "a":
        v38 := GenS2{F0: 204, F1: 255, F2: 127}
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
