package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }

func main() {
    var v1 uint32 = 940
    var v2 uint16 = 859
    var v3 uint = 2
    v4 := []string{("a" + "hello"), "foo bar", "a"}
    v5 := GenS1{F0: 131, F1: 164}
    if true {
        if true {
            if false {
                var v6 string = "foo bar"
                switch ("" + v6) {
                case "hello":
                    v3 += (v3 + (v3 / (((v3 % ((v3) | 1))) | 1)))
                case "\n\t":
                    var v7 int16 = 1000
                    _ = v7
                case "a":
                    var v8 float64 = ((1e10 + (0.1 * 3.14159)) * -2.25)
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 bool = ((v3 * v3) == 255)
                        v4 = append(v4, "a")
                        _ = fi9
                        _ = v10
                    }
                }
                _ = v6
            }
            _ = len(v4)
        }
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int = (fi11 | (fi11 >> ((v3) & 63)))
            v12 -= v12
            _ = fi11
            _ = v12
        }
    }
    v2 += 0
    var v13 uint64 = 18446744073709551615
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint64 = v13
        switch (fi14 * (fi14 / (((fi14 * 1)) | 1))) {
        case 324:
            var v16 int16 = -816
            _ = v16
        case -1:
            if ((9223372036854775807 >= 127) || false) {
                for fi17 := 0; fi17 < 11; fi17++ {
                    var v18 string = "a"
                    v15 = v15
                    _ = fi17
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 16; fi19++ {
                    fmt.Println(fi19, v3, v2)
                    switch (("foo bar" + ((("foo bar" + "\n\t") + ("" + "\n\t")) + (("hello" + "foo bar") + "foo bar"))) + "a") {
                    case "a":
                        fmt.Println(v2, v13)
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    if false {
                        if false {
                            var v20 bool = false
                            v21 := GenS1{F0: (v15 + (v15 ^ v15)), F1: 2}
                            var v22 float32 = 3.14159
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        if ((v3 * (v3 << ((v3) & 63))) != (v3 << (((v3 % (((v3 / (((v3 ^ (^(v3)))) | 1))) | 1))) & 63))) {
                            var v23 int64 = -1
                            _ = v23
                        }
                    }
                    _ = fi19
                }
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                switch "\n\t" {
                case "hello":
                    for fi24 := 0; fi24 < 17; fi24++ {
                        if len(v4) > 0 {
                            v4[0] = ("hello" + "foo bar")
                        }
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        v25 := GenS2{F0: -614}
                        _ = fi24
                        _ = v25
                    }
                case "\n\t":
                    _ = len(v4)
                case "foo bar":
                    fmt.Println(v15)
                default:
                    for fi26 := 0; fi26 < 5; fi26++ {
                        var v27 string = ""
                        fmt.Println(v27)
                        v28 := GenS3{F0: 1000, F1: v1}
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
            }
        case 712:
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 bool = false
                v31 := GenS2{F0: 42}
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        v32 := GenS4{F0: v3, F1: (v13 - v13), F2: -128}
        _ = fi14
        _ = v15
        _ = v32
    }
    v2 = (^(v2))
    v3 += (^(v3))
    v3 = 255
    for fi33 := 0; fi33 < 17; fi33++ {
        v4 = append(v4, "foo bar")
        _ = fi33
    }
    v34 := GenS4{F0: (^(v3)), F1: v13, F2: -35}
    if true {
        var v35 float32 = float32(1.5)
        _ = v35
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v36 := GenS4{F0: uint(v2), F1: v13, F2: 42}
    var v37 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
