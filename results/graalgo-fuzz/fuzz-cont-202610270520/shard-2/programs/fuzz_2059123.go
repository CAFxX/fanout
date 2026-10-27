package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int64 = 0
    v2 := make(map[string]bool)
    var v3 int16 = -834
    fmt.Println(v3, v3)
    if true {
        delete(v2, "a")
        for fi4 := 0; fi4 < 9; fi4++ {
            v5 := GenS1{F0: false, F1: false, F2: 318}
            _ = fi4
            _ = v5
        }
        var v6 float64 = 1.5
        _ = v6
    }
    fmt.Println(v3, v3)
    switch -9223372036854775808 {
    case 32768:
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 float64 = 0.1
            var v9 bool = false
            if ((fi7 < (fi7 << ((738) & 63))) || true) {
                v1 = (v1 % ((v1) | 1))
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 float32 = 3.14159
                    var v12 uint32 = 2147483648
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 uint32 = 871
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    case 32767:
        v1 += v1
    case 2147483647:
        var v15 int64 = v1
        _ = v15
    }
    v16 := []int{}
    v3 = v3
    var v17 uint32 = 65535
    switch ((("a" + "a") + "hello") + "") {
    case "hello":
        if true {
            if (!true) {
                v17 -= 937
                v18 := GenS2{F0: (v1 * v1)}
                switch 1 {
                case 2147483647:
                    v19 := GenS3{F0: (v3 / (((-(v3))) | 1)), F1: 443}
                    _ = v19
                case -1:
                    var v20 uint = 1
                    _ = v20
                default:
                    var v21 float32 = 0.0
                    _ = v21
                }
                _ = v18
            } else {
                v22 := v2[("foo bar" + "a")]
                if (!v22) {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        _ = len(v2)
                        v24 := GenS4{F0: v17, F1: 128}
                        v25 := GenS1{F0: v22, F1: false, F2: 1}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    var v26 bool = v22
                    _ = v26
                } else {
                    var v27 int32 = 2147483647
                    switch "hello" {
                    case "hello":
                        fmt.Println(v22, v27)
                    case "\n\t":
                        for fi28 := 0; fi28 < 13; fi28++ {
                            switch "a" {
                            case "\n\t":
                                v16 = append(v16, (fi28 % ((fi28) | 1)))
                            case "a":
                                if (!(v22 || true)) {
                                    for fi29 := 0; fi29 < 7; fi29++ {
                                        v1 -= (v1 >> ((319) & 63))
                                        var v30 int64 = (v1 / (((v1 / ((v1) | 1))) | 1))
                                        _ = fi29
                                        _ = v30
                                    }
                                    var v31 uint = 255
                                    _ = v31
                                }
                            }
                            _ = fi28
                        }
                    case "foo bar":
                        var v32 int = 2147483647
                        _ = v32
                    }
                    _ = v27
                }
                if true {
                    v17 -= 4294967295
                    var v33 uint64 = 2
                    v1 = v1
                    _ = v33
                }
                _ = v22
            }
            if (!true) {
                v34 := GenS3{F0: 1, F1: 644}
                _ = v34
            }
            var v35 uint = 33
            _ = v35
        } else {
            switch -2147483648 {
            case 128:
                var v36 int32 = 0
                _ = v36
            case 127:
                fmt.Println(v17, v1)
            case -930:
                v16 = append(v16, -2147483648)
            default:
                var v37 uint16 = 1
                _ = v37
            }
            fmt.Println(v3)
        }
    case "a":
        switch "foo bar" {
        case "hello":
            for fi38 := 0; fi38 < 7; fi38++ {
                v39 := GenS2{F0: (v1 * v1)}
                _ = fi38
                _ = v39
            }
        default:
            v1 += (v1 >> ((955) & 63))
        }
    default:
        v40 := GenS1{F0: ((^(v3)) >= (v3 % ((v3) | 1))), F1: (65535 != 454), F2: 2}
        _ = v40
    }
    var v41 float32 = 1e10
    var v42 string = ("a" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v41
    _ = v42
}
