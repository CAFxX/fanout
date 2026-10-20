package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 bool = (1.5 >= 1.5)
    v5 := []string{((((((("hello" + ("a" + ("" + ("a" + "\n\t")))) + (("" + ("\n\t" + ("\n\t" + "foo bar"))) + ((("foo bar" + ("hello" + ("a" + ("foo bar" + ("foo bar" + ""))))) + "") + ((("a" + "a") + "\n\t") + "hello")))) + (("foo bar" + "foo bar") + "a")) + (("a" + "hello") + "")) + (("" + "\n\t") + ((("hello" + "foo bar") + "") + "a"))) + (((("hello" + ("hello" + ((("a" + ("a" + "foo bar")) + "") + ""))) + (("hello" + ((("hello" + "foo bar") + "a") + ("foo bar" + "hello"))) + "\n\t")) + "\n\t") + "foo bar")) + "a"), "a", "\n\t", "hello"}
    var v6 float32 = 0.0
    v4 = (0 != 9223372036854775807)
    v4 = false
    fmt.Println(v4)
    if v4 {
        var v7 uint64 = 0
        _ = v7
    }
    if v4 {
        fmt.Println(v6, v4)
        switch -1 {
        case 1:
            var v8 int = 32768
            _ = v8
        case 4294967295:
            v9 := GenS1{F0: 4294967295, F1: 81, F2: 2}
            _ = v9
        case 32768:
            if v4 {
                if (!(!((42 == 127) && (32767 <= 1)))) {
                    v10 := v2["foo bar"]
                    v4 = (1 >= 344)
                    switch "\n\t" {
                    case "hello":
                        v11 := GenS2{F0: -9223372036854775808, F1: 73, F2: 169}
                        _ = v11
                    }
                    _ = v10
                }
                for fi12 := 0; fi12 < 11; fi12++ {
                    if false {
                        var v13 int8 = 42
                        _ = v13
                    }
                    v6 -= v6
                    fmt.Println(fi12, v6)
                    _ = fi12
                }
                v4 = ((-2147483648 <= 42) && (v4 || (((v6 - (float32(1e10) / ((v6 / (v6 + v6)) - (v6 - 0.0)))) + (-2.25 / ((v6 * float32(1e10)) / (v6 - float32((1.5 / 1.5)))))) != v6)))
            } else {
                v6 = ((0.1 + (3.14159 / 0.1)) / (v6 - v6))
                if (!v4) {
                    var v14 int64 = 1
                    if (!v4) {
                        fmt.Println(v4, v14)
                        switch -9223372036854775808 {
                        case 128:
                            fmt.Println(v6, v6, v14)
                        default:
                            switch "hello" {
                            case "foo bar":
                                v14 -= (v14 >> ((413) & 63))
                            case "hello":
                                var v15 float32 = v6
                                _ = v15
                            case "":
                                for fi16 := 0; fi16 < 12; fi16++ {
                                    v17 := GenS2{F0: 801, F1: 128, F2: 342}
                                    _ = fi16
                                    _ = v17
                                }
                            default:
                                var v18 uint = 18446744073709551615
                                _ = v18
                            }
                        }
                        for fi19 := 0; fi19 < 18; fi19++ {
                            v20 := GenS3{F0: 0}
                            v4 = (v4 && (128 <= 246))
                            var v21 uint32 = 863
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = v14
                }
            }
        default:
            v6 = v6
        }
    } else {
        v4 = (!(-661 > 0))
    }
    v22 := GenS4{F0: 42}
    switch 1 {
    case -209:
        _ = len(v3)
    default:
        v23 := GenS4{F0: 2147483647}
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
