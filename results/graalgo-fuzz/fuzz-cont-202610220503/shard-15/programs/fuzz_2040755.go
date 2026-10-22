package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 string }
func genFn0(p0 uint32) uint {
    var v5 float64 = 0.1
    if (false || false) {
        v5 += v5
        fmt.Println(v5)
        v5 = (1.5 + (((0.1 + (v5 - v5)) * (v5 - v5)) - 1e10))
    } else {
        var v6 bool = false
        switch "\n\t" {
        case "a":
            var v7 int64 = 163
            _ = v7
        default:
            var v8 uint8 = 42
            _ = v8
        }
        _ = v6
    }
    if (true || true) {
        v9 := []int{2147483647}
        switch 9223372036854775807 {
        case 1:
            var v10 float32 = 0.0
            _ = v10
        }
        _ = v9
    } else {
        p0 -= p0
        p0 += (p0 << ((776) & 63))
    }
    v11 := GenS1{F0: -1}
    fmt.Println(p0)
    _ = p0
    _ = v11
    _ = v5
    return 0
}
func genFn1() uint16 {
    v12 := []int{-848, 128}
    var v13 string = "a"
    switch 128 {
    case 42:
        v12 = append(v12, 42)
    case -1:
        _ = len(v12)
    default:
        var v14 uint32 = 2147483648
        _ = v14
    }
    var v15 float64 = 3.14159
    v12 = append(v12, -220)
    _ = v12
    _ = v13
    _ = v15
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = float32((3.14159 - 1e10))
    var v3 string = "a"
    v4 := make(map[int]bool)
    var v16 uint32 = 4294967295
    var v17 uint = 4294967295
    if (true || false) {
        _ = len(v4)
    } else {
        v18 := v4[-365]
        _ = v18
    }
    v19 := GenS1{F0: 42}
    fmt.Println(v3, v16, v16)
    v17 += genFn0((v16 ^ 434))
    var v20 int64 = -71
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 13; fi22++ {
            switch "hello" {
            case "hello":
                if (3.14159 < -2.25) {
                    fmt.Println(v2, v20, fi21)
                    _ = len(v1)
                } else {
                    var v23 string = v3
                    fmt.Println(v23)
                    _ = v23
                }
            }
            if ("foo bar" != v3) {
                var v24 string = v3
                _ = v24
            }
            v16 -= v16
            _ = fi22
        }
        _ = fi21
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        switch v3 {
        case "hello":
            v26 := []int{-865}
            _ = v26
        default:
            fmt.Println(v3, fi25)
        }
        _ = fi25
    }
    if (1000 != -1) {
        switch 128 {
        case 128:
            switch ("\n\t" + v3) {
            case "":
                if (-1 >= 1) {
                    var v27 string = (v3 + "hello")
                    switch 127 {
                    case 78:
                        var v28 int32 = -2147483648
                        _ = v28
                    case 2147483647:
                        var v29 int = -943
                        _ = v29
                    case 625:
                        switch 128 {
                        case 132:
                            for fi30 := 0; fi30 < 12; fi30++ {
                                var v31 float64 = 1e10
                                _ = fi30
                                _ = v31
                            }
                        case -816:
                            v3 = v27
                        case 42:
                            switch (v27 + ((v3 + (("" + "a") + v3)) + "\n\t")) {
                            case "hello":
                                var v32 uint32 = v16
                                _ = v32
                            }
                        default:
                            for fi33 := 0; fi33 < 5; fi33++ {
                                v34 := GenS2{F0: 42}
                                var v35 int = fi33
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        }
                    }
                    for fi36 := 0; fi36 < 10; fi36++ {
                        fmt.Println(v3)
                        v16 += v16
                        v3 = ((v3 + "hello") + "\n\t")
                        _ = fi36
                    }
                    _ = v27
                } else {
                    var v37 uint64 = 4294967295
                    var v38 int = 128
                    _ = v37
                    _ = v38
                }
            case "a":
                if ((133 == 127) || (((v3 + v3) + v3) > "\n\t")) {
                    for fi39 := 0; fi39 < 8; fi39++ {
                        v3 = "hello"
                        _ = fi39
                    }
                    fmt.Println(v3, v3)
                    v16 = v16
                } else {
                    v40 := GenS2{F0: 0}
                    v4[1] = false
                    _ = v40
                }
            }
        case 715:
            var v41 int16 = 22
            _ = v41
        case 127:
            var v42 float32 = v2
            _ = v42
        default:
            var v43 uint8 = 2
            _ = v43
        }
        v44 := GenS3{F0: 881, F1: -32768, F2: v3}
        _ = v44
    }
    v45 := []int{-174, -1, 279, 128}
    if len(v45) > 0 {
        v45[0] = -1
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v17, v2, v16)
    case "a":
        v3 = v3
    default:
        var v46 int = -2147483648
        _ = v46
    }
    v47 := v4[4294967295]
    v20 -= (v20 & v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v47
}
