package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 bool }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint16, p1 string, p2 int64) uint64 {
    p1 = p1
    if true {
        p0 = 0
        fmt.Println(p2)
        if false {
            if true {
                v7 := GenS1{F0: (!(false || true)), F1: 244}
                _ = v7
            }
            var v8 int64 = (p2 + -319)
            v9 := []int{1}
            _ = v8
            _ = v9
        } else {
            var v10 uint8 = 2
            fmt.Println(p2)
            switch p1 {
            case "\n\t":
                if (204 > 775) {
                    var v11 uint16 = (p0 / (((p0 << ((0) & 63))) | 1))
                    switch 1 {
                    case -1:
                        var v12 int = 0
                        _ = v12
                    case -2147483648:
                        for fi13 := 0; fi13 < 15; fi13++ {
                            var v14 int64 = p2
                            p0 = v11
                            p1 = p1
                            _ = fi13
                            _ = v14
                        }
                    }
                    v15 := []int{-9223372036854775808, 441}
                    _ = v11
                    _ = v15
                }
            case "a":
                fmt.Println(v10, p2)
            }
            _ = v10
        }
    }
    if false {
        var v16 int64 = (^(p2))
        var v17 uint64 = 0
        if false {
            v16 = p2
        }
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            switch (fi18 << ((0) & 63)) {
            case -9223372036854775808:
                v19 := GenS1{F0: (true || (2 <= 179)), F1: 0}
                _ = v19
            case 42:
                var v20 int64 = 42
                _ = v20
            case 9223372036854775807:
                var v21 int8 = 127
                _ = v21
            default:
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 uint = 2
                    v24 := GenS2{F0: -128, F1: fi18, F2: (!(!true))}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            p1 = p1
            _ = fi18
        }
        var v25 float32 = float32(1.5)
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 int16 = 32767
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := []int{4294967295}
            var v30 int32 = 2147483647
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    var v31 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return 42
}
func genFn1() uint32 {
    var v32 int = 0
    fmt.Println(v32)
    v32 = v32
    _ = v32
    return 2
}

func main() {
    var v1 uint16 = 0
    v2 := []string{"a", "\n\t", "\n\t"}
    var v3 float32 = 3.14159
    v4 := make(map[int]string)
    var v5 uint16 = v1
    var v6 float32 = v3
    v33 := GenS1{F0: true, F1: 207}
    for fi34 := 0; fi34 < 10; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            if true {
                fmt.Println(fi35, v3, fi34)
            }
            switch -2147483648 {
            case -2147483648:
                var v36 float32 = ((v6 / -2.25) + v6)
                _ = v36
            case 4294967295:
                v1 -= v5
            case 32768:
                v37 := GenS3{F0: 2, F1: 1, F2: 2147483648}
                _ = v37
            }
            _ = fi35
        }
        _ = fi34
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        if (502 < 42) {
            v39 := GenS4{F0: -1, F1: 179}
            _ = v39
        }
        v40 := GenS3{F0: 18446744073709551615, F1: -2147483648, F2: 4294967295}
        _ = fi38
        _ = v40
    }
    v41 := GenS3{F0: 18446744073709551615, F1: -47, F2: 4294967295}
    fmt.Println(v1)
    var v42 int8 = 1
    if true {
        v43 := GenS1{F0: ("a" >= "a"), F1: 29}
        _ = v43
    }
    v42 -= v42
    v44 := GenS2{F0: 0, F1: 42, F2: true}
    v45 := GenS5{F0: -1}
    var v46 int32 = -2147483648
    var v47 int16 = 0
    var v48 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
