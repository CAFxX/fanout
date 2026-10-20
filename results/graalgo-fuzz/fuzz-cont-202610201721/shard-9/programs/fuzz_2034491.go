package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint = 0
    var v2 uint = 2
    var v3 uint64 = 4294967295
    var v4 uint32 = 2147483647
    if false {
        v5 := []int{-325, -9223372036854775808, 646}
        v6 := GenS1{F0: (!true)}
        _ = v5
        _ = v6
    } else {
        var v7 float32 = 0.1
        if (-377 >= 9223372036854775807) {
            if (true || true) {
                v8 := []int{557, -895}
                if true {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        for fi10 := 0; fi10 < 17; fi10++ {
                            v11 := GenS1{F0: false}
                            switch fi10 {
                            case 42:
                                v1 = 42
                            case 432:
                                v7 += ((1e10 / v7) - (0.0 - v7))
                            case 127:
                                fmt.Println(v7, v1, v4)
                            }
                            var v12 uint8 = 194
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        _ = fi9
                    }
                    switch 9223372036854775807 {
                    case 1:
                        v8 = append(v8, -485)
                    default:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                    v13 := GenS1{F0: (true && ((v4 != uint32(v1)) || (true && (true && true))))}
                    _ = v13
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v14 string = ""
                        _ = v14
                    }
                }
                v2 = v1
                _ = v8
            } else {
                v15 := []int{}
                var v16 float64 = 0.1
                _ = v15
                _ = v16
            }
            switch "foo bar" {
            case "hello":
                var v17 uint16 = 1
                _ = v17
            default:
                for fi18 := 0; fi18 < 20; fi18++ {
                    v2 += uint(v4)
                    _ = fi18
                }
            }
            v19 := GenS1{F0: (42 != 42)}
            _ = v19
        }
        if false {
            v2 -= v2
            if false {
                v20 := GenS2{F0: ""}
                v21 := GenS2{F0: "\n\t"}
                _ = v20
                _ = v21
            } else {
                v22 := []int{-780, 42, -924, -253}
                v23 := GenS3{F0: "a"}
                _ = v22
                _ = v23
            }
            fmt.Println(v4)
        } else {
            v24 := []int{42, 127, 2147483647, 0}
            _ = v24
        }
        _ = v7
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        switch "\n\t" {
        case "a":
            if (-802 >= 1) {
                v26 := GenS4{F0: -1}
                _ = v26
            } else {
                var v27 uint32 = 65535
                v4 = v27
                _ = v27
            }
        case "hello":
            v4 = v4
        case "":
            v3 = uint64(v2)
        }
        var v28 uint64 = v3
        v3 = v3
        _ = fi25
        _ = v28
    }
    if false {
        v29 := []int{-169, 127, 4294967295, 32768}
        _ = v29
    }
    var v30 float32 = 0.1
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 uint64 = (v3 ^ 0)
        _ = fi31
        _ = v32
    }
    var v33 float32 = (v30 * v30)
    fmt.Println(v2)
    var v34 uint8 = 128
    v30 = (v33 - v33)
    v35 := GenS5{F0: 0}
    if true {
        fmt.Println(v1, v30, v2)
        v36 := []int{127, 735, -2147483648, -520}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
