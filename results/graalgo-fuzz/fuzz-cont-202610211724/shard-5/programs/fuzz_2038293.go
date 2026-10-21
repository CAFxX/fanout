package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint64 }
func genFn0() uint16 {
    var v5 uint16 = 65535
    var v6 uint8 = 42
    v5 += 255
    _ = v5
    _ = v6
    return (v5 - 855)
}
func genFn1() uint32 {
    var v7 int64 = -9223372036854775808
    v8 := []int{}
    v7 = 857
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 42
    v3 := make(map[int]string)
    var v4 bool = ((!true) && ((127 > 128) && true))
    if false {
        for fi9 := 0; fi9 < 14; fi9++ {
            if (("\n\t" + "hello") == ("a" + "\n\t")) {
                v4 = v4
                for fi10 := 0; fi10 < 12; fi10++ {
                    if false {
                        if true {
                            var v11 int64 = (^(v2))
                            v12 := GenS1{F0: genFn1(), F1: 2147483647}
                            _ = v11
                            _ = v12
                        } else {
                            v13 := v3[(fi9 << ((1) & 63))]
                            switch (v13 + "hello") {
                            case "hello":
                                switch "a" {
                                case "foo bar":
                                    var v14 string = v13
                                    _ = v14
                                case "\n\t":
                                    if false {
                                        v2 = (v2 + (v2 << ((65535) & 63)))
                                    }
                                default:
                                    fmt.Println(fi10, v13)
                                }
                            case "":
                                var v15 string = v13
                                _ = v15
                            }
                            v16 := GenS2{F0: 32767, F1: (v4 || ((42 >= 740) || false)), F2: 32767}
                            _ = v13
                            _ = v16
                        }
                    }
                    _ = fi10
                }
            } else {
                if false {
                    var v17 uint8 = 1
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 12; fi18++ {
                        switch (fi9 ^ (fi18 >> ((676) & 63))) {
                        case -9223372036854775808:
                            var v19 int32 = 42
                            _ = v19
                        case 9223372036854775807:
                            fmt.Println(fi18, v2)
                        }
                        _ = fi18
                    }
                    _ = len(v1)
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    delete(v1, (fi9 >> ((42) & 63)))
                    v21 := GenS2{F0: 112, F1: true, F2: 1}
                    v22 := GenS3{F0: 74}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := v1[(fi9 / (((fi9 << ((18446744073709551615) & 63))) | 1))]
                _ = v23
            }
            _ = fi9
        }
        var v24 int16 = 0
        _ = v24
    }
    var v25 uint = 1
    v26 := []int{9223372036854775807, 1, -9223372036854775808, 1}
    var v27 int16 = 0
    var v28 float64 = 0.1
    v27 = v27
    v29 := GenS4{F0: 102, F1: true, F2: 4294967295}
    v30 := GenS3{F0: 65}
    v31 := GenS3{F0: 0}
    v32 := GenS5{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
