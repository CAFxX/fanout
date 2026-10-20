package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint8 = 0
    var v2 int16 = 42
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    for fi5 := 0; fi5 < 15; fi5++ {
        if (float32((3.14159 - float64(-2.25))) == 1e10) {
            for fi6 := 0; fi6 < 1; fi6++ {
                v2 = (v2 / ((v2) | 1))
                switch "a" {
                case "foo bar":
                    var v7 uint64 = 18446744073709551615
                    _ = v7
                default:
                    delete(v3, (fi6 << ((4294967295) & 63)))
                }
                _ = fi6
            }
            fmt.Println(v1)
            v8 := v3[(fi5 + (^(fi5)))]
            _ = v8
        } else {
            v2 -= (v2 | v2)
        }
        v1 = 128
        v9 := []int{-1, (fi5 % ((128) | 1)), (-(fi5)), fi5}
        _ = fi5
        _ = v9
    }
    switch ("foo bar" + "hello") {
    case "":
        fmt.Println(v2)
    default:
        switch (("a" + ("" + "hello")) + "foo bar") {
        case "":
            var v10 uint16 = 0
            _ = v10
        }
    }
    v2 -= v2
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            switch (fi11 / ((fi11) | 1)) {
            case -1:
                v13 := GenS1{F0: -857}
                v2 -= v2
                if (((true && false) && (128 >= (fi11 / (((fi11 ^ -216)) | 1)))) && true) {
                    fmt.Println(v1)
                    v3[fi11] = ("\n\t" + "foo bar")
                }
                _ = v13
            default:
                fmt.Println(v1)
            }
            v2 = 0
            v14 := GenS2{F0: fi12}
            _ = fi12
            _ = v14
        }
        v1 = v1
        switch (fi11 - 128) {
        case 2147483647:
            fmt.Println(v2, fi11)
        case 127:
            v15 := v4["foo bar"]
            _ = v15
        default:
            v16 := GenS1{F0: -2147483648}
            _ = v16
        }
        _ = fi11
    }
    var v17 uint16 = 42
    for fi18 := 0; fi18 < 2; fi18++ {
        switch fi18 {
        case 368:
            if true {
                for fi19 := 0; fi19 < 2; fi19++ {
                    v20 := GenS3{F0: "a", F1: 42, F2: 0}
                    var v21 int8 = 68
                    var v22 uint = 234
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                if (2147483647 >= 42) {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        var v24 bool = (2147483648 >= 1)
                        _ = fi23
                        _ = v24
                    }
                    v2 = 32767
                }
            }
        default:
            var v25 uint32 = 1
            _ = v25
        }
        v1 = v1
        v26 := GenS4{F0: 354, F1: 4294967295}
        _ = fi18
        _ = v26
    }
    v27 := []int{}
    if ("hello" <= "a") {
        v1 = v1
        if ("" > "\n\t") {
            v1 -= 2
        } else {
            v28 := GenS2{F0: -498}
            delete(v3, 32768)
            _ = v28
        }
    } else {
        var v29 int16 = -32768
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
