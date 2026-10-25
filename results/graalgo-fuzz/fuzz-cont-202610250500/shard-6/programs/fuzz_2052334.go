package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 uint64, p1 int64) int64 {
    if ((p0 | (^(p0))) > (p0 ^ 42)) {
        var v5 uint8 = 128
        switch "" {
        case "hello":
            v5 = uint8(p0)
        }
        _ = v5
    }
    var v6 uint = 2
    _ = p0
    _ = p1
    _ = v6
    return (-(p1))
}

func main() {
    var v1 int8 = 100
    var v2 int = -1
    var v3 uint16 = 248
    v4 := make(map[int]string)
    v3 += 1
    v7 := []int{127, (v2 % ((804) | 1)), (v2 - -9223372036854775808)}
    v8 := GenS1{F0: 128}
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int16 = -296
        _ = fi9
        _ = v10
    }
    var v11 float64 = (0.0 * 0.0)
    var v12 bool = (1e10 >= (3.14159 - 0.1))
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch (v2 << ((0) & 63)) {
    case -814:
        v3 += (v3 | v3)
    case 32767:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case -985:
        v13 := GenS2{F0: "", F1: (v2 % ((v2) | 1))}
        _ = v13
    }
    delete(v4, v2)
    for fi14 := 0; fi14 < 20; fi14++ {
        v4[v2] = ""
        if true {
            var v15 uint8 = 0
            var v16 int16 = -32768
            for fi17 := 0; fi17 < 17; fi17++ {
                fmt.Println(v2, v11)
                _ = fi17
            }
            _ = v15
            _ = v16
        } else {
            fmt.Println(v11, fi14, v2)
            fmt.Println(v2, v1)
        }
        _ = fi14
    }
    v7 = append(v7, v2)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v12)
    switch "foo bar" {
    case "a":
        if (1 < 1) {
            var v18 uint64 = 784
            v2 = (v2 << ((2) & 63))
            _ = v18
        }
    }
    if true {
        for fi19 := 0; fi19 < 20; fi19++ {
            switch fi19 {
            case 2147483647:
                fmt.Println(v12, v1, v2)
            case -318:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            _ = fi19
        }
        v20 := GenS1{F0: 255}
        if ((v2 / ((462) | 1)) != (v2 >> ((321) & 63))) {
            var v21 uint = 42
            if false {
                v22 := GenS3{F0: (v12 && (!true)), F1: 1, F2: "hello"}
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int16 = -606
                    v11 = v11
                    v4[v2] = ""
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
            _ = v21
        } else {
            var v25 uint16 = (v3 % ((v3) | 1))
            if (42 > 65535) {
                v1 = (v1 ^ (v1 | 42))
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if true {
                    v2 += -626
                    var v26 string = ""
                    _ = v26
                }
            }
            var v27 float32 = 0.1
            _ = v25
            _ = v27
        }
        _ = v20
    }
    v2 = (v2 & v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
