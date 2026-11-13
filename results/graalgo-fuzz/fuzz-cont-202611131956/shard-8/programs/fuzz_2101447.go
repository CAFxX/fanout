package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    v3 := []string{"\n\t"}
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    var v6 uint8 = 255
    _ = len(v4)
    switch "\n\t" {
    case "a":
        v6 -= v6
    case "\n\t":
        if true {
            fmt.Println(v6)
        } else {
            var v7 int8 = -32
            v7 = v7
            v6 += (v6 / ((v6) | 1))
            _ = v7
        }
    default:
        var v8 float64 = 0.1
        _ = v8
    }
    if (!((!(!(!true))) && (42 >= 42))) {
        if ((!(3.14159 >= (1.5 * 1e10))) && (true && true)) {
            if (!(((("a" + "") + ((("\n\t" + "\n\t") + "a") + "foo bar")) != "") && (-1 != 1))) {
                if false {
                    for fi9 := 0; fi9 < 19; fi9++ {
                        for fi10 := 0; fi10 < 5; fi10++ {
                            v6 += (^(v6))
                            _ = fi10
                        }
                        if true {
                            v11 := GenS1{F0: false}
                            var v12 bool = ((!true) && (!false))
                            _ = v11
                            _ = v12
                        } else {
                            v13 := GenS2{F0: -128, F1: 2147483648, F2: 1}
                            _ = v13
                        }
                        var v14 uint16 = 65535
                        _ = fi9
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    fmt.Println(v6, fi15)
                    var v16 uint = 18446744073709551615
                    _ = fi15
                    _ = v16
                }
                var v17 string = "foo bar"
                _ = v17
            } else {
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(fi18)
                    v6 += v6
                    _ = len(v4)
                    _ = fi18
                }
            }
            switch "\n\t" {
            case "foo bar":
                if (876 > 32767) {
                    fmt.Println(v6)
                } else {
                    if true {
                        fmt.Println(v6)
                    } else {
                        var v19 int32 = -1
                        v6 = v6
                        _ = v19
                    }
                    var v20 uint64 = 2
                    _ = v20
                }
            case "a":
                var v21 int = -2147483648
                _ = v21
            default:
                v22 := GenS3{F0: 4294967295, F1: 1}
                _ = v22
            }
        }
        if (1 != 32767) {
            var v23 uint32 = 1
            v6 = (^(v6))
            _ = v23
        }
        switch 32767 {
        case 32768:
            v24 := GenS3{F0: 255, F1: 581}
            _ = v24
        case 0:
            for fi25 := 0; fi25 < 13; fi25++ {
                if false {
                    fmt.Println(v6, v6)
                    fmt.Println(v6)
                    v2[2147483647] = 0
                }
                _ = fi25
            }
        case 9223372036854775807:
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 int = (fi26 * fi26)
                _ = fi26
                _ = v27
            }
        }
    } else {
        v6 = 1
        var v28 float64 = 0.0
        _ = v28
    }
    v6 = v6
    v29 := GenS4{F0: 42}
    v30 := GenS3{F0: 255, F1: -1}
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
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
