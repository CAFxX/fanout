package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int32 }
func genFn0(p0 int8, p1 uint) int64 {
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint8 = 161
        v7 := GenS1{F0: (p1 < (p1 * (p1 >> (((p1 % ((p1) | 1))) & 63)))), F1: 2, F2: -745}
        p0 += p0
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint = (p1 | 18446744073709551615)
    _ = p0
    _ = p1
    _ = v8
    return -1
}
func genFn1(p0 uint64) uint16 {
    var v9 uint32 = 65535
    if false {
        var v10 int32 = 0
        _ = v10
    }
    var v11 int8 = -48
    switch -223 {
    case -1:
        v11 = v11
    case -9223372036854775808:
        v11 -= (v11 << ((18446744073709551615) & 63))
    case -642:
        for fi12 := 0; fi12 < 1; fi12++ {
            p0 -= (p0 | p0)
            v13 := []int{(-(fi12)), (fi12 + fi12), (fi12 & fi12)}
            switch fi12 {
            case -2147483648:
                var v14 float64 = float64(1e10)
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = v11
    _ = v9
    return 255
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[string]bool)
    var v3 int = -9223372036854775808
    v4 := []uint32{2, 42}
    fmt.Println(v3)
    v3 += (v3 * (v3 >> ((2) & 63)))
    _ = len(v2)
    v3 -= 0
    var v15 float64 = float64((0.0 * -2.25))
    var v16 int32 = -1
    if (float32((0.0 - float64(-2.25))) == ((1e10 + 0.1) + (1.5 * float32(float64(1e10))))) {
        if false {
            var v17 int32 = v16
            _ = v17
        } else {
            var v18 float64 = (1.5 - float64(0.1))
            var v19 int64 = 0
            var v20 int = v3
            _ = v18
            _ = v19
            _ = v20
        }
    } else {
        v21 := GenS1{F0: true, F1: 1, F2: -2147483648}
        _ = v21
    }
    switch (v3 / (((v3 | -498)) | 1)) {
    case 32768:
        fmt.Println(v3)
    case 128:
        if (-1 == -1) {
            fmt.Println(v3, v15, v15)
            var v22 uint64 = 4294967295
            v23 := GenS1{F0: (v16 < -2147483648), F1: 42, F2: (v16 / ((v16) | 1))}
            _ = v22
            _ = v23
        }
    case -348:
        var v24 uint16 = 814
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v1 = append(v1, ("a" + "\n\t"))
        var v26 bool = (!true)
        fmt.Println(v15, v26, v15)
        _ = fi25
        _ = v26
    }
    if true {
        var v27 uint8 = 127
        v3 += (v3 << ((255) & 63))
        _ = v27
    }
    var v28 uint16 = genFn1(574)
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 int32 = -1
        v28 = v28
        _ = fi29
        _ = v30
    }
    if false {
        var v31 int = (v3 - (v3 / ((128) | 1)))
        var v32 uint64 = 2
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        for fi34 := 0; fi34 < 5; fi34++ {
            fmt.Println(v3)
            if false {
                delete(v2, "hello")
                v35 := GenS1{F0: false, F1: 1, F2: v16}
                _ = v35
            } else {
                v36 := GenS1{F0: true, F1: 257, F2: (v16 & v16)}
                var v37 int16 = 1
                _ = v36
                _ = v37
            }
            v3 += (fi33 + fi34)
            _ = fi34
        }
        var v38 uint16 = 0
        var v39 uint64 = 0
        _ = fi33
        _ = v38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
