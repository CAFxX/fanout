package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint64, p2 uint8) int8 {
    fmt.Println(p0, p0, p1)
    var v4 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int64, p1 uint) uint32 {
    var v5 uint = 0
    for fi6 := 0; fi6 < 20; fi6++ {
        if ((false || false) && (true || true)) {
            fmt.Println(p1)
            fmt.Println(v5, p1, fi6)
            fmt.Println(v5, p0)
        }
        p1 = (v5 + p1)
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint64 = 916
            p1 += (v5 ^ v5)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    fmt.Println(p0, p1)
    for fi9 := 0; fi9 < 15; fi9++ {
        switch fi9 {
        case 42:
            if false {
                v10 := GenS1{F0: -97}
                _ = v10
            } else {
                var v11 float32 = (((0.0 / 0.1) + 0.1) * -2.25)
                _ = v11
            }
        case 592:
            v12 := GenS1{F0: 7}
            _ = v12
        }
        _ = fi9
    }
    v13 := GenS2{F0: true, F1: 562}
    _ = p0
    _ = p1
    _ = v13
    _ = v5
    return 883
}

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 42
    v3 := make(map[uint32]string)
    if (0 == -200) {
        var v14 bool = false
        v15 := GenS3{F0: 0, F1: 733}
        v16 := v3[(^(v2))]
        _ = v14
        _ = v15
        _ = v16
    } else {
        v1 -= 255
        v1 = v1
    }
    v17 := GenS2{F0: true, F1: 1}
    var v18 bool = false
    v19 := v3[4294967295]
    var v20 int8 = 91
    v18 = false
    v19 = "\n\t"
    v1 = (^(v1))
    for fi21 := 0; fi21 < 15; fi21++ {
        fmt.Println(fi21, v19)
        v20 = 0
        _ = fi21
    }
    v22 := []int{-527}
    v22 = append(v22, 32767)
    if v18 {
        v23 := GenS3{F0: 0, F1: 962}
        _ = v23
    } else {
        if (v18 && v18) {
            for fi24 := 0; fi24 < 10; fi24++ {
                v2 -= 4294967295
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 int32 = 453
                    v19 = (("a" + v19) + "a")
                    _ = fi25
                    _ = v26
                }
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = fi24
            }
            fmt.Println(v19, v20, v20)
            switch -9223372036854775808 {
            case -1:
                v2 = (^(v2))
            case 42:
                switch 0 {
                case -602:
                    for fi27 := 0; fi27 < 1; fi27++ {
                        if (365 >= 32767) {
                            v1 = (^(v1))
                            v28 := v3[42]
                            var v29 string = v28
                            _ = v28
                            _ = v29
                        }
                        _ = fi27
                    }
                default:
                    fmt.Println(v19)
                }
            default:
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 uint = 42
                    v18 = ((-(fi30)) > (fi30 / ((fi30) | 1)))
                    _ = fi30
                    _ = v31
                }
            }
        } else {
            var v32 uint32 = 42
            v33 := GenS2{F0: (!(((((v19 + "hello") + "foo bar") + "hello") + (v19 + v19)) <= ("foo bar" + (v19 + v19)))), F1: v1}
            v34 := GenS4{F0: 0}
            _ = v32
            _ = v33
            _ = v34
        }
    }
    delete(v3, (v2 + genFn1(9223372036854775807, uint(v2))))
    var v35 int32 = 952
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v35
}
