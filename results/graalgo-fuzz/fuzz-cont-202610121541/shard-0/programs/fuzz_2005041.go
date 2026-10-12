package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
type GenS3 struct { F0 uint32; F1 bool; F2 bool }
type GenS4 struct { F0 string; F1 bool; F2 uint8 }
func genFn0(p0 int, p1 int16) uint16 {
    var v6 float32 = 0.1
    if (0 > 42) {
        if (2147483648 != 2147483648) {
            switch "a" {
            case "foo bar":
                fmt.Println(p1)
            case "\n\t":
                var v7 int32 = -1
                _ = v7
            case "hello":
                v8 := []int{p0}
                _ = v8
            }
        } else {
            var v9 int = -781
            _ = v9
        }
    }
    switch ("\n\t" + "foo bar") {
    case "":
        p0 += p0
    case "foo bar":
        for fi10 := 0; fi10 < 2; fi10++ {
            switch ("foo bar" + "a") {
            case "":
                if (!(2 == 508)) {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        var v12 uint8 = 1
                        var v13 uint16 = 340
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    if true {
                        v14 := []int{(fi10 % (((fi10 | 697)) | 1)), (p0 << ((42) & 63)), -2147483648, -9223372036854775808}
                        p1 = p1
                        _ = v14
                    }
                } else {
                    fmt.Println(v6)
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v16 := GenS1{F0: 1, F1: 1}
                        _ = fi15
                        _ = v16
                    }
                    v6 = v6
                }
            }
            v6 += (v6 + 0.1)
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 1.5
    v3 := make(map[int]bool)
    var v4 int = -136
    var v5 int = (v4 * (v4 >> ((499) & 63)))
    fmt.Println(v5, v5)
    v17 := []int{v4, v5, 0, v5}
    switch v4 {
    case 0:
        switch "foo bar" {
        case "":
            v18 := GenS2{F0: 2147483648, F1: 127, F2: 65535}
            _ = v18
        case "\n\t":
            for fi19 := 0; fi19 < 11; fi19++ {
                for fi20 := 0; fi20 < 6; fi20++ {
                    if (-9223372036854775808 != 1) {
                        switch (v5 | v4) {
                        case 822:
                            v21 := GenS1{F0: 1, F1: 42}
                            _ = v21
                        }
                    }
                    _ = fi20
                }
                _ = fi19
            }
        }
    case 9223372036854775807:
        if (false && false) {
            var v22 uint8 = 255
            _ = v22
        } else {
            fmt.Println(v2)
        }
    case -109:
        for fi23 := 0; fi23 < 13; fi23++ {
            v24 := GenS2{F0: 42, F1: 0, F2: 697}
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 uint64 = 2
                if true {
                    v1[1] = 2
                    if false {
                        v2 = v2
                        var v27 uint64 = (^(v26))
                        _ = v27
                    } else {
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            var v28 int = fi23
                            _ = v28
                        case "\n\t":
                            v29 := GenS3{F0: 1, F1: (690 > 4294967295), F2: true}
                            _ = v29
                        case "hello":
                            if len(v17) > 0 {
                                v17[0] = fi25
                            }
                        }
                        var v30 uint = 65535
                        _ = v30
                    }
                }
                v31 := GenS1{F0: v26, F1: 4294967295}
                _ = fi25
                _ = v26
                _ = v31
            }
            switch 128 {
            case -1:
                v2 -= 1.5
            case 42:
                v5 += (-(fi23))
            default:
                v2 = 0.1
            }
            _ = fi23
            _ = v24
        }
    }
    v32 := GenS2{F0: 2147483647, F1: 108, F2: 4294967295}
    v2 = v2
    var v33 uint = 2
    v2 += (3.14159 / v2)
    v34 := GenS4{F0: "\n\t", F1: (false || true), F2: 182}
    v4 = v5
    v17 = append(v17, 117)
    if true {
        v35 := GenS2{F0: 993, F1: 127, F2: v33}
        _ = v35
    } else {
        v5 += v5
    }
    v36 := GenS3{F0: 117, F1: true, F2: false}
    _ = len(v17)
    var v37 int8 = 100
    v1[2] = 783
    var v38 string = ""
    if true {
        v33 = (v33 ^ v33)
    }
    var v39 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
