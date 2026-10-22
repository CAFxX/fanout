package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint32; F2 uint }
func genFn0(p0 uint, p1 int32) uint8 {
    if (false || true) {
        v5 := GenS1{F0: 1, F1: 1, F2: (p0 & 42)}
        var v6 uint = (p0 >> (((p0 >> (((p0 | p0)) & 63))) & 63))
        v7 := []int{2147483647}
        _ = v5
        _ = v6
        _ = v7
    }
    if (!false) {
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 string = "a"
                var v11 uint16 = 936
                if (0.1 > 1e10) {
                    v11 += 2
                    if (!(!true)) {
                        v12 := GenS2{F0: "\n\t", F1: 4}
                        v10 = ((((v10 + (v10 + ("foo bar" + "a"))) + ("foo bar" + v10)) + v10) + "\n\t")
                        var v13 int16 = 1000
                        _ = v12
                        _ = v13
                    }
                    var v14 float64 = (-2.25 - 0.1)
                    _ = v14
                } else {
                    var v15 bool = (p1 <= p1)
                    var v16 int8 = int8(p1)
                    _ = v15
                    _ = v16
                }
                v11 -= (v11 - v11)
                switch (("" + v10) + "\n\t") {
                case "":
                    switch "a" {
                    case "":
                        if true {
                            v17 := []int{-282, fi9, -782, fi8}
                            fmt.Println(p1, p1, p0)
                            _ = v17
                        }
                    case "foo bar":
                        var v18 float32 = (1.5 + (0.0 + ((-2.25 - -2.25) - 0.0)))
                        _ = v18
                    }
                case "\n\t":
                    var v19 bool = (42 > 628)
                    _ = v19
                }
                p0 += 42
                var v20 int32 = (p1 >> ((p0) & 63))
                if false {
                    v21 := GenS1{F0: 42, F1: 2, F2: 2}
                    v22 := GenS3{F0: 255}
                    v20 += v20
                    _ = v21
                    _ = v22
                } else {
                    var v23 int8 = 127
                    var v24 float64 = (1.5 + 1e10)
                    _ = v23
                    _ = v24
                }
                _ = fi9
                _ = v10
                _ = v11
                _ = v20
            }
            v25 := GenS2{F0: ("a" + ((("hello" + ("\n\t" + (("hello" + "") + "hello"))) + "\n\t") + "foo bar")), F1: 127}
            _ = fi8
            _ = v25
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{}
    v3 := []string{("" + "a"), "\n\t", "\n\t"}
    var v4 float32 = ((-2.25 + float32(0.1)) - float32((3.14159 / (((3.14159 - 1e10) - 0.0) + 0.0))))
    fmt.Println(v1)
    for fi26 := 0; fi26 < 12; fi26++ {
        v1 = 256
        _ = fi26
    }
    if true {
        var v27 int16 = 42
        v28 := GenS4{F0: ("a" + ""), F1: 558, F2: 0}
        _ = v27
        _ = v28
    }
    v4 -= v4
    var v29 uint64 = (v1 << ((486) & 63))
    var v30 uint8 = 124
    v3 = append(v3, "foo bar")
    var v31 uint64 = (v29 - 778)
    v31 -= (v29 ^ v31)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
