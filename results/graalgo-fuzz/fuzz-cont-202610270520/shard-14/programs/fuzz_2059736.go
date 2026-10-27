package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint16 }
func genFn0(p0 string) int {
    p0 = "foo bar"
    fmt.Println(p0)
    v5 := []int{}
    for fi6 := 0; fi6 < 10; fi6++ {
        if len(v5) > 0 {
            v5[0] = (fi6 / (((fi6 ^ (fi6 * (fi6 / ((fi6) | 1))))) | 1))
        }
        v7 := GenS1{F0: 42, F1: 2147483647, F2: ("\n\t" + ("foo bar" + p0))}
        var v8 string = p0
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return -484
}
func genFn1(p0 uint, p1 int) int64 {
    fmt.Println(p0, p0)
    switch p1 {
    case -676:
        var v9 uint16 = 2
        _ = v9
    case 42:
        if (1 == 1) {
            v10 := GenS2{F0: 303, F1: (("hello" + "\n\t") + "a")}
            _ = v10
        }
    case 0:
        v11 := GenS3{F0: (("" + "\n\t") + ""), F1: 0}
        _ = v11
    default:
        var v12 int32 = 432
        _ = v12
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]bool)
    var v3 int8 = (v1 ^ v1)
    var v4 uint = 0
    if (128 >= 42) {
        var v13 bool = true
        switch genFn0("\n\t") {
        case 127:
            v14 := GenS1{F0: 854, F1: 341, F2: (("a" + "") + "a")}
            _ = v14
        case 128:
            if false {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v1 = (v3 >> ((v4) & 63))
                    _ = fi15
                }
                v16 := GenS3{F0: "foo bar", F1: 43}
                _ = v16
            }
        }
        fmt.Println(v1)
        _ = v13
    } else {
        var v17 bool = false
        v18 := GenS4{F0: "", F1: "foo bar"}
        v19 := GenS5{F0: 255, F1: 1, F2: 400}
        _ = v17
        _ = v18
        _ = v19
    }
    switch ("foo bar" + "") {
    case "\n\t":
        var v20 float32 = 3.14159
        _ = v20
    default:
        switch genFn0("") {
        case -720:
            delete(v2, 42)
        case -811:
            v21 := GenS3{F0: ("\n\t" + "\n\t"), F1: 42}
            _ = v21
        }
    }
    if (102 >= (v1 / (((v3 << ((v4) & 63))) | 1))) {
        delete(v2, 0)
        if (!(0.0 >= 0.0)) {
            for fi22 := 0; fi22 < 2; fi22++ {
                if false {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        var v24 string = ""
                        var v25 int16 = 32767
                        v26 := []int{(fi22 & 4294967295), (fi23 | (fi22 << (((v4 / (((^(v4))) | 1))) & 63))), genFn0(v24)}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS1{F0: 0, F1: 2, F2: "hello"}
                    for fi28 := 0; fi28 < 18; fi28++ {
                        if true {
                            var v29 bool = (((!false) || (false && (("a" + "\n\t") >= "\n\t"))) && false)
                            _ = v29
                        }
                        var v30 uint16 = 269
                        _ = fi28
                        _ = v30
                    }
                    _ = v27
                } else {
                    switch fi22 {
                    case 281:
                        for fi31 := 0; fi31 < 1; fi31++ {
                            if true {
                                v32 := v2[fi31]
                                _ = v32
                            } else {
                                v33 := GenS6{F0: 32767, F1: 0, F2: 1}
                                _ = len(v2)
                                _ = v33
                            }
                            var v34 uint32 = 42
                            _ = fi31
                            _ = v34
                        }
                    case 42:
                        v3 = v1
                    default:
                        fmt.Println(v4, v4, v3)
                    }
                    for fi35 := 0; fi35 < 9; fi35++ {
                        var v36 float64 = (1.5 / (-2.25 * (0.0 + float64(float32(3.14159)))))
                        _ = fi35
                        _ = v36
                    }
                    v3 = (-(v1))
                }
                v37 := GenS3{F0: "", F1: 180}
                _ = fi22
                _ = v37
            }
            var v38 bool = (0 <= 1)
            switch 42 {
            case 503:
                v39 := []int{32768}
                _ = v39
            default:
                var v40 uint32 = 2
                _ = v40
            }
            _ = v38
        }
    } else {
        if false {
            v3 = (^(v1))
        }
    }
    _ = len(v2)
    v1 = (v1 / ((v1) | 1))
    v41 := GenS1{F0: 32767, F1: 2147483647, F2: ("a" + "")}
    v4 += v4
    delete(v2, 0)
    v42 := GenS1{F0: 0, F1: 823, F2: ("hello" + "a")}
    var v43 uint32 = 721
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
