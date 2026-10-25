package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 bool }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 string }

func main() {
    var v1 int32 = 42
    var v2 float32 = 1.5
    v3 := []int{1, 4294967295}
    var v4 string = "hello"
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 2}
        v7 := GenS1{F0: 18446744073709551615}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 923 {
    case 32767:
        fmt.Println(v4)
    case 2147483647:
        switch (v4 + (v4 + v4)) {
        case "foo bar":
            if (472 >= 997) {
                fmt.Println(v2, v1)
                v8 := GenS1{F0: 854}
                _ = v8
            }
        case "\n\t":
            var v9 uint16 = 32767
            _ = v9
        }
    case 9223372036854775807:
        v4 = ""
    default:
        var v10 string = "a"
        _ = v10
    }
    if (true && (-128 == -128)) {
        v11 := GenS1{F0: 454}
        var v12 int64 = 0
        _ = v11
        _ = v12
    } else {
        var v13 bool = (0 > 42)
        v4 = ""
        _ = v13
    }
    if false {
        var v14 int64 = -9223372036854775808
        var v15 float64 = (3.14159 + 0.1)
        _ = v14
        _ = v15
    }
    fmt.Println(v4)
    fmt.Println(v2, v1)
    var v16 uint8 = 0
    for fi17 := 0; fi17 < 18; fi17++ {
        if (false && false) {
            v4 = ("" + v4)
            fmt.Println(v2, v2, v1)
        } else {
            var v18 float32 = (((float32((0.0 * ((0.1 / (-2.25 + 0.1)) + 3.14159))) * v2) - 1.5) / 1.0)
            if true {
                v1 -= (v1 & (v1 << ((255) & 63)))
            }
            _ = v18
        }
        v3 = append(v3, 32768)
        _ = fi17
    }
    _ = len(v3)
    if ((390 != 422) || true) {
        v4 = "a"
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            switch "hello" {
            case "a":
                for fi20 := 0; fi20 < 19; fi20++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v16 = 255
                    if ((255 == 4294967295) || false) {
                        _ = len(v3)
                        v21 := GenS2{F0: 0, F1: -128, F2: (true || (65535 >= 4294967295))}
                        v3 = append(v3, 128)
                        var v22 uint16 = 1
                        v23 := GenS3{F0: (42 > v16), F1: -1}
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi20
                }
            }
            _ = fi19
        }
        v4 = v4
    }
    var v24 bool = true
    switch 42 {
    case 4294967295:
        switch 4294967295 {
        case -9223372036854775808:
            switch v4 {
            case "foo bar":
                for fi25 := 0; fi25 < 1; fi25++ {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 string = (v4 + "")
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = fi25
                }
            }
        case 990:
            var v28 uint32 = 42
            _ = v28
        case 0:
            var v29 int16 = -421
            _ = v29
        }
    }
    v30 := GenS4{F0: (v4 + ((v4 + v4) + v4))}
    v2 = (-2.25 * v2)
    switch 42 {
    case -755:
        var v31 uint8 = v16
        _ = v31
    case 32767:
        var v32 uint = 667
        _ = v32
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if len(v3) > 0 {
        v3[0] = 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
