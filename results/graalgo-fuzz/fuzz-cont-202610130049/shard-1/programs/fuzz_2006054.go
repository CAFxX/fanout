package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int64, p1 uint8, p2 int16) uint32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        switch "foo bar" {
        case "":
            var v7 int8 = -1
            _ = v7
        case "\n\t":
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int8 = 127
                p2 -= (p2 << ((74) & 63))
                p0 = p0
                if (true || (true && (527 <= 42))) {
                    v10 := GenS1{F0: (v9 * v9), F1: int16(p1), F2: (v9 * (v9 << ((82) & 63)))}
                    var v11 int64 = p0
                    _ = v10
                    _ = v11
                } else {
                    if (((p0 % ((p0) | 1)) <= p0) || (198 >= 4294967295)) {
                        if ((!true) || (1 != 1)) {
                            var v12 uint8 = p1
                            var v13 int16 = 1000
                            var v14 int8 = v9
                            var v15 uint32 = 0
                            var v16 int16 = p2
                            _ = v12
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                }
                var v17 float32 = 3.14159
                _ = fi8
                _ = v17
                _ = v9
            }
        case "foo bar":
            fmt.Println(fi6, p2, fi6)
        default:
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS1{F0: 127, F1: p2, F2: -26}
                _ = fi18
                _ = v19
            }
        }
        v20 := []int{fi6, fi6, fi6}
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := GenS1{F0: 63, F1: (p2 << ((4294967295) & 63)), F2: -128}
            if true {
                var v23 uint64 = 900
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        _ = fi6
        _ = v20
    }
    v24 := GenS2{F0: 2}
    p2 -= 0
    p2 = (p2 + (^(p2)))
    fmt.Println(p0, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 2
}

func main() {
    var v1 uint32 = 0
    var v2 float32 = 0.1
    var v3 uint = 4294967295
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]bool)
    v2 = (float32(((3.14159 / 1e10) * 0.0)) - v2)
    v2 -= 3.14159
    v3 -= 65535
    v25 := v5[genFn0(-847, 42, -1)]
    v26 := GenS1{F0: -1, F1: 214, F2: 100}
    var v27 bool = ((!v25) || false)
    v1 = 4294967295
    v27 = true
    var v28 int8 = 1
    if v27 {
        v29 := GenS2{F0: (^(v3))}
        v3 = 2
        v1 = v1
        _ = v29
    } else {
        v30 := GenS1{F0: 127, F1: 1000, F2: 42}
        v3 -= (v3 << ((63) & 63))
        _ = v30
    }
    fmt.Println(v1)
    v4[(v1 * (v1 + v1))] = 630
    v25 = (genFn0(1, 48, 32767) == v1)
    for fi31 := 0; fi31 < 9; fi31++ {
        delete(v4, v1)
        if (!v27) {
            for fi32 := 0; fi32 < 1; fi32++ {
                switch fi32 {
                case 128:
                    v33 := GenS3{F0: 255}
                    _ = v33
                case 635:
                    fmt.Println(v28, v27, v27)
                }
                v34 := GenS4{F0: 1}
                _ = fi32
                _ = v34
            }
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
