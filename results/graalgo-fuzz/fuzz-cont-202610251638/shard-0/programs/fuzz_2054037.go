package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint64; F1 string }

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]bool)
    var v3 string = v1
    var v4 int64 = -9223372036854775808
    var v5 int = 4294967295
    var v6 bool = (127 == 1)
    _ = len(v2)
    var v7 float32 = -2.25
    v8 := []int{v5, v5}
    switch v1 {
    case "hello":
        switch ("hello" + ("a" + (("a" + "hello") + "a"))) {
        case "\n\t":
            v9 := GenS1{F0: (v5 << ((18446744073709551615) & 63))}
            _ = v9
        case "foo bar":
            v10 := GenS1{F0: v5}
            _ = v10
        default:
            for fi11 := 0; fi11 < 3; fi11++ {
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int16 = 752
                    if (!(!false)) {
                        var v14 uint8 = 255
                        if v6 {
                            var v15 int32 = 249
                            _ = v15
                        } else {
                            v3 = ("a" + ("foo bar" + ("foo bar" + v1)))
                            if len(v8) > 0 {
                                v8[0] = (fi12 >> ((255) & 63))
                            }
                            var v16 int = 1
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = fi12
                    _ = v13
                }
                var v17 int32 = 2147483647
                v18 := GenS2{F0: 42, F1: -120}
                _ = fi11
                _ = v17
                _ = v18
            }
        }
    case "\n\t":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        v4 = v4
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if true {
            if v6 {
                if (!v6) {
                    var v20 string = v3
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 17; fi21++ {
                var v22 int32 = 0
                v7 += (v7 / (v7 + v7))
                _ = fi21
                _ = v22
            }
        }
        fmt.Println(v6)
        _ = fi19
    }
    var v23 float32 = -2.25
    switch v5 {
    case -176:
        for fi24 := 0; fi24 < 9; fi24++ {
            v3 = ""
            v25 := GenS1{F0: (fi24 | v5)}
            _ = fi24
            _ = v25
        }
    case 1:
        if len(v8) > 0 {
            v8[0] = v5
        }
    case 442:
        var v26 int64 = (v4 | v4)
        _ = v26
    }
    var v27 uint16 = 255
    v28 := v2[747]
    if len(v8) > 0 {
        v8[0] = (v5 >> ((670) & 63))
    }
    var v29 uint = 428
    if false {
        v30 := GenS3{F0: 0, F1: v28}
        var v31 string = v1
        _ = v30
        _ = v31
    } else {
        v28 = (int8(v5) < 100)
        var v32 int = (v5 * 32767)
        _ = v32
    }
    v2[2] = (!(v28 || (true || (1 < -32768))))
    switch v5 {
    case 1:
        v5 = 0
    case -260:
        v33 := GenS4{F0: 4294967295, F1: ((v1 + v1) + v3)}
        _ = v33
    case 9223372036854775807:
        _ = len(v2)
    default:
        v4 += (v4 ^ v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
