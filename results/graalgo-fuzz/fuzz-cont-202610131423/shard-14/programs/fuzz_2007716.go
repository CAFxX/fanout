package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int16 }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int32, p1 int8) int16 {
    var v7 bool = (128 == 96)
    var v8 int16 = -32768
    p0 = (-(p0))
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int8 = 100
        v10 += p1
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "hello"
    v3 := make(map[uint32]string)
    var v4 bool = false
    var v5 bool = v4
    v6 := make(map[string]bool)
    if true {
        var v11 string = "\n\t"
        _ = v11
    } else {
        var v12 int = -9223372036854775808
        var v13 float64 = float64((-2.25 + -2.25))
        v2 = (v2 + "\n\t")
        _ = v12
        _ = v13
    }
    v14 := []int{-9223372036854775808}
    v2 = ((v2 + ((v2 + ("\n\t" + (v2 + v2))) + (v2 + "a"))) + "foo bar")
    switch 393 {
    case 128:
        for fi15 := 0; fi15 < 9; fi15++ {
            if false {
                fmt.Println(fi15, v5, v4)
                var v16 bool = v5
                v2 = ("foo bar" + "hello")
                _ = v16
            }
            for fi17 := 0; fi17 < 4; fi17++ {
                switch 291 {
                case -9223372036854775808:
                    v18 := GenS1{F0: "hello", F1: (-2.25 > (float32((1e10 / 1e10)) / (float32(float64(1.5)) + 1e10))), F2: 0}
                    _ = v18
                case 730:
                    v1[62] = "hello"
                case 457:
                    var v19 float64 = -2.25
                    switch (-(fi17)) {
                    case 4294967295:
                        v20 := GenS2{F0: 32767, F1: 2, F2: 42}
                        _ = v20
                    case 1:
                        v21 := GenS1{F0: (v2 + ("foo bar" + v2)), F1: false, F2: 1}
                        _ = v21
                    default:
                        if true {
                            v2 = "\n\t"
                            v4 = v5
                            v22 := GenS1{F0: v2, F1: v5, F2: 508}
                            _ = v22
                        }
                    }
                    if (0 < 100) {
                        switch (v2 + ((v2 + (v2 + v2)) + "\n\t")) {
                        case "hello":
                            var v23 uint64 = 42
                            _ = v23
                        case "":
                            v2 = ""
                        }
                    }
                    v5 = false
                    _ = v19
                default:
                    var v24 float32 = 3.14159
                    _ = v24
                }
                v5 = v5
                _ = len(v14)
                _ = fi17
            }
            v2 = v2
            _ = fi15
        }
    case 9223372036854775807:
        v4 = (!false)
    case 2147483647:
        v25 := GenS1{F0: ("\n\t" + "\n\t"), F1: false, F2: genFn0(760, 0)}
        _ = v25
    }
    fmt.Println(v2, v4)
    v4 = (!v4)
    switch 32768 {
    case 14:
        v26 := GenS2{F0: -32768, F1: 4294967295, F2: 42}
        _ = v26
    case 453:
        if v4 {
            fmt.Println(v2, v4, v4)
            v2 = v2
            var v27 float64 = -2.25
            _ = v27
        }
    }
    v28 := GenS3{F0: 8, F1: 127}
    fmt.Println(v4, v2, v4)
    v5 = v5
    var v29 uint8 = 1
    var v30 string = v2
    var v31 float32 = -2.25
    v14 = append(v14, 600)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
