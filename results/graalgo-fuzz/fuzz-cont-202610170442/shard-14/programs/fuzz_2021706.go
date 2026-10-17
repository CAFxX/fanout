package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    var v4 int64 = -659
    var v5 string = "hello"
    v6 := []int{-1, 416, -906}
    if ((!true) && false) {
        switch v5 {
        case "\n\t":
            v7 := GenS1{F0: 283}
            _ = v7
        }
    }
    if (255 > 255) {
        var v8 uint16 = 279
        var v9 bool = (false || true)
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int = (^(fi10))
        if false {
            fmt.Println(v3)
            if ((127 <= 255) && true) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 uint64 = 4294967295
                    if (100 == 100) {
                        v6 = append(v6, fi10)
                        var v14 int8 = 100
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v14
                    } else {
                        v3 = 18446744073709551615
                    }
                    var v15 int = fi10
                    _ = fi12
                    _ = v13
                    _ = v15
                }
            }
            if true {
                v16 := GenS1{F0: 2147483647}
                var v17 int8 = 127
                _ = v16
                _ = v17
            }
        }
        if true {
            v5 = ""
            v5 = (v5 + "a")
        }
        _ = fi10
        _ = v11
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            switch ("\n\t" + "a") {
            case "\n\t":
                if false {
                    switch (fi18 / ((fi19) | 1)) {
                    case 1:
                        if (!(v3 < 18446744073709551615)) {
                            v20 := GenS2{F0: 887, F1: 52, F2: fi18}
                            var v21 uint = 42
                            v22 := GenS1{F0: 917}
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS1{F0: -709}
                            _ = len(v2)
                            _ = v23
                        }
                    case -269:
                        var v24 int32 = 2147483647
                        _ = v24
                    }
                    v6 = append(v6, -9223372036854775808)
                }
                v25 := GenS2{F0: 2, F1: -128, F2: fi19}
                _ = v25
            case "foo bar":
                v3 = v3
            case "":
                v3 += v3
                var v26 int = fi18
                _ = v26
            default:
                v27 := GenS1{F0: 42}
                _ = v27
            }
            if true {
                switch (fi18 >> ((153) & 63)) {
                case 946:
                    var v28 float32 = 1.5
                    _ = v28
                case 0:
                    switch v5 {
                    case "hello":
                        var v29 uint32 = 65535
                        _ = v29
                    case "a":
                        if false {
                            var v30 int64 = v4
                            _ = len(v6)
                            var v31 string = "a"
                            _ = v30
                            _ = v31
                        } else {
                            v32 := GenS1{F0: -1}
                            fmt.Println(v4)
                            _ = v32
                        }
                    case "":
                        v33 := GenS3{F0: 1}
                        _ = v33
                    }
                case 853:
                    v5 = v5
                }
                v5 = "a"
                v3 -= 464
                if (-32768 == 1) {
                    v1[160] = "\n\t"
                    v4 -= v4
                    v34 := v1[65535]
                    v35 := GenS4{F0: -128, F1: -1}
                    _ = v34
                    _ = v35
                }
            }
            v4 = v4
            var v36 float64 = float64(float32(1.5))
            _ = fi19
            _ = v36
        }
        switch (fi18 >> ((1) & 63)) {
        case 2147483647:
            v5 = (("" + v5) + ((v5 + ("a" + v5)) + (((v5 + v5) + "a") + v5)))
        case 32768:
            v5 = v5
        default:
            v3 += 719
        }
        _ = fi18
    }
    v4 = (v4 ^ -240)
    var v37 uint32 = 42
    var v38 float32 = 3.14159
    v39 := v1[v37]
    var v40 uint16 = 32767
    var v41 uint64 = 0
    if (((float64(v38) + 3.14159) / float64((-2.25 + (v38 * 3.14159)))) > ((3.14159 * ((1e10 + 0.0) * (1e10 / 1e10))) - (((0.0 * float64(((1e10 - (v38 * v38)) + 0.0))) / float64(v38)) * 0.0))) {
        switch "\n\t" {
        case "\n\t":
            v4 = (v4 / (((v4 % (((v4 + 853)) | 1))) | 1))
        case "hello":
            fmt.Println(v3)
        case "":
            v42 := GenS3{F0: -1}
            _ = v42
        }
        var v43 int = 484
        if true {
            v3 = v41
            v4 -= -650
            v43 = (v43 / (((v43 << ((42) & 63))) | 1))
        }
        _ = v43
    } else {
        fmt.Println(v5, v5)
        var v44 uint16 = (v40 | 65535)
        v45 := GenS4{F0: 92, F1: 2147483647}
        _ = v44
        _ = v45
    }
    switch 32768 {
    case 4294967295:
        v2[(v37 - 1)] = true
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
