package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint8 }
func genFn0(p0 bool, p1 int, p2 uint) bool {
    for fi5 := 0; fi5 < 13; fi5++ {
        switch "\n\t" {
        case "\n\t":
            for fi6 := 0; fi6 < 20; fi6++ {
                p1 += (fi5 - (fi5 * (p1 >> ((p2) & 63))))
                _ = fi6
            }
        default:
            v7 := []int{(^(p1)), (p1 / (((p1 - 32767)) | 1))}
            _ = v7
        }
        _ = fi5
    }
    p2 -= 16
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    return (!(!p0))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 uint = 0
    v4 := make(map[int]bool)
    v3 += v3
    v3 = 1
    var v8 int = 0
    if (true && (!true)) {
        switch ("\n\t" + ("" + "a")) {
        case "hello":
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: 418, F1: uint64(v3)}
                v8 += 4294967295
                _ = fi9
                _ = v10
            }
        case "":
            var v11 int8 = 2
            _ = v11
        case "a":
            v3 -= 18446744073709551615
        default:
            v8 -= v8
        }
    }
    var v12 uint = (v3 * (^(v3)))
    switch v8 {
    case 663:
        fmt.Println(v3)
    case -264:
        for fi13 := 0; fi13 < 20; fi13++ {
            switch "\n\t" {
            case "hello":
                v4[-580] = (1 < 32767)
            case "a":
                if ((0 < -1) || true) {
                    switch fi13 {
                    case 1:
                        var v14 float64 = (-2.25 / (1e10 - 1.5))
                        _ = v14
                    }
                    var v15 int16 = 234
                    v16 := v2["foo bar"]
                    _ = v15
                    _ = v16
                } else {
                    fmt.Println(fi13)
                    v3 -= v12
                    if (!((0.1 > float64(-2.25)) || false)) {
                        switch fi13 {
                        case 86:
                            v8 = fi13
                        case 42:
                            v1[""] = ((("a" + "\n\t") + ((("" + "\n\t") + "") + ("" + "a"))) + "\n\t")
                        default:
                            v3 -= (v3 % ((928) | 1))
                        }
                        v8 += (v8 % ((0) | 1))
                    } else {
                        _ = len(v4)
                    }
                }
            case "":
                var v17 int32 = 2147483647
                _ = v17
            }
            v18 := GenS2{F0: 4294967295, F1: 255, F2: -1}
            _ = fi13
            _ = v18
        }
    default:
        v12 -= 0
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v8, fi19, v3)
        _ = fi19
    }
    var v20 uint = v12
    v21 := GenS3{F0: 42}
    v22 := GenS4{F0: 989, F1: 9223372036854775807, F2: 42}
    var v23 uint16 = 32767
    v20 = 255
    v23 -= v23
    for fi24 := 0; fi24 < 2; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            v3 -= 1
            _ = fi25
        }
        switch -1 {
        case 4294967295:
            var v26 int = 32768
            _ = v26
        default:
            v23 -= v23
        }
        var v27 int = 584
        _ = fi24
        _ = v27
    }
    var v28 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v8
}
