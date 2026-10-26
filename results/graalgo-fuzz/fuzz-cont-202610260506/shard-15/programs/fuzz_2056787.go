package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint16 }
func genFn0() int {
    var v7 float32 = 0.0
    v7 = float32((((float64(1.5) / 0.1) / 1.0) * ((((3.14159 + -2.25) + float64((v7 - (float32((0.1 + -2.25)) / v7)))) / (0.1 - float64((float32((1e10 * ((1.5 - 0.0) * -2.25))) - v7)))) - 3.14159)))
    for fi8 := 0; fi8 < 12; fi8++ {
        v7 -= 0.0
        if (!true) {
            if (!true) {
                var v9 uint32 = 4294967295
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 string = "foo bar"
                    v11 = v11
                    if false {
                        if false {
                            v7 -= v7
                        }
                        v12 := []int{(fi10 / ((fi10) | 1)), (fi8 + (^(fi8)))}
                        var v13 int32 = -2147483648
                        _ = v12
                        _ = v13
                    } else {
                        var v14 uint64 = 132
                        v15 := []int{1, 127}
                        _ = v14
                        _ = v15
                    }
                    _ = fi10
                    _ = v11
                }
                v9 = (^(v9))
                _ = v9
            } else {
                v7 += -2.25
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 uint = 18446744073709551615
                    v18 := GenS1{F0: -368}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            if true {
                var v19 uint = 0
                v19 = v19
                _ = v19
            }
            if false {
                v20 := GenS2{F0: 418, F1: 18446744073709551615, F2: -69}
                if (!("hello" <= "foo bar")) {
                    if (!true) {
                        var v21 bool = (-2.25 <= (v7 * float32(float64(-2.25))))
                        v7 = 3.14159
                        _ = v21
                    } else {
                        v7 = v7
                    }
                    switch (("\n\t" + "hello") + ("" + ("hello" + (("foo bar" + "\n\t") + ("" + "foo bar"))))) {
                    case "hello":
                        v7 = v7
                    case "a":
                        v22 := []int{fi8, (fi8 << ((2) & 63)), 32768, fi8}
                        _ = v22
                    }
                } else {
                    v23 := GenS3{F0: 42}
                    _ = v23
                }
                _ = v20
            }
        } else {
            v24 := GenS4{F0: 42, F1: 255, F2: 238}
            _ = v24
        }
        _ = fi8
    }
    _ = v7
    return 128
}

func main() {
    var v1 string = (((("foo bar" + "a") + "") + "") + "a")
    var v2 uint8 = 2
    var v3 uint64 = 0
    var v4 uint32 = 4294967295
    var v5 float32 = float32(-2.25)
    var v6 uint16 = 65535
    for fi25 := 0; fi25 < 7; fi25++ {
        v26 := []int{(fi25 >> ((uint(v3)) & 63))}
        _ = fi25
        _ = v26
    }
    var v27 int64 = -508
    v28 := []int{}
    for fi29 := 0; fi29 < 9; fi29++ {
        v5 = float32(0.0)
        var v30 int16 = -596
        _ = fi29
        _ = v30
    }
    v2 = (v2 >> ((378) & 63))
    v4 = v4
    v3 += (v3 - (v3 * 42))
    v31 := GenS1{F0: (v27 & -1)}
    var v32 int16 = -1
    v33 := GenS3{F0: v4}
    v34 := GenS3{F0: v4}
    var v35 uint32 = v4
    var v36 uint = 65535
    var v37 bool = (-46 <= -1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
