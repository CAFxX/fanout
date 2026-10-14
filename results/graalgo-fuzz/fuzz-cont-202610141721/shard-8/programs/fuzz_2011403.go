package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8; F1 int; F2 int8 }

func main() {
    v1 := []string{("" + "hello"), "foo bar", "hello", ("hello" + "")}
    var v2 int16 = -1
    var v3 uint32 = 42
    v4 := []uint32{(v3 / (((v3 * v3)) | 1))}
    var v5 int8 = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 = append(v4, 1)
        _ = fi6
    }
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: v2, F1: v2, F2: 328}
    for fi9 := 0; fi9 < 14; fi9++ {
        v7 = (v7 % ((-1) | 1))
        if true {
            v2 -= (v2 >> ((uint(v3)) & 63))
            var v10 int32 = 1
            var v11 uint = 65535
            _ = v10
            _ = v11
        } else {
            var v12 int64 = (v7 >> ((42) & 63))
            fmt.Println(v2, v12)
            _ = v12
        }
        switch "a" {
        case "\n\t":
            v7 = v7
        }
        _ = fi9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v13 := GenS1{F0: v2, F1: (^(v2)), F2: 255}
    _ = len(v4)
    if (float32(0.0) >= (0.0 - (0.1 * 1e10))) {
        _ = len(v1)
        var v14 uint32 = (v3 - v3)
        if len(v4) > 0 {
            v4[0] = (v14 / (((v14 & (^(v3)))) | 1))
        }
        _ = v14
    } else {
        switch "\n\t" {
        case "a":
            if (!(v5 >= (v5 ^ 0))) {
                var v15 bool = false
                var v16 int32 = 42
                _ = v15
                _ = v16
            }
        }
        var v17 bool = (94 <= 1)
        var v18 string = "hello"
        _ = v17
        _ = v18
    }
    var v19 float32 = (0.0 + 0.0)
    var v20 uint16 = 65535
    if ((255 <= 0) || ("foo bar" <= (("\n\t" + "foo bar") + "foo bar"))) {
        if (!true) {
            _ = len(v1)
            var v21 int8 = 83
            if (0.0 > float32(float64(float32((((3.14159 + 1e10) * (float64(v19) - ((((1e10 + (1.5 * ((((0.0 + (0.0 / 1.0)) * 0.1) + 1.5) - ((float64((v19 + 1.5)) * float64(((v19 * v19) * (-2.25 + (float32((((0.0 + 3.14159) * 0.0) / 3.14159)) + ((v19 * (float32(float64((v19 - v19))) - (v19 * 0.0))) * ((-2.25 * (v19 + (v19 * v19))) - (1e10 + v19)))))))) - (float64(float32(((0.1 - 0.0) / (0.0 / (1e10 + (0.1 * float64((v19 * v19)))))))) + 3.14159))))) * (1e10 - 0.1)) / 1.5) * (3.14159 - 0.1)))) * 0.1))))) {
                var v22 int8 = (v5 | v5)
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 int16 = (v2 << ((4294967295) & 63))
                    _ = fi23
                    _ = v24
                }
                v1 = append(v1, "\n\t")
                _ = v22
            }
            _ = v21
        }
        var v25 bool = false
        _ = v25
    } else {
        var v26 int32 = -1
        switch (("" + ((("a" + "\n\t") + "foo bar") + "foo bar")) + "foo bar") {
        case "hello":
            v2 -= 0
        case "\n\t":
            v27 := GenS2{F0: 4294967295, F1: v26}
            _ = v27
        }
        _ = v26
    }
    v28 := GenS1{F0: 32767, F1: v2, F2: 381}
    if false {
        var v29 float64 = 1e10
        _ = len(v1)
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 uint16 = v20
            var v32 int64 = (v7 | -1)
            v29 = v29
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    } else {
        _ = len(v1)
        if len(v4) > 0 {
            v4[0] = 42
        }
    }
    fmt.Println(v2)
    switch -9223372036854775808 {
    case 128:
        var v33 int32 = -1
        _ = v33
    }
    if (32767 != -98) {
        if true {
            fmt.Println(v19, v20)
            switch "hello" {
            case "foo bar":
                v1 = append(v1, "a")
            case "\n\t":
                for fi34 := 0; fi34 < 16; fi34++ {
                    v35 := GenS3{F0: v5, F1: fi34, F2: (v5 - int8(v20))}
                    v2 = 294
                    _ = fi34
                    _ = v35
                }
            }
        }
        v1 = append(v1, "a")
        v1 = append(v1, ((("hello" + "foo bar") + "hello") + "hello"))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
