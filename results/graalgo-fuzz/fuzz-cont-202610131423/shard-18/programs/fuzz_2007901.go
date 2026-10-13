package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    v3 := make(map[string]string)
    var v4 int = -2147483648
    if false {
        var v5 uint16 = (v2 - v2)
        var v6 int = (v4 * v4)
        _ = v5
        _ = v6
    }
    v4 += (v4 * v4)
    v4 -= (v4 | (v4 - v4))
    var v7 uint64 = 1
    v8 := GenS1{F0: 42, F1: 4294967295}
    var v9 int8 = 1
    v10 := GenS2{F0: 126, F1: v4, F2: 362}
    var v11 int = int(v7)
    if (673 < 1000) {
        var v12 string = ("\n\t" + ("\n\t" + "a"))
        var v13 float64 = float64((-2.25 * (-2.25 + ((float32(-2.25) - 1.5) / (((1e10 - (0.0 + 0.1)) / 1e10) + ((3.14159 / 1.5) + 1e10))))))
        for fi14 := 0; fi14 < 2; fi14++ {
            if true {
                fmt.Println(v7, v4, v11)
                fmt.Println(v7, v4)
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS3{F0: v9, F1: -57}
                    v2 = v2
                    _ = fi15
                    _ = v16
                }
            } else {
                var v17 int64 = 0
                _ = v17
            }
            _ = fi14
        }
        _ = v12
        _ = v13
    } else {
        if (-1 <= 90) {
            var v18 bool = false
            _ = v18
        } else {
            fmt.Println(v11)
            var v19 string = ((("" + (("foo bar" + ("foo bar" + (((("foo bar" + ("hello" + ((("a" + "a") + (("a" + "hello") + "\n\t")) + "foo bar"))) + (("\n\t" + "hello") + (("hello" + "hello") + "\n\t"))) + ("a" + ((("a" + (("a" + "foo bar") + "hello")) + ("\n\t" + "foo bar")) + "\n\t"))) + ("hello" + "hello")))) + "a")) + ("hello" + (("" + ("a" + "a")) + ("\n\t" + "")))) + "\n\t")
            _ = v19
        }
    }
    switch "\n\t" {
    case "":
        v20 := []int{v11}
        _ = v20
    case "hello":
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS2{F0: 18446744073709551615, F1: (v11 ^ 1), F2: 42}
            v23 := []int{(-(v4)), (v11 | int(v7)), (fi21 << ((65535) & 63)), int(v2)}
            v2 += (^(v2))
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    switch v11 {
    case 0:
        var v24 uint64 = (v7 / ((uint64(v2)) | 1))
        _ = v24
    case 9223372036854775807:
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 bool = (!(1 > -519))
            var v27 int8 = v9
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v28 := GenS4{F0: v7}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
