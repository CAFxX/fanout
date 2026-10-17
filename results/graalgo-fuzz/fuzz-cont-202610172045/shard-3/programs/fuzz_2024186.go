package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 string }
func genFn0(p0 int32, p1 uint16) string {
    if false {
        p0 = (^(p0))
        p1 += 458
        p0 = p0
    }
    p1 += p1
    var v7 uint8 = 186
    _ = p0
    _ = p1
    _ = v7
    return "\n\t"
}

func main() {
    v1 := []int{32767}
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 uint16 = 639
    var v5 float32 = 1e10
    var v6 int = 127
    if false {
        for fi8 := 0; fi8 < 12; fi8++ {
            v1 = append(v1, (fi8 ^ (v6 * 9223372036854775807)))
            v4 = 32767
            v9 := GenS1{F0: -1, F1: 250, F2: 0}
            _ = fi8
            _ = v9
        }
        v6 += -363
        fmt.Println(v5, v6)
    } else {
        fmt.Println(v6)
    }
    switch (v6 ^ (v6 | v6)) {
    case 2147483647:
        var v10 uint32 = 288
        _ = v10
    case 32767:
        fmt.Println(v5)
    case -2147483648:
        v1 = append(v1, v6)
    }
    v4 = (v4 / ((v4) | 1))
    v1 = append(v1, (v6 % ((32768) | 1)))
    v6 -= (v6 & v6)
    switch v6 {
    case 0:
        v11 := GenS2{F0: 1000, F1: (v6 + v6)}
        _ = v11
    case 312:
        var v12 int8 = 1
        _ = v12
    case 4294967295:
        v13 := GenS2{F0: -1, F1: (v6 % ((42) | 1))}
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        switch (v6 - 796) {
        case -1:
            _ = len(v1)
        case -712:
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS3{F0: (0 <= 4294967295), F1: 9223372036854775807}
                _ = fi15
                _ = v16
            }
        }
        if false {
            var v17 bool = (false || (("foo bar" + "foo bar") > (("hello" + "") + "hello")))
            fmt.Println(v5)
            v18 := GenS4{F0: 1}
            _ = v17
            _ = v18
        }
        _ = fi14
    }
    _ = len(v1)
    var v19 int16 = -1
    var v20 int64 = -747
    var v21 bool = true
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            v4 = (v4 & (v4 * 255))
            if (0 < 127) {
                if v21 {
                    switch (fi22 / ((fi23) | 1)) {
                    case 4294967295:
                        fmt.Println(fi22, v6, v6)
                    case 155:
                        v20 = v20
                        var v24 uint16 = v4
                        switch (fi23 / ((982) | 1)) {
                        case 32767:
                            delete(v2, "foo bar")
                        default:
                            var v25 uint8 = 127
                            _ = v25
                        }
                        if false {
                            var v26 int16 = (v19 + (v19 + -64))
                            v21 = v21
                            v1 = append(v1, fi23)
                            _ = v26
                        } else {
                            v27 := GenS5{F0: 1, F1: 42}
                            _ = v27
                        }
                        _ = v24
                    case 32767:
                        var v28 int8 = -22
                        _ = v28
                    default:
                        v29 := GenS4{F0: -147}
                        _ = v29
                    }
                    var v30 uint = 2
                    _ = v30
                } else {
                    _ = len(v1)
                }
                var v31 int64 = (v20 << ((65535) & 63))
                _ = v31
            }
            if v21 {
                switch ("" + ("a" + ("hello" + "foo bar"))) {
                case "hello":
                    v19 -= v19
                case "foo bar":
                    v32 := GenS6{F0: ("\n\t" + "a")}
                    _ = v32
                case "":
                    var v33 float64 = 0.0
                    _ = v33
                }
            } else {
                switch (fi22 - -719) {
                case -9223372036854775808:
                    var v34 bool = false
                    _ = v34
                case 32768:
                    v35 := GenS1{F0: -1, F1: 1, F2: 2}
                    _ = v35
                case 128:
                    v36 := v3[""]
                    _ = v36
                default:
                    v4 += 1
                }
            }
            _ = fi23
        }
        fmt.Println(v21)
        _ = fi22
    }
    var v37 float32 = float32(0.0)
    var v38 float32 = ((v5 + 1e10) * 0.0)
    fmt.Println(v37, v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
