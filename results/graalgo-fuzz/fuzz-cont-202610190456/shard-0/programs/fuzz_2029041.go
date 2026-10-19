package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint16; F2 int }
type GenS4 struct { F0 uint; F1 uint64 }
func genFn0(p0 int8, p1 int, p2 int8) uint16 {
    v5 := []int{}
    fmt.Println(p0)
    var v6 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]bool)
    v3 := []int{-2147483648, 32767, -9223372036854775808, 259}
    var v4 string = ((("" + "hello") + "foo bar") + "\n\t")
    if false {
        v4 = (v4 + v4)
        var v7 float64 = 0.0
        _ = v7
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(v4, v4)
            var v9 float32 = (0.0 - (-2.25 + 1e10))
            if (fi8 > fi8) {
                for fi10 := 0; fi10 < 8; fi10++ {
                    fmt.Println(v4)
                    if false {
                        v11 := v2[fi8]
                        _ = v11
                    } else {
                        switch v4 {
                        case "a":
                            fmt.Println(fi10, v4)
                        }
                    }
                    _ = fi10
                }
                for fi12 := 0; fi12 < 15; fi12++ {
                    switch (("" + v4) + v4) {
                    case "a":
                        var v13 uint32 = 42
                        _ = v13
                    case "foo bar":
                        v9 = 1.5
                        v14 := v2[(^(fi8))]
                        _ = v14
                    case "hello":
                        v9 -= ((v9 / 1.5) + v9)
                    default:
                        var v15 int16 = 262
                        if ((1 >= 2147483647) || true) {
                            if (!false) {
                                v4 = v4
                            } else {
                                switch fi8 {
                                case -9223372036854775808:
                                    var v16 bool = true
                                    _ = v16
                                }
                                v15 = 147
                            }
                        } else {
                            var v17 float64 = 1.5
                            v18 := GenS1{F0: 0, F1: v4}
                            _ = v17
                            _ = v18
                        }
                        v15 += (^(v15))
                        _ = v15
                    }
                    switch ("hello" + (v4 + v4)) {
                    case "":
                        var v19 int = (fi12 - -2147483648)
                        _ = v19
                    }
                    _ = fi12
                }
                v20 := GenS1{F0: -15, F1: v4}
                _ = v20
            }
            _ = fi8
            _ = v9
        }
        if true {
            if (340 == 1) {
                for fi21 := 0; fi21 < 20; fi21++ {
                    v4 = ((v4 + ((v4 + (("foo bar" + "") + v4)) + v4)) + "hello")
                    v22 := GenS2{F0: fi21}
                    _ = fi21
                    _ = v22
                }
                var v23 uint32 = 2
                v1 = v1
                _ = v23
            } else {
                var v24 uint64 = v1
                _ = v24
            }
            v25 := GenS2{F0: 42}
            if true {
                var v26 int16 = 42
                var v27 float32 = 3.14159
                v27 = v27
                _ = v26
                _ = v27
            } else {
                if false {
                    v4 = (("foo bar" + "a") + "foo bar")
                    for fi28 := 0; fi28 < 20; fi28++ {
                        v4 = v4
                        v29 := GenS1{F0: 1, F1: "hello"}
                        _ = fi28
                        _ = v29
                    }
                }
                v1 = v1
            }
            _ = v25
        }
        fmt.Println(v1)
    }
    v2[32767] = false
    var v30 uint16 = 121
    var v31 int = 997
    v32 := GenS3{F0: -929, F1: (v30 ^ genFn0(-128, v31, 0)), F2: v31}
    if true {
        var v33 int16 = int16(v1)
        var v34 uint = uint(v1)
        _ = v33
        _ = v34
    } else {
        var v35 uint = 299
        v36 := GenS4{F0: (^(v35)), F1: (v1 ^ (v1 & (^(v1))))}
        _ = v35
        _ = v36
    }
    _ = len(v2)
    v4 = v4
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 string = (v4 + v4)
        var v39 uint32 = 61
        var v40 float64 = 1.5
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
