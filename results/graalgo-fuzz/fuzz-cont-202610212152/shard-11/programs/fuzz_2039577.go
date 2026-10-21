package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS3 struct { F0 int; F1 int; F2 uint64 }
type GenS4 struct { F0 string; F1 int16; F2 bool }
func genFn0(p0 int8, p1 int) uint16 {
    var v5 int8 = p0
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: 42, F1: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 858
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{42, 2}
    var v3 int32 = -1
    var v4 uint = 590
    fmt.Println(v3, v1)
    v8 := GenS1{F0: -112, F1: -464}
    switch -1 {
    case 94:
        switch 514 {
        case 0:
            _ = len(v2)
        case 700:
            _ = len(v2)
        case -131:
            v2 = append(v2, 153)
        default:
            if true {
                var v9 uint64 = 2
                v10 := GenS2{F0: 867, F1: 2, F2: -32768}
                _ = v10
                _ = v9
            } else {
                var v11 int64 = (v1 | v1)
                _ = v11
            }
        }
    case 9223372036854775807:
        v1 -= v1
    }
    if false {
        fmt.Println(v3, v3, v4)
    }
    switch -1 {
    case 851:
        var v12 int8 = 42
        _ = v12
    case -9223372036854775808:
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    case 32768:
        switch ("hello" + ("a" + ("\n\t" + "hello"))) {
        case "foo bar":
            _ = len(v2)
        case "":
            fmt.Println(v4)
        default:
            if (!true) {
                v13 := GenS3{F0: -1, F1: 4294967295, F2: 2}
                _ = v13
            }
        }
    default:
        v3 = 220
    }
    fmt.Println(v1)
    switch ("foo bar" + "\n\t") {
    case "a":
        var v14 uint8 = 127
        _ = v14
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        if true {
            var v15 int8 = 94
            _ = v15
        } else {
            for fi16 := 0; fi16 < 17; fi16++ {
                if true {
                    _ = len(v2)
                    v1 = v1
                }
                for fi17 := 0; fi17 < 17; fi17++ {
                    if (!(0.0 >= (0.0 * 0.0))) {
                        v18 := GenS1{F0: -786, F1: 1}
                        var v19 int8 = 127
                        _ = v18
                        _ = v19
                    }
                    var v20 int16 = 0
                    var v21 uint64 = 2
                    v1 = v1
                    v20 += int16(v4)
                    _ = fi17
                    _ = v20
                    _ = v21
                }
                v4 = (v4 | v4)
                _ = fi16
            }
            var v22 bool = (true || (-32768 >= 42))
            _ = v22
        }
    default:
        v1 = v1
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        if false {
            v1 = (v1 - v1)
        }
        v24 := GenS1{F0: -1, F1: -879}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 17; fi26++ {
            fmt.Println(v3, fi26)
            if (!false) {
                fmt.Println(v4, v1)
                if false {
                    var v27 string = (("" + (("" + "a") + "\n\t")) + (("a" + ("\n\t" + "")) + "foo bar"))
                    v1 = (v1 * v1)
                    v28 := GenS2{F0: 2, F1: 989, F2: 1000}
                    switch (-(fi26)) {
                    case 32768:
                        _ = len(v2)
                    }
                    _ = v27
                    _ = v28
                }
                v29 := GenS2{F0: 1, F1: 65535, F2: -32768}
                _ = v29
            } else {
                var v30 uint64 = 18446744073709551615
                var v31 uint16 = 65535
                switch "\n\t" {
                case "hello":
                    v32 := GenS4{F0: "hello", F1: -741, F2: (-257 < 32767)}
                    _ = v32
                case "\n\t":
                    v31 = v31
                }
                switch (fi26 & int(v3)) {
                case 127:
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                }
                _ = v30
                _ = v31
            }
            _ = fi26
        }
        switch (fi25 ^ 2147483647) {
        case 127:
            fmt.Println(fi25, v4, v1)
        }
        v1 = (v1 * (v1 + -1))
        _ = fi25
    }
    var v33 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v8
}
