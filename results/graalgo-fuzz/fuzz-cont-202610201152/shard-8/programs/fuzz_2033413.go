package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int, p2 string) int64 {
    p1 = 32768
    if (p2 != (p2 + "foo bar")) {
        if ((!("" != "\n\t")) || true) {
            var v4 int16 = 1000
            _ = v4
        }
        v5 := []int{p1, p1}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 255
    var v3 int8 = -124
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: (v3 & v3), F1: -316, F2: 1}
        fmt.Println(v1, v1, fi6)
        _ = fi6
        _ = v7
    }
    v3 -= (v3 | (v3 >> ((0) & 63)))
    fmt.Println(v2, v3)
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(v3, v2, v3)
        if (!(!((2147483648 <= 4294967295) && (!false)))) {
            v9 := GenS1{F0: -3, F1: -438, F2: 2}
            if true {
                var v10 uint = uint(v1)
                _ = v10
            }
            _ = v9
        }
        _ = fi8
    }
    v11 := []int{9223372036854775807, 127, -756, -496}
    for fi12 := 0; fi12 < 17; fi12++ {
        switch int(v3) {
        case 1:
            v1 = (^(v1))
        case 2147483647:
            var v13 string = ("hello" + "hello")
            _ = v13
        case -1:
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS2{F0: 65535}
                _ = fi14
                _ = v15
            }
        default:
            var v16 uint64 = v1
            _ = v16
        }
        if false {
            v3 = v3
            var v17 int32 = 0
            switch "hello" {
            case "\n\t":
                v18 := GenS2{F0: uint32(v2)}
                _ = v18
            case "foo bar":
                var v19 bool = (!false)
                _ = v19
            case "":
                v20 := GenS2{F0: 1}
                _ = v20
            default:
                fmt.Println(fi12)
            }
            _ = v17
        }
        _ = fi12
    }
    if len(v11) > 0 {
        v11[0] = 916
    }
    fmt.Println(v3, v2, v3)
    v2 = 127
    switch ("hello" + "") {
    case "":
        switch -477 {
        case 186:
            var v21 uint16 = 0
            _ = v21
        }
    default:
        if (genFn0(v2, 128, ("foo bar" + "")) != 456) {
            var v22 uint = 4294967295
            _ = v22
        } else {
            v11 = append(v11, 32768)
        }
    }
    var v23 uint = 411
    v23 = 65535
    if (("\n\t" + ("hello" + (("" + ("a" + ("hello" + "foo bar"))) + "\n\t"))) < "a") {
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 uint64 = uint64(v23)
            _ = fi24
            _ = v25
        }
    } else {
        v26 := GenS1{F0: v3, F1: 991, F2: 1}
        v2 += v2
        _ = v26
    }
    if false {
        if true {
            v27 := GenS2{F0: 65535}
            _ = v27
        } else {
            v3 = (v3 & 100)
        }
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            fmt.Println(fi29, fi28)
            var v30 uint32 = 0
            v2 += v2
            _ = fi29
            _ = v30
        }
        fmt.Println(v2, fi28, v23)
        _ = fi28
    }
    var v31 uint64 = (v1 / ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v31
}
