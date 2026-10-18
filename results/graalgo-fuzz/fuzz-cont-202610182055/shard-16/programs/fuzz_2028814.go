package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int }
type GenS5 struct { F0 int; F1 uint32 }
func genFn0(p0 uint16, p1 uint32, p2 int32) string {
    v4 := []int{}
    if true {
        p1 = p1
        var v5 uint16 = (^(p0))
        switch "a" {
        case "foo bar":
            v6 := GenS1{F0: -287, F1: 100}
            _ = v6
        }
        _ = v5
    }
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}
func genFn1(p0 uint8, p1 uint) string {
    p1 -= p1
    fmt.Println(p0)
    if (42 < 1000) {
        p0 -= 127
        switch 42 {
        case 32767:
            p0 -= 114
        case 0:
            if true {
                p0 = (p0 - p0)
                p0 = p0
                for fi7 := 0; fi7 < 20; fi7++ {
                    v8 := []int{32767, (fi7 >> (((p1 / ((p1) | 1))) & 63))}
                    _ = fi7
                    _ = v8
                }
            } else {
                p1 = (p1 - p1)
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := []int{127, (fi9 ^ (fi9 / ((fi9) | 1))), fi9}
                    p1 = 65535
                    fmt.Println(p0)
                    _ = fi9
                    _ = v10
                }
                v11 := GenS2{F0: (false && false)}
                _ = v11
            }
        case -1:
            v12 := []int{}
            _ = v12
        default:
            if true {
                var v13 int64 = 0
                v13 = v13
                _ = v13
            }
        }
    }
    _ = p0
    _ = p1
    return ("hello" + (("\n\t" + ("a" + "a")) + ""))
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint16 = 0
    var v3 float32 = ((0.1 * (((0.1 * 1.5) - 0.0) / (0.0 + 1.5))) + -2.25)
    switch 32768 {
    case 127:
        var v14 uint16 = (v2 << ((0) & 63))
        _ = v14
    }
    v3 -= float32((0.1 - 1e10))
    v1 = 692
    v1 += 2
    if (0 == 49) {
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(v1)
            _ = fi15
        }
        if true {
            v16 := GenS1{F0: 1, F1: 100}
            _ = v16
        } else {
            v17 := []int{1, -642}
            fmt.Println(v1, v2)
            var v18 int8 = 0
            _ = v17
            _ = v18
        }
    } else {
        switch "hello" {
        case "hello":
            v1 += (^(v1))
        case "":
            if (255 < 1) {
                v19 := []int{9223372036854775807, 0}
                var v20 int8 = -1
                for fi21 := 0; fi21 < 5; fi21++ {
                    v20 += v20
                    _ = fi21
                }
                _ = v19
                _ = v20
            }
        }
        var v22 float64 = 0.0
        for fi23 := 0; fi23 < 2; fi23++ {
            switch (-(fi23)) {
            case 128:
                v24 := GenS3{F0: 127, F1: 194, F2: 42}
                _ = v24
            case 32768:
                for fi25 := 0; fi25 < 14; fi25++ {
                    var v26 float64 = v22
                    if true {
                        switch 32767 {
                        case 42:
                            v1 = 0
                        case 127:
                            var v27 string = "\n\t"
                            _ = v27
                        case -689:
                            v26 = (1.5 / 3.14159)
                        }
                        var v28 int64 = -9223372036854775808
                        v1 = v1
                        _ = v28
                    } else {
                        if (1 >= 1) {
                            var v29 int = fi25
                            if false {
                                var v30 string = "a"
                                _ = v30
                            }
                            if false {
                                v31 := GenS4{F0: v1, F1: (^(v2)), F2: fi25}
                                var v32 uint = 255
                                v33 := GenS5{F0: (^(v29)), F1: 467}
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                var v34 bool = (uint8(v2) <= 2)
                                var v35 uint32 = 42
                                _ = v34
                                _ = v35
                            }
                            _ = v29
                        }
                        fmt.Println(v26, fi25)
                        var v36 uint = 42
                        _ = v36
                    }
                    var v37 float64 = 0.0
                    v37 = ((3.14159 + (v37 * (1e10 / (0.0 + v37)))) / ((float64(float32(3.14159)) - v26) + ((0.0 * v22) - ((v26 / -2.25) - (v22 / (v37 + 3.14159))))))
                    _ = fi25
                    _ = v26
                    _ = v37
                }
            case 133:
                v2 = 255
            default:
                v38 := GenS5{F0: (fi23 >> ((18446744073709551615) & 63)), F1: 2147483648}
                _ = v38
            }
            _ = fi23
        }
        _ = v22
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        v40 := []int{-885, (-(fi39))}
        _ = fi39
        _ = v40
    }
    var v41 uint32 = 744
    v3 -= v3
    v42 := GenS2{F0: ("\n\t" != "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
}
