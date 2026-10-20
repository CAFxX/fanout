package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint64, p1 int8) int16 {
    switch "" {
    case "":
        var v6 float64 = 0.1
        _ = v6
    default:
        if (1 <= 42) {
            v7 := GenS1{F0: 42, F1: 0, F2: ("foo bar" + "a")}
            var v8 int32 = 1
            _ = v7
            _ = v8
        } else {
            v9 := []int{-9223372036854775808, -2147483648, -329}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := []int{int(p0), fi10, fi10, fi10}
        if false {
            if (false || true) {
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 int64 = -349
                    switch fi12 {
                    case 0:
                        p0 -= (p0 % ((1) | 1))
                    default:
                        if (0 <= 2) {
                            v13 -= v13
                        }
                    }
                    switch (fi12 & (fi10 << ((1) & 63))) {
                    case 42:
                        if (((!true) || (3.14159 >= (((float32((((0.1 / 0.1) + 0.1) / (float64(1e10) + float64(0.1)))) / 1.0) - -2.25) - 1.5))) || true) {
                            var v14 uint32 = 2147483648
                            _ = v14
                        }
                    default:
                        p1 -= (-(p1))
                    }
                    var v15 float32 = (float32((1.5 * (1.5 + (1e10 + (0.1 * 0.0))))) / 3.14159)
                    p1 = p1
                    var v16 int64 = (v13 << ((255) & 63))
                    _ = fi12
                    _ = v13
                    _ = v15
                    _ = v16
                }
            } else {
                var v17 int = 32767
                _ = v17
            }
            var v18 uint8 = 67
            _ = v18
        } else {
            switch "" {
            case "":
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS1{F0: 0, F1: 4294967295, F2: ""}
                    _ = len(v11)
                    _ = fi19
                    _ = v20
                }
            case "foo bar":
                var v21 int = -668
                _ = v21
            case "\n\t":
                p1 -= 127
            }
        }
        _ = fi10
        _ = v11
    }
    p1 -= (-(p1))
    v22 := []int{}
    _ = p0
    _ = p1
    _ = v22
    return -374
}

func main() {
    var v1 int32 = 385
    v2 := []uint32{0, 312}
    var v3 uint32 = 65535
    var v4 int = 42
    var v5 float32 = (float32(0.1) * float32(-2.25))
    var v23 bool = (!(!true))
    if v23 {
        v5 += ((-2.25 - (v5 - ((((v5 - v5) - (v5 / (v5 - (v5 * v5)))) - float32(-2.25)) + v5))) / v5)
    }
    var v24 string = ("a" + "foo bar")
    v25 := GenS1{F0: 1, F1: 106, F2: v24}
    for fi26 := 0; fi26 < 8; fi26++ {
        if (0 >= (v1 % ((v1) | 1))) {
            if len(v2) > 0 {
                v2[0] = (v3 % ((v3) | 1))
            }
            fmt.Println(v4, fi26, fi26)
            v27 := GenS1{F0: 32767, F1: 1, F2: ("foo bar" + ((v24 + v24) + "\n\t"))}
            _ = v27
        } else {
            v2 = append(v2, v3)
        }
        v28 := GenS2{F0: 408, F1: 1000}
        _ = fi26
        _ = v28
    }
    var v29 int64 = 1
    if (296 >= 2) {
        var v30 bool = (!true)
        v4 = v4
        _ = v30
    } else {
        v31 := GenS3{F0: 0}
        _ = v31
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v32 uint32 = 4294967295
    v1 = v1
    var v33 float64 = (1e10 * -2.25)
    v34 := GenS2{F0: (v29 % (((-(v29))) | 1)), F1: int16(v29)}
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 float64 = float64((v5 * (0.1 / v5)))
        for fi37 := 0; fi37 < 14; fi37++ {
            fmt.Println(v1, v3, v33)
            if (0 > 2) {
                if v23 {
                    v33 = v33
                    v29 = v29
                    var v38 int16 = 42
                    _ = v38
                } else {
                    var v39 bool = v23
                    _ = v39
                }
                var v40 float64 = 0.0
                var v41 uint8 = 127
                _ = v40
                _ = v41
            }
            if ((344 < 829) || v23) {
                v42 := GenS2{F0: (v29 | -281), F1: 1000}
                _ = v42
            }
            _ = fi37
        }
        v29 -= (-(v29))
        _ = fi35
        _ = v36
    }
    var v43 int32 = v1
    v44 := GenS1{F0: -32768, F1: 0, F2: "hello"}
    v2 = append(v2, (v32 & v3))
    _ = len(v2)
    v23 = v23
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
