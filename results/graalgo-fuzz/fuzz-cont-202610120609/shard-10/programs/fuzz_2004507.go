package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
func genFn0() int32 {
    var v7 float32 = (1.5 * 0.0)
    v8 := GenS1{F0: (32767 < 1), F1: -2}
    var v9 int32 = 327
    _ = v7
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int16 = 578
    var v3 int64 = 76
    var v4 int32 = 42
    var v5 string = "hello"
    v6 := make(map[string]string)
    fmt.Println(v1, v5, v3)
    for fi10 := 0; fi10 < 12; fi10++ {
        if false {
            for fi11 := 0; fi11 < 13; fi11++ {
                v3 = -1
                var v12 float32 = 0.1
                var v13 bool = false
                if ((42 > 18446744073709551615) && false) {
                    var v14 uint32 = 268
                    _ = v14
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            var v15 int8 = 1
            v3 = v3
            _ = v15
        } else {
            switch "foo bar" {
            case "hello":
                var v16 int64 = v3
                _ = v16
            case "\n\t":
                if ((v4 + 1) <= (v4 + (v4 / (((v4 & v4)) | 1)))) {
                    switch ("a" + "hello") {
                    case "\n\t":
                        for fi17 := 0; fi17 < 11; fi17++ {
                            var v18 uint16 = 281
                            v19 := v6[((v5 + (v5 + (v5 + v5))) + "a")]
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    case "foo bar":
                        v2 = (v2 % (((v2 ^ v2)) | 1))
                    case "a":
                        var v20 int16 = (-(v2))
                        _ = v20
                    }
                    var v21 int32 = -439
                    _ = v21
                }
            case "foo bar":
                v1 -= (v1 >> (((v1 & (v1 % ((1) | 1)))) & 63))
            default:
                v1 = (v1 * 18446744073709551615)
            }
            v4 = v4
        }
        _ = fi10
    }
    if (0 < -131) {
        if (false && false) {
            if (true && true) {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 int = 32767
                    _ = fi22
                    _ = v23
                }
            } else {
                var v24 int16 = 42
                _ = v24
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                if false {
                    v5 = v5
                    v26 := GenS2{F0: 0, F1: (v4 ^ (v4 - genFn0())), F2: (v3 << ((v1) & 63))}
                    _ = v26
                } else {
                    v5 = (v5 + ((v5 + v5) + "foo bar"))
                    v3 += 42
                }
                _ = fi25
            }
        } else {
            var v27 int = -547
            var v28 uint64 = 983
            v4 -= -1
            _ = v27
            _ = v28
        }
        var v29 float32 = -2.25
        v29 += (float32((0.0 * 1.5)) - v29)
        _ = v29
    }
    fmt.Println(v2, v1, v3)
    v30 := GenS3{F0: 1}
    fmt.Println(v1, v2)
    if true {
        v31 := GenS4{F0: 40}
        fmt.Println(v1)
        _ = v31
    } else {
        v32 := GenS1{F0: (!((float64(1.5) < float64((1.5 / 1e10))) && (false && (!true)))), F1: 0}
        _ = v32
    }
    var v33 uint8 = 1
    v34 := []int{128}
    var v35 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
