package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 42
    var v3 bool = (-1 != -2147483648)
    var v4 uint8 = 255
    for fi5 := 0; fi5 < 19; fi5++ {
        switch "" {
        case "\n\t":
            v6 := GenS1{F0: fi5}
            _ = v6
        default:
            for fi7 := 0; fi7 < 15; fi7++ {
                v4 -= (^(v4))
                if (!false) {
                    v3 = false
                    v8 := []int{fi5, fi7}
                    switch (fi7 << ((v2) & 63)) {
                    case 32768:
                        v2 = 1
                    case 389:
                        var v9 uint32 = 1
                        _ = v9
                    default:
                        v3 = v3
                    }
                    _ = v8
                }
                _ = fi7
            }
        }
        if (!((0.1 / (3.14159 + -2.25)) != (3.14159 / (-2.25 + float32(3.14159))))) {
            var v10 bool = false
            _ = v10
        } else {
            var v11 int8 = -1
            delete(v1, "hello")
            _ = v11
        }
        _ = fi5
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        if false {
            fmt.Println(v3, v2)
            var v13 uint8 = (v4 | 127)
            if v3 {
                v14 := GenS2{F0: 0, F1: v13, F2: 95}
                fmt.Println(fi12, fi12)
                for fi15 := 0; fi15 < 19; fi15++ {
                    v4 -= v4
                    v1[(("a" + (("foo bar" + "") + ("a" + "a"))) + "hello")] = fi15
                    var v16 bool = false
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            _ = v13
        } else {
            v17 := GenS3{F0: 4294967295, F1: (^(v4)), F2: 172}
            var v18 bool = (0.0 >= 1.5)
            v19 := []int{(fi12 >> (((v2 % ((v2) | 1))) & 63)), fi12}
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi12
    }
    if v3 {
        v4 = v4
        v2 = 42
    }
    v20 := GenS4{F0: (v4 * (v4 - (v4 & 1)))}
    v21 := GenS4{F0: 2}
    v3 = v3
    v22 := []int{-939, 56}
    fmt.Println(v3)
    v2 += 545
    if (((float32(0.0) - 1.5) + (1e10 / 3.14159)) < 3.14159) {
        fmt.Println(v2, v4, v3)
        if (v3 && v3) {
            var v23 int64 = 946
            var v24 uint32 = 636
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 int32 = 42
                v22 = append(v22, fi25)
                if len(v22) > 0 {
                    v22[0] = int(v24)
                }
                _ = fi25
                _ = v26
            }
            _ = v23
            _ = v24
        } else {
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        }
        if (true && true) {
            if v3 {
                v1["a"] = 42
                v4 = (^(v4))
                v2 = v2
            }
            v1["a"] = int(v4)
        }
    } else {
        v27 := GenS1{F0: 4294967295}
        var v28 string = "\n\t"
        v29 := GenS4{F0: 99}
        _ = v27
        _ = v28
        _ = v29
    }
    switch -9223372036854775808 {
    case 2147483647:
        v3 = (-895 >= 9223372036854775807)
    case 32767:
        switch -1 {
        case 9223372036854775807:
            switch "hello" {
            case "hello":
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            case "foo bar":
                v2 = v2
            case "\n\t":
                var v30 float32 = 3.14159
                _ = v30
            }
        case 2147483647:
            var v31 uint16 = 948
            _ = v31
        case 856:
            var v32 int16 = -416
            _ = v32
        default:
            v4 -= v4
        }
    }
    var v33 uint16 = 65535
    fmt.Println(v4)
    var v34 float64 = 1.5
    switch "" {
    case "foo bar":
        v35 := GenS2{F0: 42, F1: 0, F2: 171}
        _ = v35
    default:
        var v36 float64 = 3.14159
        _ = v36
    }
    if v3 {
        var v37 uint = v2
        _ = v37
    }
    var v38 float64 = v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v38
    _ = v4
}
