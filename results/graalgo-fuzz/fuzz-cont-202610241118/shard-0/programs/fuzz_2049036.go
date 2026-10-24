package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int64 }
func genFn0() int16 {
    v7 := []int{4294967295, 32768}
    var v8 uint = 75
    var v9 uint32 = 65535
    var v10 float64 = 3.14159
    var v11 int8 = 50
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int, p1 uint8, p2 int16) int {
    p0 = (p0 << ((65535) & 63))
    fmt.Println(p1, p2)
    v12 := GenS1{F0: "a", F1: (p1 * p1)}
    v13 := []int{(p0 >> ((1) & 63)), (p0 * (p0 % ((127) | 1))), -751, p0}
    var v14 uint8 = (p1 - p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return p0
}

func main() {
    var v1 uint16 = 255
    var v2 int16 = 1000
    var v3 int8 = 42
    var v4 uint16 = (v1 % (((v1 % ((v1) | 1))) | 1))
    v5 := make(map[uint32]string)
    var v6 int32 = -845
    if ((1 >= 2) && true) {
        v1 -= (v4 ^ v4)
    } else {
        v15 := GenS2{F0: 1, F1: 18446744073709551615}
        fmt.Println(v1)
        _ = v15
    }
    if false {
        var v16 uint = uint(v1)
        switch "hello" {
        case "foo bar":
            v17 := GenS2{F0: -9223372036854775808, F1: (v16 % ((262) | 1))}
            _ = v17
        case "a":
            var v18 uint = v16
            _ = v18
        default:
            if true {
                if (!true) {
                    delete(v5, 247)
                    fmt.Println(v3)
                    var v19 uint = (v16 + (v16 - v16))
                    _ = v19
                }
            } else {
                v16 += (v16 ^ (^(v16)))
            }
        }
        fmt.Println(v16, v6)
        _ = v16
    }
    v20 := GenS2{F0: 369, F1: 65535}
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 bool = true
        for fi23 := 0; fi23 < 16; fi23++ {
            v3 = (v3 * v3)
            fmt.Println(v3)
            _ = fi23
        }
        v3 -= 100
        _ = fi21
        _ = v22
    }
    v24 := GenS3{F0: 166}
    if false {
        v4 += v1
        switch -2147483648 {
        case 338:
            v6 -= (-(v6))
        case 23:
            fmt.Println(v1)
        }
    }
    var v25 uint16 = 0
    for fi26 := 0; fi26 < 15; fi26++ {
        if (!false) {
            fmt.Println(v2, v1, fi26)
            for fi27 := 0; fi27 < 19; fi27++ {
                if (0.0 == 3.14159) {
                    fmt.Println(v2)
                    var v28 float64 = 3.14159
                    _ = v28
                }
                var v29 uint = 2
                _ = fi27
                _ = v29
            }
            fmt.Println(v6, v2)
        }
        v2 += 878
        _ = fi26
    }
    var v30 int = 42
    fmt.Println(v6, v3)
    var v31 uint32 = 799
    var v32 string = "foo bar"
    var v33 uint8 = 111
    v34 := []int{genFn1(genFn1(v30, 255, v2), 1, v2), genFn1((v30 | -851), 254, (v2 << ((42) & 63))), genFn1((v30 + -2147483648), v33, (v2 - (v2 + -1))), (^(v30))}
    for fi35 := 0; fi35 < 16; fi35++ {
        if false {
            for fi36 := 0; fi36 < 17; fi36++ {
                fmt.Println(fi36, v4)
                v2 = genFn0()
                _ = fi36
            }
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 int64 = -9223372036854775808
                v39 := GenS4{F0: v32}
                switch v30 {
                case 32768:
                    var v40 int32 = (v6 * v6)
                    v30 = (-(fi35))
                    delete(v5, (v31 + (^(v31))))
                    _ = v40
                case 2147483647:
                    if true {
                        v32 = ((v32 + ("" + v32)) + v32)
                    } else {
                        if false {
                            v2 = -32768
                        }
                    }
                    switch (fi35 >> ((2) & 63)) {
                    case 4294967295:
                        v33 = (v33 >> ((295) & 63))
                    }
                    var v41 uint32 = 1
                    _ = v41
                case -1:
                    var v42 int = 725
                    _ = v42
                default:
                    if true {
                        var v43 int16 = int16(v6)
                        _ = v43
                    } else {
                        if true {
                            if (v31 != (v31 & 6)) {
                                v44 := GenS5{F0: (v31 | uint32(v33)), F1: 127}
                                switch fi35 {
                                case 32767:
                                    _ = len(v34)
                                }
                                _ = v44
                            } else {
                                fmt.Println(v32, v1, v2)
                            }
                            v45 := GenS3{F0: (v33 & v33)}
                            _ = v45
                        }
                        _ = len(v34)
                        v33 = v33
                    }
                }
                if ((v32 + "hello") < "\n\t") {
                    v46 := GenS6{F0: -875}
                    var v47 int32 = v6
                    v48 := GenS2{F0: (v38 << ((65535) & 63)), F1: uint(v4)}
                    _ = v46
                    _ = v47
                    _ = v48
                } else {
                    var v49 uint = 18446744073709551615
                    _ = v49
                }
                switch (fi37 >> ((18446744073709551615) & 63)) {
                case -845:
                    v3 += (v3 & (v3 / ((100) | 1)))
                case 32768:
                    fmt.Println(v6, v25, v6)
                default:
                    fmt.Println(v30)
                }
                _ = fi37
                _ = v38
                _ = v39
            }
        }
        _ = fi35
    }
    if (-923 == v6) {
        v2 += 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
