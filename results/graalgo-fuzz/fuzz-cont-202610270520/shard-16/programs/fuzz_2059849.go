package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint64; F2 int32 }
type GenS4 struct { F0 int8; F1 bool }
func genFn0(p0 string, p1 uint) int8 {
    switch p0 {
    case "\n\t":
        p0 = p0
    case "a":
        if false {
            p0 = ("" + ((p0 + p0) + (("hello" + "") + "hello")))
            switch 2147483647 {
            case 465:
                p0 = p0
            case 128:
                if true {
                    p1 -= 1
                    var v5 uint64 = 2
                    _ = v5
                }
            }
            for fi6 := 0; fi6 < 8; fi6++ {
                fmt.Println(p0, p0, p1)
                for fi7 := 0; fi7 < 1; fi7++ {
                    var v8 int16 = -32768
                    fmt.Println(fi7, fi6, p0)
                    v8 = v8
                    _ = fi7
                    _ = v8
                }
                v9 := GenS1{F0: p0}
                _ = fi6
                _ = v9
            }
        } else {
            fmt.Println(p0)
        }
    case "foo bar":
        for fi10 := 0; fi10 < 5; fi10++ {
            switch p0 {
            case "a":
                var v11 bool = (1 <= 233)
                _ = v11
            case "":
                if true {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        p1 += (p1 + (p1 % (((p1 << (((p1 - 402)) & 63))) | 1)))
                        v13 := []int{}
                        _ = fi12
                        _ = v13
                    }
                    var v14 uint8 = 42
                    _ = v14
                }
            case "foo bar":
                for fi15 := 0; fi15 < 5; fi15++ {
                    v16 := []int{-1, (fi15 * (fi15 ^ fi10))}
                    if true {
                        p0 = ("hello" + ((("a" + (p0 + "hello")) + p0) + ((((p0 + p0) + "a") + "a") + p0)))
                        var v17 int64 = 986
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
            default:
                var v18 uint8 = 128
                _ = v18
            }
            _ = fi10
        }
    default:
        switch 32767 {
        case 32767:
            for fi19 := 0; fi19 < 2; fi19++ {
                fmt.Println(fi19, p0, p0)
                v20 := GenS1{F0: (p0 + "")}
                _ = fi19
                _ = v20
            }
        case 42:
            p1 = p1
        case 4294967295:
            v21 := GenS1{F0: (p0 + "\n\t")}
            _ = v21
        default:
            if false {
                switch 127 {
                case 32767:
                    var v22 float64 = (((1.5 / ((((1.5 + 3.14159) + 0.1) + float64(0.1)) + -2.25)) * 1.5) + float64(((((((((float32(0.0) * 3.14159) / 1e10) * 0.1) - ((((0.1 * 1.5) + float32(1e10)) - 1e10) / 3.14159)) - float32((((3.14159 / 1.5) + 3.14159) / 3.14159))) * 3.14159) + 1e10) - 0.0)))
                    _ = v22
                case 4294967295:
                    v23 := []int{0, 1, 4294967295, 887}
                    _ = v23
                }
                var v24 string = (p0 + "hello")
                _ = v24
            }
        }
    }
    p0 = ""
    _ = p0
    _ = p1
    return 50
}
func genFn1(p0 uint32, p1 string, p2 int) int {
    v25 := []int{}
    var v26 float32 = 0.1
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    if true {
        p1 = p1
        if ((true && ((0 <= -32768) && (204 <= 0))) || true) {
            var v27 string = (p1 + (((p1 + (p1 + "hello")) + p1) + ((p1 + p1) + p1)))
            v28 := GenS2{F0: 32767}
            v26 = v26
            _ = v27
            _ = v28
        }
    }
    v29 := GenS3{F0: "hello", F1: 2, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    _ = v29
    return p2
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = -1
    v3 := make(map[uint32]bool)
    var v4 int64 = (v2 * -237)
    v30 := GenS2{F0: -237}
    fmt.Println(v1, v1)
    var v31 uint32 = 42
    var v32 uint16 = 468
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            v31 += v31
            v35 := GenS4{F0: 42, F1: (1 >= 107)}
            if ((v4 < v2) || (!true)) {
                v1 = 4294967295
            } else {
                var v36 uint8 = 42
                if false {
                    var v37 bool = (1.5 < 0.0)
                    var v38 int16 = -1
                    _ = v37
                    _ = v38
                } else {
                    if (false && (127 <= 0)) {
                        _ = len(v3)
                        v4 -= v4
                        switch "a" {
                        case "hello":
                            v39 := GenS3{F0: "\n\t", F1: (^(v1)), F2: 42}
                            _ = v39
                        case "":
                            v32 -= uint16(v31)
                        default:
                            if ((fi34 & (-(fi34))) <= genFn1((^(v31)), "a", genFn1((^(v31)), (("\n\t" + "a") + ""), int(v32)))) {
                                v40 := []int{(fi34 | genFn1((v31 + 605), "hello", fi33)), 2147483647, (fi34 / (((fi33 % ((genFn1(0, "foo bar", 0)) | 1))) | 1)), genFn1((v31 / ((65535) | 1)), "a", genFn1((v31 % ((v31) | 1)), "foo bar", genFn1(v31, "a", genFn1((^(v31)), (("hello" + "a") + "\n\t"), (fi34 | fi33)))))}
                                if len(v40) >= 2 {
                                    v40 = v40[1:len(v40)]
                                }
                                v36 = (v36 << ((65535) & 63))
                                _ = v40
                            }
                        }
                    }
                    var v41 int16 = 1000
                    _ = len(v3)
                    _ = v41
                }
                _ = v36
            }
            _ = fi34
            _ = v35
        }
        v42 := []int{}
        _ = fi33
        _ = v42
    }
    switch 233 {
    case 42:
        delete(v3, v31)
    default:
        var v43 float64 = float64(-2.25)
        _ = v43
    }
    v44 := []int{32767}
    v3[v31] = (false || false)
    v2 -= 168
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v44
}
