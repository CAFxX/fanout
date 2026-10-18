package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint32; F1 string; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 int8 = 127
    var v4 uint = 18446744073709551615
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{int(v4), (fi6 ^ 1), 128, fi6}
        _ = fi6
        _ = v7
    }
    switch -995 {
    case 4294967295:
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(fi8)
            v9 := GenS1{F0: (v4 << (((v4 << (((^(v4))) & 63))) & 63))}
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int64 = -9223372036854775808
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    case 2147483647:
        v3 = v3
    case -379:
        v12 := GenS2{F0: v4, F1: -9223372036854775808}
        _ = v12
    default:
        v13 := []int{-696, 66, -737, 4294967295}
        _ = v13
    }
    v3 = 86
    v3 += (v3 << ((978) & 63))
    v14 := GenS3{F0: 42}
    v15 := GenS1{F0: 272}
    var v16 int16 = -936
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v3, fi17)
        switch fi17 {
        case 420:
            var v18 uint64 = 42
            _ = v18
        case 32767:
            v19 := GenS2{F0: v4, F1: int64(v16)}
            _ = v19
        case 2147483647:
            switch ("a" + "") {
            case "":
                var v20 int32 = 516
                _ = v20
            case "foo bar":
                var v21 uint32 = 2147483648
                _ = v21
            case "hello":
                var v22 uint64 = 1
                _ = v22
            }
        default:
            if (!(float32(1e10) == 1e10)) {
                var v23 uint32 = 2
                var v24 int32 = 42
                fmt.Println(v4)
                _ = v23
                _ = v24
            } else {
                var v25 int64 = 265
                var v26 uint32 = 2
                if true {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 string = "hello"
                        if (634 != v4) {
                            var v29 uint = (v4 - (v4 + (v4 + v4)))
                            _ = v29
                        } else {
                            delete(v1, (v26 + 65535))
                            v30 := GenS4{F0: 0, F1: v28}
                            _ = v30
                        }
                        v16 -= (v16 >> ((1) & 63))
                        _ = fi27
                        _ = v28
                    }
                } else {
                    var v31 int8 = int8(v4)
                    v32 := GenS4{F0: 18446744073709551615, F1: ("a" + "a")}
                    _ = v31
                    _ = v32
                }
                _ = v25
                _ = v26
            }
        }
        if true {
            v33 := []int{}
            _ = v33
        }
        _ = fi17
    }
    fmt.Println(v3, v3)
    if true {
        var v34 bool = (0 == 4294967295)
        _ = v34
    } else {
        v4 += (v4 * (v4 >> ((v4) & 63)))
        var v35 int16 = (-(v16))
        var v36 int = 42
        _ = v35
        _ = v36
    }
    fmt.Println(v4, v4)
    for fi37 := 0; fi37 < 7; fi37++ {
        fmt.Println(v3)
        switch (fi37 % ((fi37) | 1)) {
        case 42:
            if true {
                v38 := []int{fi37, -2147483648, (^(fi37))}
                _ = v38
            } else {
                v3 -= (v3 + (v3 + -128))
            }
        case 127:
            delete(v2, "foo bar")
        case 549:
            v4 = (v4 | v4)
        default:
            v3 = -90
        }
        if false {
            v4 = (^(v4))
        } else {
            var v39 uint64 = 2
            var v40 uint16 = 919
            _ = v39
            _ = v40
        }
        _ = fi37
    }
    var v41 bool = (!false)
    switch "foo bar" {
    case "":
        v42 := GenS5{F0: 887, F1: "", F2: v16}
        _ = v42
    default:
        switch -1 {
        case -2147483648:
            var v43 int = -9223372036854775808
            _ = v43
        case 263:
            v44 := []int{-2147483648, 2147483647, 1, 4294967295}
            _ = v44
        case -403:
            var v45 uint32 = 2147483647
            _ = v45
        default:
            v4 = v4
        }
    }
    v41 = false
    var v46 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v5
}
