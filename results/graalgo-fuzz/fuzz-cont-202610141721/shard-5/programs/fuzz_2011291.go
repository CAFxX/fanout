package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 string = ("hello" + ("hello" + ""))
    v3 := make(map[uint32]string)
    v4 := []uint32{}
    var v5 bool = (false && (!false))
    v5 = v5
    v6 := GenS1{F0: 0, F1: -1}
    v5 = true
    if (!("\n\t" < ("hello" + ""))) {
        fmt.Println(v5)
    }
    if false {
        v2 = "a"
        v1 += v1
        switch 9223372036854775807 {
        case -670:
            v7 := GenS2{F0: 0, F1: 128, F2: -1}
            _ = v7
        case -9223372036854775808:
            if (42 != 9223372036854775807) {
                var v8 float64 = 1.5
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 int64 = -914
                    delete(v3, 13)
                    switch (^(fi9)) {
                    case 4294967295:
                        v3[546] = v2
                    case 2147483647:
                        fmt.Println(v5, v5)
                    }
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                switch 127 {
                case 4294967295:
                    v2 = v2
                case -1:
                    v5 = ("hello" > ("\n\t" + v2))
                case 330:
                    v5 = true
                default:
                    var v11 string = ("a" + v2)
                    _ = v11
                }
                switch "foo bar" {
                case "\n\t":
                    v12 := GenS3{F0: 409}
                    _ = v12
                case "":
                    v5 = v5
                case "hello":
                    for fi13 := 0; fi13 < 5; fi13++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi13
                    }
                }
            }
        case 4294967295:
            switch 333 {
            case 127:
                for fi14 := 0; fi14 < 20; fi14++ {
                    v1 = 18446744073709551615
                    var v15 uint8 = 71
                    _ = fi14
                    _ = v15
                }
            }
        }
    }
    v4 = append(v4, 434)
    switch v2 {
    case "":
        var v16 uint64 = 2
        _ = v16
    case "hello":
        var v17 int16 = 1000
        _ = v17
    case "foo bar":
        var v18 float64 = (0.1 + (1.5 - float64(float32(-2.25))))
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        _ = len(v3)
        if true {
            v1 = (^(v1))
        }
        _ = fi19
    }
    switch "\n\t" {
    case "a":
        for fi20 := 0; fi20 < 20; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                v1 = (v1 - 18446744073709551615)
                v2 = ""
                _ = fi21
            }
            var v22 float32 = 0.0
            _ = fi20
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS1{F0: 65535, F1: 42}
        if (3.14159 >= 3.14159) {
            var v25 int16 = -981
            for fi26 := 0; fi26 < 16; fi26++ {
                fmt.Println(v25, v25, v5)
                _ = fi26
            }
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    if true {
        v3[404] = ("hello" + v2)
        var v27 uint64 = (^(v1))
        var v28 int32 = 42
        _ = v27
        _ = v28
    } else {
        fmt.Println(v2, v1, v2)
        var v29 string = (v2 + ((v2 + (v2 + ("\n\t" + (v2 + "")))) + (("a" + "hello") + "hello")))
        for fi30 := 0; fi30 < 3; fi30++ {
            v31 := GenS4{F0: -71, F1: -959}
            v32 := GenS5{F0: (65535 != 0), F1: v1, F2: 145}
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    }
    if ((65535 != 2) || v5) {
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 uint32 = 65535
            _ = fi33
            _ = v34
        }
        var v35 int32 = 42
        _ = v35
    } else {
        var v36 uint32 = 129
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
