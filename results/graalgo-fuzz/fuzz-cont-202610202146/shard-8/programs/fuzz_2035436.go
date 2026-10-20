package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 bool }

func main() {
    var v1 uint64 = 568
    v2 := make(map[int]uint32)
    var v3 uint32 = 0
    v4 := []int{127, 9223372036854775807}
    var v5 uint64 = 42
    var v6 uint16 = 2
    var v7 float32 = (float32((-2.25 / 0.1)) / 1.5)
    var v8 string = "foo bar"
    for fi9 := 0; fi9 < 18; fi9++ {
        if true {
            v10 := GenS1{F0: (v3 << ((2) & 63))}
            _ = v10
        } else {
            if false {
                v8 = v8
                var v11 int32 = 2147483647
                _ = len(v2)
                _ = v11
            } else {
                var v12 int = (fi9 << ((0) & 63))
                _ = v12
            }
            var v13 int16 = 42
            _ = v13
        }
        _ = fi9
    }
    var v14 int64 = 1
    v15 := GenS2{F0: ((v5 ^ 0) <= 4294967295)}
    fmt.Println(v14, v7)
    v8 = v8
    if (true && (!false)) {
        if len(v4) > 0 {
            v4[0] = -1
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(v1, v8)
            fmt.Println(v1, v1)
            _ = fi16
        }
    }
    var v17 uint32 = (v3 | (v3 >> ((4294967295) & 63)))
    v6 = v6
    v3 = v17
    var v18 uint32 = (v3 - (v3 * 2147483648))
    v3 -= (v17 * uint32(v6))
    var v19 int32 = -1
    if (128 == 9223372036854775807) {
        var v20 int = -1
        _ = v20
    }
    switch 1 {
    case 1:
        v17 = 42
    case 42:
        var v21 uint16 = v6
        _ = v21
    case 319:
        if (v8 > v8) {
            fmt.Println(v14, v6, v1)
            v19 = (v19 | (^(v19)))
            var v22 uint32 = v18
            _ = v22
        } else {
            v19 -= 42
            switch 42 {
            case 32767:
                switch (v8 + "a") {
                case "":
                    switch v8 {
                    case "foo bar":
                        switch ("" + "a") {
                        case "a":
                            var v23 int64 = -1
                            _ = v23
                        case "foo bar":
                            for fi24 := 0; fi24 < 17; fi24++ {
                                v7 = 0.1
                                switch (fi24 | (fi24 >> ((0) & 63))) {
                                case 127:
                                    switch ("a" + v8) {
                                    case "":
                                        for fi25 := 0; fi25 < 4; fi25++ {
                                            switch fi25 {
                                            case 128:
                                                if true {
                                                    if true {
                                                        v7 = ((v7 + (((float32(3.14159) * 0.0) * (0.1 + ((float32(3.14159) / 1.0) - float32((((-2.25 * 1e10) * 0.0) * -2.25))))) * (0.0 - v7))) + (v7 + v7))
                                                        delete(v2, fi24)
                                                        var v26 uint32 = v18
                                                        _ = v26
                                                    }
                                                } else {
                                                    switch fi24 {
                                                    case 4294967295:
                                                        v3 -= (v18 >> ((255) & 63))
                                                        var v27 bool = ((1e10 / 1e10) >= 0.1)
                                                        fmt.Println(v17)
                                                        _ = v27
                                                    case 0:
                                                        fmt.Println(v19, v8)
                                                    default:
                                                        fmt.Println(v8, v19)
                                                    }
                                                }
                                            case -2147483648:
                                                if (true || false) {
                                                    var v28 uint = 353
                                                    _ = v28
                                                } else {
                                                    v19 = v19
                                                    v29 := GenS3{F0: "\n\t", F1: (100 >= -54)}
                                                    _ = v29
                                                }
                                            }
                                            _ = fi25
                                        }
                                    case "hello":
                                        v1 = v5
                                    case "\n\t":
                                        v6 += (v6 + (v6 * (v6 * v6)))
                                    }
                                case 32767:
                                    v30 := GenS1{F0: (v18 ^ (v17 / ((v18) | 1)))}
                                    _ = v30
                                case 4294967295:
                                    var v31 int32 = v19
                                    _ = v31
                                }
                                _ = fi24
                            }
                        case "hello":
                            v32 := GenS1{F0: (v3 % (((v18 << ((0) & 63))) | 1))}
                            _ = v32
                        }
                    case "hello":
                        for fi33 := 0; fi33 < 9; fi33++ {
                            v19 -= 42
                            var v34 uint16 = (^(v6))
                            _ = fi33
                            _ = v34
                        }
                    }
                }
            case 32768:
                if false {
                    v4 = append(v4, -583)
                    v3 = uint32(v5)
                } else {
                    var v35 uint = 516
                    v19 = v19
                    _ = v35
                }
            }
            v5 -= v5
        }
    default:
        fmt.Println(v14, v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
