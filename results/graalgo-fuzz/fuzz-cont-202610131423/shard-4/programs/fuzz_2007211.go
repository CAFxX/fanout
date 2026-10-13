package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint8) uint16 {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{}
        var v7 float32 = 0.0
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 += 197
    var v8 uint16 = 32767
    v8 = (v8 << ((0) & 63))
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int8 = 46
        v11 := GenS1{F0: 563}
        switch "foo bar" {
        case "":
            v10 = v10
        case "\n\t":
            if (42 <= 42) {
                fmt.Println(fi9, v10)
                switch "" {
                case "foo bar":
                    v12 := GenS1{F0: 0}
                    _ = v12
                case "a":
                    switch (fi9 * (fi9 & (-(fi9)))) {
                    case 256:
                        var v13 uint32 = 1
                        _ = v13
                    case 128:
                        v14 := GenS2{F0: fi9, F1: p0, F2: 591}
                        _ = v14
                    case -705:
                        p0 = (p0 | 42)
                    }
                case "hello":
                    var v15 int64 = 9223372036854775807
                    _ = v15
                }
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v8
    return v8
}
func genFn1() int8 {
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int16 = -1
        if true {
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := GenS2{F0: (fi16 >> ((255) & 63)), F1: 128, F2: 1}
                var v20 uint64 = 2
                _ = fi18
                _ = v19
                _ = v20
            }
            v21 := []int{32767}
            v22 := GenS3{F0: 74, F1: 1}
            _ = v21
            _ = v22
        }
        _ = fi16
        _ = v17
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 uint64 = 90
        _ = fi23
        _ = v24
    }
    return -128
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 uint16 = 848
    var v4 uint32 = 2147483648
    v25 := []int{668, -747, 4294967295, 128}
    v26 := GenS3{F0: v4, F1: int8(v4)}
    var v27 int = 0
    delete(v1, (v27 - v27))
    v3 -= 65535
    v3 -= 255
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 string = "a"
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: (v27 + (-(v27))), F1: 2, F2: 42}
    v3 = genFn0(255)
    v31 := v1[32767]
    fmt.Println(v27, v27, v27)
    var v32 float32 = 3.14159
    delete(v1, 321)
    for fi33 := 0; fi33 < 12; fi33++ {
        v34 := GenS2{F0: (fi33 - v31), F1: 223, F2: 70}
        for fi35 := 0; fi35 < 13; fi35++ {
            _ = len(v1)
            switch (("foo bar" + (((("foo bar" + "a") + "") + "a") + "hello")) + ((((("hello" + "a") + "foo bar") + "") + "a") + ("\n\t" + "a"))) {
            case "hello":
                switch 128 {
                case 127:
                    if (2147483647 >= 2147483647) {
                        var v36 int64 = 436
                        _ = v36
                    }
                default:
                    var v37 uint32 = (v4 % (((v4 - v4)) | 1))
                    _ = v37
                }
            case "":
                var v38 float32 = v32
                _ = v38
            case "foo bar":
                v32 += float32(1.5)
            }
            var v39 int32 = -573
            _ = fi35
            _ = v39
        }
        _ = fi33
        _ = v34
    }
    fmt.Println(v27, v32)
    fmt.Println(v31, v31, v27)
    v4 -= 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
