package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 float64 = 1.5
    var v2 uint = 18446744073709551615
    var v3 float32 = 0.0
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    var v5 string = ("\n\t" + ((((("" + ("hello" + ("\n\t" + ("\n\t" + ("foo bar" + "\n\t"))))) + "hello") + ("\n\t" + (((("" + "") + (("a" + "hello") + "a")) + "hello") + "foo bar"))) + "a") + "\n\t"))
    v2 = (v2 / ((0) | 1))
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint32 = 2147483648
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 string = (v5 + (("\n\t" + v5) + v5))
        _ = fi8
        _ = v9
    }
    if true {
        switch 23 {
        case 32767:
            v10 := []int{-917}
            _ = v10
        case 1:
            var v11 uint32 = 2147483648
            _ = v11
        case -20:
            if false {
                v3 = 0.0
            }
        }
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := []int{(^(fi12)), fi12}
            v2 = 42
            v1 -= v1
            _ = fi12
            _ = v13
        }
        v14 := []int{-2147483648, 32768}
        _ = v14
    }
    var v15 int16 = 1000
    if true {
        var v16 float32 = 1e10
        v17 := []int{128, 32767, -90}
        _ = v16
        _ = v17
    } else {
        if false {
            v15 -= -983
        } else {
            for fi18 := 0; fi18 < 14; fi18++ {
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        v19 := []int{}
                        _ = v19
                    case "hello":
                        var v20 uint32 = 1
                        _ = v20
                    }
                    v21 := GenS1{F0: 42, F1: 126}
                    _ = v21
                } else {
                    fmt.Println(v15)
                }
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := []int{fi22}
                    v2 -= v2
                    var v24 uint16 = 385
                    fmt.Println(v2)
                    v25 := GenS2{F0: 255, F1: 900, F2: -599}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi18
            }
        }
    }
    v26 := GenS3{F0: 128, F1: 65535, F2: 1}
    if true {
        var v27 int = 4294967295
        _ = v27
    } else {
        var v28 uint16 = 42
        fmt.Println(v28, v3)
        var v29 uint8 = 127
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v3 = (((3.14159 + float32(1.5)) / ((v3 * 0.1) - float32(float64(1e10)))) + 1e10)
        _ = fi30
    }
    if false {
        if true {
            var v31 uint16 = 8
            v31 -= 2
            v32 := GenS3{F0: 199, F1: (v31 + 65535), F2: 0}
            _ = v31
            _ = v32
        } else {
            if false {
                if (false && false) {
                    v33 := []int{499, 32767}
                    var v34 int64 = 647
                    if true {
                        v35 := GenS3{F0: 0, F1: 42, F2: 2147483647}
                        if len(v33) > 0 {
                            v33[0] = 0
                        }
                        v2 = 683
                        _ = v35
                    }
                    _ = v33
                    _ = v34
                }
                v3 = ((((v3 + (v3 * float32(1.5))) * (float32(v1) * v3)) - v3) / v3)
                switch v5 {
                case "hello":
                    fmt.Println(v2, v15, v5)
                }
            } else {
                v36 := []int{32768, 42, 0}
                fmt.Println(v15, v15)
                _ = v36
            }
            v37 := GenS4{F0: 0, F1: 0}
            var v38 string = v5
            _ = v37
            _ = v38
        }
        var v39 int8 = int8(v2)
        fmt.Println(v5, v1)
        _ = v39
    }
    var v40 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v40
    _ = v5
}
