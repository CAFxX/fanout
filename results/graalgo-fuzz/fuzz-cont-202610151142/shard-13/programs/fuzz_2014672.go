package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
func genFn0(p0 int64, p1 uint16) uint32 {
    p0 = p0
    if (42 > 32768) {
        var v7 int64 = p0
        _ = v7
    } else {
        switch "foo bar" {
        case "foo bar":
            var v8 int = 32768
            _ = v8
        case "hello":
            switch 32768 {
            case 1:
                v9 := GenS1{F0: 1}
                _ = v9
            case 32768:
                var v10 uint8 = 27
                _ = v10
            case 119:
                fmt.Println(p1)
            }
        case "":
            for fi11 := 0; fi11 < 6; fi11++ {
                for fi12 := 0; fi12 < 5; fi12++ {
                    p0 = p0
                    v13 := []int{}
                    _ = fi12
                    _ = v13
                }
                var v14 int32 = int32(p0)
                _ = fi11
                _ = v14
            }
        }
        v15 := GenS1{F0: -1}
        if false {
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(p1, fi16)
                v17 := GenS1{F0: (fi16 | (fi16 % ((32767) | 1)))}
                _ = fi16
                _ = v17
            }
            switch 9223372036854775807 {
            case -2147483648:
                for fi18 := 0; fi18 < 20; fi18++ {
                    if false {
                        if false {
                            p1 = 255
                            var v19 uint = 18446744073709551615
                            _ = v19
                        }
                        v20 := []int{(-(fi18)), (fi18 >> ((2) & 63)), (^(fi18))}
                        v21 := GenS1{F0: -1}
                        _ = v20
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 uint16 = 32767
                            _ = fi22
                            _ = v23
                        }
                        p1 -= p1
                    }
                    p0 = -837
                    _ = fi18
                }
            default:
                if (!false) {
                    var v24 float64 = 0.0
                    for fi25 := 0; fi25 < 2; fi25++ {
                        p1 += (p1 ^ 255)
                        v26 := GenS2{F0: (p1 | p1), F1: 128}
                        _ = fi25
                        _ = v26
                    }
                    var v27 bool = true
                    _ = v24
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 2; fi28++ {
                p0 += (-(p0))
                _ = fi28
            }
        }
        _ = v15
    }
    v29 := GenS3{F0: "", F1: 42}
    var v30 uint8 = 2
    v30 -= v30
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    return 213
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]string)
    var v3 uint64 = 42
    v4 := make(map[int]int)
    var v5 float32 = 0.0
    v6 := []string{("\n\t" + "foo bar"), "hello", "a"}
    v5 = v5
    v5 -= float32(0.1)
    var v31 uint64 = v3
    _ = len(v2)
    v2["a"] = ""
    v5 = (((0.0 + float32((1.5 / float64(-2.25)))) + v5) - float32(3.14159))
    var v32 uint32 = 42
    v33 := GenS1{F0: -530}
    v3 = (v31 % ((v3) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
