package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int = 0
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := []string{"hello", ("\n\t" + "foo bar"), "hello", (((("hello" + "a") + "hello") + "") + "hello")}
    var v6 bool = false
    v7 := GenS1{F0: 1, F1: -9223372036854775808, F2: 42}
    v6 = false
    v6 = v6
    switch ("a" + ("foo bar" + "a")) {
    case "\n\t":
        for fi8 := 0; fi8 < 13; fi8++ {
            if v6 {
                _ = len(v5)
                switch ("a" + "") {
                case "a":
                    var v9 float32 = (3.14159 - (0.0 + 1e10))
                    _ = v9
                case "foo bar":
                    var v10 string = "foo bar"
                    _ = v10
                case "hello":
                    v11 := GenS2{F0: "\n\t", F1: 42}
                    _ = v11
                default:
                    v1 -= 32767
                }
            } else {
                v1 = -2147483648
            }
            fmt.Println(v6, v6, fi8)
            if (v6 && ("" <= "hello")) {
                if (!(32767 < 65535)) {
                    switch "foo bar" {
                    case "foo bar":
                        v1 = (fi8 - 32767)
                    case "":
                        var v12 int16 = 32767
                        _ = v12
                    case "\n\t":
                        for fi13 := 0; fi13 < 3; fi13++ {
                            delete(v2, v1)
                            var v14 uint8 = 128
                            _ = fi13
                            _ = v14
                        }
                    }
                }
            }
            _ = fi8
        }
    case "":
        switch ("hello" + ("a" + "foo bar")) {
        case "hello":
            var v15 string = "a"
            _ = v15
        case "\n\t":
            var v16 int32 = 384
            _ = v16
        case "a":
            v1 -= (v1 | (v1 ^ 731))
        default:
            fmt.Println(v1, v1)
        }
    case "foo bar":
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := GenS3{F0: "hello"}
            _ = fi17
            _ = v18
        }
    default:
        switch "foo bar" {
        case "":
            fmt.Println(v1, v1)
        }
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        switch (v1 + fi19) {
        case 0:
            for fi20 := 0; fi20 < 13; fi20++ {
                fmt.Println(fi19, fi20, fi19)
                _ = fi20
            }
        }
        v21 := v2[v1]
        if (!v6) {
            v5 = append(v5, "")
        } else {
            v22 := GenS2{F0: "foo bar", F1: 2}
            for fi23 := 0; fi23 < 14; fi23++ {
                v21 -= (^(v21))
                delete(v2, v1)
                if false {
                    if v6 {
                        var v24 int8 = -10
                        _ = v24
                    }
                    if v6 {
                        var v25 float32 = (((0.1 + -2.25) - -2.25) - (float32(-2.25) / 1.5))
                        v26 := GenS3{F0: ("a" + ("foo bar" + (("foo bar" + "\n\t") + "\n\t")))}
                        var v27 int64 = 1
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = len(v2)
                } else {
                    v1 += fi23
                    v1 -= fi23
                    if len(v5) > 0 {
                        v5[0] = "foo bar"
                    }
                    v21 += (v21 << ((18446744073709551615) & 63))
                }
                v6 = v6
                _ = fi23
            }
            _ = v22
        }
        _ = fi19
        _ = v21
    }
    v28 := GenS1{F0: 44, F1: -1, F2: 255}
    for fi29 := 0; fi29 < 7; fi29++ {
        if false {
            v30 := GenS4{F0: 1, F1: 2, F2: 42}
            _ = v30
        } else {
            switch fi29 {
            case 2147483647:
                if (!(true || (-2.25 <= ((1e10 / 0.1) * 1e10)))) {
                    if len(v5) > 0 {
                        v5[0] = "hello"
                    }
                }
            }
            switch v1 {
            case 127:
                v31 := GenS5{F0: 0, F1: -54}
                _ = v31
            case 42:
                for fi32 := 0; fi32 < 8; fi32++ {
                    var v33 float32 = (1e10 - 1e10)
                    var v34 uint = 2
                    fmt.Println(v6)
                    if (9223372036854775807 < 9223372036854775807) {
                        v35 := GenS4{F0: 0, F1: 1, F2: 1}
                        _ = v35
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            default:
                fmt.Println(fi29)
            }
            var v36 int64 = 682
            _ = v36
        }
        fmt.Println(v1)
        var v37 bool = false
        _ = fi29
        _ = v37
    }
    fmt.Println(v1, v1)
    v1 = (v1 & v1)
    if v6 {
        for fi38 := 0; fi38 < 9; fi38++ {
            for fi39 := 0; fi39 < 7; fi39++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v40 int64 = -510
                v6 = true
                v41 := GenS3{F0: ""}
                _ = fi39
                _ = v40
                _ = v41
            }
            var v42 int8 = 0
            fmt.Println(fi38, v42, v6)
            _ = fi38
            _ = v42
        }
        var v43 int8 = 18
        _ = v43
    }
    v44 := GenS4{F0: 2, F1: 279, F2: 1}
    if len(v5) > 0 {
        v5[0] = ("\n\t" + (("" + "\n\t") + ""))
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
