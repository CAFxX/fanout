package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 107
    var v3 float32 = 1e10
    var v4 uint64 = 42
    var v5 int16 = 32767
    var v6 uint8 = (v2 << ((uint(v4)) & 63))
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint16 = 255
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (^(v2)), F1: -1, F2: v5}
    var v10 int32 = 729
    var v11 bool = (-2147483648 <= 0)
    var v12 uint8 = 219
    v12 = v6
    v10 += (v10 & (v10 ^ (v10 - 1)))
    switch 342 {
    case 32767:
        v2 -= (v2 * v12)
    default:
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    var v14 uint8 = v12
    for fi15 := 0; fi15 < 12; fi15++ {
        if (v11 || (v4 > (v4 & 605))) {
            var v16 uint8 = (v12 % ((v12) | 1))
            fmt.Println(fi15, v16, v16)
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := v1[fi17]
                var v19 float32 = -2.25
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v16
        } else {
            var v20 uint32 = 2147483648
            if v11 {
                for fi21 := 0; fi21 < 17; fi21++ {
                    v11 = ((v10 & -1) <= 1)
                    v12 = v12
                    v22 := GenS1{F0: (v2 & (^(v14))), F1: 42, F2: 1}
                    _ = fi21
                    _ = v22
                }
                switch fi15 {
                case -741:
                    var v23 uint64 = 2
                    _ = v23
                case 758:
                    fmt.Println(v2)
                }
                switch ("hello" + "hello") {
                case "\n\t":
                    fmt.Println(v20, fi15)
                }
            } else {
                v24 := GenS1{F0: v6, F1: -9223372036854775808, F2: (^(v5))}
                if (!(!v11)) {
                    v12 = v2
                    var v25 uint8 = v12
                    if (447 == 413) {
                        v2 = 2
                    }
                    _ = v25
                }
                _ = v24
            }
            _ = v20
        }
        var v26 uint = 0
        _ = fi15
        _ = v26
    }
    v27 := GenS2{F0: -636, F1: 9223372036854775807}
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 14; fi29++ {
            if v11 {
                var v30 uint8 = v6
                v31 := []int{-9223372036854775808, (fi28 >> ((607) & 63)), (^(fi28)), fi28}
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
                if (126 < 127) {
                    if (("hello" <= "") && false) {
                        if v11 {
                            v2 = 2
                            var v32 string = "a"
                            v12 = 2
                            v6 -= v2
                            fmt.Println(v4)
                            v33 := GenS3{F0: 9223372036854775807}
                            _ = v32
                            _ = v33
                        }
                    } else {
                        var v34 uint16 = 818
                        var v35 uint16 = uint16(v30)
                        _ = v34
                        _ = v35
                    }
                } else {
                    fmt.Println(v30, v3)
                    var v36 bool = v11
                    if (-1 < 100) {
                        if v11 {
                            fmt.Println(fi28, v12, v4)
                        }
                    }
                    v6 = (v14 ^ v6)
                    if (42 < 2) {
                        var v37 int = (fi28 + fi29)
                        var v38 int16 = v5
                        _ = v37
                        _ = v38
                    }
                    _ = v36
                }
                _ = v30
                _ = v31
            } else {
                fmt.Println(v11)
            }
            v2 = v14
            var v39 int32 = (v10 ^ v10)
            _ = fi29
            _ = v39
        }
        v40 := []int{-2147483648}
        for fi41 := 0; fi41 < 1; fi41++ {
            v42 := GenS2{F0: 42, F1: -9223372036854775808}
            v43 := GenS3{F0: -273}
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi28
        _ = v40
    }
    var v44 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v9
}
