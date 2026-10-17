package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 string, p1 int8, p2 uint16) uint8 {
    if true {
        p1 = (p1 * 42)
        var v5 int = 115
        v6 := GenS1{F0: true}
        _ = v5
        _ = v6
    } else {
        if (338 == 2) {
            p1 += p1
        }
    }
    v7 := []int{0, 4294967295, 533}
    var v8 int16 = 1
    var v9 string = (((p0 + p0) + ("hello" + ("\n\t" + p0))) + ("a" + ((("foo bar" + "\n\t") + p0) + "foo bar")))
    v8 = 481
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 string) string {
    p0 = p0
    var v10 uint32 = 0
    _ = p0
    _ = v10
    return "\n\t"
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]uint32)
    var v3 uint8 = 190
    v4 := make(map[int]string)
    var v11 string = "\n\t"
    v3 = (^(v3))
    delete(v2, ((v11 + ("hello" + v11)) + "a"))
    if (!(!false)) {
        for fi12 := 0; fi12 < 12; fi12++ {
            v3 -= 19
            v3 += genFn0("hello", 1, 32767)
            fmt.Println(fi12, v11)
            _ = fi12
        }
        if (-1 > -1) {
            v13 := GenS2{F0: "", F1: 313}
            _ = v13
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            if false {
                v3 = v3
            } else {
                fmt.Println(v11)
            }
            fmt.Println(v11, v3)
            _ = fi14
        }
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            fmt.Println(fi15)
            v4[(fi15 << ((4294967295) & 63))] = ("hello" + v11)
            v1 = append(v1, 0)
            _ = fi15
        }
        var v16 uint = 881
        var v17 int64 = 1
        _ = v16
        _ = v17
    }
    if (2147483647 != 814) {
        switch 359 {
        case 32768:
            if (65535 != 42) {
                v18 := GenS3{F0: 722}
                _ = v18
            } else {
                var v19 string = ((v11 + "foo bar") + v11)
                _ = v19
            }
        case -9223372036854775808:
            v20 := GenS1{F0: true}
            _ = v20
        case -756:
            v11 = (v11 + ((("\n\t" + (v11 + v11)) + "foo bar") + "\n\t"))
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 string = "foo bar"
            fmt.Println(fi21)
            v3 += 42
            _ = fi21
            _ = v22
        }
    }
    var v23 bool = true
    v4[2147483647] = ((v11 + ((v11 + ("\n\t" + v11)) + v11)) + v11)
    for fi24 := 0; fi24 < 16; fi24++ {
        v25 := v2[(v11 + ((v11 + "") + ""))]
        v26 := GenS4{F0: v11, F1: true}
        v1 = append(v1, (^(v25)))
        _ = fi24
        _ = v25
        _ = v26
    }
    if (0 >= 32767) {
        v27 := GenS1{F0: v23}
        _ = v27
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
