package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -32768
    var v3 int32 = 0
    var v4 uint32 = 1
    var v5 float64 = 0.0
    v3 += (^(v3))
    v6 := GenS1{F0: 32767}
    v2 -= v2
    v7 := GenS2{F0: 2147483647, F1: -128}
    v8 := GenS1{F0: v2}
    switch 128 {
    case 443:
        if (10 > 963) {
            fmt.Println(v5, v4, v2)
        } else {
            fmt.Println(v3, v2)
            var v9 int64 = 649
            v10 := []int{32767, 32767}
            _ = v10
            _ = v9
        }
    case 128:
        for fi11 := 0; fi11 < 19; fi11++ {
            v5 -= v5
            for fi12 := 0; fi12 < 2; fi12++ {
                switch (fi11 & 543) {
                case -685:
                    var v13 uint16 = 924
                    _ = v13
                case 128:
                    fmt.Println(v3, v3)
                case 127:
                    delete(v1, (fi11 | fi11))
                }
                _ = fi12
            }
            _ = fi11
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v2 += 42
        var v15 bool = false
        _ = fi14
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 2; fi16++ {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 uint32 = (v4 * 2)
                v19 := []int{fi17, fi17}
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
    } else {
        fmt.Println(v3)
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 int16 = v2
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v2)
                v2 -= (v2 << ((42) & 63))
                var v23 uint = 42
                var v24 uint = v23
                _ = fi22
                _ = v23
                _ = v24
            }
            if true {
                fmt.Println(v5)
            } else {
                switch fi20 {
                case 846:
                    for fi25 := 0; fi25 < 10; fi25++ {
                        var v26 uint = uint(v4)
                        var v27 uint = v26
                        var v28 int32 = 0
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
                if (!false) {
                    var v29 uint16 = 8
                    _ = v29
                }
                v30 := GenS3{F0: (fi20 ^ fi20), F1: ""}
                _ = v30
            }
            _ = fi20
            _ = v21
        }
    }
    var v31 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
