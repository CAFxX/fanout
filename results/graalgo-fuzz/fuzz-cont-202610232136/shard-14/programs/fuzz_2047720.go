package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{65535, 135, 1, 2147483648}
    var v3 int32 = -1
    v4 := []uint32{}
    var v5 int32 = v3
    var v6 uint8 = 227
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v5, v5)
    }
    v5 -= (v3 << ((65535) & 63))
    switch "foo bar" {
    case "hello":
        v7 := GenS1{F0: 2147483647, F1: 885}
        _ = v7
    default:
        v8 := GenS1{F0: 765, F1: 2147483647}
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: (-(fi9)), F1: 828}
        v11 := GenS1{F0: fi9, F1: 4294967295}
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v6, v5, v3)
    switch -499 {
    case 32768:
        for fi12 := 0; fi12 < 17; fi12++ {
            switch -1 {
            case 337:
                v3 -= v3
            }
            switch (fi12 & (fi12 / ((fi12) | 1))) {
            case 2147483647:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            switch (("hello" + "hello") + "") {
            case "a":
                v13 := GenS2{F0: 0, F1: 42}
                _ = v13
            case "hello":
                var v14 bool = false
                _ = v14
            case "foo bar":
                v1[448] = (-(fi12))
            default:
                var v15 float32 = ((3.14159 * 0.1) + 3.14159)
                _ = v15
            }
            _ = fi12
        }
    case 26:
        if false {
            switch (("foo bar" + "a") + ("foo bar" + ((("foo bar" + "a") + ("\n\t" + (("" + "hello") + ("hello" + ("foo bar" + ("foo bar" + (("foo bar" + ("a" + "\n\t")) + ((("" + "a") + ("" + ("foo bar" + "\n\t"))) + ("a" + (("\n\t" + "hello") + "\n\t")))))))))) + "hello"))) {
            case "a":
                var v16 string = "a"
                _ = v16
            default:
                switch -270 {
                case -2147483648:
                    if len(v4) > 0 {
                        v4[0] = 65535
                    }
                }
            }
            v5 = (v5 * 2147483647)
        }
    }
    v3 -= -1
    v3 = v5
    switch "hello" {
    case "":
        for fi17 := 0; fi17 < 5; fi17++ {
            if false {
                v18 := GenS3{F0: uint32(v6)}
                v3 -= (v3 + v5)
                if (42 >= 2) {
                    if (false && true) {
                        var v19 int32 = 2147483647
                        v5 = v5
                        _ = v19
                    } else {
                        v20 := GenS3{F0: 2147483648}
                        v21 := GenS2{F0: 1, F1: v5}
                        v6 += v6
                        _ = v20
                        _ = v21
                    }
                }
                _ = v18
            } else {
                switch (fi17 * fi17) {
                case -298:
                    var v22 uint32 = 4294967295
                    _ = v22
                case -1:
                    if false {
                        v23 := GenS3{F0: 42}
                        if (65535 <= 386) {
                            fmt.Println(v5, v6, v5)
                            var v24 float64 = (1.5 + 1e10)
                            _ = v24
                        }
                        if (1e10 != (0.1 / 3.14159)) {
                            var v25 uint8 = (v6 & v6)
                            _ = v25
                        }
                        _ = v23
                    }
                }
            }
            v26 := GenS1{F0: (fi17 % ((fi17) | 1)), F1: 65535}
            switch "foo bar" {
            case "hello":
                v3 = (^(v5))
            case "\n\t":
                v3 = 0
            }
            _ = fi17
            _ = v26
        }
    case "foo bar":
        v1[22] = 178
    case "hello":
        var v27 int32 = (^(v3))
        _ = v27
    default:
        switch "hello" {
        case "foo bar":
            v28 := GenS3{F0: 65535}
            _ = v28
        case "a":
            if true {
                v1[65535] = 770
                v29 := GenS2{F0: 1, F1: v3}
                if (42 != 1) {
                    v6 = v6
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v29
            }
        case "hello":
            v30 := GenS1{F0: 9223372036854775807, F1: 1}
            _ = v30
        }
    }
    var v31 bool = ("" <= "")
    var v32 uint = 18446744073709551615
    delete(v1, 42)
    var v33 int64 = 0
    v34 := GenS3{F0: 2}
    if (!v31) {
        var v35 bool = (!v31)
        _ = v35
    }
    v1[2147483647] = 32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
