package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 bool, p1 int8) uint32 {
    var v6 uint64 = 42
    v7 := GenS1{F0: 1, F1: -9223372036854775808}
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 = true
        if p0 {
            if (0 <= -32768) {
                var v9 int32 = 0
                v10 := GenS1{F0: (v9 << ((39) & 63)), F1: -9223372036854775808}
                var v11 float64 = (-2.25 * 3.14159)
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 int8 = 1
            v13 := []int{(fi8 & (-(fi8)))}
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    switch -9223372036854775808 {
    case 128:
        p0 = (!p0)
    default:
        if p0 {
            p1 += p1
            switch "foo bar" {
            case "hello":
                for fi14 := 0; fi14 < 19; fi14++ {
                    v15 := []int{fi14, 32767, fi14, -1}
                    fmt.Println(v6)
                    if p0 {
                        for fi16 := 0; fi16 < 20; fi16++ {
                            var v17 string = ("" + "foo bar")
                            _ = fi16
                            _ = v17
                        }
                        var v18 int16 = -1
                        var v19 int16 = (v18 ^ 301)
                        _ = v18
                        _ = v19
                    }
                    _ = fi14
                    _ = v15
                }
            case "a":
                var v20 int32 = 551
                _ = v20
            case "\n\t":
                if p0 {
                    v21 := GenS1{F0: 42, F1: 423}
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v23 := []int{fi22, fi22}
                        for fi24 := 0; fi24 < 19; fi24++ {
                            var v25 int = fi22
                            _ = fi24
                            _ = v25
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                }
            default:
                var v26 int32 = -2147483648
                _ = v26
            }
        } else {
            var v27 uint16 = 961
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 bool = false
                v30 := GenS1{F0: 2147483647, F1: -9223372036854775808}
                var v31 uint = 255
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            if p0 {
                p1 += 42
            }
            _ = v27
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 737
}

func main() {
    var v1 uint8 = 127
    var v2 uint = 296
    v3 := []string{}
    v4 := []string{"foo bar"}
    var v5 int8 = 126
    var v32 bool = false
    fmt.Println(v32, v5)
    switch -2147483648 {
    case 128:
        v33 := GenS2{F0: 42}
        _ = v33
    case 9223372036854775807:
        v32 = v32
    case 32767:
        v1 = (v1 + 42)
    }
    if (0 <= 65535) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v34 int8 = v5
        _ = v34
    }
    fmt.Println(v1, v32)
    v35 := GenS3{F0: 684, F1: -354}
    for fi36 := 0; fi36 < 13; fi36++ {
        fmt.Println(v2, fi36)
        var v37 bool = ((-1 <= 32767) || (true && ((((0.1 * (3.14159 + -2.25)) / 1e10) / 1e10) <= 1e10)))
        if v32 {
            switch (fi36 % ((fi36) | 1)) {
            case 128:
                switch (fi36 / ((fi36) | 1)) {
                case 879:
                    fmt.Println(v5, v1)
                case 0:
                    v3 = append(v3, "")
                }
            case 854:
                switch "a" {
                case "foo bar":
                    var v38 float32 = 1.5
                    _ = v38
                case "a":
                    v4 = append(v4, "")
                }
            case 833:
                var v39 float32 = (1e10 - 1e10)
                _ = v39
            default:
                _ = len(v3)
            }
        } else {
            var v40 uint = 255
            fmt.Println(v2, v5)
            for fi41 := 0; fi41 < 3; fi41++ {
                switch "\n\t" {
                case "":
                    var v42 int8 = (-(v5))
                    _ = v42
                default:
                    v40 = v2
                }
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                var v43 uint64 = 42
                _ = fi41
                _ = v43
            }
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    v1 = v1
    v1 += 155
    v3 = append(v3, "foo bar")
    var v44 int32 = -214
    v2 = (v2 << ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v44
    _ = v5
}
