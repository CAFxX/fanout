package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint }
func genFn0(p0 uint8) uint {
    p0 -= p0
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(fi6, p0)
        _ = fi6
    }
    _ = p0
    return 1
}
func genFn1(p0 int16) uint32 {
    v7 := GenS1{F0: "hello"}
    switch (((("hello" + "") + "") + ((((((("a" + "\n\t") + "a") + ((("a" + "") + ("" + "\n\t")) + (("\n\t" + ("hello" + "\n\t")) + ""))) + "a") + (("foo bar" + ("\n\t" + "\n\t")) + "\n\t")) + "") + ("foo bar" + ("foo bar" + "\n\t")))) + "hello") {
    case "":
        var v8 uint16 = 1
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint64 = 444
        _ = fi9
        _ = v10
    }
    if (!(!false)) {
        v11 := GenS2{F0: 255, F1: 4294967295}
        _ = v11
    } else {
        p0 += (p0 | 0)
    }
    _ = p0
    _ = v7
    return 252
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := []uint32{1, 2147483647, 2147483647}
    var v5 uint16 = 835
    v12 := GenS1{F0: "foo bar"}
    var v13 uint8 = 102
    var v14 int32 = int32(v13)
    v15 := GenS1{F0: ("hello" + "foo bar")}
    v14 += -997
    for fi16 := 0; fi16 < 17; fi16++ {
        if false {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 int16 = int16(v14)
                switch fi17 {
                case 42:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    v18 = v18
                }
                v13 = 0
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                if false {
                    var v20 int32 = (v14 << ((0) & 63))
                    fmt.Println(v20, fi19, v14)
                    _ = v20
                } else {
                    var v21 uint32 = 4294967295
                    v3[v21] = fi19
                    var v22 int64 = 841
                    _ = v21
                    _ = v22
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi19
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                v5 += v5
                _ = fi23
            }
        }
        _ = fi16
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := v3[223]
            if false {
                switch (("hello" + "\n\t") + ("foo bar" + (("a" + "") + (("foo bar" + (("foo bar" + "") + "")) + (("foo bar" + "\n\t") + "hello"))))) {
                case "":
                    fmt.Println(fi24, v5, v14)
                case "\n\t":
                    v27 := GenS3{F0: v5, F1: 2147483647, F2: 1}
                    _ = v27
                case "a":
                    v14 = v14
                default:
                    if (!false) {
                        var v28 int = (fi24 ^ fi25)
                        fmt.Println(fi25)
                        _ = v28
                    } else {
                        v29 := GenS2{F0: (v13 * v13), F1: uint(v5)}
                        if ((4294967295 > genFn1(745)) || false) {
                            v5 = (^(v5))
                            v14 = v14
                        }
                        _ = v29
                    }
                }
                _ = len(v4)
            }
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    var v30 int64 = int64(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
