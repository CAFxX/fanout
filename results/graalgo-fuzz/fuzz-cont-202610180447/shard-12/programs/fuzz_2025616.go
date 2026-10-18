package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint8; F1 int }
func genFn0() uint8 {
    v7 := []int{913, 32767, 32767, 42}
    v8 := GenS1{F0: 1}
    _ = v7
    _ = v8
    return 155
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 1
    var v3 float32 = 0.1
    var v4 int16 = -1
    var v5 int16 = v4
    var v6 bool = (false || true)
    delete(v1, "a")
    var v9 uint16 = v2
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS2{F0: 2, F1: v5, F2: 1}
        delete(v1, "a")
        fmt.Println(v9, v2)
        _ = fi10
        _ = v11
    }
    fmt.Println(v2, v5)
    switch -2147483648 {
    case 128:
        fmt.Println(v9)
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        if false {
            if v6 {
                for fi13 := 0; fi13 < 19; fi13++ {
                    if true {
                        var v14 int16 = 0
                        v2 -= 255
                        _ = v14
                    }
                    fmt.Println(v2)
                    v15 := GenS2{F0: uint32(v2), F1: (v5 ^ v5), F2: 651}
                    _ = fi13
                    _ = v15
                }
                v4 = (v4 * v4)
            } else {
                if (-35 != -128) {
                    var v16 int32 = -311
                    _ = v16
                } else {
                    _ = len(v1)
                    if true {
                        switch "hello" {
                        case "hello":
                            v17 := GenS3{F0: 629}
                            _ = v17
                        }
                        fmt.Println(v9, v9)
                    }
                    fmt.Println(v3, v2, v5)
                }
                var v18 int64 = 42
                _ = v18
            }
        }
        var v19 float32 = 1.5
        if true {
            v20 := GenS4{F0: 307}
            var v21 float32 = 0.0
            _ = v20
            _ = v21
        }
        _ = fi12
        _ = v19
    }
    v22 := GenS4{F0: 42}
    var v23 bool = false
    switch "\n\t" {
    case "foo bar":
        var v24 int64 = 1
        _ = v24
    case "hello":
        if v23 {
            v25 := GenS1{F0: 85}
            _ = v25
        }
    }
    v26 := GenS3{F0: 18446744073709551615}
    if v6 {
        v3 = 1.5
        fmt.Println(v9, v9, v5)
    } else {
        switch "\n\t" {
        case "":
            v9 = 0
        case "a":
            for fi27 := 0; fi27 < 2; fi27++ {
                if (v23 && (((1.5 - 0.0) >= (1e10 - float64(-2.25))) || v23)) {
                    v28 := GenS3{F0: 2}
                    if true {
                        _ = len(v1)
                        var v29 uint32 = 102
                        var v30 uint16 = 42
                        _ = v29
                        _ = v30
                    }
                    _ = v28
                }
                v23 = (false && v23)
                switch (fi27 * 623) {
                case 2147483647:
                    v5 = v4
                case -2147483648:
                    v31 := GenS2{F0: 42, F1: 665, F2: -2147483648}
                    _ = v31
                case 625:
                    v32 := GenS3{F0: 0}
                    _ = v32
                }
                _ = fi27
            }
        default:
            v9 = 255
        }
    }
    v4 += -802
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 uint16 = (v9 >> ((65535) & 63))
        for fi35 := 0; fi35 < 19; fi35++ {
            v36 := []int{9223372036854775807, (fi35 >> ((305) & 63)), 0}
            if v23 {
                v1[""] = "\n\t"
                if v6 {
                    fmt.Println(v4, v34, v4)
                    v37 := GenS5{F0: 1, F1: 1}
                    v3 += v3
                    _ = v37
                } else {
                    if ((v5 * (v5 & v4)) != 0) {
                        var v38 uint8 = 42
                        _ = v38
                    } else {
                        v2 = (v34 % (((v9 & 2)) | 1))
                        v39 := v1["foo bar"]
                        if len(v36) > 0 {
                            v36[0] = int(v34)
                        }
                        _ = v39
                    }
                    var v40 uint64 = 1
                    _ = v40
                }
                v36 = append(v36, fi33)
            }
            v41 := GenS6{F0: genFn0(), F1: -229}
            _ = fi35
            _ = v36
            _ = v41
        }
        _ = fi33
        _ = v34
    }
    if (!(433 < -647)) {
        var v42 uint32 = 977
        v6 = v6
        v43 := GenS6{F0: 127, F1: -391}
        _ = v42
        _ = v43
    } else {
        v2 = v2
        v44 := []int{}
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
