package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int64; F1 uint16; F2 uint64 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint8 = 127
    v3 := make(map[uint32]bool)
    var v4 string = "a"
    v5 := make(map[int]int)
    if false {
        v2 += 243
        fmt.Println(v4, v2, v4)
    } else {
        v6 := GenS1{F0: 65535, F1: 4294967295}
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        switch fi7 {
        case -143:
            _ = len(v3)
        case -9223372036854775808:
            var v8 bool = (1 >= -1)
            _ = v8
        case -446:
            fmt.Println(v4, v1)
        default:
            v1 = ((((0.1 + v1) + v1) - 0.0) * (v1 / -2.25))
        }
        v4 = (v4 + ("" + "a"))
        v1 -= 0.0
        _ = fi7
    }
    var v9 string = ("a" + (v4 + ((v4 + v4) + v4)))
    var v10 bool = true
    fmt.Println(v1)
    if false {
        for fi11 := 0; fi11 < 17; fi11++ {
            if v10 {
                v4 = (("a" + v9) + v4)
            }
            if ((1 < (^(v2))) || (830 >= 2147483647)) {
                for fi12 := 0; fi12 < 17; fi12++ {
                    fmt.Println(v4, fi12, fi11)
                    _ = len(v3)
                    v13 := GenS2{F0: v2, F1: 1}
                    _ = fi12
                    _ = v13
                }
            }
            _ = fi11
        }
        switch ((("a" + (((v9 + "") + v9) + v4)) + (("a" + (v4 + v4)) + v9)) + ("hello" + v9)) {
        case "a":
            var v14 string = (v9 + (v9 + v9))
            _ = v14
        case "":
            v15 := GenS3{F0: int64(v2), F1: 32767, F2: 723}
            _ = v15
        case "hello":
            switch "" {
            case "foo bar":
                for fi16 := 0; fi16 < 4; fi16++ {
                    fmt.Println(v4, fi16)
                    _ = fi16
                }
            default:
                v1 = v1
            }
        }
        fmt.Println(v1, v10, v9)
    }
    switch 127 {
    case 2147483647:
        v17 := v3[1]
        _ = v17
    case 32767:
        fmt.Println(v2, v2, v1)
    }
    var v18 uint = 1
    v18 += 18446744073709551615
    switch -2147483648 {
    case 4294967295:
        if ((v18 % (((v18 & 1)) | 1)) != v18) {
            fmt.Println(v10, v18)
        }
    case 127:
        _ = len(v3)
    }
    if (0.0 < (float64(v1) + ((1e10 + (float64(v1) + (-2.25 * 0.1))) * (-2.25 * 0.0)))) {
        for fi19 := 0; fi19 < 20; fi19++ {
            switch (fi19 / (((fi19 << ((v18) & 63))) | 1)) {
            case 2147483647:
                if v10 {
                    fmt.Println(v9)
                    v2 -= v2
                    _ = len(v5)
                }
            }
            v10 = v10
            _ = fi19
        }
        if (true && v10) {
            for fi20 := 0; fi20 < 3; fi20++ {
                switch (v4 + (v9 + v4)) {
                case "":
                    v21 := v5[(fi20 | fi20)]
                    _ = v21
                case "hello":
                    v1 += ((((-2.25 - v1) / 1.5) * float32((-2.25 / (0.1 * float64(v1))))) + v1)
                default:
                    var v22 uint16 = 32767
                    _ = v22
                }
                v4 = (v4 + "foo bar")
                if v10 {
                    delete(v5, 1)
                } else {
                    if (true && (("hello" + "hello") < "hello")) {
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v1 -= 0.1
                            _ = fi23
                        }
                    }
                }
                _ = fi20
            }
            var v24 float64 = ((float64(float32(((1e10 / 3.14159) - 0.0))) * ((float64(v1) - 3.14159) * -2.25)) / 1.5)
            _ = v24
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 uint16 = 2
                if (false || (v18 != (v18 & (v18 * v18)))) {
                    v27 := []int{(^(fi25)), fi25, fi25}
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
            delete(v3, 243)
        }
    } else {
        var v28 uint16 = 1
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
