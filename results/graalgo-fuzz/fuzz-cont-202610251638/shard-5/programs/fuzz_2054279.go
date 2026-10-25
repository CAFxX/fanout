package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    v1 := []int{42, -9223372036854775808, 128, 42}
    v2 := []int{-2147483648}
    v3 := []string{"", ("hello" + ""), "hello", "\n\t"}
    var v4 uint8 = 242
    var v5 int16 = 42
    if true {
        var v6 int8 = 127
        for fi7 := 0; fi7 < 19; fi7++ {
            v5 = v5
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            _ = fi7
        }
        _ = v6
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 bool = (42 >= 136)
        var v10 int16 = 42
        fmt.Println(fi8, v5)
        _ = fi8
        _ = v10
        _ = v9
    }
    v5 = -1
    var v11 float32 = (((1e10 * -2.25) * 0.0) + 0.0)
    v12 := GenS1{F0: 1, F1: 1, F2: (1e10 <= 0.0)}
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 float32 = float32(1.5)
        var v15 int8 = 127
        v2 = append(v2, (fi13 >> ((42) & 63)))
        _ = fi13
        _ = v14
        _ = v15
    }
    if (3.14159 <= float64(float32(3.14159))) {
        if true {
            var v16 float64 = float64(0.1)
            _ = v16
        } else {
            if true {
                v17 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: false}
                var v18 int = 127
                _ = v17
                _ = v18
            }
            var v19 bool = false
            v11 = float32(3.14159)
            _ = v19
        }
    }
    var v20 string = ("a" + "foo bar")
    if (!(!false)) {
        for fi21 := 0; fi21 < 8; fi21++ {
            switch v20 {
            case "foo bar":
                var v22 float64 = 0.1
                _ = v22
            case "":
                fmt.Println(v4)
            }
            v5 = v5
            v5 = (v5 % (((v5 + -565)) | 1))
            _ = fi21
        }
        v20 = ("foo bar" + v20)
    } else {
        v23 := GenS1{F0: -1, F1: 783, F2: (!(42 == 18446744073709551615))}
        _ = v23
    }
    v5 = 42
    var v24 int64 = 42
    var v25 int8 = 1
    var v26 float32 = 0.1
    if (float64(float32(3.14159)) != (0.0 + 0.1)) {
        var v27 int = -9223372036854775808
        if len(v3) > 0 {
            v3[0] = (v20 + v20)
        }
        if true {
            v11 += v11
        }
        _ = v27
    }
    if false {
        v28 := GenS2{F0: -2147483648, F1: true}
        _ = v28
    } else {
        if false {
            fmt.Println(v5, v25, v20)
            switch v20 {
            case "foo bar":
                v29 := GenS2{F0: -986, F1: true}
                _ = v29
            case "a":
                switch (v20 + v20) {
                case "\n\t":
                    v11 = ((v11 + (0.1 - (float32(((1e10 * 1e10) - 1.5)) - 1e10))) + v26)
                case "":
                    fmt.Println(v4, v4, v24)
                case "foo bar":
                    var v30 uint = 4294967295
                    _ = v30
                }
            }
        } else {
            if len(v2) > 0 {
                v2[0] = 0
            }
            var v31 int8 = (v25 << ((568) & 63))
            _ = v31
        }
        _ = len(v1)
    }
    var v32 int = 9223372036854775807
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
