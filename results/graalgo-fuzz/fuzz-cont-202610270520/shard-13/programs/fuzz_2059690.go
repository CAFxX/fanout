package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 string }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }

func main() {
    var v1 float32 = (1.5 + ((1.5 - -2.25) * (((3.14159 / 1.0) / 1.0) * float32(float64(1e10)))))
    var v2 float32 = float32((-2.25 + float64((v1 - (v1 / 1.0)))))
    var v3 float64 = (1.5 + (3.14159 + (0.1 / 1.0)))
    var v4 int32 = 480
    for fi5 := 0; fi5 < 9; fi5++ {
        switch ((("foo bar" + ((("\n\t" + "foo bar") + (("" + "\n\t") + "a")) + "foo bar")) + "a") + "") {
        case "\n\t":
            fmt.Println(fi5, v3)
        default:
            if ("\n\t" >= ("foo bar" + ("foo bar" + "a"))) {
                switch fi5 {
                case 32767:
                    if (1000 >= 1000) {
                        v6 := GenS1{F0: 179, F1: 255, F2: ("a" + ("hello" + "a"))}
                        _ = v6
                    }
                case -347:
                    var v7 uint32 = 2
                    _ = v7
                case -525:
                    v8 := []int{fi5, (fi5 ^ 147), 4294967295}
                    _ = v8
                }
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 float32 = 1.5
                    v2 = 1.5
                    _ = fi9
                    _ = v10
                }
                if false {
                    if false {
                        if false {
                            v11 := []int{fi5, fi5}
                            for fi12 := 0; fi12 < 5; fi12++ {
                                v11 = append(v11, fi12)
                                _ = fi12
                            }
                            _ = v11
                        }
                    }
                    var v13 uint8 = 0
                    _ = v13
                } else {
                    v2 += v2
                    var v14 uint64 = 18446744073709551615
                    var v15 int32 = -986
                    _ = v14
                    _ = v15
                }
            }
        }
        _ = fi5
    }
    v1 = (1.5 / (v1 / v1))
    var v16 uint64 = 4294967295
    fmt.Println(v16, v3, v4)
    fmt.Println(v3, v2)
    var v17 int = -9223372036854775808
    var v18 int32 = 2147483647
    for fi19 := 0; fi19 < 1; fi19++ {
        fmt.Println(v1, v17)
        var v20 int32 = (-(v4))
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v18 = v18
        _ = fi21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 uint64 = v16
        v24 := GenS2{F0: "", F1: "hello"}
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint32 = 4294967295
    v26 := GenS1{F0: 42, F1: 841, F2: "hello"}
    for fi27 := 0; fi27 < 19; fi27++ {
        v17 = fi27
        v28 := []int{(v17 / (((v17 / (((fi27 / ((v17) | 1))) | 1))) | 1)), v17, 32768}
        _ = fi27
        _ = v28
    }
    v3 = (0.1 * v3)
    v29 := GenS3{F0: 1000, F1: 136, F2: 1}
    v30 := []int{v17}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
