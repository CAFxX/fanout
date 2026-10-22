package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 int32; F2 int64 }
type GenS3 struct { F0 bool }
func genFn0() int64 {
    var v4 uint32 = 42
    v5 := GenS1{F0: (v4 - (v4 ^ 2)), F1: "", F2: -128}
    var v6 float64 = (3.14159 / 1.5)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int16 = -305
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v4 = 65535
        _ = fi9
    }
    _ = v4
    _ = v5
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 int8 = 0
    v3 += (v3 + v3)
    v3 += (v3 + v3)
    var v10 int = 128
    v11 := []int{2147483647, 212, v10, (-(v10))}
    if true {
        v12 := GenS2{F0: 1, F1: -2147483648, F2: 1}
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := v2[("a" + "")]
            var v15 uint8 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(v10, v10)
        _ = v12
    } else {
        v16 := GenS1{F0: 745, F1: "", F2: -128}
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        if false {
            var v18 uint = 0
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := GenS3{F0: true}
                switch (-(fi17)) {
                case 2147483647:
                    v1[("\n\t" + (("" + "foo bar") + ("hello" + "")))] = (((!(false && (!((255 == uint16(v18)) && true)))) || (false || false)) || (genFn0() == -1))
                case 4294967295:
                    fmt.Println(fi19, v10)
                default:
                    var v21 uint64 = 1
                    _ = v21
                }
                if (false && false) {
                    if false {
                        switch (^(fi19)) {
                        case 270:
                            v22 := GenS3{F0: false}
                            _ = v22
                        case -936:
                            var v23 uint64 = 2
                            var v24 uint32 = 65535
                            var v25 int8 = int8(v18)
                            _ = v23
                            _ = v24
                            _ = v25
                        case -338:
                            if (true && ((-670 < -9223372036854775808) && false)) {
                                var v26 int32 = 1
                                v27 := GenS1{F0: 1, F1: "hello", F2: int8(v10)}
                                v3 = v3
                                v10 = -1
                                _ = v26
                                _ = v27
                            } else {
                                _ = len(v11)
                                var v28 string = "hello"
                                var v29 uint32 = 364
                                v28 = v28
                                v30 := GenS1{F0: v29, F1: "hello", F2: -27}
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                        }
                    } else {
                        v31 := v1["foo bar"]
                        fmt.Println(fi17)
                        _ = v31
                    }
                } else {
                    if false {
                        v18 = (v18 << (((v18 % (((v18 ^ 0)) | 1))) & 63))
                    }
                    v10 -= v10
                    v32 := GenS2{F0: 42, F1: 495, F2: -9223372036854775808}
                    switch "hello" {
                    case "\n\t":
                        v33 := v2["hello"]
                        _ = v33
                    case "hello":
                        switch 302 {
                        case 2147483647:
                            v34 := GenS2{F0: -1, F1: 2147483647, F2: -961}
                            _ = v34
                        case 127:
                            v3 += (v3 ^ 0)
                        }
                    default:
                        fmt.Println(fi17)
                    }
                    var v35 int = (fi17 & fi19)
                    var v36 int64 = 1
                    v18 = v18
                    var v37 int64 = 1
                    _ = v32
                    _ = v35
                    _ = v36
                    _ = v37
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            v3 += v3
        }
        _ = fi17
    }
    fmt.Println(v10, v3)
    v11 = append(v11, v10)
    v38 := GenS3{F0: (v3 < (v3 & 1))}
    var v39 uint8 = 2
    v40 := GenS2{F0: -534, F1: 42, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
}
