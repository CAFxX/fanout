package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint64, p2 uint16) uint8 {
    switch ("hello" + ("foo bar" + ("" + "hello"))) {
    case "\n\t":
        for fi7 := 0; fi7 < 10; fi7++ {
            p2 = (p2 * p2)
            p0 = (p0 / (((p2 & 32767)) | 1))
            _ = fi7
        }
    default:
        p1 += (p1 % ((2) | 1))
    }
    fmt.Println(p0)
    p2 += 69
    if true {
        p0 += (^(p2))
        v8 := []int{2147483647}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 131
    var v3 uint = 992
    var v4 int = -9223372036854775808
    var v5 float32 = float32((1e10 / 1.5))
    v6 := []int{}
    switch (v4 * int(v1)) {
    case 9223372036854775807:
        var v9 float64 = -2.25
        _ = v9
    case 1:
        for fi10 := 0; fi10 < 15; fi10++ {
            v3 += uint(v2)
            v4 -= fi10
            v11 := GenS1{F0: (v1 | v1), F1: 2147483647}
            _ = fi10
            _ = v11
        }
    case 32768:
        if (true && (((127 < 42) && false) && true)) {
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 uint32 = 4294967295
                _ = fi12
                _ = v13
            }
        }
    default:
        v4 = (-(v4))
    }
    switch (("" + "a") + "") {
    case "a":
        fmt.Println(v4, v1, v1)
    default:
        v6 = append(v6, v4)
    }
    _ = len(v6)
    fmt.Println(v4)
    if false {
        _ = len(v6)
        for fi14 := 0; fi14 < 5; fi14++ {
            if false {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if false {
                    fmt.Println(v1, v2, fi14)
                    var v15 uint8 = genFn0(65535, (v2 * v2), 65535)
                    _ = v15
                } else {
                    if true {
                        var v16 uint64 = (v2 >> ((v3) & 63))
                        var v17 uint16 = 255
                        fmt.Println(v1, fi14, v5)
                        _ = v16
                        _ = v17
                    }
                    v3 = v3
                }
                if (!true) {
                    var v18 int = v4
                    v1 = v1
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if false {
                            fmt.Println(v5, v1, v5)
                            switch (("foo bar" + "hello") + "a") {
                            case "a":
                                v20 := GenS2{F0: 978, F1: (!(((-2.25 / 0.1) - ((float64((v5 * (((float32(((0.1 - (-2.25 - (-2.25 + 0.1))) / 1.5)) - v5) * 0.1) - 0.1))) - float64(-2.25)) + float64(-2.25))) != 3.14159))}
                                _ = v20
                            case "foo bar":
                                if false {
                                    if (((1.5 + (v5 - v5)) - (v5 + v5)) <= 1.5) {
                                        var v21 uint8 = 2
                                        v5 = v5
                                        _ = v21
                                    }
                                }
                            case "\n\t":
                                var v22 string = "hello"
                                _ = v22
                            default:
                                switch (fi19 >> (((v3 % ((v3) | 1))) & 63)) {
                                case -134:
                                    v23 := GenS3{F0: int64(v2)}
                                    _ = v23
                                case 32767:
                                    var v24 uint64 = v2
                                    _ = v24
                                }
                                fmt.Println(v4, v1, fi14)
                            }
                        } else {
                            var v25 uint8 = 255
                            v1 = (v1 & (v1 << ((v3) & 63)))
                            _ = v25
                        }
                        v6 = append(v6, fi19)
                        _ = fi19
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            }
            fmt.Println(fi14, v1, v4)
            _ = fi14
        }
    } else {
        switch 0 {
        case 0:
            fmt.Println(v2, v3)
        case -786:
            fmt.Println(v2)
        default:
            v2 -= (v2 << ((v3) & 63))
        }
        var v26 string = ((("a" + "") + "a") + "a")
        v3 = (v3 % ((v3) | 1))
        _ = v26
    }
    if (false || true) {
        if false {
            v27 := GenS3{F0: -9223372036854775808}
            var v28 string = ("hello" + "a")
            _ = v27
            _ = v28
        } else {
            v3 -= (v3 % ((0) | 1))
            fmt.Println(v4)
        }
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 bool = false
        _ = len(v6)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi29
        _ = v30
    }
    v31 := GenS3{F0: 1}
    for fi32 := 0; fi32 < 20; fi32++ {
        v33 := GenS2{F0: uint(v1), F1: true}
        var v34 int64 = 1
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
