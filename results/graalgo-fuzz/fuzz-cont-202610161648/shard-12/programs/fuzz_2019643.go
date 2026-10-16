package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint32) uint {
    switch ("a" + "foo bar") {
    case "foo bar":
        var v6 uint64 = 18446744073709551615
        _ = v6
    default:
        var v7 float32 = ((float32(0.1) - (0.0 - 1e10)) * 0.0)
        _ = v7
    }
    var v8 int32 = 2147483647
    var v9 int = -134
    var v10 uint8 = 42
    p0 -= p0
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    v3 := []uint32{41, 2147483648}
    v4 := []int{0, -1, 0, 1}
    v5 := make(map[string]uint32)
    var v11 uint32 = 469
    v12 := GenS1{F0: 0, F1: "foo bar"}
    v11 = 65535
    if ((1.5 + -2.25) >= (-2.25 * (1e10 / 3.14159))) {
        v11 -= (v11 + (v11 * (^(v11))))
    }
    if ((4294967295 == genFn0(v11)) || (-2.25 >= 1.5)) {
        v13 := GenS2{F0: 111, F1: 255, F2: (^(v11))}
        v14 := GenS2{F0: 41, F1: uint(v11), F2: 65535}
        _ = v13
        _ = v14
    } else {
        switch 994 {
        case 1:
            for fi15 := 0; fi15 < 1; fi15++ {
                if len(v3) > 0 {
                    v3[0] = v11
                }
                _ = fi15
            }
        case 32768:
            var v16 uint32 = (v11 % ((v11) | 1))
            _ = v16
        case 257:
            for fi17 := 0; fi17 < 13; fi17++ {
                fmt.Println(fi17)
                _ = fi17
            }
        default:
            switch 127 {
            case 4294967295:
                v11 += v11
            }
        }
    }
    switch "hello" {
    case "":
        for fi18 := 0; fi18 < 12; fi18++ {
            for fi19 := 0; fi19 < 8; fi19++ {
                if false {
                    fmt.Println(v11, fi19, v11)
                    v20 := v1[""]
                    var v21 string = "\n\t"
                    _ = v20
                    _ = v21
                }
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi19
            }
            if false {
                fmt.Println(fi18, v11)
                var v22 int = fi18
                _ = v22
            }
            _ = fi18
        }
    case "foo bar":
        v23 := GenS3{F0: 2147483647, F1: 2}
        _ = v23
    case "hello":
        v24 := GenS2{F0: 223, F1: genFn0((v11 * v11)), F2: 0}
        _ = v24
    }
    delete(v1, "a")
    v25 := GenS1{F0: 32767, F1: "hello"}
    v26 := GenS3{F0: -102, F1: 0}
    var v27 int16 = 1
    for fi28 := 0; fi28 < 1; fi28++ {
        v27 = 0
        _ = fi28
    }
    var v29 int = -879
    for fi30 := 0; fi30 < 4; fi30++ {
        v31 := GenS2{F0: 42, F1: 4294967295, F2: (v11 & v11)}
        _ = fi30
        _ = v31
    }
    v32 := GenS4{F0: v11, F1: v11}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
