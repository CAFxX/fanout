package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint64) uint {
    if (1 <= 0) {
        if false {
            var v5 int8 = 42
            p1 = p1
            switch (("hello" + "\n\t") + "") {
            case "foo bar":
                v6 := GenS1{F0: false}
                _ = v6
            case "hello":
                v7 := GenS1{F0: (950 < 679)}
                _ = v7
            case "":
                for fi8 := 0; fi8 < 15; fi8++ {
                    v9 := []int{}
                    v10 := GenS2{F0: -1, F1: 255, F2: -1}
                    var v11 int = fi8
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
            default:
                for fi12 := 0; fi12 < 18; fi12++ {
                    fmt.Println(p1, p0)
                    _ = fi12
                }
            }
            _ = v5
        } else {
            var v13 int16 = -1
            var v14 float32 = (0.0 + (float32(((1e10 / 1e10) + (1.5 + -2.25))) - 0.0))
            _ = v13
            _ = v14
        }
        p0 -= (p0 | 2147483647)
    }
    switch (("\n\t" + "a") + "a") {
    case "foo bar":
        var v15 int16 = -1
        _ = v15
    case "\n\t":
        p0 += (p0 * p0)
    default:
        v16 := []int{4294967295}
        _ = v16
    }
    v17 := []int{127, -2147483648, -445, 1}
    if true {
        if false {
            p1 += p1
            if len(v17) > 0 {
                v17[0] = 32767
            }
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            if true {
                p1 = (p1 & 588)
                var v19 bool = true
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                if false {
                    var v21 int16 = 554
                    _ = v21
                }
                _ = fi20
            }
            _ = fi18
        }
    }
    _ = p0
    _ = p1
    _ = v17
    return 99
}
func genFn1(p0 uint, p1 uint) int64 {
    fmt.Println(p0, p0)
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := []int{(fi22 ^ 128), 0, 32767}
        if false {
            if len(v23) > 0 {
                v23[0] = 128
            }
        }
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(fi24, fi24, fi24)
            p1 = p0
            _ = fi25
        }
        v26 := []int{(^(fi24))}
        _ = fi24
        _ = v26
    }
    var v27 uint32 = 65535
    v28 := []int{42, -517, 0}
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return -385
}

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 853
    var v3 int32 = -636
    var v4 uint64 = 827
    var v29 int32 = v3
    if (-128 == 0) {
        var v30 uint16 = v2
        _ = v30
    } else {
        v31 := []int{32768}
        _ = v31
    }
    v32 := GenS3{F0: -112}
    v4 += (v4 * 42)
    v29 = v29
    v33 := []int{-9223372036854775808}
    v3 -= v3
    v29 -= (v3 / ((v29) | 1))
    v1 -= (v1 / ((65535) | 1))
    v4 = v4
    v2 = (v2 << (((v1 * (v1 - (v1 | (v1 << ((v1) & 63)))))) & 63))
    var v34 string = "foo bar"
    for fi35 := 0; fi35 < 13; fi35++ {
        if (v34 > (((("hello" + v34) + v34) + v34) + v34)) {
            for fi36 := 0; fi36 < 15; fi36++ {
                v37 := GenS2{F0: 155, F1: 32767, F2: (v3 - -2147483648)}
                v29 += v3
                _ = fi36
                _ = v37
            }
            var v38 int = (fi35 % (((fi35 * (fi35 % ((fi35) | 1)))) | 1))
            if len(v33) > 0 {
                v33[0] = v38
            }
            _ = v38
        }
        v2 += (v2 << ((v1) & 63))
        v1 = 182
        _ = fi35
    }
    switch 0 {
    case 0:
        if len(v33) > 0 {
            v33[0] = 9223372036854775807
        }
    case 42:
        v1 += (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
