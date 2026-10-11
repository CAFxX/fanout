package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 string }

func main() {
    var v1 int = -1
    var v2 uint8 = 1
    v3 := []string{"hello", "a", (("hello" + ((("foo bar" + "foo bar") + "\n\t") + (("foo bar" + "a") + "foo bar"))) + "hello"), "\n\t"}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (v1 & -634) {
    case 32768:
        fmt.Println(v1)
    case 233:
        fmt.Println(v2, v2)
    case -1:
        v4 := GenS1{F0: -406}
        _ = v4
    }
    var v5 uint8 = 255
    var v6 int8 = 42
    _ = len(v3)
    var v7 uint32 = 649
    v2 -= v2
    v8 := GenS2{F0: 4294967295}
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            if false {
                fmt.Println(fi10, fi9, v7)
                v6 -= (v6 % ((v6) | 1))
            }
            v11 := GenS1{F0: -1}
            _ = fi10
            _ = v11
        }
        if true {
            var v12 int32 = -1
            for fi13 := 0; fi13 < 19; fi13++ {
                if true {
                    v12 = (v12 % ((v12) | 1))
                }
                _ = fi13
            }
            var v14 bool = false
            _ = v12
            _ = v14
        } else {
            v6 = v6
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 float64 = (1e10 - 0.0)
                fmt.Println(v2, v7, fi9)
                if (0 != 2) {
                    fmt.Println(v16)
                } else {
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v2)
                    case "foo bar":
                        switch "a" {
                        case "foo bar":
                            var v17 int8 = (-(v6))
                            _ = v17
                        case "":
                            var v18 int8 = (-(v6))
                            _ = v18
                        }
                    case "a":
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                    }
                    if (675 != -32768) {
                        v19 := GenS2{F0: 4294967295}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v19
                    } else {
                        var v20 float64 = (float64(float32(1e10)) + v16)
                        v3 = append(v3, "hello")
                        _ = v20
                    }
                }
                _ = fi15
                _ = v16
            }
        }
        var v21 uint64 = 269
        _ = fi9
        _ = v21
    }
    if (65535 == 397) {
        var v22 int8 = 48
        fmt.Println(v7, v7)
        _ = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS3{F0: false}
        if false {
            switch "" {
            case "foo bar":
                v25 := GenS4{F0: (v6 % ((v6) | 1))}
                _ = v25
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                for fi26 := 0; fi26 < 11; fi26++ {
                    fmt.Println(v7)
                    _ = fi26
                }
            }
            if ((0.0 < 3.14159) && true) {
                v27 := GenS4{F0: v6}
                switch (v1 - (fi23 ^ (fi23 & fi23))) {
                case 32767:
                    v28 := GenS4{F0: (v6 & (^(v6)))}
                    _ = v28
                case 9223372036854775807:
                    for fi29 := 0; fi29 < 8; fi29++ {
                        if false {
                            v30 := GenS5{F0: 42, F1: 1}
                            _ = len(v3)
                            _ = v30
                        } else {
                            var v31 int = 32767
                            _ = v31
                        }
                        _ = fi29
                    }
                default:
                    v32 := GenS6{F0: "hello"}
                    _ = v32
                }
                _ = v27
            }
        } else {
            var v33 int32 = 42
            switch (v1 >> ((876) & 63)) {
            case 2147483647:
                if true {
                    v1 = (fi23 + (v1 + -9223372036854775808))
                }
            }
            _ = v33
        }
        _ = fi23
        _ = v24
    }
    if (true || false) {
        var v34 int32 = 0
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
