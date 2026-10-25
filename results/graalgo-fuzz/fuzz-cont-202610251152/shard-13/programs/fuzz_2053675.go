package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint8) uint16 {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 1, F1: ("foo bar" + "a"), F2: false}
        p0 -= 148
        var v7 float32 = (0.0 - float32(((0.1 / 3.14159) - 1e10)))
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int8 = 42
    p1 += p1
    _ = p0
    _ = p1
    _ = v8
    return 324
}
func genFn1(p0 uint32, p1 int, p2 uint32) int16 {
    var v9 uint16 = 623
    fmt.Println(p2, p0)
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(p1, p2, p2)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 855
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{1}
    var v3 float32 = -2.25
    var v4 int64 = 234
    v3 = v3
    var v11 bool = false
    v11 = v11
    fmt.Println(v11)
    var v12 float32 = float32(float64((v3 + (((1e10 / v3) / v3) - ((((1e10 + 1.5) + v3) + v3) + v3)))))
    v1[4294967295] = false
    fmt.Println(v3, v4)
    if v11 {
        for fi13 := 0; fi13 < 2; fi13++ {
            v12 += 1e10
            v14 := GenS2{F0: 42, F1: 1}
            var v15 uint32 = 4294967295
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                v3 = -2.25
                _ = fi17
            }
            fmt.Println(v4, v4)
            _ = fi16
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            if (((!((v3 + v12) >= (1.5 + v12))) && false) && v11) {
                switch "a" {
                case "\n\t":
                    if (18446744073709551615 > 4294967295) {
                        if (!(v11 || (158 == 128))) {
                            for fi19 := 0; fi19 < 13; fi19++ {
                                var v20 string = "a"
                                switch (-(fi19)) {
                                case 4294967295:
                                    v12 = ((float32((3.14159 - (0.1 - 1.5))) / (v12 + 1e10)) + v3)
                                case -574:
                                    v20 = v20
                                }
                                fmt.Println(fi18, fi19, v12)
                                var v21 uint16 = 295
                                delete(v1, 1)
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            fmt.Println(fi18, v3)
                        }
                        v3 += ((v12 * 0.0) / (float32(0.0) / v3))
                        if v11 {
                            v22 := GenS1{F0: 1, F1: "foo bar", F2: v11}
                            var v23 uint16 = genFn0(2, 37)
                            for fi24 := 0; fi24 < 6; fi24++ {
                                var v25 int8 = -128
                                var v26 int32 = -1
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            _ = v22
                            _ = v23
                        }
                    }
                }
            }
            v3 -= 0.1
            _ = fi18
        }
    }
    switch "a" {
    case "":
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 int = (fi27 + fi27)
            switch v28 {
            case 32767:
                var v29 uint = 151
                _ = v29
            default:
                switch "\n\t" {
                case "\n\t":
                    if (!v11) {
                        fmt.Println(v11, fi27)
                    }
                case "hello":
                    v1[4294967295] = v11
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            for fi30 := 0; fi30 < 9; fi30++ {
                v2 = append(v2, 2147483647)
                switch fi27 {
                case 1:
                    v31 := GenS3{F0: 309}
                    _ = v31
                }
                switch fi27 {
                case 128:
                    fmt.Println(fi30)
                default:
                    var v32 int8 = 100
                    _ = v32
                }
                _ = fi30
            }
            _ = fi27
            _ = v28
        }
    case "a":
        v12 -= (-2.25 + (v12 * v12))
    case "\n\t":
        switch -926 {
        case -1:
            v33 := GenS2{F0: 127, F1: -2147483648}
            _ = v33
        }
    }
    var v34 uint64 = 1
    v3 = v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
