package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 int8, p1 uint32) string {
    var v5 int16 = -32768
    fmt.Println(p0, p1)
    for fi6 := 0; fi6 < 2; fi6++ {
        v5 = (v5 / (((v5 / (((v5 + (v5 % ((1) | 1)))) | 1))) | 1))
        _ = fi6
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    return "a"
}
func genFn1(p0 int8, p1 string, p2 bool) uint64 {
    var v7 int = 32768
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 89
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = 1000
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    v2 = (v2 - v2)
    v9 := GenS1{F0: 42}
    var v10 string = "hello"
    var v11 string = (((v10 + ("hello" + v10)) + v10) + v10)
    fmt.Println(v10, v10, v1)
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v2 -= v2
            _ = fi13
        }
        _ = fi12
    }
    v14 := []int{-894}
    var v15 bool = (71 <= -27)
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 string = v11
        v10 = (((v11 + v17) + (v11 + v10)) + (v10 + ("a" + "a")))
        if ((true && (!v15)) && (!(!v15))) {
            if (v15 && true) {
                v18 := GenS1{F0: 2}
                for fi19 := 0; fi19 < 18; fi19++ {
                    v20 := GenS2{F0: 0, F1: "hello"}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
        } else {
            if (!(v1 <= 786)) {
                if ((v1 % (((v1 ^ v1)) | 1)) <= v1) {
                    v11 = v17
                }
                fmt.Println(v15, fi16)
            } else {
                fmt.Println(v15, fi16, v10)
                fmt.Println(v2)
                v1 = (v1 + v1)
            }
            switch "" {
            case "a":
                for fi21 := 0; fi21 < 7; fi21++ {
                    if (-1 > 0) {
                        v17 = v11
                        var v22 int = fi21
                        var v23 int64 = v1
                        _ = v22
                        _ = v23
                    }
                    v10 = v17
                    fmt.Println(v11, v17)
                    _ = fi21
                }
            case "\n\t":
                for fi24 := 0; fi24 < 16; fi24++ {
                    v2 -= 1
                    if v15 {
                        v15 = (!((v1 + v1) >= -1))
                        var v25 uint16 = 2
                        var v26 int16 = v2
                        switch 50 {
                        case -484:
                            v27 := GenS1{F0: 0}
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            _ = v27
                        case -2147483648:
                            v28 := GenS1{F0: 948}
                            _ = v28
                        case 32768:
                            _ = len(v3)
                        default:
                            v17 = v17
                        }
                        v1 += -881
                        _ = v25
                        _ = v26
                    } else {
                        var v29 int64 = v1
                        _ = v29
                    }
                    v30 := GenS3{F0: 18446744073709551615, F1: 32768, F2: 1}
                    _ = fi24
                    _ = v30
                }
            case "":
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            v31 := GenS4{F0: 42, F1: v2, F2: 0}
            _ = v31
        }
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
