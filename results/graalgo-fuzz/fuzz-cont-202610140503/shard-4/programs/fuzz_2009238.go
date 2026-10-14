package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 bool; F1 int8; F2 int64 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint }
func genFn0(p0 uint8) uint64 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{(fi6 - (fi6 + fi6)), (^(fi6))}
        v8 := GenS1{F0: 65535, F1: 0, F2: -128}
        var v9 int64 = 0
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 bool = true
    p0 = (^(p0))
    _ = p0
    _ = v10
    return 4294967295
}
func genFn1(p0 int8, p1 uint8) int8 {
    var v11 uint = 388
    var v12 int8 = -128
    if true {
        fmt.Println(v12, v12)
    } else {
        fmt.Println(v12, v11, p1)
        if true {
            fmt.Println(v11)
        }
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := []int{fi13, fi13}
        _ = len(v14)
        if false {
            switch "" {
            case "\n\t":
                p1 -= (p1 / ((26) | 1))
            }
            v11 = uint(p1)
            p1 += p1
        } else {
            if false {
                var v15 int32 = 2147483647
                _ = v15
            }
            if (824 == 1) {
                v12 += -73
                if false {
                    v12 -= (^(p0))
                    switch (^(fi13)) {
                    case 32767:
                        for fi16 := 0; fi16 < 16; fi16++ {
                            if (fi13 >= (fi13 / (((-(fi13))) | 1))) {
                                if len(v14) > 0 {
                                    v14[0] = -767
                                }
                                fmt.Println(fi13, p0)
                                var v17 float64 = float64(1.5)
                                _ = v17
                            }
                            var v18 float32 = ((-2.25 / 1.5) - 0.0)
                            _ = fi16
                            _ = v18
                        }
                    case 42:
                        v11 += v11
                    }
                } else {
                    var v19 float32 = 3.14159
                    v20 := GenS1{F0: 122, F1: 4294967295, F2: v12}
                    v21 := GenS1{F0: 0, F1: 42, F2: 45}
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
            v11 += (v11 / ((980) | 1))
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return (v12 ^ (p0 >> ((1) & 63)))
}

func main() {
    var v1 uint64 = 42
    var v2 bool = false
    var v3 int = 2147483647
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    v22 := GenS2{F0: v3, F1: 42, F2: 1}
    fmt.Println(v3, v1)
    for fi23 := 0; fi23 < 11; fi23++ {
        v4[42] = 32767
        v2 = v2
        _ = fi23
    }
    if (!v2) {
        switch "\n\t" {
        case "a":
            var v24 int64 = 42
            _ = v24
        case "":
            v25 := GenS3{F0: -1, F1: v1}
            _ = v25
        case "foo bar":
            fmt.Println(v1)
        default:
            _ = len(v5)
        }
    } else {
        v1 = v1
        var v26 bool = v2
        v1 = v1
        _ = v26
    }
    delete(v4, 986)
    if (v3 <= v3) {
        fmt.Println(v2, v2)
    } else {
        for fi27 := 0; fi27 < 7; fi27++ {
            switch ((((("" + ("a" + ("hello" + "foo bar"))) + "\n\t") + (("foo bar" + "") + ("\n\t" + ("hello" + (("foo bar" + "foo bar") + "a"))))) + ("\n\t" + "foo bar")) + "hello") {
            case "a":
                v1 += v1
            case "":
                switch 0 {
                case 32768:
                    v2 = false
                default:
                    v3 += (fi27 << ((65535) & 63))
                }
            case "hello":
                v28 := GenS4{F0: false, F1: -1, F2: 640}
                _ = v28
            }
            v29 := GenS5{F0: 42, F1: (((("a" + ("hello" + "\n\t")) + "a") + "hello") + ("\n\t" + ""))}
            _ = fi27
            _ = v29
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            if (2147483648 == 2) {
                v3 = (v3 & (v3 + v3))
                var v31 int = 32767
                _ = v31
            }
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 uint8 = 42
                var v34 int16 = 1
                _ = fi32
                _ = v33
                _ = v34
            }
            for fi35 := 0; fi35 < 1; fi35++ {
                if (v1 >= genFn0(37)) {
                    switch (fi35 | (^(fi30))) {
                    case 0:
                        fmt.Println(fi30)
                    default:
                        v2 = (1000 == 32767)
                    }
                    if false {
                        fmt.Println(v3, v1)
                    }
                    var v36 uint8 = 0
                    _ = v36
                } else {
                    var v37 int8 = 1
                    var v38 uint32 = 913
                    _ = v37
                    _ = v38
                }
                v39 := GenS3{F0: 2147483647, F1: (v1 - v1)}
                v40 := GenS2{F0: (fi35 ^ v3), F1: 76, F2: int8(v1)}
                _ = fi35
                _ = v39
                _ = v40
            }
            _ = fi30
        }
        var v41 uint16 = 153
        _ = v41
    }
    var v42 uint8 = 127
    var v43 int8 = 1
    if ((((-2.25 <= 0.0) || (!true)) || true) && (2 < 32767)) {
        v44 := GenS6{F0: 1}
        var v45 uint = 18446744073709551615
        _ = v44
        _ = v45
    } else {
        switch (-(v3)) {
        case -2147483648:
            v46 := v5[661]
            _ = v46
        case 1:
            v47 := []int{v3}
            _ = v47
        case 128:
            for fi48 := 0; fi48 < 7; fi48++ {
                v42 = v42
                _ = fi48
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
