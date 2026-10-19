package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 string; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 int8; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 string; F1 int; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 70
    v3 := []string{"a", "\n\t", ("a" + "\n\t"), "hello"}
    var v4 uint = 65535
    v5 := make(map[string]uint32)
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    v6 := GenS1{F0: 42}
    v7 := GenS2{F0: (!(!(true && (!(-2147483648 <= 4294967295))))), F1: "foo bar", F2: (0 <= 0)}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS3{F0: 524}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int = fi10
        var v12 int = (v11 | 173)
        if true {
            v13 := GenS1{F0: 9223372036854775807}
            v14 := v5["a"]
            if false {
                fmt.Println(v14, fi10)
                var v15 uint = 128
                _ = v15
            }
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    switch ("\n\t" + ("" + "hello")) {
    case "\n\t":
        v16 := GenS3{F0: 32767}
        _ = v16
    case "":
        delete(v1, ("\n\t" + ("foo bar" + "hello")))
    case "foo bar":
        for fi17 := 0; fi17 < 2; fi17++ {
            fmt.Println(v4, v2, fi17)
            var v18 string = "hello"
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        delete(v5, "\n\t")
        if false {
            switch "\n\t" {
            case "hello":
                for fi20 := 0; fi20 < 14; fi20++ {
                    if true {
                        _ = len(v3)
                        v2 = 286
                    }
                    v21 := GenS4{F0: 42, F1: 4294967295, F2: -779}
                    _ = fi20
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 14; fi22++ {
                    fmt.Println(fi19, fi22, v2)
                    _ = fi22
                }
            }
        }
        _ = fi19
    }
    var v23 uint = (v4 % ((4294967295) | 1))
    v3 = append(v3, "foo bar")
    fmt.Println(v2, v23)
    fmt.Println(v2, v2)
    if (!(0 <= -2147483648)) {
        v24 := GenS5{F0: 142}
        _ = v24
    }
    switch 32768 {
    case 128:
        var v25 int = -1
        _ = v25
    case -2147483648:
        fmt.Println(v2)
    }
    if (-1 != 42) {
        delete(v5, "\n\t")
        fmt.Println(v2, v2, v2)
        fmt.Println(v2, v23, v2)
    } else {
        var v26 int16 = 32767
        fmt.Println(v23, v26, v26)
        fmt.Println(v4)
        _ = v26
    }
    if false {
        if false {
            v27 := GenS3{F0: 0}
            _ = v27
        } else {
            v3 = append(v3, ("foo bar" + ""))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    var v28 int32 = 0
    if ((v23 / ((v4) | 1)) == (v4 + (v23 + 88))) {
        if (!(!("a" <= "foo bar"))) {
            v29 := GenS6{F0: "", F1: 52, F2: 127}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 11; fi30++ {
                switch "\n\t" {
                case "\n\t":
                    if false {
                        var v31 uint16 = 227
                        _ = v31
                    }
                case "hello":
                    v5["\n\t"] = 2147483647
                case "foo bar":
                    var v32 int16 = 1
                    _ = v32
                }
                _ = fi30
            }
            var v33 uint8 = 2
            fmt.Println(v4)
            _ = v33
        }
    } else {
        var v34 string = (("\n\t" + ("foo bar" + "hello")) + "a")
        if (true || true) {
            if true {
                v35 := GenS7{F0: (v28 | (v28 + (v28 ^ (v28 << (((v4 / (((v23 << (((v4 ^ 826)) & 63))) | 1))) & 63))))), F1: 0, F2: (v2 / (((v2 | v2)) | 1))}
                switch 42 {
                case 162:
                    switch 32767 {
                    case 896:
                        var v36 float32 = -2.25
                        _ = v36
                    case 2147483647:
                        fmt.Println(v23)
                    }
                case -31:
                    v37 := GenS8{F0: v2, F1: v2}
                    _ = v37
                case 24:
                    v38 := GenS9{F0: ("foo bar" + v34), F1: 66, F2: v2}
                    _ = v38
                }
                v2 = (v2 + v2)
                _ = v35
            }
            v28 = (v28 - (v28 / (((v28 >> (((^(v23))) & 63))) | 1)))
        } else {
            var v39 uint32 = 4294967295
            var v40 int32 = (v28 << (((v4 ^ v23)) & 63))
            var v41 uint = (v23 % ((109) | 1))
            _ = v39
            _ = v40
            _ = v41
        }
        var v42 int32 = (v28 - (^(v28)))
        _ = v34
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
