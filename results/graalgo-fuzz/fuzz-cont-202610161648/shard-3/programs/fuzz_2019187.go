package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint64) uint {
    for fi4 := 0; fi4 < 12; fi4++ {
        p0 += p0
        v5 := []int{fi4}
        p0 -= p0
        _ = fi4
        _ = v5
    }
    switch "" {
    case "\n\t":
        var v6 uint = 0
        _ = v6
    default:
        switch 32767 {
        case 4294967295:
            var v7 uint = 255
            _ = v7
        case 341:
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int = 42
                v10 := []int{}
                _ = fi8
                _ = v10
                _ = v9
            }
        case 2147483647:
            p0 = 802
        default:
            var v11 int = 670
            _ = v11
        }
    }
    var v12 uint8 = 2
    _ = p0
    _ = v12
    return 2
}

func main() {
    var v1 float64 = (0.1 / (float64(float32(((1e10 / 1.5) / 1.0))) - 0.1))
    v2 := make(map[uint32]string)
    var v3 uint16 = 287
    v13 := GenS1{F0: -9223372036854775808, F1: 106, F2: 1}
    for fi14 := 0; fi14 < 5; fi14++ {
        if true {
            var v15 uint32 = uint32(v3)
            v16 := GenS2{F0: 730, F1: -992, F2: -1}
            fmt.Println(fi14, v1)
            _ = v15
            _ = v16
        } else {
            v2[977] = "\n\t"
            v17 := GenS2{F0: 1, F1: -967, F2: 0}
            var v18 uint = 0
            _ = v17
            _ = v18
        }
        var v19 uint8 = 128
        _ = fi14
        _ = v19
    }
    var v20 int64 = 1
    v1 = (0.1 - 3.14159)
    v1 = (((((v1 - -2.25) / 1.0) + v1) - float64(3.14159)) - 3.14159)
    if (42 >= 0) {
        if (("a" + "hello") != (("a" + "a") + ("foo bar" + "hello"))) {
            v3 -= 2
            switch 448 {
            case 128:
                switch "" {
                case "a":
                    v21 := v2[2147483647]
                    _ = v21
                default:
                    delete(v2, 0)
                }
            case 446:
                v22 := v2[0]
                _ = v22
            case -2147483648:
                v3 = 208
            default:
                switch -9223372036854775808 {
                case 32767:
                    var v23 uint32 = 65535
                    _ = v23
                case 26:
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v1, v3)
                    case "\n\t":
                        v3 = v3
                    }
                }
            }
        }
        if ((!false) && ("foo bar" <= "\n\t")) {
            var v24 int16 = 457
            var v25 string = ("\n\t" + ("foo bar" + ""))
            _ = v24
            _ = v25
        }
        if (3.14159 == ((float64(((0.0 - ((0.1 + 1.5) - 0.0)) / 0.1)) / 0.1) * v1)) {
            var v26 int16 = 720
            fmt.Println(v3, v20, v26)
            switch 505 {
            case 0:
                fmt.Println(v20, v1, v3)
            }
            _ = v26
        } else {
            var v27 int16 = -32768
            v1 = 0.1
            for fi28 := 0; fi28 < 10; fi28++ {
                fmt.Println(fi28)
                _ = fi28
            }
            _ = v27
        }
    }
    delete(v2, 646)
    fmt.Println(v1, v3)
    var v29 uint32 = 65535
    var v30 int16 = -1
    v3 = (v3 ^ uint16(v29))
    switch "foo bar" {
    case "a":
        for fi31 := 0; fi31 < 19; fi31++ {
            if true {
                var v32 uint8 = 128
                _ = v32
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                v34 := []int{}
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
                _ = fi33
                _ = v34
            }
            v2[0] = (("\n\t" + "a") + "\n\t")
            _ = fi31
        }
    case "\n\t":
        var v35 uint16 = 255
        _ = v35
    case "foo bar":
        v36 := []int{2147483647, 32768, 4294967295}
        _ = v36
    default:
        v37 := GenS3{F0: 60, F1: v29}
        _ = v37
    }
    v38 := []int{}
    v39 := GenS3{F0: 7, F1: v29}
    var v40 string = "a"
    if len(v38) >= 2 {
        v38 = v38[1:len(v38)]
    }
    v30 -= (v30 * v30)
    v29 += v29
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v40
}
