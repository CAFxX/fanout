package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int8; F2 uint32 }
func genFn0(p0 int8) bool {
    fmt.Println(p0)
    p0 += (p0 - (p0 >> ((0) & 63)))
    switch 32767 {
    case 4294967295:
        var v7 uint8 = 127
        _ = v7
    case 9223372036854775807:
        if (9223372036854775807 < 1) {
            fmt.Println(p0)
            v8 := []int{127}
            v9 := GenS1{F0: 630}
            _ = v8
            _ = v9
        } else {
            if ((-(p0)) != 122) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := []int{(fi10 | 127), fi10, fi10}
                    v12 := GenS1{F0: 720}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := []int{(fi13 | (^(fi13))), (fi13 / ((int(p0)) | 1))}
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 9; fi15++ {
                    v16 := GenS2{F0: (("foo bar" + ("foo bar" + "a")) + "hello"), F1: 0}
                    v17 := GenS3{F0: 100, F1: -56}
                    p0 += p0
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            v18 := GenS4{F0: 32768}
            p0 += (p0 & (p0 / ((p0) | 1)))
            _ = v18
        }
    default:
        v19 := []int{42}
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := []int{(fi20 ^ 918), fi20, fi20}
        _ = fi20
        _ = v21
    }
    fmt.Println(p0)
    _ = p0
    return false
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    v3 := make(map[uint32]uint32)
    var v4 string = "a"
    var v5 uint16 = 2
    v6 := make(map[string]string)
    if false {
        v5 = v5
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        if true {
            v5 = (^(v5))
            switch -140 {
            case 1:
                var v23 float64 = (1.5 / 1e10)
                _ = v23
            case -744:
                var v24 float32 = -2.25
                _ = v24
            case -307:
                var v25 uint16 = v5
                _ = v25
            }
        }
        v5 = (v5 / ((v5) | 1))
        _ = fi22
    }
    if (!(true && false)) {
        v26 := GenS4{F0: 783}
        var v27 uint16 = (v5 / (((^(v5))) | 1))
        _ = v26
        _ = v27
    }
    v5 = 1
    v4 = ""
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 += (v5 >> ((4294967295) & 63))
    var v28 string = v4
    if true {
        fmt.Println(v5, v4, v4)
        var v29 int8 = -1
        v30 := v1[768]
        _ = v29
        _ = v30
    }
    delete(v1, 210)
    v31 := GenS5{F0: 1, F1: -1, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
