package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 string; F2 int16 }
type GenS4 struct { F0 string; F1 int16; F2 uint64 }
func genFn0(p0 int) int {
    v7 := GenS1{F0: -738, F1: "", F2: -9223372036854775808}
    p0 -= p0
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(p0, p0)
        v9 := []int{-2147483648, p0, (p0 << ((390) & 63))}
        _ = fi8
        _ = v9
    }
    v10 := []int{-9223372036854775808}
    _ = p0
    _ = v10
    _ = v7
    return p0
}
func genFn1(p0 int, p1 uint32, p2 uint16) int64 {
    v11 := GenS2{F0: (p1 - (p1 * p1)), F1: p2, F2: (p2 | p2)}
    if (true && false) {
        var v12 int32 = 0
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[uint32]bool)
    var v3 string = ""
    var v4 uint8 = 2
    var v5 bool = (!("hello" > "\n\t"))
    var v6 float64 = (0.0 + 1.5)
    v2[225] = (0 == 120)
    v1 = (!v5)
    var v13 uint64 = 18446744073709551615
    v14 := GenS1{F0: 42, F1: v3, F2: 9223372036854775807}
    for fi15 := 0; fi15 < 3; fi15++ {
        switch fi15 {
        case -2147483648:
            v2[0] = (9223372036854775807 != -9223372036854775808)
        case 32768:
            for fi16 := 0; fi16 < 8; fi16++ {
                v5 = (1 >= -1)
                var v17 uint64 = v13
                fmt.Println(v13, v5)
                var v18 int32 = 0
                v18 -= -595
                _ = fi16
                _ = v17
                _ = v18
            }
        default:
            if true {
                var v19 int64 = 995
                _ = v19
            } else {
                if (!v5) {
                    v13 = (^(v13))
                    switch (fi15 * (^(fi15))) {
                    case 0:
                        v4 -= v4
                    }
                }
                v20 := v2[2147483647]
                switch ((v3 + (v3 + "\n\t")) + v3) {
                case "a":
                    var v21 float64 = v6
                    _ = v21
                }
                _ = v20
            }
        }
        v6 = 1.5
        var v22 uint = 1
        _ = fi15
        _ = v22
    }
    switch genFn0(-9223372036854775808) {
    case -180:
        v23 := GenS3{F0: v4, F1: v3, F2: 42}
        _ = v23
    case 589:
        delete(v2, 2147483647)
    }
    if (798 <= 32767) {
        var v24 uint = 148
        if v1 {
            for fi25 := 0; fi25 < 7; fi25++ {
                v3 = ""
                _ = fi25
            }
        }
        _ = v24
    } else {
        v26 := GenS3{F0: v4, F1: v3, F2: 32767}
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        v28 := []int{9223372036854775807, genFn0((fi27 + (fi27 * (-(fi27))))), (fi27 >> ((0) & 63)), fi27}
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint64 = (v13 | (v13 * (v13 / (((v13 & 0)) | 1))))
        v1 = (!(42 != 564))
        switch "" {
        case "hello":
            if false {
                v4 -= 140
            }
        case "":
            for fi31 := 0; fi31 < 11; fi31++ {
                v1 = v1
                var v32 float32 = 0.1
                _ = fi31
                _ = v32
            }
        }
        _ = fi29
        _ = v30
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := v2[4294967295]
        for fi35 := 0; fi35 < 11; fi35++ {
            v36 := []int{(fi35 / (((fi33 - (fi35 - 2147483647))) | 1)), 2147483647, (fi35 << ((534) & 63)), fi33}
            v4 = 128
            _ = fi35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    switch ("foo bar" + v3) {
    case "\n\t":
        var v37 float32 = 1e10
        _ = v37
    case "foo bar":
        v38 := GenS1{F0: -155, F1: "foo bar", F2: 42}
        _ = v38
    }
    v39 := GenS3{F0: v4, F1: "hello", F2: 32767}
    var v40 int32 = 1
    if (388 >= 32767) {
        var v41 uint32 = uint32(v4)
        if v5 {
            v41 += v41
            if (v40 == v40) {
                v40 = v40
                v41 = (^(v41))
                v13 += 4294967295
            }
        }
        var v42 int32 = -294
        _ = v41
        _ = v42
    }
    var v43 int = -1
    if (false && false) {
        var v44 int16 = 42
        v45 := GenS4{F0: v3, F1: 1000, F2: (v13 & 4294967295)}
        _ = v44
        _ = v45
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
