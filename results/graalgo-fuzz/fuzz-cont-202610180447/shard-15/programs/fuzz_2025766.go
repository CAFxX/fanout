package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0() uint16 {
    switch "" {
    case "a":
        v6 := GenS1{F0: "a", F1: false}
        _ = v6
    case "foo bar":
        var v7 int16 = 820
        _ = v7
    default:
        v8 := GenS1{F0: "hello", F1: (1.5 <= -2.25)}
        _ = v8
    }
    var v9 float32 = 1e10
    fmt.Println(v9)
    _ = v9
    return 32767
}
func genFn1() uint16 {
    switch -129 {
    case -38:
        v10 := []int{236, 32768}
        _ = v10
    case -358:
        var v11 string = ""
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := GenS2{F0: 1}
            var v14 uint16 = genFn0()
            v14 -= genFn0()
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        var v15 string = "hello"
        var v16 uint8 = 255
        for fi17 := 0; fi17 < 6; fi17++ {
            if false {
                var v18 int = 169
                if true {
                    v18 += (v18 + v18)
                }
                _ = v18
            }
            v19 := []int{fi17, (fi17 % (((fi17 ^ fi17)) | 1)), fi17, 2147483647}
            if len(v19) > 0 {
                v19[0] = (fi17 % ((fi17) | 1))
            }
            _ = fi17
            _ = v19
        }
        _ = v15
        _ = v16
    }
    return 0
}

func main() {
    var v1 uint = 65535
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    var v4 int32 = 2147483647
    v5 := make(map[uint32]int)
    switch -624 {
    case 32767:
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 uint32 = 42
            for fi22 := 0; fi22 < 16; fi22++ {
                switch "a" {
                case "":
                    if false {
                        var v23 uint32 = 4294967295
                        if (uint(v21) <= (v1 * v1)) {
                            if (42 != (fi22 ^ (fi20 ^ (fi22 * fi20)))) {
                                v24 := v5[v21]
                                var v25 int = fi20
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS1{F0: ("" + "hello"), F1: ((v1 > (v1 % ((800) | 1))) && false)}
                            if false {
                                v27 := []int{(-(fi22)), fi20, (fi22 >> (((^(v1))) & 63)), (-(fi22))}
                                v4 = 727
                                _ = v27
                            } else {
                                if (!(v2 == ((v2 / v2) + v2))) {
                                    var v28 uint32 = v21
                                    v29 := GenS1{F0: ("" + "\n\t"), F1: (!(((false && (false && (!(!true)))) || true) && ((fi22 % ((fi20) | 1)) != fi22)))}
                                    v30 := []int{}
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                v31 := GenS1{F0: (("foo bar" + ("foo bar" + "hello")) + "hello"), F1: (-21 <= 42)}
                                delete(v3, (^(v21)))
                                _ = v31
                            }
                            if (!(false && true)) {
                                v23 = 42
                                v5[2147483647] = (fi20 * (fi20 & -9223372036854775808))
                                if ((^(v21)) <= v21) {
                                    fmt.Println(v21, fi22)
                                    v23 = uint32(v1)
                                }
                                var v32 int32 = (v4 * int32(v1))
                                _ = v32
                            }
                            _ = v26
                        }
                        _ = v23
                    }
                }
                var v33 uint16 = genFn0()
                fmt.Println(fi20, v33, v4)
                v34 := GenS3{F0: true, F1: 32767}
                _ = fi22
                _ = v33
                _ = v34
            }
            var v35 float32 = float32(1e10)
            _ = fi20
            _ = v21
            _ = v35
        }
    case -816:
        v1 = v1
    case 128:
        v3[2] = false
    }
    v2 -= -2.25
    for fi36 := 0; fi36 < 6; fi36++ {
        v37 := []int{42}
        _ = fi36
        _ = v37
    }
    var v38 string = "\n\t"
    var v39 int32 = 42
    var v40 uint16 = 496
    v41 := []int{-409, 73, -1}
    var v42 string = ((v38 + v38) + (v38 + "hello"))
    v2 = v2
    v43 := GenS1{F0: (v42 + v38), F1: true}
    for fi44 := 0; fi44 < 15; fi44++ {
        v42 = "foo bar"
        v2 = ((v2 * float32(-2.25)) * v2)
        if len(v41) > 0 {
            v41[0] = (fi44 ^ (fi44 << ((v1) & 63)))
        }
        _ = fi44
    }
    if (!((v4 == v39) && true)) {
        if false {
            for fi45 := 0; fi45 < 6; fi45++ {
                v46 := GenS4{F0: genFn1(), F1: 537}
                v47 := GenS2{F0: 2}
                _ = fi45
                _ = v46
                _ = v47
            }
        } else {
            v4 = 0
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
