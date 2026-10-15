package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int64 = -141
    v2 := []int{1, 0, -843}
    var v3 int32 = 1
    var v4 uint = 1
    switch 42 {
    case 1:
        var v5 string = "\n\t"
        _ = v5
    default:
        switch "" {
        case "a":
            var v6 float32 = 3.14159
            _ = v6
        default:
            var v7 uint = v4
            _ = v7
        }
    }
    switch "" {
    case "\n\t":
        if (0.1 <= 1.5) {
            if false {
                var v8 uint64 = 4294967295
                _ = v8
            }
            if (false && (true && (1 != 0))) {
                var v9 string = ""
                for fi10 := 0; fi10 < 15; fi10++ {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        _ = len(v2)
                        v3 = (v3 >> (((v4 % (((v4 % ((v4) | 1))) | 1))) & 63))
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 uint32 = 42
                        v14 := GenS1{F0: 100, F1: 0, F2: (fi10 >> ((v4) & 63))}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                }
                _ = v9
            } else {
                v15 := GenS1{F0: -32, F1: 1, F2: -9223372036854775808}
                if ((true && (32767 == 2)) || true) {
                    if (-1 > (v3 - v3)) {
                        fmt.Println(v1, v3)
                        var v16 uint = v4
                        v3 += (^(v3))
                        _ = v16
                    } else {
                        v3 -= v3
                        v4 = v4
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v1 = (v1 & 9223372036854775807)
                _ = v15
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                v4 -= v4
                for fi18 := 0; fi18 < 9; fi18++ {
                    fmt.Println(v3)
                    var v19 int32 = v3
                    v3 += (v3 - -2147483648)
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        }
    case "foo bar":
        fmt.Println(v1, v4)
    }
    fmt.Println(v4, v1)
    v20 := GenS1{F0: -1, F1: 286, F2: -2147483648}
    v2 = append(v2, 2147483647)
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 uint64 = 4294967295
    if false {
        var v22 float64 = (0.0 / -2.25)
        _ = v22
    } else {
        switch 4294967295 {
        case 32768:
            if (!true) {
                if len(v2) > 0 {
                    v2[0] = -1
                }
                v23 := GenS2{F0: v4}
                v3 = (v3 - (v3 * int32(v21)))
                _ = v23
            }
        default:
            for fi24 := 0; fi24 < 5; fi24++ {
                v2 = append(v2, 32767)
                _ = fi24
            }
        }
        for fi25 := 0; fi25 < 16; fi25++ {
            if len(v2) > 0 {
                v2[0] = fi25
            }
            fmt.Println(v21, v3)
            _ = fi25
        }
        if (-900 <= 377) {
            _ = len(v2)
            _ = len(v2)
            if true {
                v26 := GenS1{F0: 0, F1: (v21 ^ (v21 / ((v21) | 1))), F2: 4294967295}
                for fi27 := 0; fi27 < 15; fi27++ {
                    switch (^(fi27)) {
                    case 32767:
                        v21 -= v21
                    case 9223372036854775807:
                        for fi28 := 0; fi28 < 1; fi28++ {
                            fmt.Println(v21, v3)
                            v29 := GenS2{F0: v4}
                            _ = fi28
                            _ = v29
                        }
                    case 42:
                        for fi30 := 0; fi30 < 9; fi30++ {
                            var v31 float32 = ((0.1 * (3.14159 + 1.5)) + 1e10)
                            v32 := GenS1{F0: 0, F1: 1, F2: fi30}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    default:
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v3 += (v3 << (((^(v4))) & 63))
                            v2 = append(v2, fi27)
                            if ((true || true) || ((true && false) || false)) {
                                fmt.Println(v3, v3)
                            } else {
                                var v34 int16 = -69
                                var v35 bool = (128 > 45)
                                _ = v34
                                _ = v35
                            }
                            _ = fi33
                        }
                    }
                    v36 := GenS2{F0: 2}
                    _ = fi27
                    _ = v36
                }
                _ = v26
            } else {
                if true {
                    fmt.Println(v4, v3, v1)
                    var v37 uint32 = 197
                    var v38 uint32 = 562
                    _ = v37
                    _ = v38
                } else {
                    v39 := GenS3{F0: 32767, F1: uint(v21), F2: 128}
                    _ = len(v2)
                    v40 := GenS4{F0: 65535}
                    _ = v39
                    _ = v40
                }
            }
        } else {
            v41 := GenS3{F0: 0, F1: (^(v4)), F2: 42}
            _ = v41
        }
    }
    fmt.Println(v21, v3, v1)
    v42 := GenS5{F0: 32767, F1: 2}
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v42
}
