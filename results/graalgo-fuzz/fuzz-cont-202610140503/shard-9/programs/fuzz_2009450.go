package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint32, p2 int) bool {
    if ((!(0 == -1)) && (false && (2147483647 >= -1))) {
        if ((676 < 1) && false) {
            p2 -= p2
        } else {
            if (false || false) {
                v5 := GenS1{F0: 232}
                _ = v5
            } else {
                switch "a" {
                case "foo bar":
                    fmt.Println(p0, p2)
                case "\n\t":
                    for fi6 := 0; fi6 < 4; fi6++ {
                        fmt.Println(p0, p2, p1)
                        switch "hello" {
                        case "":
                            p0 = (p0 >> ((255) & 63))
                        case "a":
                            p0 += p0
                        case "foo bar":
                            var v7 int16 = 32767
                            _ = v7
                        }
                        _ = fi6
                    }
                case "hello":
                    if false {
                        var v8 int32 = -1
                        _ = v8
                    }
                default:
                    v9 := GenS2{F0: p0}
                    _ = v9
                }
                for fi10 := 0; fi10 < 2; fi10++ {
                    switch "\n\t" {
                    case "":
                        var v11 uint64 = 875
                        _ = v11
                    case "foo bar":
                        v12 := GenS2{F0: 1000}
                        _ = v12
                    default:
                        if (0 >= 4294967295) {
                            var v13 float32 = (1.5 * 0.0)
                            var v14 uint16 = 453
                            _ = v13
                            _ = v14
                        }
                    }
                    v15 := []int{(p2 / (((fi10 ^ p2)) | 1)), (-(fi10))}
                    if false {
                        switch fi10 {
                        case 127:
                            if (0 <= 42) {
                                var v16 uint32 = p1
                                if len(v15) > 0 {
                                    v15[0] = 0
                                }
                                _ = v16
                            }
                        }
                    }
                    _ = fi10
                    _ = v15
                }
            }
            v17 := GenS3{F0: 1}
            p0 += (p0 * -32768)
            _ = v17
        }
        v18 := GenS1{F0: 77}
        v19 := GenS4{F0: 63}
        _ = v18
        _ = v19
    } else {
        var v20 int32 = -1
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := GenS1{F0: 2}
            p2 = 204
            _ = fi22
            _ = v23
        }
        v24 := GenS2{F0: 32767}
        if (-1 != 94) {
            var v25 int16 = (p0 - (p0 ^ 42))
            _ = v25
        }
        _ = fi21
        _ = v24
    }
    v26 := []int{(p2 / ((p2) | 1)), 42}
    p0 = (-(p0))
    var v27 int = (p2 & (p2 ^ p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return false
}

func main() {
    v1 := []string{}
    var v2 float64 = 1e10
    var v3 float32 = 0.0
    v4 := make(map[int]uint32)
    fmt.Println(v3)
    var v28 uint = 470
    v1 = append(v1, "\n\t")
    var v29 int64 = 848
    if (65535 >= 32767) {
        v30 := v4[1]
        if ((4294967295 <= 65535) || true) {
            v30 -= 241
        } else {
            fmt.Println(v2, v2, v30)
        }
        var v31 string = ("\n\t" + "hello")
        _ = v30
        _ = v31
    } else {
        delete(v4, 9223372036854775807)
        v29 = 1
        fmt.Println(v2, v3, v29)
    }
    var v32 uint32 = 1
    var v33 uint8 = 220
    _ = len(v4)
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 uint64 = 2
        _ = fi34
        _ = v35
    }
    switch 0 {
    case 42:
        v1 = append(v1, ("\n\t" + (("a" + "a") + ((("a" + "foo bar") + "hello") + "foo bar"))))
    case 128:
        v36 := GenS2{F0: -32768}
        _ = v36
    default:
        v37 := GenS4{F0: 18446744073709551615}
        _ = v37
    }
    v32 += 69
    var v38 int = -1
    var v39 string = ("a" + ("" + ""))
    v39 = v39
    v4[v38] = (v32 >> ((v28) & 63))
    v1 = append(v1, "")
    if true {
        for fi40 := 0; fi40 < 16; fi40++ {
            if (v39 != v39) {
                fmt.Println(v39, v28)
                v32 -= (^(v32))
                fmt.Println(v28)
            } else {
                if false {
                    var v41 uint64 = 4294967295
                    _ = v41
                } else {
                    v28 += 2
                    v33 += v33
                }
                var v42 int = 42
                _ = v42
            }
            var v43 string = "\n\t"
            _ = fi40
            _ = v43
        }
        switch (v38 - v38) {
        case 127:
            var v44 int = (v38 / (((v38 * (v38 + (v38 + v38)))) | 1))
            _ = v44
        }
        var v45 int64 = v29
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v39
    _ = v4
}
