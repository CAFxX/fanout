package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 int32, p1 int64) int {
    fmt.Println(p0, p0)
    v4 := GenS1{F0: (p1 | -1)}
    p1 += 0
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}
func genFn1() uint8 {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint = 18446744073709551615
        v6 += 967
        _ = fi5
        _ = v6
    }
    switch -2147483648 {
    case -9223372036854775808:
        var v7 int = 263
        _ = v7
    case 41:
        var v8 uint = 2
        _ = v8
    case 4294967295:
        if true {
            var v9 uint32 = 1
            _ = v9
        } else {
            if false {
                var v10 uint8 = 211
                var v11 string = "\n\t"
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 uint = 65535
                    _ = fi12
                    _ = v13
                }
                _ = v10
                _ = v11
            } else {
                v14 := GenS1{F0: 9223372036854775807}
                v15 := []int{}
                _ = v14
                _ = v15
            }
        }
    default:
        var v16 int64 = 0
        _ = v16
    }
    v17 := []int{-9223372036854775808, -1, -243, 615}
    var v18 int8 = 42
    v19 := GenS2{F0: 946}
    _ = v17
    _ = v18
    _ = v19
    return 42
}

func main() {
    var v1 uint32 = 65535
    var v2 uint = 979
    var v3 uint16 = 65535
    var v20 int16 = 1000
    v2 += v2
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := []int{fi21, -9223372036854775808, 970}
        _ = len(v22)
        var v23 bool = false
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v3, v3, v20)
    v1 = 42
    if false {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 int8 = 1
            if true {
                switch (-(fi24)) {
                case 42:
                    var v26 int64 = -9223372036854775808
                    _ = v26
                case 32768:
                    if true {
                        switch fi24 {
                        case 42:
                            var v27 uint = v2
                            _ = v27
                        case 9223372036854775807:
                            fmt.Println(fi24, v1, v25)
                        case 32767:
                            v1 = 42
                        }
                        v3 = v3
                    }
                case 32767:
                    fmt.Println(v25, fi24, v20)
                }
            }
            _ = fi24
            _ = v25
        }
        if (true && false) {
            if ((v1 | v1) <= (v1 + 216)) {
                if false {
                    var v28 int = genFn0(-2147483648, -264)
                    var v29 string = "\n\t"
                    v1 = 321
                    _ = v28
                    _ = v29
                } else {
                    var v30 string = "foo bar"
                    switch 2147483647 {
                    case 42:
                        v31 := GenS3{F0: (v2 / ((4294967295) | 1))}
                        _ = v31
                    case 128:
                        v2 -= v2
                    case 1:
                        var v32 bool = false
                        _ = v32
                    default:
                        v33 := GenS4{F0: 2}
                        _ = v33
                    }
                    _ = v30
                }
                v34 := []int{4294967295, 50, -915, 160}
                _ = v34
            }
            var v35 string = "hello"
            _ = v35
        } else {
            v2 = (v2 / ((v2) | 1))
        }
    }
    v36 := GenS3{F0: v2}
    for fi37 := 0; fi37 < 11; fi37++ {
        fmt.Println(v1)
        _ = fi37
    }
    v38 := GenS5{F0: 42, F1: 2147483647, F2: 2}
    v39 := []int{}
    if ((!true) || ((-128 < 127) && true)) {
        if (("hello" + ("hello" + ("foo bar" + "\n\t"))) <= "hello") {
            var v40 int16 = 0
            v41 := GenS3{F0: 0}
            var v42 string = ("" + "\n\t")
            _ = v40
            _ = v41
            _ = v42
        }
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v38
    _ = v39
}
