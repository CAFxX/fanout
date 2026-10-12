package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0(p0 int, p1 bool) int8 {
    var v4 int8 = 100
    v5 := GenS1{F0: (v4 + v4), F1: v4}
    for fi6 := 0; fi6 < 16; fi6++ {
        switch "foo bar" {
        case "a":
            v7 := []int{9223372036854775807, p0, (fi6 << ((2) & 63)), (fi6 - (^(p0)))}
            _ = v7
        case "foo bar":
            switch "foo bar" {
            case "hello":
                if ((-2147483648 == 2147483647) && false) {
                    fmt.Println(p1)
                    var v8 uint32 = 846
                    _ = v8
                } else {
                    v9 := []int{(p0 * (p0 | 32768)), 32768, -9223372036854775808}
                    var v10 int8 = 42
                    var v11 uint64 = 18446744073709551615
                    _ = v10
                    _ = v11
                    _ = v9
                }
            }
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            switch ("a" + ("a" + "\n\t")) {
            case "foo bar":
                if p1 {
                    fmt.Println(p0)
                    var v13 int64 = 0
                    var v14 string = (("\n\t" + "foo bar") + ("" + (("\n\t" + ("\n\t" + "\n\t")) + "hello")))
                    _ = v13
                    _ = v14
                }
            }
            var v15 float64 = -2.25
            _ = fi12
            _ = v15
        }
        _ = fi6
    }
    p1 = false
    if true {
        for fi16 := 0; fi16 < 8; fi16++ {
            fmt.Println(v4, fi16, fi16)
            var v17 int = -113
            _ = fi16
            _ = v17
        }
        var v18 int = -912
        _ = v18
    } else {
        var v19 float64 = (1.5 + 1e10)
        p1 = (255 == 4294967295)
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return v4
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 42
    var v3 uint = v2
    delete(v1, -1)
    var v20 int16 = 0
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 uint16 = 65535
            var v24 int = (fi21 * (fi21 & fi21))
            var v25 int32 = -2147483648
            v26 := v1[32767]
            v27 := GenS2{F0: v26, F1: 2, F2: 0}
            v28 := GenS1{F0: -110, F1: 0}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
            _ = v26
            _ = v27
            _ = v28
        }
        fmt.Println(fi21)
        var v29 uint8 = 42
        _ = fi21
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 uint32 = 921
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := GenS2{F0: (v31 / ((v31) | 1)), F1: 162, F2: (fi30 - (fi30 % (((fi32 % ((-9223372036854775808) | 1))) | 1)))}
            fmt.Println(v3, fi32)
            v3 = v2
            v34 := GenS3{F0: 1, F1: -1}
            _ = fi32
            _ = v33
            _ = v34
        }
        v31 = 1
        _ = fi30
        _ = v31
    }
    v3 += (v2 ^ (v2 >> ((v2) & 63)))
    v20 -= 1
    v35 := GenS1{F0: 79, F1: genFn0(-9223372036854775808, false)}
    delete(v1, 341)
    var v36 uint16 = 42
    if true {
        v36 += (v36 >> (((v3 + v3)) & 63))
        var v37 int8 = 77
        var v38 bool = (false && (62 == 4294967295))
        _ = v37
        _ = v38
    } else {
        fmt.Println(v3)
        fmt.Println(v36, v3)
        if ((float64(0.1) / (-2.25 + 0.1)) >= -2.25) {
            fmt.Println(v3)
        }
    }
    if false {
        var v39 uint64 = 794
        _ = v39
    } else {
        for fi40 := 0; fi40 < 7; fi40++ {
            var v41 uint16 = v36
            _ = fi40
            _ = v41
        }
    }
    var v42 uint = 65535
    v43 := GenS2{F0: 829, F1: 42, F2: 128}
    v2 -= 18446744073709551615
    fmt.Println(v2)
    var v44 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v42
    _ = v43
    _ = v44
}
