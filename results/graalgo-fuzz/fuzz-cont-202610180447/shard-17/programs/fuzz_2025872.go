package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint64; F2 bool }
func genFn0() uint32 {
    if false {
        var v5 uint16 = 1
        v5 = 686
        for fi6 := 0; fi6 < 20; fi6++ {
            v5 += 1
            var v7 uint = 4294967295
            fmt.Println(v5, fi6, fi6)
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    var v8 uint32 = 0
    _ = v8
    return v8
}

func main() {
    var v1 bool = (1 == 4294967295)
    var v2 uint16 = 857
    v3 := []uint32{940, 4294967295, 1, 674}
    var v4 uint = 65535
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 9223372036854775807, F1: v2}
        v11 := GenS1{F0: -313, F1: (v2 >> ((v4) & 63))}
        v4 = (v4 + v4)
        _ = fi9
        _ = v10
        _ = v11
    }
    if (("foo bar" != "hello") || (32767 > 1000)) {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint8 = uint8(v4)
            _ = fi12
            _ = v13
        }
        switch ("\n\t" + "\n\t") {
        case "":
            if (v1 && v1) {
                var v14 int32 = 42
                v15 := GenS1{F0: 863, F1: (v2 << (((v4 | 861)) & 63))}
                if v1 {
                    v3 = append(v3, 4294967295)
                    if false {
                        var v16 uint64 = 42
                        _ = v16
                    }
                    switch 42 {
                    case 32768:
                        fmt.Println(v1, v14)
                    default:
                        _ = len(v3)
                    }
                } else {
                    v17 := GenS1{F0: 42, F1: (v2 | (v2 ^ (v2 / (((v2 - 829)) | 1))))}
                    fmt.Println(v14, v2, v2)
                    _ = v17
                }
                _ = v14
                _ = v15
            }
        case "a":
            var v18 string = "a"
            _ = v18
        case "\n\t":
            for fi19 := 0; fi19 < 4; fi19++ {
                if len(v3) > 0 {
                    v3[0] = 42
                }
                switch fi19 {
                case -1:
                    _ = len(v3)
                }
                v2 = (v2 + 561)
                _ = fi19
            }
        }
        if (42 > -2147483648) {
            v4 -= v4
        }
    } else {
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 uint16 = (v2 * v2)
            _ = fi20
            _ = v21
        }
        switch "" {
        case "a":
            var v22 uint8 = 2
            _ = v22
        default:
            v3 = append(v3, 4294967295)
        }
    }
    if ("hello" < "\n\t") {
        v23 := GenS2{F0: 255}
        var v24 bool = v1
        _ = v23
        _ = v24
    }
    fmt.Println(v4)
    var v25 uint16 = 0
    v26 := GenS2{F0: 127}
    v4 = v4
    var v27 int16 = 32767
    v27 += 1
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v28 int = int(v4)
        _ = v28
    default:
        fmt.Println(v1, v4)
    }
    v29 := GenS3{F0: 1, F1: 65535}
    var v30 int = 602
    switch v30 {
    case 811:
        if (!v1) {
            v27 = (v27 >> (((v4 - 808)) & 63))
            fmt.Println(v1)
        }
    case 0:
        switch v30 {
        case 845:
            _ = len(v3)
        }
    default:
        v31 := GenS4{F0: (!(true || (v30 < (v30 * (^(v30)))))), F1: 1, F2: ("\n\t" > ("a" + "\n\t"))}
        _ = v31
    }
    v32 := GenS1{F0: 1, F1: (^(v2))}
    var v33 uint16 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
}
