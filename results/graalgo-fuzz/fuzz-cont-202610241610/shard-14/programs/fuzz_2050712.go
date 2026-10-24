package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int16; F1 int16; F2 uint16 }
func genFn0(p0 int16) uint8 {
    var v7 uint32 = 2147483647
    v7 -= (v7 >> ((0) & 63))
    _ = p0
    _ = v7
    return 127
}

func main() {
    var v1 int = 0
    v2 := []string{"", "foo bar", ((("a" + "") + ("a" + (("a" + ("" + "")) + "a"))) + "")}
    var v3 uint = 4294967295
    var v4 uint64 = 1
    var v5 int64 = 9223372036854775807
    v6 := []uint32{}
    v4 += (v4 % (((v4 + (v4 * 4294967295))) | 1))
    switch (v1 - (-(v1))) {
    case -257:
        if ((42 > 1) || true) {
            var v8 int8 = 80
            v3 = (^(v3))
            var v9 uint8 = 42
            _ = v8
            _ = v9
        }
    case -473:
        if (2 > 110) {
            var v10 uint8 = genFn0(141)
            _ = v10
        } else {
            if false {
                var v11 int64 = v5
                var v12 uint8 = 0
                _ = v11
                _ = v12
            } else {
                var v13 int = v1
                _ = v13
            }
        }
    }
    switch 128 {
    case 9223372036854775807:
        v14 := GenS1{F0: (v5 * (v5 - v5))}
        _ = v14
    case 128:
        var v15 uint8 = 77
        _ = v15
    case 32767:
        if len(v6) > 0 {
            v6[0] = 819
        }
    default:
        if true {
            if (-339 >= 73) {
                v4 = v4
            } else {
                for fi16 := 0; fi16 < 10; fi16++ {
                    v17 := GenS1{F0: 9223372036854775807}
                    _ = fi16
                    _ = v17
                }
                v4 = v4
            }
            var v18 string = ("\n\t" + "foo bar")
            switch (v1 ^ (v1 & v1)) {
            case 9223372036854775807:
                if (-32768 <= 1) {
                    var v19 int64 = v5
                    _ = v19
                }
            }
            _ = v18
        } else {
            fmt.Println(v5)
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(v5, v1, v4)
                if len(v2) > 0 {
                    v2[0] = ""
                }
                _ = fi20
            }
            fmt.Println(v3, v3, v4)
        }
    }
    if true {
        v4 = 1
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 uint64 = 18446744073709551615
            var v23 int32 = 2147483647
            var v24 int = v1
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        var v25 uint8 = 0
        _ = v25
    }
    switch "\n\t" {
    case "\n\t":
        var v26 int8 = 42
        _ = v26
    case "hello":
        if (((0.1 + (((1e10 / 1.0) - 0.1) - 0.0)) - 0.1) != (3.14159 - (((0.0 / (0.1 * -2.25)) / (0.0 + 1e10)) + -2.25))) {
            var v27 int64 = (v5 * (v5 / (((v5 % ((v5) | 1))) | 1)))
            v6 = append(v6, 2147483647)
            switch v1 {
            case 32768:
                var v28 int64 = (v27 << (((v3 + 0)) & 63))
                _ = v28
            case -942:
                if len(v6) > 0 {
                    v6[0] = 624
                }
            }
            _ = v27
        }
    default:
        v29 := GenS1{F0: (v5 & (^(v5)))}
        _ = v29
    }
    var v30 int32 = -1
    _ = len(v2)
    v31 := GenS2{F0: 255}
    if false {
        v32 := GenS3{F0: 1, F1: 32767}
        fmt.Println(v3)
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS1{F0: (v5 - v5)}
        switch (^(fi33)) {
        case 2147483647:
            for fi35 := 0; fi35 < 5; fi35++ {
                v36 := GenS3{F0: 255, F1: 0}
                v37 := GenS4{F0: 0, F1: -851, F2: 775}
                v2 = append(v2, ("foo bar" + "hello"))
                _ = fi35
                _ = v36
                _ = v37
            }
        case 128:
            switch ("a" + "a") {
            case "a":
                v4 += 370
            }
        }
        if true {
            var v38 int64 = 1
            _ = v38
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
