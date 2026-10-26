package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint32, p1 int32) uint8 {
    v7 := []int{559, 270, 507}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = 1
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 37
}
func genFn1(p0 int16, p1 uint8) int16 {
    if (false && (true && true)) {
        var v10 float32 = float32(0.0)
        var v11 string = "a"
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 string = ""
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    return -308
}

func main() {
    var v1 float32 = (1e10 - 0.0)
    var v2 int32 = 40
    var v3 float32 = 0.0
    var v4 int8 = 0
    var v5 uint8 = 166
    v6 := make(map[string]bool)
    switch 815 {
    case 0:
        v14 := []int{4294967295, 32768}
        _ = v14
    }
    _ = len(v6)
    if (true && true) {
        if (("" + (("a" + "foo bar") + "")) > "a") {
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 int8 = -128
                _ = fi15
                _ = v16
            }
        }
    } else {
        v17 := GenS1{F0: -589}
        _ = v17
    }
    fmt.Println(v2, v5, v3)
    var v18 int16 = 32767
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 float32 = 1.5
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS1{F0: (^(fi21))}
        for fi23 := 0; fi23 < 6; fi23++ {
            if false {
                var v24 int = fi23
                v4 = (v4 % (((v4 ^ 41)) | 1))
                switch fi21 {
                case 128:
                    v24 -= fi23
                }
                _ = v24
            }
            v25 := GenS1{F0: int(v4)}
            _ = fi23
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    var v26 int8 = (-(v4))
    if ((("" + (("a" + "hello") + "")) + "foo bar") != "foo bar") {
        v27 := []int{32768, -1}
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
        _ = v27
    }
    var v28 uint8 = (v5 + (v5 % (((v5 >> ((1) & 63))) | 1)))
    v29 := GenS1{F0: -602}
    v4 -= (v4 >> ((2) & 63))
    var v30 int32 = (v2 * v2)
    if ((0 != int8(v18)) || (1 > 2147483647)) {
        v31 := GenS1{F0: 0}
        v32 := GenS1{F0: -865}
        _ = v31
        _ = v32
    } else {
        v33 := GenS2{F0: 1}
        _ = v33
    }
    var v34 uint64 = 800
    v35 := GenS3{F0: 65535, F1: 255}
    var v36 int8 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
