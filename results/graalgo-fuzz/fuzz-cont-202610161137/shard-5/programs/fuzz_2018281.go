package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int8; F1 uint32; F2 string }
type GenS8 struct { F0 uint64; F1 int32; F2 int }
func genFn0() int8 {
    var v7 uint8 = 0
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint32 = 2147483647
        v7 = v7
        var v10 int64 = -9223372036854775808
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{9223372036854775807, -233, 128, 2147483647}
    switch ("hello" + "hello") {
    case "":
        if (((2 <= 32767) && (42 == 1)) || (!(65535 != 42))) {
            fmt.Println(v7)
            var v12 uint = 65535
            if len(v11) > 0 {
                v11[0] = 128
            }
            _ = v12
        }
    case "hello":
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := GenS1{F0: uint32(v7), F1: -644, F2: 32767}
            _ = fi13
            _ = v14
        }
    case "\n\t":
        v7 = (v7 << ((2) & 63))
    default:
        if true {
            var v15 int = 32768
            for fi16 := 0; fi16 < 3; fi16++ {
                v17 := GenS1{F0: 42, F1: fi16, F2: 1}
                if (!true) {
                    v18 := GenS2{F0: int32(fi16), F1: 42}
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
            v11 = append(v11, v15)
            _ = v15
        } else {
            if true {
                v19 := GenS3{F0: 42, F1: 32768}
                _ = v19
            } else {
                fmt.Println(v7)
                v7 += (v7 % ((1) | 1))
                var v20 float64 = (float64(1e10) + 1e10)
                _ = v20
            }
            v11 = append(v11, 4294967295)
        }
    }
    _ = v11
    _ = v7
    return 100
}

func main() {
    v1 := []int{98, 1, 32767, 4294967295}
    var v2 int64 = 42
    v3 := []int{-1, 0, 734}
    var v4 bool = (1 == 2147483647)
    var v5 int8 = 96
    v6 := []string{"", ("hello" + "foo bar"), "\n\t", ("a" + "\n\t")}
    var v21 int = 127
    if (v4 && (false || ((true || v4) || v4))) {
        var v22 float64 = 1.5
        if (67 >= 184) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v22
    }
    if (-1 != 1) {
        v5 = v5
        v1 = append(v1, 128)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v4 = true
    switch v21 {
    case 128:
        for fi23 := 0; fi23 < 19; fi23++ {
            fmt.Println(v5, v5, v5)
            switch "hello" {
            case "a":
                if v4 {
                    var v24 string = ""
                    _ = v24
                } else {
                    switch (("a" + "\n\t") + "a") {
                    case "":
                        fmt.Println(v21, v2)
                    }
                    switch 767 {
                    case 32768:
                        v25 := GenS4{F0: 255, F1: (v5 >> ((42) & 63))}
                        _ = v25
                    default:
                        var v26 uint8 = 42
                        _ = v26
                    }
                }
            }
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 string = "foo bar"
                v1 = append(v1, fi27)
                _ = fi27
                _ = v28
            }
            _ = fi23
        }
    case -186:
        var v29 uint64 = 18446744073709551615
        _ = v29
    }
    v21 = v21
    var v30 int32 = -1
    v4 = (0 >= 255)
    switch (("hello" + "a") + "") {
    case "\n\t":
        v31 := GenS5{F0: 0, F1: v2}
        _ = v31
    }
    switch "\n\t" {
    case "a":
        if false {
            v2 = (v2 / (((v2 + 0)) | 1))
            v2 += v2
            v30 -= 42
        } else {
            if ((v30 ^ v30) >= -13) {
                fmt.Println(v5)
                var v32 int16 = 42
                _ = v32
            } else {
                v30 -= (v30 << ((18446744073709551615) & 63))
                v33 := GenS4{F0: 148, F1: v5}
                _ = v33
            }
        }
    }
    v2 -= 0
    var v34 bool = (431 > 65535)
    switch v21 {
    case 32768:
        v30 += int32(v21)
    case 42:
        if v4 {
            for fi35 := 0; fi35 < 17; fi35++ {
                if ((fi35 / ((fi35) | 1)) == (v21 % ((v21) | 1))) {
                    var v36 bool = (v4 || v34)
                    var v37 int8 = v5
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 12; fi38++ {
                    v5 += v5
                    v39 := GenS6{F0: 0, F1: int16(v2)}
                    var v40 uint64 = 18446744073709551615
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                var v41 int32 = (v30 | v30)
                _ = fi35
                _ = v41
            }
        }
    default:
        if v4 {
            var v42 uint = 2
            _ = v42
        }
    }
    v43 := GenS7{F0: (v5 ^ v5), F1: 4294967295, F2: "a"}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v21 -= v21
    switch v21 {
    case 2147483647:
        v4 = (1 == 170)
    case -2147483648:
        v4 = true
    default:
        var v44 uint = 4294967295
        _ = v44
    }
    for fi45 := 0; fi45 < 7; fi45++ {
        v5 += v5
        fmt.Println(v5)
        v46 := GenS8{F0: 18446744073709551615, F1: (v30 ^ v30), F2: int(v30)}
        _ = fi45
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
