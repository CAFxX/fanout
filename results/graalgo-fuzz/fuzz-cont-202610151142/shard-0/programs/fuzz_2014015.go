package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int8, p1 uint32) int16 {
    p0 = (^(p0))
    for fi6 := 0; fi6 < 4; fi6++ {
        switch fi6 {
        case -1:
            if (2 < uint64(p1)) {
                p0 = (p0 / (((p0 % (((p0 * p0)) | 1))) | 1))
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return -32768
}
func genFn1(p0 int8, p1 string) uint {
    if false {
        switch -1 {
        case 127:
            switch 0 {
            case 2147483647:
                fmt.Println(p1, p0)
            case -8:
                if (p0 >= p0) {
                    p1 = p1
                    for fi7 := 0; fi7 < 18; fi7++ {
                        for fi8 := 0; fi8 < 3; fi8++ {
                            p1 = "foo bar"
                            v9 := []int{fi7, (fi7 << ((65535) & 63)), fi7}
                            _ = len(v9)
                            _ = fi8
                            _ = v9
                        }
                        _ = fi7
                    }
                } else {
                    p1 = p1
                    var v10 int = 128
                    _ = v10
                }
            }
        case -208:
            v11 := []int{1}
            _ = v11
        case 32768:
            var v12 uint = 0
            _ = v12
        }
        p1 = p1
        switch "\n\t" {
        case "foo bar":
            var v13 int16 = -320
            _ = v13
        case "\n\t":
            switch ("foo bar" + "hello") {
            case "hello":
                v14 := GenS1{F0: 42, F1: 4294967295, F2: 2}
                _ = v14
            case "\n\t":
                p1 = ""
            case "foo bar":
                var v15 int32 = 1
                _ = v15
            }
        default:
            var v16 int8 = (p0 >> ((255) & 63))
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := []int{}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(p1)
            _ = fi19
        }
        fmt.Println(p1)
    }
    v20 := GenS1{F0: 2147483647, F1: 0, F2: 1}
    fmt.Println(p1, p1)
    for fi21 := 0; fi21 < 14; fi21++ {
        switch (p1 + p1) {
        case "":
            if true {
                v22 := []int{fi21, (fi21 << ((0) & 63))}
                v22 = append(v22, (fi21 - fi21))
                if (!(false || true)) {
                    var v23 int64 = 295
                    _ = v23
                }
                _ = v22
            } else {
                v24 := []int{fi21, fi21, (fi21 << ((255) & 63))}
                _ = v24
            }
        }
        p0 = (p0 << ((0) & 63))
        _ = fi21
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 bool = false
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v20
    return 0
}

func main() {
    v1 := []uint32{65535}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 18446744073709551615
    v4 := []string{}
    var v5 float64 = ((((-2.25 / float64(0.1)) * ((0.1 + 0.1) + 0.0)) * (((0.1 * (0.1 + float64(0.1))) / (1e10 + ((float64((0.0 / 1e10)) / (float64(float32(float64((0.1 / 1e10)))) / (float64(-2.25) / -2.25))) - 3.14159))) * ((1e10 * 1.5) + ((3.14159 * ((1e10 / 3.14159) * 0.0)) * float64((3.14159 + (-2.25 * (0.1 / 0.1)))))))) / (-2.25 / (0.1 - ((0.1 + (-2.25 + (0.1 - 0.1))) - (-2.25 * (1.5 - ((3.14159 + 0.0) * float64(1.5))))))))
    if (793 != 436) {
        var v27 uint32 = 1
        v28 := v2[42]
        _ = v27
        _ = v28
    }
    var v29 uint32 = 619
    fmt.Println(v5, v3, v29)
    var v30 int8 = -74
    v31 := GenS2{F0: 32767, F1: 1}
    var v32 uint16 = 0
    switch 9223372036854775807 {
    case 2147483647:
        if false {
            for fi33 := 0; fi33 < 5; fi33++ {
                fmt.Println(v3)
                _ = fi33
            }
        }
    case 128:
        fmt.Println(v5, v3, v30)
    case -9223372036854775808:
        fmt.Println(v30)
    default:
        fmt.Println(v5)
    }
    var v34 uint32 = 67
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
