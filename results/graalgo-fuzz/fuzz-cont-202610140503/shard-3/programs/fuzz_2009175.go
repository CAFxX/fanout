package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int16; F2 int }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 int16; F1 int16; F2 uint8 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint = 0
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{fi5, fi5}
        v2 = (v2 >> ((v2) & 63))
        _ = fi5
        _ = v6
    }
    v1 = v1
    var v7 uint32 = 2
    var v8 string = "hello"
    v8 = (v8 + "hello")
    for fi9 := 0; fi9 < 2; fi9++ {
        v7 += v7
        v7 = v7
        _ = fi9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        if (1 < -1) {
            v7 = 2147483647
            v11 := []int{(fi10 << ((v2) & 63)), fi10, fi10}
            _ = v11
        } else {
            if true {
                v1 = v1
                v1 = (-(v1))
            } else {
                for fi12 := 0; fi12 < 17; fi12++ {
                    switch v8 {
                    case "foo bar":
                        if (126 < 1) {
                            v13 := GenS1{F0: (false && (-275 < (v1 + (v1 | (^(v1))))))}
                            _ = v13
                        } else {
                            _ = len(v3)
                            v14 := GenS1{F0: (0 != v1)}
                            _ = v14
                        }
                    case "hello":
                        var v15 float32 = ((3.14159 / 1.0) + 0.1)
                        _ = v15
                    case "":
                        var v16 uint16 = 42
                        _ = v16
                    default:
                        v7 = (v7 / (((v7 ^ v7)) | 1))
                    }
                    var v17 uint16 = 1
                    if false {
                        if true {
                            var v18 int16 = int16(v1)
                            var v19 float64 = 1e10
                            _ = v18
                            _ = v19
                        }
                        switch 42 {
                        case 1:
                            var v20 bool = false
                            _ = v20
                        case 0:
                            var v21 uint8 = 69
                            _ = v21
                        default:
                            if (!(2 <= 42)) {
                                var v22 uint32 = 0
                                v23 := GenS2{F0: (v2 + (v2 * (v2 >> ((4294967295) & 63)))), F1: -1, F2: fi12}
                                _ = v22
                                _ = v23
                            }
                        }
                        var v24 int16 = -1
                        _ = v24
                    } else {
                        var v25 int16 = int16(v2)
                        switch (fi12 >> (((v2 + v2)) & 63)) {
                        case -2147483648:
                            v8 = v8
                        }
                        _ = v25
                    }
                    _ = fi12
                    _ = v17
                }
            }
            if (!false) {
                switch fi10 {
                case 127:
                    v1 += (^(v1))
                case 209:
                    for fi26 := 0; fi26 < 7; fi26++ {
                        if false {
                            v27 := v3["foo bar"]
                            var v28 int = fi26
                            v29 := []int{}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS2{F0: v2, F1: -32768, F2: (fi26 >> ((v2) & 63))}
                        v31 := GenS3{F0: 0, F1: -305, F2: "a"}
                        _ = fi26
                        _ = v30
                        _ = v31
                    }
                }
            } else {
                if ((true || ((v7 >> ((v2) & 63)) < v7)) && true) {
                    v1 -= v1
                    var v32 int32 = -1
                    _ = v32
                }
                if true {
                    var v33 int64 = 269
                    _ = v33
                }
            }
        }
        _ = fi10
    }
    v34 := []int{-348}
    switch "" {
    case "\n\t":
        var v35 float64 = (float64(1.5) - (float64(0.1) - ((1.5 - 1.5) * 1.5)))
        _ = v35
    default:
        switch "\n\t" {
        case "foo bar":
            v2 = (v2 + v2)
        case "hello":
            switch v8 {
            case "\n\t":
                var v36 bool = false
                _ = v36
            case "hello":
                v37 := v4["\n\t"]
                _ = v37
            default:
                v38 := GenS1{F0: true}
                _ = v38
            }
        }
    }
    v39 := v3["foo bar"]
    var v40 float64 = 3.14159
    v41 := GenS4{F0: -870, F1: 1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v7
    _ = v8
}
