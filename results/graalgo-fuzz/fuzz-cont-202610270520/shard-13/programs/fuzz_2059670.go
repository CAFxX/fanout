package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int16; F1 int8; F2 uint16 }

func main() {
    v1 := []int{4294967295, 308, 1}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 bool = true
    v4 = false
    switch ("foo bar" + "a") {
    case "\n\t":
        for fi5 := 0; fi5 < 10; fi5++ {
            switch (fi5 >> ((2) & 63)) {
            case 32767:
                for fi6 := 0; fi6 < 11; fi6++ {
                    v7 := GenS1{F0: 1, F1: "", F2: 2}
                    _ = len(v1)
                    if (!(v4 && (!v4))) {
                        var v8 uint = 954
                        v8 = (v8 & v8)
                        v9 := GenS1{F0: 9223372036854775807, F1: ("hello" + (("\n\t" + ("foo bar" + "hello")) + "\n\t")), F2: 931}
                        _ = v8
                        _ = v9
                    } else {
                        v10 := GenS2{F0: 1, F1: ("\n\t" + "hello"), F2: -128}
                        v11 := v3[0]
                        if ("\n\t" >= ("\n\t" + "")) {
                            if (v4 || (v4 && true)) {
                                v4 = v4
                            }
                            v12 := GenS2{F0: 4294967295, F1: ("hello" + ""), F2: 106}
                            v1 = append(v1, fi6)
                            switch "\n\t" {
                            case "a":
                                if v4 {
                                    var v13 uint32 = 1
                                    v4 = (217 < 2147483647)
                                    v14 := GenS1{F0: -1, F1: ("" + "foo bar"), F2: 156}
                                    _ = v13
                                    _ = v14
                                }
                            case "\n\t":
                                var v15 int32 = 0
                                _ = v15
                            case "foo bar":
                                v11 = (fi5 - fi6)
                            }
                            _ = v12
                        } else {
                            if ((42 == 65535) && ((1 >= 140) && true)) {
                                v4 = true
                                v11 = (fi6 >> ((1) & 63))
                            }
                            v11 = (-(fi6))
                        }
                        _ = v10
                        _ = v11
                    }
                    var v16 int8 = 0
                    _ = fi6
                    _ = v16
                    _ = v7
                }
            case -9223372036854775808:
                if v4 {
                    switch fi5 {
                    case -6:
                        if v4 {
                            for fi17 := 0; fi17 < 9; fi17++ {
                                v4 = v4
                                v4 = ((!(0.0 >= 3.14159)) || (0 < 0))
                                fmt.Println(fi17, fi17, fi17)
                                _ = fi17
                            }
                            v4 = ((-440 != -50) || v4)
                        }
                    default:
                        for fi18 := 0; fi18 < 19; fi18++ {
                            var v19 float32 = (0.1 / 0.1)
                            var v20 int = 32768
                            switch fi18 {
                            case 109:
                                v4 = v4
                            }
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                    var v21 int8 = -36
                    v4 = v4
                    _ = v21
                } else {
                    v22 := GenS1{F0: 562, F1: "hello", F2: 1}
                    v23 := GenS3{F0: 160}
                    _ = len(v1)
                    _ = v22
                    _ = v23
                }
            case 254:
                v24 := GenS4{F0: 0, F1: 32}
                _ = v24
            default:
                var v25 int16 = 32767
                _ = v25
            }
            _ = fi5
        }
    case "foo bar":
        fmt.Println(v4)
    default:
        delete(v3, 65535)
    }
    v4 = true
    var v26 float64 = 1.5
    _ = len(v3)
    var v27 uint8 = 255
    v28 := GenS4{F0: 32767, F1: 65535}
    switch 32768 {
    case -980:
        v1 = append(v1, 128)
    default:
        var v29 uint32 = 302
        _ = v29
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    if (9223372036854775807 != 0) {
        fmt.Println(v26, v27, v26)
        _ = len(v1)
    } else {
        var v30 uint64 = 1
        var v31 string = "hello"
        v4 = true
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        switch fi32 {
        case -9223372036854775808:
            v33 := GenS5{F0: 1000, F1: -128, F2: 42}
            _ = v33
        default:
            switch "" {
            case "\n\t":
                var v34 uint = 0
                _ = v34
            case "":
                v35 := GenS4{F0: 42, F1: 32767}
                _ = v35
            case "a":
                var v36 uint16 = 111
                _ = v36
            }
        }
        if (uint16(v27) > 1) {
            for fi37 := 0; fi37 < 5; fi37++ {
                v26 = 3.14159
                v26 -= ((v26 * float64(3.14159)) * (1.5 * v26))
                _ = fi37
            }
        }
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
