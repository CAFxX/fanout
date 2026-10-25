package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int64 }
type GenS2 struct { F0 string; F1 string; F2 int32 }
func genFn0(p0 uint) int8 {
    switch int(p0) {
    case 42:
        p0 = p0
    }
    p0 = p0
    switch -219 {
    case -9223372036854775808:
        p0 -= p0
    case 32768:
        for fi7 := 0; fi7 < 12; fi7++ {
            switch "hello" {
            case "a":
                var v8 int32 = 570
                _ = v8
            case "":
                if (-1 >= -42) {
                    var v9 uint16 = 42
                    _ = v9
                } else {
                    v10 := []int{fi7, 9223372036854775807, fi7}
                    var v11 float32 = (((0.0 / (-2.25 * 1.5)) / float32((float64(-2.25) / 3.14159))) * float32((-2.25 - 0.1)))
                    _ = v10
                    _ = v11
                }
            case "\n\t":
                for fi12 := 0; fi12 < 15; fi12++ {
                    var v13 int = (fi12 >> ((p0) & 63))
                    if ((true || true) && (42 > 42)) {
                        v14 := []int{9223372036854775807, (fi7 % (((-(fi12))) | 1)), fi12, (fi12 + (fi12 % ((-975) | 1)))}
                        v13 -= (fi12 | (fi12 * (fi7 + (fi12 + (^(fi12))))))
                        _ = v14
                    } else {
                        p0 -= (p0 >> ((p0) & 63))
                    }
                    var v15 float32 = 0.0
                    _ = fi12
                    _ = v13
                    _ = v15
                }
            }
            _ = fi7
        }
    }
    _ = p0
    return 100
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 27
    v3 := make(map[string]uint32)
    v4 := []string{"\n\t", "", "\n\t", ("a" + ("foo bar" + "a"))}
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    var v16 uint = 65535
    if ((true || false) || false) {
        v3[""] = 613
        v4 = append(v4, "\n\t")
        var v17 int64 = 42
        _ = v17
    }
    if ((v16 / ((4294967295) | 1)) == v16) {
        if (1000 <= 32767) {
            v18 := GenS1{F0: 9223372036854775807, F1: "a", F2: -1}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v18
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v16 = v16
    }
    var v19 int64 = -9223372036854775808
    for fi20 := 0; fi20 < 4; fi20++ {
        if (!true) {
            if (0 <= 0) {
                fmt.Println(fi20, v19)
                v21 := GenS2{F0: "foo bar", F1: "a", F2: 42}
                _ = v21
            }
        }
        _ = fi20
    }
    v2 = v2
    v2 = 142
    v2 -= v2
    if false {
        for fi22 := 0; fi22 < 16; fi22++ {
            if (false || false) {
                delete(v5, fi22)
                var v23 uint8 = (^(v2))
                _ = v23
            } else {
                for fi24 := 0; fi24 < 15; fi24++ {
                    if (0 >= 843) {
                        v16 = (^(v16))
                        v6[2] = (fi24 + fi24)
                    } else {
                        if true {
                            v25 := GenS1{F0: (fi24 ^ (fi22 % ((int(v19)) | 1))), F1: ("hello" + "hello"), F2: -1}
                            fmt.Println(fi22, fi22)
                            _ = v25
                        } else {
                            v3["a"] = 65535
                            fmt.Println(v2, v16, fi22)
                        }
                        fmt.Println(fi22)
                        v26 := GenS1{F0: (fi22 & fi24), F1: ("foo bar" + "a"), F2: (v19 << ((v16) & 63))}
                        _ = v26
                    }
                    _ = fi24
                }
            }
            fmt.Println(v16, fi22)
            _ = fi22
        }
        v27 := GenS1{F0: int(v16), F1: (("" + ("hello" + "foo bar")) + ("hello" + "\n\t")), F2: (v19 | (-(v19)))}
        v16 -= v16
        _ = v27
    } else {
        fmt.Println(v2, v16)
    }
    v16 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
