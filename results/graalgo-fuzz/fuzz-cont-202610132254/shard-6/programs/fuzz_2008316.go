package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
func genFn0(p0 string) int8 {
    switch ((p0 + (p0 + "\n\t")) + ("a" + ("\n\t" + "foo bar"))) {
    case "foo bar":
        var v5 float32 = 1.5
        _ = v5
    case "a":
        var v6 float32 = -2.25
        _ = v6
    case "":
        switch -1 {
        case 4294967295:
            p0 = (p0 + "hello")
        case 32767:
            p0 = "foo bar"
        default:
            for fi7 := 0; fi7 < 1; fi7++ {
                v8 := []int{fi7}
                var v9 int16 = 42
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    }
    switch p0 {
    case "hello":
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int8 = 1
                    v12 = (v12 + -102)
                    _ = fi11
                    _ = v12
                }
                p0 = ("" + (("\n\t" + (p0 + p0)) + ("foo bar" + "\n\t")))
                p0 = p0
                _ = fi10
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS1{F0: 4294967295, F1: 1, F2: 939}
                v15 := []int{-2147483648, (fi13 << ((42) & 63)), (fi13 >> ((631) & 63)), fi13}
                _ = fi13
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 bool = true
                fmt.Println(fi16, v17)
                _ = fi16
                _ = v17
            }
        } else {
            if (2 != 440) {
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 uint64 = 0
                    p0 = "foo bar"
                    p0 = ("hello" + ("" + "\n\t"))
                    _ = fi18
                    _ = v19
                }
                p0 = p0
            } else {
                v20 := GenS2{F0: 255, F1: 6, F2: 100}
                fmt.Println(p0)
                _ = v20
            }
            fmt.Println(p0)
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(p0)
                _ = fi21
            }
        }
    default:
        var v22 bool = (0.0 == 0.1)
        _ = v22
    }
    var v23 uint64 = 42
    _ = p0
    _ = v23
    return -51
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 uint32 = 594
    _ = len(v3)
    if (!false) {
        fmt.Println(v4)
        var v24 uint8 = 127
        var v25 bool = ((0.0 < (((-2.25 + 1e10) / 1e10) - (((0.1 / 1.0) - (0.0 + 0.1)) / 1.0))) || (!(v4 != 0)))
        _ = v24
        _ = v25
    }
    if (-2147483648 > -1) {
        var v26 uint64 = 2
        _ = v26
    }
    v27 := GenS1{F0: 0, F1: 0, F2: 255}
    var v28 int = -545
    var v29 int16 = 42
    for fi30 := 0; fi30 < 20; fi30++ {
        v31 := GenS2{F0: 1, F1: 164, F2: 0}
        _ = fi30
        _ = v31
    }
    var v32 int32 = 683
    fmt.Println(v29, v4)
    if true {
        switch ("\n\t" + "\n\t") {
        case "a":
            v29 = v29
        case "\n\t":
            v33 := GenS3{F0: v28}
            _ = v33
        case "foo bar":
            if ("" > "\n\t") {
                v34 := GenS2{F0: 395, F1: 1, F2: 15}
                v32 = (v32 - (v32 & v32))
                _ = v34
            }
        default:
            if (true && (42 == 2)) {
                _ = len(v3)
                var v35 uint32 = 2147483648
                _ = v35
            }
        }
        v36 := []int{int(v32), (v28 / ((v28) | 1)), 32767}
        v37 := GenS3{F0: (v28 >> ((2) & 63))}
        _ = v36
        _ = v37
    }
    v38 := GenS4{F0: (v28 % ((v28) | 1))}
    var v39 int64 = int64(v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
}
