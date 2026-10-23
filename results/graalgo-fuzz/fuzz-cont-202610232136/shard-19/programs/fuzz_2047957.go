package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint; F1 uint32; F2 int8 }
type GenS3 struct { F0 string; F1 int64; F2 int }
type GenS4 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int16) int32 {
    var v6 uint8 = 42
    if true {
        v6 = v6
        p0 += p0
    }
    if false {
        v7 := GenS1{F0: 2147483648, F1: (p0 / ((-602) | 1))}
        v6 -= 1
        var v8 uint = uint(v6)
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 uint8 = 2
    v5 := make(map[int]uint32)
    fmt.Println(v4)
    v4 = v4
    switch "" {
    case "hello":
        fmt.Println(v4)
    case "foo bar":
        for fi9 := 0; fi9 < 5; fi9++ {
            if (!false) {
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 int = 4294967295
                    var v12 uint16 = 983
                    v11 = (fi9 / (((v11 << ((606) & 63))) | 1))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                if (!true) {
                    if (2 <= 2147483647) {
                        v4 += (v4 ^ (v4 - (^(v4))))
                        for fi13 := 0; fi13 < 17; fi13++ {
                            if ((2 != 247) || true) {
                                v4 = v4
                                fmt.Println(fi13)
                            }
                            var v14 uint32 = 782
                            _ = fi13
                            _ = v14
                        }
                        v15 := GenS1{F0: 0, F1: 32767}
                        _ = v15
                    }
                    var v16 uint16 = uint16(v4)
                    v16 = 300
                    _ = v16
                } else {
                    var v17 uint64 = 4294967295
                    _ = v17
                }
                switch "" {
                case "\n\t":
                    v4 = 177
                case "hello":
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 int = (fi9 | fi18)
                        _ = fi18
                        _ = v19
                    }
                default:
                    delete(v3, (fi9 & fi9))
                }
            }
            v20 := []int{}
            v4 = (v4 / ((v4) | 1))
            _ = fi9
            _ = v20
        }
    case "a":
        v21 := []int{32767, 9223372036854775807}
        _ = v21
    default:
        var v22 int16 = 42
        _ = v22
    }
    var v23 uint16 = 856
    for fi24 := 0; fi24 < 14; fi24++ {
        v25 := GenS2{F0: 255, F1: 2147483648, F2: 1}
        _ = fi24
        _ = v25
    }
    fmt.Println(v4, v4)
    fmt.Println(v23)
    v26 := GenS3{F0: "", F1: 9223372036854775807, F2: 235}
    v1[2] = 42
    for fi27 := 0; fi27 < 8; fi27++ {
        if false {
            delete(v2, 4294967295)
        } else {
            for fi28 := 0; fi28 < 15; fi28++ {
                v29 := GenS4{F0: (v23 % (((v23 | (v23 / ((v23) | 1)))) | 1)), F1: 2}
                _ = fi28
                _ = v29
            }
            if false {
                var v30 uint32 = 2147483648
                switch 9223372036854775807 {
                case -145:
                    var v31 uint64 = 4294967295
                    _ = v31
                case 127:
                    v3[fi27] = ""
                }
                _ = v30
            } else {
                v23 -= 671
                var v32 int = -408
                fmt.Println(v23, v23)
                _ = v32
            }
            if (!(true && (-865 >= -65))) {
                v33 := GenS3{F0: "", F1: 0, F2: fi27}
                var v34 uint8 = (v4 | (^(v4)))
                v34 = (v4 << ((330) & 63))
                _ = v33
                _ = v34
            } else {
                for fi35 := 0; fi35 < 7; fi35++ {
                    v23 = (v23 % (((v23 - 628)) | 1))
                    var v36 float64 = (3.14159 * 1e10)
                    v5[(fi27 ^ (fi35 >> ((2) & 63)))] = 136
                    _ = fi35
                    _ = v36
                }
                fmt.Println(v4, v23)
            }
        }
        if (749 < 255) {
            var v37 float64 = (0.0 * 1.5)
            v38 := GenS2{F0: 0, F1: 0, F2: 127}
            _ = v37
            _ = v38
        } else {
            v4 += (v4 ^ v4)
            var v39 string = ("hello" + ("" + "foo bar"))
            for fi40 := 0; fi40 < 13; fi40++ {
                v4 += v4
                _ = fi40
            }
            _ = v39
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
