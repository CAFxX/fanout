package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint64; F1 int; F2 int }
func genFn0() int16 {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint = 1
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: 127, F1: 1, F2: 1000}
            var v9 bool = (-1 < -128)
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        var v10 int = 42
        v11 := GenS1{F0: 1, F1: int8(v10), F2: 32767}
        var v12 int32 = 42
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 float64 = (3.14159 - 3.14159)
    v14 := GenS2{F0: true, F1: -9223372036854775808}
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 359
    var v3 int8 = -128
    v4 := []uint32{2147483648, 788, 65535, 130}
    if true {
        var v15 bool = (260 > 1)
        v15 = false
        v2 = v2
        _ = v15
    } else {
        v2 = (v2 ^ 530)
        if ((42 == 784) && false) {
            var v16 uint8 = 143
            v17 := GenS3{F0: 1000, F1: (^(v3))}
            switch -2147483648 {
            case -237:
                for fi18 := 0; fi18 < 20; fi18++ {
                    if true {
                        for fi19 := 0; fi19 < 5; fi19++ {
                            v2 = (v2 | 42)
                            _ = fi19
                        }
                    } else {
                        v16 = v16
                        _ = len(v1)
                    }
                    _ = fi18
                }
            case 128:
                v3 -= (v3 & v3)
            default:
                var v20 uint32 = 2147483648
                _ = v20
            }
            _ = v16
            _ = v17
        }
    }
    v21 := GenS4{F0: 651, F1: -2147483648, F2: 9223372036854775807}
    var v22 string = "\n\t"
    if true {
        fmt.Println(v3, v22, v3)
        _ = len(v1)
    }
    switch "a" {
    case "foo bar":
        if true {
            var v23 uint8 = 2
            switch 4294967295 {
            case 997:
                v3 -= v3
            case 396:
                if false {
                    var v24 bool = true
                    for fi25 := 0; fi25 < 5; fi25++ {
                        v26 := GenS1{F0: (v23 % (((^(v23))) | 1)), F1: (v3 - v3), F2: int16(v2)}
                        fmt.Println(v2, fi25, v23)
                        v4 = append(v4, 42)
                        _ = fi25
                        _ = v26
                    }
                    var v27 string = v22
                    _ = v24
                    _ = v27
                }
            case 4294967295:
                v28 := GenS1{F0: (^(v23)), F1: (^(v3)), F2: 32767}
                _ = v28
            }
            _ = v23
        }
    }
    v22 = ""
    if false {
        if (!(!true)) {
            v2 -= v2
            if (!false) {
                v3 = v3
            }
        } else {
            fmt.Println(v22, v22, v22)
        }
        _ = len(v1)
    } else {
        if true {
            v22 = ((v22 + "a") + (((v22 + v22) + "hello") + v22))
        } else {
            v2 = (v2 / ((v2) | 1))
            var v29 int16 = -32768
            var v30 float64 = (float64(0.1) + (-2.25 + 0.0))
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(v22, v3, fi31)
        v32 := GenS4{F0: 0, F1: 721, F2: fi31}
        if (42 >= 2) {
            switch fi31 {
            case -558:
                if false {
                    var v33 int8 = v3
                    var v34 bool = true
                    _ = v33
                    _ = v34
                }
            case -508:
                v3 += (v3 << ((571) & 63))
            default:
                v35 := GenS3{F0: -32768, F1: (v3 << ((255) & 63))}
                _ = v35
            }
            v2 += 65535
        }
        _ = fi31
        _ = v32
    }
    v2 = v2
    v3 = -101
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
