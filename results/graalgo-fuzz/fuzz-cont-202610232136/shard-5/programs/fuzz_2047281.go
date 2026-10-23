package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
func genFn0(p0 uint16, p1 uint8, p2 int) int32 {
    var v5 uint8 = (p1 - 254)
    p1 = v5
    fmt.Println(p2)
    var v6 int32 = 468
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v6 + v6)
}
func genFn1(p0 int16, p1 int32, p2 uint32) uint {
    v7 := GenS1{F0: -1, F1: 2}
    var v8 float32 = ((float32(0.0) + 0.1) * (3.14159 / 1.0))
    v9 := []int{}
    var v10 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{}
    var v3 uint = 42
    var v4 uint16 = 32767
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint16 = (v4 * (v4 << ((162) & 63)))
            v13 := v1[2147483647]
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 int32 = 927
    fmt.Println(v14, v14, v3)
    if true {
        var v15 int8 = -128
        if ((("\n\t" + ("" + "\n\t")) != "a") || true) {
            v15 += v15
        }
        var v16 uint8 = uint8(v4)
        _ = v15
        _ = v16
    }
    _ = len(v1)
    switch ((("foo bar" + "") + ("hello" + "foo bar")) + "a") {
    case "hello":
        v1[37] = 0
    case "\n\t":
        for fi17 := 0; fi17 < 1; fi17++ {
            switch 0 {
            case 292:
                for fi18 := 0; fi18 < 11; fi18++ {
                    _ = len(v2)
                    _ = fi18
                }
            case 32767:
                if (9223372036854775807 < fi17) {
                    v2 = append(v2, fi17)
                    var v19 bool = (("foo bar" <= "foo bar") && (false && true))
                    if (255 >= 1) {
                        var v20 uint32 = 65535
                        if ((("\n\t" + "hello") + "a") > ("\n\t" + "a")) {
                            if (v19 || v19) {
                                if v19 {
                                    for fi21 := 0; fi21 < 12; fi21++ {
                                        switch fi17 {
                                        case 167:
                                            switch ("\n\t" + "a") {
                                            case "\n\t":
                                                v20 = v20
                                            case "foo bar":
                                                if ((!((int64(v3) == int64(v14)) && v19)) && (("a" <= "foo bar") || false)) {
                                                    v22 := GenS1{F0: 1, F1: 0}
                                                    var v23 int64 = -1
                                                    _ = v22
                                                    _ = v23
                                                }
                                            case "a":
                                                if len(v2) > 0 {
                                                    v2[0] = fi21
                                                }
                                            default:
                                                v24 := GenS2{F0: 42}
                                                _ = v24
                                            }
                                        case -9223372036854775808:
                                            v25 := GenS2{F0: 4294967295}
                                            _ = v25
                                        }
                                        _ = fi21
                                    }
                                    switch fi17 {
                                    case 0:
                                        _ = len(v2)
                                    case -9223372036854775808:
                                        v4 = (v4 % ((65535) | 1))
                                    default:
                                        v14 = 509
                                    }
                                }
                                v26 := GenS2{F0: 1}
                                _ = v26
                            } else {
                                var v27 int8 = 106
                                var v28 int8 = (v27 + 100)
                                _ = v27
                                _ = v28
                            }
                        } else {
                            v4 -= (v4 >> (((^(v3))) & 63))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = v20
                    }
                    _ = v19
                } else {
                    var v29 int16 = 1000
                    _ = v29
                }
            }
            var v30 int64 = int64(v4)
            _ = fi17
            _ = v30
        }
    case "a":
        var v31 uint16 = v4
        _ = v31
    }
    var v32 uint8 = 128
    switch ("foo bar" + "hello") {
    case "\n\t":
        if true {
            v14 += (-(v14))
        }
    default:
        var v33 uint = (^(v3))
        _ = v33
    }
    if false {
        v3 += (v3 << (((v3 * 634)) & 63))
        v34 := GenS3{F0: -9223372036854775808, F1: (v3 & v3), F2: (v32 << (((v3 ^ 172)) & 63))}
        _ = v34
    } else {
        if ((v32 < 131) || (!(!((false && (0 <= 42)) && (0.0 >= (3.14159 * 0.1)))))) {
            fmt.Println(v3, v32)
        } else {
            v3 -= 4294967295
        }
    }
    delete(v1, 683)
    for fi35 := 0; fi35 < 15; fi35++ {
        v36 := GenS1{F0: -1, F1: (v32 * (v32 * v32))}
        var v37 int32 = v14
        _ = fi35
        _ = v36
        _ = v37
    }
    v14 += 382
    var v38 int16 = -1
    var v39 string = "a"
    var v40 int8 = int8(v32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
