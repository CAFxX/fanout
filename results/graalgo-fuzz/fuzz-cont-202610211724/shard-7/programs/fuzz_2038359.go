package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0(p0 uint) uint8 {
    switch "foo bar" {
    case "\n\t":
        switch -2147483648 {
        case 127:
            var v7 int8 = 100
            _ = v7
        case 32768:
            v8 := GenS1{F0: -642, F1: 1}
            _ = v8
        default:
            var v9 string = ("a" + "a")
            _ = v9
        }
    case "foo bar":
        if true {
            v10 := []int{728, 128, 32767}
            p0 = (p0 << ((p0) & 63))
            _ = v10
        }
    default:
        var v11 int = 1
        _ = v11
    }
    p0 = (^(p0))
    v12 := GenS1{F0: 127, F1: 0}
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{fi13, (fi13 ^ fi13)}
        if true {
            var v15 string = (("\n\t" + "a") + "hello")
            if true {
                v16 := GenS1{F0: (fi13 & 128), F1: 42}
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 uint32 = 1
                    v19 := GenS1{F0: fi13, F1: 2}
                    var v20 float32 = (1e10 * 1.5)
                    if true {
                        v21 := GenS2{F0: int32(fi13), F1: 0}
                        _ = v21
                    } else {
                        switch "foo bar" {
                        case "":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                if len(v14) > 0 {
                    v14[0] = fi13
                }
                _ = v16
            } else {
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 uint32 = 65535
                    v14 = append(v14, (fi13 | (fi13 ^ fi22)))
                    _ = fi22
                    _ = v23
                }
                v24 := GenS1{F0: fi13, F1: 110}
                _ = v24
            }
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    if (42 == -1) {
        var v25 uint64 = 522
        _ = v25
    }
    _ = p0
    _ = v12
    return 128
}
func genFn1() uint {
    v26 := GenS2{F0: 2147483647, F1: 328}
    if true {
        var v27 int64 = 42
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := GenS3{F0: false}
            fmt.Println(v27)
            v30 := []int{(fi28 & fi28), fi28, fi28, (fi28 & (fi28 >> ((18446744073709551615) & 63)))}
            _ = fi28
            _ = v29
            _ = v30
        }
        v27 += (^(v27))
        _ = v27
    } else {
        var v31 bool = true
        _ = v31
    }
    v32 := []int{}
    for fi33 := 0; fi33 < 9; fi33++ {
        _ = len(v32)
        _ = fi33
    }
    _ = v26
    _ = v32
    return 534
}

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{998, 1, 28, 2147483648}
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v5 float32 = 1e10
    var v6 string = ((v1 + (v1 + ((v1 + "hello") + v1))) + "\n\t")
    var v34 uint16 = 446
    v2 = append(v2, 341)
    _ = len(v3)
    fmt.Println(v34)
    v6 = v1
    v4[4294967295] = 4294967295
    v35 := GenS4{F0: 42, F1: 42, F2: 2}
    var v36 int16 = -746
    if (42 < -1) {
        var v37 bool = true
        switch (v6 + v6) {
        case "":
            var v38 int8 = 0
            _ = v38
        case "foo bar":
            v36 = v36
        }
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
