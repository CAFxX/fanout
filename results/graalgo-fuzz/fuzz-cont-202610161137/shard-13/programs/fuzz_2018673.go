package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint = 255
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 16; fi4++ {
        fmt.Println(fi4, fi4)
        if false {
            if (0.0 <= float64(3.14159)) {
                fmt.Println(fi4, v1)
                if false {
                    for fi5 := 0; fi5 < 19; fi5++ {
                        var v6 int = 32767
                        _ = fi5
                        _ = v6
                    }
                } else {
                    v1 -= v1
                    if true {
                        var v7 string = ("\n\t" + "hello")
                        v7 = "foo bar"
                        var v8 int8 = 95
                        _ = v7
                        _ = v8
                    } else {
                        var v9 uint = v1
                        _ = v9
                    }
                    v10 := GenS1{F0: fi4, F1: (("foo bar" + "\n\t") + "a")}
                    _ = v10
                }
                var v11 string = "a"
                _ = v11
            }
            switch (fi4 % ((fi4) | 1)) {
            case -2147483648:
                _ = len(v2)
            }
            v1 = (v1 + v1)
        }
        _ = fi4
    }
    v12 := GenS1{F0: 42, F1: ""}
    fmt.Println(v1)
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS2{F0: 127, F1: 1000}
        v15 := GenS1{F0: (fi13 / (((fi13 ^ 168)) | 1)), F1: "hello"}
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int32 = 42
            switch "a" {
            case "hello":
                var v18 int32 = v17
                _ = v18
            case "":
                if (v17 != (v17 - v17)) {
                    switch "foo bar" {
                    case "":
                        switch "foo bar" {
                        case "foo bar":
                            var v19 uint = 4294967295
                            _ = v19
                        case "\n\t":
                            var v20 bool = false
                            _ = v20
                        }
                    case "\n\t":
                        var v21 uint64 = 1
                        _ = v21
                    case "hello":
                        var v22 int8 = -105
                        _ = v22
                    }
                    var v23 uint8 = 0
                    _ = v23
                }
            default:
                switch "foo bar" {
                case "\n\t":
                    var v24 int = 32767
                    _ = v24
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            v1 = (v1 >> ((78) & 63))
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v25 uint16 = 507
    if (!true) {
        fmt.Println(v25, v25)
        switch -9223372036854775808 {
        case 0:
            fmt.Println(v1, v1)
        }
        var v26 uint16 = (^(v25))
        _ = v26
    } else {
        v27 := GenS1{F0: -1, F1: "a"}
        _ = v27
    }
    var v28 bool = (!true)
    var v29 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
}
