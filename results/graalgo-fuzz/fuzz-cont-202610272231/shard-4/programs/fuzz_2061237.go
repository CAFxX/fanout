package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
func genFn0(p0 int32, p1 bool) int8 {
    for fi7 := 0; fi7 < 11; fi7++ {
        p1 = (1 > 42)
        _ = fi7
    }
    fmt.Println(p0, p0)
    if (true || ((!p1) || true)) {
        if ((("hello" + (("a" + "\n\t") + ("foo bar" + "hello"))) + "a") == ("\n\t" + ("\n\t" + "\n\t"))) {
            var v8 uint16 = 0
            _ = v8
        }
        p1 = false
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            fmt.Println(p0)
            switch "hello" {
            case "\n\t":
                if p1 {
                    v10 := []int{-1}
                    if (!true) {
                        var v11 int16 = 0
                        _ = v11
                    }
                    _ = v10
                }
            case "a":
                var v12 uint = 0
                _ = v12
            case "":
                for fi13 := 0; fi13 < 10; fi13++ {
                    p0 = (p0 << ((1) & 63))
                    p1 = ((0.1 - 3.14159) <= (1.5 + -2.25))
                    v14 := []int{(^(fi9))}
                    _ = fi13
                    _ = v14
                }
            default:
                if (float32(1e10) > (0.0 / (3.14159 / ((1.5 / 1.5) / ((((1.5 * (0.1 / 0.1)) - -2.25) / ((3.14159 + 0.0) / 3.14159)) - 1.5))))) {
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v16 := []int{}
                        if len(v16) > 0 {
                            v16[0] = -725
                        }
                        _ = len(v16)
                        _ = fi15
                        _ = v16
                    }
                    p1 = (float64(-2.25) != 0.0)
                }
            }
            switch (fi9 / (((fi9 - fi9)) | 1)) {
            case -830:
                v17 := GenS1{F0: -273, F1: "hello"}
                _ = v17
            case -315:
                var v18 uint16 = 640
                _ = v18
            case -1:
                if ((128 != (fi9 + fi9)) && ((42 >= 2) || (!(!false)))) {
                    var v19 float32 = 1e10
                    v20 := GenS2{F0: "hello"}
                    _ = v19
                    _ = v20
                } else {
                    var v21 uint = 251
                    _ = v21
                }
            }
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    return -128
}

func main() {
    var v1 float32 = (float32((3.14159 * 3.14159)) - (0.1 / (((3.14159 + 0.0) / 1.0) / -2.25)))
    var v2 int64 = -56
    v3 := []int{-994, 4294967295}
    var v4 uint64 = 2
    var v5 int16 = -32768
    v6 := make(map[string]bool)
    if (false && (true && true)) {
        for fi22 := 0; fi22 < 10; fi22++ {
            v1 += (v1 - v1)
            _ = fi22
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 bool = (!true)
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := v6["\n\t"]
            v26 = v24
            fmt.Println(v1)
            var v27 float64 = 0.0
            v28 := GenS3{F0: (fi23 << ((926) & 63))}
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        switch "\n\t" {
        case "hello":
            fmt.Println(v1)
        case "\n\t":
            v29 := v6[(("\n\t" + "a") + "\n\t")]
            _ = v29
        case "a":
            var v30 int64 = v2
            _ = v30
        default:
            var v31 float64 = -2.25
            _ = v31
        }
        _ = fi23
        _ = v24
    }
    v32 := v6["foo bar"]
    v5 += v5
    for fi33 := 0; fi33 < 9; fi33++ {
        v32 = false
        v4 = (v4 - v4)
        _ = len(v6)
        _ = fi33
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 uint = 0
        if (false || (!(3.14159 >= -2.25))) {
            v5 = v5
        } else {
            if v32 {
                if v32 {
                    if false {
                        fmt.Println(v35)
                        v36 := GenS1{F0: (v2 & v2), F1: "\n\t"}
                        _ = v36
                    } else {
                        v35 = 931
                    }
                    v35 = (^(v35))
                } else {
                    var v37 uint16 = 2
                    v38 := GenS1{F0: (v2 << ((uint(v37)) & 63)), F1: ("\n\t" + "a")}
                    _ = v37
                    _ = v38
                }
                v4 = v4
            }
            fmt.Println(fi34, v35)
            switch fi34 {
            case 2147483647:
                v35 += v35
            case 633:
                fmt.Println(v32)
            default:
                var v39 uint16 = 2
                _ = v39
            }
        }
        _ = fi34
        _ = v35
    }
    v40 := GenS4{F0: v32}
    switch -9223372036854775808 {
    case -9223372036854775808:
        v4 += v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
