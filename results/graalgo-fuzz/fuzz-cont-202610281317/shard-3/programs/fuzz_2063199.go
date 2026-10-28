package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0() int64 {
    var v4 float32 = 1e10
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int16 = 1
        v7 := []int{}
        var v8 uint16 = 255
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        switch fi9 {
        case 2147483647:
            if false {
                v10 := GenS1{F0: 42}
                _ = v10
            } else {
                if true {
                    var v11 uint8 = 255
                    var v12 uint16 = 0
                    if (v4 >= (v4 + (float32(1e10) + (v4 / (float32(0.1) + ((v4 / 0.1) / 1.5)))))) {
                        if false {
                            v13 := GenS2{F0: v12}
                            fmt.Println(v11)
                            var v14 uint16 = (v12 % ((v12) | 1))
                            _ = v13
                            _ = v14
                        } else {
                            v15 := []int{(fi9 + fi9)}
                            _ = v15
                        }
                        v16 := []int{fi9, (fi9 >> ((uint(v12)) & 63))}
                        _ = v16
                    } else {
                        var v17 string = ("foo bar" + "\n\t")
                        v18 := GenS1{F0: 1}
                        _ = v17
                        _ = v18
                    }
                    _ = v11
                    _ = v12
                } else {
                    v4 += 1.5
                }
                if (false && (!((68 == 100) || (2 <= 57)))) {
                    switch "hello" {
                    case "":
                        v4 -= ((v4 / 0.1) / (float32((0.1 + 1.5)) * -2.25))
                    case "a":
                        v19 := []int{fi9}
                        _ = v19
                    }
                    v20 := []int{fi9, (fi9 * fi9), (fi9 & -434), (-(fi9))}
                    _ = v20
                } else {
                    v21 := []int{}
                    v4 += v4
                    _ = v21
                }
            }
        case 1:
            v4 += (((v4 * 0.1) + (v4 + v4)) / (1e10 + v4))
        case 42:
            v22 := GenS2{F0: 488}
            _ = v22
        default:
            switch (fi9 >> ((0) & 63)) {
            case 128:
                var v23 int16 = 32767
                _ = v23
            case 949:
                v4 = (0.0 / 1.0)
            case 127:
                fmt.Println(v4, fi9)
            }
        }
        v24 := GenS2{F0: 339}
        _ = fi9
        _ = v24
    }
    v25 := GenS1{F0: 42}
    v4 -= -2.25
    _ = v25
    _ = v4
    return 0
}
func genFn1(p0 uint8) uint {
    switch "" {
    case "a":
        p0 += p0
    }
    v26 := []int{9223372036854775807, -9223372036854775808}
    _ = p0
    _ = v26
    return 65535
}

func main() {
    var v1 bool = true
    var v2 uint64 = 441
    var v3 float32 = (-2.25 - (0.1 + 3.14159))
    if v1 {
        v2 -= (v2 % ((533) | 1))
    }
    v3 -= float32(float64(v3))
    var v27 uint32 = 623
    fmt.Println(v1)
    v3 = -2.25
    var v28 int8 = -1
    var v29 int32 = 430
    var v30 int64 = 1
    v28 = (^(v28))
    fmt.Println(v1, v29, v28)
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 uint = genFn1(214)
        fmt.Println(v1, v28)
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
