package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -1
    v3 := make(map[string]bool)
    var v4 uint16 = 32767
    v2 = v2
    v5 := GenS1{F0: 534}
    _ = len(v1)
    v2 -= v2
    v6 := GenS1{F0: 0}
    if true {
        if ((0.1 / 1.5) > (3.14159 - (((3.14159 / 1e10) * (3.14159 * -2.25)) + float32((1e10 + 0.0))))) {
            v7 := GenS1{F0: 42}
            v2 = v2
            v8 := []int{1, 42, 1, 4294967295}
            _ = v7
            _ = v8
        } else {
            var v9 int8 = (v2 ^ (v2 / (((^(v2))) | 1)))
            v10 := GenS2{F0: 128}
            v9 += v9
            _ = v10
            _ = v9
        }
        v2 = v2
    } else {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 float32 = ((0.0 / 3.14159) + ((0.0 + (1.5 - (0.0 * (1e10 - -2.25)))) + 1e10))
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 string = ("foo bar" + "hello")
                switch fi11 {
                case 931:
                    if true {
                        if (!(v4 != (v4 % (((v4 << ((65535) & 63))) | 1)))) {
                            if (254 != 127) {
                                fmt.Println(v4, v12, v4)
                                v14 = ((v14 + v14) + (("\n\t" + v14) + "hello"))
                                v15 := GenS2{F0: fi11}
                                if true {
                                    v4 = 33
                                    var v16 int32 = -1
                                    _ = v16
                                } else {
                                    switch ("foo bar" + v14) {
                                    case "foo bar":
                                        switch (("\n\t" + (v14 + "foo bar")) + "a") {
                                        case "":
                                            v14 = (((v14 + v14) + v14) + ((((((v14 + "") + v14) + v14) + v14) + v14) + v14))
                                        }
                                    default:
                                        var v17 float32 = 0.1
                                        _ = v17
                                    }
                                }
                                _ = v15
                            }
                        }
                        v18 := GenS1{F0: 0}
                        v19 := GenS1{F0: 535}
                        _ = v18
                        _ = v19
                    } else {
                        v12 = ((v12 / (0.1 + v12)) + (((v12 - float32(0.1)) / float32(3.14159)) + ((((0.1 * v12) + float32(1e10)) / (v12 - ((1e10 * ((float32(1e10) / -2.25) + ((1e10 + 0.1) / (1.5 * float32(3.14159))))) / v12))) * (v12 / v12))))
                    }
                default:
                    var v20 int = fi11
                    _ = v20
                }
                fmt.Println(fi13, fi13, v2)
                _ = fi13
                _ = v14
            }
            var v21 int8 = (^(v2))
            _ = fi11
            _ = v12
            _ = v21
        }
        v22 := GenS2{F0: -45}
        var v23 float32 = 0.1
        _ = v22
        _ = v23
    }
    v24 := []int{-706}
    fmt.Println(v2, v2)
    fmt.Println(v2, v4)
    fmt.Println(v4)
    for fi25 := 0; fi25 < 7; fi25++ {
        v3[(("a" + "hello") + "a")] = (512 >= 2147483647)
        _ = fi25
    }
    fmt.Println(v4, v2)
    fmt.Println(v4, v2)
    v26 := GenS3{F0: 1000, F1: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
