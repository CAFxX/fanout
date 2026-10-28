package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]bool)
    var v3 float32 = -2.25
    v4 := GenS1{F0: 18446744073709551615, F1: (528 < (v1 / (((^(v1))) | 1)))}
    v1 -= v1
    v5 := GenS1{F0: 4294967295, F1: true}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint = 18446744073709551615
        switch "foo bar" {
        case "foo bar":
            v3 = ((v3 / v3) - (1.5 / v3))
        case "hello":
            v8 := GenS1{F0: v7, F1: (false && (1 != 65535))}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    v9 := []int{32767}
    fmt.Println(v3)
    v1 = 948
    if len(v9) > 0 {
        v9[0] = 32768
    }
    if true {
        v1 = v1
    } else {
        v10 := GenS2{F0: 0, F1: 0}
        if (3.14159 <= 3.14159) {
            for fi11 := 0; fi11 < 7; fi11++ {
                switch "a" {
                case "a":
                    var v12 float64 = -2.25
                    _ = v12
                default:
                    fmt.Println(v1)
                }
                _ = fi11
            }
            var v13 int = 32768
            _ = v13
        }
        _ = v10
    }
    if ((0.0 == 1.5) || (255 >= 0)) {
        for fi14 := 0; fi14 < 10; fi14++ {
            delete(v2, 3)
            _ = fi14
        }
        fmt.Println(v3, v1)
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            switch fi15 {
            case 42:
                switch "\n\t" {
                case "":
                    v17 := GenS2{F0: 127, F1: fi16}
                    _ = v17
                case "foo bar":
                    v18 := GenS1{F0: 1, F1: false}
                    _ = v18
                default:
                    var v19 int64 = -34
                    _ = v19
                }
            }
            _ = fi16
        }
        var v20 float64 = 0.0
        _ = fi15
        _ = v20
    }
    v3 = v3
    switch 276 {
    case 32767:
        var v21 int16 = -32768
        _ = v21
    case -45:
        v22 := GenS1{F0: 984, F1: false}
        _ = v22
    case -2147483648:
        switch "a" {
        case "\n\t":
            for fi23 := 0; fi23 < 2; fi23++ {
                fmt.Println(fi23, fi23)
                v1 = (^(v1))
                if true {
                    if false {
                        v1 = 2
                    } else {
                        if false {
                            var v24 uint64 = (v1 - (v1 | 42))
                            var v25 float64 = 1e10
                            switch 4294967295 {
                            case 202:
                                fmt.Println(v25, v1)
                            case 811:
                                v26 := GenS1{F0: 877, F1: false}
                                _ = v26
                            default:
                                fmt.Println(fi23, v25)
                            }
                            _ = v24
                            _ = v25
                        } else {
                            if true {
                                var v27 uint64 = (v1 / ((v1) | 1))
                                _ = v27
                            }
                        }
                        for fi28 := 0; fi28 < 11; fi28++ {
                            v29 := v2[0]
                            var v30 uint64 = v1
                            var v31 string = "hello"
                            v30 -= (v30 % (((v1 | 800)) | 1))
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    }
                    if (("\n\t" + ("foo bar" + "foo bar")) > "hello") {
                        if false {
                            var v32 bool = false
                            v33 := GenS2{F0: 85, F1: 32768}
                            _ = v32
                            _ = v33
                        }
                    }
                    var v34 int64 = -721
                    _ = v34
                }
                _ = fi23
            }
        default:
            v35 := v2[2147483648]
            _ = v35
        }
    }
    v36 := GenS2{F0: 0, F1: 4294967295}
    v1 -= 181
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v9
}
