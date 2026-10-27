package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 int) int8 {
    var v4 int64 = 990
    v5 := GenS1{F0: 0}
    var v6 float32 = (float32(-2.25) - -2.25)
    switch ("hello" + ("\n\t" + "a")) {
    case "foo bar":
        if (128 >= 255) {
            var v7 int64 = v4
            _ = v7
        }
    case "a":
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: -1}
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := []int{-941, 4294967295, 42, fi10}
                v4 = -9223372036854775808
                v11 = append(v11, p0)
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int64, p1 uint64) uint16 {
    p1 = (p1 * 686)
    var v12 uint = 18446744073709551615
    var v13 uint32 = 2147483647
    var v14 int8 = -36
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return 494
}

func main() {
    var v1 float32 = (float32((0.1 - 0.1)) / -2.25)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v15 uint16 = 65535
    for fi16 := 0; fi16 < 13; fi16++ {
        delete(v3, 65535)
        _ = fi16
    }
    var v17 int = 280
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int32 = -1
            var v21 uint16 = 463
            var v22 string = ""
            v23 := []int{(fi19 / ((fi19) | 1)), 374, 863}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
            _ = v23
        }
        v24 := GenS2{F0: 1, F1: -285}
        switch ("a" + ("\n\t" + "")) {
        case "a":
            v1 = 1e10
        case "\n\t":
            var v25 int8 = 1
            _ = v25
        default:
            _ = len(v3)
        }
        _ = fi18
        _ = v24
    }
    fmt.Println(v15, v1, v17)
    v26 := []int{v17, v17, (v17 % ((-829) | 1)), 4294967295}
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    if true {
        var v27 float64 = (1e10 - float64(0.0))
        if true {
            v15 -= 42
            v28 := GenS2{F0: 684, F1: -2147483648}
            _ = v28
        } else {
            v17 += (v17 & v17)
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 float32 = v1
                v31 := GenS3{F0: ((((v27 * v27) * 1.5) >= (v27 / ((v27 / 1e10) * (v27 * (float64(v1) - (v27 + (float64((float32(v27) + v1)) - float64((v30 / (((0.0 + v30) * float32((float64(((v1 + (v1 * ((1e10 + (0.1 + 0.0)) / (v1 - 0.0)))) * 1.5)) + v27))) / v1)))))))))) && true)}
                _ = fi29
                _ = v30
                _ = v31
            }
            switch "" {
            case "hello":
                if true {
                    v17 += 32768
                }
            case "foo bar":
                for fi32 := 0; fi32 < 10; fi32++ {
                    v15 = v15
                    _ = fi32
                }
            case "":
                v27 = v27
            default:
                var v33 float32 = (3.14159 * float32(-2.25))
                _ = v33
            }
        }
        var v34 float32 = 0.1
        _ = v27
        _ = v34
    }
    if ((!true) || false) {
        var v35 int32 = -1
        switch "a" {
        case "":
            for fi36 := 0; fi36 < 10; fi36++ {
                var v37 int8 = 118
                _ = fi36
                _ = v37
            }
        case "hello":
            _ = len(v26)
        case "a":
            if false {
                var v38 int16 = 42
                _ = v38
            }
        }
        _ = v35
    }
    if true {
        var v39 int64 = 0
        v26 = append(v26, (v17 ^ v17))
        var v40 int64 = (-(v39))
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v26
    _ = v3
}
