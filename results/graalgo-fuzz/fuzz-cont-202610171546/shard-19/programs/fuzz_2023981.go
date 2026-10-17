package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int; F1 bool; F2 uint64 }
type GenS6 struct { F0 string; F1 int64; F2 string }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    var v3 int8 = 93
    var v4 float32 = ((1e10 / 1e10) + 1e10)
    var v5 float32 = v4
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 2, F1: 0}
    if true {
        v5 = v5
        v8 := []int{}
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                switch 128 {
                case 128:
                    fmt.Println(v4, fi9, v3)
                }
                var v11 int = fi10
                var v12 bool = true
                _ = fi10
                _ = v11
                _ = v12
            }
            switch ("foo bar" + ((("\n\t" + "") + "") + ("foo bar" + "\n\t"))) {
            case "\n\t":
                switch (-(fi9)) {
                case 641:
                    fmt.Println(v5, fi9)
                case -675:
                    if false {
                        var v13 int = (fi9 + fi9)
                        if (1 != 127) {
                            switch "a" {
                            case "a":
                                v4 = float32((float64(((-2.25 / (0.1 * (v5 + v4))) - (v5 / 1.0))) - 0.1))
                            case "\n\t":
                                v5 = v4
                            case "foo bar":
                                var v14 int16 = 32767
                                _ = v14
                            default:
                                v1 = v1
                            }
                            switch int(v3) {
                            case 128:
                                v15 := GenS1{F0: 324, F1: 1}
                                _ = v15
                            case -757:
                                v1 = ((!(1 >= 100)) || v1)
                            }
                            v8 = append(v8, v13)
                        }
                        _ = v13
                    }
                }
            case "":
                for fi16 := 0; fi16 < 16; fi16++ {
                    if v1 {
                        fmt.Println(v5, v4)
                        var v17 int = (fi16 << ((281) & 63))
                        _ = v17
                    }
                    v5 = v5
                    switch "" {
                    case "hello":
                        v18 := GenS1{F0: 1, F1: 42}
                        _ = v18
                    default:
                        v19 := GenS2{F0: -849, F1: -367}
                        _ = v19
                    }
                    _ = fi16
                }
            case "hello":
                switch (fi9 % (((fi9 * fi9)) | 1)) {
                case -796:
                    delete(v2, 1)
                }
            default:
                var v20 float32 = v5
                _ = v20
            }
            var v21 uint64 = 42
            _ = fi9
            _ = v21
        }
        _ = v8
    }
    if (v1 && ((("hello" + ("foo bar" + (("foo bar" + ("" + ("hello" + "foo bar"))) + "a"))) + (((("foo bar" + "\n\t") + "a") + (("foo bar" + "foo bar") + ("foo bar" + "foo bar"))) + (((((("a" + "") + "hello") + "hello") + ("\n\t" + (("a" + "hello") + "\n\t"))) + "\n\t") + "\n\t"))) == ("" + "hello"))) {
        v22 := GenS3{F0: "", F1: "a"}
        v3 = 100
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int64 = -47
        fmt.Println(v24, fi23)
        _ = fi23
        _ = v24
    }
    v6[628] = 453
    switch "" {
    case "a":
        fmt.Println(v1)
    case "foo bar":
        v25 := v2[0]
        _ = v25
    case "\n\t":
        v3 = 0
    }
    if false {
        v26 := []int{-9223372036854775808, -2147483648}
        _ = v26
    } else {
        var v27 uint = 2
        if false {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                var v28 int = 42
                _ = v28
            case "hello":
                switch 32768 {
                case -446:
                    var v29 int = 892
                    _ = v29
                case 1:
                    var v30 int = 4294967295
                    _ = v30
                default:
                    switch (((("hello" + "\n\t") + "a") + "") + ("a" + "hello")) {
                    case "hello":
                        v31 := GenS3{F0: "\n\t", F1: ("\n\t" + "a")}
                        _ = v31
                    case "\n\t":
                        var v32 string = "hello"
                        _ = v32
                    case "a":
                        v33 := GenS4{F0: 0, F1: 42}
                        _ = v33
                    }
                }
            case "":
                v4 -= (3.14159 * float32(((((((0.1 / ((((((-2.25 * -2.25) + 0.0) / ((((3.14159 * float64(1e10)) - 1e10) - 3.14159) + 0.0)) - float64(v5)) / -2.25) - 1e10)) * 3.14159) * float64(float32((((1.5 + -2.25) * (3.14159 * 0.0)) + 1.5)))) + 1e10) - (-2.25 + (-2.25 / 1.0))) * 1.5)))
            }
            v3 -= -115
        }
        _ = v27
    }
    v34 := GenS5{F0: -736, F1: (619 < 42), F2: 0}
    v35 := []int{}
    fmt.Println(v4, v1, v1)
    if (0 >= 0) {
        v36 := GenS6{F0: "", F1: 1, F2: "a"}
        fmt.Println(v1)
        _ = v36
    } else {
        var v37 string = ("" + ("" + "foo bar"))
        if len(v35) > 0 {
            v35[0] = 2147483647
        }
        v35 = append(v35, -805)
        _ = v37
    }
    fmt.Println(v4)
    v3 = v3
    v38 := GenS1{F0: 2, F1: 0}
    v35 = append(v35, 4294967295)
    var v39 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
