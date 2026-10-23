package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int; F2 bool }
func genFn0(p0 uint16) int32 {
    for fi4 := 0; fi4 < 8; fi4++ {
        if true {
            var v5 int16 = 0
            _ = v5
        }
        _ = fi4
    }
    var v6 uint8 = 0
    var v7 uint16 = p0
    switch "foo bar" {
    case "\n\t":
        v6 -= 255
    case "a":
        p0 = v7
    }
    if (false || (-218 == 389)) {
        v6 = uint8(p0)
    } else {
        v8 := GenS1{F0: 1}
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return -760
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "hello"
    v3 := make(map[uint32]uint32)
    fmt.Println(v2)
    v2 = "foo bar"
    var v9 bool = (0.1 > (3.14159 / (((0.1 * 0.1) * 1.5) / 0.1)))
    v9 = true
    v2 = ((v2 + (v2 + ((("a" + v2) + v2) + ("hello" + "hello")))) + "")
    for fi10 := 0; fi10 < 10; fi10++ {
        v2 = "hello"
        _ = fi10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        switch v2 {
        case "":
            if v9 {
                _ = len(v3)
                v2 = ""
                v2 = (("a" + v2) + (v2 + ""))
            }
        case "\n\t":
            var v12 float64 = 0.0
            _ = v12
        case "a":
            v9 = v9
        default:
            v2 = ((v2 + "a") + v2)
        }
        fmt.Println(v9)
        _ = fi11
    }
    switch -425 {
    case -1:
        fmt.Println(v9, v2)
    default:
        switch v2 {
        case "foo bar":
            v13 := GenS2{F0: 42}
            _ = v13
        case "hello":
            if v9 {
                if (!(!(42 > 1))) {
                    _ = len(v3)
                    v2 = v2
                    if false {
                        for fi14 := 0; fi14 < 13; fi14++ {
                            v9 = v9
                            v2 = v2
                            v15 := GenS2{F0: 2}
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 6; fi16++ {
                            if false {
                                var v17 int64 = 1
                                _ = v17
                            }
                            fmt.Println(fi16, v9, fi16)
                            v2 = v2
                            _ = fi16
                        }
                        switch "a" {
                        case "hello":
                            v2 = (v2 + v2)
                        case "\n\t":
                            v2 = "a"
                        case "":
                            if (726 <= 886) {
                                v3[42] = 953
                                var v18 uint8 = 217
                                fmt.Println(v2, v9)
                                _ = v18
                            }
                        }
                    } else {
                        v19 := GenS1{F0: 4294967295}
                        v20 := []int{}
                        _ = v19
                        _ = v20
                    }
                } else {
                    v21 := GenS3{F0: 443, F1: 128, F2: false}
                    v22 := GenS1{F0: 1}
                    _ = v21
                    _ = v22
                }
            } else {
                v9 = true
            }
        case "\n\t":
            v23 := []int{-223, 127}
            _ = v23
        default:
            var v24 float64 = 1.5
            _ = v24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
