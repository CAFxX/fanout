package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }

func main() {
    v1 := []uint32{65535, 736}
    v2 := []string{("\n\t" + (("a" + ("hello" + "foo bar")) + "a")), "\n\t"}
    var v3 string = "\n\t"
    var v4 uint32 = 2147483647
    var v5 int64 = 0
    var v6 int8 = 109
    if true {
        switch 493 {
        case 216:
            v6 += v6
        case 996:
            var v7 uint8 = 128
            _ = v7
        case 4294967295:
            for fi8 := 0; fi8 < 13; fi8++ {
                if (false && false) {
                    v5 -= v5
                    fmt.Println(v4, fi8)
                    if ((int32(v4) >= 0) && true) {
                        var v9 int8 = v6
                        for fi10 := 0; fi10 < 9; fi10++ {
                            var v11 int16 = -812
                            switch v3 {
                            case "hello":
                                var v12 float64 = 0.1
                                _ = v12
                            default:
                                v13 := GenS1{F0: -72, F1: 42}
                                _ = v13
                            }
                            _ = fi10
                            _ = v11
                        }
                        _ = v9
                    } else {
                        v14 := GenS2{F0: 65535}
                        _ = v14
                    }
                } else {
                    switch v3 {
                    case "foo bar":
                        v15 := GenS1{F0: fi8, F1: 732}
                        _ = v15
                    }
                    if false {
                        v16 := GenS2{F0: 600}
                        switch "" {
                        case "foo bar":
                            v5 = v5
                        case "hello":
                            v2 = append(v2, ((v3 + "foo bar") + ((("\n\t" + ("" + v3)) + v3) + ("foo bar" + "a"))))
                        case "\n\t":
                            var v17 int64 = v5
                            _ = v17
                        }
                        var v18 int16 = 1
                        _ = v16
                        _ = v18
                    }
                }
                v1 = append(v1, (v4 ^ 42))
                _ = fi8
            }
        default:
            v5 = (v5 ^ (^(v5)))
        }
        v3 = v3
        var v19 float32 = 1e10
        _ = v19
    } else {
        if len(v2) > 0 {
            v2[0] = v3
        }
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        switch (fi20 + fi20) {
        case 127:
            var v21 uint32 = v4
            _ = v21
        }
        _ = fi20
    }
    if false {
        v5 -= v5
    } else {
        v5 += 9223372036854775807
    }
    var v22 int64 = v5
    v6 = v6
    v3 = ""
    if true {
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                switch (fi24 ^ 295) {
                case 32767:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 962:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 2147483647:
                    v25 := GenS1{F0: fi24, F1: 2}
                    _ = v25
                }
                if (!(false && (false && (false && true)))) {
                    var v26 int = fi23
                    _ = v26
                } else {
                    if true {
                        v22 += 908
                        fmt.Println(v3, v6, v3)
                        var v27 bool = false
                        _ = v27
                    }
                }
                _ = fi24
            }
            _ = fi23
        }
        v28 := GenS3{F0: 42}
        _ = v28
    }
    fmt.Println(v22, v6, v4)
    var v29 int = 9223372036854775807
    switch v29 {
    case 9223372036854775807:
        var v30 uint32 = (v4 - (v4 + v4))
        _ = v30
    case 0:
        v3 = "foo bar"
    default:
        v31 := GenS1{F0: (v29 / ((v29) | 1)), F1: 517}
        _ = v31
    }
    var v32 float32 = 1.5
    switch "" {
    case "":
        switch (v29 << ((630) & 63)) {
        case 826:
            if false {
                fmt.Println(v32, v4)
            } else {
                v6 += v6
                var v33 int16 = -1
                v4 = 184
                _ = v33
            }
        case 32768:
            for fi34 := 0; fi34 < 13; fi34++ {
                v22 = (v5 >> ((60) & 63))
                _ = fi34
            }
        case 127:
            var v35 float64 = ((3.14159 * float64(1.5)) + ((float64(v32) + -2.25) / 1.5))
            _ = v35
        default:
            var v36 uint = 224
            _ = v36
        }
    case "foo bar":
        var v37 int64 = (v5 + 9223372036854775807)
        _ = v37
    default:
        fmt.Println(v5)
    }
    var v38 uint32 = (v4 >> ((381) & 63))
    if (!(!true)) {
        fmt.Println(v6, v3)
        switch "a" {
        case "foo bar":
            v5 -= 55
        default:
            if len(v1) > 0 {
                v1[0] = (v38 ^ 0)
            }
        }
        var v39 float64 = (0.0 - (0.1 - 1.5))
        _ = v39
    }
    v40 := GenS1{F0: v29, F1: uint64(v4)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
