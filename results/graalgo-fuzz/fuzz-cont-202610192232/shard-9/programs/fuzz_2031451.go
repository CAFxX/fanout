package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 string; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint32 }
func genFn0() string {
    if false {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: -580, F1: 62}
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 bool = (0 <= 421)
                v9 = (128 >= 2)
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{(fi10 % ((fi10) | 1)), fi10}
        v11 = append(v11, 4294967295)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 int64 = -1
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS1{F0: v13, F1: -102}
            var v16 int8 = 116
            _ = fi14
            _ = v15
            _ = v16
        }
        v13 += (v13 - (v13 * v13))
        _ = fi12
        _ = v13
    }
    return "\n\t"
}

func main() {
    var v1 int64 = 0
    var v2 uint64 = 1
    v3 := []string{""}
    v4 := make(map[int]bool)
    var v5 float32 = 1e10
    v5 = -2.25
    if false {
        fmt.Println(v1, v2)
        v1 = v1
    }
    v17 := GenS2{F0: 0, F1: (("a" + "hello") + ("\n\t" + ("" + (("foo bar" + (("" + "a") + ("" + ("a" + "\n\t")))) + ("foo bar" + ("" + ("hello" + ("\n\t" + "foo bar")))))))), F2: 42}
    if (!((v5 - ((((v5 * v5) * v5) + (3.14159 - (((v5 * 1.5) / 1.0) * 3.14159))) - v5)) != ((v5 + 0.0) * v5))) {
        fmt.Println(v2, v2)
    } else {
        v18 := GenS3{F0: true, F1: (v2 / (((v2 + 999)) | 1))}
        fmt.Println(v5, v1, v5)
        _ = v18
    }
    v1 = (^(v1))
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 float32 = (0.0 - float32(0.0))
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := v4[fi21]
        var v23 bool = v22
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 uint64 = v2
            v1 = v1
            _ = fi24
            _ = v25
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    var v26 uint = 145
    v5 -= ((((v5 + v5) - v5) + ((v5 * 0.0) - 3.14159)) * v5)
    for fi27 := 0; fi27 < 7; fi27++ {
        if true {
            _ = len(v4)
        } else {
            v28 := GenS2{F0: 4294967295, F1: "hello", F2: 0}
            var v29 bool = true
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v28
            _ = v29
        }
        v2 -= (v2 * (v2 / ((v2) | 1)))
        var v30 float32 = 0.0
        _ = fi27
        _ = v30
    }
    v31 := GenS3{F0: (42 >= 614), F1: (v2 % (((v2 * v2)) | 1))}
    if true {
        _ = len(v4)
    }
    var v32 uint64 = (^(v2))
    if false {
        v3 = append(v3, ("hello" + ("\n\t" + "hello")))
        v33 := GenS4{F0: "hello", F1: 2, F2: 1000}
        var v34 bool = true
        _ = v33
        _ = v34
    } else {
        fmt.Println(v32)
    }
    _ = len(v4)
    v5 = (1.5 * (float32(1.5) / 1.0))
    for fi35 := 0; fi35 < 20; fi35++ {
        if false {
            if true {
                switch (fi35 >> ((v26) & 63)) {
                case -1:
                    for fi36 := 0; fi36 < 16; fi36++ {
                        v37 := GenS5{F0: 201, F1: 1, F2: 128}
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        _ = fi36
                        _ = v37
                    }
                case 42:
                    fmt.Println(fi35, v32, v1)
                default:
                    var v38 float32 = (1e10 + 3.14159)
                    _ = v38
                }
            }
            var v39 string = "\n\t"
            v40 := GenS6{F0: 322}
            _ = v39
            _ = v40
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
