package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 42
    var v3 float64 = (3.14159 - (float64(((3.14159 + (((0.1 * 1e10) - 1.5) - 3.14159)) + 3.14159)) * (1.5 / 0.1)))
    if (0 > 42) {
        for fi4 := 0; fi4 < 16; fi4++ {
            v5 := GenS1{F0: "hello", F1: 0, F2: 1}
            v6 := v1[(fi4 & fi4)]
            var v7 int8 = (v2 | v2)
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    }
    v8 := v1[1]
    switch v8 {
    case 42:
        delete(v1, (v8 - v8))
    }
    v9 := GenS1{F0: "hello", F1: 291, F2: 634}
    var v10 uint16 = 349
    switch ("hello" + "hello") {
    case "a":
        v8 = v8
    case "\n\t":
        var v11 int16 = -32768
        _ = v11
    }
    var v12 uint16 = v10
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v12)
        _ = fi13
    }
    v14 := GenS2{F0: v12, F1: 196, F2: v2}
    v8 -= (v8 | int(v2))
    v15 := GenS1{F0: "", F1: 255, F2: 42}
    v10 = (v10 | v12)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v3, v10, v8)
    case "":
        if true {
            var v16 string = (((("foo bar" + "a") + "hello") + "\n\t") + "foo bar")
            _ = v16
        } else {
            for fi17 := 0; fi17 < 15; fi17++ {
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 uint = 0
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        }
    default:
        if (-9223372036854775808 <= -611) {
            v20 := GenS1{F0: "", F1: 550, F2: -32768}
            var v21 float64 = float64(-2.25)
            _ = v20
            _ = v21
        } else {
            switch 994 {
            case 42:
                v22 := GenS3{F0: -344, F1: -642}
                _ = v22
            case 32767:
                _ = len(v1)
            default:
                if true {
                    v23 := v1[(-(v8))]
                    v24 := []int{(-(v23)), 127, v8}
                    _ = v23
                    _ = v24
                } else {
                    var v25 int8 = (v2 / ((v2) | 1))
                    for fi26 := 0; fi26 < 5; fi26++ {
                        if false {
                            var v27 int = (fi26 ^ (-(fi26)))
                            _ = v27
                        }
                        v28 := GenS3{F0: 868, F1: 1}
                        _ = fi26
                        _ = v28
                    }
                    switch v8 {
                    case 127:
                        for fi29 := 0; fi29 < 3; fi29++ {
                            var v30 int = 649
                            v31 := GenS4{F0: 848, F1: -2147483648}
                            switch "\n\t" {
                            case "a":
                                v10 += v12
                            case "\n\t":
                                v32 := GenS3{F0: -1, F1: 1}
                                _ = v32
                            default:
                                for fi33 := 0; fi33 < 10; fi33++ {
                                    v8 = (-(fi33))
                                    v34 := GenS5{F0: v10}
                                    fmt.Println(fi33, v2)
                                    v35 := []int{(^(v30)), 128, v8, v8}
                                    _ = fi33
                                    _ = v34
                                    _ = v35
                                }
                            }
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    default:
                        var v36 float32 = (1.5 + 1.5)
                        _ = v36
                    }
                    _ = v25
                }
            }
        }
    }
    var v37 int64 = -1
    for fi38 := 0; fi38 < 19; fi38++ {
        fmt.Println(v8, v12)
        _ = fi38
    }
    var v39 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v8
    _ = v9
}
