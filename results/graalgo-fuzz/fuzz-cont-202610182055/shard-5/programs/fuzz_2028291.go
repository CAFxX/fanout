package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
func genFn0(p0 int8) int {
    v5 := []int{}
    for fi6 := 0; fi6 < 9; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 0
    var v2 bool = ((0 >= 94) || true)
    var v3 uint = 18446744073709551615
    v4 := make(map[string]string)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int16 = -1
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := []int{(fi7 - (fi7 - fi9)), (fi7 >> (((v3 % ((603) | 1))) & 63)), (fi7 % ((int(v8)) | 1)), fi7}
            v3 = 4294967295
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 uint64 = 999
    v3 += v3
    _ = len(v4)
    v12 := []int{139, genFn0(1), 0}
    switch "hello" {
    case "\n\t":
        _ = len(v4)
    }
    delete(v4, ("hello" + ("\n\t" + "\n\t")))
    v13 := GenS1{F0: 2, F1: 570}
    var v14 uint16 = 1
    for fi15 := 0; fi15 < 11; fi15++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        if v2 {
            if false {
                var v16 uint16 = 81
                _ = len(v12)
                _ = v16
            } else {
                v11 = (v11 - v1)
                v17 := v4["hello"]
                _ = v17
            }
        }
        _ = fi15
    }
    if v2 {
        v18 := GenS1{F0: 624, F1: 162}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        if v2 {
            var v19 int32 = int32(v1)
            _ = v19
        }
        _ = v18
    } else {
        switch int(v14) {
        case -21:
            fmt.Println(v1)
        case 32768:
            if true {
                v3 += (^(v3))
                var v20 int = -566
                _ = v20
            } else {
                if v2 {
                    var v21 int32 = -2147483648
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v2 = ((false || (-491 <= -238)) && false)
                        var v23 uint8 = 127
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                } else {
                    switch 42 {
                    case 106:
                        for fi24 := 0; fi24 < 7; fi24++ {
                            v25 := GenS1{F0: 255, F1: 2147483648}
                            if (!false) {
                                v4[""] = "foo bar"
                                fmt.Println(fi24, v1, v3)
                            } else {
                                var v26 string = ""
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 13; fi27++ {
                                v3 = v3
                                switch ("" + "") {
                                case "":
                                    v1 = v1
                                }
                                _ = fi27
                            }
                            _ = fi24
                            _ = v25
                        }
                    case 127:
                        if len(v12) > 0 {
                            v12[0] = -9223372036854775808
                        }
                    case 350:
                        var v28 bool = v2
                        _ = v28
                    }
                }
                v1 = v1
            }
        default:
            v29 := GenS2{F0: 1}
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        fmt.Println(v11, fi30)
        for fi31 := 0; fi31 < 11; fi31++ {
            if ((v3 | v3) <= (v3 * (v3 + (^(v3))))) {
                v14 -= 65535
            }
            var v32 uint64 = (v1 & v11)
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v3 -= (v3 << ((v3) & 63))
    if v2 {
        for fi33 := 0; fi33 < 3; fi33++ {
            v34 := GenS3{F0: (v14 + 598), F1: "a", F2: -571}
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 8; fi35++ {
            var v36 int32 = 2147483647
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
