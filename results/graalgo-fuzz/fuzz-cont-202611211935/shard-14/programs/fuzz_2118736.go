package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 int32) int64 {
    switch (("foo bar" + "foo bar") + "foo bar") {
    case "":
        var v6 uint8 = 255
        _ = v6
    case "a":
        if ((!((42 < 32767) && (94 <= 2))) && true) {
            var v7 uint32 = 4294967295
            var v8 uint64 = 42
            _ = v7
            _ = v8
        }
    case "foo bar":
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{fi9, (fi9 << ((42) & 63))}
            switch "\n\t" {
            case "hello":
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case "":
                p0 += (p0 - 0)
            default:
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 int8 = -1
                    _ = fi11
                    _ = v12
                }
            }
            var v13 float64 = (-2.25 * (0.1 + 0.1))
            _ = fi9
            _ = v10
            _ = v13
        }
    default:
        p0 = 2147483647
    }
    v14 := []int{}
    _ = p0
    _ = v14
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 int32 = 0
    var v4 int = 128
    var v5 int32 = (v3 * v3)
    if false {
        v3 -= 2147483647
        var v15 float64 = ((3.14159 * 0.1) / -2.25)
        _ = v15
    } else {
        v4 = 4294967295
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(v4, fi16)
        var v17 int64 = -1
        var v18 float64 = 3.14159
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v3 += v5
        v5 -= v5
        _ = fi19
    }
    switch ("hello" + "foo bar") {
    case "":
        for fi20 := 0; fi20 < 12; fi20++ {
            if false {
                var v21 uint8 = 1
                v2[2147483648] = 129
                _ = v21
            } else {
                v22 := GenS1{F0: 1, F1: 127, F2: (v5 / (((v3 + v3)) | 1))}
                fmt.Println(fi20, v4)
                _ = v22
            }
            _ = fi20
        }
    case "foo bar":
        for fi23 := 0; fi23 < 4; fi23++ {
            delete(v1, "hello")
            var v24 bool = false
            _ = fi23
            _ = v24
        }
    case "hello":
        if false {
            v25 := GenS2{F0: 2, F1: v4}
            v26 := GenS2{F0: 255, F1: 128}
            _ = v25
            _ = v26
        }
    default:
        if (749 <= 2) {
            v5 -= (v5 | (-(v5)))
        } else {
            v5 -= v5
            var v27 int16 = 42
            if false {
                if true {
                    var v28 int8 = -15
                    var v29 int64 = int64(v28)
                    _ = v28
                    _ = v29
                }
            }
            _ = v27
        }
    }
    var v30 uint16 = 42
    v30 -= 65535
    if false {
        v31 := GenS1{F0: 1000, F1: 1, F2: -1}
        var v32 float32 = 1e10
        _ = v31
        _ = v32
    } else {
        v33 := []int{(v4 & v4)}
        var v34 uint16 = (v30 & (^(v30)))
        if (false || true) {
            fmt.Println(v5)
        }
        _ = v33
        _ = v34
    }
    var v35 int16 = 131
    var v36 int64 = -9223372036854775808
    for fi37 := 0; fi37 < 2; fi37++ {
        for fi38 := 0; fi38 < 10; fi38++ {
            var v39 uint8 = 128
            v40 := GenS2{F0: (v39 % (((v39 - (v39 | (v39 | v39)))) | 1)), F1: (^(fi37))}
            var v41 int = fi37
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        if (127 > 1) {
            v3 += int32(v35)
            for fi42 := 0; fi42 < 6; fi42++ {
                v43 := []int{}
                fmt.Println(fi42, fi42, fi37)
                _ = fi42
                _ = v43
            }
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
