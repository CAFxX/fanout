package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 uint64 = 42
    v2 := []uint32{2, 520, 1, 996}
    var v3 uint = 944
    var v4 int8 = 127
    var v5 int32 = -741
    fmt.Println(v4, v5, v3)
    if false {
        fmt.Println(v5)
        switch 4294967295 {
        case 32768:
            if false {
                if false {
                    v1 = (v1 * v1)
                }
                var v6 float64 = (1e10 * 0.1)
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 uint8 = 1
                    switch "" {
                    case "foo bar":
                        v9 := GenS1{F0: (!(false || true)), F1: 128}
                        _ = v9
                    case "":
                        fmt.Println(v6, v3)
                    default:
                        v10 := GenS2{F0: 162, F1: (v4 + (v4 ^ (v4 >> (((^(v3))) & 63)))), F2: 4294967295}
                        _ = v10
                    }
                    _ = fi7
                    _ = v8
                }
                _ = v6
            } else {
                v2 = append(v2, 2147483647)
                v5 = int32(v4)
                for fi11 := 0; fi11 < 4; fi11++ {
                    v12 := GenS3{F0: (v4 == v4)}
                    v3 = 492
                    _ = fi11
                    _ = v12
                }
            }
        case 1:
            switch 821 {
            case 127:
                v13 := GenS4{F0: (!true), F1: "\n\t"}
                _ = v13
            }
        case 2147483647:
            var v14 int64 = 42
            _ = v14
        }
        v15 := GenS2{F0: 107, F1: (^(v4)), F2: 18446744073709551615}
        _ = v15
    } else {
        var v16 uint8 = 2
        var v17 int64 = 42
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: ("hello" >= "hello")}
    v5 = 1
    v2 = append(v2, 139)
    v1 = v1
    v19 := GenS1{F0: (true || false), F1: -439}
    fmt.Println(v5)
    var v20 uint8 = 255
    var v21 int64 = -1
    var v22 uint8 = 2
    if true {
        v4 = (v4 >> (((v3 / ((v3) | 1))) & 63))
        var v23 float64 = 3.14159
        _ = v23
    } else {
        var v24 int = 127
        _ = v24
    }
    var v25 int = -440
    v20 = v22
    switch (("a" + "a") + "hello") {
    case "foo bar":
        if (!(!false)) {
            v26 := GenS1{F0: (1 == 473), F1: (v25 + (v25 + v25))}
            switch "foo bar" {
            case "hello":
                v2 = append(v2, 2147483647)
            case "a":
                v20 -= v20
            default:
                if false {
                    if len(v2) > 0 {
                        v2[0] = 651
                    }
                }
            }
            v22 = (v20 ^ 128)
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        v3 = v3
        switch fi27 {
        case 255:
            if false {
                var v28 int = -429
                var v29 int8 = (v4 | v4)
                if (true && true) {
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    for fi30 := 0; fi30 < 14; fi30++ {
                        var v31 uint64 = (^(v1))
                        _ = fi30
                        _ = v31
                    }
                    if (v4 >= int8(v21)) {
                        fmt.Println(v4)
                        var v32 int64 = 9223372036854775807
                        var v33 string = "a"
                        _ = v32
                        _ = v33
                    } else {
                        v34 := GenS5{F0: (v20 % (((v20 & 0)) | 1))}
                        v35 := GenS2{F0: 2, F1: (v29 >> (((v3 - v3)) & 63)), F2: v1}
                        fmt.Println(v28)
                        _ = v34
                        _ = v35
                    }
                }
                _ = v28
                _ = v29
            } else {
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v36 float64 = 0.0
                    _ = v36
                }
            }
        default:
            var v37 int16 = 1000
            _ = v37
        }
        _ = fi27
    }
    v20 = 127
    v38 := GenS6{F0: v4}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
