package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 string; F1 uint32; F2 int16 }

func main() {
    var v1 uint = 168
    v2 := []uint32{42}
    v3 := make(map[uint32]string)
    var v4 int16 = 42
    v5 := make(map[int]int)
    if (255 < 32767) {
        switch "hello" {
        case "\n\t":
            v6 := GenS1{F0: -32768, F1: v1, F2: 2}
            _ = v6
        case "hello":
            v4 = 32767
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            fmt.Println(v1)
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 int32 = 471
    delete(v5, 2147483647)
    if false {
        if false {
            switch ((("\n\t" + ("\n\t" + "foo bar")) + "") + "a") {
            case "\n\t":
                if ((!((3.14159 + (0.0 * 0.1)) == -2.25)) && false) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch "\n\t" {
                    case "":
                        _ = len(v3)
                    case "hello":
                        v4 -= 0
                    default:
                        var v8 int64 = 720
                        _ = v8
                    }
                }
            }
            v9 := GenS1{F0: -32768, F1: v1, F2: 263}
            _ = v9
        }
        v7 -= (v7 % (((-(v7))) | 1))
        var v10 uint16 = 432
        _ = v10
    }
    fmt.Println(v4, v4)
    for fi11 := 0; fi11 < 13; fi11++ {
        v1 = v1
        if false {
            v4 = (v4 + -902)
            v3[42] = ("hello" + ("" + "\n\t"))
            var v12 float32 = 1e10
            _ = v12
        } else {
            fmt.Println(v4, fi11)
            if true {
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 uint = 681
                    fmt.Println(v7, v14)
                    switch (("a" + "foo bar") + "\n\t") {
                    case "a":
                        v1 -= v1
                    case "\n\t":
                        var v15 float64 = 1e10
                        fmt.Println(v15, fi13)
                        var v16 bool = true
                        _ = v15
                        _ = v16
                    default:
                        if ((v4 & -497) < int16(fi13)) {
                            v7 = -2147483648
                            var v17 uint32 = 0
                            v18 := GenS2{F0: "", F1: v17, F2: v4}
                            if (-41 <= -9223372036854775808) {
                                var v19 bool = (65535 == (v17 / (((v17 / (((v17 & v17)) | 1))) | 1)))
                                v19 = (true && v19)
                                var v20 uint64 = 4294967295
                                _ = v19
                                _ = v20
                            }
                            v21 := GenS1{F0: (-(v4)), F1: 65535, F2: 1}
                            _ = v17
                            _ = v18
                            _ = v21
                        }
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                        v14 -= (v14 << ((v14) & 63))
                    }
                    _ = fi13
                    _ = v14
                }
                v22 := GenS1{F0: (^(v4)), F1: 1, F2: 253}
                v7 = (v7 - v7)
                _ = v22
            }
        }
        var v23 uint = 42
        _ = fi11
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        fmt.Println(v7, v4)
        _ = fi24
    }
    var v25 int8 = 59
    var v26 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
