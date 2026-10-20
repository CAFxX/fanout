package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 0
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    v5 := []uint32{305}
    var v6 float32 = 3.14159
    var v7 int16 = 780
    switch -1 {
    case 32767:
        if len(v5) > 0 {
            v5[0] = 65535
        }
    case 128:
        var v8 uint = 42
        _ = v8
    case 9223372036854775807:
        for fi9 := 0; fi9 < 14; fi9++ {
            switch (fi9 | 39) {
            case -1:
                v7 = (v7 + (v7 ^ v7))
            case -827:
                v3[42] = (fi9 & fi9)
            }
            v1["foo bar"] = false
            _ = fi9
        }
    }
    if len(v5) > 0 {
        v5[0] = 639
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        switch ("" + "\n\t") {
        case "\n\t":
            for fi11 := 0; fi11 < 2; fi11++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v3[4294967295] = (fi10 ^ fi11)
                v7 = (v7 * v7)
                _ = fi11
            }
        }
        _ = fi10
    }
    fmt.Println(v2, v7)
    v7 = (v7 / (((v7 + (v7 & (v7 % ((v7) | 1))))) | 1))
    if (true || false) {
        var v12 int = 0
        switch "" {
        case "\n\t":
            var v13 int32 = 2147483647
            _ = v13
        default:
            var v14 int32 = 2147483647
            _ = v14
        }
        if true {
            fmt.Println(v12)
            var v15 float32 = float32((0.1 / 0.1))
            _ = v15
        } else {
            v16 := GenS1{F0: "foo bar"}
            _ = v16
        }
        _ = v12
    } else {
        var v17 uint = 755
        fmt.Println(v2, v2, v17)
        _ = v17
    }
    var v18 float64 = 0.0
    var v19 uint16 = 694
    switch 0 {
    case -9223372036854775808:
        switch -411 {
        case 2147483647:
            var v20 uint32 = 65535
            _ = v20
        }
    case 288:
        v6 = v6
    case -2147483648:
        v5 = append(v5, 42)
    default:
        v21 := GenS2{F0: -2147483648}
        _ = v21
    }
    if len(v5) > 0 {
        v5[0] = 951
    }
    switch "" {
    case "hello":
        var v22 uint8 = 0
        _ = v22
    default:
        if ((v19 ^ 42) <= (^(v19))) {
            if len(v5) > 0 {
                v5[0] = 65535
            }
            var v23 string = (("a" + ("\n\t" + (("foo bar" + "foo bar") + ("" + "\n\t")))) + ("hello" + ((("hello" + "a") + "foo bar") + "hello")))
            _ = v23
        } else {
            v24 := GenS3{F0: false, F1: v19, F2: -276}
            switch -1 {
            case 128:
                fmt.Println(v6, v2, v18)
            case 0:
                switch -15 {
                case 127:
                    v25 := GenS3{F0: ("hello" <= ("" + (("foo bar" + "a") + "a"))), F1: v19, F2: -1}
                    _ = v25
                case 359:
                    _ = len(v5)
                case -9223372036854775808:
                    var v26 float64 = v18
                    _ = v26
                default:
                    switch -376 {
                    case 677:
                        v19 -= v19
                    case 4294967295:
                        v7 -= (v7 | (v7 % (((v7 & v7)) | 1)))
                    default:
                        delete(v1, "foo bar")
                    }
                }
            case -2147483648:
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 int32 = -2147483648
                    _ = fi27
                    _ = v28
                }
            }
            _ = v24
        }
    }
    v7 = (v7 - v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
