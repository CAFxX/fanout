package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 string }
func genFn0(p0 uint8, p1 uint64) int64 {
    var v4 int32 = 1
    if (((p1 >> ((0) & 63)) == p1) && (false || ((p1 - 4294967295) != (p1 >> ((255) & 63))))) {
        p0 = p0
        v5 := GenS1{F0: 32768, F1: (v4 + v4), F2: "hello"}
        p0 = 5
        _ = v5
    }
    if (!false) {
        v6 := []int{349}
        _ = v6
    }
    v7 := []int{}
    p1 = p1
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 42
}
func genFn1(p0 int, p1 bool, p2 uint8) uint32 {
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 int = p0
            v11 := []int{-9223372036854775808, p0}
            fmt.Println(p2, fi9)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        p2 = (^(p2))
        _ = fi12
    }
    var v13 int8 = 30
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 72
}

func main() {
    v1 := []int{}
    var v2 int = 734
    var v3 bool = false
    v14 := GenS1{F0: 468, F1: 0, F2: "foo bar"}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v15 string = "foo bar"
    if v3 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch "foo bar" {
    case "a":
        v3 = (32767 <= 32767)
    }
    if v3 {
        var v16 float64 = 3.14159
        v15 = v15
        if v3 {
            v2 = 32767
            var v17 int = (v2 + v2)
            switch ((v15 + ("foo bar" + ((((((v15 + v15) + (v15 + v15)) + v15) + "a") + v15) + (v15 + "hello")))) + (v15 + v15)) {
            case "hello":
                var v18 int16 = -1
                _ = v18
            default:
                for fi19 := 0; fi19 < 14; fi19++ {
                    switch (v15 + v15) {
                    case "":
                        var v20 float64 = v16
                        _ = v20
                    case "hello":
                        switch "hello" {
                        case "\n\t":
                            for fi21 := 0; fi21 < 15; fi21++ {
                                v1 = append(v1, 42)
                                var v22 float32 = 1e10
                                fmt.Println(v22, fi21, v22)
                                _ = fi21
                                _ = v22
                            }
                        }
                    }
                    _ = fi19
                }
            }
            _ = v17
        } else {
            if true {
                if v3 {
                    if v3 {
                        _ = len(v1)
                        for fi23 := 0; fi23 < 9; fi23++ {
                            v15 = v15
                            switch v15 {
                            case "a":
                                fmt.Println(v15)
                            case "foo bar":
                                _ = len(v1)
                            case "\n\t":
                                fmt.Println(v3, v16, v2)
                            }
                            _ = fi23
                        }
                    }
                    var v24 uint8 = 42
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 18; fi25++ {
                v16 -= ((v16 * v16) - v16)
                v1 = append(v1, 32768)
                _ = fi25
            }
        }
        _ = v16
    }
    var v26 int32 = 2147483647
    var v27 int16 = -32768
    fmt.Println(v2)
    var v28 string = (v15 + ("foo bar" + v15))
    v27 = v27
    v29 := GenS1{F0: (v2 % ((-9223372036854775808) | 1)), F1: v26, F2: "foo bar"}
    var v30 uint16 = 65535
    switch ((v28 + v15) + v28) {
    case "":
        if (!(-1 <= -324)) {
            switch v28 {
            case "\n\t":
                v2 -= (v2 * (v2 & v2))
            case "hello":
                v2 = (v2 & v2)
            case "foo bar":
                if (((-9223372036854775808 >= -178) && v3) && (1 > v30)) {
                    fmt.Println(v30)
                } else {
                    var v31 int = v2
                    fmt.Println(v2, v27)
                    _ = v31
                }
            default:
                fmt.Println(v15, v28)
            }
            fmt.Println(v15)
            var v32 string = ("" + "\n\t")
            _ = v32
        } else {
            var v33 int8 = 46
            _ = v33
        }
    case "hello":
        v28 = "foo bar"
    case "a":
        v34 := GenS1{F0: (v2 * 1), F1: 42, F2: v15}
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
