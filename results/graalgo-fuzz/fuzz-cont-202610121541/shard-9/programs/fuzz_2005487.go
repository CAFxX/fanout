package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int64; F1 int16; F2 uint32 }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    var v3 int = -2147483648
    var v4 uint32 = 42
    var v5 uint64 = 42
    v6 := []string{("hello" + ((("\n\t" + "hello") + "") + ((((("hello" + "\n\t") + "a") + "hello") + "hello") + "a")))}
    var v7 int16 = -1
    switch v3 {
    case 1:
        for fi8 := 0; fi8 < 15; fi8++ {
            v4 += 2
            v4 = v4
            _ = fi8
        }
    case 42:
        switch (-(v3)) {
        case -1:
            if (-160 >= int64(v4)) {
                v9 := GenS1{F0: false, F1: -1, F2: 65535}
                _ = v9
            } else {
                if true {
                    switch "foo bar" {
                    case "":
                        v7 = int16(v3)
                    }
                } else {
                    for fi10 := 0; fi10 < 20; fi10++ {
                        v5 = (^(v5))
                        fmt.Println(fi10, v3, v3)
                        _ = fi10
                    }
                    var v11 int16 = v7
                    _ = v11
                }
            }
        case -752:
            fmt.Println(v5, v5)
        case 1:
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 int16 = (-(v7))
                v5 = v5
                _ = fi12
                _ = v13
            }
        }
    }
    var v14 int16 = v7
    var v15 uint32 = v4
    var v16 uint8 = 1
    if true {
        var v17 uint16 = 1
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: (!true), F1: (v3 ^ 744), F2: 748}
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        switch v3 {
        case 42:
            v3 += (v3 - 45)
        case 127:
            delete(v1, (v3 ^ v3))
        }
        v6 = append(v6, "")
    }
    if false {
        switch (("a" + "hello") + "a") {
        case "\n\t":
            v20 := GenS1{F0: true, F1: (v3 % ((v3) | 1)), F2: 65535}
            _ = v20
        case "foo bar":
            v21 := GenS1{F0: (!true), F1: v3, F2: 2}
            _ = v21
        }
    } else {
        if (!true) {
            if (("" == (("foo bar" + "foo bar") + "foo bar")) && (uint8(v4) < v16)) {
                var v22 uint32 = (v4 + 4294967295)
                _ = v22
            }
        } else {
            fmt.Println(v2)
            switch (v3 % ((32767) | 1)) {
            case -9223372036854775808:
                v23 := GenS2{F0: v16}
                _ = v23
            case 42:
                if (2147483647 == (^(v2))) {
                    v7 += v14
                    var v24 uint = 42
                    _ = v24
                } else {
                    switch (-(v3)) {
                    case 0:
                        if true {
                            var v25 uint = 233
                            for fi26 := 0; fi26 < 13; fi26++ {
                                var v27 uint32 = (v15 & 956)
                                var v28 uint8 = v16
                                if ("hello" >= "\n\t") {
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        if (fi26 <= -771) {
                                            for fi29 := 0; fi29 < 16; fi29++ {
                                                if true {
                                                    var v30 uint32 = 2147483648
                                                    if (v2 > 1) {
                                                        v6 = append(v6, "\n\t")
                                                        fmt.Println(v28)
                                                        fmt.Println(v15)
                                                    } else {
                                                        var v31 int16 = v14
                                                        v28 -= (^(v16))
                                                        _ = v31
                                                    }
                                                    var v32 int16 = v7
                                                    _ = v30
                                                    _ = v32
                                                } else {
                                                    if true {
                                                        v33 := GenS3{F0: 1, F1: (^(v16))}
                                                        var v34 bool = (-1 > int8(v28))
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        if (1e10 > (1e10 - float64(float32((float64(float32(1.5)) * 0.0))))) {
                                                            v35 := GenS3{F0: (^(v28)), F1: (^(v28))}
                                                            _ = v35
                                                        } else {
                                                            var v36 int = -2147483648
                                                            v1[(fi29 ^ (v3 - fi29))] = false
                                                            _ = v36
                                                        }
                                                    }
                                                    var v37 uint16 = 818
                                                    var v38 uint8 = (^(v28))
                                                    _ = v37
                                                    _ = v38
                                                }
                                                _ = len(v1)
                                                var v39 int = (fi26 & (fi26 << ((v25) & 63)))
                                                v40 := v1[32767]
                                                fmt.Println(v2, v16)
                                                _ = fi29
                                                _ = v39
                                                _ = v40
                                            }
                                            if false {
                                                v14 = (v7 & (v7 + v7))
                                                v2 -= v2
                                                v7 = -32768
                                            }
                                        }
                                    default:
                                        switch 1 {
                                        case 0:
                                            var v41 int16 = v14
                                            _ = v41
                                        case -9223372036854775808:
                                            v5 = (v5 | (v5 | 18446744073709551615))
                                        }
                                    }
                                } else {
                                    v42 := GenS4{F0: int64(v7), F1: 1000, F2: (v4 ^ 2)}
                                    _ = len(v6)
                                    _ = v42
                                }
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            v4 = v15
                            _ = v25
                        }
                    case 32768:
                        for fi43 := 0; fi43 < 3; fi43++ {
                            v44 := GenS5{F0: "hello"}
                            _ = fi43
                            _ = v44
                        }
                    }
                }
            default:
                switch v3 {
                case -691:
                    var v45 uint = 255
                    _ = v45
                case 4294967295:
                    delete(v1, (v3 << ((1) & 63)))
                default:
                    fmt.Println(v2)
                }
            }
        }
    }
    switch 1 {
    case 800:
        var v46 float32 = float32(1e10)
        _ = v46
    default:
        v47 := GenS3{F0: v16, F1: v16}
        _ = v47
    }
    for fi48 := 0; fi48 < 2; fi48++ {
        fmt.Println(v3)
        if false {
            v6 = append(v6, "foo bar")
        } else {
            v2 += (v2 - v2)
            v49 := v1[0]
            switch 128 {
            case -803:
                var v50 uint16 = 0
                _ = v50
            }
            _ = v49
        }
        _ = fi48
    }
    var v51 int16 = (v14 >> ((18446744073709551615) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
}
