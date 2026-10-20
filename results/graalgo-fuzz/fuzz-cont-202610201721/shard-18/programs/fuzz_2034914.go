package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int8; F2 int64 }
func genFn0(p0 int8) int8 {
    v6 := GenS1{F0: p0}
    if false {
        switch 42 {
        case -226:
            if false {
                p0 = 100
                var v7 int8 = (^(p0))
                _ = v7
            } else {
                var v8 int16 = -556
                _ = v8
            }
        case -2147483648:
            v9 := GenS1{F0: (p0 * (p0 * (p0 << ((255) & 63))))}
            _ = v9
        case -112:
            p0 -= p0
        }
        var v10 int16 = 0
        v11 := []int{32767}
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                switch "\n\t" {
                case "":
                    v14 := GenS1{F0: 42}
                    _ = v14
                }
                switch 32768 {
                case 643:
                    v15 := GenS1{F0: 127}
                    _ = v15
                case -940:
                    fmt.Println(p0)
                case 580:
                    v16 := []int{42, fi12, 2147483647, -350}
                    _ = v16
                }
                _ = fi13
            }
            var v17 float32 = 3.14159
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint = 1
                var v20 uint8 = 255
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi12
            _ = v17
        }
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 float32 = 0.1
            v23 = v23
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v24 := GenS2{F0: 2}
    _ = p0
    _ = v24
    _ = v6
    return 62
}
func genFn1(p0 int8, p1 uint16) int {
    v25 := GenS3{F0: "foo bar", F1: -2, F2: 1}
    switch ("hello" + ("" + "foo bar")) {
    case "foo bar":
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 uint32 = 2147483647
            switch (-(fi26)) {
            case 9223372036854775807:
                var v28 uint16 = p1
                _ = v28
            case 42:
                v27 = 497
            case -2147483648:
                v29 := []int{-2147483648, (fi26 & 127)}
                _ = v29
            }
            _ = fi26
            _ = v27
        }
    case "\n\t":
        p1 = p1
    default:
        var v30 int = 128
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v25
    return 4294967295
}

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 int16 = 32767
    var v31 uint64 = 2
    var v32 bool = (!true)
    v5 -= v5
    var v33 int32 = 128
    v1 += (v1 - (v1 / (((v33 - -1)) | 1)))
    var v34 uint64 = v31
    v33 = -2147483648
    v5 = v5
    v32 = v32
    for fi35 := 0; fi35 < 7; fi35++ {
        if true {
            v33 += 797
            switch (fi35 % ((int(v5)) | 1)) {
            case 42:
                v5 = 138
            case -9223372036854775808:
                fmt.Println(v32, v34, v31)
            }
        }
        _ = fi35
    }
    switch 643 {
    case 1:
        if false {
            var v36 int = 9223372036854775807
            _ = v36
        } else {
            var v37 uint8 = 42
            _ = v37
        }
    default:
        var v38 int16 = (v5 - 32767)
        _ = v38
    }
    v4[-9223372036854775808] = 927
    var v39 uint8 = 1
    if true {
        _ = len(v3)
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        if (((128 > 255) && v32) || ((42 < uint32(v39)) && v32)) {
            for fi41 := 0; fi41 < 14; fi41++ {
                switch fi40 {
                case 842:
                    v33 = 450
                default:
                    v42 := GenS2{F0: (v39 << ((374) & 63))}
                    _ = v42
                }
                fmt.Println(fi41)
                _ = fi41
            }
        }
        v43 := []int{(fi40 | genFn1(127, 691))}
        _ = fi40
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v5
}
