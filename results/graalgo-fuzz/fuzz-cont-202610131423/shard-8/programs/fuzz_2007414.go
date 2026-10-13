package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int64; F1 int16; F2 string }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = "foo bar"
    var v3 int8 = 127
    v4 := make(map[int]uint32)
    var v5 string = v2
    var v6 uint8 = 128
    v7 := GenS1{F0: 1, F1: 508}
    if false {
        v8 := GenS2{F0: ("\n\t" + v2), F1: 0}
        v3 = (^(v3))
        v9 := GenS3{F0: (-(v1)), F1: -137, F2: v2}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint = 117
            _ = fi11
            _ = v12
        }
        if true {
            v13 := GenS1{F0: -119, F1: 65535}
            for fi14 := 0; fi14 < 10; fi14++ {
                v5 = ("a" + v5)
                _ = fi14
            }
            _ = v13
        } else {
            v2 = v5
            var v15 float32 = 0.1
            _ = v15
        }
        _ = fi10
    }
    var v16 string = ""
    for fi17 := 0; fi17 < 9; fi17++ {
        if (2 == 127) {
            v6 += (v6 << ((582) & 63))
            v6 += (v6 + v6)
            var v18 uint16 = 42
            _ = v18
        } else {
            v5 = v2
            for fi19 := 0; fi19 < 13; fi19++ {
                v2 = "a"
                _ = fi19
            }
            if (!(!(2147483647 != 17))) {
                var v20 uint16 = 1
                fmt.Println(v2, v3)
                _ = v20
            } else {
                if false {
                    var v21 float32 = ((3.14159 + float32(0.1)) / 3.14159)
                    _ = v21
                } else {
                    v22 := []int{fi17, 612, fi17, fi17}
                    for fi23 := 0; fi23 < 4; fi23++ {
                        v22 = append(v22, (fi17 ^ fi17))
                        v2 = "hello"
                        if false {
                            v16 = "\n\t"
                            switch "" {
                            case "":
                                switch fi17 {
                                case 653:
                                    v24 := GenS3{F0: int64(fi23), F1: 1, F2: v5}
                                    v16 = (v16 + ("" + v16))
                                    fmt.Println(v1, v2)
                                    _ = v24
                                case -301:
                                    if len(v22) > 0 {
                                        v22[0] = (fi17 / ((4294967295) | 1))
                                    }
                                case 9223372036854775807:
                                    v4[(fi17 - 1)] = 65535
                                }
                            case "a":
                                switch v5 {
                                case "\n\t":
                                    _ = len(v4)
                                }
                            case "foo bar":
                                var v25 int32 = 0
                                _ = v25
                            default:
                                var v26 int8 = (^(v3))
                                _ = v26
                            }
                        } else {
                            var v27 uint64 = 0
                            _ = v27
                        }
                        _ = fi23
                    }
                    for fi28 := 0; fi28 < 5; fi28++ {
                        v29 := GenS1{F0: 0, F1: 476}
                        _ = fi28
                        _ = v29
                    }
                    _ = v22
                }
                v30 := GenS2{F0: ("\n\t" + v16), F1: 1}
                _ = v30
            }
        }
        v2 = "a"
        _ = fi17
    }
    var v31 uint16 = 65535
    var v32 float64 = 1e10
    v1 = (v1 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
