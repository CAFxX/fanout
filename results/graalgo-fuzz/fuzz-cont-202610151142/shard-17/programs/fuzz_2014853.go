package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
func genFn0() uint16 {
    var v7 uint8 = 80
    fmt.Println(v7)
    if (1 > 127) {
        v8 := []int{-314, 128, -9223372036854775808, -1}
        if true {
            v7 = (v7 / ((v7) | 1))
        } else {
            v7 += v7
            v7 -= (v7 * (^(v7)))
            var v9 uint32 = 20
            _ = v9
        }
        var v10 string = ""
        _ = v10
        _ = v8
    } else {
        v7 = v7
        v7 = 255
    }
    _ = v7
    return 65535
}
func genFn1() int8 {
    var v11 int8 = -128
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v11 += v11
            fmt.Println(v11, fi12)
            _ = fi13
        }
        var v14 uint16 = 55
        _ = fi12
        _ = v14
    }
    if (127 >= v11) {
        if false {
            v15 := GenS1{F0: (true || ((((((float64(1.5) * (((-2.25 / 1e10) * 1.5) / 1.0)) * 1.5) + 1e10) / 1e10) / 3.14159) > (3.14159 - 1e10)))}
            v11 = (v11 / ((v11) | 1))
            _ = v15
        } else {
            if true {
                v16 := GenS1{F0: false}
                var v17 float64 = ((0.1 - (0.0 - 3.14159)) + -2.25)
                _ = v16
                _ = v17
            } else {
                v18 := GenS1{F0: false}
                switch "\n\t" {
                case "\n\t":
                    if (!(0 >= 992)) {
                        var v19 float64 = 1e10
                        v20 := []int{-1, 885}
                        v21 := GenS1{F0: false}
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int16 = 1
                        v23 := []int{-165, 32767, 1}
                        _ = v22
                        _ = v23
                    }
                case "":
                    var v24 int = 752
                    _ = v24
                default:
                    var v25 int32 = -700
                    _ = v25
                }
                switch (("\n\t" + (("\n\t" + "hello") + "")) + "a") {
                case "":
                    v11 -= (v11 * -1)
                }
                _ = v18
            }
        }
    } else {
        var v26 float32 = (1.5 + 1e10)
        v27 := []int{-167}
        _ = v26
        _ = v27
    }
    _ = v11
    return (^(v11))
}

func main() {
    var v1 float64 = (-2.25 / 1e10)
    var v2 uint = 18446744073709551615
    var v3 uint32 = 65535
    v4 := make(map[uint32]string)
    var v5 float64 = float64(3.14159)
    var v6 string = "hello"
    if (127 > -730) {
        var v28 int = -229
        fmt.Println(v1, v28)
        _ = v28
    } else {
        v3 = (v3 % (((v3 ^ (v3 * 125))) | 1))
        v29 := GenS2{F0: 4294967295}
        _ = v29
    }
    var v30 int16 = -844
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 uint32 = 0
        fmt.Println(v3, v3)
        _ = fi31
        _ = v32
    }
    v3 = 65535
    var v33 uint64 = 91
    for fi34 := 0; fi34 < 20; fi34++ {
        v1 += 0.1
        var v35 float32 = 1e10
        var v36 bool = false
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 int16 = (v30 % ((v30) | 1))
    v33 += 1
    v38 := []int{}
    var v39 uint32 = v3
    fmt.Println(v1, v3)
    for fi40 := 0; fi40 < 19; fi40++ {
        v2 = 1
        v3 = v39
        fmt.Println(fi40, v37, fi40)
        _ = fi40
    }
    v6 = (v6 + "hello")
    v2 += uint(v33)
    v41 := GenS1{F0: false}
    for fi42 := 0; fi42 < 4; fi42++ {
        if false {
            delete(v4, 42)
            for fi43 := 0; fi43 < 16; fi43++ {
                switch -175 {
                case -912:
                    fmt.Println(v2, fi42, v2)
                }
                var v44 int16 = v30
                _ = fi43
                _ = v44
            }
        }
        _ = fi42
    }
    fmt.Println(v33, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
