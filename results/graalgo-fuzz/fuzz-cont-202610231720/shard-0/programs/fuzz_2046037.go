package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 string; F2 int16 }
func genFn0(p0 int16) uint8 {
    p0 -= -32768
    p0 += -32768
    p0 += (p0 / (((-(p0))) | 1))
    _ = p0
    return 42
}
func genFn1(p0 string) int8 {
    var v7 uint8 = 152
    v8 := []int{-1, 2147483647, 32767, 4294967295}
    v9 := GenS1{F0: -191}
    for fi10 := 0; fi10 < 16; fi10++ {
        if ((v7 / ((128) | 1)) > v7) {
            var v11 int8 = 127
            switch "foo bar" {
            case "\n\t":
                switch ("\n\t" + p0) {
                case "foo bar":
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 float32 = ((-2.25 * float32(1.5)) * -2.25)
                        if true {
                            v7 = 0
                            fmt.Println(p0, v13)
                            p0 = ""
                        }
                        _ = fi12
                        _ = v13
                    }
                case "hello":
                    v14 := GenS2{F0: 155, F1: 756, F2: 65535}
                    _ = v14
                }
            }
            v7 += (v7 & v7)
            _ = v11
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 bool = ((!true) && true)
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 float32 = ((float32((-2.25 * (0.0 - 3.14159))) * 0.0) / 3.14159)
            p0 = ("hello" + p0)
            _ = fi17
            _ = v18
        }
        _ = fi10
    }
    p0 = p0
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return int8(v7)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    v3 := []uint32{176, 4294967295, 2147483648}
    var v4 int32 = 0
    var v5 float64 = 1e10
    v6 := make(map[int]int)
    var v19 float64 = (((v5 / (float64(float32(3.14159)) - 1.5)) + v5) * v5)
    for fi20 := 0; fi20 < 20; fi20++ {
        v5 = (((1.5 + ((((v5 / v19) * v5) - v5) + (float64(0.1) / v5))) / 1.0) + 1.5)
        _ = fi20
    }
    v21 := GenS3{F0: 731, F1: "", F2: 0}
    var v22 float64 = v19
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int16 = 283
            v25 -= -843
            v26 := v2[733]
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    var v27 int16 = -465
    for fi28 := 0; fi28 < 6; fi28++ {
        v19 = (v5 * 3.14159)
        _ = fi28
    }
    var v29 uint16 = 515
    for fi30 := 0; fi30 < 15; fi30++ {
        if false {
            for fi31 := 0; fi31 < 5; fi31++ {
                v19 = (0.0 - 1.5)
                if (1 != 127) {
                    fmt.Println(v22, v5, v19)
                }
                _ = fi31
            }
            fmt.Println(fi30, v27, v27)
        } else {
            v22 -= float64(3.14159)
            v5 += -2.25
        }
        _ = fi30
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        v27 = (v27 * v27)
        v5 = float64((0.0 * (1.5 - 1e10)))
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
