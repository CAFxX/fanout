package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 int32 }
type GenS4 struct { F0 int }
func genFn0(p0 uint64, p1 int) uint {
    fmt.Println(p0)
    v7 := GenS1{F0: -992, F1: ""}
    v8 := GenS1{F0: 32767, F1: "foo bar"}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int64, p1 uint32, p2 int16) uint16 {
    if false {
        if (true && false) {
            v9 := GenS1{F0: -32768, F1: "foo bar"}
            fmt.Println(p2)
            _ = v9
        } else {
            var v10 int32 = 0
            v11 := []int{279}
            _ = v10
            _ = v11
        }
        if (p2 > 246) {
            var v12 int32 = 2147483647
            v13 := GenS2{F0: 0, F1: 4294967295}
            _ = v12
            _ = v13
        }
        var v14 int64 = p0
        _ = v14
    } else {
        if true {
            var v15 int16 = (^(p2))
            v16 := GenS2{F0: -109, F1: p1}
            for fi17 := 0; fi17 < 19; fi17++ {
                switch "foo bar" {
                case "\n\t":
                    if false {
                        v18 := []int{int(v15), (^(fi17))}
                        fmt.Println(p2, fi17, p2)
                        _ = v18
                    } else {
                        p0 += p0
                        v19 := GenS3{F0: fi17, F1: 2, F2: 722}
                        _ = v19
                    }
                }
                _ = fi17
            }
            _ = v15
            _ = v16
        } else {
            var v20 int32 = 2147483647
            v20 = v20
            _ = v20
        }
    }
    var v21 string = ("foo bar" + "hello")
    for fi22 := 0; fi22 < 14; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 float64 = float64(0.0)
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 2
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    v5 := []uint32{0}
    var v6 uint16 = 32767
    fmt.Println(v1, v6)
    var v25 uint16 = v6
    _ = len(v4)
    switch 32767 {
    case -567:
        var v26 string = "a"
        _ = v26
    default:
        var v27 float32 = (1.5 - float32(-2.25))
        _ = v27
    }
    if false {
        var v28 int64 = 0
        var v29 bool = (!false)
        v29 = (837 != 759)
        _ = v28
        _ = v29
    }
    if (127 <= -1) {
        var v30 int8 = -50
        var v31 uint16 = (v25 % ((v25) | 1))
        _ = v30
        _ = v31
    }
    var v32 uint16 = (v25 / (((v25 * 88)) | 1))
    v32 += (v25 >> ((2) & 63))
    if (!(false && (!true))) {
        if (!(genFn1(900, 42, -986) > (v32 >> ((552) & 63)))) {
            var v33 uint8 = 2
            v5 = append(v5, 875)
            _ = v33
        } else {
            for fi34 := 0; fi34 < 13; fi34++ {
                _ = len(v5)
                _ = fi34
            }
            fmt.Println(v6)
        }
    } else {
        var v35 float64 = -2.25
        _ = v35
    }
    var v36 uint8 = 1
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 12; fi38++ {
            v39 := GenS4{F0: fi38}
            fmt.Println(v1)
            var v40 string = ""
            _ = fi38
            _ = v39
            _ = v40
        }
        switch (fi37 % ((int(v36)) | 1)) {
        case 0:
            v1 += v1
        case 2147483647:
            if (false || (!((1e10 >= -2.25) && (!false)))) {
                v25 += (^(v6))
                if true {
                    var v41 float64 = (float64(float32((1.5 * 1.5))) * 1.5)
                    _ = v41
                }
                v42 := GenS1{F0: -723, F1: "\n\t"}
                _ = v42
            }
        default:
            switch (fi37 - fi37) {
            case 4294967295:
                fmt.Println(v36)
            case 0:
                var v43 float64 = 3.14159
                _ = v43
            case -519:
                v44 := GenS4{F0: 32768}
                _ = v44
            }
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
