package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int8; F1 int16; F2 uint }
type GenS4 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0(p0 uint64, p1 string, p2 int16) uint64 {
    v6 := GenS1{F0: 84, F1: 2}
    switch p1 {
    case "hello":
        p1 = p1
    case "":
        v7 := []int{32767}
        _ = v7
    case "a":
        p0 -= (p0 - p0)
    }
    if (!(!(!true))) {
        if (!(32 < -1)) {
            p2 -= (p2 % (((p2 | -340)) | 1))
            if true {
                var v8 uint32 = uint32(p0)
                _ = v8
            } else {
                if (true && (1.5 >= -2.25)) {
                    p2 = (p2 >> ((42) & 63))
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 uint64 = p0
                        var v11 int16 = (p2 & (-(p2)))
                        v12 := []int{(fi9 & (fi9 ^ 273)), (^(fi9)), 839, (-(fi9))}
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                } else {
                    v13 := GenS2{F0: p2, F1: 1}
                    p1 = (("foo bar" + "hello") + "a")
                    switch 587 {
                    case -9223372036854775808:
                        p0 += (p0 / (((p0 & p0)) | 1))
                    case 1:
                        switch 2147483647 {
                        case 32767:
                            fmt.Println(p1)
                        default:
                            p2 = (p2 - (p2 - (^(p2))))
                        }
                    }
                    _ = v13
                }
                v14 := []int{4294967295}
                if (65535 == 42) {
                    p2 = p2
                } else {
                    var v15 uint8 = 0
                    _ = v15
                }
                _ = v14
            }
        } else {
            p2 = (p2 * p2)
            fmt.Println(p0, p2, p0)
        }
        v16 := []int{1, 32767, 287}
        for fi17 := 0; fi17 < 10; fi17++ {
            _ = len(v16)
            switch p1 {
            case "a":
                var v18 string = p1
                _ = v18
            case "hello":
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 uint16 = 32767
                    v16 = append(v16, 42)
                    if true {
                        p0 += p0
                    } else {
                        var v21 bool = (fi17 == 394)
                        if v21 {
                            fmt.Println(p1, fi17)
                            fmt.Println(fi19)
                            p2 = (p2 << ((207) & 63))
                            var v22 int64 = -1
                            _ = v22
                        } else {
                            var v23 int16 = 1000
                            _ = v23
                        }
                        switch "a" {
                        case "":
                            p1 = p1
                        case "\n\t":
                            v24 := GenS2{F0: p2, F1: 9223372036854775807}
                            _ = v24
                        default:
                            var v25 string = (("foo bar" + ("hello" + p1)) + p1)
                            fmt.Println(v21, v20, v25)
                            var v26 int64 = 42
                            _ = v25
                            _ = v26
                        }
                        var v27 int8 = -128
                        p0 -= p0
                        _ = v21
                        _ = v27
                    }
                    _ = fi19
                    _ = v20
                }
            case "\n\t":
                fmt.Println(p0)
            default:
                var v28 float32 = 1.5
                _ = v28
            }
            p0 += 251
            _ = fi17
        }
        _ = v16
    }
    v29 := GenS3{F0: 100, F1: (p2 - -337), F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v6
    return 192
}

func main() {
    v1 := []string{"", ("" + ""), "hello"}
    var v2 int16 = 1000
    var v3 int16 = v2
    var v4 uint = 408
    var v5 int8 = -128
    fmt.Println(v3, v4)
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS4{F0: genFn0(18446744073709551615, "foo bar", v3), F1: (^(v2)), F2: 1}
        var v32 int32 = -1
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int8 = (v5 >> ((v4) & 63))
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println(v4, v3)
    var v34 uint8 = 0
    if true {
        var v35 string = ("" + (("foo bar" + ("" + "")) + "a"))
        var v36 int16 = v3
        _ = len(v1)
        _ = v35
        _ = v36
    }
    var v37 uint = 677
    v3 = (v2 ^ v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v5
}
