package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
type GenS3 struct { F0 uint64; F1 int32; F2 int8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint8) uint16 {
    for fi7 := 0; fi7 < 10; fi7++ {
        if false {
            v8 := GenS1{F0: -9223372036854775808, F1: "", F2: 685}
            var v9 int64 = 0
            fmt.Println(v9)
            _ = v8
            _ = v9
        } else {
            fmt.Println(fi7, p0)
            var v10 int16 = 253
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: 535, F1: ("\n\t" + (("hello" + "a") + "a")), F2: 65535}
            if false {
                var v13 bool = true
                _ = v13
            } else {
                v14 := []int{fi7}
                p0 = p0
                v15 := GenS2{F0: p0, F1: 4294967295, F2: 1}
                _ = v14
                _ = v15
            }
            p0 -= (p0 << ((290) & 63))
            p0 = 2
            if (2 == 0) {
                v16 := GenS2{F0: 128, F1: 828, F2: (fi7 + fi7)}
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        p0 -= 233
        _ = fi7
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        p0 -= 68
        var v18 int32 = 1
        var v19 uint = 1
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if (((!false) || true) && true) {
            var v21 int = fi20
            _ = v21
        }
        if (42 >= 0) {
            for fi22 := 0; fi22 < 10; fi22++ {
                if true {
                    p0 = p0
                    p0 += (p0 << ((0) & 63))
                    fmt.Println(p0, fi20, p0)
                }
                _ = fi22
            }
            v23 := GenS3{F0: 2, F1: 1, F2: 0}
            for fi24 := 0; fi24 < 2; fi24++ {
                p0 = p0
                v25 := GenS1{F0: (fi20 | fi24), F1: ("foo bar" + "hello"), F2: 255}
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        _ = fi20
    }
    v26 := GenS1{F0: -1, F1: "a", F2: 1}
    v27 := GenS4{F0: (p0 + p0)}
    _ = p0
    _ = v26
    _ = v27
    return 42
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 int8 = 127
    v5 := []uint32{4294967295, 572, 931, 65535}
    var v6 int32 = 737
    v28 := GenS3{F0: 1, F1: v6, F2: v4}
    for fi29 := 0; fi29 < 7; fi29++ {
        fmt.Println(v4)
        var v30 uint32 = 65535
        _ = fi29
        _ = v30
    }
    switch "\n\t" {
    case "":
        var v31 int16 = 0
        _ = v31
    }
    var v32 string = "a"
    v6 += v6
    var v33 uint = 0
    for fi34 := 0; fi34 < 16; fi34++ {
        v35 := GenS1{F0: -9223372036854775808, F1: (v32 + v32), F2: 0}
        v32 = (v32 + (v32 + ("\n\t" + v32)))
        _ = fi34
        _ = v35
    }
    var v36 int16 = 0
    v6 += v6
    for fi37 := 0; fi37 < 11; fi37++ {
        v6 = 760
        _ = fi37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
