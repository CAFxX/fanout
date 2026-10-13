package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 string }
type GenS2 struct { F0 int32; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint8 }
func genFn0(p0 string, p1 int8) uint64 {
    var v5 int32 = -1
    switch 399 {
    case 128:
        p1 += p1
    case -2147483648:
        if (!true) {
            var v6 int = 9223372036854775807
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 int8 = p1
                _ = fi7
                _ = v8
            }
            p0 = "hello"
            _ = v6
        }
    case 4294967295:
        for fi9 := 0; fi9 < 14; fi9++ {
            p1 = (p1 & (p1 | p1))
            _ = fi9
        }
    default:
        var v10 uint8 = 127
        _ = v10
    }
    v5 -= 42
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 int16 = -1
    var v2 float64 = (0.1 + 1.5)
    var v3 uint64 = 2
    var v4 int32 = -1
    var v11 int8 = -83
    v11 -= v11
    v3 = (v3 & v3)
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS1{F0: 755, F1: 119, F2: "hello"}
        v14 := []int{-1, fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 int64 = 506
            _ = fi15
            _ = v16
        }
    } else {
        v2 -= 3.14159
        for fi17 := 0; fi17 < 19; fi17++ {
            v3 = (^(v3))
            v1 = v1
            _ = fi17
        }
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            v3 -= (v3 << ((834) & 63))
        case "\n\t":
            switch -9223372036854775808 {
            case -1:
                v18 := GenS1{F0: 1, F1: v11, F2: (("" + "hello") + (("foo bar" + "foo bar") + ""))}
                _ = v18
            case 9223372036854775807:
                fmt.Println(v4, v2)
            case 128:
                switch -906 {
                case -2147483648:
                    v19 := GenS2{F0: v4, F1: 181, F2: (v3 >> ((42) & 63))}
                    _ = v19
                }
            default:
                v4 -= v4
            }
        case "a":
            v20 := []int{2147483647, 42}
            _ = v20
        }
    }
    v1 = v1
    var v21 string = "a"
    var v22 int8 = -106
    for fi23 := 0; fi23 < 20; fi23++ {
        switch (fi23 % ((438) | 1)) {
        case 281:
            v24 := GenS3{F0: 2, F1: 136}
            _ = v24
        case -9223372036854775808:
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 uint8 = 22
                _ = fi25
                _ = v26
            }
        case -1:
            fmt.Println(v1, v22, v3)
        default:
            switch fi23 {
            case -2147483648:
                v27 := []int{(fi23 ^ int(v1)), fi23, (fi23 & 32768), fi23}
                _ = v27
            case 736:
                var v28 bool = (false && (486 == 42))
                _ = v28
            case -1:
                switch ((("\n\t" + (v21 + ((("a" + v21) + v21) + v21))) + v21) + "\n\t") {
                case "foo bar":
                    var v29 uint8 = 128
                    _ = v29
                case "hello":
                    v30 := GenS1{F0: -533, F1: -14, F2: ("hello" + v21)}
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 8; fi31++ {
                        fmt.Println(fi23, v21, v4)
                        _ = fi31
                    }
                }
            default:
                fmt.Println(v22, v4)
            }
        }
        var v32 uint32 = 2147483648
        _ = fi23
        _ = v32
    }
    v33 := []int{466}
    v3 += v3
    _ = len(v33)
    if ((v1 / ((1000) | 1)) < v1) {
        if false {
            v34 := GenS3{F0: 2, F1: 2}
            _ = v34
        } else {
            switch -1 {
            case 0:
                fmt.Println(v11, v22)
            case 1:
                if len(v33) > 0 {
                    v33[0] = 318
                }
            default:
                v21 = v21
            }
            var v35 string = ((v21 + "foo bar") + "foo bar")
            var v36 bool = false
            _ = v35
            _ = v36
        }
        var v37 int16 = (v1 << ((255) & 63))
        _ = v37
    } else {
        v21 = ("" + v21)
    }
    fmt.Println(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
