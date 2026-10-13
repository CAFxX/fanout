package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int32 }
type GenS2 struct { F0 int32; F1 int32; F2 uint }
type GenS3 struct { F0 uint; F1 int8; F2 string }

func main() {
    var v1 uint32 = 363
    v2 := []string{(("foo bar" + ("\n\t" + "hello")) + "a"), ""}
    var v3 int = 622
    v4 := []uint32{}
    var v5 string = (("" + "a") + ("" + "a"))
    switch 1 {
    case 42:
        switch v3 {
        case 403:
            var v6 int16 = 0
            _ = v6
        case 2147483647:
            fmt.Println(v1, v3)
        case 570:
            var v7 string = v5
            _ = v7
        default:
            var v8 int = v3
            _ = v8
        }
    case 9223372036854775807:
        if len(v4) > 0 {
            v4[0] = v1
        }
    default:
        var v9 uint = 416
        _ = v9
    }
    fmt.Println(v1, v3)
    fmt.Println(v3, v5, v1)
    var v10 uint16 = 2
    var v11 string = v5
    for fi12 := 0; fi12 < 10; fi12++ {
        if true {
            var v13 int8 = 26
            _ = v13
        }
        _ = fi12
    }
    switch int(v10) {
    case 32768:
        for fi14 := 0; fi14 < 4; fi14++ {
            switch (fi14 >> ((117) & 63)) {
            case -647:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                for fi15 := 0; fi15 < 14; fi15++ {
                    switch (v3 + (fi14 / ((fi15) | 1))) {
                    case 128:
                        fmt.Println(fi14, v1)
                    case 127:
                        v16 := GenS1{F0: 255, F1: -1, F2: 243}
                        _ = v16
                    default:
                        var v17 uint8 = 128
                        fmt.Println(v10)
                        fmt.Println(v17, fi15, v3)
                        var v18 int64 = 9223372036854775807
                        _ = v17
                        _ = v18
                    }
                    v3 -= (fi15 % ((1) | 1))
                    _ = fi15
                }
            }
            v19 := GenS1{F0: 255, F1: 50, F2: 0}
            _ = fi14
            _ = v19
        }
    }
    var v20 uint16 = 287
    if len(v2) > 0 {
        v2[0] = v5
    }
    if (true || true) {
        v5 = (((v5 + v5) + ("" + ("" + v5))) + "foo bar")
        var v21 int = v3
        _ = v21
    } else {
        var v22 uint32 = v1
        switch "hello" {
        case "foo bar":
            var v23 uint16 = uint16(v22)
            _ = v23
        case "hello":
            var v24 int16 = 32767
            _ = v24
        case "":
            for fi25 := 0; fi25 < 13; fi25++ {
                for fi26 := 0; fi26 < 12; fi26++ {
                    if (true && (true || true)) {
                        v5 = "hello"
                    }
                    if ((0.1 / 1e10) > 1.5) {
                        var v27 uint32 = v1
                        if false {
                            v28 := GenS2{F0: 2147483647, F1: -306, F2: 4294967295}
                            var v29 int = 32767
                            _ = v28
                            _ = v29
                        } else {
                            v22 = v22
                            if false {
                                var v30 uint64 = 18446744073709551615
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v30
                            }
                            v27 += 2147483647
                            fmt.Println(v1)
                        }
                        if false {
                            var v31 uint = 18446744073709551615
                            _ = v31
                        }
                        v32 := GenS2{F0: 1, F1: 42, F2: 18446744073709551615}
                        _ = v27
                        _ = v32
                    } else {
                        v5 = ("\n\t" + v11)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi26
                }
                if false {
                    for fi33 := 0; fi33 < 11; fi33++ {
                        fmt.Println(v5)
                        v34 := GenS2{F0: -1, F1: 42, F2: 2}
                        v1 = v22
                        _ = fi33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 6; fi35++ {
                        v22 = v1
                        fmt.Println(v22, v11)
                        _ = fi35
                    }
                    fmt.Println(v1)
                }
                v36 := GenS3{F0: 0, F1: 1, F2: v11}
                _ = fi25
                _ = v36
            }
        }
        v37 := GenS2{F0: 0, F1: -2147483648, F2: 18446744073709551615}
        _ = v22
        _ = v37
    }
    if true {
        var v38 int64 = 1
        _ = v38
    }
    v1 -= 398
    v5 = ("" + "")
    v3 += v3
    fmt.Println(v20, v11, v5)
    var v39 int8 = -1
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
