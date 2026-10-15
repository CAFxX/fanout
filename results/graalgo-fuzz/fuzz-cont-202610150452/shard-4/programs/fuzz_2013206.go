package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint8; F1 uint; F2 int64 }
func genFn0(p0 int64, p1 uint) uint {
    p1 = (p1 % (((p1 - (p1 << ((255) & 63)))) | 1))
    var v6 int64 = (p0 * p0)
    _ = p0
    _ = p1
    _ = v6
    return (p1 | p1)
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (255 != 255)
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    var v5 int64 = -1
    v5 = v5
    v7 := GenS1{F0: 71, F1: 42}
    switch -736 {
    case -458:
        if ((int16(v5) > 32767) || v2) {
            delete(v4, "foo bar")
        } else {
            if (v5 < v5) {
                fmt.Println(v2)
                switch "" {
                case "foo bar":
                    v5 -= (v5 >> ((214) & 63))
                default:
                    var v8 int32 = 0
                    _ = v8
                }
            }
        }
    case 521:
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS2{F0: v2, F1: ("hello" + "\n\t")}
                fmt.Println(v5, v5)
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: -1, F1: 65535}
            v2 = v2
            _ = fi9
            _ = v12
        }
    case -654:
        v5 -= -904
    default:
        v2 = ((0.0 > float32(float64(3.14159))) || ((^(v5)) > (v5 + v5)))
    }
    var v13 int64 = -568
    if (!(!(1 == -1))) {
        v5 -= (v13 % ((-1) | 1))
        var v14 bool = v2
        _ = v14
    }
    var v15 int8 = 100
    var v16 uint = 0
    fmt.Println(v16)
    var v17 int16 = 1
    v18 := v3["\n\t"]
    if v2 {
        fmt.Println(v5, v15, v15)
    } else {
        if v2 {
            v19 := GenS3{F0: 0, F1: v16, F2: v5}
            _ = v19
        } else {
            var v20 float32 = (0.0 - ((3.14159 * (0.1 + (-2.25 / 1.0))) / 1e10))
            var v21 uint64 = 1
            _ = v20
            _ = v21
        }
    }
    switch ((("hello" + ("a" + ((("\n\t" + "a") + "hello") + "\n\t"))) + "foo bar") + "\n\t") {
    case "\n\t":
        v22 := GenS3{F0: 231, F1: (^(v16)), F2: v13}
        _ = v22
    default:
        if v2 {
            v23 := []int{(v18 & (v18 % ((-438) | 1)))}
            _ = v23
        } else {
            switch -2147483648 {
            case 538:
                var v24 uint8 = 128
                _ = v24
            case 127:
                if v2 {
                    fmt.Println(v13, v17, v13)
                    v25 := GenS3{F0: 42, F1: (v16 << ((v16) & 63)), F2: (v13 * (v5 / (((v5 << (((^(v16))) & 63))) | 1)))}
                    _ = v25
                } else {
                    var v26 uint16 = 911
                    fmt.Println(v15, v13, v18)
                    var v27 int64 = v13
                    _ = v26
                    _ = v27
                }
            case 42:
                v28 := v4[("\n\t" + "foo bar")]
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                if ("hello" < "a") {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        v1[520] = (fi29 + fi30)
                        fmt.Println(v18, v5)
                        _ = fi30
                    }
                    if true {
                        var v31 int32 = 1
                        v32 := []int{fi29}
                        v32 = append(v32, -413)
                        _ = v31
                        _ = v32
                    } else {
                        switch "a" {
                        case "":
                            v33 := []int{(v18 | v18), fi29, (v18 | (fi29 - 32768))}
                            _ = v33
                        case "foo bar":
                            var v34 int = int(v16)
                            _ = v34
                        default:
                            for fi35 := 0; fi35 < 17; fi35++ {
                                v16 += v16
                                var v36 int8 = (^(v15))
                                _ = fi35
                                _ = v36
                            }
                        }
                    }
                    v17 = (v17 * int16(v13))
                } else {
                    v37 := []int{4294967295}
                    _ = v37
                }
                var v38 int16 = (v17 - (v17 >> ((v16) & 63)))
                _ = fi29
                _ = v38
            }
            v39 := GenS1{F0: 65, F1: 65535}
            _ = v39
        }
    }
    v40 := GenS3{F0: 1, F1: (^(v16)), F2: v13}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v7
}
