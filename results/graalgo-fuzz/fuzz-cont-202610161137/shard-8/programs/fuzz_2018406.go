package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int8; F2 string }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 int16, p1 string) uint {
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            v8 := []int{fi7, -9223372036854775808}
            p1 = "hello"
            _ = v8
        }
        _ = fi7
    }
    if true {
        if (127 < 230) {
            v9 := []int{-9223372036854775808, 128}
            _ = v9
        } else {
            v10 := GenS1{F0: 1}
            var v11 bool = (!true)
            v11 = (v11 || v11)
            _ = v10
            _ = v11
        }
        p0 -= 431
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS1{F0: 907}
            _ = fi12
            _ = v13
        }
    }
    switch (p1 + (p1 + "foo bar")) {
    case "a":
        fmt.Println(p1, p1)
    }
    v14 := GenS2{F0: (p0 & 381)}
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(fi15, p1)
        for fi16 := 0; fi16 < 5; fi16++ {
            p0 += (^(p0))
            v17 := []int{(fi15 & fi16), (fi15 % ((377) | 1))}
            switch p1 {
            case "a":
                if false {
                    v18 := GenS3{F0: fi15, F1: -87, F2: p1}
                    fmt.Println(p0, fi16)
                    if false {
                        var v19 int64 = 224
                        _ = v19
                    }
                    _ = v18
                } else {
                    var v20 int = (fi16 ^ -1)
                    var v21 int = -9223372036854775808
                    if true {
                        switch (fi15 + v21) {
                        case 4294967295:
                            v20 = fi15
                        case 42:
                            if (2 != 42) {
                                var v22 int16 = -32768
                                switch p1 {
                                case "":
                                    var v23 int32 = -103
                                    _ = v23
                                case "\n\t":
                                    var v24 uint16 = 2
                                    _ = v24
                                case "hello":
                                    var v25 float32 = 3.14159
                                    _ = v25
                                default:
                                    if false {
                                        var v26 int8 = -1
                                        var v27 float32 = 1.5
                                        p1 = (p1 + "foo bar")
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                var v28 int64 = 593
                                _ = v22
                                _ = v28
                            }
                        default:
                            var v29 float64 = ((float64(float32((1.5 + 1.5))) / -2.25) / 1.0)
                            _ = v29
                        }
                    }
                    _ = v20
                    _ = v21
                }
            case "hello":
                var v30 float32 = (-2.25 + float32(0.1))
                _ = v30
            case "\n\t":
                v31 := GenS4{F0: (fi15 + fi15), F1: 32767}
                _ = v31
            }
            _ = fi16
            _ = v17
        }
        p1 = ("foo bar" + p1)
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = v14
    return 1
}

func main() {
    var v1 uint16 = 321
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := []string{"hello", "a"}
    v5 := make(map[int]string)
    var v6 uint32 = 2
    fmt.Println(v6)
    var v32 float64 = 0.1
    v32 += v32
    var v33 int32 = 0
    if (true || (-1 > int8(v1))) {
        var v34 uint16 = 42
        v2["hello"] = 561
        var v35 int32 = v33
        _ = v34
        _ = v35
    } else {
        v6 -= v6
        var v36 float64 = 0.0
        fmt.Println(v36, v36, v6)
        _ = v36
    }
    v32 = 0.0
    fmt.Println(v32, v32)
    v32 = v32
    var v37 uint16 = (^(v1))
    var v38 int16 = 1000
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
