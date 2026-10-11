package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64) int32 {
    p0 = (p0 * p0)
    if false {
        var v4 float32 = 3.14159
        _ = v4
    }
    _ = p0
    return 246
}

func main() {
    var v1 int = -700
    var v2 int8 = -1
    v3 := []string{("a" + ((("\n\t" + "\n\t") + "a") + "")), "hello"}
    fmt.Println(v2, v2)
    switch (-(v1)) {
    case 2147483647:
        if ((v1 * 128) < v1) {
            for fi5 := 0; fi5 < 9; fi5++ {
                var v6 int = (v1 % ((fi5) | 1))
                var v7 float64 = (0.0 + 1.5)
                _ = fi5
                _ = v6
                _ = v7
            }
            var v8 int16 = 522
            var v9 uint32 = 2
            _ = v8
            _ = v9
        }
    case -1:
        if (279 >= 2) {
            for fi10 := 0; fi10 < 15; fi10++ {
                if (886 <= 2) {
                    _ = len(v3)
                }
                v2 = 0
                _ = fi10
            }
        }
    case -9223372036854775808:
        var v11 float32 = float32(float64((1e10 + (0.0 - (3.14159 + (0.0 + 3.14159))))))
        _ = v11
    default:
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 float64 = 0.0
            v13 = 3.14159
            v1 += (fi12 ^ v1)
            _ = fi12
            _ = v13
        }
    }
    var v14 float32 = float32(3.14159)
    v15 := GenS1{F0: true}
    v16 := GenS2{F0: (!true)}
    var v17 int16 = -32768
    v18 := GenS2{F0: true}
    if (-1 >= 0) {
        if true {
            v19 := GenS1{F0: false}
            _ = v19
        } else {
            v20 := GenS2{F0: true}
            if len(v3) > 0 {
                v3[0] = ("a" + "\n\t")
            }
            v17 = 0
            _ = v20
        }
    }
    var v21 string = (((("" + "foo bar") + "a") + "hello") + (("" + "foo bar") + ""))
    if false {
        var v22 int = (v1 - 128)
        if (false && false) {
            if false {
                switch ("\n\t" + (v21 + ("foo bar" + v21))) {
                case "":
                    var v23 uint64 = 18446744073709551615
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 19; fi24++ {
                        if (-786 > 1) {
                            switch 4294967295 {
                            case 0:
                                switch 1 {
                                case 296:
                                    fmt.Println(fi24, v22, fi24)
                                case -2147483648:
                                    for fi25 := 0; fi25 < 4; fi25++ {
                                        if ((!false) || (v17 < 42)) {
                                            var v26 uint32 = 2147483647
                                            v27 := GenS2{F0: (v14 < (((3.14159 + 1e10) / 1e10) * v14))}
                                            _ = v26
                                            _ = v27
                                        }
                                        fmt.Println(v21, fi24, fi24)
                                        _ = fi25
                                    }
                                case -1:
                                    v1 -= (fi24 - fi24)
                                default:
                                    v3 = append(v3, ((v21 + "\n\t") + (v21 + v21)))
                                }
                            default:
                                v22 = (v22 % ((v1) | 1))
                            }
                        }
                        v21 = v21
                        _ = fi24
                    }
                }
                switch v21 {
                case "":
                    var v28 uint16 = 435
                    _ = v28
                default:
                    if (-1 <= -346) {
                        var v29 uint64 = 1
                        _ = v29
                    }
                }
                v22 -= (v1 - v22)
            }
            fmt.Println(v14, v21, v14)
        }
        for fi30 := 0; fi30 < 19; fi30++ {
            _ = len(v3)
            _ = fi30
        }
        _ = v22
    }
    fmt.Println(v21, v1)
    v17 = v17
    if (((v17 >> ((4294967295) & 63)) == 1000) && false) {
        v14 -= v14
        var v31 bool = (!true)
        _ = v31
    } else {
        for fi32 := 0; fi32 < 15; fi32++ {
            v33 := GenS3{F0: 0}
            switch "hello" {
            case "a":
                var v34 uint16 = 255
                _ = v34
            }
            _ = fi32
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
