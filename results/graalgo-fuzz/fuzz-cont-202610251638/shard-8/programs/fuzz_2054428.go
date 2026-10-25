package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 bool; F1 uint8; F2 string }
type GenS6 struct { F0 uint32; F1 int }
func genFn0(p0 int64) uint {
    p0 -= p0
    fmt.Println(p0)
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 42, F1: (fi4 ^ (fi4 >> ((4294967295) & 63))), F2: 633}
        p0 -= (p0 ^ p0)
        p0 += 9223372036854775807
        _ = fi4
        _ = v5
    }
    _ = p0
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 int16 = -32768
    fmt.Println(v3)
    for fi6 := 0; fi6 < 1; fi6++ {
        v3 = v3
        v3 = 0
        switch fi6 {
        case 32767:
            fmt.Println(fi6, fi6)
        case 132:
            var v7 int = fi6
            _ = v7
        case 42:
            v3 += -289
        }
        _ = fi6
    }
    if false {
        switch "a" {
        case "hello":
            v8 := []int{-875, 32768, -2147483648, 11}
            _ = v8
        case "foo bar":
            fmt.Println(v3)
        case "a":
            switch "a" {
            case "a":
                var v9 int = -9223372036854775808
                _ = v9
            case "":
                v10 := GenS2{F0: -505, F1: 42}
                _ = v10
            default:
                switch -1 {
                case -764:
                    v11 := GenS2{F0: 42, F1: 4294967295}
                    _ = v11
                case 1:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v13 := GenS3{F0: "foo bar", F1: -2147483648, F2: 0}
                        v14 := []int{}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 4294967295:
                    var v15 int16 = 0
                    _ = v15
                }
            }
        }
    } else {
        v3 = (^(v3))
        var v16 int = 208
        v3 -= v3
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(fi17, fi17)
        _ = fi17
    }
    fmt.Println(v3)
    switch 2147483647 {
    case -338:
        v18 := GenS4{F0: 255, F1: ("\n\t" + "hello")}
        _ = v18
    default:
        fmt.Println(v3)
    }
    v3 = (v3 ^ (v3 - v3))
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 float32 = (((1.5 - -2.25) + (1e10 / ((3.14159 / 1.0) - (0.0 * 3.14159)))) - -2.25)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        if false {
            v22 := GenS3{F0: "hello", F1: -2147483648, F2: 1}
            for fi23 := 0; fi23 < 6; fi23++ {
                if true {
                    v24 := GenS1{F0: 994, F1: 2147483647, F2: 2}
                    _ = v24
                } else {
                    var v25 string = "hello"
                    _ = v25
                }
                if (true && false) {
                    v26 := GenS5{F0: true, F1: 42, F2: "a"}
                    _ = v26
                }
                _ = fi23
            }
            _ = v22
        } else {
            var v27 uint8 = 128
            _ = v27
        }
        _ = fi21
    }
    var v28 int = 696
    for fi29 := 0; fi29 < 1; fi29++ {
        _ = len(v1)
        _ = len(v1)
        _ = fi29
    }
    var v30 uint32 = 65535
    v3 -= 42
    var v31 uint8 = uint8(v30)
    if ((v28 ^ (v28 ^ (v28 + v28))) > (v28 + v28)) {
        var v32 uint16 = uint16(v31)
        _ = v32
    } else {
        v30 = 2147483648
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 int64 = -9223372036854775808
            for fi35 := 0; fi35 < 1; fi35++ {
                switch int(v3) {
                case 115:
                    v30 -= v30
                    var v36 int64 = (v34 * v34)
                    if (!(18446744073709551615 <= 1)) {
                        v31 = (v31 << ((403) & 63))
                        delete(v2, -745)
                    } else {
                        fmt.Println(v36)
                        if (!((v28 >> ((1) & 63)) < -680)) {
                            var v37 uint8 = (v31 >> ((65535) & 63))
                            v38 := []int{-137, (v28 | 233), (v28 + v28), 0}
                            _ = v37
                            _ = v38
                        } else {
                            v36 += 4
                            v36 = (v36 - v36)
                        }
                        v39 := GenS5{F0: (v31 != 2), F1: 128, F2: ""}
                        v28 = fi35
                        _ = v39
                    }
                    _ = v36
                default:
                    v40 := GenS6{F0: (v30 >> ((42) & 63)), F1: fi35}
                    _ = v40
                }
                v41 := v1[v30]
                v28 += (fi33 & v28)
                v42 := v1[2]
                v3 -= v3
                _ = fi35
                _ = v41
                _ = v42
            }
            if ("a" > "foo bar") {
                fmt.Println(v34, v30, v3)
            } else {
                v43 := v1[v30]
                for fi44 := 0; fi44 < 2; fi44++ {
                    var v45 uint16 = 0
                    v3 = v3
                    _ = fi44
                    _ = v45
                }
                v30 = 2147483647
                _ = v43
            }
            _ = fi33
            _ = v34
        }
        var v46 float64 = -2.25
        _ = v46
    }
    var v47 int16 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v47
}
