package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 string }
type GenS3 struct { F0 string; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint }

func main() {
    var v1 uint16 = 1
    var v2 bool = true
    var v3 float32 = 0.1
    v4 := []int{1, -707, 351}
    var v5 float32 = v3
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 bool = (1 != 385)
        if v2 {
            fmt.Println(v2, v2, v1)
        } else {
            v4 = append(v4, fi6)
        }
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v5 += v3
        _ = fi8
    }
    v1 = v1
    v9 := GenS1{F0: -1}
    v10 := GenS1{F0: 381}
    v4 = append(v4, -1)
    v3 -= (v5 + 1e10)
    var v11 float32 = (1e10 * v5)
    v1 -= 973
    switch -596 {
    case 128:
        for fi12 := 0; fi12 < 6; fi12++ {
            v11 -= v3
            v13 := GenS2{F0: 2, F1: v1, F2: ((("" + "hello") + ("foo bar" + "\n\t")) + "foo bar")}
            if v2 {
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 bool = false
                    v1 = v1
                    if (v2 && ((v1 ^ (^(v1))) >= v1)) {
                        v3 -= (v5 * v3)
                        v16 := GenS1{F0: 42}
                        if (0 <= 32767) {
                            fmt.Println(v1, fi12, v15)
                            _ = len(v4)
                            if v2 {
                                if (!(true || ("\n\t" >= ""))) {
                                    _ = len(v4)
                                    _ = len(v4)
                                }
                                if v15 {
                                    v5 = float32(0.1)
                                    fmt.Println(v2, v15, v11)
                                }
                                var v17 uint64 = 18446744073709551615
                                _ = v17
                            }
                        } else {
                            var v18 uint64 = 2
                            _ = v18
                        }
                        _ = v16
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                fmt.Println(v1, v11)
                v19 := GenS2{F0: 0, F1: v1, F2: (("hello" + "foo bar") + "")}
                v5 += (((v5 * float32(1e10)) / 1.0) + v11)
                _ = v19
            }
            _ = fi12
            _ = v13
        }
    case 2147483647:
        v11 += v11
    case -689:
        v20 := GenS3{F0: ("foo bar" + "hello"), F1: 220, F2: 42}
        _ = v20
    }
    fmt.Println(v11)
    if (!(v2 || false)) {
        switch "a" {
        case "":
            var v21 uint16 = 42
            _ = v21
        case "foo bar":
            v22 := GenS2{F0: 2, F1: (v1 | 255), F2: "a"}
            _ = v22
        default:
            v23 := GenS4{F0: (v1 >> ((784) & 63)), F1: 65535}
            _ = v23
        }
    }
    v11 -= v5
    _ = len(v4)
    var v24 float32 = 0.1
    var v25 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
