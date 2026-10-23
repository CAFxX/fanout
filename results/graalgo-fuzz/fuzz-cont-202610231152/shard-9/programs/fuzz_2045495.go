package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 bool }
type GenS5 struct { F0 int16; F1 bool }
func genFn0(p0 uint32, p1 int16, p2 uint16) int8 {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int8 = -128
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (^(p0))}
    v8 := GenS2{F0: 0, F1: ("" + ("a" + "hello"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint8 {
    v9 := []int{}
    var v10 float32 = ((0.0 + float32(1e10)) + 3.14159)
    var v11 int16 = 0
    _ = v10
    _ = v11
    _ = v9
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 2147483647
    var v3 float64 = float64(0.1)
    v4 := make(map[string]string)
    var v12 int = 9223372036854775807
    v3 = v3
    v13 := v4[(("" + ("a" + "a")) + "")]
    v14 := GenS1{F0: 65535}
    v15 := GenS2{F0: 2, F1: "a"}
    fmt.Println(v12)
    v13 = v13
    if false {
        if true {
            var v16 int8 = -1
            _ = v16
        } else {
            if (v3 < ((-2.25 - v3) / ((v3 + 1.5) * v3))) {
                var v17 int16 = -686
                _ = v17
            }
            switch -1 {
            case 746:
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := GenS3{F0: v2, F1: 1000}
                    fmt.Println(fi18, fi18, v13)
                    v13 = ""
                    _ = fi18
                    _ = v19
                }
            default:
                var v20 int64 = -53
                _ = v20
            }
            var v21 uint8 = 128
            _ = v21
        }
        if true {
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := v4[v13]
                _ = fi22
                _ = v23
            }
            v3 -= float64((1e10 - 1e10))
        } else {
            fmt.Println(v13, v2, v12)
            var v24 string = v13
            switch (v12 ^ 1) {
            case 42:
                for fi25 := 0; fi25 < 1; fi25++ {
                    _ = len(v4)
                    _ = fi25
                }
            case 0:
                if (((0.1 <= 0.1) || false) && true) {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        fmt.Println(fi26, v3, v3)
                        for fi27 := 0; fi27 < 15; fi27++ {
                            v28 := []int{(fi26 / ((2147483647) | 1)), fi26, fi26}
                            v29 := GenS3{F0: v2, F1: -161}
                            var v30 int64 = -1
                            var v31 uint16 = 32767
                            v24 = ("" + ("hello" + (v24 + "foo bar")))
                            if false {
                                var v32 float32 = -2.25
                                v33 := GenS1{F0: 96}
                                var v34 int64 = (v30 + (v30 & v30))
                                _ = v32
                                _ = v33
                                _ = v34
                            } else {
                                fmt.Println(v30, fi27)
                                var v35 uint32 = 191
                                v12 -= fi27
                                _ = v35
                            }
                            fmt.Println(v3, fi27)
                            _ = fi27
                            _ = v28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        var v36 int16 = 1000
                        _ = fi26
                        _ = v36
                    }
                    v37 := GenS4{F0: 42, F1: 2, F2: false}
                    _ = v37
                }
            }
            _ = v24
        }
    }
    delete(v4, v13)
    switch v12 {
    case -2147483648:
        var v38 int = (v12 >> ((0) & 63))
        _ = v38
    case 9223372036854775807:
        for fi39 := 0; fi39 < 14; fi39++ {
            for fi40 := 0; fi40 < 17; fi40++ {
                v2 = (v2 / (((v2 | (v2 ^ -771))) | 1))
                if false {
                    fmt.Println(fi40)
                }
                _ = fi40
            }
            for fi41 := 0; fi41 < 10; fi41++ {
                v12 += fi41
                _ = fi41
            }
            var v42 float32 = (-2.25 + 1e10)
            _ = fi39
            _ = v42
        }
    }
    var v43 int64 = 42
    var v44 int16 = -1
    v3 = 3.14159
    for fi45 := 0; fi45 < 14; fi45++ {
        v3 = ((((-2.25 - float64(float32(float64(3.14159)))) * v3) / (float64(0.1) + v3)) - (3.14159 + 1.5))
        v46 := GenS3{F0: v2, F1: (v44 ^ v44)}
        _ = fi45
        _ = v46
    }
    var v47 int8 = -69
    v4[""] = "hello"
    for fi48 := 0; fi48 < 5; fi48++ {
        v49 := GenS5{F0: v44, F1: false}
        var v50 int8 = (v47 + (v47 << ((65535) & 63)))
        if ((v12 - 1) <= fi48) {
            v51 := []int{fi48}
            delete(v4, "\n\t")
            _ = v51
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v47
}
