package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int, p1 int32) uint64 {
    switch p0 {
    case -784:
        v4 := []int{p0}
        _ = v4
    case 42:
        var v5 int32 = p1
        _ = v5
    case 4294967295:
        var v6 int16 = -32768
        _ = v6
    default:
        if (2 != 18446744073709551615) {
            for fi7 := 0; fi7 < 12; fi7++ {
                p0 -= (fi7 ^ (p0 - (fi7 - 4294967295)))
                if false {
                    v8 := GenS1{F0: -954, F1: 0}
                    var v9 int64 = -531
                    v10 := []int{p0}
                    _ = v10
                    _ = v8
                    _ = v9
                }
                var v11 int16 = -96
                _ = fi7
                _ = v11
            }
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: -9223372036854775808, F1: 67}
                fmt.Println(p1)
                var v14 string = ""
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := []int{p0}
            _ = v15
        }
    }
    v16 := GenS2{F0: 800}
    switch ("a" + "foo bar") {
    case "":
        if false {
            v17 := []int{-851, 847, (p0 - (-(p0)))}
            if false {
                p0 += (p0 % ((127) | 1))
            }
            _ = v17
        } else {
            for fi18 := 0; fi18 < 15; fi18++ {
                for fi19 := 0; fi19 < 10; fi19++ {
                    fmt.Println(p1)
                    switch (-(fi19)) {
                    case 919:
                        var v20 uint16 = 32767
                        _ = v20
                    case 128:
                        p0 += fi18
                    }
                    _ = fi19
                }
                for fi21 := 0; fi21 < 6; fi21++ {
                    if (true && true) {
                        var v22 int8 = -115
                        var v23 int32 = p1
                        if false {
                            fmt.Println(v23)
                            var v24 uint16 = 255
                            _ = v24
                        }
                        if false {
                            var v25 uint = 65535
                            _ = v25
                        } else {
                            fmt.Println(fi18, fi21, p0)
                            var v26 string = ""
                            var v27 uint32 = 2147483648
                            fmt.Println(v26, fi18, v22)
                            _ = v26
                            _ = v27
                        }
                        var v28 int8 = 0
                        _ = v22
                        _ = v23
                        _ = v28
                    } else {
                        p0 = fi18
                    }
                    _ = fi21
                }
                if false {
                    p1 = (p1 >> ((42) & 63))
                }
                _ = fi18
            }
            var v29 uint32 = 65535
            var v30 uint8 = 25
            _ = v29
            _ = v30
        }
    default:
        switch (^(p0)) {
        case 993:
            fmt.Println(p0, p1)
        case -909:
            switch "foo bar" {
            case "":
                var v31 uint = 255
                _ = v31
            case "hello":
                var v32 int16 = -564
                _ = v32
            }
        default:
            v33 := []int{p0, (p0 ^ -28)}
            _ = v33
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v16
    return 832
}
func genFn1(p0 int16, p1 uint8, p2 string) int64 {
    if ((true || (false || true)) || ((0.1 * 1e10) != float64(3.14159))) {
        var v34 int16 = (-(p0))
        _ = v34
    } else {
        v35 := []int{448, -9223372036854775808, 32768}
        if true {
            fmt.Println(p2, p1)
            fmt.Println(p2)
            p2 = "a"
        } else {
            var v36 uint64 = 1
            _ = v36
        }
        _ = v35
    }
    var v37 int16 = p0
    if true {
        p0 -= v37
    } else {
        if ((false || true) && (4294967295 == 65535)) {
            var v38 int64 = 217
            _ = v38
        } else {
            var v39 int = 651
            switch (p2 + "a") {
            case "foo bar":
                var v40 int16 = v37
                _ = v40
            default:
                v41 := GenS1{F0: -378, F1: 114}
                _ = v41
            }
            _ = v39
        }
        v42 := GenS1{F0: -451, F1: 100}
        _ = v42
    }
    var v43 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v43
    return 851
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = -2.25
    var v3 float32 = v2
    v3 += ((v2 * float32(1e10)) - (float32(0.0) - (3.14159 / ((v3 * (3.14159 / (((v3 / ((v2 / float32((float64(v3) / (-2.25 - 1e10)))) - 0.0)) / 1.5) / v2))) + v2))))
    v44 := GenS1{F0: -1, F1: -128}
    if true {
        v3 -= 3.14159
        v45 := GenS3{F0: -29}
        _ = v45
    } else {
        _ = len(v1)
    }
    var v46 uint32 = 2147483648
    v46 = (v46 ^ (v46 | (v46 << ((18446744073709551615) & 63))))
    v2 = 1e10
    fmt.Println(v2)
    v47 := []int{-9223372036854775808, 32768}
    v48 := v1[0]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v46
    _ = v47
    _ = v48
}
