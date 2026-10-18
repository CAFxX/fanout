package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int32; F2 int64 }

func main() {
    v1 := []string{"\n\t"}
    var v2 uint = 457
    var v3 string = ""
    var v4 string = ""
    var v5 int64 = 9223372036854775807
    fmt.Println(v5)
    if false {
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                if (((^(fi6)) <= fi7) || true) {
                    fmt.Println(fi7, v5)
                    var v8 uint = (v2 + v2)
                    if false {
                        v8 += v8
                        v4 = (v4 + (v4 + v3))
                    } else {
                        if (127 != -1) {
                            v9 := GenS1{F0: (v4 + v4), F1: 42, F2: 183}
                            fmt.Println(fi6, fi6, fi7)
                            fmt.Println(v2, v4)
                            _ = v9
                        } else {
                            v8 = v2
                            var v10 int32 = -2147483648
                            _ = v10
                        }
                    }
                    _ = v8
                }
                fmt.Println(v4, fi7, v5)
                if ((1.5 - ((3.14159 + 0.0) / -2.25)) == ((1e10 * 1e10) / float64(((1.5 / float32((3.14159 / 3.14159))) * 3.14159)))) {
                    var v11 uint8 = 2
                    _ = v11
                } else {
                    var v12 uint8 = 123
                    _ = v12
                }
                _ = fi7
            }
            _ = fi6
        }
        fmt.Println(v4)
    }
    v5 += v5
    var v13 int16 = 0
    if (0 == 42) {
        var v14 int8 = 94
        fmt.Println(v14, v13)
        v13 = -983
        _ = v14
    } else {
        switch 127 {
        case -1:
            var v15 int32 = -456
            _ = v15
        case -214:
            v16 := GenS1{F0: v3, F1: 2, F2: 128}
            _ = v16
        }
        switch (("hello" + "foo bar") + v3) {
        case "foo bar":
            v4 = (((v3 + ((("hello" + ("a" + "\n\t")) + v3) + "a")) + v3) + ("hello" + v3))
        default:
            for fi17 := 0; fi17 < 19; fi17++ {
                _ = len(v1)
                if (v2 > v2) {
                    if true {
                        if (true && (!false)) {
                            v18 := GenS1{F0: v3, F1: 1, F2: 0}
                            _ = v18
                        }
                    }
                    var v19 uint = 0
                    _ = v19
                } else {
                    var v20 int16 = -189
                    _ = v20
                }
                _ = fi17
            }
        }
        switch -485 {
        case 1:
            v21 := GenS1{F0: v4, F1: uint8(v2), F2: 88}
            _ = v21
        case 42:
            v13 = (v13 << ((303) & 63))
        case 128:
            var v22 int8 = 0
            _ = v22
        }
    }
    var v23 int8 = 42
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int32 = 1
        _ = fi24
        _ = v25
    }
    var v26 string = "hello"
    fmt.Println(v3)
    fmt.Println(v4)
    switch (((v3 + v3) + "") + v3) {
    case "foo bar":
        var v27 uint8 = uint8(v2)
        _ = v27
    case "":
        _ = len(v1)
    }
    switch -674 {
    case 1:
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 float32 = (1e10 - float32(1.5))
            if true {
                fmt.Println(fi28)
            } else {
                v23 = (v23 >> ((0) & 63))
                for fi30 := 0; fi30 < 16; fi30++ {
                    switch fi30 {
                    case 4294967295:
                        _ = len(v1)
                    }
                    _ = fi30
                }
                _ = len(v1)
            }
            _ = fi28
            _ = v29
        }
    case 0:
        var v31 uint16 = 255
        _ = v31
    case 4294967295:
        v32 := GenS2{F0: int16(v23), F1: (^(v23))}
        _ = v32
    }
    v33 := GenS3{F0: uint8(v2), F1: 18446744073709551615}
    if len(v1) > 0 {
        v1[0] = v3
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        switch v26 {
        case "\n\t":
            v35 := GenS4{F0: (("" + ("a" + "")) + (v26 + ((v3 + v26) + v26))), F1: 42}
            _ = v35
        case "foo bar":
            for fi36 := 0; fi36 < 3; fi36++ {
                fmt.Println(v5, v2, v5)
                _ = fi36
            }
        default:
            v37 := GenS5{F0: 127, F1: -413, F2: int64(fi34)}
            _ = v37
        }
        var v38 uint32 = 0
        var v39 uint16 = 1
        _ = fi34
        _ = v38
        _ = v39
    }
    switch v4 {
    case "foo bar":
        v26 = ("a" + "a")
    case "a":
        v13 = v13
    case "":
        v13 = (-(v13))
    }
    var v40 int64 = v5
    if true {
        v40 -= v5
    } else {
        v1 = append(v1, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v5
}
