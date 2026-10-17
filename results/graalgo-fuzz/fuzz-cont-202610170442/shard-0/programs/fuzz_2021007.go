package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint64 }
func genFn0() int {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int16 = -32768
        var v7 int64 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: 94, F2: 128}
    if (1 > 0) {
        if (-2.25 <= (0.0 / float32((1.5 / 0.1)))) {
            var v9 uint8 = 78
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 uint32 = 4294967295
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS1{F0: fi12, F1: 2, F2: 2}
            var v14 uint = 2
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        v15 := GenS1{F0: -742, F1: 42, F2: 184}
        var v16 int32 = -2147483648
        _ = v15
        _ = v16
    }
    _ = v8
    return -651
}
func genFn1(p0 uint16, p1 uint64, p2 int32) int {
    var v17 int32 = p2
    switch "hello" {
    case "foo bar":
        p0 = p0
    case "hello":
        switch 128 {
        case -17:
            switch "a" {
            case "a":
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 uint = 4294967295
                    _ = fi18
                    _ = v19
                }
            }
        case 322:
            for fi20 := 0; fi20 < 3; fi20++ {
                v17 -= (^(v17))
                if true {
                    v21 := []int{42, genFn0()}
                    fmt.Println(v17, p2, v17)
                    v22 := GenS2{F0: fi20, F1: 255, F2: true}
                    _ = v21
                    _ = v22
                } else {
                    p1 = p1
                    v23 := GenS1{F0: fi20, F1: 128, F2: 1}
                    var v24 int32 = (^(p2))
                    _ = v23
                    _ = v24
                }
                fmt.Println(fi20, v17, p2)
                _ = fi20
            }
        case 756:
            for fi25 := 0; fi25 < 8; fi25++ {
                v26 := []int{int(p0), fi25, (fi25 | fi25)}
                if false {
                    v26 = append(v26, fi25)
                }
                _ = fi25
                _ = v26
            }
        }
    }
    p0 = (p0 * p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 128
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = (float32(float64(3.14159)) + (-2.25 / 1.5))
    var v3 uint32 = 4294967295
    v4 := make(map[string]uint32)
    v2 = v2
    v3 -= (v3 | v3)
    v3 = (^(v3))
    v27 := v1["a"]
    if ((-228 == 32767) || true) {
        var v28 uint8 = 50
        _ = v28
    }
    if (-2147483648 <= 311) {
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 float32 = v2
            if true {
                for fi31 := 0; fi31 < 12; fi31++ {
                    fmt.Println(fi31, v30)
                    var v32 float64 = ((0.0 / -2.25) + -2.25)
                    delete(v1, (v27 + (v27 + v27)))
                    var v33 uint64 = 345
                    switch fi31 {
                    case 9223372036854775807:
                        fmt.Println(v33, v2, v3)
                    }
                    if (!(-1 >= 700)) {
                        var v34 uint64 = uint64(v3)
                        v35 := GenS3{F0: -1}
                        _ = v34
                        _ = v35
                    }
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            } else {
                var v36 uint = 1
                _ = v36
            }
            _ = fi29
            _ = v30
        }
    }
    v37 := GenS4{F0: v3, F1: -830, F2: 1}
    fmt.Println(v3)
    if ((false && false) && false) {
        switch (v27 + "a") {
        case "":
            v38 := GenS1{F0: -67, F1: 1, F2: 20}
            _ = v38
        case "\n\t":
            fmt.Println(v3, v3)
        case "foo bar":
            var v39 uint8 = 0
            _ = v39
        default:
            var v40 uint64 = 4294967295
            _ = v40
        }
        v41 := v4[(v27 + "foo bar")]
        for fi42 := 0; fi42 < 10; fi42++ {
            delete(v1, v27)
            _ = fi42
        }
        _ = v41
    } else {
        var v43 uint64 = 0
        v44 := GenS3{F0: 9223372036854775807}
        v2 = v2
        _ = v43
        _ = v44
    }
    v45 := v1[((v27 + "hello") + v27)]
    fmt.Println(v3, v3, v2)
    for fi46 := 0; fi46 < 3; fi46++ {
        v47 := []int{(fi46 << ((327) & 63)), (fi46 / (((fi46 ^ 2147483647)) | 1)), (fi46 | 1), genFn0()}
        _ = fi46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v45
}
