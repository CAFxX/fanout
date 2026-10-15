package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint }
func genFn0(p0 uint32, p1 int32, p2 int16) uint64 {
    p2 += p2
    var v6 uint64 = 385
    if ("\n\t" >= "foo bar") {
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: -1, F1: 473}
                _ = fi8
                _ = v9
            }
            p1 += (p1 + int32(v6))
            switch (-(fi7)) {
            case -2147483648:
                v10 := GenS2{F0: 42, F1: (v6 * v6)}
                _ = v10
            default:
                switch (fi7 * fi7) {
                case 9223372036854775807:
                    var v11 uint16 = 1
                    _ = v11
                case -201:
                    var v12 bool = ((0.1 / 1.5) == -2.25)
                    _ = v12
                case -2147483648:
                    var v13 uint64 = (v6 - v6)
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 16; fi14++ {
                        switch ("foo bar" + "") {
                        case "\n\t":
                            p2 += 1000
                            var v15 int8 = -128
                            var v16 float32 = 1e10
                            v15 -= 127
                            _ = v15
                            _ = v16
                        case "a":
                            v6 = (v6 & (v6 << ((873) & 63)))
                        case "":
                            p2 = (p2 / ((p2) | 1))
                        }
                        _ = fi14
                    }
                }
            }
            _ = fi7
        }
    } else {
        v17 := GenS3{F0: 42, F1: 543}
        var v18 float64 = 1.5
        _ = v17
        _ = v18
    }
    fmt.Println(p1, v6, p1)
    v19 := GenS1{F0: 1, F1: 186}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v6
    return 344
}

func main() {
    var v1 float32 = 0.1
    var v2 float64 = 3.14159
    var v3 int = 885
    var v4 int32 = -1
    var v5 uint16 = 255
    var v20 int64 = 1
    v5 = (v5 & (v5 / ((v5) | 1)))
    for fi21 := 0; fi21 < 4; fi21++ {
        v1 -= 0.0
        _ = fi21
    }
    var v22 string = (("foo bar" + "a") + ("\n\t" + ((("\n\t" + ("hello" + "hello")) + "a") + (("foo bar" + (("foo bar" + "") + ("hello" + "a"))) + ("foo bar" + "a")))))
    if (0.0 != (0.0 * ((float64(v1) / (((v2 / v2) * 0.0) / v2)) + 0.1))) {
        v23 := GenS3{F0: 100, F1: 42}
        fmt.Println(v2)
        v1 = (v1 * v1)
        _ = v23
    }
    switch int(v20) {
    case 1:
        for fi24 := 0; fi24 < 2; fi24++ {
            if false {
                v25 := GenS2{F0: v20, F1: 42}
                v22 = v22
                _ = v25
            } else {
                v26 := []int{fi24, (fi24 % ((v3) | 1)), (v3 ^ (fi24 - (fi24 % ((fi24) | 1)))), fi24}
                _ = v26
            }
            _ = fi24
        }
    case 189:
        var v27 int16 = 0
        _ = v27
    case 2147483647:
        var v28 string = v22
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 int64 = 203
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    v3 = v3
    var v32 int8 = 1
    v33 := []int{}
    var v34 uint16 = (^(v5))
    for fi35 := 0; fi35 < 15; fi35++ {
        fmt.Println(v4, v2, v1)
        _ = fi35
    }
    v4 -= (v4 >> ((4294967295) & 63))
    v2 -= v2
    v36 := GenS4{F0: -1, F1: 255, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
}
