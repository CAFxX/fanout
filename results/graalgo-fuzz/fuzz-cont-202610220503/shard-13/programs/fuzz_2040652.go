package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int) int {
    p0 = p0
    p0 -= 4294967295
    p0 = (p0 >> ((736) & 63))
    _ = p0
    return 127
}
func genFn1(p0 int8, p1 uint8) uint32 {
    var v6 float64 = 1e10
    v7 := GenS1{F0: -32768, F1: 1000}
    fmt.Println(p0, p0, v6)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 1
    var v3 uint16 = 255
    var v4 uint8 = 42
    var v5 int16 = -32768
    v3 += (v3 / ((uint16(v4)) | 1))
    fmt.Println(v4, v5)
    delete(v1, "")
    fmt.Println(v4)
    v5 -= v5
    var v8 int32 = 662
    if true {
        v5 = v5
        if true {
            for fi9 := 0; fi9 < 19; fi9++ {
                v5 += (^(v5))
                for fi10 := 0; fi10 < 15; fi10++ {
                    v2 += (v2 - v2)
                    if false {
                        switch (fi10 | (fi9 ^ genFn0((fi9 - genFn0((fi9 ^ fi10)))))) {
                        case 9223372036854775807:
                            if false {
                                var v11 uint64 = 18446744073709551615
                                _ = v11
                            } else {
                                var v12 int8 = -1
                                var v13 int16 = -721
                                _ = v12
                                _ = v13
                            }
                        case -431:
                            v14 := GenS1{F0: v5, F1: v5}
                            _ = v14
                        case 59:
                            fmt.Println(fi9, fi10, fi9)
                        default:
                            v4 -= 1
                        }
                        v4 -= 127
                    }
                    fmt.Println(v4, v8, v5)
                    _ = fi10
                }
                _ = fi9
            }
            for fi15 := 0; fi15 < 2; fi15++ {
                for fi16 := 0; fi16 < 5; fi16++ {
                    if ((("\n\t" + "\n\t") >= "hello") && (!true)) {
                        v17 := GenS2{F0: (v3 >> ((65535) & 63))}
                        v18 := GenS3{F0: v3, F1: genFn1(121, v4), F2: -1}
                        switch fi15 {
                        case 2147483647:
                            switch "\n\t" {
                            case "a":
                                if (!(((!true) || (false && (false || (false && false)))) && true)) {
                                    if true {
                                        fmt.Println(fi15, fi16, v4)
                                    } else {
                                        if true {
                                            var v19 float64 = -2.25
                                            _ = v19
                                        }
                                    }
                                }
                            case "":
                                v8 += 0
                            default:
                                var v20 int = 1
                                _ = v20
                            }
                        default:
                            if ((true || true) && true) {
                                var v21 int8 = -1
                                if true {
                                    var v22 uint64 = 2
                                    _ = v22
                                }
                                if (v21 <= (v21 << ((569) & 63))) {
                                    var v23 uint8 = v4
                                    _ = v23
                                }
                                _ = v21
                            }
                            v24 := GenS4{F0: 978}
                            _ = v24
                        }
                        _ = v17
                        _ = v18
                    } else {
                        var v25 uint16 = (v3 >> ((796) & 63))
                        fmt.Println(v25, v2, v4)
                        var v26 uint64 = 4294967295
                        var v27 int8 = -112
                        v28 := GenS3{F0: v25, F1: 2147483648, F2: (v27 & v27)}
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    var v29 uint64 = 192
                    var v30 float32 = 0.0
                    _ = fi16
                    _ = v29
                    _ = v30
                }
                v1["a"] = (("hello" + "") + "foo bar")
                _ = fi15
            }
        } else {
            for fi31 := 0; fi31 < 3; fi31++ {
                switch ("hello" + "foo bar") {
                case "a":
                    fmt.Println(v4)
                case "":
                    v5 = (v5 >> ((42) & 63))
                case "\n\t":
                    v32 := GenS3{F0: (v3 << ((625) & 63)), F1: 0, F2: -120}
                    _ = v32
                }
                v33 := []int{fi31, (-(fi31))}
                _ = fi31
                _ = v33
            }
        }
    }
    var v34 int16 = 1
    fmt.Println(v2)
    if false {
        v35 := GenS4{F0: 18446744073709551615}
        v34 += v34
        _ = v35
    } else {
        v5 += (v5 * int16(v2))
        v34 = (v34 >> ((255) & 63))
        var v36 uint32 = 1
        _ = v36
    }
    fmt.Println(v34)
    v2 = (v2 + (v2 ^ (v2 % ((v2) | 1))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v8
}
