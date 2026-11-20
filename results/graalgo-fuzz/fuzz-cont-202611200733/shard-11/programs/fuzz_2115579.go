package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 809
    var v3 int32 = -2147483648
    var v4 uint = 552
    v5 := []int{}
    switch "\n\t" {
    case "foo bar":
        if false {
            var v6 uint32 = v2
            for fi7 := 0; fi7 < 7; fi7++ {
                v2 += v6
                for fi8 := 0; fi8 < 19; fi8++ {
                    v3 = v3
                    fmt.Println(fi7)
                    _ = fi8
                }
                for fi9 := 0; fi9 < 10; fi9++ {
                    v4 = 677
                    _ = fi9
                }
                _ = fi7
            }
            _ = v6
        }
    default:
        fmt.Println(v3)
    }
    if true {
        v3 = 42
        fmt.Println(v2)
        if false {
            delete(v1, (v2 ^ (^(v2))))
            var v10 uint = (v4 % ((uint(v2)) | 1))
            for fi11 := 0; fi11 < 9; fi11++ {
                for fi12 := 0; fi12 < 7; fi12++ {
                    if false {
                        var v13 int64 = 390
                        var v14 int16 = -32768
                        switch "\n\t" {
                        case "":
                            v4 -= v4
                        case "hello":
                            v10 -= (v4 % ((2) | 1))
                        case "\n\t":
                            _ = len(v1)
                        default:
                            var v15 int = (fi12 * (-(fi12)))
                            _ = v15
                        }
                        _ = v13
                        _ = v14
                    }
                    delete(v1, v2)
                    fmt.Println(fi11)
                    if (-570 < 0) {
                        switch ("a" + "hello") {
                        case "\n\t":
                            v2 -= v2
                        default:
                            v4 = 2
                        }
                        var v16 int64 = 409
                        _ = v16
                    } else {
                        _ = len(v1)
                        v4 += v10
                        switch 127 {
                        case -2147483648:
                            var v17 float32 = 1e10
                            _ = v17
                        case 9223372036854775807:
                            v18 := GenS1{F0: 19}
                            _ = v18
                        case 558:
                            var v19 int32 = v3
                            _ = v19
                        }
                    }
                    _ = fi12
                }
                var v20 bool = true
                var v21 uint32 = (v2 / ((v2) | 1))
                _ = fi11
                _ = v20
                _ = v21
            }
            _ = v10
        }
    }
    v1[v2] = "a"
    var v22 bool = (!true)
    v2 -= (v2 | (v2 & v2))
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v4, fi23, fi23)
        _ = fi23
    }
    v3 = (v3 >> (((v4 % ((v4) | 1))) & 63))
    v2 = (v2 ^ 732)
    var v24 int = 265
    var v25 string = ("hello" + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
