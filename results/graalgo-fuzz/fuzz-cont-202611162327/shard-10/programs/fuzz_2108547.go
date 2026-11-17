package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 int; F1 string }

func main() {
    var v1 bool = false
    var v2 float64 = 1.5
    var v3 float32 = 0.0
    fmt.Println(v3, v3)
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 255, F1: ("foo bar" + "hello")}
            fmt.Println(fi6, v4, v3)
            v8 := []int{}
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 911, F1: "foo bar"}
        _ = fi5
        _ = v9
    }
    v10 := GenS1{F0: v4, F1: "a"}
    if false {
        if (v1 || v1) {
            var v11 int8 = 118
            _ = v11
        }
        v4 -= v4
        var v12 uint32 = 0
        _ = v12
    }
    if (295 > 0) {
        var v13 float64 = v2
        _ = v13
    }
    v14 := GenS2{F0: 4294967295, F1: true, F2: 0}
    switch -1 {
    case -2147483648:
        var v15 uint16 = v4
        _ = v15
    default:
        if ((127 != 100) || false) {
            for fi16 := 0; fi16 < 17; fi16++ {
                v1 = v1
                v2 = v2
                v17 := GenS3{F0: (fi16 >> ((4294967295) & 63)), F1: ("hello" + ("\n\t" + "\n\t"))}
                _ = fi16
                _ = v17
            }
            v18 := GenS1{F0: v4, F1: (("hello" + (("hello" + (("foo bar" + ("a" + "foo bar")) + ("foo bar" + ("a" + ("" + ("hello" + "")))))) + ("foo bar" + "\n\t"))) + "foo bar")}
            if false {
                fmt.Println(v1)
                if ((-139 != 42) && (((v2 < -2.25) && (2 != 1)) || v1)) {
                    if (255 == 0) {
                        v19 := []int{127, 401, -9223372036854775808}
                        _ = v19
                    } else {
                        v20 := []int{-563, 32768}
                        v3 -= v3
                        _ = v20
                    }
                } else {
                    v21 := GenS2{F0: 0, F1: false, F2: 468}
                    v4 = v4
                    v1 = (((("foo bar" + "\n\t") + ("hello" + "foo bar")) + (("foo bar" + ("hello" + "a")) + "\n\t")) <= "\n\t")
                    _ = v21
                }
            } else {
                switch 1 {
                case -59:
                    v22 := GenS3{F0: 0, F1: ""}
                    _ = v22
                case 42:
                    v23 := GenS3{F0: 871, F1: (((("hello" + ("" + ("" + "foo bar"))) + "hello") + "") + ("a" + (("a" + "\n\t") + "hello")))}
                    _ = v23
                case -9223372036854775808:
                    v24 := []int{}
                    _ = v24
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    v26 := []int{(fi25 ^ fi25), (fi25 / ((4294967295) | 1))}
                    var v27 bool = (((0.1 * (-2.25 / (v2 / v2))) - v2) >= float64(v3))
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
            _ = v18
        } else {
            if v1 {
                var v28 int32 = 2147483647
                var v29 uint = 255
                _ = v28
                _ = v29
            } else {
                v2 = float64((((float32(((float64(v3) / (v2 / v2)) - (v2 - (0.0 / ((v2 + v2) + v2))))) - 3.14159) / ((3.14159 * v3) + 1e10)) + 0.1))
            }
            v3 = 0.1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
