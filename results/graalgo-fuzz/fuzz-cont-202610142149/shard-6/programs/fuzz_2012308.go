package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int16; F1 uint64; F2 int8 }

func main() {
    v1 := []string{"\n\t", ("a" + ("hello" + (("\n\t" + "") + ""))), (((("" + ("\n\t" + "foo bar")) + ("a" + "")) + "a") + ("\n\t" + "\n\t"))}
    var v2 float32 = 1e10
    var v3 uint32 = 42
    for fi4 := 0; fi4 < 20; fi4++ {
        v3 = (^(v3))
        if true {
            var v5 float64 = 3.14159
            _ = v5
        }
        var v6 int = fi4
        _ = fi4
        _ = v6
    }
    var v7 float32 = v2
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 2, F1: 843}
        v3 -= (^(v3))
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 42
    var v11 int = -779
    _ = len(v1)
    var v12 int32 = 1
    switch (v11 ^ 32767) {
    case 128:
        v3 = v3
    case 32768:
        for fi13 := 0; fi13 < 1; fi13++ {
            fmt.Println(v2)
            var v14 uint16 = uint16(v3)
            v14 += v14
            _ = fi13
            _ = v14
        }
    case 4294967295:
        v10 = v10
    }
    if false {
        var v15 int = (v11 + 2147483647)
        var v16 int8 = -1
        if true {
            for fi17 := 0; fi17 < 2; fi17++ {
                if (((v16 >> ((0) & 63)) >= v16) && false) {
                    switch fi17 {
                    case 32768:
                        v12 -= (v12 & (v12 ^ v12))
                    }
                    var v18 uint16 = 0
                    _ = v18
                } else {
                    var v19 uint64 = 322
                    v20 := GenS2{F0: uint(v19), F1: true}
                    _ = v19
                    _ = v20
                }
                var v21 int = int(v12)
                if true {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v15 += fi17
                        v15 = fi17
                        var v23 int32 = -1
                        _ = fi22
                        _ = v23
                    }
                } else {
                    fmt.Println(v3, v3, v2)
                    v16 = v16
                }
                _ = fi17
                _ = v21
            }
            var v24 bool = (255 == 760)
            _ = v24
        }
        _ = v15
        _ = v16
    } else {
        var v25 uint64 = uint64(v10)
        var v26 int8 = -40
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 uint32 = (^(v3))
        v29 := GenS1{F0: 10, F1: v12}
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := GenS2{F0: 42, F1: (!(!true))}
    switch (v11 >> ((1) & 63)) {
    case 42:
        switch (("foo bar" + "\n\t") + (("foo bar" + "hello") + ("a" + "\n\t"))) {
        case "foo bar":
            if (!false) {
                if (!true) {
                    _ = len(v1)
                    var v31 int64 = -13
                    switch v11 {
                    case 805:
                        var v32 uint16 = 42
                        _ = v32
                    case 9223372036854775807:
                        switch "\n\t" {
                        case "hello":
                            v2 += -2.25
                        case "a":
                            for fi33 := 0; fi33 < 7; fi33++ {
                                v11 -= v11
                                v10 += 42
                                for fi34 := 0; fi34 < 15; fi34++ {
                                    fmt.Println(v2, v3)
                                    _ = fi34
                                }
                                _ = fi33
                            }
                        case "\n\t":
                            v35 := GenS3{F0: 42}
                            _ = v35
                        }
                    case 32768:
                        fmt.Println(v11, v7)
                    default:
                        switch (v11 - 0) {
                        case -227:
                            v36 := GenS4{F0: 127, F1: int16(v31)}
                            _ = v36
                        case -628:
                            _ = len(v1)
                        }
                    }
                    _ = v31
                }
                v37 := GenS5{F0: 1, F1: 4294967295, F2: -61}
                _ = v37
            }
        case "hello":
            v12 += -581
        case "\n\t":
            fmt.Println(v12, v12, v2)
        }
    case 9223372036854775807:
        var v38 float64 = (1.5 - float64(0.0))
        _ = v38
    }
    if true {
        var v39 int16 = 1
        _ = v39
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch (v11 - (v11 + v11)) {
    case -1:
        var v40 int32 = v12
        _ = v40
    case 127:
        var v41 uint64 = 4294967295
        _ = v41
    case 573:
        var v42 uint8 = v10
        _ = v42
    default:
        fmt.Println(v3, v11)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v7
}
