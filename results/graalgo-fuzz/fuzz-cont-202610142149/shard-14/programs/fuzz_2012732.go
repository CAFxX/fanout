package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int8; F1 int32 }

func main() {
    var v1 uint = 104
    var v2 float32 = 3.14159
    var v3 uint64 = 658
    var v4 uint8 = 42
    for fi5 := 0; fi5 < 4; fi5++ {
        switch 127 {
        case 127:
            if (true || true) {
                switch fi5 {
                case -930:
                    var v6 uint64 = v3
                    _ = v6
                }
            } else {
                for fi7 := 0; fi7 < 15; fi7++ {
                    v8 := []int{fi7, fi7}
                    _ = fi7
                    _ = v8
                }
            }
        case -9223372036854775808:
            v2 = -2.25
        }
        var v9 bool = false
        if (!(32767 != -32768)) {
            fmt.Println(v4)
        } else {
            if (!(v9 && true)) {
                if v9 {
                    if (126 >= uint8(v1)) {
                        v10 := []int{(fi5 / ((-773) | 1)), 550, (^(fi5))}
                        var v11 int32 = -337
                        _ = v10
                        _ = v11
                    } else {
                        var v12 uint16 = uint16(v3)
                        _ = v12
                    }
                }
            } else {
                v13 := []int{(fi5 - fi5), (fi5 >> ((v1) & 63)), fi5}
                _ = v13
            }
            var v14 bool = v9
            var v15 uint16 = 1
            _ = v14
            _ = v15
        }
        _ = fi5
        _ = v9
    }
    v1 += 2
    var v16 float32 = v2
    v16 = 0.1
    if true {
        if (v2 >= (v2 / (((0.0 - (v2 + (float32(1e10) - (0.0 + 1.5)))) + -2.25) * 3.14159))) {
            v2 += -2.25
        } else {
            if (false && ((v1 % (((v1 ^ v1)) | 1)) > 615)) {
                if (v2 < v2) {
                    if false {
                        fmt.Println(v4, v1)
                    }
                }
            }
            v17 := []int{32767, 2147483647, -334, 590}
            _ = v17
        }
        var v18 int16 = 1000
        var v19 bool = true
        _ = v18
        _ = v19
    }
    var v20 uint8 = (v4 | (v4 + 1))
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS1{F0: 810, F1: 0}
            _ = fi22
            _ = v23
        }
        if true {
            v16 -= v16
            v24 := GenS2{F0: (uint64(v20) > v3), F1: 806}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 16; fi25++ {
                v3 -= 4294967295
                if (948 > -592) {
                    v26 := GenS2{F0: ((v20 / ((v4) | 1)) >= v4), F1: 2147483647}
                    var v27 string = ""
                    _ = v26
                    _ = v27
                }
                v28 := []int{fi21, fi25}
                var v29 int8 = 1
                _ = fi25
                _ = v28
                _ = v29
            }
        }
        v16 = v16
        _ = fi21
    }
    v4 -= (v4 * 1)
    v2 = float32(0.0)
    if false {
        for fi30 := 0; fi30 < 10; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                v32 := GenS3{F0: 94, F1: -1}
                _ = fi31
                _ = v32
            }
            v33 := []int{127, fi30}
            _ = fi30
            _ = v33
        }
        var v34 string = (((("foo bar" + "a") + (("a" + "foo bar") + "foo bar")) + "hello") + "")
        v35 := []int{2147483647, 287, 994, 127}
        _ = v34
        _ = v35
    } else {
        v36 := []int{}
        v37 := GenS1{F0: 65535, F1: 4294967295}
        v2 = (v2 / 1e10)
        _ = v36
        _ = v37
    }
    v38 := GenS4{F0: int8(v20), F1: 899}
    for fi39 := 0; fi39 < 18; fi39++ {
        v2 = (v16 + 3.14159)
        fmt.Println(v3)
        var v40 uint32 = 2147483648
        _ = fi39
        _ = v40
    }
    var v41 float64 = 1.5
    v1 = (v1 % ((v1) | 1))
    var v42 string = "foo bar"
    fmt.Println(v41, v16, v3)
    v16 -= ((float32(v41) * ((3.14159 + v2) + v16)) + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v42
}
