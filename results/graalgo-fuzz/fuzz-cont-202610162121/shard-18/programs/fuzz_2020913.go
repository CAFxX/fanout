package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
func genFn0(p0 string, p1 bool, p2 int16) string {
    p0 = "hello"
    var v5 float64 = (0.0 + 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "a"
}

func main() {
    var v1 float64 = 0.1
    v2 := []string{"", "foo bar", (("a" + "a") + ("" + ("a" + "a")))}
    v3 := make(map[int]uint32)
    v4 := []int{0}
    for fi6 := 0; fi6 < 15; fi6++ {
        if false {
            switch fi6 {
            case 9223372036854775807:
                fmt.Println(fi6, fi6)
            case 127:
                switch (fi6 | (fi6 << ((729) & 63))) {
                case -9223372036854775808:
                    v7 := GenS1{F0: (fi6 >> ((819) & 63)), F1: fi6, F2: ("foo bar" + "\n\t")}
                    _ = v7
                case 42:
                    v1 = v1
                default:
                    fmt.Println(fi6, fi6)
                }
            }
        } else {
            switch fi6 {
            case 42:
                v1 -= v1
            case 940:
                v1 = 1.5
            }
            v1 -= (v1 * v1)
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int16 = 790
                var v10 uint64 = 42
                if false {
                    v1 += v1
                    var v11 int16 = (v9 - v9)
                    _ = v11
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi6
    }
    var v12 bool = true
    v13 := GenS1{F0: 127, F1: 42, F2: ""}
    var v14 int64 = -9223372036854775808
    var v15 uint8 = 157
    var v16 uint8 = 2
    v17 := GenS1{F0: 615, F1: 9223372036854775807, F2: ""}
    delete(v3, 9223372036854775807)
    if (39 >= 0) {
        for fi18 := 0; fi18 < 5; fi18++ {
            fmt.Println(v15, v14)
            switch "hello" {
            case "":
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 int16 = -170
                    v21 := GenS1{F0: 32767, F1: 4294967295, F2: "foo bar"}
                    fmt.Println(v12, v14)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case "\n\t":
                v14 -= -1
            case "a":
                v3[fi18] = 99
            default:
                for fi22 := 0; fi22 < 5; fi22++ {
                    fmt.Println(fi22, v14, v15)
                    v23 := GenS1{F0: fi22, F1: fi18, F2: (("hello" + "\n\t") + "hello")}
                    v24 := GenS1{F0: (fi22 - fi22), F1: fi18, F2: "hello"}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            var v25 float32 = (1e10 - -2.25)
            _ = fi18
            _ = v25
        }
        v26 := v3[-2147483648]
        v27 := GenS1{F0: 0, F1: 144, F2: ("" + "\n\t")}
        _ = v26
        _ = v27
    }
    v16 = v16
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 int16 = 31
        _ = fi28
        _ = v29
    }
    var v30 float32 = (float32(v1) / -2.25)
    switch "foo bar" {
    case "\n\t":
        delete(v3, 1)
    case "hello":
        if v12 {
            v31 := GenS1{F0: -41, F1: 32767, F2: ""}
            _ = len(v3)
            _ = len(v3)
            _ = v31
        } else {
            v1 -= ((float64(v30) * (v1 / (0.0 * v1))) + 0.0)
            var v32 int32 = -681
            if (1e10 <= (v1 * v1)) {
                v4 = append(v4, -1)
                switch "foo bar" {
                case "":
                    var v33 uint = 567
                    _ = v33
                default:
                    v14 -= v14
                }
                if v12 {
                    delete(v3, 345)
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v3[-402] = 65535
                        _ = fi34
                    }
                    fmt.Println(v12)
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v36 := v3[(fi35 % ((-843) | 1))]
                        v37 := GenS1{F0: (fi35 / (((fi35 % ((fi35) | 1))) | 1)), F1: (fi35 * fi35), F2: ""}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                }
            }
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
