package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint16; F2 int }
type GenS5 struct { F0 bool; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 uint64 = 0
    v2 := []uint32{}
    v3 := make(map[uint32]uint32)
    v1 = 4294967295
    if (true || true) {
        fmt.Println(v1)
        for fi4 := 0; fi4 < 5; fi4++ {
            var v5 float32 = (1.5 - 0.0)
            var v6 uint64 = (v1 ^ 18446744073709551615)
            var v7 int = -267
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    }
    v3[0] = 42
    switch 0 {
    case -2147483648:
        fmt.Println(v1)
    case -1:
        switch -419 {
        case 9223372036854775807:
            switch -329 {
            case 32767:
                var v8 int32 = 2147483647
                _ = v8
            case 0:
                v9 := v3[2147483648]
                _ = v9
            }
        case -1:
            v10 := GenS1{F0: 32767}
            _ = v10
        default:
            v2 = append(v2, 42)
        }
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v1 -= v1
        v12 := GenS1{F0: 128}
        _ = fi11
        _ = v12
    }
    switch 128 {
    case -678:
        v1 = (v1 | (v1 >> ((0) & 63)))
    default:
        fmt.Println(v1)
    }
    if false {
        v13 := GenS1{F0: 2147483647}
        _ = v13
    } else {
        if false {
            switch 127 {
            case -124:
                v1 -= (v1 / (((^(v1))) | 1))
            case -519:
                var v14 float32 = 3.14159
                _ = v14
            }
            v15 := GenS1{F0: 32768}
            for fi16 := 0; fi16 < 19; fi16++ {
                for fi17 := 0; fi17 < 3; fi17++ {
                    if true {
                        switch 346 {
                        case -696:
                            v1 = (v1 & (v1 * (v1 + (^(v1)))))
                        case 4294967295:
                            v18 := GenS2{F0: 42}
                            _ = v18
                        case 127:
                            var v19 int32 = 358
                            _ = v19
                        default:
                            var v20 uint16 = 2
                            _ = v20
                        }
                    } else {
                        var v21 uint32 = 0
                        _ = v21
                    }
                    v22 := v3[uint32(v1)]
                    var v23 string = "hello"
                    _ = fi17
                    _ = v22
                    _ = v23
                }
                v24 := GenS2{F0: -9223372036854775808}
                var v25 int64 = 538
                _ = fi16
                _ = v24
                _ = v25
            }
            _ = v15
        } else {
            if false {
                var v26 int8 = 100
                _ = v26
            }
            v27 := GenS2{F0: -778}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v27
        }
    }
    fmt.Println(v1)
    var v28 bool = true
    var v29 float32 = 0.1
    for fi30 := 0; fi30 < 9; fi30++ {
        v28 = false
        switch (fi30 | (^(fi30))) {
        case 4294967295:
            switch "" {
            case "foo bar":
                for fi31 := 0; fi31 < 10; fi31++ {
                    if v28 {
                        v29 -= v29
                        var v32 bool = v28
                        _ = v32
                    }
                    var v33 string = "foo bar"
                    var v34 string = v33
                    var v35 uint32 = 0
                    _ = len(v2)
                    if false {
                        switch ((v33 + "foo bar") + ((v33 + "foo bar") + (("a" + "a") + "a"))) {
                        case "\n\t":
                            switch "foo bar" {
                            case "a":
                                fmt.Println(fi31, fi31)
                            default:
                                switch 128 {
                                case 127:
                                    var v36 uint16 = 255
                                    _ = v36
                                case 32767:
                                    v29 += v29
                                case 42:
                                    v33 = v33
                                }
                            }
                        case "a":
                            v37 := GenS3{F0: v35}
                            _ = v37
                        default:
                            v29 = float32(float64(3.14159))
                        }
                        var v38 float64 = (1.5 + 0.0)
                        _ = v38
                    } else {
                        v29 += (0.1 + (v29 - v29))
                        if len(v2) > 0 {
                            v2[0] = v35
                        }
                        v39 := GenS4{F0: v33, F1: 1, F2: 1}
                        v40 := GenS5{F0: v28, F1: 418, F2: 8}
                        _ = v39
                        _ = v40
                    }
                    var v41 int8 = 124
                    _ = fi31
                    _ = v33
                    _ = v34
                    _ = v35
                    _ = v41
                }
            case "":
                v42 := GenS4{F0: "hello", F1: 1, F2: -9223372036854775808}
                _ = v42
            case "hello":
                if true {
                    var v43 int32 = 2147483647
                    v44 := GenS6{F0: 2, F1: 42}
                    var v45 int16 = -670
                    _ = v43
                    _ = v44
                    _ = v45
                }
            default:
                v29 = float32(0.1)
            }
        case 0:
            switch "\n\t" {
            case "a":
                var v46 int8 = 100
                _ = v46
            }
        case 92:
            fmt.Println(v29, v28)
        }
        var v47 uint64 = v1
        _ = fi30
        _ = v47
    }
    var v48 int16 = -32768
    v29 = v29
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v48
}
