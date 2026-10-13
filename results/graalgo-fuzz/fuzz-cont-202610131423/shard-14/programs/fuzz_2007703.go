package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 string }
func genFn0() uint32 {
    v7 := GenS1{F0: 2147483647, F1: 0}
    var v8 uint64 = 100
    switch "hello" {
    case "hello":
        v9 := GenS2{F0: 42, F1: 42, F2: "foo bar"}
        _ = v9
    case "a":
        var v10 int16 = -122
        _ = v10
    case "":
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint8 = 2
            v13 := []int{(fi11 * (fi11 * fi11))}
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        var v14 bool = false
        _ = v14
    }
    _ = v7
    _ = v8
    return 285
}

func main() {
    var v1 int = 696
    v2 := make(map[uint32]string)
    v3 := []uint32{977}
    v4 := []string{"a", (((("\n\t" + ("hello" + "\n\t")) + (("foo bar" + "hello") + "foo bar")) + "hello") + "\n\t")}
    var v5 float64 = 0.1
    v6 := make(map[int]int)
    fmt.Println(v5)
    var v15 uint64 = 732
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            if true {
                v1 += fi18
                var v20 uint = 65535
                _ = v20
            }
            _ = fi19
        }
        _ = fi18
    }
    v15 = v15
    v15 = (v15 * (v15 * (v15 / ((42) | 1))))
    _ = len(v6)
    for fi21 := 0; fi21 < 8; fi21++ {
        if false {
            var v22 uint64 = (v15 ^ 18446744073709551615)
            if true {
                if true {
                    v15 += v15
                    v15 -= v22
                    v1 = (fi21 / (((-(v1))) | 1))
                } else {
                    for fi23 := 0; fi23 < 4; fi23++ {
                        var v24 uint32 = 248
                        _ = fi23
                        _ = v24
                    }
                    v22 += (v22 << ((4294967295) & 63))
                    v15 += (^(v22))
                }
                v25 := v6[fi21]
                _ = v25
            }
            v1 = int(v22)
            _ = v22
        } else {
            if false {
                for fi26 := 0; fi26 < 16; fi26++ {
                    v1 -= fi21
                    v5 = v5
                    fmt.Println(v1, v5)
                    _ = fi26
                }
                fmt.Println(v15)
            }
            for fi27 := 0; fi27 < 15; fi27++ {
                if false {
                    v1 = (fi21 | fi21)
                    var v28 uint16 = 255
                    _ = v28
                } else {
                    var v29 uint = 4294967295
                    fmt.Println(v29)
                    var v30 int16 = int16(fi27)
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
        }
        if true {
            var v31 string = "foo bar"
            _ = v31
        } else {
            switch "foo bar" {
            case "foo bar":
                v32 := GenS2{F0: v15, F1: 0, F2: ("a" + "foo bar")}
                _ = v32
            case "\n\t":
                fmt.Println(v15, v15, v15)
            case "hello":
                var v33 uint64 = v15
                _ = v33
            }
        }
        _ = fi21
    }
    v5 += 1.5
    v34 := v2[395]
    fmt.Println(v5)
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS3{F0: "a"}
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
