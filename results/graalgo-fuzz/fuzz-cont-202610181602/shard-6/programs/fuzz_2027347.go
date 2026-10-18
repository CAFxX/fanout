package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 string }
func genFn0(p0 string) uint32 {
    var v7 uint = 0
    fmt.Println(p0, v7)
    p0 = p0
    if false {
        for fi8 := 0; fi8 < 4; fi8++ {
            if false {
                v9 := []int{}
                switch ("a" + p0) {
                case "\n\t":
                    v10 := GenS1{F0: 538, F1: 0, F2: 255}
                    _ = v10
                case "a":
                    v11 := GenS2{F0: 2147483647}
                    _ = v11
                case "hello":
                    v7 -= 0
                default:
                    if len(v9) > 0 {
                        v9[0] = fi8
                    }
                }
                _ = v9
            } else {
                for fi12 := 0; fi12 < 2; fi12++ {
                    p0 = "hello"
                    v13 := []int{(fi12 ^ fi12), 0, -632, -9223372036854775808}
                    switch (p0 + p0) {
                    case "":
                        v14 := GenS3{F0: -2147483648, F1: 1}
                        _ = v14
                    case "hello":
                        var v15 float64 = (1.5 / (1.5 * 3.14159))
                        _ = v15
                    case "\n\t":
                        v16 := GenS4{F0: v7, F1: 4294967295, F2: ((p0 + p0) + ("a" + "foo bar"))}
                        _ = v16
                    default:
                        v17 := GenS3{F0: 1, F1: -28}
                        _ = v17
                    }
                    _ = fi12
                    _ = v13
                }
                v18 := []int{32768, (fi8 - 42)}
                _ = v18
            }
            switch p0 {
            case "":
                if true {
                    fmt.Println(v7, p0)
                } else {
                    p0 = (p0 + "\n\t")
                }
            case "hello":
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 uint32 = 2147483647
                    p0 = (p0 + p0)
                    _ = fi19
                    _ = v20
                }
            case "a":
                for fi21 := 0; fi21 < 5; fi21++ {
                    if false {
                        var v22 float64 = 3.14159
                        _ = v22
                    } else {
                        p0 = p0
                    }
                    v7 += (v7 - (v7 / (((v7 % (((v7 << ((v7) & 63))) | 1))) | 1)))
                    _ = fi21
                }
            default:
                v23 := []int{-771, fi8, fi8, fi8}
                _ = v23
            }
            _ = fi8
        }
    }
    _ = p0
    _ = v7
    return 2
}

func main() {
    var v1 float64 = ((0.1 + (1.5 - 1e10)) * 1e10)
    v2 := make(map[string]int)
    var v3 float64 = 1e10
    v4 := make(map[int]string)
    var v5 uint16 = 255
    var v6 uint64 = 2
    v24 := []int{128}
    _ = len(v4)
    var v25 int64 = 1
    if (2 > 351) {
        var v26 uint8 = 1
        if true {
            delete(v2, ((("foo bar" + "a") + "\n\t") + ("" + (("foo bar" + ("" + "")) + (("" + "a") + "hello")))))
            fmt.Println(v25, v5)
            fmt.Println(v6)
        } else {
            v1 = v3
            if (true || true) {
                fmt.Println(v6, v3, v6)
            } else {
                v27 := GenS4{F0: 42, F1: 162, F2: ("\n\t" + "hello")}
                _ = v27
            }
            v28 := GenS4{F0: 211, F1: 2147483648, F2: ""}
            _ = v28
        }
        v26 += 67
        _ = v26
    } else {
        _ = len(v24)
    }
    var v29 float32 = (3.14159 / 1.5)
    var v30 uint16 = v5
    switch "hello" {
    case "hello":
        for fi31 := 0; fi31 < 9; fi31++ {
            v6 -= v6
            _ = fi31
        }
    case "foo bar":
        var v32 int8 = 42
        _ = v32
    default:
        if (v5 != v30) {
            var v33 uint64 = 873
            _ = v33
        } else {
            if (449 < 32767) {
                for fi34 := 0; fi34 < 4; fi34++ {
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                    _ = fi34
                }
            } else {
                v6 = v6
            }
            var v35 uint64 = 987
            v5 += (v5 - v5)
            _ = v35
        }
    }
    var v36 uint8 = 116
    for fi37 := 0; fi37 < 16; fi37++ {
        var v38 float64 = (v3 * -2.25)
        fmt.Println(v1)
        _ = fi37
        _ = v38
    }
    v36 += 255
    var v39 uint16 = (v5 - (v30 | v30))
    v5 = uint16(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
