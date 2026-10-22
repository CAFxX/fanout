package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0(p0 uint64, p1 string, p2 uint8) uint {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 float64 = (1e10 - 0.1)
        p1 = p1
        fmt.Println(p0, p1, p2)
        _ = fi4
        _ = v5
    }
    p0 -= (p0 << ((255) & 63))
    p2 = (p2 - p2)
    p1 = "a"
    for fi6 := 0; fi6 < 1; fi6++ {
        p1 = "\n\t"
        switch (p1 + p1) {
        case "":
            var v7 uint32 = 168
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: 18446744073709551615, F1: 2, F2: -1}
            _ = v8
        default:
            var v9 uint32 = 540
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 315
}
func genFn1() int8 {
    if ("\n\t" > "\n\t") {
        v10 := GenS1{F0: 42, F1: 927, F2: 0}
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(fi11)
            _ = fi11
        }
        _ = v10
    }
    var v12 int = 0
    switch (v12 ^ (v12 + v12)) {
    case 32768:
        v13 := GenS1{F0: 1, F1: 848, F2: 1}
        _ = v13
    }
    _ = v12
    return 127
}

func main() {
    var v1 uint64 = 257
    var v2 int64 = 1
    var v3 int16 = 0
    fmt.Println(v3, v1)
    var v14 uint = genFn0(v1, ("foo bar" + "foo bar"), 255)
    v15 := GenS2{F0: ((v1 - v1) > (v1 ^ 18446744073709551615))}
    v2 += 1
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            v1 += v1
            _ = fi17
        }
        _ = fi16
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS1{F0: v14, F1: 65, F2: v2}
        if ((-(v2)) < (-(v2))) {
            for fi20 := 0; fi20 < 13; fi20++ {
                v2 += 1
                v2 = v2
                if (457 != 65535) {
                    switch "" {
                    case "hello":
                        switch (fi18 | (-(fi20))) {
                        case 2147483647:
                            if (2 < 0) {
                                var v21 float64 = (0.0 + 3.14159)
                                if true {
                                    if (!false) {
                                        v22 := []int{1, fi20}
                                        v23 := GenS2{F0: (genFn0(v1, ("hello" + "a"), 0) == genFn0(4294967295, "foo bar", 1))}
                                        var v24 int32 = -2147483648
                                        v22 = append(v22, fi18)
                                        var v25 int32 = v24
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    v26 := []int{fi20}
                                    var v27 uint64 = (v1 << (((v14 * (^(v14)))) & 63))
                                    v3 -= v3
                                    var v28 float64 = (float64(1.5) * (3.14159 + v21))
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 float64 = v21
                                    var v30 bool = ((^(fi18)) < fi20)
                                    _ = v29
                                    _ = v30
                                }
                                v3 += (v3 >> ((genFn0(106, "\n\t", 132)) & 63))
                                _ = v21
                            }
                        }
                    case "":
                        var v31 float32 = (1e10 - (1e10 * 1e10))
                        _ = v31
                    default:
                        v32 := []int{fi18, 785, fi20, (fi18 - (fi18 | 0))}
                        _ = v32
                    }
                } else {
                    switch (fi18 - 121) {
                    case 128:
                        var v33 int16 = v3
                        _ = v33
                    case 0:
                        v14 = v14
                    }
                    v34 := GenS3{F0: (v3 - v3), F1: -360}
                    var v35 float64 = 1.5
                    _ = v34
                    _ = v35
                }
                _ = fi20
            }
        }
        var v36 int8 = genFn1()
        _ = fi18
        _ = v19
        _ = v36
    }
    var v37 uint16 = 255
    var v38 uint16 = v37
    v39 := GenS2{F0: false}
    var v40 int64 = v2
    var v41 int64 = (v40 + (v2 % ((42) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
