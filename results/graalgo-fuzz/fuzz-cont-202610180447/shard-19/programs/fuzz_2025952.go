package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint8, p1 int16) uint64 {
    p1 = p1
    v5 := []int{9223372036854775807}
    var v6 uint16 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1() int32 {
    v7 := GenS1{F0: -2147483648, F1: 2, F2: (!false)}
    v8 := []int{-9223372036854775808, 127, 128, 32768}
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int = fi9
        var v11 uint64 = 18446744073709551615
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    return -297
}

func main() {
    var v1 bool = (1 < 255)
    v2 := make(map[uint32]string)
    var v3 uint16 = 1
    v4 := []int{639, -466, 0}
    fmt.Println(v3, v1)
    switch "a" {
    case "hello":
        for fi12 := 0; fi12 < 18; fi12++ {
            switch (fi12 - (fi12 << ((42) & 63))) {
            case 4294967295:
                v13 := GenS1{F0: 9223372036854775807, F1: 2, F2: false}
                _ = v13
            case 42:
                fmt.Println(v3)
            case -2147483648:
                v14 := GenS1{F0: (fi12 ^ fi12), F1: 1, F2: (!v1)}
                _ = v14
            default:
                if ("foo bar" <= (("" + "\n\t") + "hello")) {
                    var v15 uint16 = v3
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                v3 = v3
                _ = fi16
            }
            _ = fi12
        }
    default:
        v17 := GenS2{F0: 2, F1: (!v1)}
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        if v1 {
            if (2147483647 != 939) {
                var v19 uint = 255
                if true {
                    v20 := GenS1{F0: fi18, F1: 191, F2: ((v3 >= (v3 + v3)) || v1)}
                    _ = v20
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v22 int64 = 0
                    if true {
                        v23 := GenS1{F0: (fi21 & (^(fi18))), F1: genFn0(42, 260), F2: v1}
                        switch (fi21 / ((2147483647) | 1)) {
                        case -7:
                            fmt.Println(fi21)
                        default:
                            v3 -= (v3 >> (((v19 + v19)) & 63))
                        }
                        _ = v23
                    }
                    var v24 bool = (-468 <= 754)
                    var v25 string = "foo bar"
                    v19 = v19
                    if false {
                        var v26 int8 = -128
                        var v27 int64 = int64(v26)
                        var v28 uint = v19
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(fi18, v25)
                        var v29 uint16 = 42
                        v25 = (v25 + "hello")
                        _ = v29
                    }
                    v30 := GenS3{F0: 198}
                    _ = fi21
                    _ = v22
                    _ = v24
                    _ = v25
                    _ = v30
                }
                _ = v19
            }
        }
        var v31 float64 = (float64((0.0 - 1e10)) * ((-2.25 + 0.1) * 1.5))
        _ = fi18
        _ = v31
    }
    switch 127 {
    case 4294967295:
        v1 = v1
    case 367:
        var v32 uint64 = 1
        _ = v32
    }
    var v33 string = ""
    for fi34 := 0; fi34 < 18; fi34++ {
        fmt.Println(v3, v3)
        _ = fi34
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 bool = (!(v33 == v33))
        delete(v2, 0)
        v1 = (((v3 ^ v3) > (v3 + (v3 & v3))) && (627 != -1))
        _ = fi35
        _ = v36
    }
    var v37 uint = 42
    v38 := GenS4{F0: 170, F1: genFn0(0, 914)}
    if v1 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v39 string = (("a" + "\n\t") + "a")
        delete(v2, 303)
        _ = v39
    } else {
        v1 = ((32767 != 1) || true)
        v40 := GenS1{F0: 32767, F1: 1, F2: false}
        _ = v40
    }
    switch (v33 + v33) {
    case "foo bar":
        for fi41 := 0; fi41 < 6; fi41++ {
            for fi42 := 0; fi42 < 11; fi42++ {
                v43 := GenS1{F0: (fi41 + fi42), F1: 425, F2: ((v1 && (2 >= 18446744073709551615)) || (0 <= -32768))}
                v1 = v1
                _ = fi42
                _ = v43
            }
            v33 = "foo bar"
            if len(v4) > 0 {
                v4[0] = fi41
            }
            _ = fi41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
}
