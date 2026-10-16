package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 uint8; F1 int16; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0() uint16 {
    if false {
        var v7 uint = 576
        v8 := []int{}
        for fi9 := 0; fi9 < 5; fi9++ {
            if len(v8) > 0 {
                v8[0] = (fi9 | (-(fi9)))
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    } else {
        if true {
            var v10 int = 128
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := []int{(fi11 ^ (-(fi11)))}
                var v13 uint16 = 255
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        } else {
            var v14 string = "a"
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS1{F0: 9223372036854775807, F1: false, F2: v14}
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS1{F0: 0, F1: true, F2: v14}
                var v19 int16 = 1
                var v20 uint16 = 602
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = v14
        }
        v21 := GenS1{F0: 1, F1: (354 <= 42), F2: (("hello" + (("a" + "hello") + "a")) + "a")}
        var v22 int16 = 32767
        _ = v21
        _ = v22
    }
    var v23 uint16 = 0
    _ = v23
    return (^(v23))
}

func main() {
    v1 := []int{42, 32768}
    var v2 float64 = ((1.5 * ((1.5 - 1.5) * 1e10)) * (0.1 + 1.5))
    var v3 uint16 = 32767
    var v4 uint32 = 0
    var v5 uint = 65535
    v6 := []int{}
    v5 = (v5 - 300)
    var v24 uint = (v5 & 42)
    v25 := GenS2{F0: 2, F1: 0, F2: (v24 & 1)}
    for fi26 := 0; fi26 < 7; fi26++ {
        fmt.Println(v4)
        _ = fi26
    }
    fmt.Println(v3)
    v4 = 1
    v3 -= (v3 - (v3 % ((1) | 1)))
    var v27 uint8 = 188
    _ = len(v1)
    var v28 float32 = (0.0 - 1e10)
    switch 247 {
    case 271:
        var v29 float64 = ((v2 - v2) * 3.14159)
        _ = v29
    default:
        v27 = (v27 & (v27 + v27))
    }
    var v30 float32 = v28
    for fi31 := 0; fi31 < 15; fi31++ {
        fmt.Println(v3, v28, v28)
        if true {
            v32 := GenS1{F0: -551, F1: true, F2: ""}
            for fi33 := 0; fi33 < 9; fi33++ {
                _ = len(v6)
                v28 += v28
                _ = fi33
            }
            var v34 uint8 = (v27 % ((v27) | 1))
            _ = v32
            _ = v34
        } else {
            switch "" {
            case "hello":
                var v35 uint8 = (v27 * (v27 << ((65535) & 63)))
                _ = v35
            case "":
                v36 := GenS2{F0: (v27 - v27), F1: -533, F2: 211}
                _ = v36
            case "foo bar":
                var v37 int = (fi31 - fi31)
                _ = v37
            default:
                _ = len(v1)
            }
            v2 += v2
            var v38 uint8 = (v27 >> (((v5 + 2)) & 63))
            _ = v38
        }
        if len(v6) > 0 {
            v6[0] = (fi31 | (fi31 / ((int(v24)) | 1)))
        }
        _ = fi31
    }
    if (genFn0() >= (v3 << ((65535) & 63))) {
        if true {
            switch "" {
            case "hello":
                if false {
                    fmt.Println(v30)
                    var v39 int16 = 42
                    _ = v39
                }
            }
            for fi40 := 0; fi40 < 17; fi40++ {
                fmt.Println(v4)
                if ((1 == 0) && true) {
                    fmt.Println(v3)
                } else {
                    if (!true) {
                        switch 1 {
                        case 9223372036854775807:
                            v2 += v2
                        }
                        for fi41 := 0; fi41 < 7; fi41++ {
                            switch "foo bar" {
                            case "foo bar":
                                v3 -= uint16(v27)
                            case "a":
                                if ((fi40 * fi40) > (fi41 / (((fi41 + -2147483648)) | 1))) {
                                    v30 -= v30
                                }
                            }
                            v4 += 65535
                            _ = fi41
                        }
                    } else {
                        v24 -= v24
                        if true {
                            _ = len(v6)
                            fmt.Println(v5)
                            v42 := GenS3{F0: int64(v24)}
                            _ = v42
                        }
                        v3 -= (v3 & 65535)
                    }
                }
                _ = fi40
            }
            switch (("foo bar" + "") + "foo bar") {
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    } else {
        for fi43 := 0; fi43 < 12; fi43++ {
            v44 := GenS2{F0: v27, F1: 32767, F2: 42}
            var v45 uint64 = uint64(v27)
            if ((3.14159 / float32(v2)) == v28) {
                v46 := GenS2{F0: v27, F1: 42, F2: (v5 / ((18446744073709551615) | 1))}
                _ = v46
            }
            _ = fi43
            _ = v44
            _ = v45
        }
    }
    var v47 uint64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
