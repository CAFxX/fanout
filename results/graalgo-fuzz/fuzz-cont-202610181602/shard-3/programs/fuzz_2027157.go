package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint32 = 238
    var v2 int64 = 376
    var v3 float32 = 1e10
    var v4 uint8 = 178
    var v5 string = "foo bar"
    var v6 int16 = -1
    for fi7 := 0; fi7 < 7; fi7++ {
        switch v5 {
        case "hello":
            fmt.Println(v5)
        }
        v8 := GenS1{F0: true, F1: 0, F2: -41}
        v9 := []int{fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 = v5
    if false {
        var v10 uint32 = (v1 & (v1 % (((v1 - v1)) | 1)))
        var v11 float32 = v3
        _ = v10
        _ = v11
    } else {
        if (!(0 == 42)) {
            v12 := GenS1{F0: false, F1: 1, F2: -61}
            v13 := GenS2{F0: -187}
            _ = v12
            _ = v13
        } else {
            v14 := []int{-1, 0, -2147483648, 42}
            switch v5 {
            case "":
                v1 += v1
            }
            var v15 uint32 = v1
            _ = v14
            _ = v15
        }
    }
    var v16 float32 = v3
    for fi17 := 0; fi17 < 14; fi17++ {
        v4 = (v4 & v4)
        _ = fi17
    }
    if ((((!false) && false) || true) || true) {
        var v18 float64 = (float64((v16 * ((0.1 * 0.1) + v16))) / 1.0)
        _ = v18
    }
    var v19 int16 = (v6 ^ -1)
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS1{F0: false, F1: 2, F2: 16}
        if true {
            v16 -= 3.14159
            switch (v5 + ("foo bar" + "\n\t")) {
            case "hello":
                var v22 uint64 = 18446744073709551615
                _ = v22
            default:
                v1 -= v1
            }
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 float64 = ((0.1 + (3.14159 * 3.14159)) + float64(v16))
                var v25 int16 = v19
                if false {
                    switch v5 {
                    case "":
                        fmt.Println(fi23)
                    }
                    v19 += v25
                } else {
                    var v26 string = ("hello" + "hello")
                    fmt.Println(fi23, v26)
                    if true {
                        var v27 uint32 = 919
                        var v28 string = "foo bar"
                        _ = v27
                        _ = v28
                    }
                    _ = v26
                }
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        var v29 uint16 = 65535
        _ = fi20
        _ = v21
        _ = v29
    }
    if ((false || (426 >= 1)) || true) {
        switch 48 {
        case 32767:
            v30 := []int{}
            _ = v30
        case -9223372036854775808:
            v3 -= 3.14159
        }
    }
    var v31 uint64 = 510
    v4 += 127
    switch -1 {
    case 127:
        if ((!(!false)) || ((float32((3.14159 * -2.25)) / (((v16 - (float32((3.14159 / -2.25)) - v3)) / 3.14159) * ((0.1 + ((v16 / -2.25) * ((v16 * -2.25) - (v3 / -2.25)))) - 3.14159))) > v16)) {
            var v32 int8 = -10
            _ = v32
        } else {
            var v33 int32 = 236
            fmt.Println(v31, v3)
            _ = v33
        }
    default:
        v19 -= (v6 % ((-32768) | 1))
    }
    v34 := []int{32768, 1, -2147483648}
    switch -2147483648 {
    case 32768:
        fmt.Println(v5)
    default:
        switch -856 {
        case -97:
            var v35 bool = false
            _ = v35
        }
    }
    if true {
        for fi36 := 0; fi36 < 2; fi36++ {
            switch 42 {
            case -117:
                fmt.Println(v6, v3, v4)
            }
            var v37 uint8 = v4
            v2 = 1
            _ = fi36
            _ = v37
        }
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
