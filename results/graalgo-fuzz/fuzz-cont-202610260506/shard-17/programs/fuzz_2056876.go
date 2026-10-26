package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int; F2 uint64 }
func genFn0(p0 int32, p1 int32) uint32 {
    v6 := GenS1{F0: 42}
    var v7 uint = 4294967295
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483648
}
func genFn1() string {
    v8 := GenS1{F0: 4294967295}
    var v9 uint64 = 18446744073709551615
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    v1 := []int{-260}
    v2 := []uint32{301}
    v3 := make(map[int]uint32)
    var v4 int = 1
    var v5 int16 = -836
    var v10 uint32 = 476
    var v11 string = "foo bar"
    v5 -= (v5 << ((945) & 63))
    v10 += v10
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int64 = -59
        fmt.Println(v13, v10)
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int32 = 0
            if (904 <= (v15 - v15)) {
                var v16 int64 = (v13 << ((4294967295) & 63))
                switch -2147483648 {
                case 2147483647:
                    var v17 int = v4
                    _ = v17
                }
                var v18 int32 = v15
                _ = v16
                _ = v18
            }
            var v19 uint = 42
            fmt.Println(v11, v4, v4)
            _ = fi14
            _ = v15
            _ = v19
        }
        _ = fi12
        _ = v13
    }
    var v20 int16 = (-(v5))
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            v23 := GenS1{F0: (v10 - v10)}
            v11 = ((v11 + "foo bar") + v11)
            v11 = v11
            switch "\n\t" {
            case "foo bar":
                if true {
                    var v24 int16 = -1
                    switch (-(v4)) {
                    case 1:
                        var v25 int = fi22
                        _ = v25
                    case -9223372036854775808:
                        v26 := GenS2{F0: 1, F1: 255}
                        _ = v26
                    case -172:
                        var v27 int16 = 1
                        _ = v27
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    fmt.Println(v10, v5, v11)
                    _ = v24
                } else {
                    var v28 uint32 = 65535
                    _ = len(v3)
                    if (false || true) {
                        var v29 int64 = 0
                        v30 := v3[fi22]
                        v20 = (v5 ^ 42)
                        _ = v29
                        _ = v30
                    }
                    _ = v28
                }
            case "\n\t":
                v31 := GenS3{F0: 65535}
                _ = v31
            case "a":
                switch fi22 {
                case -1:
                    var v32 string = v11
                    _ = v32
                case -691:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v33 uint32 = v10
                var v34 int32 = -851
                _ = v33
                _ = v34
            }
            _ = fi22
            _ = v23
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            v2 = append(v2, 1)
            v36 := GenS4{F0: -153, F1: v4, F2: 0}
            var v37 int = -9223372036854775808
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi21
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        v11 = ""
        v4 = v4
        var v39 int16 = v5
        _ = fi38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
