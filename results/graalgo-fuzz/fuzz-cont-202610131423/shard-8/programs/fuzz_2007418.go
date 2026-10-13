package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
func genFn0() int32 {
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 11; fi7++ {
        if (32767 != -32768) {
            v8 := GenS1{F0: 791, F1: -9223372036854775808, F2: 1}
            _ = v8
        }
        if false {
            v9 := []int{fi7, (fi7 + 42), (fi7 << ((626) & 63)), (^(fi7))}
            _ = v9
        }
        v10 := []int{(fi7 >> ((804) & 63)), fi7}
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: 883, F1: 0, F2: 1000}
        var v13 bool = true
        v6 = v6
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 255, F1: -516, F2: -1}
    _ = v14
    _ = v6
    return -87
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -242
    v3 := []int{v2, 128}
    v4 := make(map[uint32]string)
    var v5 uint64 = 42
    for fi15 := 0; fi15 < 9; fi15++ {
        if false {
            for fi16 := 0; fi16 < 9; fi16++ {
                v17 := GenS2{F0: 42}
                _ = fi16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            v2 = fi18
            v19 := GenS3{F0: 654}
            v20 := GenS2{F0: -1}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi15
    }
    var v21 float32 = (-2.25 + -2.25)
    for fi22 := 0; fi22 < 8; fi22++ {
        if len(v3) > 0 {
            v3[0] = v2
        }
        v23 := GenS1{F0: 733, F1: -1, F2: 42}
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: 42}
    v25 := GenS3{F0: 2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 uint64 = v5
            fmt.Println(v27, v2, v21)
            _ = fi26
            _ = v27
        }
    }
    if true {
        var v28 uint32 = 560
        var v29 uint16 = 65535
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: 4294967295, F1: 0, F2: 903}
    var v31 string = "a"
    switch ((v31 + v31) + "hello") {
    case "hello":
        v32 := GenS4{F0: v2}
        _ = v32
    case "foo bar":
        for fi33 := 0; fi33 < 3; fi33++ {
            if false {
                v31 = ((v31 + (v31 + v31)) + v31)
            }
            fmt.Println(v21, v21)
            v2 -= (fi33 | v2)
            _ = fi33
        }
    case "\n\t":
        switch ("foo bar" + "a") {
        case "\n\t":
            v2 -= 279
        }
    }
    v5 = 42
    switch v2 {
    case -2147483648:
        v2 -= v2
    case -85:
        var v34 int32 = genFn0()
        _ = v34
    case -635:
        var v35 bool = (255 <= 954)
        _ = v35
    default:
        var v36 uint = 2
        _ = v36
    }
    var v37 int16 = 676
    var v38 int64 = 0
    v2 -= v2
    v39 := GenS2{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
