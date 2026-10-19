package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string; F1 int; F2 int64 }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 56
    v3 := []uint32{27, 0}
    var v4 uint64 = 1
    var v5 bool = ((v2 / ((v2) | 1)) < v2)
    v2 = 566
    fmt.Println(v4)
    var v6 string = "\n\t"
    v4 += v4
    v4 = v4
    for fi7 := 0; fi7 < 1; fi7++ {
        v6 = ("\n\t" + v6)
        var v8 int64 = -1
        var v9 float64 = (3.14159 / -2.25)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = (v4 - uint64(v2))
    switch 4294967295 {
    case -1:
        for fi11 := 0; fi11 < 3; fi11++ {
            if (!v5) {
                switch v6 {
                case "hello":
                    if (!(v5 && false)) {
                        delete(v1, v6)
                        _ = len(v3)
                    }
                case "foo bar":
                    for fi12 := 0; fi12 < 6; fi12++ {
                        if (-1 != -1) {
                            var v13 int32 = -2147483648
                            _ = v13
                        }
                        fmt.Println(fi11)
                        fmt.Println(v10, v10)
                        _ = fi12
                    }
                case "a":
                    v4 = 4294967295
                default:
                    v14 := GenS1{F0: 1}
                    _ = v14
                }
                if (-9223372036854775808 <= (fi11 ^ 0)) {
                    fmt.Println(fi11, v2)
                } else {
                    v15 := GenS1{F0: 1}
                    switch "foo bar" {
                    case "":
                        if ((v4 >> ((914) & 63)) != v10) {
                            _ = len(v3)
                        }
                    default:
                        v16 := GenS1{F0: -128}
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 15; fi17++ {
                        if (100 == -128) {
                            if v5 {
                                v18 := GenS2{F0: uint32(v4), F1: 127}
                                var v19 float32 = 1.5
                                _ = v18
                                _ = v19
                            }
                            v20 := GenS3{F0: v6, F1: (fi11 ^ -1), F2: 1}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v2, v6, fi17)
                            _ = v20
                        }
                        v21 := GenS3{F0: (("\n\t" + v6) + (("a" + v6) + "")), F1: (fi17 / (((fi17 - fi17)) | 1)), F2: 538}
                        v2 = (^(v2))
                        _ = fi17
                        _ = v21
                    }
                    _ = v15
                }
            }
            v5 = v5
            _ = fi11
        }
    case 127:
        v22 := GenS3{F0: "\n\t", F1: -9223372036854775808, F2: -252}
        _ = v22
    case -2147483648:
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS4{F0: 182, F1: fi23}
            v10 += v4
            v4 += (v4 ^ 1)
            _ = fi23
            _ = v24
        }
    default:
        if v5 {
            if ((v5 || (v5 || (v5 && (3.14159 <= (float32(0.1) + (3.14159 * 1e10)))))) || (!(!v5))) {
                v25 := GenS3{F0: (v6 + "a"), F1: -1, F2: 1}
                v6 = v6
                _ = v25
            }
        } else {
            v4 += v4
            var v26 uint8 = 255
            _ = v26
        }
    }
    v2 = v2
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
