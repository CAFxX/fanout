package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }

func main() {
    v1 := []int{-1, 283}
    var v2 uint8 = 170
    var v3 uint64 = 4294967295
    v4 := make(map[string]uint32)
    var v5 int8 = -1
    v5 += v5
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            v3 = (v3 ^ 2)
            var v8 float64 = 1e10
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v5 -= v5
    fmt.Println(v2)
    var v9 uint16 = 2
    v5 += (v5 % (((v5 - v5)) | 1))
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 float64 = 1.5
        switch "\n\t" {
        case "a":
            if (!true) {
                if true {
                    if ("" != "hello") {
                        if false {
                            v5 = (-(v5))
                            if (true || true) {
                                v12 := GenS1{F0: 394, F1: v3, F2: 0}
                                _ = v12
                            } else {
                                for fi13 := 0; fi13 < 18; fi13++ {
                                    v11 = float64((1.5 / 1.0))
                                    v14 := GenS2{F0: 32767, F1: 573}
                                    delete(v4, "a")
                                    var v15 int = fi13
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                                switch (("\n\t" + ("" + ((("a" + (("foo bar" + "\n\t") + ("foo bar" + ("" + "hello")))) + ("hello" + "hello")) + (("foo bar" + "foo bar") + "a")))) + (("\n\t" + (("a" + ("" + (("foo bar" + "a") + "foo bar"))) + "hello")) + (("hello" + "hello") + "a"))) {
                                case "":
                                    v9 = v9
                                }
                                if (!true) {
                                    for fi16 := 0; fi16 < 10; fi16++ {
                                        switch fi16 {
                                        case 42:
                                            v17 := GenS2{F0: 32768, F1: 2147483648}
                                            _ = v17
                                        case 862:
                                            v2 += v2
                                        }
                                        var v18 float64 = v11
                                        delete(v4, "\n\t")
                                        _ = fi16
                                        _ = v18
                                    }
                                    v19 := v4[("a" + "hello")]
                                    var v20 float32 = float32((v11 - (3.14159 / v11)))
                                    _ = v19
                                    _ = v20
                                }
                            }
                            v11 -= ((v11 * v11) + 3.14159)
                        }
                        v9 += (v9 - (v9 & 65535))
                    }
                    v9 += 0
                }
            } else {
                fmt.Println(v9, v3)
                if (!(!false)) {
                    var v21 float32 = -2.25
                    var v22 int64 = 42
                    v23 := GenS2{F0: 127, F1: 2147483647}
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(v2, fi10)
                }
                v24 := GenS3{F0: 2}
                _ = v24
            }
        case "\n\t":
            for fi25 := 0; fi25 < 14; fi25++ {
                switch (fi10 << ((4294967295) & 63)) {
                case 386:
                    var v26 uint32 = uint32(v9)
                    _ = v26
                default:
                    v5 = v5
                }
                var v27 string = ""
                _ = fi25
                _ = v27
            }
        }
        v11 += (((float64((-2.25 / 1.0)) - v11) * 1.5) / ((1.5 - 1e10) / v11))
        _ = fi10
        _ = v11
    }
    switch "" {
    case "a":
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 uint = 42
            var v30 uint32 = 65535
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        switch "\n\t" {
        case "\n\t":
            if (1.5 > float32((3.14159 * (1e10 / 1.5)))) {
                var v31 uint32 = 1
                v32 := GenS4{F0: 911}
                _ = v31
                _ = v32
            }
        case "foo bar":
            v1 = append(v1, 127)
        case "a":
            var v33 float32 = float32(0.1)
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        fmt.Println(v5)
        _ = fi34
    }
    var v35 int = -9223372036854775808
    var v36 uint8 = v2
    _ = len(v1)
    v37 := GenS4{F0: 300}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
