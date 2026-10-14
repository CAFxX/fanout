package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint) uint {
    p0 = (p0 + p0)
    v5 := []int{}
    p0 = p0
    p0 -= (p0 & p0)
    _ = p0
    _ = v5
    return p0
}

func main() {
    var v1 uint8 = 128
    var v2 bool = true
    v3 := make(map[string]uint32)
    var v4 string = "hello"
    fmt.Println(v4)
    var v6 int64 = 698
    switch v4 {
    case "\n\t":
        v7 := []int{0, 128}
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 9; fi8++ {
            v4 = v4
            switch (fi8 >> ((2) & 63)) {
            case 9223372036854775807:
                v9 := GenS1{F0: v6}
                _ = v9
            default:
                v10 := GenS1{F0: -684}
                _ = v10
            }
            v11 := []int{(fi8 / ((fi8) | 1))}
            _ = fi8
            _ = v11
        }
    case "a":
        v1 -= v1
    default:
        v4 = "hello"
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint = 2
        v14 := GenS1{F0: v6}
        switch v4 {
        case "a":
            var v15 int16 = 986
            _ = v15
        default:
            v16 := GenS1{F0: (v6 ^ (v6 >> ((844) & 63)))}
            _ = v16
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    if v2 {
        v17 := v3[((v4 + v4) + (v4 + (v4 + (v4 + (v4 + "foo bar")))))]
        v2 = v2
        if ((128 >= 42) && v2) {
            for fi18 := 0; fi18 < 17; fi18++ {
                for fi19 := 0; fi19 < 14; fi19++ {
                    switch 1 {
                    case -2147483648:
                        var v20 uint16 = 2
                        _ = v20
                    case 32768:
                        v1 -= v1
                        switch "" {
                        case "hello":
                            var v21 string = (v4 + (v4 + v4))
                            _ = v21
                        case "":
                            var v22 int32 = 0
                            _ = v22
                        }
                    }
                    v23 := []int{(fi19 + -898), (fi19 * (fi19 + fi19)), (fi19 & -580)}
                    v6 -= -9223372036854775808
                    switch ((v4 + v4) + ("\n\t" + ("\n\t" + v4))) {
                    case "\n\t":
                        fmt.Println(fi19)
                    case "":
                        var v24 int = (fi18 << ((255) & 63))
                        _ = v24
                    default:
                        v1 -= (v1 % (((v1 ^ v1)) | 1))
                    }
                    _ = fi19
                    _ = v23
                }
                v25 := GenS2{F0: 0, F1: 42, F2: genFn0(255)}
                _ = fi18
                _ = v25
            }
            v26 := GenS1{F0: v6}
            _ = v26
        }
        _ = v17
    }
    if ((!v2) || v2) {
        if (163 == (^(v1))) {
            var v27 uint64 = 1
            var v28 uint8 = v1
            v28 = 255
            _ = v27
            _ = v28
        } else {
            v4 = "\n\t"
        }
        v2 = true
    } else {
        v1 += (v1 + v1)
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            v31 := GenS1{F0: -801}
            v6 -= 0
            switch v4 {
            case "":
                if v2 {
                    switch int(v1) {
                    case 790:
                        v32 := GenS3{F0: 0, F1: 1}
                        _ = v32
                    default:
                        v4 = "a"
                    }
                }
            case "foo bar":
                v33 := GenS2{F0: 799, F1: 42, F2: uint(v1)}
                _ = v33
            }
            _ = fi30
            _ = v31
        }
        v6 = 42
        v34 := v3[v4]
        _ = fi29
        _ = v34
    }
    v6 -= (v6 >> ((882) & 63))
    var v35 int16 = -1
    v36 := v3[v4]
    if false {
        v3["\n\t"] = (v36 - 2)
        var v37 int64 = (v6 | (v6 >> ((65535) & 63)))
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v6
}
