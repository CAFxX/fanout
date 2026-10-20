package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]int)
    var v3 int32 = -629
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    switch 4294967295 {
    case 0:
        v6 := GenS1{F0: -2147483648}
        _ = v6
    case -122:
        fmt.Println(v1)
    default:
        var v7 uint = 9
        _ = v7
    }
    v1 = v1
    var v8 bool = true
    var v9 int32 = -2147483648
    switch ("foo bar" + "a") {
    case "foo bar":
        var v10 int = 538
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 4; fi11++ {
            switch -450 {
            case 127:
                var v12 int32 = (-(v3))
                _ = v12
            }
            fmt.Println(v1, v1, v8)
            v13 := v4[-3]
            _ = fi11
            _ = v13
        }
    default:
        switch 32767 {
        case 32768:
            for fi14 := 0; fi14 < 12; fi14++ {
                fmt.Println(v9, v8, v3)
                if v8 {
                    v1 = 0
                    v9 = v3
                }
                v15 := GenS1{F0: v3}
                _ = fi14
                _ = v15
            }
        }
    }
    if (952 < 1) {
        v1 = (^(v1))
        fmt.Println(v9, v1, v1)
        v16 := GenS1{F0: v9}
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS1{F0: (v9 >> ((42) & 63))}
        v19 := []int{}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: (v9 % ((v3) | 1))}
    var v21 int = -1
    v22 := GenS2{F0: 42, F1: 9223372036854775807}
    v21 = (-(v21))
    v23 := v5[v21]
    if (!(-316 < 731)) {
        if true {
            v24 := GenS2{F0: 18446744073709551615, F1: 9223372036854775807}
            var v25 string = "foo bar"
            var v26 uint64 = 0
            _ = v24
            _ = v25
            _ = v26
        }
        v27 := []int{-698}
        _ = v27
    } else {
        v8 = v8
        v21 = (v21 % (((v21 * 9223372036854775807)) | 1))
        if v8 {
            var v28 int = v21
            _ = v28
        } else {
            for fi29 := 0; fi29 < 20; fi29++ {
                for fi30 := 0; fi30 < 18; fi30++ {
                    var v31 float32 = 3.14159
                    switch v23 {
                    case "":
                        switch 42 {
                        case 32768:
                            v3 += (-(v3))
                        case 127:
                            v32 := GenS2{F0: 42, F1: -9223372036854775808}
                            _ = v32
                        case 4294967295:
                            var v33 uint16 = 255
                            _ = v33
                        default:
                            v34 := GenS3{F0: 2, F1: 0}
                            _ = v34
                        }
                    default:
                        var v35 uint = uint(v1)
                        _ = v35
                    }
                    _ = fi30
                    _ = v31
                }
                for fi36 := 0; fi36 < 12; fi36++ {
                    var v37 int = (v21 ^ v21)
                    _ = fi36
                    _ = v37
                }
                _ = fi29
            }
            v1 -= 65535
            v38 := []int{(v21 | (v21 >> ((0) & 63)))}
            _ = v38
        }
    }
    var v39 int32 = v9
    var v40 int32 = v39
    var v41 int = (-(v21))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v8
    _ = v9
}
