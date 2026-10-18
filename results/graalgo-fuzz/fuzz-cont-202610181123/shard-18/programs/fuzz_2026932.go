package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint8 = 164
    v2 := []uint32{65535, 65535}
    var v3 float32 = 3.14159
    var v4 uint32 = 0
    fmt.Println(v3)
    if (0.0 == ((0.0 / ((1.5 + 0.1) + 0.0)) + 0.1)) {
        v1 -= v1
        if len(v2) > 0 {
            v2[0] = 0
        }
        v5 := GenS1{F0: 901, F1: v1}
        _ = v5
    } else {
        v1 += 42
        if ((1 >= -173) || false) {
            v3 = (v3 + (((v3 - float32(-2.25)) / float32(-2.25)) * 0.1))
            switch 32768 {
            case 4294967295:
                v6 := GenS1{F0: 42, F1: 76}
                _ = v6
            case 9223372036854775807:
                v3 -= 0.0
            default:
                for fi7 := 0; fi7 < 14; fi7++ {
                    fmt.Println(fi7)
                    var v8 float64 = ((1.5 / 3.14159) - 0.1)
                    var v9 float32 = 1e10
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        }
    }
    fmt.Println(v1, v3, v3)
    v3 -= v3
    for fi10 := 0; fi10 < 5; fi10++ {
        v1 = v1
        _ = fi10
    }
    v11 := GenS1{F0: 0, F1: (v1 << ((1) & 63))}
    if (4294967295 >= 1) {
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v4, v1, v1)
            v3 += float32((-2.25 + 0.0))
            v13 := GenS1{F0: 627, F1: v1}
            _ = fi12
            _ = v13
        }
        var v14 int16 = 1
        _ = len(v2)
        _ = v14
    }
    v3 += ((3.14159 * (0.0 - v3)) + 1e10)
    v1 = v1
    fmt.Println(v4, v4, v1)
    v3 -= 1.5
    fmt.Println(v4, v3)
    v15 := GenS1{F0: 59, F1: v1}
    var v16 uint16 = 0
    v17 := GenS2{F0: -32768}
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            switch ("foo bar" + "\n\t") {
            case "a":
                v2 = append(v2, uint32(v16))
            case "":
                var v20 int64 = -738
                _ = v20
            case "\n\t":
                var v21 float32 = 1e10
                _ = v21
            default:
                var v22 bool = ((-55 < -19) && false)
                _ = v22
            }
            if (-1 <= 1) {
                if (995 != 0) {
                    switch "a" {
                    case "hello":
                        if true {
                            v16 = 190
                            if false {
                                var v23 string = "foo bar"
                                _ = v23
                            } else {
                                if (0 < 9223372036854775807) {
                                    v16 += 2
                                    v4 += (v4 & 908)
                                    v24 := GenS2{F0: 946}
                                    _ = v24
                                }
                            }
                            var v25 uint64 = 0
                            _ = v25
                        } else {
                            fmt.Println(v3, v4, v16)
                            switch "\n\t" {
                            case "foo bar":
                                v3 = (3.14159 * 3.14159)
                            default:
                                var v26 string = ("hello" + "foo bar")
                                _ = v26
                            }
                            fmt.Println(v16, v4)
                        }
                    case "foo bar":
                        if ((^(v16)) <= (v16 ^ (v16 | (^(v16))))) {
                            if len(v2) > 0 {
                                v2[0] = v4
                            }
                        }
                    }
                    v4 += (^(v4))
                    var v27 bool = false
                    _ = v27
                } else {
                    v16 += 1
                    v2 = append(v2, 950)
                }
                v16 += (v16 / ((v16) | 1))
                fmt.Println(v16, fi19, fi18)
            }
            fmt.Println(fi19, fi18)
            _ = fi19
        }
        _ = len(v2)
        for fi28 := 0; fi28 < 2; fi28++ {
            v4 = v4
            if (-21 == 100) {
                if true {
                    fmt.Println(fi28, fi28)
                }
                var v29 int8 = 116
                _ = v29
            } else {
                fmt.Println(v16)
            }
            var v30 int8 = int8(v1)
            _ = fi28
            _ = v30
        }
        _ = fi18
    }
    v4 += (v4 << ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
