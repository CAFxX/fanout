package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint8; F1 int; F2 uint }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint64; F1 int8; F2 int64 }
func genFn0(p0 int64, p1 int64, p2 uint8) uint8 {
    v7 := GenS1{F0: 997, F1: (p0 | (p0 & (p0 + p1)))}
    var v8 int32 = 1
    v9 := []int{66, 737}
    if (!(!true)) {
        var v10 int32 = v8
        _ = v10
    } else {
        v9 = append(v9, 32768)
        switch 42 {
        case 32768:
            v11 := GenS1{F0: 18446744073709551615, F1: p1}
            _ = v11
        case 768:
            var v12 float64 = 0.0
            _ = v12
        }
        var v13 uint64 = 4294967295
        _ = v13
    }
    p0 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p2 / ((p2) | 1))
}
func genFn1(p0 uint16) uint8 {
    p0 -= (p0 ^ 0)
    p0 += (^(p0))
    _ = p0
    return 0
}

func main() {
    var v1 int16 = 210
    v2 := []string{"a", "hello", "hello", "a"}
    var v3 int32 = 2147483647
    var v4 uint = 86
    var v5 int32 = (v3 / (((v3 << ((65535) & 63))) | 1))
    var v6 uint32 = uint32(v4)
    v6 -= v6
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS2{F0: (v5 + v3)}
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                switch "a" {
                case "a":
                    v5 -= (v3 * (v5 & (v5 << ((v4) & 63))))
                case "":
                    if true {
                        if false {
                            if true {
                                v17 := GenS2{F0: 776}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v17
                            }
                            if (false || false) {
                                v18 := GenS1{F0: 18446744073709551615, F1: 0}
                                if (((true || false) && (true || true)) && (false || (!true))) {
                                    var v19 uint64 = 123
                                    _ = v19
                                }
                                _ = v18
                            } else {
                                var v20 int16 = (v1 / ((32767) | 1))
                                _ = v20
                            }
                        }
                    }
                    v21 := GenS3{F0: 53, F1: 0}
                    _ = v21
                case "\n\t":
                    v22 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807}
                    _ = v22
                }
                _ = fi16
            }
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v5, v6)
    for fi23 := 0; fi23 < 8; fi23++ {
        if ((!false) && true) {
            v3 += (v5 >> (((^(v4))) & 63))
            for fi24 := 0; fi24 < 4; fi24++ {
                v25 := GenS4{F0: 42, F1: (fi23 & (fi24 - -395)), F2: v4}
                fmt.Println(fi24, v5, v3)
                fmt.Println(fi24, v1)
                _ = fi24
                _ = v25
            }
        } else {
            var v26 int16 = -1
            v3 -= v3
            _ = v26
        }
        _ = fi23
    }
    if (!((42 == 1) && (65535 != 255))) {
        var v27 float32 = float32(0.1)
        v28 := GenS2{F0: v5}
        v29 := GenS1{F0: 4294967295, F1: 741}
        _ = v27
        _ = v28
        _ = v29
    }
    var v30 int64 = -132
    var v31 uint64 = 18446744073709551615
    switch "" {
    case "\n\t":
        _ = len(v2)
    case "":
        for fi32 := 0; fi32 < 15; fi32++ {
            switch (-(fi32)) {
            case 42:
                var v33 uint16 = 32767
                _ = v33
            }
            switch (^(fi32)) {
            case 42:
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 string = ("" + "")
                    if len(v2) > 0 {
                        v2[0] = ((v35 + (v35 + (("hello" + "") + v35))) + ((v35 + v35) + v35))
                    }
                    fmt.Println(v1, v30)
                    var v36 bool = false
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case -441:
                for fi37 := 0; fi37 < 5; fi37++ {
                    _ = len(v2)
                    var v38 int = (fi37 / (((fi32 + fi37)) | 1))
                    _ = fi37
                    _ = v38
                }
            }
            _ = fi32
        }
    case "a":
        for fi39 := 0; fi39 < 20; fi39++ {
            v40 := GenS5{F0: 1, F1: v1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi39
            _ = v40
        }
    default:
        v5 = v5
    }
    v41 := GenS3{F0: -1, F1: 100}
    var v42 float32 = -2.25
    v6 += v6
    v1 -= v1
    v43 := GenS2{F0: (v3 ^ (v5 % ((v5) | 1)))}
    var v44 float64 = 0.1
    v45 := GenS6{F0: v31, F1: 8, F2: v30}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
