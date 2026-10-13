package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint }
func genFn0() string {
    v4 := GenS1{F0: 9223372036854775807, F1: -578}
    var v5 uint32 = 2147483647
    if (42 > 0) {
        fmt.Println(v5)
        v6 := GenS1{F0: -9223372036854775808, F1: -328}
        _ = v6
    }
    var v7 bool = (1 == 1)
    var v8 int8 = -1
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return "hello"
}
func genFn1(p0 string) string {
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int8 = 42
        switch fi9 {
        case -638:
            if (!(!(1 > 2147483647))) {
                v11 := []int{fi9, 1, fi9, -889}
                if false {
                    var v12 uint8 = 128
                    v12 -= 167
                    _ = v12
                } else {
                    fmt.Println(v10)
                }
                _ = v11
            } else {
                v13 := []int{32767, 128}
                fmt.Println(v10, v10, fi9)
                fmt.Println(v10)
                _ = v13
            }
        default:
            if (true || ((fi9 & fi9) == fi9)) {
                var v14 uint64 = 42
                _ = v14
            }
        }
        _ = fi9
        _ = v10
    }
    v15 := []int{}
    if (!true) {
        var v16 int = -1
        _ = v16
    }
    _ = p0
    _ = v15
    return p0
}

func main() {
    var v1 int32 = 1
    var v2 int8 = -29
    var v3 uint32 = 65535
    if false {
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                v1 -= v1
                v19 := GenS1{F0: -1, F1: (^(v1))}
                _ = fi18
                _ = v19
            }
            switch ("foo bar" + "hello") {
            case "a":
                var v20 float32 = (1.5 + 1e10)
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 9; fi21++ {
                    fmt.Println(fi17)
                    var v22 int8 = (-(v2))
                    _ = fi21
                    _ = v22
                }
            }
            _ = fi17
        }
    }
    fmt.Println(v1)
    for fi23 := 0; fi23 < 7; fi23++ {
        v24 := GenS1{F0: 0, F1: 1}
        v25 := []int{(-(fi23))}
        v26 := GenS1{F0: int64(v2), F1: (v1 / (((v1 * v1)) | 1))}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v2 -= 42
        v28 := GenS1{F0: 0, F1: -878}
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 int32 = (v1 - int32(fi29))
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: 419, F1: (v1 % ((v1) | 1))}
    v3 -= 42
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 uint32 = (v3 >> ((2) & 63))
        _ = fi32
        _ = v33
    }
    if true {
        var v34 string = ""
        fmt.Println(v1, v3)
        if true {
            if false {
                v2 = 89
                v35 := GenS2{F0: 32767, F1: (v3 - v3), F2: 4294967295}
                fmt.Println(v3)
                _ = v35
            } else {
                var v36 bool = false
                for fi37 := 0; fi37 < 8; fi37++ {
                    var v38 int64 = 1
                    _ = fi37
                    _ = v38
                }
                v39 := GenS1{F0: 9223372036854775807, F1: -2147483648}
                _ = v36
                _ = v39
            }
        }
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
