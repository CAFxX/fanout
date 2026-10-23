package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint8, p1 int32) uint64 {
    v5 := GenS1{F0: 32767, F1: 9223372036854775807}
    v6 := []int{32768}
    if (false && (!false)) {
        v6 = append(v6, 273)
    } else {
        var v7 float32 = 3.14159
        _ = len(v6)
        _ = v7
    }
    if false {
        if false {
            var v8 uint32 = 2147483648
            v9 := GenS1{F0: 336, F1: 0}
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int8) uint32 {
    v10 := GenS1{F0: 231, F1: -9223372036854775808}
    var v11 int = 0
    var v12 uint64 = 1
    v13 := []int{v11, (-(v11)), (v11 | (v11 >> ((18446744073709551615) & 63))), 128}
    if (2147483647 != 42) {
        for fi14 := 0; fi14 < 5; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                _ = len(v13)
                _ = fi15
            }
            if (423 >= 902) {
                v16 := GenS2{F0: (v11 << ((255) & 63)), F1: 2}
                var v17 bool = (672 >= 2147483648)
                v18 := GenS1{F0: uint16(v12), F1: -1}
                _ = v16
                _ = v17
                _ = v18
            } else {
                if true {
                    if len(v13) > 0 {
                        v13[0] = (^(fi14))
                    }
                }
                fmt.Println(v12)
            }
            fmt.Println(fi14, v11)
            _ = fi14
        }
    } else {
        v13 = append(v13, (v11 ^ v11))
        var v19 int8 = p0
        p0 += (^(v19))
        _ = v19
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 258
}

func main() {
    var v1 int32 = 423
    var v2 uint8 = 1
    var v3 int = 0
    v4 := make(map[int]uint32)
    var v20 int64 = 42
    if false {
        fmt.Println(v1, v20, v2)
        var v21 int16 = 734
        _ = v21
    }
    fmt.Println(v2, v2)
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 uint64 = 18446744073709551615
        if false {
            v24 := GenS2{F0: (v3 >> ((18446744073709551615) & 63)), F1: 468}
            var v25 int16 = 42
            v26 := []int{}
            _ = v24
            _ = v25
            _ = v26
        }
        v27 := GenS1{F0: 42, F1: v20}
        _ = fi22
        _ = v23
        _ = v27
    }
    v1 = v1
    for fi28 := 0; fi28 < 19; fi28++ {
        switch (fi28 ^ (fi28 + -828)) {
        case -1:
            v20 -= (^(v20))
        case 9223372036854775807:
            if true {
                switch ("a" + "") {
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        if (1 <= -1) {
                            v29 := []int{fi28, 9223372036854775807, v3, v3}
                            _ = v29
                        } else {
                            fmt.Println(v2)
                        }
                    }
                case "a":
                    _ = len(v4)
                }
                v1 += (v1 % (((v1 | 42)) | 1))
            } else {
                for fi30 := 0; fi30 < 13; fi30++ {
                    var v31 uint32 = 0
                    _ = fi30
                    _ = v31
                }
                v32 := []int{v3, -799, v3}
                v33 := GenS2{F0: v3, F1: 330}
                _ = v32
                _ = v33
            }
        }
        if (true && false) {
            switch "\n\t" {
            case "hello":
                var v34 uint8 = (^(v2))
                _ = v34
            }
            v2 = v2
        } else {
            fmt.Println(v20, v1, v2)
        }
        _ = fi28
    }
    if false {
        var v35 uint16 = 779
        switch v3 {
        case 826:
            v36 := GenS3{F0: v35, F1: 1000, F2: (v35 & (v35 & 473))}
            _ = v36
        case 42:
            v37 := GenS4{F0: true, F1: v2}
            _ = v37
        }
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 float32 = (((1.5 + float32(3.14159)) * ((1.5 + -2.25) - (((1e10 + float32(3.14159)) - (0.1 / 1.0)) * float32(((-2.25 + 0.1) - 1e10))))) / 1.0)
            var v40 int64 = v20
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = v35
    } else {
        var v41 uint = 884
        v42 := GenS1{F0: 42, F1: (-(v20))}
        _ = v41
        _ = v42
    }
    v2 -= v2
    v43 := []int{}
    v44 := GenS3{F0: 0, F1: -32768, F2: 42}
    v45 := GenS4{F0: (32767 < -673), F1: v2}
    v2 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
