package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64) int64 {
    fmt.Println(p0)
    var v7 float64 = (-2.25 / 3.14159)
    var v8 float32 = float32(float64((3.14159 - 0.0)))
    _ = p0
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 uint64 = 1
    v2 := []string{}
    var v3 int8 = 42
    var v4 int16 = -1
    var v5 bool = (("foo bar" + "foo bar") == ("" + ((("\n\t" + "") + "foo bar") + "\n\t")))
    var v6 uint32 = 42
    v4 = (v4 / (((v4 * v4)) | 1))
    if (("" + "") <= "") {
        var v9 int32 = 0
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(v4, v4)
            for fi11 := 0; fi11 < 18; fi11++ {
                switch "\n\t" {
                case "a":
                    if false {
                        fmt.Println(v1, v1, fi11)
                        v1 += v1
                    } else {
                        var v12 uint8 = 42
                        var v13 int16 = 32767
                        if (!((128 != 9223372036854775807) && (!(65535 >= 0)))) {
                            var v14 int = fi10
                            fmt.Println(v12, v1)
                            _ = v14
                        } else {
                            var v15 int32 = v9
                            var v16 uint8 = 255
                            _ = v15
                            _ = v16
                        }
                        _ = v12
                        _ = v13
                    }
                }
                v5 = (41 == fi10)
                if len(v2) > 0 {
                    v2[0] = ("foo bar" + "\n\t")
                }
                _ = fi11
            }
            v9 = v9
            _ = fi10
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                switch (^(fi18)) {
                case 128:
                    switch (-(fi18)) {
                    case -9:
                        v4 = 1000
                    }
                case -2147483648:
                    if (v3 != (v3 << ((65535) & 63))) {
                        fmt.Println(v4, fi17)
                        if v5 {
                            if ((!v5) && true) {
                                fmt.Println(fi17)
                                if len(v2) > 0 {
                                    v2[0] = ((("\n\t" + "") + "foo bar") + "a")
                                }
                                v19 := GenS1{F0: (v4 >> ((65535) & 63))}
                                _ = v19
                            } else {
                                var v20 int64 = genFn0(v1)
                                var v21 string = "\n\t"
                                if len(v2) > 0 {
                                    v2[0] = v21
                                }
                                fmt.Println(v6, v1, v20)
                                _ = v20
                                _ = v21
                            }
                            v22 := GenS2{F0: 42, F1: 100, F2: -32768}
                            fmt.Println(v9, fi17, v1)
                            _ = v22
                        } else {
                            v2 = append(v2, "\n\t")
                            v6 = v6
                        }
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    } else {
                        v6 += v6
                    }
                case -130:
                    v4 = 1000
                }
                _ = fi18
            }
            _ = fi17
        }
        _ = v9
    }
    var v23 int16 = v4
    for fi24 := 0; fi24 < 13; fi24++ {
        fmt.Println(v3, v3)
        _ = fi24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS1{F0: 1}
        fmt.Println(v6)
        var v27 uint32 = (^(v6))
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := GenS2{F0: uint(v6), F1: v3, F2: v23}
    for fi29 := 0; fi29 < 14; fi29++ {
        switch -2147483648 {
        case 9223372036854775807:
            v2 = append(v2, "foo bar")
        case 128:
            v1 -= v1
        default:
            _ = len(v2)
        }
        _ = fi29
    }
    switch 4294967295 {
    case 42:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 0:
        var v30 int32 = -2147483648
        _ = v30
    default:
        v2 = append(v2, "hello")
    }
    v31 := GenS2{F0: 1, F1: (v3 - v3), F2: -32768}
    v4 = v4
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 float64 = 0.0
        v3 = v3
        _ = fi32
        _ = v33
    }
    v6 = v6
    for fi34 := 0; fi34 < 1; fi34++ {
        for fi35 := 0; fi35 < 3; fi35++ {
            if (!v5) {
                v5 = v5
                var v36 uint16 = 2
                _ = v36
            }
            _ = fi35
        }
        v3 = (v3 & (v3 % (((v3 + -1)) | 1)))
        v5 = false
        _ = fi34
    }
    v37 := GenS3{F0: 433}
    if (0 <= uint(v1)) {
        var v38 int64 = -1
        if len(v2) > 0 {
            v2[0] = (("" + "foo bar") + (("a" + "foo bar") + "hello"))
        }
        _ = v38
    } else {
        var v39 uint = 2
        _ = v39
    }
    v4 -= (v23 / ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
