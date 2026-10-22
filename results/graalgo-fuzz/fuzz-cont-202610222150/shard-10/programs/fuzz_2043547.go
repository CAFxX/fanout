package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint8 }
func genFn0() uint32 {
    var v6 bool = true
    var v7 bool = ("foo bar" == (("" + ("foo bar" + (("foo bar" + "foo bar") + "\n\t"))) + "hello"))
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{fi8, fi8, (-(fi8)), (fi8 % ((fi8) | 1))}
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint8 = 42
            switch fi10 {
            case 128:
                var v12 uint64 = 0
                _ = v12
            case -2147483648:
                _ = len(v9)
            }
            _ = fi10
            _ = v11
        }
        v6 = v7
        _ = fi8
        _ = v9
    }
    _ = v6
    _ = v7
    return 459
}

func main() {
    v1 := []string{}
    var v2 float32 = (0.1 * 0.1)
    var v3 uint8 = 127
    var v4 string = ((("foo bar" + "hello") + ("\n\t" + "\n\t")) + "")
    var v5 int64 = 602
    v1 = append(v1, v4)
    var v13 uint64 = 943
    var v14 int = 128
    v3 = (v3 ^ (v3 << ((0) & 63)))
    if true {
        v1 = append(v1, "hello")
        v15 := GenS1{F0: 42, F1: 42, F2: 2147483648}
        _ = v15
    } else {
        if (v2 < ((float32(((0.1 / 3.14159) - 1e10)) / 1.0) / 1.5)) {
            if true {
                v16 := GenS1{F0: 416, F1: 787, F2: 798}
                v17 := GenS2{F0: v3, F1: 2}
                if (("foo bar" + v4) < v4) {
                    v2 -= (0.1 - v2)
                    var v18 uint16 = 69
                    _ = v18
                } else {
                    v1 = append(v1, v4)
                    var v19 float64 = (0.1 * 3.14159)
                    _ = v19
                }
                _ = v16
                _ = v17
            } else {
                var v20 uint64 = 1
                v2 = (v2 * v2)
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            for fi22 := 0; fi22 < 16; fi22++ {
                if (65535 <= 832) {
                    v14 = fi22
                } else {
                    v23 := GenS1{F0: 600, F1: 0, F2: 439}
                    var v24 int = (fi21 ^ 676)
                    fmt.Println(fi22)
                    _ = v23
                    _ = v24
                }
                var v25 int = (fi22 + fi22)
                fmt.Println(v13, v5, v25)
                _ = fi22
                _ = v25
            }
            _ = fi21
        }
        v26 := GenS3{F0: 191}
        _ = v26
    }
    switch "\n\t" {
    case "\n\t":
        v27 := GenS4{F0: genFn0()}
        _ = v27
    case "hello":
        v14 -= (-(v14))
    case "a":
        switch (v14 / (((v14 | v14)) | 1)) {
        case 4294967295:
            var v28 int16 = 1
            _ = v28
        }
    default:
        v29 := GenS5{F0: 877, F1: 32767, F2: v13}
        _ = v29
    }
    var v30 bool = true
    var v31 string = "foo bar"
    v1 = append(v1, v31)
    if len(v1) > 0 {
        v1[0] = ("foo bar" + v31)
    }
    v1 = append(v1, ((v4 + v4) + (v4 + v31)))
    fmt.Println(v13, v3, v3)
    _ = len(v1)
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 int16 = 1000
        v14 = fi32
        _ = fi32
        _ = v33
    }
    v34 := GenS6{F0: uint8(v13)}
    v35 := GenS4{F0: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
