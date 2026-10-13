package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 string }
func genFn0(p0 bool, p1 uint16) string {
    p0 = p0
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := []int{fi4, (fi4 * 908)}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        if (!(p1 >= (^(p1)))) {
            var v7 uint32 = uint32(p1)
            v8 := []int{fi6}
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: (fi6 % ((444) | 1)), F1: 754, F2: 499}
            _ = v9
        }
        p0 = (-313 > 1)
        if ("foo bar" < "foo bar") {
            if (567 <= (p1 / ((42) | 1))) {
                switch "hello" {
                case "hello":
                    var v10 uint16 = (^(p1))
                    _ = v10
                case "a":
                    v11 := GenS2{F0: 127, F1: "hello"}
                    _ = v11
                case "":
                    v12 := []int{}
                    _ = v12
                }
                p1 = p1
                var v13 int8 = 1
                _ = v13
            }
            var v14 uint32 = 735
            _ = v14
        } else {
            var v15 uint32 = 531
            fmt.Println(fi6, p1, p0)
            v16 := GenS1{F0: (-(fi6)), F1: 42, F2: -535}
            _ = v15
            _ = v16
        }
        _ = fi6
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            p1 = (p1 & (p1 ^ p1))
            var v19 bool = (-56 >= 42)
            v20 := []int{(fi17 / ((fi18) | 1)), (fi18 / ((fi17) | 1)), 771, fi17}
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 float32 = 0.1
        _ = fi17
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        p0 = p0
        v23 := []int{(fi22 % ((127) | 1)), fi22}
        var v24 int16 = 279
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = p0
    _ = p1
    return "\n\t"
}
func genFn1() int32 {
    if (true && (-200 == 128)) {
        if (-517 > 4294967295) {
            var v25 int64 = 0
            v25 -= v25
            if false {
                fmt.Println(v25)
                if (-9223372036854775808 > v25) {
                    v26 := []int{582}
                    if (-128 <= -1) {
                        var v27 int16 = -337
                        if true {
                            fmt.Println(v25, v25)
                        } else {
                            if (true && false) {
                                if len(v26) > 0 {
                                    v26[0] = -9223372036854775808
                                }
                                fmt.Println(v25, v25)
                                var v28 float64 = 0.0
                                _ = v28
                            } else {
                                if false {
                                    if true {
                                        v29 := GenS1{F0: 0, F1: 0, F2: 1}
                                        _ = v29
                                    } else {
                                        v30 := GenS3{F0: 715}
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        _ = v30
                                    }
                                    v26 = append(v26, -392)
                                    var v31 int32 = 2147483647
                                    _ = v31
                                }
                            }
                        }
                        var v32 uint64 = 1
                        _ = v27
                        _ = v32
                    }
                    _ = v26
                }
                v33 := []int{-2147483648, 128, 199}
                _ = v33
            }
            _ = v25
        } else {
            if (42 >= 0) {
                v34 := GenS2{F0: 127, F1: ""}
                if true {
                    var v35 int32 = 0
                    v35 += -2147483648
                    v35 -= (v35 * v35)
                    _ = v35
                }
                var v36 float32 = 1e10
                _ = v34
                _ = v36
            } else {
                var v37 int16 = 42
                fmt.Println(v37)
                _ = v37
            }
        }
    }
    v38 := GenS1{F0: -805, F1: 492, F2: 42}
    _ = v38
    return 42
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    v1 = (v1 * (v1 / ((v1) | 1)))
    switch 128 {
    case 128:
        delete(v3, "foo bar")
    }
    v39 := GenS1{F0: 32768, F1: 0, F2: -1}
    v1 += (v1 ^ -9223372036854775808)
    var v40 uint8 = 127
    v41 := GenS4{F0: 1, F1: "hello"}
    v40 -= (^(v40))
    delete(v3, (("a" + ("a" + "a")) + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
}
