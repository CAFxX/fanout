package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint }

func main() {
    var v1 uint8 = 100
    v2 := []uint32{2, 2147483647, 1}
    var v3 float32 = -2.25
    var v4 uint = 4294967295
    _ = len(v2)
    if false {
        switch "hello" {
        case "":
            v5 := GenS1{F0: 721}
            _ = v5
        }
        _ = len(v2)
    }
    var v6 int16 = 1000
    switch 1 {
    case 9223372036854775807:
        v1 = (v1 >> (((v4 % ((1) | 1))) & 63))
    case -866:
        v7 := GenS1{F0: 1}
        _ = v7
    }
    var v8 float32 = v3
    v4 = (v4 >> ((v4) & 63))
    switch 128 {
    case 127:
        for fi9 := 0; fi9 < 7; fi9++ {
            if true {
                switch (("foo bar" + ("\n\t" + ("foo bar" + "hello"))) + ("\n\t" + "\n\t")) {
                case "":
                    var v10 int16 = v6
                    _ = v10
                case "hello":
                    fmt.Println(v8, v4, v1)
                case "a":
                    var v11 float64 = 1.5
                    _ = v11
                default:
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                v13 := GenS1{F0: 1}
                if false {
                    v14 := GenS2{F0: v4, F1: 42}
                    switch fi9 {
                    case 32768:
                        v8 = v3
                    }
                    _ = v14
                }
                _ = v13
            } else {
                var v15 int32 = -990
                var v16 string = "\n\t"
                var v17 int8 = int8(v6)
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi9
        }
    case 128:
        v18 := GenS1{F0: uint16(v1)}
        _ = v18
    case -1:
        var v19 float64 = 3.14159
        _ = v19
    }
    v20 := GenS2{F0: (v4 * 18446744073709551615), F1: v1}
    v8 = 1e10
    for fi21 := 0; fi21 < 1; fi21++ {
        if false {
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 int8 = -1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 float32 = float32((3.14159 * 0.1))
                if true {
                    v3 = v25
                    v25 -= 0.1
                    v6 = (v6 & 1)
                }
                _ = fi24
                _ = v25
            }
            if ((false && (4294967295 < 1)) || (v6 >= v6)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    switch fi21 {
                    case 128:
                        v1 += 109
                    }
                    v26 := GenS1{F0: 42}
                    v8 = -2.25
                    _ = v26
                }
                v3 = v3
            } else {
                v6 += (v6 >> ((4294967295) & 63))
            }
        }
        for fi27 := 0; fi27 < 9; fi27++ {
            if len(v2) > 0 {
                v2[0] = uint32(v4)
            }
            switch "" {
            case "hello":
                v28 := GenS3{F0: 2, F1: 65535}
                _ = v28
            case "\n\t":
                var v29 int64 = 393
                _ = v29
            }
            _ = fi27
        }
        _ = fi21
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v4 = (^(v4))
    case "foo bar":
        v6 += v6
    case "hello":
        for fi30 := 0; fi30 < 1; fi30++ {
            var v31 int8 = 1
            _ = fi30
            _ = v31
        }
    default:
        var v32 float32 = -2.25
        _ = v32
    }
    v33 := GenS2{F0: 18446744073709551615, F1: v1}
    if (-542 > 1) {
        _ = len(v2)
    }
    v1 += v1
    var v34 uint16 = 2
    for fi35 := 0; fi35 < 4; fi35++ {
        fmt.Println(v6, v4, v8)
        fmt.Println(v4, fi35, fi35)
        fmt.Println(v34, v4, v4)
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v6
    _ = v8
}
