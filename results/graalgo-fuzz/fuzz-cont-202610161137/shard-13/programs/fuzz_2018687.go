package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int64 }
func genFn0(p0 int32, p1 int32) int {
    fmt.Println(p1, p0)
    p0 = (p1 * p0)
    p0 -= (-(p0))
    p1 += -2147483648
    _ = p0
    _ = p1
    return -2147483648
}
func genFn1() int {
    var v4 int8 = -65
    if (!true) {
        var v5 float64 = 1e10
        v6 := []int{-701}
        fmt.Println(v5, v4)
        _ = v5
        _ = v6
    }
    var v7 bool = true
    fmt.Println(v4, v7)
    var v8 uint32 = 2147483648
    _ = v4
    _ = v7
    _ = v8
    return 348
}

func main() {
    var v1 uint32 = 948
    v2 := make(map[uint32]int)
    var v3 int8 = 1
    delete(v2, (v1 % (((v1 << ((174) & 63))) | 1)))
    v9 := GenS1{F0: 1, F1: 1}
    var v10 int = -9223372036854775808
    if false {
        for fi11 := 0; fi11 < 16; fi11++ {
            switch ("" + "hello") {
            case "a":
                v1 = (v1 - (v1 % ((v1) | 1)))
            case "foo bar":
                var v12 uint8 = 127
                _ = v12
            }
            _ = fi11
        }
    } else {
        if (false || (v3 <= (v3 - (v3 * v3)))) {
            var v13 uint32 = v1
            if false {
                v14 := GenS1{F0: 459, F1: 0}
                switch ("hello" + "foo bar") {
                case "foo bar":
                    if true {
                        var v15 int16 = 42
                        fmt.Println(v3)
                        var v16 int8 = (v3 + (v3 << ((2) & 63)))
                        _ = v15
                        _ = v16
                    } else {
                        v1 += (v1 >> ((uint(v13)) & 63))
                    }
                case "a":
                    var v17 int8 = 13
                    _ = v17
                case "":
                    v3 = v3
                default:
                    var v18 float64 = 1.5
                    _ = v18
                }
                for fi19 := 0; fi19 < 14; fi19++ {
                    v20 := GenS1{F0: 1, F1: -1}
                    v3 = (v3 ^ 1)
                    var v21 bool = true
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v14
            }
            switch "\n\t" {
            case "hello":
                _ = len(v2)
            }
            _ = v13
        } else {
            fmt.Println(v3, v3)
            for fi22 := 0; fi22 < 14; fi22++ {
                if false {
                    v1 -= 1
                }
                var v23 uint16 = 556
                var v24 uint64 = 0
                _ = fi22
                _ = v23
                _ = v24
            }
            v10 -= v10
        }
        _ = len(v2)
        v1 -= v1
    }
    if true {
        if false {
            v3 = (v3 + v3)
            v25 := []int{v10, genFn0(-536, 1)}
            _ = v25
        } else {
            var v26 int16 = 42
            switch genFn1() {
            case 42:
                fmt.Println(v10, v1)
            }
            _ = v26
        }
    }
    v1 += 4294967295
    switch "a" {
    case "hello":
        fmt.Println(v3)
    }
    fmt.Println(v3)
    var v27 float32 = float32((1.5 * (0.1 / 1e10)))
    v28 := GenS2{F0: 0, F1: (v3 - v3), F2: 71}
    var v29 int32 = -727
    var v30 string = (("foo bar" + "") + "a")
    if true {
        v27 += 1.5
        fmt.Println(v3)
        v3 -= v3
    }
    var v31 int32 = -2147483648
    if false {
        var v32 string = ((v30 + ((v30 + ("hello" + (v30 + "hello"))) + v30)) + v30)
        for fi33 := 0; fi33 < 4; fi33++ {
            if true {
                switch v10 {
                case -84:
                    v10 += fi33
                case -1:
                    v2[v1] = v10
                case 128:
                    switch (v30 + v32) {
                    case "a":
                        var v34 uint32 = 491
                        _ = v34
                    case "hello":
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v36 := []int{fi33, fi33, (fi35 / ((32768) | 1)), fi33}
                            v37 := GenS2{F0: (v10 * fi35), F1: (v3 - -128), F2: 127}
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    case "foo bar":
                        switch fi33 {
                        case -788:
                            var v38 int64 = -1
                            _ = v38
                        case 127:
                            if ((!true) && false) {
                                for fi39 := 0; fi39 < 15; fi39++ {
                                    v40 := []int{fi33, fi39}
                                    _ = fi39
                                    _ = v40
                                }
                                v10 -= 9223372036854775807
                            }
                        case 32767:
                            v29 = 2147483647
                        }
                    }
                }
                for fi41 := 0; fi41 < 5; fi41++ {
                    var v42 uint64 = 29
                    v30 = (v30 + v32)
                    if false {
                        v30 = ((v30 + ("a" + v30)) + "foo bar")
                    }
                    _ = fi41
                    _ = v42
                }
                var v43 int32 = (v29 * -1)
                _ = v43
            }
            _ = fi33
        }
        _ = v32
    } else {
        var v44 uint64 = 0
        _ = v44
    }
    _ = len(v2)
    v45 := GenS3{F0: 2147483647, F1: (v1 - v1), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v45
    _ = v9
}
