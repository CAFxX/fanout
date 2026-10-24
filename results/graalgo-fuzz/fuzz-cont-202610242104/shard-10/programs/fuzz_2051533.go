package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int8; F2 int32 }
type GenS4 struct { F0 uint }
func genFn0(p0 int32) int64 {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 457
    }
    switch "foo bar" {
    case "foo bar":
        p0 = 0
    case "\n\t":
        if (850 == -32768) {
            var v8 uint32 = 4294967295
            v8 -= v8
            p0 -= -879
            _ = v8
        }
    case "":
        var v9 float64 = 0.1
        _ = v9
    }
    var v10 int8 = 100
    _ = p0
    _ = v10
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"hello", "hello"}
    var v3 uint16 = 409
    var v4 uint8 = 128
    v5 := []string{"hello", "foo bar"}
    var v6 int8 = -106
    v1[""] = true
    var v11 uint16 = (v3 / ((670) | 1))
    switch -1 {
    case 127:
        v12 := GenS1{F0: 65535}
        _ = v12
    }
    switch "hello" {
    case "\n\t":
        if true {
            for fi13 := 0; fi13 < 7; fi13++ {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = fi13
            }
            if len(v2) > 0 {
                v2[0] = "a"
            }
        }
    case "":
        v14 := GenS2{F0: "hello", F1: 373}
        _ = v14
    case "hello":
        if true {
            v11 = 504
            var v15 bool = false
            _ = v15
        } else {
            fmt.Println(v6)
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 int = fi16
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 string = "hello"
                    v20 := GenS3{F0: v11, F1: 0, F2: 2147483647}
                    var v21 string = (v19 + v19)
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi16
                _ = v17
            }
        }
    default:
        var v22 int = 45
        _ = v22
    }
    var v23 uint = 280
    v6 = (v6 << ((2) & 63))
    if true {
        _ = len(v1)
        switch "foo bar" {
        case "a":
            for fi24 := 0; fi24 < 6; fi24++ {
                switch (fi24 & fi24) {
                case -267:
                    var v25 uint32 = uint32(v11)
                    _ = v25
                case 4294967295:
                    var v26 int16 = 32767
                    _ = v26
                default:
                    delete(v1, "hello")
                }
                _ = fi24
            }
        case "foo bar":
            _ = len(v1)
        }
    }
    if len(v5) > 0 {
        v5[0] = ("\n\t" + "")
    }
    v11 -= 1
    var v27 uint64 = 2
    for fi28 := 0; fi28 < 5; fi28++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v29 int16 = 42
        _ = fi28
        _ = v29
    }
    var v30 float64 = 1.5
    fmt.Println(v3, v23)
    var v31 uint64 = v27
    var v32 int = 2147483647
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 uint32 = 601
        var v35 uint8 = uint8(v27)
        if (v31 == v31) {
            _ = len(v2)
            switch "foo bar" {
            case "foo bar":
                v32 = (-(fi33))
            case "":
                var v36 int32 = 115
                _ = v36
            }
            if true {
                for fi37 := 0; fi37 < 5; fi37++ {
                    var v38 string = "a"
                    var v39 int8 = v6
                    v31 = (v27 * uint64(v4))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            } else {
                for fi40 := 0; fi40 < 13; fi40++ {
                    if (0 < 2) {
                        v32 += -9223372036854775808
                    } else {
                        var v41 uint32 = 2147483648
                        v42 := GenS1{F0: 2147483648}
                        _ = v41
                        _ = v42
                    }
                    fmt.Println(v35, v6, v32)
                    v43 := GenS4{F0: v23}
                    _ = fi40
                    _ = v43
                }
                var v44 uint16 = 52
                fmt.Println(v31, fi33)
                _ = v44
            }
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
