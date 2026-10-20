package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint8 }
func genFn0(p0 string, p1 uint8, p2 bool) uint {
    fmt.Println(p0, p1)
    fmt.Println(p2)
    if p2 {
        p1 -= p1
        p0 = (p0 + p0)
    } else {
        p1 -= 1
        for fi6 := 0; fi6 < 9; fi6++ {
            p1 = 19
            if p2 {
                v7 := []int{fi6, fi6, (fi6 | (fi6 % ((32767) | 1)))}
                fmt.Println(p1, fi6, fi6)
                if (!true) {
                    v7 = append(v7, (fi6 | 0))
                    var v8 uint16 = 65535
                    _ = v8
                } else {
                    p1 -= 2
                    if len(v7) > 0 {
                        v7[0] = fi6
                    }
                }
                _ = v7
            } else {
                var v9 int = fi6
                if (42 >= 0) {
                    fmt.Println(fi6)
                    var v10 uint32 = 622
                    v11 := []int{v9, (v9 - (fi6 + (fi6 << ((65535) & 63)))), fi6, v9}
                    _ = v10
                    _ = v11
                } else {
                    v9 += v9
                    fmt.Println(fi6, p0, p1)
                }
                _ = v9
            }
            p1 -= 128
            _ = fi6
        }
    }
    v12 := GenS1{F0: 842}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}
func genFn1(p0 bool, p1 int64) uint16 {
    v13 := GenS2{F0: "foo bar", F1: 1}
    if (!false) {
        if ("foo bar" > (("hello" + "\n\t") + "a")) {
            p1 -= (p1 | (p1 ^ (p1 >> ((65535) & 63))))
        }
        if (647 < 674) {
            v14 := []int{32767, 128}
            if p0 {
                v15 := GenS2{F0: "", F1: 42}
                p0 = (true && (42 != 18446744073709551615))
                if len(v14) > 0 {
                    v14[0] = 9223372036854775807
                }
                _ = v15
            } else {
                v16 := GenS2{F0: ("foo bar" + (("\n\t" + ("hello" + "\n\t")) + "a")), F1: 4294967295}
                _ = v16
            }
            var v17 uint64 = 42
            _ = v14
            _ = v17
        } else {
            switch 1 {
            case 127:
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := GenS1{F0: 54}
                    v20 := GenS1{F0: 2147483647}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            default:
                v21 := GenS3{F0: 4294967295, F1: 2, F2: 0}
                _ = v21
            }
            v22 := GenS4{F0: 441}
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = v13
    return 2
}

func main() {
    v1 := []uint32{354, 2147483647, 0, 17}
    var v2 int16 = 402
    var v3 int8 = 15
    var v4 uint16 = 32767
    v5 := []string{}
    var v23 uint = 123
    if false {
        fmt.Println(v4)
        for fi24 := 0; fi24 < 5; fi24++ {
            if (!((158 >= 128) && ((v3 & (^(v3))) < v3))) {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v23 -= (v23 % (((v23 % (((v23 ^ 255)) | 1))) | 1))
                    if (false && (!true)) {
                        v26 := GenS4{F0: (v23 - v23)}
                        var v27 int64 = -9223372036854775808
                        v2 += v2
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v2)
                    _ = fi25
                }
                var v28 uint16 = (v4 & (^(v4)))
                _ = v28
            }
            _ = fi24
        }
    } else {
        v4 = genFn1(false, -9223372036854775808)
        v3 += v3
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    _ = len(v5)
    v29 := GenS3{F0: v23, F1: 255, F2: (v2 + 1000)}
    var v30 uint = (v23 * (v23 * (v23 | (v23 & v23))))
    var v31 float32 = (3.14159 * float32((-2.25 / (((3.14159 * (3.14159 - float64(-2.25))) * -2.25) / ((3.14159 * float64(1e10)) * ((0.1 + 0.1) / 3.14159))))))
    if (0 < 42) {
        var v32 int = 42
        v33 := GenS5{F0: ("" + (("\n\t" + "hello") + "foo bar")), F1: 0}
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 int16 = 32767
            var v36 uint64 = 0
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v32
        _ = v33
    }
    var v37 uint32 = 1
    v3 = int8(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v5
}
