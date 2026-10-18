package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int8; F2 string }
type GenS5 struct { F0 int; F1 uint16; F2 int32 }
type GenS6 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint8 = 143
    var v2 int8 = 1
    v3 := []uint32{2147483647}
    v4 := GenS1{F0: 0, F1: v1, F2: uint64(v1)}
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS2{F0: 42}
        switch (fi5 + fi5) {
        case 0:
            v7 := GenS2{F0: uint32(v1)}
            _ = v7
        case 9223372036854775807:
            var v8 uint = 278
            _ = v8
        case 127:
            var v9 uint16 = 2
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    v10 := GenS3{F0: 42}
    if true {
        var v11 int8 = v2
        _ = len(v3)
        _ = v11
    }
    if (!true) {
        if (false || true) {
            fmt.Println(v2, v1)
            var v12 int8 = v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v12
        } else {
            var v13 uint16 = 32767
            var v14 uint64 = 33
            _ = v13
            _ = v14
        }
        v1 = v1
    } else {
        var v15 int8 = 100
        v2 = (v2 - 37)
        for fi16 := 0; fi16 < 14; fi16++ {
            v15 = v2
            v3 = append(v3, 777)
            _ = fi16
        }
        _ = v15
    }
    v3 = append(v3, 2)
    _ = len(v3)
    _ = len(v3)
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v1, v2)
        v18 := GenS2{F0: 877}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS1{F0: 255, F1: (v1 & v1), F2: 2}
        _ = fi19
        _ = v20
    }
    switch "foo bar" {
    case "a":
        v21 := GenS3{F0: 2}
        _ = v21
    case "":
        fmt.Println(v1, v2)
    }
    switch "hello" {
    case "hello":
        fmt.Println(v1)
    case "foo bar":
        switch "a" {
        case "a":
            if (1.5 >= (1.5 * 1e10)) {
                var v22 uint32 = 42
                _ = v22
            }
        case "":
            v23 := GenS2{F0: 2147483647}
            _ = v23
        }
    default:
        if (!(4294967295 < 351)) {
            _ = len(v3)
            v2 -= (-(v2))
        }
    }
    v3 = append(v3, 2147483648)
    for fi24 := 0; fi24 < 12; fi24++ {
        fmt.Println(v2, v2, v2)
        fmt.Println(v2)
        for fi25 := 0; fi25 < 3; fi25++ {
            v2 = (^(v2))
            _ = fi25
        }
        _ = fi24
    }
    if (-9223372036854775808 != 42) {
        var v26 int16 = 1000
        if ((v1 & v1) > (v1 / (((v1 % (((v1 + (v1 ^ v1))) | 1))) | 1))) {
            v27 := GenS4{F0: 862, F1: v2, F2: "hello"}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v28 := GenS5{F0: 42, F1: 0, F2: 1}
            _ = v27
            _ = v28
        }
        _ = v26
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v29 bool = false
        for fi30 := 0; fi30 < 15; fi30++ {
            if v29 {
                v2 = v2
                var v31 string = "foo bar"
                _ = v31
            }
            v32 := GenS5{F0: fi30, F1: 65535, F2: 42}
            _ = fi30
            _ = v32
        }
        _ = v29
    }
    if true {
        var v33 string = ("hello" + "foo bar")
        v1 -= (v1 * (v1 / (((v1 % ((255) | 1))) | 1)))
        _ = v33
    }
    if false {
        v34 := GenS6{F0: 127, F1: 65535}
        _ = v34
    }
    v1 -= (v1 + 120)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
