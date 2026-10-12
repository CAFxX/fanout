package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 int64; F2 int16 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint32, p1 bool, p2 uint32) uint16 {
    v6 := []int{-1, 32768, 999, 663}
    if p1 {
        v7 := GenS1{F0: -1, F1: 1}
        _ = v7
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 uint) uint32 {
    p0 += (p0 - 4294967295)
    if true {
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 float32 = 0.1
                var v11 uint64 = 748
                _ = fi9
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 11; fi12++ {
                if false {
                    v13 := GenS1{F0: 42, F1: 75}
                    if (!(!true)) {
                        p0 = p0
                        p0 -= p0
                    }
                    v14 := []int{fi12, (fi12 & fi8), fi8, 42}
                    _ = v13
                    _ = v14
                }
                v15 := []int{fi8}
                fmt.Println(fi12, fi8)
                _ = fi12
                _ = v15
            }
            _ = fi8
        }
        if (0.0 <= 1e10) {
            v16 := GenS2{F0: 42, F1: 9223372036854775807, F2: 815}
            _ = v16
        }
        var v17 float64 = -2.25
        _ = v17
    }
    if true {
        p0 -= (^(p0))
    }
    _ = p0
    return 65535
}

func main() {
    var v1 uint16 = 494
    v2 := make(map[uint32]uint32)
    var v3 int8 = -83
    v4 := make(map[string]bool)
    var v5 float64 = (1e10 * 0.1)
    var v18 int64 = 277
    delete(v2, 4294967295)
    var v19 uint8 = 12
    var v20 int8 = int8(v1)
    if false {
        if (42 <= 0) {
            fmt.Println(v3)
        } else {
            v21 := GenS2{F0: -808, F1: -442, F2: 882}
            var v22 int16 = -78
            var v23 int32 = 42
            _ = v21
            _ = v22
            _ = v23
        }
        if true {
            v19 -= v19
        } else {
            v24 := GenS1{F0: (^(v18)), F1: 100}
            _ = v24
        }
    }
    v25 := []int{313, 127, 854, -207}
    var v26 string = ""
    v26 = v26
    for fi27 := 0; fi27 < 19; fi27++ {
        if (1 > (v1 * v1)) {
            var v28 bool = false
            var v29 uint32 = 26
            if false {
                var v30 uint8 = v19
                v31 := GenS3{F0: -1, F1: -230}
                v26 = v26
                _ = v30
                _ = v31
            }
            _ = v28
            _ = v29
        } else {
            fmt.Println(v18, v5, fi27)
            if true {
                var v32 int8 = v20
                _ = v32
            } else {
                fmt.Println(v19)
            }
        }
        var v33 string = "foo bar"
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 uint16 = uint16(v19)
            v19 = 2
            var v36 uint = 4294967295
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi27
        _ = v33
    }
    delete(v4, ("" + v26))
    var v37 float32 = 3.14159
    var v38 int64 = v18
    switch v26 {
    case "":
        v19 -= 130
    case "hello":
        var v39 uint = 2
        _ = v39
    }
    _ = len(v25)
    v20 -= (v3 >> ((881) & 63))
    if false {
        v40 := GenS4{F0: 32767}
        v41 := GenS1{F0: 0, F1: int8(v38)}
        _ = v40
        _ = v41
    } else {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        if (!((v19 & v19) < (v19 + (^(v19))))) {
            for fi42 := 0; fi42 < 15; fi42++ {
                v38 = (v38 & -9223372036854775808)
                _ = fi42
            }
        }
        var v43 uint8 = 138
        _ = v43
    }
    v26 = "hello"
    v20 -= (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
