package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint64, p1 bool) int8 {
    var v4 uint32 = 65535
    v5 := GenS1{F0: ((p0 % ((p0) | 1)) == (p0 * (p0 & 27)))}
    if true {
        fmt.Println(p0, p1)
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int64 = -262
            _ = fi6
            _ = v7
        }
    } else {
        v8 := GenS2{F0: 255, F1: 1, F2: 255}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 bool = (0 > -55)
    var v2 uint64 = 921
    v3 := make(map[int]uint32)
    fmt.Println(v1)
    switch "a" {
    case "\n\t":
        v9 := []int{}
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := []int{fi10, 1, (fi10 & 684), 42}
            _ = fi10
            _ = v11
        }
    }
    if v1 {
        v1 = (9223372036854775807 == 1)
        if (v2 >= v2) {
            v2 = (v2 << ((828) & 63))
        }
        switch 128 {
        case 30:
            for fi12 := 0; fi12 < 6; fi12++ {
                if (!(-1 != -2147483648)) {
                    v2 = 4294967295
                    v13 := GenS1{F0: v1}
                    v2 = (v2 >> ((483) & 63))
                    _ = v13
                }
                var v14 bool = v1
                _ = fi12
                _ = v14
            }
        case 128:
            v1 = (!(2 < 32767))
        case 127:
            var v15 float32 = float32(((0.0 + (-2.25 * 0.1)) + -2.25))
            _ = v15
        }
    }
    v2 = v2
    for fi16 := 0; fi16 < 9; fi16++ {
        v3[(fi16 % ((-1) | 1))] = 42
        var v17 float32 = 1.5
        var v18 uint = 303
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 string = (("foo bar" + "") + ("\n\t" + "a"))
    var v20 int16 = -1
    switch v19 {
    case "a":
        v19 = (v19 + "foo bar")
    case "":
        delete(v3, 2147483647)
    default:
        fmt.Println(v20)
    }
    var v21 int64 = 860
    v19 = v19
    if v1 {
        fmt.Println(v19, v21)
    } else {
        fmt.Println(v19)
        for fi22 := 0; fi22 < 6; fi22++ {
            for fi23 := 0; fi23 < 3; fi23++ {
                v21 = 0
                _ = fi23
            }
            v24 := GenS3{F0: genFn0(v2, (!v1))}
            _ = fi22
            _ = v24
        }
        var v25 int64 = (^(v21))
        _ = v25
    }
    var v26 uint64 = v2
    var v27 uint16 = 255
    if (100 > -128) {
        if (true && v1) {
            var v28 float32 = ((1.5 / 1.5) + 1e10)
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 uint16 = (v27 ^ 32767)
            switch (fi30 % (((fi30 & (fi29 & int(v21)))) | 1)) {
            case 32768:
                v26 -= 507
            case -2147483648:
                var v32 int32 = -1
                _ = v32
            }
            switch fi29 {
            case 0:
                v26 -= v26
            default:
                v33 := []int{}
                _ = v33
            }
            _ = fi30
            _ = v31
        }
        if v1 {
            v34 := []int{fi29, 32767, (fi29 | (fi29 % (((fi29 >> ((4294967295) & 63))) | 1)))}
            var v35 int8 = 88
            v36 := GenS2{F0: 1, F1: -2147483648, F2: 42}
            _ = v34
            _ = v35
            _ = v36
        }
        switch fi29 {
        case 219:
            v37 := GenS2{F0: 226, F1: -1, F2: 42}
            _ = v37
        case -1:
            var v38 uint = 1
            _ = v38
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
}
