package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 uint, p1 int64) uint {
    if true {
        v5 := GenS1{F0: p1, F1: -395, F2: false}
        p1 = p1
        p1 += (^(p1))
        _ = v5
    } else {
        var v6 uint = (p0 & p0)
        _ = v6
    }
    v7 := GenS2{F0: 127, F1: (p0 - 321), F2: 793}
    v8 := GenS3{F0: (p0 * (p0 % (((p0 & (^(p0)))) | 1))), F1: 683}
    v9 := []int{}
    var v10 int32 = -222
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (p0 << ((p0) & 63))
}

func main() {
    var v1 uint64 = 462
    v2 := []string{(("a" + "hello") + ("a" + (("a" + "hello") + "a")))}
    var v3 uint32 = 957
    v4 := []int{}
    v4 = append(v4, 722)
    fmt.Println(v1)
    v11 := GenS3{F0: 0, F1: 42}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 += 0
    v1 -= (^(v1))
    fmt.Println(v1)
    v1 -= 18446744073709551615
    switch 44 {
    case 297:
        var v12 int64 = -554
        _ = v12
    case 0:
        if ((float32(3.14159) >= 0.0) || true) {
            var v13 uint8 = 255
            fmt.Println(v13, v1)
            _ = v13
        } else {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            if true {
                fmt.Println(v3, v1)
                v14 := GenS1{F0: 1, F1: 127, F2: false}
                if true {
                    var v15 bool = true
                    fmt.Println(v3, v1, v1)
                    _ = v15
                } else {
                    v3 = (v3 | (v3 / (((v3 / (((v3 | v3)) | 1))) | 1)))
                }
                _ = v14
            } else {
                fmt.Println(v1)
            }
        }
    default:
        switch "" {
        case "a":
            for fi16 := 0; fi16 < 2; fi16++ {
                v3 = (v3 + (^(v3)))
                if false {
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 int64 = -819
                        _ = fi17
                        _ = v18
                    }
                }
                _ = fi16
            }
        case "hello":
            switch -780 {
            case 128:
                var v19 string = "foo bar"
                _ = v19
            case -291:
                v1 = 561
            default:
                v1 -= v1
            }
        case "":
            v3 -= v3
        default:
            _ = len(v2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
