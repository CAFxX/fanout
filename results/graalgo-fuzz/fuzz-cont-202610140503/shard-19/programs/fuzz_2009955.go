package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint; F2 int16 }
type GenS5 struct { F0 uint32 }
func genFn0() string {
    v7 := GenS1{F0: (false || true)}
    var v8 uint8 = 128
    if (0 >= 42) {
        fmt.Println(v8)
        var v9 uint8 = (v8 ^ 128)
        v10 := GenS1{F0: true}
        _ = v10
        _ = v9
    } else {
        v11 := GenS2{F0: 42, F1: 42}
        _ = v11
    }
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 int32) uint8 {
    if (!(((146 >= 815) || false) || (false || false))) {
        if (!((((true && (1 > -12)) && false) || false) || true)) {
            var v12 uint64 = 737
            p0 = (p0 ^ (p0 ^ (p0 % ((-1) | 1))))
            for fi13 := 0; fi13 < 15; fi13++ {
                p0 = p0
                var v14 float32 = 1.5
                v15 := []int{642, 4294967295, 2147483647, -704}
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        } else {
            v16 := GenS2{F0: -32768, F1: 4294967295}
            var v17 uint16 = 0
            _ = v16
            _ = v17
        }
        v18 := []int{776, -809, 323, 128}
        for fi19 := 0; fi19 < 20; fi19++ {
            p0 = -2147483648
            _ = fi19
        }
        _ = v18
    }
    if (!true) {
        fmt.Println(p0)
        for fi20 := 0; fi20 < 13; fi20++ {
            if (p0 >= p0) {
                v21 := []int{42, 4294967295, 0}
                _ = v21
            } else {
                var v22 int16 = -1
                _ = v22
            }
            _ = fi20
        }
        v23 := GenS3{F0: 398}
        _ = v23
    } else {
        switch "\n\t" {
        case "hello":
            switch 9223372036854775807 {
            case 42:
                v24 := []int{4294967295, -204, -9223372036854775808}
                _ = v24
            case -107:
                if true {
                    v25 := []int{9223372036854775807}
                    _ = v25
                }
            case 2147483647:
                switch ("hello" + "foo bar") {
                case "":
                    v26 := []int{42, -2147483648, -587, -293}
                    _ = v26
                case "\n\t":
                    if true {
                        var v27 int8 = 99
                        for fi28 := 0; fi28 < 16; fi28++ {
                            var v29 bool = (((v27 & (v27 << ((4294967295) & 63))) > (v27 % (((v27 ^ (v27 >> ((255) & 63)))) | 1))) && false)
                            for fi30 := 0; fi30 < 2; fi30++ {
                                var v31 float32 = 0.0
                                _ = fi30
                                _ = v31
                            }
                            _ = fi28
                            _ = v29
                        }
                        _ = v27
                    }
                case "foo bar":
                    p0 -= (p0 - -695)
                default:
                    fmt.Println(p0)
                }
            }
        case "foo bar":
            switch 1 {
            case 350:
                v32 := []int{-579, -1}
                _ = v32
            }
        case "\n\t":
            switch 890 {
            case 4294967295:
                p0 = -220
            case -85:
                switch -1 {
                case -912:
                    for fi33 := 0; fi33 < 11; fi33++ {
                        v34 := []int{}
                        _ = fi33
                        _ = v34
                    }
                }
            default:
                fmt.Println(p0)
            }
        }
        p0 = (p0 * 705)
    }
    switch -2147483648 {
    case 944:
        v35 := GenS4{F0: 54, F1: 753, F2: 216}
        _ = v35
    case -777:
        p0 = (p0 / ((2147483647) | 1))
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    var v4 int32 = 406
    v5 := make(map[string]int)
    var v6 uint = 255
    fmt.Println(v4, v1)
    if true {
        switch ((("hello" + "") + ("\n\t" + "foo bar")) + (((("" + "a") + ("" + "foo bar")) + "a") + "\n\t")) {
        case "":
            fmt.Println(v1, v1, v1)
        }
        fmt.Println(v6, v1)
        switch 127 {
        case 127:
            v36 := []int{}
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        v1 -= 160
        _ = len(v3)
        _ = fi37
    }
    v38 := GenS3{F0: 0}
    v39 := GenS2{F0: 0, F1: 1}
    v4 = 0
    fmt.Println(v4)
    v1 += genFn1(v4)
    v40 := GenS3{F0: 1}
    v41 := GenS5{F0: 455}
    for fi42 := 0; fi42 < 13; fi42++ {
        fmt.Println(v4, v1, fi42)
        var v43 uint64 = 88
        var v44 uint = v6
        _ = fi42
        _ = v43
        _ = v44
    }
    var v45 uint8 = (v1 & (v1 | v1))
    v46 := GenS1{F0: (false || false)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
