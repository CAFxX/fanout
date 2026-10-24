package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 string; F1 int64; F2 uint32 }
func genFn0(p0 int16, p1 uint64, p2 uint8) int64 {
    v6 := GenS1{F0: "", F1: 460}
    var v7 uint8 = (^(p2))
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -388
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 float32 = 1e10
    v4 := []uint32{2, 0, 2}
    var v5 int16 = 0
    var v8 bool = false
    for fi9 := 0; fi9 < 7; fi9++ {
        v8 = (((-(fi9)) == (fi9 + fi9)) && (((!(2 != 128)) || v8) && v8))
        _ = fi9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 bool = (float64(0.1) == 1e10)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi11
            _ = v12
        }
        if v8 {
            for fi13 := 0; fi13 < 17; fi13++ {
                v3 -= 0.1
                _ = fi13
            }
            if (-2.25 != 1.5) {
                var v14 float32 = (float32(float64(((v3 / ((float32((1e10 / -2.25)) - (0.1 - (1e10 + 3.14159))) * v3)) + 0.0))) / 1.5)
                _ = v14
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int64 = 801
                var v17 float32 = 1e10
                v3 = float32(0.0)
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        if false {
            v5 += 32767
        } else {
            fmt.Println(fi10, v5)
        }
        _ = fi10
    }
    _ = len(v1)
    fmt.Println(v8, v8, v8)
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v8 = true
    if (v8 || (v8 || v8)) {
        v4 = append(v4, 2)
    } else {
        for fi18 := 0; fi18 < 4; fi18++ {
            v8 = false
            fmt.Println(fi18, v8, fi18)
            fmt.Println(fi18)
            _ = fi18
        }
        v19 := GenS1{F0: "hello", F1: 42}
        v5 -= v5
        _ = v19
    }
    if true {
        v20 := GenS2{F0: (v5 + (-(v5))), F1: "", F2: ("" + "foo bar")}
        v5 -= v5
        for fi21 := 0; fi21 < 18; fi21++ {
            switch (fi21 >> ((856) & 63)) {
            case 973:
                v3 -= (v3 + -2.25)
            case -2147483648:
                if false {
                    fmt.Println(v3, v3)
                    for fi22 := 0; fi22 < 11; fi22++ {
                        delete(v2, 65535)
                        var v23 int = 128
                        var v24 uint8 = 219
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            default:
                switch ("" + ("foo bar" + ("a" + ((("foo bar" + "a") + (("a" + ("" + "foo bar")) + "foo bar")) + ("hello" + ((("a" + "") + "hello") + "")))))) {
                case "\n\t":
                    v25 := v2[65535]
                    _ = v25
                }
            }
            if ((9223372036854775807 != -311) || (373 <= 827)) {
                var v26 uint32 = 2
                v27 := GenS1{F0: ("foo bar" + ("foo bar" + "")), F1: 749}
                _ = v26
                _ = v27
            } else {
                v1[2147483648] = (!v8)
                _ = len(v4)
            }
            _ = len(v4)
            _ = fi21
        }
        _ = v20
    }
    var v28 int64 = 9223372036854775807
    var v29 int = 1
    v30 := GenS1{F0: "\n\t", F1: 532}
    v31 := GenS3{F0: ("a" + "\n\t"), F1: (v28 | (^(v28))), F2: 2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v8
}
