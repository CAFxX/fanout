package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint32 }
func genFn0(p0 int8) uint8 {
    fmt.Println(p0)
    var v5 int = -2147483648
    v6 := GenS1{F0: 32767, F1: 1, F2: 18446744073709551615}
    v7 := []int{(v5 | v5), (v5 % ((v5) | 1)), (v5 * 127), (v5 ^ (^(v5)))}
    var v8 float32 = (0.1 + (-2.25 * (0.1 - ((-2.25 * -2.25) + (((0.0 * (3.14159 + 0.1)) / 3.14159) / float32(-2.25))))))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 128
}
func genFn1() int {
    v9 := GenS2{F0: 0, F1: 0, F2: 591}
    v10 := []int{}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v11 int8 = 1
    v10 = append(v10, -948)
    _ = v10
    _ = v11
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "a"
    v3 := make(map[string]uint32)
    v4 := []int{128, 0, 32768, 9223372036854775807}
    v12 := GenS2{F0: 255, F1: 155, F2: 32767}
    var v13 float32 = 1e10
    if true {
        fmt.Println(v13, v13)
        v2 = ((v2 + ("foo bar" + "\n\t")) + v2)
    }
    switch (v2 + "foo bar") {
    case "a":
        v14 := GenS2{F0: 32767, F1: 0, F2: 1}
        _ = v14
    case "hello":
        var v15 int32 = 1
        _ = v15
    default:
        if false {
            v2 = (v2 + "a")
            _ = len(v1)
            var v16 bool = (!false)
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            if false {
                if true {
                    v4 = append(v4, fi17)
                }
                if true {
                    v2 = v2
                    v13 = ((v13 + v13) + v13)
                    switch fi17 {
                    case 1:
                        var v19 uint32 = 279
                        _ = v19
                    default:
                        var v20 float32 = 3.14159
                        _ = v20
                    }
                }
                var v21 bool = true
                if (!v21) {
                    switch v2 {
                    case "a":
                        v4 = append(v4, fi18)
                    default:
                        v2 = "\n\t"
                    }
                    v13 = v13
                    v21 = v21
                }
                _ = v21
            }
            switch "\n\t" {
            case "hello":
                v13 = float32(float64(v13))
            case "a":
                switch fi18 {
                case -162:
                    var v22 int = (fi17 / (((fi18 + fi18)) | 1))
                    _ = v22
                case 32768:
                    var v23 float32 = v13
                    _ = v23
                case 2147483647:
                    fmt.Println(v13, fi18)
                }
                _ = len(v1)
            case "foo bar":
                v2 = "foo bar"
            }
            var v24 float32 = (v13 - v13)
            _ = fi18
            _ = v24
        }
        fmt.Println(fi17, fi17)
        _ = fi17
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        if (!true) {
            fmt.Println(fi25, v13, v2)
        }
        v2 = v2
        _ = fi25
    }
    var v26 int16 = 605
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v26)
    case "hello":
        var v27 float64 = ((0.1 / (1e10 + (((3.14159 - 1e10) + (0.1 - -2.25)) + 0.0))) / 1.5)
        _ = v27
    case "":
        v4 = append(v4, genFn1())
    default:
        if (42 >= 1) {
            for fi28 := 0; fi28 < 8; fi28++ {
                v26 += (v26 & v26)
                for fi29 := 0; fi29 < 12; fi29++ {
                    v13 += v13
                    var v30 bool = false
                    v31 := GenS1{F0: 65535, F1: 0, F2: 0}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
            fmt.Println(v13)
            var v32 int32 = 0
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        switch fi33 {
        case -9223372036854775808:
            var v34 uint64 = 42
            _ = v34
        }
        switch v2 {
        case "a":
            v26 += (v26 / (((v26 | (v26 & 882))) | 1))
        case "hello":
            var v35 float64 = 0.0
            _ = v35
        case "\n\t":
            delete(v1, v2)
        default:
            v36 := GenS2{F0: 255, F1: 2, F2: 512}
            _ = v36
        }
        _ = fi33
    }
    var v37 uint32 = 42
    for fi38 := 0; fi38 < 15; fi38++ {
        v2 = ("hello" + (v2 + v2))
        _ = fi38
    }
    var v39 int32 = -957
    switch "" {
    case "a":
        v39 -= (v39 << ((4294967295) & 63))
    case "\n\t":
        v40 := GenS3{F0: 2, F1: 129, F2: (v37 % (((v37 & v37)) | 1))}
        _ = v40
    }
    v41 := GenS3{F0: v37, F1: 32767, F2: v37}
    var v42 int8 = 1
    fmt.Println(v42, v2, v39)
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v41
    _ = v42
}
