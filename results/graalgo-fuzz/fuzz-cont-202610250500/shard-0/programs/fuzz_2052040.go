package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
type GenS4 struct { F0 int; F1 int64; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint16; F2 bool }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 int64 }

func main() {
    var v1 int32 = -1
    v2 := []int{742}
    var v3 float64 = (1.5 + (3.14159 / (1e10 / (1.5 - (3.14159 - 1e10)))))
    v4 := make(map[int]bool)
    switch ("hello" + ("a" + "\n\t")) {
    case "\n\t":
        var v5 int64 = 166
        _ = v5
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        if len(v2) > 0 {
            v2[0] = 127
        }
    default:
        for fi6 := 0; fi6 < 8; fi6++ {
            if false {
                if (-1 != 0) {
                    switch (fi6 >> ((255) & 63)) {
                    case 128:
                        v3 -= (v3 + (((v3 * v3) + v3) - 0.1))
                    case -2147483648:
                        if false {
                            fmt.Println(fi6, fi6)
                            v7 := GenS1{F0: 1, F1: 4294967295, F2: 32767}
                            fmt.Println(v1)
                            _ = v7
                        }
                    }
                } else {
                    v8 := GenS2{F0: 273}
                    var v9 float32 = 1.5
                    _ = v8
                    _ = v9
                }
            } else {
                fmt.Println(fi6, v1)
                fmt.Println(v1)
            }
            _ = fi6
        }
    }
    var v10 uint64 = 0
    if false {
        switch 4294967295 {
        case -1:
            switch "foo bar" {
            case "foo bar":
                if false {
                    var v11 int64 = -144
                    fmt.Println(v11, v3, v1)
                    _ = v11
                }
            case "\n\t":
                fmt.Println(v10)
            case "":
                v3 += (float64(1.5) * 0.1)
            }
        default:
            switch -198 {
            case -1:
                var v12 float32 = ((-2.25 / 1e10) - (1.5 - (0.0 / float32(((float64(0.0) - float64(float32(v3))) - 0.1)))))
                _ = v12
            case -216:
                var v13 uint8 = 42
                _ = v13
            default:
                if (!true) {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v15 := v4[(fi14 ^ 74)]
                        switch "foo bar" {
                        case "\n\t":
                            v16 := GenS3{F0: 0, F1: 255, F2: 42}
                            _ = v16
                        case "":
                            switch "" {
                            case "a":
                                v17 := GenS3{F0: 0, F1: 42, F2: -710}
                                _ = v17
                            case "hello":
                                for fi18 := 0; fi18 < 17; fi18++ {
                                    v19 := GenS4{F0: fi18, F1: 528, F2: 42}
                                    _ = fi18
                                    _ = v19
                                }
                            default:
                                var v20 string = ("a" + ((("hello" + "a") + ("a" + "")) + ""))
                                _ = v20
                            }
                        }
                        _ = fi14
                        _ = v15
                    }
                } else {
                    var v21 uint16 = 255
                    _ = len(v4)
                    switch 32768 {
                    case -723:
                        var v22 string = "\n\t"
                        _ = v22
                    }
                    _ = v21
                }
            }
        }
        _ = len(v4)
    }
    var v23 string = "a"
    var v24 int8 = 127
    if false {
        v25 := GenS5{F0: v24}
        _ = v25
    }
    var v26 float32 = 0.0
    v24 = 127
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 int64 = 0
        if len(v2) > 0 {
            v2[0] = fi27
        }
        switch v23 {
        case "a":
            fmt.Println(v24)
        case "foo bar":
            var v29 uint16 = uint16(v10)
            _ = v29
        default:
            if ((float32((v3 + v3)) / v26) <= 1e10) {
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 float32 = ((v26 - float32(v3)) / 0.1)
                    v32 := GenS6{F0: -1, F1: 290, F2: true}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v33 := GenS7{F0: 2, F1: -1}
                _ = v33
            } else {
                v34 := GenS7{F0: 25, F1: (v28 * 1)}
                v3 = (((((v3 + v3) - float64(float32((v3 / 1e10)))) - v3) * v3) - 0.0)
                _ = v34
            }
        }
        _ = fi27
        _ = v28
    }
    v35 := GenS8{F0: -41, F1: 2}
    var v36 int16 = -32768
    v37 := GenS9{F0: 1}
    if false {
        for fi38 := 0; fi38 < 18; fi38++ {
            fmt.Println(v23)
            var v39 uint32 = uint32(v10)
            var v40 bool = (!(!true))
            _ = fi38
            _ = v39
            _ = v40
        }
        if true {
            fmt.Println(v1, v3)
        } else {
            if false {
                v10 += (v10 | (^(v10)))
                v23 = v23
                for fi41 := 0; fi41 < 16; fi41++ {
                    v1 = -1
                    v4[fi41] = false
                    _ = fi41
                }
            }
        }
        var v42 uint16 = 1
        _ = v42
    } else {
        v2 = append(v2, 32768)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            fmt.Println(v36)
        }
    }
    var v43 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v43
}
