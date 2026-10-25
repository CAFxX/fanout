package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 0.0
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]string)
    var v6 uint32 = 2
    switch (v3 - (v3 | (v3 % ((v3) | 1)))) {
    case -770:
        _ = len(v4)
    case 42:
        v7 := []int{(-(v3)), (^(v3))}
        _ = v7
    }
    var v8 string = "a"
    for fi9 := 0; fi9 < 2; fi9++ {
        if false {
            var v10 int8 = 18
            var v11 int64 = 0
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    v6 = v6
    v12 := GenS1{F0: 0, F1: false}
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 bool = (8 < 42)
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 string = v8
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS2{F0: v6}
        v19 := GenS2{F0: v6}
        for fi20 := 0; fi20 < 12; fi20++ {
            if false {
                var v21 int16 = -711
                var v22 bool = true
                var v23 int64 = 75
                if ((v21 >> ((2) & 63)) == (v21 & v21)) {
                    if (v8 == "") {
                        v24 := []int{(-(fi17)), -165, v3, -142}
                        if v22 {
                            v25 := GenS2{F0: 2}
                            _ = v25
                        } else {
                            v23 = -596
                            switch v8 {
                            case "foo bar":
                                var v26 float32 = v2
                                _ = v26
                            }
                            v27 := v1[(-(fi17))]
                            _ = v27
                        }
                        var v28 bool = v22
                        _ = v24
                        _ = v28
                    } else {
                        v8 = v8
                        v8 = (v8 + v8)
                        var v29 uint16 = 32767
                        _ = v29
                    }
                }
                var v30 int32 = -349
                _ = v21
                _ = v22
                _ = v23
                _ = v30
            } else {
                switch -764 {
                case 2147483647:
                    v31 := GenS3{F0: "hello"}
                    _ = v31
                }
                var v32 uint32 = v6
                switch (v8 + v8) {
                case "":
                    fmt.Println(v32)
                default:
                    v2 = float32((-2.25 - 1e10))
                }
                _ = v32
            }
            v33 := []int{(-(fi17))}
            v33 = append(v33, (fi20 + fi20))
            _ = fi20
            _ = v33
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    var v34 float32 = ((float32((((0.1 + ((0.0 + -2.25) + (0.1 * 1e10))) - ((1e10 + float64(1.5)) - (1.5 + (-2.25 * 0.1)))) / -2.25)) - v2) + (float32((float64(1.5) + 1.5)) / v2))
    delete(v5, v6)
    fmt.Println(v3, v3, v3)
    v3 = (v3 / (((v3 | (v3 % ((v3) | 1)))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
