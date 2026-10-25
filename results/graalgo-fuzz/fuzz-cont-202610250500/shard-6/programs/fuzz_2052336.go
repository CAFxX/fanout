package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint16 }
func genFn0(p0 int8, p1 uint32) int64 {
    switch -473 {
    case 9223372036854775807:
        switch ("foo bar" + "") {
        case "foo bar":
            p1 = p1
        case "":
            p0 = p0
        case "a":
            p1 -= (p1 / ((p1) | 1))
        default:
            switch ("" + "") {
            case "":
                p1 -= (p1 | 1)
            }
        }
    case 1:
        v4 := []int{}
        _ = v4
    }
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            p1 = (^(p1))
            var v7 int32 = -2147483648
            var v8 uint32 = (^(p1))
            v9 := []int{(fi6 & (-(fi5))), fi6}
            fmt.Println(p0)
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        p1 += (p1 ^ (p1 + p1))
        fmt.Println(p0)
        p1 += (^(p1))
        _ = fi10
    }
    _ = p0
    _ = p1
    return -693
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]string)
    var v3 uint = 65535
    if true {
        switch "a" {
        case "\n\t":
            fmt.Println(v3, v1)
        case "a":
            if (216 != 298) {
                if (v3 <= (^(v3))) {
                    v11 := GenS1{F0: 42, F1: 0, F2: 94}
                    _ = v11
                }
            }
        }
        switch 42 {
        case 554:
            v12 := GenS2{F0: true, F1: 4294967295, F2: 127}
            _ = v12
        default:
            v13 := GenS3{F0: -762}
            _ = v13
        }
        v1 -= (v1 >> ((42) & 63))
    }
    var v14 int16 = 1000
    var v15 int = -819
    if false {
        var v16 bool = (false || true)
        var v17 int16 = v1
        _ = v16
        _ = v17
    }
    v18 := []int{(-(v15)), v15, v15, (v15 ^ v15)}
    v15 += -2147483648
    var v19 uint16 = 2
    switch v15 {
    case 32767:
        if (!(!true)) {
            var v20 uint = 266
            _ = v20
        }
    }
    var v21 uint = (v3 / ((v3) | 1))
    fmt.Println(v21)
    var v22 int = v15
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := GenS2{F0: (false || (v21 == 2)), F1: (^(v3)), F2: 0}
        switch (fi23 ^ -697) {
        case 4294967295:
            v3 = 2
        default:
            delete(v2, (v15 % ((-556) | 1)))
        }
        _ = fi23
        _ = v24
    }
    if len(v18) > 0 {
        v18[0] = 32768
    }
    v22 = v22
    v25 := GenS1{F0: 42, F1: 757, F2: (v19 ^ (v19 % (((v19 << ((979) & 63))) | 1)))}
    switch -9223372036854775808 {
    case -1:
        if ((((("foo bar" + "foo bar") + "hello") + ("a" + (("a" + ("\n\t" + (("hello" + "foo bar") + "\n\t"))) + "hello"))) + "") != "foo bar") {
            switch ("hello" + ("hello" + "hello")) {
            case "a":
                var v26 uint8 = 42
                _ = v26
            default:
                fmt.Println(v15)
            }
            if (0 != 18446744073709551615) {
                var v27 int = v22
                _ = v27
            }
            switch 32768 {
            case -807:
                if true {
                    fmt.Println(v19)
                } else {
                    var v28 uint64 = 42
                    for fi29 := 0; fi29 < 7; fi29++ {
                        fmt.Println(v19, fi29, v3)
                        v30 := GenS4{F0: (v28 ^ (v28 + (v28 + v28))), F1: (v19 * (v19 / (((v19 ^ 42)) | 1))), F2: (v28 | (^(v28)))}
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            case 325:
                v31 := GenS5{F0: genFn0(-1, 2147483647), F1: 1}
                _ = v31
            }
        } else {
            v1 += v14
        }
    default:
        if (!("foo bar" != ("" + ("hello" + ("\n\t" + ("hello" + "a")))))) {
            v1 -= v1
            delete(v2, v15)
        }
    }
    var v32 int64 = -9223372036854775808
    fmt.Println(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v32
}
