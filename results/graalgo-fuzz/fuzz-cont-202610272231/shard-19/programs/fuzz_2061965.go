package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 string }
type GenS5 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0(p0 int8, p1 uint64) int {
    var v7 int64 = 42
    v7 -= v7
    for fi8 := 0; fi8 < 14; fi8++ {
        if false {
            p1 = 972
            switch "a" {
            case "":
                var v9 uint = 178
                _ = v9
            }
            fmt.Println(v7, p0, p0)
        } else {
            switch "\n\t" {
            case "foo bar":
                if true {
                    v7 = 42
                } else {
                    v10 := []int{fi8, 4294967295, fi8}
                    _ = v10
                }
            default:
                v7 = (v7 / ((v7) | 1))
            }
        }
        var v11 int32 = -2147483648
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1() uint {
    var v12 int = genFn0(-119, 18446744073709551615)
    if ((true || false) || (486 > 426)) {
        var v13 uint16 = 32767
        v14 := []int{(-(v12)), -48, v12, -823}
        var v15 uint16 = (^(v13))
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 uint64 = 292
        switch "a" {
        case "":
            v17 := []int{v12, v12, (v12 + (^(v12))), 2147483647}
            _ = v17
        case "a":
            if (442 != 1) {
                v18 := GenS1{F0: 1}
                _ = v18
            }
        }
        _ = v16
    }
    fmt.Println(v12)
    _ = v12
    return 1
}

func main() {
    var v1 int64 = 0
    v2 := []uint32{1}
    var v3 uint32 = 479
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    var v6 int64 = v1
    var v19 int = 4294967295
    switch (("foo bar" + ("" + "foo bar")) + "\n\t") {
    case "a":
        var v20 uint16 = 809
        _ = v20
    }
    var v21 int32 = 2147483647
    v1 = (v6 / ((v1) | 1))
    for fi22 := 0; fi22 < 5; fi22++ {
        if (0.1 >= (float64(-2.25) * float64(3.14159))) {
            v23 := GenS2{F0: (v3 - (v3 | (v3 % ((v3) | 1)))), F1: -405}
            v24 := GenS1{F0: 0}
            _ = v23
            _ = v24
        } else {
            v6 = v1
            if (!true) {
                delete(v4, (v3 - v3))
            } else {
                if true {
                    if true {
                        fmt.Println(fi22)
                        fmt.Println(v3, v6, v6)
                    } else {
                        var v25 uint32 = (v3 << ((18446744073709551615) & 63))
                        _ = v25
                    }
                    if false {
                        switch v19 {
                        case -250:
                            var v26 int16 = 42
                            _ = v26
                        case 4294967295:
                            var v27 uint8 = 250
                            _ = v27
                        }
                        switch v19 {
                        case 32768:
                            var v28 uint64 = 18446744073709551615
                            _ = v28
                        case 825:
                            var v29 bool = true
                            _ = v29
                        case 0:
                            fmt.Println(v19)
                        }
                        var v30 bool = (v3 > (^(v3)))
                        _ = v30
                    } else {
                        var v31 bool = (false || false)
                        _ = len(v2)
                        _ = v31
                    }
                }
                v32 := GenS3{F0: (^(v1)), F1: (v3 - 2147483648)}
                fmt.Println(fi22)
                _ = v32
            }
            if ((32767 <= 42) || true) {
                switch 9223372036854775807 {
                case 128:
                    var v33 bool = false
                    _ = v33
                }
            }
        }
        v21 = (^(v21))
        v34 := GenS3{F0: v1, F1: (v3 % ((v3) | 1))}
        _ = fi22
        _ = v34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        fmt.Println(v19, v21)
        v1 = v1
        _ = fi35
    }
    fmt.Println(v3, v3, v3)
    v6 -= (^(v1))
    for fi36 := 0; fi36 < 8; fi36++ {
        if (!(42 < 1000)) {
            v21 += 2147483647
            if (!false) {
                var v37 string = ""
                v1 = (v1 - v1)
                _ = v37
            }
        }
        var v38 float32 = 0.1
        switch "a" {
        case "hello":
            v5["foo bar"] = (v3 % ((v3) | 1))
        default:
            v39 := GenS4{F0: 65535, F1: -32768, F2: ""}
            _ = v39
        }
        _ = fi36
        _ = v38
    }
    if ((4294967295 >= 4294967295) || false) {
        var v40 uint32 = v3
        v19 = -687
        if false {
            fmt.Println(v21, v1)
            fmt.Println(v21, v6)
            for fi41 := 0; fi41 < 17; fi41++ {
                v6 = (v6 ^ (v1 * v1))
                var v42 int = (fi41 + (v19 >> ((0) & 63)))
                _ = fi41
                _ = v42
            }
        }
        _ = v40
    } else {
        var v43 uint8 = 128
        _ = len(v4)
        _ = v43
    }
    v44 := GenS5{F0: -1, F1: (-(v21)), F2: 42}
    fmt.Println(v6, v19, v21)
    v45 := GenS5{F0: -1, F1: 2147483647, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
