package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int, p1 uint64) int16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(p1, p1)
        p0 -= fi5
        _ = fi5
    }
    v6 := GenS1{F0: 0, F1: 2147483647, F2: 0}
    v7 := []int{(^(p0)), -880, 42}
    p0 = 0
    v8 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int8) uint16 {
    p0 = (p0 << ((2) & 63))
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint = 4294967295
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 1
    v12 := GenS1{F0: -841, F1: 416, F2: 365}
    _ = p0
    _ = v11
    _ = v12
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 159
    var v3 int8 = 100
    var v4 float32 = (((1e10 / float32(float64(((1.5 / 3.14159) - 0.1)))) - 1e10) - 1e10)
    v13 := GenS2{F0: v3}
    var v14 uint64 = 485
    var v15 float64 = 0.1
    _ = len(v1)
    var v16 string = (("a" + "\n\t") + ((("hello" + "a") + "foo bar") + ("hello" + "a")))
    var v17 float32 = (v4 / 1e10)
    if true {
        var v18 float32 = float32(1.5)
        _ = v18
    } else {
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := v1[v16]
            _ = fi19
            _ = v20
        }
    }
    v3 -= v3
    v21 := []int{}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 bool = false
        _ = len(v1)
        if v23 {
            var v24 uint16 = 755
            _ = v24
        } else {
            var v25 int = fi22
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    switch ((v16 + (((v16 + v16) + "\n\t") + "")) + (v16 + "foo bar")) {
    case "hello":
        if false {
            if false {
                v2 = (v2 * v2)
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 int16 = 1000
                fmt.Println(v16, v14)
                if false {
                    v28 := GenS2{F0: (v3 * v3)}
                    if len(v21) > 0 {
                        v21[0] = fi26
                    }
                    _ = v28
                } else {
                    var v29 int64 = 0
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
        } else {
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 uint32 = 2147483648
                var v32 int32 = 0
                _ = fi30
                _ = v31
                _ = v32
            }
            v16 = (((("hello" + ((v16 + (("a" + "a") + "\n\t")) + "\n\t")) + "\n\t") + ((v16 + v16) + v16)) + v16)
        }
    }
    if (((false || true) || true) || true) {
        v33 := GenS3{F0: (v3 / (((v3 - v3)) | 1))}
        _ = v33
    } else {
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
