package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int32, p1 string) uint8 {
    fmt.Println(p0)
    var v6 uint8 = 208
    v7 := GenS1{F0: v6, F1: (false || false), F2: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (v6 | 255)
}
func genFn1(p0 int, p1 uint) uint64 {
    for fi8 := 0; fi8 < 7; fi8++ {
        if (0 == 65535) {
            var v9 int16 = 0
            p1 = 906
            _ = v9
        } else {
            p1 += p1
            var v10 bool = (3.14159 != (((0.1 - 1e10) * ((-2.25 / float32(1.5)) * 1e10)) / (1e10 * (1e10 / float32(-2.25)))))
            if true {
                var v11 bool = v10
                p1 = (^(p1))
                if (-9223372036854775808 > 1) {
                    p1 = p1
                    v12 := []int{42, 32767}
                    _ = v12
                }
                _ = v11
            }
            _ = v10
        }
        _ = fi8
    }
    var v13 int32 = -1
    var v14 float32 = 3.14159
    v15 := []int{p0}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 151
}

func main() {
    v1 := []uint32{0, 145, 42}
    var v2 uint16 = 440
    var v3 uint8 = 136
    v4 := make(map[string]bool)
    var v5 bool = (true || (("foo bar" + "foo bar") != "\n\t"))
    if true {
        var v16 int8 = 111
        _ = v16
    }
    v3 += 2
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint64 = 522
        v18 += (^(v18))
        var v19 int = 405
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: v3, F1: v5, F2: 611}
    var v21 int = 32768
    if v5 {
        if true {
            v22 := v4["foo bar"]
            v23 := GenS2{F0: 894}
            var v24 uint64 = genFn1((v21 & 9223372036854775807), 65535)
            _ = v22
            _ = v23
            _ = v24
        } else {
            var v25 float64 = 1e10
            var v26 uint = 0
            var v27 uint = v26
            _ = v25
            _ = v26
            _ = v27
        }
        var v28 uint8 = v3
        _ = v28
    }
    switch (v21 - (v21 / ((2147483647) | 1))) {
    case 32767:
        var v29 int16 = 32767
        _ = v29
    case 135:
        switch ("hello" + "hello") {
        case "\n\t":
            v2 -= v2
        case "foo bar":
            v30 := GenS3{F0: 4294967295}
            _ = v30
        case "hello":
            fmt.Println(v2)
        }
    default:
        if (-989 >= 1) {
            var v31 bool = true
            v21 += v21
            _ = v31
        } else {
            if (1 > 0) {
                var v32 int = int(v3)
                var v33 uint16 = uint16(v3)
                _ = v32
                _ = v33
            }
        }
    }
    var v34 float32 = (-2.25 / 1.5)
    var v35 uint = 0
    var v36 uint32 = 65535
    v3 = 2
    v5 = true
    var v37 uint16 = uint16(v3)
    var v38 uint8 = (v3 | 127)
    v39 := v4["\n\t"]
    fmt.Println(v38, v36)
    v37 = (v37 | v37)
    for fi40 := 0; fi40 < 17; fi40++ {
        if true {
            if false {
                v1 = append(v1, v36)
            } else {
                v36 = v36
                var v41 int32 = -586
                _ = v41
            }
            for fi42 := 0; fi42 < 20; fi42++ {
                var v43 string = "\n\t"
                var v44 uint8 = 198
                v45 := GenS2{F0: v36}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            var v46 uint8 = 128
            _ = v46
        }
        _ = fi40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
