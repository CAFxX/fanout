package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    var v2 bool = (!(("a" + ((("hello" + ((("\n\t" + ("hello" + (("a" + ("" + ("\n\t" + "\n\t"))) + ("" + ("hello" + "hello"))))) + ("foo bar" + "hello")) + "a")) + "a") + "\n\t")) >= "\n\t"))
    var v3 float32 = float32(((-2.25 * 3.14159) + 3.14159))
    var v4 int = 32768
    if v2 {
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 int32 = 42
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            switch 1 {
            case 2147483647:
                var v8 uint32 = 2147483647
                _ = v8
            case 117:
                v2 = (1 == 41)
            }
            var v9 uint16 = 513
            _ = fi7
            _ = v9
        }
    }
    v3 += 1e10
    switch (((("a" + "") + ("hello" + ("hello" + ((("\n\t" + ("foo bar" + ("\n\t" + "foo bar"))) + (("hello" + "a") + "")) + ("hello" + "hello"))))) + ("" + "a")) + "\n\t") {
    case "foo bar":
        var v10 bool = (v2 && (v2 || (3.14159 <= (3.14159 - 1.5))))
        _ = v10
    case "hello":
        fmt.Println(v3)
    case "a":
        v4 = -888
    default:
        for fi11 := 0; fi11 < 12; fi11++ {
            switch (v4 << ((481) & 63)) {
            case 1:
                v4 = (^(fi11))
            default:
                var v12 float64 = 0.0
                _ = v12
            }
            _ = fi11
        }
    }
    v13 := GenS1{F0: 60, F1: 42}
    if v2 {
        var v14 uint32 = 697
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := GenS1{F0: 1, F1: 42}
            for fi17 := 0; fi17 < 1; fi17++ {
                v1 = append(v1, v14)
                var v18 uint64 = 2
                var v19 uint16 = 0
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        if v2 {
            if len(v1) > 0 {
                v1[0] = 523
            }
            var v20 int32 = -124
            v20 += v20
            _ = v20
        } else {
            v3 += (1.5 * ((v3 - ((3.14159 * float32(3.14159)) / float32((1e10 * 0.1)))) - float32(0.0)))
            var v21 uint16 = 57
            _ = v21
        }
    }
    if false {
        v22 := GenS1{F0: -112, F1: 1}
        _ = v22
    } else {
        _ = len(v1)
        for fi23 := 0; fi23 < 5; fi23++ {
            v1 = append(v1, 2147483648)
            for fi24 := 0; fi24 < 6; fi24++ {
                switch "" {
                case "\n\t":
                    v25 := GenS1{F0: -1, F1: 42}
                    _ = v25
                case "":
                    switch "\n\t" {
                    case "\n\t":
                        switch -424 {
                        case 4294967295:
                            switch "foo bar" {
                            case "":
                                var v26 float32 = 1e10
                                _ = v26
                            default:
                                fmt.Println(fi23, fi23, v2)
                            }
                        default:
                            v3 += 1.5
                            fmt.Println(v4, fi24)
                            v2 = true
                        }
                    default:
                        var v27 float32 = (v3 * v3)
                        _ = v27
                    }
                default:
                    if (true || v2) {
                        v28 := GenS1{F0: 37, F1: 772}
                        if len(v1) > 0 {
                            v1[0] = 542
                        }
                        _ = v28
                    }
                }
                if (v2 && ((1e10 - v3) <= v3)) {
                    fmt.Println(v2, v3)
                    var v29 uint = 42
                    var v30 string = "a"
                    v31 := GenS2{F0: 1, F1: 20, F2: 0}
                    var v32 int16 = 42
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = fi24
            }
            if true {
                if v2 {
                    fmt.Println(v4, v2, v4)
                } else {
                    var v33 uint32 = 48
                    v34 := GenS2{F0: 127, F1: 0, F2: 42}
                    switch -752 {
                    case 127:
                        switch v4 {
                        case -671:
                            v3 -= 3.14159
                        }
                    case 128:
                        v4 = 128
                    default:
                        v35 := GenS1{F0: 0, F1: 0}
                        _ = v35
                    }
                    _ = v33
                    _ = v34
                }
            }
            _ = fi23
        }
    }
    v4 += (v4 | (v4 << ((65535) & 63)))
    var v36 uint = 1
    var v37 uint64 = 2
    v38 := GenS3{F0: -1}
    v39 := GenS1{F0: 89, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
