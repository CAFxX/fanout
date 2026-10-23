package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32; F2 int }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 0
    var v3 uint8 = 0
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    switch ("\n\t" + "hello") {
    case "\n\t":
        v7 := []int{-320, 4294967295, 4294967295, 2147483647}
        _ = v7
    case "foo bar":
        var v8 int64 = 9223372036854775807
        _ = v8
    case "hello":
        v9 := GenS1{F0: -976, F1: 2, F2: 42}
        _ = v9
    default:
        for fi10 := 0; fi10 < 5; fi10++ {
            if true {
                var v11 uint16 = 732
                if (((1.5 * 0.0) * 0.1) >= (3.14159 * 0.0)) {
                    switch fi10 {
                    case 32768:
                        v12 := v6[int(v11)]
                        _ = v12
                    default:
                        v3 += (v3 | 255)
                    }
                }
                v13 := GenS2{F0: 1000, F1: 323}
                _ = v11
                _ = v13
            } else {
                v14 := GenS1{F0: 921, F1: 445, F2: 18446744073709551615}
                for fi15 := 0; fi15 < 10; fi15++ {
                    switch -227 {
                    case 32767:
                        fmt.Println(fi10, fi10, fi15)
                    case -923:
                        var v16 uint = 18446744073709551615
                        _ = v16
                    case 127:
                        var v17 uint8 = 179
                        _ = v17
                    }
                    _ = fi15
                }
                v18 := []int{-299, (fi10 & (-(fi10))), fi10, -62}
                _ = v14
                _ = v18
            }
            _ = fi10
        }
    }
    switch -1 {
    case -810:
        var v19 bool = false
        _ = v19
    case 245:
        var v20 int = 32768
        _ = v20
    default:
        v2 += (v2 * 0)
    }
    v2 = -337
    var v21 uint16 = 2
    var v22 int64 = -685
    if false {
        _ = len(v5)
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 int8 = 1
            for fi25 := 0; fi25 < 18; fi25++ {
                switch (("a" + ("" + ("a" + "\n\t"))) + ("" + "foo bar")) {
                case "\n\t":
                    v21 = (^(v21))
                default:
                    fmt.Println(v22)
                }
                delete(v1, "a")
                if (!false) {
                    v2 = (v2 << ((65535) & 63))
                    v6[fi25] = "a"
                } else {
                    switch fi25 {
                    case 896:
                        v26 := []int{fi25}
                        _ = v26
                    case 9223372036854775807:
                        v27 := GenS1{F0: 0, F1: 0, F2: 2}
                        _ = v27
                    case 0:
                        var v28 uint = 18446744073709551615
                        _ = v28
                    }
                    var v29 uint64 = 18446744073709551615
                    _ = v29
                }
                delete(v1, (("hello" + ("foo bar" + "")) + (("hello" + "foo bar") + "hello")))
                v3 -= v3
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
    } else {
        for fi30 := 0; fi30 < 4; fi30++ {
            var v31 uint = 65535
            _ = fi30
            _ = v31
        }
        v3 = (v3 & 8)
        fmt.Println(v21, v2)
    }
    if (false || true) {
        for fi32 := 0; fi32 < 4; fi32++ {
            if false {
                v33 := GenS2{F0: -32768, F1: 2}
                _ = v33
            }
            v3 += 71
            for fi34 := 0; fi34 < 11; fi34++ {
                switch "foo bar" {
                case "":
                    var v35 uint32 = 0
                    _ = v35
                default:
                    v36 := GenS3{F0: 1, F1: true, F2: 768}
                    _ = v36
                }
                v22 = -9223372036854775808
                _ = fi34
            }
            _ = fi32
        }
        var v37 float64 = (3.14159 + 1e10)
        v38 := []int{127, 32768, -9223372036854775808}
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        v40 := GenS4{F0: v22, F1: 1, F2: (fi39 >> ((4294967295) & 63))}
        switch (fi39 / ((fi39) | 1)) {
        case -9223372036854775808:
            v21 += v21
        case 128:
            v41 := GenS5{F0: 84, F1: 1000}
            _ = v41
        default:
            v42 := []int{32768, (^(fi39))}
            _ = v42
        }
        for fi43 := 0; fi43 < 9; fi43++ {
            v44 := GenS1{F0: 419, F1: 338, F2: 764}
            v22 = (v22 - v22)
            _ = fi43
            _ = v44
        }
        _ = fi39
        _ = v40
    }
    v45 := []int{-1, 2147483647, 656}
    v46 := GenS6{F0: -9223372036854775808, F1: 42}
    if false {
        var v47 uint32 = 42
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
