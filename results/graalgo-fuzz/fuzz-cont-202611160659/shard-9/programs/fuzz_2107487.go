package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 int16; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint8; F1 string; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
func genFn0(p0 string) uint {
    fmt.Println(p0)
    switch (p0 + p0) {
    case "":
        if true {
            v7 := GenS1{F0: 0, F1: -1, F2: -146}
            _ = v7
        } else {
            p0 = p0
        }
    case "a":
        var v8 float64 = 0.0
        _ = v8
    default:
        p0 = p0
    }
    v9 := []int{}
    _ = p0
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 float64 = 3.14159
    var v2 string = "a"
    v3 := []uint32{495, 378, 4294967295, 1}
    var v4 int = 2147483647
    v5 := []int{v4}
    var v6 bool = (-9223372036854775808 == 577)
    v6 = v6
    if len(v5) > 0 {
        v5[0] = (v4 & 127)
    }
    var v10 uint8 = 42
    v6 = ((^(v10)) <= v10)
    var v11 int32 = 790
    v12 := GenS2{F0: -753, F1: 429, F2: v11}
    var v13 uint8 = v10
    v5 = append(v5, v4)
    var v14 int = v4
    switch ("a" + (("" + "") + (v2 + v2))) {
    case "a":
        for fi15 := 0; fi15 < 20; fi15++ {
            v13 = (v10 + v10)
            v10 = v13
            _ = fi15
        }
    case "\n\t":
        if (!false) {
            if (42 == 1000) {
                _ = len(v5)
                v10 = v13
            }
            var v16 uint32 = 2147483648
            if (v6 && false) {
                switch v2 {
                case "hello":
                    v17 := GenS2{F0: -32768, F1: 4294967295, F2: v11}
                    _ = v17
                case "":
                    v18 := GenS2{F0: 32767, F1: 0, F2: (v11 + v11)}
                    _ = v18
                default:
                    fmt.Println(v11, v13, v2)
                }
            } else {
                if (false && false) {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v20 := GenS3{F0: v13, F1: ((("\n\t" + (v2 + ((v2 + "\n\t") + "foo bar"))) + "hello") + v2), F2: "hello"}
                        _ = fi19
                        _ = v20
                    }
                }
            }
            _ = v16
        }
    case "foo bar":
        v21 := GenS2{F0: 32767, F1: 18446744073709551615, F2: (v11 | (-(v11)))}
        _ = v21
    default:
        var v22 uint = 255
        _ = v22
    }
    v2 = (v2 + "hello")
    for fi23 := 0; fi23 < 8; fi23++ {
        if ((!v6) || (!false)) {
            fmt.Println(v11, v2, v6)
        } else {
            if v6 {
                if false {
                    if ((972 >= 61) && (-1 < v11)) {
                        var v24 int = (v14 & 32768)
                        var v25 int8 = 127
                        _ = v24
                        _ = v25
                    }
                    v6 = v6
                    v26 := GenS1{F0: genFn0(v2), F1: (v11 << ((1) & 63)), F2: -1}
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        v13 += 42
                        v13 -= 142
                        fmt.Println(v4, fi27)
                        switch (v4 % ((v4) | 1)) {
                        case -637:
                            v28 := GenS4{F0: -287}
                            _ = v28
                        case 32768:
                            var v29 uint16 = 0
                            _ = v29
                        case 683:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v14 -= (fi27 >> ((708) & 63))
                        _ = fi27
                    }
                }
                for fi30 := 0; fi30 < 2; fi30++ {
                    v31 := GenS4{F0: 32767}
                    v32 := GenS5{F0: -1}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v4 -= -124
            }
            v1 += 1.5
        }
        v33 := GenS5{F0: -339}
        v34 := GenS1{F0: 394, F1: (v11 ^ -383), F2: int64(v13)}
        _ = fi23
        _ = v33
        _ = v34
    }
    var v35 bool = (v11 > -2147483648)
    for fi36 := 0; fi36 < 15; fi36++ {
        v35 = (1e10 > 3.14159)
        if v6 {
            var v37 float32 = 1e10
            v10 = (v10 | v10)
            var v38 float64 = v1
            _ = v37
            _ = v38
        }
        _ = fi36
    }
    if v35 {
        _ = len(v3)
        var v39 uint8 = (v10 / (((^(v13))) | 1))
        v35 = (4294967295 == 453)
        _ = v39
    } else {
        var v40 string = "hello"
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 int = 9223372036854775807
            _ = fi41
            _ = v42
        }
        fmt.Println(v2, v10, v11)
        _ = v40
    }
    var v43 int32 = (^(v11))
    switch v14 {
    case -790:
        var v44 uint64 = 889
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
