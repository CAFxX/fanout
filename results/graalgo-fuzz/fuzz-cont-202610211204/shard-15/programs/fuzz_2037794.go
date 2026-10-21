package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int64 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int32; F1 uint32; F2 uint16 }
func genFn0(p0 uint32, p1 int16) int {
    for fi7 := 0; fi7 < 12; fi7++ {
        p1 += p1
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(p0, fi7)
            var v9 uint64 = 2
            switch (-(fi8)) {
            case 42:
                p0 = (p0 << ((13) & 63))
            case 661:
                p0 = 2
                switch "hello" {
                case "foo bar":
                    v10 := GenS1{F0: (v9 % (((v9 | (v9 * (v9 * (v9 >> ((1) & 63)))))) | 1)), F1: false, F2: 628}
                    _ = v10
                default:
                    switch (fi8 & 4294967295) {
                    case 1:
                        var v11 uint = 442
                        _ = v11
                    default:
                        var v12 uint16 = 65535
                        _ = v12
                    }
                }
                fmt.Println(fi7, fi8, fi7)
                v9 -= 4294967295
                switch 1 {
                case 1:
                    v13 := []int{(fi7 ^ (fi8 - fi7)), (-(fi8)), fi7}
                    _ = v13
                case 32767:
                    p1 += p1
                case 0:
                    var v14 uint64 = v9
                    _ = v14
                default:
                    switch "foo bar" {
                    case "foo bar":
                        var v15 bool = true
                        _ = v15
                    case "\n\t":
                        v9 += 42
                    default:
                        v9 += v9
                    }
                }
            default:
                var v16 uint = 731
                _ = v16
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if (42 != 4294967295) {
        v17 := []int{2147483647, -128, 1, 0}
        for fi18 := 0; fi18 < 7; fi18++ {
            for fi19 := 0; fi19 < 4; fi19++ {
                if true {
                    v20 := GenS1{F0: 0, F1: true, F2: 0}
                    _ = v20
                } else {
                    v21 := GenS2{F0: "hello", F1: false, F2: p0}
                    _ = len(v17)
                    _ = v21
                }
                if len(v17) > 0 {
                    v17[0] = (fi19 ^ 127)
                }
                _ = fi19
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 int64 = -9223372036854775808
                var v24 float32 = float32(float64(0.1))
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 string = (("a" + "foo bar") + "foo bar")
            _ = fi18
            _ = v25
        }
        _ = v17
    } else {
        p0 = p0
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := []string{}
    var v2 int16 = 32767
    var v3 uint8 = 27
    v4 := []string{("foo bar" + ""), ("\n\t" + "hello"), "\n\t", (("a" + "a") + "foo bar")}
    var v5 int8 = -128
    v6 := make(map[string]string)
    v1 = append(v1, "foo bar")
    v26 := v6[""]
    v6[""] = v26
    var v27 uint = 42
    switch -9223372036854775808 {
    case 828:
        for fi28 := 0; fi28 < 1; fi28++ {
            v6[(v26 + "a")] = (("\n\t" + ("\n\t" + v26)) + (("hello" + "\n\t") + "foo bar"))
            v29 := GenS3{F0: v5, F1: 899, F2: v27}
            _ = fi28
            _ = v29
        }
    case 128:
        if (!false) {
            switch 32767 {
            case -2147483648:
                var v30 uint16 = 1
                _ = v30
            default:
                for fi31 := 0; fi31 < 19; fi31++ {
                    v32 := GenS2{F0: v26, F1: (2 >= 319), F2: 4294967295}
                    if false {
                        var v33 uint64 = 42
                        var v34 int32 = 0
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
            }
            v35 := GenS1{F0: 42, F1: true, F2: int64(v2)}
            _ = v35
        }
    case 127:
        v2 = v2
    default:
        if (v3 < v3) {
            if true {
                var v36 uint8 = v3
                _ = v36
            } else {
                var v37 int8 = (-(v5))
                fmt.Println(v5)
                _ = v37
            }
            v38 := GenS4{F0: v27, F1: v27}
            switch "foo bar" {
            case "a":
                v2 = (-(v2))
            case "\n\t":
                v39 := GenS5{F0: 1, F1: 65535, F2: 42}
                _ = v39
            case "hello":
                v27 -= v27
            default:
                var v40 uint16 = 0
                _ = v40
            }
            _ = v38
        } else {
            v4 = append(v4, (v26 + (v26 + v26)))
            v41 := GenS5{F0: 0, F1: 2147483647, F2: 532}
            _ = v41
        }
    }
    var v42 uint64 = 4294967295
    var v43 uint = (v27 << ((uint(v3)) & 63))
    var v44 int = 0
    v27 += (v43 - uint(v3))
    _ = len(v4)
    var v45 float64 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
