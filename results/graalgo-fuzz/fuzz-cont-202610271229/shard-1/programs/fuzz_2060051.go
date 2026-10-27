package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 uint8 = 54
    v2 := make(map[string]int)
    var v3 float64 = 1.5
    var v4 float64 = v3
    v3 += 3.14159
    var v5 int8 = 85
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 string = "a"
        v8 := []int{(fi6 ^ fi6)}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = v1
    v10 := []int{}
    var v11 uint16 = 2
    v9 = v1
    if true {
        v12 := GenS1{F0: 65535}
        fmt.Println(v1)
        switch ("a" + ("" + ("hello" + ("" + "")))) {
        case "a":
            v13 := GenS1{F0: 0}
            _ = v13
        case "\n\t":
            v14 := GenS1{F0: 174}
            _ = v14
        case "hello":
            if true {
                var v15 bool = true
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS1{F0: 373}
                    _ = fi16
                    _ = v17
                }
                v10 = append(v10, 9223372036854775807)
                _ = v15
            } else {
                v18 := v2["foo bar"]
                _ = v18
            }
        default:
            v3 -= v4
        }
        _ = v12
    }
    v19 := GenS2{F0: 916, F1: -1}
    if false {
        switch "a" {
        case "":
            fmt.Println(v11, v5)
        case "a":
            v20 := v2[(("\n\t" + ("\n\t" + "")) + "a")]
            _ = v20
        }
    }
    switch "foo bar" {
    case "\n\t":
        v21 := GenS3{F0: 1, F1: -2147483648}
        _ = v21
    }
    switch -9223372036854775808 {
    case 2147483647:
        v5 = (^(v5))
    case 4294967295:
        if true {
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v3)
                var v23 int8 = 0
                switch fi22 {
                case 2147483647:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        if (true || (uint64(v9) == 93)) {
                            var v25 uint64 = 388
                            var v26 int64 = -764
                            _ = v25
                            _ = v26
                        }
                        v9 = v1
                        v27 := GenS4{F0: "hello"}
                        _ = fi24
                        _ = v27
                    }
                case 70:
                    if ((false && true) || (42 != 100)) {
                        var v28 float64 = v4
                        var v29 uint = 825
                        if (-2147483648 > 0) {
                            fmt.Println(v5, fi22)
                            var v30 uint8 = (v9 | v1)
                            fmt.Println(v23, v29)
                            _ = v30
                        } else {
                            for fi31 := 0; fi31 < 15; fi31++ {
                                var v32 int32 = -2147483648
                                if len(v10) > 0 {
                                    v10[0] = fi31
                                }
                                if true {
                                    v1 += v1
                                    var v33 string = "\n\t"
                                    _ = v33
                                }
                                _ = fi31
                                _ = v32
                            }
                            v34 := GenS5{F0: (v11 / (((v11 / ((976) | 1))) | 1))}
                            if true {
                                var v35 bool = true
                                _ = v35
                            }
                            _ = v34
                        }
                        _ = v28
                        _ = v29
                    }
                }
                _ = fi22
                _ = v23
            }
            v1 -= v1
            var v36 bool = true
            _ = v36
        } else {
            v37 := GenS2{F0: 668, F1: -1}
            var v38 uint64 = 42
            _ = v37
            _ = v38
        }
    }
    fmt.Println(v4)
    switch "a" {
    case "":
        var v39 uint32 = 0
        _ = v39
    case "hello":
        fmt.Println(v9, v9, v9)
    default:
        fmt.Println(v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
