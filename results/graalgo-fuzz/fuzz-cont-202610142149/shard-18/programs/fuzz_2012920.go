package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 int32, p1 int16, p2 uint64) int16 {
    v4 := []int{int(p1), 0, 4294967295, -2147483648}
    fmt.Println(p2, p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        switch (fi5 & (^(fi5))) {
        case 9223372036854775807:
            switch "\n\t" {
            case "foo bar":
                p0 += p0
            case "\n\t":
                var v6 uint64 = (p2 * (p2 << ((18446744073709551615) & 63)))
                _ = v6
            default:
                var v7 int = fi5
                _ = v7
            }
        case -2147483648:
            if false {
                p0 += (p0 - (p0 << ((2) & 63)))
                p0 -= 413
            } else {
                fmt.Println(p2)
                switch ("foo bar" + "a") {
                case "a":
                    for fi8 := 0; fi8 < 8; fi8++ {
                        p0 += (^(p0))
                        _ = fi8
                    }
                case "\n\t":
                    p0 += p0
                default:
                    v4 = append(v4, (fi5 % ((32767) | 1)))
                }
            }
        case 42:
            for fi9 := 0; fi9 < 2; fi9++ {
                p1 = 1
                fmt.Println(p1, p2, fi9)
                _ = fi9
            }
        }
        _ = fi5
    }
    var v10 uint64 = (^(p2))
    var v11 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v4
    return 0
}
func genFn1(p0 int8, p1 uint8, p2 uint) uint32 {
    v12 := GenS1{F0: 0}
    for fi13 := 0; fi13 < 14; fi13++ {
        p2 += p2
        var v14 uint = 42
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 18446744073709551615
    if (1 < (p1 % ((p1) | 1))) {
        var v16 uint32 = 65535
        switch "" {
        case "a":
            p2 -= (p2 & p2)
        case "":
            p2 = (p2 << ((p2) & 63))
        case "\n\t":
            fmt.Println(p2, p2, v16)
        }
        _ = v16
    } else {
        v15 += v15
        var v17 float32 = 1.5
        v17 -= ((v17 * (v17 * v17)) * v17)
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(p1, p2, p2)
        p0 -= (p0 ^ -1)
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v15
    return 1
}

func main() {
    v1 := []string{"foo bar", "a", "a"}
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v19 int16 = -442
    v19 += v19
    var v20 float32 = 1e10
    if true {
        var v21 uint8 = 42
        _ = len(v2)
        _ = v21
    } else {
        v20 = (v20 - (v20 * (((v20 / -2.25) / 1e10) + v20)))
        for fi22 := 0; fi22 < 1; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                v19 = (v19 + (v19 & v19))
                _ = fi23
            }
            var v24 int16 = 1000
            _ = fi22
            _ = v24
        }
        v25 := v3[""]
        _ = v25
    }
    switch -2147483648 {
    case 42:
        if true {
            switch "" {
            case "":
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 uint8 = 255
                    fmt.Println(fi26, v20, v27)
                    _ = fi26
                    _ = v27
                }
            case "hello":
                fmt.Println(v20, v19)
            }
        } else {
            if ("hello" < ("a" + "a")) {
                v19 = v19
            }
            v28 := GenS1{F0: 578}
            v29 := GenS1{F0: 18446744073709551615}
            _ = v28
            _ = v29
        }
    case 127:
        var v30 uint = 4294967295
        _ = v30
    }
    v31 := GenS2{F0: -1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v19)
    v32 := GenS3{F0: -1, F1: 0}
    switch 0 {
    case -954:
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    case -1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if true {
            v20 = 0.1
            v33 := GenS3{F0: 42, F1: -2147483648}
            _ = v33
        } else {
            fmt.Println(v20)
            _ = len(v3)
            var v34 int = 127
            _ = v34
        }
    }
    var v35 uint64 = 4294967295
    v36 := GenS2{F0: -1}
    var v37 uint64 = 422
    if false {
        var v38 bool = (130 <= 127)
        _ = v38
    } else {
        var v39 uint = 639
        var v40 bool = ((("\n\t" + "foo bar") + "hello") <= (("foo bar" + "a") + ""))
        v40 = false
        _ = v39
        _ = v40
    }
    if ((128 > 2147483647) && (!true)) {
        switch -9223372036854775808 {
        case 829:
            _ = len(v3)
        case 32768:
            v35 = v35
        }
    }
    fmt.Println(v35, v35)
    fmt.Println(v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v37
}
