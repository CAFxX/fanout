package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 int8; F2 int8 }
func genFn0() int {
    var v4 uint64 = 4294967295
    if true {
        if false {
            var v5 uint = 1
            _ = v5
        }
        switch (("\n\t" + ("\n\t" + "")) + "") {
        case "\n\t":
            v4 = (v4 | (^(v4)))
        case "":
            var v6 uint = 65535
            _ = v6
        default:
            switch -2147483648 {
            case -832:
                v4 = v4
            case -415:
                v4 += (v4 ^ v4)
            case 0:
                var v7 uint8 = 128
                _ = v7
            default:
                var v8 int16 = -767
                _ = v8
            }
        }
    }
    v4 = v4
    v9 := []int{-1, -748, 4294967295, 2147483647}
    _ = v4
    _ = v9
    return 32767
}
func genFn1(p0 uint8) uint16 {
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := []int{fi10, genFn0(), 127, (-(fi10))}
        if ("hello" <= ("foo bar" + "\n\t")) {
            var v12 float64 = 0.1
            _ = v12
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            p0 = p0
            _ = fi13
        }
        _ = fi10
        _ = v11
    }
    v14 := []int{-280, -1, -924, int(p0)}
    if len(v14) > 0 {
        v14[0] = -1
    }
    v15 := GenS1{F0: 1, F1: 2147483647, F2: true}
    v16 := GenS1{F0: 116, F1: -510, F2: false}
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    var v3 int = 8
    delete(v1, v3)
    if false {
        switch "" {
        case "a":
            if true {
                if (0.0 != float32(((1e10 / (1.5 * 3.14159)) * float64((-2.25 + 1.5))))) {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v2 = (v2 >> ((1) & 63))
                        v2 -= (v2 - v2)
                        _ = fi17
                    }
                    if (42 <= -1) {
                        v3 = v3
                    } else {
                        v3 += -2147483648
                        for fi18 := 0; fi18 < 7; fi18++ {
                            var v19 int8 = 53
                            _ = fi18
                            _ = v19
                        }
                        var v20 int = v3
                        _ = v20
                    }
                }
                if (!(!true)) {
                    fmt.Println(v3, v2)
                    v21 := GenS2{F0: 1, F1: 0}
                    _ = v21
                } else {
                    if (1 > 1) {
                        var v22 int64 = v2
                        for fi23 := 0; fi23 < 16; fi23++ {
                            v22 = (v22 * (v22 % (((v2 + (v2 & 0))) | 1)))
                            _ = fi23
                        }
                        _ = v22
                    } else {
                        if true {
                            v24 := []int{32768, v3, (-(v3))}
                            _ = v24
                        }
                    }
                }
                if true {
                    for fi25 := 0; fi25 < 9; fi25++ {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v27 := GenS1{F0: 100, F1: -2147483648, F2: false}
                            if true {
                                v28 := GenS2{F0: 42, F1: 2147483647}
                                _ = v28
                            }
                            v3 = (fi26 % ((-272) | 1))
                            _ = fi26
                            _ = v27
                        }
                        _ = fi25
                    }
                } else {
                    v1[v3] = ""
                    for fi29 := 0; fi29 < 18; fi29++ {
                        v2 = v2
                        delete(v1, fi29)
                        _ = fi29
                    }
                    if (4294967295 != 2) {
                        v2 = (^(v2))
                    }
                }
            }
        case "hello":
            var v30 uint32 = 42
            _ = v30
        case "\n\t":
            v3 -= v3
        }
        if (!true) {
            fmt.Println(v2)
        } else {
            var v31 int32 = 42
            _ = v31
        }
        var v32 int32 = 0
        _ = v32
    }
    var v33 int64 = (-(v2))
    v34 := []int{127, (-(v3)), v3}
    if ((655 > 32767) || (0 != 0)) {
        v2 = v2
        fmt.Println(v3, v2, v33)
        v33 = -9223372036854775808
    }
    var v35 float32 = (3.14159 - ((((-2.25 / 1.0) - 1.5) + (0.0 / (3.14159 + -2.25))) * (0.0 - (1e10 - ((3.14159 - -2.25) * 1.5)))))
    v1[(v3 & genFn0())] = ("a" + ("a" + "foo bar"))
    switch (("hello" + (((("hello" + (("foo bar" + "foo bar") + ("\n\t" + "\n\t"))) + ("a" + "\n\t")) + "hello") + "a")) + "\n\t") {
    case "a":
        if true {
            var v36 float64 = 1e10
            _ = v36
        } else {
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
        }
    case "\n\t":
        v37 := GenS3{F0: 100}
        _ = v37
    }
    if (0 == 642) {
        v3 += v3
    }
    fmt.Println(v33, v35)
    if (100 == 127) {
        fmt.Println(v35, v35)
    }
    var v38 uint32 = 4294967295
    var v39 string = ""
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    if len(v34) > 0 {
        v34[0] = (^(v3))
    }
    fmt.Println(v38, v33)
    switch (v3 - v3) {
    case 32767:
        _ = len(v34)
    case -802:
        switch genFn0() {
        case 1:
            if (!(!false)) {
                v33 -= v2
                var v40 int = 113
                if false {
                    if len(v34) > 0 {
                        v34[0] = (-(v3))
                    }
                    var v41 bool = false
                    var v42 int64 = (v2 + (v2 >> ((393) & 63)))
                    _ = v41
                    _ = v42
                }
                _ = v40
            } else {
                v1[v3] = v39
                _ = len(v34)
                v3 -= -464
            }
        case 723:
            v35 += (((v35 * v35) + v35) * v35)
        case 128:
            var v43 uint16 = 594
            _ = v43
        }
    case 32768:
        if false {
            fmt.Println(v3, v33)
            if true {
                if (100 < 1) {
                    for fi44 := 0; fi44 < 15; fi44++ {
                        var v45 int = -9223372036854775808
                        if true {
                            v46 := GenS4{F0: genFn1(128), F1: false, F2: v45}
                            v47 := GenS3{F0: 1}
                            _ = v46
                            _ = v47
                        }
                        switch (v3 / ((genFn0()) | 1)) {
                        case 927:
                            var v48 uint32 = (^(v38))
                            _ = v48
                        case 2147483647:
                            v49 := GenS5{F0: 444, F1: -51, F2: 0}
                            _ = v49
                        }
                        _ = fi44
                        _ = v45
                    }
                    v3 -= (v3 | (-(v3)))
                } else {
                    var v50 uint8 = 2
                    var v51 int16 = 32767
                    _ = v50
                    _ = v51
                }
                v3 -= v3
                v35 -= float32(3.14159)
            } else {
                var v52 bool = ((255 < 1) || true)
                _ = v52
            }
        } else {
            v53 := GenS3{F0: -128}
            v54 := GenS2{F0: 0, F1: -436}
            v35 += v35
            _ = v53
            _ = v54
        }
    }
    v55 := v1[v3]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v55
}
