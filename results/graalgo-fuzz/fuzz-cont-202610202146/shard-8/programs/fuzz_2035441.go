package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 int8) uint8 {
    switch ("hello" + "a") {
    case "":
        var v7 float64 = float64(((-2.25 + 1e10) / (0.0 - (1.5 * -2.25))))
        _ = v7
    }
    v8 := []int{818, 32767, -987}
    _ = p0
    _ = v8
    return 128
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 32768
    var v3 int32 = 812
    var v4 uint = 65535
    v5 := make(map[int]uint32)
    var v6 uint64 = 2
    for fi9 := 0; fi9 < 5; fi9++ {
        if (false || ((float32(-2.25) / -2.25) != 1e10)) {
            var v10 int = (v2 + (fi9 % ((int(v4)) | 1)))
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int64 = int64(v6)
                var v13 uint32 = 2147483648
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(v2)
            _ = v10
        }
        if true {
            if (int8(v4) > -128) {
                var v14 int8 = -21
                _ = v14
            }
            if true {
                v4 += v4
                v15 := []int{v2, v2}
                _ = v15
            }
            if ((fi9 % ((fi9) | 1)) > (^(fi9))) {
                v3 += (v3 << ((v4) & 63))
            }
        } else {
            var v16 float64 = (1.5 - float64(0.1))
            v17 := []int{fi9, (v2 * (v2 / (((fi9 / ((fi9) | 1))) | 1))), fi9, v2}
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 uint = v4
            _ = fi18
            _ = v19
        }
        _ = fi9
    }
    v1["a"] = 1
    fmt.Println(v4, v6, v6)
    if false {
        v5[(v2 * (v2 << ((65535) & 63)))] = 4294967295
        v3 -= v3
    } else {
        var v20 uint16 = 2
        v2 = v2
        _ = v20
    }
    switch (v2 % (((-(v2))) | 1)) {
    case 32767:
        if ((("a" + "\n\t") + "a") > "hello") {
            v21 := GenS1{F0: 2, F1: 41, F2: 255}
            _ = v21
        } else {
            if (995 != 760) {
                v22 := GenS2{F0: 2, F1: 104, F2: 1}
                var v23 uint64 = 0
                _ = v22
                _ = v23
            } else {
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 float64 = 0.0
                    v6 = 18446744073709551615
                    if (!(78 > -98)) {
                        v1[""] = 388
                        var v26 int64 = 42
                        _ = v26
                    } else {
                        v6 = 42
                        v27 := GenS2{F0: uint32(v6), F1: 75, F2: 65535}
                        var v28 int = (v2 ^ (fi24 + (v2 * v2)))
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v6, v3, v4)
            }
        }
    case 32768:
        _ = len(v1)
    case -510:
        var v29 int = (-(v2))
        _ = v29
    }
    var v30 int64 = int64(v3)
    _ = len(v1)
    switch (^(v2)) {
    case 4294967295:
        var v31 uint = (v4 & v4)
        _ = v31
    }
    v6 -= v6
    v30 -= 42
    v32 := GenS3{F0: v4, F1: v6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
