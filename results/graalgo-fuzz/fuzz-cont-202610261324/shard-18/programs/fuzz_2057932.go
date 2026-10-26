package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{""}
    v3 := make(map[int]uint32)
    v4 := []uint32{42}
    var v5 uint16 = 561
    var v6 uint32 = 267
    switch ("\n\t" + ("a" + "")) {
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        var v7 uint32 = (v6 * (v6 - v6))
        _ = v7
    }
    var v8 float32 = 0.0
    switch ("hello" + "foo bar") {
    case "":
        var v9 int64 = 0
        _ = v9
    case "a":
        var v10 int32 = 42
        _ = v10
    default:
        var v11 float32 = v8
        _ = v11
    }
    v12 := GenS1{F0: v5}
    switch 32768 {
    case 32768:
        switch ("hello" + (("\n\t" + "a") + "a")) {
        case "":
            if ((^(v5)) <= v5) {
                v8 += v8
                var v13 uint32 = v6
                v5 -= v5
                _ = v13
            }
        case "foo bar":
            var v14 uint32 = (v6 ^ v6)
            _ = v14
        case "a":
            var v15 uint = 1
            _ = v15
        default:
            var v16 uint64 = 153
            _ = v16
        }
    }
    v6 = (v6 * 65535)
    fmt.Println(v8)
    if false {
        if ((false || true) && (18446744073709551615 == 2)) {
            v8 = float32(((1e10 / 1.0) + 0.0))
        }
        switch 32767 {
        case 9223372036854775807:
            var v17 int16 = -212
            _ = v17
        default:
            v5 -= v5
        }
    }
    if true {
        v18 := GenS1{F0: 2}
        if false {
            _ = len(v1)
            switch (("" + "foo bar") + "foo bar") {
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "a":
                v8 -= (((v8 - float32(3.14159)) * v8) / ((v8 / (v8 - (v8 * float32(3.14159)))) + (v8 - 0.1)))
            default:
                var v19 int64 = 9223372036854775807
                _ = v19
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v18
    }
    switch "" {
    case "":
        v8 = (-2.25 / v8)
    default:
        if (!(!(!(!("a" == ("a" + "foo bar")))))) {
            delete(v3, 42)
            switch -327 {
            case -448:
                if true {
                    v6 = 497
                }
            case 4294967295:
                var v20 bool = true
                _ = v20
            default:
                for fi21 := 0; fi21 < 19; fi21++ {
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        var v22 float32 = v8
                        _ = v22
                    default:
                        v8 += -2.25
                    }
                    _ = fi21
                }
            }
            v23 := GenS1{F0: (v5 ^ 84)}
            _ = v23
        }
    }
    var v24 int = 42
    v24 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
