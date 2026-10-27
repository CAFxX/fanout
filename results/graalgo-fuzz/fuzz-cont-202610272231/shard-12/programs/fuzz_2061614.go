package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8; F1 string; F2 int }
type GenS3 struct { F0 int8; F1 uint64; F2 int8 }
type GenS4 struct { F0 bool }
func genFn0() int8 {
    var v4 uint = 560
    fmt.Println(v4)
    var v5 uint = 0
    if (true || (((3.14159 * float32(((1.5 * (1.5 * (3.14159 * 1.5))) / 3.14159))) * (1.5 / 0.1)) == ((-2.25 / (1.5 / 1.5)) - (1.5 - float32(float64(1.5)))))) {
        if (true || (2147483647 != -1)) {
            v5 -= (v5 << (((v5 * v5)) & 63))
        }
        var v6 uint32 = 2
        _ = v6
    }
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int) uint8 {
    fmt.Println(p0)
    var v7 float64 = -2.25
    switch ("a" + "a") {
    case "\n\t":
        p0 -= (-(p0))
    case "":
        switch "a" {
        case "hello":
            v7 = v7
        }
    case "hello":
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = 859
            _ = fi8
            _ = v9
        }
    default:
        var v10 bool = true
        _ = v10
    }
    switch "hello" {
    case "a":
        switch (("foo bar" + "foo bar") + "foo bar") {
        case "foo bar":
            v11 := []int{p0, -59, p0}
            _ = v11
        case "hello":
            v12 := []int{}
            _ = v12
        case "a":
            switch (p0 & 42) {
            case 99:
                v13 := []int{(p0 + 32767), -9223372036854775808}
                _ = v13
            case -2147483648:
                p0 = p0
            default:
                p0 = p0
            }
        }
    default:
        fmt.Println(p0, v7)
    }
    var v14 int = (^(p0))
    _ = p0
    _ = v14
    _ = v7
    return 127
}

func main() {
    var v1 uint8 = 2
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    v2 = v2
    v15 := GenS1{F0: 2, F1: 32767, F2: -32768}
    v16 := GenS1{F0: 1, F1: 42, F2: 1000}
    v17 := GenS1{F0: 32767, F1: 2, F2: -32768}
    if true {
        v18 := GenS2{F0: genFn1(127), F1: ((("a" + "\n\t") + "hello") + ""), F2: -829}
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS3{F0: -1, F1: 436, F2: -23}
            var v21 int32 = 0
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v18
    }
    fmt.Println(v2)
    for fi22 := 0; fi22 < 17; fi22++ {
        if false {
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 int = (fi23 - (fi23 >> ((65535) & 63)))
                if (false && true) {
                    var v25 string = "\n\t"
                    v25 = "foo bar"
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
        } else {
            var v26 uint16 = 651
            v27 := []int{}
            if (4294967295 > 697) {
                v28 := GenS1{F0: v26, F1: v26, F2: 0}
                for fi29 := 0; fi29 < 7; fi29++ {
                    v27 = append(v27, fi29)
                    v1 -= 255
                    _ = fi29
                }
                var v30 uint32 = 2
                _ = v28
                _ = v30
            }
            _ = v26
            _ = v27
        }
        _ = fi22
    }
    fmt.Println(v1, v1)
    var v31 int8 = 127
    fmt.Println(v1, v2)
    v32 := GenS4{F0: true}
    switch -615 {
    case 127:
        v3[42] = 65535
    }
    v33 := GenS2{F0: genFn1(4294967295), F1: "a", F2: 2147483647}
    if (18446744073709551615 < 4294967295) {
        v31 = 42
        var v34 uint = 479
        _ = v34
    } else {
        for fi35 := 0; fi35 < 6; fi35++ {
            v1 -= v1
            v2 += -2.25
            for fi36 := 0; fi36 < 14; fi36++ {
                v31 -= (v31 / ((genFn0()) | 1))
                _ = fi36
            }
            _ = fi35
        }
    }
    fmt.Println(v31)
    var v37 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
}
