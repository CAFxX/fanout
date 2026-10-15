package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int }
func genFn0() int64 {
    var v6 int16 = 1
    if (2 >= 0) {
        var v7 uint16 = 452
        v6 = (v6 + -860)
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{fi8, fi8}
        _ = fi8
        _ = v9
    }
    if true {
        fmt.Println(v6)
        var v10 uint8 = 127
        _ = v10
    }
    var v11 int64 = -9223372036854775808
    _ = v11
    _ = v6
    return v11
}
func genFn1(p0 int, p1 int32) uint {
    var v12 uint16 = 551
    switch p0 {
    case -837:
        switch "a" {
        case "hello":
            p1 -= int32(v12)
        default:
            v13 := []int{32768, 1, int(p1)}
            _ = v13
        }
    case 781:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint = 397
            for fi16 := 0; fi16 < 16; fi16++ {
                v17 := []int{fi16, p0}
                var v18 int8 = -1
                v15 = (v15 * v15)
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi14
            _ = v15
        }
    default:
        p1 = p1
    }
    v19 := []int{p0, (p0 % ((-1) | 1)), (p0 % ((-1) | 1)), (p0 / (((p0 + -2147483648)) | 1))}
    _ = p0
    _ = p1
    _ = v12
    _ = v19
    return 537
}

func main() {
    var v1 int8 = 100
    v2 := []int{}
    var v3 uint16 = 255
    v4 := make(map[string]int)
    var v5 uint16 = (v3 * (v3 ^ v3))
    fmt.Println(v3, v5, v5)
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 string = ""
        v22 := GenS1{F0: 1, F1: 255, F2: -434}
        var v23 uint = 18446744073709551615
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v2 = append(v2, int(v1))
    fmt.Println(v3, v5)
    for fi24 := 0; fi24 < 8; fi24++ {
        switch 128 {
        case -218:
            var v25 uint8 = 127
            _ = v25
        }
        var v26 uint32 = 431
        switch "hello" {
        case "foo bar":
            var v27 int32 = 1
            _ = v27
        case "a":
            if (0 != 0) {
                v28 := GenS1{F0: 9223372036854775807, F1: 2, F2: -32768}
                fmt.Println(v3)
                switch fi24 {
                case 0:
                    fmt.Println(fi24, v1, v5)
                case 32767:
                    for fi29 := 0; fi29 < 20; fi29++ {
                        v30 := GenS2{F0: 128, F1: 2}
                        var v31 int32 = 1
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                case 4294967295:
                    for fi32 := 0; fi32 < 1; fi32++ {
                        v33 := GenS3{F0: (v5 * 970), F1: genFn0(), F2: 32767}
                        if true {
                            v1 -= (v1 >> ((2) & 63))
                            v34 := GenS3{F0: v3, F1: -747, F2: fi24}
                            _ = v34
                        } else {
                            v35 := GenS3{F0: 560, F1: -1, F2: fi32}
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                    }
                }
                _ = v28
            }
        case "\n\t":
            var v36 int64 = 0
            _ = v36
        }
        _ = fi24
        _ = v26
    }
    v37 := GenS1{F0: -670, F1: 127, F2: 42}
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 uint64 = 0
        v5 += v5
        var v40 string = (("\n\t" + "hello") + "foo bar")
        _ = fi38
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        var v42 uint8 = 87
        _ = fi41
        _ = v42
    }
    v43 := GenS1{F0: 9223372036854775807, F1: 1, F2: 1}
    if true {
        v44 := v4[("a" + "")]
        v1 -= (v1 * 42)
        _ = v44
    }
    var v45 int64 = 0
    if false {
        v46 := GenS1{F0: (^(v45)), F1: 42, F2: 0}
        switch ("" + "a") {
        case "a":
            var v47 string = ""
            _ = v47
        case "":
            if false {
                fmt.Println(v5, v5, v3)
            }
        }
        _ = v46
    } else {
        switch -645 {
        case -2147483648:
            var v48 int16 = 42
            _ = v48
        case 438:
            switch 127 {
            case -689:
                v45 = (v45 >> ((656) & 63))
            case 42:
                for fi49 := 0; fi49 < 2; fi49++ {
                    switch (fi49 & (fi49 << ((18446744073709551615) & 63))) {
                    case 1:
                        v50 := GenS3{F0: 255, F1: 934, F2: fi49}
                        _ = v50
                    }
                    v51 := GenS3{F0: (v3 * (^(v5))), F1: (^(v45)), F2: 4294967295}
                    _ = fi49
                    _ = v51
                }
            case 128:
                fmt.Println(v45, v5)
            }
        case 2147483647:
            var v52 float32 = 0.1
            _ = v52
        }
    }
    var v53 int64 = (^(v45))
    v53 += genFn0()
    v45 += (v53 / ((v45) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v45
    _ = v5
    _ = v53
}
