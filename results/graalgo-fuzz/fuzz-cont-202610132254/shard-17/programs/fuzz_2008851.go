package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint, p1 uint) uint8 {
    switch "foo bar" {
    case "\n\t":
        p0 -= 0
    case "a":
        var v4 uint64 = 724
        _ = v4
    }
    p1 = p0
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(p1, p1)
        _ = fi5
    }
    var v6 int8 = 35
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            v6 -= v6
            _ = fi8
        }
        switch ("foo bar" + ("foo bar" + "foo bar")) {
        case "\n\t":
            if ((v6 * (^(v6))) > (-(v6))) {
                var v9 uint16 = 1
                var v10 int32 = 1
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 11; fi11++ {
                    if (!(183 <= 42)) {
                        var v12 float32 = (0.0 - (1.5 - 0.0))
                        var v13 uint = p0
                        _ = v12
                        _ = v13
                    } else {
                        p1 -= (p1 >> (((p1 % ((42) | 1))) & 63))
                    }
                    var v14 int = fi7
                    p0 += p0
                    _ = fi11
                    _ = v14
                }
                var v15 float64 = ((3.14159 - (1e10 - 0.0)) + 1.5)
                _ = v15
            }
        case "a":
            switch "a" {
            case "":
                v16 := GenS1{F0: 2}
                _ = v16
            case "\n\t":
                p0 = (p1 + (p1 - p1))
            }
        case "":
            switch "\n\t" {
            case "":
                var v17 uint64 = 18446744073709551615
                _ = v17
            case "foo bar":
                v18 := GenS1{F0: (^(p0))}
                _ = v18
            case "\n\t":
                v19 := GenS1{F0: 301}
                _ = v19
            }
        default:
            for fi20 := 0; fi20 < 1; fi20++ {
                if (42 == 0) {
                    p0 = p0
                    var v21 float64 = 0.0
                    v22 := []int{fi7, fi7, (-(fi7)), 0}
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return 255
}
func genFn1() int32 {
    var v23 uint8 = 2
    if (((v23 - genFn0(uint(v23), 2)) < (v23 - v23)) && true) {
        var v24 uint16 = 32767
        var v25 uint64 = 42
        v25 += 580
        _ = v24
        _ = v25
    } else {
        fmt.Println(v23)
        v26 := []int{0}
        v27 := GenS1{F0: 9}
        _ = v26
        _ = v27
    }
    _ = v23
    return -297
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    var v3 int8 = -1
    v3 += v3
    v2 = v2
    v3 = 1
    var v28 int32 = 42
    switch 32768 {
    case 325:
        v29 := GenS2{F0: 2147483648}
        _ = v29
    case -1:
        for fi30 := 0; fi30 < 12; fi30++ {
            v31 := GenS3{F0: 0}
            if true {
                v32 := []int{fi30, fi30, fi30}
                _ = v32
            }
            for fi33 := 0; fi33 < 18; fi33++ {
                var v34 int64 = 42
                var v35 uint32 = 1
                v36 := GenS2{F0: 0}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi30
            _ = v31
        }
    case 1:
        var v37 uint8 = 4
        _ = v37
    default:
        var v38 uint32 = 501
        _ = v38
    }
    fmt.Println(v2, v28)
    v3 = 1
    var v39 bool = (v2 || (6 >= 127))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
}
