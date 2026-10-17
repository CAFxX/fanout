package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8; F1 int16; F2 int }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 bool; F1 int8 }

func main() {
    v1 := []int{1}
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 1
    var v4 int16 = 32767
    if ((!(true && (391 >= 506))) || (42 >= 767)) {
        for fi5 := 0; fi5 < 13; fi5++ {
            switch (fi5 / (((fi5 - (fi5 << ((uint(v3)) & 63)))) | 1)) {
            case -9223372036854775808:
                var v6 float32 = 0.0
                _ = v6
            case 132:
                for fi7 := 0; fi7 < 10; fi7++ {
                    switch ("" + ("a" + ("foo bar" + "a"))) {
                    case "foo bar":
                        v8 := GenS1{F0: true, F1: 255}
                        _ = v8
                    case "hello":
                        v2[v3] = v3
                    case "a":
                        var v9 int32 = 1
                        _ = v9
                    }
                    if false {
                        v4 = 1
                    }
                    _ = fi7
                }
            default:
                fmt.Println(v3, fi5)
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS1{F0: false, F1: 4294967295}
                if false {
                    var v12 int64 = -1
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            switch "a" {
            case "\n\t":
                var v13 int16 = v4
                _ = v13
            case "hello":
                v3 = v3
            case "foo bar":
                v4 -= (v4 & (v4 % ((v4) | 1)))
            }
            _ = fi5
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v14 := GenS1{F0: (true || false), F1: 18446744073709551615}
        _ = v14
    }
    v15 := v2[(v3 >> ((48) & 63))]
    v3 += (v15 ^ (v3 << ((42) & 63)))
    delete(v2, v3)
    switch "foo bar" {
    case "\n\t":
        switch -9223372036854775808 {
        case -199:
            v2[498] = v15
        }
    case "a":
        for fi16 := 0; fi16 < 2; fi16++ {
            if true {
                _ = len(v1)
                v15 = v15
            } else {
                v1 = append(v1, (fi16 << ((5) & 63)))
            }
            _ = len(v1)
            if (!true) {
                v17 := GenS1{F0: true, F1: 65535}
                var v18 uint = 452
                fmt.Println(v15)
                _ = v17
                _ = v18
            }
            _ = fi16
        }
    case "":
        v19 := GenS2{F0: 127, F1: (-(v4)), F2: 128}
        _ = v19
    default:
        var v20 int = 119
        _ = v20
    }
    var v21 float32 = 0.1
    var v22 uint16 = 201
    v4 += (v4 + v4)
    if false {
        v1 = append(v1, -191)
        for fi23 := 0; fi23 < 14; fi23++ {
            fmt.Println(v3, v3)
            v24 := GenS3{F0: 170, F1: -73}
            _ = fi23
            _ = v24
        }
    } else {
        switch "\n\t" {
        case "":
            v25 := GenS2{F0: int8(v3), F1: v4, F2: -664}
            _ = v25
        default:
            if true {
                v26 := GenS1{F0: true, F1: 1}
                _ = v26
            }
        }
        v22 -= 255
        switch 4294967295 {
        case -95:
            var v27 int = 0
            _ = v27
        case 286:
            v28 := GenS2{F0: -123, F1: (v4 | v4), F2: 1}
            _ = v28
        default:
            v29 := GenS4{F0: 2, F1: 839}
            _ = v29
        }
    }
    var v30 float32 = v21
    if (true || (!false)) {
        switch -2147483648 {
        case 42:
            switch ("foo bar" + ("foo bar" + ("" + ("\n\t" + "hello")))) {
            case "a":
                fmt.Println(v30)
            }
        case 127:
            switch ((("\n\t" + "foo bar") + "hello") + "hello") {
            case "a":
                v30 = v30
            case "":
                v30 += (v21 + float32(1e10))
            case "foo bar":
                v31 := GenS3{F0: 2147483647, F1: 0}
                _ = v31
            }
        case 367:
            v32 := GenS5{F0: ((201 > 255) || true), F1: 42}
            _ = v32
        default:
            v1 = append(v1, 914)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v4
}
