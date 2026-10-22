package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0() int32 {
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 uint = 462
        _ = fi4
        _ = v5
    }
    if false {
        var v6 uint16 = 215
        _ = v6
    }
    switch 67 {
    case -9223372036854775808:
        v7 := []int{32768, 32767, 0}
        _ = v7
    }
    return -598
}
func genFn1(p0 int64) uint64 {
    p0 -= (p0 & (p0 - -888))
    v8 := []int{127}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]uint32)
    v3 := []string{}
    v1 -= genFn0()
    v2["hello"] = 504
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 1}
        if ((3.14159 < (0.0 + (0.1 * (float32(1e10) - float32((((3.14159 * (1e10 - 3.14159)) + ((3.14159 - 3.14159) * -2.25)) / 1e10)))))) && ("foo bar" >= "foo bar")) {
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 float64 = 1e10
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                _ = fi11
                _ = v12
            }
        } else {
            v1 = (v1 * v1)
            fmt.Println(fi9, fi9)
        }
        var v13 uint32 = 581
        _ = fi9
        _ = v10
        _ = v13
    }
    if true {
        if true {
            if ((!true) && (1 >= 42)) {
                var v14 uint = 4294967295
                _ = v14
            } else {
                v3 = append(v3, "hello")
                var v15 float64 = (0.1 / 0.1)
                _ = v15
            }
            var v16 float32 = 0.0
            var v17 float64 = (3.14159 * -2.25)
            _ = v16
            _ = v17
        } else {
            var v18 uint32 = 2
            _ = v18
        }
        var v19 int32 = 1
        _ = v19
    } else {
        switch 2147483647 {
        case 9223372036854775807:
            switch ("" + "\n\t") {
            case "\n\t":
                v1 = v1
            case "foo bar":
                v20 := GenS2{F0: "hello", F1: 1}
                _ = v20
            case "":
                var v21 uint64 = genFn1(42)
                _ = v21
            }
        default:
            fmt.Println(v1)
        }
        var v22 uint8 = 72
        _ = v22
    }
    v23 := GenS2{F0: "foo bar", F1: 128}
    switch -198 {
    case 4294967295:
        if (true || false) {
            v3 = append(v3, "foo bar")
            var v24 uint64 = genFn1(-1)
            _ = v24
        } else {
            v3 = append(v3, "hello")
            v25 := GenS1{F0: (^(v1))}
            _ = v25
        }
    case -9223372036854775808:
        switch "foo bar" {
        case "":
            for fi26 := 0; fi26 < 4; fi26++ {
                switch "\n\t" {
                case "\n\t":
                    v27 := GenS1{F0: (v1 - 0)}
                    _ = v27
                case "foo bar":
                    fmt.Println(v1, fi26)
                case "hello":
                    v1 -= -1
                }
                var v28 bool = (1 > 1)
                fmt.Println(fi26, fi26)
                _ = fi26
                _ = v28
            }
        case "\n\t":
            var v29 uint32 = 65535
            _ = v29
        case "foo bar":
            v1 -= (-(v1))
        }
    case 32768:
        v1 += 2147483647
    }
    v1 -= (v1 & 1)
    var v30 int64 = 42
    for fi31 := 0; fi31 < 14; fi31++ {
        v1 = genFn0()
        _ = fi31
    }
    if (true || (!false)) {
        if false {
            for fi32 := 0; fi32 < 2; fi32++ {
                v33 := GenS1{F0: v1}
                if false {
                    var v34 int32 = -1
                    delete(v2, "")
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    _ = v34
                } else {
                    v35 := GenS3{F0: 1000}
                    v1 = (v1 >> ((4294967295) & 63))
                    if (!(42 == 1000)) {
                        fmt.Println(v1, fi32)
                        v3 = append(v3, ("" + "foo bar"))
                        v36 := GenS4{F0: 199, F1: (fi32 & (fi32 / ((fi32) | 1)))}
                        _ = v36
                    } else {
                        v2["foo bar"] = 573
                    }
                    _ = v35
                }
                _ = fi32
                _ = v33
            }
        } else {
            if false {
                var v37 uint8 = 42
                _ = len(v3)
                var v38 uint = 255
                _ = v37
                _ = v38
            }
        }
        v39 := v2["\n\t"]
        _ = v39
    } else {
        v1 = -575
        if (968 >= 42) {
            v30 = v30
            v40 := GenS1{F0: (v1 + v1)}
            v1 = (^(v1))
            _ = v40
        } else {
            fmt.Println(v30, v30)
            var v41 int = 1
            _ = v41
        }
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
}
