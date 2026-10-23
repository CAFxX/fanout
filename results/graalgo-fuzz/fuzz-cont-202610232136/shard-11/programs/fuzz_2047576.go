package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := []string{"\n\t", "a", (("foo bar" + "hello") + ("" + "hello"))}
    var v4 string = "\n\t"
    for fi5 := 0; fi5 < 4; fi5++ {
        if (fi5 <= fi5) {
            v6 := GenS1{F0: 2, F1: -98}
            if len(v3) > 0 {
                v3[0] = v4
            }
            if true {
                for fi7 := 0; fi7 < 5; fi7++ {
                    delete(v1, (((v4 + v4) + "a") + "foo bar"))
                    fmt.Println(fi7, v4)
                    if (false || true) {
                        var v8 bool = false
                        _ = v8
                    } else {
                        var v9 uint64 = 4294967295
                        _ = v9
                    }
                    switch v4 {
                    case "a":
                        var v10 uint32 = 127
                        _ = v10
                    default:
                        fmt.Println(fi5, fi7, fi5)
                    }
                    _ = fi7
                }
                v4 = "foo bar"
            }
            _ = v6
        } else {
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := GenS1{F0: 208, F1: 42}
                if len(v3) > 0 {
                    v3[0] = ((("hello" + "a") + v4) + (v4 + v4))
                }
                var v13 int32 = -517
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v4 = v4
        for fi14 := 0; fi14 < 10; fi14++ {
            _ = len(v3)
            if (!(!false)) {
                if (0.0 > ((1e10 + (((-2.25 - (1.5 + (0.1 - ((1.5 - 0.0) * 0.0)))) * (1.5 - -2.25)) - ((1e10 + float64((-2.25 / 3.14159))) * 1.5))) * 0.0)) {
                    v4 = v4
                    fmt.Println(fi5, fi5, fi5)
                } else {
                    v15 := GenS1{F0: 46, F1: 0}
                    if false {
                        v4 = (("a" + v4) + "\n\t")
                        v3 = append(v3, (v4 + ("\n\t" + "hello")))
                        var v16 string = "hello"
                        _ = v16
                    }
                    _ = v15
                }
                v17 := GenS2{F0: (-(fi14))}
                switch -9223372036854775808 {
                case 42:
                    if (-1 > int32(fi14)) {
                        var v18 float32 = (3.14159 - 0.0)
                        v19 := GenS2{F0: (-(fi5))}
                        _ = len(v3)
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint64 = 4294967295
                        _ = len(v2)
                        _ = v20
                    }
                case 0:
                    switch (fi14 >> ((513) & 63)) {
                    case 1:
                        fmt.Println(v4, v4, v4)
                    case -1:
                        _ = len(v3)
                    case 4294967295:
                        var v21 float64 = 1.5
                        _ = v21
                    }
                }
                _ = v17
            }
            if (true || false) {
                v22 := GenS1{F0: 4294967295, F1: -128}
                _ = v22
            }
            _ = fi14
        }
        _ = fi5
    }
    delete(v2, 4294967295)
    var v23 float32 = (0.0 - 0.1)
    var v24 int = 1
    switch v24 {
    case -2147483648:
        v24 = v24
    default:
        var v25 uint8 = 128
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint64 = 42
        v28 := GenS1{F0: 4294967295, F1: 42}
        v29 := GenS3{F0: 0, F1: (^(v27))}
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    if false {
        fmt.Println(v4, v4)
        switch "" {
        case "a":
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 string = "a"
                _ = fi30
                _ = v31
            }
        }
    }
    if false {
        var v32 int = v24
        _ = v32
    }
    v33 := GenS3{F0: 65535, F1: 4294967295}
    v4 = v4
    if false {
        switch 0 {
        case 1:
            v34 := GenS1{F0: 927, F1: 42}
            _ = v34
        case 32768:
            switch v24 {
            case -647:
                v35 := GenS4{F0: 42, F1: v4}
                _ = v35
            default:
                if (true && false) {
                    var v36 int16 = -32768
                    switch v4 {
                    case "hello":
                        v37 := GenS5{F0: -128}
                        _ = v37
                    case "foo bar":
                        for fi38 := 0; fi38 < 14; fi38++ {
                            var v39 int16 = -312
                            var v40 float32 = (v23 + 0.0)
                            v41 := GenS2{F0: 32767}
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    case "a":
                        v42 := GenS6{F0: 127, F1: 2, F2: 32767}
                        _ = v42
                    default:
                        if (v23 != v23) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    if false {
                        var v43 uint32 = 976
                        _ = v43
                    }
                    _ = v36
                }
            }
        case 9223372036854775807:
            v23 = v23
        }
        var v44 uint32 = 0
        var v45 float64 = 1.5
        _ = v44
        _ = v45
    }
    fmt.Println(v24, v4)
    var v46 bool = (!true)
    for fi47 := 0; fi47 < 9; fi47++ {
        v46 = false
        _ = fi47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v46
}
