package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint8; F2 bool }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int8; F1 int8; F2 uint32 }
type GenS8 struct { F0 int; F1 uint64; F2 uint16 }
func genFn0(p0 uint64, p1 string) uint32 {
    var v7 float32 = float32(0.0)
    switch (p1 + p1) {
    case "hello":
        p1 = "a"
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 float64 = -2.25
        fmt.Println(fi8, v9)
        v10 := GenS1{F0: 2}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(p0, v7, v7)
    if (p1 >= p1) {
        v11 := GenS2{F0: 0}
        v12 := []int{0, 638}
        var v13 int8 = 1
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return 2147483648
}
func genFn1() bool {
    switch "a" {
    case "hello":
        v14 := []int{377}
        _ = v14
    case "\n\t":
        if (true || (4294967295 > 2)) {
            var v15 uint16 = 1
            var v16 bool = (1 > 1)
            for fi17 := 0; fi17 < 18; fi17++ {
                if (0 == 100) {
                    v18 := GenS3{F0: 1, F1: 1}
                    v19 := GenS4{F0: fi17}
                    v20 := []int{32767, fi17, fi17, fi17}
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS3{F0: 284, F1: 782}
                    _ = v21
                }
                var v22 uint64 = 4294967295
                v23 := []int{fi17, fi17, (fi17 * fi17)}
                _ = fi17
                _ = v22
                _ = v23
            }
            _ = v15
            _ = v16
        }
    case "a":
        var v24 float32 = 1.5
        _ = v24
    default:
        if true {
            if ((65535 > 866) || ((true || false) || (831 != -32768))) {
                switch 4294967295 {
                case 42:
                    v25 := []int{-9223372036854775808, 32768, -1}
                    _ = v25
                case 4294967295:
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v27 := GenS5{F0: 970, F1: 2, F2: true}
                        _ = fi26
                        _ = v27
                    }
                }
                for fi28 := 0; fi28 < 15; fi28++ {
                    var v29 uint = 65535
                    _ = fi28
                    _ = v29
                }
            }
            var v30 uint32 = 4294967295
            _ = v30
        } else {
            for fi31 := 0; fi31 < 13; fi31++ {
                v32 := []int{fi31}
                var v33 int = fi31
                v34 := GenS6{F0: 0, F1: "foo bar"}
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            if true {
                var v35 float64 = 0.0
                _ = v35
            }
        }
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        fmt.Println(fi36)
        var v37 uint64 = 4294967295
        _ = fi36
        _ = v37
    }
    return true
}

func main() {
    v1 := []int{-872}
    var v2 int64 = 823
    var v3 int16 = 1000
    v4 := make(map[string]string)
    v5 := []string{"hello", ("hello" + "foo bar"), ""}
    v6 := make(map[uint32]int)
    if (!false) {
        var v38 int16 = v3
        _ = v38
    }
    var v39 int16 = (-(v3))
    if (false && true) {
        for fi40 := 0; fi40 < 6; fi40++ {
            v39 += -1
            var v41 int = (fi40 / ((0) | 1))
            v42 := GenS2{F0: 4294967295}
            _ = fi40
            _ = v41
            _ = v42
        }
        v43 := GenS2{F0: 2}
        _ = v43
    }
    switch 9223372036854775807 {
    case 1:
        var v44 int = 2147483647
        _ = v44
    case 466:
        var v45 int = 4294967295
        _ = v45
    case 32768:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        v2 = (v2 % (((v2 | 1)) | 1))
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        if (float32(3.14159) >= 1.5) {
            var v47 uint = 42
            if (!false) {
                var v48 int32 = 1
                _ = v48
            }
            v47 -= (v47 / (((v47 << ((v47) & 63))) | 1))
            _ = v47
        } else {
            if (18446744073709551615 < 18446744073709551615) {
                fmt.Println(v3, v39, v2)
                var v49 int64 = (v2 * (v2 ^ v2))
                _ = v49
            }
            v3 += v39
            switch "hello" {
            case "":
                for fi50 := 0; fi50 < 19; fi50++ {
                    v2 -= v2
                    var v51 float64 = (1e10 + (float64((-2.25 * 3.14159)) * 0.1))
                    var v52 uint64 = 18446744073709551615
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            case "hello":
                if true {
                    var v53 float64 = 1e10
                    v54 := GenS7{F0: 127, F1: 127, F2: 643}
                    _ = v53
                    _ = v54
                } else {
                    v3 -= v3
                    v3 = v39
                }
            }
        }
        _ = fi46
    }
    var v55 float64 = (0.1 / ((3.14159 + 3.14159) + 1e10))
    fmt.Println(v3, v39)
    v3 += (^(v39))
    v56 := GenS8{F0: 42, F1: 1, F2: 32767}
    var v57 uint = 2
    var v58 uint16 = 697
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
