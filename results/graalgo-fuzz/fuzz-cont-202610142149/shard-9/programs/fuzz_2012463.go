package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
func genFn0(p0 uint8) int16 {
    fmt.Println(p0)
    p0 = 164
    var v4 string = "\n\t"
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int32 = 2147483647
        v4 = (v4 + "foo bar")
        switch v4 {
        case "\n\t":
            if false {
                var v7 uint64 = 578
                v8 := GenS1{F0: -32768, F1: 42, F2: (p0 + 127)}
                _ = v7
                _ = v8
            } else {
                if true {
                    p0 -= (p0 + (p0 | p0))
                    v4 = ((v4 + v4) + ("foo bar" + "\n\t"))
                    for fi9 := 0; fi9 < 10; fi9++ {
                        v6 = (v6 % (((v6 | (v6 * v6))) | 1))
                        switch fi9 {
                        case 2147483647:
                            var v10 float32 = 1.5
                            _ = v10
                        case 127:
                            v11 := GenS2{F0: 1, F1: 42, F2: 0}
                            _ = v11
                        }
                        _ = fi9
                    }
                }
                v6 += v6
            }
        case "hello":
            v12 := GenS3{F0: v6}
            _ = v12
        default:
            p0 += 172
        }
        _ = fi5
        _ = v6
    }
    v4 = ("" + v4)
    _ = p0
    _ = v4
    return 32767
}
func genFn1(p0 int64, p1 int8) uint8 {
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int32 = -2147483648
        v15 := GenS4{F0: 0}
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v16 := []int{127}
    v3[2147483647] = 2147483648
    for fi17 := 0; fi17 < 19; fi17++ {
        fmt.Println(fi17, v1)
        v18 := GenS3{F0: 1}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            fmt.Println(fi19)
            if false {
                _ = len(v16)
                v21 := v2[((((((("a" + "hello") + "hello") + "foo bar") + ("foo bar" + ("\n\t" + (((("a" + "hello") + "hello") + "") + ("a" + ("\n\t" + "foo bar")))))) + "foo bar") + "a") + "\n\t")]
                v21 = false
                _ = v21
            } else {
                var v22 bool = (!false)
                _ = v22
            }
            if ((!true) || (53 == 32767)) {
                var v23 int16 = 27
                fmt.Println(fi19, fi20, fi20)
                _ = v23
            }
            _ = fi20
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            if ((!false) && ((!true) || true)) {
                v25 := GenS5{F0: (fi24 ^ (fi19 >> ((260) & 63)))}
                v1 -= (v1 / (((v1 << ((0) & 63))) | 1))
                _ = v25
            }
            v26 := GenS1{F0: 32767, F1: -2147483648, F2: genFn1((v1 / (((^(v1))) | 1)), -128)}
            var v27 uint = 4294967295
            _ = fi24
            _ = v26
            _ = v27
        }
        v1 += v1
        _ = fi19
    }
    if len(v16) > 0 {
        v16[0] = -464
    }
    v1 -= (v1 | v1)
    for fi28 := 0; fi28 < 17; fi28++ {
        if false {
            v1 += v1
            v1 += (^(v1))
            var v29 int8 = -76
            _ = v29
        }
        _ = fi28
    }
    delete(v3, 434)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
