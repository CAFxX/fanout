package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0(p0 int, p1 uint64) uint64 {
    v4 := GenS1{F0: -970, F1: -451}
    v5 := []int{(p0 + (p0 | (p0 | 1))), p0, -977}
    v6 := GenS2{F0: "a", F1: -2147483648, F2: 862}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return p1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -1
    v3 := make(map[uint32]string)
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(fi8, fi8, fi7)
            _ = fi8
        }
        if true {
            var v9 int8 = 100
            _ = len(v1)
            _ = v9
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := []int{(fi10 - fi10), fi10}
                v12 := GenS3{F0: 232, F1: v2}
                _ = fi10
                _ = v11
                _ = v12
            }
            v2 = (v2 >> ((65535) & 63))
        }
        _ = fi7
    }
    switch 127 {
    case 0:
        var v13 float64 = -2.25
        _ = v13
    case 1:
        if false {
            var v14 int32 = -1
            _ = v14
        }
    }
    if false {
        v2 -= (v2 >> ((91) & 63))
        fmt.Println(v2)
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(v2, fi15)
            _ = fi15
        }
    }
    v16 := GenS2{F0: (("foo bar" + "") + "a"), F1: 0, F2: 255}
    v17 := []int{32767, -204}
    if (!(!(-1 <= -1))) {
        v2 = (v2 * (^(v2)))
        v18 := GenS1{F0: (-(v2)), F1: -784}
        _ = v18
    }
    v2 = (v2 & v2)
    for fi19 := 0; fi19 < 11; fi19++ {
        if false {
            var v20 float64 = -2.25
            v21 := v3[437]
            v21 = v21
            _ = v20
            _ = v21
        }
        v22 := GenS2{F0: "hello", F1: (v2 % (((v2 % (((-(v2))) | 1))) | 1)), F2: 0}
        var v23 float64 = ((float64(3.14159) * 1.5) * 3.14159)
        _ = fi19
        _ = v22
        _ = v23
    }
    var v24 int32 = 1
    var v25 bool = (!(!(42 == -1)))
    if false {
        delete(v3, 1)
        v26 := v1[0]
        switch ("a" + "") {
        case "foo bar":
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint = 909
                _ = fi27
                _ = v28
            }
        case "":
            v29 := GenS3{F0: 129, F1: v2}
            _ = v29
        }
        _ = v26
    } else {
        _ = len(v3)
        if true {
            v25 = (((0.1 / 1e10) - 1e10) > 0.1)
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        } else {
            var v30 int8 = 5
            switch "\n\t" {
            case "":
                var v31 string = ""
                _ = v31
            case "hello":
                v30 += (v30 >> ((856) & 63))
            case "a":
                var v32 int64 = 0
                _ = v32
            default:
                var v33 bool = v25
                _ = v33
            }
            _ = v30
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
