package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 uint8; F2 uint32 }

func main() {
    var v1 int16 = 401
    v2 := []int{32768}
    var v3 int8 = -128
    v2 = append(v2, 42)
    v4 := GenS1{F0: (-(v1))}
    if false {
        if (773 > 32767) {
            var v5 int16 = v1
            v3 += v3
            v6 := GenS1{F0: v1}
            _ = v5
            _ = v6
        }
        v1 = v1
        for fi7 := 0; fi7 < 3; fi7++ {
            v1 -= (v1 / ((v1) | 1))
            _ = fi7
        }
    } else {
        v1 -= (v1 | v1)
        if false {
            fmt.Println(v3, v3)
        } else {
            v8 := GenS2{F0: (v1 / ((v1) | 1)), F1: "\n\t", F2: "hello"}
            v1 = v1
            if (-9223372036854775808 < 658) {
                if (-1 < 867) {
                    v1 += v1
                    v1 = v1
                } else {
                    v3 += (^(v3))
                    switch -9223372036854775808 {
                    case 617:
                        switch ("foo bar" + "hello") {
                        case "":
                            var v9 int = -9223372036854775808
                            _ = v9
                        default:
                            if (0 <= 32767) {
                                for fi10 := 0; fi10 < 19; fi10++ {
                                    v11 := GenS3{F0: 0}
                                    for fi12 := 0; fi12 < 9; fi12++ {
                                        var v13 int = (fi12 << ((644) & 63))
                                        _ = fi12
                                        _ = v13
                                    }
                                    switch (-(fi10)) {
                                    case 1:
                                        var v14 float32 = 0.1
                                        _ = v14
                                    case 127:
                                        for fi15 := 0; fi15 < 11; fi15++ {
                                            var v16 int = -804
                                            _ = fi15
                                            _ = v16
                                        }
                                    default:
                                        if (!(255 <= 923)) {
                                            v17 := GenS3{F0: 580}
                                            v3 -= (v3 + v3)
                                            var v18 int = -20
                                            _ = v17
                                            _ = v18
                                        }
                                    }
                                    _ = fi10
                                    _ = v11
                                }
                                v1 = v1
                            }
                        }
                    case 32767:
                        fmt.Println(v1, v3)
                    }
                }
                v1 -= 1
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 uint = 623
                    v21 := GenS4{F0: -533}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            _ = v8
        }
        var v22 bool = true
        _ = v22
    }
    v1 -= v1
    fmt.Println(v3, v1)
    var v23 int64 = -9223372036854775808
    if false {
        if (true && false) {
            v23 += -1
        }
        v2 = append(v2, -9223372036854775808)
    } else {
        if (42 == 2) {
            v23 += v23
        } else {
            var v24 int = 128
            fmt.Println(v23, v23)
            _ = v24
        }
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 += 191
        for fi25 := 0; fi25 < 2; fi25++ {
            for fi26 := 0; fi26 < 8; fi26++ {
                v1 += (v1 + v1)
                v27 := GenS5{F0: 2}
                fmt.Println(v3, fi25)
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    } else {
        switch -821 {
        case -9223372036854775808:
            v3 += int8(v1)
        case 1:
            if true {
                for fi28 := 0; fi28 < 19; fi28++ {
                    fmt.Println(fi28, v1)
                    v3 -= 107
                    _ = fi28
                }
                fmt.Println(v23, v23)
            } else {
                v29 := GenS4{F0: v23}
                var v30 float32 = float32(0.0)
                _ = v29
                _ = v30
            }
        case -442:
            v3 += v3
        }
        v31 := GenS2{F0: 42, F1: "", F2: "a"}
        if len(v2) > 0 {
            v2[0] = 827
        }
        _ = v31
    }
    var v32 int = 42
    v1 += (-(v1))
    fmt.Println(v32, v1, v1)
    for fi33 := 0; fi33 < 6; fi33++ {
        switch "hello" {
        case "\n\t":
            v32 -= 2147483647
        case "":
            var v34 uint64 = 18446744073709551615
            _ = v34
        }
        var v35 uint32 = 65535
        for fi36 := 0; fi36 < 3; fi36++ {
            v37 := GenS6{F0: 128}
            v38 := GenS7{F0: 127, F1: 2, F2: v35}
            v3 = (v3 * v3)
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi33
        _ = v35
    }
    v2 = append(v2, (v32 % (((v32 - v32)) | 1)))
    fmt.Println(v1, v3)
    v1 -= v1
    fmt.Println(v32, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
}
