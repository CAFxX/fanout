package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint; F1 int16; F2 int }
type GenS5 struct { F0 bool; F1 uint8; F2 int64 }
func genFn0() bool {
    v7 := []int{-555, 1, 337}
    switch -1 {
    case 17:
        var v8 int = 32768
        _ = v8
    case 4294967295:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case 151:
        var v9 int = 9223372036854775807
        _ = v9
    }
    var v10 float64 = (-2.25 / 0.1)
    var v11 int = 128
    _ = v10
    _ = v11
    _ = v7
    return (true && (-482 < 9223372036854775807))
}
func genFn1(p0 int8, p1 bool, p2 uint16) int64 {
    if p1 {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := []int{(fi12 % (((fi12 - 0)) | 1)), 127}
            var v14 float32 = float32(0.0)
            _ = fi12
            _ = v13
            _ = v14
        }
        if (-2147483648 >= 0) {
            v15 := []int{903, -2147483648}
            _ = v15
        } else {
            fmt.Println(p1, p0, p0)
        }
        var v16 uint = 42
        _ = v16
    } else {
        v17 := GenS1{F0: 2, F1: 351, F2: 2147483648}
        var v18 bool = true
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint32 = 65535
        var v21 string = (("\n\t" + "foo bar") + ("" + ""))
        p0 = (p0 & (p0 ^ (^(p0))))
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 1000, F1: 98}
    v23 := []int{-155}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return int64(p0)
}

func main() {
    var v1 uint8 = 127
    v2 := []string{((("\n\t" + ((("" + "foo bar") + "hello") + "foo bar")) + "\n\t") + ("a" + "a")), ((("a" + ("" + "")) + "a") + ((("a" + ("" + "foo bar")) + "foo bar") + "a"))}
    var v3 uint16 = uint16(v1)
    var v4 uint64 = 782
    var v5 uint8 = v1
    v6 := []uint32{42, 342}
    fmt.Println(v1, v1, v3)
    var v24 uint64 = v4
    v1 = 128
    v1 = v1
    if (false || (!true)) {
        for fi25 := 0; fi25 < 1; fi25++ {
            if len(v6) > 0 {
                v6[0] = 42
            }
            v26 := GenS3{F0: 652, F1: (333 == 534)}
            _ = fi25
            _ = v26
        }
        fmt.Println(v4)
    }
    v27 := GenS4{F0: 1, F1: 42, F2: 32768}
    v28 := GenS5{F0: false, F1: (v5 % (((v1 + v1)) | 1)), F2: 1}
    if (!true) {
        var v29 uint64 = (v24 & v4)
        v30 := GenS4{F0: 42, F1: -32768, F2: 128}
        _ = v29
        _ = v30
    } else {
        var v31 uint32 = 712
        _ = len(v2)
        switch 32768 {
        case 502:
            switch ("a" + ("a" + ("" + "\n\t"))) {
            case "\n\t":
                v5 -= 255
            case "a":
                v32 := GenS1{F0: (v24 + 4294967295), F1: 65535, F2: v31}
                _ = v32
            case "foo bar":
                for fi33 := 0; fi33 < 6; fi33++ {
                    fmt.Println(fi33)
                    if true {
                        v2 = append(v2, ("" + "a"))
                        if (1 >= v5) {
                            v31 += (v31 << ((888) & 63))
                        } else {
                            var v34 int64 = -363
                            _ = v34
                        }
                    } else {
                        var v35 int16 = 137
                        fmt.Println(v5)
                        _ = v35
                    }
                    _ = fi33
                }
            default:
                for fi36 := 0; fi36 < 16; fi36++ {
                    v3 -= v3
                    var v37 string = "a"
                    _ = fi36
                    _ = v37
                }
            }
        }
        _ = v31
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
