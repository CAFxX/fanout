package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0(p0 uint32, p1 uint, p2 int32) string {
    p0 = 4294967295
    switch ("\n\t" + (((("hello" + "") + "a") + ("\n\t" + "foo bar")) + "a")) {
    case "\n\t":
        for fi4 := 0; fi4 < 7; fi4++ {
            var v5 int8 = int8(p0)
            v6 := []int{-2147483648, fi4}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi4
            _ = v5
            _ = v6
        }
    case "hello":
        var v7 uint = (p1 >> ((255) & 63))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1(p0 int32) int32 {
    for fi8 := 0; fi8 < 16; fi8++ {
        if (0 == 51) {
            var v9 int64 = 0
            v10 := []int{-981, (fi8 & 1), (-(fi8))}
            _ = v10
            _ = v9
        }
        fmt.Println(p0, fi8)
        _ = fi8
    }
    var v11 int32 = (p0 % (((p0 / ((p0) | 1))) | 1))
    _ = p0
    _ = v11
    return -1
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := GenS1{F0: "a", F1: 500}
        _ = fi12
        _ = v13
    }
    switch (("" + "a") + "hello") {
    case "":
        v1 = v1
    }
    switch "hello" {
    case "a":
        v14 := []int{9223372036854775807, -869}
        _ = v14
    default:
        v15 := GenS1{F0: "", F1: 2}
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if true {
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 string = ""
                v19 := []int{fi17, (^(fi16))}
                v18 = (v18 + v18)
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            delete(v2, (fi16 << ((255) & 63)))
        }
        v1 -= v1
        var v20 uint64 = 4294967295
        _ = fi16
        _ = v20
    }
    var v21 int64 = 1
    delete(v2, -989)
    var v22 uint8 = 255
    v21 = v21
    v21 += (v21 + v21)
    fmt.Println(v22, v21)
    if (-2147483648 != -109) {
        var v23 int64 = (v21 & (v21 / ((v21) | 1)))
        var v24 float64 = (1e10 + float64(-2.25))
        if (v22 == v22) {
            v23 -= v21
            v25 := []int{32767}
            var v26 float64 = float64(3.14159)
            _ = v25
            _ = v26
        }
        _ = v23
        _ = v24
    }
    v1 += 32767
    v27 := []int{9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
}
