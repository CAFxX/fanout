package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 bool; F1 uint }

func main() {
    v1 := []string{"hello", "foo bar", (("foo bar" + "") + ("a" + ("foo bar" + "foo bar")))}
    var v2 bool = true
    var v3 int8 = 42
    var v4 uint = 65535
    var v5 string = ((("foo bar" + "foo bar") + (("foo bar" + "a") + "hello")) + "")
    var v6 uint32 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 float64 = float64((-2.25 * -2.25))
    v8 := GenS1{F0: 0, F1: 982, F2: -1}
    if (("a" + v5) > "\n\t") {
        var v9 uint = 65535
        fmt.Println(v3)
        for fi10 := 0; fi10 < 17; fi10++ {
            switch v5 {
            case "\n\t":
                switch -164 {
                case 2147483647:
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                }
            case "a":
                v3 = (v3 + int8(fi10))
            default:
                if (331 >= (fi10 - fi10)) {
                    switch "\n\t" {
                    case "\n\t":
                        var v11 uint = 4294967295
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 10; fi12++ {
                        if false {
                            var v13 bool = v2
                            switch (fi12 ^ fi10) {
                            case 592:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v14 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: (v3 * (v3 >> ((679) & 63)))}
                            _ = v13
                            _ = v14
                        }
                        _ = fi12
                    }
                    v15 := GenS1{F0: 166, F1: 227, F2: (v3 * (v3 << ((v9) & 63)))}
                    _ = v15
                } else {
                    if v2 {
                        if (v5 == ("foo bar" + "foo bar")) {
                            var v16 uint = v9
                            fmt.Println(v7)
                            _ = v16
                        }
                        v17 := GenS2{F0: 127, F1: (("hello" + v5) + ((v5 + (v5 + v5)) + v5))}
                        _ = v17
                    }
                    var v18 int64 = -1
                    if true {
                        var v19 bool = v2
                        switch fi10 {
                        case 128:
                            for fi20 := 0; fi20 < 12; fi20++ {
                                switch fi10 {
                                case -406:
                                    switch (^(fi10)) {
                                    case 0:
                                        v1 = append(v1, (((v5 + (v5 + v5)) + v5) + ((v5 + "\n\t") + (((v5 + v5) + (("foo bar" + "foo bar") + (v5 + v5))) + v5))))
                                    default:
                                        var v21 uint64 = 0
                                        _ = v21
                                    }
                                default:
                                    var v22 float64 = v7
                                    _ = v22
                                }
                                _ = fi20
                            }
                        case 32768:
                            v7 -= 1e10
                        default:
                            var v23 uint16 = uint16(v4)
                            _ = v23
                        }
                        v4 = v4
                        _ = v19
                    }
                    _ = v18
                }
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 int32 = int32(v9)
                v2 = true
                _ = fi24
                _ = v25
            }
            _ = fi10
        }
        _ = v9
    } else {
        v6 = (v6 | 511)
        var v26 int = -9223372036854775808
        _ = v26
    }
    fmt.Println(v2)
    v3 = v3
    var v27 uint8 = 181
    switch v5 {
    case "a":
        var v28 string = v5
        _ = v28
    case "\n\t":
        switch v5 {
        case "foo bar":
            if true {
                v29 := GenS1{F0: 4294967295, F1: 859, F2: (^(v3))}
                _ = v29
            }
        case "":
            var v30 float32 = 0.1
            _ = v30
        case "\n\t":
            v31 := GenS2{F0: v27, F1: v5}
            _ = v31
        }
    default:
        var v32 int64 = 0
        _ = v32
    }
    var v33 string = "\n\t"
    v34 := GenS1{F0: 717, F1: 42, F2: 127}
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS3{F0: true, F1: (v4 >> ((1) & 63))}
        _ = fi35
        _ = v36
    }
    var v37 uint32 = v6
    for fi38 := 0; fi38 < 2; fi38++ {
        v2 = true
        if true {
            var v39 uint16 = 2
            v40 := GenS2{F0: (v27 >> ((uint(v27)) & 63)), F1: v5}
            _ = v39
            _ = v40
        }
        _ = fi38
    }
    var v41 uint8 = (v27 | v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
