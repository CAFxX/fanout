package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]int)
    v3 := []string{"hello", "hello", "hello", "foo bar"}
    var v4 bool = (false && (0 > 2))
    v5 := make(map[string]bool)
    var v6 int16 = 32767
    v7 := v5[(("a" + "\n\t") + "hello")]
    var v8 uint32 = 65535
    v9 := GenS1{F0: "a"}
    for fi10 := 0; fi10 < 14; fi10++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v3 = append(v3, ("a" + (("\n\t" + "foo bar") + "\n\t")))
        _ = fi11
    }
    v7 = (false || ((-1 <= -1) || false))
    switch 9223372036854775807 {
    case 4294967295:
        v12 := GenS2{F0: (v8 / (((v8 << ((71) & 63))) | 1)), F1: -128, F2: uint8(v8)}
        _ = v12
    case 32767:
        if true {
            fmt.Println(v1, v4, v8)
        }
    case 9223372036854775807:
        var v13 int32 = 940
        _ = v13
    }
    fmt.Println(v1, v8, v1)
    delete(v5, "a")
    fmt.Println(v1, v8, v8)
    if (v4 && false) {
        switch "a" {
        case "foo bar":
            _ = len(v3)
        case "":
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 int64 = 42
                _ = fi14
                _ = v15
            }
        case "a":
            var v16 float64 = 0.1
            _ = v16
        default:
            if (!true) {
                for fi17 := 0; fi17 < 7; fi17++ {
                    if v4 {
                        var v18 int64 = 1
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v20 := GenS2{F0: (v8 / ((v8) | 1)), F1: 73, F2: 127}
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        v6 = 0
                        v4 = (("" + ("a" + ("" + "foo bar"))) < "\n\t")
                    }
                    for fi21 := 0; fi21 < 13; fi21++ {
                        v1 = 2
                        v22 := GenS1{F0: ""}
                        _ = fi21
                        _ = v22
                    }
                    _ = fi17
                }
            }
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            switch "\n\t" {
            case "foo bar":
                for fi24 := 0; fi24 < 11; fi24++ {
                    v8 -= 42
                    if v4 {
                        var v25 string = (((("foo bar" + "hello") + "hello") + "a") + (("\n\t" + ((("a" + "\n\t") + ("a" + "hello")) + ("" + "\n\t"))) + "a"))
                        _ = v25
                    } else {
                        v26 := GenS3{F0: 42, F1: fi23}
                        var v27 uint = 255
                        fmt.Println(fi23)
                        _ = v26
                        _ = v27
                    }
                    _ = fi24
                }
            default:
                for fi28 := 0; fi28 < 8; fi28++ {
                    if (v4 && true) {
                        switch "" {
                        case "foo bar":
                            var v29 int16 = v6
                            _ = v29
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                        case "\n\t":
                            if (false || false) {
                                switch (fi23 << ((2) & 63)) {
                                case -2147483648:
                                    var v30 float64 = float64((3.14159 + 1e10))
                                    _ = v30
                                case 128:
                                    switch 2147483647 {
                                    case -84:
                                        var v31 float64 = float64((float32((0.0 - 0.0)) - ((1.5 + 1.5) + -2.25)))
                                        _ = v31
                                    }
                                case 0:
                                    if false {
                                        fmt.Println(v7, v7, v6)
                                        switch fi28 {
                                        case 570:
                                            var v32 uint = 18446744073709551615
                                            _ = v32
                                        case -1:
                                            v33 := GenS4{F0: 100}
                                            _ = v33
                                        case 32767:
                                            v2[(v8 ^ v8)] = fi28
                                        }
                                    } else {
                                        _ = len(v3)
                                    }
                                default:
                                    v34 := GenS1{F0: "foo bar"}
                                    _ = v34
                                }
                                v6 = v6
                                v6 -= v6
                                var v35 bool = true
                                v36 := GenS2{F0: 42, F1: -128, F2: 132}
                                v8 -= v8
                                _ = v35
                                _ = v36
                            }
                        }
                        v1 = v1
                        var v37 int64 = 871
                        _ = v37
                    } else {
                        fmt.Println(v6, fi28)
                        v8 = 2147483647
                        switch "\n\t" {
                        case "a":
                            switch "foo bar" {
                            case "":
                                var v38 int16 = (v6 ^ v6)
                                _ = v38
                            }
                        }
                        v5["foo bar"] = (!v7)
                        var v39 uint8 = 39
                        _ = v39
                    }
                    v7 = true
                    _ = fi28
                }
            }
            var v40 string = ("a" + "\n\t")
            _ = fi23
            _ = v40
        }
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 uint16 = (v1 >> ((835) & 63))
            _ = fi41
            _ = v42
        }
    } else {
        var v43 uint8 = 42
        v8 = (v8 << ((65535) & 63))
        var v44 int64 = -9223372036854775808
        _ = v43
        _ = v44
    }
    var v45 int = -1
    v6 = (v6 << ((255) & 63))
    var v46 int = 9223372036854775807
    v6 = (-(v6))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
