package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint64; F1 uint16; F2 string }
func genFn0(p0 uint32, p1 int32) int64 {
    switch ("\n\t" + (("" + "hello") + "\n\t")) {
    case "\n\t":
        var v4 uint16 = 0
        _ = v4
    case "foo bar":
        for fi5 := 0; fi5 < 2; fi5++ {
            if (-2.25 == -2.25) {
                if (559 <= 18446744073709551615) {
                    for fi6 := 0; fi6 < 10; fi6++ {
                        switch (fi6 - (fi6 ^ (-(fi6)))) {
                        case 127:
                            v7 := GenS1{F0: 1, F1: 2, F2: 0}
                            if true {
                                v8 := []int{fi6, -2147483648, -9223372036854775808, 128}
                                var v9 uint8 = 127
                                fmt.Println(v9, p0)
                                _ = v8
                                _ = v9
                            }
                            if ((!false) && true) {
                                var v10 float32 = ((0.1 * -2.25) * 3.14159)
                                _ = v10
                            }
                            _ = v7
                        case 42:
                            v11 := []int{}
                            _ = v11
                        case -38:
                            var v12 string = ("" + (("a" + (("hello" + "foo bar") + "foo bar")) + ""))
                            _ = v12
                        }
                        _ = fi6
                    }
                }
                if ((100 > 84) && ((^(p1)) != (p1 >> ((1) & 63)))) {
                    p1 -= (p1 - (p1 + (p1 % ((p1) | 1))))
                    fmt.Println(p1, p0)
                } else {
                    fmt.Println(p1, p0, p0)
                    if (false && false) {
                        v13 := GenS2{F0: 342}
                        p1 += p1
                        v14 := GenS1{F0: 255, F1: 65535, F2: 703}
                        _ = v13
                        _ = v14
                    }
                }
                v15 := GenS2{F0: 2}
                _ = v15
            } else {
                v16 := []int{(fi5 % (((fi5 + 128)) | 1)), (fi5 << ((42) & 63))}
                fmt.Println(fi5, fi5, p0)
                _ = v16
            }
            _ = fi5
        }
    }
    var v17 int32 = p1
    v18 := []int{-1}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 42
}

func main() {
    var v1 int64 = -1
    var v2 bool = (false && true)
    v3 := make(map[string]uint32)
    v3[""] = 977
    fmt.Println(v1, v2)
    switch ("" + ("\n\t" + "a")) {
    case "a":
        v19 := GenS2{F0: 42}
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 string = ""
        v22 := []int{(fi20 & (fi20 + fi20))}
        _ = fi20
        _ = v21
        _ = v22
    }
    switch "\n\t" {
    case "foo bar":
        switch -9223372036854775808 {
        case 0:
            v1 = v1
        case 4294967295:
            v1 -= (v1 % ((-9223372036854775808) | 1))
        }
    }
    if v2 {
        switch "a" {
        case "":
            switch "\n\t" {
            case "hello":
                if (!((v2 || ((0.1 / 3.14159) <= (0.0 / (float32(3.14159) + 1.5)))) && v2)) {
                    var v23 uint16 = 32767
                    var v24 int32 = -2147483648
                    _ = v23
                    _ = v24
                }
            case "\n\t":
                for fi25 := 0; fi25 < 4; fi25++ {
                    delete(v3, "")
                    _ = fi25
                }
            default:
                v26 := GenS3{F0: 1}
                _ = v26
            }
        case "hello":
            var v27 int32 = 42
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 10; fi28++ {
                v2 = (-379 < 0)
                var v29 float32 = (3.14159 - 0.0)
                v29 -= 1e10
                _ = fi28
                _ = v29
            }
        }
    }
    v30 := GenS1{F0: 42, F1: 512, F2: 1}
    switch "" {
    case "hello":
        v31 := GenS4{F0: 3}
        _ = v31
    default:
        v2 = true
    }
    v32 := GenS5{F0: (!((v1 + (-(v1))) <= (v1 >> ((42) & 63)))), F1: 128, F2: 42}
    for fi33 := 0; fi33 < 3; fi33++ {
        for fi34 := 0; fi34 < 12; fi34++ {
            v35 := GenS6{F0: 894, F1: 42, F2: ("\n\t" + "\n\t")}
            var v36 string = ("" + "")
            _ = len(v3)
            fmt.Println(v2, fi33, fi33)
            var v37 uint = 111
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        _ = fi33
    }
    if false {
        v1 = 593
        var v38 bool = v2
        v1 = (v1 * (v1 + (v1 & (-(v1)))))
        _ = v38
    } else {
        v39 := GenS3{F0: -1}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
}
