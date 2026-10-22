package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint64; F1 int32; F2 uint32 }
func genFn0(p0 int16) bool {
    var v4 uint64 = 4294967295
    for fi5 := 0; fi5 < 12; fi5++ {
        switch "" {
        case "hello":
            switch fi5 {
            case -2147483648:
                v4 -= v4
            case 1:
                p0 = (p0 - int16(v4))
            }
        }
        p0 = p0
        _ = fi5
    }
    _ = p0
    _ = v4
    return true
}

func main() {
    var v1 uint64 = 0
    var v2 float32 = 3.14159
    var v3 float32 = (0.0 / v2)
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 18; fi7++ {
        switch ("foo bar" + ("a" + "a")) {
        case "hello":
            v8 := GenS1{F0: 2147483647}
            _ = v8
        case "":
            switch fi7 {
            case 330:
                var v9 bool = (109 == 127)
                _ = v9
            case 825:
                switch ("hello" + "") {
                case "\n\t":
                    v3 = v3
                default:
                    v1 = (v1 ^ (v1 + (v1 % (((v1 + (v1 % ((4294967295) | 1)))) | 1))))
                }
            default:
                var v10 uint = 65535
                _ = v10
            }
        default:
            if (false && (("hello" + ("foo bar" + "foo bar")) >= ("\n\t" + ((("hello" + "foo bar") + "a") + "\n\t")))) {
                v11 := GenS2{F0: 845, F1: false}
                _ = v11
            }
        }
        if (true || (1 >= 97)) {
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := []int{32767, (fi7 / (((fi7 | fi7)) | 1)), (-(fi12)), 4294967295}
                v6 = 1.5
                if true {
                    v1 -= (v1 & v1)
                    v2 += v2
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                _ = fi12
                _ = v13
            }
        } else {
            v3 = v3
            v14 := []int{9223372036854775807, (fi7 - fi7), (fi7 << ((1) & 63))}
            _ = v14
        }
        fmt.Println(fi7)
        _ = fi7
    }
    v15 := GenS2{F0: 32767, F1: (v6 < 3.14159)}
    v16 := []int{-157, 9, 4294967295}
    v2 = v3
    switch (("\n\t" + "hello") + "hello") {
    case "\n\t":
        v17 := GenS1{F0: -2147483648}
        _ = v17
    case "hello":
        fmt.Println(v6, v3, v1)
    case "":
        v6 = float64((float32(v6) * float32(v6)))
    }
    fmt.Println(v3, v1, v3)
    v18 := GenS3{F0: 4294967295, F1: 0, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v6
}
