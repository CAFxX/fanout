package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 bool, p1 uint8) int16 {
    if (!p0) {
        p0 = p0
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{(fi6 >> ((240) & 63))}
        var v8 uint64 = 1
        fmt.Println(p1, fi6, v8)
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: p1}
    v10 := GenS1{F0: 42, F1: p1}
    v11 := []int{-98, 127, 32767}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return -32768
}

func main() {
    v1 := []string{"", ""}
    var v2 int16 = 42
    var v3 int16 = v2
    var v4 uint = 905
    var v5 float64 = (1.5 / 1.5)
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint32 = 134
        _ = fi12
        _ = v13
    }
    v2 = genFn0((0 <= -1), 125)
    var v14 int = 32768
    v14 += -833
    if (v2 > 32767) {
        if len(v1) > 0 {
            v1[0] = ("hello" + "hello")
        }
        v4 += (^(v4))
    }
    if true {
        var v15 uint8 = 144
        for fi16 := 0; fi16 < 16; fi16++ {
            if len(v1) > 0 {
                v1[0] = ((("\n\t" + "foo bar") + "a") + "\n\t")
            }
            if (!(!(29 > 1))) {
                var v17 uint32 = 799
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                _ = v17
            }
            _ = fi16
        }
        var v18 uint16 = 65535
        _ = v15
        _ = v18
    }
    v19 := GenS1{F0: 2, F1: 127}
    v14 = (v14 * v14)
    switch 1 {
    case 4294967295:
        for fi20 := 0; fi20 < 13; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                if (true || (!(225 != 2))) {
                    fmt.Println(v3, v3)
                    var v22 uint32 = 4294967295
                    _ = v22
                } else {
                    if ("" != "a") {
                        var v23 int8 = -1
                        _ = v23
                    }
                    if (4294967295 == 42) {
                        var v24 uint = v4
                        _ = v24
                    }
                    var v25 uint = v4
                    v26 := GenS2{F0: genFn0(("a" != "\n\t"), 210)}
                    _ = v25
                    _ = v26
                }
                if false {
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        v27 := GenS1{F0: 65535, F1: 2}
                        _ = v27
                    case 0:
                        v3 = v3
                        v28 := GenS1{F0: 764, F1: 1}
                        _ = v28
                    }
                }
                _ = fi21
            }
            fmt.Println(v5, v4)
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 string = "\n\t"
                fmt.Println(v5)
                _ = fi29
                _ = v30
            }
            _ = fi20
        }
    case -811:
        v31 := GenS3{F0: 7}
        _ = v31
    }
    var v32 uint32 = 622
    var v33 uint32 = v32
    var v34 int32 = -1
    v1 = append(v1, "foo bar")
    fmt.Println(v2, v14, v33)
    v35 := GenS1{F0: 270, F1: 42}
    v36 := GenS1{F0: 79, F1: 127}
    if false {
        v5 += (v5 - v5)
        switch (v14 ^ v14) {
        case 9223372036854775807:
            if len(v1) > 0 {
                v1[0] = ((("" + "") + (("" + (("\n\t" + "") + "foo bar")) + "foo bar")) + ("" + "foo bar"))
            }
        case 128:
            var v37 uint16 = 887
            _ = v37
        default:
            v38 := GenS4{F0: v34}
            _ = v38
        }
    } else {
        if true {
            v33 = v32
        }
        v14 += int(v33)
    }
    var v39 string = ("hello" + "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
