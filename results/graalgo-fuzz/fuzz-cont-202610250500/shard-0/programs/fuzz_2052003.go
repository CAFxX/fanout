package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
func genFn0(p0 int8) uint32 {
    v5 := []int{725, -2147483648}
    var v6 uint16 = 0
    _ = p0
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 string, p1 int16, p2 int16) bool {
    fmt.Println(p1)
    v7 := []int{42, 32767, 127, 121}
    var v8 uint64 = 53
    var v9 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    v3 := make(map[uint32]bool)
    var v4 uint16 = 2
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 string = "\n\t"
        var v12 uint32 = 65535
        fmt.Println(fi10)
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{}
    var v14 int64 = 9223372036854775807
    v15 := GenS1{F0: 0, F1: v4}
    var v16 int64 = (-(v14))
    for fi17 := 0; fi17 < 3; fi17++ {
        v4 -= (v4 / (((^(v4))) | 1))
        v2 = 42
        for fi18 := 0; fi18 < 9; fi18++ {
            v16 += v16
            switch fi17 {
            case 231:
                var v19 uint = uint(v4)
                _ = v19
            case 0:
                v20 := GenS2{F0: 296, F1: -1, F2: 0}
                _ = v20
            default:
                if true {
                    v4 = (v4 & v4)
                    var v21 string = "a"
                    var v22 int16 = 42
                    _ = v21
                    _ = v22
                }
            }
            _ = fi18
        }
        _ = fi17
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            v4 -= (v4 & 2)
            _ = fi24
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint16 = v4
        if (float32(0.0) > (((1.5 - 0.1) - 0.1) / -2.25)) {
            if false {
                if (!(("a" + "a") > ("a" + ("a" + ("hello" + ((("" + ("" + ("" + "hello"))) + "foo bar") + (("\n\t" + ((("foo bar" + ("foo bar" + "hello")) + "a") + "")) + "hello"))))))) {
                    for fi27 := 0; fi27 < 11; fi27++ {
                        if true {
                            if false {
                                if false {
                                    var v28 int16 = 0
                                    _ = v28
                                }
                                v26 -= v4
                            }
                            v29 := GenS1{F0: 1, F1: v4}
                            switch (^(fi27)) {
                            case -539:
                                fmt.Println(v14, v16)
                            case -998:
                                fmt.Println(fi27)
                            }
                            _ = v29
                        }
                        _ = fi27
                    }
                }
                v26 += (v26 << ((uint(v2)) & 63))
            } else {
                v30 := GenS3{F0: 593, F1: 1, F2: (v16 - (v16 * (v14 >> ((945) & 63))))}
                v31 := GenS4{F0: (^(fi25))}
                _ = v30
                _ = v31
            }
            v32 := GenS5{F0: (-(v16))}
            var v33 uint16 = v26
            _ = v32
            _ = v33
        } else {
            v34 := GenS3{F0: -1, F1: 47, F2: v14}
            var v35 int16 = -1
            var v36 uint8 = 116
            _ = v34
            _ = v35
            _ = v36
        }
        _ = fi25
        _ = v26
    }
    var v37 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
