package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int16; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64) uint64 {
    var v7 uint64 = 1
    v8 := GenS1{F0: 1, F1: 42}
    fmt.Println(v7, v7)
    v9 := []int{2147483647}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 168
}
func genFn1(p0 uint32, p1 uint) int {
    v10 := GenS1{F0: uint8(p0), F1: p1}
    fmt.Println(p1)
    v11 := []int{}
    fmt.Println(p0, p0)
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int64 = 42
        var v14 int32 = 42
        if len(v11) > 0 {
            v11[0] = -362
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 string = ("foo bar" + "\n\t")
    v5 := make(map[uint32]string)
    v6 := []int{127, 127, -218}
    switch v4 {
    case "hello":
        var v15 int = 2147483647
        _ = v15
    }
    v4 = (v4 + ((v4 + v4) + v4))
    switch (("hello" + v4) + (v4 + (v4 + v4))) {
    case "":
        var v16 int = 311
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint64 = 18446744073709551615
        v1 += v1
        v19 := GenS1{F0: (v1 | 0), F1: 4294967295}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 1, F1: -1, F2: 4294967295}
    v21 := GenS3{F0: 50}
    _ = len(v6)
    v4 = v4
    for fi22 := 0; fi22 < 17; fi22++ {
        if true {
            v23 := GenS2{F0: -615, F1: 9223372036854775807, F2: 743}
            _ = v23
        } else {
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 uint64 = 4294967295
                if (3.14159 != 1e10) {
                    if false {
                        v6 = append(v6, fi22)
                    }
                    v2[genFn1(0, 0)] = 128
                    v1 = 1
                }
                v26 := GenS2{F0: 0, F1: 42, F2: 18446744073709551615}
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        _ = fi22
    }
    if (0 < -72) {
        for fi27 := 0; fi27 < 10; fi27++ {
            if (false || true) {
                fmt.Println(fi27)
                v1 = v1
            }
            if false {
                if false {
                    v1 = v1
                    if len(v6) > 0 {
                        v6[0] = fi27
                    }
                }
                v5[0] = v4
            } else {
                v4 = v4
                switch fi27 {
                case 128:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if false {
                            v4 = "foo bar"
                        }
                        _ = fi28
                    }
                }
                v4 = v4
            }
            _ = fi27
        }
    } else {
        for fi29 := 0; fi29 < 7; fi29++ {
            v30 := GenS1{F0: (v1 % ((v1) | 1)), F1: 885}
            _ = fi29
            _ = v30
        }
        switch 9223372036854775807 {
        case 32767:
            var v31 bool = (false && (4294967295 <= 460))
            _ = v31
        case 32768:
            switch 128 {
            case 466:
                v4 = "\n\t"
            case 127:
                v5[42] = ("a" + v4)
            case 9223372036854775807:
                for fi32 := 0; fi32 < 16; fi32++ {
                    var v33 uint8 = v1
                    _ = fi32
                    _ = v33
                }
            }
        }
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 bool = false
        _ = fi34
        _ = v35
    }
    v4 = v4
    v36 := GenS3{F0: 584}
    var v37 uint = 12
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
