package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 uint, p2 uint) int {
    v6 := []int{}
    v7 := GenS1{F0: uint16(p2), F1: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{("a" + "a"), "\n\t", ("\n\t" + ("\n\t" + "")), "hello"}
    v3 := make(map[string]bool)
    var v4 uint16 = 1
    v5 := []int{-1, -1, 128}
    v4 -= (v4 << ((199) & 63))
    if false {
        if false {
            var v8 uint8 = 107
            if false {
                v9 := GenS1{F0: (v4 | v4), F1: "foo bar"}
                _ = v9
            } else {
                var v10 int32 = -627
                _ = v10
            }
            var v11 uint8 = v8
            _ = v11
            _ = v8
        } else {
            if (59 <= 522) {
                fmt.Println(v4)
                fmt.Println(v4)
                v4 += 485
            } else {
                v4 = v4
                v4 -= 2
            }
            var v12 int64 = 0
            _ = v12
        }
        if (false && true) {
            for fi13 := 0; fi13 < 6; fi13++ {
                for fi14 := 0; fi14 < 19; fi14++ {
                    var v15 float64 = 0.1
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            var v16 int8 = 127
            _ = v16
        }
        v4 += (v4 & (v4 << ((1) & 63)))
    } else {
        var v17 uint8 = 186
        _ = v17
    }
    var v18 int32 = 1
    var v19 string = ("\n\t" + ("" + ""))
    var v20 uint64 = 4294967295
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            _ = len(v3)
            switch (v19 + v19) {
            case "hello":
                var v23 string = v19
                _ = v23
            case "\n\t":
                if true {
                    v24 := GenS2{F0: 0}
                    _ = v24
                }
            case "":
                var v25 uint64 = v20
                _ = v25
            }
            fmt.Println(fi22)
            v26 := GenS1{F0: (v4 | (v4 >> ((42) & 63))), F1: v19}
            _ = fi22
            _ = v26
        }
        v18 = -1
        _ = fi21
    }
    _ = len(v2)
    v20 = (v20 ^ v20)
    fmt.Println(v20, v20, v20)
    v18 = v18
    switch ("foo bar" + "a") {
    case "\n\t":
        var v27 string = "foo bar"
        _ = v27
    case "hello":
        v18 = (v18 % (((-(v18))) | 1))
    case "a":
        if (false && true) {
            v19 = v19
            var v28 uint16 = v4
            _ = v28
        }
    default:
        v20 = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
