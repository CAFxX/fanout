package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }
type GenS3 struct { F0 int64; F1 bool; F2 int32 }
type GenS4 struct { F0 int; F1 uint16; F2 int }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -18
    var v3 uint = 586
    var v4 uint = v3
    v5 := []int{32767, -631, 954, 32768}
    var v6 uint8 = 2
    v7 := GenS1{F0: 189}
    v8 := GenS1{F0: 65535}
    if (1000 >= -486) {
        var v9 int64 = int64(v2)
        _ = v9
    }
    v3 = (v4 % ((v3) | 1))
    var v10 uint64 = 1
    v11 := GenS2{F0: false, F1: -1, F2: (v10 - (v10 * 1))}
    if false {
        v12 := GenS3{F0: 9223372036854775807, F1: true, F2: 2147483647}
        switch (((("a" + "") + (("a" + "foo bar") + "\n\t")) + ("a" + "a")) + "foo bar") {
        case "\n\t":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "foo bar":
            v2 = (v2 - (v2 ^ 86))
        default:
            v13 := v1[2]
            _ = v13
        }
        switch "a" {
        case "foo bar":
            if true {
                _ = len(v1)
                if false {
                    fmt.Println(v10)
                    v2 = -1
                }
                var v14 int32 = -1
                _ = v14
            } else {
                var v15 uint16 = 2
                var v16 bool = true
                _ = v15
                _ = v16
            }
        case "":
            v4 -= v3
        }
        _ = v12
    } else {
        if true {
            fmt.Println(v10, v10)
            v4 += uint(v10)
            fmt.Println(v4, v4, v10)
        }
        v17 := GenS1{F0: 71}
        _ = v17
    }
    var v18 int32 = 2147483647
    v6 = 127
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            switch (("a" + "foo bar") + (("" + "hello") + ("foo bar" + "foo bar"))) {
            case "a":
                var v21 int8 = (v2 ^ (v2 ^ v2))
                switch 127 {
                case -874:
                    v4 = (v3 | v3)
                case -308:
                    var v22 float32 = 0.0
                    _ = v22
                case -2147483648:
                    var v23 uint16 = 2
                    _ = v23
                }
                _ = v21
            case "":
                var v24 int32 = 1
                _ = v24
            case "foo bar":
                var v25 int16 = 42
                _ = v25
            default:
                if true {
                    var v26 int32 = -2147483648
                    v27 := GenS4{F0: fi19, F1: 65535, F2: (fi20 << ((972) & 63))}
                    if true {
                        if (!(false || true)) {
                            var v28 string = ""
                            v10 = 18446744073709551615
                            v28 = v28
                            _ = v28
                        }
                        if true {
                            _ = len(v5)
                            _ = len(v5)
                            v29 := GenS2{F0: (32767 > 42), F1: v26, F2: 4294967295}
                            v6 = (^(v6))
                            _ = v29
                        }
                        var v30 uint32 = 2
                        _ = v30
                    } else {
                        var v31 int16 = -704
                        _ = v31
                    }
                    if (-397 == -1) {
                        var v32 bool = (true && ((v6 / (((^(v6))) | 1)) < v6))
                        switch "\n\t" {
                        case "foo bar":
                            var v33 int64 = 1
                            _ = v33
                        }
                        var v34 uint16 = 902
                        _ = v32
                        _ = v34
                    }
                    v35 := v1[uint32(v10)]
                    v36 := GenS5{F0: false, F1: (v6 >= uint8(v3))}
                    switch 127 {
                    case 2147483647:
                        v18 = (-(v26))
                    }
                    _ = v26
                    _ = v27
                    _ = v35
                    _ = v36
                }
            }
            _ = fi20
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
