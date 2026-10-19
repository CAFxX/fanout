package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 int16 }

func main() {
    v1 := []int{32768, -1, 736, 4294967295}
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint64 = 2
        var v9 int = fi7
        if false {
            var v10 int16 = 42
            _ = v10
        } else {
            var v11 int16 = -1
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint64 = 456
        _ = fi12
        _ = v13
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    var v14 bool = true
    if (!v14) {
        if (-779 > 9223372036854775807) {
            v14 = false
        } else {
            var v15 int8 = -128
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS1{F0: 155, F1: 4294967295, F2: 859}
                v3["foo bar"] = 841
                for fi18 := 0; fi18 < 3; fi18++ {
                    if ((366 == -694) || v14) {
                        var v19 bool = v14
                        _ = v19
                    } else {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v15, fi18)
                        default:
                            v15 = -31
                        }
                    }
                    v3[(("\n\t" + "") + ("\n\t" + "\n\t"))] = 2147483648
                    var v20 string = ("foo bar" + ("foo bar" + ""))
                    _ = fi18
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        fmt.Println(v14)
        if false {
            var v21 int8 = -1
            _ = v21
        } else {
            v4[-2147483648] = -203
            var v22 int16 = 1000
            switch (("hello" + ("a" + "")) + "") {
            case "\n\t":
                v14 = (!v14)
            case "a":
                v23 := GenS1{F0: 18446744073709551615, F1: 133, F2: 42}
                _ = v23
            case "hello":
                if (!v14) {
                    if true {
                        if (false || v14) {
                            fmt.Println(v22, v22)
                            v14 = v14
                            v24 := GenS2{F0: 0, F1: 1, F2: 42}
                            _ = v24
                        }
                        var v25 int8 = 1
                        if (225 != 4294967295) {
                            var v26 uint = 0
                            for fi27 := 0; fi27 < 6; fi27++ {
                                for fi28 := 0; fi28 < 17; fi28++ {
                                    v26 = 4294967295
                                    _ = fi28
                                }
                                _ = fi27
                            }
                            var v29 int64 = 1
                            _ = v26
                            _ = v29
                        }
                        _ = v25
                    }
                } else {
                    var v30 uint16 = 255
                    _ = v30
                }
            }
            _ = v22
        }
    }
    v14 = v14
    for fi31 := 0; fi31 < 16; fi31++ {
        for fi32 := 0; fi32 < 11; fi32++ {
            fmt.Println(v14)
            var v33 uint64 = 18446744073709551615
            _ = fi32
            _ = v33
        }
        _ = len(v1)
        if (42 == 1) {
            v34 := GenS1{F0: 0, F1: 255, F2: -9223372036854775808}
            v35 := GenS3{F0: -123, F1: 18446744073709551615, F2: 1}
            _ = v34
            _ = v35
        } else {
            var v36 float32 = (((float32((float64(0.1) / 0.1)) / 3.14159) * 3.14159) / ((1e10 / 3.14159) / 3.14159))
            switch "hello" {
            case "foo bar":
                _ = len(v4)
            case "a":
                var v37 uint8 = 70
                _ = v37
            }
            _ = v36
        }
        _ = fi31
    }
    v14 = v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
