package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 bool }
func genFn0(p0 uint, p1 uint16) string {
    for fi6 := 0; fi6 < 9; fi6++ {
        switch ((("" + "") + (("a" + "foo bar") + ("" + "foo bar"))) + "foo bar") {
        case "\n\t":
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := []int{32767, -1, (fi7 - fi7), (fi6 & 127)}
                _ = fi7
                _ = v8
            }
        case "a":
            v9 := GenS1{F0: 0}
            _ = v9
        }
        _ = fi6
    }
    if (!false) {
        p0 -= (p0 - p0)
        v10 := GenS1{F0: -2}
        if false {
            var v11 uint16 = p1
            _ = v11
        }
        _ = v10
    } else {
        p1 = (^(p1))
        v12 := []int{-1, 1, 302, -28}
        v13 := GenS2{F0: 127}
        _ = v12
        _ = v13
    }
    switch 42 {
    case 42:
        var v14 uint = uint(p1)
        _ = v14
    case -1:
        v15 := []int{720, 127, 0, 32768}
        _ = v15
    default:
        switch 4294967295 {
        case 128:
            switch "a" {
            case "a":
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 uint16 = (^(p1))
                    _ = fi16
                    _ = v17
                }
            case "hello":
                p0 -= 2
            }
        case 9223372036854775807:
            var v18 int16 = 0
            _ = v18
        }
    }
    _ = p0
    _ = p1
    return "a"
}
func genFn1(p0 int64) uint16 {
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 uint16 = 642
    var v4 bool = (true && (-1 < 332))
    v5 := make(map[string]int)
    var v19 int64 = -778
    v20 := []int{}
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS2{F0: 255}
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint = 2
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    if ((v4 || true) || v4) {
        var v25 float64 = (float64(0.0) * 1.5)
        _ = v25
    } else {
        v26 := GenS3{F0: 36, F1: 1000}
        v27 := GenS4{F0: 2147483648}
        switch 9223372036854775807 {
        case 2147483647:
            fmt.Println(v4, v3, v3)
        case 128:
            switch int(v19) {
            case -1:
                var v28 int64 = v19
                _ = v28
            }
        default:
            var v29 int64 = v19
            _ = v29
        }
        _ = v26
        _ = v27
    }
    if (v4 && v4) {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        if v4 {
            v30 := GenS2{F0: 14}
            _ = v30
        }
        v3 += (v3 % ((genFn1(677)) | 1))
    }
    if len(v20) > 0 {
        v20[0] = 42
    }
    v31 := GenS2{F0: 1}
    var v32 bool = (v3 <= (^(v3)))
    _ = len(v20)
    v32 = ("\n\t" == "hello")
    for fi33 := 0; fi33 < 6; fi33++ {
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 float64 = -2.25
            _ = fi34
            _ = v35
        }
        if true {
            v20 = append(v20, (fi33 << ((255) & 63)))
            v36 := GenS1{F0: 2147483647}
            var v37 int8 = -128
            _ = v36
            _ = v37
        } else {
            var v38 uint = 4294967295
            if (!(-9223372036854775808 == v19)) {
                v38 += (v38 << ((1) & 63))
            } else {
                var v39 int8 = -99
                for fi40 := 0; fi40 < 10; fi40++ {
                    if (!(255 >= (v38 & v38))) {
                        _ = len(v1)
                    } else {
                        var v41 float32 = ((3.14159 + 0.1) + 0.1)
                        _ = v41
                    }
                    fmt.Println(v19)
                    fmt.Println(fi33, v39)
                    v42 := GenS5{F0: 2, F1: true}
                    _ = fi40
                    _ = v42
                }
                _ = v39
            }
            _ = v38
        }
        v2["hello"] = fi33
        _ = fi33
    }
    var v43 int8 = 100
    switch 32768 {
    case 128:
        var v44 uint8 = 0
        _ = v44
    case -464:
        v2["hello"] = 32767
    default:
        var v45 bool = v32
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v5
}
