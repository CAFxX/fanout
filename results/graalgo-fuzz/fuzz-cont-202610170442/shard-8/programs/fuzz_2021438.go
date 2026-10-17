package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
func genFn0(p0 uint64, p1 int64, p2 int16) uint64 {
    if (true || true) {
        var v6 string = "foo bar"
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    var v9 int16 = p2
    if (255 < 65535) {
        if true {
            v10 := GenS1{F0: (^(p2)), F1: 861}
            var v11 uint8 = 127
            if ((p2 >> ((4294967295) & 63)) > p2) {
                var v12 int64 = p1
                if (v11 > 1) {
                    var v13 int = 0
                    var v14 int16 = (v9 + (p2 & p2))
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 1; fi15++ {
                        var v16 int = (fi15 & fi15)
                        _ = fi15
                        _ = v16
                    }
                }
                v17 := GenS2{F0: 65535, F1: (p1 << ((1) & 63)), F2: 453}
                _ = v12
                _ = v17
            } else {
                fmt.Println(v9, p1)
            }
            _ = v10
            _ = v11
        } else {
            v9 = v9
            var v18 string = ""
            _ = v18
        }
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}
func genFn1(p0 uint16) int16 {
    var v19 uint8 = 152
    if (false && ("\n\t" >= "\n\t")) {
        p0 = 65535
    }
    v19 = (v19 ^ (v19 - v19))
    _ = p0
    _ = v19
    return 32767
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint16 = 2
    var v20 int8 = 1
    v3[0] = (v1 + "hello")
    v1 = ((v1 + "") + "hello")
    switch ("hello" + v1) {
    case "\n\t":
        if ((1.5 - 0.1) < (1.5 + 1.5)) {
            if (!true) {
                v5 += v5
            } else {
                fmt.Println(v2, v20, v1)
            }
        } else {
            var v21 int = -654
            v22 := GenS1{F0: genFn1((v5 - (v5 | v5))), F1: 1}
            var v23 uint = 18446744073709551615
            _ = v21
            _ = v22
            _ = v23
        }
    case "":
        v24 := GenS2{F0: 42, F1: 42, F2: 1}
        _ = v24
    case "hello":
        for fi25 := 0; fi25 < 18; fi25++ {
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 int8 = 0
                var v28 float32 = 0.1
                v1 = ""
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi25
        }
    default:
        switch 4294967295 {
        case 1:
            v29 := []int{-253, 32767, -2147483648}
            _ = v29
        }
    }
    switch v1 {
    case "":
        var v30 bool = (128 < 24)
        _ = v30
    case "hello":
        v5 -= 255
    }
    var v31 uint = 1
    for fi32 := 0; fi32 < 12; fi32++ {
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v33 float32 = 1e10
            _ = v33
        case 9223372036854775807:
            switch fi32 {
            case -2147483648:
                fmt.Println(fi32, v5)
            case 32768:
                var v34 uint32 = 65535
                _ = v34
            }
        default:
            var v35 uint32 = 4294967295
            _ = v35
        }
        _ = fi32
    }
    fmt.Println(v2)
    var v36 float64 = float64(1e10)
    var v37 int16 = 941
    var v38 uint16 = (v5 / (((v5 / ((v5) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
