package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int16; F2 int32 }
type GenS7 struct { F0 uint32; F1 uint; F2 bool }
func genFn0(p0 int8, p1 int32, p2 uint) int8 {
    v4 := []int{2147483647, 32767}
    var v5 int8 = 93
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int8, p1 int) uint32 {
    v6 := GenS1{F0: -32768}
    switch (p1 << ((65535) & 63)) {
    case -571:
        p0 -= 1
    case -1:
        v7 := GenS2{F0: 372, F1: 2, F2: 65535}
        _ = v7
    case 42:
        var v8 int32 = 1
        _ = v8
    default:
        p1 = (p1 - -426)
    }
    v9 := []int{(p1 / ((32767) | 1)), 2147483647, 0}
    var v10 int64 = int64(p0)
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(p1, p0, p1)
        fmt.Println(p0)
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v9
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = -2147483648
    var v3 int = v2
    v12 := GenS3{F0: 4294967295}
    if false {
        v2 = -1
        if (true || ("a" >= ("a" + "a"))) {
            switch "hello" {
            case "":
                v1 = 0
            case "foo bar":
                var v13 int64 = 906
                _ = v13
            case "\n\t":
                switch (v2 ^ v2) {
                case 32768:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 int64 = 42
                        var v16 float32 = (float32(0.0) + 1.5)
                        v2 = 4294967295
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                default:
                    var v17 float32 = (3.14159 / 1.0)
                    _ = v17
                }
            }
        } else {
            v18 := GenS4{F0: 127}
            if true {
                if false {
                    v19 := []int{v2}
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    v3 -= (v3 << ((0) & 63))
                    _ = v19
                }
            }
            _ = v18
        }
    }
    var v20 uint64 = 0
    var v21 string = (("foo bar" + "a") + "a")
    v3 -= (-(v3))
    var v22 uint = 4294967295
    v22 = v22
    var v23 uint16 = 65535
    v24 := GenS5{F0: v2}
    v21 = v21
    fmt.Println(v23, v23)
    if true {
        fmt.Println(v2)
        v2 = v2
        if (false || false) {
            if false {
                for fi25 := 0; fi25 < 5; fi25++ {
                    v26 := GenS4{F0: 127}
                    switch v3 {
                    case 32768:
                        if true {
                            switch v21 {
                            case "a":
                                var v27 int64 = 1
                                _ = v27
                            case "foo bar":
                                v28 := GenS2{F0: 1, F1: 4294967295, F2: uint(v20)}
                                _ = v28
                            case "hello":
                                var v29 float32 = 3.14159
                                _ = v29
                            }
                            v3 += (v3 >> ((581) & 63))
                        }
                    default:
                        v30 := GenS6{F0: (v23 << ((v22) & 63)), F1: -474, F2: -299}
                        _ = v30
                    }
                    _ = fi25
                    _ = v26
                }
                v31 := GenS4{F0: 2}
                v1 = v20
                _ = v31
            }
        } else {
            switch (v3 / (((v3 & 1)) | 1)) {
            case 4294967295:
                v32 := GenS3{F0: (v1 / ((884) | 1))}
                _ = v32
            case 32768:
                v21 = "a"
            default:
                v23 = v23
            }
            var v33 float64 = float64(0.1)
            _ = v33
        }
    }
    if true {
        v34 := GenS7{F0: 633, F1: (v22 * uint(v23)), F2: true}
        _ = v34
    } else {
        fmt.Println(v22)
        var v35 uint = (v22 % ((v22) | 1))
        for fi36 := 0; fi36 < 3; fi36++ {
            fmt.Println(v21, v23)
            _ = fi36
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
