package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 string) int8 {
    fmt.Println(p0)
    if false {
        v5 := []int{32767, -62, 127, 776}
        _ = v5
    } else {
        v6 := []int{-1}
        _ = v6
    }
    var v7 uint8 = 127
    if ((false && (42 == 839)) || (true && (-334 >= 32768))) {
        var v8 uint16 = 246
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int64 = 0
            v10 = v10
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        fmt.Println(p0, v7)
        var v11 float64 = (3.14159 - ((float64((0.1 + (1e10 / -2.25))) - 1.5) / (0.0 + (-2.25 - 0.0))))
        _ = v11
    }
    switch 661 {
    case -637:
        v7 = (v7 ^ (v7 + v7))
    case 127:
        if true {
            if true {
                v12 := []int{42, -9223372036854775808, 1}
                var v13 int16 = 862
                _ = v12
                _ = v13
            } else {
                v14 := []int{32768, 0}
                _ = v14
            }
            fmt.Println(v7, v7)
        }
    case -8:
        p0 = ("\n\t" + (p0 + "\n\t"))
    }
    _ = p0
    _ = v7
    return 100
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{462, 4294967295, 65535, 65535}
    v3 := make(map[int]int)
    var v4 uint = 255
    if (42 < int8(v4)) {
        var v15 int = 127
        _ = v15
    }
    var v16 bool = (!(127 < 2))
    switch 297 {
    case 341:
        v1 += v1
    case -1:
        var v17 string = "a"
        _ = v17
    case 0:
        v2 = append(v2, 134)
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int = fi18
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            switch 109 {
            case -9223372036854775808:
                switch "\n\t" {
                case "foo bar":
                    var v22 uint64 = 1
                    _ = v22
                case "":
                    if true {
                        v1 = 2
                        switch (-(fi21)) {
                        case 42:
                            switch (fi21 & int(v1)) {
                            case 9223372036854775807:
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                            }
                        default:
                            fmt.Println(fi21, fi18)
                        }
                        v23 := GenS1{F0: -1, F1: fi21, F2: 2}
                        var v24 bool = ((!v16) && (!false))
                        _ = v23
                        _ = v24
                    }
                    v1 = 65535
                    v25 := GenS2{F0: 1}
                    _ = v25
                }
            }
            v2 = append(v2, 331)
            var v26 uint16 = v1
            _ = fi21
            _ = v26
        }
        v27 := GenS2{F0: 0}
        _ = fi18
        _ = v27
    }
    if len(v2) > 0 {
        v2[0] = 423
    }
    fmt.Println(v4, v1, v4)
    for fi28 := 0; fi28 < 13; fi28++ {
        if (-1 < 0) {
            for fi29 := 0; fi29 < 6; fi29++ {
                switch "a" {
                case "a":
                    var v30 int32 = 2147483647
                    _ = v30
                }
                _ = fi29
            }
            v16 = (4294967295 > 885)
        }
        _ = fi28
    }
    if (63 > 127) {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 int32 = 2147483647
            _ = fi31
            _ = v32
        }
        v4 = v4
        var v33 uint64 = 18446744073709551615
        _ = v33
    } else {
        fmt.Println(v1, v1)
    }
    v16 = (32767 > -9223372036854775808)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v34 int64 = 42
    for fi35 := 0; fi35 < 15; fi35++ {
        _ = len(v3)
        _ = fi35
    }
    _ = len(v3)
    v36 := GenS3{F0: 0}
    fmt.Println(v16, v34)
    v16 = (v16 && (!(255 != 46)))
    v16 = (((v4 + v4) <= 2) && true)
    v34 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
}
