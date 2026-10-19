package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint8, p2 int16) int8 {
    switch "a" {
    case "hello":
        p1 = p1
    }
    var v4 uint8 = 66
    p2 -= -32768
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            p0 += (^(p0))
            _ = fi6
        }
        v7 := []int{(fi5 / (((fi5 % (((fi5 / ((fi5) | 1))) | 1))) | 1)), 42, -9223372036854775808, fi5}
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 float64 = (((((3.14159 / 3.14159) + 3.14159) + -2.25) + 1.5) + 0.1)
            _ = fi8
            _ = v9
        }
        _ = fi5
        _ = v7
    }
    if false {
        v10 := []int{}
        _ = v10
    } else {
        p0 = (-(p0))
        v11 := GenS1{F0: (-(p0))}
        p0 = (-(p0))
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 100
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 uint32 = 2147483647
    v12 := GenS2{F0: 2147483647, F1: 65535, F2: -1}
    v3 = v3
    v3 = (v3 ^ v3)
    var v13 uint64 = 4294967295
    if false {
        v14 := GenS3{F0: -2147483648, F1: 42, F2: 891}
        if (!(!true)) {
            for fi15 := 0; fi15 < 16; fi15++ {
                for fi16 := 0; fi16 < 3; fi16++ {
                    _ = len(v1)
                    v17 := []int{fi15}
                    _ = fi16
                    _ = v17
                }
                v18 := GenS4{F0: (v3 & v3), F1: 2}
                _ = fi15
                _ = v18
            }
            v19 := GenS3{F0: -403, F1: (v13 * (^(v13))), F2: 0}
            v1[v3] = (("\n\t" + ("a" + ("\n\t" + "\n\t"))) + ("foo bar" + ((("" + ("hello" + (("a" + "hello") + "foo bar"))) + ("\n\t" + "hello")) + "hello")))
            _ = v19
        }
        var v20 float64 = 3.14159
        _ = v14
        _ = v20
    } else {
        var v21 uint16 = 228
        fmt.Println(v3, v13)
        if ("" > ("hello" + "\n\t")) {
            var v22 uint32 = v3
            _ = v22
        } else {
            if false {
                v13 = v13
                fmt.Println(v13)
            } else {
                for fi23 := 0; fi23 < 19; fi23++ {
                    switch (fi23 % ((-568) | 1)) {
                    case 0:
                        v13 -= v13
                    case 32768:
                        fmt.Println(v3, v13)
                    default:
                        var v24 float32 = 0.1
                        _ = v24
                    }
                    _ = fi23
                }
                fmt.Println(v21, v3, v21)
            }
            v25 := []int{32767, -781, 32768}
            var v26 int32 = -1
            _ = v25
            _ = v26
        }
        _ = v21
    }
    var v27 int = -711
    switch ("a" + "") {
    case "hello":
        _ = len(v2)
    case "foo bar":
        switch "hello" {
        case "foo bar":
            var v28 uint64 = (v13 << ((18446744073709551615) & 63))
            _ = v28
        case "":
            var v29 float64 = (float64((0.1 / 0.1)) / float64((float32(((((3.14159 - (float64(0.1) - (0.1 / 1.0))) + 1e10) - float64(1e10)) - (((3.14159 + -2.25) * (0.1 / float64(-2.25))) * -2.25))) / ((1e10 * 0.0) - float32(-2.25)))))
            _ = v29
        default:
            if true {
                var v30 int8 = 42
                _ = v30
            } else {
                for fi31 := 0; fi31 < 17; fi31++ {
                    v27 += 32768
                    var v32 uint64 = 1
                    _ = fi31
                    _ = v32
                }
                if true {
                    var v33 int8 = 100
                    _ = v33
                }
            }
        }
    }
    var v34 float64 = float64(((3.14159 * 0.0) / (0.1 * 1.5)))
    var v35 uint32 = v3
    var v36 int64 = 9223372036854775807
    delete(v1, v3)
    var v37 uint32 = 2147483647
    v38 := GenS5{F0: 128}
    v13 = (v13 >> ((930) & 63))
    fmt.Println(v35, v34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
