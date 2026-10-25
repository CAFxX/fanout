package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }
type GenS2 struct { F0 int8; F1 uint8; F2 int }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]string)
    var v3 float32 = 0.1
    var v4 int64 = 1
    var v5 int8 = 1
    v6 := GenS1{F0: 40, F1: (v4 / (((v4 ^ v4)) | 1)), F2: -2147483648}
    var v7 int64 = (v4 * v4)
    v4 += v4
    var v8 uint32 = 0
    switch -275 {
    case 980:
        for fi9 := 0; fi9 < 19; fi9++ {
            v8 = (v8 * v8)
            switch "foo bar" {
            case "":
                if true {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        delete(v2, "")
                        var v11 float64 = 0.0
                        var v12 uint16 = 2
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    if false {
                        var v13 int16 = -32768
                        v4 += (v4 % ((v7) | 1))
                        v14 := GenS1{F0: 1, F1: (^(v4)), F2: (v1 >> ((65535) & 63))}
                        _ = v13
                        _ = v14
                    }
                } else {
                    v15 := GenS1{F0: (v8 ^ v8), F1: v7, F2: 52}
                    _ = v15
                }
            case "\n\t":
                var v16 int64 = v4
                _ = v16
            case "hello":
                v17 := []int{1, int(v4), (fi9 + 436)}
                _ = v17
            default:
                var v18 float32 = (float32(-2.25) - -2.25)
                _ = v18
            }
            _ = fi9
        }
    case 2147483647:
        var v19 float32 = v3
        _ = v19
    default:
        fmt.Println(v4, v5, v7)
    }
    v1 -= (v1 % ((-1) | 1))
    v20 := GenS1{F0: v8, F1: v4, F2: (v1 % ((861) | 1))}
    v21 := GenS2{F0: -79, F1: 42, F2: 2147483647}
    var v22 int = 32767
    var v23 uint16 = 2
    fmt.Println(v5)
    fmt.Println(v22, v1)
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS2{F0: (-(v5)), F1: uint8(v8), F2: (v22 | (-(fi24)))}
            _ = len(v2)
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            switch (("hello" + "foo bar") + (((("\n\t" + "\n\t") + ("a" + "hello")) + (("foo bar" + "hello") + "")) + ("\n\t" + ("hello" + "")))) {
            case "foo bar":
                switch "\n\t" {
                case "a":
                    var v28 float64 = -2.25
                    v4 += v7
                    fmt.Println(v8, v28)
                    _ = v28
                case "\n\t":
                    switch fi27 {
                    case -149:
                        v29 := GenS3{F0: 18446744073709551615, F1: v5}
                        _ = v29
                    case 515:
                        v7 -= 863
                    case 0:
                        v5 += -1
                    default:
                        if true {
                            delete(v2, "a")
                            v30 := GenS4{F0: 4294967295, F1: true, F2: ""}
                            _ = v30
                        }
                    }
                case "":
                    var v31 int32 = (v1 + (-(v1)))
                    _ = v31
                }
            case "a":
                v1 = v1
            case "":
                if false {
                    v32 := GenS3{F0: 2, F1: v5}
                    v33 := v2[""]
                    _ = v32
                    _ = v33
                }
            }
            _ = fi27
        }
        var v34 uint32 = v8
        _ = fi24
        _ = v34
    }
    var v35 uint16 = 1
    var v36 int8 = v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
