package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483647
    v3 := make(map[int]string)
    v4 := []string{"hello"}
    v2 -= 2
    var v5 uint64 = 4294967295
    var v6 int8 = -127
    if (true && (v6 <= (v6 + (v6 ^ int8(v5))))) {
        fmt.Println(v2)
        if true {
            v6 += (v6 - v6)
            var v7 float32 = 1.5
            _ = v7
        } else {
            var v8 uint8 = 1
            _ = v8
        }
        v9 := GenS1{F0: v2, F1: "hello", F2: 65535}
        _ = v9
    } else {
        v2 = (^(v2))
        v5 += v5
        v10 := GenS1{F0: v2, F1: (("" + ("a" + ((("hello" + "\n\t") + "\n\t") + "\n\t"))) + "hello"), F2: v2}
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if false {
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            v12 := GenS1{F0: v2, F1: ("" + ("\n\t" + "hello")), F2: (v2 >> ((4294967295) & 63))}
            _ = v12
        } else {
            v2 = (v2 / ((695) | 1))
            v5 -= 18446744073709551615
            switch "a" {
            case "a":
                switch (^(fi11)) {
                case 937:
                    if true {
                        if (!(3.14159 == 3.14159)) {
                            for fi13 := 0; fi13 < 10; fi13++ {
                                v14 := GenS1{F0: (v2 + v2), F1: "a", F2: v2}
                                switch "\n\t" {
                                case "\n\t":
                                    var v15 uint64 = v5
                                    _ = v15
                                case "":
                                    var v16 float32 = (0.0 * 1.5)
                                    _ = v16
                                case "hello":
                                    v17 := GenS1{F0: 134, F1: ("foo bar" + "a"), F2: (v2 ^ v2)}
                                    _ = v17
                                }
                                _ = fi13
                                _ = v14
                            }
                            v18 := GenS2{F0: 2, F1: (-(fi11)), F2: 1}
                            var v19 uint32 = 2147483647
                            _ = v18
                            _ = v19
                        }
                    }
                default:
                    if (1 < (v5 << ((255) & 63))) {
                        var v20 uint16 = 255
                        _ = v20
                    }
                }
            case "":
                switch "" {
                case "":
                    for fi21 := 0; fi21 < 12; fi21++ {
                        if (!true) {
                            v22 := GenS3{F0: -2147483648}
                            v23 := GenS1{F0: (v2 & 1), F1: (("" + ("\n\t" + "foo bar")) + "hello"), F2: 2147483647}
                            _ = v22
                            _ = v23
                        }
                        v6 = 0
                        _ = fi21
                    }
                }
            }
        }
        var v24 uint64 = 737
        v25 := GenS4{F0: 642}
        _ = fi11
        _ = v24
        _ = v25
    }
    if false {
        if false {
            var v26 uint64 = uint64(v2)
            v26 = 149
            _ = v26
        } else {
            v2 += (v2 | v2)
            var v27 int8 = (v6 + (v6 % ((v6) | 1)))
            for fi28 := 0; fi28 < 10; fi28++ {
                v6 = 21
                v27 -= 127
                _ = fi28
            }
            _ = v27
        }
        v2 = v2
    } else {
        if ((("\n\t" + "\n\t") + "hello") != "a") {
            fmt.Println(v5)
            fmt.Println(v5, v6)
            v1[(v2 | uint32(v5))] = false
        }
    }
    v5 = (v5 % ((v5) | 1))
    var v29 int = 32768
    var v30 string = (("\n\t" + ("hello" + "")) + ("a" + ("foo bar" + "\n\t")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
