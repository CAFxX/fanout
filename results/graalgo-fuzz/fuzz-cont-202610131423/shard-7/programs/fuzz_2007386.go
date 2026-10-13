package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint; F2 uint64 }
func genFn0(p0 int) bool {
    v5 := GenS1{F0: (128 != 663), F1: 255}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return (true || false)
}

func main() {
    v1 := []string{"\n\t", ("hello" + "foo bar"), "\n\t"}
    v2 := make(map[uint32]bool)
    var v3 int16 = -1
    var v4 bool = false
    v4 = true
    fmt.Println(v3, v4)
    var v6 uint64 = 1
    if false {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v4, v3)
            var v7 uint = 489
            _ = v7
        } else {
            v4 = (!((v3 % (((v3 - (^(v3)))) | 1)) != (^(v3))))
            fmt.Println(v6, v6)
            v8 := v2[uint32(v6)]
            _ = v8
        }
    }
    v6 = (v6 | (^(v6)))
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 float32 = 1e10
        v10 -= ((0.1 + float32((1.5 / (((0.1 + 0.0) - 1e10) + 1e10)))) / 1.0)
        var v11 uint64 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    v6 = 42
    v4 = v4
    if v4 {
        var v12 int = 9223372036854775807
        switch "hello" {
        case "\n\t":
            var v13 uint64 = 42
            _ = v13
        case "a":
            var v14 int64 = -271
            _ = v14
        }
        v12 = v12
        _ = v12
    }
    v15 := GenS2{F0: 42, F1: 1, F2: 32767}
    for fi16 := 0; fi16 < 15; fi16++ {
        _ = len(v1)
        _ = len(v1)
        _ = fi16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        switch (fi17 / (((fi17 - fi17)) | 1)) {
        case 127:
            switch ("" + "hello") {
            case "a":
                for fi18 := 0; fi18 < 17; fi18++ {
                    if (!v4) {
                        v19 := GenS2{F0: fi17, F1: 2, F2: 65535}
                        _ = v19
                    }
                    v20 := GenS2{F0: (fi17 | (fi18 % (((-(fi18))) | 1))), F1: 2147483647, F2: 32767}
                    _ = fi18
                    _ = v20
                }
            }
        case 248:
            switch (fi17 << ((445) & 63)) {
            case 1:
                var v21 uint32 = 694
                _ = v21
            case -687:
                v3 -= (v3 | v3)
            case 32767:
                var v22 uint16 = 0
                _ = v22
            default:
                var v23 float64 = 3.14159
                _ = v23
            }
        case -259:
            switch fi17 {
            case 167:
                fmt.Println(v6, v4)
            case 9223372036854775807:
                v24 := GenS3{F0: 2}
                _ = v24
            case -9223372036854775808:
                delete(v2, 1)
            default:
                fmt.Println(v3, fi17, v6)
            }
        default:
            v25 := GenS4{F0: 127, F1: 65535, F2: v6}
            _ = v25
        }
        _ = fi17
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        v6 -= v6
        if (((3.14159 - (0.1 * 3.14159)) + 0.1) >= ((0.1 + -2.25) * 1e10)) {
            if v4 {
                v27 := GenS2{F0: fi26, F1: 12, F2: 84}
                var v28 uint32 = 2
                fmt.Println(v4)
                _ = v27
                _ = v28
            } else {
                if false {
                    fmt.Println(v4)
                    fmt.Println(v4, v6, v4)
                    var v29 uint64 = (v6 | v6)
                    _ = v29
                }
                v6 -= (v6 << ((18446744073709551615) & 63))
            }
            delete(v2, 2147483647)
        } else {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            var v30 float32 = 1e10
            _ = v30
        }
        fmt.Println(v6, fi26)
        _ = fi26
    }
    if (42 <= 1) {
        var v31 int64 = 42
        _ = v31
    } else {
        v4 = v4
        v4 = v4
        _ = len(v2)
    }
    if (!(!((0.1 < 0.1) && (!v4)))) {
        v32 := GenS3{F0: 95}
        delete(v2, 208)
        _ = v32
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        fmt.Println(v3, v4)
        var v34 int32 = 2147483647
        _ = fi33
        _ = v34
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
