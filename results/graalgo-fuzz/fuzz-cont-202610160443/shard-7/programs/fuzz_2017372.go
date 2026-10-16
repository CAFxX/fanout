package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int; F2 int64 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint32) int16 {
    switch 128 {
    case 749:
        v7 := GenS1{F0: ("\n\t" + "a")}
        _ = v7
    case 4294967295:
        if (false || (!(-155 == 2147483647))) {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 float32 = 1e10
                _ = fi8
                _ = v9
            }
            if false {
                v10 := GenS1{F0: "a"}
                _ = v10
            } else {
                p0 -= (^(p0))
                p0 = p0
            }
        }
    case 32767:
        var v11 uint64 = 4294967295
        _ = v11
    default:
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 uint32 = (^(p0))
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    var v15 bool = false
    fmt.Println(v15, v15)
    v16 := []int{-1, 42}
    var v17 float64 = -2.25
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 1000
}

func main() {
    v1 := []string{"\n\t", "hello", "a"}
    var v2 uint32 = 558
    v3 := make(map[string]int)
    var v4 int8 = 100
    var v5 uint16 = 65535
    var v6 bool = false
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int32 = 2147483647
        for fi20 := 0; fi20 < 16; fi20++ {
            v1 = append(v1, ("foo bar" + (("a" + ("hello" + "")) + "")))
            var v21 float64 = -2.25
            _ = len(v1)
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            if (!v6) {
                v19 -= (v19 + v19)
                if (0.1 != (0.0 * 1e10)) {
                    var v23 int32 = (^(v19))
                    var v24 float64 = -2.25
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(v4)
                    var v25 uint8 = 200
                    fmt.Println(v4, v2)
                    _ = v25
                }
            } else {
                if v6 {
                    var v26 int = (-(fi22))
                    if (!(v6 || (!false))) {
                        fmt.Println(v5, v2)
                    } else {
                        var v27 int = fi18
                        var v28 int64 = 9223372036854775807
                        _ = v27
                        _ = v28
                    }
                    v29 := GenS1{F0: "a"}
                    _ = v26
                    _ = v29
                } else {
                    switch fi18 {
                    case -230:
                        v6 = v6
                    case -45:
                        v30 := GenS2{F0: v2, F1: 4294967295, F2: -456}
                        _ = v30
                    case 1:
                        fmt.Println(v4, v4, v5)
                    }
                }
                var v31 int8 = -34
                _ = v31
            }
            v6 = (!v6)
            v32 := GenS3{F0: 65535, F1: (fi22 * -1)}
            _ = fi22
            _ = v32
        }
        _ = fi18
        _ = v19
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        v5 = v5
        fmt.Println(v6, v4)
        v3[(("\n\t" + "\n\t") + "foo bar")] = fi33
        _ = fi33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        switch (fi34 - (^(fi34))) {
        case -9223372036854775808:
            fmt.Println(v6, fi34, v2)
        case 265:
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 string = ""
                _ = fi35
                _ = v36
            }
        case 554:
            v2 = (v2 % ((v2) | 1))
        }
        var v37 int32 = 0
        _ = fi34
        _ = v37
    }
    var v38 uint8 = 13
    if false {
        fmt.Println(v4)
        v6 = (!v6)
        v4 = (v4 >> ((348) & 63))
    } else {
        if v6 {
            if (v6 && true) {
                v39 := GenS4{F0: (^(v4))}
                v4 += 114
                v2 = uint32(v38)
                _ = v39
            }
            var v40 int8 = v4
            var v41 int16 = -799
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            fmt.Println(v6, v6, fi42)
            var v43 uint8 = v38
            v5 += (v5 + (v5 + 0))
            _ = fi42
            _ = v43
        }
        switch ("" + "hello") {
        case "\n\t":
            v1 = append(v1, "\n\t")
        case "a":
            v44 := GenS5{F0: 42}
            _ = v44
        case "foo bar":
            v45 := GenS6{F0: 0, F1: (v4 & (v4 * v4))}
            _ = v45
        }
    }
    v2 = 1
    fmt.Println(v5, v5, v4)
    switch "\n\t" {
    case "a":
        v38 = v38
    case "":
        var v46 uint32 = (v2 & 31)
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
