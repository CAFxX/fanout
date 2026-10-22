package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]int)
    var v3 bool = (2147483647 >= -2147483648)
    fmt.Println(v1, v3)
    if false {
        var v4 int8 = 100
        _ = v4
    }
    var v5 int16 = -32768
    if ((true || ((false && (v3 && (!((42 != 952) && (311 != 2))))) && v3)) && true) {
        switch -2147483648 {
        case 0:
            v6 := GenS1{F0: 2147483648, F1: 2}
            _ = v6
        case 2147483647:
            v3 = (true && (!(1.5 != v1)))
        }
        v3 = (42 <= 2)
    } else {
        switch ("hello" + (("hello" + ((("\n\t" + "hello") + "") + "foo bar")) + "\n\t")) {
        case "hello":
            switch "a" {
            case "a":
                v5 += (v5 + v5)
            case "hello":
                if v3 {
                    if v3 {
                        var v7 int64 = -1
                        v3 = v3
                        v8 := []int{42, 0, 42}
                        _ = v7
                        _ = v8
                    }
                    var v9 int64 = -1
                    v2[-991] = 42
                    _ = v9
                } else {
                    var v10 int16 = (v5 * v5)
                    switch 1 {
                    case 9223372036854775807:
                        fmt.Println(v1, v3)
                    default:
                        fmt.Println(v3)
                    }
                    var v11 float64 = 0.1
                    _ = v10
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 bool = (2147483648 < 2147483647)
                    if (1 <= 0) {
                        v5 -= v5
                    } else {
                        fmt.Println(v5, v1, v1)
                    }
                    var v14 int32 = -2147483648
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        case "":
            var v15 bool = (!(79 != -685))
            _ = v15
        case "a":
            var v16 float32 = 1.5
            _ = v16
        default:
            v17 := []int{-874, -1}
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            switch fi18 {
            case -2147483648:
                var v19 uint16 = 180
                _ = v19
            case 42:
                switch (fi18 * fi18) {
                case 42:
                    v3 = v3
                case -1:
                    if (128 <= 0) {
                        fmt.Println(v1, fi18, v1)
                        for fi20 := 0; fi20 < 4; fi20++ {
                            v3 = v3
                            v5 -= (v5 << ((42) & 63))
                            _ = fi20
                        }
                    }
                case -9223372036854775808:
                    if v3 {
                        _ = len(v2)
                        v21 := []int{fi18, fi18}
                        _ = v21
                    }
                default:
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v5 = (v5 + v5)
                        var v23 int8 = -128
                        _ = fi22
                        _ = v23
                    }
                }
            default:
                switch -66 {
                case -859:
                    fmt.Println(v5, v1, fi18)
                case 2147483647:
                    v24 := []int{fi18}
                    _ = v24
                }
            }
            _ = fi18
        }
    }
    var v25 bool = ((2147483647 < 128) && true)
    v2[32767] = -9223372036854775808
    v5 += (-(v5))
    v26 := GenS1{F0: 2, F1: 0}
    v27 := []int{1, 4294967295, -750, 858}
    var v28 int64 = 478
    if v25 {
        v29 := GenS2{F0: 32768, F1: 28}
        v30 := GenS3{F0: 912, F1: 1}
        if false {
            var v31 uint = 65535
            var v32 int8 = -128
            _ = v31
            _ = v32
        } else {
            switch "hello" {
            case "\n\t":
                v28 += (v28 / ((v28) | 1))
            case "hello":
                v33 := GenS4{F0: (2 <= 4294967295), F1: -2147483648}
                _ = v33
            default:
                switch 228 {
                case 1:
                    if (!v25) {
                        switch 1 {
                        case 1:
                            fmt.Println(v25, v1)
                        case 128:
                            v34 := GenS1{F0: 189, F1: 42}
                            _ = v34
                        default:
                            var v35 int64 = (v28 * 42)
                            _ = v35
                        }
                        var v36 uint16 = 2
                        _ = v36
                    } else {
                        if (v25 && v3) {
                            if false {
                                fmt.Println(v25, v3, v28)
                            } else {
                                fmt.Println(v28)
                            }
                            v37 := GenS5{F0: -651, F1: "\n\t"}
                            _ = v37
                        }
                        v38 := GenS6{F0: 231}
                        _ = v38
                    }
                case 4294967295:
                    v5 = (v5 % ((v5) | 1))
                case -9223372036854775808:
                    for fi39 := 0; fi39 < 17; fi39++ {
                        v40 := GenS5{F0: (^(fi39)), F1: ("foo bar" + ("a" + "foo bar"))}
                        for fi41 := 0; fi41 < 2; fi41++ {
                            v42 := GenS5{F0: fi39, F1: "a"}
                            _ = fi41
                            _ = v42
                        }
                        _ = fi39
                        _ = v40
                    }
                default:
                    v1 += v1
                }
            }
        }
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v5
}
