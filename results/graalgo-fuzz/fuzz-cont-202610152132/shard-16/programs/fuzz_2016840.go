package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]int)
    var v3 uint = 568
    if false {
        switch -9223372036854775808 {
        case 1:
            v4 := GenS1{F0: 65535, F1: (v3 << ((65535) & 63))}
            _ = v4
        }
    } else {
        v3 -= (v3 % (((v3 + (v3 | v3))) | 1))
        for fi5 := 0; fi5 < 10; fi5++ {
            switch "\n\t" {
            case "hello":
                v6 := GenS2{F0: -752, F1: -2147483648, F2: (v1 * int64(fi5))}
                _ = v6
            case "":
                v7 := GenS3{F0: (0.1 >= 0.0)}
                _ = v7
            case "a":
                switch fi5 {
                case 127:
                    v8 := []int{}
                    _ = v8
                }
            default:
                v3 = 365
            }
            for fi9 := 0; fi9 < 9; fi9++ {
                if true {
                    v10 := GenS4{F0: int(v1), F1: "foo bar", F2: 255}
                    v1 = (-(v1))
                    v3 += v3
                    _ = v10
                }
                var v11 int8 = 66
                v11 -= v11
                _ = fi9
                _ = v11
            }
            _ = fi5
        }
        var v12 int8 = 100
        _ = v12
    }
    v13 := GenS3{F0: (true && (true || false))}
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(v3, v1, fi14)
            _ = fi14
        }
        var v15 uint32 = 2147483648
        _ = len(v2)
        _ = v15
    } else {
        _ = len(v2)
        for fi16 := 0; fi16 < 18; fi16++ {
            v1 = v1
            if true {
                var v17 uint64 = 0
                _ = v17
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                v1 -= (v1 % (((v1 >> ((v3) & 63))) | 1))
                if (false || true) {
                    v19 := GenS2{F0: 32767, F1: -2147483648, F2: v1}
                    _ = v19
                } else {
                    v3 += (v3 / ((v3) | 1))
                }
                v20 := GenS2{F0: -32768, F1: 2147483647, F2: v1}
                _ = fi18
                _ = v20
            }
            _ = fi16
        }
        if true {
            for fi21 := 0; fi21 < 1; fi21++ {
                if true {
                    v22 := GenS5{F0: 127}
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v1 -= 0
                        var v24 uint64 = 1
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                } else {
                    v3 += (v3 ^ v3)
                    v25 := []int{(fi21 ^ -1), (^(fi21)), (fi21 % (((fi21 | 32768)) | 1))}
                    v26 := v2["a"]
                    _ = v25
                    _ = v26
                }
                _ = fi21
            }
            for fi27 := 0; fi27 < 4; fi27++ {
                v2[""] = fi27
                var v28 uint16 = 2
                if ((!(((-2.25 - 1e10) / -2.25) != ((1.5 - ((((float32(1e10) / (((1.5 - 1.5) / 3.14159) + 1.5)) * 0.1) + ((float32((0.1 - -2.25)) * (-2.25 - 1.5)) - 1.5)) / 3.14159)) / (-2.25 * ((1.5 * float32(1.5)) + 1e10))))) && true) {
                    var v29 uint64 = 216
                    var v30 int32 = 1
                    _ = v29
                    _ = v30
                } else {
                    if (!true) {
                        v28 = (v28 + 2)
                        var v31 float32 = 1.5
                        _ = v31
                    } else {
                        v3 = 233
                        v1 = v1
                    }
                    v28 += 255
                    for fi32 := 0; fi32 < 17; fi32++ {
                        v28 += v28
                        switch (fi27 << (((^(v3))) & 63)) {
                        case 42:
                            v3 -= v3
                        default:
                            var v33 uint = v3
                            _ = v33
                        }
                        switch fi32 {
                        case 128:
                            v1 -= v1
                        }
                        v34 := []int{fi32, fi32, (fi27 * (fi32 * (fi27 | 128))), fi32}
                        _ = fi32
                        _ = v34
                    }
                }
                _ = fi27
                _ = v28
            }
            v35 := GenS3{F0: true}
            _ = v35
        }
    }
    fmt.Println(v1, v1)
    v36 := []int{-538, 2147483647, 32768, -707}
    for fi37 := 0; fi37 < 5; fi37++ {
        if ((fi37 >> ((v3) & 63)) == 141) {
            v3 -= (v3 | v3)
            var v38 int8 = 100
            switch "\n\t" {
            case "":
                v1 = 42
            case "\n\t":
                var v39 int8 = -1
                _ = v39
            default:
                var v40 int8 = -1
                _ = v40
            }
            _ = v38
        }
        _ = fi37
    }
    v41 := v2[("" + "")]
    v42 := GenS4{F0: -601, F1: "foo bar", F2: 476}
    v1 -= v1
    _ = len(v36)
    v43 := GenS6{F0: uint64(v3), F1: 104}
    var v44 int8 = -1
    v36 = append(v36, (v41 >> ((42) & 63)))
    var v45 uint64 = 42
    v46 := GenS3{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
