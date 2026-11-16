package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0(p0 int32, p1 int32) uint8 {
    p0 = p0
    var v5 uint32 = 398
    var v6 bool = ((p0 ^ p0) >= (p1 ^ p0))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int16) uint {
    var v7 uint32 = 2147483647
    v7 += (v7 >> ((134) & 63))
    v8 := GenS1{F0: 42, F1: -1}
    var v9 string = "foo bar"
    var v10 float64 = 0.0
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float32 = 0.0
    var v2 int = -1
    v3 := make(map[int]int)
    var v4 uint64 = 217
    v1 -= ((v1 * v1) + 1e10)
    v1 += v1
    switch "foo bar" {
    case "foo bar":
        for fi11 := 0; fi11 < 20; fi11++ {
            v3[2147483647] = (fi11 | fi11)
            _ = fi11
        }
    }
    fmt.Println(v2, v2, v2)
    v1 = v1
    v12 := GenS1{F0: (^(v4)), F1: 42}
    var v13 uint32 = 2147483648
    if (v13 < (v13 / ((uint32(v4)) | 1))) {
        var v14 int64 = 492
        delete(v3, v2)
        _ = v14
    } else {
        switch (v2 * (v2 - (v2 ^ (v2 % ((128) | 1))))) {
        case 32767:
            if (!(-1 != -1)) {
                if false {
                    fmt.Println(v13, v4, v4)
                    switch (("foo bar" + ("hello" + ("" + ("a" + "")))) + (((("" + "\n\t") + "foo bar") + "a") + "hello")) {
                    case "foo bar":
                        v15 := []int{-708, v2, v2}
                        _ = v15
                    case "\n\t":
                        fmt.Println(v2, v1, v4)
                    }
                } else {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v17 := v3[v2]
                        delete(v3, v2)
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS2{F0: 0}
                    _ = v18
                }
            }
        }
        if (!true) {
            v4 += v4
            fmt.Println(v4, v2)
            for fi19 := 0; fi19 < 16; fi19++ {
                v4 = v4
                v20 := GenS3{F0: v13, F1: 252, F2: -905}
                v4 += (v4 % ((v4) | 1))
                _ = fi19
                _ = v20
            }
        }
        v21 := GenS2{F0: -358}
        _ = v21
    }
    var v22 uint64 = v4
    var v23 int8 = 127
    v23 = v23
    for fi24 := 0; fi24 < 2; fi24++ {
        switch (fi24 / ((fi24) | 1)) {
        case 42:
            switch (v2 ^ v2) {
            case -1:
                var v25 bool = ((1.5 >= 0.0) && true)
                _ = v25
            case 132:
                switch (v2 & -845) {
                case -274:
                    var v26 float32 = v1
                    _ = v26
                case -543:
                    var v27 int32 = 42
                    _ = v27
                default:
                    if true {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            v29 := GenS2{F0: 0}
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        var v30 int = (fi24 << ((939) & 63))
                        fmt.Println(v23, v2)
                        v4 -= 569
                        _ = v30
                    }
                }
            case 32767:
                v31 := GenS4{F0: "\n\t", F1: "a"}
                _ = v31
            }
        case 32767:
            v32 := GenS5{F0: 1000}
            _ = v32
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            v23 -= v23
            v34 := GenS6{F0: 255, F1: 42, F2: v4}
            if (42 != 9223372036854775807) {
                v22 += (v22 * 867)
            } else {
                var v35 int = fi24
                v3[fi33] = fi24
                v36 := GenS2{F0: 32767}
                v37 := GenS1{F0: 1, F1: 1}
                var v38 int16 = -32768
                v39 := GenS1{F0: (v4 << ((517) & 63)), F1: 896}
                v1 = (float32(1.5) / v1)
                _ = v35
                _ = v36
                _ = v37
                _ = v38
                _ = v39
            }
            var v40 bool = false
            v13 += 1
            _ = fi33
            _ = v34
            _ = v40
        }
        _ = fi24
    }
    var v41 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
}
