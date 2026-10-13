package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 int8 = 127
    var v3 float32 = float32(-2.25)
    if (640 <= 1) {
        if true {
            for fi4 := 0; fi4 < 2; fi4++ {
                v5 := []int{}
                _ = fi4
                _ = v5
            }
            for fi6 := 0; fi6 < 2; fi6++ {
                if true {
                    var v7 uint32 = 117
                    v7 -= (v7 - v7)
                    if (0 < 42) {
                        if (!((3.14159 + 1e10) < (1e10 / float64((v3 * v3))))) {
                            v1 = "a"
                            v8 := []int{(^(fi6)), 1}
                            _ = v8
                        } else {
                            switch fi6 {
                            case 512:
                                v3 += v3
                            case 42:
                                for fi9 := 0; fi9 < 12; fi9++ {
                                    var v10 uint8 = 91
                                    v11 := []int{2147483647}
                                    v12 := GenS1{F0: (v2 >> ((1) & 63))}
                                    v2 += 42
                                    _ = fi9
                                    _ = v10
                                    _ = v11
                                    _ = v12
                                }
                            }
                            for fi13 := 0; fi13 < 20; fi13++ {
                                v3 -= (1e10 / v3)
                                v14 := []int{fi13, fi13, fi13}
                                _ = fi13
                                _ = v14
                            }
                            switch ("\n\t" + "hello") {
                            case "":
                                v15 := GenS1{F0: 0}
                                _ = v15
                            }
                        }
                        v16 := GenS2{F0: 883}
                        if false {
                            var v17 uint64 = 986
                            var v18 int = fi6
                            if (((517 >= 546) && ((v3 - v3) != 0.1)) || true) {
                                v19 := []int{(v18 - (v18 << ((0) & 63))), (v18 ^ (v18 & v18)), fi6}
                                v2 -= v2
                                _ = v19
                            }
                            _ = v17
                            _ = v18
                        }
                        _ = v16
                    } else {
                        v3 -= v3
                        fmt.Println(v3, v2, v1)
                    }
                    _ = v7
                }
                _ = fi6
            }
        } else {
            fmt.Println(v2)
            v3 = 1.5
        }
        var v20 bool = false
        fmt.Println(v3, v3, v3)
        _ = v20
    }
    var v21 float64 = 1e10
    v22 := []int{-2147483648, 32767, 42, -143}
    var v23 int8 = v2
    v21 = v21
    var v24 uint16 = 2
    var v25 uint8 = 0
    if (!true) {
        if len(v22) > 0 {
            v22[0] = 1
        }
        var v26 int16 = 42
        v22 = append(v22, -359)
        _ = v26
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v27 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
}
