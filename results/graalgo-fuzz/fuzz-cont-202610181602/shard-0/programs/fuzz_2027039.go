package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 string }
func genFn0(p0 uint32) uint64 {
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{48, fi7, (fi7 << ((144) & 63)), -1}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 100, F1: -2147483648, F2: -982}
    v10 := GenS2{F0: (42 != 255), F1: p0, F2: 374}
    _ = p0
    _ = v10
    _ = v6
    _ = v9
    return 556
}
func genFn1(p0 int32, p1 uint64) int {
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    var v13 string = (("\n\t" + (("hello" + "hello") + (("a" + "foo bar") + "\n\t"))) + "hello")
    var v14 int64 = -9223372036854775808
    v15 := []int{}
    v15 = append(v15, -240)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return int(p1)
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]string)
    var v3 uint64 = (v1 * v1)
    var v4 int8 = 127
    var v5 float64 = 1.5
    v3 -= 42
    v16 := GenS3{F0: 923, F1: ("a" + ("hello" + "")), F2: 1}
    delete(v2, 665)
    if (genFn0(781) < v1) {
        fmt.Println(v1, v5)
    } else {
        fmt.Println(v3, v1, v5)
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        if false {
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 bool = (genFn1(2147483647, (v1 * v3)) <= genFn1(-648, genFn0(4294967295)))
                _ = fi18
                _ = v19
            }
            if false {
                v20 := GenS2{F0: (true && (false && (127 >= 51))), F1: 65535, F2: 734}
                if false {
                    var v21 float64 = float64((1.5 + 1.5))
                    if (true || ((((1.5 * 3.14159) * 0.0) - 3.14159) > 1e10)) {
                        fmt.Println(v4)
                    } else {
                        fmt.Println(v1, v21, v4)
                        for fi22 := 0; fi22 < 12; fi22++ {
                            fmt.Println(v3, v1)
                            var v23 float64 = 0.0
                            if ((true || false) && (1 < 0)) {
                                var v24 uint64 = v1
                                _ = v24
                            }
                            _ = fi22
                            _ = v23
                        }
                        v25 := []int{fi17, fi17}
                        _ = v25
                    }
                    _ = v21
                }
                _ = v20
            } else {
                v3 = v1
            }
        } else {
            v1 += (v3 << ((1) & 63))
            var v26 uint32 = 320
            _ = v26
        }
        v4 += v4
        v3 -= genFn0(971)
        _ = fi17
    }
    var v27 string = "hello"
    var v28 int = 558
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 15; fi30++ {
            switch fi30 {
            case 42:
                var v31 uint32 = 591
                _ = v31
            default:
                if (v3 != (v1 & genFn0(2))) {
                    v32 := v2[fi29]
                    var v33 uint16 = 32767
                    var v34 uint16 = (v33 << ((65535) & 63))
                    if false {
                        v35 := []int{v28, fi29, fi29, v28}
                        _ = v35
                    }
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    v27 = ("a" + (v27 + "a"))
                }
                v5 = v5
                var v36 string = v27
                switch v36 {
                case "\n\t":
                    v36 = v27
                case "":
                    v27 = (v27 + (v36 + "foo bar"))
                default:
                    if true {
                        switch (^(v28)) {
                        case -9223372036854775808:
                            fmt.Println(v4, fi29, v28)
                        case 820:
                            fmt.Println(v3, v1)
                        }
                        var v37 uint = 252
                        _ = v37
                    }
                }
                var v38 int8 = (v4 + (v4 - -68))
                _ = v36
                _ = v38
            }
            var v39 float32 = 0.1
            _ = fi30
            _ = v39
        }
        v40 := GenS3{F0: 739, F1: (((((v27 + v27) + v27) + "a") + "hello") + (v27 + "\n\t")), F2: 9223372036854775807}
        _ = fi29
        _ = v40
    }
    v41 := GenS4{F0: 485, F1: (v28 | (^(v28)))}
    delete(v2, v28)
    fmt.Println(v27)
    var v42 uint32 = 0
    v43 := GenS5{F0: "foo bar"}
    v44 := GenS2{F0: false, F1: 1, F2: 1}
    var v45 float64 = (v5 * v5)
    fmt.Println(v5, v27)
    var v46 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
