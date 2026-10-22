package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 bool }

func main() {
    var v1 float64 = (float64(1.5) * (0.1 / 0.1))
    var v2 string = ("a" + "a")
    var v3 uint32 = 0
    var v4 string = v2
    v5 := make(map[uint32]string)
    fmt.Println(v3, v3, v2)
    fmt.Println(v4, v4)
    if (false || (v4 >= v4)) {
        v6 := []int{32767, int(v3), 591}
        _ = v6
    }
    var v7 uint64 = 320
    v4 = v2
    fmt.Println(v4, v1, v4)
    v8 := GenS1{F0: true}
    v7 -= (v7 * v7)
    switch (v2 + v2) {
    case "\n\t":
        var v9 int16 = 0
        _ = v9
    }
    if true {
        if (-1 > 1) {
            v1 -= 1e10
        } else {
            var v10 int8 = 42
            v11 := []int{-729, 495, -1, 4294967295}
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v3)
            if false {
                var v13 int8 = 0
                v14 := GenS2{F0: 128, F1: (v3 - (v3 % ((v3) | 1))), F2: 31}
                v15 := GenS3{F0: 100, F1: 1}
                _ = v13
                _ = v14
                _ = v15
            } else {
                var v16 int32 = -1
                switch (fi12 ^ (fi12 * 42)) {
                case -2147483648:
                    var v17 int8 = -86
                    _ = v17
                case -686:
                    v18 := GenS1{F0: (v7 != (v7 << ((1) & 63)))}
                    _ = v18
                }
                _ = v16
            }
            v4 = v2
            _ = fi12
        }
    } else {
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := []int{(fi19 ^ 0)}
            _ = fi19
            _ = v20
        }
    }
    v21 := GenS1{F0: ((!true) || true)}
    v22 := GenS4{F0: 335, F1: 255}
    fmt.Println(v4, v2)
    fmt.Println(v4, v3, v1)
    for fi23 := 0; fi23 < 17; fi23++ {
        v4 = v4
        if true {
            switch fi23 {
            case 127:
                var v24 uint = 18446744073709551615
                _ = v24
            case 1:
                var v25 uint32 = (^(v3))
                _ = v25
            case -314:
                var v26 string = "a"
                _ = v26
            default:
                fmt.Println(v2, v7, v3)
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                fmt.Println(v3)
                _ = fi27
            }
            if (false || true) {
                switch "hello" {
                case "foo bar":
                    delete(v5, v3)
                case "a":
                    v28 := GenS2{F0: 42, F1: (^(v3)), F2: 2}
                    _ = v28
                default:
                    var v29 uint32 = 65535
                    _ = v29
                }
                switch 127 {
                case 2147483647:
                    delete(v5, v3)
                case 471:
                    for fi30 := 0; fi30 < 16; fi30++ {
                        if false {
                            if (!(!true)) {
                                v7 = v7
                            }
                        } else {
                            if false {
                                var v31 uint64 = v7
                                _ = len(v5)
                                switch (fi30 % (((fi30 / ((fi23) | 1))) | 1)) {
                                case 341:
                                    fmt.Println(v1, fi30, v7)
                                case -1:
                                    v32 := v5[(^(v3))]
                                    _ = v32
                                case -359:
                                    v33 := []int{(fi30 ^ -84), (^(fi30))}
                                    _ = v33
                                }
                                v4 = v2
                                _ = v31
                            }
                            v34 := []int{}
                            var v35 int32 = 999
                            _ = v34
                            _ = v35
                        }
                        _ = fi30
                    }
                case -388:
                    var v36 float32 = (1e10 + float32(v1))
                    _ = v36
                }
                v1 = (v1 * ((0.0 - float64((3.14159 + (1e10 + (1e10 * 0.1))))) - v1))
            }
        } else {
            for fi37 := 0; fi37 < 13; fi37++ {
                v38 := []int{(-(fi23)), -2147483648}
                v39 := v5[2]
                if ((v1 * v1) <= v1) {
                    if len(v38) > 0 {
                        v38[0] = fi23
                    }
                    var v40 bool = false
                    var v41 int32 = 911
                    _ = v40
                    _ = v41
                } else {
                    v3 += v3
                    v3 = (v3 & (v3 % (((^(v3))) | 1)))
                }
                _ = fi37
                _ = v38
                _ = v39
            }
        }
        switch (v2 + v4) {
        case "foo bar":
            v42 := GenS5{F0: false}
            _ = v42
        case "a":
            var v43 int8 = -93
            _ = v43
        case "hello":
            if (-1 < 9223372036854775807) {
                v2 = ("foo bar" + (v2 + "\n\t"))
                var v44 string = v4
                _ = v44
            }
        default:
            v45 := GenS1{F0: (!false)}
            _ = v45
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
