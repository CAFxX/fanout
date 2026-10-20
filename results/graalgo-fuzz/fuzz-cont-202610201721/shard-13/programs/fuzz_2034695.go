package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS5 struct { F0 bool; F1 string; F2 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64; F1 int64; F2 int }
func genFn0(p0 uint32, p1 uint32) int64 {
    v5 := []int{4294967295}
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 538, F1: 1}
        fmt.Println(p0)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 uint64, p1 uint) int32 {
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int = fi8
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        p0 -= uint64(p1)
        for fi11 := 0; fi11 < 2; fi11++ {
            if true {
                p1 += 4294967295
                fmt.Println(fi10, fi11)
                p1 -= p1
            } else {
                var v12 int64 = genFn0(42, 782)
                v13 := GenS2{F0: -128, F1: -1, F2: 2}
                v14 := GenS3{F0: 127, F1: 2, F2: 32767}
                p1 = 18446744073709551615
                switch fi11 {
                case 266:
                    switch (fi11 * 606) {
                    case 2147483647:
                        if false {
                            p0 += p0
                        }
                    case -1:
                        var v15 int8 = 0
                        _ = v15
                    }
                }
                _ = v12
                _ = v13
                _ = v14
            }
            p0 += (^(p0))
            _ = fi11
        }
        _ = fi10
    }
    if true {
        fmt.Println(p1, p0)
        var v16 uint16 = 684
        if (255 == 84) {
            v17 := []int{-2147483648, 42, 1, 4294967295}
            if true {
                switch 812 {
                case -9223372036854775808:
                    fmt.Println(v16, p0, p1)
                case -2147483648:
                    switch -1 {
                    case 9223372036854775807:
                        switch "\n\t" {
                        case "":
                            if ("foo bar" <= "hello") {
                                fmt.Println(v16, p1, v16)
                                var v18 int16 = -1
                                _ = v18
                            }
                        case "hello":
                            v17 = append(v17, 2147483647)
                        default:
                            var v19 uint8 = 251
                            _ = v19
                        }
                    case -2147483648:
                        switch -9223372036854775808 {
                        case 32767:
                            if (1 != int8(p0)) {
                                p0 = 42
                            }
                        case -686:
                            v17 = append(v17, -475)
                        }
                    }
                }
                switch 42 {
                case -444:
                    var v20 bool = (42 >= 9223372036854775807)
                    _ = v20
                case 46:
                    v21 := GenS4{F0: 92, F1: 302, F2: 42}
                    _ = v21
                case -1:
                    v22 := GenS1{F0: -605, F1: 127}
                    _ = v22
                }
            }
            _ = v17
        } else {
            fmt.Println(p1)
            v23 := GenS5{F0: false, F1: ("hello" + "foo bar"), F2: (^(p0))}
            _ = v23
        }
        _ = v16
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]string)
    var v3 bool = true
    v4 := make(map[string]uint32)
    var v24 uint16 = 42
    fmt.Println(v1, v1, v3)
    fmt.Println(v3)
    switch -2147483648 {
    case 115:
        v25 := GenS6{F0: 127}
        _ = v25
    case 128:
        for fi26 := 0; fi26 < 9; fi26++ {
            v27 := []int{(fi26 + fi26), 42}
            if v3 {
                if v3 {
                    v28 := GenS7{F0: fi26}
                    v1 = 0
                    _ = v28
                }
                fmt.Println(v3, fi26)
            }
            fmt.Println(v1, v3)
            _ = fi26
            _ = v27
        }
    default:
        v24 -= v24
    }
    delete(v4, "hello")
    fmt.Println(v24)
    v29 := GenS8{F0: 127}
    v24 += 1
    v30 := GenS9{F0: -1, F1: -773, F2: 0}
    var v31 int32 = -1
    v32 := []int{0, 42, -276}
    v33 := GenS6{F0: 233}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
