package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := []string{"hello"}
    var v2 int8 = -82
    var v3 int8 = -23
    var v4 uint16 = 65535
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]bool)
    if true {
        var v7 int64 = -9223372036854775808
        v7 = v7
        var v8 uint = 4294967295
        _ = v7
        _ = v8
    }
    fmt.Println(v3)
    _ = len(v6)
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            if false {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                v11 := GenS1{F0: 1, F1: 65535}
                var v12 int64 = -1
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            v2 += 100
            v3 += 58
            _ = fi13
        }
        _ = fi9
    }
    v4 = (v4 >> ((2) & 63))
    var v14 float64 = float64(-2.25)
    v15 := v6[1]
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            if (v3 < v3) {
                v2 += -128
                var v18 int32 = 2147483647
                _ = v18
            } else {
                fmt.Println(v3, v2)
                var v19 int = fi17
                var v20 int = 1
                _ = v19
                _ = v20
            }
            var v21 uint8 = 1
            _ = fi17
            _ = v21
        }
        var v22 float64 = 1e10
        switch "" {
        case "hello":
            v2 = (v3 << ((553) & 63))
        default:
            switch fi16 {
            case 32767:
                var v23 string = ("hello" + "")
                _ = v23
            case 127:
                var v24 uint64 = 18446744073709551615
                _ = v24
            default:
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 int32 = 2147483647
                    if v15 {
                        v27 := v5[536]
                        var v28 float32 = 1e10
                        var v29 float64 = float64(3.14159)
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        var v30 int = (fi25 >> ((1) & 63))
                        switch "foo bar" {
                        case "":
                            var v31 uint32 = 298
                            _ = v31
                        case "a":
                            fmt.Println(fi25, v22, v15)
                        }
                        _ = v30
                    }
                    fmt.Println(v14, v26)
                    _ = fi25
                    _ = v26
                }
            }
        }
        _ = fi16
        _ = v22
    }
    switch 521 {
    case -749:
        var v32 int32 = 823
        _ = v32
    case 0:
        v33 := v6[0]
        _ = v33
    default:
        for fi34 := 0; fi34 < 6; fi34++ {
            for fi35 := 0; fi35 < 8; fi35++ {
                switch "hello" {
                case "\n\t":
                    if (124 < 128) {
                        fmt.Println(v4)
                        var v36 uint8 = 153
                        var v37 uint = 18446744073709551615
                        _ = v36
                        _ = v37
                    }
                case "":
                    v15 = v15
                }
                v1 = append(v1, "")
                _ = fi35
            }
            _ = fi34
        }
    }
    v38 := GenS1{F0: 127, F1: v4}
    var v39 float32 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
