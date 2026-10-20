package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint32 = 2
    v2 := []uint32{v1, (v1 % ((v1) | 1))}
    var v3 string = "a"
    var v4 string = (v3 + v3)
    v5 := GenS1{F0: (false || ((v1 ^ v1) != v1)), F1: 42}
    v4 = ((v4 + ((v4 + "a") + "")) + "\n\t")
    v3 = v4
    v6 := GenS2{F0: 42, F1: (!(-9223372036854775808 < -2147483648)), F2: 0}
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS2{F0: 9223372036854775807, F1: true, F2: 1}
        for fi9 := 0; fi9 < 17; fi9++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v10 string = v4
            var v11 uint16 = 2
            if (!true) {
                if true {
                    switch 1 {
                    case 1:
                        switch (v10 + v3) {
                        case "foo bar":
                            if false {
                                var v12 bool = ((true && true) && (1 == 1))
                                v13 := GenS2{F0: -1, F1: v12, F2: -1}
                                _ = len(v2)
                                _ = v12
                                _ = v13
                            } else {
                                v3 = ((("" + v3) + v10) + (v4 + v4))
                            }
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    case 32768:
                        v10 = v4
                    case -510:
                        v1 = v1
                    default:
                        v11 -= (v11 + v11)
                    }
                    var v14 uint16 = v11
                    if len(v2) > 0 {
                        v2[0] = (v1 / ((v1) | 1))
                    }
                    _ = v14
                }
            }
            var v15 int = -112
            _ = fi9
            _ = v10
            _ = v11
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    if true {
        v1 = (v1 / ((v1) | 1))
        v4 = v4
        var v16 uint16 = 2
        _ = v16
    } else {
        switch 32768 {
        case 32767:
            var v17 int8 = -1
            _ = v17
        }
        fmt.Println(v1)
        fmt.Println(v3, v3, v1)
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS3{F0: 885}
        if (4294967295 > 4294967295) {
            v1 = v1
        }
        switch -278 {
        case -136:
            var v20 uint8 = 127
            _ = v20
        case 128:
            v21 := GenS1{F0: (42 >= 0), F1: 944}
            _ = v21
        default:
            v1 = v1
        }
        _ = fi18
        _ = v19
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = v1
    }
    if true {
        var v22 uint = 0
        for fi23 := 0; fi23 < 10; fi23++ {
            _ = len(v2)
            _ = fi23
        }
        if false {
            var v24 string = (v4 + (v3 + v3))
            _ = v24
        }
        _ = v22
    } else {
        var v25 uint = 529
        _ = v25
    }
    if (-1 > int8(v1)) {
        switch "hello" {
        case "foo bar":
            v3 = "a"
        case "":
            if false {
                v26 := GenS2{F0: 42, F1: (!(0 == 42)), F2: -32768}
                _ = v26
            }
        default:
            var v27 int16 = 0
            _ = v27
        }
    } else {
        var v28 int32 = 0
        v1 += (v1 & (v1 ^ 2))
        _ = v28
    }
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
