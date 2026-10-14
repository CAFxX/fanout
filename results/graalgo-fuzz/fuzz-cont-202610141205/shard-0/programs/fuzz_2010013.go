package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{-871, 128, 32768}
    var v3 uint16 = 840
    var v4 float32 = -2.25
    v5 := make(map[uint32]bool)
    fmt.Println(v3)
    var v6 string = "foo bar"
    v6 = ("\n\t" + "foo bar")
    if true {
        v7 := GenS1{F0: 271}
        _ = v7
    }
    if false {
        fmt.Println(v4, v6, v3)
        v4 = 0.0
        v8 := v5[2]
        _ = v8
    }
    v6 = ""
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: -1}
        var v11 uint16 = v3
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v4 += (v4 * (float32(-2.25) / (1e10 * float32(0.1))))
    for fi12 := 0; fi12 < 7; fi12++ {
        if (!(false && true)) {
            if true {
                v13 := GenS1{F0: -963}
                _ = v13
            }
            var v14 float32 = float32(((3.14159 * (0.1 + 3.14159)) * ((float64((v4 / (v4 - float32(1e10)))) * float64(1.5)) + (-2.25 + (0.1 * (3.14159 / ((1e10 / 1.0) * 1.5)))))))
            v15 := GenS1{F0: -754}
            _ = v14
            _ = v15
        } else {
            switch "foo bar" {
            case "a":
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 uint64 = 1
                    v3 += 0
                    var v18 uint32 = 42
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v19 := GenS2{F0: 42, F1: 512, F2: 859}
            var v20 int = -819
            _ = v19
            _ = v20
        }
        _ = fi12
    }
    if (!(true || true)) {
        v1 -= (^(v1))
        v21 := GenS2{F0: 405, F1: 948, F2: 2147483648}
        _ = v21
    } else {
        var v22 int16 = 1000
        switch ("\n\t" + (v6 + v6)) {
        case "foo bar":
            if false {
                if true {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v2 = append(v2, fi23)
                        _ = fi23
                    }
                }
                if true {
                    if true {
                        v1 += v1
                    }
                    v1 -= (v1 + 75)
                    _ = len(v2)
                } else {
                    v1 = v1
                    for fi24 := 0; fi24 < 17; fi24++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi24
                    }
                }
            }
        case "a":
            var v25 uint64 = 364
            _ = v25
        default:
            v6 = (("\n\t" + "") + ("foo bar" + (v6 + "hello")))
        }
        _ = v22
    }
    var v26 uint16 = 32767
    var v27 uint8 = 1
    var v28 uint32 = 42
    fmt.Println(v28)
    switch v6 {
    case "a":
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 string = "hello"
            v26 += v26
            _ = fi29
            _ = v30
        }
    case "\n\t":
        v4 += 3.14159
    case "hello":
        var v31 int32 = 42
        _ = v31
    default:
        if false {
            fmt.Println(v6, v1, v26)
            delete(v5, 2147483647)
        } else {
            v32 := GenS1{F0: 1}
            _ = v32
        }
    }
    v33 := GenS1{F0: -2147483648}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
